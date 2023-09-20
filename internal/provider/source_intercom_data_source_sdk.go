// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceIntercomDataSourceModel) RefreshFromGetResponse(resp *shared.SourceIntercomGetResponse) {
	r.Configuration.AccessToken = types.StringValue(resp.Configuration.AccessToken)
	if resp.Configuration.ClientID != nil {
		r.Configuration.ClientID = types.StringValue(*resp.Configuration.ClientID)
	} else {
		r.Configuration.ClientID = types.StringNull()
	}
	if resp.Configuration.ClientSecret != nil {
		r.Configuration.ClientSecret = types.StringValue(*resp.Configuration.ClientSecret)
	} else {
		r.Configuration.ClientSecret = types.StringNull()
	}
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	r.Configuration.StartDate = types.StringValue(resp.Configuration.StartDate.Format(time.RFC3339))
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
