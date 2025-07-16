# AlertOps Basic Setup Example

This example demonstrates how to set up a basic AlertOps environment using Terraform. It creates a complete working configuration with users, groups, schedules, escalation policies, workflows, and integrations.

## What This Example Creates

- **3 Users**: Primary on-call, secondary on-call, and team lead
- **2 Groups**: On-call engineers and management
- **1 Schedule**: Weekday schedule for on-call rotation
- **1 Escalation Policy**: Two-tier escalation (engineers → management)
- **1 Workflow**: Basic alert processing
- **1 Integration**: API integration for receiving alerts

## Prerequisites

1. **AlertOps Account**: You need an active AlertOps account
2. **API Key**: Generate an API key from AlertOps UI
3. **Terraform**: Install Terraform 0.13 or later

## Getting Your API Key

1. Log into your AlertOps account
2. Navigate to **Settings** → **API Management**
3. Click **Generate API Key**
4. Copy the generated key (you'll need it for configuration)

## Quick Start

### 1. Clone and Navigate

```bash
git clone https://github.com/alertops/terraform-provider-alertops.git
cd terraform-provider-alertops/examples/basic-setup
```

### 2. Configure Variables

Copy the example variables file and update it with your information:

```bash
cp terraform.tfvars.example terraform.tfvars
```

Edit `terraform.tfvars`:

```hcl
# Required: Your AlertOps API Key
alertops_api_key = "your-actual-api-key-here"

# Your company's email domain
company_domain = "yourcompany.com"

# Environment identifier
environment = "dev"
```

### 3. Initialize and Apply

```bash
# Initialize Terraform
terraform init

# Review the planned changes
terraform plan

# Apply the configuration
terraform apply
```

### 4. Verify Setup

After successful deployment, you'll see output showing all created resources:

```
setup_summary = {
  "environment" = "dev"
  "escalation_policy_created" = "standard-escalation-dev"
  "groups_created" = {
    "management" = "management-dev"
    "oncall_engineers" = "oncall-engineers-dev"
  }
  "integration_created" = "api-alerts-dev"
  "schedule_created" = "weekday-oncall-dev"
  "users_created" = {
    "primary_oncall" = "primary-oncall-dev"
    "secondary_oncall" = "secondary-oncall-dev"
    "team_lead" = "team-lead-dev"
  }
  "workflow_created" = "basic-alert-processor-dev"
}
```

## Configuration Details

### Users Created

| User | Email | Contact Methods | Purpose |
|------|-------|----------------|---------|
| `primary-oncall-{env}` | `oncall-primary@{domain}` | Email + SMS | First responder |
| `secondary-oncall-{env}` | `oncall-secondary@{domain}` | Email | Backup responder |
| `team-lead-{env}` | `team-lead@{domain}` | Email | Escalation contact |

### Escalation Flow

1. **Primary Role** (2 minutes between members, 10 minutes total)
   - On-call engineers group
2. **Secondary Role** (5 minutes between members, 15 minutes total)
   - Management group

### Schedule Configuration

- **Days**: Monday through Friday
- **Time Zone**: Eastern Time
- **Members**: All users in the on-call engineers group

## Customization

### Changing Contact Information

Update the contact methods in the user resources:

```hcl
contact_methods {
  contact_method_name = "Email-Official"
  email {
    email_address = "your-actual-email@company.com"
  }
  enabled = true
  sequence = 1
}
```

### Adding Phone Numbers

Add phone contact methods:

```hcl
contact_methods {
  contact_method_name = "Phone-Official"
  phone {
    phone_number = "555-0123"
    country_code = "1"
  }
  enabled = true
  sequence = 3
}
```

### Modifying Schedule

Change the schedule days:

```hcl
schedule_weekdays {
  mon = true
  tue = true
  wed = true
  thu = true
  fri = true
  sat = true  # Add Saturday
  sun = false
}
```

### Multiple Environments

Deploy different environments by changing the `environment` variable:

```bash
# Development environment
terraform apply -var="environment=dev"

# Staging environment
terraform apply -var="environment=staging"

# Production environment
terraform apply -var="environment=prod"
```

## Testing Your Setup

### 1. Test API Integration

After deployment, find your integration endpoint in the AlertOps UI and send a test alert:

```bash
curl -X POST "https://api.alertops.com/webhook/your-integration-id" \
  -H "Content-Type: application/json" \
  -d '{
    "subject": "Test Alert",
    "message": "This is a test alert from Terraform setup",
    "priority": "High"
  }'
```

### 2. Verify Escalation

1. Create a test alert
2. Verify primary on-call receives notification
3. Wait for escalation timeout
4. Verify secondary escalation occurs

## Cleanup

To remove all created resources:

```bash
terraform destroy
```

**⚠️ Warning**: This will permanently delete all AlertOps resources created by this configuration.

## Next Steps

Once you have the basic setup working:

1. **Add More Users**: Create additional team members
2. **Complex Schedules**: Set up rotating schedules with time restrictions
3. **Multiple Integrations**: Add integrations for different monitoring tools
4. **Custom Workflows**: Create workflows for specific alert types
5. **Advanced Escalation**: Set up complex escalation policies with multiple tiers

## Troubleshooting

### Common Issues

**Invalid API Key**
```
Error: authentication failed
```
- Verify your API key is correct
- Ensure the API key has necessary permissions

**Resource Already Exists**
```
Error: resource already exists
```
- Check if resources with the same names exist in AlertOps
- Change the `environment` variable to use a unique suffix

**Invalid Contact Method**
```
Error: contact_method_name must be one of: [...]
```
- Use only supported contact method names
- See the main documentation for valid contact method types

### Getting Help

- 📖 [AlertOps Documentation](https://docs.alertops.com)
- 💬 [Community Forum](https://community.alertops.com)
- 📧 [Support Email](mailto:support@alertops.com)
- 🐛 [GitHub Issues](https://github.com/alertops/terraform-provider-alertops/issues)

## Related Examples

- [Complete Environment](../complete-environment/) - Full production setup
- [Integration Hub](../integrations/) - Multiple integration types
- [Multi-Team Setup](../multi-team/) - Complex team structures 