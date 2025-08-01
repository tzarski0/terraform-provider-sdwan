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
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type TrafficDataPolicyDefinition struct {
	Id            types.String                           `tfsdk:"id"`
	Version       types.Int64                            `tfsdk:"version"`
	Type          types.String                           `tfsdk:"type"`
	Name          types.String                           `tfsdk:"name"`
	Description   types.String                           `tfsdk:"description"`
	DefaultAction types.String                           `tfsdk:"default_action"`
	Sequences     []TrafficDataPolicyDefinitionSequences `tfsdk:"sequences"`
}

type TrafficDataPolicyDefinitionSequences struct {
	Id            types.Int64                                         `tfsdk:"id"`
	Name          types.String                                        `tfsdk:"name"`
	Type          types.String                                        `tfsdk:"type"`
	IpType        types.String                                        `tfsdk:"ip_type"`
	BaseAction    types.String                                        `tfsdk:"base_action"`
	MatchEntries  []TrafficDataPolicyDefinitionSequencesMatchEntries  `tfsdk:"match_entries"`
	ActionEntries []TrafficDataPolicyDefinitionSequencesActionEntries `tfsdk:"action_entries"`
}

type TrafficDataPolicyDefinitionSequencesMatchEntries struct {
	Type                             types.String `tfsdk:"type"`
	ApplicationListId                types.String `tfsdk:"application_list_id"`
	ApplicationListVersion           types.Int64  `tfsdk:"application_list_version"`
	DnsApplicationListId             types.String `tfsdk:"dns_application_list_id"`
	DnsApplicationListVersion        types.Int64  `tfsdk:"dns_application_list_version"`
	IcmpMessage                      types.String `tfsdk:"icmp_message"`
	Dns                              types.String `tfsdk:"dns"`
	Dscp                             types.String `tfsdk:"dscp"`
	PacketLength                     types.Int64  `tfsdk:"packet_length"`
	Plp                              types.String `tfsdk:"plp"`
	Protocol                         types.String `tfsdk:"protocol"`
	SourceDataPrefixListId           types.String `tfsdk:"source_data_prefix_list_id"`
	SourceDataPrefixListVersion      types.Int64  `tfsdk:"source_data_prefix_list_version"`
	SourceIp                         types.String `tfsdk:"source_ip"`
	SourcePort                       types.String `tfsdk:"source_port"`
	DestinationDataPrefixListId      types.String `tfsdk:"destination_data_prefix_list_id"`
	DestinationDataPrefixListVersion types.Int64  `tfsdk:"destination_data_prefix_list_version"`
	DestinationIp                    types.String `tfsdk:"destination_ip"`
	DestinationPort                  types.String `tfsdk:"destination_port"`
	DestinationRegion                types.String `tfsdk:"destination_region"`
	Tcp                              types.String `tfsdk:"tcp"`
	TrafficTo                        types.String `tfsdk:"traffic_to"`
}
type TrafficDataPolicyDefinitionSequencesActionEntries struct {
	Type                            types.String                                                     `tfsdk:"type"`
	Cflowd                          types.Bool                                                       `tfsdk:"cflowd"`
	Counter                         types.String                                                     `tfsdk:"counter"`
	DreOptimization                 types.Bool                                                       `tfsdk:"dre_optimization"`
	FallbackToRouting               types.Bool                                                       `tfsdk:"fallback_to_routing"`
	Log                             types.Bool                                                       `tfsdk:"log"`
	LossCorrection                  types.String                                                     `tfsdk:"loss_correction"`
	LossCorrectionFec               types.String                                                     `tfsdk:"loss_correction_fec"`
	LossCorrectionPacketDuplication types.String                                                     `tfsdk:"loss_correction_packet_duplication"`
	LossCorrectionFecThreshold      types.String                                                     `tfsdk:"loss_correction_fec_threshold"`
	NatPool                         types.String                                                     `tfsdk:"nat_pool"`
	NatPoolId                       types.Int64                                                      `tfsdk:"nat_pool_id"`
	RedirectDns                     types.String                                                     `tfsdk:"redirect_dns"`
	RedirectDnsType                 types.String                                                     `tfsdk:"redirect_dns_type"`
	RedirectDnsAddress              types.String                                                     `tfsdk:"redirect_dns_address"`
	ServiceNodeGroup                types.String                                                     `tfsdk:"service_node_group"`
	SecureInternetGateway           types.Bool                                                       `tfsdk:"secure_internet_gateway"`
	TcpOptimization                 types.Bool                                                       `tfsdk:"tcp_optimization"`
	SetParameters                   []TrafficDataPolicyDefinitionSequencesActionEntriesSetParameters `tfsdk:"set_parameters"`
	NatParameters                   []TrafficDataPolicyDefinitionSequencesActionEntriesNatParameters `tfsdk:"nat_parameters"`
}

