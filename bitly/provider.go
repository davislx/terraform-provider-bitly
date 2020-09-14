package bitly

import (
  "context"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

  "github.com/retgits/bitly/client"
)

// Provider -
func Provider() *schema.Provider {
  return &schema.Provider{
    Schema: map[string]*schema.Schema{
			"bitly_token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BITLY_TOKEN", nil),
      },
    },
    ResourcesMap: map[string]*schema.Resource{
      "bitly_bitlink": resourceBitlink(),
    },
    DataSourcesMap: map[string]*schema.Resource{},
    ConfigureContextFunc: providerConfigure,
  }
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	bitlyToken := d.Get("bitly_token").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics


	if bitlyToken == "" {
    diags = append(diags, diag.Diagnostic{
      Severity: diag.Error,
      Summary:  "Unable to create Bitly client",
      Detail:   "Access Token for Bitly client not found",
    })

    return nil, diags
  } else {
		c := client.NewClient().WithAccessToken(bitlyToken)
		return c, diags
	}
}