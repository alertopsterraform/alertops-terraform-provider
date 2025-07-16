@echo off
echo Setting AlertOps API key...
REM Replace YOUR_API_KEY_HERE with your actual AlertOps API key
set TF_VAR_alertops_api_key=YOUR_API_KEY_HERE
echo API key set. You can now run terraform commands.
echo.
echo Usage:
echo   terraform plan
echo   terraform apply
echo   terraform destroy 