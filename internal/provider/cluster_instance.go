package provider

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	temboclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &temboInstanceResource{}
	_ resource.ResourceWithConfigure = &temboInstanceResource{}
)

const (
	DefaultReplicas = 1
)

// temboInstanceResourceModel maps the resource schema data.
type temboInstanceResourceModel struct {
	InstanceID      types.String     `tfsdk:"instance_id"`
	InstanceName    types.String     `tfsdk:"instance_name"`
	OrgId           types.String     `tfsdk:"org_id"`
	CPU             types.String     `tfsdk:"cpu"`
	StackType       types.String     `tfsdk:"stack_type"`
	Environment     types.String     `tfsdk:"environment"`
	Replicas        types.Int64      `tfsdk:"replicas"`
	Memory          types.String     `tfsdk:"memory"`
	Storage         types.String     `tfsdk:"storage"`
	LastUpdated     types.String     `tfsdk:"last_updated"`
	State           types.String     `tfsdk:"state"`
	ExtraDomainsRw  []types.String   `tfsdk:"extra_domains_rw"`
	PostgresConfigs []PostGresConfig `tfsdk:"postgres_configs"`
	TrunkInstalls   []TrunkInstall   `tfsdk:"trunk_installs"`
}

type PostGresConfig struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type TrunkInstall struct {
	Name    types.String `tfsdk:"name"`
	Version types.String `tfsdk:"version"`
}

// Configure adds the provider configured data to the resource.
func (r *temboInstanceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	temboInstanceConfig, ok := req.ProviderData.(instanceConfig)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *instanceConfig, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.temboInstanceConfig = temboInstanceConfig
}

// NewTemboInstanceResource is a helper function to simplify the provider implementation.
func NewTemboInstanceResource() resource.Resource {
	return &temboInstanceResource{}
}

// temboInstanceResource is the resource implementation.
type temboInstanceResource struct {
	temboInstanceConfig instanceConfig
}

// Tembo CLuster Configuration.
type instanceConfig struct {
	client      *temboclient.APIClient
	accessToken string
}

// Metadata returns the resource type name.
func (r *temboInstanceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_instance"
}

