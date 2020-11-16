package pixela

import (
	"context"
	"fmt"

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
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"graphs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"unit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"color": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"timezone": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"purge_cache_urls": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"self_sufficient": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_secret": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"publish_optional_data": {
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
	}

	graphs := flattenGraphsData(&result.Graphs)
	if err := d.Set("graphs", graphs); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(fmt.Sprintf("%s/graphs", client.UserName))

	return diags
}

// to manual mapping because terraform cannot use camel case fields.
// Pixela responses has it.
func flattenGraphsData(graphs *[]pixela.GraphDefinition) []interface{} {
	if graphs != nil {
		gs := make([]interface{}, len(*graphs), len(*graphs))

		for i, graph := range *graphs {
			g := make(map[string]interface{})
			g["id"] = graph.ID
			g["name"] = graph.Name
			g["unit"] = graph.Unit
			g["type"] = graph.Type
			g["color"] = graph.Color
			g["timezone"] = graph.TimeZone
			// below fields is used camel case in json tags.
			g["purge_cache_urls"] = graph.PurgeCacheURLs
			g["self_sufficient"] = graph.SelfSufficient
			g["is_secret"] = graph.IsSecret
			g["publish_optional_data"] = graph.PublishOptionalData

			gs[i] = g
		}

		return gs
	}

	return make([]interface{}, 0)
}
