// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceOktaDataSourceModel) RefreshFromGetResponse(resp *shared.SourceOktaGetResponse) {
	if resp.Configuration.Credentials == nil {
		r.Configuration.Credentials = nil
	} else {
		r.Configuration.Credentials = &SourceOktaAuthorizationMethod{}
		if resp.Configuration.Credentials.SourceOktaAuthorizationMethodAPIToken != nil {
			r.Configuration.Credentials.SourceOktaAuthorizationMethodAPIToken = &SourceMondayAuthorizationMethodAPIToken{}
			r.Configuration.Credentials.SourceOktaAuthorizationMethodAPIToken.APIToken = types.StringValue(resp.Configuration.Credentials.SourceOktaAuthorizationMethodAPIToken.APIToken)
			r.Configuration.Credentials.SourceOktaAuthorizationMethodAPIToken.AuthType = types.StringValue(string(resp.Configuration.Credentials.SourceOktaAuthorizationMethodAPIToken.AuthType))
		}
		if resp.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20 != nil {
			r.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20 = &SourceOktaAuthorizationMethodOAuth20{}
			r.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20.AuthType = types.StringValue(string(resp.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20.AuthType))
			r.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20.ClientID = types.StringValue(resp.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20.ClientID)
			r.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20.ClientSecret = types.StringValue(resp.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20.ClientSecret)
			r.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20.RefreshToken = types.StringValue(resp.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20.RefreshToken)
		}
		if resp.Configuration.Credentials.SourceOktaUpdateAuthorizationMethodAPIToken != nil {
			r.Configuration.Credentials.SourceOktaUpdateAuthorizationMethodAPIToken = &SourceMondayAuthorizationMethodAPIToken{}
		}
		if resp.Configuration.Credentials.SourceOktaUpdateAuthorizationMethodOAuth20 != nil {
			r.Configuration.Credentials.SourceOktaUpdateAuthorizationMethodOAuth20 = &SourceOktaAuthorizationMethodOAuth20{}
		}
	}
	if resp.Configuration.Domain != nil {
		r.Configuration.Domain = types.StringValue(*resp.Configuration.Domain)
	} else {
		r.Configuration.Domain = types.StringNull()
	}
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	if resp.Configuration.StartDate != nil {
		r.Configuration.StartDate = types.StringValue(*resp.Configuration.StartDate)
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
