# AlertOps Terraform Provider - Comprehensive Test Plan

## Executive Summary

This document outlines the complete testing strategy for the AlertOps Terraform provider implementation. The goal is to deliver a production-ready provider supporting 6 core resources within 2 hours while maintaining high quality and reliability.

## Current Status Overview

### ✅ Completed Resources (33%)
- **Users Resource**: Full CRUD operations, complex contact methods, proper authentication
- **Groups Resource**: Full CRUD operations, basic structure with user members (topics/attributes are separate resources)

### 🚧 Current Phase (50%)
- **Schedules Resource**: Implementation complete, ready for testing
- **Escalation Policies Resource**: Pending implementation  
- **Workflows Resource**: Pending implementation
- **Inbound Integrations Resource**: Pending implementation
- **Services Resource**: Pending implementation

## Phase 1: Immediate Testing (Groups + Users Integration)

### 1.1 Pre-Test Validation
**Timeline**: 5 minutes
**Objective**: Ensure test environment is ready

```bash
# Environment Checks
✅ AlertOps API key set: $env:TF_VAR_alertops_api_key
✅ Provider built and installed via dev_overrides
✅ .terraformrc configured correctly
✅ Test directory clean and ready
```

**Validation Commands:**
```bash
cd test
terraform validate
terraform version
echo $env:TF_VAR_alertops_api_key | measure-object -character
```

### 1.2 Phase 1A: User Resource Testing
**Timeline**: 5 minutes  
**Objective**: Validate user creation as foundation for groups

**Test Steps:**
1. **Create User Resource**
   ```bash
   terraform plan -target=alertops_user.test_user_for_group
   terraform apply -target=alertops_user.test_user_for_group
   ```

2. **Validation Criteria:**
   - ✅ User created with correct username: `terraform-group-test-user`
   - ✅ Contact methods properly configured (Email-Official, Phone-Official)
   - ✅ User ID returned and stored in state
   - ✅ API response shows HTTP 200/201 status

### 1.3 Phase 1B: Group Resource Testing  
**Timeline**: 10 minutes
**Objective**: Validate complex group structure and user integration

**Test Steps:**
1. **Create Group Resource**
   ```bash
   terraform plan -target=alertops_group.iterative_test_group
   terraform apply -target=alertops_group.iterative_test_group
   ```

2. **Validation Criteria:**
   - ✅ Group created with correct name: `terraform-iterative-test-group`
   - ✅ User properly referenced as member by username
   - ✅ Member roles correctly set: ["Primary", "Manager"]
   - ✅ Group created successfully (minimal configuration)
   - ✅ Topics and attributes skipped (will be separate resources)
   - ✅ Group ID returned and stored in state

### 1.4 Integration Validation
**Timeline**: 5 minutes
**Objective**: Verify user-group relationship integrity

**Test Commands:**
```bash
terraform output phase1_user_info
terraform output phase2_group_info  
terraform output integration_test
terraform output debug_info
```

**Success Criteria:**
- ✅ `integration_test.dependency_works = true`
- ✅ User and group IDs are valid integers
- ✅ Group member reference matches user username exactly
- ✅ No Terraform state inconsistencies

### 1.5 CRUD Operations Testing
**Timeline**: 15 minutes
**Objective**: Validate full lifecycle management

**Update Testing:**
1. Modify user contact method
2. Add/remove group member roles
3. Update group attributes
4. Apply changes and verify

**Destroy Testing:**
```bash
terraform destroy
```
- ✅ Group destroyed first (dependency order)
- ✅ User destroyed second  
- ✅ Clean state file after destruction
- ✅ No orphaned resources in AlertOps

## Phase 2: Resource Implementation Testing (Schedules → Services)

### 2.1 Schedules Resource Testing
**Timeline**: 20 minutes
**Dependencies**: Users, Groups resources working

