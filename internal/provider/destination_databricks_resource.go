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
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DestinationDatabricksResource{}
var _ resource.ResourceWithImportState = &DestinationDatabricksResource{}

func NewDestinationDatabricksResource() resource.Resource {
	return &DestinationDatabricksResource{}
}

// DestinationDatabricksResource defines the resource implementation.
type DestinationDatabricksResource struct {
	client *sdk.SDK
}

// DestinationDatabricksResourceModel describes the resource data model.
type DestinationDatabricksResourceModel struct {
	Configuration   DestinationDatabricks `tfsdk:"configuration"`
	DestinationID   types.String          `tfsdk:"destination_id"`
	DestinationType types.String          `tfsdk:"destination_type"`
	Name            types.String          `tfsdk:"name"`
	WorkspaceID     types.String          `tfsdk:"workspace_id"`
}

func (r *DestinationDatabricksResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_databricks"
}

func (r *DestinationDatabricksResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationDatabricks Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"accept_terms": schema.BoolAttribute{
						Required:    true,
						Description: `You must agree to the Databricks JDBC Driver <a href="https://databricks.com/jdbc-odbc-driver-license">Terms & Conditions</a> to use this connector.`,
					},
					"data_source": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"destination_databricks_data_source_recommended_managed_tables": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"data_source_type": schema.StringAttribute{
										Required: true,
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
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"data_source_type": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"S3_STORAGE",
											),
										},
										Description: `must be one of ["S3_STORAGE"]`,
									},
									"file_name_pattern": schema.StringAttribute{
										Optional:    true,
										Description: `The pattern allows you to set the file-name format for the S3 staging file(s)`,
									},
									"s3_access_key_id": schema.StringAttribute{
										Required:    true,
										Description: `The Access Key Id granting allow one to access the above S3 staging bucket. Airbyte requires Read and Write permissions to the given bucket.`,
									},
									"s3_bucket_name": schema.StringAttribute{
										Required:    true,
										Description: `The name of the S3 bucket to use for intermittent staging of the data.`,
									},
									"s3_bucket_path": schema.StringAttribute{
										Required:    true,
										Description: `The directory under the S3 bucket where data will be written.`,
									},
									"s3_bucket_region": schema.StringAttribute{
										Required: true,
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
										Required:    true,
										Description: `The corresponding secret to the above access key id.`,
									},
								},
								Description: `Storage on which the delta lake is built.`,
							},
							"destination_databricks_data_source_azure_blob_storage": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"azure_blob_storage_account_name": schema.StringAttribute{
										Required:    true,
										Description: `The account's name of the Azure Blob Storage.`,
									},
									"azure_blob_storage_container_name": schema.StringAttribute{
										Required:    true,
										Description: `The name of the Azure blob storage container.`,
									},
									"azure_blob_storage_endpoint_domain_name": schema.StringAttribute{
										Optional:    true,
										Description: `This is Azure Blob Storage endpoint domain name. Leave default value (or leave it empty if run container from command line) to use Microsoft native from example.`,
									},
									"azure_blob_storage_sas_token": schema.StringAttribute{
										Required:    true,
										Description: `Shared access signature (SAS) token to grant limited access to objects in your storage account.`,
									},
									"data_source_type": schema.StringAttribute{
										Required: true,
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
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"data_source_type": schema.StringAttribute{
										Required: true,
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
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"data_source_type": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"S3_STORAGE",
											),
										},
										Description: `must be one of ["S3_STORAGE"]`,
									},
									"file_name_pattern": schema.StringAttribute{
										Optional:    true,
										Description: `The pattern allows you to set the file-name format for the S3 staging file(s)`,
									},
									"s3_access_key_id": schema.StringAttribute{
										Required:    true,
										Description: `The Access Key Id granting allow one to access the above S3 staging bucket. Airbyte requires Read and Write permissions to the given bucket.`,
									},
									"s3_bucket_name": schema.StringAttribute{
										Required:    true,
										Description: `The name of the S3 bucket to use for intermittent staging of the data.`,
									},
									"s3_bucket_path": schema.StringAttribute{
										Required:    true,
										Description: `The directory under the S3 bucket where data will be written.`,
									},
									"s3_bucket_region": schema.StringAttribute{
										Required: true,
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
										Required:    true,
										Description: `The corresponding secret to the above access key id.`,
									},
								},
								Description: `Storage on which the delta lake is built.`,
							},
							"destination_databricks_update_data_source_azure_blob_storage": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"azure_blob_storage_account_name": schema.StringAttribute{
										Required:    true,
										Description: `The account's name of the Azure Blob Storage.`,
									},
									"azure_blob_storage_container_name": schema.StringAttribute{
										Required:    true,
										Description: `The name of the Azure blob storage container.`,
									},
									"azure_blob_storage_endpoint_domain_name": schema.StringAttribute{
										Optional:    true,
										Description: `This is Azure Blob Storage endpoint domain name. Leave default value (or leave it empty if run container from command line) to use Microsoft native from example.`,
									},
									"azure_blob_storage_sas_token": schema.StringAttribute{
										Required:    true,
										Description: `Shared access signature (SAS) token to grant limited access to objects in your storage account.`,
									},
									"data_source_type": schema.StringAttribute{
										Required: true,
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
						Optional:    true,
						Description: `The name of the catalog. If not specified otherwise, the "hive_metastore" will be used.`,
					},
					"databricks_http_path": schema.StringAttribute{
						Required:    true,
						Description: `Databricks Cluster HTTP Path.`,
					},
					"databricks_personal_access_token": schema.StringAttribute{
						Required:    true,
						Description: `Databricks Personal Access Token for making authenticated requests.`,
					},
					"databricks_port": schema.StringAttribute{
						Optional:    true,
						Description: `Databricks Cluster Port.`,
					},
					"databricks_server_hostname": schema.StringAttribute{
						Required:    true,
						Description: `Databricks Cluster Server Hostname.`,
					},
					"destination_type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"databricks",
							),
						},
						Description: `must be one of ["databricks"]`,
					},
					"enable_schema_evolution": schema.BoolAttribute{
						Optional:    true,
						Description: `Support schema evolution for all streams. If "false", the connector might fail when a stream's schema changes.`,
					},
					"purge_staging_data": schema.BoolAttribute{
						Optional:    true,
						Description: `Default to 'true'. Switch it to 'false' for debugging purpose.`,
					},
					"schema": schema.StringAttribute{
						Optional:    true,
						Description: `The default schema tables are written. If not specified otherwise, the "default" will be used.`,
					},
				},
			},
			"destination_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
			},
			"destination_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Required: true,
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

