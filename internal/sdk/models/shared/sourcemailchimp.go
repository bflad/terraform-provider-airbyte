// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type SourceMailchimpSchemasAuthType string

const (
	SourceMailchimpSchemasAuthTypeApikey SourceMailchimpSchemasAuthType = "apikey"
)

func (e SourceMailchimpSchemasAuthType) ToPointer() *SourceMailchimpSchemasAuthType {
	return &e
}
func (e *SourceMailchimpSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "apikey":
		*e = SourceMailchimpSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMailchimpSchemasAuthType: %v", v)
	}
}

type SourceMailchimpAPIKey struct {
	// Mailchimp API Key. See the <a href="https://docs.airbyte.com/integrations/sources/mailchimp">docs</a> for information on how to generate this key.
	Apikey   string                         `json:"apikey"`
	authType SourceMailchimpSchemasAuthType `const:"apikey" json:"auth_type"`
}

func (s SourceMailchimpAPIKey) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMailchimpAPIKey) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMailchimpAPIKey) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

func (o *SourceMailchimpAPIKey) GetAuthType() SourceMailchimpSchemasAuthType {
	return SourceMailchimpSchemasAuthTypeApikey
}

type SourceMailchimpAuthType string

const (
	SourceMailchimpAuthTypeOauth20 SourceMailchimpAuthType = "oauth2.0"
)

func (e SourceMailchimpAuthType) ToPointer() *SourceMailchimpAuthType {
	return &e
}
func (e *SourceMailchimpAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceMailchimpAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMailchimpAuthType: %v", v)
	}
}

type SourceMailchimpOAuth20 struct {
	// An access token generated using the above client ID and secret.
	AccessToken string                  `json:"access_token"`
	authType    SourceMailchimpAuthType `const:"oauth2.0" json:"auth_type"`
	// The Client ID of your OAuth application.
	ClientID *string `json:"client_id,omitempty"`
	// The Client Secret of your OAuth application.
	ClientSecret *string `json:"client_secret,omitempty"`
}

func (s SourceMailchimpOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMailchimpOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMailchimpOAuth20) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceMailchimpOAuth20) GetAuthType() SourceMailchimpAuthType {
	return SourceMailchimpAuthTypeOauth20
}

func (o *SourceMailchimpOAuth20) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *SourceMailchimpOAuth20) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

type SourceMailchimpAuthenticationType string

const (
	SourceMailchimpAuthenticationTypeSourceMailchimpOAuth20 SourceMailchimpAuthenticationType = "source-mailchimp_OAuth2.0"
	SourceMailchimpAuthenticationTypeSourceMailchimpAPIKey  SourceMailchimpAuthenticationType = "source-mailchimp_API Key"
)

type SourceMailchimpAuthentication struct {
	SourceMailchimpOAuth20 *SourceMailchimpOAuth20
	SourceMailchimpAPIKey  *SourceMailchimpAPIKey

	Type SourceMailchimpAuthenticationType
}

func CreateSourceMailchimpAuthenticationSourceMailchimpOAuth20(sourceMailchimpOAuth20 SourceMailchimpOAuth20) SourceMailchimpAuthentication {
	typ := SourceMailchimpAuthenticationTypeSourceMailchimpOAuth20

	return SourceMailchimpAuthentication{
		SourceMailchimpOAuth20: &sourceMailchimpOAuth20,
		Type:                   typ,
	}
}

func CreateSourceMailchimpAuthenticationSourceMailchimpAPIKey(sourceMailchimpAPIKey SourceMailchimpAPIKey) SourceMailchimpAuthentication {
	typ := SourceMailchimpAuthenticationTypeSourceMailchimpAPIKey

	return SourceMailchimpAuthentication{
		SourceMailchimpAPIKey: &sourceMailchimpAPIKey,
		Type:                  typ,
	}
}

func (u *SourceMailchimpAuthentication) UnmarshalJSON(data []byte) error {

	var sourceMailchimpAPIKey SourceMailchimpAPIKey = SourceMailchimpAPIKey{}
	if err := utils.UnmarshalJSON(data, &sourceMailchimpAPIKey, "", true, true); err == nil {
		u.SourceMailchimpAPIKey = &sourceMailchimpAPIKey
		u.Type = SourceMailchimpAuthenticationTypeSourceMailchimpAPIKey
		return nil
	}

	var sourceMailchimpOAuth20 SourceMailchimpOAuth20 = SourceMailchimpOAuth20{}
	if err := utils.UnmarshalJSON(data, &sourceMailchimpOAuth20, "", true, true); err == nil {
		u.SourceMailchimpOAuth20 = &sourceMailchimpOAuth20
		u.Type = SourceMailchimpAuthenticationTypeSourceMailchimpOAuth20
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceMailchimpAuthentication", string(data))
}

func (u SourceMailchimpAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceMailchimpOAuth20 != nil {
		return utils.MarshalJSON(u.SourceMailchimpOAuth20, "", true)
	}

	if u.SourceMailchimpAPIKey != nil {
		return utils.MarshalJSON(u.SourceMailchimpAPIKey, "", true)
	}

	return nil, errors.New("could not marshal union type SourceMailchimpAuthentication: all fields are null")
}

type Mailchimp string

const (
	MailchimpMailchimp Mailchimp = "mailchimp"
)

func (e Mailchimp) ToPointer() *Mailchimp {
	return &e
}
func (e *Mailchimp) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mailchimp":
		*e = Mailchimp(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Mailchimp: %v", v)
	}
}

type SourceMailchimp struct {
	Credentials *SourceMailchimpAuthentication `json:"credentials,omitempty"`
	sourceType  Mailchimp                      `const:"mailchimp" json:"sourceType"`
	// The date from which you want to start syncing data for Incremental streams. Only records that have been created or modified since this date will be synced. If left blank, all data will by synced.
	StartDate *time.Time `json:"start_date,omitempty"`
}

func (s SourceMailchimp) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMailchimp) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceMailchimp) GetCredentials() *SourceMailchimpAuthentication {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceMailchimp) GetSourceType() Mailchimp {
	return MailchimpMailchimp
}

func (o *SourceMailchimp) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}