**Implementation Strategy:**
1. **API Discovery**: Research AlertOps Schedules API structure
2. **Schema Mapping**: Create Go models matching API schema
3. **Resource Implementation**: CRUD operations with proper validation
4. **Integration Testing**: Test schedules referencing users/groups

**Test Scenarios:**
- Basic schedule creation with time periods
- Schedule with user assignments
- Schedule with group assignments  
- Schedule with complex rotation patterns
- Schedule updates and modifications

### 2.2 Escalation Policies Resource Testing
**Timeline**: 20 minutes  
**Dependencies**: Users, Groups, Schedules resources working

**Test Scenarios:**
- Basic escalation policy creation
- Multi-level escalation (Level 1 → Level 2 → Level 3)
- Integration with schedules for on-call routing
- Timeout configurations and retry logic
- Policy updates and rule modifications

### 2.3 Workflows Resource Testing
**Timeline**: 15 minutes
**Dependencies**: Users, Groups, Escalation Policies working

**Test Scenarios:**
- Basic workflow creation
- Workflow with escalation policy integration
- Conditional logic and branching
- Workflow triggers and conditions
- Workflow updates and modifications

### 2.4 Inbound Integrations Resource Testing  
**Timeline**: 15 minutes
**Dependencies**: Workflows resource working

**Test Scenarios:**
- Basic integration creation (email, webhook, API)
- Integration with workflow assignment
- Authentication and security settings
- Integration configuration updates
- Integration enable/disable functionality

### 2.5 Services Resource Testing
**Timeline**: 15 minutes
**Dependencies**: All previous resources working

**Test Scenarios:**
- Basic service creation
- Service with integration assignments
- Service with escalation policy assignments
- Service hierarchy and dependencies
- Service configuration updates

## Phase 3: Cross-Resource Integration Testing

### 3.1 Complete Workflow Testing
**Timeline**: 20 minutes
**Objective**: Test end-to-end AlertOps configuration via Terraform

**Test Scenario: "Complete Incident Management Setup"**
```hcl
# Create comprehensive test covering all resources
resource "alertops_user" "oncall_engineer" { ... }
resource "alertops_group" "engineering_team" { ... }
resource "alertops_schedule" "oncall_schedule" { ... }
resource "alertops_escalation_policy" "critical_escalation" { ... }
resource "alertops_workflow" "incident_workflow" { ... }
resource "alertops_inbound_integration" "monitoring_webhook" { ... }
resource "alertops_service" "production_service" { ... }
```

**Validation Criteria:**
- ✅ All resources create successfully in dependency order
- ✅ Cross-resource references work correctly
- ✅ Complex nested structures handle properly
- ✅ Update operations maintain referential integrity
- ✅ Destroy operations clean up in reverse order

### 3.2 Data Sources Testing
**Timeline**: 15 minutes
**Objective**: Validate data source functionality for resource lookups

**Test Scenarios:**
- Look up existing users by username/email
- Look up existing groups by name
- Look up schedules, policies, workflows by name/ID
- Validate data source outputs match resource attributes

## Phase 4: Edge Cases and Error Handling

### 4.1 API Error Handling Testing
**Timeline**: 15 minutes

**Test Scenarios:**
- Invalid API key handling
- Rate limiting and retry logic
- Network timeout scenarios
- Invalid field values (malformed data)
- Resource not found scenarios
- Concurrent modification conflicts

### 4.2 Terraform State Management Testing
**Timeline**: 10 minutes

**Test Scenarios:**
- State import functionality
- State migration scenarios  
- Resource drift detection and correction
- Partial configuration failures
- State corruption recovery

### 4.3 Validation and Input Testing  
**Timeline**: 10 minutes

**Test Scenarios:**
- Invalid contact method types
- Invalid user roles and permissions
- Invalid schedule time formats
- Invalid escalation timing configurations
- Required field validation
- Field length and format constraints

## Phase 5: Performance and Scale Testing

