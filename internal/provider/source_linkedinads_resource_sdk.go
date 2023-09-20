// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	customTypes "airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceLinkedinAdsResourceModel) ToCreateSDKType() *shared.SourceLinkedinAdsCreateRequest {
	var accountIds []int64 = nil
	for _, accountIdsItem := range r.Configuration.AccountIds {
		accountIds = append(accountIds, accountIdsItem.ValueInt64())
	}
	var adAnalyticsReports []shared.SourceLinkedinAdsAdAnalyticsReportConfiguration = nil
	for _, adAnalyticsReportsItem := range r.Configuration.AdAnalyticsReports {
		name := adAnalyticsReportsItem.Name.ValueString()
		pivotBy := shared.SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory(adAnalyticsReportsItem.PivotBy.ValueString())
		timeGranularity := shared.SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity(adAnalyticsReportsItem.TimeGranularity.ValueString())
		adAnalyticsReports = append(adAnalyticsReports, shared.SourceLinkedinAdsAdAnalyticsReportConfiguration{
			Name:            name,
			PivotBy:         pivotBy,
			TimeGranularity: timeGranularity,
		})
	}
	var credentials *shared.SourceLinkedinAdsAuthentication
	if r.Configuration.Credentials != nil {
		var sourceLinkedinAdsAuthenticationOAuth20 *shared.SourceLinkedinAdsAuthenticationOAuth20
		if r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20 != nil {
			authMethod := new(shared.SourceLinkedinAdsAuthenticationOAuth20AuthMethod)
			if !r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.AuthMethod.IsUnknown() && !r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.AuthMethod.IsNull() {
				*authMethod = shared.SourceLinkedinAdsAuthenticationOAuth20AuthMethod(r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.AuthMethod.ValueString())
			} else {
				authMethod = nil
			}
			clientID := r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.RefreshToken.ValueString()
			sourceLinkedinAdsAuthenticationOAuth20 = &shared.SourceLinkedinAdsAuthenticationOAuth20{
				AuthMethod:   authMethod,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceLinkedinAdsAuthenticationOAuth20 != nil {
			credentials = &shared.SourceLinkedinAdsAuthentication{
				SourceLinkedinAdsAuthenticationOAuth20: sourceLinkedinAdsAuthenticationOAuth20,
			}
		}
		var sourceLinkedinAdsAuthenticationAccessToken *shared.SourceLinkedinAdsAuthenticationAccessToken
		if r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken != nil {
			accessToken := r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken.AccessToken.ValueString()
			authMethod1 := new(shared.SourceLinkedinAdsAuthenticationAccessTokenAuthMethod)
			if !r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken.AuthMethod.IsUnknown() && !r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken.AuthMethod.IsNull() {
				*authMethod1 = shared.SourceLinkedinAdsAuthenticationAccessTokenAuthMethod(r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken.AuthMethod.ValueString())
			} else {
				authMethod1 = nil
			}
			sourceLinkedinAdsAuthenticationAccessToken = &shared.SourceLinkedinAdsAuthenticationAccessToken{
				AccessToken: accessToken,
				AuthMethod:  authMethod1,
			}
		}
		if sourceLinkedinAdsAuthenticationAccessToken != nil {
			credentials = &shared.SourceLinkedinAdsAuthentication{
				SourceLinkedinAdsAuthenticationAccessToken: sourceLinkedinAdsAuthenticationAccessToken,
			}
		}
	}
	sourceType := shared.SourceLinkedinAdsLinkedinAds(r.Configuration.SourceType.ValueString())
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	configuration := shared.SourceLinkedinAds{
		AccountIds:         accountIds,
		AdAnalyticsReports: adAnalyticsReports,
		Credentials:        credentials,
		SourceType:         sourceType,
		StartDate:          startDate,
	}
	name1 := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceLinkedinAdsCreateRequest{
		Configuration: configuration,
		Name:          name1,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceLinkedinAdsResourceModel) ToGetSDKType() *shared.SourceLinkedinAdsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceLinkedinAdsResourceModel) ToUpdateSDKType() *shared.SourceLinkedinAdsPutRequest {
	var accountIds []int64 = nil
	for _, accountIdsItem := range r.Configuration.AccountIds {
		accountIds = append(accountIds, accountIdsItem.ValueInt64())
	}
	var adAnalyticsReports []shared.SourceLinkedinAdsUpdateAdAnalyticsReportConfiguration = nil
	for _, adAnalyticsReportsItem := range r.Configuration.AdAnalyticsReports {
		name := adAnalyticsReportsItem.Name.ValueString()
		pivotBy := shared.SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotCategory(adAnalyticsReportsItem.PivotBy.ValueString())
		timeGranularity := shared.SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularity(adAnalyticsReportsItem.TimeGranularity.ValueString())
		adAnalyticsReports = append(adAnalyticsReports, shared.SourceLinkedinAdsUpdateAdAnalyticsReportConfiguration{
			Name:            name,
			PivotBy:         pivotBy,
			TimeGranularity: timeGranularity,
		})
	}
	var credentials *shared.SourceLinkedinAdsUpdateAuthentication
	if r.Configuration.Credentials != nil {
		var sourceLinkedinAdsUpdateAuthenticationOAuth20 *shared.SourceLinkedinAdsUpdateAuthenticationOAuth20
		if r.Configuration.Credentials.SourceLinkedinAdsUpdateAuthenticationOAuth20 != nil {
			authMethod := new(shared.SourceLinkedinAdsUpdateAuthenticationOAuth20AuthMethod)
			if !r.Configuration.Credentials.SourceLinkedinAdsUpdateAuthenticationOAuth20.AuthMethod.IsUnknown() && !r.Configuration.Credentials.SourceLinkedinAdsUpdateAuthenticationOAuth20.AuthMethod.IsNull() {
				*authMethod = shared.SourceLinkedinAdsUpdateAuthenticationOAuth20AuthMethod(r.Configuration.Credentials.SourceLinkedinAdsUpdateAuthenticationOAuth20.AuthMethod.ValueString())
			} else {
				authMethod = nil
			}
			clientID := r.Configuration.Credentials.SourceLinkedinAdsUpdateAuthenticationOAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceLinkedinAdsUpdateAuthenticationOAuth20.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SourceLinkedinAdsUpdateAuthenticationOAuth20.RefreshToken.ValueString()
			sourceLinkedinAdsUpdateAuthenticationOAuth20 = &shared.SourceLinkedinAdsUpdateAuthenticationOAuth20{
				AuthMethod:   authMethod,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceLinkedinAdsUpdateAuthenticationOAuth20 != nil {
			credentials = &shared.SourceLinkedinAdsUpdateAuthentication{
				SourceLinkedinAdsUpdateAuthenticationOAuth20: sourceLinkedinAdsUpdateAuthenticationOAuth20,
			}
		}
		var sourceLinkedinAdsUpdateAuthenticationAccessToken *shared.SourceLinkedinAdsUpdateAuthenticationAccessToken
		if r.Configuration.Credentials.SourceLinkedinAdsUpdateAuthenticationAccessToken != nil {
			accessToken := r.Configuration.Credentials.SourceLinkedinAdsUpdateAuthenticationAccessToken.AccessToken.ValueString()
			authMethod1 := new(shared.SourceLinkedinAdsUpdateAuthenticationAccessTokenAuthMethod)
			if !r.Configuration.Credentials.SourceLinkedinAdsUpdateAuthenticationAccessToken.AuthMethod.IsUnknown() && !r.Configuration.Credentials.SourceLinkedinAdsUpdateAuthenticationAccessToken.AuthMethod.IsNull() {
				*authMethod1 = shared.SourceLinkedinAdsUpdateAuthenticationAccessTokenAuthMethod(r.Configuration.Credentials.SourceLinkedinAdsUpdateAuthenticationAccessToken.AuthMethod.ValueString())
			} else {
				authMethod1 = nil
			}
			sourceLinkedinAdsUpdateAuthenticationAccessToken = &shared.SourceLinkedinAdsUpdateAuthenticationAccessToken{
				AccessToken: accessToken,
				AuthMethod:  authMethod1,
			}
		}
		if sourceLinkedinAdsUpdateAuthenticationAccessToken != nil {
			credentials = &shared.SourceLinkedinAdsUpdateAuthentication{
				SourceLinkedinAdsUpdateAuthenticationAccessToken: sourceLinkedinAdsUpdateAuthenticationAccessToken,
			}
		}
	}
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	configuration := shared.SourceLinkedinAdsUpdate{
		AccountIds:         accountIds,
		AdAnalyticsReports: adAnalyticsReports,
		Credentials:        credentials,
		StartDate:          startDate,
	}
	name1 := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceLinkedinAdsPutRequest{
		Configuration: configuration,
		Name:          name1,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceLinkedinAdsResourceModel) ToDeleteSDKType() *shared.SourceLinkedinAdsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceLinkedinAdsResourceModel) RefreshFromGetResponse(resp *shared.SourceLinkedinAdsGetResponse) {
	r.Configuration.AccountIds = nil
	for _, v := range resp.Configuration.AccountIds {
		r.Configuration.AccountIds = append(r.Configuration.AccountIds, types.Int64Value(v))
	}
	r.Configuration.AdAnalyticsReports = nil
	for _, adAnalyticsReportsItem := range resp.Configuration.AdAnalyticsReports {
		var adAnalyticsReports1 SourceLinkedinAdsAdAnalyticsReportConfiguration
		adAnalyticsReports1.Name = types.StringValue(adAnalyticsReportsItem.Name)
		adAnalyticsReports1.PivotBy = types.StringValue(string(adAnalyticsReportsItem.PivotBy))
		adAnalyticsReports1.TimeGranularity = types.StringValue(string(adAnalyticsReportsItem.TimeGranularity))
		r.Configuration.AdAnalyticsReports = append(r.Configuration.AdAnalyticsReports, adAnalyticsReports1)
	}
	if resp.Configuration.Credentials == nil {
		r.Configuration.Credentials = nil
	} else {
		r.Configuration.Credentials = &SourceLinkedinAdsAuthentication{}
		if resp.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken != nil {
			r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken = &SourceLinkedinAdsAuthenticationAccessToken{}
			r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken.AccessToken = types.StringValue(resp.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken.AccessToken)
			if resp.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken.AuthMethod != nil {
				r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken.AuthMethod = types.StringValue(string(*resp.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken.AuthMethod))
			} else {
				r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken.AuthMethod = types.StringNull()
			}
		}
		if resp.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20 != nil {
			r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20 = &SourceLinkedinAdsAuthenticationOAuth20{}
			if resp.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.AuthMethod != nil {
				r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.AuthMethod = types.StringValue(string(*resp.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.AuthMethod))
			} else {
				r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.AuthMethod = types.StringNull()
			}
			r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.ClientID = types.StringValue(resp.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.ClientID)
			r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.ClientSecret = types.StringValue(resp.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.ClientSecret)
			r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.RefreshToken = types.StringValue(resp.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.RefreshToken)
		}
		if resp.Configuration.Credentials.SourceLinkedinAdsUpdateAuthenticationAccessToken != nil {
			r.Configuration.Credentials.SourceLinkedinAdsUpdateAuthenticationAccessToken = &SourceLinkedinAdsAuthenticationAccessToken{}
		}
		if resp.Configuration.Credentials.SourceLinkedinAdsUpdateAuthenticationOAuth20 != nil {
			r.Configuration.Credentials.SourceLinkedinAdsUpdateAuthenticationOAuth20 = &SourceLinkedinAdsAuthenticationOAuth20{}
		}
	}
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	r.Configuration.StartDate = types.StringValue(resp.Configuration.StartDate.String())
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

func (r *SourceLinkedinAdsResourceModel) RefreshFromCreateResponse(resp *shared.SourceLinkedinAdsGetResponse) {
	r.RefreshFromGetResponse(resp)
}
