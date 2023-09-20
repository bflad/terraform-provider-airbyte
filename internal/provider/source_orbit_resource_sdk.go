// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceOrbitResourceModel) ToCreateSDKType() *shared.SourceOrbitCreateRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	sourceType := shared.SourceOrbitOrbit(r.Configuration.SourceType.ValueString())
	startDate := new(string)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate = r.Configuration.StartDate.ValueString()
	} else {
		startDate = nil
	}
	workspace := r.Configuration.Workspace.ValueString()
	configuration := shared.SourceOrbit{
		APIToken:   apiToken,
		SourceType: sourceType,
		StartDate:  startDate,
		Workspace:  workspace,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceOrbitCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOrbitResourceModel) ToGetSDKType() *shared.SourceOrbitCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceOrbitResourceModel) ToUpdateSDKType() *shared.SourceOrbitPutRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	startDate := new(string)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate = r.Configuration.StartDate.ValueString()
	} else {
		startDate = nil
	}
	workspace := r.Configuration.Workspace.ValueString()
	configuration := shared.SourceOrbitUpdate{
		APIToken:  apiToken,
		StartDate: startDate,
		Workspace: workspace,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceOrbitPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOrbitResourceModel) ToDeleteSDKType() *shared.SourceOrbitCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceOrbitResourceModel) RefreshFromGetResponse(resp *shared.SourceOrbitGetResponse) {
	r.Configuration.APIToken = types.StringValue(resp.Configuration.APIToken)
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	if resp.Configuration.StartDate != nil {
		r.Configuration.StartDate = types.StringValue(*resp.Configuration.StartDate)
	} else {
		r.Configuration.StartDate = types.StringNull()
	}
	r.Configuration.Workspace = types.StringValue(resp.Configuration.Workspace)
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

func (r *SourceOrbitResourceModel) RefreshFromCreateResponse(resp *shared.SourceOrbitGetResponse) {
	r.RefreshFromGetResponse(resp)
}
