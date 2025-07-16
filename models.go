package main

// User represents an AlertOps user
type User struct {
	UserID        int             `json:"user_id,omitempty"`
	UserName      string          `json:"user_name"`
	FirstName     string          `json:"first_name"`
	LastName      string          `json:"last_name"`
	Locale        string          `json:"locale,omitempty"`
	TimeZone      string          `json:"time_zone,omitempty"`
	Type          string          `json:"type,omitempty"`
	ExternalID    string          `json:"external_id,omitempty"`
	LastLoginDate string          `json:"last_login_date,omitempty"`
	ContactMethods []ContactMethod `json:"contact_methods,omitempty"`
	Roles         []string        `json:"roles,omitempty"`
}

// Group represents an AlertOps group structure matching the actual API
type Group struct {
	GroupID        int                    `json:"group_id,omitempty"`
	GroupName      string                 `json:"group_name"`
	Dynamic        bool                   `json:"dynamic,omitempty"`
	Description    []string               `json:"description,omitempty"`
	Members        []GroupMember          `json:"members,omitempty"`
	ContactMethods []GroupContactMethod   `json:"contact_methods,omitempty"`
	Topics         []string               `json:"topics,omitempty"`
	Attributes     []GroupAttribute       `json:"attributes,omitempty"`
}

// GroupMember represents a member of a group (user or another group)
type GroupMember struct {
	MemberType string   `json:"member_type"` // "User" or "Group"
	Member     string   `json:"member"`      // username or group name
	Sequence   int      `json:"sequence"`
	Roles      []string `json:"roles,omitempty"` // e.g., ["Primary", "Manager"]
}

// GroupContactMethod represents contact methods for groups
type GroupContactMethod struct {
	ContactMethodName string `json:"contact_method_name"`
	EmailAddress      string `json:"email_address,omitempty"`
	CountryCode       string `json:"country_code,omitempty"`
	PhoneNumber       string `json:"phone_number,omitempty"`
	Extension         string `json:"extension,omitempty"`
	URL               string `json:"url,omitempty"`
	GetAlertUpdate    bool   `json:"get_alert_update,omitempty"`
	Enabled           bool   `json:"enabled,omitempty"`
	Sequence          int    `json:"sequence"`
}

// GroupAttribute represents custom attributes for groups
type GroupAttribute struct {
	AttributeName  string `json:"attribute_name"`
	AttributeValue string `json:"attribute_value"`
}

// GroupsResponse represents the API response for listing groups
type GroupsResponse struct {
	Limit  int     `json:"limit"`
	Offset int     `json:"offset"`
	Groups []Group `json:"groups"`
}

// ContactMethod represents a user's contact method
type ContactMethod struct {
	ContactMethodName       string             `json:"contact_method_name"`
	Email                   *EmailContact      `json:"email,omitempty"`
	Phone                   *PhoneContact      `json:"phone,omitempty"`
	SMS                     *SMSContact        `json:"sms,omitempty"`
	Gateway                 *GatewayContact    `json:"gateway,omitempty"`
	SlackDM                 *SlackDMContact    `json:"slack_dm,omitempty"`
	WaitTimeInMins          int                `json:"wait_time_in_mins,omitempty"`
	Repeat                  bool               `json:"repeat,omitempty"`
	RepeatTimes             int                `json:"repeat_times,omitempty"`
	RepeatMinutes           int                `json:"repeat_minutes,omitempty"`
	NotificationTime24x7    bool               `json:"notification_time24x7,omitempty"`
	NotificationTimes       []NotificationTime `json:"notification_times,omitempty"`
	Enabled                 bool               `json:"enabled,omitempty"`
	Sequence                int                `json:"sequence,omitempty"`
}

// Valid contact method types in AlertOps
var ValidContactMethodTypes = []string{
	"Email-Official",
	"Phone-Official",
	"SMS-Official",
	"Email-Official-SMS Gateway",
	"Email-Personal",
	"Email-Personal-SMS Gateway",
	"Phone-Official-Mobile",
	"Phone-Personal",
	"Phone-Personal-Mobile",
	"SMS-Personal",
}

// EmailContact represents email contact information
type EmailContact struct {
	EmailAddress string `json:"email_address"`
}

// PhoneContact represents phone contact information
type PhoneContact struct {
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
	Extension   string `json:"extension,omitempty"`
}

// SMSContact represents SMS contact information
type SMSContact struct {
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
}

// GatewayContact represents gateway contact information
type GatewayContact struct {
	Provider string `json:"provider"`
	Address  string `json:"address"`
}

// SlackDMContact represents Slack DM contact information
type SlackDMContact struct {
	MemberID string `json:"member_id"`
}

