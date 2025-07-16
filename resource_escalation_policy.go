package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEscalationPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEscalationPolicyCreate,
		ReadContext:   resourceEscalationPolicyRead,
		UpdateContext: resourceEscalationPolicyUpdate,
		DeleteContext: resourceEscalationPolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"escalation_policy_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The unique identifier for the escalation policy",
			},
			"escalation_policy_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the escalation policy",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the escalation policy",
			},
			"priority": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Priority level for the escalation policy",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether the escalation policy is enabled",
			},
			"quick_launch": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether quick launch is enabled",
			},
			"notify_using_centralized_settings": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether to use centralized notification settings",
			},
			"wait_time_before_notifying_next_group_in_min": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Wait time before notifying next group in minutes",
			},
			"member_roles": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Member roles configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member_role_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of member role",
						},
						"wait_time_between_members_in_mins": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Wait time between members in minutes",
						},
						"role_wait_time_in_mins": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Role wait time in minutes",
						},
						"no_of_retries": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of retries",
						},
						"retry_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Retry interval",
						},
						"contact_methods": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Contact methods for this member role",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"contact_method_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The contact method name",
									},
									"wait_time_in_mins": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Wait time in minutes",
									},
									"repeat": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to repeat the contact",
									},
									"repeat_times": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Number of times to repeat",
									},
									"repeat_minutes": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Minutes between repeats",
									},
									"to_bcc_or_cc": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "BCC or CC setting",
									},
									"sequence": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Sequence number",
									},
								},
							},
						},
					},
				},
			},
			"group_contact_notifications": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Group contact notifications configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"contact_methods": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Contact methods for group notifications",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"contact_method_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The contact method name",
									},
									"wait_time_in_mins": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Wait time in minutes",
									},
									"to_bcc_or_cc": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "BCC or CC setting",
									},
								},
							},
						},
					},
				},
			},
			"workflows": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Associated workflows",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"workflow_id": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The workflow ID",
						},
						"workflow_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The workflow name",
						},
					},
				},
			},
			"outbound_integrations": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Outbound integrations configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"outbound_integration_id": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The outbound integration ID",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The integration name",
						},
						"interval_in_sec": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Interval in seconds",
						},
						"actions": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Actions for the integration",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The action name",
									},
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether the action is enabled",
									},
								},
							},
						},
					},
				},
			},
			"outbound_actions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Outbound actions configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action_id": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The action ID",
						},
						"action_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The action name",
						},
					},
				},
			},
			"options": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Options configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acknowledgement": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Acknowledgement settings",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"phone": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable phone acknowledgement",
									},
									"sms": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable SMS acknowledgement",
									},
									"email": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable email acknowledgement",
									},
									"group_chat": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable group chat acknowledgement",
									},
								},
							},
						},
						"assignment": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Assignment settings",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"phone": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable phone assignment",
									},
									"sms": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable SMS assignment",
									},
									"email": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable email assignment",
									},
									"group_chat": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable group chat assignment",
									},
								},
							},
						},
						"escalate": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Escalate settings",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"phone": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable phone escalation",
									},
									"sms": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable SMS escalation",
									},
									"email": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable email escalation",
									},
									"group_chat": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable group chat escalation",
									},
								},
							},
						},
						"close": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Close settings",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"phone": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable phone close",
									},
									"sms": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable SMS close",
									},
									"email": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable email close",
									},
									"group_chat": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable group chat close",
									},
								},
							},
						},
						"notification_settings": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Notification settings",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"email": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Email notification setting",
									},
									"phone": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Phone notification setting",
									},
									"sms": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "SMS notification setting",
									},
								},
							},
						},
						"escalation_policy_name_for_reply": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Escalation policy name for reply",
						},
						"sla_in_hours": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "SLA in hours",
						},
						"message_text": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Message text",
						},
						"include_alert_id_in_subject": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Include alert ID in subject",
						},
						"one_email_per_message": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "One email per message",
						},
						"one_message_per_recipient": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "One message per recipient",
						},
						"sequence_group_first": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Sequence group first",
						},
						"alert_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Alert type",
						},
						"recipients": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Recipients list",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"recipient_type_id": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Recipient type ID",
									},
									"recipient_id": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Recipient ID",
									},
									"recipient_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Recipient name",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// CRUD operations for escalation policies

func resourceEscalationPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	escalationPolicy := EscalationPolicy{
		EscalationPolicyName:           d.Get("escalation_policy_name").(string),
		Enabled:                        d.Get("enabled").(bool),
		QuickLaunch:                    d.Get("quick_launch").(bool),
		NotifyUsingCentralizedSettings: d.Get("notify_using_centralized_settings").(bool),
	}

	if v, ok := d.GetOk("description"); ok {
		escalationPolicy.Description = v.(string)
	}
	if v, ok := d.GetOk("priority"); ok {
		escalationPolicy.Priority = v.(string)
	}
	if v, ok := d.GetOk("wait_time_before_notifying_next_group_in_min"); ok {
		escalationPolicy.WaitTimeBeforeNotifyingNextGroupInMin = v.(int)
	}

	// Handle nested structures
	if v, ok := d.GetOk("member_roles"); ok {
		escalationPolicy.MemberRoles = expandMemberRoles(v.([]interface{}))
	}
	if v, ok := d.GetOk("group_contact_notifications"); ok {
		escalationPolicy.GroupContactNotifications = expandGroupContactNotifications(v.([]interface{}))
	}
	if v, ok := d.GetOk("workflows"); ok {
		escalationPolicy.Workflows = expandEscalationPolicyWorkflows(v.([]interface{}))
	}
	if v, ok := d.GetOk("outbound_integrations"); ok {
		escalationPolicy.OutboundIntegrations = expandOutboundIntegrations(v.([]interface{}))
	}
	if v, ok := d.GetOk("outbound_actions"); ok {
		escalationPolicy.OutboundActions = expandOutboundActions(v.([]interface{}))
	}
	if v, ok := d.GetOk("options"); ok {
		escalationPolicy.Options = expandOptions(v.([]interface{}))
	}

	// Create escalation policy via API
	var result EscalationPolicy
	err := client.post(ctx, "/api/v2/escalation_policies", escalationPolicy, &result)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating escalation policy: %v", err))
	}

	d.SetId(strconv.Itoa(result.EscalationPolicyID))

	return resourceEscalationPolicyRead(ctx, d, meta)
}

func resourceEscalationPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	escalationPolicyID := d.Id()
	var escalationPolicy EscalationPolicy
	err := client.get(ctx, fmt.Sprintf("/api/v2/escalation_policies/%s", escalationPolicyID), &escalationPolicy)
	if err != nil {
		if err.Error() == "404" {
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error reading escalation policy: %v", err))
	}

	// Set all the escalation policy attributes
	d.Set("escalation_policy_id", escalationPolicy.EscalationPolicyID)
	d.Set("escalation_policy_name", escalationPolicy.EscalationPolicyName)
	d.Set("description", escalationPolicy.Description)
	d.Set("priority", escalationPolicy.Priority)
	d.Set("enabled", escalationPolicy.Enabled)
	d.Set("quick_launch", escalationPolicy.QuickLaunch)
	d.Set("notify_using_centralized_settings", escalationPolicy.NotifyUsingCentralizedSettings)
	d.Set("wait_time_before_notifying_next_group_in_min", escalationPolicy.WaitTimeBeforeNotifyingNextGroupInMin)

	// Set nested structures
	d.Set("member_roles", flattenMemberRoles(escalationPolicy.MemberRoles))
	d.Set("group_contact_notifications", flattenGroupContactNotifications(escalationPolicy.GroupContactNotifications))
	d.Set("workflows", flattenEscalationPolicyWorkflows(escalationPolicy.Workflows))
	d.Set("outbound_integrations", flattenOutboundIntegrations(escalationPolicy.OutboundIntegrations))
	d.Set("outbound_actions", flattenOutboundActions(escalationPolicy.OutboundActions))
	d.Set("options", flattenOptions(escalationPolicy.Options))

	return nil
}

func resourceEscalationPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	escalationPolicyID := d.Id()
	escalationPolicy := EscalationPolicy{
		EscalationPolicyID:             d.Get("escalation_policy_id").(int),
		EscalationPolicyName:           d.Get("escalation_policy_name").(string),
		Enabled:                        d.Get("enabled").(bool),
		QuickLaunch:                    d.Get("quick_launch").(bool),
		NotifyUsingCentralizedSettings: d.Get("notify_using_centralized_settings").(bool),
	}

	if v, ok := d.GetOk("description"); ok {
		escalationPolicy.Description = v.(string)
	}
	if v, ok := d.GetOk("priority"); ok {
		escalationPolicy.Priority = v.(string)
	}
	if v, ok := d.GetOk("wait_time_before_notifying_next_group_in_min"); ok {
		escalationPolicy.WaitTimeBeforeNotifyingNextGroupInMin = v.(int)
	}

	// Handle nested structures
	if v, ok := d.GetOk("member_roles"); ok {
		escalationPolicy.MemberRoles = expandMemberRoles(v.([]interface{}))
	}
	if v, ok := d.GetOk("group_contact_notifications"); ok {
		escalationPolicy.GroupContactNotifications = expandGroupContactNotifications(v.([]interface{}))
	}
	if v, ok := d.GetOk("workflows"); ok {
		escalationPolicy.Workflows = expandEscalationPolicyWorkflows(v.([]interface{}))
	}
	if v, ok := d.GetOk("outbound_integrations"); ok {
		escalationPolicy.OutboundIntegrations = expandOutboundIntegrations(v.([]interface{}))
	}
	if v, ok := d.GetOk("outbound_actions"); ok {
		escalationPolicy.OutboundActions = expandOutboundActions(v.([]interface{}))
	}
	if v, ok := d.GetOk("options"); ok {
		escalationPolicy.Options = expandOptions(v.([]interface{}))
	}

	// Update escalation policy via API
	err := client.put(ctx, fmt.Sprintf("/api/v2/escalation_policies/%s", escalationPolicyID), escalationPolicy, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating escalation policy: %v", err))
	}

	return resourceEscalationPolicyRead(ctx, d, meta)
}

func resourceEscalationPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	escalationPolicyID := d.Id()
	err := client.delete(ctx, fmt.Sprintf("/api/v2/escalation_policies/%s", escalationPolicyID))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting escalation policy: %v", err))
	}

	d.SetId("")
	return nil
}

// Helper functions for expanding and flattening nested structures

// expandMemberRoles converts Terraform data to MemberRole structs
func expandMemberRoles(memberRolesData []interface{}) []EscalationPolicyMemberRole {
	if len(memberRolesData) == 0 {
		return nil
	}

	memberRoles := make([]EscalationPolicyMemberRole, len(memberRolesData))
	for i, memberRoleData := range memberRolesData {
		memberRoleMap := memberRoleData.(map[string]interface{})
		memberRoles[i] = EscalationPolicyMemberRole{
			MemberRoleType: memberRoleMap["member_role_type"].(string),
		}
		
		if v, ok := memberRoleMap["wait_time_between_members_in_mins"]; ok {
			memberRoles[i].WaitTimeBetweenMembersInMins = v.(int)
		}
		if v, ok := memberRoleMap["role_wait_time_in_mins"]; ok {
			memberRoles[i].RoleWaitTimeInMins = v.(int)
		}
		if v, ok := memberRoleMap["no_of_retries"]; ok {
			memberRoles[i].NoOfRetries = v.(int)
		}
		if v, ok := memberRoleMap["retry_interval"]; ok {
			memberRoles[i].RetryInterval = v.(int)
		}
		
		// Handle contact methods
		if v, ok := memberRoleMap["contact_methods"]; ok && v != nil {
			memberRoles[i].ContactMethods = expandEscalationPolicyContactMethods(v.([]interface{}))
		}
	}
	return memberRoles
}

// flattenMemberRoles converts MemberRole structs to Terraform data
func flattenMemberRoles(memberRoles []EscalationPolicyMemberRole) []map[string]interface{} {
	if len(memberRoles) == 0 {
		return nil
	}

	result := make([]map[string]interface{}, len(memberRoles))
	for i, memberRole := range memberRoles {
		result[i] = map[string]interface{}{
			"member_role_type":                 memberRole.MemberRoleType,
			"wait_time_between_members_in_mins": memberRole.WaitTimeBetweenMembersInMins,
			"role_wait_time_in_mins":           memberRole.RoleWaitTimeInMins,
			"no_of_retries":                    memberRole.NoOfRetries,
			"retry_interval":                   memberRole.RetryInterval,
			"contact_methods":                  flattenEscalationPolicyContactMethods(memberRole.ContactMethods),
		}
	}
	return result
}

