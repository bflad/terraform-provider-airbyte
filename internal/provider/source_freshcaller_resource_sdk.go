// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceFreshcallerResourceModel) ToCreateSDKType() *shared.SourceFreshcallerCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	domain := r.Configuration.Domain.ValueString()
	requestsPerMinute := new(int64)
	if !r.Configuration.RequestsPerMinute.IsUnknown() && !r.Configuration.RequestsPerMinute.IsNull() {
		*requestsPerMinute = r.Configuration.RequestsPerMinute.ValueInt64()
	} else {
		requestsPerMinute = nil
	}
	sourceType := shared.SourceFreshcallerFreshcaller(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	syncLagMinutes := new(int64)
	if !r.Configuration.SyncLagMinutes.IsUnknown() && !r.Configuration.SyncLagMinutes.IsNull() {
		*syncLagMinutes = r.Configuration.SyncLagMinutes.ValueInt64()
	} else {
		syncLagMinutes = nil
	}
	configuration := shared.SourceFreshcaller{
		APIKey:            apiKey,
		Domain:            domain,
		RequestsPerMinute: requestsPerMinute,
		SourceType:        sourceType,
		StartDate:         startDate,
		SyncLagMinutes:    syncLagMinutes,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFreshcallerCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFreshcallerResourceModel) ToGetSDKType() *shared.SourceFreshcallerCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFreshcallerResourceModel) ToUpdateSDKType() *shared.SourceFreshcallerPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	domain := r.Configuration.Domain.ValueString()
	requestsPerMinute := new(int64)
	if !r.Configuration.RequestsPerMinute.IsUnknown() && !r.Configuration.RequestsPerMinute.IsNull() {
		*requestsPerMinute = r.Configuration.RequestsPerMinute.ValueInt64()
	} else {
		requestsPerMinute = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	syncLagMinutes := new(int64)
	if !r.Configuration.SyncLagMinutes.IsUnknown() && !r.Configuration.SyncLagMinutes.IsNull() {
		*syncLagMinutes = r.Configuration.SyncLagMinutes.ValueInt64()
	} else {
		syncLagMinutes = nil
	}
	configuration := shared.SourceFreshcallerUpdate{
		APIKey:            apiKey,
		Domain:            domain,
		RequestsPerMinute: requestsPerMinute,
		StartDate:         startDate,
		SyncLagMinutes:    syncLagMinutes,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFreshcallerPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFreshcallerResourceModel) ToDeleteSDKType() *shared.SourceFreshcallerCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFreshcallerResourceModel) RefreshFromGetResponse(resp *shared.SourceFreshcallerGetResponse) {
	r.Configuration.APIKey = types.StringValue(resp.Configuration.APIKey)
	r.Configuration.Domain = types.StringValue(resp.Configuration.Domain)
	if resp.Configuration.RequestsPerMinute != nil {
		r.Configuration.RequestsPerMinute = types.Int64Value(*resp.Configuration.RequestsPerMinute)
	} else {
		r.Configuration.RequestsPerMinute = types.Int64Null()
	}
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	r.Configuration.StartDate = types.StringValue(resp.Configuration.StartDate.Format(time.RFC3339))
	if resp.Configuration.SyncLagMinutes != nil {
		r.Configuration.SyncLagMinutes = types.Int64Value(*resp.Configuration.SyncLagMinutes)
	} else {
		r.Configuration.SyncLagMinutes = types.Int64Null()
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

func (r *SourceFreshcallerResourceModel) RefreshFromCreateResponse(resp *shared.SourceFreshcallerGetResponse) {
	r.RefreshFromGetResponse(resp)
}
