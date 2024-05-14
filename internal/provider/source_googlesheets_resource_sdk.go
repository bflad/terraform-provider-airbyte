// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceGoogleSheetsResourceModel) ToSharedSourceGoogleSheetsCreateRequest() *shared.SourceGoogleSheetsCreateRequest {
	batchSize := new(int64)
	if !r.Configuration.BatchSize.IsUnknown() && !r.Configuration.BatchSize.IsNull() {
		*batchSize = r.Configuration.BatchSize.ValueInt64()
	} else {
		batchSize = nil
	}
	var credentials shared.SourceGoogleSheetsAuthentication
	var sourceGoogleSheetsAuthenticateViaGoogleOAuth *shared.SourceGoogleSheetsAuthenticateViaGoogleOAuth
	if r.Configuration.Credentials.AuthenticateViaGoogleOAuth != nil {
		clientID := r.Configuration.Credentials.AuthenticateViaGoogleOAuth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.AuthenticateViaGoogleOAuth.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.AuthenticateViaGoogleOAuth.RefreshToken.ValueString()
		sourceGoogleSheetsAuthenticateViaGoogleOAuth = &shared.SourceGoogleSheetsAuthenticateViaGoogleOAuth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
		}
	}
	if sourceGoogleSheetsAuthenticateViaGoogleOAuth != nil {
		credentials = shared.SourceGoogleSheetsAuthentication{
			SourceGoogleSheetsAuthenticateViaGoogleOAuth: sourceGoogleSheetsAuthenticateViaGoogleOAuth,
		}
	}
	var sourceGoogleSheetsServiceAccountKeyAuthentication *shared.SourceGoogleSheetsServiceAccountKeyAuthentication
	if r.Configuration.Credentials.ServiceAccountKeyAuthentication != nil {
		serviceAccountInfo := r.Configuration.Credentials.ServiceAccountKeyAuthentication.ServiceAccountInfo.ValueString()
		sourceGoogleSheetsServiceAccountKeyAuthentication = &shared.SourceGoogleSheetsServiceAccountKeyAuthentication{
			ServiceAccountInfo: serviceAccountInfo,
		}
	}
	if sourceGoogleSheetsServiceAccountKeyAuthentication != nil {
		credentials = shared.SourceGoogleSheetsAuthentication{
			SourceGoogleSheetsServiceAccountKeyAuthentication: sourceGoogleSheetsServiceAccountKeyAuthentication,
		}
	}
	namesConversion := new(bool)
	if !r.Configuration.NamesConversion.IsUnknown() && !r.Configuration.NamesConversion.IsNull() {
		*namesConversion = r.Configuration.NamesConversion.ValueBool()
	} else {
		namesConversion = nil
	}
	spreadsheetID := r.Configuration.SpreadsheetID.ValueString()
	configuration := shared.SourceGoogleSheets{
		BatchSize:       batchSize,
		Credentials:     credentials,
		NamesConversion: namesConversion,
		SpreadsheetID:   spreadsheetID,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleSheetsCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleSheetsResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceGoogleSheetsResourceModel) ToSharedSourceGoogleSheetsPutRequest() *shared.SourceGoogleSheetsPutRequest {
	batchSize := new(int64)
	if !r.Configuration.BatchSize.IsUnknown() && !r.Configuration.BatchSize.IsNull() {
		*batchSize = r.Configuration.BatchSize.ValueInt64()
	} else {
		batchSize = nil
	}
	var credentials shared.SourceGoogleSheetsUpdateAuthentication
	var sourceGoogleSheetsUpdateAuthenticateViaGoogleOAuth *shared.SourceGoogleSheetsUpdateAuthenticateViaGoogleOAuth
	if r.Configuration.Credentials.AuthenticateViaGoogleOAuth != nil {
		clientID := r.Configuration.Credentials.AuthenticateViaGoogleOAuth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.AuthenticateViaGoogleOAuth.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.AuthenticateViaGoogleOAuth.RefreshToken.ValueString()
		sourceGoogleSheetsUpdateAuthenticateViaGoogleOAuth = &shared.SourceGoogleSheetsUpdateAuthenticateViaGoogleOAuth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
		}
	}
	if sourceGoogleSheetsUpdateAuthenticateViaGoogleOAuth != nil {
		credentials = shared.SourceGoogleSheetsUpdateAuthentication{
			SourceGoogleSheetsUpdateAuthenticateViaGoogleOAuth: sourceGoogleSheetsUpdateAuthenticateViaGoogleOAuth,
		}
	}
	var sourceGoogleSheetsUpdateServiceAccountKeyAuthentication *shared.SourceGoogleSheetsUpdateServiceAccountKeyAuthentication
	if r.Configuration.Credentials.ServiceAccountKeyAuthentication != nil {
		serviceAccountInfo := r.Configuration.Credentials.ServiceAccountKeyAuthentication.ServiceAccountInfo.ValueString()
		sourceGoogleSheetsUpdateServiceAccountKeyAuthentication = &shared.SourceGoogleSheetsUpdateServiceAccountKeyAuthentication{
			ServiceAccountInfo: serviceAccountInfo,
		}
	}
	if sourceGoogleSheetsUpdateServiceAccountKeyAuthentication != nil {
		credentials = shared.SourceGoogleSheetsUpdateAuthentication{
			SourceGoogleSheetsUpdateServiceAccountKeyAuthentication: sourceGoogleSheetsUpdateServiceAccountKeyAuthentication,
		}
	}
	namesConversion := new(bool)
	if !r.Configuration.NamesConversion.IsUnknown() && !r.Configuration.NamesConversion.IsNull() {
		*namesConversion = r.Configuration.NamesConversion.ValueBool()
	} else {
		namesConversion = nil
	}
	spreadsheetID := r.Configuration.SpreadsheetID.ValueString()
	configuration := shared.SourceGoogleSheetsUpdate{
		BatchSize:       batchSize,
		Credentials:     credentials,
		NamesConversion: namesConversion,
		SpreadsheetID:   spreadsheetID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleSheetsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
