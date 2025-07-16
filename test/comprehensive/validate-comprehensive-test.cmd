@echo off
REM ============================================================================= 
REM AlertOps Terraform Provider - Comprehensive Test Validator
REM =============================================================================
REM Validates the comprehensive CRUD test configuration without creating resources

echo.
echo ============================================================================= 
echo AlertOps Terraform Provider - Configuration Validator
echo ============================================================================= 
echo Validating comprehensive-crud-test.tf configuration...
echo.

REM Navigate to test directory
cd /d "%~dp0"
echo Current directory: %CD%
echo.

REM Check if configuration file exists
if not exist "comprehensive-crud-test.tf" (
    echo ❌ comprehensive-crud-test.tf not found!
    pause
    exit /b 1
)
echo ✅ Configuration file found

REM Initialize terraform if needed
if not exist ".terraform" (
    echo Initializing Terraform...
    terraform init
    if errorlevel 1 (
        echo ❌ Terraform init failed!
        pause
        exit /b 1
    )
)
echo ✅ Terraform initialized

REM Validate syntax
echo.
echo ============================================================================= 
echo STEP 1: Syntax Validation
echo ============================================================================= 
terraform validate comprehensive-crud-test.tf
if errorlevel 1 (
    echo ❌ Syntax validation failed!
    pause
    exit /b 1
)
echo ✅ Syntax validation passed

REM Format check
echo.
echo ============================================================================= 
echo STEP 2: Format Check
echo ============================================================================= 
terraform fmt -check comprehensive-crud-test.tf
if errorlevel 1 (
    echo ⚠️  Formatting issues found. Running terraform fmt...
    terraform fmt comprehensive-crud-test.tf
    echo ✅ Formatting fixed
) else (
    echo ✅ Formatting is correct
)

REM Plan validation (dry run)
echo.
echo ============================================================================= 
echo STEP 3: Plan Validation
echo ============================================================================= 
echo Creating plan to validate resource dependencies...

REM Check if API key is available for planning
if "%ALERTOPS_API_KEY%"=="" (
    echo ⚠️  ALERTOPS_API_KEY not set - skipping plan validation
    echo   (This is OK for syntax checking)
) else (
    echo ✅ API key available - running plan validation...
    terraform plan -var-file=terraform.tfvars comprehensive-crud-test.tf -out=validation.tfplan
    if errorlevel 1 (
        echo ❌ Plan validation failed!
        pause
        exit /b 1
    )
    echo ✅ Plan validation passed
    
    REM Clean up plan file
    if exist "validation.tfplan" del validation.tfplan
)

echo.
echo ============================================================================= 
echo CONFIGURATION SUMMARY
echo ============================================================================= 
echo Analyzing configuration structure...

REM Count resources
for /f %%i in ('findstr /c:"resource \"alertops_" comprehensive-crud-test.tf') do set resource_count=%%i
echo Resources defined: %resource_count%

REM List resource types
echo.
echo Resource types found:
findstr "resource \"alertops_" comprehensive-crud-test.tf | findstr /o ".*"

echo.
echo ============================================================================= 
echo ✅ VALIDATION COMPLETED SUCCESSFULLY!
echo ============================================================================= 
echo The comprehensive-crud-test.tf configuration is valid and ready to use.
echo.
echo Next steps:
echo 1. Ensure ALERTOPS_API_KEY environment variable is set
echo 2. Review terraform.tfvars for any required variables
echo 3. Run: run-comprehensive-crud-test.cmd
echo.
pause 