// Schema defines the schema for the resource.
func (r *temboInstanceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"instance_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"instance_name": schema.StringAttribute{
				Required: true,
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"cpu": schema.StringAttribute{
				Required: true,
			},
			"stack_type": schema.StringAttribute{
				Required: true,
			},
			"replicas": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(DefaultReplicas),
			},
			"environment": schema.StringAttribute{
				Required: true,
			},
			"memory": schema.StringAttribute{
				Required: true,
			},
			"storage": schema.StringAttribute{
				Required: true,
			},
			"state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"extra_domains_rw": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"postgres_configs": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
						},
						"value": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"trunk_installs": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
						},
						"version": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *temboInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan temboInstanceResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call API to Create Tembo Instance
	createInstance := *temboclient.NewCreateInstance(
		temboclient.Cpu(plan.CPU.ValueString()),
		temboclient.Environment(plan.Environment.ValueString()),
		plan.InstanceName.ValueString(),
		temboclient.Memory(plan.Memory.ValueString()),
		temboclient.EntityType(plan.StackType.ValueString()),
		temboclient.Storage(plan.Storage.ValueString()))

	if !plan.Replicas.IsNull() {
		createInstance.SetReplicas(int32(plan.Replicas.ValueInt64()))
	}

	createInstance.SetExtraDomainsRw(getExtraDomainRW(plan.ExtraDomainsRw))

	createInstance.SetPostgresConfigs(getPgConfig(plan.PostgresConfigs))

	createInstance.SetTrunkInstalls(getTrunkInstall(plan.TrunkInstalls))

	// TODO: Figure out a better way to set this so it doesn't have to be be called in each method.
	ctx = context.WithValue(ctx, temboclient.ContextAccessToken, r.temboInstanceConfig.accessToken)

	instanceRequest := r.temboInstanceConfig.client.InstanceApi.CreateInstance(ctx,
		plan.OrgId.ValueString())

	instance, response, err := instanceRequest.CreateInstance(createInstance).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Tembo Instance:",
			"Could not create Tembo Instance, unexpected error: "+getErrorFromResponse(response),
		)
		return
	}

	// Wait until it's created.
	for {
		instanceState := getInstanceState(r, ctx, plan.OrgId.ValueString(), instance.GetInstanceId(), &resp.Diagnostics)

		if instanceState == temboclient.ERROR || instanceState == temboclient.UP {
			break
		}

		time.Sleep(10 * time.Second)
		log.Printf("[INFO] Waiting for Tembo instance %s to be UP", plan.InstanceName)
	}

	log.Printf("[INFO] Tembo instance %s has been successfully created", plan.InstanceName)

	// Map response body to schema and populate Computed attribute values
	instance.SetState(temboclient.UP)
	setTemboInstanceResourceModel(&plan, instance, true, &resp.Diagnostics)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *temboInstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state *temboInstanceResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, temboclient.ContextAccessToken, r.temboInstanceConfig.accessToken)
	// Get refreshed Instance value from API
	instance, response, err := r.temboInstanceConfig.client.InstanceApi.GetInstance(ctx, state.OrgId.ValueString(), state.InstanceID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Tembo Instance",
			"Could not read Tembo Instance ID "+state.InstanceID.ValueString()+": "+getErrorFromResponse(response),
		)
		return
	}

	// Overwrite items with refreshed state
	setTemboInstanceResourceModel(state, instance, false, &resp.Diagnostics)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *temboInstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan temboInstanceResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	updateInstance := *temboclient.NewUpdateInstance(
		temboclient.Cpu(plan.CPU.ValueString()),
		temboclient.Environment(plan.Environment.ValueString()),
		temboclient.Memory((plan.Memory.ValueString())),
		int32(plan.Replicas.ValueInt64()),
		temboclient.Storage(plan.Storage.ValueString()))

	updateInstance.SetExtraDomainsRw(getExtraDomainRW(plan.ExtraDomainsRw))
	updateInstance.SetPostgresConfigs(getPgConfig(plan.PostgresConfigs))
	updateInstance.SetTrunkInstalls(getTrunkInstall(plan.TrunkInstalls))

	log.Printf("[INFO] Tembo instanceID %s", plan)

	ctx = context.WithValue(ctx, temboclient.ContextAccessToken, r.temboInstanceConfig.accessToken)

	// Update existing Instance
	_, response, err := r.temboInstanceConfig.client.InstanceApi.PutInstance(
		ctx,
		plan.OrgId.ValueString(),
		plan.InstanceID.ValueString()).UpdateInstance(updateInstance).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Tembo Instance",
			"Could not update Tembo Instance ID "+plan.InstanceID.ValueString()+": "+getErrorFromResponse(response),
		)
		return
	}

	// Wait until it's updated
	for {
		instance, err := getInstance(r, ctx, plan.OrgId.ValueString(), plan.InstanceID.ValueString(), &resp.Diagnostics)

		if err != nil {
			return
		}

		if updateInstance.GetCpu() == instance.GetCpu() &&
			updateInstance.GetMemory() == instance.GetMemory() &&
			updateInstance.GetStorage() == instance.GetStorage() &&
			updateInstance.GetReplicas() == instance.GetReplicas() &&
			instance.GetState() == temboclient.UP {
			// Update resource state with updated items and timestamp
			setTemboInstanceResourceModel(&plan, &instance, true, &resp.Diagnostics)
			break
		}
		time.Sleep(10 * time.Second)

		log.Printf("[INFO] Waiting for Tembo instance %s to be updated", plan.InstanceName)
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *temboInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state temboInstanceResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, temboclient.ContextAccessToken, r.temboInstanceConfig.accessToken)

	// Delete existing Tembo Instance
	_, response, err := (*r.temboInstanceConfig.client.InstanceApi).DeleteInstance(ctx, state.OrgId.ValueString(), state.InstanceID.ValueString()).Execute()
	if err != nil {

		resp.Diagnostics.AddError(
			"Error Deleting Tembo Instance",
			"Could not delete instance, unexpected error: "+getErrorFromResponse(response),
		)
		return
	}

	// Wait until it's deleted
	for {
		instanceState := getInstanceState(r, ctx, state.OrgId.ValueString(), state.InstanceID.ValueString(), &resp.Diagnostics)

		if instanceState == temboclient.ERROR || instanceState == temboclient.DELETED {
			break
		}

		time.Sleep(10 * time.Second)
		log.Printf("[INFO] Waiting for Tembo instance %s to be DELETED", state.InstanceName)
	}

	log.Printf("[INFO] Tembo instance %s has been successfully deleted", state.InstanceName)
}

