// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourceMondayUpdateSchemasAuthType string

const (
	SourceMondayUpdateSchemasAuthTypeAPIToken SourceMondayUpdateSchemasAuthType = "api_token"
)

func (e SourceMondayUpdateSchemasAuthType) ToPointer() *SourceMondayUpdateSchemasAuthType {
	return &e
}

func (e *SourceMondayUpdateSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_token":
		*e = SourceMondayUpdateSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMondayUpdateSchemasAuthType: %v", v)
	}
}

type APIToken struct {
	// API Token for making authenticated requests.
	APIToken string                            `json:"api_token"`
	authType SourceMondayUpdateSchemasAuthType `const:"api_token" json:"auth_type"`
}

func (a APIToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *APIToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *APIToken) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *APIToken) GetAuthType() SourceMondayUpdateSchemasAuthType {
	return SourceMondayUpdateSchemasAuthTypeAPIToken
}

type SourceMondayUpdateAuthType string

const (
	SourceMondayUpdateAuthTypeOauth20 SourceMondayUpdateAuthType = "oauth2.0"
)

func (e SourceMondayUpdateAuthType) ToPointer() *SourceMondayUpdateAuthType {
	return &e
}

func (e *SourceMondayUpdateAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceMondayUpdateAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMondayUpdateAuthType: %v", v)
	}
}

type SourceMondayUpdateOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken string                     `json:"access_token"`
	authType    SourceMondayUpdateAuthType `const:"oauth2.0" json:"auth_type"`
	// The Client ID of your OAuth application.
	ClientID string `json:"client_id"`
	// The Client Secret of your OAuth application.
	ClientSecret string `json:"client_secret"`
	// Slug/subdomain of the account, or the first part of the URL that comes before .monday.com
	Subdomain *string `default:"" json:"subdomain"`
}

func (s SourceMondayUpdateOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMondayUpdateOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMondayUpdateOAuth20) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceMondayUpdateOAuth20) GetAuthType() SourceMondayUpdateAuthType {
	return SourceMondayUpdateAuthTypeOauth20
}

func (o *SourceMondayUpdateOAuth20) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceMondayUpdateOAuth20) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceMondayUpdateOAuth20) GetSubdomain() *string {
	if o == nil {
		return nil
	}
	return o.Subdomain
}

type SourceMondayUpdateAuthorizationMethodType string

const (
	SourceMondayUpdateAuthorizationMethodTypeOAuth20  SourceMondayUpdateAuthorizationMethodType = "OAuth20"
	SourceMondayUpdateAuthorizationMethodTypeAPIToken SourceMondayUpdateAuthorizationMethodType = "APIToken"
)

type SourceMondayUpdateAuthorizationMethod struct {
	OAuth20  *SourceMondayUpdateOAuth20
	APIToken *APIToken

	Type SourceMondayUpdateAuthorizationMethodType
}

func CreateSourceMondayUpdateAuthorizationMethodOAuth20(oAuth20 SourceMondayUpdateOAuth20) SourceMondayUpdateAuthorizationMethod {
	typ := SourceMondayUpdateAuthorizationMethodTypeOAuth20

	return SourceMondayUpdateAuthorizationMethod{
		OAuth20: &oAuth20,
		Type:    typ,
	}
}

func CreateSourceMondayUpdateAuthorizationMethodAPIToken(apiToken APIToken) SourceMondayUpdateAuthorizationMethod {
	typ := SourceMondayUpdateAuthorizationMethodTypeAPIToken

	return SourceMondayUpdateAuthorizationMethod{
		APIToken: &apiToken,
		Type:     typ,
	}
}

func (u *SourceMondayUpdateAuthorizationMethod) UnmarshalJSON(data []byte) error {

	apiToken := new(APIToken)
	if err := utils.UnmarshalJSON(data, &apiToken, "", true, true); err == nil {
		u.APIToken = apiToken
		u.Type = SourceMondayUpdateAuthorizationMethodTypeAPIToken
		return nil
	}

	oAuth20 := new(SourceMondayUpdateOAuth20)
	if err := utils.UnmarshalJSON(data, &oAuth20, "", true, true); err == nil {
		u.OAuth20 = oAuth20
		u.Type = SourceMondayUpdateAuthorizationMethodTypeOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMondayUpdateAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.OAuth20 != nil {
		return utils.MarshalJSON(u.OAuth20, "", true)
	}

	if u.APIToken != nil {
		return utils.MarshalJSON(u.APIToken, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceMondayUpdate struct {
	Credentials *SourceMondayUpdateAuthorizationMethod `json:"credentials,omitempty"`
}

func (o *SourceMondayUpdate) GetCredentials() *SourceMondayUpdateAuthorizationMethod {
	if o == nil {
		return nil
	}
	return o.Credentials
}
