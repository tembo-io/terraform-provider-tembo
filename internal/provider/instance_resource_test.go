package provider

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestTemboInstanceResource(t *testing.T) {

	instanceName := generateInstanceName()
	orgId := os.Getenv("ORG_ID")
	resourceName := "tembo_instance.test"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testProviderConfig() + testInstanceResourceCreateConfig(instanceName, orgId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "instance_name", instanceName),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgId),
					resource.TestCheckResourceAttr(resourceName, "cpu", "1"),
					resource.TestCheckResourceAttr(resourceName, "stack_type", "Standard"),
					resource.TestCheckResourceAttr(resourceName, "replicas", "1"),
					resource.TestCheckResourceAttr(resourceName, "environment", "dev"),
					resource.TestCheckResourceAttr(resourceName, "memory", "4Gi"),
					resource.TestCheckResourceAttr(resourceName, "storage", "10Gi"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(resourceName, "last_updated"),
				),
			},
			// ImportState testing
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "instance_id",
				ImportStateIdFunc:                    testTemboInstanceClientImportStateIDFunc(resourceName),
				ImportStateVerifyIgnore:              []string{"last_updated", "trunk_installs", "extensions"},
			},
			// Update and Read testing
			{
				Config: testProviderConfig() + testInstanceResourceUpdateConfig(instanceName, orgId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "instance_name", instanceName),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgId),
					resource.TestCheckResourceAttr(resourceName, "cpu", "2"),
					resource.TestCheckResourceAttr(resourceName, "stack_type", "Standard"),
					resource.TestCheckResourceAttr(resourceName, "replicas", "2"),
					resource.TestCheckResourceAttr(resourceName, "environment", "dev"),
					resource.TestCheckResourceAttr(resourceName, "memory", "8Gi"),
					resource.TestCheckResourceAttr(resourceName, "storage", "10Gi"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(resourceName, "last_updated"),
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
	trunk_installs = [
	{
		name    = "pgmq"
		version = "0.24.0"
	}]
	extensions = [{
	name        = "plperl"
	description = "PL/Perl procedural language"
		locations = [{
			database = "app"
			schema   = "public"
			version  = "1.0"
			enabled  = false
			},
			{
			database = "postgres"
			schema   = "public"
			version  = "1.0"
			enabled  = true
		}]
	}]
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
	trunk_installs = [
		{
          name    = "pgmq"
		  version = "0.24.0"
		},
		{
          name    = "pg_stat_statements"
		  version = "1.10.0"
		}]
	extensions = [{
	name        = "plperl"
	description = "PL/Perl procedural language"
		locations = [{
			database = "app"
			schema   = "public"
			version  = "1.0"
			enabled  = false
			},
			{
			database = "postgres"
			schema   = "public"
			version  = "1.0"
			enabled  = true
		}]
	}]
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

func testTemboInstanceClientImportStateIDFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return "", errors.New("No Tembo Instance ID set")
		}

		orgId := rs.Primary.Attributes["org_id"]
		instanceId := rs.Primary.Attributes["instance_id"]

		return fmt.Sprintf("%s/%s", orgId, instanceId), nil
	}
}
