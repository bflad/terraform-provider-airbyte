// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceHarvestResourceModel) ToCreateSDKType() *shared.SourceHarvestCreateRequest {
	accountID := r.Configuration.AccountID.ValueString()
	var credentials *shared.SourceHarvestAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var sourceHarvestAuthenticateViaHarvestOAuth *shared.SourceHarvestAuthenticateViaHarvestOAuth
		if r.Configuration.Credentials.AuthenticateViaHarvestOAuth != nil {
			var additionalProperties interface{}
			if !r.Configuration.Credentials.AuthenticateViaHarvestOAuth.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.AuthenticateViaHarvestOAuth.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.AuthenticateViaHarvestOAuth.AdditionalProperties.ValueString()), &additionalProperties)
			}
			clientID := r.Configuration.Credentials.AuthenticateViaHarvestOAuth.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.AuthenticateViaHarvestOAuth.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.AuthenticateViaHarvestOAuth.RefreshToken.ValueString()
			sourceHarvestAuthenticateViaHarvestOAuth = &shared.SourceHarvestAuthenticateViaHarvestOAuth{
				AdditionalProperties: additionalProperties,
				ClientID:             clientID,
				ClientSecret:         clientSecret,
				RefreshToken:         refreshToken,
			}
		}
		if sourceHarvestAuthenticateViaHarvestOAuth != nil {
			credentials = &shared.SourceHarvestAuthenticationMechanism{
				AuthenticateViaHarvestOAuth: sourceHarvestAuthenticateViaHarvestOAuth,
			}
		}
		var sourceHarvestAuthenticateWithPersonalAccessToken *shared.SourceHarvestAuthenticateWithPersonalAccessToken
		if r.Configuration.Credentials.AuthenticateWithPersonalAccessToken != nil {
			var additionalProperties1 interface{}
			if !r.Configuration.Credentials.AuthenticateWithPersonalAccessToken.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.AuthenticateWithPersonalAccessToken.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.AuthenticateWithPersonalAccessToken.AdditionalProperties.ValueString()), &additionalProperties1)
			}
			apiToken := r.Configuration.Credentials.AuthenticateWithPersonalAccessToken.APIToken.ValueString()
			sourceHarvestAuthenticateWithPersonalAccessToken = &shared.SourceHarvestAuthenticateWithPersonalAccessToken{
				AdditionalProperties: additionalProperties1,
				APIToken:             apiToken,
			}
		}
		if sourceHarvestAuthenticateWithPersonalAccessToken != nil {
			credentials = &shared.SourceHarvestAuthenticationMechanism{
				AuthenticateWithPersonalAccessToken: sourceHarvestAuthenticateWithPersonalAccessToken,
			}
		}
	}
	replicationEndDate := new(time.Time)
	if !r.Configuration.ReplicationEndDate.IsUnknown() && !r.Configuration.ReplicationEndDate.IsNull() {
		*replicationEndDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.ReplicationEndDate.ValueString())
	} else {
		replicationEndDate = nil
	}
	replicationStartDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.ReplicationStartDate.ValueString())
	configuration := shared.SourceHarvest{
		AccountID:            accountID,
		Credentials:          credentials,
		ReplicationEndDate:   replicationEndDate,
		ReplicationStartDate: replicationStartDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceHarvestCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceHarvestResourceModel) ToGetSDKType() *shared.SourceHarvestCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceHarvestResourceModel) ToUpdateSDKType() *shared.SourceHarvestPutRequest {
	accountID := r.Configuration.AccountID.ValueString()
	var credentials *shared.SourceHarvestUpdateAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var authenticateViaHarvestOAuth *shared.AuthenticateViaHarvestOAuth
		if r.Configuration.Credentials.AuthenticateViaHarvestOAuth != nil {
			var additionalProperties interface{}
			if !r.Configuration.Credentials.AuthenticateViaHarvestOAuth.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.AuthenticateViaHarvestOAuth.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.AuthenticateViaHarvestOAuth.AdditionalProperties.ValueString()), &additionalProperties)
			}
			clientID := r.Configuration.Credentials.AuthenticateViaHarvestOAuth.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.AuthenticateViaHarvestOAuth.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.AuthenticateViaHarvestOAuth.RefreshToken.ValueString()
			authenticateViaHarvestOAuth = &shared.AuthenticateViaHarvestOAuth{
				AdditionalProperties: additionalProperties,
				ClientID:             clientID,
				ClientSecret:         clientSecret,
				RefreshToken:         refreshToken,
			}
		}
		if authenticateViaHarvestOAuth != nil {
			credentials = &shared.SourceHarvestUpdateAuthenticationMechanism{
				AuthenticateViaHarvestOAuth: authenticateViaHarvestOAuth,
			}
		}
		var sourceHarvestUpdateAuthenticateWithPersonalAccessToken *shared.SourceHarvestUpdateAuthenticateWithPersonalAccessToken
		if r.Configuration.Credentials.AuthenticateWithPersonalAccessToken != nil {
			var additionalProperties1 interface{}
			if !r.Configuration.Credentials.AuthenticateWithPersonalAccessToken.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.AuthenticateWithPersonalAccessToken.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.AuthenticateWithPersonalAccessToken.AdditionalProperties.ValueString()), &additionalProperties1)
			}
			apiToken := r.Configuration.Credentials.AuthenticateWithPersonalAccessToken.APIToken.ValueString()
			sourceHarvestUpdateAuthenticateWithPersonalAccessToken = &shared.SourceHarvestUpdateAuthenticateWithPersonalAccessToken{
				AdditionalProperties: additionalProperties1,
				APIToken:             apiToken,
			}
		}
		if sourceHarvestUpdateAuthenticateWithPersonalAccessToken != nil {
			credentials = &shared.SourceHarvestUpdateAuthenticationMechanism{
				AuthenticateWithPersonalAccessToken: sourceHarvestUpdateAuthenticateWithPersonalAccessToken,
			}
		}
	}
	replicationEndDate := new(time.Time)
	if !r.Configuration.ReplicationEndDate.IsUnknown() && !r.Configuration.ReplicationEndDate.IsNull() {
		*replicationEndDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.ReplicationEndDate.ValueString())
	} else {
		replicationEndDate = nil
	}
	replicationStartDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.ReplicationStartDate.ValueString())
	configuration := shared.SourceHarvestUpdate{
		AccountID:            accountID,
		Credentials:          credentials,
		ReplicationEndDate:   replicationEndDate,
		ReplicationStartDate: replicationStartDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceHarvestPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceHarvestResourceModel) ToDeleteSDKType() *shared.SourceHarvestCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceHarvestResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceHarvestResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
