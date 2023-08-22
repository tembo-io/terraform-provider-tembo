package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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

// Configure adds the provider configured client to the resource.
func (r *temboClusterResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*temboclient.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *tembo.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// NewTemboClusterResource is a helper function to simplify the provider implementation.
func NewTemboClusterResource() resource.Resource {
	return &temboClusterResource{}
}

// temboClusterResource is the resource implementation.
type temboClusterResource struct {
	client *temboclient.APIClient
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
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *temboClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan *temboClusterResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call API to Create Tembo Cluster
	createCluster := *temboclient.NewCreateCluster(temboclient.Cpu(plan.CPU.ValueString()), temboclient.Environment(plan.Environment.ValueString()), plan.ClusterName.ValueString(), temboclient.Memory(plan.Memory.ValueString()), temboclient.Storage(plan.Storage.ValueString()))
	instance := r.client.InstancesApi.CreateInstance(ctx, plan.OrganizationId.ValueString(), temboclient.EntityType(plan.Stack.ValueString()))

	cluster, _, err := instance.CreateCluster(createCluster).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Tembo Cluster:",
			"Could not create Tembo Cluster, unexpected error: "+err.Error(),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	setTemboClusterResourceModel(plan, cluster, true)

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

	// Get refreshed Cluster value from API
	cluster, _, err := r.client.InstancesApi.GetInstance(ctx, state.OrganizationId.ValueString(), state.Stack.ValueString(), state.ClusterID.ValueString()).Execute()
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

func setTemboClusterResourceModel(clusterResourceModel *temboClusterResourceModel, cluster *temboclient.ReadCluster, updateComputedValue bool) {
	if updateComputedValue {
		clusterResourceModel.ClusterID = types.StringValue(cluster.GetInstanceId())
		clusterResourceModel.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
	}

	clusterResourceModel.ClusterName = types.StringValue(cluster.InstanceName)
	clusterResourceModel.OrganizationId = types.StringValue(cluster.GetOrganizationId())
	clusterResourceModel.CPU = types.StringValue(string(cluster.GetCpu()))
	clusterResourceModel.Stack = types.StringValue(string(cluster.EntityType))
	clusterResourceModel.Environment = types.StringValue(string(cluster.GetEnvironment()))
	clusterResourceModel.Memory = types.StringValue(string(cluster.GetMemory()))
	clusterResourceModel.Storage = types.StringValue(string(cluster.GetStorage()))
	clusterResourceModel.State = types.StringValue(string(cluster.GetState()))
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *temboClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *temboClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
