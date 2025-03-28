---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_system_ipv4_device_access_feature Data Source - terraform-provider-sdwan"
subcategory: "Features - System"
description: |-
  This data source can read the System IPv4 Device Access Feature.
---

# sdwan_system_ipv4_device_access_feature (Data Source)

This data source can read the System IPv4 Device Access Feature.

## Example Usage

```terraform
data "sdwan_system_ipv4_device_access_feature" "example" {
  id                 = "f6b2c44c-693c-4763-b010-895aa3d236bd"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `id` (String) The id of the Feature

### Read-Only

- `default_action` (String) Default Action
- `description` (String) The description of the Feature
- `name` (String) The name of the Feature
- `sequences` (Attributes List) Device Access Control List (see [below for nested schema](#nestedatt--sequences))
- `version` (Number) The version of the Feature

<a id="nestedatt--sequences"></a>
### Nested Schema for `sequences`

Read-Only:

- `base_action` (String) Base Action
- `destination_data_prefix_list_id` (String)
- `destination_ip_prefix_list` (Set of String) Destination Data IP Prefix List
- `destination_ip_prefix_list_variable` (String) Variable name
- `device_access_port` (Number) device access protocol
- `id` (Number) Sequence Id
- `name` (String) Sequence Name
- `source_data_prefix_list_id` (String)
- `source_ip_prefix_list` (Set of String) Source Data IP Prefix List
- `source_ip_prefix_list_variable` (String) Variable name
- `source_ports` (Set of Number) Source Port List
