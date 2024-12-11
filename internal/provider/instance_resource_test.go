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
					resource.TestCheckResourceAttr(resourceName, "postgres_configs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "connection_pooler.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "connection_pooler.pooler.pool_mode", "transaction"),
					resource.TestCheckResourceAttr(resourceName, "connection_pooler.pooler.parameters.max_client_conn", "20"),
					resource.TestCheckResourceAttr(resourceName, "connection_pooler.pooler.parameters.default_pool_size", "100"),
				),
			},
			// Wait for 10 minutes to ensure that the initial backup has gone through
			{
				Config: testProviderConfig() + testInstanceResourceCreateConfig(instanceName, orgId),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						fmt.Println("Waiting for 10 minutes before checking first_recoverability_time...")
						time.Sleep(10 * time.Minute)
						return nil
					},
				),
			},
			// Check that first_recoverability_time is set
			{
				Config: testProviderConfig() + testInstanceResourceCreateConfig(instanceName, orgId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "first_recoverability_time"),
					func(s *terraform.State) error {
						fmt.Println("first_recoverability_time has been set successfully")
						return nil
					},
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
					resource.TestCheckResourceAttr(resourceName, "replicas", "1"),
					resource.TestCheckResourceAttr(resourceName, "environment", "dev"),
					resource.TestCheckResourceAttr(resourceName, "memory", "8Gi"),
					resource.TestCheckResourceAttr(resourceName, "storage", "10Gi"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(resourceName, "last_updated"),
					resource.TestCheckResourceAttr(resourceName, "postgres_configs.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "connection_pooler.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "connection_pooler.pooler.pool_mode", "transaction"),
					resource.TestCheckResourceAttr(resourceName, "connection_pooler.pooler.parameters.max_client_conn", "20"),
					resource.TestCheckResourceAttr(resourceName, "connection_pooler.pooler.parameters.default_pool_size", "100"),
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
	connection_pooler = {
		enabled = true,
		pooler = {
		  pool_mode = "transaction",
		  parameters = {
			max_client_conn   = "20"
			default_pool_size = "100"
		  }
		}
	  }
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
	replicas        = 1
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
		  value = "16MB"
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
	},
	{
		"name" : "pltclu",
		"description" : "PL/TclU untrusted procedural language",
		"locations" : [
		{
			"database" : "app",
			"schema" : "public",
			"version" : "1.0",
			"enabled" : false,
			"error" : false,
			"error_message" : null
		},
		{
			"database" : "postgres",
			"schema" : "public",
			"version" : "1.0",
			"enabled" : false,
			"error" : false,
			"error_message" : null
		}
		]
	}]
	connection_pooler = {
		enabled = true,
		pooler = {
		  pool_mode = "transaction",
		  parameters = {
			max_client_conn   = "20"
			default_pool_size = "100"
		  }
		}
	  }
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