// NotificationTime represents notification time settings
type NotificationTime struct {
	NotificationTimeID int    `json:"notification_time_id,omitempty"`
	Name              string `json:"name"`
	Sunday            bool   `json:"sunday"`
	Monday            bool   `json:"monday"`
	Tuesday           bool   `json:"tuesday"`
	Wednesday         bool   `json:"wednesday"`
	Thursday          bool   `json:"thursday"`
	Friday            bool   `json:"friday"`
	Saturday          bool   `json:"saturday"`
	StartHour         int    `json:"start_hour"`
	StartMinute       int    `json:"start_minute"`
	EndHour           int    `json:"end_hour"`
	EndMinute         int    `json:"end_minute"`
}

// UserCreateRequest represents the request body for creating a user
type UserCreateRequest struct {
	UserName       string          `json:"user_name"`
	FirstName      string          `json:"first_name"`
	LastName       string          `json:"last_name"`
	Locale         string          `json:"locale,omitempty"`
	TimeZone       string          `json:"time_zone,omitempty"`
	Type           string          `json:"type,omitempty"`
	ExternalID     string          `json:"external_id,omitempty"`
	ContactMethods []ContactMethod `json:"contact_methods,omitempty"`
	Roles          []string        `json:"roles,omitempty"`
}

// UserUpdateRequest represents the request body for updating a user (same as create)
type UserUpdateRequest = UserCreateRequest

// UserListResponse represents the response for listing users
type UserListResponse struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Users  []User `json:"users"`
}

// Schedule represents a schedule in AlertOps
type Schedule struct {
	ScheduleID               int              `json:"schedule_id,omitempty"`
	Group                    string           `json:"group"`
	ScheduleName             string           `json:"schedule_name"`
	ScheduleType             string           `json:"schedule_type"`
	Continuous               bool             `json:"continuous"`
	TimeZone                 string           `json:"time_zone"`
	Color                    string           `json:"color,omitempty"`
	StartDate                *ScheduleDate    `json:"start_date,omitempty"`
	EndDate                  *ScheduleDate    `json:"end_date,omitempty"`
	StartWeekday             string           `json:"start_weekday,omitempty"`
	EndWeekday               string           `json:"end_weekday,omitempty"`
	ScheduleWeekdays         *ScheduleWeekdays `json:"schedule_weekdays,omitempty"`
	RotateFrequency          string           `json:"rotate_frequency,omitempty"`
	RotateDaily              *RotateDaily     `json:"rotate_daily,omitempty"`
	RotateWeekly             *RotateWeekly    `json:"rotate_weekly,omitempty"`
	RotateMonthly            *RotateMonthly   `json:"rotate_monthly,omitempty"`
	RepeatSchedule           *RepeatSchedule  `json:"repeat_schedule,omitempty"`
	IncludeAllUsersInGroup   bool             `json:"include_all_users_in_group"`
	Users                    []ScheduleUser   `json:"users,omitempty"`
	Enabled                  bool             `json:"enabled"`
	IsHolidayNotify          bool             `json:"is_holiday_notify"`
}

// ScheduleDate represents a date with hour and minute
type ScheduleDate struct {
	Date   string `json:"date"`
	Hour   int    `json:"hour"`
	Minute int    `json:"minute"`
}

// ScheduleTime represents a time with hour and minute
type ScheduleTime struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

// ScheduleWeekdays represents which weekdays are included
type ScheduleWeekdays struct {
	Sun bool `json:"sun"`
	Mon bool `json:"mon"`
	Tue bool `json:"tue"`
	Wed bool `json:"wed"`
	Thu bool `json:"thu"`
	Fri bool `json:"fri"`
	Sat bool `json:"sat"`
}

// RotateDaily represents daily rotation configuration
type RotateDaily struct {
	RotateXUsers  int           `json:"rotate_x_users"`
	RotateAtTime  *ScheduleTime `json:"rotate_at_time"`
	EveryXDays    int           `json:"every_x_days"`
}

// RotateWeekly represents weekly rotation configuration
type RotateWeekly struct {
	RotateXUsers        int           `json:"rotate_x_users"`
	RotateAtTime        *ScheduleTime `json:"rotate_at_time"`
	EveryXWeeks         int           `json:"every_x_weeks"`
	RotateAtDayOfWeek   string        `json:"rotate_at_day_of_week"`
}

// RotateMonthly represents monthly rotation configuration
type RotateMonthly struct {
	RotateXUsers  int           `json:"rotate_x_users"`
	RotateAtTime  *ScheduleTime `json:"rotate_at_time"`
	EveryXMonths  int           `json:"every_x_months"`
}

// RepeatSchedule represents schedule repetition configuration
type RepeatSchedule struct {
	EveryXWeeks      int    `json:"every_x_weeks"`
	RepeatUntilDate  string `json:"repeat_until_date"`
}

// ScheduleUser represents a user assignment in a schedule
type ScheduleUser struct {
	User string `json:"user"`
	Role string `json:"role"`
}

// ScheduleListResponse represents the response for listing schedules
type ScheduleListResponse struct {
	Limit     int        `json:"limit"`
	Schedules []Schedule `json:"schedules"`
	Offset    int        `json:"offset"`
}

