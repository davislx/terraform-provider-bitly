package bitly

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/retgits/bitly/client"
	"github.com/retgits/bitly/client/bitlinks"
)

func resourceBitlink() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBitlinkCreate,
		ReadContext:   resourceBitlinkRead,
		UpdateContext: resourceBitlinkUpdate,
		DeleteContext: resourceBitlinkDelete,
		Schema: map[string]*schema.Schema{
			"link": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"long_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_guid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceBitlinkCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	bc := bitlinks.New(c)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	longURL := d.Get("long_url").(string)
	payload := bitlinks.Bitlink {
		LongURL: longURL,
	}

	res, err := bc.CreateBitlink(&payload)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(res.ID)

	resourceBitlinkRead(ctx, d, m)

	return diags
}

func resourceBitlinkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// c := m.(*client.Client)
	// bc := bitlinks.New(c)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// bitlinkID := d.Id()

	// res, err := bc.RetrieveBitlink()
	// if err != nil {
	// 	return diag.FromErr(err)
	// }

	return diags
}

func resourceBitlinkUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// c := m.(*client.Client)
	// bc := bitlinks.New(c)

	// bitlinkID := d.Id()

	return resourceBitlinkRead(ctx, d, m)
}

func resourceBitlinkDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// c := m.(*client.Client)
	// bc := bitlinks.New(c)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// bitlinkID := d.Id()


	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
