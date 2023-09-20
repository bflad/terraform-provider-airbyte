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
var _ datasource.DataSource = &SourceChartmogulDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceChartmogulDataSource{}

func NewSourceChartmogulDataSource() datasource.DataSource {
	return &SourceChartmogulDataSource{}
}

// SourceChartmogulDataSource is the data source implementation.
type SourceChartmogulDataSource struct {
	client *sdk.SDK
}

// SourceChartmogulDataSourceModel describes the data model.
type SourceChartmogulDataSourceModel struct {
	Configuration SourceChartmogul `tfsdk:"configuration"`
	Name          types.String     `tfsdk:"name"`
	SecretID      types.String     `tfsdk:"secret_id"`
	SourceID      types.String     `tfsdk:"source_id"`
	WorkspaceID   types.String     `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceChartmogulDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_chartmogul"
}

// Schema defines the schema for the data source.
func (r *SourceChartmogulDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceChartmogul DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"api_key": schema.StringAttribute{
						Computed:    true,
						Description: `Your Chartmogul API key. See <a href="https://help.chartmogul.com/hc/en-us/articles/4407796325906-Creating-and-Managing-API-keys#creating-an-api-key"> the docs </a> for info on how to obtain this.`,
					},
					"interval": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"day",
								"week",
								"month",
								"quarter",
							),
						},
						MarkdownDescription: `must be one of ["day", "week", "month", "quarter"]` + "\n" +
							`Some APIs such as <a href="https://dev.chartmogul.com/reference/endpoint-overview-metrics-api">Metrics</a> require intervals to cluster data.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"chartmogul",
							),
						},
						Description: `must be one of ["chartmogul"]`,
					},
					"start_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
						Description: `UTC date and time in the format 2017-01-25T00:00:00Z. When feasible, any data before this date will not be replicated.`,
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

func (r *SourceChartmogulDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceChartmogulDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceChartmogulDataSourceModel
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
	request := operations.GetSourceChartmogulRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceChartmogul(ctx, request)
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
	if res.SourceChartmogulGetResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.SourceChartmogulGetResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
