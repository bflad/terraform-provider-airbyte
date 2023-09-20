// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceMailjetSmsDataSourceModel) RefreshFromGetResponse(resp *shared.SourceMailjetSmsGetResponse) {
	if resp.Configuration.EndDate != nil {
		r.Configuration.EndDate = types.Int64Value(*resp.Configuration.EndDate)
	} else {
		r.Configuration.EndDate = types.Int64Null()
	}
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	if resp.Configuration.StartDate != nil {
		r.Configuration.StartDate = types.Int64Value(*resp.Configuration.StartDate)
	} else {
		r.Configuration.StartDate = types.Int64Null()
	}
	r.Configuration.Token = types.StringValue(resp.Configuration.Token)
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
