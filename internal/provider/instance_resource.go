package provider

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	temboclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &temboInstanceResource{}
	_ resource.ResourceWithConfigure   = &temboInstanceResource{}
	_ resource.ResourceWithImportState = &temboInstanceResource{}
)

const (
	DefaultReplicas = 1
)

// temboInstanceResourceModel maps the resource schema data.
type temboInstanceResourceModel struct {
	InstanceID              types.String      `tfsdk:"instance_id"`
	InstanceName            types.String      `tfsdk:"instance_name"`
	OrgId                   types.String      `tfsdk:"org_id"`
	CPU                     types.String      `tfsdk:"cpu"`
	StackType               types.String      `tfsdk:"stack_type"`
	Environment             types.String      `tfsdk:"environment"`
	Replicas                types.Int64       `tfsdk:"replicas"`
	Memory                  types.String      `tfsdk:"memory"`
	Storage                 types.String      `tfsdk:"storage"`
	LastUpdated             types.String      `tfsdk:"last_updated"`
	State                   types.String      `tfsdk:"state"`
	ExtraDomainsRw          []types.String    `tfsdk:"extra_domains_rw"`
	PostgresConfigs         []KeyValue        `tfsdk:"postgres_configs"`
	TrunkInstalls           []TrunkInstall    `tfsdk:"trunk_installs"`
	Extensions              []Extension       `tfsdk:"extensions"`
	IpAllowList             []types.String    `tfsdk:"ip_allow_list"`
	ConnectionPooler        *ConnectionPooler `tfsdk:"connection_pooler"`
	Restore                 *Restore          `tfsdk:"restore"`
	FirstRecoverabilityTime types.String      `tfsdk:"first_recoverability_time"`
	ProviderId              types.String      `tfsdk:"provider_id"`
	RegionId                types.String      `tfsdk:"region_id"`
	Autoscaling             *AutoscalingModel `tfsdk:"autoscaling"`
	PGVersion               types.String      `tfsdk:"pg_version"`
}

type Restore struct {
	InstanceId         types.String `tfsdk:"instance_id"`
	RecoveryTargetTime types.String `tfsdk:"recovery_target_time"`
}

type ConnectionPooler struct {
	Enabled types.Bool `tfsdk:"enabled"`
	Pooler  PgBouncer  `tfsdk:"pooler"`
}
type PgBouncer struct {
	Parameters map[string]string `tfsdk:"parameters"`
	PoolMode   types.String      `tfsdk:"pool_mode"`
}

type KeyValue struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type TrunkInstall struct {
	Name    types.String `tfsdk:"name"`
	Version types.String `tfsdk:"version"`
}

type Extension struct {
	Description types.String               `tfsdk:"description"`
	Locations   []ExtensionInstallLocation `tfsdk:"locations"`
	Name        string                     `tfsdk:"name"`
}

type ExtensionInstallLocation struct {
	Database types.String `tfsdk:"database"`
	Enabled  types.Bool   `tfsdk:"enabled"`
	Schema   types.String `tfsdk:"schema"`
	Version  types.String `tfsdk:"version"`
}

type AutoscalingModel struct {
	Autostop *AutoscalingAutostopModel `tfsdk:"autostop"`
	Storage  *AutoscalingStorageModel  `tfsdk:"storage"`
}

