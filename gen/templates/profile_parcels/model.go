//go:build ignore
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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

{{- $name := camelCase .Name}}
type {{camelCase .Name}} struct {
	Id types.String `tfsdk:"id"`
	Version types.Int64 `tfsdk:"version"`
	Name types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
{{- range .Attributes}}
{{- if eq .Type "List"}}
	{{toGoName .TfName}} []{{$name}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "StringList"}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- end}}
{{- end}}
}

{{ range .Attributes}}
{{- $childName := toGoName .TfName}}
{{- if eq .Type "List"}}
type {{$name}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if eq .Type "List"}}
	{{toGoName .TfName}} []{{$name}}{{$childName}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "StringList"}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{ end}}

{{ range .Attributes}}
{{- $childName := toGoName .TfName}}
{{- if eq .Type "List"}}
{{ range .Attributes}}
{{- $childChildName := toGoName .TfName}}
{{- if eq .Type "List"}}
type {{$name}}{{$childName}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if eq .Type "List"}}
	{{toGoName .TfName}} []{{$name}}{{$childName}}{{$childChildName}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "StringList"}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{ end}}

{{ range .Attributes}}
{{- $childName := toGoName .TfName}}
{{- if eq .Type "List"}}
{{ range .Attributes}}
{{- $childChildName := toGoName .TfName}}
{{- if eq .Type "List"}}
{{ range .Attributes}}
{{- if eq .Type "List"}}
type {{$name}}{{$childName}}{{$childChildName}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if eq .Type "StringList"}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{ end}}

func (data {{camelCase .Name}}) getModel() string {
	return "{{.Model}}"
}

func (data {{camelCase .Name}}) getPath() string {
	{{- if hasReference .Attributes}}
		return fmt.Sprintf("{{.RestEndpoint}}"{{range .Attributes}}{{if .Reference}}, data.{{toGoName .TfName}}.Value{{.Type}}(){{end}}{{end}})
	{{- else}}
		return "{{.RestEndpoint}}"
	{{- end}}
}

func (data {{camelCase .Name}}) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	{{- range .Attributes}}
	{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (eq .Type "StringList")) (not .Reference)}}
	{{if .Variable}}
	if !data.{{toGoName .TfName}}Variable.IsNull() {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "variable")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", data.{{toGoName .TfName}}Variable.ValueString())
	} else {{end}}{{if or .DefaultValue (not .ParcelMandatory)}}if data.{{toGoName .TfName}}.IsNull() {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "default")
		{{if .DefaultValuePresent}}body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}}){{end}}
	} else{{else}}if true{{end}} {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "global")
		{{- if eq .Type "StringList"}}
		var values []string
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", values)
		{{- else}}
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", data.{{toGoName .TfName}}.Value{{.Type}}())
		{{- end}}
	}
	{{- else if eq .Type "List"}}
	for _, item := range data.{{toGoName .TfName}} {
		itemBody := ""
		{{- range .Attributes}}
		{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (eq .Type "StringList")}}
		{{if .Variable}}
		if !item.{{toGoName .TfName}}Variable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", item.{{toGoName .TfName}}Variable.ValueString())
		} else {{end}}{{if or .DefaultValue (not .ParcelMandatory)}}if item.{{toGoName .TfName}}.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "default")
			{{if .DefaultValuePresent}}itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}}){{end}}
		} else{{else}}if true{{end}} {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "global")
			{{- if eq .Type "StringList"}}
			var values []string
			item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", values)
			{{- else}}
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", item.{{toGoName .TfName}}.Value{{.Type}}())
			{{- end}}
		}
		{{- else if eq .Type "List"}}
		for _, childItem := range item.{{toGoName .TfName}} {
			itemChildBody := ""
			{{- range .Attributes}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (eq .Type "StringList")}}
			{{if .Variable}}
			if !childItem.{{toGoName .TfName}}Variable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "variable")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", childItem.{{toGoName .TfName}}Variable.ValueString())
			} else {{end}}{{if or .DefaultValue (not .ParcelMandatory)}}if childItem.{{toGoName .TfName}}.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "default")
				{{if .DefaultValuePresent}}itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}}){{end}}
			} else{{else}}if true{{end}} {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "global")
				{{- if eq .Type "StringList"}}
				var values []string
				childItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", values)
				{{- else}}
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", childItem.{{toGoName .TfName}}.Value{{.Type}}())
				{{- end}}
			}
			{{- else if eq .Type "List"}}
			for _, childChildItem := range childItem.{{toGoName .TfName}} {
				itemChildChildBody := ""
				{{- range .Attributes}}
				{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (eq .Type "StringList")}}
				{{if .Variable}}
				if !childChildItem.{{toGoName .TfName}}Variable.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "variable")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", childChildItem.{{toGoName .TfName}}Variable.ValueString())
				} else {{end}}{{if or .DefaultValue (not .ParcelMandatory)}}if childChildItem.{{toGoName .TfName}}.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "default")
					{{if .DefaultValuePresent}}itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}}){{end}}
				} else{{else}}if true{{end}} {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "global")
					{{- if eq .Type "StringList"}}
					var values []string
					childChildItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", values)
					{{- else}}
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", childChildItem.{{toGoName .TfName}}.Value{{.Type}}())
					{{- end}}
				}
				{{- end}}
				{{- end}}
				itemChildBody, _ = sjson.SetRaw(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemChildChildBody)
			}
			{{- end}}
			{{- end}}
			itemBody, _ = sjson.SetRaw(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemChildBody)
		}
		{{- end}}
		{{- end}}
		body, _ = sjson.SetRaw(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemBody)
	}
	{{- end}}
	{{- end}}
	return body
}

