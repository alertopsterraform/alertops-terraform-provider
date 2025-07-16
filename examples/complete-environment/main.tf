# =============================================================================
# AlertOps Terraform Provider - Complete Environment Example
# =============================================================================
# This example demonstrates a comprehensive AlertOps setup for a production
# environment with multiple teams, complex schedules, and advanced integrations.

terraform {
  required_version = ">= 0.13"
  required_providers {
    alertops = {
      source  = "alertops/alertops"
      version = "~> 1.0"
    }
  }
}

provider "alertops" {
  api_key = var.alertops_api_key
}

# =============================================================================
# VARIABLES
# =============================================================================

variable "alertops_api_key" {
  description = "AlertOps API Key"
  type        = string
  sensitive   = true
}

variable "organization_name" {
  description = "Organization name (used in resource naming)"
  type        = string
  default     = "acme"
}

variable "environment" {
  description = "Environment (dev, staging, prod)"
  type        = string
  default     = "prod"
}

variable "company_domain" {
  description = "Company email domain"
  type        = string
  default     = "acme.com"
}

# =============================================================================
# LOCALS
# =============================================================================

locals {
  # Resource naming convention
  prefix = "${var.organization_name}-${var.environment}"
  
  # Common tags/labels
  common_tags = {
    Organization = var.organization_name
    Environment  = var.environment
    ManagedBy    = "terraform"
  }
}

# =============================================================================
# INFRASTRUCTURE TEAM
# =============================================================================

# Infrastructure Engineers
resource "alertops_user" "infra_lead" {
  user_name  = "${local.prefix}-infra-lead"
  first_name = "Infrastructure"
  last_name  = "Lead"
  locale     = "en-US"
  time_zone  = "(UTC-05:00) Eastern Time (US & Canada)"
  type       = "Standard"
  
  contact_methods {
    contact_method_name = "Email-Official"
    email {
      email_address = "infra-lead@${var.company_domain}"
    }
    enabled = true
    sequence = 1
  }
  
  contact_methods {
    contact_method_name = "SMS-Official"
    sms {
      phone_number = "555-1001"
      country_code = "1"
    }
    enabled = true
    sequence = 2
  }
  
  roles = ["Basic"]
}

resource "alertops_user" "infra_engineer_1" {
  user_name  = "${local.prefix}-infra-eng-1"
  first_name = "Sarah"
  last_name  = "DevOps"
  locale     = "en-US"
  time_zone  = "(UTC-05:00) Eastern Time (US & Canada)"
  type       = "Standard"
  
  contact_methods {
    contact_method_name = "Email-Official"
    email {
      email_address = "sarah.devops@${var.company_domain}"
    }
    enabled = true
    sequence = 1
  }
  
  contact_methods {
    contact_method_name = "Phone-Official"
    phone {
      phone_number = "555-1002"
      country_code = "1"
    }
    enabled = true
    sequence = 2
  }
  
  roles = ["Basic"]
}

resource "alertops_user" "infra_engineer_2" {
  user_name  = "${local.prefix}-infra-eng-2"
  first_name = "Mike"
  last_name  = "SRE"
  locale     = "en-US"
  time_zone  = "(UTC-08:00) Pacific Time (US & Canada)"
  type       = "Standard"
  
  contact_methods {
    contact_method_name = "Email-Official"
    email {
      email_address = "mike.sre@${var.company_domain}"
    }
    enabled = true
    sequence = 1
  }
  
  roles = ["Basic"]
}

# =============================================================================
# APPLICATION TEAM
# =============================================================================

resource "alertops_user" "app_lead" {
  user_name  = "${local.prefix}-app-lead"
  first_name = "Application"
  last_name  = "Lead"
  locale     = "en-US"
  time_zone  = "(UTC-06:00) Central Time (US & Canada)"
  type       = "Standard"
  
  contact_methods {
    contact_method_name = "Email-Official"
    email {
      email_address = "app-lead@${var.company_domain}"
    }
    enabled = true
    sequence = 1
  }
  
  roles = ["Basic"]
}

resource "alertops_user" "app_developer_1" {
  user_name  = "${local.prefix}-app-dev-1"
  first_name = "Alice"
  last_name  = "Developer"
  locale     = "en-US"
  time_zone  = "(UTC-06:00) Central Time (US & Canada)"
  type       = "Standard"
  
  contact_methods {
    contact_method_name = "Email-Official"
    email {
      email_address = "alice.dev@${var.company_domain}"
    }
    enabled = true
    sequence = 1
  }
  
  roles = ["Basic"]
}

