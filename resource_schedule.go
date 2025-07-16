package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSchedule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceScheduleCreate,
		ReadContext:   resourceScheduleRead,
		UpdateContext: resourceScheduleUpdate,
		DeleteContext: resourceScheduleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"schedule_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The unique identifier for the schedule",
			},
			"group": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The group this schedule belongs to",
			},
			"schedule_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the schedule",
			},
			"schedule_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of schedule",
			},
			"continuous": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether the schedule is continuous",
			},
			"time_zone": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The timezone for the schedule",
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Color for the schedule display",
			},
			"start_date": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Start date configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"date": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Start date (YYYY-MM-DD format)",
						},
						"hour": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Start hour (0-23)",
						},
						"minute": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Start minute (0-59)",
						},
					},
				},
			},
			"end_date": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "End date configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"date": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "End date (YYYY-MM-DD format)",
						},
						"hour": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "End hour (0-23)",
						},
						"minute": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "End minute (0-59)",
						},
					},
				},
			},
			"start_weekday": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Starting weekday",
			},
			"end_weekday": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Ending weekday",
			},
			"schedule_weekdays": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Weekdays configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sun": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Include Sunday",
						},
						"mon": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Include Monday",
						},
						"tue": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Include Tuesday",
						},
						"wed": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Include Wednesday",
						},
						"thu": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Include Thursday",
						},
						"fri": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Include Friday",
						},
						"sat": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Include Saturday",
						},
					},
				},
			},
			"rotate_frequency": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Rotation frequency (daily, weekly, monthly)",
			},
			"rotate_daily": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Daily rotation configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rotate_x_users": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Number of users to rotate",
						},
						"rotate_at_time": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Required:    true,
							Description: "Time to rotate",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hour": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Hour (0-23)",
									},
									"minute": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Minute (0-59)",
									},
								},
							},
						},
						"every_x_days": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Rotate every X days",
						},
					},
				},
			},
			"rotate_weekly": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Weekly rotation configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rotate_x_users": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Number of users to rotate",
						},
						"rotate_at_time": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Required:    true,
							Description: "Time to rotate",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hour": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Hour (0-23)",
									},
									"minute": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Minute (0-59)",
									},
								},
							},
						},
						"every_x_weeks": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Rotate every X weeks",
						},
						"rotate_at_day_of_week": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Day of week to rotate",
						},
					},
				},
			},
			"rotate_monthly": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Monthly rotation configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rotate_x_users": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Number of users to rotate",
						},
						"rotate_at_time": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Required:    true,
							Description: "Time to rotate",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hour": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Hour (0-23)",
									},
									"minute": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Minute (0-59)",
									},
								},
							},
						},
						"every_x_months": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Rotate every X months",
						},
					},
				},
			},
			"repeat_schedule": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Schedule repetition configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"every_x_weeks": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Repeat every X weeks",
						},
						"repeat_until_date": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Repeat until date",
						},
					},
				},
			},
			"include_all_users_in_group": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Include all users from the group",
			},
			"users": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Users assigned to this schedule",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Username",
						},
						"role": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "User role in schedule",
						},
					},
				},
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether the schedule is enabled",
			},
			"is_holiday_notify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to notify on holidays",
			},
			"debug_request_json": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "JSON payload sent to AlertOps API (for debugging)",
			},
		},
	}
}

func resourceScheduleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	schedule := Schedule{
		Group:                    d.Get("group").(string),
		ScheduleName:             d.Get("schedule_name").(string),
		ScheduleType:             d.Get("schedule_type").(string),
		Continuous:               d.Get("continuous").(bool),
		TimeZone:                 d.Get("time_zone").(string),
		IncludeAllUsersInGroup:   d.Get("include_all_users_in_group").(bool),
		Enabled:                  d.Get("enabled").(bool),
		IsHolidayNotify:          d.Get("is_holiday_notify").(bool),
	}

	if v, ok := d.GetOk("color"); ok {
		schedule.Color = v.(string)
	}

	if v, ok := d.GetOk("start_weekday"); ok {
		schedule.StartWeekday = v.(string)
	}

	if v, ok := d.GetOk("end_weekday"); ok {
		schedule.EndWeekday = v.(string)
	}

	if v, ok := d.GetOk("rotate_frequency"); ok {
		schedule.RotateFrequency = v.(string)
	}

	// Handle nested objects
	if v, ok := d.GetOk("start_date"); ok {
		schedule.StartDate = expandScheduleDate(v.([]interface{}))
	}

	if v, ok := d.GetOk("end_date"); ok {
		schedule.EndDate = expandScheduleDate(v.([]interface{}))
	}

	if v, ok := d.GetOk("schedule_weekdays"); ok {
		schedule.ScheduleWeekdays = expandScheduleWeekdays(v.([]interface{}))
	}

	if v, ok := d.GetOk("rotate_daily"); ok {
		schedule.RotateDaily = expandRotateDaily(v.([]interface{}))
	}

	if v, ok := d.GetOk("rotate_weekly"); ok {
		schedule.RotateWeekly = expandRotateWeekly(v.([]interface{}))
	}

	if v, ok := d.GetOk("rotate_monthly"); ok {
		schedule.RotateMonthly = expandRotateMonthly(v.([]interface{}))
	}

	if v, ok := d.GetOk("repeat_schedule"); ok {
		schedule.RepeatSchedule = expandRepeatSchedule(v.([]interface{}))
	}

	if v, ok := d.GetOk("users"); ok {
		schedule.Users = expandScheduleUsers(v.([]interface{}))
	}

	// Store the request JSON for debugging
	requestJSON, _ := json.Marshal(schedule)
	d.Set("debug_request_json", string(requestJSON))

	var createdSchedule Schedule
	err := client.post(ctx, "/api/v2/schedules", schedule, &createdSchedule)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create schedule: %w", err))
	}

	d.SetId(strconv.Itoa(createdSchedule.ScheduleID))
	d.Set("schedule_id", createdSchedule.ScheduleID)

	return resourceScheduleRead(ctx, d, meta)
}

func resourceScheduleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	scheduleID := d.Id()
	groupID := d.Get("group").(string)
	var schedule Schedule
	err := client.get(ctx, fmt.Sprintf("/api/v2/schedules/%s/%s", groupID, scheduleID), &schedule)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read schedule: %w", err))
	}

	d.Set("schedule_id", schedule.ScheduleID)
	d.Set("group", schedule.Group)
	d.Set("schedule_name", schedule.ScheduleName)
	d.Set("schedule_type", schedule.ScheduleType)
	d.Set("continuous", schedule.Continuous)
	d.Set("time_zone", schedule.TimeZone)
	d.Set("color", schedule.Color)
	d.Set("start_weekday", schedule.StartWeekday)
	d.Set("end_weekday", schedule.EndWeekday)
	d.Set("rotate_frequency", schedule.RotateFrequency)
	d.Set("include_all_users_in_group", schedule.IncludeAllUsersInGroup)
	d.Set("enabled", schedule.Enabled)
	d.Set("is_holiday_notify", schedule.IsHolidayNotify)

	// Set nested objects
	if schedule.StartDate != nil {
		d.Set("start_date", flattenScheduleDate(schedule.StartDate))
	}

	if schedule.EndDate != nil {
		d.Set("end_date", flattenScheduleDate(schedule.EndDate))
	}

	if schedule.ScheduleWeekdays != nil {
		d.Set("schedule_weekdays", flattenScheduleWeekdays(schedule.ScheduleWeekdays))
	}

	if schedule.RotateDaily != nil {
		d.Set("rotate_daily", flattenRotateDaily(schedule.RotateDaily))
	}

	if schedule.RotateWeekly != nil {
		d.Set("rotate_weekly", flattenRotateWeekly(schedule.RotateWeekly))
	}

	if schedule.RotateMonthly != nil {
		d.Set("rotate_monthly", flattenRotateMonthly(schedule.RotateMonthly))
	}

	if schedule.RepeatSchedule != nil {
		d.Set("repeat_schedule", flattenRepeatSchedule(schedule.RepeatSchedule))
	}

	if len(schedule.Users) > 0 {
		d.Set("users", flattenScheduleUsers(schedule.Users))
	}

	return nil
}

func resourceScheduleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	scheduleID := d.Id()
	schedule := Schedule{
		ScheduleID:               d.Get("schedule_id").(int),
		Group:                    d.Get("group").(string),
		ScheduleName:             d.Get("schedule_name").(string),
		ScheduleType:             d.Get("schedule_type").(string),
		Continuous:               d.Get("continuous").(bool),
		TimeZone:                 d.Get("time_zone").(string),
		IncludeAllUsersInGroup:   d.Get("include_all_users_in_group").(bool),
		Enabled:                  d.Get("enabled").(bool),
		IsHolidayNotify:          d.Get("is_holiday_notify").(bool),
	}

	if v, ok := d.GetOk("color"); ok {
		schedule.Color = v.(string)
	}

	if v, ok := d.GetOk("start_weekday"); ok {
		schedule.StartWeekday = v.(string)
	}

	if v, ok := d.GetOk("end_weekday"); ok {
		schedule.EndWeekday = v.(string)
	}

	if v, ok := d.GetOk("rotate_frequency"); ok {
		schedule.RotateFrequency = v.(string)
	}

	// Handle nested objects
	if v, ok := d.GetOk("start_date"); ok {
		schedule.StartDate = expandScheduleDate(v.([]interface{}))
	}

	if v, ok := d.GetOk("end_date"); ok {
		schedule.EndDate = expandScheduleDate(v.([]interface{}))
	}

	if v, ok := d.GetOk("schedule_weekdays"); ok {
		schedule.ScheduleWeekdays = expandScheduleWeekdays(v.([]interface{}))
	}

	if v, ok := d.GetOk("rotate_daily"); ok {
		schedule.RotateDaily = expandRotateDaily(v.([]interface{}))
	}

	if v, ok := d.GetOk("rotate_weekly"); ok {
		schedule.RotateWeekly = expandRotateWeekly(v.([]interface{}))
	}

	if v, ok := d.GetOk("rotate_monthly"); ok {
		schedule.RotateMonthly = expandRotateMonthly(v.([]interface{}))
	}

	if v, ok := d.GetOk("repeat_schedule"); ok {
		schedule.RepeatSchedule = expandRepeatSchedule(v.([]interface{}))
	}

	if v, ok := d.GetOk("users"); ok {
		schedule.Users = expandScheduleUsers(v.([]interface{}))
	}

	// Store the request JSON for debugging
	requestJSON, _ := json.Marshal(schedule)
	d.Set("debug_request_json", string(requestJSON))

	// AlertOps API returns 204 No Content for updates, so we don't expect a response body
	groupID := d.Get("group").(string)
	err := client.put(ctx, fmt.Sprintf("/api/v2/schedules/%s/%s", groupID, scheduleID), schedule, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update schedule: %w", err))
	}

	return resourceScheduleRead(ctx, d, meta)
}

func resourceScheduleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	scheduleID := d.Id()
	groupID := d.Get("group").(string)
	err := client.delete(ctx, fmt.Sprintf("/api/v2/schedules/%s/%s", groupID, scheduleID))
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete schedule: %w", err))
	}

	d.SetId("")
	return nil
}

// Helper functions for expanding/flattening nested structures

func expandScheduleDate(dates []interface{}) *ScheduleDate {
	if len(dates) == 0 {
		return nil
	}

	date := dates[0].(map[string]interface{})
	return &ScheduleDate{
		Date:   date["date"].(string),
		Hour:   date["hour"].(int),
		Minute: date["minute"].(int),
	}
}

func flattenScheduleDate(date *ScheduleDate) []interface{} {
	if date == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"date":   date.Date,
			"hour":   date.Hour,
			"minute": date.Minute,
		},
	}
}

func expandScheduleTime(times []interface{}) *ScheduleTime {
	if len(times) == 0 {
		return nil
	}

	time := times[0].(map[string]interface{})
	return &ScheduleTime{
		Hour:   time["hour"].(int),
		Minute: time["minute"].(int),
	}
}

func flattenScheduleTime(time *ScheduleTime) []interface{} {
	if time == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"hour":   time.Hour,
			"minute": time.Minute,
		},
	}
}

func expandScheduleWeekdays(weekdays []interface{}) *ScheduleWeekdays {
	if len(weekdays) == 0 {
		return nil
	}

	wd := weekdays[0].(map[string]interface{})
	return &ScheduleWeekdays{
		Sun: wd["sun"].(bool),
		Mon: wd["mon"].(bool),
		Tue: wd["tue"].(bool),
		Wed: wd["wed"].(bool),
		Thu: wd["thu"].(bool),
		Fri: wd["fri"].(bool),
		Sat: wd["sat"].(bool),
	}
}

func flattenScheduleWeekdays(weekdays *ScheduleWeekdays) []interface{} {
	if weekdays == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"sun": weekdays.Sun,
			"mon": weekdays.Mon,
			"tue": weekdays.Tue,
			"wed": weekdays.Wed,
			"thu": weekdays.Thu,
			"fri": weekdays.Fri,
			"sat": weekdays.Sat,
		},
	}
}

