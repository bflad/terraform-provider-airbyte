// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourcePipedriveDataSourceModel) RefreshFromGetResponse(resp *shared.SourcePipedriveGetResponse) {
	if r.Configuration.Authorization == nil {
		r.Configuration.Authorization = &SourcePipedriveAPIKeyAuthentication{}
	}
	if resp.Configuration.Authorization == nil {
		r.Configuration.Authorization = nil
	} else {
		r.Configuration.Authorization = &SourcePipedriveAPIKeyAuthentication{}
		r.Configuration.Authorization.APIToken = types.StringValue(resp.Configuration.Authorization.APIToken)
		r.Configuration.Authorization.AuthType = types.StringValue(string(resp.Configuration.Authorization.AuthType))
	}
	r.Configuration.ReplicationStartDate = types.StringValue(resp.Configuration.ReplicationStartDate.Format(time.RFC3339))
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	r.Name = types.StringValue(resp.Name)
	if resp.SecretID != nil {
		r.SecretID = types.StringValue(*resp.SecretID)
	} else {
		r.SecretID = types.StringNull()
	}
	if resp.SourceID != nil {
		r.SourceID = types.StringValue(*resp.SourceID)
	} else {
		r.SourceID = types.StringNull()
	}
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
