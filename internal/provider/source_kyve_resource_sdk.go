// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceKyveResourceModel) ToCreateSDKType() *shared.SourceKyveCreateRequest {
	maxPages := new(int64)
	if !r.Configuration.MaxPages.IsUnknown() && !r.Configuration.MaxPages.IsNull() {
		*maxPages = r.Configuration.MaxPages.ValueInt64()
	} else {
		maxPages = nil
	}
	pageSize := new(int64)
	if !r.Configuration.PageSize.IsUnknown() && !r.Configuration.PageSize.IsNull() {
		*pageSize = r.Configuration.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	poolIds := r.Configuration.PoolIds.ValueString()
	startIds := r.Configuration.StartIds.ValueString()
	urlBase := new(string)
	if !r.Configuration.URLBase.IsUnknown() && !r.Configuration.URLBase.IsNull() {
		*urlBase = r.Configuration.URLBase.ValueString()
	} else {
		urlBase = nil
	}
	configuration := shared.SourceKyve{
		MaxPages: maxPages,
		PageSize: pageSize,
		PoolIds:  poolIds,
		StartIds: startIds,
		URLBase:  urlBase,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceKyveCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceKyveResourceModel) ToGetSDKType() *shared.SourceKyveCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceKyveResourceModel) ToUpdateSDKType() *shared.SourceKyvePutRequest {
	maxPages := new(int64)
	if !r.Configuration.MaxPages.IsUnknown() && !r.Configuration.MaxPages.IsNull() {
		*maxPages = r.Configuration.MaxPages.ValueInt64()
	} else {
		maxPages = nil
	}
	pageSize := new(int64)
	if !r.Configuration.PageSize.IsUnknown() && !r.Configuration.PageSize.IsNull() {
		*pageSize = r.Configuration.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	poolIds := r.Configuration.PoolIds.ValueString()
	startIds := r.Configuration.StartIds.ValueString()
	urlBase := new(string)
	if !r.Configuration.URLBase.IsUnknown() && !r.Configuration.URLBase.IsNull() {
		*urlBase = r.Configuration.URLBase.ValueString()
	} else {
		urlBase = nil
	}
	configuration := shared.SourceKyveUpdate{
		MaxPages: maxPages,
		PageSize: pageSize,
		PoolIds:  poolIds,
		StartIds: startIds,
		URLBase:  urlBase,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceKyvePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceKyveResourceModel) ToDeleteSDKType() *shared.SourceKyveCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceKyveResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceKyveResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
