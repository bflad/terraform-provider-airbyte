// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

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
) // Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceHubspotResource{}
var _ resource.ResourceWithImportState = &SourceHubspotResource{}

func NewSourceHubspotResource() resource.Resource {
	return &SourceHubspotResource{}
}

// SourceHubspotResource defines the resource implementation.
type SourceHubspotResource struct {
	client *sdk.SDK
}

// SourceHubspotResourceModel describes the resource data model.
type SourceHubspotResourceModel struct {
	Configuration SourceHubspot `tfsdk:"configuration"`
	Name          types.String  `tfsdk:"name"`
	SecretID      types.String  `tfsdk:"secret_id"`
	SourceID      types.String  `tfsdk:"source_id"`
	SourceType    types.String  `tfsdk:"source_type"`
	WorkspaceID   types.String  `tfsdk:"workspace_id"`
}

func (r *SourceHubspotResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_hubspot"
}

func (r *SourceHubspotResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceHubspot Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"source_hubspot_authentication_o_auth": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"client_id": schema.StringAttribute{
										Required:    true,
										Description: `The Client ID of your HubSpot developer application. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this ID.`,
									},
									"client_secret": schema.StringAttribute{
										Required:    true,
										Description: `The client secret for your HubSpot developer application. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this secret.`,
									},
									"credentials_title": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"OAuth Credentials",
											),
										},
										MarkdownDescription: `must be one of [OAuth Credentials]` + "\n" +
											`Name of the credentials`,
									},
									"refresh_token": schema.StringAttribute{
										Required:    true,
										Description: `Refresh token to renew an expired access token. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this token.`,
									},
								},
								Description: `Choose how to authenticate to HubSpot.`,
							},
							"source_hubspot_authentication_private_app": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Required:    true,
										Description: `HubSpot Access token. See the <a href="https://developers.hubspot.com/docs/api/private-apps">Hubspot docs</a> if you need help finding this token.`,
									},
									"credentials_title": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Private App Credentials",
											),
										},
										MarkdownDescription: `must be one of [Private App Credentials]` + "\n" +
											`Name of the credentials set`,
									},
								},
								Description: `Choose how to authenticate to HubSpot.`,
							},
							"source_hubspot_update_authentication_o_auth": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"client_id": schema.StringAttribute{
										Required:    true,
										Description: `The Client ID of your HubSpot developer application. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this ID.`,
									},
									"client_secret": schema.StringAttribute{
										Required:    true,
										Description: `The client secret for your HubSpot developer application. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this secret.`,
									},
									"credentials_title": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"OAuth Credentials",
											),
										},
										MarkdownDescription: `must be one of [OAuth Credentials]` + "\n" +
											`Name of the credentials`,
									},
									"refresh_token": schema.StringAttribute{
										Required:    true,
										Description: `Refresh token to renew an expired access token. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this token.`,
									},
								},
								Description: `Choose how to authenticate to HubSpot.`,
							},
							"source_hubspot_update_authentication_private_app": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Required:    true,
										Description: `HubSpot Access token. See the <a href="https://developers.hubspot.com/docs/api/private-apps">Hubspot docs</a> if you need help finding this token.`,
									},
									"credentials_title": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Private App Credentials",
											),
										},
										MarkdownDescription: `must be one of [Private App Credentials]` + "\n" +
											`Name of the credentials set`,
									},
								},
								Description: `Choose how to authenticate to HubSpot.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Choose how to authenticate to HubSpot.`,
					},
					"source_type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"hubspot",
							),
						},
						Description: `must be one of [hubspot]`,
					},
					"start_date": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
						Description: `UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.`,
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
				Optional:    true,
				Description: `Optional secretID obtained through the public API OAuth redirect flow.`,
			},
			"source_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
			},
			"source_type": schema.StringAttribute{
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

func (r *SourceHubspotResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceHubspotResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceHubspotResourceModel
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
	res, err := r.client.Sources.CreateSourceHubspot(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
	data.RefreshFromCreateResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceHubspotResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceHubspotResourceModel
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
	request := operations.GetSourceHubspotRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceHubspot(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
	data.RefreshFromGetResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceHubspotResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceHubspotResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceHubspotPutRequest := data.ToUpdateSDKType()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceHubspotRequest{
		SourceHubspotPutRequest: sourceHubspotPutRequest,
		SourceID:                sourceID,
	}
	res, err := r.client.Sources.PutSourceHubspot(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
	getRequest := operations.GetSourceHubspotRequest{
		SourceID: sourceId1,
	}
	getResponse, err := r.client.Sources.GetSourceHubspot(ctx, getRequest)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
	if getResponse.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(getResponse.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceHubspotResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceHubspotResourceModel
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
	request := operations.DeleteSourceHubspotRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceHubspot(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *SourceHubspotResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("source_id"), req, resp)
}
