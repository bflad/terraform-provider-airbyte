// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationAzureBlobStorageDataSourceModel) RefreshFromGetResponse(resp *shared.DestinationAzureBlobStorageGetResponse) {
	r.Configuration.AzureBlobStorageAccountKey = types.StringValue(resp.Configuration.AzureBlobStorageAccountKey)
	r.Configuration.AzureBlobStorageAccountName = types.StringValue(resp.Configuration.AzureBlobStorageAccountName)
	if resp.Configuration.AzureBlobStorageContainerName != nil {
		r.Configuration.AzureBlobStorageContainerName = types.StringValue(*resp.Configuration.AzureBlobStorageContainerName)
	} else {
		r.Configuration.AzureBlobStorageContainerName = types.StringNull()
	}
	if resp.Configuration.AzureBlobStorageEndpointDomainName != nil {
		r.Configuration.AzureBlobStorageEndpointDomainName = types.StringValue(*resp.Configuration.AzureBlobStorageEndpointDomainName)
	} else {
		r.Configuration.AzureBlobStorageEndpointDomainName = types.StringNull()
	}
	if resp.Configuration.AzureBlobStorageOutputBufferSize != nil {
		r.Configuration.AzureBlobStorageOutputBufferSize = types.Int64Value(*resp.Configuration.AzureBlobStorageOutputBufferSize)
	} else {
		r.Configuration.AzureBlobStorageOutputBufferSize = types.Int64Null()
	}
	if resp.Configuration.AzureBlobStorageSpillSize != nil {
		r.Configuration.AzureBlobStorageSpillSize = types.Int64Value(*resp.Configuration.AzureBlobStorageSpillSize)
	} else {
		r.Configuration.AzureBlobStorageSpillSize = types.Int64Null()
	}
	r.Configuration.DestinationType = types.StringValue(string(resp.Configuration.DestinationType))
	if resp.Configuration.Format.DestinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues != nil {
		r.Configuration.Format.DestinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues = &DestinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues{}
		r.Configuration.Format.DestinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues.Flattening = types.StringValue(string(resp.Configuration.Format.DestinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues.Flattening))
		r.Configuration.Format.DestinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues.FormatType = types.StringValue(string(resp.Configuration.Format.DestinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues.FormatType))
	}
	if resp.Configuration.Format.DestinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON != nil {
		r.Configuration.Format.DestinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON = &DestinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON{}
		r.Configuration.Format.DestinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON.FormatType = types.StringValue(string(resp.Configuration.Format.DestinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON.FormatType))
	}
	if resp.Configuration.Format.DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues != nil {
		r.Configuration.Format.DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues = &DestinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues{}
	}
	if resp.Configuration.Format.DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON != nil {
		r.Configuration.Format.DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON = &DestinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON{}
	}
	if resp.DestinationID != nil {
		r.DestinationID = types.StringValue(*resp.DestinationID)
	} else {
		r.DestinationID = types.StringNull()
	}
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
