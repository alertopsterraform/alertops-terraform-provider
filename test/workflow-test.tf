# AlertOps Workflow Resource Test Configuration
# =============================================================================
# Tests the Workflow resource CRUD operations
# Note: Uses configuration from schedule-test.tf for provider setup

# WORKFLOW RESOURCE TEST
# =============================================================================

resource "alertops_workflow" "test_workflow" {
  workflow_name        = "UPDATED CRUD terraform-test-workflow-${local.test_suffix}"
  workflow_type        = "Alert"
  enabled              = true
  alert_type           = "Standard Alert"
  scheduled            = false
  recurrence_interval  = 0

  conditions {
    type     = "start"
    match    = "all"
    name     = "AlertStatus"
    operator = "is"
    value    = "Closed"
  }

  actions {
    name                       = "Outbound Service Notification"
    value                      = "Generic Service Desk - Close Incident"
    webhook_url                = null
    send_to_original_recipients = false
    send_to_sender             = false
    send_to_owner              = false
    launch_new_thread          = false
    subject                    = ""
    message_text               = ""
    users                      = []
    groups                     = []
  }

}

# Test Outputs for Verification
output "workflow_creation_success" {
  value = "✅ Workflow created: ${alertops_workflow.test_workflow.workflow_name} (ID: ${alertops_workflow.test_workflow.workflow_id})"
}

output "workflow_test_results" {
  value = {
    test_suffix         = local.test_suffix
    workflow_id         = alertops_workflow.test_workflow.workflow_id
    workflow_name       = alertops_workflow.test_workflow.workflow_name
    workflow_type       = alertops_workflow.test_workflow.workflow_type
    enabled             = alertops_workflow.test_workflow.enabled
    alert_type          = alertops_workflow.test_workflow.alert_type
    scheduled           = alertops_workflow.test_workflow.scheduled
    recurrence_interval = alertops_workflow.test_workflow.recurrence_interval
    is_used             = alertops_workflow.test_workflow.is_used
    is_bidirection      = alertops_workflow.test_workflow.is_bidirection
    conditions_count    = length(alertops_workflow.test_workflow.conditions)
    actions_count       = length(alertops_workflow.test_workflow.actions)
  }
}

output "workflow_debug" {
  value = {
    workflow_request_json = alertops_workflow.test_workflow.debug_request_json
  }
  description = "Debug information showing the JSON sent to AlertOps API"
} 