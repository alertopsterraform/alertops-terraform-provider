# AlertOps Escalation Policy Resource Test Configuration
# =============================================================================
# Tests the Escalation Policy resource CRUD operations
# Note: Uses configuration from schedule-test.tf for provider setup

# ESCALATION POLICY RESOURCE TEST
# =============================================================================

resource "alertops_escalation_policy" "test_escalation_policy" {
  escalation_policy_name           = "terraform-test-escalation-policy-${local.test_suffix}"
  description                      = "Test escalation policy created by Terraform"
  priority                         = "High"
  enabled                          = true
  quick_launch                     = true
  notify_using_centralized_settings = false  # User settings mode - no contact methods in member roles
  wait_time_before_notifying_next_group_in_min = 5

  member_roles {
    member_role_type                = "Primary"
    wait_time_between_members_in_mins = 2
    role_wait_time_in_mins          = 10
    # In user settings mode, these should be 0 and contact_methods should be empty
    no_of_retries                   = 0
    retry_interval                  = 0
    # No contact_methods block - will use user's own contact preferences
  }

  member_roles {
    member_role_type                = "Secondary"
    wait_time_between_members_in_mins = 5
    role_wait_time_in_mins          = 15
    # In user settings mode, these should be 0 and contact_methods should be empty
    no_of_retries                   = 0
    retry_interval                  = 0
    # No contact_methods block - will use user's own contact preferences
  }

  # Remove group_contact_notifications since "Email-Official" doesn't exist
  # group_contact_notifications {
  #   contact_methods {
  #     contact_method_name = "Email-Official"
  #     wait_time_in_mins   = 0
  #   }
  # }

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

    # Remove invalid escalation_policy_name_for_reply
    # escalation_policy_name_for_reply = "Test Policy Reply"
    sla_in_hours                     = 4
    message_text                     = "This is a test escalation policy message"
    include_alert_id_in_subject      = true
    one_email_per_message            = false
    one_message_per_recipient        = true
    sequence_group_first             = true
    alert_type                       = "Standard Alert"
  }
}

# REALISTIC ESCALATION POLICY (Based on AlertOps Production Example)
# =============================================================================

resource "alertops_escalation_policy" "realistic_escalation_policy" {
  escalation_policy_name           = "Escalate after 5 mins (Pri->Sec->Mgr) if not assigned/closed - ${local.test_suffix}"
  description                      = "Notify members at 5 minute intervals using user preferences before escalating to the next member"
  priority                         = "Critical"
  enabled                          = true
  quick_launch                     = false
  notify_using_centralized_settings = false  # User settings mode
  wait_time_before_notifying_next_group_in_min = 1

  # Primary escalation member (0 minutes)
  member_roles {
    member_role_type                = "Primary"
    wait_time_between_members_in_mins = 0
    role_wait_time_in_mins          = 0
    # User settings mode - no retries or contact methods
    no_of_retries                   = 0
    retry_interval                  = 0
  }

  # Secondary escalation member (+5 minutes)
  member_roles {
    member_role_type                = "Secondary" 
    wait_time_between_members_in_mins = 0
    role_wait_time_in_mins          = 5
    # User settings mode - no retries or contact methods
    no_of_retries                   = 0
    retry_interval                  = 0
  }

  # Manager escalation member (+10 minutes)
  member_roles {
    member_role_type                = "Manager"
    wait_time_between_members_in_mins = 0
    role_wait_time_in_mins          = 10
    # User settings mode - no retries or contact methods
    no_of_retries                   = 0
    retry_interval                  = 0
  }

  # Include workflow references as in production example
  workflows {
    workflow_id   = 32752
    workflow_name = "Notify Owner 1 Hr Before SLA If Not Resolved"
  }

  workflows {
    workflow_id   = 32747
    workflow_name = "Stop Delivery on Assignment or Closed"
  }

  # Production-like options configuration
  options {
    acknowledgement {
      phone      = false
      sms        = false
      email      = false
      group_chat = true  # Only Group Chat for acknowledgement
    }

    assignment {
      phone      = true
      sms        = true
      email      = true
      group_chat = true  # All channels for assignment
    }

    escalate {
      phone      = true
      sms        = true
      email      = true
      group_chat = false  # No Group Chat for escalation
    }

    close {
      phone      = true
      sms        = true
      email      = true
      group_chat = true  # All channels for close
    }

    notification_settings {
      email = "Email Using Long Text"
      phone = "VOIP Using Short Text"
      sms   = "SMS Using Short Text"
    }

    sla_in_hours                     = 1  # 1 hour SLA as in production
    message_text                     = "Critical escalation - please respond immediately"
    include_alert_id_in_subject      = true
    one_email_per_message            = false
    one_message_per_recipient        = true
    sequence_group_first             = true
    alert_type                       = "Standard Alert"
  }
}

# Test Outputs for Verification
output "escalation_policy_creation_success" {
  value = "✅ Escalation Policy created: ${alertops_escalation_policy.test_escalation_policy.escalation_policy_name} (ID: ${alertops_escalation_policy.test_escalation_policy.escalation_policy_id})"
}

