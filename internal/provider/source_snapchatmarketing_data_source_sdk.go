// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceSnapchatMarketingDataSourceModel) RefreshFromGetResponse(resp *shared.SourceSnapchatMarketingGetResponse) {
	r.Configuration.ClientID = types.StringValue(resp.Configuration.ClientID)
	r.Configuration.ClientSecret = types.StringValue(resp.Configuration.ClientSecret)
	if resp.Configuration.EndDate != nil {
		r.Configuration.EndDate = types.StringValue(resp.Configuration.EndDate.String())
	} else {
		r.Configuration.EndDate = types.StringNull()
	}
	r.Configuration.RefreshToken = types.StringValue(resp.Configuration.RefreshToken)
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	if resp.Configuration.StartDate != nil {
		r.Configuration.StartDate = types.StringValue(resp.Configuration.StartDate.String())
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
