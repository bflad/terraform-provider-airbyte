// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	customTypes "github.com/airbytehq/terraform-provider-airbyte/internal/sdk/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSnapchatMarketingResourceModel) ToSharedSourceSnapchatMarketingCreateRequest() *shared.SourceSnapchatMarketingCreateRequest {
	actionReportTime := new(shared.SourceSnapchatMarketingActionReportTime)
	if !r.Configuration.ActionReportTime.IsUnknown() && !r.Configuration.ActionReportTime.IsNull() {
		*actionReportTime = shared.SourceSnapchatMarketingActionReportTime(r.Configuration.ActionReportTime.ValueString())
	} else {
		actionReportTime = nil
	}
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	endDate := new(customTypes.Date)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		endDate = customTypes.MustNewDateFromString(r.Configuration.EndDate.ValueString())
	} else {
		endDate = nil
	}
	refreshToken := r.Configuration.RefreshToken.ValueString()
	startDate := new(customTypes.Date)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		startDate = customTypes.MustNewDateFromString(r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	swipeUpAttributionWindow := new(shared.SourceSnapchatMarketingSwipeUpAttributionWindow)
	if !r.Configuration.SwipeUpAttributionWindow.IsUnknown() && !r.Configuration.SwipeUpAttributionWindow.IsNull() {
		*swipeUpAttributionWindow = shared.SourceSnapchatMarketingSwipeUpAttributionWindow(r.Configuration.SwipeUpAttributionWindow.ValueString())
	} else {
		swipeUpAttributionWindow = nil
	}
	viewAttributionWindow := new(shared.SourceSnapchatMarketingViewAttributionWindow)
	if !r.Configuration.ViewAttributionWindow.IsUnknown() && !r.Configuration.ViewAttributionWindow.IsNull() {
		*viewAttributionWindow = shared.SourceSnapchatMarketingViewAttributionWindow(r.Configuration.ViewAttributionWindow.ValueString())
	} else {
		viewAttributionWindow = nil
	}
	configuration := shared.SourceSnapchatMarketing{
		ActionReportTime:         actionReportTime,
		ClientID:                 clientID,
		ClientSecret:             clientSecret,
		EndDate:                  endDate,
		RefreshToken:             refreshToken,
		StartDate:                startDate,
		SwipeUpAttributionWindow: swipeUpAttributionWindow,
		ViewAttributionWindow:    viewAttributionWindow,
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
	out := shared.SourceSnapchatMarketingCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSnapchatMarketingResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceSnapchatMarketingResourceModel) ToSharedSourceSnapchatMarketingPutRequest() *shared.SourceSnapchatMarketingPutRequest {
	actionReportTime := new(shared.ActionReportTime)
	if !r.Configuration.ActionReportTime.IsUnknown() && !r.Configuration.ActionReportTime.IsNull() {
		*actionReportTime = shared.ActionReportTime(r.Configuration.ActionReportTime.ValueString())
	} else {
		actionReportTime = nil
	}
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	endDate := new(customTypes.Date)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		endDate = customTypes.MustNewDateFromString(r.Configuration.EndDate.ValueString())
	} else {
		endDate = nil
	}
	refreshToken := r.Configuration.RefreshToken.ValueString()
	startDate := new(customTypes.Date)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		startDate = customTypes.MustNewDateFromString(r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	swipeUpAttributionWindow := new(shared.SwipeUpAttributionWindow)
	if !r.Configuration.SwipeUpAttributionWindow.IsUnknown() && !r.Configuration.SwipeUpAttributionWindow.IsNull() {
		*swipeUpAttributionWindow = shared.SwipeUpAttributionWindow(r.Configuration.SwipeUpAttributionWindow.ValueString())
	} else {
		swipeUpAttributionWindow = nil
	}
	viewAttributionWindow := new(shared.ViewAttributionWindow)
	if !r.Configuration.ViewAttributionWindow.IsUnknown() && !r.Configuration.ViewAttributionWindow.IsNull() {
		*viewAttributionWindow = shared.ViewAttributionWindow(r.Configuration.ViewAttributionWindow.ValueString())
	} else {
		viewAttributionWindow = nil
	}
	configuration := shared.SourceSnapchatMarketingUpdate{
		ActionReportTime:         actionReportTime,
		ClientID:                 clientID,
		ClientSecret:             clientSecret,
		EndDate:                  endDate,
		RefreshToken:             refreshToken,
		StartDate:                startDate,
		SwipeUpAttributionWindow: swipeUpAttributionWindow,
		ViewAttributionWindow:    viewAttributionWindow,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSnapchatMarketingPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
