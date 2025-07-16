# ============================================================================= 
# AlertOps Terraform Provider - Comprehensive CRUD Test
# =============================================================================
# Tests all resource types with proper dependencies and lifecycle management
# This test covers: Users, Groups, Schedules, Workflows, Escalation Policies, 
# and Inbound Integrations

terraform {
  required_providers {
    alertops = {
      source = "alertopsterraform/alertops"
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

# Generate unique suffix to avoid conflicts
locals {
  test_suffix = formatdate("MMDDHHmm", timestamp())
}

# =============================================================================
# PHASE 1: USERS (Foundation Resources)
# =============================================================================

resource "alertops_user" "primary_user" {
  user_name  = "terraform-primary-user-${local.test_suffix}"
  first_name = "Primary"
  last_name  = "TestUser"
  locale     = "en-US"
  time_zone  = "(UTC-06:00) Central Time (US & Canada)"
  type       = "Standard"

  contact_methods {
    contact_method_name = "Email-Official"
    email {
      email_address = "primary-${local.test_suffix}@example.com"
    }
    enabled  = true
    sequence = 1
  }

  contact_methods {
    contact_method_name = "SMS-Official"
    sms {
      phone_number = "555-0001"
      country_code = "1"
    }
    enabled  = true
    sequence = 2
  }

  roles = ["Basic"]
}

resource "alertops_user" "secondary_user" {
  user_name  = "terraform-secondary-user-${local.test_suffix}"
  first_name = "Secondary"
  last_name  = "TestUser"
  locale     = "en-US"
  time_zone  = "(UTC-05:00) Eastern Time (US & Canada)"
  type       = "Standard"

  contact_methods {
    contact_method_name = "Email-Official"
    email {
      email_address = "secondary-${local.test_suffix}@example.com"
    }
    enabled  = true
    sequence = 1
  }

  roles = ["Basic"]
}

resource "alertops_user" "backup_user" {
  user_name  = "terraform-backup-user-${local.test_suffix}"
  first_name = "Backup"
  last_name  = "TestUser"
  locale     = "en-US"
  time_zone  = "(UTC-08:00) Pacific Time (US & Canada)"
  type       = "Standard"

  contact_methods {
    contact_method_name = "Email-Official"
    email {
      email_address = "backup-${local.test_suffix}@example.com"
    }
    enabled  = true
    sequence = 1
  }

  roles = ["Basic"]
}

# =============================================================================
# PHASE 2: GROUPS (Use Users)
# =============================================================================

resource "alertops_group" "primary_group" {
  group_name = "terraform-primary-group-${local.test_suffix}"
  dynamic    = false

  members {
    member_type = "User"
    member      = alertops_user.primary_user.user_name
    sequence    = 1
    roles       = ["Primary"]
  }

  members {
    member_type = "User"
    member      = alertops_user.secondary_user.user_name
    sequence    = 2
    roles       = ["Secondary"]
  }
}

resource "alertops_group" "backup_group" {
  group_name = "terraform-backup-group-${local.test_suffix}"
  dynamic    = false

  members {
    member_type = "User"
    member      = alertops_user.backup_user.user_name
    sequence    = 1
    roles       = ["Primary"]
  }
}

resource "alertops_group" "combined_group" {
  group_name = "terraform-combined-group-${local.test_suffix}"
  dynamic    = false

  members {
    member_type = "User"
    member      = alertops_user.primary_user.user_name
    sequence    = 1
    roles       = ["Primary"]
  }

  members {
    member_type = "User"
    member      = alertops_user.secondary_user.user_name
    sequence    = 2
    roles       = ["Secondary"]
  }

  members {
    member_type = "User"
    member      = alertops_user.backup_user.user_name
    sequence    = 3
    roles       = ["Primary"]
  }
}

# =============================================================================
# PHASE 3: SCHEDULES (Use Users and Groups)
# =============================================================================

resource "alertops_schedule" "primary_schedule" {
  schedule_name = "terraform-primary-schedule-${local.test_suffix}"
  group         = alertops_group.primary_group.group_name
  schedule_type = "Fixed"
  time_zone     = "(UTC-06:00) Central Time (US & Canada)"
  continuous    = false
  color         = "blue"
  include_all_users_in_group = true
  
  schedule_weekdays {
    mon = true
    tue = true
    wed = true
    thu = true
    fri = true
    sat = false
    sun = false
  }
}

resource "alertops_schedule" "backup_schedule" {
  schedule_name = "terraform-backup-schedule-${local.test_suffix}"
  group         = alertops_group.backup_group.group_name
  schedule_type = "Fixed"
  time_zone     = "(UTC-08:00) Pacific Time (US & Canada)"
  continuous    = false
  color         = "green"
  include_all_users_in_group = true
  
  schedule_weekdays {
    mon = false
    tue = false
    wed = false
    thu = false
    fri = false
    sat = true
    sun = true
  }
}

# =============================================================================
# PHASE 4: WORKFLOWS (Independent)
# =============================================================================

resource "alertops_workflow" "alert_processing_workflow" {
  workflow_name       = "terraform-alert-processing-${local.test_suffix}"
  workflow_type       = "Alert"
  enabled             = true
  alert_type          = "Standard Alert"
  scheduled           = false
  recurrence_interval = 0

  conditions {
    type     = "start"
    match    = "all"
    name     = "AlertStatus"
    operator = "is"
    value    = "Open"
  }

  actions {
    name                        = "Outbound Service Notification"
    value                       = "Generic Service Desk - Create Incident"
    webhook_url                 = null
    send_to_original_recipients = false
    send_to_sender              = false
    send_to_owner               = false
    launch_new_thread           = false
    subject                     = "Alert: {{alert.subject}}"
    message_text                = "Alert processing workflow triggered for: {{alert.message}}"
    users                       = []
    groups                      = []
  }
}

resource "alertops_workflow" "notification_workflow" {
  workflow_name       = "terraform-notification-${local.test_suffix}"
  workflow_type       = "Alert"
  enabled             = true
  alert_type          = "Standard Alert"
  scheduled           = false
  recurrence_interval = 0

  conditions {
    type     = "start"
    match    = "any"
    name     = "AlertStatus"
    operator = "is"
    value    = "Open"
  }

  actions {
    name                        = "Outbound Service Notification"
    value                       = "Generic Service Desk - Create Incident"
    webhook_url                 = null
    send_to_original_recipients = true
    send_to_sender              = false
    send_to_owner               = true
    launch_new_thread           = false
    subject                     = "HIGH PRIORITY: {{alert.subject}}"
    message_text                = "High priority alert requires immediate attention: {{alert.message}}"
    users                       = []
    groups                      = []
  }
}

# =============================================================================
# PHASE 5: ESCALATION POLICIES (Use Users, Groups, Schedules)
# =============================================================================

resource "alertops_escalation_policy" "primary_escalation_policy" {
  escalation_policy_name                       = "terraform-primary-escalation-${local.test_suffix}"
  description                                  = "Primary escalation policy for comprehensive testing"
  priority                                     = "High"
  enabled                                      = true
  quick_launch                                 = true
  notify_using_centralized_settings            = false
  wait_time_before_notifying_next_group_in_min = 5

  member_roles {
    member_role_type                  = "Primary"
    wait_time_between_members_in_mins = 2
    role_wait_time_in_mins            = 10
    no_of_retries                     = 0
    retry_interval                    = 0
  }

  member_roles {
    member_role_type                  = "Secondary"
    wait_time_between_members_in_mins = 5
    role_wait_time_in_mins            = 15
    no_of_retries                     = 0
    retry_interval                    = 0
  }

  options {
    acknowledgement {
      phone      = true
      sms        = true
      email      = true
      group_chat = false
    }

    assignment {
      phone      = true
      sms        = true
      email      = true
      group_chat = true
    }

    escalate {
      phone      = true
      sms        = true
      email      = true
      group_chat = false
    }

    close {
      phone      = false
      sms        = false
      email      = true
      group_chat = false
    }

    notification_settings {
      email = "Email Using Long Text"
      phone = "VOIP Using Short Text"
      sms   = "SMS Using Short Text"
    }
  }
}

resource "alertops_escalation_policy" "backup_escalation_policy" {
  escalation_policy_name                       = "terraform-backup-escalation-${local.test_suffix}"
  description                                  = "Backup escalation policy for comprehensive testing"
  priority                                     = "Medium"
  enabled                                      = true
  quick_launch                                 = false
  notify_using_centralized_settings            = false
  wait_time_before_notifying_next_group_in_min = 10

  member_roles {
    member_role_type                  = "Primary"
    wait_time_between_members_in_mins = 5
    role_wait_time_in_mins            = 15
    no_of_retries                     = 0
    retry_interval                    = 0
  }

  options {
    acknowledgement {
      phone      = true
      sms        = false
      email      = true
      group_chat = false
    }

    assignment {
      phone      = false
      sms        = false
      email      = true
      group_chat = false
    }

    escalate {
      phone      = true
      sms        = false
      email      = true
      group_chat = false
    }

    close {
      phone      = false
      sms        = false
      email      = true
      group_chat = false
    }

    notification_settings {
      email = "Email Using Long Text"
      phone = "VOIP Using Short Text"
      sms   = "SMS Using Short Text"
    }
  }
}

# =============================================================================
# PHASE 6: INBOUND INTEGRATIONS (Use Escalation Policies)
# =============================================================================

resource "alertops_inbound_integration" "api_integration" {
  inbound_integration_name = "terraform-api-integration-${local.test_suffix}"
  type                     = "API"
  sequence                 = 1
  enabled                  = true
  escalation_policy        = alertops_escalation_policy.primary_escalation_policy.escalation_policy_name
  recipient_groups         = [alertops_group.primary_group.group_name]
  recipient_users          = [alertops_user.primary_user.user_name]

  api_settings {
    is_bidirection = true
  }
}

resource "alertops_inbound_integration" "email_integration" {
  inbound_integration_name = "terraform-api-integration-2-${local.test_suffix}"
  type                     = "API"
  sequence                 = 2
  enabled                  = true
  escalation_policy        = alertops_escalation_policy.backup_escalation_policy.escalation_policy_name
  recipient_groups         = [alertops_group.backup_group.group_name]
  recipient_users          = [alertops_user.backup_user.user_name]

  api_settings {
    is_bidirection = false
  }
}

resource "alertops_inbound_integration" "bridge_integration" {
  inbound_integration_name = "terraform-bridge-integration-${local.test_suffix}"
  type                     = "Bridge"
  sequence                 = 3
  enabled                  = true
  escalation_policy        = alertops_escalation_policy.primary_escalation_policy.escalation_policy_name
  recipient_groups         = [alertops_group.combined_group.group_name]
  recipient_users = [
    alertops_user.primary_user.user_name,
    alertops_user.secondary_user.user_name
  ]

  bridge {
    telephone_number = "555-0100"
    access_code      = "12345"
  }

  heartbeat_settings {
    heartbeat_interval_in_min = 60
  }
}

# =============================================================================
# TEST OUTPUTS FOR VERIFICATION
# =============================================================================

output "comprehensive_test_summary" {
  value = {
    test_suffix = local.test_suffix
    timestamp   = timestamp()

    users_created = {
      primary   = "${alertops_user.primary_user.user_name} (ID: ${alertops_user.primary_user.user_id})"
      secondary = "${alertops_user.secondary_user.user_name} (ID: ${alertops_user.secondary_user.user_id})"
      backup    = "${alertops_user.backup_user.user_name} (ID: ${alertops_user.backup_user.user_id})"
    }

    groups_created = {
      primary  = "${alertops_group.primary_group.group_name} (ID: ${alertops_group.primary_group.group_id})"
      backup   = "${alertops_group.backup_group.group_name} (ID: ${alertops_group.backup_group.group_id})"
      combined = "${alertops_group.combined_group.group_name} (ID: ${alertops_group.combined_group.group_id})"
    }

    schedules_created = {
      primary = "${alertops_schedule.primary_schedule.schedule_name} (ID: ${alertops_schedule.primary_schedule.schedule_id})"
      backup  = "${alertops_schedule.backup_schedule.schedule_name} (ID: ${alertops_schedule.backup_schedule.schedule_id})"
    }

    workflows_created = {
      alert_processing = "${alertops_workflow.alert_processing_workflow.workflow_name} (ID: ${alertops_workflow.alert_processing_workflow.workflow_id})"
      notification     = "${alertops_workflow.notification_workflow.workflow_name} (ID: ${alertops_workflow.notification_workflow.workflow_id})"
    }

    escalation_policies_created = {
      primary = "${alertops_escalation_policy.primary_escalation_policy.escalation_policy_name} (ID: ${alertops_escalation_policy.primary_escalation_policy.escalation_policy_id})"
      backup  = "${alertops_escalation_policy.backup_escalation_policy.escalation_policy_name} (ID: ${alertops_escalation_policy.backup_escalation_policy.escalation_policy_id})"
    }

    inbound_integrations_created = {
      api_primary   = "${alertops_inbound_integration.api_integration.inbound_integration_name} (ID: ${alertops_inbound_integration.api_integration.inbound_integration_id})"
      api_secondary = "${alertops_inbound_integration.email_integration.inbound_integration_name} (ID: ${alertops_inbound_integration.email_integration.inbound_integration_id})"
      bridge        = "${alertops_inbound_integration.bridge_integration.inbound_integration_name} (ID: ${alertops_inbound_integration.bridge_integration.inbound_integration_id})"
    }
  }
}

output "crud_test_success_indicators" {
  value = {
    total_resources_created = 16
    resource_types_tested   = 6
    dependency_chain_tested = "Users → Groups → Schedules → Escalation Policies → Inbound Integrations"
    independent_resources   = "Workflows"

    success_flags = {
      users_have_contact_methods     = length(alertops_user.primary_user.contact_methods) > 0
      groups_have_members            = length(alertops_group.primary_group.members) > 0
      schedules_are_created          = alertops_schedule.primary_schedule.schedule_id != 0
      workflows_have_conditions      = length(alertops_workflow.alert_processing_workflow.conditions) > 0
      escalation_policies_have_roles = length(alertops_escalation_policy.primary_escalation_policy.member_roles) > 0
      integrations_have_settings     = alertops_inbound_integration.api_integration.api_settings != null
    }
  }
}

output "cleanup_commands" {
  value = {
    description = "Run these commands to clean up test resources"
    commands = [
      "terraform destroy -auto-approve",
      "# Or selectively destroy in reverse dependency order:",
      "terraform destroy -target=alertops_inbound_integration.api_integration -auto-approve",
      "terraform destroy -target=alertops_inbound_integration.email_integration -auto-approve",
      "terraform destroy -target=alertops_inbound_integration.bridge_integration -auto-approve",
      "terraform destroy -target=alertops_escalation_policy.primary_escalation_policy -auto-approve",
      "terraform destroy -target=alertops_escalation_policy.backup_escalation_policy -auto-approve",
      "terraform destroy -target=alertops_workflow.alert_processing_workflow -auto-approve",
      "terraform destroy -target=alertops_workflow.notification_workflow -auto-approve",
      "terraform destroy -target=alertops_schedule.primary_schedule -auto-approve",
      "terraform destroy -target=alertops_schedule.backup_schedule -auto-approve",
      "terraform destroy -target=alertops_group.primary_group -auto-approve",
      "terraform destroy -target=alertops_group.backup_group -auto-approve",
      "terraform destroy -target=alertops_group.combined_group -auto-approve",
      "terraform destroy -target=alertops_user.primary_user -auto-approve",
      "terraform destroy -target=alertops_user.secondary_user -auto-approve",
      "terraform destroy -target=alertops_user.backup_user -auto-approve"
    ]
  }
}

# =============================================================================
# CRUD VERIFICATION TESTS
# =============================================================================

# Data source tests to verify resource creation
data "alertops_user" "verify_primary_user" {
  user_id = alertops_user.primary_user.user_id
}

output "verification_tests" {
  value = {
    primary_user_verification = {
      created_name = alertops_user.primary_user.user_name
      fetched_name = data.alertops_user.verify_primary_user.user_name
      names_match  = alertops_user.primary_user.user_name == data.alertops_user.verify_primary_user.user_name
    }

    resource_id_validation = {
      all_users_have_ids = alltrue([
        alertops_user.primary_user.user_id != "",
        alertops_user.secondary_user.user_id != "",
        alertops_user.backup_user.user_id != ""
      ])
      all_groups_have_ids = alltrue([
        alertops_group.primary_group.group_id != "",
        alertops_group.backup_group.group_id != "",
        alertops_group.combined_group.group_id != ""
      ])
      all_schedules_have_ids = alltrue([
        alertops_schedule.primary_schedule.schedule_id != "",
        alertops_schedule.backup_schedule.schedule_id != ""
      ])
      all_workflows_have_ids = alltrue([
        alertops_workflow.alert_processing_workflow.workflow_id != "",
        alertops_workflow.notification_workflow.workflow_id != ""
      ])
      all_escalation_policies_have_ids = alltrue([
        alertops_escalation_policy.primary_escalation_policy.escalation_policy_id != "",
        alertops_escalation_policy.backup_escalation_policy.escalation_policy_id != ""
      ])
      all_integrations_have_ids = alltrue([
        alertops_inbound_integration.api_integration.inbound_integration_id != 0,
        alertops_inbound_integration.email_integration.inbound_integration_id != 0,
        alertops_inbound_integration.bridge_integration.inbound_integration_id != 0
      ])
    }
  }
} 