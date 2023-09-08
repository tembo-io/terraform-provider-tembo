package provider

import (
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testProviderConfig is a shared configuration to combine with the actual
// test configuration so the Tembo client is properly configured.
// It is also possible to use the TEMBO_ environment variables instead,
// such as updating the Makefile and running the testing through that tool.
func testProviderConfig() string {
	host := os.Getenv("TEMBO_HOST")
	access_token := os.Getenv("TEMBO_ACCESS_TOKEN")

	return fmt.Sprintf(`
provider "tembo" {
  host     = "%v"
  access_token = "%v"
}`, host, access_token)

}

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"tembo": providerserver.NewProtocol6WithError(New("test")()),
}
