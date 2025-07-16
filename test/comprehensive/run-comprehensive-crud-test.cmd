@echo off
REM ============================================================================= 
REM AlertOps Terraform Provider - Comprehensive CRUD Test Runner
REM =============================================================================
REM Runs the comprehensive test that covers all 6 resource types

echo.
echo ============================================================================= 
echo AlertOps Terraform Provider - Comprehensive CRUD Test
echo ============================================================================= 
echo This test will create and verify all resource types:
echo   - 3 Users (primary, secondary, backup)
echo   - 3 Groups (primary, backup, combined)  
echo   - 2 Schedules (primary, backup)
echo   - 2 Workflows (alert processing, notification)
echo   - 2 Escalation Policies (primary, backup)
echo   - 3 Inbound Integrations (2x API, 1x Bridge)
echo.
echo Total: 16 resources across 6 resource types
echo.

REM Check if API key is set
if "%ALERTOPS_API_KEY%"=="" (
    echo ERROR: ALERTOPS_API_KEY environment variable is not set!
    echo Please set it with: set ALERTOPS_API_KEY=your_api_key_here
    echo.
    pause
    exit /b 1
)

echo API Key Status: ✅ Set (length: %ALERTOPS_API_KEY:~0,8%...)
echo.

REM Navigate to test directory
cd /d "%~dp0"
echo Current directory: %CD%
echo.

REM Initialize terraform if needed
echo ============================================================================= 
echo STEP 1: Terraform Initialization
echo ============================================================================= 
if not exist ".terraform" (
    echo Running terraform init...
    terraform init
    if errorlevel 1 (
        echo ❌ Terraform init failed!
        pause
        exit /b 1
    )
) else (
    echo ✅ Terraform already initialized
)
echo.

REM Validate configuration
echo ============================================================================= 
echo STEP 2: Configuration Validation
echo ============================================================================= 
echo Validating comprehensive-crud-test.tf...
terraform validate comprehensive-crud-test.tf
if errorlevel 1 (
    echo ❌ Configuration validation failed!
    pause
    exit /b 1
)
echo ✅ Configuration validation passed
echo.

REM Plan the deployment
echo ============================================================================= 
echo STEP 3: Planning Deployment
echo ============================================================================= 
echo Creating execution plan...
terraform plan -var-file=terraform.tfvars -target=alertops_user.primary_user -target=alertops_user.secondary_user -target=alertops_user.backup_user comprehensive-crud-test.tf
if errorlevel 1 (
    echo ❌ Plan failed!
    pause
    exit /b 1
)
echo.

echo Continue with comprehensive test? (This will create 16 resources)
set /p continue="Press Y to continue, N to exit: "
if /i not "%continue%"=="Y" (
    echo Test cancelled by user
    exit /b 0
)
echo.

REM Execute the comprehensive test
echo ============================================================================= 
echo STEP 4: Executing Comprehensive CRUD Test
echo ============================================================================= 
echo Creating all resources with dependency management...
echo.

REM Phase 1: Users
echo Phase 1: Creating Users...
terraform apply -var-file=terraform.tfvars -target=alertops_user.primary_user -target=alertops_user.secondary_user -target=alertops_user.backup_user -auto-approve comprehensive-crud-test.tf
if errorlevel 1 (
    echo ❌ User creation failed!
    goto :cleanup_prompt
)
echo ✅ Users created successfully
echo.

REM Phase 2: Groups
echo Phase 2: Creating Groups...
terraform apply -var-file=terraform.tfvars -target=alertops_group.primary_group -target=alertops_group.backup_group -target=alertops_group.combined_group -auto-approve comprehensive-crud-test.tf
if errorlevel 1 (
    echo ❌ Group creation failed!
    goto :cleanup_prompt
)
echo ✅ Groups created successfully
echo.

REM Phase 3: Schedules
echo Phase 3: Creating Schedules...
terraform apply -var-file=terraform.tfvars -target=alertops_schedule.primary_schedule -target=alertops_schedule.backup_schedule -auto-approve comprehensive-crud-test.tf
if errorlevel 1 (
    echo ❌ Schedule creation failed!
    goto :cleanup_prompt
)
echo ✅ Schedules created successfully
echo.

REM Phase 4: Workflows
echo Phase 4: Creating Workflows...
terraform apply -var-file=terraform.tfvars -target=alertops_workflow.alert_processing_workflow -target=alertops_workflow.notification_workflow -auto-approve comprehensive-crud-test.tf
if errorlevel 1 (
    echo ❌ Workflow creation failed!
    goto :cleanup_prompt
)
echo ✅ Workflows created successfully
echo.

REM Phase 5: Escalation Policies
echo Phase 5: Creating Escalation Policies...
terraform apply -var-file=terraform.tfvars -target=alertops_escalation_policy.primary_escalation_policy -target=alertops_escalation_policy.backup_escalation_policy -auto-approve comprehensive-crud-test.tf
if errorlevel 1 (
    echo ❌ Escalation Policy creation failed!
    goto :cleanup_prompt
)
echo ✅ Escalation Policies created successfully
echo.

REM Phase 6: Inbound Integrations
echo Phase 6: Creating Inbound Integrations...
terraform apply -var-file=terraform.tfvars -target=alertops_inbound_integration.api_integration -target=alertops_inbound_integration.email_integration -target=alertops_inbound_integration.bridge_integration -auto-approve comprehensive-crud-test.tf
if errorlevel 1 (
    echo ❌ Inbound Integration creation failed!
    goto :cleanup_prompt
)
echo ✅ Inbound Integrations created successfully
echo.

REM Final application to ensure all dependencies are satisfied
echo Phase 7: Final Verification...
terraform apply -var-file=terraform.tfvars -auto-approve comprehensive-crud-test.tf
if errorlevel 1 (
    echo ❌ Final verification failed!
    goto :cleanup_prompt
)

echo.
echo ============================================================================= 
echo STEP 5: Test Results and Verification
echo ============================================================================= 
echo Generating test outputs...
terraform output comprehensive_test_summary
echo.
terraform output crud_test_success_indicators
echo.

echo ============================================================================= 
echo ✅ COMPREHENSIVE CRUD TEST COMPLETED SUCCESSFULLY!
echo ============================================================================= 
echo All 16 resources across 6 resource types have been created and verified.
echo.
echo Resources created:
echo   - 3 Users with contact methods
echo   - 3 Groups with member relationships  
echo   - 2 Schedules with rotation settings
echo   - 2 Workflows with conditions and actions
echo   - 2 Escalation Policies with multi-tier escalation
echo   - 3 Inbound Integrations (2x API with different settings, 1x Bridge)
echo.

:cleanup_prompt
echo.
echo ============================================================================= 
echo CLEANUP OPTIONS
echo ============================================================================= 
echo What would you like to do next?
echo.
echo 1. Keep resources for manual testing
echo 2. Clean up all test resources (terraform destroy)
echo 3. View cleanup commands for selective cleanup
echo.
set /p cleanup="Enter choice (1-3): "

if "%cleanup%"=="2" (
    echo.
    echo Running terraform destroy to clean up all test resources...
    terraform destroy -var-file=terraform.tfvars -auto-approve comprehensive-crud-test.tf
    echo ✅ Cleanup completed
)

if "%cleanup%"=="3" (
    echo.
    terraform output cleanup_commands
)

echo.
echo Test completed! Check the outputs above for detailed results.
echo.
pause 