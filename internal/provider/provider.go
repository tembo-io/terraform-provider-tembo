package provider

import (
	"context"
	"net/url"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	temboclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
	tembodataclient "github.com/tembo-io/terraform-provider-tembo/tembodataclient"
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
	Host              types.String `tfsdk:"host"`
	HostURLScheme     types.String `tfsdk:"host_url_scheme"`
	DataHost          types.String `tfsdk:"data_host"`
	DataHostURLScheme types.String `tfsdk:"data_host_url_scheme"`
	AccessToken       types.String `tfsdk:"access_token"`
}

// Create a struct to hold both clients.
type providerData struct {
	apiClient   *temboclient.APIClient
	dataClient  *tembodataclient.APIClient
	accessToken string
}

// Metadata returns the provider type name.
func (p *temboProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "tembo"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *temboProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Use the Tembo provider to interact with the resources supported by Tembo Cloud.",
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				MarkdownDescription: "Tembo API the provider should connect to. By default it connects to Tembo Cloud Production API.",
				Optional:            true,
			},
			"host_url_scheme": schema.StringAttribute{
				MarkdownDescription: "Url scheme for Tembo API. By default it is https.",
				Optional:            true,
			},
			"data_host": schema.StringAttribute{
				MarkdownDescription: "Tembo Data API the provider should connect to. By default it connects to Tembo Data Cloud Production API.",
				Optional:            true,
			},
			"data_host_url_scheme": schema.StringAttribute{
				MarkdownDescription: "Url scheme for Tembo Data API. By default it is https.",
				Optional:            true,
			},
			"access_token": schema.StringAttribute{
				MarkdownDescription: "Access Token generated using steps [here](https://tembo.io/docs/tembo-cloud/api#create-a-long-lived-api-token).",
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
			"Unknown Tembo API Host",
			"The provider cannot create the tembo API client as there is an unknown configuration value for the tembo API Host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TEMBO_HOST environment variable.",
		)
	}

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("data_host"),
			"Unknown Tembo API Data Host",
			"The provider cannot create the tembo API client as there is an unknown configuration value for the Tembo API Data Host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TEMBO_DATA_HOST environment variable.",
		)
	}

	if config.AccessToken.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("access_token"),
			"Unknown Tembo API Access Token",
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
	host_url_scheme := os.Getenv("TEMBO_HOST_URL_SCHEME")
	data_host := os.Getenv("TEMBO_DATA_HOST")
	data_host_url_scheme := os.Getenv("TEMBO_DATA_HOST_URL_SCHEME")
	access_token := os.Getenv("TEMBO_ACCESS_TOKEN")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if host == "" {
		host = "https://api.coredb.io"
	}

	if !config.HostURLScheme.IsNull() {
		host_url_scheme = config.HostURLScheme.ValueString()
	}

	if host_url_scheme == "" {
		host_url_scheme = "https"
	}

	if !config.DataHost.IsNull() {
		data_host = config.DataHost.ValueString()
	}

	if data_host == "" {
		data_host = "https://api.data-1.use1.tembo.io"
	}

	if !config.DataHostURLScheme.IsNull() {
		data_host_url_scheme = config.DataHostURLScheme.ValueString()
	}

	if data_host_url_scheme == "" {
		data_host_url_scheme = "https"
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

	hostUrl, err := url.Parse(host)
	if err != nil {
		panic(err)
	}

	configuration.Scheme = host_url_scheme

	configuration.Host = hostUrl.Host

	// Create a new tembo client using the configuration values
	client := temboclient.NewAPIClient(configuration)

	data_configuration := tembodataclient.NewConfiguration()

	dataHostUrl, err := url.Parse(data_host)
	if err != nil {
		panic(err)
	}

	data_configuration.Scheme = data_host_url_scheme

	data_configuration.Host = dataHostUrl.Host

	// Create a new tembo data client using the configuration values
	dataclient := tembodataclient.NewAPIClient(data_configuration)

	// Make the tembo client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = providerData{
		apiClient:   client,
		dataClient:  dataclient,
		accessToken: access_token,
	}

	resp.ResourceData = instanceConfig{
		client:      client,
		accessToken: access_token,
	}
}

// DataSources defines the data sources implemented in the provider.
func (p *temboProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewTemboInstanceSecretsDataSource,
		NewTemboInstanceSecretDataSource,
		NewTemboInstanceDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *temboProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewTemboInstanceResource,
	}
}
