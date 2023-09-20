// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourcePolygonStockAPIDataSourceModel) RefreshFromGetResponse(resp *shared.SourcePolygonStockAPIGetResponse) {
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
