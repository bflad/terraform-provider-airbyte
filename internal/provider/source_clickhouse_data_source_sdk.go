// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceClickhouseDataSourceModel) RefreshFromGetResponse(resp *shared.SourceClickhouseGetResponse) {
	r.Configuration.Database = types.StringValue(resp.Configuration.Database)
	r.Configuration.Host = types.StringValue(resp.Configuration.Host)
	if resp.Configuration.Password != nil {
		r.Configuration.Password = types.StringValue(*resp.Configuration.Password)
	} else {
		r.Configuration.Password = types.StringNull()
	}
	r.Configuration.Port = types.Int64Value(resp.Configuration.Port)
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	if resp.Configuration.TunnelMethod == nil {
		r.Configuration.TunnelMethod = nil
	} else {
		r.Configuration.TunnelMethod = &SourceClickhouseSSHTunnelMethod{}
		if resp.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodNoTunnel != nil {
			r.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodNoTunnel = &DestinationClickhouseSSHTunnelMethodNoTunnel{}
			r.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodNoTunnel.TunnelMethod = types.StringValue(string(resp.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodNoTunnel.TunnelMethod))
		}
		if resp.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodPasswordAuthentication != nil {
			r.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodPasswordAuthentication = &DestinationClickhouseSSHTunnelMethodPasswordAuthentication{}
			r.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodPasswordAuthentication.TunnelHost = types.StringValue(resp.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodPasswordAuthentication.TunnelHost)
			r.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodPasswordAuthentication.TunnelMethod = types.StringValue(string(resp.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodPasswordAuthentication.TunnelMethod))
			r.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodPasswordAuthentication.TunnelPort = types.Int64Value(resp.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodPasswordAuthentication.TunnelPort)
			r.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodPasswordAuthentication.TunnelUser = types.StringValue(resp.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodPasswordAuthentication.TunnelUser)
			r.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodPasswordAuthentication.TunnelUserPassword = types.StringValue(resp.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodPasswordAuthentication.TunnelUserPassword)
		}
		if resp.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodSSHKeyAuthentication != nil {
			r.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodSSHKeyAuthentication = &DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication{}
			r.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodSSHKeyAuthentication.SSHKey = types.StringValue(resp.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodSSHKeyAuthentication.SSHKey)
			r.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodSSHKeyAuthentication.TunnelHost = types.StringValue(resp.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodSSHKeyAuthentication.TunnelHost)
			r.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodSSHKeyAuthentication.TunnelMethod = types.StringValue(string(resp.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodSSHKeyAuthentication.TunnelMethod))
			r.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodSSHKeyAuthentication.TunnelPort = types.Int64Value(resp.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodSSHKeyAuthentication.TunnelPort)
			r.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodSSHKeyAuthentication.TunnelUser = types.StringValue(resp.Configuration.TunnelMethod.SourceClickhouseSSHTunnelMethodSSHKeyAuthentication.TunnelUser)
		}
		if resp.Configuration.TunnelMethod.SourceClickhouseUpdateSSHTunnelMethodNoTunnel != nil {
			r.Configuration.TunnelMethod.SourceClickhouseUpdateSSHTunnelMethodNoTunnel = &DestinationClickhouseSSHTunnelMethodNoTunnel{}
		}
		if resp.Configuration.TunnelMethod.SourceClickhouseUpdateSSHTunnelMethodPasswordAuthentication != nil {
			r.Configuration.TunnelMethod.SourceClickhouseUpdateSSHTunnelMethodPasswordAuthentication = &DestinationClickhouseSSHTunnelMethodPasswordAuthentication{}
		}
		if resp.Configuration.TunnelMethod.SourceClickhouseUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
			r.Configuration.TunnelMethod.SourceClickhouseUpdateSSHTunnelMethodSSHKeyAuthentication = &DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication{}
		}
	}
	r.Configuration.Username = types.StringValue(resp.Configuration.Username)
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
