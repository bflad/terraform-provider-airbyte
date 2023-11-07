// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type DestinationSnowflakeUpdateSchemasAuthType string

const (
	DestinationSnowflakeUpdateSchemasAuthTypeUsernameAndPassword DestinationSnowflakeUpdateSchemasAuthType = "Username and Password"
)

func (e DestinationSnowflakeUpdateSchemasAuthType) ToPointer() *DestinationSnowflakeUpdateSchemasAuthType {
	return &e
}

func (e *DestinationSnowflakeUpdateSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Username and Password":
		*e = DestinationSnowflakeUpdateSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationSnowflakeUpdateSchemasAuthType: %v", v)
	}
}

type UsernameAndPassword struct {
	authType *DestinationSnowflakeUpdateSchemasAuthType `const:"Username and Password" json:"auth_type"`
	// Enter the password associated with the username.
	Password string `json:"password"`
}

func (u UsernameAndPassword) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UsernameAndPassword) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *UsernameAndPassword) GetAuthType() *DestinationSnowflakeUpdateSchemasAuthType {
	return DestinationSnowflakeUpdateSchemasAuthTypeUsernameAndPassword.ToPointer()
}

func (o *UsernameAndPassword) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

type DestinationSnowflakeUpdateAuthType string

const (
	DestinationSnowflakeUpdateAuthTypeKeyPairAuthentication DestinationSnowflakeUpdateAuthType = "Key Pair Authentication"
)

func (e DestinationSnowflakeUpdateAuthType) ToPointer() *DestinationSnowflakeUpdateAuthType {
	return &e
}

func (e *DestinationSnowflakeUpdateAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Key Pair Authentication":
		*e = DestinationSnowflakeUpdateAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationSnowflakeUpdateAuthType: %v", v)
	}
}

type KeyPairAuthentication struct {
	authType *DestinationSnowflakeUpdateAuthType `const:"Key Pair Authentication" json:"auth_type"`
	// RSA Private key to use for Snowflake connection. See the <a href="https://docs.airbyte.com/integrations/destinations/snowflake">docs</a> for more information on how to obtain this key.
	PrivateKey string `json:"private_key"`
	// Passphrase for private key
	PrivateKeyPassword *string `json:"private_key_password,omitempty"`
}

func (k KeyPairAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(k, "", false)
}

func (k *KeyPairAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &k, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *KeyPairAuthentication) GetAuthType() *DestinationSnowflakeUpdateAuthType {
	return DestinationSnowflakeUpdateAuthTypeKeyPairAuthentication.ToPointer()
}

func (o *KeyPairAuthentication) GetPrivateKey() string {
	if o == nil {
		return ""
	}
	return o.PrivateKey
}

func (o *KeyPairAuthentication) GetPrivateKeyPassword() *string {
	if o == nil {
		return nil
	}
	return o.PrivateKeyPassword
}

type DestinationSnowflakeUpdateSchemasCredentialsAuthType string

const (
	DestinationSnowflakeUpdateSchemasCredentialsAuthTypeOAuth20 DestinationSnowflakeUpdateSchemasCredentialsAuthType = "OAuth2.0"
)

func (e DestinationSnowflakeUpdateSchemasCredentialsAuthType) ToPointer() *DestinationSnowflakeUpdateSchemasCredentialsAuthType {
	return &e
}

func (e *DestinationSnowflakeUpdateSchemasCredentialsAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OAuth2.0":
		*e = DestinationSnowflakeUpdateSchemasCredentialsAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationSnowflakeUpdateSchemasCredentialsAuthType: %v", v)
	}
}

type OAuth20 struct {
	// Enter you application's Access Token
	AccessToken string                                                `json:"access_token"`
	authType    *DestinationSnowflakeUpdateSchemasCredentialsAuthType `const:"OAuth2.0" json:"auth_type"`
	// Enter your application's Client ID
	ClientID *string `json:"client_id,omitempty"`
	// Enter your application's Client secret
	ClientSecret *string `json:"client_secret,omitempty"`
	// Enter your application's Refresh Token
	RefreshToken string `json:"refresh_token"`
}

func (o OAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *OAuth20) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *OAuth20) GetAuthType() *DestinationSnowflakeUpdateSchemasCredentialsAuthType {
	return DestinationSnowflakeUpdateSchemasCredentialsAuthTypeOAuth20.ToPointer()
}

func (o *OAuth20) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *OAuth20) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

func (o *OAuth20) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type AuthorizationMethodType string

const (
	AuthorizationMethodTypeOAuth20               AuthorizationMethodType = "OAuth20"
	AuthorizationMethodTypeKeyPairAuthentication AuthorizationMethodType = "KeyPairAuthentication"
	AuthorizationMethodTypeUsernameAndPassword   AuthorizationMethodType = "UsernameAndPassword"
)

