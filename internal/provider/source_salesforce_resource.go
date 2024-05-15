// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	speakeasy_objectplanmodifier "github.com/airbytehq/terraform-provider-airbyte/internal/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "github.com/airbytehq/terraform-provider-airbyte/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/airbytehq/terraform-provider-airbyte/internal/provider/types"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/operations"
	"github.com/airbytehq/terraform-provider-airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceSalesforceResource{}
var _ resource.ResourceWithImportState = &SourceSalesforceResource{}

func NewSourceSalesforceResource() resource.Resource {
	return &SourceSalesforceResource{}
}

// SourceSalesforceResource defines the resource implementation.
type SourceSalesforceResource struct {
	client *sdk.SDK
}

// SourceSalesforceResourceModel describes the resource data model.
type SourceSalesforceResourceModel struct {
	Configuration tfTypes.SourceSalesforce `tfsdk:"configuration"`
	DefinitionID  types.String             `tfsdk:"definition_id"`
	Name          types.String             `tfsdk:"name"`
	SecretID      types.String             `tfsdk:"secret_id"`
	SourceID      types.String             `tfsdk:"source_id"`
	SourceType    types.String             `tfsdk:"source_type"`
	WorkspaceID   types.String             `tfsdk:"workspace_id"`
}

func (r *SourceSalesforceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_salesforce"
}

func (r *SourceSalesforceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceSalesforce Resource",
		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Required: true,
				Attributes: map[string]schema.Attribute{
					"client_id": schema.StringAttribute{
						Required:    true,
						Description: `Enter your Salesforce developer application's <a href="https://developer.salesforce.com/forums/?id=9062I000000DLgbQAG">Client ID</a>`,
					},
					"client_secret": schema.StringAttribute{
						Required:    true,
						Description: `Enter your Salesforce developer application's <a href="https://developer.salesforce.com/forums/?id=9062I000000DLgbQAG">Client secret</a>`,
					},
					"force_use_bulk_api": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `Toggle to use Bulk API (this might cause empty fields for some streams). Default: false`,
					},
					"is_sandbox": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `Toggle if you're using a <a href="https://help.salesforce.com/s/articleView?id=sf.deploy_sandboxes_parent.htm&type=5">Salesforce Sandbox</a>. Default: false`,
					},
					"refresh_token": schema.StringAttribute{
						Required:    true,
						Sensitive:   true,
						Description: `Enter your application's <a href="https://developer.salesforce.com/docs/atlas.en-us.mobile_sdk.meta/mobile_sdk/oauth_refresh_token_flow.htm">Salesforce Refresh Token</a> used for Airbyte to access your Salesforce account.`,
					},
					"start_date": schema.StringAttribute{
						Optional:    true,
						Description: `Enter the date (or date-time) in the YYYY-MM-DD or YYYY-MM-DDTHH:mm:ssZ format. Airbyte will replicate the data updated on and after this date. If this field is blank, Airbyte will replicate the data for last two years.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"stream_slice_step": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("P30D"),
						Description: `The size of the time window (ISO8601 duration) to slice requests. Default: "P30D"`,
					},
					"streams_criteria": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"criteria": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Default:     stringdefault.StaticString("contains"),
									Description: `must be one of ["starts with", "ends with", "contains", "exacts", "starts not with", "ends not with", "not contains", "not exacts"]; Default: "contains"`,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"starts with",
											"ends with",
											"contains",
											"exacts",
											"starts not with",
											"ends not with",
											"not contains",
											"not exacts",
										),
									},
								},
								"value": schema.StringAttribute{
									Required: true,
								},
							},
						},
						Description: `Add filters to select only required stream based on ` + "`" + `SObject` + "`" + ` name. Use this field to filter which tables are displayed by this connector. This is useful if your Salesforce account has a large number of tables (>1000), in which case you may find it easier to navigate the UI and speed up the connector's performance if you restrict the tables displayed by this connector.`,
					},
				},
			},
			"definition_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed. `,
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `Name of the source e.g. dev-mysql-instance.`,
			},
			"secret_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed. `,
			},
			"source_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"source_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"workspace_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required: true,
			},
		},
	}
}

func (r *SourceSalesforceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceSalesforceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceSalesforceResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := data.ToSharedSourceSalesforceCreateRequest()
	res, err := r.client.Sources.CreateSourceSalesforce(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSourceResponse(res.SourceResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	sourceID := data.SourceID.ValueString()
	request1 := operations.GetSourceSalesforceRequest{
		SourceID: sourceID,
	}
	res1, err := r.client.Sources.GetSourceSalesforce(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if res1.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedSourceResponse(res1.SourceResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceSalesforceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceSalesforceResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	sourceID := data.SourceID.ValueString()
	request := operations.GetSourceSalesforceRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceSalesforce(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSourceResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceSalesforceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceSalesforceResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceSalesforcePutRequest := data.ToSharedSourceSalesforcePutRequest()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceSalesforceRequest{
		SourceSalesforcePutRequest: sourceSalesforcePutRequest,
		SourceID:                   sourceID,
	}
	res, err := r.client.Sources.PutSourceSalesforce(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	sourceId1 := data.SourceID.ValueString()
	request1 := operations.GetSourceSalesforceRequest{
		SourceID: sourceId1,
	}
	res1, err := r.client.Sources.GetSourceSalesforce(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if res1.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedSourceResponse(res1.SourceResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceSalesforceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceSalesforceResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	sourceID := data.SourceID.ValueString()
	request := operations.DeleteSourceSalesforceRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceSalesforce(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *SourceSalesforceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("source_id"), req.ID)...)
}
