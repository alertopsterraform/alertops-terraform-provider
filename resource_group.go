package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Group models are defined in models.go

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGroupCreate,
		ReadContext:   resourceGroupRead,
		UpdateContext: resourceGroupUpdate,
		DeleteContext: resourceGroupDelete,

		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Unique identifier for the group",
			},
			"group_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the group",
			},
			"dynamic": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether the group is dynamic",
			},
			"description": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of description strings for the group",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"members": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of group members (users or other groups)",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Type of member: 'User' or 'Group'",
						},
						"member": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Username or group name",
						},
						"sequence": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Sequence order",
						},
						"roles": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Roles for this member (e.g., Primary, Manager)",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"contact_methods": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Contact methods for the group",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"contact_method_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the contact method",
						},
						"email_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Email address (for email-based methods)",
						},
						"country_code": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Country code (for phone/SMS methods)",
						},
						"phone_number": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Phone number (for phone/SMS methods)",
						},
						"extension": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Phone extension",
						},
						"url": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "URL (for webhook/Slack methods)",
						},
						"get_alert_update": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to get alert updates",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Whether the contact method is enabled",
						},
						"sequence": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Sequence order",
						},
					},
				},
			},
			"topics": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of topics associated with the group",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"attributes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Custom attributes for the group",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the attribute",
						},
						"attribute_value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Value of the attribute",
						},
					},
				},
			},
			"debug_request_json": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "JSON representation of the request sent to AlertOps API (for debugging)",
			},
		},
	}
}

func resourceGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	group := Group{
		GroupName: d.Get("group_name").(string),
		Dynamic:   d.Get("dynamic").(bool),
	}

	// Handle description array
	if v, ok := d.GetOk("description"); ok {
		descList := v.([]interface{})
		descriptions := make([]string, len(descList))
		for i, desc := range descList {
			descriptions[i] = desc.(string)
		}
		group.Description = descriptions
	}

	// Handle members
	if v, ok := d.GetOk("members"); ok {
		group.Members = expandGroupMembers(v.([]interface{}))
	}

	// Handle contact methods
	if v, ok := d.GetOk("contact_methods"); ok {
		group.ContactMethods = expandGroupContactMethods(v.([]interface{}))
	}

	// Handle topics
	if v, ok := d.GetOk("topics"); ok {
		topicsList := v.([]interface{})
		topics := make([]string, len(topicsList))
		for i, topic := range topicsList {
			topics[i] = topic.(string)
		}
		group.Topics = topics
	}

	// Handle attributes
	if v, ok := d.GetOk("attributes"); ok {
		group.Attributes = expandGroupAttributes(v.([]interface{}))
	}

	// Store the request JSON for debugging
	requestJSON, _ := json.Marshal(group)
	d.Set("debug_request_json", string(requestJSON))

	var createdGroup Group
	err := client.post(ctx, "/api/v2/groups", group, &createdGroup)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create group: %w", err))
	}

	d.SetId(strconv.Itoa(createdGroup.GroupID))
	d.Set("group_id", createdGroup.GroupID)

	return resourceGroupRead(ctx, d, meta)
}

func resourceGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	groupID := d.Id()
	var group Group
	err := client.get(ctx, fmt.Sprintf("/api/v2/groups/%s", groupID), &group)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read group: %w", err))
	}

	d.Set("group_id", group.GroupID)
	d.Set("group_name", group.GroupName)
	d.Set("dynamic", group.Dynamic)
	d.Set("description", group.Description)
	d.Set("topics", group.Topics)

	if len(group.Members) > 0 {
		d.Set("members", flattenGroupMembers(group.Members))
	}

	if len(group.ContactMethods) > 0 {
		d.Set("contact_methods", flattenGroupContactMethods(group.ContactMethods))
	}

	if len(group.Attributes) > 0 {
		d.Set("attributes", flattenGroupAttributes(group.Attributes))
	}

	return nil
}

func resourceGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	groupID := d.Id()
	group := Group{
		GroupID:   d.Get("group_id").(int),
		GroupName: d.Get("group_name").(string),
		Dynamic:   d.Get("dynamic").(bool),
	}

	// Handle description array
	if v, ok := d.GetOk("description"); ok {
		descList := v.([]interface{})
		descriptions := make([]string, len(descList))
		for i, desc := range descList {
			descriptions[i] = desc.(string)
		}
		group.Description = descriptions
	}

	// Handle members
	if v, ok := d.GetOk("members"); ok {
		group.Members = expandGroupMembers(v.([]interface{}))
	}

	// Handle contact methods
	if v, ok := d.GetOk("contact_methods"); ok {
		group.ContactMethods = expandGroupContactMethods(v.([]interface{}))
	}

	// Handle topics
	if v, ok := d.GetOk("topics"); ok {
		topicsList := v.([]interface{})
		topics := make([]string, len(topicsList))
		for i, topic := range topicsList {
			topics[i] = topic.(string)
		}
		group.Topics = topics
	}

	// Handle attributes
	if v, ok := d.GetOk("attributes"); ok {
		group.Attributes = expandGroupAttributes(v.([]interface{}))
	}

	// Store the request JSON for debugging
	requestJSON, _ := json.Marshal(group)
	d.Set("debug_request_json", string(requestJSON))

	// AlertOps API returns 204 No Content for updates, so we don't expect a response body
	err := client.put(ctx, fmt.Sprintf("/api/v2/groups/%s", groupID), group, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update group: %w", err))
	}

	return resourceGroupRead(ctx, d, meta)
}

func resourceGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	groupID := d.Id()
	err := client.delete(ctx, fmt.Sprintf("/api/v2/groups/%s", groupID))
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete group: %w", err))
	}

	d.SetId("")
	return nil
}

// Helper functions for expanding and flattening complex structures

func expandGroupMembers(members []interface{}) []GroupMember {
	if len(members) == 0 {
		return nil
	}

	result := make([]GroupMember, len(members))
	for i, memberData := range members {
		member := memberData.(map[string]interface{})
		
		groupMember := GroupMember{
			MemberType: member["member_type"].(string),
			Member:     member["member"].(string),
			Sequence:   member["sequence"].(int),
		}

		if roles, ok := member["roles"]; ok && roles != nil {
			rolesList := roles.([]interface{})
			groupMember.Roles = make([]string, len(rolesList))
			for j, role := range rolesList {
				groupMember.Roles[j] = role.(string)
			}
		}

		result[i] = groupMember
	}
	return result
}

func flattenGroupMembers(members []GroupMember) []interface{} {
	if len(members) == 0 {
		return nil
	}

	result := make([]interface{}, len(members))
	for i, member := range members {
		memberMap := map[string]interface{}{
			"member_type": member.MemberType,
			"member":      member.Member,
			"sequence":    member.Sequence,
		}

		if len(member.Roles) > 0 {
			roles := make([]interface{}, len(member.Roles))
			for j, role := range member.Roles {
				roles[j] = role
			}
			memberMap["roles"] = roles
		}

		result[i] = memberMap
	}
	return result
}

func expandGroupContactMethods(contactMethods []interface{}) []GroupContactMethod {
	if len(contactMethods) == 0 {
		return nil
	}

	result := make([]GroupContactMethod, len(contactMethods))
	for i, contactMethodData := range contactMethods {
		cm := contactMethodData.(map[string]interface{})
		
		contactMethod := GroupContactMethod{
			ContactMethodName: cm["contact_method_name"].(string),
			Sequence:          cm["sequence"].(int),
		}

		if email, ok := cm["email_address"]; ok && email != nil {
			contactMethod.EmailAddress = email.(string)
		}
		if country, ok := cm["country_code"]; ok && country != nil {
			contactMethod.CountryCode = country.(string)
		}
		if phone, ok := cm["phone_number"]; ok && phone != nil {
			contactMethod.PhoneNumber = phone.(string)
		}
		if ext, ok := cm["extension"]; ok && ext != nil {
			contactMethod.Extension = ext.(string)
		}
		if url, ok := cm["url"]; ok && url != nil {
			contactMethod.URL = url.(string)
		}
		if update, ok := cm["get_alert_update"]; ok && update != nil {
			contactMethod.GetAlertUpdate = update.(bool)
		}
		if enabled, ok := cm["enabled"]; ok && enabled != nil {
			contactMethod.Enabled = enabled.(bool)
		}

		result[i] = contactMethod
	}
	return result
}

func flattenGroupContactMethods(contactMethods []GroupContactMethod) []interface{} {
	if len(contactMethods) == 0 {
		return nil
	}

	result := make([]interface{}, len(contactMethods))
	for i, cm := range contactMethods {
		contactMethodMap := map[string]interface{}{
			"contact_method_name": cm.ContactMethodName,
			"sequence":            cm.Sequence,
			"enabled":             cm.Enabled,
		}

		if cm.EmailAddress != "" {
			contactMethodMap["email_address"] = cm.EmailAddress
		}
		if cm.CountryCode != "" {
			contactMethodMap["country_code"] = cm.CountryCode
		}
		if cm.PhoneNumber != "" {
			contactMethodMap["phone_number"] = cm.PhoneNumber
		}
		if cm.Extension != "" {
			contactMethodMap["extension"] = cm.Extension
		}
		if cm.URL != "" {
			contactMethodMap["url"] = cm.URL
		}
		if cm.GetAlertUpdate {
			contactMethodMap["get_alert_update"] = cm.GetAlertUpdate
		}

		result[i] = contactMethodMap
	}
	return result
}

func expandGroupAttributes(attributes []interface{}) []GroupAttribute {
	if len(attributes) == 0 {
		return nil
	}

	result := make([]GroupAttribute, len(attributes))
	for i, attrData := range attributes {
		attr := attrData.(map[string]interface{})
		
		result[i] = GroupAttribute{
			AttributeName:  attr["attribute_name"].(string),
			AttributeValue: attr["attribute_value"].(string),
		}
	}
	return result
}

func flattenGroupAttributes(attributes []GroupAttribute) []interface{} {
	if len(attributes) == 0 {
		return nil
	}

	result := make([]interface{}, len(attributes))
	for i, attr := range attributes {
		result[i] = map[string]interface{}{
			"attribute_name":  attr.AttributeName,
			"attribute_value": attr.AttributeValue,
		}
	}
	return result
} 