type AutoscalingAutostopModel struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type AutoscalingStorageModel struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	Maximum types.String `tfsdk:"maximum"`
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
		Description: "Resource for creating a tembo instance.",
		Attributes: map[string]schema.Attribute{
			"instance_id": schema.StringAttribute{
				MarkdownDescription: "Unique ID for the instance generated by Tembo",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				MarkdownDescription: "Last updated date time in UTC",
				Computed:            true,
			},
			"instance_name": schema.StringAttribute{
				MarkdownDescription: "Unique name of the instance",
				Required:            true,
			},
			"org_id": schema.StringAttribute{
				MarkdownDescription: "Id of the organization in which the instance will be created",
				Required:            true,
			},
			"cpu": schema.StringAttribute{
				MarkdownDescription: "CPU",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf(getStringArrayForType(temboclient.AllowedCpuEnumValues)...)},
			},
			"stack_type": schema.StringAttribute{
				MarkdownDescription: "Stack type for the instance.",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf(getStringArrayForType(temboclient.AllowedStackTypeEnumValues)...)},
			},
			"replicas": schema.Int64Attribute{
				MarkdownDescription: "Instance replicas",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(DefaultReplicas),
			},
			"environment": schema.StringAttribute{
				MarkdownDescription: "Environment",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf(getStringArrayForType(temboclient.AllowedEnvironmentEnumValues)...)},
			},
			"memory": schema.StringAttribute{
				MarkdownDescription: "Memory",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf(getStringArrayForType(temboclient.AllowedMemoryEnumValues)...)},
			},
			"storage": schema.StringAttribute{
				MarkdownDescription: "Storage",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf(getStringArrayForType(temboclient.AllowedStorageEnumValues)...)},
			},
			"state": schema.StringAttribute{
				MarkdownDescription: "Instance state. Values: Submitted, Up, Configuring, Error, Restarting, Starting, Stopping, Stopped, Deleting, Deleted",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"extra_domains_rw": schema.ListAttribute{
				MarkdownDescription: "Custom domain. Read more [here](https://tembo.io/docs/tembo-cloud/custom-domains/)",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"postgres_configs": schema.ListNestedAttribute{
				MarkdownDescription: "Postgres configs",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Postgres config name",
							Required:            true,
						},
						"value": schema.StringAttribute{
							MarkdownDescription: "Postgres config value",
							Optional:            true,
						},
					},
				},
			},
			"trunk_installs": schema.ListNestedAttribute{
				MarkdownDescription: "Trunk installs",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Trunk install name",
							Required:            true,
						},
						"version": schema.StringAttribute{
							MarkdownDescription: "Trunk install version",
							Optional:            true,
						},
					},
				},
			},
			"extensions": schema.ListNestedAttribute{
				MarkdownDescription: "Extensions to install in the instance",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Extension Name",
							Required:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Extension Description",
							Optional:            true,
						},
						"locations": schema.ListNestedAttribute{
							MarkdownDescription: "Locations for extension",
							Required:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"database": schema.StringAttribute{
										MarkdownDescription: "Database to install extension",
										Required:            true,
									},
									"enabled": schema.BoolAttribute{
										MarkdownDescription: "Enabled",
										Required:            true,
									},
									"schema": schema.StringAttribute{
										MarkdownDescription: "The name of the schema",
										Optional:            true,
									},
									"version": schema.StringAttribute{
										MarkdownDescription: "The version of the extension to install",
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
			"ip_allow_list": schema.ListAttribute{
				MarkdownDescription: "Allowed IP list",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"connection_pooler": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Required: true,
					},
					"pooler": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"pool_mode": schema.StringAttribute{
								Required: true,
							},
							"parameters": schema.MapAttribute{
								MarkdownDescription: "Parameters",
								Required:            true,
								ElementType:         types.StringType,
							},
						},
					},
				},
				Optional: true,
			},
			"restore": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"instance_id": schema.StringAttribute{
						Required: true,
					},
					"recovery_target_time": schema.StringAttribute{
						Optional: true,
					},
				},
				Optional: true,
			},
			"first_recoverability_time": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The time at which the instance first became recoverable",
			},
			"provider_id": schema.StringAttribute{
				MarkdownDescription: "The Cloud Provider ID where the instance will be created",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("aws"),
			},
			"region_id": schema.StringAttribute{
				MarkdownDescription: "The cloud provider region where the instance will be created",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("us-east-1"),
			},
			"autoscaling": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"autostop": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Optional: true,
								Computed: true,
								Default:  booldefault.StaticBool(false),
							},
						},
					},
					"storage": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Optional: true,
								Computed: true,
								Default:  booldefault.StaticBool(false),
							},
							"maximum": schema.StringAttribute{
								Optional: true,
							},
						},
					},
				},
			},
			"pg_version": schema.StringAttribute{
				MarkdownDescription: "PostgresSQL version to deploy to Tembo Cloud",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("17"),
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

	var instanceId string
	var hasError bool

	isRestoreMode := (plan.Restore != nil)

	if isRestoreMode {
		ctx, instanceId, hasError = restoreInstance(plan, ctx, r, resp)
	} else {
		ctx, instanceId, hasError = createInstance(plan, ctx, r, resp)
	}

	if hasError {
		return
	}

	// Wait until it's created.
	for {
		latestInstance, err := getInstance(r, ctx, plan.OrgId.ValueString(), instanceId, &resp.Diagnostics)

		if err != nil {
			return
		}

		if latestInstance.GetState() == temboclient.ERROR ||
			latestInstance.GetState() == temboclient.UP {

			// Map response body to schema and populate Computed attribute values
			setTemboInstanceResourceModel(&plan, &latestInstance, true, isRestoreMode, &resp.Diagnostics)

			break
		}

		time.Sleep(10 * time.Second)
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func createInstance(plan temboInstanceResourceModel,
	ctx context.Context, r *temboInstanceResource,
	resp *resource.CreateResponse) (context.Context, string, bool) {
	createInstance := *temboclient.NewCreateInstance(
		temboclient.Cpu(plan.CPU.ValueString()),
		temboclient.Environment(plan.Environment.ValueString()),
		plan.InstanceName.ValueString(),
		temboclient.Memory(plan.Memory.ValueString()),
		temboclient.StackType(plan.StackType.ValueString()),
		temboclient.Storage(plan.Storage.ValueString()))

	if !plan.Replicas.IsNull() {
		createInstance.SetReplicas(int32(plan.Replicas.ValueInt64()))
	}

	createInstance.SetExtraDomainsRw(getStringArray(plan.ExtraDomainsRw))

	createInstance.SetPostgresConfigs(getPgConfig(plan.PostgresConfigs))

	createInstance.SetTrunkInstalls(getTemboTrunkInstalls(plan.TrunkInstalls))

	createInstance.SetExtensions(getTemboExtensions(plan.Extensions))

	createInstance.SetIpAllowList(getStringArray(plan.IpAllowList))

	createInstance.SetProviderId(plan.ProviderId.ValueString())
	createInstance.SetRegionId(plan.RegionId.ValueString())

	if plan.ConnectionPooler != nil {
		createInstance.SetConnectionPooler(*getConnectionPooler(plan.ConnectionPooler))
	}

	if plan.Autoscaling != nil {
		autoscaling := temboclient.NewPatchAutoscaling()

		// Handle Autostop
		if plan.Autoscaling.Autostop != nil {
			autostop := temboclient.NewAutoStop(plan.Autoscaling.Autostop.Enabled.ValueBool())
			autoscaling.SetAutostop(*autostop)
		}

		// Handle Storage
		if plan.Autoscaling.Storage != nil {
			storage := temboclient.NewAutoscalingStorage(plan.Autoscaling.Storage.Enabled.ValueBool())
			if !plan.Autoscaling.Storage.Maximum.IsNull() {
				storage.SetMaximum(temboclient.Storage(plan.Autoscaling.Storage.Maximum.ValueString()))
			}
			autoscaling.SetStorage(*storage)
		}

		createInstance.SetAutoscaling(*autoscaling)
	}

	if !plan.PGVersion.IsNull() {
		pgVersion, err := strconv.ParseInt(plan.PGVersion.ValueString(), 10, 32)
		if err != nil {
			resp.Diagnostics.AddError("Error parsing the PostgreSQL version",
				"Could not parse pg_version, unexpected error: "+err.Error(),
			)
			return nil, "", true
		}
		v := int32(pgVersion)
		createInstance.SetPgVersion(v)
	}

	ctx = context.WithValue(ctx, temboclient.ContextAccessToken, r.temboInstanceConfig.accessToken)

	instanceRequest := r.temboInstanceConfig.client.InstanceAPI.CreateInstance(ctx,
		plan.OrgId.ValueString())

	instance, response, err := instanceRequest.CreateInstance(createInstance).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Tembo Instance",
			"Could not create instance, unexpected error: "+err.Error()+" details:"+getErrorFromResponse(response),
		)
		return nil, "", true
	}

	return ctx, instance.GetInstanceId(), false
}