output "realistic_escalation_policy_creation_success" {
  value = "✅ Realistic Escalation Policy created: ${alertops_escalation_policy.realistic_escalation_policy.escalation_policy_name} (ID: ${alertops_escalation_policy.realistic_escalation_policy.escalation_policy_id})"
}

output "escalation_policy_test_results" {
  value = {
    test_suffix                          = local.test_suffix
    escalation_policy_id                 = alertops_escalation_policy.test_escalation_policy.escalation_policy_id
    escalation_policy_name               = alertops_escalation_policy.test_escalation_policy.escalation_policy_name
    description                          = alertops_escalation_policy.test_escalation_policy.description
    priority                             = alertops_escalation_policy.test_escalation_policy.priority
    enabled                              = alertops_escalation_policy.test_escalation_policy.enabled
    quick_launch                         = alertops_escalation_policy.test_escalation_policy.quick_launch
    notify_using_centralized_settings    = alertops_escalation_policy.test_escalation_policy.notify_using_centralized_settings
    wait_time_before_notifying_next_group = alertops_escalation_policy.test_escalation_policy.wait_time_before_notifying_next_group_in_min
    member_roles_count                   = length(alertops_escalation_policy.test_escalation_policy.member_roles)
    has_group_notifications              = length(alertops_escalation_policy.test_escalation_policy.group_contact_notifications) > 0
    has_options                          = length(alertops_escalation_policy.test_escalation_policy.options) > 0
  }
  description = "Test results for Escalation Policy resource"
}

output "realistic_escalation_policy_test_results" {
  value = {
    test_suffix                          = local.test_suffix
    escalation_policy_id                 = alertops_escalation_policy.realistic_escalation_policy.escalation_policy_id
    escalation_policy_name               = alertops_escalation_policy.realistic_escalation_policy.escalation_policy_name
    description                          = alertops_escalation_policy.realistic_escalation_policy.description
    priority                             = alertops_escalation_policy.realistic_escalation_policy.priority
    enabled                              = alertops_escalation_policy.realistic_escalation_policy.enabled
    quick_launch                         = alertops_escalation_policy.realistic_escalation_policy.quick_launch
    notify_using_centralized_settings    = alertops_escalation_policy.realistic_escalation_policy.notify_using_centralized_settings
    wait_time_before_notifying_next_group = alertops_escalation_policy.realistic_escalation_policy.wait_time_before_notifying_next_group_in_min
    member_roles_count                   = length(alertops_escalation_policy.realistic_escalation_policy.member_roles)
    workflows_count                      = length(alertops_escalation_policy.realistic_escalation_policy.workflows)
    has_group_notifications              = length(alertops_escalation_policy.realistic_escalation_policy.group_contact_notifications) > 0
    has_options                          = length(alertops_escalation_policy.realistic_escalation_policy.options) > 0
  }
  description = "Test results for Realistic Escalation Policy resource (based on production example)"
}

output "escalation_policy_configuration_summary" {
  value = {
    total_member_roles     = length(alertops_escalation_policy.test_escalation_policy.member_roles)
    total_contact_methods  = sum([for role in alertops_escalation_policy.test_escalation_policy.member_roles : length(role.contact_methods)])
    has_acknowledgement    = length(alertops_escalation_policy.test_escalation_policy.options) > 0 ? length(alertops_escalation_policy.test_escalation_policy.options[0].acknowledgement) > 0 : false
    has_assignment         = length(alertops_escalation_policy.test_escalation_policy.options) > 0 ? length(alertops_escalation_policy.test_escalation_policy.options[0].assignment) > 0 : false
    has_escalate           = length(alertops_escalation_policy.test_escalation_policy.options) > 0 ? length(alertops_escalation_policy.test_escalation_policy.options[0].escalate) > 0 : false
    has_close              = length(alertops_escalation_policy.test_escalation_policy.options) > 0 ? length(alertops_escalation_policy.test_escalation_policy.options[0].close) > 0 : false
    has_notification_settings = length(alertops_escalation_policy.test_escalation_policy.options) > 0 ? length(alertops_escalation_policy.test_escalation_policy.options[0].notification_settings) > 0 : false
  }
  description = "Configuration summary showing the complexity of the escalation policy"
}

output "realistic_escalation_policy_configuration_summary" {
  value = {
    escalation_flow = "Primary (0min) → Secondary (+5min) → Manager (+10min)"
    timing_strategy = "5-minute intervals with user preferences"
    acknowledgement_channels = "Group Chat Only"
    assignment_channels = "Phone, SMS, Email, Group Chat"
    escalate_channels = "Phone, SMS, Email (No Group Chat)"
    close_channels = "Phone, SMS, Email, Group Chat"
    notification_preferences = {
      email = "Email Using Long Text"
      phone = "VOIP Using Short Text"
      sms   = "SMS Using Short Text"
    }
    sla_hours = 0
    reply_policy = "All hands email only"
    workflows_attached = length(alertops_escalation_policy.realistic_escalation_policy.workflows)
  }
  description = "Realistic escalation policy configuration matching production patterns"
} 