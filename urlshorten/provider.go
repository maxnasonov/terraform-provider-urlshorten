package urlshorten

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider definition of schema(configuration), resources(CRUD) operations and dataSources(query)
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"service": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SHORTEN_SERVICE", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SHORTEN_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"shorten_url": dataSourceShorthenURL(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
