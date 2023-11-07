// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

// DestinationVerticaUpdateSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationVerticaUpdateSchemasTunnelMethodTunnelMethod string

const (
	DestinationVerticaUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth DestinationVerticaUpdateSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationVerticaUpdateSchemasTunnelMethodTunnelMethod) ToPointer() *DestinationVerticaUpdateSchemasTunnelMethodTunnelMethod {
	return &e
}

func (e *DestinationVerticaUpdateSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationVerticaUpdateSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationVerticaUpdateSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

// DestinationVerticaUpdatePasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationVerticaUpdatePasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationVerticaUpdateSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationVerticaUpdatePasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationVerticaUpdatePasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationVerticaUpdatePasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationVerticaUpdatePasswordAuthentication) GetTunnelMethod() DestinationVerticaUpdateSchemasTunnelMethodTunnelMethod {
	return DestinationVerticaUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *DestinationVerticaUpdatePasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationVerticaUpdatePasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationVerticaUpdatePasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationVerticaUpdateSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationVerticaUpdateSchemasTunnelMethod string

const (
	DestinationVerticaUpdateSchemasTunnelMethodSSHKeyAuth DestinationVerticaUpdateSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationVerticaUpdateSchemasTunnelMethod) ToPointer() *DestinationVerticaUpdateSchemasTunnelMethod {
	return &e
}

func (e *DestinationVerticaUpdateSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationVerticaUpdateSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationVerticaUpdateSchemasTunnelMethod: %v", v)
	}
}

// DestinationVerticaUpdateSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationVerticaUpdateSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationVerticaUpdateSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationVerticaUpdateSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationVerticaUpdateSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationVerticaUpdateSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationVerticaUpdateSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationVerticaUpdateSSHKeyAuthentication) GetTunnelMethod() DestinationVerticaUpdateSchemasTunnelMethod {
	return DestinationVerticaUpdateSchemasTunnelMethodSSHKeyAuth
}

func (o *DestinationVerticaUpdateSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationVerticaUpdateSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationVerticaUpdateTunnelMethod - No ssh tunnel needed to connect to database
type DestinationVerticaUpdateTunnelMethod string

const (
	DestinationVerticaUpdateTunnelMethodNoTunnel DestinationVerticaUpdateTunnelMethod = "NO_TUNNEL"
)

func (e DestinationVerticaUpdateTunnelMethod) ToPointer() *DestinationVerticaUpdateTunnelMethod {
	return &e
}

func (e *DestinationVerticaUpdateTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationVerticaUpdateTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationVerticaUpdateTunnelMethod: %v", v)
	}
}

// DestinationVerticaUpdateNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationVerticaUpdateNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationVerticaUpdateTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationVerticaUpdateNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationVerticaUpdateNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationVerticaUpdateNoTunnel) GetTunnelMethod() DestinationVerticaUpdateTunnelMethod {
	return DestinationVerticaUpdateTunnelMethodNoTunnel
}

type DestinationVerticaUpdateSSHTunnelMethodType string

const (
	DestinationVerticaUpdateSSHTunnelMethodTypeNoTunnel               DestinationVerticaUpdateSSHTunnelMethodType = "NoTunnel"
	DestinationVerticaUpdateSSHTunnelMethodTypeSSHKeyAuthentication   DestinationVerticaUpdateSSHTunnelMethodType = "SSHKeyAuthentication"
	DestinationVerticaUpdateSSHTunnelMethodTypePasswordAuthentication DestinationVerticaUpdateSSHTunnelMethodType = "PasswordAuthentication"
)

type DestinationVerticaUpdateSSHTunnelMethod struct {
	NoTunnel               *DestinationVerticaUpdateNoTunnel
	SSHKeyAuthentication   *DestinationVerticaUpdateSSHKeyAuthentication
	PasswordAuthentication *DestinationVerticaUpdatePasswordAuthentication

	Type DestinationVerticaUpdateSSHTunnelMethodType
}

func CreateDestinationVerticaUpdateSSHTunnelMethodNoTunnel(noTunnel DestinationVerticaUpdateNoTunnel) DestinationVerticaUpdateSSHTunnelMethod {
	typ := DestinationVerticaUpdateSSHTunnelMethodTypeNoTunnel

	return DestinationVerticaUpdateSSHTunnelMethod{
		NoTunnel: &noTunnel,
		Type:     typ,
	}
}

func CreateDestinationVerticaUpdateSSHTunnelMethodSSHKeyAuthentication(sshKeyAuthentication DestinationVerticaUpdateSSHKeyAuthentication) DestinationVerticaUpdateSSHTunnelMethod {
	typ := DestinationVerticaUpdateSSHTunnelMethodTypeSSHKeyAuthentication

	return DestinationVerticaUpdateSSHTunnelMethod{
		SSHKeyAuthentication: &sshKeyAuthentication,
		Type:                 typ,
	}
}

func CreateDestinationVerticaUpdateSSHTunnelMethodPasswordAuthentication(passwordAuthentication DestinationVerticaUpdatePasswordAuthentication) DestinationVerticaUpdateSSHTunnelMethod {
	typ := DestinationVerticaUpdateSSHTunnelMethodTypePasswordAuthentication

	return DestinationVerticaUpdateSSHTunnelMethod{
		PasswordAuthentication: &passwordAuthentication,
		Type:                   typ,
	}
}

func (u *DestinationVerticaUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	noTunnel := new(DestinationVerticaUpdateNoTunnel)
	if err := utils.UnmarshalJSON(data, &noTunnel, "", true, true); err == nil {
		u.NoTunnel = noTunnel
		u.Type = DestinationVerticaUpdateSSHTunnelMethodTypeNoTunnel
		return nil
	}

	sshKeyAuthentication := new(DestinationVerticaUpdateSSHKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &sshKeyAuthentication, "", true, true); err == nil {
		u.SSHKeyAuthentication = sshKeyAuthentication
		u.Type = DestinationVerticaUpdateSSHTunnelMethodTypeSSHKeyAuthentication
		return nil
	}

	passwordAuthentication := new(DestinationVerticaUpdatePasswordAuthentication)
	if err := utils.UnmarshalJSON(data, &passwordAuthentication, "", true, true); err == nil {
		u.PasswordAuthentication = passwordAuthentication
		u.Type = DestinationVerticaUpdateSSHTunnelMethodTypePasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationVerticaUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.NoTunnel != nil {
		return utils.MarshalJSON(u.NoTunnel, "", true)
	}

	if u.SSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.SSHKeyAuthentication, "", true)
	}

	if u.PasswordAuthentication != nil {
		return utils.MarshalJSON(u.PasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationVerticaUpdate struct {
	// Name of the database.
	Database string `json:"database"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// Port of the database.
	Port *int64 `default:"5433" json:"port"`
	// Schema for vertica destination
	Schema string `json:"schema"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationVerticaUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}

func (d DestinationVerticaUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationVerticaUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationVerticaUpdate) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationVerticaUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationVerticaUpdate) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationVerticaUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *DestinationVerticaUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DestinationVerticaUpdate) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *DestinationVerticaUpdate) GetTunnelMethod() *DestinationVerticaUpdateSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *DestinationVerticaUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
