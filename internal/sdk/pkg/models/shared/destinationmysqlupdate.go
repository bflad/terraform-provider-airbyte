// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

// DestinationMysqlUpdateSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationMysqlUpdateSchemasTunnelMethodTunnelMethod string

const (
	DestinationMysqlUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth DestinationMysqlUpdateSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationMysqlUpdateSchemasTunnelMethodTunnelMethod) ToPointer() *DestinationMysqlUpdateSchemasTunnelMethodTunnelMethod {
	return &e
}

func (e *DestinationMysqlUpdateSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationMysqlUpdateSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMysqlUpdateSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

// DestinationMysqlUpdatePasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMysqlUpdatePasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationMysqlUpdateSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationMysqlUpdatePasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMysqlUpdatePasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMysqlUpdatePasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationMysqlUpdatePasswordAuthentication) GetTunnelMethod() DestinationMysqlUpdateSchemasTunnelMethodTunnelMethod {
	return DestinationMysqlUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *DestinationMysqlUpdatePasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationMysqlUpdatePasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationMysqlUpdatePasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationMysqlUpdateSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationMysqlUpdateSchemasTunnelMethod string

const (
	DestinationMysqlUpdateSchemasTunnelMethodSSHKeyAuth DestinationMysqlUpdateSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationMysqlUpdateSchemasTunnelMethod) ToPointer() *DestinationMysqlUpdateSchemasTunnelMethod {
	return &e
}

func (e *DestinationMysqlUpdateSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationMysqlUpdateSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMysqlUpdateSchemasTunnelMethod: %v", v)
	}
}

// DestinationMysqlUpdateSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMysqlUpdateSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationMysqlUpdateSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationMysqlUpdateSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMysqlUpdateSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMysqlUpdateSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationMysqlUpdateSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationMysqlUpdateSSHKeyAuthentication) GetTunnelMethod() DestinationMysqlUpdateSchemasTunnelMethod {
	return DestinationMysqlUpdateSchemasTunnelMethodSSHKeyAuth
}

func (o *DestinationMysqlUpdateSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationMysqlUpdateSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationMysqlUpdateTunnelMethod - No ssh tunnel needed to connect to database
type DestinationMysqlUpdateTunnelMethod string

const (
	DestinationMysqlUpdateTunnelMethodNoTunnel DestinationMysqlUpdateTunnelMethod = "NO_TUNNEL"
)

func (e DestinationMysqlUpdateTunnelMethod) ToPointer() *DestinationMysqlUpdateTunnelMethod {
	return &e
}

func (e *DestinationMysqlUpdateTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationMysqlUpdateTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMysqlUpdateTunnelMethod: %v", v)
	}
}

// DestinationMysqlUpdateNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMysqlUpdateNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationMysqlUpdateTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationMysqlUpdateNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMysqlUpdateNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMysqlUpdateNoTunnel) GetTunnelMethod() DestinationMysqlUpdateTunnelMethod {
	return DestinationMysqlUpdateTunnelMethodNoTunnel
}

type DestinationMysqlUpdateSSHTunnelMethodType string

const (
	DestinationMysqlUpdateSSHTunnelMethodTypeNoTunnel               DestinationMysqlUpdateSSHTunnelMethodType = "NoTunnel"
	DestinationMysqlUpdateSSHTunnelMethodTypeSSHKeyAuthentication   DestinationMysqlUpdateSSHTunnelMethodType = "SSHKeyAuthentication"
	DestinationMysqlUpdateSSHTunnelMethodTypePasswordAuthentication DestinationMysqlUpdateSSHTunnelMethodType = "PasswordAuthentication"
)

type DestinationMysqlUpdateSSHTunnelMethod struct {
	NoTunnel               *DestinationMysqlUpdateNoTunnel
	SSHKeyAuthentication   *DestinationMysqlUpdateSSHKeyAuthentication
	PasswordAuthentication *DestinationMysqlUpdatePasswordAuthentication

	Type DestinationMysqlUpdateSSHTunnelMethodType
}

func CreateDestinationMysqlUpdateSSHTunnelMethodNoTunnel(noTunnel DestinationMysqlUpdateNoTunnel) DestinationMysqlUpdateSSHTunnelMethod {
	typ := DestinationMysqlUpdateSSHTunnelMethodTypeNoTunnel

	return DestinationMysqlUpdateSSHTunnelMethod{
		NoTunnel: &noTunnel,
		Type:     typ,
	}
}

func CreateDestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication(sshKeyAuthentication DestinationMysqlUpdateSSHKeyAuthentication) DestinationMysqlUpdateSSHTunnelMethod {
	typ := DestinationMysqlUpdateSSHTunnelMethodTypeSSHKeyAuthentication

	return DestinationMysqlUpdateSSHTunnelMethod{
		SSHKeyAuthentication: &sshKeyAuthentication,
		Type:                 typ,
	}
}

func CreateDestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication(passwordAuthentication DestinationMysqlUpdatePasswordAuthentication) DestinationMysqlUpdateSSHTunnelMethod {
	typ := DestinationMysqlUpdateSSHTunnelMethodTypePasswordAuthentication

	return DestinationMysqlUpdateSSHTunnelMethod{
		PasswordAuthentication: &passwordAuthentication,
		Type:                   typ,
	}
}

func (u *DestinationMysqlUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	noTunnel := new(DestinationMysqlUpdateNoTunnel)
	if err := utils.UnmarshalJSON(data, &noTunnel, "", true, true); err == nil {
		u.NoTunnel = noTunnel
		u.Type = DestinationMysqlUpdateSSHTunnelMethodTypeNoTunnel
		return nil
	}

	sshKeyAuthentication := new(DestinationMysqlUpdateSSHKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &sshKeyAuthentication, "", true, true); err == nil {
		u.SSHKeyAuthentication = sshKeyAuthentication
		u.Type = DestinationMysqlUpdateSSHTunnelMethodTypeSSHKeyAuthentication
		return nil
	}

	passwordAuthentication := new(DestinationMysqlUpdatePasswordAuthentication)
	if err := utils.UnmarshalJSON(data, &passwordAuthentication, "", true, true); err == nil {
		u.PasswordAuthentication = passwordAuthentication
		u.Type = DestinationMysqlUpdateSSHTunnelMethodTypePasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMysqlUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
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

type DestinationMysqlUpdate struct {
	// Name of the database.
	Database string `json:"database"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// Port of the database.
	Port *int64 `default:"3306" json:"port"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationMysqlUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}

func (d DestinationMysqlUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMysqlUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMysqlUpdate) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationMysqlUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationMysqlUpdate) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationMysqlUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *DestinationMysqlUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DestinationMysqlUpdate) GetTunnelMethod() *DestinationMysqlUpdateSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *DestinationMysqlUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
