package pixela

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var testAccProvider *schema.Provider
var testAccProviderFactories map[string]func() (*schema.Provider, error)

func init() {
	testAccProvider = Provider()
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"pixela": func() (*schema.Provider, error) {
			return testAccProvider, nil
		},
	}
}
