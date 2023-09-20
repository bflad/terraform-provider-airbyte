// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceMyHoursDataSourceModel) RefreshFromGetResponse(resp *shared.SourceMyHoursGetResponse) {
	r.Configuration.Email = types.StringValue(resp.Configuration.Email)
	if resp.Configuration.LogsBatchSize != nil {
		r.Configuration.LogsBatchSize = types.Int64Value(*resp.Configuration.LogsBatchSize)
	} else {
		r.Configuration.LogsBatchSize = types.Int64Null()
	}
	r.Configuration.Password = types.StringValue(resp.Configuration.Password)
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	r.Configuration.StartDate = types.StringValue(resp.Configuration.StartDate)
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