func restoreInstance(plan temboInstanceResourceModel,
	ctx context.Context, r *temboInstanceResource,
	resp *resource.CreateResponse) (context.Context, string, bool) {
	restoreInstance := *temboclient.NewRestoreInstance(
		plan.InstanceName.ValueString(),
		*temboclient.NewRestore(plan.Restore.InstanceId.ValueString()),
	)

	restoreInstance.SetCpu(temboclient.Cpu(plan.CPU.ValueString()))
	restoreInstance.SetEnvironment(temboclient.Environment(plan.Environment.ValueString()))
	restoreInstance.SetMemory(temboclient.Memory(plan.Memory.ValueString()))
	restoreInstance.SetStorage(temboclient.Storage(plan.Storage.ValueString()))

	restoreInstance.SetExtraDomainsRw(getStringArray(plan.ExtraDomainsRw))

	if plan.ConnectionPooler != nil {
		restoreInstance.SetConnectionPooler(*getConnectionPooler(plan.ConnectionPooler))
	}

	ctx = context.WithValue(ctx, temboclient.ContextAccessToken, r.temboInstanceConfig.accessToken)

	instanceRequest := r.temboInstanceConfig.client.InstanceAPI.RestoreInstance(ctx,
		plan.OrgId.ValueString())

	instance, response, err := instanceRequest.RestoreInstance(restoreInstance).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error restoring Tembo Instance",
			"Could not restore instance, unexpected error: "+err.Error()+" details:"+getErrorFromResponse(response),
		)
		return nil, "", true
	}
	return ctx, instance.GetInstanceId(), false
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
	instance, response, err := r.temboInstanceConfig.client.InstanceAPI.GetInstance(ctx, state.OrgId.ValueString(), state.InstanceID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Tembo Instance",
			"Could not read instance, unexpected error: "+err.Error()+" details:"+getErrorFromResponse(response),
		)

		return
	}

	// Overwrite items with refreshed state
	setTemboInstanceResourceModel(state, instance, false, false, &resp.Diagnostics)

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
	updateInstance := *temboclient.NewPatchInstance()

	cpu, err := temboclient.NewCpuFromValue(plan.CPU.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Tembo Instance",
			err.Error(),
		)
		return
	}
	updateInstance.SetCpu(*cpu)

	env, err := temboclient.NewEnvironmentFromValue(plan.Environment.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Tembo Instance",
			err.Error(),
		)
		return
	}
	updateInstance.SetEnvironment(*env)

	memory, err := temboclient.NewMemoryFromValue(plan.Memory.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Tembo Instance",
			err.Error(),
		)
		return
	}
	updateInstance.SetMemory(*memory)

	updateInstance.SetReplicas(int32(plan.Replicas.ValueInt64()))

	storage, err := temboclient.NewStorageFromValue(plan.Storage.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Tembo Instance",
			err.Error(),
		)
		return
	}
	updateInstance.SetStorage(*storage)

	updateInstance.SetExtraDomainsRw(getStringArray(plan.ExtraDomainsRw))
	updateInstance.SetPostgresConfigs(getPgConfig(plan.PostgresConfigs))
	updateInstance.SetTrunkInstalls(getTemboTrunkInstalls(plan.TrunkInstalls))
	updateInstance.SetExtensions(getTemboExtensions(plan.Extensions))
	updateInstance.SetIpAllowList(getStringArray(plan.IpAllowList))

	if plan.ConnectionPooler != nil {
		updateInstance.SetConnectionPooler(*getConnectionPooler(plan.ConnectionPooler))
	}

	ctx = context.WithValue(ctx, temboclient.ContextAccessToken, r.temboInstanceConfig.accessToken)

	// Update existing Instance
	_, response, err := r.temboInstanceConfig.client.InstanceAPI.PatchInstance(
		ctx,
		plan.OrgId.ValueString(),
		plan.InstanceID.ValueString()).PatchInstance(updateInstance).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Tembo Instance",
			"Could not update instance, unexpected error: "+err.Error()+" details:"+getErrorFromResponse(response),
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
			setTemboInstanceResourceModel(&plan, &instance, true, false, &resp.Diagnostics)
			break
		}
		time.Sleep(10 * time.Second)
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
	_, response, err := (r.temboInstanceConfig.client.InstanceAPI).DeleteInstance(ctx, state.OrgId.ValueString(), state.InstanceID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Tembo Instance",
			"Could not delete instance, unexpected error: "+err.Error()+" details:"+getErrorFromResponse(response),
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
	}
}

