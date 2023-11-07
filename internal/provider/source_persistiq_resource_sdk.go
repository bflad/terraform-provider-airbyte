// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourcePersistiqResourceModel) ToCreateSDKType() *shared.SourcePersistiqCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	configuration := shared.SourcePersistiq{
		APIKey: apiKey,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePersistiqCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePersistiqResourceModel) ToGetSDKType() *shared.SourcePersistiqCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePersistiqResourceModel) ToUpdateSDKType() *shared.SourcePersistiqPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	configuration := shared.SourcePersistiqUpdate{
		APIKey: apiKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePersistiqPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePersistiqResourceModel) ToDeleteSDKType() *shared.SourcePersistiqCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePersistiqResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourcePersistiqResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
