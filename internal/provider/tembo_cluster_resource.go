package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	temboclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &temboClusterResource{}
)

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
			"cluster_name": schema.StringAttribute{
				Required: true,
			},
			"organization_id": schema.StringAttribute{
				Required: true,
			},
			"cpu": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *temboClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	createCluster := *temboclient.NewCreateCluster(temboclient.Cpu("1"), temboclient.Environment("dev"), "InstanceName_example", temboclient.Memory("1Gi"), temboclient.Storage("10Gi")) // CreateCluster |
	instance := r.client.InstancesApi.CreateInstance(context.Background(), "test", temboclient.STANDARD)

	_, respo, err := instance.CreateCluster(createCluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.CreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", respo)
	}
	// Set state to fully populated data
}

// Read refreshes the Terraform state with the latest data.
func (r *temboClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *temboClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *temboClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
