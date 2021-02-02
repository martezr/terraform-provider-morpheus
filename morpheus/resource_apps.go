package morpheus

import (
	"errors"
	"fmt"
	"github.com/gomorpheus/morpheus-go-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	//"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
)

func resourceApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppCreate,
		Read:   resourceAppRead,
		Update: resourceAppUpdate,
		Delete: resourceAppDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group": {
				Type:     schema.TypeMap,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"blueprintid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"environment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceAppCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*morpheus.Client)
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	environment := d.Get("environment").(string)
	blueprintid := d.Get("blueprintid").(string)

	var group map[string]interface{}
	if d.Get("group") != nil {
		group = d.Get("group").(map[string]interface{})
	}

	req := &morpheus.Request{
		Body: map[string]interface{}{
			"app": map[string]interface{}{
				"name":        name,
				"description": description,
				"environment": environment,
				"group":       group,
				"blueprintId": blueprintid,
			},
		},
	}
	log.Printf("######API BODY######: ", req.Body)
	resp, err := client.CreateApp(req)
	if err != nil {
		log.Printf("API FAILURE:", resp, err)
		return err
	}
	log.Printf("API RESPONSE: ", resp)

	result := resp.Result.(*morpheus.CreateAppResult)
	app := result.App
	// Successfully created resource, now set id
	d.SetId(int64ToString(app.ID))

	return resourceAppRead(d, meta)
}

func resourceAppRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*morpheus.Client)
	id := d.Id()
	name := d.Get("name").(string)

	// lookup by name if we do not have an id yet
	var resp *morpheus.Response
	var err error
	if id == "" && name != "" {
		resp, err = client.FindAppByName(name)
	} else if id != "" {
		resp, err = client.GetApp(toInt64(id), &morpheus.Request{})
	} else {
		return errors.New("App cannot be read without name or id")
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
	result := resp.Result.(*morpheus.GetAppResult)
	app := result.App
	if app != nil {
		d.SetId(int64ToString(app.ID))
		d.Set("name", app.Name)
		d.Set("description", app.Description)
		d.Set("environment", app.Environment)
		d.Set("group", app.Group)
		d.Set("blueprintid", app.BlueprintID)
		// todo: more fields
	} else {
		log.Println(app)
		return fmt.Errorf("read operation: app not found in response data") // should not happen
	}

	return nil
}

func resourceAppUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*morpheus.Client)
	id := d.Id()
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	environment := d.Get("environment").(string)
	group := d.Get("group")
	blueprintid := d.Get("blueprintid").(string)

	req := &morpheus.Request{
		Body: map[string]interface{}{
			"app": map[string]interface{}{
				"name":        name,
				"description": description,
				"environment": environment,
				"group":       group,
				"blueprintid": blueprintid,
			},
		},
	}
	resp, err := client.UpdateApp(toInt64(id), req)
	if err != nil {
		log.Printf("API FAILURE:", resp, err)
		return err
	}
	log.Printf("API RESPONSE: ", resp)
	result := resp.Result.(*morpheus.UpdateAppResult)
	account := result.App
	// Successfully updated resource, now set id
	// err, it should not have changed though..
	d.SetId(int64ToString(account.ID))
	return resourceAppRead(d, meta)
}

func resourceAppDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*morpheus.Client)
	id := d.Id()
	req := &morpheus.Request{}
	resp, err := client.DeleteApp(toInt64(id), req)
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
