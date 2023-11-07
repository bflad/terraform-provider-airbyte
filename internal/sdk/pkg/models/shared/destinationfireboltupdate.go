// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type DestinationFireboltUpdateSchemasMethod string

const (
	DestinationFireboltUpdateSchemasMethodS3 DestinationFireboltUpdateSchemasMethod = "S3"
)

func (e DestinationFireboltUpdateSchemasMethod) ToPointer() *DestinationFireboltUpdateSchemasMethod {
	return &e
}

func (e *DestinationFireboltUpdateSchemasMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "S3":
		*e = DestinationFireboltUpdateSchemasMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationFireboltUpdateSchemasMethod: %v", v)
	}
}

// ExternalTableViaS3 - Loading method used to select the way data will be uploaded to Firebolt
type ExternalTableViaS3 struct {
	// AWS access key granting read and write access to S3.
	AwsKeyID string `json:"aws_key_id"`
	// Corresponding secret part of the AWS Key
	AwsKeySecret string                                 `json:"aws_key_secret"`
	method       DestinationFireboltUpdateSchemasMethod `const:"S3" json:"method"`
	// The name of the S3 bucket.
	S3Bucket string `json:"s3_bucket"`
	// Region name of the S3 bucket.
	S3Region string `json:"s3_region"`
}

func (e ExternalTableViaS3) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *ExternalTableViaS3) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ExternalTableViaS3) GetAwsKeyID() string {
	if o == nil {
		return ""
	}
	return o.AwsKeyID
}

func (o *ExternalTableViaS3) GetAwsKeySecret() string {
	if o == nil {
		return ""
	}
	return o.AwsKeySecret
}

func (o *ExternalTableViaS3) GetMethod() DestinationFireboltUpdateSchemasMethod {
	return DestinationFireboltUpdateSchemasMethodS3
}

func (o *ExternalTableViaS3) GetS3Bucket() string {
	if o == nil {
		return ""
	}
	return o.S3Bucket
}

func (o *ExternalTableViaS3) GetS3Region() string {
	if o == nil {
		return ""
	}
	return o.S3Region
}

type DestinationFireboltUpdateMethod string

const (
	DestinationFireboltUpdateMethodSQL DestinationFireboltUpdateMethod = "SQL"
)

func (e DestinationFireboltUpdateMethod) ToPointer() *DestinationFireboltUpdateMethod {
	return &e
}

func (e *DestinationFireboltUpdateMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SQL":
		*e = DestinationFireboltUpdateMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationFireboltUpdateMethod: %v", v)
	}
}

// SQLInserts - Loading method used to select the way data will be uploaded to Firebolt
type SQLInserts struct {
	method DestinationFireboltUpdateMethod `const:"SQL" json:"method"`
}

func (s SQLInserts) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SQLInserts) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SQLInserts) GetMethod() DestinationFireboltUpdateMethod {
	return DestinationFireboltUpdateMethodSQL
}

type DestinationFireboltUpdateLoadingMethodType string

const (
	DestinationFireboltUpdateLoadingMethodTypeSQLInserts         DestinationFireboltUpdateLoadingMethodType = "SQLInserts"
	DestinationFireboltUpdateLoadingMethodTypeExternalTableViaS3 DestinationFireboltUpdateLoadingMethodType = "ExternalTableViaS3"
)

type DestinationFireboltUpdateLoadingMethod struct {
	SQLInserts         *SQLInserts
	ExternalTableViaS3 *ExternalTableViaS3

	Type DestinationFireboltUpdateLoadingMethodType
}

func CreateDestinationFireboltUpdateLoadingMethodSQLInserts(sqlInserts SQLInserts) DestinationFireboltUpdateLoadingMethod {
	typ := DestinationFireboltUpdateLoadingMethodTypeSQLInserts

	return DestinationFireboltUpdateLoadingMethod{
		SQLInserts: &sqlInserts,
		Type:       typ,
	}
}

func CreateDestinationFireboltUpdateLoadingMethodExternalTableViaS3(externalTableViaS3 ExternalTableViaS3) DestinationFireboltUpdateLoadingMethod {
	typ := DestinationFireboltUpdateLoadingMethodTypeExternalTableViaS3

	return DestinationFireboltUpdateLoadingMethod{
		ExternalTableViaS3: &externalTableViaS3,
		Type:               typ,
	}
}

func (u *DestinationFireboltUpdateLoadingMethod) UnmarshalJSON(data []byte) error {

	sqlInserts := new(SQLInserts)
	if err := utils.UnmarshalJSON(data, &sqlInserts, "", true, true); err == nil {
		u.SQLInserts = sqlInserts
		u.Type = DestinationFireboltUpdateLoadingMethodTypeSQLInserts
		return nil
	}

	externalTableViaS3 := new(ExternalTableViaS3)
	if err := utils.UnmarshalJSON(data, &externalTableViaS3, "", true, true); err == nil {
		u.ExternalTableViaS3 = externalTableViaS3
		u.Type = DestinationFireboltUpdateLoadingMethodTypeExternalTableViaS3
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationFireboltUpdateLoadingMethod) MarshalJSON() ([]byte, error) {
	if u.SQLInserts != nil {
		return utils.MarshalJSON(u.SQLInserts, "", true)
	}

	if u.ExternalTableViaS3 != nil {
		return utils.MarshalJSON(u.ExternalTableViaS3, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationFireboltUpdate struct {
	// Firebolt account to login.
	Account *string `json:"account,omitempty"`
	// The database to connect to.
	Database string `json:"database"`
	// Engine name or url to connect to.
	Engine *string `json:"engine,omitempty"`
	// The host name of your Firebolt database.
	Host *string `json:"host,omitempty"`
	// Loading method used to select the way data will be uploaded to Firebolt
	LoadingMethod *DestinationFireboltUpdateLoadingMethod `json:"loading_method,omitempty"`
	// Firebolt password.
	Password string `json:"password"`
	// Firebolt email address you use to login.
	Username string `json:"username"`
}

func (o *DestinationFireboltUpdate) GetAccount() *string {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *DestinationFireboltUpdate) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationFireboltUpdate) GetEngine() *string {
	if o == nil {
		return nil
	}
	return o.Engine
}

func (o *DestinationFireboltUpdate) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *DestinationFireboltUpdate) GetLoadingMethod() *DestinationFireboltUpdateLoadingMethod {
	if o == nil {
		return nil
	}
	return o.LoadingMethod
}

func (o *DestinationFireboltUpdate) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *DestinationFireboltUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
