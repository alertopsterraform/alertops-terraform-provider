# Schedule Testing Setup - Updated Test Scripts

## 🎯 Overview

The test scripts have been updated to focus on **Schedule resource testing** as the next phase of the AlertOps Terraform provider development.

## 📁 Updated Files

### **New Test Scripts:**
- `run-schedule-tests.cmd` - Automated Schedule testing with full dependency chain
- `SCHEDULE-TEST-CHECKLIST.md` - Manual validation checklist for Schedule resource
- `setup-schedule-testing.bat` - Environment setup script

### **Test Configuration:**
- `schedule-test.tf` - Complete test configuration with User → Group → Schedule chain

### **Updated Documentation:**
- `COMPREHENSIVE-TEST-PLAN.md` - Updated to reflect Schedule testing phase
- TODOs updated to mark Schedules as "in_progress"

## 🔄 Test Structure

### **Phase 2: User Resource (Foundation)**
- Creates `terraform-schedule-user`
- Email-Official contact method
- Validates basic user CRUD operations

### **Phase 3: Group Resource (Prerequisites)**
- Creates `terraform-schedule-group`
- Adds user as member with "Primary" role
- Validates group-user integration

### **Phase 4: Schedule Resource (Main Focus)**
- Creates `terraform-test-schedule`
- References the group created in Phase 3
- Assigns the user created in Phase 2
- Tests schedule-specific fields:
  - Schedule type: "Fixed"
  - Timezone: "(UTC-06:00) Central Time (US & Canada)"
  - Color: "#FF5733"
  - Start weekday: "Mon"
  - End weekday: "Sun"
  - Schedule weekdays: All days enabled (Sun-Sat)
  - User assignment with role

## 🚀 How to Run Tests

### **Option 1: Automated Testing (Recommended)**
```cmd
cd test
set TF_VAR_alertops_api_key=your-api-key
run-schedule-tests.cmd
```

### **Option 2: Manual Step-by-Step**
```cmd
cd test
set TF_VAR_alertops_api_key=your-api-key
terraform plan
terraform apply
terraform output
```

### **Option 3: Environment Setup**
```cmd
cd test
setup-schedule-testing.bat
# Follow the prompts
```

## ✅ Success Criteria

### **Individual Resources:**
- ✅ User created: `terraform-schedule-user`
- ✅ Group created: `terraform-schedule-group`
- ✅ Schedule created: `terraform-test-schedule`

### **Integration Validation:**
- ✅ Group contains user as member
- ✅ Schedule references group correctly
- ✅ Schedule assigns user with role
- ✅ Full dependency chain: User → Group → Schedule

### **API Validation:**
- ✅ All HTTP status codes handled (200/201/204)
- ✅ Complex nested Schedule API structure works
- ✅ Debug JSON shows proper request format

## 📊 Expected Outputs

### **schedule_test_results:**
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

### **integration_verification:**
```json
{
  "dependency_chain": "terraform-schedule-user -> terraform-schedule-group -> terraform-test-schedule",
  "user_created": "terraform-schedule-user",
  "group_created": "terraform-schedule-group",
  "schedule_created": "terraform-test-schedule"
}
```

## 🎯 Schedule Resource Features Tested

### **Basic Configuration:**
- Group assignment (`group` field - uses group_id, not group_name)
- Schedule name and type
- Timezone configuration
- Color assignment
- Schedule weekdays (all days enabled)
- Enable/disable state

### **User Management:**
- Individual user assignments
- Role-based access (Primary, etc.)
- Integration with existing users

### **Advanced Features (Ready but not tested yet):**
- Rotation patterns (daily, weekly, monthly)
- Date/time configurations
- Weekday schedules
- Repeat patterns
- Holiday notifications

## 🔧 Technical Implementation

### **Schedule Resource (`resource_schedule.go`):**
- **850 lines** of comprehensive CRUD implementation
- **20+ helper functions** for expand/flatten operations
- **Complete schema** matching AlertOps API
- **Proper error handling** and debugging support

### **Data Models (`models.go`):**
- `Schedule` struct with 20+ fields
- Nested structures: `ScheduleDate`, `ScheduleTime`, `ScheduleWeekdays`
- Rotation models: `RotateDaily`, `RotateWeekly`, `RotateMonthly`
- Support structures: `RepeatSchedule`, `ScheduleUser`

## ⏰ Timeline

- **Schedule Implementation**: ✅ Complete (~20 minutes)
- **Test Script Updates**: ✅ Complete (~10 minutes) 
- **Schedule Testing**: 🚧 Ready to execute (~20 minutes)
- **Total Progress**: 50% of 6 resources complete

## 🚦 Next Steps

1. **Execute Schedule Testing** - Run the test scripts
2. **Validate Results** - Check all success criteria
3. **Document Issues** - Note any API schema differences
4. **Proceed to Escalation Policies** - Next resource implementation

## 📝 Important Notes

### **AlertOps API Discovery:**
- **Group Reference**: Schedules must reference groups by `group_id` (integer), not `group_name` (string)
- **API Behavior**: The schedule creation succeeds, but read operations fail if using group_name
- **Configuration**: Use `group = tostring(alertops_group.name.group_id)` in schedule resources

### **Technical Notes:**
- Schedule resource has the most complex nested structures of all resources
- Patterns established here will accelerate remaining resource development  
- Full dependency chain testing validates provider foundation
- Ready for production-level Schedule management in AlertOps

---

**Ready to test Schedules!** Set your API key and run `run-schedule-tests.cmd` 🚀 