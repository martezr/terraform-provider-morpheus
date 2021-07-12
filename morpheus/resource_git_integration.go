package morpheus

import (
	"context"
	"fmt"
	"log"

	"github.com/gomorpheus/morpheus-go-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGitIntegration() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides a Morpheus git integration resource",
		CreateContext: resourceGitIntegrationCreate,
		ReadContext:   resourceGitIntegrationRead,
		UpdateContext: resourceGitIntegrationUpdate,
		DeleteContext: resourceGitIntegrationDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Description: "The ID of the integration",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the integration",
				Required:    true,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Description: "Whether the integration is enabled or not",
				Optional:    true,
				Default:     true,
			},
			"url": {
				Type:        schema.TypeString,
				Description: "The description of the integration",
				Required:    true,
			},
			"default_branch": {
				Type:        schema.TypeBool,
				Description: "Whether the integration is enabled or not",
				Optional:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "Whether the integration is enabled or not",
				Optional:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Whether the integration is enabled or not",
				Optional:    true,
			},
			"access_token": {
				Type:        schema.TypeString,
				Description: "The access token used to access the git repository",
				Optional:    true,
			},
			"key_pair_id": {
				Type:        schema.TypeInt,
				Description: "The id of the key pair used to access the git repository",
				Optional:    true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceGitIntegrationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*morpheus.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Get("name").(string)
	req := &morpheus.Request{
		Body: map[string]interface{}{
			"integration": map[string]interface{}{
				"name":            name,
				"type":            "git",
				"service_url":     "",
				"serviceUsername": "",
				"servicePassword": "",
			},
		},
	}

	resp, err := client.CreateIntegration(req)
	if err != nil {
		log.Printf("API FAILURE: %s - %s", resp, err)
		return diag.FromErr(err)
	}

	log.Printf("API RESPONSE: %s", resp)

	result := resp.Result.(*morpheus.CreateIntegrationResult)
	integration := result.Integration
	// Successfully created resource, now set id
	d.SetId(int64ToString(integration.ID))

	return diags
}

func resourceGitIntegrationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*morpheus.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	name := d.Get("name").(string)

	// lookup by name if we do not have an id yet
	var resp *morpheus.Response
	var err error
	if id == "" && name != "" {
		resp, err = client.FindIntegrationByName(name)
	} else if id != "" {
		resp, err = client.GetIntegration(toInt64(id), &morpheus.Request{})
	} else {
		return diag.Errorf("Integration cannot be read without name or id")
	}

	if err != nil {
		// 404 is ok?
		if resp != nil && resp.StatusCode == 404 {
			log.Printf("API 404: %s - %s", resp, err)
			return nil
		} else {
			log.Printf("API FAILURE: %s - %s", resp, err)
			return diag.FromErr(err)
		}
	}
	log.Printf("API RESPONSE: %s", resp)

	// store resource data
	result := resp.Result.(*morpheus.GetIntegrationResult)
	integration := result.Integration
	if integration != nil {
		d.SetId(int64ToString(integration.ID))
		d.Set("name", integration.Name)
		d.Set("enabled", integration.Enabled)

	} else {
		log.Println(integration)
		err := fmt.Errorf("read operation: integration not found in response data") // should not happen
		return diag.FromErr(err)
	}

	return diags
}

func resourceGitIntegrationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*morpheus.Client)

	id := d.Id()
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	code := d.Get("code").(string)

	req := &morpheus.Request{
		Body: map[string]interface{}{
			"integration": map[string]interface{}{
				"active":      d.Get("active").(bool),
				"name":        name,
				"description": description,
				"code":        code,
				"visibility":  d.Get("visibility").(string),
			},
		},
	}
	resp, err := client.UpdateIntegration(toInt64(id), req)
	if err != nil {
		log.Printf("API FAILURE: %s - %s", resp, err)
		return diag.FromErr(err)
	}
	log.Printf("API RESPONSE: %s", resp)
	result := resp.Result.(*morpheus.UpdateIntegrationResult)
	integration := result.Integration
	// Successfully updated resource, now set id
	// err, it should not have changed though..
	d.SetId(int64ToString(integration.ID))
	return resourceGitIntegrationRead(ctx, d, meta)
}

func resourceGitIntegrationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*morpheus.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	req := &morpheus.Request{}
	resp, err := client.DeleteIntegration(toInt64(id), req)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Printf("API 404: %s - %s", resp, err)
			return nil
		} else {
			log.Printf("API FAILURE: %s - %s", resp, err)
			return diag.FromErr(err)
		}
	}
	log.Printf("API RESPONSE: %s", resp)
	d.SetId("")
	return diags
}

type GitIntegration struct {
	Success     bool `json:"success"`
	Integration struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		Enabled         bool   `json:"enabled"`
		Type            string `json:"type"`
		Integrationtype struct {
			ID   int    `json:"id"`
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"integrationType"`
		URL      string `json:"url"`
		Isplugin bool   `json:"isPlugin"`
		Config   struct {
			Defaultbranch string `json:"defaultBranch"`
		} `json:"config"`
		Status           string      `json:"status"`
		Statusdate       interface{} `json:"statusDate"`
		Statusmessage    interface{} `json:"statusMessage"`
		Lastsync         interface{} `json:"lastSync"`
		Lastsyncduration interface{} `json:"lastSyncDuration"`
	} `json:"integration"`
}
