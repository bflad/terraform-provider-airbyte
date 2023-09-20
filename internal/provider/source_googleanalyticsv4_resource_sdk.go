// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	customTypes "airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceGoogleAnalyticsV4ResourceModel) ToCreateSDKType() *shared.SourceGoogleAnalyticsV4CreateRequest {
	var credentials *shared.SourceGoogleAnalyticsV4Credentials
	if r.Configuration.Credentials != nil {
		var sourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth *shared.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth
		if r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth != nil {
			accessToken := new(string)
			if !r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AccessToken.IsUnknown() && !r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AccessToken.IsNull() {
				*accessToken = r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AccessToken.ValueString()
			} else {
				accessToken = nil
			}
			authType := new(shared.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauthAuthType)
			if !r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AuthType.IsNull() {
				*authType = shared.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauthAuthType(r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AuthType.ValueString())
			} else {
				authType = nil
			}
			clientID := r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.RefreshToken.ValueString()
			sourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth = &shared.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth{
				AccessToken:  accessToken,
				AuthType:     authType,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth != nil {
			credentials = &shared.SourceGoogleAnalyticsV4Credentials{
				SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth: sourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth,
			}
		}
		var sourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication *shared.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication
		if r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication != nil {
			authType1 := new(shared.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthenticationAuthType)
			if !r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication.AuthType.IsNull() {
				*authType1 = shared.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthenticationAuthType(r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication.AuthType.ValueString())
			} else {
				authType1 = nil
			}
			credentialsJSON := r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication.CredentialsJSON.ValueString()
			sourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication = &shared.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication{
				AuthType:        authType1,
				CredentialsJSON: credentialsJSON,
			}
		}
		if sourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication != nil {
			credentials = &shared.SourceGoogleAnalyticsV4Credentials{
				SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication: sourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication,
			}
		}
	}
	customReports := new(string)
	if !r.Configuration.CustomReports.IsUnknown() && !r.Configuration.CustomReports.IsNull() {
		*customReports = r.Configuration.CustomReports.ValueString()
	} else {
		customReports = nil
	}
	sourceType := shared.SourceGoogleAnalyticsV4GoogleAnalyticsV4(r.Configuration.SourceType.ValueString())
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	viewID := r.Configuration.ViewID.ValueString()
	windowInDays := new(int64)
	if !r.Configuration.WindowInDays.IsUnknown() && !r.Configuration.WindowInDays.IsNull() {
		*windowInDays = r.Configuration.WindowInDays.ValueInt64()
	} else {
		windowInDays = nil
	}
	configuration := shared.SourceGoogleAnalyticsV4{
		Credentials:   credentials,
		CustomReports: customReports,
		SourceType:    sourceType,
		StartDate:     startDate,
		ViewID:        viewID,
		WindowInDays:  windowInDays,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleAnalyticsV4CreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleAnalyticsV4ResourceModel) ToGetSDKType() *shared.SourceGoogleAnalyticsV4CreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGoogleAnalyticsV4ResourceModel) ToUpdateSDKType() *shared.SourceGoogleAnalyticsV4PutRequest {
	var credentials *shared.SourceGoogleAnalyticsV4UpdateCredentials
	if r.Configuration.Credentials != nil {
		var sourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth *shared.SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth
		if r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth != nil {
			accessToken := new(string)
			if !r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth.AccessToken.IsUnknown() && !r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth.AccessToken.IsNull() {
				*accessToken = r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth.AccessToken.ValueString()
			} else {
				accessToken = nil
			}
			authType := new(shared.SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauthAuthType)
			if !r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth.AuthType.IsNull() {
				*authType = shared.SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauthAuthType(r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth.AuthType.ValueString())
			} else {
				authType = nil
			}
			clientID := r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth.RefreshToken.ValueString()
			sourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth = &shared.SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth{
				AccessToken:  accessToken,
				AuthType:     authType,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth != nil {
			credentials = &shared.SourceGoogleAnalyticsV4UpdateCredentials{
				SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth: sourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth,
			}
		}
		var sourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthentication *shared.SourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthentication
		if r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthentication != nil {
			authType1 := new(shared.SourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthenticationAuthType)
			if !r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthentication.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthentication.AuthType.IsNull() {
				*authType1 = shared.SourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthenticationAuthType(r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthentication.AuthType.ValueString())
			} else {
				authType1 = nil
			}
			credentialsJSON := r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthentication.CredentialsJSON.ValueString()
			sourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthentication = &shared.SourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthentication{
				AuthType:        authType1,
				CredentialsJSON: credentialsJSON,
			}
		}
		if sourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthentication != nil {
			credentials = &shared.SourceGoogleAnalyticsV4UpdateCredentials{
				SourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthentication: sourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthentication,
			}
		}
	}
	customReports := new(string)
	if !r.Configuration.CustomReports.IsUnknown() && !r.Configuration.CustomReports.IsNull() {
		*customReports = r.Configuration.CustomReports.ValueString()
	} else {
		customReports = nil
	}
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	viewID := r.Configuration.ViewID.ValueString()
	windowInDays := new(int64)
	if !r.Configuration.WindowInDays.IsUnknown() && !r.Configuration.WindowInDays.IsNull() {
		*windowInDays = r.Configuration.WindowInDays.ValueInt64()
	} else {
		windowInDays = nil
	}
	configuration := shared.SourceGoogleAnalyticsV4Update{
		Credentials:   credentials,
		CustomReports: customReports,
		StartDate:     startDate,
		ViewID:        viewID,
		WindowInDays:  windowInDays,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleAnalyticsV4PutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleAnalyticsV4ResourceModel) ToDeleteSDKType() *shared.SourceGoogleAnalyticsV4CreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGoogleAnalyticsV4ResourceModel) RefreshFromGetResponse(resp *shared.SourceGoogleAnalyticsV4GetResponse) {
	if resp.Configuration.Credentials == nil {
		r.Configuration.Credentials = nil
	} else {
		r.Configuration.Credentials = &SourceGoogleAnalyticsV4Credentials{}
		if resp.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth != nil {
			r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth = &SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth{}
			if resp.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AccessToken != nil {
				r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AccessToken = types.StringValue(*resp.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AccessToken)
			} else {
				r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AccessToken = types.StringNull()
			}
			if resp.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AuthType != nil {
				r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AuthType = types.StringValue(string(*resp.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AuthType))
			} else {
				r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AuthType = types.StringNull()
			}
			r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.ClientID = types.StringValue(resp.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.ClientID)
			r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.ClientSecret = types.StringValue(resp.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.ClientSecret)
			r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.RefreshToken = types.StringValue(resp.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.RefreshToken)
		}
		if resp.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication != nil {
			r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication = &SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication{}
			if resp.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication.AuthType != nil {
				r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication.AuthType = types.StringValue(string(*resp.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication.AuthType))
			} else {
				r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication.AuthType = types.StringNull()
			}
			r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication.CredentialsJSON = types.StringValue(resp.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication.CredentialsJSON)
		}
		if resp.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth != nil {
			r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsAuthenticateViaGoogleOauth = &SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth{}
		}
		if resp.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthentication != nil {
			r.Configuration.Credentials.SourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthentication = &SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication{}
		}
	}
	if resp.Configuration.CustomReports != nil {
		r.Configuration.CustomReports = types.StringValue(*resp.Configuration.CustomReports)
	} else {
		r.Configuration.CustomReports = types.StringNull()
	}
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	r.Configuration.StartDate = types.StringValue(resp.Configuration.StartDate.String())
	r.Configuration.ViewID = types.StringValue(resp.Configuration.ViewID)
	if resp.Configuration.WindowInDays != nil {
		r.Configuration.WindowInDays = types.Int64Value(*resp.Configuration.WindowInDays)
	} else {
		r.Configuration.WindowInDays = types.Int64Null()
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

func (r *SourceGoogleAnalyticsV4ResourceModel) RefreshFromCreateResponse(resp *shared.SourceGoogleAnalyticsV4GetResponse) {
	r.RefreshFromGetResponse(resp)
}
