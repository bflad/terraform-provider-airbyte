// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceBingAdsDataSourceModel) RefreshFromGetResponse(resp *shared.SourceBingAdsGetResponse) {
	if resp.Configuration.AuthMethod != nil {
		r.Configuration.AuthMethod = types.StringValue(string(*resp.Configuration.AuthMethod))
	} else {
		r.Configuration.AuthMethod = types.StringNull()
	}
	r.Configuration.ClientID = types.StringValue(resp.Configuration.ClientID)
	if resp.Configuration.ClientSecret != nil {
		r.Configuration.ClientSecret = types.StringValue(*resp.Configuration.ClientSecret)
	} else {
		r.Configuration.ClientSecret = types.StringNull()
	}
	r.Configuration.DeveloperToken = types.StringValue(resp.Configuration.DeveloperToken)
	if resp.Configuration.LookbackWindow != nil {
		r.Configuration.LookbackWindow = types.Int64Value(*resp.Configuration.LookbackWindow)
	} else {
		r.Configuration.LookbackWindow = types.Int64Null()
	}
	r.Configuration.RefreshToken = types.StringValue(resp.Configuration.RefreshToken)
	r.Configuration.ReportsStartDate = types.StringValue(resp.Configuration.ReportsStartDate.String())
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	if resp.Configuration.TenantID != nil {
		r.Configuration.TenantID = types.StringValue(*resp.Configuration.TenantID)
	} else {
		r.Configuration.TenantID = types.StringNull()
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
