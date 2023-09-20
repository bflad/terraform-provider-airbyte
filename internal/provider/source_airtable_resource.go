// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

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
var _ resource.Resource = &SourceAirtableResource{}
var _ resource.ResourceWithImportState = &SourceAirtableResource{}

func NewSourceAirtableResource() resource.Resource {
	return &SourceAirtableResource{}
}

// SourceAirtableResource defines the resource implementation.
type SourceAirtableResource struct {
	client *sdk.SDK
}

// SourceAirtableResourceModel describes the resource data model.
type SourceAirtableResourceModel struct {
	Configuration SourceAirtable `tfsdk:"configuration"`
	Name          types.String   `tfsdk:"name"`
	SecretID      types.String   `tfsdk:"secret_id"`
	SourceID      types.String   `tfsdk:"source_id"`
	WorkspaceID   types.String   `tfsdk:"workspace_id"`
}

func (r *SourceAirtableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_airtable"
}

func (r *SourceAirtableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceAirtable Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(),
				},
				Required: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							speakeasy_objectplanmodifier.SuppressDiff(),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source_airtable_authentication_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									speakeasy_objectplanmodifier.SuppressDiff(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(),
										},
										Optional:    true,
										Description: `Access Token for making authenticated requests.`,
									},
									"auth_method": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(),
										},
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oauth2.0",
											),
										},
										Description: `must be one of ["oauth2.0"]`,
									},
									"client_id": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(),
										},
										Required:    true,
										Description: `The client ID of the Airtable developer application.`,
									},
									"client_secret": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(),
										},
										Required:    true,
										Description: `The client secret the Airtable developer application.`,
									},
									"refresh_token": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(),
										},
										Required:    true,
										Description: `The key to refresh the expired access token.`,
									},
									"token_expiry_date": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(),
										},
										Optional: true,
										Validators: []validator.String{
											validators.IsRFC3339(),
										},
										Description: `The date-time when the access token should be refreshed.`,
									},
								},
							},
							"source_airtable_authentication_personal_access_token": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									speakeasy_objectplanmodifier.SuppressDiff(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"api_key": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(),
										},
										Required:    true,
										Description: `The Personal Access Token for the Airtable account. See the <a href="https://airtable.com/developers/web/guides/personal-access-tokens">Support Guide</a> for more information on how to obtain this token.`,
									},
									"auth_method": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(),
										},
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"api_key",
											),
										},
										Description: `must be one of ["api_key"]`,
									},
								},
							},
							"source_airtable_update_authentication_o_auth2_0": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Optional:    true,
										Description: `Access Token for making authenticated requests.`,
									},
									"auth_method": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oauth2.0",
											),
										},
										Description: `must be one of ["oauth2.0"]`,
									},
									"client_id": schema.StringAttribute{
										Required:    true,
										Description: `The client ID of the Airtable developer application.`,
									},
									"client_secret": schema.StringAttribute{
										Required:    true,
										Description: `The client secret the Airtable developer application.`,
									},
									"refresh_token": schema.StringAttribute{
										Required:    true,
										Description: `The key to refresh the expired access token.`,
									},
									"token_expiry_date": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsRFC3339(),
										},
										Description: `The date-time when the access token should be refreshed.`,
									},
								},
							},
							"source_airtable_update_authentication_personal_access_token": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"api_key": schema.StringAttribute{
										Required:    true,
										Description: `The Personal Access Token for the Airtable account. See the <a href="https://airtable.com/developers/web/guides/personal-access-tokens">Support Guide</a> for more information on how to obtain this token.`,
									},
									"auth_method": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"api_key",
											),
										},
										Description: `must be one of ["api_key"]`,
									},
								},
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(),
						},
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"airtable",
							),
						},
						Description: `must be one of ["airtable"]`,
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

func (r *SourceAirtableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceAirtableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceAirtableResourceModel
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
	res, err := r.client.Sources.CreateSourceAirtable(ctx, request)
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
	if res.SourceAirtableGetResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.SourceAirtableGetResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceAirtableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceAirtableResourceModel
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
	request := operations.GetSourceAirtableRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceAirtable(ctx, request)
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
	if res.SourceAirtableGetResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.SourceAirtableGetResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceAirtableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceAirtableResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceAirtablePutRequest := data.ToUpdateSDKType()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceAirtableRequest{
		SourceAirtablePutRequest: sourceAirtablePutRequest,
		SourceID:                 sourceID,
	}
	res, err := r.client.Sources.PutSourceAirtable(ctx, request)
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
	getRequest := operations.GetSourceAirtableRequest{
		SourceID: sourceId1,
	}
	getResponse, err := r.client.Sources.GetSourceAirtable(ctx, getRequest)
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
	if getResponse.SourceAirtableGetResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(getResponse.RawResponse))
		return
	}
	data.RefreshFromGetResponse(getResponse.SourceAirtableGetResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceAirtableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceAirtableResourceModel
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
	request := operations.DeleteSourceAirtableRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceAirtable(ctx, request)
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

func (r *SourceAirtableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("source_id"), req, resp)
}
