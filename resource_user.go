package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "User ID",
			},
			"debug_request_json": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DEBUG: The exact JSON that will be sent to AlertOps API",
			},
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username for the user",
			},
			"first_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "First name of the user",
			},
			"last_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Last name of the user",
			},
			"locale": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "en-US",
				Description: "Locale of the user",
			},
			"time_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Time zone of the user",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Standard",
				Description: "Type of the user (Standard, etc.)",
			},
			"external_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "External ID for the user",
			},
			"last_login_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last login date",
			},
			"contact_methods": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Contact methods for the user",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"contact_method_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the contact method (Email-Official, Phone-Official, SMS-Official, Email-Official-SMS Gateway, Email-Personal, Email-Personal-SMS Gateway, Phone-Official-Mobile, Phone-Personal, Phone-Personal-Mobile, SMS-Personal)",
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								v := val.(string)
								for _, valid := range ValidContactMethodTypes {
									if v == valid {
										return
									}
								}
								errs = append(errs, fmt.Errorf("contact_method_name must be one of: %v", ValidContactMethodTypes))
								return
							},
						},
						"email": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Email contact",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"email_address": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Email address",
									},
								},
							},
						},
						"phone": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Phone contact",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"country_code": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Country code",
									},
									"phone_number": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Phone number",
									},
									"extension": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Extension",
									},
								},
							},
						},
						"sms": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "SMS contact",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"country_code": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Country code",
									},
									"phone_number": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Phone number",
									},
								},
							},
						},
						"wait_time_in_mins": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Wait time in minutes",
						},
						"repeat": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to repeat notifications",
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
						"notification_time24x7": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "24x7 notification",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Whether contact method is enabled",
						},
						"sequence": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Sequence order",
						},
					},
				},
			},
			"roles": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "User roles",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	user := UserCreateRequest{
		UserName:  d.Get("user_name").(string),
		FirstName: d.Get("first_name").(string),
		LastName:  d.Get("last_name").(string),
		Locale:    d.Get("locale").(string),
		TimeZone:  d.Get("time_zone").(string),
		Type:      d.Get("type").(string),
	}

	if externalID, ok := d.GetOk("external_id"); ok {
		user.ExternalID = externalID.(string)
	}

	if contactMethods := d.Get("contact_methods").([]interface{}); len(contactMethods) > 0 {
		user.ContactMethods = expandContactMethods(contactMethods)
	}

	if roles := d.Get("roles").([]interface{}); len(roles) > 0 {
		user.Roles = expandStringSlice(roles)
	}

	// Generate the exact JSON that will be sent to AlertOps
	requestJSON := "failed to marshal request"
	if jsonBytes, jsonErr := json.Marshal(user); jsonErr == nil {
		requestJSON = string(jsonBytes)
		log.Printf("=== DEBUG REQUEST JSON ===\n%s", requestJSON)
	}

	// Set the debug field so it shows in terraform plan/apply  
	d.Set("debug_request_json", requestJSON)
	
	// Force log what we're about to send
	log.Printf("🔥🔥🔥 ABOUT TO SEND TO ALERTOPS 🔥🔥🔥")
	log.Printf("URL: https://api.alertops.com/api/v2/users")
	log.Printf("HEADERS: api-key: [REDACTED]")
	log.Printf("BODY: %s", requestJSON)

	var result User
	err := client.post(ctx, "/api/v2/users", user, &result)
	if err != nil {
		log.Printf("=== DEBUG ERROR ===\n%v", err)
		return diag.Errorf("🚨🚨🚨 NEW BINARY IS RUNNING - FIXED AUTHENTICATION! 🚨🚨🚨\nREQUEST JSON:\n%s\nERROR: %v", requestJSON, err)
	}

	d.SetId(strconv.Itoa(result.UserID))
	return resourceUserRead(ctx, d, meta)
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	var user User
	err := client.get(ctx, fmt.Sprintf("/api/v2/users/%s", d.Id()), &user)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read user: %w", err))
	}

	d.Set("user_id", user.UserID)
	d.Set("user_name", user.UserName)
	d.Set("first_name", user.FirstName)
	d.Set("last_name", user.LastName)
	d.Set("locale", user.Locale)
	d.Set("time_zone", user.TimeZone)
	d.Set("type", user.Type)
	d.Set("external_id", user.ExternalID)
	d.Set("last_login_date", user.LastLoginDate)

	if len(user.ContactMethods) > 0 {
		d.Set("contact_methods", flattenContactMethods(user.ContactMethods))
	}

	if len(user.Roles) > 0 {
		d.Set("roles", user.Roles)
	}

	return nil
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	user := UserUpdateRequest{
		UserName:  d.Get("user_name").(string),
		FirstName: d.Get("first_name").(string),
		LastName:  d.Get("last_name").(string),
		Locale:    d.Get("locale").(string),
		TimeZone:  d.Get("time_zone").(string),
		Type:      d.Get("type").(string),
	}

	if externalID, ok := d.GetOk("external_id"); ok {
		user.ExternalID = externalID.(string)
	}

	if contactMethods := d.Get("contact_methods").([]interface{}); len(contactMethods) > 0 {
		user.ContactMethods = expandContactMethods(contactMethods)
	}

	if roles := d.Get("roles").([]interface{}); len(roles) > 0 {
		user.Roles = expandStringSlice(roles)
	}

	err := client.put(ctx, fmt.Sprintf("/api/v2/users/%s", d.Id()), user, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update user: %w", err))
	}

	return resourceUserRead(ctx, d, meta)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	err := client.delete(ctx, fmt.Sprintf("/api/v2/users/%s", d.Id()))
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete user: %w", err))
	}

	d.SetId("")
	return nil
}

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "User ID",
			},
			"user_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Username for the user",
			},
			"first_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "First name of the user",
			},
			"last_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last name of the user",
			},
			"locale": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Locale of the user",
			},
			"time_zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time zone of the user",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of the user",
			},
			"external_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "External ID for the user",
			},
			"last_login_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last login date",
			},
			"contact_methods": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Contact methods for the user",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"contact_method_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the contact method (Email-Official, Phone-Official, SMS-Official, etc.)",
						},
						"email": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Email contact",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"email_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Email address",
									},
								},
							},
						},
						"phone": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Phone contact",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"country_code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Country code",
									},
									"phone_number": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Phone number",
									},
									"extension": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Extension",
									},
								},
							},
						},
						"sms": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "SMS contact",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"country_code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Country code",
									},
									"phone_number": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Phone number",
									},
								},
							},
						},
						"wait_time_in_mins": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Wait time in minutes",
						},
						"repeat": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to repeat notifications",
						},
						"repeat_times": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of times to repeat",
						},
						"repeat_minutes": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Minutes between repeats",
						},
						"notification_time24x7": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "24x7 notification",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether contact method is enabled",
						},
						"sequence": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Sequence order",
						},
					},
				},
			},
			"roles": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User roles",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	// If user_id is provided, fetch by ID
	if userID, ok := d.GetOk("user_id"); ok {
		var user User
		err := client.get(ctx, fmt.Sprintf("/api/v2/users/%d", userID.(int)), &user)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to read user: %w", err))
		}

		d.SetId(strconv.Itoa(user.UserID))
		d.Set("user_id", user.UserID)
		d.Set("user_name", user.UserName)
		d.Set("first_name", user.FirstName)
		d.Set("last_name", user.LastName)
		d.Set("locale", user.Locale)
		d.Set("time_zone", user.TimeZone)
		d.Set("type", user.Type)
		d.Set("external_id", user.ExternalID)
		d.Set("last_login_date", user.LastLoginDate)
		
		if len(user.ContactMethods) > 0 {
			d.Set("contact_methods", flattenContactMethods(user.ContactMethods))
		}
		if len(user.Roles) > 0 {
			d.Set("roles", user.Roles)
		}

		return nil
	}

	// Otherwise, search by user_name
	var listResponse UserListResponse
	err := client.get(ctx, "/api/v2/users", &listResponse)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to list users: %w", err))
	}

	userName := d.Get("user_name").(string)

	for _, user := range listResponse.Users {
		if userName != "" && user.UserName == userName {
			d.SetId(strconv.Itoa(user.UserID))
			d.Set("user_id", user.UserID)
			d.Set("user_name", user.UserName)
			d.Set("first_name", user.FirstName)
			d.Set("last_name", user.LastName)
			d.Set("locale", user.Locale)
			d.Set("time_zone", user.TimeZone)
			d.Set("type", user.Type)
			d.Set("external_id", user.ExternalID)
			d.Set("last_login_date", user.LastLoginDate)
			
			if len(user.ContactMethods) > 0 {
				d.Set("contact_methods", flattenContactMethods(user.ContactMethods))
			}
			if len(user.Roles) > 0 {
				d.Set("roles", user.Roles)
			}

			return nil
		}
	}

	return diag.FromErr(fmt.Errorf("user not found"))
}

