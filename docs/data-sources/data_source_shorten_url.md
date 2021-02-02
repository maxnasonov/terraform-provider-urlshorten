---
page_title: "shorten_url Data Source - terraform-provider-urlshorten"
subcategory: ""
description: |-
  The shorten_url data source uses a service and gives a short url for the given url.
---

# Data Source shorten_url

The shorten_url data source uses a service and gives a short url for the given url.

## Example Usage

```hcl
data "shorten_url" "response" {
  provider = urlshorten
  url     = "https://www.iana.org/dnssec/archive/launch-status-updates"
}
```

## Argument Reference

- `url` - (Required) URL to short.

## Attributes Reference

The following attributes are exported.

- `new_url` - The short version of your URL.
