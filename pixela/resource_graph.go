package pixela

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGraph() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGraphCreate,
		ReadContext:   resourceGraphRead,
		UpdateContext: resourceGraphUpdate,
		DeleteContext: resourceGraphDelete,
		/*
		   {
		     "id": "test-graph",
		     "name": "graph-name",
		     "unit": "commit",
		     "type": "int",
		     "color": "shibafu",
		     "timezone": "Asia/Tokyo",
		     "purgeCacheURLs": [
		       "https://camo.githubusercontent.com/xxx/xxxx"
		     ],
		     "selfSufficient": "increment",
		     "isSecret": false,
		     "publishOptionalData": true
		   }
		*/
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"unit": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"color": {
				Type:     schema.TypeString,
				Required: true,
			},
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "UTC",
			},
			"purge_cache_urls": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"self_sufficient": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "none",
			},
			"is_secret": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"publish_optional_data": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceGraphCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceGraphRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceGraphUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceGraphRead(ctx, d, m)
}

func resourceGraphDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}
