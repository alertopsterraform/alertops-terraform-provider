package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInboundIntegration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInboundIntegrationCreate,
		ReadContext:   resourceInboundIntegrationRead,
		UpdateContext: resourceInboundIntegrationUpdate,
		DeleteContext: resourceInboundIntegrationDelete,

		Schema: map[string]*schema.Schema{
			"inbound_integration_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The unique identifier for the inbound integration",
			},
			"inbound_integration_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the inbound integration",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of inbound integration (e.g., API, Email, Chat, etc.)",
			},
			"sequence": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The sequence order of the integration",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether the integration is enabled",
			},
			"escalation_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The escalation policy for this integration",
			},
			"recipient_groups": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of recipient groups",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"recipient_users": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of recipient users",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"inbound_template_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The template ID for the inbound integration",
			},
			"mail_box": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mailbox for email integrations",
			},
			"bridge": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Bridge configuration for telephone integrations",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"telephone_number": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Telephone number for bridge",
						},
						"access_code": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Access code for bridge",
						},
					},
				},
			},
			"api_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "API settings configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_bidirection": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether the integration is bidirectional",
						},
						"url_mapping": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "URL mapping configuration",
							Elem:        getURLMappingSchema(),
						},
						"alert_tags": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Alert tags configuration",
							Elem:        getAlertTagsSchema(),
						},
						"delaying_or_grouping": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Delaying and grouping rules",
							Elem:        getDelayingOrGroupingSchema(),
						},
						"filters_to_match_json_or_form_fields": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Filters to match JSON or form fields",
							Elem:        getFiltersSchema(),
						},
						"escalation_policy_override": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Escalation policy override settings",
							Elem:        getEscalationPolicyOverrideSchema(),
						},
						"dynamic_recipient_groups": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Dynamic recipient groups",
							Elem:        getDynamicRecipientGroupSchema(),
						},
					},
				},
			},
			"email_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Email settings configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email_mapping": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Email mapping configuration",
							Elem:        getEmailMappingSchema(),
						},
						"alert_tags": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Email alert tags configuration",
							Elem:        getEmailAlertTagsSchema(),
						},
						"delaying_or_grouping": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Email delaying and grouping rules",
							Elem:        getDelayingOrGroupingSchema(),
						},
						"filters_to_match_incoming_emails": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Filters to match incoming emails",
							Elem:        getEmailFiltersSchema(),
						},
						"escalation_policy_override": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Email escalation policy override settings",
							Elem:        getEscalationPolicyOverrideSchema(),
						},
						"dynamic_recipient_groups": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Email dynamic recipient groups",
							Elem:        getDynamicRecipientGroupSchema(),
						},
					},
				},
			},
			"chat_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Chat settings configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url_mapping": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Chat URL mapping configuration",
							Elem:        getChatURLMappingSchema(),
						},
						"escalation_policy_override": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Chat escalation policy override settings",
							Elem:        getEscalationPolicyOverrideSchema(),
						},
					},
				},
			},
			"heartbeat_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Heartbeat settings configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"heartbeat_interval_in_min": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Heartbeat interval in minutes",
						},
					},
				},
			},
		},
	}
}

// Helper function to get URL mapping schema
func getURLMappingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "HTTP method",
			},
			"content": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Content mapping",
			},
			"source": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source field",
			},
			"source_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source name field",
			},
			"static": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the source is static",
			},
			"source_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source value",
			},
			"source_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source ID field",
			},
			"source_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source URL field",
			},
			"severity": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Severity mapping",
			},
			"source_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source status field",
			},
			"assignee": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Assignee field",
			},
			"long_text": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Long text field",
			},
			"short_text": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Short text field",
			},
			"subject": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Subject field",
			},
			"recipient_user": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Recipient user field",
			},
			"recipient_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Recipient group field",
			},
			"topic": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Topic field",
			},
			"sample_data": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Sample data",
			},
			"open_alert_when": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Condition for opening alerts",
				Elem:        getConditionSchema(),
			},
			"close_alert_when": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Condition for closing alerts",
				Elem:        getSimpleConditionSchema(),
			},
			"update_alert_when": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Condition for updating alerts",
				Elem:        getSimpleConditionSchema(),
			},
			"custom_alert_fields": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Custom alert fields",
				Elem:        getCustomAlertFieldSchema(),
			},
			"attachments": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Attachment settings",
				Elem:        getAttachmentsSchema(),
			},
		},
	}
}

