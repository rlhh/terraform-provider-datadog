package datadog

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestDatadogIntegrationPagerduty_import(t *testing.T) {
	resourceName := "datadog_integration_pagerduty.pd"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatadogIntegrationPagerdutyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckDatadogIntegrationPagerdutyConfigImported,
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccCheckDatadogIntegrationPagerdutyConfigImported = `
resource "datadog_integration_pagerduty" "pd" {
  service 
	{
		service_name = "test_service",
		service_key  = "*****"
	}

  service
	{
		service_name = "test_service_2",
		service_key  = "*****",
	}
  
  subdomain = "testdomain"
  api_token = "*****"
}
`
