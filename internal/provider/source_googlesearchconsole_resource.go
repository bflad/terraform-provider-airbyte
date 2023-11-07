// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk"

	speakeasy_stringplanmodifier "github.com/airbytehq/terraform-provider-airbyte/internal/planmodifiers/stringplanmodifier"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/operations"
	"github.com/airbytehq/terraform-provider-airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceGoogleSearchConsoleResource{}
var _ resource.ResourceWithImportState = &SourceGoogleSearchConsoleResource{}

func NewSourceGoogleSearchConsoleResource() resource.Resource {
	return &SourceGoogleSearchConsoleResource{}
}

// SourceGoogleSearchConsoleResource defines the resource implementation.
type SourceGoogleSearchConsoleResource struct {
	client *sdk.SDK
}

// SourceGoogleSearchConsoleResourceModel describes the resource data model.
type SourceGoogleSearchConsoleResourceModel struct {
	Configuration SourceGoogleSearchConsole `tfsdk:"configuration"`
	Name          types.String              `tfsdk:"name"`
	SecretID      types.String              `tfsdk:"secret_id"`
	SourceID      types.String              `tfsdk:"source_id"`
	SourceType    types.String              `tfsdk:"source_type"`
	WorkspaceID   types.String              `tfsdk:"workspace_id"`
}

func (r *SourceGoogleSearchConsoleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_google_search_console"
}

func (r *SourceGoogleSearchConsoleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceGoogleSearchConsole Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"authorization": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"o_auth": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Optional:    true,
										Description: `Access token for making authenticated requests. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.`,
									},
									"client_id": schema.StringAttribute{
										Required:    true,
										Description: `The client ID of your Google Search Console developer application. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.`,
									},
									"client_secret": schema.StringAttribute{
										Required:    true,
										Description: `The client secret of your Google Search Console developer application. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.`,
									},
									"refresh_token": schema.StringAttribute{
										Required:    true,
										Description: `The token for obtaining a new access token. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.`,
									},
								},
							},
							"service_account_key_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"email": schema.StringAttribute{
										Required:    true,
										Description: `The email of the user which has permissions to access the Google Workspace Admin APIs.`,
									},
									"service_account_info": schema.StringAttribute{
										Required:    true,
										Description: `The JSON key of the service account to use for authorization. Read more <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys">here</a>.`,
									},
								},
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"custom_reports": schema.StringAttribute{
						Optional:    true,
						Description: `(DEPRCATED) A JSON array describing the custom reports you want to sync from Google Search Console. See our <a href='https://docs.airbyte.com/integrations/sources/google-search-console'>documentation</a> for more information on formulating custom reports.`,
					},
					"custom_reports_array": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"dimensions": schema.ListAttribute{
									Required:    true,
									ElementType: types.StringType,
									Description: `A list of dimensions (country, date, device, page, query)`,
								},
								"name": schema.StringAttribute{
									Required:    true,
									Description: `The name of the custom report, this name would be used as stream name`,
								},
							},
						},
						Description: `You can add your Custom Analytics report by creating one.`,
					},
					"data_state": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"final",
								"all",
							),
						},
						MarkdownDescription: `must be one of ["final", "all"]; Default: "final"` + "\n" +
							`If set to 'final', the returned data will include only finalized, stable data. If set to 'all', fresh data will be included. When using Incremental sync mode, we do not recommend setting this parameter to 'all' as it may cause data loss. More information can be found in our <a href='https://docs.airbyte.com/integrations/source/google-search-console'>full documentation</a>.`,
					},
					"end_date": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
						Description: `UTC date in the format YYYY-MM-DD. Any data created after this date will not be replicated. Must be greater or equal to the start date field. Leaving this field blank will replicate all data from the start date onward.`,
					},
					"site_urls": schema.ListAttribute{
						Required:    true,
						ElementType: types.StringType,
						Description: `The URLs of the website property attached to your GSC account. Learn more about properties <a href="https://support.google.com/webmasters/answer/34592?hl=en">here</a>.`,
					},
					"start_date": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
						MarkdownDescription: `Default: "2021-01-01"` + "\n" +
							`UTC date in the format YYYY-MM-DD. Any data before this date will not be replicated.`,
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
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
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

func (r *SourceGoogleSearchConsoleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceGoogleSearchConsoleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceGoogleSearchConsoleResourceModel
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

	request := data.ToCreateSDKType()
	res, err := r.client.Sources.CreateSourceGoogleSearchConsole(ctx, request)
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
	data.RefreshFromCreateResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceGoogleSearchConsoleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceGoogleSearchConsoleResourceModel
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
	request := operations.GetSourceGoogleSearchConsoleRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceGoogleSearchConsole(ctx, request)
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
	data.RefreshFromGetResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceGoogleSearchConsoleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceGoogleSearchConsoleResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceGoogleSearchConsolePutRequest := data.ToUpdateSDKType()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceGoogleSearchConsoleRequest{
		SourceGoogleSearchConsolePutRequest: sourceGoogleSearchConsolePutRequest,
		SourceID:                            sourceID,
	}
	res, err := r.client.Sources.PutSourceGoogleSearchConsole(ctx, request)
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
	getRequest := operations.GetSourceGoogleSearchConsoleRequest{
		SourceID: sourceId1,
	}
	getResponse, err := r.client.Sources.GetSourceGoogleSearchConsole(ctx, getRequest)
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
	if getResponse.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(getResponse.RawResponse))
		return
	}
	data.RefreshFromGetResponse(getResponse.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceGoogleSearchConsoleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceGoogleSearchConsoleResourceModel
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
	request := operations.DeleteSourceGoogleSearchConsoleRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceGoogleSearchConsole(ctx, request)
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

func (r *SourceGoogleSearchConsoleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("source_id"), req.ID)...)
}
