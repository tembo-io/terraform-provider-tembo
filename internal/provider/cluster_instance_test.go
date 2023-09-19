package provider

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestTemboInstanceResource(t *testing.T) {

	instanceName := generateInstanceName()
	orgId := os.Getenv("ORG_ID")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testProviderConfig() + testInstanceResourceCreateConfig(instanceName, orgId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tembo_instance.test", "instance_name", instanceName),
					resource.TestCheckResourceAttr("tembo_instance.test", "org_id", orgId),
					resource.TestCheckResourceAttr("tembo_instance.test", "cpu", "1"),
					resource.TestCheckResourceAttr("tembo_instance.test", "stack_type", "Standard"),
					resource.TestCheckResourceAttr("tembo_instance.test", "replicas", "1"),
					resource.TestCheckResourceAttr("tembo_instance.test", "environment", "dev"),
					resource.TestCheckResourceAttr("tembo_instance.test", "memory", "4Gi"),
					resource.TestCheckResourceAttr("tembo_instance.test", "storage", "10Gi"),
					resource.TestCheckResourceAttrSet("tembo_instance.test", "instance_id"),
					resource.TestCheckResourceAttrSet("tembo_instance.test", "last_updated"),
					resource.TestCheckResourceAttr("tembo_instance.test", "postgres_configs.#", "1"),
				),
			},
			// TODO: Add ImportState testing
			// Update and Read testing
			{
				Config: testProviderConfig() + testInstanceResourceUpdateConfig(instanceName, orgId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tembo_instance.test", "instance_name", instanceName),
					resource.TestCheckResourceAttr("tembo_instance.test", "org_id", orgId),
					resource.TestCheckResourceAttr("tembo_instance.test", "cpu", "2"),
					resource.TestCheckResourceAttr("tembo_instance.test", "stack_type", "Standard"),
					resource.TestCheckResourceAttr("tembo_instance.test", "replicas", "2"),
					resource.TestCheckResourceAttr("tembo_instance.test", "environment", "dev"),
					resource.TestCheckResourceAttr("tembo_instance.test", "memory", "8Gi"),
					resource.TestCheckResourceAttr("tembo_instance.test", "storage", "10Gi"),
					resource.TestCheckResourceAttrSet("tembo_instance.test", "instance_id"),
					resource.TestCheckResourceAttrSet("tembo_instance.test", "last_updated"),
					resource.TestCheckResourceAttr("tembo_instance.test", "postgres_configs.#", "2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testInstanceResourceCreateConfig(instanceName string, orgId string) string {
	return fmt.Sprintf(`
resource "tembo_instance" "test" {
	instance_name   = "%v"
	org_id          = "%v"
	cpu             = "1"
	stack_type      = "Standard"
	replicas        = 1
	environment     = "dev"
	memory          = "4Gi"
	storage         = "10Gi"
	postgres_configs = [
		{
		  name = "max_connections"
		  value = "200"
		}
	  ]
	  trunk_installs = [
		{
          name    = "pgmq"
		  version = "0.24.0"
		}
	  ]
  }
	`, instanceName, orgId)
}

func testInstanceResourceUpdateConfig(instanceName string, orgId string) string {
	return fmt.Sprintf(`
resource "tembo_instance" "test" {
	instance_name   = "%v"
	org_id 			= "%v"
	cpu             = "2"
	stack_type      = "Standard"
	replicas        = 2
	environment     = "dev"
	memory          = "8Gi"
	storage         = "10Gi"
	postgres_configs = [
		{
		  name = "max_connections"
		  value = "200"
		},
		{
		  name = "wal_buffers"
		  value = "10"
		}
	  ]
	trunk_installs = [
		{
          name    = "pgmq"
		  version = "0.24.0"
		},
		{
          name    = "pg_stat_statements"
		  version = "1.10.0"
		},
	  ]
  }
	`, instanceName, orgId)
}

func generateInstanceName() string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return fmt.Sprintf("tf-test-%v", string(b))
}
