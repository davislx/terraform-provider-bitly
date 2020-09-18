package bitly

import (
	"context"
	"log"

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
			"link": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"long_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"title": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceBitlinkCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	bc := bitlinks.New(c)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	payload := bitlinks.Bitlink{
		LongURL: d.Get("long_url").(string),
		Domain:  d.Get("domain").(string),
		Title:   d.Get("title").(string),
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
	c := m.(*client.Client)
	bc := bitlinks.New(c)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	bitlinkID := d.Id()

	res, err := bc.RetrieveBitlink(bitlinkID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("long_url", res.LongURL)
	d.Set("link", res.Link)
	d.Set("domain", "bit.ly")
	d.Set("title", res.Title)
	return diags
}

func resourceBitlinkUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	bc := bitlinks.New(c)

	bitlinkID := d.Id()

	var bitlinksDetails bitlinks.BitlinkDetails

	// Note: non-paying customer of bitly cannot change the `long_url`
	if d.HasChange("long_url") {
		bitlinksDetails.LongURL = d.Get("long_url").(string)
	}
	if d.HasChange("title") {
		bitlinksDetails.Title = d.Get("title").(string)
	}

	ret, err := bc.UpdateBitlink(bitlinkID, &bitlinksDetails)
	log.Printf("Update output: %+v\n", ret)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceBitlinkRead(ctx, d, m)
}

func resourceBitlinkDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	bc := bitlinks.New(c)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	bitlinkID := d.Id()

	// Bitly only supports soft delete
	bc.UpdateBitlink(bitlinkID, &bitlinks.BitlinkDetails{Archived: true})

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
