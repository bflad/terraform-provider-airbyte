// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourceRetentlySchemasAuthType string

const (
	SourceRetentlySchemasAuthTypeToken SourceRetentlySchemasAuthType = "Token"
)

func (e SourceRetentlySchemasAuthType) ToPointer() *SourceRetentlySchemasAuthType {
	return &e
}

func (e *SourceRetentlySchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Token":
		*e = SourceRetentlySchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceRetentlySchemasAuthType: %v", v)
	}
}

// SourceRetentlyAuthenticateWithAPIToken - Choose how to authenticate to Retently
type SourceRetentlyAuthenticateWithAPIToken struct {
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
	// Retently API Token. See the <a href="https://app.retently.com/settings/api/tokens">docs</a> for more information on how to obtain this key.
	APIKey   string                         `json:"api_key"`
	authType *SourceRetentlySchemasAuthType `const:"Token" json:"auth_type,omitempty"`
}

func (s SourceRetentlyAuthenticateWithAPIToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceRetentlyAuthenticateWithAPIToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceRetentlyAuthenticateWithAPIToken) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SourceRetentlyAuthenticateWithAPIToken) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceRetentlyAuthenticateWithAPIToken) GetAuthType() *SourceRetentlySchemasAuthType {
	return SourceRetentlySchemasAuthTypeToken.ToPointer()
}

type SourceRetentlyAuthType string

const (
	SourceRetentlyAuthTypeClient SourceRetentlyAuthType = "Client"
)

func (e SourceRetentlyAuthType) ToPointer() *SourceRetentlyAuthType {
	return &e
}

func (e *SourceRetentlyAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceRetentlyAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceRetentlyAuthType: %v", v)
	}
}

// SourceRetentlyAuthenticateViaRetentlyOAuth - Choose how to authenticate to Retently
type SourceRetentlyAuthenticateViaRetentlyOAuth struct {
	AdditionalProperties interface{}             `additionalProperties:"true" json:"-"`
	authType             *SourceRetentlyAuthType `const:"Client" json:"auth_type,omitempty"`
	// The Client ID of your Retently developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Retently developer application.
	ClientSecret string `json:"client_secret"`
	// Retently Refresh Token which can be used to fetch new Bearer Tokens when the current one expires.
	RefreshToken string `json:"refresh_token"`
}

func (s SourceRetentlyAuthenticateViaRetentlyOAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceRetentlyAuthenticateViaRetentlyOAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceRetentlyAuthenticateViaRetentlyOAuth) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SourceRetentlyAuthenticateViaRetentlyOAuth) GetAuthType() *SourceRetentlyAuthType {
	return SourceRetentlyAuthTypeClient.ToPointer()
}

func (o *SourceRetentlyAuthenticateViaRetentlyOAuth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceRetentlyAuthenticateViaRetentlyOAuth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceRetentlyAuthenticateViaRetentlyOAuth) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceRetentlyAuthenticationMechanismType string

const (
	SourceRetentlyAuthenticationMechanismTypeAuthenticateViaRetentlyOAuth SourceRetentlyAuthenticationMechanismType = "AuthenticateViaRetentlyOAuth"
	SourceRetentlyAuthenticationMechanismTypeAuthenticateWithAPIToken     SourceRetentlyAuthenticationMechanismType = "AuthenticateWithAPIToken"
)

type SourceRetentlyAuthenticationMechanism struct {
	AuthenticateViaRetentlyOAuth *SourceRetentlyAuthenticateViaRetentlyOAuth
	AuthenticateWithAPIToken     *SourceRetentlyAuthenticateWithAPIToken

	Type SourceRetentlyAuthenticationMechanismType
}

func CreateSourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth(authenticateViaRetentlyOAuth SourceRetentlyAuthenticateViaRetentlyOAuth) SourceRetentlyAuthenticationMechanism {
	typ := SourceRetentlyAuthenticationMechanismTypeAuthenticateViaRetentlyOAuth

	return SourceRetentlyAuthenticationMechanism{
		AuthenticateViaRetentlyOAuth: &authenticateViaRetentlyOAuth,
		Type:                         typ,
	}
}

func CreateSourceRetentlyAuthenticationMechanismAuthenticateWithAPIToken(authenticateWithAPIToken SourceRetentlyAuthenticateWithAPIToken) SourceRetentlyAuthenticationMechanism {
	typ := SourceRetentlyAuthenticationMechanismTypeAuthenticateWithAPIToken

	return SourceRetentlyAuthenticationMechanism{
		AuthenticateWithAPIToken: &authenticateWithAPIToken,
		Type:                     typ,
	}
}

func (u *SourceRetentlyAuthenticationMechanism) UnmarshalJSON(data []byte) error {

	authenticateWithAPIToken := new(SourceRetentlyAuthenticateWithAPIToken)
	if err := utils.UnmarshalJSON(data, &authenticateWithAPIToken, "", true, true); err == nil {
		u.AuthenticateWithAPIToken = authenticateWithAPIToken
		u.Type = SourceRetentlyAuthenticationMechanismTypeAuthenticateWithAPIToken
		return nil
	}

	authenticateViaRetentlyOAuth := new(SourceRetentlyAuthenticateViaRetentlyOAuth)
	if err := utils.UnmarshalJSON(data, &authenticateViaRetentlyOAuth, "", true, true); err == nil {
		u.AuthenticateViaRetentlyOAuth = authenticateViaRetentlyOAuth
		u.Type = SourceRetentlyAuthenticationMechanismTypeAuthenticateViaRetentlyOAuth
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceRetentlyAuthenticationMechanism) MarshalJSON() ([]byte, error) {
	if u.AuthenticateViaRetentlyOAuth != nil {
		return utils.MarshalJSON(u.AuthenticateViaRetentlyOAuth, "", true)
	}

	if u.AuthenticateWithAPIToken != nil {
		return utils.MarshalJSON(u.AuthenticateWithAPIToken, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type Retently string

const (
	RetentlyRetently Retently = "retently"
)

func (e Retently) ToPointer() *Retently {
	return &e
}

func (e *Retently) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "retently":
		*e = Retently(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Retently: %v", v)
	}
}

type SourceRetently struct {
	// Choose how to authenticate to Retently
	Credentials *SourceRetentlyAuthenticationMechanism `json:"credentials,omitempty"`
	sourceType  *Retently                              `const:"retently" json:"sourceType,omitempty"`
}

func (s SourceRetently) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceRetently) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceRetently) GetCredentials() *SourceRetentlyAuthenticationMechanism {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceRetently) GetSourceType() *Retently {
	return RetentlyRetently.ToPointer()
}
