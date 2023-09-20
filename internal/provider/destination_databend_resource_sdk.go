// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationDatabendResourceModel) ToCreateSDKType() *shared.DestinationDatabendCreateRequest {
	database := r.Configuration.Database.ValueString()
	destinationType := shared.DestinationDatabendDatabend(r.Configuration.DestinationType.ValueString())
	host := r.Configuration.Host.ValueString()
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	table := new(string)
	if !r.Configuration.Table.IsUnknown() && !r.Configuration.Table.IsNull() {
		*table = r.Configuration.Table.ValueString()
	} else {
		table = nil
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationDatabend{
		Database:        database,
		DestinationType: destinationType,
		Host:            host,
		Password:        password,
		Port:            port,
		Table:           table,
		Username:        username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDatabendCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDatabendResourceModel) ToGetSDKType() *shared.DestinationDatabendCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationDatabendResourceModel) ToUpdateSDKType() *shared.DestinationDatabendPutRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	table := new(string)
	if !r.Configuration.Table.IsUnknown() && !r.Configuration.Table.IsNull() {
		*table = r.Configuration.Table.ValueString()
	} else {
		table = nil
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationDatabendUpdate{
		Database: database,
		Host:     host,
		Password: password,
		Port:     port,
		Table:    table,
		Username: username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDatabendPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDatabendResourceModel) ToDeleteSDKType() *shared.DestinationDatabendCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationDatabendResourceModel) RefreshFromGetResponse(resp *shared.DestinationDatabendGetResponse) {
	r.Configuration.Database = types.StringValue(resp.Configuration.Database)
	r.Configuration.DestinationType = types.StringValue(string(resp.Configuration.DestinationType))
	r.Configuration.Host = types.StringValue(resp.Configuration.Host)
	if resp.Configuration.Password != nil {
		r.Configuration.Password = types.StringValue(*resp.Configuration.Password)
	} else {
		r.Configuration.Password = types.StringNull()
	}
	if resp.Configuration.Port != nil {
		r.Configuration.Port = types.Int64Value(*resp.Configuration.Port)
	} else {
		r.Configuration.Port = types.Int64Null()
	}
	if resp.Configuration.Table != nil {
		r.Configuration.Table = types.StringValue(*resp.Configuration.Table)
	} else {
		r.Configuration.Table = types.StringNull()
	}
	r.Configuration.Username = types.StringValue(resp.Configuration.Username)
	if resp.DestinationID != nil {
		r.DestinationID = types.StringValue(*resp.DestinationID)
	} else {
		r.DestinationID = types.StringNull()
	}
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationDatabendResourceModel) RefreshFromCreateResponse(resp *shared.DestinationDatabendGetResponse) {
	r.RefreshFromGetResponse(resp)
}
