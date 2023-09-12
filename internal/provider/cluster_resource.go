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
	_ resource.Resource              = &temboClusterResource{}
	_ resource.ResourceWithConfigure = &temboClusterResource{}
)

// temboClusterResourceModel maps the resource schema data.
type temboClusterResourceModel struct {
	ClusterID      types.String `tfsdk:"cluster_id"`
	ClusterName    types.String `tfsdk:"cluster_name"`
	OrganizationId types.String `tfsdk:"organization_id"`
	CPU            types.String `tfsdk:"cpu"`
	Stack          types.String `tfsdk:"stack"`
	Environment    types.String `tfsdk:"environment"`
	Memory         types.String `tfsdk:"memory"`
	Storage        types.String `tfsdk:"storage"`
	LastUpdated    types.String `tfsdk:"last_updated"`
	State          types.String `tfsdk:"state"`
}

// Configure adds the provider configured data to the resource.
func (r *temboClusterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	temboClusterConfig, ok := req.ProviderData.(clusterConfig)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *clusterConfig, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.temboClusterConfig = temboClusterConfig
}

// NewTemboClusterResource is a helper function to simplify the provider implementation.
func NewTemboClusterResource() resource.Resource {
	return &temboClusterResource{}
}

// temboClusterResource is the resource implementation.
type temboClusterResource struct {
	temboClusterConfig clusterConfig
}

// Tembo CLuster Configuration.
type clusterConfig struct {
	client      *temboclient.APIClient
	accessToken string
}

