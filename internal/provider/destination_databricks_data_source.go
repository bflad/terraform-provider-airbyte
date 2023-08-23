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
var _ datasource.DataSource = &DestinationDatabricksDataSource{}
var _ datasource.DataSourceWithConfigure = &DestinationDatabricksDataSource{}

func NewDestinationDatabricksDataSource() datasource.DataSource {
	return &DestinationDatabricksDataSource{}
}

// DestinationDatabricksDataSource is the data source implementation.
type DestinationDatabricksDataSource struct {
	client *sdk.SDK
}

// DestinationDatabricksDataSourceModel describes the data model.
type DestinationDatabricksDataSourceModel struct {
	Configuration DestinationDatabricks1 `tfsdk:"configuration"`
	DestinationID types.String           `tfsdk:"destination_id"`
	Name          types.String           `tfsdk:"name"`
	WorkspaceID   types.String           `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *DestinationDatabricksDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_databricks"
}

// Schema defines the schema for the data source.
func (r *DestinationDatabricksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationDatabricks DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"accept_terms": schema.BoolAttribute{
						Computed:    true,
						Description: `You must agree to the Databricks JDBC Driver <a href="https://databricks.com/jdbc-odbc-driver-license">Terms & Conditions</a> to use this connector.`,
					},
					"data_source": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"destination_databricks_data_source_recommended_managed_tables": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"data_source_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"MANAGED_TABLES_STORAGE",
											),
										},
										Description: `must be one of ["MANAGED_TABLES_STORAGE"]`,
									},
								},
								Description: `Storage on which the delta lake is built.`,
							},
							"destination_databricks_data_source_amazon_s3": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"data_source_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"S3_STORAGE",
											),
										},
										Description: `must be one of ["S3_STORAGE"]`,
									},
									"file_name_pattern": schema.StringAttribute{
										Computed:    true,
										Description: `The pattern allows you to set the file-name format for the S3 staging file(s)`,
									},
									"s3_access_key_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Access Key Id granting allow one to access the above S3 staging bucket. Airbyte requires Read and Write permissions to the given bucket.`,
									},
									"s3_bucket_name": schema.StringAttribute{
										Computed:    true,
										Description: `The name of the S3 bucket to use for intermittent staging of the data.`,
									},
									"s3_bucket_path": schema.StringAttribute{
										Computed:    true,
										Description: `The directory under the S3 bucket where data will be written.`,
									},
									"s3_bucket_region": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"",
												"us-east-1",
												"us-east-2",
												"us-west-1",
												"us-west-2",
												"af-south-1",
												"ap-east-1",
												"ap-south-1",
												"ap-northeast-1",
												"ap-northeast-2",
												"ap-northeast-3",
												"ap-southeast-1",
												"ap-southeast-2",
												"ca-central-1",
												"cn-north-1",
												"cn-northwest-1",
												"eu-central-1",
												"eu-north-1",
												"eu-south-1",
												"eu-west-1",
												"eu-west-2",
												"eu-west-3",
												"sa-east-1",
												"me-south-1",
												"us-gov-east-1",
												"us-gov-west-1",
											),
										},
										MarkdownDescription: `must be one of ["", "us-east-1", "us-east-2", "us-west-1", "us-west-2", "af-south-1", "ap-east-1", "ap-south-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-southeast-1", "ap-southeast-2", "ca-central-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-north-1", "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3", "sa-east-1", "me-south-1", "us-gov-east-1", "us-gov-west-1"]` + "\n" +
											`The region of the S3 staging bucket to use if utilising a copy strategy.`,
									},
									"s3_secret_access_key": schema.StringAttribute{
										Computed:    true,
										Description: `The corresponding secret to the above access key id.`,
									},
								},
								Description: `Storage on which the delta lake is built.`,
							},
							"destination_databricks_data_source_azure_blob_storage": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"azure_blob_storage_account_name": schema.StringAttribute{
										Computed:    true,
										Description: `The account's name of the Azure Blob Storage.`,
									},
									"azure_blob_storage_container_name": schema.StringAttribute{
										Computed:    true,
										Description: `The name of the Azure blob storage container.`,
									},
									"azure_blob_storage_endpoint_domain_name": schema.StringAttribute{
										Computed:    true,
										Description: `This is Azure Blob Storage endpoint domain name. Leave default value (or leave it empty if run container from command line) to use Microsoft native from example.`,
									},
									"azure_blob_storage_sas_token": schema.StringAttribute{
										Computed:    true,
										Description: `Shared access signature (SAS) token to grant limited access to objects in your storage account.`,
									},
									"data_source_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"AZURE_BLOB_STORAGE",
											),
										},
										Description: `must be one of ["AZURE_BLOB_STORAGE"]`,
									},
								},
								Description: `Storage on which the delta lake is built.`,
							},
							"destination_databricks_update_data_source_recommended_managed_tables": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"data_source_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"MANAGED_TABLES_STORAGE",
											),
										},
										Description: `must be one of ["MANAGED_TABLES_STORAGE"]`,
									},
								},
								Description: `Storage on which the delta lake is built.`,
							},
							"destination_databricks_update_data_source_amazon_s3": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"data_source_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"S3_STORAGE",
											),
										},
										Description: `must be one of ["S3_STORAGE"]`,
									},
									"file_name_pattern": schema.StringAttribute{
										Computed:    true,
										Description: `The pattern allows you to set the file-name format for the S3 staging file(s)`,
									},
									"s3_access_key_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Access Key Id granting allow one to access the above S3 staging bucket. Airbyte requires Read and Write permissions to the given bucket.`,
									},
									"s3_bucket_name": schema.StringAttribute{
										Computed:    true,
										Description: `The name of the S3 bucket to use for intermittent staging of the data.`,
									},
									"s3_bucket_path": schema.StringAttribute{
										Computed:    true,
										Description: `The directory under the S3 bucket where data will be written.`,
									},
									"s3_bucket_region": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"",
												"us-east-1",
												"us-east-2",
												"us-west-1",
												"us-west-2",
												"af-south-1",
												"ap-east-1",
												"ap-south-1",
												"ap-northeast-1",
												"ap-northeast-2",
												"ap-northeast-3",
												"ap-southeast-1",
												"ap-southeast-2",
												"ca-central-1",
												"cn-north-1",
												"cn-northwest-1",
												"eu-central-1",
												"eu-north-1",
												"eu-south-1",
												"eu-west-1",
												"eu-west-2",
												"eu-west-3",
												"sa-east-1",
												"me-south-1",
												"us-gov-east-1",
												"us-gov-west-1",
											),
										},
										MarkdownDescription: `must be one of ["", "us-east-1", "us-east-2", "us-west-1", "us-west-2", "af-south-1", "ap-east-1", "ap-south-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-southeast-1", "ap-southeast-2", "ca-central-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-north-1", "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3", "sa-east-1", "me-south-1", "us-gov-east-1", "us-gov-west-1"]` + "\n" +
											`The region of the S3 staging bucket to use if utilising a copy strategy.`,
									},
									"s3_secret_access_key": schema.StringAttribute{
										Computed:    true,
										Description: `The corresponding secret to the above access key id.`,
									},
								},
								Description: `Storage on which the delta lake is built.`,
							},
							"destination_databricks_update_data_source_azure_blob_storage": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"azure_blob_storage_account_name": schema.StringAttribute{
										Computed:    true,
										Description: `The account's name of the Azure Blob Storage.`,
									},
									"azure_blob_storage_container_name": schema.StringAttribute{
										Computed:    true,
										Description: `The name of the Azure blob storage container.`,
									},
									"azure_blob_storage_endpoint_domain_name": schema.StringAttribute{
										Computed:    true,
										Description: `This is Azure Blob Storage endpoint domain name. Leave default value (or leave it empty if run container from command line) to use Microsoft native from example.`,
									},
									"azure_blob_storage_sas_token": schema.StringAttribute{
										Computed:    true,
										Description: `Shared access signature (SAS) token to grant limited access to objects in your storage account.`,
									},
									"data_source_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"AZURE_BLOB_STORAGE",
											),
										},
										Description: `must be one of ["AZURE_BLOB_STORAGE"]`,
									},
								},
								Description: `Storage on which the delta lake is built.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Storage on which the delta lake is built.`,
					},
					"database": schema.StringAttribute{
						Computed:    true,
						Description: `The name of the catalog. If not specified otherwise, the "hive_metastore" will be used.`,
					},
					"databricks_http_path": schema.StringAttribute{
						Computed:    true,
						Description: `Databricks Cluster HTTP Path.`,
					},
					"databricks_personal_access_token": schema.StringAttribute{
						Computed:    true,
						Description: `Databricks Personal Access Token for making authenticated requests.`,
					},
					"databricks_port": schema.StringAttribute{
						Computed:    true,
						Description: `Databricks Cluster Port.`,
					},
					"databricks_server_hostname": schema.StringAttribute{
						Computed:    true,
						Description: `Databricks Cluster Server Hostname.`,
					},
					"destination_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"databricks",
							),
						},
						Description: `must be one of ["databricks"]`,
					},
					"enable_schema_evolution": schema.BoolAttribute{
						Computed:    true,
						Description: `Support schema evolution for all streams. If "false", the connector might fail when a stream's schema changes.`,
					},
					"purge_staging_data": schema.BoolAttribute{
						Computed:    true,
						Description: `Default to 'true'. Switch it to 'false' for debugging purpose.`,
					},
					"schema": schema.StringAttribute{
						Computed:    true,
						Description: `The default schema tables are written. If not specified otherwise, the "default" will be used.`,
					},
				},
			},
			"destination_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *DestinationDatabricksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *DestinationDatabricksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *DestinationDatabricksDataSourceModel
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

	destinationID := data.DestinationID.ValueString()
	request := operations.GetDestinationDatabricksRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationDatabricks(ctx, request)
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
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