// Helper functions for contact methods
func expandContactMethods(methods []interface{}) []ContactMethod {
	result := make([]ContactMethod, len(methods))
	
	for i, method := range methods {
		methodMap := method.(map[string]interface{})
		cm := ContactMethod{
			ContactMethodName: methodMap["contact_method_name"].(string),
		}

		if emailList, ok := methodMap["email"].([]interface{}); ok && len(emailList) > 0 {
			emailMap := emailList[0].(map[string]interface{})
			cm.Email = &EmailContact{
				EmailAddress: emailMap["email_address"].(string),
			}
		}

		if phoneList, ok := methodMap["phone"].([]interface{}); ok && len(phoneList) > 0 {
			phoneMap := phoneList[0].(map[string]interface{})
			phone := &PhoneContact{
				CountryCode: phoneMap["country_code"].(string),
				PhoneNumber: phoneMap["phone_number"].(string),
			}
			if ext, ok := phoneMap["extension"].(string); ok {
				phone.Extension = ext
			}
			cm.Phone = phone
		}

		if smsList, ok := methodMap["sms"].([]interface{}); ok && len(smsList) > 0 {
			smsMap := smsList[0].(map[string]interface{})
			cm.SMS = &SMSContact{
				CountryCode: smsMap["country_code"].(string),
				PhoneNumber: smsMap["phone_number"].(string),
			}
		}

		if wait, ok := methodMap["wait_time_in_mins"].(int); ok {
			cm.WaitTimeInMins = wait
		}
		if repeat, ok := methodMap["repeat"].(bool); ok {
			cm.Repeat = repeat
		}
		if times, ok := methodMap["repeat_times"].(int); ok {
			cm.RepeatTimes = times
		}
		if minutes, ok := methodMap["repeat_minutes"].(int); ok {
			cm.RepeatMinutes = minutes
		}
		if notif24x7, ok := methodMap["notification_time24x7"].(bool); ok {
			cm.NotificationTime24x7 = notif24x7
		}
		if enabled, ok := methodMap["enabled"].(bool); ok {
			cm.Enabled = enabled
		}
		if seq, ok := methodMap["sequence"].(int); ok {
			cm.Sequence = seq
		}

		result[i] = cm
	}

	return result
}

