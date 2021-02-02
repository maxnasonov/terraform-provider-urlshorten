package urlshorten

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	service := d.Get("service").(string)
	password := d.Get("password").(string)

	var diags diag.Diagnostics

	if (service != "") && (password != "") {
		c, err := NewClient(service, password)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create Shorten client",
				Detail:   "Unable to authorize user for Shorten client",
			})
			return nil, diags
		}

		return c, diags
	}
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to create Shorten client",
		Detail:   "Required values are service and password",
	})
	return nil, diags
}
