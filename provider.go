package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ALERTOPS_API_KEY", nil),
				Description: "AlertOps API Key",
				Sensitive:   true,
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ALERTOPS_BASE_URL", "https://api.alertops.com"),
				Description: "AlertOps API Base URL",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"alertops_user":                 resourceUser(),
			"alertops_group":                resourceGroup(),
			"alertops_schedule":             resourceSchedule(),
			"alertops_workflow":             resourceWorkflow(),
			"alertops_escalation_policy":    resourceEscalationPolicy(),
			"alertops_inbound_integration":  resourceInboundIntegration(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"alertops_user": dataSourceUser(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get("api_key").(string)
	baseURL := d.Get("base_url").(string)

	log.Printf("DEBUG: Provider received API key: '%s' (length: %d)", apiKey, len(apiKey))
	log.Printf("DEBUG: Provider received base URL: '%s'", baseURL)
	
	// Add obvious error if API key is empty
	if apiKey == "" {
		log.Printf("ERROR: API key is EMPTY!")
	}

	var diags diag.Diagnostics

	if apiKey == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create AlertOps client",
			Detail:   "API key is required",
		})
		return nil, diags
	}



	client := NewClient(apiKey, baseURL)
	
	return client, diags
} 