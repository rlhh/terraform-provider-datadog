package datadog

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	datadog "github.com/zorkian/go-datadog-api"
)

func TestAccDatadogIntegrationPagerduty_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatadogIntegrationPagerdutyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDatadogIntegrationPagerdutyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatadogIntegrationPagerdutyExists("datadog_integration_pagerduty.foo"),
					resource.TestCheckResourceAttr(
						"datadog_integration_pagerduty.foo", "subdomain", "testdomain"),
					resource.TestCheckResourceAttr(
						"datadog_integration_pagerduty.foo", "api_token", "*****"),
					resource.TestCheckResourceAttr(
						"datadog_integration_pagerduty.foo", "service.0.service_name", "test_service"),
					resource.TestCheckResourceAttr(
						"datadog_integration_pagerduty.foo", "service.0.service_key", "*****"),
				),
			},
		},
	})
}

func testAccCheckDatadogIntegrationPagerdutyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*datadog.Client)
		if err := datadogIntegrationPagerdutyExistsHelper(s, client); err != nil {
			return err
		}
		return nil
	}
}

func datadogIntegrationPagerdutyExistsHelper(s *terraform.State, client *datadog.Client) error {
	if _, err := client.GetIntegrationPD(); err != nil {
		return fmt.Errorf("Received an error retrieving integration pagerduty %s", err)
	}
	return nil
}

func testAccCheckDatadogIntegrationPagerdutyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*datadog.Client)

	return client.DeleteIntegrationPD()
}

const testAccCheckDatadogIntegrationPagerdutyConfig = `
 resource "datadog_integration_pagerduty" "foo" {
   service
     {
         service_name = "test_service",
         service_key  = "*****",
     }

   subdomain = "testdomain"
   api_token = "*****"
 }
 `
