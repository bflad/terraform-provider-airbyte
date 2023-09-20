// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceNotionResourceModel) ToCreateSDKType() *shared.SourceNotionCreateRequest {
	var credentials *shared.SourceNotionAuthenticateUsing
	if r.Configuration.Credentials != nil {
		var sourceNotionAuthenticateUsingOAuth20 *shared.SourceNotionAuthenticateUsingOAuth20
		if r.Configuration.Credentials.SourceNotionAuthenticateUsingOAuth20 != nil {
			accessToken := r.Configuration.Credentials.SourceNotionAuthenticateUsingOAuth20.AccessToken.ValueString()
			authType := shared.SourceNotionAuthenticateUsingOAuth20AuthType(r.Configuration.Credentials.SourceNotionAuthenticateUsingOAuth20.AuthType.ValueString())
			clientID := r.Configuration.Credentials.SourceNotionAuthenticateUsingOAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceNotionAuthenticateUsingOAuth20.ClientSecret.ValueString()
			sourceNotionAuthenticateUsingOAuth20 = &shared.SourceNotionAuthenticateUsingOAuth20{
				AccessToken:  accessToken,
				AuthType:     authType,
				ClientID:     clientID,
				ClientSecret: clientSecret,
			}
		}
		if sourceNotionAuthenticateUsingOAuth20 != nil {
			credentials = &shared.SourceNotionAuthenticateUsing{
				SourceNotionAuthenticateUsingOAuth20: sourceNotionAuthenticateUsingOAuth20,
			}
		}
		var sourceNotionAuthenticateUsingAccessToken *shared.SourceNotionAuthenticateUsingAccessToken
		if r.Configuration.Credentials.SourceNotionAuthenticateUsingAccessToken != nil {
			authType1 := shared.SourceNotionAuthenticateUsingAccessTokenAuthType(r.Configuration.Credentials.SourceNotionAuthenticateUsingAccessToken.AuthType.ValueString())
			token := r.Configuration.Credentials.SourceNotionAuthenticateUsingAccessToken.Token.ValueString()
			sourceNotionAuthenticateUsingAccessToken = &shared.SourceNotionAuthenticateUsingAccessToken{
				AuthType: authType1,
				Token:    token,
			}
		}
		if sourceNotionAuthenticateUsingAccessToken != nil {
			credentials = &shared.SourceNotionAuthenticateUsing{
				SourceNotionAuthenticateUsingAccessToken: sourceNotionAuthenticateUsingAccessToken,
			}
		}
	}
	sourceType := shared.SourceNotionNotion(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceNotion{
		Credentials: credentials,
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
	out := shared.SourceNotionCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceNotionResourceModel) ToGetSDKType() *shared.SourceNotionCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceNotionResourceModel) ToUpdateSDKType() *shared.SourceNotionPutRequest {
	var credentials *shared.SourceNotionUpdateAuthenticateUsing
	if r.Configuration.Credentials != nil {
		var sourceNotionUpdateAuthenticateUsingOAuth20 *shared.SourceNotionUpdateAuthenticateUsingOAuth20
		if r.Configuration.Credentials.SourceNotionUpdateAuthenticateUsingOAuth20 != nil {
			accessToken := r.Configuration.Credentials.SourceNotionUpdateAuthenticateUsingOAuth20.AccessToken.ValueString()
			authType := shared.SourceNotionUpdateAuthenticateUsingOAuth20AuthType(r.Configuration.Credentials.SourceNotionUpdateAuthenticateUsingOAuth20.AuthType.ValueString())
			clientID := r.Configuration.Credentials.SourceNotionUpdateAuthenticateUsingOAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceNotionUpdateAuthenticateUsingOAuth20.ClientSecret.ValueString()
			sourceNotionUpdateAuthenticateUsingOAuth20 = &shared.SourceNotionUpdateAuthenticateUsingOAuth20{
				AccessToken:  accessToken,
				AuthType:     authType,
				ClientID:     clientID,
				ClientSecret: clientSecret,
			}
		}
		if sourceNotionUpdateAuthenticateUsingOAuth20 != nil {
			credentials = &shared.SourceNotionUpdateAuthenticateUsing{
				SourceNotionUpdateAuthenticateUsingOAuth20: sourceNotionUpdateAuthenticateUsingOAuth20,
			}
		}
		var sourceNotionUpdateAuthenticateUsingAccessToken *shared.SourceNotionUpdateAuthenticateUsingAccessToken
		if r.Configuration.Credentials.SourceNotionUpdateAuthenticateUsingAccessToken != nil {
			authType1 := shared.SourceNotionUpdateAuthenticateUsingAccessTokenAuthType(r.Configuration.Credentials.SourceNotionUpdateAuthenticateUsingAccessToken.AuthType.ValueString())
			token := r.Configuration.Credentials.SourceNotionUpdateAuthenticateUsingAccessToken.Token.ValueString()
			sourceNotionUpdateAuthenticateUsingAccessToken = &shared.SourceNotionUpdateAuthenticateUsingAccessToken{
				AuthType: authType1,
				Token:    token,
			}
		}
		if sourceNotionUpdateAuthenticateUsingAccessToken != nil {
			credentials = &shared.SourceNotionUpdateAuthenticateUsing{
				SourceNotionUpdateAuthenticateUsingAccessToken: sourceNotionUpdateAuthenticateUsingAccessToken,
			}
		}
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceNotionUpdate{
		Credentials: credentials,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceNotionPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceNotionResourceModel) ToDeleteSDKType() *shared.SourceNotionCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceNotionResourceModel) RefreshFromGetResponse(resp *shared.SourceNotionGetResponse) {
	if resp.Configuration.Credentials == nil {
		r.Configuration.Credentials = nil
	} else {
		r.Configuration.Credentials = &SourceNotionAuthenticateUsing{}
		if resp.Configuration.Credentials.SourceNotionAuthenticateUsingAccessToken != nil {
			r.Configuration.Credentials.SourceNotionAuthenticateUsingAccessToken = &SourceNotionAuthenticateUsingAccessToken{}
			r.Configuration.Credentials.SourceNotionAuthenticateUsingAccessToken.AuthType = types.StringValue(string(resp.Configuration.Credentials.SourceNotionAuthenticateUsingAccessToken.AuthType))
			r.Configuration.Credentials.SourceNotionAuthenticateUsingAccessToken.Token = types.StringValue(resp.Configuration.Credentials.SourceNotionAuthenticateUsingAccessToken.Token)
		}
		if resp.Configuration.Credentials.SourceNotionAuthenticateUsingOAuth20 != nil {
			r.Configuration.Credentials.SourceNotionAuthenticateUsingOAuth20 = &SourceNotionAuthenticateUsingOAuth20{}
			r.Configuration.Credentials.SourceNotionAuthenticateUsingOAuth20.AccessToken = types.StringValue(resp.Configuration.Credentials.SourceNotionAuthenticateUsingOAuth20.AccessToken)
			r.Configuration.Credentials.SourceNotionAuthenticateUsingOAuth20.AuthType = types.StringValue(string(resp.Configuration.Credentials.SourceNotionAuthenticateUsingOAuth20.AuthType))
			r.Configuration.Credentials.SourceNotionAuthenticateUsingOAuth20.ClientID = types.StringValue(resp.Configuration.Credentials.SourceNotionAuthenticateUsingOAuth20.ClientID)
			r.Configuration.Credentials.SourceNotionAuthenticateUsingOAuth20.ClientSecret = types.StringValue(resp.Configuration.Credentials.SourceNotionAuthenticateUsingOAuth20.ClientSecret)
		}
		if resp.Configuration.Credentials.SourceNotionUpdateAuthenticateUsingAccessToken != nil {
			r.Configuration.Credentials.SourceNotionUpdateAuthenticateUsingAccessToken = &SourceNotionAuthenticateUsingAccessToken{}
		}
		if resp.Configuration.Credentials.SourceNotionUpdateAuthenticateUsingOAuth20 != nil {
			r.Configuration.Credentials.SourceNotionUpdateAuthenticateUsingOAuth20 = &SourceNotionAuthenticateUsingOAuth20{}
		}
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

func (r *SourceNotionResourceModel) RefreshFromCreateResponse(resp *shared.SourceNotionGetResponse) {
	r.RefreshFromGetResponse(resp)
}
