---
page_title: "morpheus_power_schedule_policy Resource - terraform-provider-morpheus"
subcategory: ""
description: |-
  Provides a Morpheus power schedule policy resource
---

# morpheus_power_schedule_policy

Provides a Morpheus power schedule policy resource

## Example Usage

Creating the policy with a global scope:

```terraform
resource "morpheus_power_schedule_policy" "tf_example_power_schedule_policy_global" {
  name                         = "tf_example_power_schedule_policy_global"
  description                  = "terraform example global power schedule policy"
  enabled                      = true
  enforcement_type             = "fixed"
  power_schedule_id            = 2
  hide_power_schedule_if_fixed = true
  scope                        = "global"
}
```

Creating the policy with a cloud scope:

```terraform
resource "morpheus_power_schedule_policy" "tf_example_power_schedule_policy_cloud" {
  name                         = "tf_example_power_schedule_policy_cloud"
  description                  = "terraform example cloud power schedule policy"
  enabled                      = true
  enforcement_type             = "fixed"
  power_schedule_id            = 2
  hide_power_schedule_if_fixed = true
  scope                        = "cloud"
  cloud_id                     = 1
}
```

Creating the policy with a group scope:

```terraform
resource "morpheus_power_schedule_policy" "tf_example_power_schedule_policy_group" {
  name                         = "tf_example_power_schedule_policy_group"
  description                  = "terraform example group power schedule policy"
  enabled                      = true
  enforcement_type             = "fixed"
  power_schedule_id            = 2
  hide_power_schedule_if_fixed = true
  scope                        = "group"
  group_id                     = 1
}
```

Creating the policy with a role scope:

```terraform
resource "morpheus_power_schedule_policy" "tf_example_power_schedule_policy_role" {
  name                         = "tf_example_power_schedule_policy_role"
  description                  = "terraform example role power schedule policy"
  enabled                      = true
  enforcement_type             = "fixed"
  power_schedule_id            = 2
  hide_power_schedule_if_fixed = true
  scope                        = "role"
  role_id                      = 1
  apply_to_each_user           = true
}
```

Creating the policy with a user scope:

```terraform
resource "morpheus_power_schedule_policy" "tf_example_power_schedule_policy_user" {
  name                         = "tf_example_power_schedule_policy_user"
  description                  = "terraform example user power schedule policy"
  enabled                      = true
  enforcement_type             = "fixed"
  power_schedule_id            = 2
  hide_power_schedule_if_fixed = true
  scope                        = "user"
  user_id                      = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enforcement_type` (String) The enforcement type of the policy (fixed, user)
- `name` (String) The name of the power schedule policy
- `power_schedule_id` (Number) The ID of the power schedule to associate with the policy
- `scope` (String) The filter or scope that the policy is applied to (global, group, cloud, user, role)

### Optional

- `apply_to_each_user` (Boolean) Whether to assign the policy at the individual user level to all users assigned the associated role
- `cloud_id` (Number) The id of the cloud associated with the cloud scoped filter
- `description` (String) The description of the power schedule policy
- `enabled` (Boolean) Whether the policy is enabled
- `group_id` (Number) The id of the group associated with the group scoped filter
- `hide_power_schedule_if_fixed` (Boolean) Whether to hide the power schedule option on the instance provisioning wizard if the enforcement type is fixed
- `role_id` (Number) The id of the role associated with the role scoped filter
- `tenant_ids` (List of Number) A list of tenant IDs to assign the policy to
- `user_id` (Number) The id of the user associated with the user scoped filter

### Read-Only

- `id` (String) The ID of the power schedule policy

## Import

Import is supported using the following syntax:

```shell
terraform import morpheus_power_schedule_policy.tf_example_power_schedule_policy 1
```
