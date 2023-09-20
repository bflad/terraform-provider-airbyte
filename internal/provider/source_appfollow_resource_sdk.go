// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceAppfollowResourceModel) ToCreateSDKType() *shared.SourceAppfollowCreateRequest {
	apiSecret := new(string)
	if !r.Configuration.APISecret.IsUnknown() && !r.Configuration.APISecret.IsNull() {
		*apiSecret = r.Configuration.APISecret.ValueString()
	} else {
		apiSecret = nil
	}
	sourceType := shared.SourceAppfollowAppfollow(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceAppfollow{
		APISecret:  apiSecret,
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
	out := shared.SourceAppfollowCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAppfollowResourceModel) ToGetSDKType() *shared.SourceAppfollowCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAppfollowResourceModel) ToUpdateSDKType() *shared.SourceAppfollowPutRequest {
	apiSecret := new(string)
	if !r.Configuration.APISecret.IsUnknown() && !r.Configuration.APISecret.IsNull() {
		*apiSecret = r.Configuration.APISecret.ValueString()
	} else {
		apiSecret = nil
	}
	configuration := shared.SourceAppfollowUpdate{
		APISecret: apiSecret,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAppfollowPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAppfollowResourceModel) ToDeleteSDKType() *shared.SourceAppfollowCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAppfollowResourceModel) RefreshFromGetResponse(resp *shared.SourceAppfollowGetResponse) {
	if resp.Configuration.APISecret != nil {
		r.Configuration.APISecret = types.StringValue(*resp.Configuration.APISecret)
	} else {
		r.Configuration.APISecret = types.StringNull()
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

func (r *SourceAppfollowResourceModel) RefreshFromCreateResponse(resp *shared.SourceAppfollowGetResponse) {
	r.RefreshFromGetResponse(resp)
}