# =============================================================================
# EXECUTIVE TEAM
# =============================================================================

resource "alertops_user" "cto" {
  user_name  = "${local.prefix}-cto"
  first_name = "CTO"
  last_name  = "Executive"
  locale     = "en-US"
  time_zone  = "(UTC-05:00) Eastern Time (US & Canada)"
  type       = "Standard"
  
  contact_methods {
    contact_method_name = "Email-Official"
    email {
      email_address = "cto@${var.company_domain}"
    }
    enabled = true
    sequence = 1
  }
  
  roles = ["Basic"]
}

# =============================================================================
# GROUPS
# =============================================================================

# Infrastructure team group
resource "alertops_group" "infrastructure_team" {
  group_name = "${local.prefix}-infrastructure-team"
  dynamic    = false
  
  members {
    member_type = "User"
    member      = alertops_user.infra_lead.user_name
    sequence    = 1
    roles       = ["Primary"]
  }
  
  members {
    member_type = "User"
    member      = alertops_user.infra_engineer_1.user_name
    sequence    = 2
    roles       = ["Secondary"]
  }
  
  members {
    member_type = "User"
    member      = alertops_user.infra_engineer_2.user_name
    sequence    = 3
    roles       = ["Secondary"]
  }
}

# Application team group
resource "alertops_group" "application_team" {
  group_name = "${local.prefix}-application-team"
  dynamic    = false
  
  members {
    member_type = "User"
    member      = alertops_user.app_lead.user_name
    sequence    = 1
    roles       = ["Primary"]
  }
  
  members {
    member_type = "User"
    member      = alertops_user.app_developer_1.user_name
    sequence    = 2
    roles       = ["Secondary"]
  }
}

# Executive escalation group
resource "alertops_group" "executives" {
  group_name = "${local.prefix}-executives"
  dynamic    = false
  
  members {
    member_type = "User"
    member      = alertops_user.cto.user_name
    sequence    = 1
    roles       = ["Primary"]
  }
}

# Combined on-call group
resource "alertops_group" "all_oncall" {
  group_name = "${local.prefix}-all-oncall"
  dynamic    = false
  
  members {
    member_type = "User"
    member      = alertops_user.infra_lead.user_name
    sequence    = 1
    roles       = ["Primary"]
  }
  
  members {
    member_type = "User"
    member      = alertops_user.infra_engineer_1.user_name
    sequence    = 2
    roles       = ["Primary"]
  }
  
  members {
    member_type = "User"
    member      = alertops_user.app_lead.user_name
    sequence    = 3
    roles       = ["Primary"]
  }
}

# =============================================================================
# SCHEDULES
# =============================================================================

# Business hours schedule
resource "alertops_schedule" "business_hours" {
  schedule_name               = "${local.prefix}-business-hours"
  group                      = alertops_group.application_team.group_name
  schedule_type              = "Fixed"
  time_zone                  = "(UTC-05:00) Eastern Time (US & Canada)"
  continuous                 = false
  color                      = "green"
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

# 24/7 infrastructure schedule
resource "alertops_schedule" "infrastructure_247" {
  schedule_name               = "${local.prefix}-infrastructure-247"
  group                      = alertops_group.infrastructure_team.group_name
  schedule_type              = "Fixed"
  time_zone                  = "(UTC-05:00) Eastern Time (US & Canada)"
  continuous                 = true
  color                      = "red"
  include_all_users_in_group = true
  
  schedule_weekdays {
    mon = true
    tue = true
    wed = true
    thu = true
    fri = true
    sat = true
    sun = true
  }
}

# Weekend schedule
resource "alertops_schedule" "weekend_coverage" {
  schedule_name               = "${local.prefix}-weekend-coverage"
  group                      = alertops_group.all_oncall.group_name
  schedule_type              = "Fixed"
  time_zone                  = "(UTC-05:00) Eastern Time (US & Canada)"
  continuous                 = false
  color                      = "blue"
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
# WORKFLOWS
# =============================================================================

# Critical alert workflow
resource "alertops_workflow" "critical_alerts" {
  workflow_name       = "${local.prefix}-critical-alerts"
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
    send_to_original_recipients = true
    send_to_sender              = false
    send_to_owner               = true
    launch_new_thread           = false
    subject                     = "🚨 CRITICAL: {{alert.subject}}"
    message_text                = "Critical alert in ${var.environment}: {{alert.message}}"
    users                       = []
    groups                      = []
  }
}

# Auto-close workflow
resource "alertops_workflow" "auto_close" {
  workflow_name       = "${local.prefix}-auto-close"
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
    value    = "Closed"
  }

  actions {
    name                        = "Outbound Service Notification"
    value                       = "Generic Service Desk - Close Incident"
    webhook_url                 = null
    send_to_original_recipients = false
    send_to_sender              = false
    send_to_owner               = false
    launch_new_thread           = false
    subject                     = "✅ RESOLVED: {{alert.subject}}"
    message_text                = "Alert resolved in ${var.environment}: {{alert.message}}"
    users                       = []
    groups                      = []
  }
}

