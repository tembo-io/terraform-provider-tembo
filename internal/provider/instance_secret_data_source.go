package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	tembodataclient "github.com/tembo-io/terraform-provider-tembo/tembodataclient"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &temboInstanceSecret{}
	_ datasource.DataSourceWithConfigure = &temboInstanceSecret{}
)

// NewTemboInstanceSecretDataSource is a helper function to simplify the provider implementation.
func NewTemboInstanceSecretDataSource() datasource.DataSource {
	return &temboInstanceSecret{}
}

// TemboInstanceSecret is the data source implementation.
type temboInstanceSecret struct {
	temboInstanceConfig instanceSecretsConfig
}

// Metadata returns the data source type name.
func (d *temboInstanceSecret) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_instance_secret"
}

// Schema defines the schema for the data source.
func (d *temboInstanceSecret) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data Source for Tembo Instance Secret.",
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				MarkdownDescription: "Id of the organization in which the instance will be created",
				Required:            true,
			},
			"instance_id": schema.StringAttribute{
				MarkdownDescription: "Unique ID for the instance generated by Tembo",
				Required:            true,
			},
			"secret_name": schema.StringAttribute{
				MarkdownDescription: "Secret name",
				Required:            true,
			},
			"secrets": schema.MapAttribute{
				MarkdownDescription: "Secret Key/Values",
				Computed:            true,
				ElementType:         types.StringType,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *temboInstanceSecret) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(providerData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected providerData, got: %T.", req.ProviderData),
		)
		return
	}

	d.temboInstanceConfig = instanceSecretsConfig{
		client:      data.dataClient,
		accessToken: data.accessToken,
	}
}

// temboInstanceSecretModel maps the data source schema data.
type temboInstanceSecretModel struct {
	OrgId      types.String            `tfsdk:"org_id"`
	InstanceId types.String            `tfsdk:"instance_id"`
	SecretName types.String            `tfsdk:"secret_name"`
	Secrets    map[string]types.String `tfsdk:"secrets"`
}

// Read refreshes the Terraform state with the latest data.
func (d *temboInstanceSecret) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Get current state
	var state temboInstanceSecretModel

	ctx = context.WithValue(ctx, tembodataclient.ContextAccessToken, d.temboInstanceConfig.accessToken)

	var orgId string
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("org_id"), &orgId)...)

	var instanceId string
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("instance_id"), &instanceId)...)

	var secretName string
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("secret_name"), &secretName)...)

	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, fmt.Sprintf("error reading terraform plan %v", resp.Diagnostics.Errors()))
		return
	}

	availableSecrets, _, err := d.temboInstanceConfig.client.SecretsAPI.GetSecretNamesV1(ctx, orgId, instanceId).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Tembo Instance Available Secrets",
			err.Error(),
		)
		return
	}

	// Get refreshed Instance value from API
	secret, _, err := d.temboInstanceConfig.client.SecretsAPI.GetSecretV1(ctx, orgId, instanceId, secretName).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Tembo Instance Secret",
			err.Error(),
		)
		return
	}

	localSecret := make(map[string]types.String)

	if len(availableSecrets) > 0 {
		for _, aSecret := range availableSecrets {
			if aSecret.Name == secretName {
				for _, possibleKey := range aSecret.PossibleKeys {
					localSecret[possibleKey] = types.StringValue(secret[possibleKey])
				}
				state.Secrets = localSecret
			}
		}
	}

	// Set refreshed state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
