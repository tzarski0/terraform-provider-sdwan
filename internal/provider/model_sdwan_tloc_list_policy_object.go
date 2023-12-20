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
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type TLOCListPolicyObject struct {
	Id      types.String                  `tfsdk:"id"`
	Version types.Int64                   `tfsdk:"version"`
	Name    types.String                  `tfsdk:"name"`
	Entries []TLOCListPolicyObjectEntries `tfsdk:"entries"`
}

type TLOCListPolicyObjectEntries struct {
	TlocIp        types.String `tfsdk:"tloc_ip"`
	Color         types.String `tfsdk:"color"`
	Encapsulation types.String `tfsdk:"encapsulation"`
	Preference    types.Int64  `tfsdk:"preference"`
}

func (data TLOCListPolicyObject) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "tloc")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "entries", []interface{}{})
		for _, item := range data.Entries {
			itemBody := ""
			if !item.TlocIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "tloc", item.TlocIp.ValueString())
			}
			if !item.Color.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "color", item.Color.ValueString())
			}
			if !item.Encapsulation.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "encap", item.Encapsulation.ValueString())
			}
			if !item.Preference.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "preference", fmt.Sprint(item.Preference.ValueInt64()))
			}
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

func (data *TLOCListPolicyObject) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("entries"); value.Exists() && len(value.Array()) > 0 {
		data.Entries = make([]TLOCListPolicyObjectEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TLOCListPolicyObjectEntries{}
			if cValue := v.Get("tloc"); cValue.Exists() {
				item.TlocIp = types.StringValue(cValue.String())
			} else {
				item.TlocIp = types.StringNull()
			}
			if cValue := v.Get("color"); cValue.Exists() {
				item.Color = types.StringValue(cValue.String())
			} else {
				item.Color = types.StringNull()
			}
			if cValue := v.Get("encap"); cValue.Exists() {
				item.Encapsulation = types.StringValue(cValue.String())
			} else {
				item.Encapsulation = types.StringNull()
			}
			if cValue := v.Get("preference"); cValue.Exists() {
				item.Preference = types.Int64Value(cValue.Int())
			} else {
				item.Preference = types.Int64Null()
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}

func (data *TLOCListPolicyObject) hasChanges(ctx context.Context, state *TLOCListPolicyObject) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if len(data.Entries) != len(state.Entries) {
		hasChanges = true
	} else {
		for i := range data.Entries {
			if !data.Entries[i].TlocIp.Equal(state.Entries[i].TlocIp) {
				hasChanges = true
			}
			if !data.Entries[i].Color.Equal(state.Entries[i].Color) {
				hasChanges = true
			}
			if !data.Entries[i].Encapsulation.Equal(state.Entries[i].Encapsulation) {
				hasChanges = true
			}
			if !data.Entries[i].Preference.Equal(state.Entries[i].Preference) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}