func setTemboInstanceResourceModel(instanceResourceModel *temboInstanceResourceModel,
	instance *temboclient.Instance, updateComputedValue bool, diagnostics *diag.Diagnostics) {
	if updateComputedValue {
		instanceResourceModel.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
	}

	instanceResourceModel.InstanceID = types.StringValue(instance.GetInstanceId())
	instanceResourceModel.InstanceName = types.StringValue(instance.InstanceName)
	instanceResourceModel.OrgId = types.StringValue(instance.GetOrganizationId())
	instanceResourceModel.CPU = types.StringValue(string(instance.GetCpu()))
	instanceResourceModel.StackType = types.StringValue(string(instance.StackType))
	instanceResourceModel.Environment = types.StringValue(string(instance.GetEnvironment()))
	instanceResourceModel.Memory = types.StringValue(string(instance.GetMemory()))
	instanceResourceModel.Storage = types.StringValue(string(instance.GetStorage()))
	instanceResourceModel.State = types.StringValue(string(instance.GetState()))
	instanceResourceModel.Replicas = types.Int64Value(int64(instance.GetReplicas()))

	if len(instance.ExtraDomainsRw) > 0 {
		var localExtraDomainsRw []basetypes.StringValue
		for _, domain := range instance.ExtraDomainsRw {
			localExtraDomainsRw = append(localExtraDomainsRw, types.StringValue(domain))
		}
		instanceResourceModel.ExtraDomainsRw = localExtraDomainsRw
	}

	if len(instance.PostgresConfigs) > 0 {
		var localPGConfigs []PostGresConfig
		for _, pgConfig := range instance.PostgresConfigs {
			localPGConfigs = append(localPGConfigs, PostGresConfig{Name: types.StringValue(pgConfig.Name), Value: types.StringValue(pgConfig.Value)})
		}
		instanceResourceModel.PostgresConfigs = localPGConfigs
	}

	if len(instance.TrunkInstalls) > 0 {
		for _, trunkInstall := range instance.TrunkInstalls {
			if trunkInstall.Error {
				diagnostics.AddError(
					"Error with Trunk Install:",
					"unexpected error: "+trunkInstall.GetErrorMessage(),
				)
			}
		}
	}
}

func getExtraDomainRW(extraDomainRWs []basetypes.StringValue) []string {
	var localExtraDomainRW []string
	if len(extraDomainRWs) > 0 {
		for _, extraDomainRW := range extraDomainRWs {
			localExtraDomainRW = append(localExtraDomainRW, extraDomainRW.ValueString())
		}
	}
	return localExtraDomainRW
}

func getPgConfig(postgresConfigs []PostGresConfig) []temboclient.PgConfig {
	var localPGConfigs []temboclient.PgConfig
	if len(postgresConfigs) > 0 {
		for _, pgConfig := range postgresConfigs {
			localPGConfigs = append(localPGConfigs, temboclient.PgConfig{Name: pgConfig.Name.ValueString(), Value: pgConfig.Value.ValueString()})
		}
	}
	return localPGConfigs
}

func getTrunkInstall(trunkInstalls []TrunkInstall) []temboclient.TrunkInstall {
	var localTrunkInstalls []temboclient.TrunkInstall
	if len(trunkInstalls) > 0 {
		for _, trunkInstall := range trunkInstalls {
			localTrunkInstalls = append(localTrunkInstalls, fetchTrunkInstall(trunkInstall))
		}
	}
	return localTrunkInstalls
}

func fetchTrunkInstall(trunkInstall TrunkInstall) temboclient.TrunkInstall {
	localTrunkInstall := temboclient.TrunkInstall{
		Name: trunkInstall.Name.ValueString(),
	}
	localTrunkInstall.SetVersion(trunkInstall.Version.ValueString())

	return localTrunkInstall
}

func getErrorFromResponse(response *http.Response) string {
	localVarBody, _ := io.ReadAll(response.Body)
	response.Body.Close()

	return string(localVarBody)
}

func getInstanceState(r *temboInstanceResource, ctx context.Context,
	orgId string, instanceId string, diagnostics *diag.Diagnostics) temboclient.State {
	refreshInstance, response, err := r.temboInstanceConfig.client.InstanceApi.GetInstance(ctx, orgId, instanceId).Execute()
	if err != nil {
		diagnostics.AddError(
			"Error Reading Tembo Instance State",
			"Could not read Tembo Instance ID "+instanceId+": "+getErrorFromResponse(response),
		)
		return temboclient.ERROR
	}

	return refreshInstance.GetState()
}

func getInstance(r *temboInstanceResource, ctx context.Context,
	orgId string, instanceId string, diagnostics *diag.Diagnostics) (temboclient.Instance, error) {
	refreshInstance, response, err := r.temboInstanceConfig.client.InstanceApi.GetInstance(ctx, orgId, instanceId).Execute()
	if err != nil {
		diagnostics.AddError(
			"Error Reading Tembo Instance State",
			"Could not read Tembo Instance ID "+instanceId+": "+getErrorFromResponse(response),
		)
	}

	return *refreshInstance, err
}
