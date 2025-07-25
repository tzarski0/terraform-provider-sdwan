// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceSdwanTransportRoutingBGPProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "as_number", "429"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "router_id", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "propagate_as_path", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "propagate_community", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "external_routes_distance", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "internal_routes_distance", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "local_routes_distance", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "keepalive_time", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "hold_time", "180"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "always_compare_med", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "deterministic_med", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "missing_med_as_worst", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "compare_router_id", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "multipath_relax", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.description", "neighbor1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.remote_as", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.local_as", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.keepalive_time", "40"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.hold_time", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.update_source_interface", "GigabitEthernet0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.next_hop_self", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.send_community", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.send_extended_community", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.ebgp_multihop", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.send_label", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.explicit_null", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.as_override", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.allowas_in_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.address_families.0.family_type", "ipv4-unicast"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.address_families.0.policy_type", "restart"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.address_families.0.restart_max_number_of_prefixes", "2000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.address_families.0.restart_threshold", "75"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_neighbors.0.address_families.0.restart_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.address", "2001::1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.description", "neighbor2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.remote_as", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.local_as", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.keepalive_time", "180"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.hold_time", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.update_source_interface", "Loopback1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.next_hop_self", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.send_community", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.send_extended_community", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.ebgp_multihop", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.as_override", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.allowas_in_number", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.address_families.0.family_type", "ipv6-unicast"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.address_families.0.max_number_of_prefixes", "2000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.address_families.0.threshold", "75"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.address_families.0.policy_type", "restart"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_neighbors.0.address_families.0.restart_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_aggregate_addresses.0.network_address", "10.10.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_aggregate_addresses.0.subnet_mask", "255.255.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_aggregate_addresses.0.as_set_path", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_aggregate_addresses.0.summary_only", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_networks.0.network_address", "10.10.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_networks.0.subnet_mask", "255.255.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_eibgp_maximum_paths", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_originate", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv4_table_map_filter", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_aggregate_addresses.0.aggregate_prefix", "3001::1/128"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_aggregate_addresses.0.as_set_path", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_aggregate_addresses.0.summary_only", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_networks.0.network_prefix", "2001:0DB8:0000:000b::/64"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_eibgp_maximum_paths", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_originate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "ipv6_table_map_filter", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_routing_bgp_feature.test", "mpls_interfaces.0.interface_name", "GigabitEthernet1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanTransportRoutingBGPPrerequisitesProfileParcelConfig + testAccDataSourceSdwanTransportRoutingBGPProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanTransportRoutingBGPPrerequisitesProfileParcelConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanTransportRoutingBGPProfileParcelConfig() string {
	config := `resource "sdwan_transport_routing_bgp_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	as_number = 429` + "\n"
	config += `	router_id = "1.2.3.4"` + "\n"
	config += `	propagate_as_path = false` + "\n"
	config += `	propagate_community = false` + "\n"
	config += `	external_routes_distance = 20` + "\n"
	config += `	internal_routes_distance = 200` + "\n"
	config += `	local_routes_distance = 20` + "\n"
	config += `	keepalive_time = 60` + "\n"
	config += `	hold_time = 180` + "\n"
	config += `	always_compare_med = false` + "\n"
	config += `	deterministic_med = false` + "\n"
	config += `	missing_med_as_worst = false` + "\n"
	config += `	compare_router_id = false` + "\n"
	config += `	multipath_relax = false` + "\n"
	config += `	ipv4_neighbors = [{` + "\n"
	config += `	  address = "1.2.3.4"` + "\n"
	config += `	  description = "neighbor1"` + "\n"
	config += `	  shutdown = false` + "\n"
	config += `	  remote_as = 200` + "\n"
	config += `	  local_as = 200` + "\n"
	config += `	  keepalive_time = 40` + "\n"
	config += `	  hold_time = 200` + "\n"
	config += `	  update_source_interface = "GigabitEthernet0"` + "\n"
	config += `	  next_hop_self = false` + "\n"
	config += `	  send_community = true` + "\n"
	config += `	  send_extended_community = true` + "\n"
	config += `	  ebgp_multihop = 1` + "\n"
	config += `	  password = "myPassword"` + "\n"
	config += `	  send_label = true` + "\n"
	config += `	  explicit_null = false` + "\n"
	config += `	  as_override = false` + "\n"
	config += `	  allowas_in_number = 1` + "\n"
	config += `	  address_families = [{` + "\n"
	config += `		family_type = "ipv4-unicast"` + "\n"
	config += `		policy_type = "restart"` + "\n"
	config += `		restart_max_number_of_prefixes = 2000` + "\n"
	config += `		restart_threshold = 75` + "\n"
	config += `		restart_interval = 30` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_neighbors = [{` + "\n"
	config += `	  address = "2001::1"` + "\n"
	config += `	  description = "neighbor2"` + "\n"
	config += `	  shutdown = false` + "\n"
	config += `	  remote_as = 200` + "\n"
	config += `	  local_as = 200` + "\n"
	config += `	  keepalive_time = 180` + "\n"
	config += `	  hold_time = 60` + "\n"
	config += `	  update_source_interface = "Loopback1"` + "\n"
	config += `	  next_hop_self = true` + "\n"
	config += `	  send_community = true` + "\n"
	config += `	  send_extended_community = true` + "\n"
	config += `	  ebgp_multihop = 3` + "\n"
	config += `	  password = "myPassword"` + "\n"
	config += `	  as_override = true` + "\n"
	config += `	  allowas_in_number = 3` + "\n"
	config += `	  address_families = [{` + "\n"
	config += `		family_type = "ipv6-unicast"` + "\n"
	config += `		max_number_of_prefixes = 2000` + "\n"
	config += `		threshold = 75` + "\n"
	config += `		policy_type = "restart"` + "\n"
	config += `		restart_interval = 30` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_aggregate_addresses = [{` + "\n"
	config += `	  network_address = "10.10.0.0"` + "\n"
	config += `	  subnet_mask = "255.255.0.0"` + "\n"
	config += `	  as_set_path = false` + "\n"
	config += `	  summary_only = false` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_networks = [{` + "\n"
	config += `	  network_address = "10.10.0.0"` + "\n"
	config += `	  subnet_mask = "255.255.0.0"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_eibgp_maximum_paths = 1` + "\n"
	config += `	ipv4_originate = false` + "\n"
	config += `	ipv4_table_map_filter = false` + "\n"
	config += `	ipv6_aggregate_addresses = [{` + "\n"
	config += `	  aggregate_prefix = "3001::1/128"` + "\n"
	config += `	  as_set_path = false` + "\n"
	config += `	  summary_only = false` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_networks = [{` + "\n"
	config += `	  network_prefix = "2001:0DB8:0000:000b::/64"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_eibgp_maximum_paths = 2` + "\n"
	config += `	ipv6_originate = true` + "\n"
	config += `	ipv6_table_map_filter = false` + "\n"
	config += `	mpls_interfaces = [{` + "\n"
	config += `	  interface_name = "GigabitEthernet1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_transport_routing_bgp_feature" "test" {
			id = sdwan_transport_routing_bgp_feature.test.id
			feature_profile_id = sdwan_transport_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
