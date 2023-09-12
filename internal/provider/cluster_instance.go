package provider

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	temboclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &temboInstanceResource{}
	_ resource.ResourceWithConfigure = &temboInstanceResource{}
)

// temboInstanceResourceModel maps the resource schema data.
type temboInstanceResourceModel struct {
	InstanceID   types.String `tfsdk:"instance_id"`
	InstanceName types.String `tfsdk:"instance_name"`
	OrgId        types.String `tfsdk:"org_id"`
	CPU          types.String `tfsdk:"cpu"`
	StackType    types.String `tfsdk:"stack_type"`
	Environment  types.String `tfsdk:"environment"`
	Memory       types.String `tfsdk:"memory"`
	Storage      types.String `tfsdk:"storage"`
	LastUpdated  types.String `tfsdk:"last_updated"`
	State        types.String `tfsdk:"state"`
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

	// TODO: Figure out a better way to set this so it doesn't have to be be called in each method.
	ctx = context.WithValue(ctx, temboclient.ContextAccessToken, r.temboInstanceConfig.accessToken)

	instanceRequest := r.temboInstanceConfig.client.InstanceApi.CreateInstance(ctx,
		plan.OrgId.ValueString())

	instance, _, err := instanceRequest.CreateInstance(createInstance).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Tembo Instance:",
			"Could not create Tembo Instance, unexpected error: "+err.Error(),
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
	setTemboInstanceResourceModel(&plan, instance, true)

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
	instance, _, err := r.temboInstanceConfig.client.InstanceApi.GetInstance(ctx, state.OrgId.ValueString(), state.InstanceID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Tembo Instance",
			"Could not read Tembo Instance ID "+state.InstanceID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite items with refreshed state
	setTemboInstanceResourceModel(state, instance, false)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func setTemboInstanceResourceModel(instanceResourceModel *temboInstanceResourceModel, instance *temboclient.Instance, updateComputedValue bool) {
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
		1,
		temboclient.Storage(plan.Storage.ValueString()))

	log.Printf("[INFO] Tembo instanceID %s", plan)

	ctx = context.WithValue(ctx, temboclient.ContextAccessToken, r.temboInstanceConfig.accessToken)

	// Update existing Instance
	_, _, err := r.temboInstanceConfig.client.InstanceApi.PutInstance(
		ctx,
		plan.OrgId.ValueString(),
		plan.InstanceID.ValueString()).UpdateInstance(updateInstance).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Tembo Instance",
			"Could not update Tembo Instance ID "+plan.InstanceID.ValueString()+": "+err.Error(),
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
			updateInstance.GetStorage() == instance.GetStorage() {

			// Update resource state with updated items and timestamp
			setTemboInstanceResourceModel(&plan, &instance, true)
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
	_, _, err := (*r.temboInstanceConfig.client.InstanceApi).DeleteInstance(ctx, state.OrgId.ValueString(), state.InstanceID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Tembo Instance",
			"Could not delete instance, unexpected error: "+err.Error(),
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

func getInstanceState(r *temboInstanceResource, ctx context.Context,
	orgId string, instanceId string, diagnostics *diag.Diagnostics) temboclient.State {
	refreshInstance, _, err := r.temboInstanceConfig.client.InstanceApi.GetInstance(ctx, orgId, instanceId).Execute()
	if err != nil {
		diagnostics.AddError(
			"Error Reading Tembo Instance State",
			"Could not read Tembo Instance ID "+instanceId+": "+err.Error(),
		)
		return temboclient.ERROR
	}

	return refreshInstance.GetState()
}

func getInstance(r *temboInstanceResource, ctx context.Context,
	orgId string, instanceId string, diagnostics *diag.Diagnostics) (temboclient.Instance, error) {
	refreshInstance, _, err := r.temboInstanceConfig.client.InstanceApi.GetInstance(ctx, orgId, instanceId).Execute()
	if err != nil {
		diagnostics.AddError(
			"Error Reading Tembo Instance State",
			"Could not read Tembo Instance ID "+instanceId+": "+err.Error(),
		)
	}

	return *refreshInstance, err
}
