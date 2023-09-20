// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationMysqlResourceModel) ToCreateSDKType() *shared.DestinationMysqlCreateRequest {
	database := r.Configuration.Database.ValueString()
	destinationType := shared.DestinationMysqlMysql(r.Configuration.DestinationType.ValueString())
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := r.Configuration.Port.ValueInt64()
	var tunnelMethod *shared.DestinationMysqlSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var destinationMysqlSSHTunnelMethodNoTunnel *shared.DestinationMysqlSSHTunnelMethodNoTunnel
		if r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodNoTunnel != nil {
			tunnelMethod1 := shared.DestinationMysqlSSHTunnelMethodNoTunnelTunnelMethod(r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodNoTunnel.TunnelMethod.ValueString())
			destinationMysqlSSHTunnelMethodNoTunnel = &shared.DestinationMysqlSSHTunnelMethodNoTunnel{
				TunnelMethod: tunnelMethod1,
			}
		}
		if destinationMysqlSSHTunnelMethodNoTunnel != nil {
			tunnelMethod = &shared.DestinationMysqlSSHTunnelMethod{
				DestinationMysqlSSHTunnelMethodNoTunnel: destinationMysqlSSHTunnelMethodNoTunnel,
			}
		}
		var destinationMysqlSSHTunnelMethodSSHKeyAuthentication *shared.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication
		if r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication != nil {
			sshKey := r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication.SSHKey.ValueString()
			tunnelHost := r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication.TunnelHost.ValueString()
			tunnelMethod2 := shared.DestinationMysqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication.TunnelMethod.ValueString())
			tunnelPort := r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication.TunnelPort.ValueInt64()
			tunnelUser := r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication.TunnelUser.ValueString()
			destinationMysqlSSHTunnelMethodSSHKeyAuthentication = &shared.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication{
				SSHKey:       sshKey,
				TunnelHost:   tunnelHost,
				TunnelMethod: tunnelMethod2,
				TunnelPort:   tunnelPort,
				TunnelUser:   tunnelUser,
			}
		}
		if destinationMysqlSSHTunnelMethodSSHKeyAuthentication != nil {
			tunnelMethod = &shared.DestinationMysqlSSHTunnelMethod{
				DestinationMysqlSSHTunnelMethodSSHKeyAuthentication: destinationMysqlSSHTunnelMethodSSHKeyAuthentication,
			}
		}
		var destinationMysqlSSHTunnelMethodPasswordAuthentication *shared.DestinationMysqlSSHTunnelMethodPasswordAuthentication
		if r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication != nil {
			tunnelHost1 := r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication.TunnelHost.ValueString()
			tunnelMethod3 := shared.DestinationMysqlSSHTunnelMethodPasswordAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication.TunnelMethod.ValueString())
			tunnelPort1 := r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication.TunnelPort.ValueInt64()
			tunnelUser1 := r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication.TunnelUser.ValueString()
			tunnelUserPassword := r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication.TunnelUserPassword.ValueString()
			destinationMysqlSSHTunnelMethodPasswordAuthentication = &shared.DestinationMysqlSSHTunnelMethodPasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelMethod:       tunnelMethod3,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if destinationMysqlSSHTunnelMethodPasswordAuthentication != nil {
			tunnelMethod = &shared.DestinationMysqlSSHTunnelMethod{
				DestinationMysqlSSHTunnelMethodPasswordAuthentication: destinationMysqlSSHTunnelMethodPasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationMysql{
		Database:        database,
		DestinationType: destinationType,
		Host:            host,
		JdbcURLParams:   jdbcURLParams,
		Password:        password,
		Port:            port,
		TunnelMethod:    tunnelMethod,
		Username:        username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationMysqlCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationMysqlResourceModel) ToGetSDKType() *shared.DestinationMysqlCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationMysqlResourceModel) ToUpdateSDKType() *shared.DestinationMysqlPutRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := r.Configuration.Port.ValueInt64()
	var tunnelMethod *shared.DestinationMysqlUpdateSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var destinationMysqlUpdateSSHTunnelMethodNoTunnel *shared.DestinationMysqlUpdateSSHTunnelMethodNoTunnel
		if r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodNoTunnel != nil {
			tunnelMethod1 := shared.DestinationMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod(r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodNoTunnel.TunnelMethod.ValueString())
			destinationMysqlUpdateSSHTunnelMethodNoTunnel = &shared.DestinationMysqlUpdateSSHTunnelMethodNoTunnel{
				TunnelMethod: tunnelMethod1,
			}
		}
		if destinationMysqlUpdateSSHTunnelMethodNoTunnel != nil {
			tunnelMethod = &shared.DestinationMysqlUpdateSSHTunnelMethod{
				DestinationMysqlUpdateSSHTunnelMethodNoTunnel: destinationMysqlUpdateSSHTunnelMethodNoTunnel,
			}
		}
		var destinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication *shared.DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication
		if r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
			sshKey := r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication.SSHKey.ValueString()
			tunnelHost := r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication.TunnelHost.ValueString()
			tunnelMethod2 := shared.DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication.TunnelMethod.ValueString())
			tunnelPort := r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication.TunnelPort.ValueInt64()
			tunnelUser := r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication.TunnelUser.ValueString()
			destinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication = &shared.DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication{
				SSHKey:       sshKey,
				TunnelHost:   tunnelHost,
				TunnelMethod: tunnelMethod2,
				TunnelPort:   tunnelPort,
				TunnelUser:   tunnelUser,
			}
		}
		if destinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
			tunnelMethod = &shared.DestinationMysqlUpdateSSHTunnelMethod{
				DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication: destinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication,
			}
		}
		var destinationMysqlUpdateSSHTunnelMethodPasswordAuthentication *shared.DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication
		if r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication != nil {
			tunnelHost1 := r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication.TunnelHost.ValueString()
			tunnelMethod3 := shared.DestinationMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication.TunnelMethod.ValueString())
			tunnelPort1 := r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication.TunnelPort.ValueInt64()
			tunnelUser1 := r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication.TunnelUser.ValueString()
			tunnelUserPassword := r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication.TunnelUserPassword.ValueString()
			destinationMysqlUpdateSSHTunnelMethodPasswordAuthentication = &shared.DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelMethod:       tunnelMethod3,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if destinationMysqlUpdateSSHTunnelMethodPasswordAuthentication != nil {
			tunnelMethod = &shared.DestinationMysqlUpdateSSHTunnelMethod{
				DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication: destinationMysqlUpdateSSHTunnelMethodPasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationMysqlUpdate{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		TunnelMethod:  tunnelMethod,
		Username:      username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationMysqlPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationMysqlResourceModel) ToDeleteSDKType() *shared.DestinationMysqlCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationMysqlResourceModel) RefreshFromGetResponse(resp *shared.DestinationMysqlGetResponse) {
	r.Configuration.Database = types.StringValue(resp.Configuration.Database)
	r.Configuration.DestinationType = types.StringValue(string(resp.Configuration.DestinationType))
	r.Configuration.Host = types.StringValue(resp.Configuration.Host)
	if resp.Configuration.JdbcURLParams != nil {
		r.Configuration.JdbcURLParams = types.StringValue(*resp.Configuration.JdbcURLParams)
	} else {
		r.Configuration.JdbcURLParams = types.StringNull()
	}
	if resp.Configuration.Password != nil {
		r.Configuration.Password = types.StringValue(*resp.Configuration.Password)
	} else {
		r.Configuration.Password = types.StringNull()
	}
	r.Configuration.Port = types.Int64Value(resp.Configuration.Port)
	if resp.Configuration.TunnelMethod == nil {
		r.Configuration.TunnelMethod = nil
	} else {
		r.Configuration.TunnelMethod = &DestinationMysqlSSHTunnelMethod{}
		if resp.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodNoTunnel != nil {
			r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodNoTunnel = &DestinationClickhouseSSHTunnelMethodNoTunnel{}
			r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodNoTunnel.TunnelMethod = types.StringValue(string(resp.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodNoTunnel.TunnelMethod))
		}
		if resp.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication != nil {
			r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication = &DestinationClickhouseSSHTunnelMethodPasswordAuthentication{}
			r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication.TunnelHost = types.StringValue(resp.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication.TunnelHost)
			r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication.TunnelMethod = types.StringValue(string(resp.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication.TunnelMethod))
			r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication.TunnelPort = types.Int64Value(resp.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication.TunnelPort)
			r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication.TunnelUser = types.StringValue(resp.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication.TunnelUser)
			r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication.TunnelUserPassword = types.StringValue(resp.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodPasswordAuthentication.TunnelUserPassword)
		}
		if resp.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication != nil {
			r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication = &DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication{}
			r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication.SSHKey = types.StringValue(resp.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication.SSHKey)
			r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication.TunnelHost = types.StringValue(resp.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication.TunnelHost)
			r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication.TunnelMethod = types.StringValue(string(resp.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication.TunnelMethod))
			r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication.TunnelPort = types.Int64Value(resp.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication.TunnelPort)
			r.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication.TunnelUser = types.StringValue(resp.Configuration.TunnelMethod.DestinationMysqlSSHTunnelMethodSSHKeyAuthentication.TunnelUser)
		}
		if resp.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodNoTunnel != nil {
			r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodNoTunnel = &DestinationClickhouseSSHTunnelMethodNoTunnel{}
		}
		if resp.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication != nil {
			r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication = &DestinationClickhouseSSHTunnelMethodPasswordAuthentication{}
		}
		if resp.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
			r.Configuration.TunnelMethod.DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication = &DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication{}
		}
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

func (r *DestinationMysqlResourceModel) RefreshFromCreateResponse(resp *shared.DestinationMysqlGetResponse) {
	r.RefreshFromGetResponse(resp)
}
