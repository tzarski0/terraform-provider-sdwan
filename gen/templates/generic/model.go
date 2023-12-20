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
{{- if .HasVersion}}
	Version types.Int64 `tfsdk:"version"`
{{- end}}
{{- range .Attributes}}
{{- if not .Value}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
	{{toGoName .TfName}} []{{$name}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if or (eq .Type "StringList") (eq .Type "Versions")}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "Version"}}
	{{toGoName .TfName}} types.Int64 `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}

{{ range .Attributes}}
{{- if not .Value}}
{{- $childName := toGoName .TfName}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
type {{$name}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
	{{toGoName .TfName}} []{{$name}}{{$childName}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if or (eq .Type "StringList") (eq .Type "Versions")}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "Version"}}
	{{toGoName .TfName}} types.Int64 `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{ end}}

{{ range .Attributes}}
{{- if not .Value}}
{{- $childName := toGoName .TfName}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
{{ range .Attributes}}
{{- if not .Value}}
{{- $childChildName := toGoName .TfName}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
type {{$name}}{{$childName}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
	{{toGoName .TfName}} []{{$name}}{{$childName}}{{$childChildName}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if or (eq .Type "StringList") (eq .Type "Versions")}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "Version"}}
	{{toGoName .TfName}} types.Int64 `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
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

{{ range .Attributes}}
{{- if not .Value}}
{{- $childName := toGoName .TfName}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
{{ range .Attributes}}
{{- if not .Value}}
{{- $childChildName := toGoName .TfName}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
{{ range .Attributes}}
{{- if not .Value}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
type {{$name}}{{$childName}}{{$childChildName}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if or (eq .Type "StringList") (eq .Type "Versions")}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "Version"}}
	{{toGoName .TfName}} types.Int64 `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{ end}}

func (data {{camelCase .Name}}) toBody(ctx context.Context) string {
	body := ""
	{{- range .Attributes}}
	{{- if .Value}}
	if true{{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
	}
	{{- else if not .TfOnly}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
	if {{if not .AlwaysInclude}}!data.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}data.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
	}
	{{- else if eq .Type "Bool"}}
	if {{if not .AlwaysInclude}}!data.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
		if {{.BoolEmptyString}} && data.{{toGoName .TfName}}.Value{{.Type}}() {
			body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", "")
		} else {
			body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}data.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
		}
	}
	{{- else if eq .Type "StringList"}}
	if {{if not .AlwaysInclude}}!data.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
		var values []string
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
	}
	{{- else if or (eq .Type "List") (eq .Type "Set")}}
	if true{{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
		for _, item := range data.{{toGoName .TfName}} {
			itemBody := ""
			{{- range .Attributes}}
			{{- if .Value}}
			if true{{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
			}
			{{- else if not .TfOnly}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
			if {{if not .AlwaysInclude}}!item.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}item.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
			}
			{{- else if eq .Type "Bool"}}
			if {{if not .AlwaysInclude}}!item.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
				if {{.BoolEmptyString}} && item.{{toGoName .TfName}}.Value{{.Type}}() {
					itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", "")
				} else {
					itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}item.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
				}
			}
			{{- else if eq .Type "StringList"}}
			if {{if not .AlwaysInclude}}!item.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
				var values []string
				item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
			}
			{{- else if or (eq .Type "List") (eq .Type "Set")}}
			if true{{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
				for _, childItem := range item.{{toGoName .TfName}} {
					itemChildBody := ""
					{{- range .Attributes}}
					{{- if .Value}}
					if true{{if ne .ConditionalAttribute.Name ""}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
					}
					{{- else if not .TfOnly}}
					{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
					if {{if not .AlwaysInclude}}!childItem.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}childItem.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
					}
					{{- else if eq .Type "Bool"}}
					if {{if not .AlwaysInclude}}!childItem.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
						if {{.BoolEmptyString}} && childItem.{{toGoName .TfName}}.Value{{.Type}}() {
							itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", "")
						} else {
							itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}childItem.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
						}
					}
					{{- else if eq .Type "StringList"}}
					if {{if not .AlwaysInclude}}!childItem.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
						var values []string
						childItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
					}
					{{- else if or (eq .Type "List") (eq .Type "Set")}}
					if true{{if ne .ConditionalAttribute.Name ""}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
						for _, childChildItem := range childItem.{{toGoName .TfName}} {
							itemChildChildBody := ""
							{{- range .Attributes}}
							{{- if .Value}}
							if true{{if ne .ConditionalAttribute.Name ""}} && childChildItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
							}
							{{- else if not .TfOnly}}
							{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
							if {{if not .AlwaysInclude}}!childChildItem.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && childChildItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}childChildItem.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
							}
							{{- else if eq .Type "Bool"}}
							if {{if not .AlwaysInclude}}!childChildItem.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && childChildItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
								if {{.BoolEmptyString}} && childChildItem.{{toGoName .TfName}}.Value{{.Type}}() {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", "")
								} else {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}childChildItem.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
								}
							}
							{{- else if eq .Type "StringList"}}
							if {{if not .AlwaysInclude}}!childChildItem.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && childChildItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
								var values []string
								childChildItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
							}
							{{- end}}
							{{- end}}
							{{- end}}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemChildChildBody)
						}
					}
					{{- end}}
					{{- end}}
					{{- end}}
					itemBody, _ = sjson.SetRaw(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemChildBody)
				}
			}
			{{- end}}
			{{- end}}
			{{- end}}
			body, _ = sjson.SetRaw(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemBody)
		}
	}
	{{- end}}
	{{- end}}
	{{- end}}
	return body
}