// Workflow represents a workflow in AlertOps
type Workflow struct {
	WorkflowID           int                 `json:"workflow_id,omitempty"`
	WorkflowName         string              `json:"workflow_name"`
	WorkflowType         string              `json:"workflow_type"`
	Enabled              bool                `json:"enabled"`
	AlertType            string              `json:"alert_type"`
	Scheduled            bool                `json:"scheduled"`
	RecurrenceInterval   int                 `json:"recurrence_interval"`
	Conditions           []WorkflowCondition `json:"conditions"`
	Actions              []WorkflowAction    `json:"actions"`
	IsUsed               bool                `json:"is_used,omitempty"`
	IsBidirection        bool                `json:"is_bidirection,omitempty"`
}

// WorkflowCondition represents a condition in a workflow
type WorkflowCondition struct {
	Type     string `json:"type"`
	Match    string `json:"match"`
	Name     string `json:"name"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
	ListID   string `json:"list_id,omitempty"`
}

// WorkflowAction represents an action in a workflow
type WorkflowAction struct {
	Name                     string   `json:"name"`
	Value                    string   `json:"value,omitempty"`
	WebhookURL               string   `json:"webhook_url,omitempty"`
	SendToOriginalRecipients bool     `json:"send_to_original_recipients,omitempty"`
	SendToSender             bool     `json:"send_to_sender,omitempty"`
	SendToOwner              bool     `json:"send_to_owner,omitempty"`
	LaunchNewThread          bool     `json:"launch_new_thread,omitempty"`
	Subject                  string   `json:"subject,omitempty"`
	MessageText              string   `json:"message_text,omitempty"`
	Users                    []string `json:"users,omitempty"`
	Groups                   []string `json:"groups,omitempty"`
}

// WorkflowListResponse represents the response for listing workflows
type WorkflowListResponse struct {
	Total     int        `json:"total"`
	Limit     int        `json:"limit"`
	Offset    int        `json:"offset"`
	Workflows []Workflow `json:"workflows"`
}

// EscalationPolicy represents an escalation policy
type EscalationPolicy struct {
	EscalationPolicyID                       int                                        `json:"escalation_policy_id,omitempty"`
	EscalationPolicyName                     string                                     `json:"escalation_policy_name"`
	Description                              string                                     `json:"description,omitempty"`
	Priority                                 string                                     `json:"priority,omitempty"`
	Enabled                                  bool                                       `json:"enabled"`
	QuickLaunch                              bool                                       `json:"quick_launch"`
	NotifyUsingCentralizedSettings           bool                                       `json:"notify_using_centralized_settings"`
	MemberRoles                              []EscalationPolicyMemberRole               `json:"member_roles,omitempty"`
	WaitTimeBeforeNotifyingNextGroupInMin    int                                        `json:"wait_time_before_notifying_next_group_in_min,omitempty"`
	GroupContactNotifications                *EscalationPolicyGroupContactNotifications `json:"group_contact_notifications,omitempty"`
	Workflows                                []EscalationPolicyWorkflow                 `json:"workflows,omitempty"`
	OutboundIntegrations                     []EscalationPolicyOutboundIntegration      `json:"outbound_integrations,omitempty"`
	OutboundActions                          []EscalationPolicyOutboundAction           `json:"outbound_actions,omitempty"`
	Options                                  *EscalationPolicyOptions                   `json:"options,omitempty"`
}

// EscalationPolicyMemberRole represents a member role in an escalation policy
type EscalationPolicyMemberRole struct {
	MemberRoleType                 string                             `json:"member_role_type"`
	WaitTimeBetweenMembersInMins   int                                `json:"wait_time_between_members_in_mins,omitempty"`
	RoleWaitTimeInMins             int                                `json:"role_wait_time_in_mins,omitempty"`
	NoOfRetries                    int                                `json:"no_of_retries,omitempty"`
	RetryInterval                  int                                `json:"retry_interval,omitempty"`
	ContactMethods                 []EscalationPolicyContactMethod    `json:"contact_methods,omitempty"`
}

// EscalationPolicyContactMethod represents a contact method in escalation policy
type EscalationPolicyContactMethod struct {
	ContactMethodName string `json:"contact_method_name"`
	WaitTimeInMins    int    `json:"wait_time_in_mins,omitempty"`
	Repeat            bool   `json:"repeat,omitempty"`
	RepeatTimes       int    `json:"repeat_times,omitempty"`
	RepeatMinutes     int    `json:"repeat_minutes,omitempty"`
	ToBccOrCc         string `json:"to_bcc_or_cc,omitempty"`
	Sequence          int    `json:"sequence,omitempty"`
}

// EscalationPolicyGroupContactNotifications represents group contact notifications
type EscalationPolicyGroupContactNotifications struct {
	ContactMethods []EscalationPolicyGroupContactMethod `json:"contact_methods,omitempty"`
}

// EscalationPolicyGroupContactMethod represents a contact method for group notifications
type EscalationPolicyGroupContactMethod struct {
	ContactMethodName string `json:"contact_method_name"`
	WaitTimeInMins    int    `json:"wait_time_in_mins,omitempty"`
	ToBccOrCc         string `json:"to_bcc_or_cc,omitempty"`
}

// EscalationPolicyWorkflow represents a workflow reference in escalation policy
type EscalationPolicyWorkflow struct {
	WorkflowID   int    `json:"workflow_id"`
	WorkflowName string `json:"workflow_name,omitempty"`
}

// EscalationPolicyOutboundIntegration represents an outbound integration
type EscalationPolicyOutboundIntegration struct {
	OutboundIntegrationID int                                        `json:"outbound_integration_id"`
	Name                  string                                     `json:"name,omitempty"`
	IntervalInSec         int                                        `json:"interval_in_sec,omitempty"`
	Actions               []EscalationPolicyOutboundIntegrationAction `json:"actions,omitempty"`
}

// EscalationPolicyOutboundIntegrationAction represents an action in outbound integration
type EscalationPolicyOutboundIntegrationAction struct {
	ActionName string `json:"action_name"`
	Enabled    bool   `json:"enabled"`
}

// EscalationPolicyOutboundAction represents an outbound action
type EscalationPolicyOutboundAction struct {
	ActionID   int    `json:"action_id"`
	ActionName string `json:"action_name,omitempty"`
}

// EscalationPolicyOptions represents the options configuration
type EscalationPolicyOptions struct {
	Acknowledgement                 *EscalationPolicyOptionSettings `json:"acknowledgement,omitempty"`
	Assignment                      *EscalationPolicyOptionSettings `json:"assignment,omitempty"`
	Escalate                        *EscalationPolicyOptionSettings `json:"escalate,omitempty"`
	Close                           *EscalationPolicyOptionSettings `json:"close,omitempty"`
	NotificationSettings            *EscalationPolicyNotificationSettings `json:"notification_settings,omitempty"`
	EscalationPolicyNameForReply    string                                 `json:"escalation_policy_name_for_reply,omitempty"`
	SlaInHours                      float64                                `json:"sla_in_hours,omitempty"`
	MessageText                     string                                 `json:"message_text,omitempty"`
	IncludeAlertIDInSubject         bool                                   `json:"include_alert_id_in_subject,omitempty"`
	OneEmailPerMessage              bool                                   `json:"one_email_per_message,omitempty"`
	OneMessagePerRecipient          bool                                   `json:"one_message_per_recipient,omitempty"`
	SequenceGroupFirst              bool                                   `json:"sequence_group_first,omitempty"`
	AlertType                       string                                 `json:"alert_type,omitempty"`
	Recipients                      []EscalationPolicyRecipient            `json:"recipients,omitempty"`
}

// EscalationPolicyOptionSettings represents boolean settings for various options
type EscalationPolicyOptionSettings struct {
	Phone     bool `json:"phone,omitempty"`
	SMS       bool `json:"sms,omitempty"`
	Email     bool `json:"email,omitempty"`
	GroupChat bool `json:"group_chat,omitempty"`
}

// EscalationPolicyNotificationSettings represents notification settings
type EscalationPolicyNotificationSettings struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
	SMS   string `json:"sms,omitempty"`
}

// EscalationPolicyRecipient represents a recipient in the options
type EscalationPolicyRecipient struct {
	RecipientTypeID int    `json:"recipient_type_id"`
	RecipientID     int    `json:"recipient_id"`
	RecipientName   string `json:"recipient_name,omitempty"`
}

// EscalationPolicyListResponse represents the response for listing escalation policies
type EscalationPolicyListResponse struct {
	Total              int                `json:"total"`
	Limit              int                `json:"limit"`
	Offset             int                `json:"offset"`
	EscalationPolicies []EscalationPolicy `json:"escalation_policies"`
} 

// ============================================================================
// INBOUND INTEGRATION MODELS
// ============================================================================

// InboundIntegration represents an inbound integration in AlertOps
type InboundIntegration struct {
	InboundIntegrationID   int                          `json:"inbound_integration_id,omitempty"`
	InboundIntegrationName string                       `json:"inbound_integration_name"`
	Type                   string                       `json:"type"`
	Sequence               int                          `json:"sequence,omitempty"`
	Enabled                bool                         `json:"enabled,omitempty"`
	EscalationPolicy       string                       `json:"escalation_policy,omitempty"`
	RecipientGroups        []string                     `json:"recipient_groups,omitempty"`
	RecipientUsers         []string                     `json:"recipient_users,omitempty"`
	Bridge                 *InboundIntegrationBridge    `json:"bridge,omitempty"`
	InboundTemplateID      int                          `json:"inbound_template_id,omitempty"`
	APISettings            *InboundIntegrationAPISettings `json:"api_settings,omitempty"`
	MailBox                string                       `json:"mail_box,omitempty"`
	EmailSettings          *InboundIntegrationEmailSettings `json:"email_settings,omitempty"`
	ChatSettings           *InboundIntegrationChatSettings `json:"chat_settings,omitempty"`
	HeartbeatSettings      *InboundIntegrationHeartbeatSettings `json:"heartbeat_settings,omitempty"`
}

// InboundIntegrationBridge represents bridge settings
type InboundIntegrationBridge struct {
	TelephoneNumber string `json:"telephone_number,omitempty"`
	AccessCode      string `json:"access_code,omitempty"`
}

// InboundIntegrationAPISettings represents API settings for inbound integration
type InboundIntegrationAPISettings struct {
	IsBidirection                 bool                                          `json:"is_bidirection,omitempty"`
	URLMapping                    *InboundIntegrationURLMapping                 `json:"url_mapping,omitempty"`
	AlertTags                     *InboundIntegrationAlertTags                  `json:"alert_tags,omitempty"`
	DelayingOrGrouping            *InboundIntegrationDelayingOrGrouping         `json:"delaying_or_grouping,omitempty"`
	FiltersToMatchJSONOrFormFields *InboundIntegrationFilters                   `json:"filters_to_match_json_or_form_fields,omitempty"`
	EscalationPolicyOverride      *InboundIntegrationEscalationPolicyOverride  `json:"escalation_policy_override,omitempty"`
	DynamicRecipientGroups        []InboundIntegrationDynamicRecipientGroup    `json:"dynamic_recipient_groups,omitempty"`
}

// InboundIntegrationURLMapping represents URL mapping settings
type InboundIntegrationURLMapping struct {
	Method              string                                    `json:"method,omitempty"`
	Content             string                                    `json:"content,omitempty"`
	Source              string                                    `json:"source,omitempty"`
	SourceName          string                                    `json:"source_name,omitempty"`
	Static              bool                                      `json:"static,omitempty"`
	SourceValue         string                                    `json:"source_value,omitempty"`
	SourceID            string                                    `json:"source_id,omitempty"`
	SourceURL           string                                    `json:"source_url,omitempty"`
	Severity            string                                    `json:"severity,omitempty"`
	SourceStatus        string                                    `json:"source_status,omitempty"`
	Assignee            string                                    `json:"assignee,omitempty"`
	OpenAlertWhen       *InboundIntegrationCondition              `json:"open_alert_when,omitempty"`
	CloseAlertWhen      *InboundIntegrationSimpleCondition        `json:"close_alert_when,omitempty"`
	UpdateAlertWhen     *InboundIntegrationSimpleCondition        `json:"update_alert_when,omitempty"`
	LongText            string                                    `json:"long_text,omitempty"`
	ShortText           string                                    `json:"short_text,omitempty"`
	Subject             string                                    `json:"subject,omitempty"`
	RecipientUser       string                                    `json:"recipient_user,omitempty"`
	RecipientGroup      string                                    `json:"recipient_group,omitempty"`
	Topic               string                                    `json:"topic,omitempty"`
	SampleData          string                                    `json:"sample_data,omitempty"`
	SampleFieldValue    map[string]interface{}                    `json:"sample_field_value,omitempty"`
	CustomAlertFields   []InboundIntegrationCustomAlertField      `json:"custom_alert_fields,omitempty"`
	Attachments         *InboundIntegrationAttachments            `json:"attachments,omitempty"`
}

// InboundIntegrationCondition represents a condition with field name and values
type InboundIntegrationCondition struct {
	FieldName string   `json:"field_name,omitempty"`
	Type      string   `json:"type,omitempty"`
	Values    []string `json:"values,omitempty"`
}

// InboundIntegrationSimpleCondition represents a simple condition with type and values
type InboundIntegrationSimpleCondition struct {
	Type   string   `json:"type,omitempty"`
	Values []string `json:"values,omitempty"`
}

// InboundIntegrationCustomAlertField represents custom alert field
type InboundIntegrationCustomAlertField struct {
	AttributeName  string `json:"attribute_name,omitempty"`
	AttributeValue string `json:"attribute_value,omitempty"`
	Required       bool   `json:"required,omitempty"`
}

// InboundIntegrationAttachments represents attachment settings
type InboundIntegrationAttachments struct {
	BasePath     string `json:"base_path,omitempty"`
	URL          string `json:"url,omitempty"`
	FileName     string `json:"file_name,omitempty"`
	IsLink       bool   `json:"is_link,omitempty"`
	IsCollection bool   `json:"is_collection,omitempty"`
}

// InboundIntegrationAlertTags represents alert tags
type InboundIntegrationAlertTags struct {
	BusinessService string `json:"business_service,omitempty"`
	ComponentType   string `json:"component_type,omitempty"`
	ComponentName   string `json:"component_name,omitempty"`
	DataCenter      string `json:"data_center,omitempty"`
	Environment     string `json:"environment,omitempty"`
	ProblemType     string `json:"problem_type,omitempty"`
}

// InboundIntegrationDelayingOrGrouping represents delaying and grouping rules
type InboundIntegrationDelayingOrGrouping struct {
	DelayingRule *InboundIntegrationDelayingRule `json:"delaying_rule,omitempty"`
	GroupingRule *InboundIntegrationGroupingRule `json:"grouping_rule,omitempty"`
}

// InboundIntegrationDelayingRule represents delaying rule
type InboundIntegrationDelayingRule struct {
	Delay                                          bool                                                              `json:"delay,omitempty"`
	DelayNotificationsForEveryXAlerts              *InboundIntegrationDelayEveryXAlerts                              `json:"delay_notifications_for_every_x_alerts,omitempty"`
	DelayNotificationsForEveryXMinutes             *InboundIntegrationDelayEveryXMinutes                             `json:"delay_notifications_for_every_x_minutes,omitempty"`
	DelayNotificationsForEveryXAlertsWithinXMinutes *InboundIntegrationDelayEveryXAlertsWithinXMinutes                `json:"delay_notifications_for_every_x_alerts_within_x_minutes,omitempty"`
	DelayNotificationsUntilSupportHours            *InboundIntegrationDelayUntilSupportHours                         `json:"delay_notifications_until_support_hours,omitempty"`
}

// InboundIntegrationDelayEveryXAlerts represents delay every X alerts
type InboundIntegrationDelayEveryXAlerts struct {
	EveryXAlerts int `json:"every_x_alerts,omitempty"`
}

// InboundIntegrationDelayEveryXMinutes represents delay every X minutes
type InboundIntegrationDelayEveryXMinutes struct {
	EveryXMinutes int `json:"every_x_minutes,omitempty"`
}

// InboundIntegrationDelayEveryXAlertsWithinXMinutes represents delay every X alerts within X minutes
type InboundIntegrationDelayEveryXAlertsWithinXMinutes struct {
	EveryXAlerts  int `json:"every_x_alerts,omitempty"`
	EveryXMinutes int `json:"every_x_minutes,omitempty"`
}

// InboundIntegrationDelayUntilSupportHours represents delay until support hours
type InboundIntegrationDelayUntilSupportHours struct {
	WeeklySchedules []InboundIntegrationWeeklySchedule `json:"weekly_schedules,omitempty"`
}

// InboundIntegrationWeeklySchedule represents weekly schedule
type InboundIntegrationWeeklySchedule struct {
	Name       string                              `json:"name,omitempty"`
	DaysOfWeek *InboundIntegrationDaysOfWeek       `json:"days_of_week,omitempty"`
	StartTime  *InboundIntegrationTime             `json:"start_time,omitempty"`
	EndTime    *InboundIntegrationTime             `json:"end_time,omitempty"`
}

// InboundIntegrationDaysOfWeek represents days of week
type InboundIntegrationDaysOfWeek struct {
	Sun bool `json:"sun,omitempty"`
	Mon bool `json:"mon,omitempty"`
	Tue bool `json:"tue,omitempty"`
	Wed bool `json:"wed,omitempty"`
	Thu bool `json:"thu,omitempty"`
	Fri bool `json:"fri,omitempty"`
	Sat bool `json:"sat,omitempty"`
}

// InboundIntegrationTime represents time
type InboundIntegrationTime struct {
	Hour   int `json:"hour,omitempty"`
	Minute int `json:"minute,omitempty"`
}

// InboundIntegrationGroupingRule represents grouping rule
type InboundIntegrationGroupingRule struct {
	Group                    bool                                          `json:"group,omitempty"`
	GroupingWithInXMinutes   *InboundIntegrationGroupingWithInXMinutes     `json:"grouping_with_in_x_minutes,omitempty"`
}

// InboundIntegrationGroupingWithInXMinutes represents grouping within X minutes
type InboundIntegrationGroupingWithInXMinutes struct {
	EveryXMinutes int `json:"every_x_minutes,omitempty"`
}

// InboundIntegrationFilters represents filters to match JSON or form fields
type InboundIntegrationFilters struct {
	AddAllFilter *InboundIntegrationFilterSet `json:"add_all_filter,omitempty"`
	AddAnyFilter *InboundIntegrationFilterSet `json:"add_any_filter,omitempty"`
}

// InboundIntegrationFilterSet represents a set of filters
type InboundIntegrationFilterSet struct {
	Filters []InboundIntegrationFilter `json:"filters,omitempty"`
}

// InboundIntegrationFilter represents a filter
type InboundIntegrationFilter struct {
	FilterID  int                                 `json:"filter_id,omitempty"`
	Condition *InboundIntegrationFilterCondition  `json:"condition,omitempty"`
	Not       bool                                `json:"not,omitempty"`
}

// InboundIntegrationFilterCondition represents filter condition
type InboundIntegrationFilterCondition struct {
	FieldName string `json:"field_name,omitempty"`
	Type      string `json:"type,omitempty"`
}

// InboundIntegrationEscalationPolicyOverride represents escalation policy override
type InboundIntegrationEscalationPolicyOverride struct {
	BasedOnTimeOfDay  *InboundIntegrationEscalationPolicyTimeOfDay    `json:"based_on_time_of_day,omitempty"`
	BasedOnSourceData []InboundIntegrationEscalationPolicySourceData  `json:"based_on_source_data,omitempty"`
}

// InboundIntegrationEscalationPolicyTimeOfDay represents time-based escalation policy override
type InboundIntegrationEscalationPolicyTimeOfDay struct {
	Week                   *InboundIntegrationDaysOfWeek `json:"week,omitempty"`
	StartTime              *InboundIntegrationTime       `json:"start_time,omitempty"`
	EndTime                *InboundIntegrationTime       `json:"end_time,omitempty"`
	EscalationPolicyID     string                        `json:"escalation_policy_id,omitempty"`
	EscalationPolicyName   string                        `json:"escalation_policy_name,omitempty"`
}

// InboundIntegrationEscalationPolicySourceData represents source data-based escalation policy override
type InboundIntegrationEscalationPolicySourceData struct {
	Condition              *InboundIntegrationSourceDataCondition `json:"condition,omitempty"`
	EscalationPolicyID     string                                  `json:"escalation_policy_id,omitempty"`
	EscalationPolicyName   string                                  `json:"escalation_policy_name,omitempty"`
}

// InboundIntegrationSourceDataCondition represents source data condition
type InboundIntegrationSourceDataCondition struct {
	FieldName string `json:"field_name,omitempty"`
	Type      string `json:"type,omitempty"`
	Value     string `json:"value,omitempty"`
}

// InboundIntegrationDynamicRecipientGroup represents dynamic recipient group
type InboundIntegrationDynamicRecipientGroup struct {
	Condition      *InboundIntegrationSourceDataCondition `json:"condition,omitempty"`
	RecipientGroup string                                  `json:"recipient_group,omitempty"`
	GroupName      string                                  `json:"group_name,omitempty"`
}

// ============================================================================
// EMAIL SETTINGS MODELS
// ============================================================================

// InboundIntegrationEmailSettings represents email settings for inbound integration
type InboundIntegrationEmailSettings struct {
	EmailMapping                    *InboundIntegrationEmailMapping              `json:"email_mapping,omitempty"`
	AlertTags                       *InboundIntegrationEmailAlertTags            `json:"alert_tags,omitempty"`
	DelayingOrGrouping              *InboundIntegrationDelayingOrGrouping        `json:"delaying_or_grouping,omitempty"`
	FiltersToMatchIncomingEmails    *InboundIntegrationEmailFilters              `json:"filters_to_match_incoming_emails,omitempty"`
	EscalationPolicyOverride        *InboundIntegrationEscalationPolicyOverride `json:"escalation_policy_override,omitempty"`
	DynamicRecipientGroups          []InboundIntegrationDynamicRecipientGroup   `json:"dynamic_recipient_groups,omitempty"`
}

// InboundIntegrationEmailMapping represents email mapping settings
type InboundIntegrationEmailMapping struct {
	EveryIncomingEmailWillOpenAnAlert bool                                             `json:"every_incoming_email_will_open_an_alert,omitempty"`
	SourceName                        *InboundIntegrationEmailField                    `json:"source_name,omitempty"`
	SourceIdentifier                  *InboundIntegrationEmailField                    `json:"source_identifier,omitempty"`
	OpenAlertWhen                     *InboundIntegrationCondition                     `json:"open_alert_when,omitempty"`
	CloseAlertWhen                    *InboundIntegrationSimpleCondition               `json:"close_alert_when,omitempty"`
	UpdateAlertWhen                   *InboundIntegrationSimpleCondition               `json:"update_alert_when,omitempty"`
	IgnoreDuplicates                  bool                                             `json:"ignore_duplicates,omitempty"`
	LongText                          *InboundIntegrationEmailField                    `json:"long_text,omitempty"`
	ShortText                         *InboundIntegrationEmailField                    `json:"short_text,omitempty"`
	SourceURL                         *InboundIntegrationEmailField                    `json:"source_url,omitempty"`
	AssigneeMailOfficial              *InboundIntegrationEmailField                    `json:"assignee_mail_official,omitempty"`
	RecipientUser                     *InboundIntegrationEmailField                    `json:"recipient_user,omitempty"`
	RecipientGroups                   *InboundIntegrationEmailField                    `json:"recipient_groups,omitempty"`
	Topic                             *InboundIntegrationEmailField                    `json:"topic,omitempty"`
	CustomAlertFields                 []InboundIntegrationEmailCustomAlertField       `json:"custom_alert_fields,omitempty"`
	LongMessageText                   string                                           `json:"long_message_text,omitempty"`
	ShortMessageText                  string                                           `json:"short_message_text,omitempty"`
	SampleData                        string                                           `json:"sample_data,omitempty"`
}

// InboundIntegrationEmailField represents email field with tags
type InboundIntegrationEmailField struct {
	FieldName string `json:"field_name,omitempty"`
	StartTag  string `json:"start_tag,omitempty"`
	EndTag    string `json:"end_tag,omitempty"`
}

// InboundIntegrationEmailCustomAlertField represents email custom alert field
type InboundIntegrationEmailCustomAlertField struct {
	AttributeName     string `json:"attribute_name,omitempty"`
	AttributeValue    string `json:"attribute_value,omitempty"`
	Required          bool   `json:"required,omitempty"`
	AttributeDataType string `json:"attribute_data_type,omitempty"`
	StartTag          string `json:"start_tag,omitempty"`
	EndTag            string `json:"end_tag,omitempty"`
}

// InboundIntegrationEmailAlertTags represents email alert tags with field mapping
type InboundIntegrationEmailAlertTags struct {
	BusinessService *InboundIntegrationEmailField `json:"business_service,omitempty"`
	ComponentType   *InboundIntegrationEmailField `json:"component_type,omitempty"`
	ComponentName   *InboundIntegrationEmailField `json:"component_name,omitempty"`
	DataCenter      *InboundIntegrationEmailField `json:"data_center,omitempty"`
	Environment     *InboundIntegrationEmailField `json:"environment,omitempty"`
	ProblemType     *InboundIntegrationEmailField `json:"problem_type,omitempty"`
}

// InboundIntegrationEmailFilters represents email filters
type InboundIntegrationEmailFilters struct {
	SubjectFilters               []InboundIntegrationEmailSubjectFilter       `json:"subject_filters,omitempty"`
	BodyFilters                  []InboundIntegrationEmailBodyFilter          `json:"body_filters,omitempty"`
	SenderOrRecipientFilters     []InboundIntegrationEmailRecipientFilter     `json:"sender_or_recipient_filters,omitempty"`
	PriorityFilters              []InboundIntegrationEmailPriorityFilter      `json:"priority_filters,omitempty"`
}

// InboundIntegrationEmailSubjectFilter represents subject filter
type InboundIntegrationEmailSubjectFilter struct {
	FilterID  int                                      `json:"filter_id,omitempty"`
	Condition *InboundIntegrationEmailFilterCondition  `json:"condition,omitempty"`
	And       bool                                     `json:"and,omitempty"`
	Not       bool                                     `json:"not,omitempty"`
}

// InboundIntegrationEmailBodyFilter represents body filter
type InboundIntegrationEmailBodyFilter struct {
	FilterID  int                                      `json:"filter_id,omitempty"`
	Condition *InboundIntegrationEmailFilterCondition  `json:"condition,omitempty"`
	And       bool                                     `json:"and,omitempty"`
	Not       bool                                     `json:"not,omitempty"`
}

// InboundIntegrationEmailRecipientFilter represents recipient filter
type InboundIntegrationEmailRecipientFilter struct {
	FilterID         int    `json:"filter_id,omitempty"`
	RecipientName    string `json:"recipient_name,omitempty"`
	RecipientAddress string `json:"recipient_address,omitempty"`
	RecipientType    string `json:"recipient_type,omitempty"`
	And              bool   `json:"and,omitempty"`
	Not              bool   `json:"not,omitempty"`
}

// InboundIntegrationEmailPriorityFilter represents priority filter
type InboundIntegrationEmailPriorityFilter struct {
	FilterID int    `json:"filter_id,omitempty"`
	Priority string `json:"priority,omitempty"`
	And      bool   `json:"and,omitempty"`
	Not      bool   `json:"not,omitempty"`
}

// InboundIntegrationEmailFilterCondition represents email filter condition
type InboundIntegrationEmailFilterCondition struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// ============================================================================
// CHAT AND HEARTBEAT SETTINGS MODELS
// ============================================================================

// InboundIntegrationChatSettings represents chat settings for inbound integration
type InboundIntegrationChatSettings struct {
	URLMapping               *InboundIntegrationChatURLMapping                `json:"url_mapping,omitempty"`
	EscalationPolicyOverride *InboundIntegrationEscalationPolicyOverride     `json:"escalation_policy_override,omitempty"`
}

// InboundIntegrationChatURLMapping represents chat URL mapping
type InboundIntegrationChatURLMapping struct {
	Source      string `json:"source,omitempty"`
	SourceName  string `json:"source_name,omitempty"`
	Static      bool   `json:"static,omitempty"`
	SourceValue string `json:"source_value,omitempty"`
}

// InboundIntegrationHeartbeatSettings represents heartbeat settings
type InboundIntegrationHeartbeatSettings struct {
	HeartbeatIntervalInMin int `json:"heartbeat_interval_in_min,omitempty"`
} 