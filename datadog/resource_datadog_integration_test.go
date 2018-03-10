package datadog

import (
	"github.com/hashicorp/terraform/terraform"
	datadog "github.com/zorkian/go-datadog-api"
)

func testAccCheckDatadogIntegrationPagerdutyDestroy(s *terraform.State) error {
	_ = testAccProvider.Meta().(*datadog.Client)

	/*
		if err := datadogUserDestroyHelper(s, client); err != nil {
			return err
		}
	*/
	return nil
}
