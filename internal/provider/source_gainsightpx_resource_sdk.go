// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceGainsightPxResourceModel) ToCreateSDKType() *shared.SourceGainsightPxCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	sourceType := shared.SourceGainsightPxGainsightPx(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceGainsightPx{
		APIKey:     apiKey,
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
	out := shared.SourceGainsightPxCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGainsightPxResourceModel) ToGetSDKType() *shared.SourceGainsightPxCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGainsightPxResourceModel) ToUpdateSDKType() *shared.SourceGainsightPxPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	configuration := shared.SourceGainsightPxUpdate{
		APIKey: apiKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGainsightPxPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGainsightPxResourceModel) ToDeleteSDKType() *shared.SourceGainsightPxCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGainsightPxResourceModel) RefreshFromGetResponse(resp *shared.SourceGainsightPxGetResponse) {
	r.Configuration.APIKey = types.StringValue(resp.Configuration.APIKey)
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

func (r *SourceGainsightPxResourceModel) RefreshFromCreateResponse(resp *shared.SourceGainsightPxGetResponse) {
	r.RefreshFromGetResponse(resp)
}
