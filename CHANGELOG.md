# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial release preparation
- Documentation improvements
- Example configurations

## [1.0.0] - 2024-01-15

### Added
- **Users Resource**: Complete CRUD operations for AlertOps users
  - Contact methods (Email, SMS, Phone, Gateway)
  - User roles and permissions
  - Time zone and locale settings
  - Data source for user lookups

- **Groups Resource**: Group management with dynamic membership
  - Static and dynamic group types
  - Member management with roles
  - Contact method configurations
  - Group hierarchy support

- **Schedules Resource**: On-call schedule management
  - Fixed and rotating schedule types
  - Weekday/weekend configurations
  - Time zone support
  - Rotation settings and restrictions

- **Workflows Resource**: Alert processing automation
  - Condition-based alert routing
  - Custom actions and notifications
  - Webhook integrations
  - Template variables support

- **Escalation Policies Resource**: Multi-tier escalation management
  - Primary, secondary, tertiary escalation roles
  - Flexible timing configurations
  - Notification method preferences
  - Auto-close and snooze options

- **Inbound Integrations Resource**: Alert ingestion management
  - API integrations with bidirectional support
  - Email integrations with mailbox configuration
  - Bridge integrations for phone access
  - Heartbeat monitoring settings

### Documentation
- Comprehensive README with quick start guide
- Provider documentation and resource references
- Multiple example configurations
- Production deployment patterns
- Troubleshooting guides

### Examples
- **Basic Setup**: Simple getting-started configuration
- **Complete Environment**: Production-ready multi-team setup
- **Integration Hub**: Multiple integration types demonstration

### Infrastructure
- Automated build and release pipeline
- Cross-platform binary distribution
- Terraform Registry preparation
- GitHub Actions CI/CD
- GoReleaser configuration

## [0.1.0] - 2024-01-01

### Added
- Initial provider structure
- Basic user resource implementation
- Authentication and API client setup
- Development environment configuration

---

## Release Notes

### v1.0.0 - Initial Public Release

This is the first production-ready release of the AlertOps Terraform Provider. It includes complete support for all major AlertOps resources and is ready for use in production environments.

**Key Features:**
- ✅ 6 Resource Types Supported
- ✅ 1 Data Source Available  
- ✅ Production-Ready Examples
- ✅ Comprehensive Documentation
- ✅ Cross-Platform Binaries
- ✅ Terraform Registry Compatible

**Supported AlertOps Features:**
- User management with contact methods
- Group organization and membership
- On-call schedule configuration
- Alert workflow automation
- Multi-tier escalation policies
- Inbound integration management

**Getting Started:**
```hcl
terraform {
  required_providers {
    alertops = {
      source  = "alertops/alertops"
      version = "~> 1.0"
    }
  }
}

provider "alertops" {
  api_key = var.alertops_api_key
}
```

**Migration Guide:**
This is the initial public release, so no migration is required.

**Known Limitations:**
- Email integrations require valid mailbox configuration
- Bridge integrations limited to basic phone access
- Some advanced workflow features may require additional configuration

**Upcoming Features (v1.1.0):**
- Enhanced integration filtering
- Advanced schedule rotation patterns
- Custom notification templates
- Bulk resource operations
- Additional data sources

For detailed usage instructions, see the [documentation](https://github.com/alertops/terraform-provider-alertops/tree/main/docs) and [examples](https://github.com/alertops/terraform-provider-alertops/tree/main/examples). 