func expandRotateDaily(rotations []interface{}) *RotateDaily {
	if len(rotations) == 0 {
		return nil
	}

	rotation := rotations[0].(map[string]interface{})
	rd := &RotateDaily{
		RotateXUsers: rotation["rotate_x_users"].(int),
		EveryXDays:   rotation["every_x_days"].(int),
	}

	if v, ok := rotation["rotate_at_time"]; ok {
		rd.RotateAtTime = expandScheduleTime(v.([]interface{}))
	}

	return rd
}

func flattenRotateDaily(rotation *RotateDaily) []interface{} {
	if rotation == nil {
		return []interface{}{}
	}

	result := map[string]interface{}{
		"rotate_x_users": rotation.RotateXUsers,
		"every_x_days":   rotation.EveryXDays,
	}

	if rotation.RotateAtTime != nil {
		result["rotate_at_time"] = flattenScheduleTime(rotation.RotateAtTime)
	}

	return []interface{}{result}
}

func expandRotateWeekly(rotations []interface{}) *RotateWeekly {
	if len(rotations) == 0 {
		return nil
	}

	rotation := rotations[0].(map[string]interface{})
	rw := &RotateWeekly{
		RotateXUsers:      rotation["rotate_x_users"].(int),
		EveryXWeeks:       rotation["every_x_weeks"].(int),
		RotateAtDayOfWeek: rotation["rotate_at_day_of_week"].(string),
	}

	if v, ok := rotation["rotate_at_time"]; ok {
		rw.RotateAtTime = expandScheduleTime(v.([]interface{}))
	}

	return rw
}

func flattenRotateWeekly(rotation *RotateWeekly) []interface{} {
	if rotation == nil {
		return []interface{}{}
	}

	result := map[string]interface{}{
		"rotate_x_users":        rotation.RotateXUsers,
		"every_x_weeks":         rotation.EveryXWeeks,
		"rotate_at_day_of_week": rotation.RotateAtDayOfWeek,
	}

	if rotation.RotateAtTime != nil {
		result["rotate_at_time"] = flattenScheduleTime(rotation.RotateAtTime)
	}

	return []interface{}{result}
}

func expandRotateMonthly(rotations []interface{}) *RotateMonthly {
	if len(rotations) == 0 {
		return nil
	}

	rotation := rotations[0].(map[string]interface{})
	rm := &RotateMonthly{
		RotateXUsers: rotation["rotate_x_users"].(int),
		EveryXMonths: rotation["every_x_months"].(int),
	}

	if v, ok := rotation["rotate_at_time"]; ok {
		rm.RotateAtTime = expandScheduleTime(v.([]interface{}))
	}

	return rm
}

func flattenRotateMonthly(rotation *RotateMonthly) []interface{} {
	if rotation == nil {
		return []interface{}{}
	}

	result := map[string]interface{}{
		"rotate_x_users":  rotation.RotateXUsers,
		"every_x_months":  rotation.EveryXMonths,
	}

	if rotation.RotateAtTime != nil {
		result["rotate_at_time"] = flattenScheduleTime(rotation.RotateAtTime)
	}

	return []interface{}{result}
}

func expandRepeatSchedule(repeats []interface{}) *RepeatSchedule {
	if len(repeats) == 0 {
		return nil
	}

	repeat := repeats[0].(map[string]interface{})
	return &RepeatSchedule{
		EveryXWeeks:     repeat["every_x_weeks"].(int),
		RepeatUntilDate: repeat["repeat_until_date"].(string),
	}
}

func flattenRepeatSchedule(repeat *RepeatSchedule) []interface{} {
	if repeat == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"every_x_weeks":     repeat.EveryXWeeks,
			"repeat_until_date": repeat.RepeatUntilDate,
		},
	}
}

func expandScheduleUsers(users []interface{}) []ScheduleUser {
	if len(users) == 0 {
		return nil
	}

	scheduleUsers := make([]ScheduleUser, len(users))
	for i, user := range users {
		u := user.(map[string]interface{})
		scheduleUsers[i] = ScheduleUser{
			User: u["user"].(string),
			Role: u["role"].(string),
		}
	}

	return scheduleUsers
}

func flattenScheduleUsers(users []ScheduleUser) []interface{} {
	if len(users) == 0 {
		return []interface{}{}
	}

	result := make([]interface{}, len(users))
	for i, user := range users {
		result[i] = map[string]interface{}{
			"user": user.User,
			"role": user.Role,
		}
	}

	return result
} 