func (r *DestinationDatabricksResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DestinationDatabricksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DestinationDatabricksResourceModel
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
	res, err := r.client.Destinations.CreateDestinationDatabricks(ctx, request)
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
	data.RefreshFromCreateResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationDatabricksResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DestinationDatabricksResourceModel
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

func (r *DestinationDatabricksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DestinationDatabricksResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	destinationDatabricksPutRequest := data.ToUpdateSDKType()
	destinationID := data.DestinationID.ValueString()
	request := operations.PutDestinationDatabricksRequest{
		DestinationDatabricksPutRequest: destinationDatabricksPutRequest,
		DestinationID:                   destinationID,
	}
	res, err := r.client.Destinations.PutDestinationDatabricks(ctx, request)
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
	destinationId1 := data.DestinationID.ValueString()
	getRequest := operations.GetDestinationDatabricksRequest{
		DestinationID: destinationId1,
	}
	getResponse, err := r.client.Destinations.GetDestinationDatabricks(ctx, getRequest)
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
	if getResponse.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(getResponse.RawResponse))
		return
	}
	data.RefreshFromGetResponse(getResponse.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationDatabricksResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DestinationDatabricksResourceModel
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

	destinationID := data.DestinationID.ValueString()
	request := operations.DeleteDestinationDatabricksRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.DeleteDestinationDatabricks(ctx, request)
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

func (r *DestinationDatabricksResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("destination_id"), req, resp)
}
