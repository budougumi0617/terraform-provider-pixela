package pixela

import (
	"context"
	"strconv"
	"time"

	pixela "github.com/ebc-2in2crc/pixela4go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGraphs() *schema.Resource {
	/**
	* Example JSON
	 {
	   "graphs": [
	     {
	       "id": "egiu",
	       "name": "English grammar in use",
	       "unit": "unit",
	       "type": "int",
	       "color": "ajisai",
	       "timezone": "Asia/Tokyo",
	       "purgeCacheURLs": null,
	       "selfSufficient": "none",
	       "isSecret": true,
	       "publishOptionalData": false
	     }
	   ]
	 }
	*/
	return &schema.Resource{
		ReadContext: dataSourceGraphsRead,
		Schema: map[string]*schema.Schema{
			"graphs": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"unit": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"color": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"timezone": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"purge_cache_urls": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"self_sufficient": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_secret": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"publish_optional_data": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGraphsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*pixela.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	result, err := client.Graph().GetAll()
	if err != nil {
		return diag.FromErr(err)
		})
	}

	if err := d.Set("graphs", result.Graphs); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
