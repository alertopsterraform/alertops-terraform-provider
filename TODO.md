# AlertOps Terraform Provider - TODO List

## 🎯 Project Goal
Build a complete Terraform provider for AlertOps to manage Users, Groups, Schedules, Escalation Policies, Workflows, Inbound Integrations, and Services via CRUD operations.

## ✅ COMPLETED

### ✅ User Resource - FULLY FUNCTIONAL
- [x] **Core Infrastructure**
  - [x] Go module setup with terraform-plugin-sdk/v2
  - [x] HTTP client with proper authentication (`api-key` header)
  - [x] Retry logic and error handling
  - [x] Debug logging for API calls

- [x] **User Resource Implementation**
  - [x] Complete User struct matching AlertOps API v2 schema
  - [x] Full CRUD operations (Create, Read, Update, Delete)
  - [x] Complex contact methods with validation
  - [x] Proper Terraform schema with nested blocks
  - [x] HTTP status code handling (200/201/204)

- [x] **Authentication & API Integration**
  - [x] Fixed authentication to use `api-key` header (not Authorization Bearer)
  - [x] Proper JSON request/response handling
  - [x] Error messages with full request/response details
  - [x] Support for AlertOps API v2 endpoints

- [x] **Development Environment**
  - [x] Local development with dev_overrides
  - [x] Testing scripts and documentation
  - [x] Validation for predefined contact method types

### ✅ Groups Resource - FULLY FUNCTIONAL
- [x] **Groups Resource Implementation**
  - [x] Complete Group struct matching AlertOps API v2 schema
  - [x] Full CRUD operations (Create, Read, Update, Delete)
  - [x] Group members support (user ID references)
  - [x] Escalation policy integration
  - [x] Proper Terraform schema with computed fields
  - [x] Debug JSON output for API troubleshooting

- [x] **Integration Testing**
  - [x] Combined User + Group testing
  - [x] Resource dependencies (user created first, then added to group)
  - [x] Schema validation working correctly
  - [x] All fields properly mapped

## 🔄 IN PROGRESS

### 🎯 Next: Schedules Resource
**Priority**: HIGH - Critical for on-call management

## 📋 PENDING IMPLEMENTATION

### 1. Schedules Resource  
- [ ] Research AlertOps Schedules API structure (`/api/v2/schedules`)
- [ ] Create `resource_schedule.go` with CRUD operations
- [ ] Handle schedule rules and time zones
- [ ] Test CRUD operations

### 2. Escalation Policies Resource
- [ ] Research AlertOps Escalation Policies API structure
- [ ] Create `resource_escalation_policy.go`
- [ ] Handle escalation rules and user/group references
- [ ] Test CRUD operations

### 3. Workflows Resource
- [ ] Research AlertOps Workflows API structure
- [ ] Create `resource_workflow.go` 
- [ ] Handle workflow steps and conditions
- [ ] Test CRUD operations

### 4. Inbound Integrations Resource
- [ ] Research AlertOps Inbound Integrations API structure
- [ ] Create `resource_inbound_integration.go`
- [ ] Handle integration types and configurations
- [ ] Test CRUD operations

### 5. Services Resource
- [ ] Research AlertOps Services API structure
- [ ] Create `resource_service.go`
- [ ] Handle service dependencies and configurations
- [ ] Test CRUD operations

## 🔮 FUTURE ENHANCEMENTS

### Data Sources
- [ ] `data.alertops_user` - User lookup ✅ (DONE)
- [ ] `data.alertops_group` - Group lookup  
- [ ] `data.alertops_schedule` - Schedule lookup
- [ ] Additional data sources for all resources

### Advanced Features
- [ ] Import functionality for existing AlertOps resources
- [ ] Comprehensive validation and error handling
- [ ] Resource relationships and dependencies
- [ ] Bulk operations support

### Documentation & Testing
- [ ] Complete API documentation
- [ ] Integration test suites
- [ ] Example configurations
- [ ] Best practices guide

## 📝 Technical Notes

### Authentication
- ✅ Uses `api-key` header (not Authorization Bearer)
- ✅ API key from terraform.tfvars or environment variable

### HTTP Status Codes
- ✅ POST (CREATE): Accepts 200, 201, 204
- ✅ PUT (UPDATE): Accepts 200, 204  
- ✅ DELETE: Accepts 200, 204
- ✅ GET (READ): Accepts 200

### Dev Environment
- ✅ Uses dev_overrides to bypass Terraform registry
- ✅ Skip `terraform init` - go directly to validate/plan/apply
- ✅ Environment: `TF_CLI_CONFIG_FILE=C:\alertopsterraform\test\.terraformrc`

## 🚀 Ready to Continue

**Current Status**: Users and Groups resources are complete and fully tested. Ready to implement Schedules resource next.

**Next Action**: Research AlertOps Schedules API structure and begin implementation.

## 🏆 Achievement Summary

✅ **2 out of 6 core resources completed** (33% done)
✅ **Full CRUD operations** working for Users and Groups
✅ **Authentication & API integration** solid foundation
✅ **Resource dependencies** working (Groups can reference Users)
✅ **Dev environment** streamlined for rapid development

**Estimated completion**: 33% done, ~2 hours remaining for remaining 4 resources 