// expandEscalationPolicyContactMethods converts Terraform data to ContactMethod structs
func expandEscalationPolicyContactMethods(contactMethodsData []interface{}) []EscalationPolicyContactMethod {
	if len(contactMethodsData) == 0 {
		return nil
	}

	contactMethods := make([]EscalationPolicyContactMethod, len(contactMethodsData))
	for i, contactMethodData := range contactMethodsData {
		contactMethodMap := contactMethodData.(map[string]interface{})
		contactMethods[i] = EscalationPolicyContactMethod{
			ContactMethodName: contactMethodMap["contact_method_name"].(string),
		}
		
		if v, ok := contactMethodMap["wait_time_in_mins"]; ok {
			contactMethods[i].WaitTimeInMins = v.(int)
		}
		if v, ok := contactMethodMap["repeat"]; ok {
			contactMethods[i].Repeat = v.(bool)
		}
		if v, ok := contactMethodMap["repeat_times"]; ok {
			contactMethods[i].RepeatTimes = v.(int)
		}
		if v, ok := contactMethodMap["repeat_minutes"]; ok {
			contactMethods[i].RepeatMinutes = v.(int)
		}
		if v, ok := contactMethodMap["to_bcc_or_cc"]; ok && v.(string) != "" {
			contactMethods[i].ToBccOrCc = v.(string)
		}
		if v, ok := contactMethodMap["sequence"]; ok {
			contactMethods[i].Sequence = v.(int)
		}
	}
	return contactMethods
}

// flattenEscalationPolicyContactMethods converts ContactMethod structs to Terraform data
func flattenEscalationPolicyContactMethods(contactMethods []EscalationPolicyContactMethod) []map[string]interface{} {
	if len(contactMethods) == 0 {
		return nil
	}

	result := make([]map[string]interface{}, len(contactMethods))
	for i, contactMethod := range contactMethods {
		result[i] = map[string]interface{}{
			"contact_method_name": contactMethod.ContactMethodName,
			"wait_time_in_mins":   contactMethod.WaitTimeInMins,
			"repeat":              contactMethod.Repeat,
			"repeat_times":        contactMethod.RepeatTimes,
			"repeat_minutes":      contactMethod.RepeatMinutes,
			"to_bcc_or_cc":        contactMethod.ToBccOrCc,
			"sequence":            contactMethod.Sequence,
		}
	}
	return result
}

// expandGroupContactNotifications converts Terraform data to GroupContactNotifications struct
func expandGroupContactNotifications(groupNotificationsData []interface{}) *EscalationPolicyGroupContactNotifications {
	if len(groupNotificationsData) == 0 {
		return nil
	}

	// Check if the first element is nil
	if groupNotificationsData[0] == nil {
		return nil
	}

	groupNotificationsMap := groupNotificationsData[0].(map[string]interface{})
	result := &EscalationPolicyGroupContactNotifications{}
	
	if v, ok := groupNotificationsMap["contact_methods"]; ok && v != nil {
		result.ContactMethods = expandEscalationPolicyGroupContactMethods(v.([]interface{}))
	}
	
	return result
}

// flattenGroupContactNotifications converts GroupContactNotifications struct to Terraform data
func flattenGroupContactNotifications(groupNotifications *EscalationPolicyGroupContactNotifications) []map[string]interface{} {
	if groupNotifications == nil {
		return nil
	}

	return []map[string]interface{}{
		{
			"contact_methods": flattenEscalationPolicyGroupContactMethods(groupNotifications.ContactMethods),
		},
	}
}

// expandEscalationPolicyGroupContactMethods converts Terraform data to GroupContactMethod structs
func expandEscalationPolicyGroupContactMethods(contactMethodsData []interface{}) []EscalationPolicyGroupContactMethod {
	if len(contactMethodsData) == 0 {
		return nil
	}

	contactMethods := make([]EscalationPolicyGroupContactMethod, len(contactMethodsData))
	for i, contactMethodData := range contactMethodsData {
		contactMethodMap := contactMethodData.(map[string]interface{})
		contactMethods[i] = EscalationPolicyGroupContactMethod{
			ContactMethodName: contactMethodMap["contact_method_name"].(string),
		}
		
		if v, ok := contactMethodMap["wait_time_in_mins"]; ok {
			contactMethods[i].WaitTimeInMins = v.(int)
		}
		if v, ok := contactMethodMap["to_bcc_or_cc"]; ok && v.(string) != "" {
			contactMethods[i].ToBccOrCc = v.(string)
		}
	}
	return contactMethods
}

