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
var _ resource.Resource = &SourceLinkedinPagesResource{}
var _ resource.ResourceWithImportState = &SourceLinkedinPagesResource{}

func NewSourceLinkedinPagesResource() resource.Resource {
	return &SourceLinkedinPagesResource{}
}

// SourceLinkedinPagesResource defines the resource implementation.
type SourceLinkedinPagesResource struct {
	client *sdk.SDK
}

// SourceLinkedinPagesResourceModel describes the resource data model.
type SourceLinkedinPagesResourceModel struct {
	Configuration SourceLinkedinPages `tfsdk:"configuration"`
	Name          types.String        `tfsdk:"name"`
	SecretID      types.String        `tfsdk:"secret_id"`
	SourceID      types.String        `tfsdk:"source_id"`
	WorkspaceID   types.String        `tfsdk:"workspace_id"`
}

func (r *SourceLinkedinPagesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_linkedin_pages"
}

func (r *SourceLinkedinPagesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceLinkedinPages Resource",

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
							"source_linkedin_pages_authentication_access_token": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									speakeasy_objectplanmodifier.SuppressDiff(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(),
										},
										Required:    true,
										Description: `The token value generated using the LinkedIn Developers OAuth Token Tools. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-pages/">docs</a> to obtain yours.`,
									},
									"auth_method": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(),
										},
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"access_token",
											),
										},
										Description: `must be one of ["access_token"]`,
									},
								},
							},
							"source_linkedin_pages_authentication_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									speakeasy_objectplanmodifier.SuppressDiff(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"auth_method": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(),
										},
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oAuth2.0",
											),
										},
										Description: `must be one of ["oAuth2.0"]`,
									},
									"client_id": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(),
										},
										Required:    true,
										Description: `The client ID of the LinkedIn developer application.`,
									},
									"client_secret": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(),
										},
										Required:    true,
										Description: `The client secret of the LinkedIn developer application.`,
									},
									"refresh_token": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(),
										},
										Required:    true,
										Description: `The token value generated using the LinkedIn Developers OAuth Token Tools. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-pages/">docs</a> to obtain yours.`,
									},
								},
							},
							"source_linkedin_pages_update_authentication_access_token": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Required:    true,
										Description: `The token value generated using the LinkedIn Developers OAuth Token Tools. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-pages/">docs</a> to obtain yours.`,
									},
									"auth_method": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"access_token",
											),
										},
										Description: `must be one of ["access_token"]`,
									},
								},
							},
							"source_linkedin_pages_update_authentication_o_auth2_0": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"auth_method": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oAuth2.0",
											),
										},
										Description: `must be one of ["oAuth2.0"]`,
									},
									"client_id": schema.StringAttribute{
										Required:    true,
										Description: `The client ID of the LinkedIn developer application.`,
									},
									"client_secret": schema.StringAttribute{
										Required:    true,
										Description: `The client secret of the LinkedIn developer application.`,
									},
									"refresh_token": schema.StringAttribute{
										Required:    true,
										Description: `The token value generated using the LinkedIn Developers OAuth Token Tools. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-pages/">docs</a> to obtain yours.`,
									},
								},
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"org_id": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(),
						},
						Required:    true,
						Description: `Specify the Organization ID`,
					},
					"source_type": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(),
						},
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"linkedin-pages",
							),
						},
						Description: `must be one of ["linkedin-pages"]`,
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

func (r *SourceLinkedinPagesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceLinkedinPagesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceLinkedinPagesResourceModel
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
	res, err := r.client.Sources.CreateSourceLinkedinPages(ctx, request)
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
	if res.SourceLinkedinPagesGetResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.SourceLinkedinPagesGetResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceLinkedinPagesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceLinkedinPagesResourceModel
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
	request := operations.GetSourceLinkedinPagesRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceLinkedinPages(ctx, request)
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
	if res.SourceLinkedinPagesGetResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.SourceLinkedinPagesGetResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceLinkedinPagesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceLinkedinPagesResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceLinkedinPagesPutRequest := data.ToUpdateSDKType()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceLinkedinPagesRequest{
		SourceLinkedinPagesPutRequest: sourceLinkedinPagesPutRequest,
		SourceID:                      sourceID,
	}
	res, err := r.client.Sources.PutSourceLinkedinPages(ctx, request)
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
	getRequest := operations.GetSourceLinkedinPagesRequest{
		SourceID: sourceId1,
	}
	getResponse, err := r.client.Sources.GetSourceLinkedinPages(ctx, getRequest)
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
	if getResponse.SourceLinkedinPagesGetResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(getResponse.RawResponse))
		return
	}
	data.RefreshFromGetResponse(getResponse.SourceLinkedinPagesGetResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceLinkedinPagesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceLinkedinPagesResourceModel
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
	request := operations.DeleteSourceLinkedinPagesRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceLinkedinPages(ctx, request)
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

func (r *SourceLinkedinPagesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("source_id"), req, resp)
}
