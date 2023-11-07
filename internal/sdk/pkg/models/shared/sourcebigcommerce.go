// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type Bigcommerce string

const (
	BigcommerceBigcommerce Bigcommerce = "bigcommerce"
)

func (e Bigcommerce) ToPointer() *Bigcommerce {
	return &e
}

func (e *Bigcommerce) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bigcommerce":
		*e = Bigcommerce(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Bigcommerce: %v", v)
	}
}

type SourceBigcommerce struct {
	// Access Token for making authenticated requests.
	AccessToken string      `json:"access_token"`
	sourceType  Bigcommerce `const:"bigcommerce" json:"sourceType"`
	// The date you would like to replicate data. Format: YYYY-MM-DD.
	StartDate string `json:"start_date"`
	// The hash code of the store. For https://api.bigcommerce.com/stores/HASH_CODE/v3/, The store's hash code is 'HASH_CODE'.
	StoreHash string `json:"store_hash"`
}

func (s SourceBigcommerce) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceBigcommerce) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceBigcommerce) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceBigcommerce) GetSourceType() Bigcommerce {
	return BigcommerceBigcommerce
}

func (o *SourceBigcommerce) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}

func (o *SourceBigcommerce) GetStoreHash() string {
	if o == nil {
		return ""
	}
	return o.StoreHash
}
