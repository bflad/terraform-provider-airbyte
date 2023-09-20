// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	customTypes "airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourcePolygonStockAPIResourceModel) ToCreateSDKType() *shared.SourcePolygonStockAPICreateRequest {
	adjusted := new(string)
	if !r.Configuration.Adjusted.IsUnknown() && !r.Configuration.Adjusted.IsNull() {
		*adjusted = r.Configuration.Adjusted.ValueString()
	} else {
		adjusted = nil
	}
	apiKey := r.Configuration.APIKey.ValueString()
	endDate := customTypes.MustDateFromString(r.Configuration.EndDate.ValueString())
	limit := new(int64)
	if !r.Configuration.Limit.IsUnknown() && !r.Configuration.Limit.IsNull() {
		*limit = r.Configuration.Limit.ValueInt64()
	} else {
		limit = nil
	}
	multiplier := r.Configuration.Multiplier.ValueInt64()
	sort := new(string)
	if !r.Configuration.Sort.IsUnknown() && !r.Configuration.Sort.IsNull() {
		*sort = r.Configuration.Sort.ValueString()
	} else {
		sort = nil
	}
	sourceType := shared.SourcePolygonStockAPIPolygonStockAPI(r.Configuration.SourceType.ValueString())
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	stocksTicker := r.Configuration.StocksTicker.ValueString()
	timespan := r.Configuration.Timespan.ValueString()
	configuration := shared.SourcePolygonStockAPI{
		Adjusted:     adjusted,
		APIKey:       apiKey,
		EndDate:      endDate,
		Limit:        limit,
		Multiplier:   multiplier,
		Sort:         sort,
		SourceType:   sourceType,
		StartDate:    startDate,
		StocksTicker: stocksTicker,
		Timespan:     timespan,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePolygonStockAPICreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePolygonStockAPIResourceModel) ToGetSDKType() *shared.SourcePolygonStockAPICreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePolygonStockAPIResourceModel) ToUpdateSDKType() *shared.SourcePolygonStockAPIPutRequest {
	adjusted := new(string)
	if !r.Configuration.Adjusted.IsUnknown() && !r.Configuration.Adjusted.IsNull() {
		*adjusted = r.Configuration.Adjusted.ValueString()
	} else {
		adjusted = nil
	}
	apiKey := r.Configuration.APIKey.ValueString()
	endDate := customTypes.MustDateFromString(r.Configuration.EndDate.ValueString())
	limit := new(int64)
	if !r.Configuration.Limit.IsUnknown() && !r.Configuration.Limit.IsNull() {
		*limit = r.Configuration.Limit.ValueInt64()
	} else {
		limit = nil
	}
	multiplier := r.Configuration.Multiplier.ValueInt64()
	sort := new(string)
	if !r.Configuration.Sort.IsUnknown() && !r.Configuration.Sort.IsNull() {
		*sort = r.Configuration.Sort.ValueString()
	} else {
		sort = nil
	}
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	stocksTicker := r.Configuration.StocksTicker.ValueString()
	timespan := r.Configuration.Timespan.ValueString()
	configuration := shared.SourcePolygonStockAPIUpdate{
		Adjusted:     adjusted,
		APIKey:       apiKey,
		EndDate:      endDate,
		Limit:        limit,
		Multiplier:   multiplier,
		Sort:         sort,
		StartDate:    startDate,
		StocksTicker: stocksTicker,
		Timespan:     timespan,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePolygonStockAPIPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePolygonStockAPIResourceModel) ToDeleteSDKType() *shared.SourcePolygonStockAPICreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePolygonStockAPIResourceModel) RefreshFromGetResponse(resp *shared.SourcePolygonStockAPIGetResponse) {
	if resp.Configuration.Adjusted != nil {
		r.Configuration.Adjusted = types.StringValue(*resp.Configuration.Adjusted)
	} else {
		r.Configuration.Adjusted = types.StringNull()
	}
	r.Configuration.APIKey = types.StringValue(resp.Configuration.APIKey)
	r.Configuration.EndDate = types.StringValue(resp.Configuration.EndDate.String())
	if resp.Configuration.Limit != nil {
		r.Configuration.Limit = types.Int64Value(*resp.Configuration.Limit)
	} else {
		r.Configuration.Limit = types.Int64Null()
	}
	r.Configuration.Multiplier = types.Int64Value(resp.Configuration.Multiplier)
	if resp.Configuration.Sort != nil {
		r.Configuration.Sort = types.StringValue(*resp.Configuration.Sort)
	} else {
		r.Configuration.Sort = types.StringNull()
	}
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	r.Configuration.StartDate = types.StringValue(resp.Configuration.StartDate.String())
	r.Configuration.StocksTicker = types.StringValue(resp.Configuration.StocksTicker)
	r.Configuration.Timespan = types.StringValue(resp.Configuration.Timespan)
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

func (r *SourcePolygonStockAPIResourceModel) RefreshFromCreateResponse(resp *shared.SourcePolygonStockAPIGetResponse) {
	r.RefreshFromGetResponse(resp)
}