func flattenContactMethods(methods []ContactMethod) []interface{} {
	result := make([]interface{}, len(methods))

	for i, method := range methods {
		m := map[string]interface{}{
			"contact_method_name":    method.ContactMethodName,
			"wait_time_in_mins":      method.WaitTimeInMins,
			"repeat":                 method.Repeat,
			"repeat_times":           method.RepeatTimes,
			"repeat_minutes":         method.RepeatMinutes,
			"notification_time24x7":  method.NotificationTime24x7,
			"enabled":                method.Enabled,
			"sequence":               method.Sequence,
		}

		if method.Email != nil {
			m["email"] = []interface{}{
				map[string]interface{}{
					"email_address": method.Email.EmailAddress,
				},
			}
		}

		if method.Phone != nil {
			phone := map[string]interface{}{
				"country_code": method.Phone.CountryCode,
				"phone_number": method.Phone.PhoneNumber,
			}
			if method.Phone.Extension != "" {
				phone["extension"] = method.Phone.Extension
			}
			m["phone"] = []interface{}{phone}
		}

		if method.SMS != nil {
			m["sms"] = []interface{}{
				map[string]interface{}{
					"country_code": method.SMS.CountryCode,
					"phone_number": method.SMS.PhoneNumber,
				},
			}
		}

		result[i] = m
	}

	return result
}

func expandStringSlice(slice []interface{}) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = v.(string)
	}
	return result
} 