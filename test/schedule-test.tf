terraform {
  required_providers {
    alertops = {
      source  = "alertopsterraform/alertops"
    }
  }
}

provider "alertops" {
  api_key = var.alertops_api_key
}

variable "alertops_api_key" {
  description = "AlertOps API key"
  type        = string
  sensitive   = true
}

# Generate unique suffix to avoid conflicts with existing resources
locals {
  test_suffix = formatdate("MMDDHHmm", timestamp())
}

# =============================================================================
# PHASE 1: USER AND GROUP (Prerequisites for Schedule)
# =============================================================================

resource "alertops_user" "schedule_test_user" {
  user_name  = "terraform-schedule-user-${local.test_suffix}"
  first_name = "Schedule"
  last_name  = "TestUser"
  locale     = "en-US"
  time_zone  = "(UTC-06:00) Central Time (US & Canada)"
  type       = "Standard"
  
  contact_methods {
    contact_method_name = "Email-Official"
    email {
      email_address = "schedule-user-${local.test_suffix}@example.com"
    }
    enabled = true
    sequence = 1
  }
  
  roles = ["Basic"]
}

resource "alertops_group" "schedule_test_group" {
  group_name = "terraform-schedule-group-${local.test_suffix}"
  dynamic    = false
  
  description = [
    "Test group for schedule testing",
    "Created for Schedule resource validation"
  ]
  
  members {
    member_type = "User"
    member      = alertops_user.schedule_test_user.user_name
    sequence    = 1
    roles       = ["Primary"]
  }
  
  # Explicit dependency: Group depends on User
  depends_on = [alertops_user.schedule_test_user]
}

# =============================================================================
# PHASE 2: SCHEDULE TESTING (Simplified)
# =============================================================================

resource "alertops_schedule" "simple_schedule" {
  group          = tostring(alertops_group.schedule_test_group.group_id)
  schedule_name  = "terraform-test-schedule-${local.test_suffix}"
  schedule_type  = "Fixed"
  continuous     = false
  time_zone      = "(UTC-06:00) Central Time (US & Canada)"
  color          = "#FF5733"
  
  # Required start and end weekdays
  start_weekday  = "Mon"
  end_weekday    = "Sun"
  
  # Required schedule weekdays - all days enabled
  schedule_weekdays {
    sun = true
    mon = true
    tue = true
    wed = true
    thu = true
    fri = true
    sat = true
  }
  
  # Basic configuration without complex rotation
  include_all_users_in_group = false
  enabled                   = true
  is_holiday_notify         = false
  
  # Assign the test user to the schedule
  users {
    user = alertops_user.schedule_test_user.user_name
    role = "Primary"
  }
  
  # Explicit dependencies: Schedule depends on Group and User
  # Ensures deletion order: Schedule → Group → User
  depends_on = [
    alertops_group.schedule_test_group,
    alertops_user.schedule_test_user
  ]
}

# =============================================================================
# OUTPUTS FOR VERIFICATION
# =============================================================================

output "schedule_test_results" {
  value = {
    user_id      = alertops_user.schedule_test_user.user_id
    user_name    = alertops_user.schedule_test_user.user_name
    group_id     = alertops_group.schedule_test_group.group_id
    group_name   = alertops_group.schedule_test_group.group_name
    schedule_id   = alertops_schedule.simple_schedule.schedule_id
    schedule_name = alertops_schedule.simple_schedule.schedule_name
  }
  description = "Test results for Schedule resource"
}

output "schedule_debug" {
  value = {
    schedule_request_json = alertops_schedule.simple_schedule.debug_request_json
  }
  description = "Debug information for schedule API request"
}

output "integration_verification" {
  value = {
    test_suffix         = local.test_suffix
    user_created        = alertops_user.schedule_test_user.user_name
    group_created       = alertops_group.schedule_test_group.group_name
    group_id_created    = alertops_group.schedule_test_group.group_id
    schedule_created    = alertops_schedule.simple_schedule.schedule_name
    group_has_user      = alertops_group.schedule_test_group.members[0].member
    schedule_has_user   = alertops_schedule.simple_schedule.users[0].user
    schedule_uses_group = alertops_schedule.simple_schedule.group
    dependency_chain    = "${alertops_user.schedule_test_user.user_name} -> ${alertops_group.schedule_test_group.group_name}(${alertops_group.schedule_test_group.group_id}) -> ${alertops_schedule.simple_schedule.schedule_name}"
  }
  description = "Verification that all dependencies work correctly with unique suffix"
} 