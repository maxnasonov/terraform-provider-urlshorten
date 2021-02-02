---
page_title: "Provider: Shorten URL"
subcategory: ""
description: |-
  Terraform provider for interacting with Shorten URL services.
---

# Shorten URL Provider

The Shorten URL provider uses a service and shortens the given url.

## Example Usage

The Shorten URL client could be generated with the following parameters:

```hcl
  # Configure provider with your service details
  provider "urlshorten" {
    # Shorten URL service's password
    password = "admin123"
    # it can be set using the environment variable SHORTEN_PASSWORD

    # Service's name
    service = "tinyurl"
    # it can be set using the environment variable SHORTEN_SERVICE
  }

  # Retrieve project's templates
  data "shorten_url" "response" {
    provider = urlshorten
    url     = "https://www.iana.org/dnssec/archive/launch-status-updates"
  }
  output "shorten_url_response" {
    value = data.shorten_url.response
  }
```

Do not keep your authentication password in HashiCorp for production environments, use environment variables.

## Argument Reference

- **service** - (Required) - Service name.
- **password** - (Required) - Password for the service.
