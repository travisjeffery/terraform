package datadog

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	datadog "gopkg.in/zorkian/go-datadog-api.v2"
)

func TestAccDatadogMetricMetadata_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDatadogMetricMetadataConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatadogMetricMetadataExists("datadog_monitor.foo"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "name", "name for monitor foo"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "message", "some message Notify: @hipchat-channel"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "type", "metric alert"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "query", "avg(last_1h):avg:aws.ec2.cpu{environment:foo,host:foo} by {host} > 2"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "notify_no_data", "false"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "new_host_delay", "600"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "evaluation_delay", "700"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "renotify_interval", "60"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "thresholds.warning", "1.0"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "thresholds.critical", "2.0"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "require_full_window", "true"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "locked", "false"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "tags.0", "foo:bar"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "tags.1", "baz"),
				),
			},
		},
	})
}

func TestAccDatadogMetricMetadata_Updated(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDatadogMetricMetadataConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatadogMetricMetadataExists("datadog_monitor.foo"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "name", "name for monitor foo"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "message", "some message Notify: @hipchat-channel"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "escalation_message", "the situation has escalated @pagerduty"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "query", "avg(last_1h):avg:aws.ec2.cpu{environment:foo,host:foo} by {host} > 2"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "type", "metric alert"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "notify_no_data", "false"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "new_host_delay", "600"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "evaluation_delay", "700"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "renotify_interval", "60"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "thresholds.warning", "1.0"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "thresholds.critical", "2.0"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "notify_audit", "false"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "timeout_h", "60"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "include_tags", "true"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "require_full_window", "true"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "locked", "false"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "tags.0", "foo:bar"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "tags.1", "baz"),
				),
			},
			{
				Config: testAccCheckDatadogMetricMetadataConfigUpdated,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatadogMetricMetadataExists("datadog_monitor.foo"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "name", "name for monitor bar"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "message", "a different message Notify: @hipchat-channel"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "query", "avg(last_1h):avg:aws.ec2.cpu{environment:bar,host:bar} by {host} > 3"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "escalation_message", "the situation has escalated! @pagerduty"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "type", "metric alert"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "notify_no_data", "true"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "new_host_delay", "900"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "evaluation_delay", "800"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "no_data_timeframe", "20"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "renotify_interval", "40"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "thresholds.ok", "0.0"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "thresholds.warning", "1.0"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "thresholds.critical", "3.0"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "notify_audit", "true"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "timeout_h", "70"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "include_tags", "false"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "silenced.*", "0"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "require_full_window", "false"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "locked", "true"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "tags.0", "baz:qux"),
					resource.TestCheckResourceAttr(
						"datadog_metricMetadata.foo", "tags.1", "quux"),
				),
			},
		},
	})
}

func testAccCheckDatadogMetricMetadataExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*datadog.Client)
		if err := existsHelper(s, client); err != nil {
			return err
		}
		return nil
	}
}

const testAccCheckDatadogMetricMetadataConfig = `
resource "datadog_metric_metadata" "foo" {
  short_name = "short name for metric_metadata foo"
  type = "gauge"
  description = "some description"
  unit = "operation"
  per_unit = "second"
  statsd_interval = "1s"
}
`

const testAccCheckDatadogMetricMetadataConfigUpdated = `
resource "datadog_metric_Metadata" "foo" {
  short_name = "a different short name for metric_metadata foo"
  type = "gauge"
  description = "some description"
  unit = "operation"
  per_unit = "second"
  statsd_interval = "1s"
}
`
