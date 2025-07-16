# AlertOps Provider Test Checklist

## Quick Validation Checklist

### ✅ Pre-Test Setup
- [ ] API key set: `echo %TF_VAR_alertops_api_key%`
- [ ] Provider built and installed in plugin directory
- [ ] .terraformrc dev_overrides configured
- [ ] Test directory clean (no terraform.tfstate)

### ✅ Phase 1A: User Resource Testing
- [ ] `terraform plan -target=alertops_user.test_user_for_group` succeeds
- [ ] `terraform apply -target=alertops_user.test_user_for_group` succeeds
- [ ] User ID is a valid integer (not null/empty)
- [ ] Username matches: `terraform-group-test-user`
- [ ] Contact methods created: Email-Official, Phone-Official
- [ ] Output `phase1_user_info` shows correct data

### ✅ Phase 1B: Group Resource Testing  
- [ ] `terraform plan -target=alertops_group.iterative_test_group` succeeds
- [ ] `terraform apply -target=alertops_group.iterative_test_group` succeeds
- [ ] Group ID is a valid integer (not null/empty)
- [ ] Group name matches: `terraform-iterative-test-group`
- [ ] User referenced as member correctly
- [ ] Member roles: ["Primary", "Manager"]
- [ ] Group created successfully (minimal configuration)
- [ ] Topics and attributes skipped (will be separate resources later)
- [ ] Output `phase2_group_info` shows correct data

### ✅ Integration Validation
- [ ] Output `integration_test.dependency_works` = true
- [ ] Output `integration_test.user_created_first` matches username
- [ ] Output `integration_test.group_references_user` matches username
- [ ] Debug JSON shows proper API request structure
- [ ] No Terraform state inconsistencies

### ✅ Complete Configuration Test
- [ ] `terraform plan` shows no changes (state matches reality)
- [ ] `terraform apply` completes without errors
- [ ] All outputs display correctly
- [ ] `terraform show` displays both resources correctly

### ✅ CRUD Operations (Optional)
- [ ] Update user contact method → apply succeeds
- [ ] Update group attributes → apply succeeds  
- [ ] `terraform destroy` removes group first, then user
- [ ] Clean state after destroy (no resources remain)
- [ ] No orphaned resources in AlertOps dashboard

## Expected Outputs Format

### phase1_user_info
```json
{
  "user_id": 12345,
  "user_name": "terraform-group-test-user", 
  "full_name": "Group TestUser"
}
```

### phase2_group_info
```json
{
  "group_id": 67890,
  "group_name": "terraform-iterative-test-group",
  "dynamic": false
}
```

### integration_test
```json
{
  "user_created_first": "terraform-group-test-user",
  "group_references_user": "terraform-group-test-user", 
  "dependency_works": true
}
```

## Common Issues & Solutions

### ❌ "Invalid API Key" Error
- **Problem**: Wrong header format or invalid key
- **Solution**: Verify API key, check authentication in client.go

### ❌ "Resource Not Found" Error  
- **Problem**: User doesn't exist when creating group
- **Solution**: Ensure user created first, check dependency chain

### ❌ "Invalid Contact Method Type" Error
- **Problem**: Using unsupported contact method type
- **Solution**: Use only: Email-Official, Phone-Official, SMS-Official, etc.

### ❌ Terraform State Issues
- **Problem**: State corruption or inconsistency
- **Solution**: Remove terraform.tfstate, start fresh

### ❌ Provider Not Found Error
- **Problem**: dev_overrides not working
- **Solution**: Check .terraformrc path, rebuild provider

## Success Criteria Summary

**Users Resource:** ✅ CRUD operations working with complex contact methods  
**Groups Resource:** ✅ CRUD operations working with complex nested structures  
**Integration:** ✅ User-Group relationships working correctly  
**API Compatibility:** ✅ All HTTP status codes handled properly (200/201/204)  
**State Management:** ✅ Terraform state consistency maintained  

## Next Steps After Success

1. **Document any API schema differences discovered**
2. **Proceed to Schedules resource implementation**  
3. **Use same testing pattern for remaining resources**
4. **Build comprehensive end-to-end test scenario**

## Time Tracking

- **Pre-Test Setup**: ~5 minutes
- **User Testing**: ~5 minutes  
- **Group Testing**: ~10 minutes
- **Integration Validation**: ~5 minutes
- **CRUD Testing**: ~15 minutes (optional)
- **Total**: ~25-40 minutes

**Goal**: Complete Phase 1 testing in under 30 minutes to stay on track for 2-hour completion target. 