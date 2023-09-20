// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationBigqueryDenormalizedDataSourceModel) RefreshFromGetResponse(resp *shared.DestinationBigqueryDenormalizedGetResponse) {
	if resp.Configuration.BigQueryClientBufferSizeMb != nil {
		r.Configuration.BigQueryClientBufferSizeMb = types.Int64Value(*resp.Configuration.BigQueryClientBufferSizeMb)
	} else {
		r.Configuration.BigQueryClientBufferSizeMb = types.Int64Null()
	}
	if resp.Configuration.CredentialsJSON != nil {
		r.Configuration.CredentialsJSON = types.StringValue(*resp.Configuration.CredentialsJSON)
	} else {
		r.Configuration.CredentialsJSON = types.StringNull()
	}
	r.Configuration.DatasetID = types.StringValue(resp.Configuration.DatasetID)
	if resp.Configuration.DatasetLocation != nil {
		r.Configuration.DatasetLocation = types.StringValue(string(*resp.Configuration.DatasetLocation))
	} else {
		r.Configuration.DatasetLocation = types.StringNull()
	}
	r.Configuration.DestinationType = types.StringValue(string(resp.Configuration.DestinationType))
	if resp.Configuration.LoadingMethod == nil {
		r.Configuration.LoadingMethod = nil
	} else {
		r.Configuration.LoadingMethod = &DestinationBigqueryDenormalizedLoadingMethod{}
		if resp.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging != nil {
			r.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging = &DestinationBigqueryDenormalizedLoadingMethodGCSStaging{}
			if resp.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.Credential.DestinationBigqueryDenormalizedLoadingMethodGCSStagingCredentialHMACKey != nil {
				r.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.Credential.DestinationBigqueryDenormalizedLoadingMethodGCSStagingCredentialHMACKey = &DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey{}
				r.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.Credential.DestinationBigqueryDenormalizedLoadingMethodGCSStagingCredentialHMACKey.CredentialType = types.StringValue(string(resp.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.Credential.DestinationBigqueryDenormalizedLoadingMethodGCSStagingCredentialHMACKey.CredentialType))
				r.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.Credential.DestinationBigqueryDenormalizedLoadingMethodGCSStagingCredentialHMACKey.HmacKeyAccessID = types.StringValue(resp.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.Credential.DestinationBigqueryDenormalizedLoadingMethodGCSStagingCredentialHMACKey.HmacKeyAccessID)
				r.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.Credential.DestinationBigqueryDenormalizedLoadingMethodGCSStagingCredentialHMACKey.HmacKeySecret = types.StringValue(resp.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.Credential.DestinationBigqueryDenormalizedLoadingMethodGCSStagingCredentialHMACKey.HmacKeySecret)
			}
			if resp.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.FileBufferCount != nil {
				r.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.FileBufferCount = types.Int64Value(*resp.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.FileBufferCount)
			} else {
				r.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.FileBufferCount = types.Int64Null()
			}
			r.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.GcsBucketName = types.StringValue(resp.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.GcsBucketName)
			r.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.GcsBucketPath = types.StringValue(resp.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.GcsBucketPath)
			if resp.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.KeepFilesInGcsBucket != nil {
				r.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.KeepFilesInGcsBucket = types.StringValue(string(*resp.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.KeepFilesInGcsBucket))
			} else {
				r.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.KeepFilesInGcsBucket = types.StringNull()
			}
			r.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.Method = types.StringValue(string(resp.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodGCSStaging.Method))
		}
		if resp.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodStandardInserts != nil {
			r.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodStandardInserts = &DestinationBigqueryLoadingMethodStandardInserts{}
			r.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodStandardInserts.Method = types.StringValue(string(resp.Configuration.LoadingMethod.DestinationBigqueryDenormalizedLoadingMethodStandardInserts.Method))
		}
		if resp.Configuration.LoadingMethod.DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging != nil {
			r.Configuration.LoadingMethod.DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging = &DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging{}
		}
		if resp.Configuration.LoadingMethod.DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts != nil {
			r.Configuration.LoadingMethod.DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts = &DestinationBigqueryLoadingMethodStandardInserts{}
		}
	}
	r.Configuration.ProjectID = types.StringValue(resp.Configuration.ProjectID)
	if resp.DestinationID != nil {
		r.DestinationID = types.StringValue(*resp.DestinationID)
	} else {
		r.DestinationID = types.StringNull()
	}
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
