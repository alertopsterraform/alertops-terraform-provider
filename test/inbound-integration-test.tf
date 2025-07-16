# AlertOps Inbound Integration Resource Test Configuration
# =============================================================================
# Tests the Inbound Integration resource CRUD operations
# Note: Uses configuration from schedule-test.tf for provider setup

# Use the test suffix from the schedule test configuration
# (test_suffix is already defined in schedule-test.tf)

# BASIC API INBOUND INTEGRATION
# =============================================================================

resource "alertops_inbound_integration" "test_api_integration" {
  inbound_integration_name = "terraform-test-api-integration-${local.test_suffix}"
  type                    = "API"
  sequence                = 1
  enabled                 = true
  escalation_policy       = "Default Policy"

  recipient_groups = [
    "Operations Team",
    "DevOps Team"
  ]

  recipient_users = [
    "admin@example.com"
  ]

  api_settings {
    is_bidirection = true
  }

  heartbeat_settings {
    heartbeat_interval_in_min = 15
  }
}

# EMAIL INBOUND INTEGRATION
# =============================================================================

resource "alertops_inbound_integration" "test_email_integration" {
  inbound_integration_name = "terraform-test-email-integration-${local.test_suffix}"
  type                    = "Email"
  sequence                = 2
  enabled                 = true
  escalation_policy       = "Email Policy"
  mail_box                = "alerts@example.com"

  recipient_groups = [
    "Support Team"
  ]
}

# BRIDGE (TELEPHONE) INBOUND INTEGRATION
# =============================================================================

resource "alertops_inbound_integration" "test_bridge_integration" {
  inbound_integration_name = "terraform-test-bridge-integration-${local.test_suffix}"
  type                    = "Bridge"
  sequence                = 3
  enabled                 = true
  escalation_policy       = "Bridge Policy"

  bridge {
    telephone_number = "+1-555-123-4567"
    access_code      = "12345"
  }

  heartbeat_settings {
    heartbeat_interval_in_min = 30
  }
}

# OUTPUTS FOR TESTING
# =============================================================================

output "inbound_integration_api_results" {
  description = "Results from API inbound integration creation"
  value = {
    integration_id   = alertops_inbound_integration.test_api_integration.inbound_integration_id
    integration_name = alertops_inbound_integration.test_api_integration.inbound_integration_name
    type            = alertops_inbound_integration.test_api_integration.type
    enabled         = alertops_inbound_integration.test_api_integration.enabled
    has_api_settings = length(alertops_inbound_integration.test_api_integration.api_settings) > 0
    has_heartbeat   = length(alertops_inbound_integration.test_api_integration.heartbeat_settings) > 0
    test_suffix     = local.test_suffix
  }
}

output "inbound_integration_email_results" {
  description = "Results from Email inbound integration creation"
  value = {
    integration_id   = alertops_inbound_integration.test_email_integration.inbound_integration_id
    integration_name = alertops_inbound_integration.test_email_integration.inbound_integration_name
    type            = alertops_inbound_integration.test_email_integration.type
    enabled         = alertops_inbound_integration.test_email_integration.enabled
    mail_box        = alertops_inbound_integration.test_email_integration.mail_box
    test_suffix     = local.test_suffix
  }
}

output "inbound_integration_bridge_results" {
  description = "Results from Bridge inbound integration creation"
  value = {
    integration_id   = alertops_inbound_integration.test_bridge_integration.inbound_integration_id
    integration_name = alertops_inbound_integration.test_bridge_integration.inbound_integration_name
    type            = alertops_inbound_integration.test_bridge_integration.type
    enabled         = alertops_inbound_integration.test_bridge_integration.enabled
    has_bridge      = length(alertops_inbound_integration.test_bridge_integration.bridge) > 0
    has_heartbeat   = length(alertops_inbound_integration.test_bridge_integration.heartbeat_settings) > 0
    test_suffix     = local.test_suffix
  }
}

output "inbound_integration_creation_success" {
  description = "Success message for inbound integration creation"
  value = "✅ Inbound Integrations created: API (${alertops_inbound_integration.test_api_integration.inbound_integration_name}), Email (${alertops_inbound_integration.test_email_integration.inbound_integration_name}), Bridge (${alertops_inbound_integration.test_bridge_integration.inbound_integration_name})"
}

output "inbound_integration_summary" {
  description = "Summary of created inbound integrations"
  value = {
    total_integrations = 3
    api_integration = {
      name = alertops_inbound_integration.test_api_integration.inbound_integration_name
      type = alertops_inbound_integration.test_api_integration.type
      id   = alertops_inbound_integration.test_api_integration.inbound_integration_id
    }
    email_integration = {
      name = alertops_inbound_integration.test_email_integration.inbound_integration_name
      type = alertops_inbound_integration.test_email_integration.type
      id   = alertops_inbound_integration.test_email_integration.inbound_integration_id
    }
    bridge_integration = {
      name = alertops_inbound_integration.test_bridge_integration.inbound_integration_name
      type = alertops_inbound_integration.test_bridge_integration.type
      id   = alertops_inbound_integration.test_bridge_integration.inbound_integration_id
    }
  }
} 