---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_application_priority_qos_policy Resource - terraform-provider-sdwan"
subcategory: "Policies"
description: |-
  This resource can manage a Application Priority QoS Policy.
  Minimum SD-WAN Manager version: 20.12.0
---

# sdwan_application_priority_qos_policy (Resource)

This resource can manage a Application Priority QoS Policy.
  - Minimum SD-WAN Manager version: `20.12.0`

## Example Usage

```terraform
resource "sdwan_application_priority_qos_policy" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  target_interface   = ["{{interface_var_1}}"]
  qos_schedulers = [
    {
      drops           = "tail-drop"
      queue           = "0"
      bandwidth       = "10"
      scheduling_type = "llq"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `name` (String) The name of the Policy

### Optional

- `description` (String) The description of the Policy
- `qos_schedulers` (Attributes List) qosSchedulers (see [below for nested schema](#nestedatt--qos_schedulers))
- `target_interface` (Set of String) interfaces
- `target_interface_variable` (String) Variable name

### Read-Only

- `id` (String) The id of the Policy
- `version` (Number) The version of the Policy

<a id="nestedatt--qos_schedulers"></a>
### Nested Schema for `qos_schedulers`

Optional:

- `bandwidth` (String) bandwidthPercent
- `drops` (String) drops
- `forwarding_class_id` (String)
- `queue` (String) queue
- `scheduling_type` (String) scheduling

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Expected import identifier with the format: "application_priority_qos_policy_id,feature_profile_id"
terraform import sdwan_application_priority_qos_policy.example "f6b2c44c-693c-4763-b010-895aa3d236bd,f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
```
