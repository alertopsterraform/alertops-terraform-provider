# AlertOps Terraform Provider - Comprehensive CRUD Test

## Overview

The comprehensive CRUD test validates all 6 resource types in the AlertOps Terraform provider with proper dependency management and lifecycle testing. This is the most complete test available, creating a realistic AlertOps environment with all resource types working together.

## What It Tests

### Resource Types (6 total)
1. **Users** - 3 users with different contact methods and time zones
2. **Groups** - 3 groups with different member configurations
3. **Schedules** - 2 schedules with rotation settings and restrictions
4. **Workflows** - 2 workflows for alert processing and notifications
5. **Escalation Policies** - 2 policies with multi-tier escalation
6. **Inbound Integrations** - 3 integrations (2x API, 1x Bridge)

### Total Resources Created: 16

## Test Architecture

```
Users (Foundation)
  ├── Primary User (Email + SMS)
  ├── Secondary User (Email)
  └── Backup User (Email)
      ↓
Groups (Uses Users)
  ├── Primary Group (Primary + Secondary Users)
  ├── Backup Group (Backup User)
  └── Combined Group (All Users)
      ↓
Schedules (Uses Users)
  ├── Primary Schedule (Primary + Secondary, 9-5)
  └── Backup Schedule (Backup User, 5-9)
      ↓
Workflows (Independent)
  ├── Alert Processing Workflow
  └── Notification Workflow
      ↓
Escalation Policies (Uses Users + Groups + Schedules)
  ├── Primary Policy (3-tier: Users → Groups → Schedules)
  └── Backup Policy (Simple: Groups only)
      ↓
Inbound Integrations (Uses Escalation Policies)
  ├── API Integration 1 (bidirectional, 15min heartbeat)
  ├── API Integration 2 (unidirectional, 30min heartbeat)
  └── Bridge Integration (with phone access, 60min heartbeat)
```

## Prerequisites

1. **AlertOps API Key**: Set environment variable
   ```cmd
   set ALERTOPS_API_KEY=your_api_key_here
   ```

2. **Terraform Variables**: Ensure `terraform.tfvars` exists with:
   ```hcl
   alertops_api_key = "your_api_key_here"
   ```

3. **Provider Binary**: Compiled `terraform-provider-alertops.exe` in parent directory

## Running the Test

### Option 1: Automated Test Runner (Recommended)
```cmd
cd test
run-comprehensive-crud-test.cmd
```

This script will:
- Validate configuration
- Create resources in proper dependency order
- Show progress for each phase
- Generate comprehensive outputs
- Offer cleanup options

### Option 2: Manual Execution
```cmd
cd test

# Validate first
validate-comprehensive-test.cmd

# Then run manually
terraform init
terraform plan comprehensive-crud-test.tf
terraform apply comprehensive-crud-test.tf
```

### Option 3: Validation Only
```cmd
cd test
validate-comprehensive-test.cmd
```

## Test Phases

The test runs in 7 phases with proper dependency management:

1. **Users** - Foundation resources with contact methods
2. **Groups** - Member relationships and group structures  
3. **Schedules** - On-call rotations with time restrictions
4. **Workflows** - Alert processing and notification logic
5. **Escalation Policies** - Multi-tier escalation with all member types
6. **Inbound Integrations** - Different integration types with settings
7. **Verification** - Final validation and output generation

## Expected Results

### Success Indicators
- ✅ All 16 resources created successfully
- ✅ All resource IDs populated (non-empty)
- ✅ Dependency relationships working (groups have members, etc.)
- ✅ Different configurations tested (contact methods, schedules, etc.)
- ✅ Data source verification (can fetch created users)

### Comprehensive Outputs
The test generates detailed outputs showing:
- All created resource names and IDs
- Success flags for key functionality
- Dependency chain validation
- Resource type coverage summary

## Cleanup

### Automatic Cleanup
The test runner offers 3 cleanup options:
1. Keep resources for manual testing
2. Automatic cleanup (`terraform destroy`)
3. View selective cleanup commands

### Manual Cleanup
```cmd
# Complete cleanup
terraform destroy -auto-approve comprehensive-crud-test.tf

# Selective cleanup (reverse dependency order)
terraform destroy -target=alertops_inbound_integration.api_integration -auto-approve
terraform destroy -target=alertops_escalation_policy.primary_escalation_policy -auto-approve
# ... etc
```

## Troubleshooting

### Common Issues

1. **API Key Not Set**
   ```
   ERROR: ALERTOPS_API_KEY environment variable is not set!
   ```
   **Solution**: Set the environment variable before running

2. **Validation Failed**
   ```
   ❌ Configuration validation failed!
   ```
   **Solution**: Run `validate-comprehensive-test.cmd` to identify issues

3. **Dependency Errors**
   ```
   ❌ Group creation failed!
   ```
   **Solution**: Ensure users were created first, or run phases individually

4. **Resource Conflicts**
   ```
   Error: Resource already exists
   ```
   **Solution**: The test uses timestamp-based suffixes to avoid conflicts

### Debugging

1. **Enable Terraform Debug**:
   ```cmd
   set TF_LOG=DEBUG
   terraform apply comprehensive-crud-test.tf
   ```

2. **Check Individual Phases**:
   ```cmd
   terraform plan -target=alertops_user.primary_user comprehensive-crud-test.tf
   ```

3. **Validate API Connection**:
   ```cmd
   terraform plan comprehensive-crud-test.tf
   ```

## File Structure

```
test/
├── comprehensive-crud-test.tf           # Main test configuration
├── run-comprehensive-crud-test.cmd      # Automated test runner
├── validate-comprehensive-test.cmd      # Validation script
├── COMPREHENSIVE-CRUD-TEST-README.md    # This documentation
├── terraform.tfvars                     # Variable definitions
└── .terraform/                          # Terraform state and plugins
```

## Integration with Existing Tests

This comprehensive test can run alongside existing individual tests:
- `inbound-integration-test.tf` - Focuses on inbound integrations only
- `escalation-policy-test.tf` - Focuses on escalation policies only
- `schedule-test.tf` - Focuses on schedules only
- `workflow-test.tf` - Focuses on workflows only

Each test uses unique suffixes to avoid resource conflicts.

## Best Practices

1. **Always validate first** using `validate-comprehensive-test.cmd`
2. **Use the automated runner** for consistent results
3. **Clean up resources** after testing to avoid API limits
4. **Check outputs** to verify all functionality
5. **Run individual phases** if debugging specific issues

## Support

For issues with the comprehensive test:
1. Check this README for troubleshooting steps
2. Run validation script to identify configuration issues
3. Review individual test files for simpler examples
4. Check AlertOps API documentation for resource requirements 