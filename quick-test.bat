@echo off
echo ===============================
echo AlertOps Provider Quick Test
echo ===============================
echo.

echo Step 1: Building provider...
go build -o terraform-provider-alertops.exe
if %errorlevel% neq 0 (
    echo ERROR: Failed to build provider
    exit /b 1
)

echo Step 2: Installing provider...
mkdir "%APPDATA%\terraform.d\plugins\registry.terraform.io\alertopsterraform\alertops\1.0.0\windows_amd64" 2>nul
copy terraform-provider-alertops.exe "%APPDATA%\terraform.d\plugins\registry.terraform.io\alertopsterraform\alertops\1.0.0\windows_amd64\terraform-provider-alertops.exe" >nul

echo Step 3: Checking API key...
if "%ALERTOPS_API_KEY%"=="" (
    echo ERROR: Please set ALERTOPS_API_KEY environment variable
    echo Example: set ALERTOPS_API_KEY=your-api-key-here
    exit /b 1
)

echo Step 4: Changing to test directory...
cd test

echo Step 5: Initializing Terraform...
terraform init
if %errorlevel% neq 0 (
    echo ERROR: Terraform init failed
    cd ..
    exit /b 1
)

echo Step 6: Validating configuration...
terraform validate
if %errorlevel% neq 0 (
    echo ERROR: Configuration validation failed
    cd ..
    exit /b 1
)

cd ..

echo.
echo ===============================
echo Ready for testing!
echo ===============================
echo.
echo Next steps:
echo 1. cd test
echo 2. Edit simple-test.tf to update user_name
echo 3. Run: terraform plan
echo 4. Run: terraform apply
echo 5. Test updates by modifying simple-test.tf
echo 6. Run: terraform destroy (to clean up)
echo.
echo For detailed testing, see TESTING.md
echo. 