// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceGitlabDataSourceModel) RefreshFromGetResponse(resp *shared.SourceGitlabGetResponse) {
	if resp.Configuration.APIURL != nil {
		r.Configuration.APIURL = types.StringValue(*resp.Configuration.APIURL)
	} else {
		r.Configuration.APIURL = types.StringNull()
	}
	if resp.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20 != nil {
		r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20 = &SourceGitlabAuthorizationMethodOAuth20{}
		r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.AccessToken = types.StringValue(resp.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.AccessToken)
		if resp.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.AuthType != nil {
			r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.AuthType = types.StringValue(string(*resp.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.AuthType))
		} else {
			r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.AuthType = types.StringNull()
		}
		r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.ClientID = types.StringValue(resp.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.ClientID)
		r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.ClientSecret = types.StringValue(resp.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.ClientSecret)
		r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.RefreshToken = types.StringValue(resp.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.RefreshToken)
		r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.TokenExpiryDate = types.StringValue(resp.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.TokenExpiryDate.Format(time.RFC3339))
	}
	if resp.Configuration.Credentials.SourceGitlabAuthorizationMethodPrivateToken != nil {
		r.Configuration.Credentials.SourceGitlabAuthorizationMethodPrivateToken = &SourceGitlabAuthorizationMethodPrivateToken{}
		r.Configuration.Credentials.SourceGitlabAuthorizationMethodPrivateToken.AccessToken = types.StringValue(resp.Configuration.Credentials.SourceGitlabAuthorizationMethodPrivateToken.AccessToken)
		if resp.Configuration.Credentials.SourceGitlabAuthorizationMethodPrivateToken.AuthType != nil {
			r.Configuration.Credentials.SourceGitlabAuthorizationMethodPrivateToken.AuthType = types.StringValue(string(*resp.Configuration.Credentials.SourceGitlabAuthorizationMethodPrivateToken.AuthType))
		} else {
			r.Configuration.Credentials.SourceGitlabAuthorizationMethodPrivateToken.AuthType = types.StringNull()
		}
	}
	if resp.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20 != nil {
		r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20 = &SourceGitlabAuthorizationMethodOAuth20{}
	}
	if resp.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodPrivateToken != nil {
		r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodPrivateToken = &SourceGitlabAuthorizationMethodPrivateToken{}
	}
	if resp.Configuration.Groups != nil {
		r.Configuration.Groups = types.StringValue(*resp.Configuration.Groups)
	} else {
		r.Configuration.Groups = types.StringNull()
	}
	if resp.Configuration.Projects != nil {
		r.Configuration.Projects = types.StringValue(*resp.Configuration.Projects)
	} else {
		r.Configuration.Projects = types.StringNull()
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
