# Testing Guide for AlertOps Terraform Provider

## Important: Dev Overrides Behavior

**⚠️ KEY INSIGHT: When using dev_overrides, DO NOT run `terraform init`!**

Terraform's dev_overrides feature bypasses the normal provider installation process. Running `terraform init` will cause errors because it tries to download the provider from the registry, which doesn't exist for our local development version.

## Quick Testing Steps

1. **Build the provider** (if not already done):
   ```bash
   go build -o terraform-provider-alertops.exe
   mkdir -p "$APPDATA/terraform.d/plugins/alertopsterraform/alertops/1.0.0/windows_amd64"
   copy terraform-provider-alertops.exe "$APPDATA/terraform.d/plugins/alertopsterraform/alertops/1.0.0/windows_amd64/"
   ```

2. **Use the convenient batch file**:
   ```bash
   cd test
   terraform-dev.bat validate    # Validate configuration
   terraform-dev.bat plan        # Show execution plan  
   terraform-dev.bat apply       # Apply changes
   ```

3. **Manual approach** (if preferred):
   ```bash
   cd test
   set TF_CLI_CONFIG_FILE=C:\alertopsterraform\test\.terraformrc
   terraform validate
   terraform plan
   terraform apply
   ```

## Configuration Files

- `test/.terraformrc` - Dev overrides configuration
- `test/simple-test.tf` - Simple user creation test
- `test/terraform.tfvars` - API key configuration
- `test/terraform-dev.bat` - Convenient testing script

## Expected Behavior

✅ **Correct**: You should see this warning when everything works:
```
Warning: Provider development overrides are in effect
The following provider development overrides are set in the CLI configuration:
 - alertopsterraform/alertops in C:\Users\kamal\AppData\Roaming\terraform.d\plugins\alertopsterraform\alertops\1.0.0\windows_amd64
```

❌ **Incorrect**: If you see this error, dev_overrides aren't working:
```
Error: Failed to query available provider packages
Could not retrieve the list of available versions for provider alertopsterraform/alertops
```

## Troubleshooting

1. **Ensure .terraformrc is configured correctly**:
   ```hcl
   provider_installation {
     dev_overrides {
       "alertopsterraform/alertops" = "C:\\Users\\kamal\\AppData\\Roaming\\terraform.d\\plugins\\alertopsterraform\\alertops\\1.0.0\\windows_amd64"
     }
     direct {}
   }
   ```

2. **Verify provider binary exists**:
   ```bash
   dir "C:\Users\kamal\AppData\Roaming\terraform.d\plugins\alertopsterraform\alertops\1.0.0\windows_amd64\terraform-provider-alertops.exe"
   ```

3. **Set environment variable correctly**:
   ```bash
   set TF_CLI_CONFIG_FILE=C:\alertopsterraform\test\.terraformrc
   ```

4. **Skip terraform init** - Go directly to validate/plan/apply

## API Testing

The provider requires a valid AlertOps API key. Set it in `test/terraform.tfvars`:
```hcl
alertops_api_key = "your-api-key-here"
```

Test with a simple user creation to verify API connectivity. 