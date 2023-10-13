package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestInstanceSecretsDataSource(t *testing.T) {
	instanceName := generateInstanceName()
	orgId := os.Getenv("ORG_ID")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testProviderConfig() + testInstanceResourceCreateConfig(instanceName, orgId) + testInstanceSecretsCreateConfig(orgId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.tembo_instance_secrets.test_ds", "available_secrets.#", "4"),
					resource.TestCheckResourceAttr("data.tembo_instance_secrets.test_ds", "available_secrets.0.name", "app-role"),
					resource.TestCheckResourceAttr("data.tembo_instance_secrets.test_ds", "available_secrets.0.possible_keys.#", "2"),
					resource.TestCheckResourceAttr("data.tembo_instance_secrets.test_ds", "available_secrets.0.possible_keys.0", "username"),
					resource.TestCheckResourceAttr("data.tembo_instance_secrets.test_ds", "available_secrets.0.possible_keys.1", "password"),
					resource.TestCheckResourceAttr("data.tembo_instance_secrets.test_ds", "available_secrets.1.name", "readonly-role"),
					resource.TestCheckResourceAttr("data.tembo_instance_secrets.test_ds", "available_secrets.1.possible_keys.#", "2"),
					resource.TestCheckResourceAttr("data.tembo_instance_secrets.test_ds", "available_secrets.1.possible_keys.0", "username"),
					resource.TestCheckResourceAttr("data.tembo_instance_secrets.test_ds", "available_secrets.1.possible_keys.1", "password"),
					resource.TestCheckResourceAttr("data.tembo_instance_secrets.test_ds", "available_secrets.2.name", "superuser-role"),
					resource.TestCheckResourceAttr("data.tembo_instance_secrets.test_ds", "available_secrets.2.possible_keys.#", "2"),
					resource.TestCheckResourceAttr("data.tembo_instance_secrets.test_ds", "available_secrets.2.possible_keys.0", "username"),
					resource.TestCheckResourceAttr("data.tembo_instance_secrets.test_ds", "available_secrets.2.possible_keys.1", "password"),
					resource.TestCheckResourceAttr("data.tembo_instance_secrets.test_ds", "available_secrets.3.name", "certificate"),
					resource.TestCheckResourceAttr("data.tembo_instance_secrets.test_ds", "available_secrets.3.possible_keys.#", "1"),
					resource.TestCheckResourceAttr("data.tembo_instance_secrets.test_ds", "available_secrets.3.possible_keys.0", "ca.crt"),
				),
			},
		},
	})
}

func testInstanceSecretsCreateConfig(orgId string) string {
	return fmt.Sprintf(`
	data "tembo_instance_secrets" "test_ds" {
		org_id      = "%v"
		instance_id = tembo_instance.test.instance_id	  
	}`, orgId)
}
