// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceRecurlyResourceModel) ToCreateSDKType() *shared.SourceRecurlyCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	beginTime := new(string)
	if !r.Configuration.BeginTime.IsUnknown() && !r.Configuration.BeginTime.IsNull() {
		*beginTime = r.Configuration.BeginTime.ValueString()
	} else {
		beginTime = nil
	}
	endTime := new(string)
	if !r.Configuration.EndTime.IsUnknown() && !r.Configuration.EndTime.IsNull() {
		*endTime = r.Configuration.EndTime.ValueString()
	} else {
		endTime = nil
	}
	sourceType := shared.SourceRecurlyRecurly(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceRecurly{
		APIKey:     apiKey,
		BeginTime:  beginTime,
		EndTime:    endTime,
		SourceType: sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceRecurlyCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceRecurlyResourceModel) ToGetSDKType() *shared.SourceRecurlyCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceRecurlyResourceModel) ToUpdateSDKType() *shared.SourceRecurlyPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	beginTime := new(string)
	if !r.Configuration.BeginTime.IsUnknown() && !r.Configuration.BeginTime.IsNull() {
		*beginTime = r.Configuration.BeginTime.ValueString()
	} else {
		beginTime = nil
	}
	endTime := new(string)
	if !r.Configuration.EndTime.IsUnknown() && !r.Configuration.EndTime.IsNull() {
		*endTime = r.Configuration.EndTime.ValueString()
	} else {
		endTime = nil
	}
	configuration := shared.SourceRecurlyUpdate{
		APIKey:    apiKey,
		BeginTime: beginTime,
		EndTime:   endTime,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceRecurlyPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceRecurlyResourceModel) ToDeleteSDKType() *shared.SourceRecurlyCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceRecurlyResourceModel) RefreshFromGetResponse(resp *shared.SourceRecurlyGetResponse) {
	r.Configuration.APIKey = types.StringValue(resp.Configuration.APIKey)
	if resp.Configuration.BeginTime != nil {
		r.Configuration.BeginTime = types.StringValue(*resp.Configuration.BeginTime)
	} else {
		r.Configuration.BeginTime = types.StringNull()
	}
	if resp.Configuration.EndTime != nil {
		r.Configuration.EndTime = types.StringValue(*resp.Configuration.EndTime)
	} else {
		r.Configuration.EndTime = types.StringNull()
	}
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
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

func (r *SourceRecurlyResourceModel) RefreshFromCreateResponse(resp *shared.SourceRecurlyGetResponse) {
	r.RefreshFromGetResponse(resp)
}