# =============================================================================
# ESCALATION POLICIES
# =============================================================================

# Infrastructure escalation policy
resource "alertops_escalation_policy" "infrastructure_escalation" {
  escalation_policy_name                       = "${local.prefix}-infrastructure-escalation"
  description                                  = "Multi-tier escalation for infrastructure alerts"
  priority                                     = "High"
  enabled                                      = true
  quick_launch                                 = true
  notify_using_centralized_settings            = false
  wait_time_before_notifying_next_group_in_min = 5

  member_roles {
    member_role_type                  = "Primary"
    wait_time_between_members_in_mins = 1
    role_wait_time_in_mins           = 5
    no_of_retries                    = 0
    retry_interval                   = 0
  }

  member_roles {
    member_role_type                  = "Secondary"
    wait_time_between_members_in_mins = 2
    role_wait_time_in_mins           = 10
    no_of_retries                    = 0
    retry_interval                   = 0
  }

  member_roles {
    member_role_type                  = "Tertiary"
    wait_time_between_members_in_mins = 5
    role_wait_time_in_mins           = 15
    no_of_retries                    = 0
    retry_interval                   = 0
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

# Application escalation policy
resource "alertops_escalation_policy" "application_escalation" {
  escalation_policy_name                       = "${local.prefix}-application-escalation"
  description                                  = "Business hours escalation for application alerts"
  priority                                     = "Medium"
  enabled                                      = true
  quick_launch                                 = false
  notify_using_centralized_settings            = false
  wait_time_before_notifying_next_group_in_min = 10

  member_roles {
    member_role_type                  = "Primary"
    wait_time_between_members_in_mins = 5
    role_wait_time_in_mins           = 15
    no_of_retries                    = 0
    retry_interval                   = 0
  }

  member_roles {
    member_role_type                  = "Secondary"
    wait_time_between_members_in_mins = 10
    role_wait_time_in_mins           = 30
    no_of_retries                    = 0
    retry_interval                   = 0
  }

  options {
    acknowledgement {
      phone      = false
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
# INBOUND INTEGRATIONS
# =============================================================================

# Infrastructure monitoring integration
resource "alertops_inbound_integration" "infrastructure_monitoring" {
  inbound_integration_name = "${local.prefix}-infrastructure-monitoring"
  type                    = "API"
  sequence                = 1
  enabled                 = true
  escalation_policy       = alertops_escalation_policy.infrastructure_escalation.escalation_policy_name
  recipient_groups        = [alertops_group.infrastructure_team.group_name]
  recipient_users         = [alertops_user.infra_lead.user_name]

  api_settings {
    is_bidirection = true
  }
}

# Application monitoring integration
resource "alertops_inbound_integration" "application_monitoring" {
  inbound_integration_name = "${local.prefix}-application-monitoring"
  type                    = "API"
  sequence                = 2
  enabled                 = true
  escalation_policy       = alertops_escalation_policy.application_escalation.escalation_policy_name
  recipient_groups        = [alertops_group.application_team.group_name]
  recipient_users         = [alertops_user.app_lead.user_name]

  api_settings {
    is_bidirection = false
  }
}

# Bridge integration for phone escalation
resource "alertops_inbound_integration" "executive_bridge" {
  inbound_integration_name = "${local.prefix}-executive-bridge"
  type                    = "Bridge"
  sequence                = 3
  enabled                 = true
  escalation_policy       = alertops_escalation_policy.infrastructure_escalation.escalation_policy_name
  recipient_groups        = [alertops_group.executives.group_name]
  recipient_users         = [alertops_user.cto.user_name]

  bridge {
    telephone_number = "555-0100"
    access_code      = "12345"
  }

  heartbeat_settings {
    heartbeat_interval_in_min = 60
  }
}

# =============================================================================
# OUTPUTS
# =============================================================================

output "environment_summary" {
  description = "Summary of the complete AlertOps environment"
  value = {
    organization = var.organization_name
    environment  = var.environment
    
    teams_created = {
      infrastructure = {
        lead      = alertops_user.infra_lead.user_name
        engineers = [
          alertops_user.infra_engineer_1.user_name,
          alertops_user.infra_engineer_2.user_name
        ]
        group = alertops_group.infrastructure_team.group_name
      }
      
      application = {
        lead       = alertops_user.app_lead.user_name
        developers = [alertops_user.app_developer_1.user_name]
        group      = alertops_group.application_team.group_name
      }
      
      executives = {
        members = [alertops_user.cto.user_name]
        group   = alertops_group.executives.group_name
      }
    }
    
    schedules_created = {
      business_hours     = alertops_schedule.business_hours.schedule_name
      infrastructure_247 = alertops_schedule.infrastructure_247.schedule_name
      weekend_coverage   = alertops_schedule.weekend_coverage.schedule_name
    }
    
    escalation_policies = {
      infrastructure = alertops_escalation_policy.infrastructure_escalation.escalation_policy_name
      application    = alertops_escalation_policy.application_escalation.escalation_policy_name
    }
    
    workflows_created = {
      critical_alerts = alertops_workflow.critical_alerts.workflow_name
      auto_close      = alertops_workflow.auto_close.workflow_name
    }
    
    integrations_created = {
      infrastructure_monitoring = alertops_inbound_integration.infrastructure_monitoring.inbound_integration_name
      application_monitoring    = alertops_inbound_integration.application_monitoring.inbound_integration_name
      executive_bridge          = alertops_inbound_integration.executive_bridge.inbound_integration_name
    }
  }
}

output "team_contact_summary" {
  description = "Team contact information for verification"
  value = {
    infrastructure_team = {
      lead = "${alertops_user.infra_lead.first_name} ${alertops_user.infra_lead.last_name} (infra-lead@${var.company_domain})"
      members = [
        "${alertops_user.infra_engineer_1.first_name} ${alertops_user.infra_engineer_1.last_name} (sarah.devops@${var.company_domain})",
        "${alertops_user.infra_engineer_2.first_name} ${alertops_user.infra_engineer_2.last_name} (mike.sre@${var.company_domain})"
      ]
    }
    
    application_team = {
      lead = "${alertops_user.app_lead.first_name} ${alertops_user.app_lead.last_name} (app-lead@${var.company_domain})"
      members = [
        "${alertops_user.app_developer_1.first_name} ${alertops_user.app_developer_1.last_name} (alice.dev@${var.company_domain})"
      ]
    }
    
    executives = {
      cto = "${alertops_user.cto.first_name} ${alertops_user.cto.last_name} (cto@${var.company_domain})"
    }
  }
}

output "integration_endpoints" {
  description = "Integration endpoints for monitoring tools"
  value = {
    infrastructure_monitoring = "Configure your infrastructure monitoring to send alerts to the '${alertops_inbound_integration.infrastructure_monitoring.inbound_integration_name}' integration"
    application_monitoring    = "Configure your application monitoring to send alerts to the '${alertops_inbound_integration.application_monitoring.inbound_integration_name}' integration"
    executive_bridge          = "Executive bridge available at ${alertops_inbound_integration.executive_bridge.bridge}"
  }
}

output "next_steps" {
  description = "Recommended next steps"
  value = {
    description = "Your complete AlertOps environment has been set up!"
    actions = [
      "1. Configure monitoring tools to use the integration endpoints",
      "2. Test escalation flows with each team",
      "3. Set up rotation schedules for each team",
      "4. Train team members on AlertOps procedures",
      "5. Customize workflows for specific alert types",
      "6. Set up additional integrations as needed"
    ]
  }
} 