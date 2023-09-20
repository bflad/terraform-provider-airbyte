// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceZuoraResourceModel) ToCreateSDKType() *shared.SourceZuoraCreateRequest {
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	dataQuery := shared.SourceZuoraDataQueryType(r.Configuration.DataQuery.ValueString())
	sourceType := shared.SourceZuoraZuora(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	tenantEndpoint := shared.SourceZuoraTenantEndpointLocation(r.Configuration.TenantEndpoint.ValueString())
	windowInDays := new(string)
	if !r.Configuration.WindowInDays.IsUnknown() && !r.Configuration.WindowInDays.IsNull() {
		*windowInDays = r.Configuration.WindowInDays.ValueString()
	} else {
		windowInDays = nil
	}
	configuration := shared.SourceZuora{
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		DataQuery:      dataQuery,
		SourceType:     sourceType,
		StartDate:      startDate,
		TenantEndpoint: tenantEndpoint,
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
	out := shared.SourceZuoraCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZuoraResourceModel) ToGetSDKType() *shared.SourceZuoraCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZuoraResourceModel) ToUpdateSDKType() *shared.SourceZuoraPutRequest {
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	dataQuery := shared.SourceZuoraUpdateDataQueryType(r.Configuration.DataQuery.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	tenantEndpoint := shared.SourceZuoraUpdateTenantEndpointLocation(r.Configuration.TenantEndpoint.ValueString())
	windowInDays := new(string)
	if !r.Configuration.WindowInDays.IsUnknown() && !r.Configuration.WindowInDays.IsNull() {
		*windowInDays = r.Configuration.WindowInDays.ValueString()
	} else {
		windowInDays = nil
	}
	configuration := shared.SourceZuoraUpdate{
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		DataQuery:      dataQuery,
		StartDate:      startDate,
		TenantEndpoint: tenantEndpoint,
		WindowInDays:   windowInDays,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZuoraPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZuoraResourceModel) ToDeleteSDKType() *shared.SourceZuoraCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZuoraResourceModel) RefreshFromGetResponse(resp *shared.SourceZuoraGetResponse) {
	r.Configuration.ClientID = types.StringValue(resp.Configuration.ClientID)
	r.Configuration.ClientSecret = types.StringValue(resp.Configuration.ClientSecret)
	r.Configuration.DataQuery = types.StringValue(string(resp.Configuration.DataQuery))
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	r.Configuration.StartDate = types.StringValue(resp.Configuration.StartDate)
	r.Configuration.TenantEndpoint = types.StringValue(string(resp.Configuration.TenantEndpoint))
	if resp.Configuration.WindowInDays != nil {
		r.Configuration.WindowInDays = types.StringValue(*resp.Configuration.WindowInDays)
	} else {
		r.Configuration.WindowInDays = types.StringNull()
	}
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

func (r *SourceZuoraResourceModel) RefreshFromCreateResponse(resp *shared.SourceZuoraGetResponse) {
	r.RefreshFromGetResponse(resp)
}