### 5.1 Bulk Operations Testing
**Timeline**: 10 minutes

**Test Scenarios:**
- Create 10+ users simultaneously
- Create multiple groups with complex memberships
- Batch update operations
- Large configuration file processing
- Provider memory usage monitoring

### 5.2 API Rate Limiting Testing
**Timeline**: 10 minutes

**Test Scenarios:**
- Rapid successive API calls
- Retry logic validation under load
- Exponential backoff testing
- Graceful degradation under rate limits

## Phase 6: Documentation and Usability Testing

### 6.1 Documentation Validation
**Timeline**: 10 minutes

**Test Coverage:**
- ✅ All resource schemas documented
- ✅ Example configurations provided
- ✅ Common use cases covered
- ✅ Error messages are clear and actionable
- ✅ Terraform registry compatibility

### 6.2 User Experience Testing
**Timeline**: 10 minutes

**Test Scenarios:**
- First-time user onboarding
- Common configuration patterns
- Error message clarity
- Debugging and troubleshooting flows

## Risk Mitigation and Rollback Strategy

### High-Risk Areas
1. **API Authentication Changes**: Could break all operations
2. **Complex Nested Structures**: Groups, escalation policies complexity
3. **Resource Dependencies**: Improper dependency chains
4. **State Management**: Data corruption or inconsistency

### Rollback Plans
1. **Code Rollback**: Git revert to last known working state
2. **State Rollback**: Terraform state backup and restore
3. **Resource Cleanup**: Manual AlertOps resource cleanup scripts
4. **Provider Rollback**: Revert to previous provider version

## Success Criteria Summary

### Individual Resource Criteria
Each resource must pass:
- ✅ Create operation (HTTP 200/201)
- ✅ Read operation (state refresh)
- ✅ Update operation (field modifications)
- ✅ Delete operation (HTTP 200/204)
- ✅ Import functionality
- ✅ Data source lookup

### Integration Criteria  
- ✅ Cross-resource references work
- ✅ Dependency ordering correct
- ✅ Complex nested structures handle properly
- ✅ End-to-end workflow validation

### Quality Criteria
- ✅ No memory leaks or performance issues
- ✅ Proper error handling and user messaging
- ✅ Code coverage >80% for critical paths
- ✅ Documentation complete and accurate

## Testing Timeline Summary

| Phase | Duration | Cumulative | Status |
|-------|----------|------------|---------|
| Groups + Users Integration | 25 min | 25 min | 🚧 Current |
| Schedules Implementation | 20 min | 45 min | ⏳ Next |
| Escalation Policies | 20 min | 65 min | ⏳ Pending |
| Workflows | 15 min | 80 min | ⏳ Pending |
| Inbound Integrations | 15 min | 95 min | ⏳ Pending |
| Services | 15 min | 110 min | ⏳ Pending |
| Cross-Resource Integration | 20 min | 130 min | ⏳ Pending |
| Edge Cases & Error Handling | 35 min | 165 min | ⏳ Pending |
| Performance & Scale | 20 min | 185 min | ⏳ Pending |
| Documentation & UX | 20 min | 205 min | ⏳ Pending |

**Total Testing Time**: 3.5 hours (including implementation)
**Implementation Time**: ~1.5 hours  
**Pure Testing Time**: ~2 hours

## Next Immediate Actions

1. **Execute Groups Testing** (Now)
   ```bash
   set TF_VAR_alertops_api_key=your-api-key
   cd test
   terraform plan
   terraform apply
   terraform output
   ```

2. **Analyze Results** (5 min)
   - Verify all success criteria met
   - Document any issues or adjustments needed
   - Prepare for Schedules implementation

3. **Proceed to Schedules** (If Groups pass)
   - Research AlertOps Schedules API
   - Implement Schedules resource
   - Test Schedules with Users/Groups integration

The comprehensive test plan ensures systematic validation of all components while maintaining rapid development pace toward the 2-hour completion goal. 