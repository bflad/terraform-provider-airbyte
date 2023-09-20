// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceAirtableDataSourceModel) RefreshFromGetResponse(resp *shared.SourceAirtableGetResponse) {
	if resp.Configuration.Credentials == nil {
		r.Configuration.Credentials = nil
	} else {
		r.Configuration.Credentials = &SourceAirtableAuthentication{}
		if resp.Configuration.Credentials.SourceAirtableAuthenticationOAuth20 != nil {
			r.Configuration.Credentials.SourceAirtableAuthenticationOAuth20 = &SourceAirtableAuthenticationOAuth20{}
			if resp.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.AccessToken != nil {
				r.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.AccessToken = types.StringValue(*resp.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.AccessToken)
			} else {
				r.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.AccessToken = types.StringNull()
			}
			if resp.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.AuthMethod != nil {
				r.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.AuthMethod = types.StringValue(string(*resp.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.AuthMethod))
			} else {
				r.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.AuthMethod = types.StringNull()
			}
			r.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.ClientID = types.StringValue(resp.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.ClientID)
			r.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.ClientSecret = types.StringValue(resp.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.ClientSecret)
			r.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.RefreshToken = types.StringValue(resp.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.RefreshToken)
			if resp.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.TokenExpiryDate != nil {
				r.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.TokenExpiryDate = types.StringValue(resp.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.TokenExpiryDate.Format(time.RFC3339))
			} else {
				r.Configuration.Credentials.SourceAirtableAuthenticationOAuth20.TokenExpiryDate = types.StringNull()
			}
		}
		if resp.Configuration.Credentials.SourceAirtableAuthenticationPersonalAccessToken != nil {
			r.Configuration.Credentials.SourceAirtableAuthenticationPersonalAccessToken = &SourceAirtableAuthenticationPersonalAccessToken{}
			r.Configuration.Credentials.SourceAirtableAuthenticationPersonalAccessToken.APIKey = types.StringValue(resp.Configuration.Credentials.SourceAirtableAuthenticationPersonalAccessToken.APIKey)
			if resp.Configuration.Credentials.SourceAirtableAuthenticationPersonalAccessToken.AuthMethod != nil {
				r.Configuration.Credentials.SourceAirtableAuthenticationPersonalAccessToken.AuthMethod = types.StringValue(string(*resp.Configuration.Credentials.SourceAirtableAuthenticationPersonalAccessToken.AuthMethod))
			} else {
				r.Configuration.Credentials.SourceAirtableAuthenticationPersonalAccessToken.AuthMethod = types.StringNull()
			}
		}
		if resp.Configuration.Credentials.SourceAirtableUpdateAuthenticationOAuth20 != nil {
			r.Configuration.Credentials.SourceAirtableUpdateAuthenticationOAuth20 = &SourceAirtableAuthenticationOAuth20{}
		}
		if resp.Configuration.Credentials.SourceAirtableUpdateAuthenticationPersonalAccessToken != nil {
			r.Configuration.Credentials.SourceAirtableUpdateAuthenticationPersonalAccessToken = &SourceAirtableAuthenticationPersonalAccessToken{}
		}
	}
	if resp.Configuration.SourceType != nil {
		r.Configuration.SourceType = types.StringValue(string(*resp.Configuration.SourceType))
	} else {
		r.Configuration.SourceType = types.StringNull()
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
