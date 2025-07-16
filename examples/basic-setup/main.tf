# =============================================================================
# AlertOps Terraform Provider - Basic Setup Example
# =============================================================================
# This example demonstrates the basic setup of AlertOps resources:
# - Creating users with contact methods
# - Setting up groups with proper membership
# - Basic escalation policy configuration

terraform {
  required_version = ">= 0.13"
  required_providers {
    alertops = {
      source  = "alertops/alertops"
      version = "~> 1.0"
    }
  }
}

# Configure the AlertOps Provider
provider "alertops" {
  api_key = var.alertops_api_key
  # base_url = "https://api.alertops.com"  # Optional: defaults to official API
}

# =============================================================================
# VARIABLES
# =============================================================================

variable "alertops_api_key" {
  description = "AlertOps API Key - Get from Settings > API Management"
  type        = string
  sensitive   = true
}

variable "company_domain" {
  description = "Your company domain for email addresses"
  type        = string
  default     = "example.com"
}

variable "environment" {
  description = "Environment name (dev, staging, prod)"
  type        = string
  default     = "dev"
}

# =============================================================================
# USERS
# =============================================================================

# Primary on-call engineer
resource "alertops_user" "primary_oncall" {
  user_name  = "primary-oncall-${var.environment}"
  first_name = "Primary"
  last_name  = "OnCall"
  locale     = "en-US"
  time_zone  = "(UTC-05:00) Eastern Time (US & Canada)"
  type       = "Standard"
  
  # Email contact method
  contact_methods {
    contact_method_name = "Email-Official"
    email {
      email_address = "oncall-primary@${var.company_domain}"
    }
    enabled = true
    sequence = 1
  }
  
  # SMS contact method
  contact_methods {
    contact_method_name = "SMS-Official"
    sms {
      phone_number = "555-0001"
      country_code = "1"
    }
    enabled = true
    sequence = 2
  }
  
  roles = ["Basic"]
}

# Secondary on-call engineer
resource "alertops_user" "secondary_oncall" {
  user_name  = "secondary-oncall-${var.environment}"
  first_name = "Secondary"
  last_name  = "OnCall"
  locale     = "en-US"
  time_zone  = "(UTC-08:00) Pacific Time (US & Canada)"
  type       = "Standard"
  
  contact_methods {
    contact_method_name = "Email-Official"
    email {
      email_address = "oncall-secondary@${var.company_domain}"
    }
    enabled = true
    sequence = 1
  }
  
  roles = ["Basic"]
}

# Team lead
resource "alertops_user" "team_lead" {
  user_name  = "team-lead-${var.environment}"
  first_name = "Team"
  last_name  = "Lead"
  locale     = "en-US"
  time_zone  = "(UTC-06:00) Central Time (US & Canada)"
  type       = "Standard"
  
  contact_methods {
    contact_method_name = "Email-Official"
    email {
      email_address = "team-lead@${var.company_domain}"
    }
    enabled = true
    sequence = 1
  }
  
  roles = ["Basic"]
}

# =============================================================================
# GROUPS
# =============================================================================

# On-call engineers group
resource "alertops_group" "oncall_engineers" {
  group_name = "oncall-engineers-${var.environment}"
  dynamic    = false
  
  members {
    member_type = "User"
    member      = alertops_user.primary_oncall.user_name
    sequence    = 1
    roles       = ["Primary"]
  }
  
  members {
    member_type = "User"
    member      = alertops_user.secondary_oncall.user_name
    sequence    = 2
    roles       = ["Secondary"]
  }
}

# Management group
resource "alertops_group" "management" {
  group_name = "management-${var.environment}"
  dynamic    = false
  
  members {
    member_type = "User"
    member      = alertops_user.team_lead.user_name
    sequence    = 1
    roles       = ["Primary"]
  }
}

# =============================================================================
# SCHEDULE
# =============================================================================

# Basic weekday schedule
resource "alertops_schedule" "weekday_oncall" {
  schedule_name               = "weekday-oncall-${var.environment}"
  group                      = alertops_group.oncall_engineers.group_name
  schedule_type              = "Fixed"
  time_zone                  = "(UTC-05:00) Eastern Time (US & Canada)"
  continuous                 = false
  color                      = "blue"
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

# =============================================================================
# ESCALATION POLICY
# =============================================================================

# Basic escalation policy
resource "alertops_escalation_policy" "standard_escalation" {
  escalation_policy_name                       = "standard-escalation-${var.environment}"
  description                                  = "Standard escalation policy for ${var.environment} environment"
  priority                                     = "High"
  enabled                                      = true
  quick_launch                                 = true
  notify_using_centralized_settings            = false
  wait_time_before_notifying_next_group_in_min = 5

  # Primary role - on-call engineers
  member_roles {
    member_role_type                  = "Primary"
    wait_time_between_members_in_mins = 2
    role_wait_time_in_mins           = 10
    no_of_retries                    = 0
    retry_interval                   = 0
  }

  # Secondary role - management escalation
  member_roles {
    member_role_type                  = "Secondary"
    wait_time_between_members_in_mins = 5
    role_wait_time_in_mins           = 15
    no_of_retries                    = 0
    retry_interval                   = 0
  }

  # Notification options
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
      group_chat = false
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

# =============================================================================
# WORKFLOW
# =============================================================================

# Basic alert processing workflow
resource "alertops_workflow" "alert_processor" {
  workflow_name       = "basic-alert-processor-${var.environment}"
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
    message_text                = "Alert received in ${var.environment}: {{alert.message}}"
    users                       = []
    groups                      = []
  }
}

# =============================================================================
# INBOUND INTEGRATION
# =============================================================================

# API integration for receiving alerts
resource "alertops_inbound_integration" "api_alerts" {
  inbound_integration_name = "api-alerts-${var.environment}"
  type                    = "API"
  sequence                = 1
  enabled                 = true
  escalation_policy       = alertops_escalation_policy.standard_escalation.escalation_policy_name
  recipient_groups        = [alertops_group.oncall_engineers.group_name]
  recipient_users         = [alertops_user.primary_oncall.user_name]

  api_settings {
    is_bidirection = true
  }
}

# =============================================================================
# OUTPUTS
# =============================================================================

output "setup_summary" {
  description = "Summary of created AlertOps resources"
  value = {
    environment = var.environment
    
    users_created = {
      primary_oncall   = alertops_user.primary_oncall.user_name
      secondary_oncall = alertops_user.secondary_oncall.user_name
      team_lead       = alertops_user.team_lead.user_name
    }
    
    groups_created = {
      oncall_engineers = alertops_group.oncall_engineers.group_name
      management      = alertops_group.management.group_name
    }
    
    schedule_created = alertops_schedule.weekday_oncall.schedule_name
    escalation_policy_created = alertops_escalation_policy.standard_escalation.escalation_policy_name
    workflow_created = alertops_workflow.alert_processor.workflow_name
    integration_created = alertops_inbound_integration.api_alerts.inbound_integration_name
  }
}

output "next_steps" {
  description = "Next steps after basic setup"
  value = {
    description = "Your basic AlertOps setup is complete!"
    next_actions = [
      "1. Test your API integration endpoint: Use the integration settings to send test alerts",
      "2. Configure monitoring tools: Point your monitoring systems to the API integration",
      "3. Test escalation: Create a test alert to verify the escalation flow",
      "4. Customize schedules: Add more complex rotation schedules as needed",
      "5. Add more integrations: Set up additional integrations for different alert sources"
    ]
    documentation = "https://docs.alertops.com"
  }
} 