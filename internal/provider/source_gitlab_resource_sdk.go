// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceGitlabResourceModel) ToCreateSDKType() *shared.SourceGitlabCreateRequest {
	apiURL := new(string)
	if !r.Configuration.APIURL.IsUnknown() && !r.Configuration.APIURL.IsNull() {
		*apiURL = r.Configuration.APIURL.ValueString()
	} else {
		apiURL = nil
	}
	var credentials shared.SourceGitlabAuthorizationMethod
	var sourceGitlabAuthorizationMethodOAuth20 *shared.SourceGitlabAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20 != nil {
		accessToken := r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.AccessToken.ValueString()
		authType := new(shared.SourceGitlabAuthorizationMethodOAuth20AuthType)
		if !r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.AuthType.IsNull() {
			*authType = shared.SourceGitlabAuthorizationMethodOAuth20AuthType(r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.AuthType.ValueString())
		} else {
			authType = nil
		}
		clientID := r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.TokenExpiryDate.ValueString())
		sourceGitlabAuthorizationMethodOAuth20 = &shared.SourceGitlabAuthorizationMethodOAuth20{
			AccessToken:     accessToken,
			AuthType:        authType,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceGitlabAuthorizationMethodOAuth20 != nil {
		credentials = shared.SourceGitlabAuthorizationMethod{
			SourceGitlabAuthorizationMethodOAuth20: sourceGitlabAuthorizationMethodOAuth20,
		}
	}
	var sourceGitlabAuthorizationMethodPrivateToken *shared.SourceGitlabAuthorizationMethodPrivateToken
	if r.Configuration.Credentials.SourceGitlabAuthorizationMethodPrivateToken != nil {
		accessToken1 := r.Configuration.Credentials.SourceGitlabAuthorizationMethodPrivateToken.AccessToken.ValueString()
		authType1 := new(shared.SourceGitlabAuthorizationMethodPrivateTokenAuthType)
		if !r.Configuration.Credentials.SourceGitlabAuthorizationMethodPrivateToken.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceGitlabAuthorizationMethodPrivateToken.AuthType.IsNull() {
			*authType1 = shared.SourceGitlabAuthorizationMethodPrivateTokenAuthType(r.Configuration.Credentials.SourceGitlabAuthorizationMethodPrivateToken.AuthType.ValueString())
		} else {
			authType1 = nil
		}
		sourceGitlabAuthorizationMethodPrivateToken = &shared.SourceGitlabAuthorizationMethodPrivateToken{
			AccessToken: accessToken1,
			AuthType:    authType1,
		}
	}
	if sourceGitlabAuthorizationMethodPrivateToken != nil {
		credentials = shared.SourceGitlabAuthorizationMethod{
			SourceGitlabAuthorizationMethodPrivateToken: sourceGitlabAuthorizationMethodPrivateToken,
		}
	}
	groups := new(string)
	if !r.Configuration.Groups.IsUnknown() && !r.Configuration.Groups.IsNull() {
		*groups = r.Configuration.Groups.ValueString()
	} else {
		groups = nil
	}
	projects := new(string)
	if !r.Configuration.Projects.IsUnknown() && !r.Configuration.Projects.IsNull() {
		*projects = r.Configuration.Projects.ValueString()
	} else {
		projects = nil
	}
	sourceType := shared.SourceGitlabGitlab(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceGitlab{
		APIURL:      apiURL,
		Credentials: credentials,
		Groups:      groups,
		Projects:    projects,
		SourceType:  sourceType,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGitlabCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGitlabResourceModel) ToGetSDKType() *shared.SourceGitlabCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGitlabResourceModel) ToUpdateSDKType() *shared.SourceGitlabPutRequest {
	apiURL := new(string)
	if !r.Configuration.APIURL.IsUnknown() && !r.Configuration.APIURL.IsNull() {
		*apiURL = r.Configuration.APIURL.ValueString()
	} else {
		apiURL = nil
	}
	var credentials shared.SourceGitlabUpdateAuthorizationMethod
	var sourceGitlabUpdateAuthorizationMethodOAuth20 *shared.SourceGitlabUpdateAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20 != nil {
		accessToken := r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20.AccessToken.ValueString()
		authType := new(shared.SourceGitlabUpdateAuthorizationMethodOAuth20AuthType)
		if !r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20.AuthType.IsNull() {
			*authType = shared.SourceGitlabUpdateAuthorizationMethodOAuth20AuthType(r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20.AuthType.ValueString())
		} else {
			authType = nil
		}
		clientID := r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20.TokenExpiryDate.ValueString())
		sourceGitlabUpdateAuthorizationMethodOAuth20 = &shared.SourceGitlabUpdateAuthorizationMethodOAuth20{
			AccessToken:     accessToken,
			AuthType:        authType,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceGitlabUpdateAuthorizationMethodOAuth20 != nil {
		credentials = shared.SourceGitlabUpdateAuthorizationMethod{
			SourceGitlabUpdateAuthorizationMethodOAuth20: sourceGitlabUpdateAuthorizationMethodOAuth20,
		}
	}
	var sourceGitlabUpdateAuthorizationMethodPrivateToken *shared.SourceGitlabUpdateAuthorizationMethodPrivateToken
	if r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodPrivateToken != nil {
		accessToken1 := r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodPrivateToken.AccessToken.ValueString()
		authType1 := new(shared.SourceGitlabUpdateAuthorizationMethodPrivateTokenAuthType)
		if !r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodPrivateToken.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodPrivateToken.AuthType.IsNull() {
			*authType1 = shared.SourceGitlabUpdateAuthorizationMethodPrivateTokenAuthType(r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodPrivateToken.AuthType.ValueString())
		} else {
			authType1 = nil
		}
		sourceGitlabUpdateAuthorizationMethodPrivateToken = &shared.SourceGitlabUpdateAuthorizationMethodPrivateToken{
			AccessToken: accessToken1,
			AuthType:    authType1,
		}
	}
	if sourceGitlabUpdateAuthorizationMethodPrivateToken != nil {
		credentials = shared.SourceGitlabUpdateAuthorizationMethod{
			SourceGitlabUpdateAuthorizationMethodPrivateToken: sourceGitlabUpdateAuthorizationMethodPrivateToken,
		}
	}
	groups := new(string)
	if !r.Configuration.Groups.IsUnknown() && !r.Configuration.Groups.IsNull() {
		*groups = r.Configuration.Groups.ValueString()
	} else {
		groups = nil
	}
	projects := new(string)
	if !r.Configuration.Projects.IsUnknown() && !r.Configuration.Projects.IsNull() {
		*projects = r.Configuration.Projects.ValueString()
	} else {
		projects = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceGitlabUpdate{
		APIURL:      apiURL,
		Credentials: credentials,
		Groups:      groups,
		Projects:    projects,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGitlabPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGitlabResourceModel) ToDeleteSDKType() *shared.SourceGitlabCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGitlabResourceModel) RefreshFromGetResponse(resp *shared.SourceGitlabGetResponse) {
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

func (r *SourceGitlabResourceModel) RefreshFromCreateResponse(resp *shared.SourceGitlabGetResponse) {
	r.RefreshFromGetResponse(resp)
}
