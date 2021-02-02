package urlshorten

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceShorthenURL() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceShorthenURLRead,
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"new_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceShorthenURLRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Service)

	var diags diag.Diagnostics

	url := d.Get("url").(string)

	response, err := client.ShortenURL(url)
	if err != nil {
		return diag.FromErr(err)
	}
	if response == nil {
		return diag.FromErr(fmt.Errorf("Empty response"))
	}

	if err := d.Set("new_url", *response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
