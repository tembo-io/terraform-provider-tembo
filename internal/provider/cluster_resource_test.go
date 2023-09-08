package provider

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestTemboClusterResource(t *testing.T) {

	clusterName := generateClusterName()
	organizationId := os.Getenv("ORGANIZATION_ID")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testProviderConfig() + testClusterResourceCreateConfig(clusterName, organizationId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tembo_cluster.test", "cluster_name", clusterName),
					resource.TestCheckResourceAttr("tembo_cluster.test", "organization_id", organizationId),
					resource.TestCheckResourceAttr("tembo_cluster.test", "cpu", "1"),
					resource.TestCheckResourceAttr("tembo_cluster.test", "stack", "Standard"),
					resource.TestCheckResourceAttr("tembo_cluster.test", "environment", "dev"),
					resource.TestCheckResourceAttr("tembo_cluster.test", "memory", "4Gi"),
					resource.TestCheckResourceAttr("tembo_cluster.test", "storage", "10Gi"),
					resource.TestCheckResourceAttrSet("tembo_cluster.test", "cluster_id"),
					resource.TestCheckResourceAttrSet("tembo_cluster.test", "last_updated"),
				),
			},
			// TODO: Add ImportState testing
			// Update and Read testing
			{
				Config: testProviderConfig() + testClusterResourceUpdateConfig(clusterName, organizationId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tembo_cluster.test", "cluster_name", clusterName),
					resource.TestCheckResourceAttr("tembo_cluster.test", "organization_id", organizationId),
					resource.TestCheckResourceAttr("tembo_cluster.test", "cpu", "2"),
					resource.TestCheckResourceAttr("tembo_cluster.test", "stack", "Standard"),
					resource.TestCheckResourceAttr("tembo_cluster.test", "environment", "dev"),
					resource.TestCheckResourceAttr("tembo_cluster.test", "memory", "8Gi"),
					resource.TestCheckResourceAttr("tembo_cluster.test", "storage", "10Gi"),
					resource.TestCheckResourceAttrSet("tembo_cluster.test", "cluster_id"),
					resource.TestCheckResourceAttrSet("tembo_cluster.test", "last_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testClusterResourceCreateConfig(clusterName string, organizationId string) string {
	return fmt.Sprintf(`
resource "tembo_cluster" "test" {
	cluster_name    = "%v"
	organization_id = "%v"
	cpu             = "1"
	stack           = "Standard"
	environment     = "dev"
	memory          = "4Gi"
	storage         = "10Gi"
  }
	`, clusterName, organizationId)
}

func testClusterResourceUpdateConfig(clusterName string, organizationId string) string {
	return fmt.Sprintf(`
resource "tembo_cluster" "test" {
	cluster_name    = "%v"
	organization_id = "%v"
	cpu             = "2"
	stack           = "Standard"
	environment     = "dev"
	memory          = "8Gi"
	storage         = "10Gi"
  }
	`, clusterName, organizationId)
}

func generateClusterName() string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return fmt.Sprintf("tf-test-%v", string(b))
}
