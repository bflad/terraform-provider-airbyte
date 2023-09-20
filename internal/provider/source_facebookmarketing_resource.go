// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

	speakeasy_boolplanmodifier "airbyte/internal/planmodifiers/boolplanmodifier"
	speakeasy_int64planmodifier "airbyte/internal/planmodifiers/int64planmodifier"
	speakeasy_listplanmodifier "airbyte/internal/planmodifiers/listplanmodifier"
	speakeasy_objectplanmodifier "airbyte/internal/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "airbyte/internal/planmodifiers/stringplanmodifier"
	"airbyte/internal/sdk/pkg/models/operations"
	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceFacebookMarketingResource{}
var _ resource.ResourceWithImportState = &SourceFacebookMarketingResource{}

func NewSourceFacebookMarketingResource() resource.Resource {
	return &SourceFacebookMarketingResource{}
}

// SourceFacebookMarketingResource defines the resource implementation.
type SourceFacebookMarketingResource struct {
	client *sdk.SDK
}

// SourceFacebookMarketingResourceModel describes the resource data model.
type SourceFacebookMarketingResourceModel struct {
	Configuration SourceFacebookMarketing `tfsdk:"configuration"`
	Name          types.String            `tfsdk:"name"`
	SecretID      types.String            `tfsdk:"secret_id"`
	SourceID      types.String            `tfsdk:"source_id"`
	WorkspaceID   types.String            `tfsdk:"workspace_id"`
}

func (r *SourceFacebookMarketingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_facebook_marketing"
}

func (r *SourceFacebookMarketingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceFacebookMarketing Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(),
				},
				Required: true,
				Attributes: map[string]schema.Attribute{
					"access_token": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(),
						},
						Required:    true,
						Description: `The value of the generated access token. From your App’s Dashboard, click on "Marketing API" then "Tools". Select permissions <b>ads_management, ads_read, read_insights, business_management</b>. Then click on "Get token". See the <a href="https://docs.airbyte.com/integrations/sources/facebook-marketing">docs</a> for more information.`,
					},
					"account_id": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(),
						},
						Required:    true,
						Description: `The Facebook Ad account ID to use when pulling data from the Facebook Marketing API. Open your Meta Ads Manager. The Ad account ID number is in the account dropdown menu or in your browser's address bar. See the <a href="https://www.facebook.com/business/help/1492627900875762">docs</a> for more information.`,
					},
					"action_breakdowns_allow_empty": schema.BoolAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Bool{
							speakeasy_boolplanmodifier.SuppressDiff(),
						},
						Optional:    true,
						Description: `Allows action_breakdowns to be an empty list`,
					},
					"client_id": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(),
						},
						Optional:    true,
						Description: `The Client Id for your OAuth app`,
					},
					"client_secret": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(),
						},
						Optional:    true,
						Description: `The Client Secret for your OAuth app`,
					},
					"custom_insights": schema.ListNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.List{
							speakeasy_listplanmodifier.SuppressDiff(),
						},
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action_breakdowns": schema.ListAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.List{
										speakeasy_listplanmodifier.SuppressDiff(),
									},
									Optional:    true,
									ElementType: types.StringType,
									Description: `A list of chosen action_breakdowns for action_breakdowns`,
								},
								"action_report_time": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										speakeasy_stringplanmodifier.SuppressDiff(),
									},
									Optional: true,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"conversion",
											"impression",
											"mixed",
										),
									},
									MarkdownDescription: `must be one of ["conversion", "impression", "mixed"]` + "\n" +
										`Determines the report time of action stats. For example, if a person saw the ad on Jan 1st but converted on Jan 2nd, when you query the API with action_report_time=impression, you see a conversion on Jan 1st. When you query the API with action_report_time=conversion, you see a conversion on Jan 2nd.`,
								},
								"breakdowns": schema.ListAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.List{
										speakeasy_listplanmodifier.SuppressDiff(),
									},
									Optional:    true,
									ElementType: types.StringType,
									Description: `A list of chosen breakdowns for breakdowns`,
								},
								"end_date": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										speakeasy_stringplanmodifier.SuppressDiff(),
									},
									Optional: true,
									Validators: []validator.String{
										validators.IsRFC3339(),
									},
									Description: `The date until which you'd like to replicate data for this stream, in the format YYYY-MM-DDT00:00:00Z. All data generated between the start date and this end date will be replicated. Not setting this option will result in always syncing the latest data.`,
								},
								"fields": schema.ListAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.List{
										speakeasy_listplanmodifier.SuppressDiff(),
									},
									Optional:    true,
									ElementType: types.StringType,
									Description: `A list of chosen fields for fields parameter`,
								},
								"insights_lookback_window": schema.Int64Attribute{
									Computed: true,
									PlanModifiers: []planmodifier.Int64{
										speakeasy_int64planmodifier.SuppressDiff(),
									},
									Optional:    true,
									Description: `The attribution window`,
								},
								"level": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										speakeasy_stringplanmodifier.SuppressDiff(),
									},
									Optional: true,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"ad",
											"adset",
											"campaign",
											"account",
										),
									},
									MarkdownDescription: `must be one of ["ad", "adset", "campaign", "account"]` + "\n" +
										`Chosen level for API`,
								},
								"name": schema.StringAttribute{
									PlanModifiers: []planmodifier.String{
										speakeasy_stringplanmodifier.SuppressDiff(),
									},
									Required:    true,
									Description: `The name value of insight`,
								},
								"start_date": schema.StringAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.String{
										speakeasy_stringplanmodifier.SuppressDiff(),
									},
									Optional: true,
									Validators: []validator.String{
										validators.IsRFC3339(),
									},
									Description: `The date from which you'd like to replicate data for this stream, in the format YYYY-MM-DDT00:00:00Z.`,
								},
								"time_increment": schema.Int64Attribute{
									Computed: true,
									PlanModifiers: []planmodifier.Int64{
										speakeasy_int64planmodifier.SuppressDiff(),
									},
									Optional:    true,
									Description: `Time window in days by which to aggregate statistics. The sync will be chunked into N day intervals, where N is the number of days you specified. For example, if you set this value to 7, then all statistics will be reported as 7-day aggregates by starting from the start_date. If the start and end dates are October 1st and October 30th, then the connector will output 5 records: 01 - 06, 07 - 13, 14 - 20, 21 - 27, and 28 - 30 (3 days only).`,
								},
							},
						},
						Description: `A list which contains ad statistics entries, each entry must have a name and can contains fields, breakdowns or action_breakdowns. Click on "add" to fill this field.`,
					},
					"end_date": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(),
						},
						Optional: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
						Description: `The date until which you'd like to replicate data for all incremental streams, in the format YYYY-MM-DDT00:00:00Z. All data generated between the start date and this end date will be replicated. Not setting this option will result in always syncing the latest data.`,
					},
					"fetch_thumbnail_images": schema.BoolAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Bool{
							speakeasy_boolplanmodifier.SuppressDiff(),
						},
						Optional:    true,
						Description: `Set to active if you want to fetch the thumbnail_url and store the result in thumbnail_data_url for each Ad Creative.`,
					},
					"include_deleted": schema.BoolAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Bool{
							speakeasy_boolplanmodifier.SuppressDiff(),
						},
						Optional:    true,
						Description: `Set to active if you want to include data from deleted Campaigns, Ads, and AdSets.`,
					},
					"insights_lookback_window": schema.Int64Attribute{
						Computed: true,
						PlanModifiers: []planmodifier.Int64{
							speakeasy_int64planmodifier.SuppressDiff(),
						},
						Optional:    true,
						Description: `The attribution window. Facebook freezes insight data 28 days after it was generated, which means that all data from the past 28 days may have changed since we last emitted it, so you can retrieve refreshed insights from the past by setting this parameter. If you set a custom lookback window value in Facebook account, please provide the same value here.`,
					},
					"max_batch_size": schema.Int64Attribute{
						Computed: true,
						PlanModifiers: []planmodifier.Int64{
							speakeasy_int64planmodifier.SuppressDiff(),
						},
						Optional:    true,
						Description: `Maximum batch size used when sending batch requests to Facebook API. Most users do not need to set this field unless they specifically need to tune the connector to address specific issues or use cases.`,
					},
					"page_size": schema.Int64Attribute{
						Computed: true,
						PlanModifiers: []planmodifier.Int64{
							speakeasy_int64planmodifier.SuppressDiff(),
						},
						Optional:    true,
						Description: `Page size used when sending requests to Facebook API to specify number of records per page when response has pagination. Most users do not need to set this field unless they specifically need to tune the connector to address specific issues or use cases.`,
					},
					"source_type": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(),
						},
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"facebook-marketing",
							),
						},
						Description: `must be one of ["facebook-marketing"]`,
					},
					"start_date": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(),
						},
						Required: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
						Description: `The date from which you'd like to replicate data for all incremental streams, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.`,
					},
				},
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Required: true,
			},
			"secret_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Optional:    true,
				Description: `Optional secretID obtained through the public API OAuth redirect flow.`,
			},
			"source_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
			},
			"workspace_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Required: true,
			},
		},
	}
}

