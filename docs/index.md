---
page_title: "Provider: Morpheus"
subcategory: ""
description: |-
  Terraform provider for configuring Morpheus.
---

# Morpheus Provider

The Morpheus provider is used to interact with the resources supported by [Morpheus Cloud Management Platform (CMP)](https://morpheusdata.com/). The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Authentication

The Morpheus provider supports authentication via username/password or an access token. The [authentication guide](guides/auth.md) describes how to obtain client credentials.

## Example Usage

```terraform
terraform {
  required_providers {
    morpheus = {
      source  = "morpheus/morpheus"
      version = "~> 0.1"
    }
  }
}

# Configure the provider
provider "morpheus" {
  url      = "${var.morpheus_url}"
  username = "${var.morpheus_username}"
  password = "${var.morpheus_password}"       
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **url** (String) The URL of the Morpheus Data Appliance where requests will be directed.

### Optional

- **access_token** (String, Sensitive) Access Token of Morpheus user. This can be used instead of authenticating with Username and Password.
- **password** (String, Sensitive) Password of Morpheus user for authentication
- **username** (String) Username of Morpheus user for authentication