func (r *temboInstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Resource Import Invalid ID", fmt.Sprintf("wrong format of import ID (%s), use: 'org_id/instance_id'", req.ID))
		return
	}
	orgId := parts[0]
	instanceId := parts[1]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), orgId)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("instance_id"), instanceId)...)
}

func setTemboInstanceResourceModel(instanceResourceModel *temboInstanceResourceModel,
	instance *temboclient.Instance, isUpdateMode bool, isRestoreMode bool, diagnostics *diag.Diagnostics) {
	if isUpdateMode {
		instanceResourceModel.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
	}

	instanceResourceModel.InstanceID = types.StringValue(instance.GetInstanceId())
	instanceResourceModel.InstanceName = types.StringValue(instance.InstanceName)
	instanceResourceModel.OrgId = types.StringValue(instance.GetOrganizationId())
	instanceResourceModel.CPU = types.StringValue(string(instance.GetCpu()))
	instanceResourceModel.Environment = types.StringValue(string(instance.GetEnvironment()))
	instanceResourceModel.Memory = types.StringValue(string(instance.GetMemory()))
	instanceResourceModel.Storage = types.StringValue(string(instance.GetStorage()))
	instanceResourceModel.State = types.StringValue(string(instance.GetState()))

	if len(instance.ExtraDomainsRw) > 0 {
		var localExtraDomainsRw []basetypes.StringValue
		for _, domain := range instance.ExtraDomainsRw {
			localExtraDomainsRw = append(localExtraDomainsRw, types.StringValue(domain))
		}
		instanceResourceModel.ExtraDomainsRw = localExtraDomainsRw
	}

	if instance.ConnectionPooler.Get() != nil {
		var localConnectionPooler ConnectionPooler
		cp := instance.ConnectionPooler.Get()
		localConnectionPooler.Enabled = types.BoolValue(*cp.Enabled)
		localConnectionPooler.Pooler.PoolMode = types.StringValue(string(*cp.Pooler.PoolMode.Ptr()))
		localConnectionPooler.Pooler.Parameters = cp.Pooler.Parameters
		instanceResourceModel.ConnectionPooler = &localConnectionPooler
	}

	// Restore Mode only sets above fields so skipping the other fields.
	if isRestoreMode {
		return
	}

	instanceResourceModel.StackType = types.StringValue(string(instance.StackType))
	instanceResourceModel.Replicas = types.Int64Value(int64(instance.GetReplicas()))

	if len(instance.PostgresConfigs) > 0 {
		var localPGConfigs []KeyValue
		for _, pgConfig := range instance.PostgresConfigs {
			localPGConfigs = append(localPGConfigs, KeyValue{Name: types.StringValue(pgConfig.Name), Value: types.StringValue(pgConfig.Value)})
		}
		instanceResourceModel.PostgresConfigs = localPGConfigs
	}

	if isUpdateMode && len(instance.TrunkInstalls) > 0 {
		for _, trunkInstall := range instance.TrunkInstalls {
			if trunkInstall.Error {
				diagnostics.AddError(
					"Error with Trunk Install:",
					"unexpected error: "+trunkInstall.GetErrorMessage(),
				)
			}
		}
	}

	if isUpdateMode && len(instance.Extensions) > 0 {
		for _, extension := range instance.Extensions {
			if len(extension.Locations) > 0 {
				for _, elocation := range extension.Locations {
					if elocation.GetError() {
						diagnostics.AddError(
							"Error with Extension Installation:",
							"unexpected error: "+elocation.GetErrorMessage(),
						)
					}
				}
			}
		}
	}

	if len(instance.IpAllowList) > 0 {
		var localIpAllowList []basetypes.StringValue
		for _, domain := range instance.IpAllowList {
			localIpAllowList = append(localIpAllowList, types.StringValue(domain))
		}
		instanceResourceModel.IpAllowList = localIpAllowList
	}

	if instance.FirstRecoverabilityTime.IsSet() && instance.FirstRecoverabilityTime.Get() != nil {
		fmt.Println("Setting FirstRecoverabilityTime")
		instanceResourceModel.FirstRecoverabilityTime = types.StringValue(instance.FirstRecoverabilityTime.Get().Format(time.RFC3339))
	} else {
		instanceResourceModel.FirstRecoverabilityTime = types.StringNull()
	}

	if autoscaling, ok := instance.GetAutoscalingOk(); ok && autoscaling != nil {
		// Only set autoscaling if it was in the original plan
		if instanceResourceModel.Autoscaling != nil {
			autoscalingModel := &AutoscalingModel{}

			// Handle Autostop
			if autostop, ok := autoscaling.GetAutostopOk(); ok && autostop != nil {
				autoscalingModel.Autostop = &AutoscalingAutostopModel{
					Enabled: types.BoolValue(autostop.GetEnabled()),
				}
			}

			// Handle Storage
			if storage, ok := autoscaling.GetStorageOk(); ok && storage != nil {
				storageModel := &AutoscalingStorageModel{
					Enabled: types.BoolValue(storage.GetEnabled()),
				}
				// Safely handle maximum value
				if maximum, ok := storage.GetMaximumOk(); ok && maximum != nil {
					storageModel.Maximum = types.StringValue(string(*maximum))
				}
				autoscalingModel.Storage = storageModel
			}

			instanceResourceModel.Autoscaling = autoscalingModel
		}
	}
}

