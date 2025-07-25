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
	"net/url"
	"strings"
	"sync"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &ApplicationAwareRoutingPolicyDefinitionResource{}
var _ resource.ResourceWithImportState = &ApplicationAwareRoutingPolicyDefinitionResource{}

func NewApplicationAwareRoutingPolicyDefinitionResource() resource.Resource {
	return &ApplicationAwareRoutingPolicyDefinitionResource{}
}

type ApplicationAwareRoutingPolicyDefinitionResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *ApplicationAwareRoutingPolicyDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_aware_routing_policy_definition"
}

func (r *ApplicationAwareRoutingPolicyDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Application Aware Routing Policy Definition .").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the policy definition").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The description of the policy definition").String,
				Required:            true,
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of sequences").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence ID").String,
							Required:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence name").String,
							Required:            true,
						},
						"ip_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence IP type, either `ipv4`, `ipv6` or `all`").AddStringEnumDescription("ipv4", "ipv6", "all").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ipv4", "ipv6", "all"),
							},
						},
						"match_entries": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of match entries").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of match entry").AddStringEnumDescription("appList", "dnsAppList", "dns", "dscp", "plp", "protocol", "sourceDataPrefixList", "sourceIp", "sourcePort", "destinationDataPrefixList", "destinationIp", "destinationRegion", "destinationPort", "trafficTo", "icmpMessage").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("appList", "dnsAppList", "dns", "dscp", "plp", "protocol", "sourceDataPrefixList", "sourceIp", "sourcePort", "destinationDataPrefixList", "destinationIp", "destinationRegion", "destinationPort", "trafficTo", "icmpMessage"),
										},
									},
									"application_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Application list ID, Attribute conditional on `type` being equal to `appList`").String,
										Optional:            true,
									},
									"application_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Application list version").String,
										Optional:            true,
									},
									"dns_application_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("DNS Application list ID, Attribute conditional on `type` being equal to `dnsAppList`").String,
										Optional:            true,
									},
									"dns_application_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("DNS Application list version").String,
										Optional:            true,
									},
									"icmp_message": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("ICMP Message, Attribute conditional on `type` being equal to `icmpMessage`").String,
										Optional:            true,
									},
									"dns": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("DNS request or response, Attribute conditional on `type` being equal to `dns`").AddStringEnumDescription("request", "response").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("request", "response"),
										},
									},
									"dscp": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("DSCP value, Attribute conditional on `type` being equal to `dscp`").String,
										Optional:            true,
									},
									"plp": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("PLP, Attribute conditional on `type` being equal to `plp`").AddStringEnumDescription("low", "high").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("low", "high"),
										},
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP Protocol, 0-255 (Single value or multiple values separated by spaces), Attribute conditional on `type` being equal to `protocol`").String,
										Optional:            true,
									},
									"source_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source Data Prefix list ID, Attribute conditional on `type` being equal to `sourceDataPrefixList`").String,
										Optional:            true,
									},
									"source_data_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source Data Prefix list version").String,
										Optional:            true,
									},
									"source_ip": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source IP, Attribute conditional on `type` being equal to `sourceIp`").String,
										Optional:            true,
									},
									"source_port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source port, 0-65535 (Single value, range or multiple values separated by spaces), Attribute conditional on `type` being equal to `sourcePort`").String,
										Optional:            true,
									},
									"destination_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination Data Prefix list ID, Attribute conditional on `type` being equal to `destinationDataPrefixList`").String,
										Optional:            true,
									},
									"destination_data_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination Data Prefix list version").String,
										Optional:            true,
									},
									"destination_ip": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination IP, Attribute conditional on `type` being equal to `destinationIp`").String,
										Optional:            true,
									},
									"destination_port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination port, 0-65535 (Single value, range or multiple values separated by spaces), Attribute conditional on `type` being equal to `destinationPort`").String,
										Optional:            true,
									},
									"destination_region": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination region, Attribute conditional on `type` being equal to `destinationRegion`").AddStringEnumDescription("primary-region", "secondary-region", "other-region").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("primary-region", "secondary-region", "other-region"),
										},
									},
									"traffic_to": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Traffic to, Attribute conditional on `type` being equal to `trafficTo`").AddStringEnumDescription("access", "core", "service").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("access", "core", "service"),
										},
									},
								},
							},
						},
						"action_entries": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of action entries").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of action entry").AddStringEnumDescription("backupSlaPreferredColor", "count", "log", "slaClass", "cloudSaas").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("backupSlaPreferredColor", "count", "log", "slaClass", "cloudSaas"),
										},
									},
									"backup_sla_preferred_color": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Backup SLA preferred color (Single value or multiple values separated by spaces), Attribute conditional on `type` being equal to `backupSlaPreferredColor`").String,
										Optional:            true,
									},
									"counter": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Counter name, Attribute conditional on `type` being equal to `count`").String,
										Optional:            true,
									},
									"log": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Enable logging, Attribute conditional on `type` being equal to `log`").String,
										Optional:            true,
									},
									"cloud_sla": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Cloud SLA, Attribute conditional on `type` being equal to `cloudSaas`").String,
										Optional:            true,
									},
									"sla_class_parameters": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("List of SLA class parameters, Attribute conditional on `type` being equal to `slaClass`").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Type of SLA class parameter").AddStringEnumDescription("name", "preferredColor", "preferredColorGroup", "strict", "fallbackToBestPath").String,
													Required:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("name", "preferredColor", "preferredColorGroup", "strict", "fallbackToBestPath"),
													},
												},
												"sla_class_list": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("SLA class list ID, Attribute conditional on `type` being equal to `name`").String,
													Optional:            true,
												},
												"sla_class_list_version": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("SLA class list version").String,
													Optional:            true,
												},
												"preferred_color_group_list": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Preferred color group list ID, Attribute conditional on `type` being equal to `preferredColorGroup`").String,
													Optional:            true,
												},
												"preferred_color_group_list_version": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("Preferred color group list version").String,
													Optional:            true,
												},
												"preferred_color": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("preferred color (Single value or multiple values separated by spaces), Attribute conditional on `type` being equal to `preferredColor`").String,
													Optional:            true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *ApplicationAwareRoutingPolicyDefinitionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *ApplicationAwareRoutingPolicyDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplicationAwareRoutingPolicyDefinition

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("definitionId").String())
	plan.Version = types.Int64Value(0)
	plan.Type = types.StringValue("appRoute")

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *ApplicationAwareRoutingPolicyDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplicationAwareRoutingPolicyDefinition

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.ValueString()))

	res, err := r.client.Get(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if strings.Contains(res.Get("error.message").String(), "Failed to find specified resource") || strings.Contains(res.Get("error.message").String(), "Invalid template type") || strings.Contains(res.Get("error.message").String(), "Template definition not found") || strings.Contains(res.Get("error.message").String(), "Invalid Profile Id") || strings.Contains(res.Get("error.message").String(), "Invalid feature Id") || strings.Contains(res.Get("error.message").String(), "Invalid config group passed") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)
	if state.Version.IsNull() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *ApplicationAwareRoutingPolicyDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ApplicationAwareRoutingPolicyDefinition

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Name.ValueString()))

	if plan.hasChanges(ctx, &state) {
		body := plan.toBody(ctx)
		r.updateMutex.Lock()
		res, err := r.client.Put(plan.getPath()+url.QueryEscape(plan.Id.ValueString()), body)
		r.updateMutex.Unlock()
		if err != nil {
			if strings.Contains(res.Get("error.message").String(), "Failed to acquire lock") {
				resp.Diagnostics.AddWarning("Client Warning", "Failed to modify policy due to policy being locked by another change. Policy changes will not be applied. Re-run 'terraform apply' to try again.")
			} else if strings.Contains(res.Get("error.message").String(), "Template locked in edit mode") {
				resp.Diagnostics.AddWarning("Client Warning", "Failed to modify template due to template being locked by another change. Template changes will not be applied. Re-run 'terraform apply' to try again.")
			} else {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
				return
			}
		}
	} else {
		tflog.Debug(ctx, fmt.Sprintf("%s: No changes detected", plan.Name.ValueString()))
	}
	plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *ApplicationAwareRoutingPolicyDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplicationAwareRoutingPolicyDefinition

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *ApplicationAwareRoutingPolicyDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
