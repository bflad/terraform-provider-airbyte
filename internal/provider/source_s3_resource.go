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
var _ resource.Resource = &SourceS3Resource{}
var _ resource.ResourceWithImportState = &SourceS3Resource{}

func NewSourceS3Resource() resource.Resource {
	return &SourceS3Resource{}
}

// SourceS3Resource defines the resource implementation.
type SourceS3Resource struct {
	client *sdk.SDK
}

// SourceS3ResourceModel describes the resource data model.
type SourceS3ResourceModel struct {
	Configuration SourceS3     `tfsdk:"configuration"`
	Name          types.String `tfsdk:"name"`
	SecretID      types.String `tfsdk:"secret_id"`
	SourceID      types.String `tfsdk:"source_id"`
	SourceType    types.String `tfsdk:"source_type"`
	WorkspaceID   types.String `tfsdk:"workspace_id"`
}

func (r *SourceS3Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_s3"
}

func (r *SourceS3Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceS3 Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"dataset": schema.StringAttribute{
						Required:    true,
						Description: `The name of the stream you would like this source to output. Can contain letters, numbers, or underscores.`,
					},
					"format": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source_s3_file_format_avro": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"filetype": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"avro",
											),
										},
										Description: `must be one of [avro]`,
									},
								},
								Description: `This connector utilises <a href="https://fastavro.readthedocs.io/en/latest/" target="_blank">fastavro</a> for Avro parsing.`,
							},
							"source_s3_file_format_csv": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_reader_options": schema.StringAttribute{
										Optional:    true,
										Description: `Optionally add a valid JSON string here to provide additional options to the csv reader. Mappings must correspond to options <a href="https://arrow.apache.org/docs/python/generated/pyarrow.csv.ConvertOptions.html#pyarrow.csv.ConvertOptions" target="_blank">detailed here</a>. 'column_types' is used internally to handle schema so overriding that would likely cause problems.`,
									},
									"advanced_options": schema.StringAttribute{
										Optional:    true,
										Description: `Optionally add a valid JSON string here to provide additional <a href="https://arrow.apache.org/docs/python/generated/pyarrow.csv.ReadOptions.html#pyarrow.csv.ReadOptions" target="_blank">Pyarrow ReadOptions</a>. Specify 'column_names' here if your CSV doesn't have header, or if you want to use custom column names. 'block_size' and 'encoding' are already used above, specify them again here will override the values above.`,
									},
									"block_size": schema.Int64Attribute{
										Optional:    true,
										Description: `The chunk size in bytes to process at a time in memory from each file. If your data is particularly wide and failing during schema detection, increasing this should solve it. Beware of raising this too high as you could hit OOM errors.`,
									},
									"delimiter": schema.StringAttribute{
										Optional:    true,
										Description: `The character delimiting individual cells in the CSV data. This may only be a 1-character string. For tab-delimited data enter '\t'.`,
									},
									"double_quote": schema.BoolAttribute{
										Optional:    true,
										Description: `Whether two quotes in a quoted CSV value denote a single quote in the data.`,
									},
									"encoding": schema.StringAttribute{
										Optional:    true,
										Description: `The character encoding of the CSV data. Leave blank to default to <strong>UTF8</strong>. See <a href="https://docs.python.org/3/library/codecs.html#standard-encodings" target="_blank">list of python encodings</a> for allowable options.`,
									},
									"escape_char": schema.StringAttribute{
										Optional:    true,
										Description: `The character used for escaping special characters. To disallow escaping, leave this field blank.`,
									},
									"filetype": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"csv",
											),
										},
										Description: `must be one of [csv]`,
									},
									"infer_datatypes": schema.BoolAttribute{
										Optional:    true,
										Description: `Configures whether a schema for the source should be inferred from the current data or not. If set to false and a custom schema is set, then the manually enforced schema is used. If a schema is not manually set, and this is set to false, then all fields will be read as strings`,
									},
									"newlines_in_values": schema.BoolAttribute{
										Optional:    true,
										Description: `Whether newline characters are allowed in CSV values. Turning this on may affect performance. Leave blank to default to False.`,
									},
									"quote_char": schema.StringAttribute{
										Optional:    true,
										Description: `The character used for quoting CSV values. To disallow quoting, make this field blank.`,
									},
								},
								Description: `This connector utilises <a href="https: // arrow.apache.org/docs/python/generated/pyarrow.csv.open_csv.html" target="_blank">PyArrow (Apache Arrow)</a> for CSV parsing.`,
							},
							"source_s3_file_format_jsonl": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"block_size": schema.Int64Attribute{
										Optional:    true,
										Description: `The chunk size in bytes to process at a time in memory from each file. If your data is particularly wide and failing during schema detection, increasing this should solve it. Beware of raising this too high as you could hit OOM errors.`,
									},
									"filetype": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"jsonl",
											),
										},
										Description: `must be one of [jsonl]`,
									},
									"newlines_in_values": schema.BoolAttribute{
										Optional:    true,
										Description: `Whether newline characters are allowed in JSON values. Turning this on may affect performance. Leave blank to default to False.`,
									},
									"unexpected_field_behavior": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"ignore",
												"infer",
												"error",
											),
										},
										MarkdownDescription: `must be one of [ignore, infer, error]` + "\n" +
											`How JSON fields outside of explicit_schema (if given) are treated. Check <a href="https://arrow.apache.org/docs/python/generated/pyarrow.json.ParseOptions.html" target="_blank">PyArrow documentation</a> for details`,
									},
								},
								Description: `This connector uses <a href="https://arrow.apache.org/docs/python/json.html" target="_blank">PyArrow</a> for JSON Lines (jsonl) file parsing.`,
							},
							"source_s3_file_format_parquet": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"batch_size": schema.Int64Attribute{
										Optional:    true,
										Description: `Maximum number of records per batch read from the input files. Batches may be smaller if there aren’t enough rows in the file. This option can help avoid out-of-memory errors if your data is particularly wide.`,
									},
									"buffer_size": schema.Int64Attribute{
										Optional:    true,
										Description: `Perform read buffering when deserializing individual column chunks. By default every group column will be loaded fully to memory. This option can help avoid out-of-memory errors if your data is particularly wide.`,
									},
									"columns": schema.ListAttribute{
										Optional:    true,
										ElementType: types.StringType,
										Description: `If you only want to sync a subset of the columns from the file(s), add the columns you want here as a comma-delimited list. Leave it empty to sync all columns.`,
									},
									"filetype": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"parquet",
											),
										},
										Description: `must be one of [parquet]`,
									},
								},
								Description: `This connector utilises <a href="https://arrow.apache.org/docs/python/generated/pyarrow.parquet.ParquetFile.html" target="_blank">PyArrow (Apache Arrow)</a> for Parquet parsing.`,
							},
							"source_s3_update_file_format_avro": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"filetype": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"avro",
											),
										},
										Description: `must be one of [avro]`,
									},
								},
								Description: `This connector utilises <a href="https://fastavro.readthedocs.io/en/latest/" target="_blank">fastavro</a> for Avro parsing.`,
							},
							"source_s3_update_file_format_csv": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_reader_options": schema.StringAttribute{
										Optional:    true,
										Description: `Optionally add a valid JSON string here to provide additional options to the csv reader. Mappings must correspond to options <a href="https://arrow.apache.org/docs/python/generated/pyarrow.csv.ConvertOptions.html#pyarrow.csv.ConvertOptions" target="_blank">detailed here</a>. 'column_types' is used internally to handle schema so overriding that would likely cause problems.`,
									},
									"advanced_options": schema.StringAttribute{
										Optional:    true,
										Description: `Optionally add a valid JSON string here to provide additional <a href="https://arrow.apache.org/docs/python/generated/pyarrow.csv.ReadOptions.html#pyarrow.csv.ReadOptions" target="_blank">Pyarrow ReadOptions</a>. Specify 'column_names' here if your CSV doesn't have header, or if you want to use custom column names. 'block_size' and 'encoding' are already used above, specify them again here will override the values above.`,
									},
									"block_size": schema.Int64Attribute{
										Optional:    true,
										Description: `The chunk size in bytes to process at a time in memory from each file. If your data is particularly wide and failing during schema detection, increasing this should solve it. Beware of raising this too high as you could hit OOM errors.`,
									},
									"delimiter": schema.StringAttribute{
										Optional:    true,
										Description: `The character delimiting individual cells in the CSV data. This may only be a 1-character string. For tab-delimited data enter '\t'.`,
									},
									"double_quote": schema.BoolAttribute{
										Optional:    true,
										Description: `Whether two quotes in a quoted CSV value denote a single quote in the data.`,
									},
									"encoding": schema.StringAttribute{
										Optional:    true,
										Description: `The character encoding of the CSV data. Leave blank to default to <strong>UTF8</strong>. See <a href="https://docs.python.org/3/library/codecs.html#standard-encodings" target="_blank">list of python encodings</a> for allowable options.`,
									},
									"escape_char": schema.StringAttribute{
										Optional:    true,
										Description: `The character used for escaping special characters. To disallow escaping, leave this field blank.`,
									},
									"filetype": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"csv",
											),
										},
										Description: `must be one of [csv]`,
									},
									"infer_datatypes": schema.BoolAttribute{
										Optional:    true,
										Description: `Configures whether a schema for the source should be inferred from the current data or not. If set to false and a custom schema is set, then the manually enforced schema is used. If a schema is not manually set, and this is set to false, then all fields will be read as strings`,
									},
									"newlines_in_values": schema.BoolAttribute{
										Optional:    true,
										Description: `Whether newline characters are allowed in CSV values. Turning this on may affect performance. Leave blank to default to False.`,
									},
									"quote_char": schema.StringAttribute{
										Optional:    true,
										Description: `The character used for quoting CSV values. To disallow quoting, make this field blank.`,
									},
								},
								Description: `This connector utilises <a href="https: // arrow.apache.org/docs/python/generated/pyarrow.csv.open_csv.html" target="_blank">PyArrow (Apache Arrow)</a> for CSV parsing.`,
							},
							"source_s3_update_file_format_jsonl": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"block_size": schema.Int64Attribute{
										Optional:    true,
										Description: `The chunk size in bytes to process at a time in memory from each file. If your data is particularly wide and failing during schema detection, increasing this should solve it. Beware of raising this too high as you could hit OOM errors.`,
									},
									"filetype": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"jsonl",
											),
										},
										Description: `must be one of [jsonl]`,
									},
									"newlines_in_values": schema.BoolAttribute{
										Optional:    true,
										Description: `Whether newline characters are allowed in JSON values. Turning this on may affect performance. Leave blank to default to False.`,
									},
									"unexpected_field_behavior": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"ignore",
												"infer",
												"error",
											),
										},
										MarkdownDescription: `must be one of [ignore, infer, error]` + "\n" +
											`How JSON fields outside of explicit_schema (if given) are treated. Check <a href="https://arrow.apache.org/docs/python/generated/pyarrow.json.ParseOptions.html" target="_blank">PyArrow documentation</a> for details`,
									},
								},
								Description: `This connector uses <a href="https://arrow.apache.org/docs/python/json.html" target="_blank">PyArrow</a> for JSON Lines (jsonl) file parsing.`,
							},
							"source_s3_update_file_format_parquet": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"batch_size": schema.Int64Attribute{
										Optional:    true,
										Description: `Maximum number of records per batch read from the input files. Batches may be smaller if there aren’t enough rows in the file. This option can help avoid out-of-memory errors if your data is particularly wide.`,
									},
									"buffer_size": schema.Int64Attribute{
										Optional:    true,
										Description: `Perform read buffering when deserializing individual column chunks. By default every group column will be loaded fully to memory. This option can help avoid out-of-memory errors if your data is particularly wide.`,
									},
									"columns": schema.ListAttribute{
										Optional:    true,
										ElementType: types.StringType,
										Description: `If you only want to sync a subset of the columns from the file(s), add the columns you want here as a comma-delimited list. Leave it empty to sync all columns.`,
									},
									"filetype": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"parquet",
											),
										},
										Description: `must be one of [parquet]`,
									},
								},
								Description: `This connector utilises <a href="https://arrow.apache.org/docs/python/generated/pyarrow.parquet.ParquetFile.html" target="_blank">PyArrow (Apache Arrow)</a> for Parquet parsing.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `The format of the files you'd like to replicate`,
					},
					"path_pattern": schema.StringAttribute{
						Required:    true,
						Description: `A regular expression which tells the connector which files to replicate. All files which match this pattern will be replicated. Use | to separate multiple patterns. See <a href="https://facelessuser.github.io/wcmatch/glob/" target="_blank">this page</a> to understand pattern syntax (GLOBSTAR and SPLIT flags are enabled). Use pattern <strong>**</strong> to pick up all files.`,
					},
					"provider": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"aws_access_key_id": schema.StringAttribute{
								Optional:    true,
								Description: `In order to access private Buckets stored on AWS S3, this connector requires credentials with the proper permissions. If accessing publicly available data, this field is not necessary.`,
							},
							"aws_secret_access_key": schema.StringAttribute{
								Optional:    true,
								Description: `In order to access private Buckets stored on AWS S3, this connector requires credentials with the proper permissions. If accessing publicly available data, this field is not necessary.`,
							},
							"bucket": schema.StringAttribute{
								Required:    true,
								Description: `Name of the S3 bucket where the file(s) exist.`,
							},
							"endpoint": schema.StringAttribute{
								Optional:    true,
								Description: `Endpoint to an S3 compatible service. Leave empty to use AWS.`,
							},
							"path_prefix": schema.StringAttribute{
								Optional:    true,
								Description: `By providing a path-like prefix (e.g. myFolder/thisTable/) under which all the relevant files sit, we can optimize finding these in S3. This is optional but recommended if your bucket contains many folders/files which you don't need to replicate.`,
							},
							"start_date": schema.StringAttribute{
								Optional: true,
								Validators: []validator.String{
									validators.IsRFC3339(),
								},
								Description: `UTC date and time in the format 2017-01-25T00:00:00Z. Any file modified before this date will not be replicated.`,
							},
						},
						Description: `Use this to load files from S3 or S3-compatible services`,
					},
					"schema": schema.StringAttribute{
						Optional:    true,
						Description: `Optionally provide a schema to enforce, as a valid JSON string. Ensure this is a mapping of <strong>{ "column" : "type" }</strong>, where types are valid <a href="https://json-schema.org/understanding-json-schema/reference/type.html" target="_blank">JSON Schema datatypes</a>. Leave as {} to auto-infer the schema.`,
					},
					"source_type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"s3",
							),
						},
						Description: `must be one of [s3]`,
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

func (r *SourceS3Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceS3Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceS3ResourceModel
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
	res, err := r.client.Sources.CreateSourceS3(ctx, request)
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

func (r *SourceS3Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceS3ResourceModel
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
	request := operations.GetSourceS3Request{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceS3(ctx, request)
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

func (r *SourceS3Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceS3ResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceS3PutRequest := data.ToUpdateSDKType()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceS3Request{
		SourceS3PutRequest: sourceS3PutRequest,
		SourceID:           sourceID,
	}
	res, err := r.client.Sources.PutSourceS3(ctx, request)
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
	getRequest := operations.GetSourceS3Request{
		SourceID: sourceId1,
	}
	getResponse, err := r.client.Sources.GetSourceS3(ctx, getRequest)
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

func (r *SourceS3Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceS3ResourceModel
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
	request := operations.DeleteSourceS3Request{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceS3(ctx, request)
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

func (r *SourceS3Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("source_id"), req, resp)
}