// flattenEscalationPolicyGroupContactMethods converts GroupContactMethod structs to Terraform data
func flattenEscalationPolicyGroupContactMethods(contactMethods []EscalationPolicyGroupContactMethod) []map[string]interface{} {
	if len(contactMethods) == 0 {
		return nil
	}

	result := make([]map[string]interface{}, len(contactMethods))
	for i, contactMethod := range contactMethods {
		result[i] = map[string]interface{}{
			"contact_method_name": contactMethod.ContactMethodName,
			"wait_time_in_mins":   contactMethod.WaitTimeInMins,
			"to_bcc_or_cc":        contactMethod.ToBccOrCc,
		}
	}
	return result
}

// expandWorkflows converts Terraform data to Workflow structs
func expandEscalationPolicyWorkflows(workflowsData []interface{}) []EscalationPolicyWorkflow {
	if len(workflowsData) == 0 {
		return nil
	}

	workflows := make([]EscalationPolicyWorkflow, len(workflowsData))
	for i, workflowData := range workflowsData {
		workflowMap := workflowData.(map[string]interface{})
		workflows[i] = EscalationPolicyWorkflow{
			WorkflowID: workflowMap["workflow_id"].(int),
		}
		
		if v, ok := workflowMap["workflow_name"]; ok && v.(string) != "" {
			workflows[i].WorkflowName = v.(string)
		}
	}
	return workflows
}

// flattenWorkflows converts Workflow structs to Terraform data
func flattenEscalationPolicyWorkflows(workflows []EscalationPolicyWorkflow) []map[string]interface{} {
	if len(workflows) == 0 {
		return nil
	}

	result := make([]map[string]interface{}, len(workflows))
	for i, workflow := range workflows {
		result[i] = map[string]interface{}{
			"workflow_id":   workflow.WorkflowID,
			"workflow_name": workflow.WorkflowName,
		}
	}
	return result
}

// expandOutboundIntegrations converts Terraform data to OutboundIntegration structs
func expandOutboundIntegrations(integrationsData []interface{}) []EscalationPolicyOutboundIntegration {
	if len(integrationsData) == 0 {
		return nil
	}

	integrations := make([]EscalationPolicyOutboundIntegration, len(integrationsData))
	for i, integrationData := range integrationsData {
		integrationMap := integrationData.(map[string]interface{})
		integrations[i] = EscalationPolicyOutboundIntegration{
			OutboundIntegrationID: integrationMap["outbound_integration_id"].(int),
		}
		
		if v, ok := integrationMap["name"]; ok && v.(string) != "" {
			integrations[i].Name = v.(string)
		}
		if v, ok := integrationMap["interval_in_sec"]; ok {
			integrations[i].IntervalInSec = v.(int)
		}
		
		// Handle actions
		if v, ok := integrationMap["actions"]; ok && v != nil {
			integrations[i].Actions = expandOutboundIntegrationActions(v.([]interface{}))
		}
	}
	return integrations
}

// flattenOutboundIntegrations converts OutboundIntegration structs to Terraform data
func flattenOutboundIntegrations(integrations []EscalationPolicyOutboundIntegration) []map[string]interface{} {
	if len(integrations) == 0 {
		return nil
	}

	result := make([]map[string]interface{}, len(integrations))
	for i, integration := range integrations {
		result[i] = map[string]interface{}{
			"outbound_integration_id": integration.OutboundIntegrationID,
			"name":                    integration.Name,
			"interval_in_sec":         integration.IntervalInSec,
			"actions":                 flattenOutboundIntegrationActions(integration.Actions),
		}
	}
	return result
}

// expandOutboundIntegrationActions converts Terraform data to OutboundIntegrationAction structs
func expandOutboundIntegrationActions(actionsData []interface{}) []EscalationPolicyOutboundIntegrationAction {
	if len(actionsData) == 0 {
		return nil
	}

	actions := make([]EscalationPolicyOutboundIntegrationAction, len(actionsData))
	for i, actionData := range actionsData {
		actionMap := actionData.(map[string]interface{})
		actions[i] = EscalationPolicyOutboundIntegrationAction{
			ActionName: actionMap["action_name"].(string),
		}
		
		if v, ok := actionMap["enabled"]; ok {
			actions[i].Enabled = v.(bool)
		}
	}
	return actions
}