type AuthorizationMethod struct {
	OAuth20               *OAuth20
	KeyPairAuthentication *KeyPairAuthentication
	UsernameAndPassword   *UsernameAndPassword

	Type AuthorizationMethodType
}

func CreateAuthorizationMethodOAuth20(oAuth20 OAuth20) AuthorizationMethod {
	typ := AuthorizationMethodTypeOAuth20

	return AuthorizationMethod{
		OAuth20: &oAuth20,
		Type:    typ,
	}
}

func CreateAuthorizationMethodKeyPairAuthentication(keyPairAuthentication KeyPairAuthentication) AuthorizationMethod {
	typ := AuthorizationMethodTypeKeyPairAuthentication

	return AuthorizationMethod{
		KeyPairAuthentication: &keyPairAuthentication,
		Type:                  typ,
	}
}

func CreateAuthorizationMethodUsernameAndPassword(usernameAndPassword UsernameAndPassword) AuthorizationMethod {
	typ := AuthorizationMethodTypeUsernameAndPassword

	return AuthorizationMethod{
		UsernameAndPassword: &usernameAndPassword,
		Type:                typ,
	}
}

func (u *AuthorizationMethod) UnmarshalJSON(data []byte) error {

	usernameAndPassword := new(UsernameAndPassword)
	if err := utils.UnmarshalJSON(data, &usernameAndPassword, "", true, true); err == nil {
		u.UsernameAndPassword = usernameAndPassword
		u.Type = AuthorizationMethodTypeUsernameAndPassword
		return nil
	}

	keyPairAuthentication := new(KeyPairAuthentication)
	if err := utils.UnmarshalJSON(data, &keyPairAuthentication, "", true, true); err == nil {
		u.KeyPairAuthentication = keyPairAuthentication
		u.Type = AuthorizationMethodTypeKeyPairAuthentication
		return nil
	}

	oAuth20 := new(OAuth20)
	if err := utils.UnmarshalJSON(data, &oAuth20, "", true, true); err == nil {
		u.OAuth20 = oAuth20
		u.Type = AuthorizationMethodTypeOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.OAuth20 != nil {
		return utils.MarshalJSON(u.OAuth20, "", true)
	}

	if u.KeyPairAuthentication != nil {
		return utils.MarshalJSON(u.KeyPairAuthentication, "", true)
	}

	if u.UsernameAndPassword != nil {
		return utils.MarshalJSON(u.UsernameAndPassword, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationSnowflakeUpdate struct {
	Credentials *AuthorizationMethod `json:"credentials,omitempty"`
	// Enter the name of the <a href="https://docs.snowflake.com/en/sql-reference/ddl-database.html#database-schema-share-ddl">database</a> you want to sync data into
	Database string `json:"database"`
	// Enter your Snowflake account's <a href="https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#using-an-account-locator-as-an-identifier">locator</a> (in the format <account_locator>.<region>.<cloud>.snowflakecomputing.com)
	Host string `json:"host"`
	// Enter the additional properties to pass to the JDBC URL string when connecting to the database (formatted as key=value pairs separated by the symbol &). Example: key1=value1&key2=value2&key3=value3
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// The schema to write raw tables into
	RawDataSchema *string `json:"raw_data_schema,omitempty"`
	// Enter the <a href="https://docs.snowflake.com/en/user-guide/security-access-control-overview.html#roles">role</a> that you want to use to access Snowflake
	Role string `json:"role"`
	// Enter the name of the default <a href="https://docs.snowflake.com/en/sql-reference/ddl-database.html#database-schema-share-ddl">schema</a>
	Schema string `json:"schema"`
	// Enter the name of the user you want to use to access the database
	Username string `json:"username"`
	// Enter the name of the <a href="https://docs.snowflake.com/en/user-guide/warehouses-overview.html#overview-of-warehouses">warehouse</a> that you want to sync data into
	Warehouse string `json:"warehouse"`
}

func (o *DestinationSnowflakeUpdate) GetCredentials() *AuthorizationMethod {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *DestinationSnowflakeUpdate) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationSnowflakeUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationSnowflakeUpdate) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationSnowflakeUpdate) GetRawDataSchema() *string {
	if o == nil {
		return nil
	}
	return o.RawDataSchema
}

func (o *DestinationSnowflakeUpdate) GetRole() string {
	if o == nil {
		return ""
	}
	return o.Role
}

func (o *DestinationSnowflakeUpdate) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *DestinationSnowflakeUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *DestinationSnowflakeUpdate) GetWarehouse() string {
	if o == nil {
		return ""
	}
	return o.Warehouse
}
