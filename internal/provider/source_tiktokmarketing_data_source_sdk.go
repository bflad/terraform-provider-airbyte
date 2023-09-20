// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceTiktokMarketingDataSourceModel) RefreshFromGetResponse(resp *shared.SourceTiktokMarketingGetResponse) {
	if resp.Configuration.AttributionWindow != nil {
		r.Configuration.AttributionWindow = types.Int64Value(*resp.Configuration.AttributionWindow)
	} else {
		r.Configuration.AttributionWindow = types.Int64Null()
	}
	if resp.Configuration.Credentials == nil {
		r.Configuration.Credentials = nil
	} else {
		r.Configuration.Credentials = &SourceTiktokMarketingAuthenticationMethod{}
		if resp.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20 != nil {
			r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20 = &SourceTiktokMarketingAuthenticationMethodOAuth20{}
			r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AccessToken = types.StringValue(resp.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AccessToken)
			if resp.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AdvertiserID != nil {
				r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AdvertiserID = types.StringValue(*resp.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AdvertiserID)
			} else {
				r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AdvertiserID = types.StringNull()
			}
			r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AppID = types.StringValue(resp.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AppID)
			if resp.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AuthType != nil {
				r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AuthType = types.StringValue(string(*resp.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AuthType))
			} else {
				r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AuthType = types.StringNull()
			}
			r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.Secret = types.StringValue(resp.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.Secret)
		}
		if resp.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken != nil {
			r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken = &SourceTiktokMarketingAuthenticationMethodSandboxAccessToken{}
			r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken.AccessToken = types.StringValue(resp.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken.AccessToken)
			r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken.AdvertiserID = types.StringValue(resp.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken.AdvertiserID)
			if resp.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken.AuthType != nil {
				r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken.AuthType = types.StringValue(string(*resp.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken.AuthType))
			} else {
				r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken.AuthType = types.StringNull()
			}
		}
		if resp.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodOAuth20 != nil {
			r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodOAuth20 = &SourceTiktokMarketingAuthenticationMethodOAuth20{}
		}
		if resp.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken != nil {
			r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken = &SourceTiktokMarketingAuthenticationMethodSandboxAccessToken{}
		}
	}
	if resp.Configuration.EndDate != nil {
		r.Configuration.EndDate = types.StringValue(resp.Configuration.EndDate.String())
	} else {
		r.Configuration.EndDate = types.StringNull()
	}
	if resp.Configuration.IncludeDeleted != nil {
		r.Configuration.IncludeDeleted = types.BoolValue(*resp.Configuration.IncludeDeleted)
	} else {
		r.Configuration.IncludeDeleted = types.BoolNull()
	}
	if resp.Configuration.SourceType != nil {
		r.Configuration.SourceType = types.StringValue(string(*resp.Configuration.SourceType))
	} else {
		r.Configuration.SourceType = types.StringNull()
	}
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
