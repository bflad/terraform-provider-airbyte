// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationOracleDataSourceModel) RefreshFromGetResponse(resp *shared.DestinationOracleGetResponse) {
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
	if resp.Configuration.Schema != nil {
		r.Configuration.Schema = types.StringValue(*resp.Configuration.Schema)
	} else {
		r.Configuration.Schema = types.StringNull()
	}
	r.Configuration.Sid = types.StringValue(resp.Configuration.Sid)
	if resp.Configuration.TunnelMethod == nil {
		r.Configuration.TunnelMethod = nil
	} else {
		r.Configuration.TunnelMethod = &DestinationOracleSSHTunnelMethod{}
		if resp.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodNoTunnel != nil {
			r.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodNoTunnel = &DestinationClickhouseSSHTunnelMethodNoTunnel{}
			r.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodNoTunnel.TunnelMethod = types.StringValue(string(resp.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodNoTunnel.TunnelMethod))
		}
		if resp.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodPasswordAuthentication != nil {
			r.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodPasswordAuthentication = &DestinationClickhouseSSHTunnelMethodPasswordAuthentication{}
			r.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodPasswordAuthentication.TunnelHost = types.StringValue(resp.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodPasswordAuthentication.TunnelHost)
			r.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodPasswordAuthentication.TunnelMethod = types.StringValue(string(resp.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodPasswordAuthentication.TunnelMethod))
			r.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodPasswordAuthentication.TunnelPort = types.Int64Value(resp.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodPasswordAuthentication.TunnelPort)
			r.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodPasswordAuthentication.TunnelUser = types.StringValue(resp.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodPasswordAuthentication.TunnelUser)
			r.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodPasswordAuthentication.TunnelUserPassword = types.StringValue(resp.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodPasswordAuthentication.TunnelUserPassword)
		}
		if resp.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodSSHKeyAuthentication != nil {
			r.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodSSHKeyAuthentication = &DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication{}
			r.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodSSHKeyAuthentication.SSHKey = types.StringValue(resp.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodSSHKeyAuthentication.SSHKey)
			r.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodSSHKeyAuthentication.TunnelHost = types.StringValue(resp.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodSSHKeyAuthentication.TunnelHost)
			r.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodSSHKeyAuthentication.TunnelMethod = types.StringValue(string(resp.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodSSHKeyAuthentication.TunnelMethod))
			r.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodSSHKeyAuthentication.TunnelPort = types.Int64Value(resp.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodSSHKeyAuthentication.TunnelPort)
			r.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodSSHKeyAuthentication.TunnelUser = types.StringValue(resp.Configuration.TunnelMethod.DestinationOracleSSHTunnelMethodSSHKeyAuthentication.TunnelUser)
		}
		if resp.Configuration.TunnelMethod.DestinationOracleUpdateSSHTunnelMethodNoTunnel != nil {
			r.Configuration.TunnelMethod.DestinationOracleUpdateSSHTunnelMethodNoTunnel = &DestinationClickhouseSSHTunnelMethodNoTunnel{}
		}
		if resp.Configuration.TunnelMethod.DestinationOracleUpdateSSHTunnelMethodPasswordAuthentication != nil {
			r.Configuration.TunnelMethod.DestinationOracleUpdateSSHTunnelMethodPasswordAuthentication = &DestinationClickhouseSSHTunnelMethodPasswordAuthentication{}
		}
		if resp.Configuration.TunnelMethod.DestinationOracleUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
			r.Configuration.TunnelMethod.DestinationOracleUpdateSSHTunnelMethodSSHKeyAuthentication = &DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication{}
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
