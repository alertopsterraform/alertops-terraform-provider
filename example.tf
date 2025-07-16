terraform {
  required_providers {
    alertops = {
      source  = "alertopsterraform/alertops"
      version = "1.0.0"
    }
  }
}

provider "alertops" {
  api_key  = var.alertops_api_key
  base_url = "https://api.alertops.com"
}

variable "alertops_api_key" {
  description = "AlertOps API key"
  type        = string
  sensitive   = true
}

# Create a user with contact methods
resource "alertops_user" "example_user" {
  user_name  = "john.doe"
  first_name = "John"
  last_name  = "Doe"
  locale     = "en-US"
  time_zone  = "(UTC-06:00) Central Time (US & Canada)"
  type       = "Standard"
  
  contact_methods = [
    {
      contact_method_name = "Email-Official"
      email = [
        {
          email_address = "john.doe@example.com"
        }
      ]
      enabled  = true
      sequence = 1
    },
    {
      contact_method_name = "Phone-Official-Mobile"
      phone = [
        {
          country_code = "+1"
          phone_number = "5551234567"
        }
      ]
      enabled  = true
      sequence = 2
    },
    {
      contact_method_name = "SMS-Personal"
      sms = [
        {
          country_code = "+1"
          phone_number = "5551234567"
        }
      ]
      enabled  = true
      sequence = 3
    }
  ]
  
  roles = ["Standard User"]
}

# Data source to lookup a user
data "alertops_user" "lookup_user" {
  user_name = "john.doe"
  depends_on = [alertops_user.example_user]
}

# Output user information
output "user_id" {
  value = alertops_user.example_user.user_id
}

output "user_last_login" {
  value = alertops_user.example_user.last_login_date
}

output "lookup_user_id" {
  value = data.alertops_user.lookup_user.user_id
}

output "user_contact_methods" {
  value = alertops_user.example_user.contact_methods
} 