func (data *{{camelCase .Name}}) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	{{- range .Attributes}}
	{{- $cname := toGoName .TfName}}
	{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (eq .Type "StringList")) (not .Reference) (not .WriteOnly)}}
	data.{{toGoName .TfName}} = {{if eq .Type "StringList"}}types.ListNull(types.StringType){{else}}types.{{.Type}}Null(){{end}}
	{{ if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
	if t := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
		va := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
		{{if .Variable}}if t.String() == "variable" {
			data.{{toGoName .TfName}}Variable = types.StringValue(va.String())
		} else{{end}} if t.String() == "global" {
			data.{{toGoName .TfName}} = {{if eq .Type "StringList"}}helpers.GetStringList(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
		}
	}
	{{- else if eq .Type "List"}}
	if value := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
		data.{{toGoName .TfName}} = make([]{{$name}}{{toGoName .TfName}}, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := {{$name}}{{toGoName .TfName}}{}
			{{- range .Attributes}}
			{{- $ccname := toGoName .TfName}}
			{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (eq .Type "StringList")) (not .Reference) (not .WriteOnly)}}
			item.{{toGoName .TfName}} = {{if eq .Type "StringList"}}types.ListNull(types.StringType){{else}}types.{{.Type}}Null(){{end}}
			{{ if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
			if t := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
				va := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
				{{if .Variable}}if t.String() == "variable" {
					item.{{toGoName .TfName}}Variable = types.StringValue(va.String())
				} else{{end}} if t.String() == "global" {
					item.{{toGoName .TfName}} = {{if eq .Type "StringList"}}helpers.GetStringList(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
				}
			}
			{{- else if eq .Type "List"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{toGoName .TfName}}, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := {{$name}}{{$cname}}{{toGoName .TfName}}{}
					{{- range .Attributes}}
					{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (eq .Type "StringList")) (not .Reference) (not .WriteOnly)}}
					cItem.{{toGoName .TfName}} = {{if eq .Type "StringList"}}types.ListNull(types.StringType){{else}}types.{{.Type}}Null(){{end}}
					{{ if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
					if t := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
						va := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
						{{if .Variable}}if t.String() == "variable" {
							cItem.{{toGoName .TfName}}Variable = types.StringValue(va.String())
						} else{{end}} if t.String() == "global" {
							cItem.{{toGoName .TfName}} = {{if eq .Type "StringList"}}helpers.GetStringList(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
						}
					}
					{{- else if eq .Type "List"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := {{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}{}
							{{- range .Attributes}}
							{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (eq .Type "StringList")) (not .Reference) (not .WriteOnly)}}
							ccItem.{{toGoName .TfName}} = {{if eq .Type "StringList"}}types.ListNull(types.StringType){{else}}types.{{.Type}}Null(){{end}}
							{{ if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
							if t := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
								va := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
								{{if .Variable}}if t.String() == "variable" {
									ccItem.{{toGoName .TfName}}Variable = types.StringValue(va.String())
								} else{{end}} if t.String() == "global" {
									ccItem.{{toGoName .TfName}} = {{if eq .Type "StringList"}}helpers.GetStringList(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
								}
							}
							{{- end}}
							{{- end}}
							cItem.{{toGoName .TfName}} = append(cItem.{{toGoName .TfName}}, ccItem)
							return true
						})
					}
					{{- end}}
					{{- end}}
					item.{{toGoName .TfName}} = append(item.{{toGoName .TfName}}, cItem)
					return true
				})
			}
			{{- end}}
			{{- end}}
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, item)
			return true
		})
	}
	{{- end}}
	{{- end}}
}