// flattenOutboundIntegrationActions converts OutboundIntegrationAction structs to Terraform data
func flattenOutboundIntegrationActions(actions []EscalationPolicyOutboundIntegrationAction) []map[string]interface{} {
	if len(actions) == 0 {
		return nil
	}

	result := make([]map[string]interface{}, len(actions))
	for i, action := range actions {
		result[i] = map[string]interface{}{
			"action_name": action.ActionName,
			"enabled":     action.Enabled,
		}
	}
	return result
}

// expandOutboundActions converts Terraform data to OutboundAction structs
func expandOutboundActions(actionsData []interface{}) []EscalationPolicyOutboundAction {
	if len(actionsData) == 0 {
		return nil
	}

	actions := make([]EscalationPolicyOutboundAction, len(actionsData))
	for i, actionData := range actionsData {
		actionMap := actionData.(map[string]interface{})
		actions[i] = EscalationPolicyOutboundAction{
			ActionID: actionMap["action_id"].(int),
		}
		
		if v, ok := actionMap["action_name"]; ok && v.(string) != "" {
			actions[i].ActionName = v.(string)
		}
	}
	return actions
}

// flattenOutboundActions converts OutboundAction structs to Terraform data
func flattenOutboundActions(actions []EscalationPolicyOutboundAction) []map[string]interface{} {
	if len(actions) == 0 {
		return nil
	}

	result := make([]map[string]interface{}, len(actions))
	for i, action := range actions {
		result[i] = map[string]interface{}{
			"action_id":   action.ActionID,
			"action_name": action.ActionName,
		}
	}
	return result
}

// expandOptions converts Terraform data to Options struct
func expandOptions(optionsData []interface{}) *EscalationPolicyOptions {
	if len(optionsData) == 0 {
		return nil
	}

	// Check if the first element is nil
	if optionsData[0] == nil {
		return nil
	}

	optionsMap := optionsData[0].(map[string]interface{})
	options := &EscalationPolicyOptions{}
	
	if v, ok := optionsMap["acknowledgement"]; ok && v != nil {
		options.Acknowledgement = expandOptionSettings(v.([]interface{}))
	}
	if v, ok := optionsMap["assignment"]; ok && v != nil {
		options.Assignment = expandOptionSettings(v.([]interface{}))
	}
	if v, ok := optionsMap["escalate"]; ok && v != nil {
		options.Escalate = expandOptionSettings(v.([]interface{}))
	}
	if v, ok := optionsMap["close"]; ok && v != nil {
		options.Close = expandOptionSettings(v.([]interface{}))
	}
	if v, ok := optionsMap["notification_settings"]; ok && v != nil {
		options.NotificationSettings = expandNotificationSettings(v.([]interface{}))
	}
	if v, ok := optionsMap["escalation_policy_name_for_reply"]; ok && v.(string) != "" {
		options.EscalationPolicyNameForReply = v.(string)
	}
	if v, ok := optionsMap["sla_in_hours"]; ok {
		options.SlaInHours = float64(v.(int))
	}
	if v, ok := optionsMap["message_text"]; ok && v.(string) != "" {
		options.MessageText = v.(string)
	}
	if v, ok := optionsMap["include_alert_id_in_subject"]; ok {
		options.IncludeAlertIDInSubject = v.(bool)
	}
	if v, ok := optionsMap["one_email_per_message"]; ok {
		options.OneEmailPerMessage = v.(bool)
	}
	if v, ok := optionsMap["one_message_per_recipient"]; ok {
		options.OneMessagePerRecipient = v.(bool)
	}
	if v, ok := optionsMap["sequence_group_first"]; ok {
		options.SequenceGroupFirst = v.(bool)
	}
	if v, ok := optionsMap["alert_type"]; ok && v.(string) != "" {
		options.AlertType = v.(string)
	}
	if v, ok := optionsMap["recipients"]; ok && v != nil {
		options.Recipients = expandRecipients(v.([]interface{}))
	}
	
	return options
}

