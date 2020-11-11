package pixela

import (
"context"
"encoding/json"
"fmt"
"net/http"
"strconv"
"time"

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
				Type:             schema.TypeList,
				Computed:         true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:             schema.TypeString,
							Computed:         true,
						},
						"name": &schema.Schema{
							Type:             schema.TypeString,
							Computed:         true,
						},
						"unit": &schema.Schema{
							Type:             schema.TypeString,
							Computed:         true,
						},
						"type": &schema.Schema{
							Type:             schema.TypeString,
							Computed:         true,
						},
						"color": &schema.Schema{
							Type:             schema.TypeString,
							Computed:         true,
						},
						"timezone": &schema.Schema{
							Type:             schema.TypeString,
							Computed:         true,
						},
						"purgeCacheURLs": &schema.Schema{
							Type:             schema.TypeList,
							Computed:         true,
							Elem: &schema.Schema{
								Type:             schema.TypeString,
								Computed:         true,
							},
						},
						"selfSufficient": &schema.Schema{
							Type:             schema.TypeString,
							Computed:         true,
						},
						"isSecret": &schema.Schema{
							Type:             schema.TypeBool,
							Computed:         true,
						},
						"publishOptionalData": &schema.Schema{
							Type:             schema.TypeBool,
							Computed:         true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGraphsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/coffees", "http://localhost:19090"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	coffees := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&coffees)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("coffees", coffees); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}