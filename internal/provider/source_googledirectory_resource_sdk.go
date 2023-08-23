// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceGoogleDirectoryResourceModel) ToCreateSDKType() *shared.SourceGoogleDirectoryCreateRequest {
	var credentials *shared.SourceGoogleDirectoryGoogleCredentials
	if r.Configuration.Credentials != nil {
		var sourceGoogleDirectoryGoogleCredentialsSignInViaGoogleOAuth *shared.SourceGoogleDirectoryGoogleCredentialsSignInViaGoogleOAuth
		if r.Configuration.Credentials.SourceGoogleDirectoryGoogleCredentialsSignInViaGoogleOAuth != nil {
			clientID := r.Configuration.Credentials.SourceGoogleDirectoryGoogleCredentialsSignInViaGoogleOAuth.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceGoogleDirectoryGoogleCredentialsSignInViaGoogleOAuth.ClientSecret.ValueString()
			credentialsTitle := new(shared.SourceGoogleDirectoryGoogleCredentialsSignInViaGoogleOAuthCredentialsTitle)
			if !r.Configuration.Credentials.SourceGoogleDirectoryGoogleCredentialsSignInViaGoogleOAuth.CredentialsTitle.IsUnknown() && !r.Configuration.Credentials.SourceGoogleDirectoryGoogleCredentialsSignInViaGoogleOAuth.CredentialsTitle.IsNull() {
				*credentialsTitle = shared.SourceGoogleDirectoryGoogleCredentialsSignInViaGoogleOAuthCredentialsTitle(r.Configuration.Credentials.SourceGoogleDirectoryGoogleCredentialsSignInViaGoogleOAuth.CredentialsTitle.ValueString())
			} else {
				credentialsTitle = nil
			}
			refreshToken := r.Configuration.Credentials.SourceGoogleDirectoryGoogleCredentialsSignInViaGoogleOAuth.RefreshToken.ValueString()
			sourceGoogleDirectoryGoogleCredentialsSignInViaGoogleOAuth = &shared.SourceGoogleDirectoryGoogleCredentialsSignInViaGoogleOAuth{
				ClientID:         clientID,
				ClientSecret:     clientSecret,
				CredentialsTitle: credentialsTitle,
				RefreshToken:     refreshToken,
			}
		}
		if sourceGoogleDirectoryGoogleCredentialsSignInViaGoogleOAuth != nil {
			credentials = &shared.SourceGoogleDirectoryGoogleCredentials{
				SourceGoogleDirectoryGoogleCredentialsSignInViaGoogleOAuth: sourceGoogleDirectoryGoogleCredentialsSignInViaGoogleOAuth,
			}
		}
		var sourceGoogleDirectoryGoogleCredentialsServiceAccountKey *shared.SourceGoogleDirectoryGoogleCredentialsServiceAccountKey
		if r.Configuration.Credentials.SourceGoogleDirectoryGoogleCredentialsServiceAccountKey != nil {
			credentialsJSON := r.Configuration.Credentials.SourceGoogleDirectoryGoogleCredentialsServiceAccountKey.CredentialsJSON.ValueString()
			credentialsTitle1 := new(shared.SourceGoogleDirectoryGoogleCredentialsServiceAccountKeyCredentialsTitle)
			if !r.Configuration.Credentials.SourceGoogleDirectoryGoogleCredentialsServiceAccountKey.CredentialsTitle.IsUnknown() && !r.Configuration.Credentials.SourceGoogleDirectoryGoogleCredentialsServiceAccountKey.CredentialsTitle.IsNull() {
				*credentialsTitle1 = shared.SourceGoogleDirectoryGoogleCredentialsServiceAccountKeyCredentialsTitle(r.Configuration.Credentials.SourceGoogleDirectoryGoogleCredentialsServiceAccountKey.CredentialsTitle.ValueString())
			} else {
				credentialsTitle1 = nil
			}
			email := r.Configuration.Credentials.SourceGoogleDirectoryGoogleCredentialsServiceAccountKey.Email.ValueString()
			sourceGoogleDirectoryGoogleCredentialsServiceAccountKey = &shared.SourceGoogleDirectoryGoogleCredentialsServiceAccountKey{
				CredentialsJSON:  credentialsJSON,
				CredentialsTitle: credentialsTitle1,
				Email:            email,
			}
		}
		if sourceGoogleDirectoryGoogleCredentialsServiceAccountKey != nil {
			credentials = &shared.SourceGoogleDirectoryGoogleCredentials{
				SourceGoogleDirectoryGoogleCredentialsServiceAccountKey: sourceGoogleDirectoryGoogleCredentialsServiceAccountKey,
			}
		}
	}
	sourceType := shared.SourceGoogleDirectoryGoogleDirectory(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceGoogleDirectory{
		Credentials: credentials,
		SourceType:  sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleDirectoryCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleDirectoryResourceModel) ToGetSDKType() *shared.SourceGoogleDirectoryCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGoogleDirectoryResourceModel) ToUpdateSDKType() *shared.SourceGoogleDirectoryPutRequest {
	var credentials *shared.SourceGoogleDirectoryUpdateGoogleCredentials
	if r.Configuration.Credentials != nil {
		var sourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth *shared.SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth
		if r.Configuration.Credentials.SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth != nil {
			clientID := r.Configuration.Credentials.SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth.ClientSecret.ValueString()
			credentialsTitle := new(shared.SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuthCredentialsTitle)
			if !r.Configuration.Credentials.SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth.CredentialsTitle.IsUnknown() && !r.Configuration.Credentials.SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth.CredentialsTitle.IsNull() {
				*credentialsTitle = shared.SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuthCredentialsTitle(r.Configuration.Credentials.SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth.CredentialsTitle.ValueString())
			} else {
				credentialsTitle = nil
			}
			refreshToken := r.Configuration.Credentials.SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth.RefreshToken.ValueString()
			sourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth = &shared.SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth{
				ClientID:         clientID,
				ClientSecret:     clientSecret,
				CredentialsTitle: credentialsTitle,
				RefreshToken:     refreshToken,
			}
		}
		if sourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth != nil {
			credentials = &shared.SourceGoogleDirectoryUpdateGoogleCredentials{
				SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth: sourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth,
			}
		}
		var sourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey *shared.SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey
		if r.Configuration.Credentials.SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey != nil {
			credentialsJSON := r.Configuration.Credentials.SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey.CredentialsJSON.ValueString()
			credentialsTitle1 := new(shared.SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKeyCredentialsTitle)
			if !r.Configuration.Credentials.SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey.CredentialsTitle.IsUnknown() && !r.Configuration.Credentials.SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey.CredentialsTitle.IsNull() {
				*credentialsTitle1 = shared.SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKeyCredentialsTitle(r.Configuration.Credentials.SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey.CredentialsTitle.ValueString())
			} else {
				credentialsTitle1 = nil
			}
			email := r.Configuration.Credentials.SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey.Email.ValueString()
			sourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey = &shared.SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey{
				CredentialsJSON:  credentialsJSON,
				CredentialsTitle: credentialsTitle1,
				Email:            email,
			}
		}
		if sourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey != nil {
			credentials = &shared.SourceGoogleDirectoryUpdateGoogleCredentials{
				SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey: sourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey,
			}
		}
	}
	configuration := shared.SourceGoogleDirectoryUpdate{
		Credentials: credentials,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleDirectoryPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleDirectoryResourceModel) ToDeleteSDKType() *shared.SourceGoogleDirectoryCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGoogleDirectoryResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceGoogleDirectoryResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
