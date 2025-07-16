@echo off
echo ========================================
echo Setting up Schedule Testing Environment
echo ========================================
echo.

REM Check if we're in the test directory
if not exist "schedule-test.tf" (
    echo ❌ ERROR: schedule-test.tf not found
    echo Please run this script from the test directory
    pause
    exit /b 1
)

echo ✅ Found schedule-test.tf configuration
echo.

REM Clean up any existing state
echo Cleaning up existing Terraform state...
del terraform.tfstate 2>nul
del terraform.tfstate.backup 2>nul
del .terraform.lock.hcl 2>nul
rmdir /s /q .terraform 2>nul
echo ✅ Cleaned up existing state

echo.
echo ========================================
echo Schedule Testing Environment Ready!
echo ========================================
echo.
echo Available test scripts:
echo  - run-schedule-tests.cmd     (Automated full testing)
echo  - SCHEDULE-TEST-CHECKLIST.md (Manual checklist)
echo.
echo Test configuration:
echo  - User: terraform-schedule-user
echo  - Group: terraform-schedule-group  
echo  - Schedule: terraform-test-schedule
echo.
echo To start testing:
echo  1. Set API key: set TF_VAR_alertops_api_key=your-api-key
echo  2. Run: run-schedule-tests.cmd
echo.
echo Or test manually:
echo  1. terraform plan
echo  2. terraform apply
echo  3. terraform output
echo.
pause 