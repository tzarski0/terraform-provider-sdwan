---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_vpn_membership_policy_definition Resource - terraform-provider-sdwan"
subcategory: "(Classic) Centralized Policies"
description: |-
  This resource can manage a VPN Membership Policy Definition .
---

# sdwan_vpn_membership_policy_definition (Resource)

This resource can manage a VPN Membership Policy Definition .

## Example Usage

```terraform
resource "sdwan_vpn_membership_policy_definition" "example" {
  name        = "Example"
  description = "My description"
  sites = [
    {
      site_list_id = "e858e1c4-6aa8-4de7-99df-c3adbf80290d"
      vpn_list_ids = ["04fcbb0b-efbf-43d2-a04b-847d3a7b104e"]
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) The description of the policy definition
- `name` (String) The name of the policy definition
- `sites` (Attributes List) List of sites (see [below for nested schema](#nestedatt--sites))

### Read-Only

- `id` (String) The id of the object
- `type` (String) Type
- `version` (Number) The version of the object

<a id="nestedatt--sites"></a>
### Nested Schema for `sites`

Optional:

- `site_list_id` (String) Site list ID
- `site_list_version` (Number) Site list version
- `vpn_list_ids` (Set of String) VPN list IDs
- `vpn_list_versions` (List of String) VPN list versions

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import sdwan_vpn_membership_policy_definition.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```
