// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type SourceAirtableSchemasAuthMethod string

const (
	SourceAirtableSchemasAuthMethodAPIKey SourceAirtableSchemasAuthMethod = "api_key"
)

func (e SourceAirtableSchemasAuthMethod) ToPointer() *SourceAirtableSchemasAuthMethod {
	return &e
}
func (e *SourceAirtableSchemasAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_key":
		*e = SourceAirtableSchemasAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAirtableSchemasAuthMethod: %v", v)
	}
}

type SourceAirtablePersonalAccessToken struct {
	// The Personal Access Token for the Airtable account. See the <a href="https://airtable.com/developers/web/guides/personal-access-tokens">Support Guide</a> for more information on how to obtain this token.
	APIKey     string                           `json:"api_key"`
	authMethod *SourceAirtableSchemasAuthMethod `const:"api_key" json:"auth_method,omitempty"`
}

func (s SourceAirtablePersonalAccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAirtablePersonalAccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceAirtablePersonalAccessToken) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceAirtablePersonalAccessToken) GetAuthMethod() *SourceAirtableSchemasAuthMethod {
	return SourceAirtableSchemasAuthMethodAPIKey.ToPointer()
}

type SourceAirtableAuthMethod string

const (
	SourceAirtableAuthMethodOauth20 SourceAirtableAuthMethod = "oauth2.0"
)

func (e SourceAirtableAuthMethod) ToPointer() *SourceAirtableAuthMethod {
	return &e
}
func (e *SourceAirtableAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceAirtableAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAirtableAuthMethod: %v", v)
	}
}

type SourceAirtableOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken *string                   `json:"access_token,omitempty"`
	authMethod  *SourceAirtableAuthMethod `const:"oauth2.0" json:"auth_method,omitempty"`
	// The client ID of the Airtable developer application.
	ClientID string `json:"client_id"`
	// The client secret the Airtable developer application.
	ClientSecret string `json:"client_secret"`
	// The key to refresh the expired access token.
	RefreshToken string `json:"refresh_token"`
	// The date-time when the access token should be refreshed.
	TokenExpiryDate *time.Time `json:"token_expiry_date,omitempty"`
}

func (s SourceAirtableOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAirtableOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceAirtableOAuth20) GetAccessToken() *string {
	if o == nil {
		return nil
	}
	return o.AccessToken
}

func (o *SourceAirtableOAuth20) GetAuthMethod() *SourceAirtableAuthMethod {
	return SourceAirtableAuthMethodOauth20.ToPointer()
}

func (o *SourceAirtableOAuth20) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceAirtableOAuth20) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceAirtableOAuth20) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

func (o *SourceAirtableOAuth20) GetTokenExpiryDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.TokenExpiryDate
}

type SourceAirtableAuthenticationType string

const (
	SourceAirtableAuthenticationTypeSourceAirtableOAuth20             SourceAirtableAuthenticationType = "source-airtable_OAuth2.0"
	SourceAirtableAuthenticationTypeSourceAirtablePersonalAccessToken SourceAirtableAuthenticationType = "source-airtable_Personal Access Token"
)

type SourceAirtableAuthentication struct {
	SourceAirtableOAuth20             *SourceAirtableOAuth20
	SourceAirtablePersonalAccessToken *SourceAirtablePersonalAccessToken

	Type SourceAirtableAuthenticationType
}

func CreateSourceAirtableAuthenticationSourceAirtableOAuth20(sourceAirtableOAuth20 SourceAirtableOAuth20) SourceAirtableAuthentication {
	typ := SourceAirtableAuthenticationTypeSourceAirtableOAuth20

	return SourceAirtableAuthentication{
		SourceAirtableOAuth20: &sourceAirtableOAuth20,
		Type:                  typ,
	}
}

func CreateSourceAirtableAuthenticationSourceAirtablePersonalAccessToken(sourceAirtablePersonalAccessToken SourceAirtablePersonalAccessToken) SourceAirtableAuthentication {
	typ := SourceAirtableAuthenticationTypeSourceAirtablePersonalAccessToken

	return SourceAirtableAuthentication{
		SourceAirtablePersonalAccessToken: &sourceAirtablePersonalAccessToken,
		Type:                              typ,
	}
}

func (u *SourceAirtableAuthentication) UnmarshalJSON(data []byte) error {

	var sourceAirtablePersonalAccessToken SourceAirtablePersonalAccessToken = SourceAirtablePersonalAccessToken{}
	if err := utils.UnmarshalJSON(data, &sourceAirtablePersonalAccessToken, "", true, true); err == nil {
		u.SourceAirtablePersonalAccessToken = &sourceAirtablePersonalAccessToken
		u.Type = SourceAirtableAuthenticationTypeSourceAirtablePersonalAccessToken
		return nil
	}

	var sourceAirtableOAuth20 SourceAirtableOAuth20 = SourceAirtableOAuth20{}
	if err := utils.UnmarshalJSON(data, &sourceAirtableOAuth20, "", true, true); err == nil {
		u.SourceAirtableOAuth20 = &sourceAirtableOAuth20
		u.Type = SourceAirtableAuthenticationTypeSourceAirtableOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceAirtableAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceAirtableOAuth20 != nil {
		return utils.MarshalJSON(u.SourceAirtableOAuth20, "", true)
	}

	if u.SourceAirtablePersonalAccessToken != nil {
		return utils.MarshalJSON(u.SourceAirtablePersonalAccessToken, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type Airtable string

const (
	AirtableAirtable Airtable = "airtable"
)

func (e Airtable) ToPointer() *Airtable {
	return &e
}
func (e *Airtable) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "airtable":
		*e = Airtable(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Airtable: %v", v)
	}
}

type SourceAirtable struct {
	Credentials *SourceAirtableAuthentication `json:"credentials,omitempty"`
	sourceType  *Airtable                     `const:"airtable" json:"sourceType,omitempty"`
}

func (s SourceAirtable) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAirtable) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAirtable) GetCredentials() *SourceAirtableAuthentication {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceAirtable) GetSourceType() *Airtable {
	return AirtableAirtable.ToPointer()
}