func getStringArray(inputArray []basetypes.StringValue) []string {
	var localStringArray []string
	if len(inputArray) > 0 {
		for _, input := range inputArray {
			localStringArray = append(localStringArray, input.ValueString())
		}
	}
	return localStringArray
}

func getPgConfig(postgresConfigs []KeyValue) []temboclient.PgConfig {
	var localPGConfigs []temboclient.PgConfig
	if len(postgresConfigs) > 0 {
		for _, pgConfig := range postgresConfigs {
			localPGConfigs = append(localPGConfigs, temboclient.PgConfig{Name: pgConfig.Name.ValueString(), Value: pgConfig.Value.ValueString()})
		}
	}
	return localPGConfigs
}

func getTemboTrunkInstalls(trunkInstalls []TrunkInstall) []temboclient.TrunkInstall {
	var localTrunkInstalls []temboclient.TrunkInstall
	if len(trunkInstalls) > 0 {
		for _, trunkInstall := range trunkInstalls {
			localTrunkInstalls = append(localTrunkInstalls, getTemboTrunkInstall(trunkInstall))
		}
	}
	return localTrunkInstalls
}

func getTemboTrunkInstall(trunkInstall TrunkInstall) temboclient.TrunkInstall {
	localTrunkInstall := temboclient.TrunkInstall{
		Name: trunkInstall.Name.ValueString(),
	}
	localTrunkInstall.SetVersion(trunkInstall.Version.ValueString())

	return localTrunkInstall
}

