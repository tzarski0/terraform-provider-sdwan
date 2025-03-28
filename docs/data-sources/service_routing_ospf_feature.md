---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_service_routing_ospf_feature Data Source - terraform-provider-sdwan"
subcategory: "Features - Service"
description: |-
  This data source can read the Service Routing OSPF Feature.
---

# sdwan_service_routing_ospf_feature (Data Source)

This data source can read the Service Routing OSPF Feature.

## Example Usage

```terraform
data "sdwan_service_routing_ospf_feature" "example" {
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

- `areas` (Attributes List) Configure OSPF area (see [below for nested schema](#nestedatt--areas))
- `default_information_originate` (Boolean) Distribute default external route into OSPF
- `default_information_originate_always` (Boolean) Always advertise default route
- `default_information_originate_always_variable` (String) Variable name
- `default_information_originate_metric` (Number) Set metric used to generate default route <0..16777214>
- `default_information_originate_metric_type` (String) Set default route type
- `default_information_originate_metric_type_variable` (String) Variable name
- `default_information_originate_metric_variable` (String) Variable name
- `description` (String) The description of the Feature
- `distance_external` (Number) Set distance for external routes
- `distance_external_variable` (String) Variable name
- `distance_inter_area` (Number) Set distance for inter-area routes
- `distance_inter_area_variable` (String) Variable name
- `distance_intra_area` (Number) Set distance for intra-area routes
- `distance_intra_area_variable` (String) Variable name
- `name` (String) The name of the Feature
- `redistributes` (Attributes List) Redistribute routes (see [below for nested schema](#nestedatt--redistributes))
- `reference_bandwidth` (Number) Set reference bandwidth method to assign OSPF cost
- `reference_bandwidth_variable` (String) Variable name
- `rfc_1583_compatible` (Boolean) Calculate summary route cost based on RFC 1583
- `rfc_1583_compatible_variable` (String) Variable name
- `route_policy_id` (String)
- `router_id` (String) Set OSPF router ID to override system IP address
- `router_id_variable` (String) Variable name
- `router_lsas` (Attributes List) Advertise own router LSA with infinite distance (see [below for nested schema](#nestedatt--router_lsas))
- `spf_calculation_delay` (Number) Set delay from first change received until performing SPF calculation
- `spf_calculation_delay_variable` (String) Variable name
- `spf_initial_hold_time` (Number) Set initial hold time between consecutive SPF calculations
- `spf_initial_hold_time_variable` (String) Variable name
- `spf_maximum_hold_time` (Number) Set maximum hold time between consecutive SPF calculations
- `spf_maximum_hold_time_variable` (String) Variable name
- `version` (Number) The version of the Feature

<a id="nestedatt--areas"></a>
### Nested Schema for `areas`

Read-Only:

- `area_number` (Number) Set OSPF area number
- `area_number_variable` (String) Variable name
- `area_type` (String) set the area type
- `interfaces` (Attributes List) Set OSPF interface parameters (see [below for nested schema](#nestedatt--areas--interfaces))
- `no_summary` (Boolean) Do not inject interarea routes into STUB or NSSA
- `no_summary_variable` (String) Variable name
- `ranges` (Attributes List) Summarize OSPF routes at an area boundary (see [below for nested schema](#nestedatt--areas--ranges))

<a id="nestedatt--areas--interfaces"></a>
### Nested Schema for `areas.interfaces`

Read-Only:

- `authentication_type` (String) Set OSPF interface authentication type
- `authentication_type_variable` (String) Variable name
- `cost` (Number) Set cost of OSPF interface
- `cost_variable` (String) Variable name
- `dead_interval` (Number) Set interval after which neighbor is declared to be down
- `dead_interval_variable` (String) Variable name
- `designated_router_priority` (Number) Set router’s priority to be elected as designated router
- `designated_router_priority_variable` (String) Variable name
- `hello_interval` (Number) Set interval between OSPF hello packets
- `hello_interval_variable` (String) Variable name
- `lsa_retransmit_interval` (Number) Set time between retransmitting LSAs
- `lsa_retransmit_interval_variable` (String) Variable name
- `message_digest_key` (String) Set MD5 authentication key
- `message_digest_key_id` (Number) Set MD5 message digest key
- `message_digest_key_id_variable` (String) Variable name
- `message_digest_key_variable` (String) Variable name
- `name` (String) Set interface name
- `name_variable` (String) Variable name
- `network_type` (String) Set the OSPF network type
- `network_type_variable` (String) Variable name
- `passive_interface` (Boolean) Set the interface to advertise its address, but not to actively run OSPF
- `passive_interface_variable` (String) Variable name


<a id="nestedatt--areas--ranges"></a>
### Nested Schema for `areas.ranges`

Read-Only:

- `cost` (Number) Set cost for this range
- `cost_variable` (String) Variable name
- `ip_address` (String) IP Address
- `ip_address_variable` (String) Variable name
- `no_advertise` (Boolean) Do not advertise this range
- `no_advertise_variable` (String) Variable name
- `subnet_mask` (String) Subnet Mask
- `subnet_mask_variable` (String) Variable name



<a id="nestedatt--redistributes"></a>
### Nested Schema for `redistributes`

Read-Only:

- `nat_dia` (Boolean) Enable NAT DIA for redistributed routes
- `nat_dia_variable` (String) Variable name
- `protocol` (String) Set the protocol
- `protocol_variable` (String) Variable name
- `route_policy_id` (String)


<a id="nestedatt--router_lsas"></a>
### Nested Schema for `router_lsas`

Read-Only:

- `time` (Number) Set how long to advertise maximum metric after router starts up
- `time_variable` (String) Variable name
- `type` (String) Set the router LSA advertisement type