func (data *{{camelCase .Name}}) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	{{- range .Attributes}}
	{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (eq .Type "StringList")) (not .Reference) (not .WriteOnly)}}
	data.{{toGoName .TfName}} = {{if eq .Type "StringList"}}types.ListNull(types.StringType){{else}}types.{{.Type}}Null(){{end}}
	{{ if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
	if t := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
		va := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
		{{if .Variable}}if t.String() == "variable" {
			data.{{toGoName .TfName}}Variable = types.StringValue(va.String())
		} else{{end}} if t.String() == "global" {
			data.{{toGoName .TfName}} = {{if eq .Type "StringList"}}helpers.GetStringList(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
		}
	}
	{{- else if eq .Type "List"}}
	{{- $list := (toGoName .TfName)}}
	for i := range data.{{toGoName .TfName}} {
		keys := [...]string{ {{range .Attributes}}{{if .Id}}{{if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String")}}"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{end}}{{end}}{{end}} }
		keyValues := [...]string{ {{range .Attributes}}{{if .Id}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{toGoName .TfName}}.ValueBool()), {{else if eq .Type "String"}}data.{{$list}}[i].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }
		keyValuesVariables := [...]string{ {{range .Attributes}}{{if .Id}}{{if .Variable}}data.{{$list}}[i].{{toGoName .TfName}}Variable.ValueString(), {{else}}"", {{end}}{{end}}{{end}} }

		var r gjson.Result
		res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)

		{{- range .Attributes}}
		{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (eq .Type "StringList")) (not .Reference) (not .WriteOnly)}}
		data.{{$list}}[i].{{toGoName .TfName}} = {{if eq .Type "StringList"}}types.ListNull(types.StringType){{else}}types.{{.Type}}Null(){{end}}
		{{ if .Variable}}data.{{$list}}[i].{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		if t := r.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
			va := r.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
			{{if .Variable}}if t.String() == "variable" {
				data.{{$list}}[i].{{toGoName .TfName}}Variable = types.StringValue(va.String())
			} else{{end}} if t.String() == "global" {
				data.{{$list}}[i].{{toGoName .TfName}} = {{if eq .Type "StringList"}}helpers.GetStringList(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
			}
		}
		{{- else if eq .Type "List"}}
		{{- $clist := (toGoName .TfName)}}
		for ci := range data.{{$list}}[i].{{toGoName .TfName}} {
			keys := [...]string{ {{range .Attributes}}{{if .Id}}{{if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String")}}"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{end}}{{end}}{{end}} }
			keyValues := [...]string{ {{range .Attributes}}{{if .Id}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.ValueBool()), {{else if eq .Type "String"}}data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }
			keyValuesVariables := [...]string{ {{range .Attributes}}{{if .Id}}{{if .Variable}}data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}Variable.ValueString(), {{else}}"", {{end}}{{end}}{{end}} }

			var cr gjson.Result
			r.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType").String()
						vv := v.Get(keys[ik] + ".value").String()
						if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)

			{{- range .Attributes}}
			{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (eq .Type "StringList")) (not .Reference) (not .WriteOnly)}}
			data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = {{if eq .Type "StringList"}}types.ListNull(types.StringType){{else}}types.{{.Type}}Null(){{end}}
			{{ if .Variable}}data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}Variable = types.StringNull(){{end}}
			if t := cr.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
				va := cr.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
				{{if .Variable}}if t.String() == "variable" {
					data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}Variable = types.StringValue(va.String())
				} else{{end}} if t.String() == "global" {
					data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = {{if eq .Type "StringList"}}helpers.GetStringList(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
				}
			}
			{{- else if eq .Type "List"}}
			{{- $cclist := (toGoName .TfName)}}
			for cci := range data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} {
				keys := [...]string{ {{range .Attributes}}{{if .Id}}{{if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String")}}"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{end}}{{end}}{{end}} }
				keyValues := [...]string{ {{range .Attributes}}{{if .Id}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.ValueBool()), {{else if eq .Type "String"}}data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }
				keyValuesVariables := [...]string{ {{range .Attributes}}{{if .Id}}{{if .Variable}}data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}Variable.ValueString(), {{else}}"", {{end}}{{end}}{{end}} }

				var ccr gjson.Result
				cr.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").ForEach(
					func(_, v gjson.Result) bool {
						found := false
						for ik := range keys {
							tt := v.Get(keys[ik] + ".optionType").String()
							vv := v.Get(keys[ik] + ".value").String()
							if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
								found = true
								continue
							}
							found = false
							break
						}
						if found {
							ccr = v
							return false
						}
						return true
					},
				)

				{{- range .Attributes}}
				{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (eq .Type "StringList")) (not .Reference) (not .WriteOnly)}}
				data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}} = {{if eq .Type "StringList"}}types.ListNull(types.StringType){{else}}types.{{.Type}}Null(){{end}}
				{{ if .Variable}}data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				if t := ccr.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
					va := ccr.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
					{{if .Variable}}if t.String() == "variable" {
						data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}Variable = types.StringValue(va.String())
					} else{{end}} if t.String() == "global" {
						data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}} = {{if eq .Type "StringList"}}helpers.GetStringList(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
					}
				}
				{{- end}}
				{{- end}}
			}
			{{- end}}
			{{- end}}
		}
		{{- end}}
		{{- end}}
	}
	{{- end}}
	{{- end}}
}
