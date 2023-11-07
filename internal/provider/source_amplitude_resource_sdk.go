// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceAmplitudeResourceModel) ToCreateSDKType() *shared.SourceAmplitudeCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	dataRegion := new(shared.SourceAmplitudeDataRegion)
	if !r.Configuration.DataRegion.IsUnknown() && !r.Configuration.DataRegion.IsNull() {
		*dataRegion = shared.SourceAmplitudeDataRegion(r.Configuration.DataRegion.ValueString())
	} else {
		dataRegion = nil
	}
	requestTimeRange := new(int64)
	if !r.Configuration.RequestTimeRange.IsUnknown() && !r.Configuration.RequestTimeRange.IsNull() {
		*requestTimeRange = r.Configuration.RequestTimeRange.ValueInt64()
	} else {
		requestTimeRange = nil
	}
	secretKey := r.Configuration.SecretKey.ValueString()
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceAmplitude{
		APIKey:           apiKey,
		DataRegion:       dataRegion,
		RequestTimeRange: requestTimeRange,
		SecretKey:        secretKey,
		StartDate:        startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAmplitudeCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAmplitudeResourceModel) ToGetSDKType() *shared.SourceAmplitudeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAmplitudeResourceModel) ToUpdateSDKType() *shared.SourceAmplitudePutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	dataRegion := new(shared.DataRegion)
	if !r.Configuration.DataRegion.IsUnknown() && !r.Configuration.DataRegion.IsNull() {
		*dataRegion = shared.DataRegion(r.Configuration.DataRegion.ValueString())
	} else {
		dataRegion = nil
	}
	requestTimeRange := new(int64)
	if !r.Configuration.RequestTimeRange.IsUnknown() && !r.Configuration.RequestTimeRange.IsNull() {
		*requestTimeRange = r.Configuration.RequestTimeRange.ValueInt64()
	} else {
		requestTimeRange = nil
	}
	secretKey := r.Configuration.SecretKey.ValueString()
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceAmplitudeUpdate{
		APIKey:           apiKey,
		DataRegion:       dataRegion,
		RequestTimeRange: requestTimeRange,
		SecretKey:        secretKey,
		StartDate:        startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAmplitudePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAmplitudeResourceModel) ToDeleteSDKType() *shared.SourceAmplitudeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAmplitudeResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceAmplitudeResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
