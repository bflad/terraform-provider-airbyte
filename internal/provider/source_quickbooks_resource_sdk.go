// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceQuickbooksResourceModel) ToCreateSDKType() *shared.SourceQuickbooksCreateRequest {
	var credentials shared.SourceQuickbooksAuthorizationMethod
	var sourceQuickbooksOAuth20 *shared.SourceQuickbooksOAuth20
	if r.Configuration.Credentials.OAuth20 != nil {
		accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
		clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
		realmID := r.Configuration.Credentials.OAuth20.RealmID.ValueString()
		refreshToken := r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.OAuth20.TokenExpiryDate.ValueString())
		sourceQuickbooksOAuth20 = &shared.SourceQuickbooksOAuth20{
			AccessToken:     accessToken,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RealmID:         realmID,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceQuickbooksOAuth20 != nil {
		credentials = shared.SourceQuickbooksAuthorizationMethod{
			OAuth20: sourceQuickbooksOAuth20,
		}
	}
	sandbox := new(bool)
	if !r.Configuration.Sandbox.IsUnknown() && !r.Configuration.Sandbox.IsNull() {
		*sandbox = r.Configuration.Sandbox.ValueBool()
	} else {
		sandbox = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceQuickbooks{
		Credentials: credentials,
		Sandbox:     sandbox,
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
	out := shared.SourceQuickbooksCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceQuickbooksResourceModel) ToGetSDKType() *shared.SourceQuickbooksCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceQuickbooksResourceModel) ToUpdateSDKType() *shared.SourceQuickbooksPutRequest {
	var credentials shared.SourceQuickbooksUpdateAuthorizationMethod
	var sourceQuickbooksUpdateOAuth20 *shared.SourceQuickbooksUpdateOAuth20
	if r.Configuration.Credentials.OAuth20 != nil {
		accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
		clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
		realmID := r.Configuration.Credentials.OAuth20.RealmID.ValueString()
		refreshToken := r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.OAuth20.TokenExpiryDate.ValueString())
		sourceQuickbooksUpdateOAuth20 = &shared.SourceQuickbooksUpdateOAuth20{
			AccessToken:     accessToken,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RealmID:         realmID,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceQuickbooksUpdateOAuth20 != nil {
		credentials = shared.SourceQuickbooksUpdateAuthorizationMethod{
			OAuth20: sourceQuickbooksUpdateOAuth20,
		}
	}
	sandbox := new(bool)
	if !r.Configuration.Sandbox.IsUnknown() && !r.Configuration.Sandbox.IsNull() {
		*sandbox = r.Configuration.Sandbox.ValueBool()
	} else {
		sandbox = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceQuickbooksUpdate{
		Credentials: credentials,
		Sandbox:     sandbox,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceQuickbooksPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceQuickbooksResourceModel) ToDeleteSDKType() *shared.SourceQuickbooksCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceQuickbooksResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceQuickbooksResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
