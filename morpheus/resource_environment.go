package morpheus

import (
	"errors"
	"fmt"
	"github.com/gomorpheus/morpheus-go-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	//"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
)

func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnvironmentCreate,
		Read:   resourceEnvironmentRead,
		Update: resourceEnvironmentUpdate,
		Delete: resourceEnvironmentDelete,

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

func resourceEnvironmentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*morpheus.Client)
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	code := d.Get("code").(string)
	req := &morpheus.Request{
		Body: map[string]interface{}{
			"environment": map[string]interface{}{
				"name":        name,
				"description": description,
				"code":        code,
			},
		},
	}
	resp, err := client.CreateEnvironment(req)
	if err != nil {
		log.Printf("API FAILURE:", resp, err)
		return err
	}
	log.Printf("API RESPONSE: ", resp)

	result := resp.Result.(*morpheus.CreateEnvironmentResult)
	environment := result.Environment
	// Successfully created resource, now set id
	d.SetId(int64ToString(environment.ID))

	return resourceEnvironmentRead(d, meta)
}

func resourceEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*morpheus.Client)
	id := d.Id()
	name := d.Get("name").(string)

	// lookup by name if we do not have an id yet
	var resp *morpheus.Response
	var err error
	if id == "" && name != "" {
		resp, err = client.FindEnvironmentByName(name)
	} else if id != "" {
		resp, err = client.GetEnvironment(toInt64(id), &morpheus.Request{})
	} else {
		return errors.New("Environment cannot be read without name or id")
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
	result := resp.Result.(*morpheus.GetEnvironmentResult)
	environment := result.Environment
	if environment != nil {
		d.SetId(int64ToString(environment.ID))
		d.Set("name", environment.Name)
		d.Set("description", environment.Description)
		// todo: more fields
	} else {
		log.Println(environment)
		return fmt.Errorf("read operation: environment not found in response data") // should not happen
	}

	return nil
}

func resourceEnvironmentUpdate(d *schema.ResourceData, meta interface{}) error {
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
	resp, err := client.UpdateEnvironment(toInt64(id), req)
	if err != nil {
		log.Printf("API FAILURE:", resp, err)
		return err
	}
	log.Printf("API RESPONSE: ", resp)
	result := resp.Result.(*morpheus.UpdateEnvironmentResult)
	account := result.Environment
	// Successfully updated resource, now set id
	// err, it should not have changed though..
	d.SetId(int64ToString(account.ID))
	return resourceEnvironmentRead(d, meta)
}

func resourceEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*morpheus.Client)
	id := d.Id()
	req := &morpheus.Request{}
	resp, err := client.DeleteEnvironment(toInt64(id), req)
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