func (r *SourceFacebookMarketingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceFacebookMarketingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceFacebookMarketingResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
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

	request := *data.ToCreateSDKType()
	res, err := r.client.Sources.CreateSourceFacebookMarketing(ctx, request)
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
	if res.SourceFacebookMarketingGetResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.SourceFacebookMarketingGetResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceFacebookMarketingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceFacebookMarketingResourceModel
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
	request := operations.GetSourceFacebookMarketingRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceFacebookMarketing(ctx, request)
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
	if res.SourceFacebookMarketingGetResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.SourceFacebookMarketingGetResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceFacebookMarketingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceFacebookMarketingResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceFacebookMarketingPutRequest := data.ToUpdateSDKType()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceFacebookMarketingRequest{
		SourceFacebookMarketingPutRequest: sourceFacebookMarketingPutRequest,
		SourceID:                          sourceID,
	}
	res, err := r.client.Sources.PutSourceFacebookMarketing(ctx, request)
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
	sourceId1 := data.SourceID.ValueString()
	getRequest := operations.GetSourceFacebookMarketingRequest{
		SourceID: sourceId1,
	}
	getResponse, err := r.client.Sources.GetSourceFacebookMarketing(ctx, getRequest)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if getResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", getResponse))
		return
	}
	if getResponse.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", getResponse.StatusCode), debugResponse(getResponse.RawResponse))
		return
	}
	if getResponse.SourceFacebookMarketingGetResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(getResponse.RawResponse))
		return
	}
	data.RefreshFromGetResponse(getResponse.SourceFacebookMarketingGetResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceFacebookMarketingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceFacebookMarketingResourceModel
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
	request := operations.DeleteSourceFacebookMarketingRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceFacebookMarketing(ctx, request)
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

func (r *SourceFacebookMarketingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("source_id"), req, resp)
}
