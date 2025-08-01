---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_cisco_ospfv3_feature_template Resource - terraform-provider-sdwan"
subcategory: "(Classic) Feature Templates"
description: |-
  This resource can manage a Cisco OSPFv3 feature template.
  Minimum SD-WAN Manager version: 15.0.0
---

# sdwan_cisco_ospfv3_feature_template (Resource)

This resource can manage a Cisco OSPFv3 feature template.
  - Minimum SD-WAN Manager version: `15.0.0`

## Example Usage

```terraform
resource "sdwan_cisco_ospfv3_feature_template" "example" {
  name                                           = "Example"
  description                                    = "My Example"
  device_types                                   = ["vedge-C8000V"]
  ipv4_router_id                                 = "1.2.3.4"
  ipv4_auto_cost_reference_bandwidth             = 100000
  ipv4_compatible_rfc1583                        = true
  ipv4_default_information_originate             = true
  ipv4_default_information_originate_always      = true
  ipv4_default_information_originate_metric      = 100
  ipv4_default_information_originate_metric_type = "type1"
  ipv4_distance_external                         = 111
  ipv4_distance_inter_area                       = 111
  ipv4_distance_intra_area                       = 112
  ipv4_timers_spf_delay                          = 300
  ipv4_timers_spf_initial_hold                   = 2000
  ipv4_timers_spf_max_hold                       = 20000
  ipv4_distance                                  = 110
  ipv4_policy_name                               = "POLICY1"
  ipv4_filter                                    = false
  ipv4_redistributes = [
    {
      protocol     = "static"
      route_policy = "RP1"
      nat_dia      = true
    }
  ]
  ipv4_max_metric_router_lsas = [
    {
      ad_type = "on-startup"
      time    = 100
    }
  ]
  ipv4_areas = [
    {
      area_number     = 1
      stub            = false
      stub_no_summary = false
      nssa            = false
      nssa_no_summary = true
      translate       = "always"
      normal          = false
      interfaces = [
        {
          name                = "e1"
          hello_interval      = 20
          dead_interval       = 60
          retransmit_interval = 10
          cost                = 100
          network             = "point-to-point"
          passive_interface   = true
          authentication_type = "md5"
          authentication_key  = "authenticationKey"
          ipsec_spi           = 256
        }
      ]
      ranges = [
        {
          address      = "1.1.1.0/24"
          cost         = 100
          no_advertise = true
        }
      ]
    }
  ]
  ipv6_router_id                                 = "1.2.3.4"
  ipv6_auto_cost_reference_bandwidth             = 100000
  ipv6_compatible_rfc1583                        = true
  ipv6_default_information_originate             = true
  ipv6_default_information_originate_always      = true
  ipv6_default_information_originate_metric      = 100
  ipv6_default_information_originate_metric_type = "type1"
  ipv6_distance_external                         = 111
  ipv6_distance_inter_area                       = 111
  ipv6_distance_intra_area                       = 112
  ipv6_timers_spf_delay                          = 300
  ipv6_timers_spf_initial_hold                   = 2000
  ipv6_timers_spf_max_hold                       = 20000
  ipv6_distance                                  = 110
  ipv6_policy_name                               = "POLICY2"
  ipv6_filter                                    = false
  ipv6_redistributes = [
    {
      protocol     = "static"
      route_policy = "RP1"
    }
  ]
  ipv6_max_metric_router_lsas = [
    {
      ad_type = "on-startup"
      time    = 100
    }
  ]
  ipv6_areas = [
    {
      area_number     = 1
      stub            = false
      stub_no_summary = false
      nssa            = false
      nssa_no_summary = true
      translate       = "always"
      normal          = false
      interfaces = [
        {
          name                = "e1"
          hello_interval      = 20
          dead_interval       = 60
          retransmit_interval = 10
          cost                = 100
          network             = "point-to-point"
          passive_interface   = true
          authentication_type = "md5"
          authentication_key  = "authenticationKey"
          ipsec_spi           = 256
        }
      ]
      ranges = [
        {
          address      = "2001::/48"
          cost         = 100
          no_advertise = true
        }
      ]
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) The description of the feature template
- `device_types` (Set of String) List of supported device types
  - Choices: `vedge-C8000V`, `vedge-C8300-1N1S-4T2X`, `vedge-C8300-1N1S-6T`, `vedge-C8300-2N2S-6T`, `vedge-C8300-2N2S-4T2X`, `vedge-C8500-12X4QC`, `vedge-C8500-12X`, `vedge-C8500-20X6C`, `vedge-C8500L-8S4X`, `vedge-C8200-1N-4T`, `vedge-C8200L-1N-4T`
- `name` (String) The name of the feature template

### Optional

- `ipv4_areas` (Attributes List) Configure OSPF area (see [below for nested schema](#nestedatt--ipv4_areas))
- `ipv4_auto_cost_reference_bandwidth` (Number) Set reference bandwidth method to assign OSPF cost
  - Range: `1`-`4294967`
  - Default value: `100`
- `ipv4_auto_cost_reference_bandwidth_variable` (String) Variable name
- `ipv4_compatible_rfc1583` (Boolean) Calculate summary route cost based on RFC 1583
  - Default value: `true`
- `ipv4_compatible_rfc1583_variable` (String) Variable name
- `ipv4_default_information_originate` (Boolean) Distribute default external route into OSPF
  - Default value: `false`
- `ipv4_default_information_originate_always` (Boolean) Always advertise default route
  - Default value: `false`
- `ipv4_default_information_originate_always_variable` (String) Variable name
- `ipv4_default_information_originate_metric` (Number) Set metric used to generate default route <0..16777214>
  - Range: `0`-`16777214`
- `ipv4_default_information_originate_metric_type` (String) Set default route type
  - Choices: `type1`, `type2`
- `ipv4_default_information_originate_metric_type_variable` (String) Variable name
- `ipv4_default_information_originate_metric_variable` (String) Variable name
- `ipv4_distance` (Number) Distance
  - Range: `1`-`255`
  - Default value: `110`
- `ipv4_distance_external` (Number) Set distance for external routes
  - Range: `1`-`254`
  - Default value: `110`
- `ipv4_distance_external_variable` (String) Variable name
- `ipv4_distance_inter_area` (Number) Set distance for inter-area routes
  - Range: `1`-`254`
  - Default value: `110`
- `ipv4_distance_inter_area_variable` (String) Variable name
- `ipv4_distance_intra_area` (Number) Set distance for intra-area routes
  - Range: `1`-`254`
  - Default value: `110`
- `ipv4_distance_intra_area_variable` (String) Variable name
- `ipv4_distance_variable` (String) Variable name
- `ipv4_filter` (Boolean) Filter
  - Default value: `false`
- `ipv4_filter_variable` (String) Variable name
- `ipv4_max_metric_router_lsas` (Attributes List) Advertise own router LSA with infinite distance (see [below for nested schema](#nestedatt--ipv4_max_metric_router_lsas))
- `ipv4_policy_name` (String) Policy Name
- `ipv4_policy_name_variable` (String) Variable name
- `ipv4_redistributes` (Attributes List) Redistribute routes (see [below for nested schema](#nestedatt--ipv4_redistributes))
- `ipv4_router_id` (String) Set OSPF router ID to override system IP address
- `ipv4_router_id_variable` (String) Variable name
- `ipv4_timers_spf_delay` (Number) Set delay from first change received until performing SPF calculation
  - Range: `1`-`600000`
  - Default value: `200`
- `ipv4_timers_spf_delay_variable` (String) Variable name
- `ipv4_timers_spf_initial_hold` (Number) Set initial hold time between consecutive SPF calculations
  - Range: `1`-`600000`
  - Default value: `1000`
- `ipv4_timers_spf_initial_hold_variable` (String) Variable name
- `ipv4_timers_spf_max_hold` (Number) Set maximum hold time between consecutive SPF calculations
  - Range: `1`-`600000`
  - Default value: `10000`
- `ipv4_timers_spf_max_hold_variable` (String) Variable name
- `ipv6_areas` (Attributes List) Configure OSPF area (see [below for nested schema](#nestedatt--ipv6_areas))
- `ipv6_auto_cost_reference_bandwidth` (Number) Set reference bandwidth method to assign OSPF cost
  - Range: `1`-`4294967`
  - Default value: `100`
- `ipv6_auto_cost_reference_bandwidth_variable` (String) Variable name
- `ipv6_compatible_rfc1583` (Boolean) Calculate summary route cost based on RFC 1583
  - Default value: `true`
- `ipv6_compatible_rfc1583_variable` (String) Variable name
- `ipv6_default_information_originate` (Boolean) Distribute default external route into OSPF
  - Default value: `false`
- `ipv6_default_information_originate_always` (Boolean) Always advertise default route
  - Default value: `false`
- `ipv6_default_information_originate_always_variable` (String) Variable name
- `ipv6_default_information_originate_metric` (Number) Set metric used to generate default route <0..16777214>
  - Range: `0`-`16777214`
- `ipv6_default_information_originate_metric_type` (String) Set default route type
  - Choices: `type1`, `type2`
- `ipv6_default_information_originate_metric_type_variable` (String) Variable name
- `ipv6_default_information_originate_metric_variable` (String) Variable name
- `ipv6_distance` (Number) Distance
  - Range: `1`-`254`
  - Default value: `110`
- `ipv6_distance_external` (Number) Set distance for external routes
  - Range: `1`-`254`
  - Default value: `110`
- `ipv6_distance_external_variable` (String) Variable name
- `ipv6_distance_inter_area` (Number) Set distance for inter-area routes
  - Range: `1`-`254`
  - Default value: `110`
- `ipv6_distance_inter_area_variable` (String) Variable name
- `ipv6_distance_intra_area` (Number) Set distance for intra-area routes
  - Range: `1`-`254`
  - Default value: `110`
- `ipv6_distance_intra_area_variable` (String) Variable name
- `ipv6_distance_variable` (String) Variable name
- `ipv6_filter` (Boolean) Filter
  - Default value: `false`
- `ipv6_filter_variable` (String) Variable name
- `ipv6_max_metric_router_lsas` (Attributes List) Advertise own router LSA with infinite distance (see [below for nested schema](#nestedatt--ipv6_max_metric_router_lsas))
- `ipv6_policy_name` (String) Name
- `ipv6_policy_name_variable` (String) Variable name
- `ipv6_redistributes` (Attributes List) Redistribute routes (see [below for nested schema](#nestedatt--ipv6_redistributes))
- `ipv6_router_id` (String) Set OSPF router ID to override system IP address
- `ipv6_router_id_variable` (String) Variable name
- `ipv6_timers_spf_delay` (Number) Set delay from first change received until performing SPF calculation
  - Range: `0`-`600000`
  - Default value: `200`
- `ipv6_timers_spf_delay_variable` (String) Variable name
- `ipv6_timers_spf_initial_hold` (Number) Set initial hold time between consecutive SPF calculations
  - Range: `0`-`600000`
  - Default value: `1000`
- `ipv6_timers_spf_initial_hold_variable` (String) Variable name
- `ipv6_timers_spf_max_hold` (Number) Set maximum hold time between consecutive SPF calculations
  - Range: `0`-`600000`
  - Default value: `10000`
- `ipv6_timers_spf_max_hold_variable` (String) Variable name

### Read-Only

- `id` (String) The id of the feature template
- `template_type` (String) The template type
- `version` (Number) The version of the feature template

<a id="nestedatt--ipv4_areas"></a>
### Nested Schema for `ipv4_areas`

Optional:

- `area_number` (Number) Set OSPF area number
  - Range: `0`-`4294967295`
- `area_number_variable` (String) Variable name
- `interfaces` (Attributes List) Set OSPF interface parameters (see [below for nested schema](#nestedatt--ipv4_areas--interfaces))
- `normal` (Boolean) Area Type Normal
  - Default value: `false`
- `normal_variable` (String) Variable name
- `nssa` (Boolean) NSSA area
- `nssa_no_summary` (Boolean) Do not inject interarea routes into NSSA
  - Default value: `false`
- `nssa_no_summary_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `ranges` (Attributes List) Summarize OSPF routes at an area boundary (see [below for nested schema](#nestedatt--ipv4_areas--ranges))
- `stub` (Boolean) Stub area
- `stub_no_summary` (Boolean) Do not inject interarea routes into stub
  - Default value: `false`
- `stub_no_summary_variable` (String) Variable name
- `translate` (String) Always Translate LSAs on this ABR
  - Choices: `always`
- `translate_variable` (String) Variable name

<a id="nestedatt--ipv4_areas--interfaces"></a>
### Nested Schema for `ipv4_areas.interfaces`

Optional:

- `authentication_key` (String) Set OSPF interface authentication key
- `authentication_key_variable` (String) Variable name
- `authentication_type` (String) Set OSPF interface authentication type
  - Choices: `md5`, `sha1`
- `authentication_type_variable` (String) Variable name
- `cost` (Number) Set cost of OSPF interface
  - Range: `1`-`65535`
- `cost_variable` (String) Variable name
- `dead_interval` (Number) Set interval after which neighbor is declared to be down
  - Range: `1`-`65535`
  - Default value: `40`
- `dead_interval_variable` (String) Variable name
- `hello_interval` (Number) Set interval between OSPF hello packets
  - Range: `1`-`65535`
  - Default value: `10`
- `hello_interval_variable` (String) Variable name
- `ipsec_spi` (Number) Set OSPF interface authentication IPSec SPI, range 256..4294967295
  - Range: `256`-`4294967295`
- `ipsec_spi_variable` (String) Variable name
- `name` (String) Set interface name
- `name_variable` (String) Variable name
- `network` (String) Set the OSPF network type
  - Choices: `broadcast`, `point-to-point`, `non-broadcast`, `point-to-multipoint`
  - Default value: `broadcast`
- `network_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `passive_interface` (Boolean) Set the interface to advertise its address, but not to actively run OSPF
  - Default value: `false`
- `passive_interface_variable` (String) Variable name
- `retransmit_interval` (Number) Set time between retransmitting LSAs
  - Range: `1`-`65535`
  - Default value: `5`
- `retransmit_interval_variable` (String) Variable name


<a id="nestedatt--ipv4_areas--ranges"></a>
### Nested Schema for `ipv4_areas.ranges`

Optional:

- `address` (String) Set Matching Prefix
- `address_variable` (String) Variable name
- `cost` (Number) Set cost for this range
  - Range: `0`-`16777214`
- `cost_variable` (String) Variable name
- `no_advertise` (Boolean) Do not advertise this range
  - Default value: `false`
- `no_advertise_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.



<a id="nestedatt--ipv4_max_metric_router_lsas"></a>
### Nested Schema for `ipv4_max_metric_router_lsas`

Optional:

- `ad_type` (String) Set the router LSA advertisement type
  - Choices: `on-startup`
- `optional` (Boolean) Indicates if list item is considered optional.
- `time` (Number) Set how long to advertise maximum metric after router starts up
  - Range: `5`-`86400`
- `time_variable` (String) Variable name


<a id="nestedatt--ipv4_redistributes"></a>
### Nested Schema for `ipv4_redistributes`

Optional:

- `nat_dia` (Boolean) Enable NAT DIA for redistributed routes
  - Default value: `true`
- `nat_dia_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `protocol` (String) Set the protocol
  - Choices: `bgp`, `connected`, `eigrp`, `isis`, `lisp`, `nat-route`, `omp`, `static`
- `protocol_variable` (String) Variable name
- `route_policy` (String) Set route policy to apply to redistributed routes
- `route_policy_variable` (String) Variable name


<a id="nestedatt--ipv6_areas"></a>
### Nested Schema for `ipv6_areas`

Optional:

- `area_number` (Number) Set OSPF area number
  - Range: `0`-`4294967295`
- `area_number_variable` (String) Variable name
- `interfaces` (Attributes List) Set OSPF interface parameters (see [below for nested schema](#nestedatt--ipv6_areas--interfaces))
- `normal` (Boolean) Area Type Normal
  - Default value: `false`
- `normal_variable` (String) Variable name
- `nssa` (Boolean) NSSA area
- `nssa_no_summary` (Boolean) Do not inject interarea routes into NSSA
  - Default value: `false`
- `nssa_no_summary_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `ranges` (Attributes List) Summarize OSPF routes at an area boundary (see [below for nested schema](#nestedatt--ipv6_areas--ranges))
- `stub` (Boolean) Stub area
- `stub_no_summary` (Boolean) Do not inject interarea routes into stub
  - Default value: `false`
- `stub_no_summary_variable` (String) Variable name
- `translate` (String) Always translate LSAs on this ABR
  - Choices: `always`
- `translate_variable` (String) Variable name

<a id="nestedatt--ipv6_areas--interfaces"></a>
### Nested Schema for `ipv6_areas.interfaces`

Optional:

- `authentication_key` (String) Set OSPF interface authentication key
- `authentication_key_variable` (String) Variable name
- `authentication_type` (String) Set OSPF interface authentication type
  - Choices: `md5`, `sha1`
- `authentication_type_variable` (String) Variable name
- `cost` (Number) Set cost of OSPF interface
  - Range: `1`-`65535`
- `cost_variable` (String) Variable name
- `dead_interval` (Number) Set interval after which neighbor is declared to be down
  - Range: `1`-`65535`
  - Default value: `40`
- `dead_interval_variable` (String) Variable name
- `hello_interval` (Number) Set interval between OSPF hello packets
  - Range: `1`-`65535`
  - Default value: `10`
- `hello_interval_variable` (String) Variable name
- `ipsec_spi` (Number) Set OSPF interface authentication IPSec SPI, range 256..4294967295
  - Range: `256`-`4294967295`
- `ipsec_spi_variable` (String) Variable name
- `name` (String) Set interface name
- `name_variable` (String) Variable name
- `network` (String) Set the OSPF network type
  - Choices: `broadcast`, `point-to-point`, `non-broadcast`, `point-to-multipoint`
  - Default value: `broadcast`
- `network_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `passive_interface` (Boolean) Set the interface to advertise its address, but not to actively run OSPF
  - Default value: `false`
- `passive_interface_variable` (String) Variable name
- `retransmit_interval` (Number) Set time between retransmitting LSAs
  - Range: `1`-`65535`
  - Default value: `5`
- `retransmit_interval_variable` (String) Variable name


<a id="nestedatt--ipv6_areas--ranges"></a>
### Nested Schema for `ipv6_areas.ranges`

Optional:

- `address` (String) Set Matching Prefix
- `address_variable` (String) Variable name
- `cost` (Number) Set cost for this range
  - Range: `0`-`16777214`
- `cost_variable` (String) Variable name
- `no_advertise` (Boolean) Do not advertise this range
  - Default value: `false`
- `no_advertise_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.



<a id="nestedatt--ipv6_max_metric_router_lsas"></a>
### Nested Schema for `ipv6_max_metric_router_lsas`

Optional:

- `ad_type` (String) Set the router LSA advertisement type
  - Choices: `on-startup`
- `optional` (Boolean) Indicates if list item is considered optional.
- `time` (Number) Set how long to advertise maximum metric after router starts up
- `time_variable` (String) Variable name


<a id="nestedatt--ipv6_redistributes"></a>
### Nested Schema for `ipv6_redistributes`

Optional:

- `optional` (Boolean) Indicates if list item is considered optional.
- `protocol` (String) Set the protocol
  - Choices: `bgp`, `connected`, `eigrp`, `isis`, `lisp`, `nat-route`, `omp`, `static`
- `protocol_variable` (String) Variable name
- `route_policy` (String) Set route policy to apply to redistributed routes
- `route_policy_variable` (String) Variable name

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import sdwan_cisco_ospfv3_feature_template.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```
