// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SourceShopifyDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceShopifyDataSource{}

func NewSourceShopifyDataSource() datasource.DataSource {
	return &SourceShopifyDataSource{}
}

// SourceShopifyDataSource is the data source implementation.
type SourceShopifyDataSource struct {
	client *sdk.SDK
}

// SourceShopifyDataSourceModel describes the data model.
type SourceShopifyDataSourceModel struct {
	Configuration SourceShopify `tfsdk:"configuration"`
	Name          types.String  `tfsdk:"name"`
	SecretID      types.String  `tfsdk:"secret_id"`
	SourceID      types.String  `tfsdk:"source_id"`
	WorkspaceID   types.String  `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceShopifyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_shopify"
}

// Schema defines the schema for the data source.
func (r *SourceShopifyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceShopify DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_shopify_shopify_authorization_method_api_password": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"api_password": schema.StringAttribute{
										Computed:    true,
										Description: `The API Password for your private application in the ` + "`" + `Shopify` + "`" + ` store.`,
									},
									"auth_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"api_password",
											),
										},
										Description: `must be one of ["api_password"]`,
									},
								},
								Description: `API Password Auth`,
							},
							"source_shopify_shopify_authorization_method_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `The Access Token for making authenticated requests.`,
									},
									"auth_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oauth2.0",
											),
										},
										Description: `must be one of ["oauth2.0"]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Client ID of the Shopify developer application.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The Client Secret of the Shopify developer application.`,
									},
								},
								Description: `OAuth2.0`,
							},
							"source_shopify_update_shopify_authorization_method_api_password": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"api_password": schema.StringAttribute{
										Computed:    true,
										Description: `The API Password for your private application in the ` + "`" + `Shopify` + "`" + ` store.`,
									},
									"auth_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"api_password",
											),
										},
										Description: `must be one of ["api_password"]`,
									},
								},
								Description: `API Password Auth`,
							},
							"source_shopify_update_shopify_authorization_method_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `The Access Token for making authenticated requests.`,
									},
									"auth_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oauth2.0",
											),
										},
										Description: `must be one of ["oauth2.0"]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Client ID of the Shopify developer application.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The Client Secret of the Shopify developer application.`,
									},
								},
								Description: `OAuth2.0`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `The authorization method to use to retrieve data from Shopify`,
					},
					"shop": schema.StringAttribute{
						Computed:    true,
						Description: `The name of your Shopify store found in the URL. For example, if your URL was https://NAME.myshopify.com, then the name would be 'NAME' or 'NAME.myshopify.com'.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"shopify",
							),
						},
						Description: `must be one of ["shopify"]`,
					},
					"start_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
						Description: `The date you would like to replicate data from. Format: YYYY-MM-DD. Any data before this date will not be replicated.`,
					},
				},
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"secret_id": schema.StringAttribute{
				Computed:    true,
				Description: `Optional secretID obtained through the public API OAuth redirect flow.`,
			},
			"source_id": schema.StringAttribute{
				Optional: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *SourceShopifyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceShopifyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceShopifyDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
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
	request := operations.GetSourceShopifyRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceShopify(ctx, request)
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
	if res.SourceShopifyGetResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.SourceShopifyGetResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
