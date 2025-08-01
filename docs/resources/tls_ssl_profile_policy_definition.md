---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_tls_ssl_profile_policy_definition Resource - terraform-provider-sdwan"
subcategory: "(Classic) Security Policies"
description: |-
  This resource can manage a TLS SSL Profile Policy Definition .
---

# sdwan_tls_ssl_profile_policy_definition (Resource)

This resource can manage a TLS SSL Profile Policy Definition .

## Example Usage

```terraform
resource "sdwan_tls_ssl_profile_policy_definition" "example" {
  name                     = "Example"
  description              = "My description"
  mode                     = "security"
  decrypt_categories       = ["alcohol-and-tobacco"]
  never_decrypt_categories = ["auctions"]
  skip_decrypt_categories  = ["cdns"]
  decrypt_threshold        = "high-risk"
  reputation               = false
  fail_decrypt             = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) The description of the policy definition.
- `name` (String) The name of the policy definition.

### Optional

- `allow_url_list_id` (String) Allow URL list ID
- `allow_url_list_version` (Number) Allow URL list version
- `block_url_list_id` (String) Block URL list ID
- `block_url_list_version` (Number) Block URL list version
- `decrypt_categories` (Set of String) Categories that should be decrypted
- `decrypt_threshold` (String) Decrypt threshold
  - Choices: `high-risk`, `suspicious`, `moderate-risk`, `low-risk`, `trustworthy`
- `fail_decrypt` (Boolean) Fail decrypt enabled
- `mode` (String) The policy mode
  - Choices: `security`, `unified`
- `never_decrypt_categories` (Set of String) Categories that should never be decrypted
- `reputation` (Boolean) Reputation enabled
- `skip_decrypt_categories` (Set of String) Categories that should skipped

### Read-Only

- `id` (String) The id of the object
- `version` (Number) The version of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import sdwan_tls_ssl_profile_policy_definition.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```
