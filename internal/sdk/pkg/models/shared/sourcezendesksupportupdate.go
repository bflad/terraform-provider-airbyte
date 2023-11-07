// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"time"
)

type SourceZendeskSupportUpdateSchemasCredentials string

const (
	SourceZendeskSupportUpdateSchemasCredentialsAPIToken SourceZendeskSupportUpdateSchemasCredentials = "api_token"
)

func (e SourceZendeskSupportUpdateSchemasCredentials) ToPointer() *SourceZendeskSupportUpdateSchemasCredentials {
	return &e
}

func (e *SourceZendeskSupportUpdateSchemasCredentials) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_token":
		*e = SourceZendeskSupportUpdateSchemasCredentials(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskSupportUpdateSchemasCredentials: %v", v)
	}
}

// SourceZendeskSupportUpdateAPIToken - Zendesk allows two authentication methods. We recommend using `OAuth2.0` for Airbyte Cloud users and `API token` for Airbyte Open Source users.
type SourceZendeskSupportUpdateAPIToken struct {
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
	// The value of the API token generated. See our <a href="https://docs.airbyte.com/integrations/sources/zendesk-support#setup-guide">full documentation</a> for more information on generating this token.
	APIToken    string                                        `json:"api_token"`
	credentials *SourceZendeskSupportUpdateSchemasCredentials `const:"api_token" json:"credentials,omitempty"`
	// The user email for your Zendesk account.
	Email string `json:"email"`
}

func (s SourceZendeskSupportUpdateAPIToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskSupportUpdateAPIToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskSupportUpdateAPIToken) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SourceZendeskSupportUpdateAPIToken) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceZendeskSupportUpdateAPIToken) GetCredentials() *SourceZendeskSupportUpdateSchemasCredentials {
	return SourceZendeskSupportUpdateSchemasCredentialsAPIToken.ToPointer()
}

func (o *SourceZendeskSupportUpdateAPIToken) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

type SourceZendeskSupportUpdateCredentials string

const (
	SourceZendeskSupportUpdateCredentialsOauth20 SourceZendeskSupportUpdateCredentials = "oauth2.0"
)

func (e SourceZendeskSupportUpdateCredentials) ToPointer() *SourceZendeskSupportUpdateCredentials {
	return &e
}

func (e *SourceZendeskSupportUpdateCredentials) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceZendeskSupportUpdateCredentials(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskSupportUpdateCredentials: %v", v)
	}
}

// SourceZendeskSupportUpdateOAuth20 - Zendesk allows two authentication methods. We recommend using `OAuth2.0` for Airbyte Cloud users and `API token` for Airbyte Open Source users.
type SourceZendeskSupportUpdateOAuth20 struct {
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
	// The OAuth access token. See the <a href="https://developer.zendesk.com/documentation/ticketing/working-with-oauth/creating-and-using-oauth-tokens-with-the-api/">Zendesk docs</a> for more information on generating this token.
	AccessToken string `json:"access_token"`
	// The OAuth client's ID. See <a href="https://docs.searchunify.com/Content/Content-Sources/Zendesk-Authentication-OAuth-Client-ID-Secret.htm#:~:text=Get%20Client%20ID%20and%20Client%20Secret&text=Go%20to%20OAuth%20Clients%20and,will%20be%20displayed%20only%20once.">this guide</a> for more information.
	ClientID *string `json:"client_id,omitempty"`
	// The OAuth client secret. See <a href="https://docs.searchunify.com/Content/Content-Sources/Zendesk-Authentication-OAuth-Client-ID-Secret.htm#:~:text=Get%20Client%20ID%20and%20Client%20Secret&text=Go%20to%20OAuth%20Clients%20and,will%20be%20displayed%20only%20once.">this guide</a> for more information.
	ClientSecret *string                                `json:"client_secret,omitempty"`
	credentials  *SourceZendeskSupportUpdateCredentials `const:"oauth2.0" json:"credentials,omitempty"`
}

func (s SourceZendeskSupportUpdateOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskSupportUpdateOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskSupportUpdateOAuth20) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SourceZendeskSupportUpdateOAuth20) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceZendeskSupportUpdateOAuth20) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *SourceZendeskSupportUpdateOAuth20) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

func (o *SourceZendeskSupportUpdateOAuth20) GetCredentials() *SourceZendeskSupportUpdateCredentials {
	return SourceZendeskSupportUpdateCredentialsOauth20.ToPointer()
}

type SourceZendeskSupportUpdateAuthenticationType string

const (
	SourceZendeskSupportUpdateAuthenticationTypeOAuth20  SourceZendeskSupportUpdateAuthenticationType = "OAuth20"
	SourceZendeskSupportUpdateAuthenticationTypeAPIToken SourceZendeskSupportUpdateAuthenticationType = "APIToken"
)

type SourceZendeskSupportUpdateAuthentication struct {
	OAuth20  *SourceZendeskSupportUpdateOAuth20
	APIToken *SourceZendeskSupportUpdateAPIToken

	Type SourceZendeskSupportUpdateAuthenticationType
}

func CreateSourceZendeskSupportUpdateAuthenticationOAuth20(oAuth20 SourceZendeskSupportUpdateOAuth20) SourceZendeskSupportUpdateAuthentication {
	typ := SourceZendeskSupportUpdateAuthenticationTypeOAuth20

	return SourceZendeskSupportUpdateAuthentication{
		OAuth20: &oAuth20,
		Type:    typ,
	}
}

func CreateSourceZendeskSupportUpdateAuthenticationAPIToken(apiToken SourceZendeskSupportUpdateAPIToken) SourceZendeskSupportUpdateAuthentication {
	typ := SourceZendeskSupportUpdateAuthenticationTypeAPIToken

	return SourceZendeskSupportUpdateAuthentication{
		APIToken: &apiToken,
		Type:     typ,
	}
}

func (u *SourceZendeskSupportUpdateAuthentication) UnmarshalJSON(data []byte) error {

	apiToken := new(SourceZendeskSupportUpdateAPIToken)
	if err := utils.UnmarshalJSON(data, &apiToken, "", true, true); err == nil {
		u.APIToken = apiToken
		u.Type = SourceZendeskSupportUpdateAuthenticationTypeAPIToken
		return nil
	}

	oAuth20 := new(SourceZendeskSupportUpdateOAuth20)
	if err := utils.UnmarshalJSON(data, &oAuth20, "", true, true); err == nil {
		u.OAuth20 = oAuth20
		u.Type = SourceZendeskSupportUpdateAuthenticationTypeOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceZendeskSupportUpdateAuthentication) MarshalJSON() ([]byte, error) {
	if u.OAuth20 != nil {
		return utils.MarshalJSON(u.OAuth20, "", true)
	}

	if u.APIToken != nil {
		return utils.MarshalJSON(u.APIToken, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceZendeskSupportUpdate struct {
	// Zendesk allows two authentication methods. We recommend using `OAuth2.0` for Airbyte Cloud users and `API token` for Airbyte Open Source users.
	Credentials *SourceZendeskSupportUpdateAuthentication `json:"credentials,omitempty"`
	// Makes each stream read a single page of data.
	IgnorePagination *bool `default:"false" json:"ignore_pagination"`
	// The UTC date and time from which you'd like to replicate data, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
	StartDate *time.Time `json:"start_date,omitempty"`
	// This is your unique Zendesk subdomain that can be found in your account URL. For example, in https://MY_SUBDOMAIN.zendesk.com/, MY_SUBDOMAIN is the value of your subdomain.
	Subdomain string `json:"subdomain"`
}

func (s SourceZendeskSupportUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskSupportUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskSupportUpdate) GetCredentials() *SourceZendeskSupportUpdateAuthentication {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceZendeskSupportUpdate) GetIgnorePagination() *bool {
	if o == nil {
		return nil
	}
	return o.IgnorePagination
}

func (o *SourceZendeskSupportUpdate) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *SourceZendeskSupportUpdate) GetSubdomain() string {
	if o == nil {
		return ""
	}
	return o.Subdomain
}
