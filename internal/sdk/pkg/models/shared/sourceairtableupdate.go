// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"time"
)

type SourceAirtableUpdateAuthMethod string

const (
	SourceAirtableUpdateAuthMethodAPIKey SourceAirtableUpdateAuthMethod = "api_key"
)

func (e SourceAirtableUpdateAuthMethod) ToPointer() *SourceAirtableUpdateAuthMethod {
	return &e
}

func (e *SourceAirtableUpdateAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_key":
		*e = SourceAirtableUpdateAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAirtableUpdateAuthMethod: %v", v)
	}
}

type PersonalAccessToken struct {
	// The Personal Access Token for the Airtable account. See the <a href="https://airtable.com/developers/web/guides/personal-access-tokens">Support Guide</a> for more information on how to obtain this token.
	APIKey     string                          `json:"api_key"`
	authMethod *SourceAirtableUpdateAuthMethod `const:"api_key" json:"auth_method,omitempty"`
}

func (p PersonalAccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonalAccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *PersonalAccessToken) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *PersonalAccessToken) GetAuthMethod() *SourceAirtableUpdateAuthMethod {
	return SourceAirtableUpdateAuthMethodAPIKey.ToPointer()
}

type SourceAirtableUpdateSchemasAuthMethod string

const (
	SourceAirtableUpdateSchemasAuthMethodOauth20 SourceAirtableUpdateSchemasAuthMethod = "oauth2.0"
)

func (e SourceAirtableUpdateSchemasAuthMethod) ToPointer() *SourceAirtableUpdateSchemasAuthMethod {
	return &e
}

func (e *SourceAirtableUpdateSchemasAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceAirtableUpdateSchemasAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAirtableUpdateSchemasAuthMethod: %v", v)
	}
}

type SourceAirtableUpdateOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken *string                                `json:"access_token,omitempty"`
	authMethod  *SourceAirtableUpdateSchemasAuthMethod `const:"oauth2.0" json:"auth_method,omitempty"`
	// The client ID of the Airtable developer application.
	ClientID string `json:"client_id"`
	// The client secret the Airtable developer application.
	ClientSecret string `json:"client_secret"`
	// The key to refresh the expired access token.
	RefreshToken string `json:"refresh_token"`
	// The date-time when the access token should be refreshed.
	TokenExpiryDate *time.Time `json:"token_expiry_date,omitempty"`
}

func (s SourceAirtableUpdateOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAirtableUpdateOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceAirtableUpdateOAuth20) GetAccessToken() *string {
	if o == nil {
		return nil
	}
	return o.AccessToken
}

func (o *SourceAirtableUpdateOAuth20) GetAuthMethod() *SourceAirtableUpdateSchemasAuthMethod {
	return SourceAirtableUpdateSchemasAuthMethodOauth20.ToPointer()
}

func (o *SourceAirtableUpdateOAuth20) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceAirtableUpdateOAuth20) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceAirtableUpdateOAuth20) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

func (o *SourceAirtableUpdateOAuth20) GetTokenExpiryDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.TokenExpiryDate
}

type SourceAirtableUpdateAuthenticationType string

const (
	SourceAirtableUpdateAuthenticationTypeOAuth20             SourceAirtableUpdateAuthenticationType = "OAuth20"
	SourceAirtableUpdateAuthenticationTypePersonalAccessToken SourceAirtableUpdateAuthenticationType = "PersonalAccessToken"
)

type SourceAirtableUpdateAuthentication struct {
	OAuth20             *SourceAirtableUpdateOAuth20
	PersonalAccessToken *PersonalAccessToken

	Type SourceAirtableUpdateAuthenticationType
}

func CreateSourceAirtableUpdateAuthenticationOAuth20(oAuth20 SourceAirtableUpdateOAuth20) SourceAirtableUpdateAuthentication {
	typ := SourceAirtableUpdateAuthenticationTypeOAuth20

	return SourceAirtableUpdateAuthentication{
		OAuth20: &oAuth20,
		Type:    typ,
	}
}

func CreateSourceAirtableUpdateAuthenticationPersonalAccessToken(personalAccessToken PersonalAccessToken) SourceAirtableUpdateAuthentication {
	typ := SourceAirtableUpdateAuthenticationTypePersonalAccessToken

	return SourceAirtableUpdateAuthentication{
		PersonalAccessToken: &personalAccessToken,
		Type:                typ,
	}
}

func (u *SourceAirtableUpdateAuthentication) UnmarshalJSON(data []byte) error {

	personalAccessToken := new(PersonalAccessToken)
	if err := utils.UnmarshalJSON(data, &personalAccessToken, "", true, true); err == nil {
		u.PersonalAccessToken = personalAccessToken
		u.Type = SourceAirtableUpdateAuthenticationTypePersonalAccessToken
		return nil
	}

	oAuth20 := new(SourceAirtableUpdateOAuth20)
	if err := utils.UnmarshalJSON(data, &oAuth20, "", true, true); err == nil {
		u.OAuth20 = oAuth20
		u.Type = SourceAirtableUpdateAuthenticationTypeOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceAirtableUpdateAuthentication) MarshalJSON() ([]byte, error) {
	if u.OAuth20 != nil {
		return utils.MarshalJSON(u.OAuth20, "", true)
	}

	if u.PersonalAccessToken != nil {
		return utils.MarshalJSON(u.PersonalAccessToken, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceAirtableUpdate struct {
	Credentials *SourceAirtableUpdateAuthentication `json:"credentials,omitempty"`
}

func (o *SourceAirtableUpdate) GetCredentials() *SourceAirtableUpdateAuthentication {
	if o == nil {
		return nil
	}
	return o.Credentials
}