func (data *{{camelCase .Name}}) fromBody(ctx context.Context, res gjson.Result) {
	{{- if hasVersionAttribute .Attributes}}
	state := *data
	{{- end}}
	{{- range .Attributes}}
	{{- if and (not .TfOnly) (not .Value)}}
	{{- $cname := toGoName .TfName}}
	{{- if eq .Type "String"}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); value.Exists(){{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && value.String() != ""{{end}} {
		data.{{toGoName .TfName}} = types.StringValue(value.String())
	} else {
		data.{{toGoName .TfName}} = types.StringNull()
	}
	{{- else if eq .Type "Int64"}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); value.Exists(){{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && value.String() != ""{{end}} {
		data.{{toGoName .TfName}} = types.Int64Value(value.Int())
	} else {
		data.{{toGoName .TfName}} = types.Int64Null()
	}
	{{- else if eq .Type "Float64"}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); value.Exists(){{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && value.String() != ""{{end}} {
		data.{{toGoName .TfName}} = types.Float64Value(value.Float())
	} else {
		data.{{toGoName .TfName}} = types.Float64Null()
	}
	{{- else if eq .Type "Bool"}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); value.Exists(){{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && value.String() != ""{{end}} {
		if {{.BoolEmptyString}} && value.String() == "" {
			data.{{toGoName .TfName}} = types.BoolValue(true)
		} else {
			data.{{toGoName .TfName}} = types.BoolValue(value.Bool())
		}
	} else {
		data.{{toGoName .TfName}} = types.BoolNull()
	}
	{{- else if eq .Type "StringList"}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); value.Exists(){{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && value.String() != ""{{end}} {
		data.{{toGoName .TfName}} = helpers.GetStringList(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.ListNull(types.StringType)
	}
	{{- else if or (eq .Type "List") (eq .Type "Set")}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); value.Exists() && len(value.Array()) > 0{{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
		data.{{toGoName .TfName}} = make([]{{$name}}{{toGoName .TfName}}, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := {{$name}}{{toGoName .TfName}}{}
			{{- range .Attributes}}
			{{- $ccname := toGoName .TfName}}
			{{- if and (not .TfOnly) (not .Value)}}
			{{- if eq .Type "String"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); cValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cValue.String() != ""{{end}} {
				item.{{toGoName .TfName}} = types.StringValue(cValue.String())
			} else {
				item.{{toGoName .TfName}} = types.StringNull()
			}
			{{- else if eq .Type "Int64"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); cValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cValue.String() != ""{{end}} {
				item.{{toGoName .TfName}} = types.Int64Value(cValue.Int())
			} else {
				item.{{toGoName .TfName}} = types.Int64Null()
			}
			{{- else if eq .Type "Float64"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); cValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cValue.String() != ""{{end}} {
				item.{{toGoName .TfName}} = types.Float64Value(cValue.Float())
			} else {
				item.{{toGoName .TfName}} = types.Float64Null()
			}
			{{- else if eq .Type "Bool"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); cValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cValue.String() != ""{{end}} {
				if {{.BoolEmptyString}} && cValue.String() == "" {
					item.{{toGoName .TfName}} = types.BoolValue(true)
				} else {
					item.{{toGoName .TfName}} = types.BoolValue(cValue.Bool())
				}
			} else {
				item.{{toGoName .TfName}} = types.BoolNull()
			}
			{{- else if eq .Type "StringList"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); cValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cValue.String() != ""{{end}} {
				item.{{toGoName .TfName}} = helpers.GetStringList(cValue.Array())
			} else {
				item.{{toGoName .TfName}} = types.ListNull(types.StringType)
			}
			{{- else if or (eq .Type "List") (eq .Type "Set")}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); cValue.Exists() && len(cValue.Array()) > 0{{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
				item.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{toGoName .TfName}}, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := {{$name}}{{$cname}}{{toGoName .TfName}}{}
					{{- range .Attributes}}
					{{- if and (not .TfOnly) (not .Value)}}
					{{- if eq .Type "String"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); ccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && cItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && ccValue.String() != ""{{end}} {
						cItem.{{toGoName .TfName}} = types.StringValue(ccValue.String())
					} else {
						cItem.{{toGoName .TfName}} = types.StringNull()
					}
					{{- else if eq .Type "Int64"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); ccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && cItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && ccValue.String() != ""{{end}} {
						cItem.{{toGoName .TfName}} = types.Int64Value(ccValue.Int())
					} else {
						cItem.{{toGoName .TfName}} = types.Int64Null()
					}
					{{- else if eq .Type "Float64"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); ccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && cItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && ccValue.String() != ""{{end}} {
						cItem.{{toGoName .TfName}} = types.Float64Value(ccValue.Float())
					} else {
						cItem.{{toGoName .TfName}} = types.Float64Null()
					}
					{{- else if eq .Type "Bool"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); ccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && cItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && ccValue.String() != ""{{end}} {
						if {{.BoolEmptyString}} && ccValue.String() == "" {
							cItem.{{toGoName .TfName}} = types.BoolValue(true)
						} else {
							cItem.{{toGoName .TfName}} = types.BoolValue(ccValue.Bool())
						}
					} else {
						cItem.{{toGoName .TfName}} = types.BoolNull()
					}
					{{- else if eq .Type "StringList"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); ccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && cItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && ccValue.String() != ""{{end}} {
						cItem.{{toGoName .TfName}} = helpers.GetStringList(ccValue.Array())
					} else {
						cItem.{{toGoName .TfName}} = types.ListNull(types.StringType)
					}
					{{- else if or (eq .Type "List") (eq .Type "Set")}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); ccValue.Exists() && len(ccValue.Array()) > 0{{if ne .ConditionalAttribute.Name ""}} && cItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
						cItem.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := {{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}{}
							{{- range .Attributes}}
							{{- if and (not .TfOnly) (not .Value)}}
							{{- if eq .Type "String"}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); cccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && ccItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cccValue.String() != ""{{end}} {
								ccItem.{{toGoName .TfName}} = types.StringValue(cccValue.String())
							} else {
								ccItem.{{toGoName .TfName}} = types.StringNull()
							}
							{{- else if eq .Type "Int64"}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); cccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && ccItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cccValue.String() != ""{{end}} {
								ccItem.{{toGoName .TfName}} = types.Int64Value(cccValue.Int())
							} else {
								ccItem.{{toGoName .TfName}} = types.Int64Null()
							}
							{{- else if eq .Type "Float64"}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); cccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && ccItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cccValue.String() != ""{{end}} {
								ccItem.{{toGoName .TfName}} = types.Float64Value(cccValue.Float())
							} else {
								ccItem.{{toGoName .TfName}} = types.Float64Null()
							}
							{{- else if eq .Type "Bool"}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); cccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && ccItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cccValue.String() != ""{{end}} {
								if {{.BoolEmptyString}} && cccValue.String() == "" {
									ccItem.{{toGoName .TfName}} = types.BoolValue(true)
								} else {
									ccItem.{{toGoName .TfName}} = types.BoolValue(cccValue.Bool())
								}
							} else {
								ccItem.{{toGoName .TfName}} = types.BoolNull()
							}
							{{- else if eq .Type "StringList"}}
							if cccValue := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{getResponseModelName .}}"); cccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && ccItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cccValue.String() != ""{{end}} {
								ccItem.{{toGoName .TfName}} = helpers.GetStringList(cccValue.Array())
							} else {
								ccItem.{{toGoName .TfName}} = types.ListNull(types.StringType)
							}
							{{- end}}
							{{- end}}
							{{- end}}
							cItem.{{toGoName .TfName}} = append(cItem.{{toGoName .TfName}}, ccItem)
							return true
						})
					}
					{{- end}}
					{{- end}}
					{{- end}}
					item.{{toGoName .TfName}} = append(item.{{toGoName .TfName}}, cItem)
					return true
				})
			}
			{{- end}}
			{{- end}}
			{{- end}}
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, item)
			return true
		})
	}
	{{- end}}
	{{- end}}
	{{- end}}
	{{- if hasVersionAttribute .Attributes}}
	data.updateVersions(ctx, &state)
	{{- end}}
}

func (data *{{camelCase .Name}}) hasChanges(ctx context.Context, state *{{camelCase .Name}}) bool {
	hasChanges := false
	{{- range .Attributes}}
	{{- $name := toGoName .TfName}}
	{{- if and (not .Value) (not .TfOnly)}}
	{{- if and (ne .Type "List") (ne .Type "Set")}}
	if !data.{{toGoName .TfName}}.Equal(state.{{toGoName .TfName}}) {
		hasChanges = true
	}
	{{- else}}
	if len(data.{{toGoName .TfName}}) != len(state.{{toGoName .TfName}}) {
		hasChanges = true
	} else {
		for i := range data.{{toGoName .TfName}} {
			{{- range .Attributes}}
			{{- $cname := toGoName .TfName}}
			{{- if and (not .Value) (not .TfOnly)}}
			{{- if and (ne .Type "List") (ne .Type "Set")}}
			if !data.{{$name}}[i].{{toGoName .TfName}}.Equal(state.{{$name}}[i].{{toGoName .TfName}}) {
				hasChanges = true
			}
			{{- else}}
			if len(data.{{$name}}[i].{{toGoName .TfName}}) != len(state.{{$name}}[i].{{toGoName .TfName}}) {
				hasChanges = true
			} else {
				for ii := range data.{{$name}}[i].{{toGoName .TfName}} {
					{{- range .Attributes}}
					{{- $ccname := toGoName .TfName}}
					{{- if and (not .Value) (not .TfOnly)}}
					{{- if and (ne .Type "List") (ne .Type "Set")}}
					if !data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}.Equal(state.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}) {
						hasChanges = true
					}
					{{- else}}
					if len(data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}) != len(state.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}) {
						hasChanges = true
					} else {
						for iii := range data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}} {
							{{- range .Attributes}}
							{{- if and (not .Value) (not .TfOnly)}}
							if !data.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}}.Equal(state.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}}) {
								hasChanges = true
							}
							{{- end}}
							{{- end}}
						}
					}
					{{- end}}
					{{- end}}
					{{- end}}
				}
			}
			{{- end}}
			{{- end}}
			{{- end}}
		}
	}
	{{- end}}
	{{- end}}
	{{- end}}
	return hasChanges
}

{{if hasVersionAttribute .Attributes}}
func (data *{{camelCase .Name}}) updateVersions(ctx context.Context, state *{{camelCase .Name}}) {
	{{- range .Attributes}}
	{{- $name := toGoName .TfName}}
	{{- if eq .Type "Version"}}
	data.{{toGoName .TfName}} = state.{{toGoName .TfName}}
	{{- else if and (or (eq .Type "List") (eq .Type "Set")) (hasVersionAttribute .Attributes)}}
	for i := range data.{{toGoName .TfName}} {
		dataKeys := [...]string{ {{range .Attributes}}{{if .Id}}fmt.Sprintf("%v", data.{{$name}}[i].{{toGoName .TfName}}.{{if eq .Type "StringList"}}String{{else}}Value{{.Type}}{{end}}()), {{end}}{{end}} }
		stateIndex := -1
		for j := range state.{{toGoName .TfName}} {
			stateKeys := [...]string{ {{range .Attributes}}{{if .Id}}fmt.Sprintf("%v", state.{{$name}}[j].{{toGoName .TfName}}.{{if eq .Type "StringList"}}String{{else}}Value{{.Type}}{{end}}()), {{end}}{{end}} }
			if dataKeys == stateKeys {
				stateIndex = j
                break
			}
		}
		{{- range .Attributes}}
		{{- $cname := toGoName .TfName}}
		{{- if eq .Type "Version"}}
		if stateIndex > -1 {
			data.{{$name}}[i].{{toGoName .TfName}} = state.{{$name}}[stateIndex].{{toGoName .TfName}}
		} else {
			data.{{$name}}[i].{{toGoName .TfName}} = types.Int64Null()
		}
		{{- else if eq .Type "Versions"}}
		if stateIndex > -1 && !state.{{$name}}[stateIndex].{{toGoName .TfName}}.IsNull() {
			data.{{$name}}[i].{{toGoName .TfName}} = state.{{$name}}[stateIndex].{{toGoName .TfName}}
		} else {
			data.{{$name}}[i].{{toGoName .TfName}} = types.ListNull(types.StringType)
		}
		{{- else if and (or (eq .Type "List") (eq .Type "Set")) (hasVersionAttribute .Attributes)}}
		for ii := range data.{{$name}}[i].{{toGoName .TfName}} {
			cDataKeys := [...]string{ {{range .Attributes}}{{if .Id}}fmt.Sprintf("%v", data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}.{{if eq .Type "StringList"}}String{{else}}Value{{.Type}}{{end}}()), {{end}}{{end}} }
			cStateIndex := -1
			if stateIndex > -1 {
				for jj := range state.{{$name}}[stateIndex].{{toGoName .TfName}} {
					cStateKeys := [...]string{ {{range .Attributes}}{{if .Id}}fmt.Sprintf("%v", state.{{$name}}[stateIndex].{{$cname}}[jj].{{toGoName .TfName}}.{{if eq .Type "StringList"}}String{{else}}Value{{.Type}}{{end}}()), {{end}}{{end}} }
					if cDataKeys == cStateKeys {
						cStateIndex = jj
						break
					}
				}
			}
			{{- range .Attributes}}
			{{- $ccname := toGoName .TfName}}
			{{- if eq .Type "Version"}}
			if cStateIndex > -1 {
				data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}} = state.{{$name}}[stateIndex].{{$cname}}[cStateIndex].{{toGoName .TfName}}
			} else {
				data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}} = types.Int64Null()
			}
			{{- else if eq .Type "Versions"}}
			if cStateIndex > -1 && !state.{{$name}}[stateIndex].{{$cname}}[cStateIndex].{{toGoName .TfName}}.IsNull() {
				data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}} = state.{{$name}}[stateIndex].{{$cname}}[cStateIndex].{{toGoName .TfName}}
			} else {
				data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}} = types.ListNull(types.StringType)
			}
			{{- else if and (or (eq .Type "List") (eq .Type "Set")) (hasVersionAttribute .Attributes)}}
			for iii := range data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}} {
				ccDataKeys := [...]string{ {{range .Attributes}}{{if .Id}}fmt.Sprintf("%v", data.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}}.{{if eq .Type "StringList"}}String{{else}}Value{{.Type}}{{end}}()), {{end}}{{end}} }
				ccStateIndex := -1
				if cStateIndex > -1 {
					for jjj := range state.{{$name}}[stateIndex].{{$cname}}[cStateIndex].{{toGoName .TfName}} {
						ccStateKeys := [...]string{ {{range .Attributes}}{{if .Id}}fmt.Sprintf("%v", state.{{$name}}[stateIndex].{{$cname}}[cStateIndex].{{$ccname}}[jjj].{{toGoName .TfName}}.{{if eq .Type "StringList"}}String{{else}}Value{{.Type}}{{end}}()), {{end}}{{end}} }
						if ccDataKeys == ccStateKeys {
							ccStateIndex = jjj
							break
						}
					}
				}
				{{- range .Attributes}}
				{{- if eq .Type "Version"}}
				if ccStateIndex > -1 {
					data.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}} = state.{{$name}}[stateIndex].{{$cname}}[cStateIndex].{{$ccname}}[ccStateIndex].{{toGoName .TfName}}
				} else {
					data.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}} = types.Int64Null()
				}
				{{- else if eq .Type "Versions"}}
				if ccStateIndex > -1 && !state.{{$name}}[stateIndex].{{$cname}}[cStateIndex].{{$ccname}}[ccStateIndex].{{toGoName .TfName}}.IsNull() {
					data.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}} = state.{{$name}}[stateIndex].{{$cname}}[cStateIndex].{{$ccname}}[ccStateIndex].{{toGoName .TfName}}
				} else {
					data.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}} = types.ListNull(types.StringType)
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
{{end}}
