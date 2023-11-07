// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceNetsuiteResourceModel) ToCreateSDKType() *shared.SourceNetsuiteCreateRequest {
	consumerKey := r.Configuration.ConsumerKey.ValueString()
	consumerSecret := r.Configuration.ConsumerSecret.ValueString()
	var objectTypes []string = nil
	for _, objectTypesItem := range r.Configuration.ObjectTypes {
		objectTypes = append(objectTypes, objectTypesItem.ValueString())
	}
	realm := r.Configuration.Realm.ValueString()
	startDatetime := r.Configuration.StartDatetime.ValueString()
	tokenKey := r.Configuration.TokenKey.ValueString()
	tokenSecret := r.Configuration.TokenSecret.ValueString()
	windowInDays := new(int64)
	if !r.Configuration.WindowInDays.IsUnknown() && !r.Configuration.WindowInDays.IsNull() {
		*windowInDays = r.Configuration.WindowInDays.ValueInt64()
	} else {
		windowInDays = nil
	}
	configuration := shared.SourceNetsuite{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		ObjectTypes:    objectTypes,
		Realm:          realm,
		StartDatetime:  startDatetime,
		TokenKey:       tokenKey,
		TokenSecret:    tokenSecret,
		WindowInDays:   windowInDays,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceNetsuiteCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceNetsuiteResourceModel) ToGetSDKType() *shared.SourceNetsuiteCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceNetsuiteResourceModel) ToUpdateSDKType() *shared.SourceNetsuitePutRequest {
	consumerKey := r.Configuration.ConsumerKey.ValueString()
	consumerSecret := r.Configuration.ConsumerSecret.ValueString()
	var objectTypes []string = nil
	for _, objectTypesItem := range r.Configuration.ObjectTypes {
		objectTypes = append(objectTypes, objectTypesItem.ValueString())
	}
	realm := r.Configuration.Realm.ValueString()
	startDatetime := r.Configuration.StartDatetime.ValueString()
	tokenKey := r.Configuration.TokenKey.ValueString()
	tokenSecret := r.Configuration.TokenSecret.ValueString()
	windowInDays := new(int64)
	if !r.Configuration.WindowInDays.IsUnknown() && !r.Configuration.WindowInDays.IsNull() {
		*windowInDays = r.Configuration.WindowInDays.ValueInt64()
	} else {
		windowInDays = nil
	}
	configuration := shared.SourceNetsuiteUpdate{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		ObjectTypes:    objectTypes,
		Realm:          realm,
		StartDatetime:  startDatetime,
		TokenKey:       tokenKey,
		TokenSecret:    tokenSecret,
		WindowInDays:   windowInDays,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceNetsuitePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceNetsuiteResourceModel) ToDeleteSDKType() *shared.SourceNetsuiteCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceNetsuiteResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceNetsuiteResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
