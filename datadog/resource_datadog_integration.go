package datadog

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/zorkian/go-datadog-api"
)

func resourceDatadogIntegrationPagerduty() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatadogIntegrationPagerdutyCreate,
		Read:   resourceDatadogIntegrationPagerdutyRead,
		Update: resourceDatadogIntegrationPagerdutyUpdate,
		Delete: resourceDatadogIntegrationPagerdutyDelete,
		Exists: resourceDatadogIntegrationPagerdutyExists,
		Importer: &schema.ResourceImporter{
			State: resourceDatadogIntegrationPagerdutyImport,
		},

		Schema: map[string]*schema.Schema{
			"service": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "A list of service names and service keys.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"service_key": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"subdomain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"schedules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"api_token": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDatadogIntegrationPagerdutyExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	// Exists - This is called to verify a resource still exists. It is called prior to Read,
	// and lowers the burden of Read to be able to assume the resource exists.
	client := meta.(*datadog.Client)

	if _, err := client.GetIntegrationPD(); err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func resourceDatadogIntegrationPagerdutyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*datadog.Client)

	pd := &datadog.IntegrationPDRequest{}
	pd.SetSubdomain(d.Get("subdomain").(string))
	pd.SetAPIToken(d.Get("api_token").(string))

	schedules := []string{}
	for _, s := range d.Get("schedules").([]interface{}) {
		schedules = append(schedules, s.(string))
	}
	pd.Schedules = schedules

	services := []datadog.ServicePDRequest{}
	for _, sInterface := range d.Get("service").([]interface{}) {
		s := sInterface.(map[string]interface{})

		service := datadog.ServicePDRequest{}
		service.SetServiceName(s["service_name"].(string))
		service.SetServiceKey(s["service_key"].(string))

		services = append(services, service)
	}
	pd.Services = services

	if err := client.CreateIntegrationPD(pd); err != nil {
		return fmt.Errorf("Failed to create integration pagerduty using Datadog API: %s", err.Error())
	}

	pdIntegration, err := client.GetIntegrationPD()
	if err != nil {
		return fmt.Errorf("error retrieving integration pagerduty: %s", err.Error())
	}

	d.SetId(pdIntegration.GetSubdomain())

	return nil
}

func resourceDatadogIntegrationPagerdutyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*datadog.Client)

	pd, err := client.GetIntegrationPD()
	if err != nil {
		return err
	}

	services := []map[string]string{}
	for _, service := range pd.Services {

		serviceName, ok := datadog.GetStringOk(service.ServiceName)
		if !ok {
			return fmt.Errorf("Error")
		}

		serviceKey, ok := datadog.GetStringOk(service.ServiceKey)
		if !ok {
			return fmt.Errorf("Error")
		}

		services = append(services, map[string]string{
			"service_name": serviceName,
			"service_key":  serviceKey,
		})
	}

	d.Set("service", services)
	d.Set("subdomain", pd.GetSubdomain())
	d.Set("schedules", pd.Schedules)
	d.Set("api_token", pd.GetAPIToken())

	return nil
}

func resourceDatadogIntegrationPagerdutyUpdate(d *schema.ResourceData, meta interface{}) error {

	/*
		var u datadog.User
		u.SetDisabled(d.Get("disabled").(bool))
		u.SetEmail(d.Get("email").(string))
		u.SetHandle(d.Id())
		u.SetIsAdmin(d.Get("is_admin").(bool))
		u.SetName(d.Get("name").(string))
		u.SetRole(d.Get("role").(string))

		if err := client.UpdateUser(u); err != nil {
			return fmt.Errorf("error updating user: %s", err.Error())
		}

		return resourceDatadogIntegrationPagerdutyRead(d, meta)
	*/
	return nil
}

func resourceDatadogIntegrationPagerdutyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*datadog.Client)

	return client.DeleteIntegrationPD()
}

func resourceDatadogIntegrationPagerdutyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := resourceDatadogIntegrationPagerdutyRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
