package datadog

import (
	"github.com/hashicorp/terraform/terraform"
	datadog "github.com/zorkian/go-datadog-api"
)

func testAccCheckDatadogIntegrationPagerdutyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*datadog.Client)

	return client.DeleteIntegrationPD()
}
