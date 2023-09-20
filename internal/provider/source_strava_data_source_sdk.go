// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceStravaDataSourceModel) RefreshFromGetResponse(resp *shared.SourceStravaGetResponse) {
	r.Configuration.AthleteID = types.Int64Value(resp.Configuration.AthleteID)
	if resp.Configuration.AuthType != nil {
		r.Configuration.AuthType = types.StringValue(string(*resp.Configuration.AuthType))
	} else {
		r.Configuration.AuthType = types.StringNull()
	}
	r.Configuration.ClientID = types.StringValue(resp.Configuration.ClientID)
	r.Configuration.ClientSecret = types.StringValue(resp.Configuration.ClientSecret)
	r.Configuration.RefreshToken = types.StringValue(resp.Configuration.RefreshToken)
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
