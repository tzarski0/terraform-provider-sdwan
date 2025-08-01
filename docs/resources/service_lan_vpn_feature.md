---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_service_lan_vpn_feature Resource - terraform-provider-sdwan"
subcategory: "Features - Service"
description: |-
  This resource can manage a Service LAN VPN Feature.
  Minimum SD-WAN Manager version: 20.12.0
---

# sdwan_service_lan_vpn_feature (Resource)

This resource can manage a Service LAN VPN Feature.
  - Minimum SD-WAN Manager version: `20.12.0`

## Example Usage

```terraform
resource "sdwan_service_lan_vpn_feature" "example" {
  name                       = "Example"
  description                = "My Example"
  feature_profile_id         = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  vpn                        = 1
  config_description         = "VPN1"
  omp_admin_distance_ipv4    = 1
  omp_admin_distance_ipv6    = 1
  enable_sdwan_remote_access = false
  primary_dns_address_ipv4   = "1.2.3.4"
  secondary_dns_address_ipv4 = "2.3.4.5"
  primary_dns_address_ipv6   = "2001:0:0:1::0"
  secondary_dns_address_ipv6 = "2001:0:0:2::0"
  host_mappings = [
    {
      host_name   = "HOSTMAPPING1"
      list_of_ips = ["1.2.3.4"]
    }
  ]
  ipv4_static_routes = [
    {
      network_address = "1.2.3.4"
      subnet_mask     = "0.0.0.0"
      next_hops = [
        {
          address                 = "1.2.3.4"
          administrative_distance = 1
        }
      ]
    }
  ]
  ipv6_static_routes = [
    {
      prefix = "2001:0:0:1::0/12"
      next_hops = [
        {
          address                 = "2001:0:0:1::0"
          administrative_distance = 1
        }
      ]
    }
  ]
  services = [
    {
      service_type   = "FW"
      ipv4_addresses = ["1.2.3.4"]
      tracking       = true
    }
  ]
  service_routes = [
    {
      network_address = "1.2.3.4"
      subnet_mask     = "0.0.0.0"
      service         = "SIG"
      vpn             = 0
    }
  ]
  gre_routes = [
    {
      network_address = "1.2.3.4"
      subnet_mask     = "0.0.0.0"
      interface       = ["gre01"]
      vpn             = 0
    }
  ]
  ipsec_routes = [
    {
      network_address = "1.2.3.4"
      subnet_mask     = "0.0.0.0"
      interface       = ["ipsec01"]
    }
  ]
  nat_pools = [
    {
      nat_pool_name = 1
      prefix_length = 3
      range_start   = "1.2.3.4"
      range_end     = "2.3.4.5"
      overload      = true
      direction     = "inside"
    }
  ]
  nat_port_forwards = [
    {
      nat_pool_name        = 2
      source_port          = 122
      translate_port       = 330
      source_ip            = "1.2.3.4"
      translated_source_ip = "2.3.4.5"
      protocol             = "TCP"
    }
  ]
  static_nats = [
    {
      nat_pool_name        = 3
      source_ip            = "1.2.3.4"
      translated_source_ip = "2.3.4.5"
      static_nat_direction = "inside"
    }
  ]
  nat_64_v4_pools = [
    {
      name        = "NATPOOL1"
      range_start = "1.2.3.4"
      range_end   = "2.3.4.5"
      overload    = false
    }
  ]
  ipv4_import_route_targets = [
    {
      route_target = "1.1.1.3:200"
    }
  ]
  ipv4_export_route_targets = [
    {
      route_target = "1.1.1.3:200"
    }
  ]
  ipv6_import_route_targets = [
    {
      route_target = "1.1.1.3:200"
    }
  ]
  ipv6_export_route_targets = [
    {
      route_target = "1.1.1.3:200"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `name` (String) The name of the Feature

### Optional

- `advertise_omp_ipv4s` (Attributes List) OMP Advertise IPv4 (see [below for nested schema](#nestedatt--advertise_omp_ipv4s))
- `advertise_omp_ipv6s` (Attributes List) OMP Advertise IPv6 (see [below for nested schema](#nestedatt--advertise_omp_ipv6s))
- `config_description` (String) Name
- `config_description_variable` (String) Variable name
- `description` (String) The description of the Feature
- `enable_sdwan_remote_access` (Boolean) Enable SDWAN Remote Access
  - Default value: `false`
- `gre_routes` (Attributes List) IPv4 Static GRE Route (see [below for nested schema](#nestedatt--gre_routes))
- `host_mappings` (Attributes List) (see [below for nested schema](#nestedatt--host_mappings))
- `ipsec_routes` (Attributes List) IPv4 Static IPSEC Route (see [below for nested schema](#nestedatt--ipsec_routes))
- `ipv4_export_route_targets` (Attributes List) (see [below for nested schema](#nestedatt--ipv4_export_route_targets))
- `ipv4_import_route_targets` (Attributes List) (see [below for nested schema](#nestedatt--ipv4_import_route_targets))
- `ipv4_static_routes` (Attributes List) IPv4 Static Route (see [below for nested schema](#nestedatt--ipv4_static_routes))
- `ipv6_export_route_targets` (Attributes List) (see [below for nested schema](#nestedatt--ipv6_export_route_targets))
- `ipv6_import_route_targets` (Attributes List) (see [below for nested schema](#nestedatt--ipv6_import_route_targets))
- `ipv6_static_routes` (Attributes List) IPv6 Static Route (see [below for nested schema](#nestedatt--ipv6_static_routes))
- `nat_64_v4_pools` (Attributes List) NAT64 V4 Pool (see [below for nested schema](#nestedatt--nat_64_v4_pools))
- `nat_pools` (Attributes List) NAT Pool (see [below for nested schema](#nestedatt--nat_pools))
- `nat_port_forwards` (Attributes List) NAT Port Forward (see [below for nested schema](#nestedatt--nat_port_forwards))
- `omp_admin_distance_ipv4` (Number) OMP Admin Distance IPv4
  - Range: `1`-`255`
- `omp_admin_distance_ipv4_variable` (String) Variable name
- `omp_admin_distance_ipv6` (Number) OMP Admin Distance IPv6
  - Range: `1`-`255`
- `omp_admin_distance_ipv6_variable` (String) Variable name
- `primary_dns_address_ipv4` (String) Primary DNS Address (IPv4)
- `primary_dns_address_ipv4_variable` (String) Variable name
- `primary_dns_address_ipv6` (String) Primary DNS Address (IPv6)
- `primary_dns_address_ipv6_variable` (String) Variable name
- `route_leak_from_global_vpns` (Attributes List) Enable route leaking from Global to Service VPN (see [below for nested schema](#nestedatt--route_leak_from_global_vpns))
- `route_leak_from_other_services` (Attributes List) Enable route leak from another Service VPN to current Service VPN (see [below for nested schema](#nestedatt--route_leak_from_other_services))
- `route_leak_to_global_vpns` (Attributes List) Enable route leaking from Service to Global VPN (see [below for nested schema](#nestedatt--route_leak_to_global_vpns))
- `secondary_dns_address_ipv4` (String) Secondary DNS Address (IPv4)
- `secondary_dns_address_ipv4_variable` (String) Variable name
- `secondary_dns_address_ipv6` (String) Secondary DNS Address (IPv6)
- `secondary_dns_address_ipv6_variable` (String) Variable name
- `service_routes` (Attributes List) Service (see [below for nested schema](#nestedatt--service_routes))
- `services` (Attributes List) Service (see [below for nested schema](#nestedatt--services))
- `static_nats` (Attributes List) Static NAT Rules (see [below for nested schema](#nestedatt--static_nats))
- `vpn` (Number) VPN
  - Range: `1`-`65527`
  - Default value: `0`
- `vpn_variable` (String) Variable name

### Read-Only

- `id` (String) The id of the Feature
- `version` (Number) The version of the Feature

<a id="nestedatt--advertise_omp_ipv4s"></a>
### Nested Schema for `advertise_omp_ipv4s`

Optional:

- `prefixes` (Attributes List) IPv4 Prefix List (see [below for nested schema](#nestedatt--advertise_omp_ipv4s--prefixes))
- `protocol` (String) Protocol
  - Choices: `bgp`, `ospf`, `ospfv3`, `connected`, `static`, `network`, `aggregate`, `eigrp`, `lisp`, `isis`
- `protocol_variable` (String) Variable name
- `route_policy_id` (String)

<a id="nestedatt--advertise_omp_ipv4s--prefixes"></a>
### Nested Schema for `advertise_omp_ipv4s.prefixes`

Optional:

- `aggregate_only` (Boolean) Aggregate Only
  - Default value: `false`
- `network_address` (String)
- `network_address_variable` (String) Variable name
- `region` (String) Applied to Region
  - Choices: `core-and-access`, `core`, `access`
  - Default value: `core-and-access`
- `region_variable` (String) Variable name
- `subnet_mask` (String) - Choices: `255.255.255.255`, `255.255.255.254`, `255.255.255.252`, `255.255.255.248`, `255.255.255.240`, `255.255.255.224`, `255.255.255.192`, `255.255.255.128`, `255.255.255.0`, `255.255.254.0`, `255.255.252.0`, `255.255.248.0`, `255.255.240.0`, `255.255.224.0`, `255.255.192.0`, `255.255.128.0`, `255.255.0.0`, `255.254.0.0`, `255.252.0.0`, `255.240.0.0`, `255.224.0.0`, `255.192.0.0`, `255.128.0.0`, `255.0.0.0`, `254.0.0.0`, `252.0.0.0`, `248.0.0.0`, `240.0.0.0`, `224.0.0.0`, `192.0.0.0`, `128.0.0.0`, `0.0.0.0`
- `subnet_mask_variable` (String) Variable name



<a id="nestedatt--advertise_omp_ipv6s"></a>
### Nested Schema for `advertise_omp_ipv6s`

Optional:

- `prefixes` (Attributes List) IPv6 Prefix List (see [below for nested schema](#nestedatt--advertise_omp_ipv6s--prefixes))
- `protocol` (String) Protocol
  - Choices: `BGP`, `OSPF`, `Connected`, `Static`, `Network`, `Aggregate`
- `protocol_sub_type` (String) Protocol Sub Type
  - Choices: `External`
- `protocol_sub_type_variable` (String) Variable name
- `protocol_variable` (String) Variable name
- `route_policy_id` (String)

<a id="nestedatt--advertise_omp_ipv6s--prefixes"></a>
### Nested Schema for `advertise_omp_ipv6s.prefixes`

Optional:

- `aggregate_only` (Boolean) Aggregate Only
  - Default value: `false`
- `prefix` (String) IPv6 Prefix
- `prefix_variable` (String) Variable name



<a id="nestedatt--gre_routes"></a>
### Nested Schema for `gre_routes`

Optional:

- `interface` (Set of String) Interface
- `interface_variable` (String) Variable name
- `network_address` (String) IP Address
- `network_address_variable` (String) Variable name
- `subnet_mask` (String) Subnet Mask
  - Choices: `255.255.255.255`, `255.255.255.254`, `255.255.255.252`, `255.255.255.248`, `255.255.255.240`, `255.255.255.224`, `255.255.255.192`, `255.255.255.128`, `255.255.255.0`, `255.255.254.0`, `255.255.252.0`, `255.255.248.0`, `255.255.240.0`, `255.255.224.0`, `255.255.192.0`, `255.255.128.0`, `255.255.0.0`, `255.254.0.0`, `255.252.0.0`, `255.240.0.0`, `255.224.0.0`, `255.192.0.0`, `255.128.0.0`, `255.0.0.0`, `254.0.0.0`, `252.0.0.0`, `248.0.0.0`, `240.0.0.0`, `224.0.0.0`, `192.0.0.0`, `128.0.0.0`, `0.0.0.0`
- `subnet_mask_variable` (String) Variable name
- `vpn` (Number) Service


<a id="nestedatt--host_mappings"></a>
### Nested Schema for `host_mappings`

Optional:

- `host_name` (String) Hostname
- `host_name_variable` (String) Variable name
- `list_of_ips` (Set of String) List of IP
- `list_of_ips_variable` (String) Variable name


<a id="nestedatt--ipsec_routes"></a>
### Nested Schema for `ipsec_routes`

Optional:

- `interface` (Set of String) Interface
- `interface_variable` (String) Variable name
- `network_address` (String) IP Address
- `network_address_variable` (String) Variable name
- `subnet_mask` (String) Subnet Mask
  - Choices: `255.255.255.255`, `255.255.255.254`, `255.255.255.252`, `255.255.255.248`, `255.255.255.240`, `255.255.255.224`, `255.255.255.192`, `255.255.255.128`, `255.255.255.0`, `255.255.254.0`, `255.255.252.0`, `255.255.248.0`, `255.255.240.0`, `255.255.224.0`, `255.255.192.0`, `255.255.128.0`, `255.255.0.0`, `255.254.0.0`, `255.252.0.0`, `255.240.0.0`, `255.224.0.0`, `255.192.0.0`, `255.128.0.0`, `255.0.0.0`, `254.0.0.0`, `252.0.0.0`, `248.0.0.0`, `240.0.0.0`, `224.0.0.0`, `192.0.0.0`, `128.0.0.0`, `0.0.0.0`
- `subnet_mask_variable` (String) Variable name


<a id="nestedatt--ipv4_export_route_targets"></a>
### Nested Schema for `ipv4_export_route_targets`

Optional:

- `route_target` (String) Route target
- `route_target_variable` (String) Variable name


<a id="nestedatt--ipv4_import_route_targets"></a>
### Nested Schema for `ipv4_import_route_targets`

Optional:

- `route_target` (String) Route target
- `route_target_variable` (String) Variable name


<a id="nestedatt--ipv4_static_routes"></a>
### Nested Schema for `ipv4_static_routes`

Optional:

- `gateway_dhcp` (Boolean) IPv4 Route Gateway DHCP
- `network_address` (String) IP Address
- `network_address_variable` (String) Variable name
- `next_hop_with_trackers` (Attributes List) IPv4 Route Gateway Next Hop with Tracker (see [below for nested schema](#nestedatt--ipv4_static_routes--next_hop_with_trackers))
- `next_hops` (Attributes List) IPv4 Route Gateway Next Hop (see [below for nested schema](#nestedatt--ipv4_static_routes--next_hops))
- `null0` (Boolean) IPv4 Route Gateway Next Hop
- `subnet_mask` (String) Subnet Mask
  - Choices: `255.255.255.255`, `255.255.255.254`, `255.255.255.252`, `255.255.255.248`, `255.255.255.240`, `255.255.255.224`, `255.255.255.192`, `255.255.255.128`, `255.255.255.0`, `255.255.254.0`, `255.255.252.0`, `255.255.248.0`, `255.255.240.0`, `255.255.224.0`, `255.255.192.0`, `255.255.128.0`, `255.255.0.0`, `255.254.0.0`, `255.252.0.0`, `255.240.0.0`, `255.224.0.0`, `255.192.0.0`, `255.128.0.0`, `255.0.0.0`, `254.0.0.0`, `252.0.0.0`, `248.0.0.0`, `240.0.0.0`, `224.0.0.0`, `192.0.0.0`, `128.0.0.0`, `0.0.0.0`
- `subnet_mask_variable` (String) Variable name
- `vpn` (Boolean) IPv4 Route Gateway VPN

<a id="nestedatt--ipv4_static_routes--next_hop_with_trackers"></a>
### Nested Schema for `ipv4_static_routes.next_hop_with_trackers`

Optional:

- `address` (String) Address
- `address_variable` (String) Variable name
- `administrative_distance` (Number) Administrative distance
  - Range: `1`-`255`
- `administrative_distance_variable` (String) Variable name
- `tracker_id` (String)


<a id="nestedatt--ipv4_static_routes--next_hops"></a>
### Nested Schema for `ipv4_static_routes.next_hops`

Optional:

- `address` (String) Address
- `address_variable` (String) Variable name
- `administrative_distance` (Number) Administrative distance
  - Range: `1`-`255`
- `administrative_distance_variable` (String) Variable name



<a id="nestedatt--ipv6_export_route_targets"></a>
### Nested Schema for `ipv6_export_route_targets`

Optional:

- `route_target` (String) Route target
- `route_target_variable` (String) Variable name


<a id="nestedatt--ipv6_import_route_targets"></a>
### Nested Schema for `ipv6_import_route_targets`

Optional:

- `route_target` (String) Route target
- `route_target_variable` (String) Variable name


<a id="nestedatt--ipv6_static_routes"></a>
### Nested Schema for `ipv6_static_routes`

Optional:

- `nat` (String) IPv6 Nat
  - Choices: `NAT64`, `NAT66`
- `nat_variable` (String) Variable name
- `next_hops` (Attributes List) IPv6 Route Gateway Next Hop (see [below for nested schema](#nestedatt--ipv6_static_routes--next_hops))
- `null0` (Boolean) IPv6 Route Gateway Next Hop
- `prefix` (String) Prefix
- `prefix_variable` (String) Variable name

<a id="nestedatt--ipv6_static_routes--next_hops"></a>
### Nested Schema for `ipv6_static_routes.next_hops`

Optional:

- `address` (String) Address
- `address_variable` (String) Variable name
- `administrative_distance` (Number) Administrative distance
  - Range: `1`-`254`
- `administrative_distance_variable` (String) Variable name



<a id="nestedatt--nat_64_v4_pools"></a>
### Nested Schema for `nat_64_v4_pools`

Optional:

- `name` (String) NAT64 v4 Pool Name
- `name_variable` (String) Variable name
- `overload` (Boolean) NAT64 Overload
  - Default value: `false`
- `overload_variable` (String) Variable name
- `range_end` (String) NAT64 Pool Range End
- `range_end_variable` (String) Variable name
- `range_start` (String) NAT64 Pool Range Start
- `range_start_variable` (String) Variable name


<a id="nestedatt--nat_pools"></a>
### Nested Schema for `nat_pools`

Optional:

- `direction` (String) NAT Direction
  - Choices: `inside`, `outside`
- `direction_variable` (String) Variable name
- `nat_pool_name` (Number) NAT Pool Name
  - Range: `1`-`32`
- `nat_pool_name_variable` (String) Variable name
- `overload` (Boolean) NAT Overload
  - Default value: `true`
- `overload_variable` (String) Variable name
- `prefix_length` (Number) NAT Pool Prefix Length
  - Range: `1`-`32`
- `prefix_length_variable` (String) Variable name
- `range_end` (String) NAT Pool Range End
- `range_end_variable` (String) Variable name
- `range_start` (String) NAT Pool Range Start
- `range_start_variable` (String) Variable name
- `tracker_object_id` (String)


<a id="nestedatt--nat_port_forwards"></a>
### Nested Schema for `nat_port_forwards`

Optional:

- `nat_pool_name` (Number) NAT Pool Name
  - Range: `1`-`32`
- `nat_pool_name_variable` (String) Variable name
- `protocol` (String) Protocol
  - Choices: `TCP`, `UDP`
- `protocol_variable` (String) Variable name
- `source_ip` (String) Source IP Address
- `source_ip_variable` (String) Variable name
- `source_port` (Number) Source Port
- `source_port_variable` (String) Variable name
- `translate_port` (Number) Translate Port
- `translate_port_variable` (String) Variable name
- `translated_source_ip` (String) Translated Source IP Address
- `translated_source_ip_variable` (String) Variable name


<a id="nestedatt--route_leak_from_global_vpns"></a>
### Nested Schema for `route_leak_from_global_vpns`

Optional:

- `redistributions` (Attributes List) Redistribute Routes to specific Protocol on Service VPN (see [below for nested schema](#nestedatt--route_leak_from_global_vpns--redistributions))
- `route_policy_id` (String)
- `route_protocol` (String) Leak Routes of particular protocol from Global to Service VPN
  - Choices: `static`, `connected`, `bgp`, `ospf`
- `route_protocol_variable` (String) Variable name

<a id="nestedatt--route_leak_from_global_vpns--redistributions"></a>
### Nested Schema for `route_leak_from_global_vpns.redistributions`

Optional:

- `protocol` (String) Protocol to restributed leaked routes
  - Choices: `bgp`, `ospf`
- `protocol_variable` (String) Variable name
- `redistribution_policy_id` (String)



<a id="nestedatt--route_leak_from_other_services"></a>
### Nested Schema for `route_leak_from_other_services`

Optional:

- `redistributions` (Attributes List) Redistribute Route to specific Protocol on Current Service VPN (see [below for nested schema](#nestedatt--route_leak_from_other_services--redistributions))
- `route_policy_id` (String)
- `route_protocol` (String) Leak Route of particular protocol from Source Service VPN
  - Choices: `static`, `connected`, `bgp`, `ospf`
- `route_protocol_variable` (String) Variable name
- `source_vpn` (Number) Source Service VPN from where route are to be leaked
  - Range: `1`-`65530`
- `source_vpn_variable` (String) Variable name

<a id="nestedatt--route_leak_from_other_services--redistributions"></a>
### Nested Schema for `route_leak_from_other_services.redistributions`

Optional:

- `protocol` (String) Protocol to restributed leaked routes
  - Choices: `bgp`, `ospf`
- `protocol_variable` (String) Variable name
- `redistribution_policy_id` (String)



<a id="nestedatt--route_leak_to_global_vpns"></a>
### Nested Schema for `route_leak_to_global_vpns`

Optional:

- `redistributions` (Attributes List) Redistribute Routes to specific Protocol on Global VPN (see [below for nested schema](#nestedatt--route_leak_to_global_vpns--redistributions))
- `route_policy_id` (String)
- `route_protocol` (String) Leak Routes of particular protocol from Service to Global VPN
  - Choices: `static`, `connected`, `bgp`, `ospf`
- `route_protocol_variable` (String) Variable name

<a id="nestedatt--route_leak_to_global_vpns--redistributions"></a>
### Nested Schema for `route_leak_to_global_vpns.redistributions`

Optional:

- `protocol` (String) Protocol to restributed leaked routes
  - Choices: `bgp`, `ospf`
- `protocol_variable` (String) Variable name
- `redistribution_policy_id` (String)



<a id="nestedatt--service_routes"></a>
### Nested Schema for `service_routes`

Optional:

- `network_address` (String) IP Address
- `network_address_variable` (String) Variable name
- `service` (String) Service
  - Choices: `SIG`
  - Default value: `SIG`
- `service_variable` (String) Variable name
- `subnet_mask` (String) Subnet Mask
  - Choices: `255.255.255.255`, `255.255.255.254`, `255.255.255.252`, `255.255.255.248`, `255.255.255.240`, `255.255.255.224`, `255.255.255.192`, `255.255.255.128`, `255.255.255.0`, `255.255.254.0`, `255.255.252.0`, `255.255.248.0`, `255.255.240.0`, `255.255.224.0`, `255.255.192.0`, `255.255.128.0`, `255.255.0.0`, `255.254.0.0`, `255.252.0.0`, `255.240.0.0`, `255.224.0.0`, `255.192.0.0`, `255.128.0.0`, `255.0.0.0`, `254.0.0.0`, `252.0.0.0`, `248.0.0.0`, `240.0.0.0`, `224.0.0.0`, `192.0.0.0`, `128.0.0.0`, `0.0.0.0`
- `subnet_mask_variable` (String) Variable name
- `vpn` (Number) Service


<a id="nestedatt--services"></a>
### Nested Schema for `services`

Optional:

- `ipv4_addresses` (Set of String) IPv4 Addresses (Maximum: 4)
- `ipv4_addresses_variable` (String) Variable name
- `service_type` (String) Service Type
  - Choices: `FW`, `IDS`, `IDP`, `netsvc1`, `netsvc2`, `netsvc3`, `netsvc4`, `TE`, `appqoe`
- `service_type_variable` (String) Variable name
- `tracking` (Boolean) Tracking
  - Default value: `true`
- `tracking_variable` (String) Variable name


<a id="nestedatt--static_nats"></a>
### Nested Schema for `static_nats`

Optional:

- `nat_pool_name` (Number) NAT Pool Name
  - Range: `1`-`32`
- `nat_pool_name_variable` (String) Variable name
- `source_ip` (String) Source IP Address
- `source_ip_variable` (String) Variable name
- `static_nat_direction` (String) Static NAT Direction
  - Choices: `inside`, `outside`
- `static_nat_direction_variable` (String) Variable name
- `tracker_object_id` (String)
- `translated_source_ip` (String) Translated Source IP Address
- `translated_source_ip_variable` (String) Variable name

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Expected import identifier with the format: "service_lan_vpn_feature_id,feature_profile_id"
terraform import sdwan_service_lan_vpn_feature.example "f6b2c44c-693c-4763-b010-895aa3d236bd,f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
```
