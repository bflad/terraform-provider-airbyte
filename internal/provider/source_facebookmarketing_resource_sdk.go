// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceFacebookMarketingResourceModel) ToCreateSDKType() *shared.SourceFacebookMarketingCreateRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	accountID := r.Configuration.AccountID.ValueString()
	actionBreakdownsAllowEmpty := new(bool)
	if !r.Configuration.ActionBreakdownsAllowEmpty.IsUnknown() && !r.Configuration.ActionBreakdownsAllowEmpty.IsNull() {
		*actionBreakdownsAllowEmpty = r.Configuration.ActionBreakdownsAllowEmpty.ValueBool()
	} else {
		actionBreakdownsAllowEmpty = nil
	}
	clientID := new(string)
	if !r.Configuration.ClientID.IsUnknown() && !r.Configuration.ClientID.IsNull() {
		*clientID = r.Configuration.ClientID.ValueString()
	} else {
		clientID = nil
	}
	clientSecret := new(string)
	if !r.Configuration.ClientSecret.IsUnknown() && !r.Configuration.ClientSecret.IsNull() {
		*clientSecret = r.Configuration.ClientSecret.ValueString()
	} else {
		clientSecret = nil
	}
	var customInsights []shared.SourceFacebookMarketingInsightConfig = nil
	for _, customInsightsItem := range r.Configuration.CustomInsights {
		var actionBreakdowns []shared.SourceFacebookMarketingInsightConfigValidActionBreakdowns = nil
		for _, actionBreakdownsItem := range customInsightsItem.ActionBreakdowns {
			actionBreakdowns = append(actionBreakdowns, shared.SourceFacebookMarketingInsightConfigValidActionBreakdowns(actionBreakdownsItem.ValueString()))
		}
		actionReportTime := new(shared.SourceFacebookMarketingInsightConfigActionReportTime)
		if !customInsightsItem.ActionReportTime.IsUnknown() && !customInsightsItem.ActionReportTime.IsNull() {
			*actionReportTime = shared.SourceFacebookMarketingInsightConfigActionReportTime(customInsightsItem.ActionReportTime.ValueString())
		} else {
			actionReportTime = nil
		}
		var breakdowns []shared.SourceFacebookMarketingInsightConfigValidBreakdowns = nil
		for _, breakdownsItem := range customInsightsItem.Breakdowns {
			breakdowns = append(breakdowns, shared.SourceFacebookMarketingInsightConfigValidBreakdowns(breakdownsItem.ValueString()))
		}
		endDate := new(time.Time)
		if !customInsightsItem.EndDate.IsUnknown() && !customInsightsItem.EndDate.IsNull() {
			*endDate, _ = time.Parse(time.RFC3339Nano, customInsightsItem.EndDate.ValueString())
		} else {
			endDate = nil
		}
		var fields []shared.SourceFacebookMarketingInsightConfigValidEnums = nil
		for _, fieldsItem := range customInsightsItem.Fields {
			fields = append(fields, shared.SourceFacebookMarketingInsightConfigValidEnums(fieldsItem.ValueString()))
		}
		insightsLookbackWindow := new(int64)
		if !customInsightsItem.InsightsLookbackWindow.IsUnknown() && !customInsightsItem.InsightsLookbackWindow.IsNull() {
			*insightsLookbackWindow = customInsightsItem.InsightsLookbackWindow.ValueInt64()
		} else {
			insightsLookbackWindow = nil
		}
		level := new(shared.SourceFacebookMarketingInsightConfigLevel)
		if !customInsightsItem.Level.IsUnknown() && !customInsightsItem.Level.IsNull() {
			*level = shared.SourceFacebookMarketingInsightConfigLevel(customInsightsItem.Level.ValueString())
		} else {
			level = nil
		}
		name := customInsightsItem.Name.ValueString()
		startDate := new(time.Time)
		if !customInsightsItem.StartDate.IsUnknown() && !customInsightsItem.StartDate.IsNull() {
			*startDate, _ = time.Parse(time.RFC3339Nano, customInsightsItem.StartDate.ValueString())
		} else {
			startDate = nil
		}
		timeIncrement := new(int64)
		if !customInsightsItem.TimeIncrement.IsUnknown() && !customInsightsItem.TimeIncrement.IsNull() {
			*timeIncrement = customInsightsItem.TimeIncrement.ValueInt64()
		} else {
			timeIncrement = nil
		}
		customInsights = append(customInsights, shared.SourceFacebookMarketingInsightConfig{
			ActionBreakdowns:       actionBreakdowns,
			ActionReportTime:       actionReportTime,
			Breakdowns:             breakdowns,
			EndDate:                endDate,
			Fields:                 fields,
			InsightsLookbackWindow: insightsLookbackWindow,
			Level:                  level,
			Name:                   name,
			StartDate:              startDate,
			TimeIncrement:          timeIncrement,
		})
	}
	endDate1 := new(time.Time)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate1, _ = time.Parse(time.RFC3339Nano, r.Configuration.EndDate.ValueString())
	} else {
		endDate1 = nil
	}
	fetchThumbnailImages := new(bool)
	if !r.Configuration.FetchThumbnailImages.IsUnknown() && !r.Configuration.FetchThumbnailImages.IsNull() {
		*fetchThumbnailImages = r.Configuration.FetchThumbnailImages.ValueBool()
	} else {
		fetchThumbnailImages = nil
	}
	includeDeleted := new(bool)
	if !r.Configuration.IncludeDeleted.IsUnknown() && !r.Configuration.IncludeDeleted.IsNull() {
		*includeDeleted = r.Configuration.IncludeDeleted.ValueBool()
	} else {
		includeDeleted = nil
	}
	insightsLookbackWindow1 := new(int64)
	if !r.Configuration.InsightsLookbackWindow.IsUnknown() && !r.Configuration.InsightsLookbackWindow.IsNull() {
		*insightsLookbackWindow1 = r.Configuration.InsightsLookbackWindow.ValueInt64()
	} else {
		insightsLookbackWindow1 = nil
	}
	maxBatchSize := new(int64)
	if !r.Configuration.MaxBatchSize.IsUnknown() && !r.Configuration.MaxBatchSize.IsNull() {
		*maxBatchSize = r.Configuration.MaxBatchSize.ValueInt64()
	} else {
		maxBatchSize = nil
	}
	pageSize := new(int64)
	if !r.Configuration.PageSize.IsUnknown() && !r.Configuration.PageSize.IsNull() {
		*pageSize = r.Configuration.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	sourceType := shared.SourceFacebookMarketingFacebookMarketing(r.Configuration.SourceType.ValueString())
	startDate1, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceFacebookMarketing{
		AccessToken:                accessToken,
		AccountID:                  accountID,
		ActionBreakdownsAllowEmpty: actionBreakdownsAllowEmpty,
		ClientID:                   clientID,
		ClientSecret:               clientSecret,
		CustomInsights:             customInsights,
		EndDate:                    endDate1,
		FetchThumbnailImages:       fetchThumbnailImages,
		IncludeDeleted:             includeDeleted,
		InsightsLookbackWindow:     insightsLookbackWindow1,
		MaxBatchSize:               maxBatchSize,
		PageSize:                   pageSize,
		SourceType:                 sourceType,
		StartDate:                  startDate1,
	}
	name1 := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFacebookMarketingCreateRequest{
		Configuration: configuration,
		Name:          name1,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFacebookMarketingResourceModel) ToGetSDKType() *shared.SourceFacebookMarketingCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFacebookMarketingResourceModel) ToUpdateSDKType() *shared.SourceFacebookMarketingPutRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	accountID := r.Configuration.AccountID.ValueString()
	actionBreakdownsAllowEmpty := new(bool)
	if !r.Configuration.ActionBreakdownsAllowEmpty.IsUnknown() && !r.Configuration.ActionBreakdownsAllowEmpty.IsNull() {
		*actionBreakdownsAllowEmpty = r.Configuration.ActionBreakdownsAllowEmpty.ValueBool()
	} else {
		actionBreakdownsAllowEmpty = nil
	}
	clientID := new(string)
	if !r.Configuration.ClientID.IsUnknown() && !r.Configuration.ClientID.IsNull() {
		*clientID = r.Configuration.ClientID.ValueString()
	} else {
		clientID = nil
	}
	clientSecret := new(string)
	if !r.Configuration.ClientSecret.IsUnknown() && !r.Configuration.ClientSecret.IsNull() {
		*clientSecret = r.Configuration.ClientSecret.ValueString()
	} else {
		clientSecret = nil
	}
	var customInsights []shared.SourceFacebookMarketingUpdateInsightConfig = nil
	for _, customInsightsItem := range r.Configuration.CustomInsights {
		var actionBreakdowns []shared.SourceFacebookMarketingUpdateInsightConfigValidActionBreakdowns = nil
		for _, actionBreakdownsItem := range customInsightsItem.ActionBreakdowns {
			actionBreakdowns = append(actionBreakdowns, shared.SourceFacebookMarketingUpdateInsightConfigValidActionBreakdowns(actionBreakdownsItem.ValueString()))
		}
		actionReportTime := new(shared.SourceFacebookMarketingUpdateInsightConfigActionReportTime)
		if !customInsightsItem.ActionReportTime.IsUnknown() && !customInsightsItem.ActionReportTime.IsNull() {
			*actionReportTime = shared.SourceFacebookMarketingUpdateInsightConfigActionReportTime(customInsightsItem.ActionReportTime.ValueString())
		} else {
			actionReportTime = nil
		}
		var breakdowns []shared.SourceFacebookMarketingUpdateInsightConfigValidBreakdowns = nil
		for _, breakdownsItem := range customInsightsItem.Breakdowns {
			breakdowns = append(breakdowns, shared.SourceFacebookMarketingUpdateInsightConfigValidBreakdowns(breakdownsItem.ValueString()))
		}
		endDate := new(time.Time)
		if !customInsightsItem.EndDate.IsUnknown() && !customInsightsItem.EndDate.IsNull() {
			*endDate, _ = time.Parse(time.RFC3339Nano, customInsightsItem.EndDate.ValueString())
		} else {
			endDate = nil
		}
		var fields []shared.SourceFacebookMarketingUpdateInsightConfigValidEnums = nil
		for _, fieldsItem := range customInsightsItem.Fields {
			fields = append(fields, shared.SourceFacebookMarketingUpdateInsightConfigValidEnums(fieldsItem.ValueString()))
		}
		insightsLookbackWindow := new(int64)
		if !customInsightsItem.InsightsLookbackWindow.IsUnknown() && !customInsightsItem.InsightsLookbackWindow.IsNull() {
			*insightsLookbackWindow = customInsightsItem.InsightsLookbackWindow.ValueInt64()
		} else {
			insightsLookbackWindow = nil
		}
		level := new(shared.SourceFacebookMarketingUpdateInsightConfigLevel)
		if !customInsightsItem.Level.IsUnknown() && !customInsightsItem.Level.IsNull() {
			*level = shared.SourceFacebookMarketingUpdateInsightConfigLevel(customInsightsItem.Level.ValueString())
		} else {
			level = nil
		}
		name := customInsightsItem.Name.ValueString()
		startDate := new(time.Time)
		if !customInsightsItem.StartDate.IsUnknown() && !customInsightsItem.StartDate.IsNull() {
			*startDate, _ = time.Parse(time.RFC3339Nano, customInsightsItem.StartDate.ValueString())
		} else {
			startDate = nil
		}
		timeIncrement := new(int64)
		if !customInsightsItem.TimeIncrement.IsUnknown() && !customInsightsItem.TimeIncrement.IsNull() {
			*timeIncrement = customInsightsItem.TimeIncrement.ValueInt64()
		} else {
			timeIncrement = nil
		}
		customInsights = append(customInsights, shared.SourceFacebookMarketingUpdateInsightConfig{
			ActionBreakdowns:       actionBreakdowns,
			ActionReportTime:       actionReportTime,
			Breakdowns:             breakdowns,
			EndDate:                endDate,
			Fields:                 fields,
			InsightsLookbackWindow: insightsLookbackWindow,
			Level:                  level,
			Name:                   name,
			StartDate:              startDate,
			TimeIncrement:          timeIncrement,
		})
	}
	endDate1 := new(time.Time)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate1, _ = time.Parse(time.RFC3339Nano, r.Configuration.EndDate.ValueString())
	} else {
		endDate1 = nil
	}
	fetchThumbnailImages := new(bool)
	if !r.Configuration.FetchThumbnailImages.IsUnknown() && !r.Configuration.FetchThumbnailImages.IsNull() {
		*fetchThumbnailImages = r.Configuration.FetchThumbnailImages.ValueBool()
	} else {
		fetchThumbnailImages = nil
	}
	includeDeleted := new(bool)
	if !r.Configuration.IncludeDeleted.IsUnknown() && !r.Configuration.IncludeDeleted.IsNull() {
		*includeDeleted = r.Configuration.IncludeDeleted.ValueBool()
	} else {
		includeDeleted = nil
	}
	insightsLookbackWindow1 := new(int64)
	if !r.Configuration.InsightsLookbackWindow.IsUnknown() && !r.Configuration.InsightsLookbackWindow.IsNull() {
		*insightsLookbackWindow1 = r.Configuration.InsightsLookbackWindow.ValueInt64()
	} else {
		insightsLookbackWindow1 = nil
	}
	maxBatchSize := new(int64)
	if !r.Configuration.MaxBatchSize.IsUnknown() && !r.Configuration.MaxBatchSize.IsNull() {
		*maxBatchSize = r.Configuration.MaxBatchSize.ValueInt64()
	} else {
		maxBatchSize = nil
	}
	pageSize := new(int64)
	if !r.Configuration.PageSize.IsUnknown() && !r.Configuration.PageSize.IsNull() {
		*pageSize = r.Configuration.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	startDate1, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceFacebookMarketingUpdate{
		AccessToken:                accessToken,
		AccountID:                  accountID,
		ActionBreakdownsAllowEmpty: actionBreakdownsAllowEmpty,
		ClientID:                   clientID,
		ClientSecret:               clientSecret,
		CustomInsights:             customInsights,
		EndDate:                    endDate1,
		FetchThumbnailImages:       fetchThumbnailImages,
		IncludeDeleted:             includeDeleted,
		InsightsLookbackWindow:     insightsLookbackWindow1,
		MaxBatchSize:               maxBatchSize,
		PageSize:                   pageSize,
		StartDate:                  startDate1,
	}
	name1 := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFacebookMarketingPutRequest{
		Configuration: configuration,
		Name:          name1,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFacebookMarketingResourceModel) ToDeleteSDKType() *shared.SourceFacebookMarketingCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFacebookMarketingResourceModel) RefreshFromGetResponse(resp *shared.SourceFacebookMarketingGetResponse) {
	r.Configuration.AccessToken = types.StringValue(resp.Configuration.AccessToken)
	r.Configuration.AccountID = types.StringValue(resp.Configuration.AccountID)
	if resp.Configuration.ActionBreakdownsAllowEmpty != nil {
		r.Configuration.ActionBreakdownsAllowEmpty = types.BoolValue(*resp.Configuration.ActionBreakdownsAllowEmpty)
	} else {
		r.Configuration.ActionBreakdownsAllowEmpty = types.BoolNull()
	}
	if resp.Configuration.ClientID != nil {
		r.Configuration.ClientID = types.StringValue(*resp.Configuration.ClientID)
	} else {
		r.Configuration.ClientID = types.StringNull()
	}
	if resp.Configuration.ClientSecret != nil {
		r.Configuration.ClientSecret = types.StringValue(*resp.Configuration.ClientSecret)
	} else {
		r.Configuration.ClientSecret = types.StringNull()
	}
	r.Configuration.CustomInsights = nil
	for _, customInsightsItem := range resp.Configuration.CustomInsights {
		var customInsights1 SourceFacebookMarketingInsightConfig
		customInsights1.ActionBreakdowns = nil
		for _, v := range customInsightsItem.ActionBreakdowns {
			customInsights1.ActionBreakdowns = append(customInsights1.ActionBreakdowns, types.StringValue(string(v)))
		}
		if customInsightsItem.ActionReportTime != nil {
			customInsights1.ActionReportTime = types.StringValue(string(*customInsightsItem.ActionReportTime))
		} else {
			customInsights1.ActionReportTime = types.StringNull()
		}
		customInsights1.Breakdowns = nil
		for _, v := range customInsightsItem.Breakdowns {
			customInsights1.Breakdowns = append(customInsights1.Breakdowns, types.StringValue(string(v)))
		}
		if customInsightsItem.EndDate != nil {
			customInsights1.EndDate = types.StringValue(customInsightsItem.EndDate.Format(time.RFC3339))
		} else {
			customInsights1.EndDate = types.StringNull()
		}
		customInsights1.Fields = nil
		for _, v := range customInsightsItem.Fields {
			customInsights1.Fields = append(customInsights1.Fields, types.StringValue(string(v)))
		}
		if customInsightsItem.InsightsLookbackWindow != nil {
			customInsights1.InsightsLookbackWindow = types.Int64Value(*customInsightsItem.InsightsLookbackWindow)
		} else {
			customInsights1.InsightsLookbackWindow = types.Int64Null()
		}
		if customInsightsItem.Level != nil {
			customInsights1.Level = types.StringValue(string(*customInsightsItem.Level))
		} else {
			customInsights1.Level = types.StringNull()
		}
		customInsights1.Name = types.StringValue(customInsightsItem.Name)
		if customInsightsItem.StartDate != nil {
			customInsights1.StartDate = types.StringValue(customInsightsItem.StartDate.Format(time.RFC3339))
		} else {
			customInsights1.StartDate = types.StringNull()
		}
		if customInsightsItem.TimeIncrement != nil {
			customInsights1.TimeIncrement = types.Int64Value(*customInsightsItem.TimeIncrement)
		} else {
			customInsights1.TimeIncrement = types.Int64Null()
		}
		r.Configuration.CustomInsights = append(r.Configuration.CustomInsights, customInsights1)
	}
	if resp.Configuration.EndDate != nil {
		r.Configuration.EndDate = types.StringValue(resp.Configuration.EndDate.Format(time.RFC3339))
	} else {
		r.Configuration.EndDate = types.StringNull()
	}
	if resp.Configuration.FetchThumbnailImages != nil {
		r.Configuration.FetchThumbnailImages = types.BoolValue(*resp.Configuration.FetchThumbnailImages)
	} else {
		r.Configuration.FetchThumbnailImages = types.BoolNull()
	}
	if resp.Configuration.IncludeDeleted != nil {
		r.Configuration.IncludeDeleted = types.BoolValue(*resp.Configuration.IncludeDeleted)
	} else {
		r.Configuration.IncludeDeleted = types.BoolNull()
	}
	if resp.Configuration.InsightsLookbackWindow != nil {
		r.Configuration.InsightsLookbackWindow = types.Int64Value(*resp.Configuration.InsightsLookbackWindow)
	} else {
		r.Configuration.InsightsLookbackWindow = types.Int64Null()
	}
	if resp.Configuration.MaxBatchSize != nil {
		r.Configuration.MaxBatchSize = types.Int64Value(*resp.Configuration.MaxBatchSize)
	} else {
		r.Configuration.MaxBatchSize = types.Int64Null()
	}
	if resp.Configuration.PageSize != nil {
		r.Configuration.PageSize = types.Int64Value(*resp.Configuration.PageSize)
	} else {
		r.Configuration.PageSize = types.Int64Null()
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

func (r *SourceFacebookMarketingResourceModel) RefreshFromCreateResponse(resp *shared.SourceFacebookMarketingGetResponse) {
	r.RefreshFromGetResponse(resp)
}
