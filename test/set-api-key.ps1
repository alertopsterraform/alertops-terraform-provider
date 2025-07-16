# AlertOps API Key Setup
# Replace YOUR_API_KEY_HERE with your actual AlertOps API key
$env:TF_VAR_alertops_api_key = "YOUR_API_KEY_HERE"

Write-Host "AlertOps API key has been set for this session."
Write-Host ""
Write-Host "You can now run:"
Write-Host "  terraform plan"
Write-Host "  terraform apply" 
Write-Host "  terraform destroy" 