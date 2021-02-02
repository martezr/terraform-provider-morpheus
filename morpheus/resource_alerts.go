package morpheus

import (
	"errors"
	"fmt"
	"github.com/gomorpheus/morpheus-go-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	//"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
)

func resourceAlert() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlertCreate,
		Read:   resourceAlertRead,
		Update: resourceAlertUpdate,
		Delete: resourceAlertDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceAlertCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*morpheus.Client)
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	code := d.Get("code").(string)
	req := &morpheus.Request{
		Body: map[string]interface{}{
			"alert": map[string]interface{}{
				"name":        name,
				"description": description,
				"code":        code,
			},
		},
	}
	resp, err := client.CreateAlert(req)
	if err != nil {
		log.Printf("API FAILURE:", resp, err)
		return err
	}
	log.Printf("API RESPONSE: ", resp)

	result := resp.Result.(*morpheus.CreateAlertResult)
	alert := result.Alert
	// Successfully created resource, now set id
	d.SetId(int64ToString(alert.ID))

	return resourceAlertRead(d, meta)
}

func resourceAlertRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*morpheus.Client)
	id := d.Id()
	name := d.Get("name").(string)

	// lookup by name if we do not have an id yet
	var resp *morpheus.Response
	var err error
	if id == "" && name != "" {
		resp, err = client.FindAlertByName(name)
	} else if id != "" {
		resp, err = client.GetAlert(toInt64(id), &morpheus.Request{})
	} else {
		return errors.New("Alert cannot be read without name or id")
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
	result := resp.Result.(*morpheus.GetAlertResult)
	alert := result.Alert
	if alert != nil {
		d.SetId(int64ToString(alert.ID))
		d.Set("name", alert.Name)
		d.Set("description", alert.Description)
		// todo: more fields
	} else {
		log.Println(alert)
		return fmt.Errorf("read operation: alert not found in response data") // should not happen
	}

	return nil
}

func resourceAlertUpdate(d *schema.ResourceData, meta interface{}) error {
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
	resp, err := client.UpdateAlert(toInt64(id), req)
	if err != nil {
		log.Printf("API FAILURE:", resp, err)
		return err
	}
	log.Printf("API RESPONSE: ", resp)
	result := resp.Result.(*morpheus.UpdateAlertResult)
	account := result.Alert
	// Successfully updated resource, now set id
	// err, it should not have changed though..
	d.SetId(int64ToString(account.ID))
	return resourceAlertRead(d, meta)
}

func resourceAlertDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*morpheus.Client)
	id := d.Id()
	req := &morpheus.Request{}
	resp, err := client.DeleteAlert(toInt64(id), req)
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
