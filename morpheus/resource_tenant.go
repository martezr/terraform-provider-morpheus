package morpheus

import (
	"errors"
	"fmt"
	"github.com/gomorpheus/morpheus-go-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	//"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
)

func resourceTenant() *schema.Resource {
	return &schema.Resource{
		Create: resourceTenantCreate,
		Read:   resourceTenantRead,
		Update: resourceTenantUpdate,
		Delete: resourceTenantDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subdomain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceTenantCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*morpheus.Client)
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	subdomain := d.Get("subdomain").(string)
	req := &morpheus.Request{
		Body: map[string]interface{}{
			"account": map[string]interface{}{
				"name":        name,
				"description": description,
				"subdomain":   subdomain,
			},
		},
	}
	resp, err := client.CreateTenant(req)
	if err != nil {
		log.Printf("API FAILURE:", resp, err)
		return err
	}
	log.Printf("API RESPONSE: ", resp)

	result := resp.Result.(*morpheus.CreateTenantResult)
	tenant := result.Tenant
	// Successfully created resource, now set id
	d.SetId(int64ToString(tenant.ID))

	return resourceTenantRead(d, meta)
}

func resourceTenantRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*morpheus.Client)
	id := d.Id()
	name := d.Get("name").(string)

	// lookup by name if we do not have an id yet
	var resp *morpheus.Response
	var err error
	if id == "" && name != "" {
		resp, err = client.FindTenantByName(name)
	} else if id != "" {
		resp, err = client.GetTenant(toInt64(id), &morpheus.Request{})
	} else {
		return errors.New("Tenant cannot be read without name or id")
	}

	if err != nil {
		// 404 is ok?
		if resp != nil && resp.StatusCode == 404 {
			log.Printf("API 404:", resp, err)
			return nil
		} else {
			log.Printf("API FAILURE:", resp, err)
			return err
		}
	}
	log.Printf("API RESPONSE:", resp)

	// store resource data
	result := resp.Result.(*morpheus.GetTenantResult)
	tenant := result.Tenant
	if tenant != nil {
		d.SetId(int64ToString(tenant.ID))
		d.Set("name", tenant.Name)
		d.Set("description", tenant.Description)
		// todo: more fields
	} else {
		log.Println(tenant)
		return fmt.Errorf("read operation: tenant not found in response data") // should not happen
	}

	return nil
}

func resourceTenantUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*morpheus.Client)
	id := d.Id()
	name := d.Get("name").(string)
	description := d.Get("description").(string)

	req := &morpheus.Request{
		Body: map[string]interface{}{
			"accouunt": map[string]interface{}{
				"name":        name,
				"description": description,
			},
		},
	}
	resp, err := client.UpdateTenant(toInt64(id), req)
	if err != nil {
		log.Printf("API FAILURE:", resp, err)
		return err
	}
	log.Printf("API RESPONSE: ", resp)
	result := resp.Result.(*morpheus.UpdateTenantResult)
	account := result.Tenant
	// Successfully updated resource, now set id
	// err, it should not have changed though..
	d.SetId(int64ToString(account.ID))
	return resourceTenantRead(d, meta)
}

func resourceTenantDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*morpheus.Client)
	id := d.Id()
	req := &morpheus.Request{}
	resp, err := client.DeleteTenant(toInt64(id), req)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Printf("API 404:", resp, err)
			return nil
		} else {
			log.Printf("API FAILURE:", resp, err)
			return err
		}
	}
	log.Printf("API RESPONSE:", resp)
	//d.setId("") // implicit
	return nil
}
