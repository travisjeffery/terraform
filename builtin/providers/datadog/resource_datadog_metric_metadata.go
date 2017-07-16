package datadog

import (
	"fmt"
	"strings"

	datadog "gopkg.in/zorkian/go-datadog-api.v2"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDatadogMetricMetadata() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatadogMetricMetadataCreate,
		Read:   resourceDatadogMetricMetadataRead,
		Update: resourceDatadogMetricMetadataUpdate,
		Delete: resourceDatadogMetricMetadataDelete,
		Exists: resourceDatadogMetricMetadataExists,
		Importer: &schema.ResourceImporter{
			State: resourceDatadogMetricMetadataImport,
		},

		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: false,
			},
			"description": {
				Type:     schema.TypeString,
				Required: false,
			},
			"short_name": {
				Type:     schema.TypeString,
				Optional: false,
			},
			"unit": {
				Type:     schema.TypeString,
				Required: false,
			},
			"per_unit": {
				Type:     schema.TypeString,
				Required: false,
			},
			"statsd_interval": {
				Type:     schema.TypeString,
				Required: false,
			},
		},
	}
}

func buildMetricMetadataStruct(d *schema.ResourceData) (string, *datadog.MetricMetadata) {
	return d.Id(), &datadog.MetricMetadata{
		Type:           datadog.String(d.Get("type").(string)),
		Description:    datadog.String(d.Get("description").(string)),
		ShortName:      datadog.String(d.Get("short_name").(string)),
		Unit:           datadog.String(d.Get("unit").(string)),
		PerUnit:        datadog.String(d.Get("per_unit").(string)),
		StatsdInterval: datadog.String(d.Get("statsd_interval").(string)),
	}
}

func resourceDatadogMetricMetadataExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	// Exists - This is called to verify a resource still exists. It is called prior to Read,
	// and lowers the burden of Read to be able to assume the resource exists.
	client := meta.(*datadog.Client)

	i, _ := buildMetricMetadataStruct(d)

	if _, err := client.ViewMetricMetadata(i); err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func resourceDatadogMetricMetadataCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*datadog.Client)

	i, m := buildMetricMetadataStruct(d)
	m, err := client.EditMetricMetadata(i, m)
	if err != nil {
		return fmt.Errorf("error updating MetricMetadata: %s", err.Error())
	}

	d.SetId(i)

	return nil
}

func resourceDatadogMetricMetadataRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*datadog.Client)

	i, _ := buildMetricMetadataStruct(d)

	m, err := client.ViewMetricMetadata(i)
	if err != nil {
		return err
	}

	d.Set("type", m.GetType())
	d.Set("description", m.GetDescription())
	d.Set("short_name", m.GetShortName())
	d.Set("unit", m.GetUnit())
	d.Set("per_unit", m.GetPerUnit())
	d.Set("statsd_interval", m.GetStatsdInterval())

	return nil
}

func resourceDatadogMetricMetadataUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*datadog.Client)

	m := &datadog.MetricMetadata{}
	i := d.Get("metric_name").(string)

	if attr, ok := d.GetOk("type"); ok {
		m.SetType(attr.(string))
	}
	if attr, ok := d.GetOk("description"); ok {
		m.SetDescription(attr.(string))
	}
	if attr, ok := d.GetOk("short_name"); ok {
		m.SetShortName(attr.(string))
	}
	if attr, ok := d.GetOk("unit"); ok {
		m.SetUnit(attr.(string))
	}
	if attr, ok := d.GetOk("per_unit"); ok {
		m.SetPerUnit(attr.(string))
	}
	if attr, ok := d.GetOk("statsd_interval"); ok {
		m.SetStatsdInterval(attr.(string))
	}

	if _, err := client.EditMetricMetadata(i, m); err != nil {
		return fmt.Errorf("error updating MetricMetadata: %s", err.Error())
	}

	return resourceDatadogMetricMetadataRead(d, meta)
}

func resourceDatadogMetricMetadataDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDatadogMetricMetadataImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := resourceDatadogMetricMetadataRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
