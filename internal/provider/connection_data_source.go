// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/airbytehq/terraform-provider-airbyte/internal/provider/types"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ConnectionDataSource{}
var _ datasource.DataSourceWithConfigure = &ConnectionDataSource{}

func NewConnectionDataSource() datasource.DataSource {
	return &ConnectionDataSource{}
}

// ConnectionDataSource is the data source implementation.
type ConnectionDataSource struct {
	client *sdk.SDK
}

// ConnectionDataSourceModel describes the data model.
type ConnectionDataSourceModel struct {
	Configurations                   tfTypes.StreamConfigurations       `tfsdk:"configurations"`
	ConnectionID                     types.String                       `tfsdk:"connection_id"`
	DataResidency                    types.String                       `tfsdk:"data_residency"`
	DestinationID                    types.String                       `tfsdk:"destination_id"`
	Name                             types.String                       `tfsdk:"name"`
	NamespaceDefinition              types.String                       `tfsdk:"namespace_definition"`
	NamespaceFormat                  types.String                       `tfsdk:"namespace_format"`
	NonBreakingSchemaUpdatesBehavior types.String                       `tfsdk:"non_breaking_schema_updates_behavior"`
	Prefix                           types.String                       `tfsdk:"prefix"`
	Schedule                         tfTypes.ConnectionScheduleResponse `tfsdk:"schedule"`
	SourceID                         types.String                       `tfsdk:"source_id"`
	Status                           types.String                       `tfsdk:"status"`
	WorkspaceID                      types.String                       `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *ConnectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connection"
}

// Schema defines the schema for the data source.
func (r *ConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Connection DataSource",

		Attributes: map[string]schema.Attribute{
			"configurations": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"streams": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cursor_field": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `Path to the field that will be used to determine if a record is new or modified since the last sync. This field is REQUIRED if ` + "`" + `sync_mode` + "`" + ` is ` + "`" + `incremental` + "`" + ` unless there is a default.`,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"primary_key": schema.ListAttribute{
									Computed: true,
									ElementType: types.ListType{
										ElemType: types.StringType,
									},
									Description: `Paths to the fields that will be used as primary key. This field is REQUIRED if ` + "`" + `destination_sync_mode` + "`" + ` is ` + "`" + `*_dedup` + "`" + ` unless it is already supplied by the source schema.`,
								},
								"sync_mode": schema.StringAttribute{
									Computed:    true,
									Description: `must be one of ["full_refresh_overwrite", "full_refresh_append", "incremental_append", "incremental_deduped_history"]`,
								},
							},
						},
					},
				},
				Description: `A list of configured stream options for a connection.`,
			},
			"connection_id": schema.StringAttribute{
				Required: true,
			},
			"data_residency": schema.StringAttribute{
				Computed:    true,
				Description: `must be one of ["auto", "us", "eu"]`,
			},
			"destination_id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"namespace_definition": schema.StringAttribute{
				Computed:    true,
				Description: `Define the location where the data will be stored in the destination. must be one of ["source", "destination", "custom_format"]`,
			},
			"namespace_format": schema.StringAttribute{
				Computed: true,
			},
			"non_breaking_schema_updates_behavior": schema.StringAttribute{
				Computed:    true,
				Description: `Set how Airbyte handles syncs when it detects a non-breaking schema change in the source. must be one of ["ignore", "disable_connection", "propagate_columns", "propagate_fully"]`,
			},
			"prefix": schema.StringAttribute{
				Computed: true,
			},
			"schedule": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"basic_timing": schema.StringAttribute{
						Computed: true,
					},
					"cron_expression": schema.StringAttribute{
						Computed: true,
					},
					"schedule_type": schema.StringAttribute{
						Computed:    true,
						Description: `must be one of ["manual", "cron", "basic"]`,
					},
				},
				Description: `schedule for when the the connection should run, per the schedule type`,
			},
			"source_id": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: `must be one of ["active", "inactive", "deprecated"]`,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *ConnectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *ConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ConnectionDataSourceModel
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

	connectionID := data.ConnectionID.ValueString()
	request := operations.GetConnectionRequest{
		ConnectionID: connectionID,
	}
	res, err := r.client.PublicConnections.GetConnection(ctx, request)
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
	if res.ConnectionResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConnectionResponse(res.ConnectionResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