// Metadata returns the resource type name.
func (r *temboClusterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

// Schema defines the schema for the resource.
func (r *temboClusterResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"cluster_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"cluster_name": schema.StringAttribute{
				Required: true,
			},
			"organization_id": schema.StringAttribute{
				Required: true,
			},
			"cpu": schema.StringAttribute{
				Required: true,
			},
			"stack": schema.StringAttribute{
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
func (r *temboClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan temboClusterResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call API to Create Tembo Cluster
	createInstance := *temboclient.NewCreateInstance(
		temboclient.Cpu(plan.CPU.ValueString()),
		temboclient.Environment(plan.Environment.ValueString()),
		plan.ClusterName.ValueString(),
		temboclient.Memory(plan.Memory.ValueString()),
		temboclient.EntityType(plan.Stack.ValueString()),
		temboclient.Storage(plan.Storage.ValueString()))

	// TODO: Figure out a better way to set this so it doesn't have to be be called in each method.
	ctx = context.WithValue(ctx, temboclient.ContextAccessToken, r.temboClusterConfig.accessToken)

	instance := r.temboClusterConfig.client.InstanceApi.CreateInstance(ctx,
		plan.OrganizationId.ValueString())

	cluster, _, err := instance.CreateInstance(createInstance).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Tembo Cluster:",
			"Could not create Tembo Cluster, unexpected error: "+err.Error(),
		)
		return
	}

	// Wait until it's created.
	for {
		clusterState := getClusterState(r, ctx, plan.OrganizationId.ValueString(), cluster.GetInstanceId(), &resp.Diagnostics)

		if clusterState == temboclient.ERROR || clusterState == temboclient.UP {
			break
		}

		time.Sleep(10 * time.Second)
		log.Printf("[INFO] Waiting for Tembo cluster %s to be UP", plan.ClusterName)
	}

	log.Printf("[INFO] Tembo cluster %s has been successfully created", plan.ClusterName)

	// Map response body to schema and populate Computed attribute values
	cluster.SetState(temboclient.UP)
	setTemboClusterResourceModel(&plan, cluster, true)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *temboClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state *temboClusterResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, temboclient.ContextAccessToken, r.temboClusterConfig.accessToken)
	// Get refreshed Cluster value from API
	cluster, _, err := r.temboClusterConfig.client.InstanceApi.GetInstance(ctx, state.OrganizationId.ValueString(), state.ClusterID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Tembo Cluster",
			"Could not read Tembo Cluster ID "+state.ClusterID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite items with refreshed state
	setTemboClusterResourceModel(state, cluster, false)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func setTemboClusterResourceModel(clusterResourceModel *temboClusterResourceModel, instance *temboclient.Instance, updateComputedValue bool) {
	if updateComputedValue {
		clusterResourceModel.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
	}

	clusterResourceModel.ClusterID = types.StringValue(instance.GetInstanceId())
	clusterResourceModel.ClusterName = types.StringValue(instance.InstanceName)
	clusterResourceModel.OrganizationId = types.StringValue(instance.GetOrganizationId())
	clusterResourceModel.CPU = types.StringValue(string(instance.GetCpu()))
	clusterResourceModel.Stack = types.StringValue(string(instance.StackType))
	clusterResourceModel.Environment = types.StringValue(string(instance.GetEnvironment()))
	clusterResourceModel.Memory = types.StringValue(string(instance.GetMemory()))
	clusterResourceModel.Storage = types.StringValue(string(instance.GetStorage()))
	clusterResourceModel.State = types.StringValue(string(instance.GetState()))
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *temboClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan temboClusterResourceModel
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

	log.Printf("[INFO] Tembo clusterID %s", plan)

	ctx = context.WithValue(ctx, temboclient.ContextAccessToken, r.temboClusterConfig.accessToken)

	// Update existing Cluster
	_, _, err := r.temboClusterConfig.client.InstanceApi.PutInstance(
		ctx,
		plan.OrganizationId.ValueString(),
		plan.ClusterID.ValueString()).UpdateInstance(updateInstance).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Tembo Cluster",
			"Could not update Tembo Cluster ID "+plan.ClusterID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Wait until it's updated
	for {
		cluster, err := getCluster(r, ctx, plan.OrganizationId.ValueString(), plan.ClusterID.ValueString(), &resp.Diagnostics)

		if err != nil {
			return
		}

		if updateInstance.GetCpu() == cluster.GetCpu() &&
			updateInstance.GetMemory() == cluster.GetMemory() &&
			updateInstance.GetStorage() == cluster.GetStorage() {

			// Update resource state with updated items and timestamp
			setTemboClusterResourceModel(&plan, &cluster, true)
			break
		}
		time.Sleep(10 * time.Second)

		log.Printf("[INFO] Waiting for Tembo cluster %s to be updated", plan.ClusterName)
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *temboClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state temboClusterResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, temboclient.ContextAccessToken, r.temboClusterConfig.accessToken)

	// Delete existing Tembo Cluster
	_, _, err := (*r.temboClusterConfig.client.InstanceApi).DeleteInstance(ctx, state.OrganizationId.ValueString(), state.ClusterID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Tembo Cluster",
			"Could not delete cluster, unexpected error: "+err.Error(),
		)
		return
	}

	// Wait until it's deleted
	for {
		clusterState := getClusterState(r, ctx, state.OrganizationId.ValueString(), state.ClusterID.ValueString(), &resp.Diagnostics)

		if clusterState == temboclient.ERROR || clusterState == temboclient.DELETED {
			break
		}

		time.Sleep(10 * time.Second)
		log.Printf("[INFO] Waiting for Tembo cluster %s to be DELETED", state.ClusterName)
	}

	log.Printf("[INFO] Tembo cluster %s has been successfully deleted", state.ClusterName)
}

func getClusterState(r *temboClusterResource, ctx context.Context,
	organizationId string, clusterId string, diagnostics *diag.Diagnostics) temboclient.State {
	refreshCluster, _, err := r.temboClusterConfig.client.InstanceApi.GetInstance(ctx, organizationId, clusterId).Execute()
	if err != nil {
		diagnostics.AddError(
			"Error Reading Tembo Cluster State",
			"Could not read Tembo Cluster ID "+clusterId+": "+err.Error(),
		)
		return temboclient.ERROR
	}

	return refreshCluster.GetState()
}

func getCluster(r *temboClusterResource, ctx context.Context,
	organizationId string, clusterId string, diagnostics *diag.Diagnostics) (temboclient.Instance, error) {
	refreshCluster, _, err := r.temboClusterConfig.client.InstanceApi.GetInstance(ctx, organizationId, clusterId).Execute()
	if err != nil {
		diagnostics.AddError(
			"Error Reading Tembo Cluster State",
			"Could not read Tembo Cluster ID "+clusterId+": "+err.Error(),
		)
	}

	return *refreshCluster, err
}
