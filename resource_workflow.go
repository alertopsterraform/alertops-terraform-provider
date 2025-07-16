package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWorkflow() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkflowCreate,
		ReadContext:   resourceWorkflowRead,
		UpdateContext: resourceWorkflowUpdate,
		DeleteContext: resourceWorkflowDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"workflow_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The unique identifier for the workflow",
			},
			"workflow_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the workflow",
			},
			"workflow_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of workflow (Alert, Notification, Message)",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					validTypes := []string{"Alert", "Notification", "Message"}
					for _, validType := range validTypes {
						if v == validType {
							return
						}
					}
					errs = append(errs, fmt.Errorf("%q must be one of %v, got: %q", key, validTypes, v))
					return
				},
			},
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether the workflow is enabled",
			},
			"alert_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The alert type for the workflow",
			},
			"scheduled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether the workflow is scheduled",
			},
			"recurrence_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The recurrence interval for scheduled workflows",
			},
			"conditions": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The conditions that trigger this workflow",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The condition type",
						},
						"match": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The match criteria (all, any)",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The condition name/field",
						},
						"operator": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The comparison operator",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The condition value",
						},
						"list_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Optional list ID for condition",
						},
					},
				},
			},
			"actions": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The actions to execute when workflow is triggered",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The action name",
						},
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The action value",
						},
						"webhook_url": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Webhook URL for the action",
						},
						"send_to_original_recipients": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Send to original recipients",
						},
						"send_to_sender": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Send to sender",
						},
						"send_to_owner": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Send to owner",
						},
						"launch_new_thread": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Launch new thread",
						},
						"subject": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Subject for the action",
						},
						"message_text": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Message text for the action",
						},
						"users": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Users to include in the action",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"groups": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Groups to include in the action",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"is_used": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the workflow is currently being used",
			},
			"is_bidirection": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the workflow is bidirectional",
			},
			"debug_request_json": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "JSON representation of the request for debugging",
			},
		},
	}
}

// Helper functions for expanding and flattening nested structures

func expandWorkflowConditions(conditionsData []interface{}) []WorkflowCondition {
	if len(conditionsData) == 0 {
		return nil
	}

	conditions := make([]WorkflowCondition, len(conditionsData))
	for i, conditionData := range conditionsData {
		conditionMap := conditionData.(map[string]interface{})
		conditions[i] = WorkflowCondition{
			Type:     conditionMap["type"].(string),
			Match:    conditionMap["match"].(string),
			Name:     conditionMap["name"].(string),
			Operator: conditionMap["operator"].(string),
			Value:    conditionMap["value"].(string),
		}
		if v, ok := conditionMap["list_id"]; ok && v.(string) != "" {
			conditions[i].ListID = v.(string)
		}
	}
	return conditions
}

func flattenWorkflowConditions(conditions []WorkflowCondition) []map[string]interface{} {
	if len(conditions) == 0 {
		return nil
	}

	result := make([]map[string]interface{}, len(conditions))
	for i, condition := range conditions {
		result[i] = map[string]interface{}{
			"type":     condition.Type,
			"match":    condition.Match,
			"name":     condition.Name,
			"operator": condition.Operator,
			"value":    condition.Value,
			"list_id":  condition.ListID,
		}
	}
	return result
}

func expandWorkflowActions(actionsData []interface{}) []WorkflowAction {
	if len(actionsData) == 0 {
		return nil
	}

	actions := make([]WorkflowAction, len(actionsData))
	for i, actionData := range actionsData {
		actionMap := actionData.(map[string]interface{})
		actions[i] = WorkflowAction{
			Name: actionMap["name"].(string),
		}
		
		if v, ok := actionMap["value"]; ok && v.(string) != "" {
			actions[i].Value = v.(string)
		}
		if v, ok := actionMap["webhook_url"]; ok && v.(string) != "" {
			actions[i].WebhookURL = v.(string)
		}
		if v, ok := actionMap["send_to_original_recipients"]; ok {
			actions[i].SendToOriginalRecipients = v.(bool)
		}
		if v, ok := actionMap["send_to_sender"]; ok {
			actions[i].SendToSender = v.(bool)
		}
		if v, ok := actionMap["send_to_owner"]; ok {
			actions[i].SendToOwner = v.(bool)
		}
		if v, ok := actionMap["launch_new_thread"]; ok {
			actions[i].LaunchNewThread = v.(bool)
		}
		if v, ok := actionMap["subject"]; ok && v.(string) != "" {
			actions[i].Subject = v.(string)
		}
		if v, ok := actionMap["message_text"]; ok && v.(string) != "" {
			actions[i].MessageText = v.(string)
		}
		
		// Handle users list
		if v, ok := actionMap["users"]; ok {
			usersList := v.([]interface{})
			users := make([]string, len(usersList))
			for j, user := range usersList {
				users[j] = user.(string)
			}
			actions[i].Users = users
		}
		
		// Handle groups list
		if v, ok := actionMap["groups"]; ok {
			groupsList := v.([]interface{})
			groups := make([]string, len(groupsList))
			for j, group := range groupsList {
				groups[j] = group.(string)
			}
			actions[i].Groups = groups
		}
	}
	return actions
}