type TrafficDataPolicyDefinitionSequencesActionEntriesSetParameters struct {
	Type                           types.String `tfsdk:"type"`
	Dscp                           types.Int64  `tfsdk:"dscp"`
	ForwardingClass                types.String `tfsdk:"forwarding_class"`
	NextHop                        types.String `tfsdk:"next_hop"`
	LocalTlocListColor             types.String `tfsdk:"local_tloc_list_color"`
	LocalTlocListEncap             types.String `tfsdk:"local_tloc_list_encap"`
	LocalTlocListRestrict          types.Bool   `tfsdk:"local_tloc_list_restrict"`
	NextHopLoose                   types.Bool   `tfsdk:"next_hop_loose"`
	PolicerListId                  types.String `tfsdk:"policer_list_id"`
	PolicerListVersion             types.Int64  `tfsdk:"policer_list_version"`
	PreferredColorGroupList        types.String `tfsdk:"preferred_color_group_list"`
	PreferredColorGroupListVersion types.Int64  `tfsdk:"preferred_color_group_list_version"`
	TlocListId                     types.String `tfsdk:"tloc_list_id"`
	TlocListVersion                types.Int64  `tfsdk:"tloc_list_version"`
	TlocIp                         types.String `tfsdk:"tloc_ip"`
	TlocColor                      types.String `tfsdk:"tloc_color"`
	TlocEncapsulation              types.String `tfsdk:"tloc_encapsulation"`
	ServiceType                    types.String `tfsdk:"service_type"`
	ServiceVpnId                   types.Int64  `tfsdk:"service_vpn_id"`
	ServiceTlocListId              types.String `tfsdk:"service_tloc_list_id"`
	ServiceTlocListVersion         types.Int64  `tfsdk:"service_tloc_list_version"`
	ServiceTlocIp                  types.String `tfsdk:"service_tloc_ip"`
	ServiceTlocLocal               types.Bool   `tfsdk:"service_tloc_local"`
	ServiceTlocRestrict            types.Bool   `tfsdk:"service_tloc_restrict"`
	ServiceTlocColor               types.String `tfsdk:"service_tloc_color"`
	ServiceTlocEncapsulation       types.String `tfsdk:"service_tloc_encapsulation"`
	VpnId                          types.Int64  `tfsdk:"vpn_id"`
}
type TrafficDataPolicyDefinitionSequencesActionEntriesNatParameters struct {
	Type     types.String `tfsdk:"type"`
	VpnId    types.Int64  `tfsdk:"vpn_id"`
	Fallback types.Bool   `tfsdk:"fallback"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TrafficDataPolicyDefinition) getPath() string {
	return "/template/policy/definition/data/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TrafficDataPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "data")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.DefaultAction.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.type", data.DefaultAction.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "sequences", []interface{}{})
		for _, item := range data.Sequences {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceId", item.Id.ValueInt64())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceName", item.Name.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceType", item.Type.ValueString())
			}
			if !item.IpType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceIpType", item.IpType.ValueString())
			}
			if !item.BaseAction.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "baseAction", item.BaseAction.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "match.entries", []interface{}{})
				for _, childItem := range item.MatchEntries {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "field", childItem.Type.ValueString())
					}
					if !childItem.ApplicationListId.IsNull() && childItem.Type.ValueString() == "appList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.ApplicationListId.ValueString())
					}
					if !childItem.DnsApplicationListId.IsNull() && childItem.Type.ValueString() == "dnsAppList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.DnsApplicationListId.ValueString())
					}
					if !childItem.IcmpMessage.IsNull() && childItem.Type.ValueString() == "icmpMessage" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.IcmpMessage.ValueString())
					}
					if !childItem.Dns.IsNull() && childItem.Type.ValueString() == "dns" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Dns.ValueString())
					}
					if !childItem.Dscp.IsNull() && childItem.Type.ValueString() == "dscp" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.Dscp.ValueString()))
					}
					if !childItem.PacketLength.IsNull() && childItem.Type.ValueString() == "packetLength" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.PacketLength.ValueInt64()))
					}
					if !childItem.Plp.IsNull() && childItem.Type.ValueString() == "plp" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Plp.ValueString())
					}
					if !childItem.Protocol.IsNull() && childItem.Type.ValueString() == "protocol" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Protocol.ValueString())
					}
					if !childItem.SourceDataPrefixListId.IsNull() && childItem.Type.ValueString() == "sourceDataPrefixList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.SourceDataPrefixListId.ValueString())
					}
					if !childItem.SourceIp.IsNull() && childItem.Type.ValueString() == "sourceIp" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.SourceIp.ValueString())
					}
					if !childItem.SourcePort.IsNull() && childItem.Type.ValueString() == "sourcePort" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.SourcePort.ValueString())
					}
					if !childItem.DestinationDataPrefixListId.IsNull() && childItem.Type.ValueString() == "destinationDataPrefixList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.DestinationDataPrefixListId.ValueString())
					}
					if !childItem.DestinationIp.IsNull() && childItem.Type.ValueString() == "destinationIp" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.DestinationIp.ValueString())
					}
					if !childItem.DestinationPort.IsNull() && childItem.Type.ValueString() == "destinationPort" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.DestinationPort.ValueString())
					}
					if !childItem.DestinationRegion.IsNull() && childItem.Type.ValueString() == "destinationRegion" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.DestinationRegion.ValueString())
					}
					if !childItem.Tcp.IsNull() && childItem.Type.ValueString() == "tcp" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Tcp.ValueString())
					}
					if !childItem.TrafficTo.IsNull() && childItem.Type.ValueString() == "trafficTo" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.TrafficTo.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "match.entries.-1", itemChildBody)
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "actions", []interface{}{})
				for _, childItem := range item.ActionEntries {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Cflowd.IsNull() && childItem.Type.ValueString() == "cflowd" {
						if true && childItem.Cflowd.ValueBool() {
							itemChildBody, _ = sjson.Set(itemChildBody, "parameter", "")
						}
					}
					if !childItem.Counter.IsNull() && childItem.Type.ValueString() == "count" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter", childItem.Counter.ValueString())
					}
					if !childItem.DreOptimization.IsNull() && childItem.Type.ValueString() == "dreOptimization" {
						if true && childItem.DreOptimization.ValueBool() {
							itemChildBody, _ = sjson.Set(itemChildBody, "parameter", "")
						}
					}
					if !childItem.FallbackToRouting.IsNull() && childItem.Type.ValueString() == "fallbackToRouting" {
						if true && childItem.FallbackToRouting.ValueBool() {
							itemChildBody, _ = sjson.Set(itemChildBody, "parameter", "")
						}
					}
					if !childItem.Log.IsNull() && childItem.Type.ValueString() == "log" {
						if true && childItem.Log.ValueBool() {
							itemChildBody, _ = sjson.Set(itemChildBody, "parameter", "")
						}
					}
					if !childItem.LossCorrection.IsNull() && childItem.Type.ValueString() == "lossProtect" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter", childItem.LossCorrection.ValueString())
					}
					if !childItem.LossCorrectionFec.IsNull() && childItem.Type.ValueString() == "lossProtectFec" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter", childItem.LossCorrectionFec.ValueString())
					}
					if !childItem.LossCorrectionPacketDuplication.IsNull() && childItem.Type.ValueString() == "lossProtectPktDup" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter", childItem.LossCorrectionPacketDuplication.ValueString())
					}
					if !childItem.LossCorrectionFecThreshold.IsNull() && childItem.Type.ValueString() == "lossProtectFec" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.LossCorrectionFecThreshold.ValueString())
					}
					if !childItem.NatPool.IsNull() && childItem.Type.ValueString() == "nat" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter.field", childItem.NatPool.ValueString())
					}
					if !childItem.NatPoolId.IsNull() && childItem.Type.ValueString() == "nat" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter.value", childItem.NatPoolId.ValueInt64())
					}
					if !childItem.RedirectDns.IsNull() && childItem.Type.ValueString() == "redirectDns" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter.field", childItem.RedirectDns.ValueString())
					}
					if !childItem.RedirectDnsType.IsNull() && childItem.RedirectDns.ValueString() == "dnsType" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter.value", childItem.RedirectDnsType.ValueString())
					}
					if !childItem.RedirectDnsAddress.IsNull() && childItem.RedirectDns.ValueString() == "ipAddress" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter.value", childItem.RedirectDnsAddress.ValueString())
					}
					if !childItem.ServiceNodeGroup.IsNull() && childItem.Type.ValueString() == "serviceNodeGroup" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter", childItem.ServiceNodeGroup.ValueString())
					}
					if !childItem.SecureInternetGateway.IsNull() && childItem.Type.ValueString() == "sig" {
						if true && childItem.SecureInternetGateway.ValueBool() {
							itemChildBody, _ = sjson.Set(itemChildBody, "parameter", "")
						}
					}
					if !childItem.TcpOptimization.IsNull() && childItem.Type.ValueString() == "tcpOptimization" {
						if true && childItem.TcpOptimization.ValueBool() {
							itemChildBody, _ = sjson.Set(itemChildBody, "parameter", "")
						}
					}
					if true && childItem.Type.ValueString() == "set" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter", []interface{}{})
						for _, childChildItem := range childItem.SetParameters {
							itemChildChildBody := ""
							if !childChildItem.Type.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "field", childChildItem.Type.ValueString())
							}
							if !childChildItem.Dscp.IsNull() && childChildItem.Type.ValueString() == "dscp" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", fmt.Sprint(childChildItem.Dscp.ValueInt64()))
							}
							if !childChildItem.ForwardingClass.IsNull() && childChildItem.Type.ValueString() == "forwardingClass" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", childChildItem.ForwardingClass.ValueString())
							}
							if !childChildItem.NextHop.IsNull() && childChildItem.Type.ValueString() == "nextHop" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", childChildItem.NextHop.ValueString())
							}
							if !childChildItem.LocalTlocListColor.IsNull() && childChildItem.Type.ValueString() == "localTlocList" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.color", childChildItem.LocalTlocListColor.ValueString())
							}
							if !childChildItem.LocalTlocListEncap.IsNull() && childChildItem.Type.ValueString() == "localTlocList" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.encap", childChildItem.LocalTlocListEncap.ValueString())
							}
							if !childChildItem.LocalTlocListRestrict.IsNull() && childChildItem.Type.ValueString() == "localTlocList" {
								if true && childChildItem.LocalTlocListRestrict.ValueBool() {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.restrict", "")
								}
							}
							if !childChildItem.NextHopLoose.IsNull() && childChildItem.Type.ValueString() == "nextHopLoose" {
								if false && childChildItem.NextHopLoose.ValueBool() {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", "")
								} else {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", fmt.Sprint(childChildItem.NextHopLoose.ValueBool()))
								}
							}
							if !childChildItem.PolicerListId.IsNull() && childChildItem.Type.ValueString() == "policer" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "ref", childChildItem.PolicerListId.ValueString())
							}
							if !childChildItem.PreferredColorGroupList.IsNull() && childChildItem.Type.ValueString() == "preferredColorGroup" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "ref", childChildItem.PreferredColorGroupList.ValueString())
							}
							if !childChildItem.TlocListId.IsNull() && childChildItem.Type.ValueString() == "tlocList" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "ref", childChildItem.TlocListId.ValueString())
							}
							if !childChildItem.TlocIp.IsNull() && childChildItem.Type.ValueString() == "tloc" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.ip", childChildItem.TlocIp.ValueString())
							}
							if !childChildItem.TlocColor.IsNull() && childChildItem.Type.ValueString() == "tloc" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.color", childChildItem.TlocColor.ValueString())
							}
							if !childChildItem.TlocEncapsulation.IsNull() && childChildItem.Type.ValueString() == "tloc" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.encap", childChildItem.TlocEncapsulation.ValueString())
							}
							if !childChildItem.ServiceType.IsNull() && childChildItem.Type.ValueString() == "service" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.type", childChildItem.ServiceType.ValueString())
							}
							if !childChildItem.ServiceVpnId.IsNull() && childChildItem.Type.ValueString() == "service" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.vpn", fmt.Sprint(childChildItem.ServiceVpnId.ValueInt64()))
							}
							if !childChildItem.ServiceTlocListId.IsNull() && childChildItem.Type.ValueString() == "service" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.tlocList.ref", childChildItem.ServiceTlocListId.ValueString())
							}
							if !childChildItem.ServiceTlocIp.IsNull() && childChildItem.Type.ValueString() == "service" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.tloc.ip", childChildItem.ServiceTlocIp.ValueString())
							}
							if !childChildItem.ServiceTlocLocal.IsNull() && childChildItem.Type.ValueString() == "service" {
								if true && childChildItem.ServiceTlocLocal.ValueBool() {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.local", "")
								}
							}
							if !childChildItem.ServiceTlocRestrict.IsNull() && childChildItem.Type.ValueString() == "service" {
								if true && childChildItem.ServiceTlocRestrict.ValueBool() {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.restrict", "")
								}
							}
							if !childChildItem.ServiceTlocColor.IsNull() && childChildItem.Type.ValueString() == "service" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.tloc.color", childChildItem.ServiceTlocColor.ValueString())
							}
							if !childChildItem.ServiceTlocEncapsulation.IsNull() && childChildItem.Type.ValueString() == "service" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.tloc.encap", childChildItem.ServiceTlocEncapsulation.ValueString())
							}
							if !childChildItem.VpnId.IsNull() && childChildItem.Type.ValueString() == "vpn" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", fmt.Sprint(childChildItem.VpnId.ValueInt64()))
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "parameter.-1", itemChildChildBody)
						}
					}
					if true && childItem.Type.ValueString() == "nat" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter", []interface{}{})
						for _, childChildItem := range childItem.NatParameters {
							itemChildChildBody := ""
							if !childChildItem.Type.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "field", childChildItem.Type.ValueString())
							}
							if !childChildItem.VpnId.IsNull() && childChildItem.Type.ValueString() == "useVpn" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", fmt.Sprint(childChildItem.VpnId.ValueInt64()))
							}
							if !childChildItem.Fallback.IsNull() && childChildItem.Type.ValueString() == "fallback" {
								if false && childChildItem.Fallback.ValueBool() {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", "")
								} else {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", fmt.Sprint(childChildItem.Fallback.ValueBool()))
								}
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "parameter.-1", itemChildChildBody)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "actions.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "sequences.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TrafficDataPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
	state := *data
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("defaultAction.type"); value.Exists() {
		data.DefaultAction = types.StringValue(value.String())
	} else {
		data.DefaultAction = types.StringNull()
	}
	if value := res.Get("sequences"); value.Exists() && len(value.Array()) > 0 {
		data.Sequences = make([]TrafficDataPolicyDefinitionSequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TrafficDataPolicyDefinitionSequences{}
			if cValue := v.Get("sequenceId"); cValue.Exists() {
				item.Id = types.Int64Value(cValue.Int())
			} else {
				item.Id = types.Int64Null()
			}
			if cValue := v.Get("sequenceName"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("sequenceType"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			} else {
				item.Type = types.StringNull()
			}
			if cValue := v.Get("sequenceIpType"); cValue.Exists() {
				item.IpType = types.StringValue(cValue.String())
			} else {
				item.IpType = types.StringNull()
			}
			if cValue := v.Get("baseAction"); cValue.Exists() {
				item.BaseAction = types.StringValue(cValue.String())
			} else {
				item.BaseAction = types.StringNull()
			}
			if cValue := v.Get("match.entries"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.MatchEntries = make([]TrafficDataPolicyDefinitionSequencesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TrafficDataPolicyDefinitionSequencesMatchEntries{}
					if ccValue := cv.Get("field"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "appList" {
						cItem.ApplicationListId = types.StringValue(ccValue.String())
					} else {
						cItem.ApplicationListId = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "dnsAppList" {
						cItem.DnsApplicationListId = types.StringValue(ccValue.String())
					} else {
						cItem.DnsApplicationListId = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "icmpMessage" {
						cItem.IcmpMessage = types.StringValue(ccValue.String())
					} else {
						cItem.IcmpMessage = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "dns" {
						cItem.Dns = types.StringValue(ccValue.String())
					} else {
						cItem.Dns = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "dscp" {
						cItem.Dscp = types.StringValue(ccValue.String())
					} else {
						cItem.Dscp = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "packetLength" {
						cItem.PacketLength = types.Int64Value(ccValue.Int())
					} else {
						cItem.PacketLength = types.Int64Null()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "plp" {
						cItem.Plp = types.StringValue(ccValue.String())
					} else {
						cItem.Plp = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "protocol" {
						cItem.Protocol = types.StringValue(ccValue.String())
					} else {
						cItem.Protocol = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "sourceDataPrefixList" {
						cItem.SourceDataPrefixListId = types.StringValue(ccValue.String())
					} else {
						cItem.SourceDataPrefixListId = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "sourceIp" {
						cItem.SourceIp = types.StringValue(ccValue.String())
					} else {
						cItem.SourceIp = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "sourcePort" {
						cItem.SourcePort = types.StringValue(ccValue.String())
					} else {
						cItem.SourcePort = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "destinationDataPrefixList" {
						cItem.DestinationDataPrefixListId = types.StringValue(ccValue.String())
					} else {
						cItem.DestinationDataPrefixListId = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "destinationIp" {
						cItem.DestinationIp = types.StringValue(ccValue.String())
					} else {
						cItem.DestinationIp = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "destinationPort" {
						cItem.DestinationPort = types.StringValue(ccValue.String())
					} else {
						cItem.DestinationPort = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "destinationRegion" {
						cItem.DestinationRegion = types.StringValue(ccValue.String())
					} else {
						cItem.DestinationRegion = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "tcp" {
						cItem.Tcp = types.StringValue(ccValue.String())
					} else {
						cItem.Tcp = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "trafficTo" {
						cItem.TrafficTo = types.StringValue(ccValue.String())
					} else {
						cItem.TrafficTo = types.StringNull()
					}
					item.MatchEntries = append(item.MatchEntries, cItem)
					return true
				})
			} else {
				if len(item.MatchEntries) > 0 {
					item.MatchEntries = []TrafficDataPolicyDefinitionSequencesMatchEntries{}
				}
			}
			if cValue := v.Get("actions"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ActionEntries = make([]TrafficDataPolicyDefinitionSequencesActionEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TrafficDataPolicyDefinitionSequencesActionEntries{}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "cflowd" {
						if true && ccValue.String() == "" {
							cItem.Cflowd = types.BoolValue(true)
						}
					} else {
						cItem.Cflowd = types.BoolNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "count" {
						cItem.Counter = types.StringValue(ccValue.String())
					} else {
						cItem.Counter = types.StringNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "dreOptimization" {
						if true && ccValue.String() == "" {
							cItem.DreOptimization = types.BoolValue(true)
						}
					} else {
						cItem.DreOptimization = types.BoolNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "fallbackToRouting" {
						if true && ccValue.String() == "" {
							cItem.FallbackToRouting = types.BoolValue(true)
						}
					} else {
						cItem.FallbackToRouting = types.BoolNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "log" {
						if true && ccValue.String() == "" {
							cItem.Log = types.BoolValue(true)
						}
					} else {
						cItem.Log = types.BoolNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "lossProtect" {
						cItem.LossCorrection = types.StringValue(ccValue.String())
					} else {
						cItem.LossCorrection = types.StringNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "lossProtectFec" {
						cItem.LossCorrectionFec = types.StringValue(ccValue.String())
					} else {
						cItem.LossCorrectionFec = types.StringNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "lossProtectPktDup" {
						cItem.LossCorrectionPacketDuplication = types.StringValue(ccValue.String())
					} else {
						cItem.LossCorrectionPacketDuplication = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "lossProtectFec" {
						cItem.LossCorrectionFecThreshold = types.StringValue(ccValue.String())
					} else {
						cItem.LossCorrectionFecThreshold = types.StringNull()
					}
					if ccValue := cv.Get("parameter.field"); ccValue.Exists() && cItem.Type.ValueString() == "nat" {
						cItem.NatPool = types.StringValue(ccValue.String())
					} else {
						cItem.NatPool = types.StringNull()
					}
					if ccValue := cv.Get("parameter.value"); ccValue.Exists() && cItem.Type.ValueString() == "nat" {
						cItem.NatPoolId = types.Int64Value(ccValue.Int())
					} else {
						cItem.NatPoolId = types.Int64Null()
					}
					if ccValue := cv.Get("parameter.field"); ccValue.Exists() && cItem.Type.ValueString() == "redirectDns" {
						cItem.RedirectDns = types.StringValue(ccValue.String())
					} else {
						cItem.RedirectDns = types.StringNull()
					}
					if ccValue := cv.Get("parameter.value"); ccValue.Exists() && cItem.RedirectDns.ValueString() == "dnsType" {
						cItem.RedirectDnsType = types.StringValue(ccValue.String())
					} else {
						cItem.RedirectDnsType = types.StringNull()
					}
					if ccValue := cv.Get("parameter.value"); ccValue.Exists() && cItem.RedirectDns.ValueString() == "ipAddress" {
						cItem.RedirectDnsAddress = types.StringValue(ccValue.String())
					} else {
						cItem.RedirectDnsAddress = types.StringNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "serviceNodeGroup" {
						cItem.ServiceNodeGroup = types.StringValue(ccValue.String())
					} else {
						cItem.ServiceNodeGroup = types.StringNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "sig" {
						if true && ccValue.String() == "" {
							cItem.SecureInternetGateway = types.BoolValue(true)
						}
					} else {
						cItem.SecureInternetGateway = types.BoolNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "tcpOptimization" {
						if true && ccValue.String() == "" {
							cItem.TcpOptimization = types.BoolValue(true)
						}
					} else {
						cItem.TcpOptimization = types.BoolNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && len(ccValue.Array()) > 0 && cItem.Type.ValueString() == "set" {
						cItem.SetParameters = make([]TrafficDataPolicyDefinitionSequencesActionEntriesSetParameters, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := TrafficDataPolicyDefinitionSequencesActionEntriesSetParameters{}
							if cccValue := ccv.Get("field"); cccValue.Exists() {
								ccItem.Type = types.StringValue(cccValue.String())
							} else {
								ccItem.Type = types.StringNull()
							}
							if cccValue := ccv.Get("value"); cccValue.Exists() && ccItem.Type.ValueString() == "dscp" {
								ccItem.Dscp = types.Int64Value(cccValue.Int())
							} else {
								ccItem.Dscp = types.Int64Null()
							}
							if cccValue := ccv.Get("value"); cccValue.Exists() && ccItem.Type.ValueString() == "forwardingClass" {
								ccItem.ForwardingClass = types.StringValue(cccValue.String())
							} else {
								ccItem.ForwardingClass = types.StringNull()
							}
							if cccValue := ccv.Get("value"); cccValue.Exists() && ccItem.Type.ValueString() == "nextHop" {
								ccItem.NextHop = types.StringValue(cccValue.String())
							} else {
								ccItem.NextHop = types.StringNull()
							}
							if cccValue := ccv.Get("value.color"); cccValue.Exists() && ccItem.Type.ValueString() == "localTlocList" {
								ccItem.LocalTlocListColor = types.StringValue(cccValue.String())
							} else {
								ccItem.LocalTlocListColor = types.StringNull()
							}
							if cccValue := ccv.Get("value.encap"); cccValue.Exists() && ccItem.Type.ValueString() == "localTlocList" {
								ccItem.LocalTlocListEncap = types.StringValue(cccValue.String())
							} else {
								ccItem.LocalTlocListEncap = types.StringNull()
							}
							if cccValue := ccv.Get("value.restrict"); cccValue.Exists() && ccItem.Type.ValueString() == "localTlocList" {
								if true && cccValue.String() == "" {
									ccItem.LocalTlocListRestrict = types.BoolValue(true)
								}
							} else {
								ccItem.LocalTlocListRestrict = types.BoolNull()
							}
							if cccValue := ccv.Get("value"); cccValue.Exists() && ccItem.Type.ValueString() == "nextHopLoose" {
								if false && cccValue.String() == "" {
									ccItem.NextHopLoose = types.BoolValue(true)
								} else {
									ccItem.NextHopLoose = types.BoolValue(cccValue.Bool())
								}
							} else {
								ccItem.NextHopLoose = types.BoolNull()
							}
							if cccValue := ccv.Get("ref"); cccValue.Exists() && ccItem.Type.ValueString() == "policer" {
								ccItem.PolicerListId = types.StringValue(cccValue.String())
							} else {
								ccItem.PolicerListId = types.StringNull()
							}
							if cccValue := ccv.Get("ref"); cccValue.Exists() && ccItem.Type.ValueString() == "preferredColorGroup" {
								ccItem.PreferredColorGroupList = types.StringValue(cccValue.String())
							} else {
								ccItem.PreferredColorGroupList = types.StringNull()
							}
							if cccValue := ccv.Get("ref"); cccValue.Exists() && ccItem.Type.ValueString() == "tlocList" {
								ccItem.TlocListId = types.StringValue(cccValue.String())
							} else {
								ccItem.TlocListId = types.StringNull()
							}
							if cccValue := ccv.Get("value.ip"); cccValue.Exists() && ccItem.Type.ValueString() == "tloc" {
								ccItem.TlocIp = types.StringValue(cccValue.String())
							} else {
								ccItem.TlocIp = types.StringNull()
							}
							if cccValue := ccv.Get("value.color"); cccValue.Exists() && ccItem.Type.ValueString() == "tloc" {
								ccItem.TlocColor = types.StringValue(cccValue.String())
							} else {
								ccItem.TlocColor = types.StringNull()
							}
							if cccValue := ccv.Get("value.encap"); cccValue.Exists() && ccItem.Type.ValueString() == "tloc" {
								ccItem.TlocEncapsulation = types.StringValue(cccValue.String())
							} else {
								ccItem.TlocEncapsulation = types.StringNull()
							}
							if cccValue := ccv.Get("value.type"); cccValue.Exists() && ccItem.Type.ValueString() == "service" {
								ccItem.ServiceType = types.StringValue(cccValue.String())
							} else {
								ccItem.ServiceType = types.StringNull()
							}
							if cccValue := ccv.Get("value.vpn"); cccValue.Exists() && ccItem.Type.ValueString() == "service" {
								ccItem.ServiceVpnId = types.Int64Value(cccValue.Int())
							} else {
								ccItem.ServiceVpnId = types.Int64Null()
							}
							if cccValue := ccv.Get("value.tlocList.ref"); cccValue.Exists() && ccItem.Type.ValueString() == "service" {
								ccItem.ServiceTlocListId = types.StringValue(cccValue.String())
							} else {
								ccItem.ServiceTlocListId = types.StringNull()
							}
							if cccValue := ccv.Get("value.tloc.ip"); cccValue.Exists() && ccItem.Type.ValueString() == "service" {
								ccItem.ServiceTlocIp = types.StringValue(cccValue.String())
							} else {
								ccItem.ServiceTlocIp = types.StringNull()
							}
							if cccValue := ccv.Get("value.local"); cccValue.Exists() && ccItem.Type.ValueString() == "service" {
								if true && cccValue.String() == "" {
									ccItem.ServiceTlocLocal = types.BoolValue(true)
								}
							} else {
								ccItem.ServiceTlocLocal = types.BoolNull()
							}
							if cccValue := ccv.Get("value.restrict"); cccValue.Exists() && ccItem.Type.ValueString() == "service" {
								if true && cccValue.String() == "" {
									ccItem.ServiceTlocRestrict = types.BoolValue(true)
								}
							} else {
								ccItem.ServiceTlocRestrict = types.BoolNull()
							}
							if cccValue := ccv.Get("value.tloc.color"); cccValue.Exists() && ccItem.Type.ValueString() == "service" {
								ccItem.ServiceTlocColor = types.StringValue(cccValue.String())
							} else {
								ccItem.ServiceTlocColor = types.StringNull()
							}
							if cccValue := ccv.Get("value.tloc.encap"); cccValue.Exists() && ccItem.Type.ValueString() == "service" {
								ccItem.ServiceTlocEncapsulation = types.StringValue(cccValue.String())
							} else {
								ccItem.ServiceTlocEncapsulation = types.StringNull()
							}
							if cccValue := ccv.Get("value"); cccValue.Exists() && ccItem.Type.ValueString() == "vpn" {
								ccItem.VpnId = types.Int64Value(cccValue.Int())
							} else {
								ccItem.VpnId = types.Int64Null()
							}
							cItem.SetParameters = append(cItem.SetParameters, ccItem)
							return true
						})
					} else {
						if len(cItem.SetParameters) > 0 {
							cItem.SetParameters = []TrafficDataPolicyDefinitionSequencesActionEntriesSetParameters{}
						}
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && len(ccValue.Array()) > 0 && cItem.Type.ValueString() == "nat" {
						cItem.NatParameters = make([]TrafficDataPolicyDefinitionSequencesActionEntriesNatParameters, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := TrafficDataPolicyDefinitionSequencesActionEntriesNatParameters{}
							if cccValue := ccv.Get("field"); cccValue.Exists() {
								ccItem.Type = types.StringValue(cccValue.String())
							} else {
								ccItem.Type = types.StringNull()
							}
							if cccValue := ccv.Get("value"); cccValue.Exists() && ccItem.Type.ValueString() == "useVpn" {
								ccItem.VpnId = types.Int64Value(cccValue.Int())
							} else {
								ccItem.VpnId = types.Int64Null()
							}
							if cccValue := ccv.Get("value"); cccValue.Exists() && ccItem.Type.ValueString() == "fallback" {
								if false && cccValue.String() == "" {
									ccItem.Fallback = types.BoolValue(true)
								} else {
									ccItem.Fallback = types.BoolValue(cccValue.Bool())
								}
							} else {
								ccItem.Fallback = types.BoolNull()
							}
							cItem.NatParameters = append(cItem.NatParameters, ccItem)
							return true
						})
					} else {
						if len(cItem.NatParameters) > 0 {
							cItem.NatParameters = []TrafficDataPolicyDefinitionSequencesActionEntriesNatParameters{}
						}
					}
					item.ActionEntries = append(item.ActionEntries, cItem)
					return true
				})
			} else {
				if len(item.ActionEntries) > 0 {
					item.ActionEntries = []TrafficDataPolicyDefinitionSequencesActionEntries{}
				}
			}
			data.Sequences = append(data.Sequences, item)
			return true
		})
	} else {
		if len(data.Sequences) > 0 {
			data.Sequences = []TrafficDataPolicyDefinitionSequences{}
		}
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *TrafficDataPolicyDefinition) hasChanges(ctx context.Context, state *TrafficDataPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.DefaultAction.Equal(state.DefaultAction) {
		hasChanges = true
	}
	if len(data.Sequences) != len(state.Sequences) {
		hasChanges = true
	} else {
		for i := range data.Sequences {
			if !data.Sequences[i].Id.Equal(state.Sequences[i].Id) {
				hasChanges = true
			}
			if !data.Sequences[i].Name.Equal(state.Sequences[i].Name) {
				hasChanges = true
			}
			if !data.Sequences[i].Type.Equal(state.Sequences[i].Type) {
				hasChanges = true
			}
			if !data.Sequences[i].IpType.Equal(state.Sequences[i].IpType) {
				hasChanges = true
			}
			if !data.Sequences[i].BaseAction.Equal(state.Sequences[i].BaseAction) {
				hasChanges = true
			}
			if len(data.Sequences[i].MatchEntries) != len(state.Sequences[i].MatchEntries) {
				hasChanges = true
			} else {
				for ii := range data.Sequences[i].MatchEntries {
					if !data.Sequences[i].MatchEntries[ii].Type.Equal(state.Sequences[i].MatchEntries[ii].Type) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].ApplicationListId.Equal(state.Sequences[i].MatchEntries[ii].ApplicationListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DnsApplicationListId.Equal(state.Sequences[i].MatchEntries[ii].DnsApplicationListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].IcmpMessage.Equal(state.Sequences[i].MatchEntries[ii].IcmpMessage) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Dns.Equal(state.Sequences[i].MatchEntries[ii].Dns) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Dscp.Equal(state.Sequences[i].MatchEntries[ii].Dscp) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].PacketLength.Equal(state.Sequences[i].MatchEntries[ii].PacketLength) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Plp.Equal(state.Sequences[i].MatchEntries[ii].Plp) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Protocol.Equal(state.Sequences[i].MatchEntries[ii].Protocol) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].SourceDataPrefixListId.Equal(state.Sequences[i].MatchEntries[ii].SourceDataPrefixListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].SourceIp.Equal(state.Sequences[i].MatchEntries[ii].SourceIp) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].SourcePort.Equal(state.Sequences[i].MatchEntries[ii].SourcePort) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DestinationDataPrefixListId.Equal(state.Sequences[i].MatchEntries[ii].DestinationDataPrefixListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DestinationIp.Equal(state.Sequences[i].MatchEntries[ii].DestinationIp) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DestinationPort.Equal(state.Sequences[i].MatchEntries[ii].DestinationPort) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DestinationRegion.Equal(state.Sequences[i].MatchEntries[ii].DestinationRegion) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Tcp.Equal(state.Sequences[i].MatchEntries[ii].Tcp) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].TrafficTo.Equal(state.Sequences[i].MatchEntries[ii].TrafficTo) {
						hasChanges = true
					}
				}
			}
			if len(data.Sequences[i].ActionEntries) != len(state.Sequences[i].ActionEntries) {
				hasChanges = true
			} else {
				for ii := range data.Sequences[i].ActionEntries {
					if !data.Sequences[i].ActionEntries[ii].Type.Equal(state.Sequences[i].ActionEntries[ii].Type) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].Cflowd.Equal(state.Sequences[i].ActionEntries[ii].Cflowd) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].Counter.Equal(state.Sequences[i].ActionEntries[ii].Counter) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].DreOptimization.Equal(state.Sequences[i].ActionEntries[ii].DreOptimization) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].FallbackToRouting.Equal(state.Sequences[i].ActionEntries[ii].FallbackToRouting) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].Log.Equal(state.Sequences[i].ActionEntries[ii].Log) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].LossCorrection.Equal(state.Sequences[i].ActionEntries[ii].LossCorrection) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].LossCorrectionFec.Equal(state.Sequences[i].ActionEntries[ii].LossCorrectionFec) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].LossCorrectionPacketDuplication.Equal(state.Sequences[i].ActionEntries[ii].LossCorrectionPacketDuplication) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].LossCorrectionFecThreshold.Equal(state.Sequences[i].ActionEntries[ii].LossCorrectionFecThreshold) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].NatPool.Equal(state.Sequences[i].ActionEntries[ii].NatPool) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].NatPoolId.Equal(state.Sequences[i].ActionEntries[ii].NatPoolId) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].RedirectDns.Equal(state.Sequences[i].ActionEntries[ii].RedirectDns) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].RedirectDnsType.Equal(state.Sequences[i].ActionEntries[ii].RedirectDnsType) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].RedirectDnsAddress.Equal(state.Sequences[i].ActionEntries[ii].RedirectDnsAddress) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].ServiceNodeGroup.Equal(state.Sequences[i].ActionEntries[ii].ServiceNodeGroup) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].SecureInternetGateway.Equal(state.Sequences[i].ActionEntries[ii].SecureInternetGateway) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].TcpOptimization.Equal(state.Sequences[i].ActionEntries[ii].TcpOptimization) {
						hasChanges = true
					}
					if len(data.Sequences[i].ActionEntries[ii].SetParameters) != len(state.Sequences[i].ActionEntries[ii].SetParameters) {
						hasChanges = true
					} else {
						for iii := range data.Sequences[i].ActionEntries[ii].SetParameters {
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].Type.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].Type) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].Dscp.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].Dscp) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].ForwardingClass.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].ForwardingClass) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].NextHop.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].NextHop) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].LocalTlocListColor.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].LocalTlocListColor) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].LocalTlocListEncap.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].LocalTlocListEncap) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].LocalTlocListRestrict.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].LocalTlocListRestrict) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].NextHopLoose.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].NextHopLoose) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].PolicerListId.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].PolicerListId) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].PreferredColorGroupList.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].PreferredColorGroupList) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocListId.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocListId) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocIp.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocIp) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocColor.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocColor) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocEncapsulation.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocEncapsulation) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceType.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceType) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceVpnId.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceVpnId) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocListId.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocListId) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocIp.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocIp) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocLocal.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocLocal) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocRestrict.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocRestrict) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocColor.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocColor) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocEncapsulation.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocEncapsulation) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].VpnId.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].VpnId) {
								hasChanges = true
							}
						}
					}
					if len(data.Sequences[i].ActionEntries[ii].NatParameters) != len(state.Sequences[i].ActionEntries[ii].NatParameters) {
						hasChanges = true
					} else {
						for iii := range data.Sequences[i].ActionEntries[ii].NatParameters {
							if !data.Sequences[i].ActionEntries[ii].NatParameters[iii].Type.Equal(state.Sequences[i].ActionEntries[ii].NatParameters[iii].Type) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].NatParameters[iii].VpnId.Equal(state.Sequences[i].ActionEntries[ii].NatParameters[iii].VpnId) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].NatParameters[iii].Fallback.Equal(state.Sequences[i].ActionEntries[ii].NatParameters[iii].Fallback) {
								hasChanges = true
							}
						}
					}
				}
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *TrafficDataPolicyDefinition) updateVersions(ctx context.Context, state *TrafficDataPolicyDefinition) {
	for i := range data.Sequences {
		dataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].Id.ValueInt64()), fmt.Sprintf("%v", data.Sequences[i].Name.ValueString())}
		stateIndex := -1
		for j := range state.Sequences {
			stateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[j].Id.ValueInt64()), fmt.Sprintf("%v", state.Sequences[j].Name.ValueString())}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		for ii := range data.Sequences[i].MatchEntries {
			cDataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].MatchEntries[ii].Type.ValueString())}
			cStateIndex := -1
			if stateIndex > -1 {
				for jj := range state.Sequences[stateIndex].MatchEntries {
					cStateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[stateIndex].MatchEntries[jj].Type.ValueString())}
					if cDataKeys == cStateKeys {
						cStateIndex = jj
						break
					}
				}
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].ApplicationListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].ApplicationListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].ApplicationListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].DnsApplicationListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].DnsApplicationListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].DnsApplicationListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].SourceDataPrefixListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].SourceDataPrefixListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].SourceDataPrefixListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].DestinationDataPrefixListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].DestinationDataPrefixListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].DestinationDataPrefixListVersion = types.Int64Null()
			}
		}
		for ii := range data.Sequences[i].ActionEntries {
			cDataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].ActionEntries[ii].Type.ValueString())}
			cStateIndex := -1
			if stateIndex > -1 {
				for jj := range state.Sequences[stateIndex].ActionEntries {
					cStateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[stateIndex].ActionEntries[jj].Type.ValueString())}
					if cDataKeys == cStateKeys {
						cStateIndex = jj
						break
					}
				}
			}
			for iii := range data.Sequences[i].ActionEntries[ii].SetParameters {
				ccDataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].ActionEntries[ii].SetParameters[iii].Type.ValueString())}
				ccStateIndex := -1
				if cStateIndex > -1 {
					for jjj := range state.Sequences[stateIndex].ActionEntries[cStateIndex].SetParameters {
						ccStateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[stateIndex].ActionEntries[cStateIndex].SetParameters[jjj].Type.ValueString())}
						if ccDataKeys == ccStateKeys {
							ccStateIndex = jjj
							break
						}
					}
				}
				if ccStateIndex > -1 {
					data.Sequences[i].ActionEntries[ii].SetParameters[iii].PolicerListVersion = state.Sequences[stateIndex].ActionEntries[cStateIndex].SetParameters[ccStateIndex].PolicerListVersion
				} else {
					data.Sequences[i].ActionEntries[ii].SetParameters[iii].PolicerListVersion = types.Int64Null()
				}
				if ccStateIndex > -1 {
					data.Sequences[i].ActionEntries[ii].SetParameters[iii].PreferredColorGroupListVersion = state.Sequences[stateIndex].ActionEntries[cStateIndex].SetParameters[ccStateIndex].PreferredColorGroupListVersion
				} else {
					data.Sequences[i].ActionEntries[ii].SetParameters[iii].PreferredColorGroupListVersion = types.Int64Null()
				}
				if ccStateIndex > -1 {
					data.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocListVersion = state.Sequences[stateIndex].ActionEntries[cStateIndex].SetParameters[ccStateIndex].TlocListVersion
				} else {
					data.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocListVersion = types.Int64Null()
				}
				if ccStateIndex > -1 {
					data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocListVersion = state.Sequences[stateIndex].ActionEntries[cStateIndex].SetParameters[ccStateIndex].ServiceTlocListVersion
				} else {
					data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocListVersion = types.Int64Null()
				}
			}
		}
	}
}

// End of section. //template:end updateVersions