// Helper function to get condition schema
func getConditionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"field_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Field name for condition",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Condition type",
			},
			"values": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Condition values",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

// Helper function to get simple condition schema
func getSimpleConditionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Condition type",
			},
			"values": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Condition values",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

// Helper function to get custom alert field schema
func getCustomAlertFieldSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Attribute name",
			},
			"attribute_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Attribute value",
			},
			"required": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the field is required",
			},
		},
	}
}

// Helper function to get attachments schema
func getAttachmentsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"base_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Base path for attachments",
			},
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL for attachments",
			},
			"file_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name pattern",
			},
			"is_link": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether attachment is a link",
			},
			"is_collection": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether attachment is a collection",
			},
		},
	}
}

// Helper function to get alert tags schema
func getAlertTagsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"business_service": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Business service tag",
			},
			"component_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Component type tag",
			},
			"component_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Component name tag",
			},
			"data_center": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Data center tag",
			},
			"environment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Environment tag",
			},
			"problem_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Problem type tag",
			},
		},
	}
}

// PLACEHOLDER FUNCTIONS - These need to be implemented for the remaining complex schemas
func getDelayingOrGroupingSchema() *schema.Resource {
	// TODO: Implement this complex schema
	return &schema.Resource{Schema: map[string]*schema.Schema{}}
}

func getFiltersSchema() *schema.Resource {
	// TODO: Implement this complex schema
	return &schema.Resource{Schema: map[string]*schema.Schema{}}
}

func getEscalationPolicyOverrideSchema() *schema.Resource {
	// TODO: Implement this complex schema
	return &schema.Resource{Schema: map[string]*schema.Schema{}}
}

func getDynamicRecipientGroupSchema() *schema.Resource {
	// TODO: Implement this complex schema
	return &schema.Resource{Schema: map[string]*schema.Schema{}}
}

func getEmailMappingSchema() *schema.Resource {
	// TODO: Implement this complex schema
	return &schema.Resource{Schema: map[string]*schema.Schema{}}
}

func getEmailAlertTagsSchema() *schema.Resource {
	// TODO: Implement this complex schema
	return &schema.Resource{Schema: map[string]*schema.Schema{}}
}

func getEmailFiltersSchema() *schema.Resource {
	// TODO: Implement this complex schema
	return &schema.Resource{Schema: map[string]*schema.Schema{}}
}

func getChatURLMappingSchema() *schema.Resource {
	// TODO: Implement this complex schema
	return &schema.Resource{Schema: map[string]*schema.Schema{}}
}

// CRUD OPERATIONS - Basic implementations

func resourceInboundIntegrationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	inboundIntegration := InboundIntegration{
		InboundIntegrationName: d.Get("inbound_integration_name").(string),
		Type:                   d.Get("type").(string),
		Enabled:                d.Get("enabled").(bool),
	}

	if v, ok := d.GetOk("sequence"); ok {
		inboundIntegration.Sequence = v.(int)
	}
	if v, ok := d.GetOk("escalation_policy"); ok {
		inboundIntegration.EscalationPolicy = v.(string)
	}
	if v, ok := d.GetOk("recipient_groups"); ok {
		inboundIntegration.RecipientGroups = expandStringSlice(v.([]interface{}))
	}
	if v, ok := d.GetOk("recipient_users"); ok {
		inboundIntegration.RecipientUsers = expandStringSlice(v.([]interface{}))
	}
	if v, ok := d.GetOk("inbound_template_id"); ok {
		inboundIntegration.InboundTemplateID = v.(int)
	}
	if v, ok := d.GetOk("mail_box"); ok {
		inboundIntegration.MailBox = v.(string)
	}

	// Handle nested structures
	if v, ok := d.GetOk("bridge"); ok {
		inboundIntegration.Bridge = expandBridge(v.([]interface{}))
	}
	if v, ok := d.GetOk("api_settings"); ok {
		inboundIntegration.APISettings = expandAPISettings(v.([]interface{}))
	}
	if v, ok := d.GetOk("heartbeat_settings"); ok {
		inboundIntegration.HeartbeatSettings = expandHeartbeatSettings(v.([]interface{}))
	}

	// Create inbound integration via API
	var result InboundIntegration
	err := client.post(ctx, "/api/v2/integrations/inbound", inboundIntegration, &result)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating inbound integration: %v", err))
	}

	d.SetId(strconv.Itoa(result.InboundIntegrationID))

	return resourceInboundIntegrationRead(ctx, d, meta)
}

func resourceInboundIntegrationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	inboundIntegrationID := d.Id()
	
	var inboundIntegration InboundIntegration
	err := client.get(ctx, fmt.Sprintf("/api/v2/integrations/inbound/%s", inboundIntegrationID), &inboundIntegration)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading inbound integration: %v", err))
	}

	d.Set("inbound_integration_id", inboundIntegration.InboundIntegrationID)
	d.Set("inbound_integration_name", inboundIntegration.InboundIntegrationName)
	d.Set("type", inboundIntegration.Type)
	d.Set("sequence", inboundIntegration.Sequence)
	d.Set("enabled", inboundIntegration.Enabled)
	d.Set("escalation_policy", inboundIntegration.EscalationPolicy)
	d.Set("recipient_groups", inboundIntegration.RecipientGroups)
	d.Set("recipient_users", inboundIntegration.RecipientUsers)
	d.Set("inbound_template_id", inboundIntegration.InboundTemplateID)
	d.Set("mail_box", inboundIntegration.MailBox)

	// Set nested structures
	d.Set("bridge", flattenBridge(inboundIntegration.Bridge))
	d.Set("api_settings", flattenAPISettings(inboundIntegration.APISettings))
	d.Set("heartbeat_settings", flattenHeartbeatSettings(inboundIntegration.HeartbeatSettings))

	return nil
}

func resourceInboundIntegrationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	inboundIntegrationID := d.Id()
	inboundIntegration := InboundIntegration{
		InboundIntegrationID:   d.Get("inbound_integration_id").(int),
		InboundIntegrationName: d.Get("inbound_integration_name").(string),
		Type:                   d.Get("type").(string),
		Enabled:                d.Get("enabled").(bool),
	}

	if v, ok := d.GetOk("sequence"); ok {
		inboundIntegration.Sequence = v.(int)
	}
	if v, ok := d.GetOk("escalation_policy"); ok {
		inboundIntegration.EscalationPolicy = v.(string)
	}
	if v, ok := d.GetOk("recipient_groups"); ok {
		inboundIntegration.RecipientGroups = expandStringSlice(v.([]interface{}))
	}
	if v, ok := d.GetOk("recipient_users"); ok {
		inboundIntegration.RecipientUsers = expandStringSlice(v.([]interface{}))
	}
	if v, ok := d.GetOk("inbound_template_id"); ok {
		inboundIntegration.InboundTemplateID = v.(int)
	}
	if v, ok := d.GetOk("mail_box"); ok {
		inboundIntegration.MailBox = v.(string)
	}

	// Handle nested structures
	if v, ok := d.GetOk("bridge"); ok {
		inboundIntegration.Bridge = expandBridge(v.([]interface{}))
	}
	if v, ok := d.GetOk("api_settings"); ok {
		inboundIntegration.APISettings = expandAPISettings(v.([]interface{}))
	}
	if v, ok := d.GetOk("heartbeat_settings"); ok {
		inboundIntegration.HeartbeatSettings = expandHeartbeatSettings(v.([]interface{}))
	}

	// Update inbound integration via API
	err := client.put(ctx, fmt.Sprintf("/api/v2/integrations/inbound/%s", inboundIntegrationID), inboundIntegration, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating inbound integration: %v", err))
	}

	return resourceInboundIntegrationRead(ctx, d, meta)
}

func resourceInboundIntegrationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	inboundIntegrationID := d.Id()
	err := client.delete(ctx, fmt.Sprintf("/api/v2/integrations/inbound/%s", inboundIntegrationID))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting inbound integration: %v", err))
	}

	d.SetId("")
	return nil
}

// HELPER FUNCTIONS FOR EXPANDING AND FLATTENING

// expandBridge converts Terraform data to Bridge struct
func expandBridge(bridgeData []interface{}) *InboundIntegrationBridge {
	if len(bridgeData) == 0 {
		return nil
	}

	if bridgeData[0] == nil {
		return nil
	}

	bridgeMap := bridgeData[0].(map[string]interface{})
	bridge := &InboundIntegrationBridge{}

	if v, ok := bridgeMap["telephone_number"]; ok && v.(string) != "" {
		bridge.TelephoneNumber = v.(string)
	}
	if v, ok := bridgeMap["access_code"]; ok && v.(string) != "" {
		bridge.AccessCode = v.(string)
	}

	return bridge
}

// flattenBridge converts Bridge struct to Terraform data
func flattenBridge(bridge *InboundIntegrationBridge) []map[string]interface{} {
	if bridge == nil {
		return nil
	}

	return []map[string]interface{}{
		{
			"telephone_number": bridge.TelephoneNumber,
			"access_code":      bridge.AccessCode,
		},
	}
}

// expandAPISettings converts Terraform data to APISettings struct (basic implementation)
func expandAPISettings(apiSettingsData []interface{}) *InboundIntegrationAPISettings {
	if len(apiSettingsData) == 0 {
		return nil
	}

	if apiSettingsData[0] == nil {
		return nil
	}

	apiSettingsMap := apiSettingsData[0].(map[string]interface{})
	apiSettings := &InboundIntegrationAPISettings{}

	if v, ok := apiSettingsMap["is_bidirection"]; ok {
		apiSettings.IsBidirection = v.(bool)
	}

	// TODO: Implement nested structures like url_mapping, alert_tags, etc.
	// For now, only handling basic boolean field

	return apiSettings
}

// flattenAPISettings converts APISettings struct to Terraform data (basic implementation)
func flattenAPISettings(apiSettings *InboundIntegrationAPISettings) []map[string]interface{} {
	if apiSettings == nil {
		return nil
	}

	return []map[string]interface{}{
		{
			"is_bidirection": apiSettings.IsBidirection,
			// TODO: Add nested structures
		},
	}
}

// expandHeartbeatSettings converts Terraform data to HeartbeatSettings struct
func expandHeartbeatSettings(heartbeatData []interface{}) *InboundIntegrationHeartbeatSettings {
	if len(heartbeatData) == 0 {
		return nil
	}

	if heartbeatData[0] == nil {
		return nil
	}

	heartbeatMap := heartbeatData[0].(map[string]interface{})
	heartbeat := &InboundIntegrationHeartbeatSettings{}

	if v, ok := heartbeatMap["heartbeat_interval_in_min"]; ok {
		heartbeat.HeartbeatIntervalInMin = v.(int)
	}

	return heartbeat
}

// flattenHeartbeatSettings converts HeartbeatSettings struct to Terraform data
func flattenHeartbeatSettings(heartbeat *InboundIntegrationHeartbeatSettings) []map[string]interface{} {
	if heartbeat == nil {
		return nil
	}

	return []map[string]interface{}{
		{
			"heartbeat_interval_in_min": heartbeat.HeartbeatIntervalInMin,
		},
	}
} 