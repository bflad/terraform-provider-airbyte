// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceLeverHiringResourceModel) ToCreateSDKType() *shared.SourceLeverHiringCreateRequest {
	var credentials *shared.SourceLeverHiringAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var sourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth *shared.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth
		if r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth != nil {
			authType := new(shared.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuthAuthType)
			if !r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.AuthType.IsNull() {
				*authType = shared.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuthAuthType(r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.AuthType.ValueString())
			} else {
				authType = nil
			}
			clientID := new(string)
			if !r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.ClientID.IsUnknown() && !r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.ClientSecret.IsUnknown() && !r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			refreshToken := r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.RefreshToken.ValueString()
			sourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth = &shared.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth{
				AuthType:     authType,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth != nil {
			credentials = &shared.SourceLeverHiringAuthenticationMechanism{
				SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth: sourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth,
			}
		}
		var sourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey *shared.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey
		if r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey != nil {
			apiKey := r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey.APIKey.ValueString()
			authType1 := new(shared.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKeyAuthType)
			if !r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey.AuthType.IsNull() {
				*authType1 = shared.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKeyAuthType(r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey.AuthType.ValueString())
			} else {
				authType1 = nil
			}
			sourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey = &shared.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey{
				APIKey:   apiKey,
				AuthType: authType1,
			}
		}
		if sourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey != nil {
			credentials = &shared.SourceLeverHiringAuthenticationMechanism{
				SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey: sourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey,
			}
		}
	}
	environment := new(shared.SourceLeverHiringEnvironment)
	if !r.Configuration.Environment.IsUnknown() && !r.Configuration.Environment.IsNull() {
		*environment = shared.SourceLeverHiringEnvironment(r.Configuration.Environment.ValueString())
	} else {
		environment = nil
	}
	sourceType := shared.SourceLeverHiringLeverHiring(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceLeverHiring{
		Credentials: credentials,
		Environment: environment,
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
	out := shared.SourceLeverHiringCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceLeverHiringResourceModel) ToGetSDKType() *shared.SourceLeverHiringCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceLeverHiringResourceModel) ToUpdateSDKType() *shared.SourceLeverHiringPutRequest {
	var credentials *shared.SourceLeverHiringUpdateAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var sourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth *shared.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth
		if r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth != nil {
			authType := new(shared.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuthAuthType)
			if !r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth.AuthType.IsNull() {
				*authType = shared.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuthAuthType(r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth.AuthType.ValueString())
			} else {
				authType = nil
			}
			clientID := new(string)
			if !r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth.ClientID.IsUnknown() && !r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth.ClientSecret.IsUnknown() && !r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			refreshToken := r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth.RefreshToken.ValueString()
			sourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth = &shared.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth{
				AuthType:     authType,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth != nil {
			credentials = &shared.SourceLeverHiringUpdateAuthenticationMechanism{
				SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth: sourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth,
			}
		}
		var sourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKey *shared.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKey
		if r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKey != nil {
			apiKey := r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKey.APIKey.ValueString()
			authType1 := new(shared.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKeyAuthType)
			if !r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKey.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKey.AuthType.IsNull() {
				*authType1 = shared.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKeyAuthType(r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKey.AuthType.ValueString())
			} else {
				authType1 = nil
			}
			sourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKey = &shared.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKey{
				APIKey:   apiKey,
				AuthType: authType1,
			}
		}
		if sourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKey != nil {
			credentials = &shared.SourceLeverHiringUpdateAuthenticationMechanism{
				SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKey: sourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKey,
			}
		}
	}
	environment := new(shared.SourceLeverHiringUpdateEnvironment)
	if !r.Configuration.Environment.IsUnknown() && !r.Configuration.Environment.IsNull() {
		*environment = shared.SourceLeverHiringUpdateEnvironment(r.Configuration.Environment.ValueString())
	} else {
		environment = nil
	}
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceLeverHiringUpdate{
		Credentials: credentials,
		Environment: environment,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceLeverHiringPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceLeverHiringResourceModel) ToDeleteSDKType() *shared.SourceLeverHiringCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceLeverHiringResourceModel) RefreshFromGetResponse(resp *shared.SourceLeverHiringGetResponse) {
	if resp.Configuration.Credentials == nil {
		r.Configuration.Credentials = nil
	} else {
		r.Configuration.Credentials = &SourceLeverHiringAuthenticationMechanism{}
		if resp.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey != nil {
			r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey = &SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey{}
			r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey.APIKey = types.StringValue(resp.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey.APIKey)
			if resp.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey.AuthType != nil {
				r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey.AuthType = types.StringValue(string(*resp.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey.AuthType))
			} else {
				r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey.AuthType = types.StringNull()
			}
		}
		if resp.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth != nil {
			r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth = &SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth{}
			if resp.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.AuthType != nil {
				r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.AuthType = types.StringValue(string(*resp.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.AuthType))
			} else {
				r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.AuthType = types.StringNull()
			}
			if resp.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.ClientID != nil {
				r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.ClientID = types.StringValue(*resp.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.ClientID)
			} else {
				r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.ClientID = types.StringNull()
			}
			if resp.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.ClientSecret != nil {
				r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.ClientSecret = types.StringValue(*resp.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.ClientSecret)
			} else {
				r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.ClientSecret = types.StringNull()
			}
			r.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.RefreshToken = types.StringValue(resp.Configuration.Credentials.SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth.RefreshToken)
		}
		if resp.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKey != nil {
			r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverAPIKey = &SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverAPIKey{}
		}
		if resp.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth != nil {
			r.Configuration.Credentials.SourceLeverHiringUpdateAuthenticationMechanismAuthenticateViaLeverOAuth = &SourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth{}
		}
	}
	if resp.Configuration.Environment != nil {
		r.Configuration.Environment = types.StringValue(string(*resp.Configuration.Environment))
	} else {
		r.Configuration.Environment = types.StringNull()
	}
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	r.Configuration.StartDate = types.StringValue(resp.Configuration.StartDate)
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

func (r *SourceLeverHiringResourceModel) RefreshFromCreateResponse(resp *shared.SourceLeverHiringGetResponse) {
	r.RefreshFromGetResponse(resp)
}
