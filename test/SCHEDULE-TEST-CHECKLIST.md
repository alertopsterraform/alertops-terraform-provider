# AlertOps Schedule Resource Test Checklist

## Quick Validation Checklist

### ✅ Pre-Test Setup
- [ ] API key set: `echo %TF_VAR_alertops_api_key%`
- [ ] Provider built and installed in plugin directory
- [ ] .terraformrc dev_overrides configured
- [ ] Test directory clean (no terraform.tfstate)
- [ ] Schedule test configuration in place

### ✅ Phase 2: User Resource Testing (Foundation)
- [ ] `terraform plan -target=alertops_user.schedule_test_user` succeeds
- [ ] `terraform apply -target=alertops_user.schedule_test_user` succeeds
- [ ] User ID is a valid integer (not null/empty)
- [ ] Username matches: `terraform-schedule-user`
- [ ] Contact methods created: Email-Official
- [ ] Output `schedule_test_results.user_name` shows correct data

### ✅ Phase 3: Group Resource Testing (Prerequisites)  
- [ ] `terraform plan -target=alertops_group.schedule_test_group` succeeds
- [ ] `terraform apply -target=alertops_group.schedule_test_group` succeeds
- [ ] Group ID is a valid integer (not null/empty)
- [ ] Group name matches: `terraform-schedule-group`
- [ ] User referenced as member correctly
- [ ] Member roles: ["Primary"]
- [ ] Output `schedule_test_results.group_name` shows correct data

### ✅ Phase 4: Schedule Resource Testing (Main Focus)
- [ ] `terraform plan -target=alertops_schedule.simple_schedule` succeeds
- [ ] `terraform apply -target=alertops_schedule.simple_schedule` succeeds
- [ ] Schedule ID is a valid integer (not null/empty)
- [ ] Schedule name matches: `terraform-test-schedule`
- [ ] Group reference works correctly (using group_id, not group_name)
- [ ] Schedule type: "Fixed"
- [ ] Timezone: "(UTC-06:00) Central Time (US & Canada)"
- [ ] Color: "#FF5733"
- [ ] Start weekday: "Mon"
- [ ] End weekday: "Sun"
- [ ] Schedule weekdays: All days enabled (Sun-Sat = true)
- [ ] User assignment works (user + role)
- [ ] Output `schedule_test_results.schedule_name` shows correct data

### ✅ Integration Validation
- [ ] Output `integration_verification.dependency_chain` shows correct flow
- [ ] Output `integration_verification.user_created` matches username
- [ ] Output `integration_verification.group_created` matches group name  
- [ ] Output `integration_verification.schedule_created` matches schedule name
- [ ] Output `integration_verification.group_has_user` matches username
- [ ] Output `integration_verification.schedule_has_user` matches username
- [ ] Debug JSON shows proper API request structure
- [ ] No Terraform state inconsistencies

### ✅ Complete Configuration Test
- [ ] `terraform plan` shows no changes (state matches reality)
- [ ] `terraform apply` completes without errors
- [ ] All outputs display correctly
- [ ] `terraform show` displays all three resources correctly

### ✅ CRUD Operations (Optional)
- [ ] Update user contact method → apply succeeds
- [ ] Update group description → apply succeeds  
- [ ] Update schedule name/color → apply succeeds
- [ ] `terraform destroy` removes resources in correct order (Schedule → Group → User)
- [ ] Clean state after destroy (no resources remain)
- [ ] No orphaned resources in AlertOps dashboard

## Expected Outputs Format

### schedule_test_results
```json
{
  "user_id": 12345,
  "user_name": "terraform-schedule-user",
  "group_id": 67890,
  "group_name": "terraform-schedule-group", 
  "schedule_id": 11111,
  "schedule_name": "terraform-test-schedule"
}
```

### integration_verification
```json
{
  "user_created": "terraform-schedule-user",
  "group_created": "terraform-schedule-group",
  "schedule_created": "terraform-test-schedule",
  "group_has_user": "terraform-schedule-user",
  "schedule_has_user": "terraform-schedule-user",
  "dependency_chain": "terraform-schedule-user -> terraform-schedule-group -> terraform-test-schedule"
}
```

### schedule_debug
```json
{
  "schedule_request_json": "{\"group\":\"terraform-schedule-group\",\"schedule_name\":\"terraform-test-schedule\",...}"
}
```

## Common Issues & Solutions

### ❌ "Invalid API Key" Error
- **Problem**: Wrong header format or invalid key
- **Solution**: Verify API key, check authentication in client.go

### ❌ "Group Not Found" Error  
- **Problem**: Group doesn't exist when creating schedule
- **Solution**: Ensure group created first, check dependency chain
- **Important**: Use group_id (not group_name) in schedule configuration

### ❌ "Invalid Schedule Type" Error
- **Problem**: Using unsupported schedule type
- **Solution**: Verify schedule_type values with AlertOps API documentation

### ❌ "User Not Found in Schedule" Error
- **Problem**: User assignment failing
- **Solution**: Ensure user exists, check username reference

### ❌ Terraform State Issues
- **Problem**: State corruption or inconsistency
- **Solution**: Remove terraform.tfstate, start fresh

### ❌ Provider Not Found Error
- **Problem**: dev_overrides not working
- **Solution**: Check .terraformrc path, rebuild provider

## Success Criteria Summary

**Users Resource:** ✅ CRUD operations working with contact methods  
**Groups Resource:** ✅ CRUD operations working with user members  
**Schedule Resource:** ✅ CRUD operations working with group/user references  
**Integration:** ✅ User-Group-Schedule relationships working correctly  
**API Compatibility:** ✅ All HTTP status codes handled properly (200/201/204)  
**State Management:** ✅ Terraform state consistency maintained  

## Advanced Schedule Features to Test Later

### Complex Rotation Patterns
- [ ] Daily rotation with specific times
- [ ] Weekly rotation with day-of-week
- [ ] Monthly rotation patterns
- [ ] Multiple user rotations

### Date/Time Features
- [ ] Start/end date configurations
- [ ] Weekday schedules (Mon-Fri, etc.)
- [ ] Timezone handling across different zones

### Advanced Configuration
- [ ] Repeat schedule patterns
- [ ] Holiday notification settings
- [ ] Include all group users option
- [ ] Complex user role assignments

## Next Steps After Success

1. **Document any API schema differences discovered**
2. **Proceed to Escalation Policies resource implementation**  
3. **Use same testing pattern for remaining resources**
4. **Build comprehensive end-to-end test scenario**

## Time Tracking

- **Pre-Test Setup**: ~5 minutes
- **User Testing**: ~3 minutes  
- **Group Testing**: ~5 minutes
- **Schedule Testing**: ~10 minutes
- **Integration Validation**: ~5 minutes
- **CRUD Testing**: ~15 minutes (optional)
- **Total**: ~30-45 minutes

**Goal**: Complete Schedule testing in under 30 minutes to stay on track for 2-hour completion target. 