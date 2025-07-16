@echo off
echo Setting up Terraform for local development...
set TF_CLI_CONFIG_FILE=C:\alertopsterraform\test\.terraformrc

echo.
echo Available commands:
echo   %0 validate   - Validate the configuration
echo   %0 plan       - Show execution plan
echo   %0 apply      - Apply the configuration (with approval)
echo   %0 destroy    - Destroy the resources (with approval)
echo.

if "%1"=="validate" (
    echo Running terraform validate...
    terraform validate
) else if "%1"=="plan" (
    echo Running terraform plan...
    terraform plan
) else if "%1"=="apply" (
    echo Running terraform apply...
    terraform apply
) else if "%1"=="destroy" (
    echo Running terraform destroy...
    terraform destroy
) else (
    echo Please specify a command: validate, plan, apply, or destroy
    echo Example: %0 plan
) 