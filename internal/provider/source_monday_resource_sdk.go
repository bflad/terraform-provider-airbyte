// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceMondayResourceModel) ToCreateSDKType() *shared.SourceMondayCreateRequest {
	var credentials *shared.SourceMondayAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceMondayOAuth20 *shared.SourceMondayOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			subdomain := new(string)
			if !r.Configuration.Credentials.OAuth20.Subdomain.IsUnknown() && !r.Configuration.Credentials.OAuth20.Subdomain.IsNull() {
				*subdomain = r.Configuration.Credentials.OAuth20.Subdomain.ValueString()
			} else {
				subdomain = nil
			}
			sourceMondayOAuth20 = &shared.SourceMondayOAuth20{
				AccessToken:  accessToken,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				Subdomain:    subdomain,
			}
		}
		if sourceMondayOAuth20 != nil {
			credentials = &shared.SourceMondayAuthorizationMethod{
				OAuth20: sourceMondayOAuth20,
			}
		}
		var sourceMondayAPIToken *shared.SourceMondayAPIToken
		if r.Configuration.Credentials.APIToken != nil {
			apiToken := r.Configuration.Credentials.APIToken.APIToken.ValueString()
			sourceMondayAPIToken = &shared.SourceMondayAPIToken{
				APIToken: apiToken,
			}
		}
		if sourceMondayAPIToken != nil {
			credentials = &shared.SourceMondayAuthorizationMethod{
				APIToken: sourceMondayAPIToken,
			}
		}
	}
	configuration := shared.SourceMonday{
		Credentials: credentials,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMondayCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMondayResourceModel) ToGetSDKType() *shared.SourceMondayCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMondayResourceModel) ToUpdateSDKType() *shared.SourceMondayPutRequest {
	var credentials *shared.SourceMondayUpdateAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceMondayUpdateOAuth20 *shared.SourceMondayUpdateOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			subdomain := new(string)
			if !r.Configuration.Credentials.OAuth20.Subdomain.IsUnknown() && !r.Configuration.Credentials.OAuth20.Subdomain.IsNull() {
				*subdomain = r.Configuration.Credentials.OAuth20.Subdomain.ValueString()
			} else {
				subdomain = nil
			}
			sourceMondayUpdateOAuth20 = &shared.SourceMondayUpdateOAuth20{
				AccessToken:  accessToken,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				Subdomain:    subdomain,
			}
		}
		if sourceMondayUpdateOAuth20 != nil {
			credentials = &shared.SourceMondayUpdateAuthorizationMethod{
				OAuth20: sourceMondayUpdateOAuth20,
			}
		}
		var apiToken *shared.APIToken
		if r.Configuration.Credentials.APIToken != nil {
			apiToken1 := r.Configuration.Credentials.APIToken.APIToken.ValueString()
			apiToken = &shared.APIToken{
				APIToken: apiToken1,
			}
		}
		if apiToken != nil {
			credentials = &shared.SourceMondayUpdateAuthorizationMethod{
				APIToken: apiToken,
			}
		}
	}
	configuration := shared.SourceMondayUpdate{
		Credentials: credentials,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMondayPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMondayResourceModel) ToDeleteSDKType() *shared.SourceMondayCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMondayResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceMondayResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
