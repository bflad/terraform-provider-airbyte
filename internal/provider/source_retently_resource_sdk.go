// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceRetentlyResourceModel) ToCreateSDKType() *shared.SourceRetentlyCreateRequest {
	var credentials *shared.SourceRetentlyAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var sourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth *shared.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth
		if r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth != nil {
			authType := new(shared.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuthAuthType)
			if !r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.AuthType.IsNull() {
				*authType = shared.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuthAuthType(r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.AuthType.ValueString())
			} else {
				authType = nil
			}
			clientID := r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.RefreshToken.ValueString()
			var additionalProperties interface{}
			if !r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.AdditionalProperties.ValueString()), &additionalProperties)
			}
			sourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth = &shared.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth{
				AuthType:             authType,
				ClientID:             clientID,
				ClientSecret:         clientSecret,
				RefreshToken:         refreshToken,
				AdditionalProperties: additionalProperties,
			}
		}
		if sourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth != nil {
			credentials = &shared.SourceRetentlyAuthenticationMechanism{
				SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth: sourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth,
			}
		}
		var sourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken *shared.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken
		if r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken != nil {
			apiKey := r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.APIKey.ValueString()
			authType1 := new(shared.SourceRetentlyAuthenticationMechanismAuthenticateWithAPITokenAuthType)
			if !r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.AuthType.IsNull() {
				*authType1 = shared.SourceRetentlyAuthenticationMechanismAuthenticateWithAPITokenAuthType(r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.AuthType.ValueString())
			} else {
				authType1 = nil
			}
			var additionalProperties1 interface{}
			if !r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.AdditionalProperties.ValueString()), &additionalProperties1)
			}
			sourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken = &shared.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken{
				APIKey:               apiKey,
				AuthType:             authType1,
				AdditionalProperties: additionalProperties1,
			}
		}
		if sourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken != nil {
			credentials = &shared.SourceRetentlyAuthenticationMechanism{
				SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken: sourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken,
			}
		}
	}
	sourceType := new(shared.SourceRetentlyRetently)
	if !r.Configuration.SourceType.IsUnknown() && !r.Configuration.SourceType.IsNull() {
		*sourceType = shared.SourceRetentlyRetently(r.Configuration.SourceType.ValueString())
	} else {
		sourceType = nil
	}
	configuration := shared.SourceRetently{
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
	out := shared.SourceRetentlyCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceRetentlyResourceModel) ToGetSDKType() *shared.SourceRetentlyCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceRetentlyResourceModel) ToUpdateSDKType() *shared.SourceRetentlyPutRequest {
	var credentials *shared.SourceRetentlyUpdateAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var sourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth *shared.SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth
		if r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth != nil {
			authType := new(shared.SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuthAuthType)
			if !r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth.AuthType.IsNull() {
				*authType = shared.SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuthAuthType(r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth.AuthType.ValueString())
			} else {
				authType = nil
			}
			clientID := r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth.RefreshToken.ValueString()
			var additionalProperties interface{}
			if !r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth.AdditionalProperties.ValueString()), &additionalProperties)
			}
			sourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth = &shared.SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth{
				AuthType:             authType,
				ClientID:             clientID,
				ClientSecret:         clientSecret,
				RefreshToken:         refreshToken,
				AdditionalProperties: additionalProperties,
			}
		}
		if sourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth != nil {
			credentials = &shared.SourceRetentlyUpdateAuthenticationMechanism{
				SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth: sourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth,
			}
		}
		var sourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken *shared.SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken
		if r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken != nil {
			apiKey := r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken.APIKey.ValueString()
			authType1 := new(shared.SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPITokenAuthType)
			if !r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken.AuthType.IsNull() {
				*authType1 = shared.SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPITokenAuthType(r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken.AuthType.ValueString())
			} else {
				authType1 = nil
			}
			var additionalProperties1 interface{}
			if !r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken.AdditionalProperties.ValueString()), &additionalProperties1)
			}
			sourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken = &shared.SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken{
				APIKey:               apiKey,
				AuthType:             authType1,
				AdditionalProperties: additionalProperties1,
			}
		}
		if sourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken != nil {
			credentials = &shared.SourceRetentlyUpdateAuthenticationMechanism{
				SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken: sourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken,
			}
		}
	}
	configuration := shared.SourceRetentlyUpdate{
		Credentials: credentials,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceRetentlyPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceRetentlyResourceModel) ToDeleteSDKType() *shared.SourceRetentlyCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceRetentlyResourceModel) RefreshFromGetResponse(resp *shared.SourceRetentlyGetResponse) {
	if resp.Configuration.Credentials == nil {
		r.Configuration.Credentials = nil
	} else {
		r.Configuration.Credentials = &SourceRetentlyAuthenticationMechanism{}
		if resp.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth != nil {
			r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth = &SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth{}
			if resp.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.AuthType != nil {
				r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.AuthType = types.StringValue(string(*resp.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.AuthType))
			} else {
				r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.AuthType = types.StringNull()
			}
			r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.ClientID = types.StringValue(resp.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.ClientID)
			r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.ClientSecret = types.StringValue(resp.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.ClientSecret)
			r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.RefreshToken = types.StringValue(resp.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.RefreshToken)
			if r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.AdditionalProperties.IsUnknown() {
				if resp.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.AdditionalProperties == nil {
					r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.AdditionalProperties = types.StringNull()
				} else {
					additionalPropertiesResult, _ := json.Marshal(resp.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.AdditionalProperties)
					r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth.AdditionalProperties = types.StringValue(string(additionalPropertiesResult))
				}
			}
		}
		if resp.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken != nil {
			r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken = &SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken{}
			r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.APIKey = types.StringValue(resp.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.APIKey)
			if resp.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.AuthType != nil {
				r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.AuthType = types.StringValue(string(*resp.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.AuthType))
			} else {
				r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.AuthType = types.StringNull()
			}
			if r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.AdditionalProperties.IsUnknown() {
				if resp.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.AdditionalProperties == nil {
					r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.AdditionalProperties = types.StringNull()
				} else {
					additionalPropertiesResult1, _ := json.Marshal(resp.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.AdditionalProperties)
					r.Configuration.Credentials.SourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken.AdditionalProperties = types.StringValue(string(additionalPropertiesResult1))
				}
			}
		}
		if resp.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth != nil {
			r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth = &SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth{}
		}
		if resp.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken != nil {
			r.Configuration.Credentials.SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken = &SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithAPIToken{}
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

func (r *SourceRetentlyResourceModel) RefreshFromCreateResponse(resp *shared.SourceRetentlyGetResponse) {
	r.RefreshFromGetResponse(resp)
}
