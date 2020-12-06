package pixela

import (
	"fmt"
	"os"
	"testing"

	pixela "github.com/ebc-2in2crc/pixela4go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPixelaGraph_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckPixelaGraphDestroy,
		Steps:             nil,
	})
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("PIXELA_TOKEN") == "" {
		t.Fatal("PIXELA_TOKEN must be set for acceptance tests")
	}
}

// testAccCheckPixelaGraphDestroy verifies the Widget has been destroyed
func testAccCheckPixelaGraphDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	cli := testAccProvider.Meta().(*pixela.Client)

	// loop through the resources in state, verifying each widget
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "pixela_graph" {
			continue
		}

		// Retrieve our widget by referencing it's state ID for API lookup
		request := &pixela.GraphGetInput{ID: pixela.String(rs.Primary.ID)}

		response, err := cli.Graph().Get(request)
		if err != nil {
			return err
		}

		if response.IsSuccess {
			return fmt.Errorf("Graph (%s) still exists.", rs.Primary.ID)
		}

		return nil
	}

	return nil
}
