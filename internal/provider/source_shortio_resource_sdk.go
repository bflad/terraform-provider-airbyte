// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceShortioResourceModel) ToCreateSDKType() *shared.SourceShortioCreateRequest {
	domainID := r.Configuration.DomainID.ValueString()
	secretKey := r.Configuration.SecretKey.ValueString()
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceShortio{
		DomainID:  domainID,
		SecretKey: secretKey,
		StartDate: startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceShortioCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceShortioResourceModel) ToGetSDKType() *shared.SourceShortioCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceShortioResourceModel) ToUpdateSDKType() *shared.SourceShortioPutRequest {
	domainID := r.Configuration.DomainID.ValueString()
	secretKey := r.Configuration.SecretKey.ValueString()
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceShortioUpdate{
		DomainID:  domainID,
		SecretKey: secretKey,
		StartDate: startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceShortioPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceShortioResourceModel) ToDeleteSDKType() *shared.SourceShortioCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceShortioResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceShortioResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
