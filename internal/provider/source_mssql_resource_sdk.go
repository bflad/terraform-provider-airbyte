// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceMssqlResourceModel) ToCreateSDKType() *shared.SourceMssqlCreateRequest {
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
	var replicationMethod *shared.SourceMssqlUpdateMethod
	if r.Configuration.ReplicationMethod != nil {
		var sourceMssqlReadChangesUsingChangeDataCaptureCDC *shared.SourceMssqlReadChangesUsingChangeDataCaptureCDC
		if r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC != nil {
			dataToSync := new(shared.SourceMssqlDataToSync)
			if !r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.DataToSync.IsUnknown() && !r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.DataToSync.IsNull() {
				*dataToSync = shared.SourceMssqlDataToSync(r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.DataToSync.ValueString())
			} else {
				dataToSync = nil
			}
			initialWaitingSeconds := new(int64)
			if !r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.InitialWaitingSeconds.IsUnknown() && !r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.InitialWaitingSeconds.IsNull() {
				*initialWaitingSeconds = r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.InitialWaitingSeconds.ValueInt64()
			} else {
				initialWaitingSeconds = nil
			}
			snapshotIsolation := new(shared.SourceMssqlInitialSnapshotIsolationLevel)
			if !r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.SnapshotIsolation.IsUnknown() && !r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.SnapshotIsolation.IsNull() {
				*snapshotIsolation = shared.SourceMssqlInitialSnapshotIsolationLevel(r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.SnapshotIsolation.ValueString())
			} else {
				snapshotIsolation = nil
			}
			sourceMssqlReadChangesUsingChangeDataCaptureCDC = &shared.SourceMssqlReadChangesUsingChangeDataCaptureCDC{
				DataToSync:            dataToSync,
				InitialWaitingSeconds: initialWaitingSeconds,
				SnapshotIsolation:     snapshotIsolation,
			}
		}
		if sourceMssqlReadChangesUsingChangeDataCaptureCDC != nil {
			replicationMethod = &shared.SourceMssqlUpdateMethod{
				ReadChangesUsingChangeDataCaptureCDC: sourceMssqlReadChangesUsingChangeDataCaptureCDC,
			}
		}
		var sourceMssqlScanChangesWithUserDefinedCursor *shared.SourceMssqlScanChangesWithUserDefinedCursor
		if r.Configuration.ReplicationMethod.ScanChangesWithUserDefinedCursor != nil {
			sourceMssqlScanChangesWithUserDefinedCursor = &shared.SourceMssqlScanChangesWithUserDefinedCursor{}
		}
		if sourceMssqlScanChangesWithUserDefinedCursor != nil {
			replicationMethod = &shared.SourceMssqlUpdateMethod{
				ScanChangesWithUserDefinedCursor: sourceMssqlScanChangesWithUserDefinedCursor,
			}
		}
	}
	var schemas []string = nil
	for _, schemasItem := range r.Configuration.Schemas {
		schemas = append(schemas, schemasItem.ValueString())
	}
	var sslMethod *shared.SourceMssqlSSLMethod
	if r.Configuration.SslMethod != nil {
		var sourceMssqlEncryptedTrustServerCertificate *shared.SourceMssqlEncryptedTrustServerCertificate
		if r.Configuration.SslMethod.EncryptedTrustServerCertificate != nil {
			sourceMssqlEncryptedTrustServerCertificate = &shared.SourceMssqlEncryptedTrustServerCertificate{}
		}
		if sourceMssqlEncryptedTrustServerCertificate != nil {
			sslMethod = &shared.SourceMssqlSSLMethod{
				EncryptedTrustServerCertificate: sourceMssqlEncryptedTrustServerCertificate,
			}
		}
		var sourceMssqlEncryptedVerifyCertificate *shared.SourceMssqlEncryptedVerifyCertificate
		if r.Configuration.SslMethod.EncryptedVerifyCertificate != nil {
			hostNameInCertificate := new(string)
			if !r.Configuration.SslMethod.EncryptedVerifyCertificate.HostNameInCertificate.IsUnknown() && !r.Configuration.SslMethod.EncryptedVerifyCertificate.HostNameInCertificate.IsNull() {
				*hostNameInCertificate = r.Configuration.SslMethod.EncryptedVerifyCertificate.HostNameInCertificate.ValueString()
			} else {
				hostNameInCertificate = nil
			}
			sourceMssqlEncryptedVerifyCertificate = &shared.SourceMssqlEncryptedVerifyCertificate{
				HostNameInCertificate: hostNameInCertificate,
			}
		}
		if sourceMssqlEncryptedVerifyCertificate != nil {
			sslMethod = &shared.SourceMssqlSSLMethod{
				EncryptedVerifyCertificate: sourceMssqlEncryptedVerifyCertificate,
			}
		}
	}
	var tunnelMethod *shared.SourceMssqlSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var sourceMssqlNoTunnel *shared.SourceMssqlNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			sourceMssqlNoTunnel = &shared.SourceMssqlNoTunnel{}
		}
		if sourceMssqlNoTunnel != nil {
			tunnelMethod = &shared.SourceMssqlSSHTunnelMethod{
				NoTunnel: sourceMssqlNoTunnel,
			}
		}
		var sourceMssqlSSHKeyAuthentication *shared.SourceMssqlSSHKeyAuthentication
		if r.Configuration.TunnelMethod.SSHKeyAuthentication != nil {
			sshKey := r.Configuration.TunnelMethod.SSHKeyAuthentication.SSHKey.ValueString()
			tunnelHost := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelHost.ValueString()
			tunnelPort := new(int64)
			if !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsNull() {
				*tunnelPort = r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort = nil
			}
			tunnelUser := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelUser.ValueString()
			sourceMssqlSSHKeyAuthentication = &shared.SourceMssqlSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if sourceMssqlSSHKeyAuthentication != nil {
			tunnelMethod = &shared.SourceMssqlSSHTunnelMethod{
				SSHKeyAuthentication: sourceMssqlSSHKeyAuthentication,
			}
		}
		var sourceMssqlPasswordAuthentication *shared.SourceMssqlPasswordAuthentication
		if r.Configuration.TunnelMethod.PasswordAuthentication != nil {
			tunnelHost1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelHost.ValueString()
			tunnelPort1 := new(int64)
			if !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsNull() {
				*tunnelPort1 = r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort1 = nil
			}
			tunnelUser1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUser.ValueString()
			tunnelUserPassword := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUserPassword.ValueString()
			sourceMssqlPasswordAuthentication = &shared.SourceMssqlPasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if sourceMssqlPasswordAuthentication != nil {
			tunnelMethod = &shared.SourceMssqlSSHTunnelMethod{
				PasswordAuthentication: sourceMssqlPasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceMssql{
		Database:          database,
		Host:              host,
		JdbcURLParams:     jdbcURLParams,
		Password:          password,
		Port:              port,
		ReplicationMethod: replicationMethod,
		Schemas:           schemas,
		SslMethod:         sslMethod,
		TunnelMethod:      tunnelMethod,
		Username:          username,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMssqlCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMssqlResourceModel) ToGetSDKType() *shared.SourceMssqlCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMssqlResourceModel) ToUpdateSDKType() *shared.SourceMssqlPutRequest {
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
	var replicationMethod *shared.UpdateMethod
	if r.Configuration.ReplicationMethod != nil {
		var readChangesUsingChangeDataCaptureCDC *shared.ReadChangesUsingChangeDataCaptureCDC
		if r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC != nil {
			dataToSync := new(shared.DataToSync)
			if !r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.DataToSync.IsUnknown() && !r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.DataToSync.IsNull() {
				*dataToSync = shared.DataToSync(r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.DataToSync.ValueString())
			} else {
				dataToSync = nil
			}
			initialWaitingSeconds := new(int64)
			if !r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.InitialWaitingSeconds.IsUnknown() && !r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.InitialWaitingSeconds.IsNull() {
				*initialWaitingSeconds = r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.InitialWaitingSeconds.ValueInt64()
			} else {
				initialWaitingSeconds = nil
			}
			snapshotIsolation := new(shared.InitialSnapshotIsolationLevel)
			if !r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.SnapshotIsolation.IsUnknown() && !r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.SnapshotIsolation.IsNull() {
				*snapshotIsolation = shared.InitialSnapshotIsolationLevel(r.Configuration.ReplicationMethod.ReadChangesUsingChangeDataCaptureCDC.SnapshotIsolation.ValueString())
			} else {
				snapshotIsolation = nil
			}
			readChangesUsingChangeDataCaptureCDC = &shared.ReadChangesUsingChangeDataCaptureCDC{
				DataToSync:            dataToSync,
				InitialWaitingSeconds: initialWaitingSeconds,
				SnapshotIsolation:     snapshotIsolation,
			}
		}
		if readChangesUsingChangeDataCaptureCDC != nil {
			replicationMethod = &shared.UpdateMethod{
				ReadChangesUsingChangeDataCaptureCDC: readChangesUsingChangeDataCaptureCDC,
			}
		}
		var scanChangesWithUserDefinedCursor *shared.ScanChangesWithUserDefinedCursor
		if r.Configuration.ReplicationMethod.ScanChangesWithUserDefinedCursor != nil {
			scanChangesWithUserDefinedCursor = &shared.ScanChangesWithUserDefinedCursor{}
		}
		if scanChangesWithUserDefinedCursor != nil {
			replicationMethod = &shared.UpdateMethod{
				ScanChangesWithUserDefinedCursor: scanChangesWithUserDefinedCursor,
			}
		}
	}
	var schemas []string = nil
	for _, schemasItem := range r.Configuration.Schemas {
		schemas = append(schemas, schemasItem.ValueString())
	}
	var sslMethod *shared.SourceMssqlUpdateSSLMethod
	if r.Configuration.SslMethod != nil {
		var sourceMssqlUpdateEncryptedTrustServerCertificate *shared.SourceMssqlUpdateEncryptedTrustServerCertificate
		if r.Configuration.SslMethod.EncryptedTrustServerCertificate != nil {
			sourceMssqlUpdateEncryptedTrustServerCertificate = &shared.SourceMssqlUpdateEncryptedTrustServerCertificate{}
		}
		if sourceMssqlUpdateEncryptedTrustServerCertificate != nil {
			sslMethod = &shared.SourceMssqlUpdateSSLMethod{
				EncryptedTrustServerCertificate: sourceMssqlUpdateEncryptedTrustServerCertificate,
			}
		}
		var sourceMssqlUpdateEncryptedVerifyCertificate *shared.SourceMssqlUpdateEncryptedVerifyCertificate
		if r.Configuration.SslMethod.EncryptedVerifyCertificate != nil {
			hostNameInCertificate := new(string)
			if !r.Configuration.SslMethod.EncryptedVerifyCertificate.HostNameInCertificate.IsUnknown() && !r.Configuration.SslMethod.EncryptedVerifyCertificate.HostNameInCertificate.IsNull() {
				*hostNameInCertificate = r.Configuration.SslMethod.EncryptedVerifyCertificate.HostNameInCertificate.ValueString()
			} else {
				hostNameInCertificate = nil
			}
			sourceMssqlUpdateEncryptedVerifyCertificate = &shared.SourceMssqlUpdateEncryptedVerifyCertificate{
				HostNameInCertificate: hostNameInCertificate,
			}
		}
		if sourceMssqlUpdateEncryptedVerifyCertificate != nil {
			sslMethod = &shared.SourceMssqlUpdateSSLMethod{
				EncryptedVerifyCertificate: sourceMssqlUpdateEncryptedVerifyCertificate,
			}
		}
	}
	var tunnelMethod *shared.SourceMssqlUpdateSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var sourceMssqlUpdateNoTunnel *shared.SourceMssqlUpdateNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			sourceMssqlUpdateNoTunnel = &shared.SourceMssqlUpdateNoTunnel{}
		}
		if sourceMssqlUpdateNoTunnel != nil {
			tunnelMethod = &shared.SourceMssqlUpdateSSHTunnelMethod{
				NoTunnel: sourceMssqlUpdateNoTunnel,
			}
		}
		var sourceMssqlUpdateSSHKeyAuthentication *shared.SourceMssqlUpdateSSHKeyAuthentication
		if r.Configuration.TunnelMethod.SSHKeyAuthentication != nil {
			sshKey := r.Configuration.TunnelMethod.SSHKeyAuthentication.SSHKey.ValueString()
			tunnelHost := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelHost.ValueString()
			tunnelPort := new(int64)
			if !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsNull() {
				*tunnelPort = r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort = nil
			}
			tunnelUser := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelUser.ValueString()
			sourceMssqlUpdateSSHKeyAuthentication = &shared.SourceMssqlUpdateSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if sourceMssqlUpdateSSHKeyAuthentication != nil {
			tunnelMethod = &shared.SourceMssqlUpdateSSHTunnelMethod{
				SSHKeyAuthentication: sourceMssqlUpdateSSHKeyAuthentication,
			}
		}
		var sourceMssqlUpdatePasswordAuthentication *shared.SourceMssqlUpdatePasswordAuthentication
		if r.Configuration.TunnelMethod.PasswordAuthentication != nil {
			tunnelHost1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelHost.ValueString()
			tunnelPort1 := new(int64)
			if !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsNull() {
				*tunnelPort1 = r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort1 = nil
			}
			tunnelUser1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUser.ValueString()
			tunnelUserPassword := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUserPassword.ValueString()
			sourceMssqlUpdatePasswordAuthentication = &shared.SourceMssqlUpdatePasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if sourceMssqlUpdatePasswordAuthentication != nil {
			tunnelMethod = &shared.SourceMssqlUpdateSSHTunnelMethod{
				PasswordAuthentication: sourceMssqlUpdatePasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceMssqlUpdate{
		Database:          database,
		Host:              host,
		JdbcURLParams:     jdbcURLParams,
		Password:          password,
		Port:              port,
		ReplicationMethod: replicationMethod,
		Schemas:           schemas,
		SslMethod:         sslMethod,
		TunnelMethod:      tunnelMethod,
		Username:          username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMssqlPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMssqlResourceModel) ToDeleteSDKType() *shared.SourceMssqlCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMssqlResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceMssqlResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
