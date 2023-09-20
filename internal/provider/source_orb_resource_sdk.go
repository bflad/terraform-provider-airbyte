// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceOrbResourceModel) ToCreateSDKType() *shared.SourceOrbCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	lookbackWindowDays := new(int64)
	if !r.Configuration.LookbackWindowDays.IsUnknown() && !r.Configuration.LookbackWindowDays.IsNull() {
		*lookbackWindowDays = r.Configuration.LookbackWindowDays.ValueInt64()
	} else {
		lookbackWindowDays = nil
	}
	var numericEventPropertiesKeys []string = nil
	for _, numericEventPropertiesKeysItem := range r.Configuration.NumericEventPropertiesKeys {
		numericEventPropertiesKeys = append(numericEventPropertiesKeys, numericEventPropertiesKeysItem.ValueString())
	}
	planID := new(string)
	if !r.Configuration.PlanID.IsUnknown() && !r.Configuration.PlanID.IsNull() {
		*planID = r.Configuration.PlanID.ValueString()
	} else {
		planID = nil
	}
	sourceType := shared.SourceOrbOrb(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	var stringEventPropertiesKeys []string = nil
	for _, stringEventPropertiesKeysItem := range r.Configuration.StringEventPropertiesKeys {
		stringEventPropertiesKeys = append(stringEventPropertiesKeys, stringEventPropertiesKeysItem.ValueString())
	}
	subscriptionUsageGroupingKey := new(string)
	if !r.Configuration.SubscriptionUsageGroupingKey.IsUnknown() && !r.Configuration.SubscriptionUsageGroupingKey.IsNull() {
		*subscriptionUsageGroupingKey = r.Configuration.SubscriptionUsageGroupingKey.ValueString()
	} else {
		subscriptionUsageGroupingKey = nil
	}
	configuration := shared.SourceOrb{
		APIKey:                       apiKey,
		LookbackWindowDays:           lookbackWindowDays,
		NumericEventPropertiesKeys:   numericEventPropertiesKeys,
		PlanID:                       planID,
		SourceType:                   sourceType,
		StartDate:                    startDate,
		StringEventPropertiesKeys:    stringEventPropertiesKeys,
		SubscriptionUsageGroupingKey: subscriptionUsageGroupingKey,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceOrbCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOrbResourceModel) ToGetSDKType() *shared.SourceOrbCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceOrbResourceModel) ToUpdateSDKType() *shared.SourceOrbPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	lookbackWindowDays := new(int64)
	if !r.Configuration.LookbackWindowDays.IsUnknown() && !r.Configuration.LookbackWindowDays.IsNull() {
		*lookbackWindowDays = r.Configuration.LookbackWindowDays.ValueInt64()
	} else {
		lookbackWindowDays = nil
	}
	var numericEventPropertiesKeys []string = nil
	for _, numericEventPropertiesKeysItem := range r.Configuration.NumericEventPropertiesKeys {
		numericEventPropertiesKeys = append(numericEventPropertiesKeys, numericEventPropertiesKeysItem.ValueString())
	}
	planID := new(string)
	if !r.Configuration.PlanID.IsUnknown() && !r.Configuration.PlanID.IsNull() {
		*planID = r.Configuration.PlanID.ValueString()
	} else {
		planID = nil
	}
	startDate := r.Configuration.StartDate.ValueString()
	var stringEventPropertiesKeys []string = nil
	for _, stringEventPropertiesKeysItem := range r.Configuration.StringEventPropertiesKeys {
		stringEventPropertiesKeys = append(stringEventPropertiesKeys, stringEventPropertiesKeysItem.ValueString())
	}
	subscriptionUsageGroupingKey := new(string)
	if !r.Configuration.SubscriptionUsageGroupingKey.IsUnknown() && !r.Configuration.SubscriptionUsageGroupingKey.IsNull() {
		*subscriptionUsageGroupingKey = r.Configuration.SubscriptionUsageGroupingKey.ValueString()
	} else {
		subscriptionUsageGroupingKey = nil
	}
	configuration := shared.SourceOrbUpdate{
		APIKey:                       apiKey,
		LookbackWindowDays:           lookbackWindowDays,
		NumericEventPropertiesKeys:   numericEventPropertiesKeys,
		PlanID:                       planID,
		StartDate:                    startDate,
		StringEventPropertiesKeys:    stringEventPropertiesKeys,
		SubscriptionUsageGroupingKey: subscriptionUsageGroupingKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceOrbPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOrbResourceModel) ToDeleteSDKType() *shared.SourceOrbCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceOrbResourceModel) RefreshFromGetResponse(resp *shared.SourceOrbGetResponse) {
	r.Configuration.APIKey = types.StringValue(resp.Configuration.APIKey)
	if resp.Configuration.LookbackWindowDays != nil {
		r.Configuration.LookbackWindowDays = types.Int64Value(*resp.Configuration.LookbackWindowDays)
	} else {
		r.Configuration.LookbackWindowDays = types.Int64Null()
	}
	r.Configuration.NumericEventPropertiesKeys = nil
	for _, v := range resp.Configuration.NumericEventPropertiesKeys {
		r.Configuration.NumericEventPropertiesKeys = append(r.Configuration.NumericEventPropertiesKeys, types.StringValue(v))
	}
	if resp.Configuration.PlanID != nil {
		r.Configuration.PlanID = types.StringValue(*resp.Configuration.PlanID)
	} else {
		r.Configuration.PlanID = types.StringNull()
	}
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	r.Configuration.StartDate = types.StringValue(resp.Configuration.StartDate)
	r.Configuration.StringEventPropertiesKeys = nil
	for _, v := range resp.Configuration.StringEventPropertiesKeys {
		r.Configuration.StringEventPropertiesKeys = append(r.Configuration.StringEventPropertiesKeys, types.StringValue(v))
	}
	if resp.Configuration.SubscriptionUsageGroupingKey != nil {
		r.Configuration.SubscriptionUsageGroupingKey = types.StringValue(*resp.Configuration.SubscriptionUsageGroupingKey)
	} else {
		r.Configuration.SubscriptionUsageGroupingKey = types.StringNull()
	}
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

func (r *SourceOrbResourceModel) RefreshFromCreateResponse(resp *shared.SourceOrbGetResponse) {
	r.RefreshFromGetResponse(resp)
}
