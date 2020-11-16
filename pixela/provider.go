package pixela

import (
	"context"
	"fmt"

	pixela "github.com/ebc-2in2crc/pixela4go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	UserNameKey = "PIXELA_USERNAME"
	TokenKey    = "PIXELA_TOKEN"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PIXELA_USERNAME", nil),
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("PIXELA_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"pixela_graphs": dataSourceGraphs(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	un := d.Get("username").(string)
	if un == "" {
		return nil, diag.FromErr(fmt.Errorf("not find username"))
	}

	token := d.Get("token").(string)
	if token == "" {
		return nil, diag.FromErr(fmt.Errorf("not find token"))
	}

	return pixela.New(un, token), diag.Diagnostics{}

}