func getConnectionPooler(connectionPooler *ConnectionPooler) *temboclient.ConnectionPooler {
	localConnectionPooler := temboclient.ConnectionPooler{
		Enabled: connectionPooler.Enabled.ValueBoolPointer(),
	}

	localConnectionPooler.Pooler = &temboclient.PgBouncer{
		Parameters: connectionPooler.Pooler.Parameters,
		PoolMode:   (*temboclient.PoolerPgbouncerPoolMode)(connectionPooler.Pooler.PoolMode.ValueStringPointer()),
	}

	return &localConnectionPooler
}

func getTemboExtensions(extensions []Extension) []temboclient.Extension {
	var tcExtensions []temboclient.Extension

	if len(extensions) > 0 {
		for i := 0; i <= len(extensions)-1; i += 1 {
			tcExtensions = append(tcExtensions, getTemboExtension(extensions[i]))
		}
	}
	return tcExtensions
}

func getTemboExtension(extension Extension) temboclient.Extension {
	ext := temboclient.Extension{
		Name: extension.Name,
	}

	ext.SetDescription(extension.Description.ValueString())

	var location []temboclient.ExtensionInstallLocation

	if extension.Locations == nil {
		return ext
	}

	for i := 0; i <= len(extension.Locations)-1; i += 1 {
		eil := temboclient.ExtensionInstallLocation{
			Database: extension.Locations[i].Database.ValueStringPointer(),
			Enabled:  extension.Locations[i].Enabled.ValueBool(),
		}

		eil.SetSchema(extension.Locations[i].Schema.ValueString())
		eil.SetVersion(extension.Locations[i].Version.ValueString())

		location = append(location, eil)
	}
	ext.SetLocations(location)

	return ext
}

func getErrorFromResponse(response *http.Response) string {
	localVarBody, _ := io.ReadAll(response.Body)
	response.Body.Close()

	return string(localVarBody)
}

func getInstanceState(r *temboInstanceResource, ctx context.Context,
	orgId string, instanceId string, diagnostics *diag.Diagnostics) temboclient.State {
	refreshInstance, response, err := r.temboInstanceConfig.client.InstanceAPI.GetInstance(ctx, orgId, instanceId).Execute()

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
	refreshInstance, response, err := r.temboInstanceConfig.client.InstanceAPI.GetInstance(ctx, orgId, instanceId).Execute()
	if err != nil {
		diagnostics.AddError(
			"Error Reading Tembo Instance State",
			"Could not read Tembo Instance ID "+instanceId+": "+getErrorFromResponse(response),
		)
	}

	return *refreshInstance, err
}

func getStringArrayForType[T ~string](inputArray []T) []string {
	var returnArray []string
	if len(inputArray) > 0 {
		for _, input := range inputArray {
			returnArray = append(returnArray, string(input))
		}
	}
	return returnArray
}
