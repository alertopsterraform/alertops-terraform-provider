# Terraform Provider for AlertOps

[![Go Report Card](https://goreportcard.com/badge/github.com/alertops/terraform-provider-alertops)](https://goreportcard.com/report/github.com/alertops/terraform-provider-alertops)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/release/alertops/terraform-provider-alertops.svg)](https://github.com/alertops/terraform-provider-alertops/releases)

The official Terraform provider for [AlertOps](https://alertops.com), allowing you to manage your AlertOps infrastructure as code.

## Features

- **Complete AlertOps Management**: Manage users, groups, schedules, workflows, escalation policies, and integrations
- **Infrastructure as Code**: Version control and automate your AlertOps configuration
- **Production Ready**: Used by AlertOps customers in production environments
- **Comprehensive Testing**: Extensive test suite with real-world scenarios

## Supported Resources

| Resource | Description |
|----------|-------------|
| `alertops_user` | Manage AlertOps users with contact methods and roles |
| `alertops_group` | Manage groups with dynamic membership and contact settings |
| `alertops_schedule` | Create on-call schedules with rotation and time restrictions |
| `alertops_workflow` | Automate alert processing with conditions and actions |
| `alertops_escalation_policy` | Define multi-tier escalation with flexible notification options |
| `alertops_inbound_integration` | Configure API, Email, and Bridge integrations |

## Supported Data Sources

| Data Source | Description |
|-------------|-------------|
| `alertops_user` | Retrieve user information by ID or username |

## Quick Start

### 1. Installation

#### Terraform 0.13+

Add the provider to your Terraform configuration:

```hcl
terraform {
  required_providers {
    alertops = {
      source  = "alertops/alertops"
      version = "~> 1.0"
    }
  }
}
```

#### Manual Installation

Download the latest binary from [releases](https://github.com/alertops/terraform-provider-alertops/releases) and place it in your Terraform plugins directory.

### 2. Configuration

Configure the provider with your AlertOps API key:

```hcl
provider "alertops" {
  api_key  = var.alertops_api_key  # or use ALERTOPS_API_KEY env var
  base_url = "https://api.alertops.com"  # optional, defaults to official API
}

variable "alertops_api_key" {
  description = "AlertOps API Key"
  type        = string
  sensitive   = true
}
```

### 3. Your First Resource

Create a user with contact methods:

```hcl
resource "alertops_user" "oncall_engineer" {
  user_name  = "john.doe"
  first_name = "John"
  last_name  = "Doe"
  locale     = "en-US"
  time_zone  = "(UTC-05:00) Eastern Time (US & Canada)"
  type       = "Standard"
  
  contact_methods {
    contact_method_name = "Email-Official"
    email {
      email_address = "john.doe@company.com"
    }
    enabled = true
    sequence = 1
  }
  
  contact_methods {
    contact_method_name = "SMS-Official"
    sms {
      phone_number = "555-0123"
      country_code = "1"
    }
    enabled = true
    sequence = 2
  }
  
  roles = ["Basic"]
}
```

## Examples

### Basic Setup
- [Simple User and Group](examples/basic-setup/)
- [On-Call Schedule](examples/schedule/)
- [Escalation Policy](examples/escalation-policy/)

### Advanced Scenarios
- [Complete Environment](examples/complete-environment/)
- [Multi-Team Setup](examples/multi-team/)
- [Integration Hub](examples/integrations/)

### Production Patterns
- [Blue-Green Deployments](examples/blue-green/)
- [Disaster Recovery](examples/disaster-recovery/)
- [Compliance Setup](examples/compliance/)

## Authentication

### API Key

Get your API key from AlertOps:
1. Log into your AlertOps account
2. Go to **Settings** → **API Management**
3. Generate a new API key
4. Set the environment variable: `export ALERTOPS_API_KEY="your-api-key"`

### Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `ALERTOPS_API_KEY` | Your AlertOps API key | Yes |
| `ALERTOPS_BASE_URL` | API base URL (defaults to https://api.alertops.com) | No |

## Documentation

- 📖 [Provider Documentation](docs/index.md)
- 🔧 [Resource Reference](docs/resources/)
- 📊 [Data Sources](docs/data-sources/)
- 💡 [Examples](examples/)
- ❓ [FAQ](docs/faq.md)

## Development

### Building from Source

```bash
git clone https://github.com/alertops/terraform-provider-alertops.git
cd terraform-provider-alertops
go build -o terraform-provider-alertops
```

### Running Tests

```bash
# Unit tests
go test ./...

# Acceptance tests (requires API key)
TF_ACC=1 ALERTOPS_API_KEY=your-key go test ./...
```

### Local Development

1. Build the provider: `make build`
2. Create `.terraformrc` in your home directory:
```hcl
provider_installation {
  dev_overrides {
    "alertops/alertops" = "/path/to/terraform-provider-alertops"
  }
  direct {}
}
```

## Contributing

We welcome contributions! Please see:
- [Contributing Guidelines](CONTRIBUTING.md)
- [Code of Conduct](CODE_OF_CONDUCT.md)
- [Development Guide](docs/development.md)

### Reporting Issues

- 🐛 [Bug Reports](https://github.com/alertops/terraform-provider-alertops/issues/new?template=bug_report.md)
- ✨ [Feature Requests](https://github.com/alertops/terraform-provider-alertops/issues/new?template=feature_request.md)
- 📖 [Documentation Issues](https://github.com/alertops/terraform-provider-alertops/issues/new?template=documentation.md)

## Support

- 📧 **Email**: support@alertops.com
- 💬 **API**: [AlertOps API Docs](https://api.alertops.com)
- 📚 **Documentation**: [docs.alertops.com](https://docs.alertops.com)
- 🎯 **Issues**: [GitHub Issues](https://github.com/alertops/terraform-provider-alertops/issues)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for release notes and version history.

---

