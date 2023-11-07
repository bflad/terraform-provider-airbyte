// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourceClickhouseClickhouse string

const (
	SourceClickhouseClickhouseClickhouse SourceClickhouseClickhouse = "clickhouse"
)

func (e SourceClickhouseClickhouse) ToPointer() *SourceClickhouseClickhouse {
	return &e
}

func (e *SourceClickhouseClickhouse) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "clickhouse":
		*e = SourceClickhouseClickhouse(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceClickhouseClickhouse: %v", v)
	}
}

// SourceClickhouseSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type SourceClickhouseSchemasTunnelMethodTunnelMethod string

const (
	SourceClickhouseSchemasTunnelMethodTunnelMethodSSHPasswordAuth SourceClickhouseSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e SourceClickhouseSchemasTunnelMethodTunnelMethod) ToPointer() *SourceClickhouseSchemasTunnelMethodTunnelMethod {
	return &e
}

func (e *SourceClickhouseSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = SourceClickhouseSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceClickhouseSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

// SourceClickhousePasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceClickhousePasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod SourceClickhouseSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (s SourceClickhousePasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceClickhousePasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceClickhousePasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *SourceClickhousePasswordAuthentication) GetTunnelMethod() SourceClickhouseSchemasTunnelMethodTunnelMethod {
	return SourceClickhouseSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *SourceClickhousePasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *SourceClickhousePasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *SourceClickhousePasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// SourceClickhouseSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type SourceClickhouseSchemasTunnelMethod string

const (
	SourceClickhouseSchemasTunnelMethodSSHKeyAuth SourceClickhouseSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e SourceClickhouseSchemasTunnelMethod) ToPointer() *SourceClickhouseSchemasTunnelMethod {
	return &e
}

func (e *SourceClickhouseSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = SourceClickhouseSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceClickhouseSchemasTunnelMethod: %v", v)
	}
}

// SourceClickhouseSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceClickhouseSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod SourceClickhouseSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (s SourceClickhouseSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceClickhouseSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceClickhouseSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *SourceClickhouseSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *SourceClickhouseSSHKeyAuthentication) GetTunnelMethod() SourceClickhouseSchemasTunnelMethod {
	return SourceClickhouseSchemasTunnelMethodSSHKeyAuth
}

func (o *SourceClickhouseSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *SourceClickhouseSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// SourceClickhouseTunnelMethod - No ssh tunnel needed to connect to database
type SourceClickhouseTunnelMethod string

const (
	SourceClickhouseTunnelMethodNoTunnel SourceClickhouseTunnelMethod = "NO_TUNNEL"
)

func (e SourceClickhouseTunnelMethod) ToPointer() *SourceClickhouseTunnelMethod {
	return &e
}

func (e *SourceClickhouseTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = SourceClickhouseTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceClickhouseTunnelMethod: %v", v)
	}
}

// SourceClickhouseNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceClickhouseNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod SourceClickhouseTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (s SourceClickhouseNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceClickhouseNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceClickhouseNoTunnel) GetTunnelMethod() SourceClickhouseTunnelMethod {
	return SourceClickhouseTunnelMethodNoTunnel
}

type SourceClickhouseSSHTunnelMethodType string

const (
	SourceClickhouseSSHTunnelMethodTypeNoTunnel               SourceClickhouseSSHTunnelMethodType = "NoTunnel"
	SourceClickhouseSSHTunnelMethodTypeSSHKeyAuthentication   SourceClickhouseSSHTunnelMethodType = "SSHKeyAuthentication"
	SourceClickhouseSSHTunnelMethodTypePasswordAuthentication SourceClickhouseSSHTunnelMethodType = "PasswordAuthentication"
)

type SourceClickhouseSSHTunnelMethod struct {
	NoTunnel               *SourceClickhouseNoTunnel
	SSHKeyAuthentication   *SourceClickhouseSSHKeyAuthentication
	PasswordAuthentication *SourceClickhousePasswordAuthentication

	Type SourceClickhouseSSHTunnelMethodType
}

func CreateSourceClickhouseSSHTunnelMethodNoTunnel(noTunnel SourceClickhouseNoTunnel) SourceClickhouseSSHTunnelMethod {
	typ := SourceClickhouseSSHTunnelMethodTypeNoTunnel

	return SourceClickhouseSSHTunnelMethod{
		NoTunnel: &noTunnel,
		Type:     typ,
	}
}

func CreateSourceClickhouseSSHTunnelMethodSSHKeyAuthentication(sshKeyAuthentication SourceClickhouseSSHKeyAuthentication) SourceClickhouseSSHTunnelMethod {
	typ := SourceClickhouseSSHTunnelMethodTypeSSHKeyAuthentication

	return SourceClickhouseSSHTunnelMethod{
		SSHKeyAuthentication: &sshKeyAuthentication,
		Type:                 typ,
	}
}

func CreateSourceClickhouseSSHTunnelMethodPasswordAuthentication(passwordAuthentication SourceClickhousePasswordAuthentication) SourceClickhouseSSHTunnelMethod {
	typ := SourceClickhouseSSHTunnelMethodTypePasswordAuthentication

	return SourceClickhouseSSHTunnelMethod{
		PasswordAuthentication: &passwordAuthentication,
		Type:                   typ,
	}
}

func (u *SourceClickhouseSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	noTunnel := new(SourceClickhouseNoTunnel)
	if err := utils.UnmarshalJSON(data, &noTunnel, "", true, true); err == nil {
		u.NoTunnel = noTunnel
		u.Type = SourceClickhouseSSHTunnelMethodTypeNoTunnel
		return nil
	}

	sshKeyAuthentication := new(SourceClickhouseSSHKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &sshKeyAuthentication, "", true, true); err == nil {
		u.SSHKeyAuthentication = sshKeyAuthentication
		u.Type = SourceClickhouseSSHTunnelMethodTypeSSHKeyAuthentication
		return nil
	}

	passwordAuthentication := new(SourceClickhousePasswordAuthentication)
	if err := utils.UnmarshalJSON(data, &passwordAuthentication, "", true, true); err == nil {
		u.PasswordAuthentication = passwordAuthentication
		u.Type = SourceClickhouseSSHTunnelMethodTypePasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceClickhouseSSHTunnelMethod) MarshalJSON() ([]byte, error) {
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

type SourceClickhouse struct {
	// The name of the database.
	Database string `json:"database"`
	// The host endpoint of the Clickhouse cluster.
	Host string `json:"host"`
	// The password associated with this username.
	Password *string `json:"password,omitempty"`
	// The port of the database.
	Port       *int64                     `default:"8123" json:"port"`
	sourceType SourceClickhouseClickhouse `const:"clickhouse" json:"sourceType"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *SourceClickhouseSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// The username which is used to access the database.
	Username string `json:"username"`
}

func (s SourceClickhouse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceClickhouse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceClickhouse) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *SourceClickhouse) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SourceClickhouse) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourceClickhouse) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SourceClickhouse) GetSourceType() SourceClickhouseClickhouse {
	return SourceClickhouseClickhouseClickhouse
}

func (o *SourceClickhouse) GetTunnelMethod() *SourceClickhouseSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *SourceClickhouse) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