// flattenOptions converts Options struct to Terraform data
func flattenOptions(options *EscalationPolicyOptions) []map[string]interface{} {
	if options == nil {
		return nil
	}

	return []map[string]interface{}{
		{
			"acknowledgement":                    flattenOptionSettings(options.Acknowledgement),
			"assignment":                        flattenOptionSettings(options.Assignment),
			"escalate":                          flattenOptionSettings(options.Escalate),
			"close":                            flattenOptionSettings(options.Close),
			"notification_settings":             flattenNotificationSettings(options.NotificationSettings),
			"escalation_policy_name_for_reply": options.EscalationPolicyNameForReply,
			"sla_in_hours":                     int(options.SlaInHours),
			"message_text":                     options.MessageText,
			"include_alert_id_in_subject":      options.IncludeAlertIDInSubject,
			"one_email_per_message":            options.OneEmailPerMessage,
			"one_message_per_recipient":        options.OneMessagePerRecipient,
			"sequence_group_first":             options.SequenceGroupFirst,
			"alert_type":                       options.AlertType,
			"recipients":                       flattenRecipients(options.Recipients),
		},
	}
}

// expandOptionSettings converts Terraform data to OptionSettings struct
func expandOptionSettings(settingsData []interface{}) *EscalationPolicyOptionSettings {
	if len(settingsData) == 0 {
		return nil
	}

	// Check if the first element is nil
	if settingsData[0] == nil {
		return nil
	}

	settingsMap := settingsData[0].(map[string]interface{})
	settings := &EscalationPolicyOptionSettings{}
	
	if v, ok := settingsMap["phone"]; ok {
		settings.Phone = v.(bool)
	}
	if v, ok := settingsMap["sms"]; ok {
		settings.SMS = v.(bool)
	}
	if v, ok := settingsMap["email"]; ok {
		settings.Email = v.(bool)
	}
	if v, ok := settingsMap["group_chat"]; ok {
		settings.GroupChat = v.(bool)
	}
	
	return settings
}

// flattenOptionSettings converts OptionSettings struct to Terraform data
func flattenOptionSettings(settings *EscalationPolicyOptionSettings) []map[string]interface{} {
	if settings == nil {
		return nil
	}

	return []map[string]interface{}{
		{
			"phone":      settings.Phone,
			"sms":        settings.SMS,
			"email":      settings.Email,
			"group_chat": settings.GroupChat,
		},
	}
}

// expandNotificationSettings converts Terraform data to NotificationSettings struct
func expandNotificationSettings(notificationData []interface{}) *EscalationPolicyNotificationSettings {
	if len(notificationData) == 0 {
		return nil
	}

	// Check if the first element is nil
	if notificationData[0] == nil {
		return nil
	}

	notificationMap := notificationData[0].(map[string]interface{})
	settings := &EscalationPolicyNotificationSettings{}
	
	if v, ok := notificationMap["email"]; ok && v.(string) != "" {
		settings.Email = v.(string)
	}
	if v, ok := notificationMap["phone"]; ok && v.(string) != "" {
		settings.Phone = v.(string)
	}
	if v, ok := notificationMap["sms"]; ok && v.(string) != "" {
		settings.SMS = v.(string)
	}
	
	return settings
}

// flattenNotificationSettings converts NotificationSettings struct to Terraform data
func flattenNotificationSettings(settings *EscalationPolicyNotificationSettings) []map[string]interface{} {
	if settings == nil {
		return nil
	}

	return []map[string]interface{}{
		{
			"email": settings.Email,
			"phone": settings.Phone,
			"sms":   settings.SMS,
		},
	}
}

// expandRecipients converts Terraform data to Recipient structs
func expandRecipients(recipientsData []interface{}) []EscalationPolicyRecipient {
	if len(recipientsData) == 0 {
		return nil
	}

	recipients := make([]EscalationPolicyRecipient, len(recipientsData))
	for i, recipientData := range recipientsData {
		recipientMap := recipientData.(map[string]interface{})
		recipients[i] = EscalationPolicyRecipient{
			RecipientTypeID: recipientMap["recipient_type_id"].(int),
			RecipientID:     recipientMap["recipient_id"].(int),
		}
		
		if v, ok := recipientMap["recipient_name"]; ok && v.(string) != "" {
			recipients[i].RecipientName = v.(string)
		}
	}
	return recipients
}

// flattenRecipients converts Recipient structs to Terraform data
func flattenRecipients(recipients []EscalationPolicyRecipient) []map[string]interface{} {
	if len(recipients) == 0 {
		return nil
	}

	result := make([]map[string]interface{}, len(recipients))
	for i, recipient := range recipients {
		result[i] = map[string]interface{}{
			"recipient_type_id": recipient.RecipientTypeID,
			"recipient_id":      recipient.RecipientID,
			"recipient_name":    recipient.RecipientName,
		}
	}
	return result
}