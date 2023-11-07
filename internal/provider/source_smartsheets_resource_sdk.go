// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceSmartsheetsResourceModel) ToCreateSDKType() *shared.SourceSmartsheetsCreateRequest {
	var credentials shared.SourceSmartsheetsAuthorizationMethod
	var sourceSmartsheetsOAuth20 *shared.SourceSmartsheetsOAuth20
	if r.Configuration.Credentials.OAuth20 != nil {
		accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
		clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.OAuth20.TokenExpiryDate.ValueString())
		sourceSmartsheetsOAuth20 = &shared.SourceSmartsheetsOAuth20{
			AccessToken:     accessToken,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceSmartsheetsOAuth20 != nil {
		credentials = shared.SourceSmartsheetsAuthorizationMethod{
			OAuth20: sourceSmartsheetsOAuth20,
		}
	}
	var sourceSmartsheetsAPIAccessToken *shared.SourceSmartsheetsAPIAccessToken
	if r.Configuration.Credentials.APIAccessToken != nil {
		accessToken1 := r.Configuration.Credentials.APIAccessToken.AccessToken.ValueString()
		sourceSmartsheetsAPIAccessToken = &shared.SourceSmartsheetsAPIAccessToken{
			AccessToken: accessToken1,
		}
	}
	if sourceSmartsheetsAPIAccessToken != nil {
		credentials = shared.SourceSmartsheetsAuthorizationMethod{
			APIAccessToken: sourceSmartsheetsAPIAccessToken,
		}
	}
	var metadataFields []shared.SourceSmartsheetsValidenums = nil
	for _, metadataFieldsItem := range r.Configuration.MetadataFields {
		metadataFields = append(metadataFields, shared.SourceSmartsheetsValidenums(metadataFieldsItem.ValueString()))
	}
	spreadsheetID := r.Configuration.SpreadsheetID.ValueString()
	startDatetime := new(time.Time)
	if !r.Configuration.StartDatetime.IsUnknown() && !r.Configuration.StartDatetime.IsNull() {
		*startDatetime, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDatetime.ValueString())
	} else {
		startDatetime = nil
	}
	configuration := shared.SourceSmartsheets{
		Credentials:    credentials,
		MetadataFields: metadataFields,
		SpreadsheetID:  spreadsheetID,
		StartDatetime:  startDatetime,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSmartsheetsCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSmartsheetsResourceModel) ToGetSDKType() *shared.SourceSmartsheetsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSmartsheetsResourceModel) ToUpdateSDKType() *shared.SourceSmartsheetsPutRequest {
	var credentials shared.SourceSmartsheetsUpdateAuthorizationMethod
	var sourceSmartsheetsUpdateOAuth20 *shared.SourceSmartsheetsUpdateOAuth20
	if r.Configuration.Credentials.OAuth20 != nil {
		accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
		clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.OAuth20.TokenExpiryDate.ValueString())
		sourceSmartsheetsUpdateOAuth20 = &shared.SourceSmartsheetsUpdateOAuth20{
			AccessToken:     accessToken,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceSmartsheetsUpdateOAuth20 != nil {
		credentials = shared.SourceSmartsheetsUpdateAuthorizationMethod{
			OAuth20: sourceSmartsheetsUpdateOAuth20,
		}
	}
	var apiAccessToken *shared.APIAccessToken
	if r.Configuration.Credentials.APIAccessToken != nil {
		accessToken1 := r.Configuration.Credentials.APIAccessToken.AccessToken.ValueString()
		apiAccessToken = &shared.APIAccessToken{
			AccessToken: accessToken1,
		}
	}
	if apiAccessToken != nil {
		credentials = shared.SourceSmartsheetsUpdateAuthorizationMethod{
			APIAccessToken: apiAccessToken,
		}
	}
	var metadataFields []shared.Validenums = nil
	for _, metadataFieldsItem := range r.Configuration.MetadataFields {
		metadataFields = append(metadataFields, shared.Validenums(metadataFieldsItem.ValueString()))
	}
	spreadsheetID := r.Configuration.SpreadsheetID.ValueString()
	startDatetime := new(time.Time)
	if !r.Configuration.StartDatetime.IsUnknown() && !r.Configuration.StartDatetime.IsNull() {
		*startDatetime, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDatetime.ValueString())
	} else {
		startDatetime = nil
	}
	configuration := shared.SourceSmartsheetsUpdate{
		Credentials:    credentials,
		MetadataFields: metadataFields,
		SpreadsheetID:  spreadsheetID,
		StartDatetime:  startDatetime,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSmartsheetsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSmartsheetsResourceModel) ToDeleteSDKType() *shared.SourceSmartsheetsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSmartsheetsResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceSmartsheetsResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
