// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceOutbrainAmplifyResourceModel) ToCreateSDKType() *shared.SourceOutbrainAmplifyCreateRequest {
	var credentials shared.SourceOutbrainAmplifyAuthenticationMethod
	var sourceOutbrainAmplifyAuthenticationMethodAccessToken *shared.SourceOutbrainAmplifyAuthenticationMethodAccessToken
	if r.Configuration.Credentials.SourceOutbrainAmplifyAuthenticationMethodAccessToken != nil {
		accessToken := r.Configuration.Credentials.SourceOutbrainAmplifyAuthenticationMethodAccessToken.AccessToken.ValueString()
		typeVar := shared.SourceOutbrainAmplifyAuthenticationMethodAccessTokenAccessTokenIsRequiredForAuthenticationRequests(r.Configuration.Credentials.SourceOutbrainAmplifyAuthenticationMethodAccessToken.Type.ValueString())
		sourceOutbrainAmplifyAuthenticationMethodAccessToken = &shared.SourceOutbrainAmplifyAuthenticationMethodAccessToken{
			AccessToken: accessToken,
			Type:        typeVar,
		}
	}
	if sourceOutbrainAmplifyAuthenticationMethodAccessToken != nil {
		credentials = shared.SourceOutbrainAmplifyAuthenticationMethod{
			SourceOutbrainAmplifyAuthenticationMethodAccessToken: sourceOutbrainAmplifyAuthenticationMethodAccessToken,
		}
	}
	var sourceOutbrainAmplifyAuthenticationMethodUsernamePassword *shared.SourceOutbrainAmplifyAuthenticationMethodUsernamePassword
	if r.Configuration.Credentials.SourceOutbrainAmplifyAuthenticationMethodUsernamePassword != nil {
		password := r.Configuration.Credentials.SourceOutbrainAmplifyAuthenticationMethodUsernamePassword.Password.ValueString()
		typeVar1 := shared.SourceOutbrainAmplifyAuthenticationMethodUsernamePasswordBothUsernameAndPasswordIsRequiredForAuthenticationRequest(r.Configuration.Credentials.SourceOutbrainAmplifyAuthenticationMethodUsernamePassword.Type.ValueString())
		username := r.Configuration.Credentials.SourceOutbrainAmplifyAuthenticationMethodUsernamePassword.Username.ValueString()
		sourceOutbrainAmplifyAuthenticationMethodUsernamePassword = &shared.SourceOutbrainAmplifyAuthenticationMethodUsernamePassword{
			Password: password,
			Type:     typeVar1,
			Username: username,
		}
	}
	if sourceOutbrainAmplifyAuthenticationMethodUsernamePassword != nil {
		credentials = shared.SourceOutbrainAmplifyAuthenticationMethod{
			SourceOutbrainAmplifyAuthenticationMethodUsernamePassword: sourceOutbrainAmplifyAuthenticationMethodUsernamePassword,
		}
	}
	endDate := new(string)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate = r.Configuration.EndDate.ValueString()
	} else {
		endDate = nil
	}
	geoLocationBreakdown := new(shared.SourceOutbrainAmplifyGranularityForGeoLocationRegion)
	if !r.Configuration.GeoLocationBreakdown.IsUnknown() && !r.Configuration.GeoLocationBreakdown.IsNull() {
		*geoLocationBreakdown = shared.SourceOutbrainAmplifyGranularityForGeoLocationRegion(r.Configuration.GeoLocationBreakdown.ValueString())
	} else {
		geoLocationBreakdown = nil
	}
	reportGranularity := new(shared.SourceOutbrainAmplifyGranularityForPeriodicReports)
	if !r.Configuration.ReportGranularity.IsUnknown() && !r.Configuration.ReportGranularity.IsNull() {
		*reportGranularity = shared.SourceOutbrainAmplifyGranularityForPeriodicReports(r.Configuration.ReportGranularity.ValueString())
	} else {
		reportGranularity = nil
	}
	sourceType := shared.SourceOutbrainAmplifyOutbrainAmplify(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceOutbrainAmplify{
		Credentials:          credentials,
		EndDate:              endDate,
		GeoLocationBreakdown: geoLocationBreakdown,
		ReportGranularity:    reportGranularity,
		SourceType:           sourceType,
		StartDate:            startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceOutbrainAmplifyCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOutbrainAmplifyResourceModel) ToGetSDKType() *shared.SourceOutbrainAmplifyCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceOutbrainAmplifyResourceModel) ToUpdateSDKType() *shared.SourceOutbrainAmplifyPutRequest {
	var credentials shared.SourceOutbrainAmplifyUpdateAuthenticationMethod
	var sourceOutbrainAmplifyUpdateAuthenticationMethodAccessToken *shared.SourceOutbrainAmplifyUpdateAuthenticationMethodAccessToken
	if r.Configuration.Credentials.SourceOutbrainAmplifyUpdateAuthenticationMethodAccessToken != nil {
		accessToken := r.Configuration.Credentials.SourceOutbrainAmplifyUpdateAuthenticationMethodAccessToken.AccessToken.ValueString()
		typeVar := shared.SourceOutbrainAmplifyUpdateAuthenticationMethodAccessTokenAccessTokenIsRequiredForAuthenticationRequests(r.Configuration.Credentials.SourceOutbrainAmplifyUpdateAuthenticationMethodAccessToken.Type.ValueString())
		sourceOutbrainAmplifyUpdateAuthenticationMethodAccessToken = &shared.SourceOutbrainAmplifyUpdateAuthenticationMethodAccessToken{
			AccessToken: accessToken,
			Type:        typeVar,
		}
	}
	if sourceOutbrainAmplifyUpdateAuthenticationMethodAccessToken != nil {
		credentials = shared.SourceOutbrainAmplifyUpdateAuthenticationMethod{
			SourceOutbrainAmplifyUpdateAuthenticationMethodAccessToken: sourceOutbrainAmplifyUpdateAuthenticationMethodAccessToken,
		}
	}
	var sourceOutbrainAmplifyUpdateAuthenticationMethodUsernamePassword *shared.SourceOutbrainAmplifyUpdateAuthenticationMethodUsernamePassword
	if r.Configuration.Credentials.SourceOutbrainAmplifyUpdateAuthenticationMethodUsernamePassword != nil {
		password := r.Configuration.Credentials.SourceOutbrainAmplifyUpdateAuthenticationMethodUsernamePassword.Password.ValueString()
		typeVar1 := shared.SourceOutbrainAmplifyUpdateAuthenticationMethodUsernamePasswordBothUsernameAndPasswordIsRequiredForAuthenticationRequest(r.Configuration.Credentials.SourceOutbrainAmplifyUpdateAuthenticationMethodUsernamePassword.Type.ValueString())
		username := r.Configuration.Credentials.SourceOutbrainAmplifyUpdateAuthenticationMethodUsernamePassword.Username.ValueString()
		sourceOutbrainAmplifyUpdateAuthenticationMethodUsernamePassword = &shared.SourceOutbrainAmplifyUpdateAuthenticationMethodUsernamePassword{
			Password: password,
			Type:     typeVar1,
			Username: username,
		}
	}
	if sourceOutbrainAmplifyUpdateAuthenticationMethodUsernamePassword != nil {
		credentials = shared.SourceOutbrainAmplifyUpdateAuthenticationMethod{
			SourceOutbrainAmplifyUpdateAuthenticationMethodUsernamePassword: sourceOutbrainAmplifyUpdateAuthenticationMethodUsernamePassword,
		}
	}
	endDate := new(string)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate = r.Configuration.EndDate.ValueString()
	} else {
		endDate = nil
	}
	geoLocationBreakdown := new(shared.SourceOutbrainAmplifyUpdateGranularityForGeoLocationRegion)
	if !r.Configuration.GeoLocationBreakdown.IsUnknown() && !r.Configuration.GeoLocationBreakdown.IsNull() {
		*geoLocationBreakdown = shared.SourceOutbrainAmplifyUpdateGranularityForGeoLocationRegion(r.Configuration.GeoLocationBreakdown.ValueString())
	} else {
		geoLocationBreakdown = nil
	}
	reportGranularity := new(shared.SourceOutbrainAmplifyUpdateGranularityForPeriodicReports)
	if !r.Configuration.ReportGranularity.IsUnknown() && !r.Configuration.ReportGranularity.IsNull() {
		*reportGranularity = shared.SourceOutbrainAmplifyUpdateGranularityForPeriodicReports(r.Configuration.ReportGranularity.ValueString())
	} else {
		reportGranularity = nil
	}
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceOutbrainAmplifyUpdate{
		Credentials:          credentials,
		EndDate:              endDate,
		GeoLocationBreakdown: geoLocationBreakdown,
		ReportGranularity:    reportGranularity,
		StartDate:            startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceOutbrainAmplifyPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOutbrainAmplifyResourceModel) ToDeleteSDKType() *shared.SourceOutbrainAmplifyCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceOutbrainAmplifyResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceOutbrainAmplifyResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
