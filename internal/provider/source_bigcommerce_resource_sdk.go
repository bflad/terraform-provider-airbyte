// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceBigcommerceResourceModel) ToCreateSDKType() *shared.SourceBigcommerceCreateRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	sourceType := shared.SourceBigcommerceBigcommerce(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	storeHash := r.Configuration.StoreHash.ValueString()
	configuration := shared.SourceBigcommerce{
		AccessToken: accessToken,
		SourceType:  sourceType,
		StartDate:   startDate,
		StoreHash:   storeHash,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceBigcommerceCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceBigcommerceResourceModel) ToGetSDKType() *shared.SourceBigcommerceCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceBigcommerceResourceModel) ToUpdateSDKType() *shared.SourceBigcommercePutRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	startDate := r.Configuration.StartDate.ValueString()
	storeHash := r.Configuration.StoreHash.ValueString()
	configuration := shared.SourceBigcommerceUpdate{
		AccessToken: accessToken,
		StartDate:   startDate,
		StoreHash:   storeHash,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceBigcommercePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceBigcommerceResourceModel) ToDeleteSDKType() *shared.SourceBigcommerceCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceBigcommerceResourceModel) RefreshFromGetResponse(resp *shared.SourceBigcommerceGetResponse) {
	r.Configuration.AccessToken = types.StringValue(resp.Configuration.AccessToken)
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	r.Configuration.StartDate = types.StringValue(resp.Configuration.StartDate)
	r.Configuration.StoreHash = types.StringValue(resp.Configuration.StoreHash)
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

func (r *SourceBigcommerceResourceModel) RefreshFromCreateResponse(resp *shared.SourceBigcommerceGetResponse) {
	r.RefreshFromGetResponse(resp)
}
