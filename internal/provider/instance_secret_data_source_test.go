package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestInstanceSecretDataSource(t *testing.T) {
	instanceName := generateInstanceName()
	orgId := os.Getenv("ORG_ID")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testProviderConfig() + testInstanceResourceCreateConfig(instanceName, orgId) + testInstanceSecretCreateConfig(orgId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.tembo_instance_secret.test_readonly", "secrets.username"),
					resource.TestCheckResourceAttrSet("data.tembo_instance_secret.test_readonly", "secrets.password"),
					resource.TestCheckResourceAttrSet("data.tembo_instance_secret.test_certificate", "secrets.ca.crt"),
				),
			},
		},
	})
}

func testInstanceSecretCreateConfig(orgId string) string {
	return fmt.Sprintf(`
	data "tembo_instance_secret" "test_readonly" {
		org_id      = "%v"
		instance_id = tembo_instance.test.instance_id
		secret_name = "readonly-role"	  
	}
	data "tembo_instance_secret" "test_certificate" {
		org_id      = "%v"
		instance_id = tembo_instance.test.instance_id
		secret_name = "certificate"	
	}  
	`, orgId, orgId)
}
