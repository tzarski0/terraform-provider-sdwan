---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_service_lan_vpn_interface_svi_feature_associate_dhcp_server_feature Resource - terraform-provider-sdwan"
subcategory: "Features - Service"
description: |-
  This resource can manage a Service LAN VPN Interface SVI Feature Associate DHCP Server Feature .
  Minimum SD-WAN Manager version: 20.12.0
---

# sdwan_service_lan_vpn_interface_svi_feature_associate_dhcp_server_feature (Resource)

This resource can manage a Service LAN VPN Interface SVI Feature Associate DHCP Server Feature .
  - Minimum SD-WAN Manager version: `20.12.0`

## Example Usage

```terraform
resource "sdwan_service_lan_vpn_interface_svi_feature_associate_dhcp_server_feature" "example" {
  feature_profile_id                       = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  service_lan_vpn_feature_id               = "140331f6-5418-4755-a059-13c77eb96037"
  service_lan_vpn_interface_svi_feature_id = "140331f6-5418-4755-a059-13c77eb96037"
  service_dhcp_server_feature_id           = "140331f6-5418-4755-a059-13c77eb96037"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `service_dhcp_server_feature_id` (String) Service DHCP Server Feature ID
- `service_lan_vpn_feature_id` (String) Service LAN VPN Feature ID
- `service_lan_vpn_interface_svi_feature_id` (String) Service LAN VPN Interface SVI Feature ID

### Read-Only

- `id` (String) The id of the object
- `version` (Number) The version of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Expected import identifier with the format: "service_lan_vpn_interface_svi_feature_associate_dhcp_server_feature_id,feature_profile_id,service_lan_vpn_feature_id,service_lan_vpn_interface_svi_feature_id"
terraform import sdwan_service_lan_vpn_interface_svi_feature_associate_dhcp_server_feature.example "f6b2c44c-693c-4763-b010-895aa3d236bd,f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac,140331f6-5418-4755-a059-13c77eb96037,140331f6-5418-4755-a059-13c77eb96037"
```
