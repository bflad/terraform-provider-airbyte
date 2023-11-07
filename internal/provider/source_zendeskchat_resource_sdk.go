// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceZendeskChatResourceModel) ToCreateSDKType() *shared.SourceZendeskChatCreateRequest {
	var credentials *shared.SourceZendeskChatAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceZendeskChatOAuth20 *shared.SourceZendeskChatOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			accessToken := new(string)
			if !r.Configuration.Credentials.OAuth20.AccessToken.IsUnknown() && !r.Configuration.Credentials.OAuth20.AccessToken.IsNull() {
				*accessToken = r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			} else {
				accessToken = nil
			}
			clientID := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			refreshToken := new(string)
			if !r.Configuration.Credentials.OAuth20.RefreshToken.IsUnknown() && !r.Configuration.Credentials.OAuth20.RefreshToken.IsNull() {
				*refreshToken = r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
			} else {
				refreshToken = nil
			}
			sourceZendeskChatOAuth20 = &shared.SourceZendeskChatOAuth20{
				AccessToken:  accessToken,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceZendeskChatOAuth20 != nil {
			credentials = &shared.SourceZendeskChatAuthorizationMethod{
				OAuth20: sourceZendeskChatOAuth20,
			}
		}
		var sourceZendeskChatAccessToken *shared.SourceZendeskChatAccessToken
		if r.Configuration.Credentials.AccessToken != nil {
			accessToken1 := r.Configuration.Credentials.AccessToken.AccessToken.ValueString()
			sourceZendeskChatAccessToken = &shared.SourceZendeskChatAccessToken{
				AccessToken: accessToken1,
			}
		}
		if sourceZendeskChatAccessToken != nil {
			credentials = &shared.SourceZendeskChatAuthorizationMethod{
				AccessToken: sourceZendeskChatAccessToken,
			}
		}
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	subdomain := new(string)
	if !r.Configuration.Subdomain.IsUnknown() && !r.Configuration.Subdomain.IsNull() {
		*subdomain = r.Configuration.Subdomain.ValueString()
	} else {
		subdomain = nil
	}
	configuration := shared.SourceZendeskChat{
		Credentials: credentials,
		StartDate:   startDate,
		Subdomain:   subdomain,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZendeskChatCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZendeskChatResourceModel) ToGetSDKType() *shared.SourceZendeskChatCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZendeskChatResourceModel) ToUpdateSDKType() *shared.SourceZendeskChatPutRequest {
	var credentials *shared.SourceZendeskChatUpdateAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceZendeskChatUpdateOAuth20 *shared.SourceZendeskChatUpdateOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			accessToken := new(string)
			if !r.Configuration.Credentials.OAuth20.AccessToken.IsUnknown() && !r.Configuration.Credentials.OAuth20.AccessToken.IsNull() {
				*accessToken = r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			} else {
				accessToken = nil
			}
			clientID := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			refreshToken := new(string)
			if !r.Configuration.Credentials.OAuth20.RefreshToken.IsUnknown() && !r.Configuration.Credentials.OAuth20.RefreshToken.IsNull() {
				*refreshToken = r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
			} else {
				refreshToken = nil
			}
			sourceZendeskChatUpdateOAuth20 = &shared.SourceZendeskChatUpdateOAuth20{
				AccessToken:  accessToken,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceZendeskChatUpdateOAuth20 != nil {
			credentials = &shared.SourceZendeskChatUpdateAuthorizationMethod{
				OAuth20: sourceZendeskChatUpdateOAuth20,
			}
		}
		var sourceZendeskChatUpdateAccessToken *shared.SourceZendeskChatUpdateAccessToken
		if r.Configuration.Credentials.AccessToken != nil {
			accessToken1 := r.Configuration.Credentials.AccessToken.AccessToken.ValueString()
			sourceZendeskChatUpdateAccessToken = &shared.SourceZendeskChatUpdateAccessToken{
				AccessToken: accessToken1,
			}
		}
		if sourceZendeskChatUpdateAccessToken != nil {
			credentials = &shared.SourceZendeskChatUpdateAuthorizationMethod{
				AccessToken: sourceZendeskChatUpdateAccessToken,
			}
		}
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	subdomain := new(string)
	if !r.Configuration.Subdomain.IsUnknown() && !r.Configuration.Subdomain.IsNull() {
		*subdomain = r.Configuration.Subdomain.ValueString()
	} else {
		subdomain = nil
	}
	configuration := shared.SourceZendeskChatUpdate{
		Credentials: credentials,
		StartDate:   startDate,
		Subdomain:   subdomain,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZendeskChatPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZendeskChatResourceModel) ToDeleteSDKType() *shared.SourceZendeskChatCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZendeskChatResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceZendeskChatResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
