@echo off
echo Setting Terraform CLI config file...
set TF_CLI_CONFIG_FILE=C:\alertopsterraform\test\.terraformrc
echo TF_CLI_CONFIG_FILE is set to: %TF_CLI_CONFIG_FILE%
echo.
echo Running terraform init...
terraform init 