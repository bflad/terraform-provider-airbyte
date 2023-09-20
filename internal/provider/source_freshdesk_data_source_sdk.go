// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceFreshdeskDataSourceModel) RefreshFromGetResponse(resp *shared.SourceFreshdeskGetResponse) {
	r.Configuration.APIKey = types.StringValue(resp.Configuration.APIKey)
	r.Configuration.Domain = types.StringValue(resp.Configuration.Domain)
	if resp.Configuration.RequestsPerMinute != nil {
		r.Configuration.RequestsPerMinute = types.Int64Value(*resp.Configuration.RequestsPerMinute)
	} else {
		r.Configuration.RequestsPerMinute = types.Int64Null()
	}
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	if resp.Configuration.StartDate != nil {
		r.Configuration.StartDate = types.StringValue(resp.Configuration.StartDate.Format(time.RFC3339))
	} else {
		r.Configuration.StartDate = types.StringNull()
	}
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
