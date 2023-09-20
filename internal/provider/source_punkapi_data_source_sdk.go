// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourcePunkAPIDataSourceModel) RefreshFromGetResponse(resp *shared.SourcePunkAPIGetResponse) {
	r.Configuration.BrewedAfter = types.StringValue(resp.Configuration.BrewedAfter)
	r.Configuration.BrewedBefore = types.StringValue(resp.Configuration.BrewedBefore)
	if resp.Configuration.ID != nil {
		r.Configuration.ID = types.StringValue(*resp.Configuration.ID)
	} else {
		r.Configuration.ID = types.StringNull()
	}
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
