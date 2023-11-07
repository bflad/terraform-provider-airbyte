// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceWhiskyHunterResourceModel) ToCreateSDKType() *shared.SourceWhiskyHunterCreateRequest {
	configuration := shared.SourceWhiskyHunter{}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceWhiskyHunterCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceWhiskyHunterResourceModel) ToGetSDKType() *shared.SourceWhiskyHunterCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceWhiskyHunterResourceModel) ToUpdateSDKType() *shared.SourceWhiskyHunterPutRequest {
	configuration := shared.SourceWhiskyHunterUpdate{}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceWhiskyHunterPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceWhiskyHunterResourceModel) ToDeleteSDKType() *shared.SourceWhiskyHunterCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceWhiskyHunterResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceWhiskyHunterResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
