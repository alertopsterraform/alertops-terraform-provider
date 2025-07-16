# Iterative Groups Testing Guide

## Overview
This guide walks through testing the AlertOps Terraform provider's Groups functionality step by step.

## Prerequisites
1. AlertOps API key set in environment: `$env:TF_VAR_alertops_api_key="your-api-key"`
2. Provider built and available via dev_overrides
3. Clean test environment

## Test Structure
The test file `updated-group-test.tf` is organized in phases:

### Phase 1: User Creation
- Creates a single user that will be used as a group member
- Tests basic user CRUD operations
- Validates user contact methods

### Phase 2: Group Creation  
- Creates a group that references the Phase 1 user
- Tests basic group structure (members only) - contact methods, topics, and attributes will be separate resources
- Validates user-group integration

## Step-by-Step Testing

### Step 1: Test User Only (Phase 1)
```bash
cd test
terraform validate
terraform plan
terraform apply
```

**Expected Results:**
- User `terraform-group-test-user` created successfully
- Output shows user_id and user_name
- User has Email-Official and Phone-Official contact methods

### Step 2: Test User + Group (Phase 2)
The group should automatically reference the user from Phase 1.

**Expected Results:**
- Group `terraform-iterative-test-group` created successfully
- Group has the user as a member with roles ["Primary", "Manager"]
- Group created successfully (minimal configuration with user member only)
- Topics and attributes skipped for this phase
- Integration test shows dependency_works = true

### Step 3: Verify Integration
Check the outputs:
```bash
terraform output
```

**Key Validations:**
- `phase1_user_info` shows user details
- `phase2_group_info` shows group details  
- `integration_test.dependency_works` should be `true`
- `debug_info` shows the actual JSON sent to API

### Step 4: Test Updates
Try updating user or group properties and apply changes.

### Step 5: Test Deletion
```bash
terraform destroy
```

**Expected Results:**
- Group deleted first (dependency order)
- User deleted second
- Clean removal with no errors

## Troubleshooting

### If Phase 1 Fails:
- Check API key is correct
- Verify user contact methods format
- Check AlertOps API response in debug output

### If Phase 2 Fails:
- Ensure Phase 1 user exists
- Check group member reference syntax
- Verify contact method types are valid
- Check topics and attributes format

### If Integration Fails:
- Verify user_name matches between resources
- Check dependency chain is working
- Look at debug_request_json for API payload

## Success Criteria
✅ User created successfully in Phase 1  
✅ Group created successfully in Phase 2  
✅ Group correctly references user as member  
✅ All contact methods, topics, and attributes work  
✅ Integration test shows dependency_works = true  
✅ Clean destruction in reverse order  

## Next Steps
Once this passes, we can proceed to implement Schedules resource. 