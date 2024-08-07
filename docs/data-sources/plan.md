---
page_title: "morpheus_plan Data Source - terraform-provider-morpheus"
subcategory: ""
description: |-
  Provides a Morpheus plan data source.
---

# morpheus_plan (Data Source)

Provides a Morpheus plan data source.

## Example Usage

```terraform
data "morpheus_plan" "vmware" {
  name = "1 CPU, 4GB Memory"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `provision_type` (String) The name of the service plan provisiong type (i.e. - Amazon EC2, Azure, Google, Nutanix, VMware, etc.)

### Optional

- `name` (String) The name of the Morpheus plan.

### Read-Only

- `code` (String) Optional code for use with policies
- `description` (String) The description of the plan
- `id` (Number) The ID of this resource.