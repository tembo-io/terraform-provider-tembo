package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccPreCheck(t *testing.T) {
	if os.Getenv("ORG_ID") == "" {
		t.Fatal("ORG_ID must be set for acceptance tests")
	}
}

func TestInstanceDataSource(t *testing.T) {
	instanceName := generateInstanceName()
	orgId := os.Getenv("ORG_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testProviderConfig() + testInstanceResourceCreateConfig(instanceName, orgId) + testInstanceDataSourceConfig(orgId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"data.tembo_instance.test", "instance_name",
						"tembo_instance.test", "instance_name",
					),
					resource.TestCheckResourceAttrPair(
						"data.tembo_instance.test", "cpu",
						"tembo_instance.test", "cpu",
					),
					resource.TestCheckResourceAttrPair(
						"data.tembo_instance.test", "stack_type",
						"tembo_instance.test", "stack_type",
					),
					resource.TestCheckResourceAttrPair(
						"data.tembo_instance.test", "environment",
						"tembo_instance.test", "environment",
					),
					resource.TestCheckResourceAttrPair(
						"data.tembo_instance.test", "ip_allow_list",
						"tembo_instance.test", "ip_allow_list",
					),
					resource.TestCheckResourceAttr("data.tembo_instance.test", "autoscaling.storage.0.enabled", "true"),
					resource.TestCheckResourceAttr("data.tembo_instance.test", "autoscaling.storage.0.max_gb", "50Gi"),
					resource.TestCheckResourceAttr("data.tembo_instance.test", "autoscaling.autostop.0.enabled", "false"),
				),
			},
		},
	})
}

func testInstanceDataSourceConfig(orgId string) string {
	return fmt.Sprintf(`
data "tembo_instance" "test" {
	org_id      = "%v"
	instance_id = tembo_instance.test.instance_id
}
`, orgId)
}
