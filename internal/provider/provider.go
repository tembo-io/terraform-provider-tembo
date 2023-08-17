package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	temboclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &temboProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &temboProvider{
			version: version,
		}
	}
}

// temboProvider is the provider implementation.
type temboProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// temboProviderModel maps provider schema data to a Go type.
type temboProviderModel struct {
	Host        types.String `tfsdk:"host"`
	AccessToken types.String `tfsdk:"access_token"`
}

// Metadata returns the provider type name.
func (p *temboProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "tembo"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *temboProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				MarkdownDescription: "Host",
				Optional:            true,
			},
			"access_token": schema.StringAttribute{
				MarkdownDescription: "Access Token",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *temboProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config temboProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown tembo API Host",
			"The provider cannot create the tembo API client as there is an unknown configuration value for the tembo API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the tembo_HOST environment variable.",
		)
	}

	if config.AccessToken.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("access_token"),
			"Unknown tembo API Access Token",
			"The provider cannot create the tembo API client as there is an unknown configuration value for the tembo API password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the tembo_PASSWORD environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	host := os.Getenv("TEMBO_HOST")
	access_token := os.Getenv("TEMBO_ACCESS_TOKEN")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.AccessToken.IsNull() {
		access_token = config.AccessToken.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing tembo API Host",
			"The provider cannot create the tembo API client as there is a missing or empty value for the tembo API host. "+
				"Set the host value in the configuration or use the TEMBO_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if access_token == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("access_token"),
			"Missing tembo API Access Token",
			"The provider cannot create the tembo API Access Token as there is a missing or empty value for the tembo API password. "+
				"Set the password value in the configuration or use the TEMBO_ACCESS_TOKEN environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	configuration := temboclient.NewConfiguration()
	// Create a new tembo client using the configuration values
	client := temboclient.NewAPIClient(configuration)

	// Make the tembo client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

// DataSources defines the data sources implemented in the provider.
func (p *temboProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewExampleDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *temboProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewTemboClusterResource,
	}
}
