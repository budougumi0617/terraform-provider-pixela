package pixela

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/budougumi0617/pixela"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform/helper/acctest"
)

func TestAccPixelaGraph_basic(t *testing.T) {
	var graph pixela.GraphDefinition

	rName := "tf" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckPixelaGraphDestroy,
		Steps: []resource.TestStep{
			// create test
			{
				// use a dynamic configuration with the random name from above
				Config: testAccPixelaResource(rName, "ajisai"),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the widget object
					testAccCheckPixelaGraphExists("pixela_graph.basic", &graph),
					// verify remote values
					testAccCheckPixelaGraphValues(&graph, rName, "ajisai"),
					// verify local values
					resource.TestCheckResourceAttr("pixela_graph.basic", "color", "ajisai"),
					resource.TestCheckResourceAttr("pixela_graph.basic", "id", rName),
				),
			},
			// update test
			{
				Config: testAccPixelaResource(rName, "momiji"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPixelaGraphExists("pixela_graph.basic", &graph),
					testAccCheckPixelaGraphValues(&graph, rName, "momiji"),
					resource.TestCheckResourceAttr("pixela_graph.basic", "color", "momiji"),
					resource.TestCheckResourceAttr("pixela_graph.basic", "id", rName),
				),
			},
		},
	})
}

func testAccCheckPixelaGraphValues(graph *pixela.GraphDefinition, name, color string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if string(graph.ID) != name {
			return fmt.Errorf("bad active state, expected \"true\", got: %#v", graph.ID)
		}
		if string(graph.Color) != color {
			return fmt.Errorf("bad name, expected \"%s\", got: %#v", name, graph.Color)
		}
		return nil
	}
}

// testAccCheckPixelaGraphExists queries the API and retrieves the matching Widget.
func testAccCheckPixelaGraphExists(n string, graph *pixela.GraphDefinition) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := testAccProvider.Meta().(*pixela.Client)
		resp, err := conn.GetGraph(context.Background(), pixela.GraphID(rs.Primary.ID))

		if err != nil {
			return err
		}

		// If no error, assign the response Graph attribute to the graph pointer
		*graph = *resp

		if graph == nil {
			return fmt.Errorf("Graph (%s) not found", rs.Primary.ID)
		}

		return nil
	}
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("PIXELA_TOKEN") == "" {
		t.Fatal("PIXELA_TOKEN must be set for acceptance tests")
	}
}

// testAccExampleResource returns an configuration for an Example Widget with the provided name
func testAccPixelaResource(name, color string) string {
	return fmt.Sprintf(`
provider pixela {
  username = "acceptance-test"
}

resource "pixela_graph" "basic" {
  graph_id              = %q
  name                  = "update from terraform"
  unit                  = "page"
  type                  = "int"
  color                 = %q
  timezone              = "Asia/Tokyo"
  self_sufficient       = "none"
  is_secret             = true
  publish_optional_data = false
}`, name, color)
}

// testAccCheckPixelaGraphDestroy verifies the Widget has been destroyed
func testAccCheckPixelaGraphDestroy(s *terraform.State) error {
	cli := testAccProvider.Meta().(*pixela.Client)
	ctx := context.Background()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "pixela_graph" {
			continue
		}

		_, err := cli.GetGraph(ctx, pixela.GraphID(rs.Primary.ID))
		if err != nil {
			return fmt.Errorf("Graph (%s) still exists.", rs.Primary.ID)
		}

		return nil
	}

	return nil
}
