# CRUD Testing Guide for AlertOps Terraform Provider

## ✅ You've successfully tested: CREATE
- User was created with "Basic" role
- Contact methods were added
- All fields were properly set

## 🔄 Now let's test: UPDATE

### Step 1: Make a change to your configuration
Edit `simple-test.tf` and modify one or more fields. Good options:

```hcl
# Option A: Change the user's name
first_name = "Updated-Terraform"
last_name  = "UpdatedTest"

# Option B: Change the role
roles = ["Admin"]  # or ["User"] or another valid role

# Option C: Add another contact method
contact_methods {
  contact_method_name = "Email-Official"
  email {
    email_address = "terraform-test@example.com"
  }
  enabled = true
  sequence = 1
}
contact_methods {
  contact_method_name = "Phone-Official"
  phone {
    country_code = "1"
    phone_number = "555-123-4567"
  }
  enabled = true
  sequence = 2
}

# Option D: Change time zone
time_zone = "(UTC-05:00) Eastern Time (US & Canada)"
```

### Step 2: Test the UPDATE
```bash
terraform-dev.bat plan    # Should show only the changes you made
terraform-dev.bat apply   # Apply the updates
```

**What to verify:**
- ✅ Only changed fields are shown in the plan
- ✅ `user_id` stays the same
- ✅ Update succeeds without recreating the resource
- ✅ Changes appear in AlertOps dashboard

## 📖 Testing: READ/RETRIEVAL

### Step 3: Test refresh operation
```bash
# Refresh state from API (tests READ operation)
set TF_CLI_CONFIG_FILE=C:\alertopsterraform\test\.terraformrc
terraform refresh

# Show current state
terraform show
```

**What to verify:**
- ✅ Data is accurately retrieved from AlertOps API
- ✅ All fields match what's in AlertOps
- ✅ Computed fields (like `last_login_date`) are populated

### Step 4: Test data source (if implemented)
Create a new file `test-datasource.tf`:

```hcl
# Test data source lookup
data "alertops_user" "lookup_test" {
  user_name = "terraform-simple-test"  # The user we created
}

output "retrieved_user" {
  value = data.alertops_user.lookup_test
}
```

Then run:
```bash
terraform-dev.bat plan    # Should show data source will be read
terraform-dev.bat apply   # Execute the data source lookup
```

## 🗑️ Testing: DELETE

### Step 5: Test destroy operation
```bash
terraform-dev.bat destroy
```

**What to verify:**
- ✅ User is deleted from AlertOps dashboard
- ✅ Terraform state becomes empty
- ✅ No errors during deletion

## 🔄 Additional Tests

### Test IMPORT functionality
```bash
# Create a user manually in AlertOps dashboard first
# Then import it into Terraform:

terraform import alertops_user.imported_user [USER_ID_FROM_ALERTOPS]
terraform show  # Verify import worked
terraform plan  # Should show no changes if import was perfect
```

### Test validation errors
```bash
# Test invalid contact method
roles = ["InvalidRole"]  # Should fail validation
terraform-dev.bat validate  # Should show error

# Test invalid contact method type
contact_methods {
  contact_method_name = "Invalid-Type"  # Should fail
  email {
    email_address = "test@example.com"
  }
}
```

## 🐛 Debugging Tips

### Enable debug logging:
```bash
set TF_LOG=DEBUG
terraform-dev.bat plan
```

### Check API calls:
Look for HTTP requests/responses in the debug output to see exactly what's being sent to AlertOps.

### State inspection:
```bash
terraform show -json > state.json  # Export state for inspection
terraform state list              # List all resources in state
terraform state show alertops_user.simple_test  # Show specific resource
```

## ✅ Success Criteria

**UPDATE Tests:**
- [ ] Field changes are detected correctly
- [ ] Only modified fields are updated
- [ ] User ID remains constant
- [ ] Complex nested changes (contact methods) work

**READ Tests:**
- [ ] Refresh retrieves current data from API
- [ ] All fields match AlertOps dashboard
- [ ] Data source lookup works (if implemented)

**DELETE Tests:**
- [ ] Resource is cleanly removed from AlertOps
- [ ] Terraform state is cleaned up
- [ ] No orphaned data remains

**Additional Tests:**
- [ ] Import works correctly
- [ ] Validation catches errors
- [ ] Debug logging shows API calls
- [ ] State management is clean 