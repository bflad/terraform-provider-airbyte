// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type DestinationBigqueryLoadingMethod struct {
	GCSStaging      *DestinationBigqueryGCSStaging                            `tfsdk:"gcs_staging"`
	StandardInserts *DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON `tfsdk:"standard_inserts"`
}