func flattenWorkflowActions(actions []WorkflowAction) []map[string]interface{} {
	if len(actions) == 0 {
		return nil
	}

	result := make([]map[string]interface{}, len(actions))
	for i, action := range actions {
		result[i] = map[string]interface{}{
			"name":                       action.Name,
			"value":                      action.Value,
			"webhook_url":                action.WebhookURL,
			"send_to_original_recipients": action.SendToOriginalRecipients,
			"send_to_sender":             action.SendToSender,
			"send_to_owner":              action.SendToOwner,
			"launch_new_thread":          action.LaunchNewThread,
			"subject":                    action.Subject,
			"message_text":               action.MessageText,
			"users":                      action.Users,
			"groups":                     action.Groups,
		}
	}
	return result
}

func resourceWorkflowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	workflow := Workflow{
		WorkflowName:       d.Get("workflow_name").(string),
		WorkflowType:       d.Get("workflow_type").(string),
		Enabled:            d.Get("enabled").(bool),
		AlertType:          d.Get("alert_type").(string),
		Scheduled:          d.Get("scheduled").(bool),
		RecurrenceInterval: d.Get("recurrence_interval").(int),
	}

	// Handle conditions
	if v, ok := d.GetOk("conditions"); ok {
		workflow.Conditions = expandWorkflowConditions(v.([]interface{}))
	}

	// Handle actions
	if v, ok := d.GetOk("actions"); ok {
		workflow.Actions = expandWorkflowActions(v.([]interface{}))
	}

	// Store the request JSON for debugging
	requestJSON, _ := json.Marshal(workflow)
	d.Set("debug_request_json", string(requestJSON))

	var createdWorkflow Workflow
	err := client.post(ctx, "/api/v2/workflows", workflow, &createdWorkflow)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create workflow: %w", err))
	}

	d.SetId(strconv.Itoa(createdWorkflow.WorkflowID))

	return resourceWorkflowRead(ctx, d, meta)
}

func resourceWorkflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	workflowID := d.Id()
	var workflow Workflow
	err := client.get(ctx, fmt.Sprintf("/api/v2/workflows/%s", workflowID), &workflow)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read workflow: %w", err))
	}

	d.Set("workflow_id", workflow.WorkflowID)
	d.Set("workflow_name", workflow.WorkflowName)
	d.Set("workflow_type", workflow.WorkflowType)
	d.Set("enabled", workflow.Enabled)
	d.Set("alert_type", workflow.AlertType)
	d.Set("scheduled", workflow.Scheduled)
	d.Set("recurrence_interval", workflow.RecurrenceInterval)
	d.Set("is_used", workflow.IsUsed)
	d.Set("is_bidirection", workflow.IsBidirection)

	// Set nested objects
	if workflow.Conditions != nil {
		d.Set("conditions", flattenWorkflowConditions(workflow.Conditions))
	}

	if workflow.Actions != nil {
		d.Set("actions", flattenWorkflowActions(workflow.Actions))
	}

	return nil
}

func resourceWorkflowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	workflowID := d.Id()
	workflow := Workflow{
		WorkflowID:         d.Get("workflow_id").(int),
		WorkflowName:       d.Get("workflow_name").(string),
		WorkflowType:       d.Get("workflow_type").(string),
		Enabled:            d.Get("enabled").(bool),
		AlertType:          d.Get("alert_type").(string),
		Scheduled:          d.Get("scheduled").(bool),
		RecurrenceInterval: d.Get("recurrence_interval").(int),
	}

	// Handle conditions
	if v, ok := d.GetOk("conditions"); ok {
		workflow.Conditions = expandWorkflowConditions(v.([]interface{}))
	}

	// Handle actions
	if v, ok := d.GetOk("actions"); ok {
		workflow.Actions = expandWorkflowActions(v.([]interface{}))
	}

	// Store the request JSON for debugging
	requestJSON, _ := json.Marshal(workflow)
	d.Set("debug_request_json", string(requestJSON))

	// AlertOps API returns 204 No Content for updates, so we don't expect a response body
	err := client.put(ctx, fmt.Sprintf("/api/v2/workflows/%s", workflowID), workflow, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update workflow: %w", err))
	}

	return resourceWorkflowRead(ctx, d, meta)
}

func resourceWorkflowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	workflowID := d.Id()
	err := client.delete(ctx, fmt.Sprintf("/api/v2/workflows/%s", workflowID))
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete workflow: %w", err))
	}

	d.SetId("")
	return nil
} 