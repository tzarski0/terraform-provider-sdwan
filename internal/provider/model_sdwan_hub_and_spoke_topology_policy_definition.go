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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type HubAndSpokeTopology struct {
	Id             types.String                    `tfsdk:"id"`
	Version        types.Int64                     `tfsdk:"version"`
	Type           types.String                    `tfsdk:"type"`
	Name           types.String                    `tfsdk:"name"`
	Description    types.String                    `tfsdk:"description"`
	VpnListId      types.String                    `tfsdk:"vpn_list_id"`
	VpnListVersion types.Int64                     `tfsdk:"vpn_list_version"`
	Topologies     []HubAndSpokeTopologyTopologies `tfsdk:"topologies"`
}

type HubAndSpokeTopologyTopologies struct {
	Name   types.String                          `tfsdk:"name"`
	Spokes []HubAndSpokeTopologyTopologiesSpokes `tfsdk:"spokes"`
}

type HubAndSpokeTopologyTopologiesSpokes struct {
	SiteListId      types.String                              `tfsdk:"site_list_id"`
	SiteListVersion types.Int64                               `tfsdk:"site_list_version"`
	Hubs            []HubAndSpokeTopologyTopologiesSpokesHubs `tfsdk:"hubs"`
}

type HubAndSpokeTopologyTopologiesSpokesHubs struct {
	SiteListId      types.String `tfsdk:"site_list_id"`
	SiteListVersion types.Int64  `tfsdk:"site_list_version"`
}

func (data HubAndSpokeTopology) getType() string {
	return "hubAndSpoke"
}

func (data HubAndSpokeTopology) toBody(ctx context.Context) string {
	body, _ := sjson.Set("", "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	body, _ = sjson.Set(body, "type", "hubAndSpoke")
	path := "definition."
	if !data.VpnListId.IsNull() {
		body, _ = sjson.Set(body, path+"vpnList", data.VpnListId.ValueString())
	}
	if len(data.Topologies) > 0 {
		body, _ = sjson.Set(body, path+"subDefinitions", []interface{}{})
		for _, item := range data.Topologies {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if len(item.Spokes) > 0 {
				itemBody, _ = sjson.Set(itemBody, "spokes", []interface{}{})
				for _, childItem := range item.Spokes {
					itemChildBody := ""
					if !childItem.SiteListId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "siteList", childItem.SiteListId.ValueString())
					}
					if len(childItem.Hubs) > 0 {
						itemChildBody, _ = sjson.Set(itemChildBody, "hubs", []interface{}{})
						for _, childChildItem := range childItem.Hubs {
							itemChildChildBody := ""
							if !childChildItem.SiteListId.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "siteList", childChildItem.SiteListId.ValueString())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "hubs.-1", itemChildChildBody)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "spokes.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"subDefinitions.-1", itemBody)
		}
	}
	return body
}

func (data *HubAndSpokeTopology) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	path := "definition."
	if value := res.Get(path + "vpnList"); value.Exists() {
		data.VpnListId = types.StringValue(value.String())
	} else {
		data.VpnListId = types.StringNull()
	}
	if value := res.Get(path + "subDefinitions"); value.Exists() {
		data.Topologies = make([]HubAndSpokeTopologyTopologies, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := HubAndSpokeTopologyTopologies{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("spokes"); cValue.Exists() {
				item.Spokes = make([]HubAndSpokeTopologyTopologiesSpokes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := HubAndSpokeTopologyTopologiesSpokes{}
					if ccValue := cv.Get("siteList"); ccValue.Exists() {
						cItem.SiteListId = types.StringValue(ccValue.String())
					} else {
						cItem.SiteListId = types.StringNull()
					}
					if ccValue := cv.Get("hubs"); ccValue.Exists() {
						cItem.Hubs = make([]HubAndSpokeTopologyTopologiesSpokesHubs, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := HubAndSpokeTopologyTopologiesSpokesHubs{}
							if cccValue := ccv.Get("siteList"); cccValue.Exists() {
								ccItem.SiteListId = types.StringValue(cccValue.String())
							} else {
								ccItem.SiteListId = types.StringNull()
							}
							cItem.Hubs = append(cItem.Hubs, ccItem)
							return true
						})
					}
					item.Spokes = append(item.Spokes, cItem)
					return true
				})
			}
			data.Topologies = append(data.Topologies, item)
			return true
		})
	}
}

func (data *HubAndSpokeTopology) hasChanges(ctx context.Context, state *HubAndSpokeTopology) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.VpnListId.Equal(state.VpnListId) {
		hasChanges = true
	}
	if len(data.Topologies) != len(state.Topologies) {
		hasChanges = true
	} else {
		for i := range data.Topologies {
			if !data.Topologies[i].Name.Equal(state.Topologies[i].Name) {
				hasChanges = true
			}
			if len(data.Topologies[i].Spokes) != len(state.Topologies[i].Spokes) {
				hasChanges = true
			} else {
				for ii := range data.Topologies[i].Spokes {
					if !data.Topologies[i].Spokes[ii].SiteListId.Equal(state.Topologies[i].Spokes[ii].SiteListId) {
						hasChanges = true
					}
					if len(data.Topologies[i].Spokes[ii].Hubs) != len(state.Topologies[i].Spokes[ii].Hubs) {
						hasChanges = true
					} else {
						for iii := range data.Topologies[i].Spokes[ii].Hubs {
							if !data.Topologies[i].Spokes[ii].Hubs[iii].SiteListId.Equal(state.Topologies[i].Spokes[ii].Hubs[iii].SiteListId) {
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

func (data *HubAndSpokeTopology) getSpokeSiteListVersion(ctx context.Context, name, spokeId string) types.Int64 {
	for _, item := range data.Topologies {
		if item.Name.ValueString() == name {
			for _, cItem := range item.Spokes {
				if cItem.SiteListId.ValueString() == spokeId {
					return cItem.SiteListVersion
				}
			}
		}
	}
	return types.Int64Null()
}

func (data *HubAndSpokeTopology) getHubSiteListVersion(ctx context.Context, name, spokeId, hubId string) types.Int64 {
	for _, item := range data.Topologies {
		if item.Name.ValueString() == name {
			for _, cItem := range item.Spokes {
				if cItem.SiteListId.ValueString() == spokeId {
					for _, ccItem := range cItem.Hubs {
						if ccItem.SiteListId.ValueString() == hubId {
							return ccItem.SiteListVersion
						}
					}
				}
			}
		}
	}
	return types.Int64Null()
}

func (data *HubAndSpokeTopology) updateVersions(ctx context.Context, state HubAndSpokeTopology) {
	data.VpnListVersion = state.VpnListVersion
	for t := range data.Topologies {
		name := data.Topologies[t].Name.ValueString()
		for s := range data.Topologies[t].Spokes {
			spokeId := data.Topologies[t].Spokes[s].SiteListId.ValueString()
			data.Topologies[t].Spokes[s].SiteListVersion = state.getSpokeSiteListVersion(ctx, name, spokeId)
			for h := range data.Topologies[t].Spokes[s].Hubs {
				hubId := data.Topologies[t].Spokes[s].Hubs[h].SiteListId.ValueString()
				data.Topologies[t].Spokes[s].Hubs[h].SiteListVersion = state.getHubSiteListVersion(ctx, name, spokeId, hubId)
			}
		}
	}
}
