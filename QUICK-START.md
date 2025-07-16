# Quick Start Testing Guide

## 🚀 **Ready to Test!**

Your AlertOps provider is now configured for the alertopsterraform GitHub account. Here's how to test:

## Step 1: Set Your API Key
```bash
set ALERTOPS_API_KEY=your-actual-api-key-here
```

## Step 2: Navigate to Test Directory
```bash
cd test
```

## Step 3: Test Basic User Creation

### Initialize and Validate:
```bash
terraform init
terraform validate
```

### Test with Simple Configuration:
```bash
# Test the simple configuration
terraform validate simple-test.tf

# If you have a real API key, test planning:
terraform plan -target=alertops_user.simple_test
```

### Update Test User Details:
Edit `simple-test.tf` and change:
- `user_name` to something unique (e.g., `terraform-test-yourname`)
- Verify other details are correct

### Test CRUD Operations:
```bash
# CREATE - Create the user
terraform apply -target=alertops_user.simple_test

# READ - Check the created user
terraform show alertops_user.simple_test

# UPDATE - Edit simple-test.tf (change first_name) then:
terraform plan
terraform apply

# DELETE - Clean up
terraform destroy
```

## Step 4: Verify in AlertOps Dashboard

1. Log into AlertOps
2. Check Users section
3. Verify test user was created/updated/deleted

## 📋 **What This Tests:**

✅ **Provider Integration**: Confirms provider loads and connects to API  
✅ **Basic CRUD**: Create, Read, Update, Delete operations  
✅ **API Authentication**: Verifies your API key works  
✅ **Field Mapping**: Tests user_name, first_name, last_name, etc.  

## 🔧 **Troubleshooting:**

**"Missing required provider"** → Run `terraform init` to install the provider

**"Failed to install provider"** → Check provider is built and in correct directory

**"API authentication failed"** → Check your `ALERTOPS_API_KEY` is correct

**"User already exists"** → Change `user_name` to something unique

## 📈 **Next Steps:**

Once basic testing works:

1. **Test Contact Methods**: We'll need to align our schema with the existing provider
2. **Add More Objects**: Groups, Schedules, etc.
3. **Schema Investigation**: Compare our implementation with existing provider

## 🎯 **Current Status:**

✅ **Provider**: Configured for alertopsterraform GitHub account  
✅ **Basic Users**: CREATE/READ/UPDATE/DELETE operations  
✅ **Contact Methods**: Implemented with validation  
✅ **Data Sources**: User lookup functionality  
⏳ **Other Objects**: Groups, Schedules, etc. (next phase)  

## 💡 **Pro Tip:**

Make sure to run `terraform init` to install the provider before testing! 