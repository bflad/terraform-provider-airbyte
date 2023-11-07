// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type DestinationDatabricksUpdateSchemasDataSourceType string

const (
	DestinationDatabricksUpdateSchemasDataSourceTypeAzureBlobStorage DestinationDatabricksUpdateSchemasDataSourceType = "AZURE_BLOB_STORAGE"
)

func (e DestinationDatabricksUpdateSchemasDataSourceType) ToPointer() *DestinationDatabricksUpdateSchemasDataSourceType {
	return &e
}

func (e *DestinationDatabricksUpdateSchemasDataSourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AZURE_BLOB_STORAGE":
		*e = DestinationDatabricksUpdateSchemasDataSourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksUpdateSchemasDataSourceType: %v", v)
	}
}

// DestinationDatabricksUpdateAzureBlobStorage - Storage on which the delta lake is built.
type DestinationDatabricksUpdateAzureBlobStorage struct {
	// The account's name of the Azure Blob Storage.
	AzureBlobStorageAccountName string `json:"azure_blob_storage_account_name"`
	// The name of the Azure blob storage container.
	AzureBlobStorageContainerName string `json:"azure_blob_storage_container_name"`
	// This is Azure Blob Storage endpoint domain name. Leave default value (or leave it empty if run container from command line) to use Microsoft native from example.
	AzureBlobStorageEndpointDomainName *string `default:"blob.core.windows.net" json:"azure_blob_storage_endpoint_domain_name"`
	// Shared access signature (SAS) token to grant limited access to objects in your storage account.
	AzureBlobStorageSasToken string                                           `json:"azure_blob_storage_sas_token"`
	dataSourceType           DestinationDatabricksUpdateSchemasDataSourceType `const:"AZURE_BLOB_STORAGE" json:"data_source_type"`
}

func (d DestinationDatabricksUpdateAzureBlobStorage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationDatabricksUpdateAzureBlobStorage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationDatabricksUpdateAzureBlobStorage) GetAzureBlobStorageAccountName() string {
	if o == nil {
		return ""
	}
	return o.AzureBlobStorageAccountName
}

func (o *DestinationDatabricksUpdateAzureBlobStorage) GetAzureBlobStorageContainerName() string {
	if o == nil {
		return ""
	}
	return o.AzureBlobStorageContainerName
}

func (o *DestinationDatabricksUpdateAzureBlobStorage) GetAzureBlobStorageEndpointDomainName() *string {
	if o == nil {
		return nil
	}
	return o.AzureBlobStorageEndpointDomainName
}

func (o *DestinationDatabricksUpdateAzureBlobStorage) GetAzureBlobStorageSasToken() string {
	if o == nil {
		return ""
	}
	return o.AzureBlobStorageSasToken
}

func (o *DestinationDatabricksUpdateAzureBlobStorage) GetDataSourceType() DestinationDatabricksUpdateSchemasDataSourceType {
	return DestinationDatabricksUpdateSchemasDataSourceTypeAzureBlobStorage
}

type DestinationDatabricksUpdateDataSourceType string

const (
	DestinationDatabricksUpdateDataSourceTypeS3Storage DestinationDatabricksUpdateDataSourceType = "S3_STORAGE"
)

func (e DestinationDatabricksUpdateDataSourceType) ToPointer() *DestinationDatabricksUpdateDataSourceType {
	return &e
}

func (e *DestinationDatabricksUpdateDataSourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "S3_STORAGE":
		*e = DestinationDatabricksUpdateDataSourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksUpdateDataSourceType: %v", v)
	}
}

// DestinationDatabricksUpdateS3BucketRegion - The region of the S3 staging bucket to use if utilising a copy strategy.
type DestinationDatabricksUpdateS3BucketRegion string

const (
	DestinationDatabricksUpdateS3BucketRegionUnknown      DestinationDatabricksUpdateS3BucketRegion = ""
	DestinationDatabricksUpdateS3BucketRegionUsEast1      DestinationDatabricksUpdateS3BucketRegion = "us-east-1"
	DestinationDatabricksUpdateS3BucketRegionUsEast2      DestinationDatabricksUpdateS3BucketRegion = "us-east-2"
	DestinationDatabricksUpdateS3BucketRegionUsWest1      DestinationDatabricksUpdateS3BucketRegion = "us-west-1"
	DestinationDatabricksUpdateS3BucketRegionUsWest2      DestinationDatabricksUpdateS3BucketRegion = "us-west-2"
	DestinationDatabricksUpdateS3BucketRegionAfSouth1     DestinationDatabricksUpdateS3BucketRegion = "af-south-1"
	DestinationDatabricksUpdateS3BucketRegionApEast1      DestinationDatabricksUpdateS3BucketRegion = "ap-east-1"
	DestinationDatabricksUpdateS3BucketRegionApSouth1     DestinationDatabricksUpdateS3BucketRegion = "ap-south-1"
	DestinationDatabricksUpdateS3BucketRegionApNortheast1 DestinationDatabricksUpdateS3BucketRegion = "ap-northeast-1"
	DestinationDatabricksUpdateS3BucketRegionApNortheast2 DestinationDatabricksUpdateS3BucketRegion = "ap-northeast-2"
	DestinationDatabricksUpdateS3BucketRegionApNortheast3 DestinationDatabricksUpdateS3BucketRegion = "ap-northeast-3"
	DestinationDatabricksUpdateS3BucketRegionApSoutheast1 DestinationDatabricksUpdateS3BucketRegion = "ap-southeast-1"
	DestinationDatabricksUpdateS3BucketRegionApSoutheast2 DestinationDatabricksUpdateS3BucketRegion = "ap-southeast-2"
	DestinationDatabricksUpdateS3BucketRegionCaCentral1   DestinationDatabricksUpdateS3BucketRegion = "ca-central-1"
	DestinationDatabricksUpdateS3BucketRegionCnNorth1     DestinationDatabricksUpdateS3BucketRegion = "cn-north-1"
	DestinationDatabricksUpdateS3BucketRegionCnNorthwest1 DestinationDatabricksUpdateS3BucketRegion = "cn-northwest-1"
	DestinationDatabricksUpdateS3BucketRegionEuCentral1   DestinationDatabricksUpdateS3BucketRegion = "eu-central-1"
	DestinationDatabricksUpdateS3BucketRegionEuNorth1     DestinationDatabricksUpdateS3BucketRegion = "eu-north-1"
	DestinationDatabricksUpdateS3BucketRegionEuSouth1     DestinationDatabricksUpdateS3BucketRegion = "eu-south-1"
	DestinationDatabricksUpdateS3BucketRegionEuWest1      DestinationDatabricksUpdateS3BucketRegion = "eu-west-1"
	DestinationDatabricksUpdateS3BucketRegionEuWest2      DestinationDatabricksUpdateS3BucketRegion = "eu-west-2"
	DestinationDatabricksUpdateS3BucketRegionEuWest3      DestinationDatabricksUpdateS3BucketRegion = "eu-west-3"
	DestinationDatabricksUpdateS3BucketRegionSaEast1      DestinationDatabricksUpdateS3BucketRegion = "sa-east-1"
	DestinationDatabricksUpdateS3BucketRegionMeSouth1     DestinationDatabricksUpdateS3BucketRegion = "me-south-1"
	DestinationDatabricksUpdateS3BucketRegionUsGovEast1   DestinationDatabricksUpdateS3BucketRegion = "us-gov-east-1"
	DestinationDatabricksUpdateS3BucketRegionUsGovWest1   DestinationDatabricksUpdateS3BucketRegion = "us-gov-west-1"
)

func (e DestinationDatabricksUpdateS3BucketRegion) ToPointer() *DestinationDatabricksUpdateS3BucketRegion {
	return &e
}

func (e *DestinationDatabricksUpdateS3BucketRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "us-east-1":
		fallthrough
	case "us-east-2":
		fallthrough
	case "us-west-1":
		fallthrough
	case "us-west-2":
		fallthrough
	case "af-south-1":
		fallthrough
	case "ap-east-1":
		fallthrough
	case "ap-south-1":
		fallthrough
	case "ap-northeast-1":
		fallthrough
	case "ap-northeast-2":
		fallthrough
	case "ap-northeast-3":
		fallthrough
	case "ap-southeast-1":
		fallthrough
	case "ap-southeast-2":
		fallthrough
	case "ca-central-1":
		fallthrough
	case "cn-north-1":
		fallthrough
	case "cn-northwest-1":
		fallthrough
	case "eu-central-1":
		fallthrough
	case "eu-north-1":
		fallthrough
	case "eu-south-1":
		fallthrough
	case "eu-west-1":
		fallthrough
	case "eu-west-2":
		fallthrough
	case "eu-west-3":
		fallthrough
	case "sa-east-1":
		fallthrough
	case "me-south-1":
		fallthrough
	case "us-gov-east-1":
		fallthrough
	case "us-gov-west-1":
		*e = DestinationDatabricksUpdateS3BucketRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksUpdateS3BucketRegion: %v", v)
	}
}

// AmazonS3 - Storage on which the delta lake is built.
type AmazonS3 struct {
	dataSourceType DestinationDatabricksUpdateDataSourceType `const:"S3_STORAGE" json:"data_source_type"`
	// The pattern allows you to set the file-name format for the S3 staging file(s)
	FileNamePattern *string `json:"file_name_pattern,omitempty"`
	// The Access Key Id granting allow one to access the above S3 staging bucket. Airbyte requires Read and Write permissions to the given bucket.
	S3AccessKeyID string `json:"s3_access_key_id"`
	// The name of the S3 bucket to use for intermittent staging of the data.
	S3BucketName string `json:"s3_bucket_name"`
	// The directory under the S3 bucket where data will be written.
	S3BucketPath string `json:"s3_bucket_path"`
	// The region of the S3 staging bucket to use if utilising a copy strategy.
	S3BucketRegion *DestinationDatabricksUpdateS3BucketRegion `default:"" json:"s3_bucket_region"`
	// The corresponding secret to the above access key id.
	S3SecretAccessKey string `json:"s3_secret_access_key"`
}

func (a AmazonS3) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AmazonS3) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AmazonS3) GetDataSourceType() DestinationDatabricksUpdateDataSourceType {
	return DestinationDatabricksUpdateDataSourceTypeS3Storage
}

func (o *AmazonS3) GetFileNamePattern() *string {
	if o == nil {
		return nil
	}
	return o.FileNamePattern
}

func (o *AmazonS3) GetS3AccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.S3AccessKeyID
}

func (o *AmazonS3) GetS3BucketName() string {
	if o == nil {
		return ""
	}
	return o.S3BucketName
}

func (o *AmazonS3) GetS3BucketPath() string {
	if o == nil {
		return ""
	}
	return o.S3BucketPath
}

func (o *AmazonS3) GetS3BucketRegion() *DestinationDatabricksUpdateS3BucketRegion {
	if o == nil {
		return nil
	}
	return o.S3BucketRegion
}

func (o *AmazonS3) GetS3SecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.S3SecretAccessKey
}

type DataSourceType string

const (
	DataSourceTypeManagedTablesStorage DataSourceType = "MANAGED_TABLES_STORAGE"
)

func (e DataSourceType) ToPointer() *DataSourceType {
	return &e
}

func (e *DataSourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MANAGED_TABLES_STORAGE":
		*e = DataSourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DataSourceType: %v", v)
	}
}

// RecommendedManagedTables - Storage on which the delta lake is built.
type RecommendedManagedTables struct {
	dataSourceType DataSourceType `const:"MANAGED_TABLES_STORAGE" json:"data_source_type"`
}

func (r RecommendedManagedTables) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RecommendedManagedTables) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *RecommendedManagedTables) GetDataSourceType() DataSourceType {
	return DataSourceTypeManagedTablesStorage
}

type DataSourceUnionType string

const (
	DataSourceUnionTypeRecommendedManagedTables DataSourceUnionType = "RecommendedManagedTables"
	DataSourceUnionTypeAmazonS3                 DataSourceUnionType = "AmazonS3"
	DataSourceUnionTypeAzureBlobStorage         DataSourceUnionType = "AzureBlobStorage"
)

type DataSource struct {
	RecommendedManagedTables *RecommendedManagedTables
	AmazonS3                 *AmazonS3
	AzureBlobStorage         *DestinationDatabricksUpdateAzureBlobStorage

	Type DataSourceUnionType
}

func CreateDataSourceRecommendedManagedTables(recommendedManagedTables RecommendedManagedTables) DataSource {
	typ := DataSourceUnionTypeRecommendedManagedTables

	return DataSource{
		RecommendedManagedTables: &recommendedManagedTables,
		Type:                     typ,
	}
}

func CreateDataSourceAmazonS3(amazonS3 AmazonS3) DataSource {
	typ := DataSourceUnionTypeAmazonS3

	return DataSource{
		AmazonS3: &amazonS3,
		Type:     typ,
	}
}

func CreateDataSourceAzureBlobStorage(azureBlobStorage DestinationDatabricksUpdateAzureBlobStorage) DataSource {
	typ := DataSourceUnionTypeAzureBlobStorage

	return DataSource{
		AzureBlobStorage: &azureBlobStorage,
		Type:             typ,
	}
}

func (u *DataSource) UnmarshalJSON(data []byte) error {

	recommendedManagedTables := new(RecommendedManagedTables)
	if err := utils.UnmarshalJSON(data, &recommendedManagedTables, "", true, true); err == nil {
		u.RecommendedManagedTables = recommendedManagedTables
		u.Type = DataSourceUnionTypeRecommendedManagedTables
		return nil
	}

	azureBlobStorage := new(DestinationDatabricksUpdateAzureBlobStorage)
	if err := utils.UnmarshalJSON(data, &azureBlobStorage, "", true, true); err == nil {
		u.AzureBlobStorage = azureBlobStorage
		u.Type = DataSourceUnionTypeAzureBlobStorage
		return nil
	}

	amazonS3 := new(AmazonS3)
	if err := utils.UnmarshalJSON(data, &amazonS3, "", true, true); err == nil {
		u.AmazonS3 = amazonS3
		u.Type = DataSourceUnionTypeAmazonS3
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DataSource) MarshalJSON() ([]byte, error) {
	if u.RecommendedManagedTables != nil {
		return utils.MarshalJSON(u.RecommendedManagedTables, "", true)
	}

	if u.AmazonS3 != nil {
		return utils.MarshalJSON(u.AmazonS3, "", true)
	}

	if u.AzureBlobStorage != nil {
		return utils.MarshalJSON(u.AzureBlobStorage, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationDatabricksUpdate struct {
	// You must agree to the Databricks JDBC Driver <a href="https://databricks.com/jdbc-odbc-driver-license">Terms & Conditions</a> to use this connector.
	AcceptTerms *bool `default:"false" json:"accept_terms"`
	// Storage on which the delta lake is built.
	DataSource DataSource `json:"data_source"`
	// The name of the catalog. If not specified otherwise, the "hive_metastore" will be used.
	Database *string `json:"database,omitempty"`
	// Databricks Cluster HTTP Path.
	DatabricksHTTPPath string `json:"databricks_http_path"`
	// Databricks Personal Access Token for making authenticated requests.
	DatabricksPersonalAccessToken string `json:"databricks_personal_access_token"`
	// Databricks Cluster Port.
	DatabricksPort *string `default:"443" json:"databricks_port"`
	// Databricks Cluster Server Hostname.
	DatabricksServerHostname string `json:"databricks_server_hostname"`
	// Support schema evolution for all streams. If "false", the connector might fail when a stream's schema changes.
	EnableSchemaEvolution *bool `default:"false" json:"enable_schema_evolution"`
	// Default to 'true'. Switch it to 'false' for debugging purpose.
	PurgeStagingData *bool `default:"true" json:"purge_staging_data"`
	// The default schema tables are written. If not specified otherwise, the "default" will be used.
	Schema *string `default:"default" json:"schema"`
}

func (d DestinationDatabricksUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationDatabricksUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationDatabricksUpdate) GetAcceptTerms() *bool {
	if o == nil {
		return nil
	}
	return o.AcceptTerms
}

func (o *DestinationDatabricksUpdate) GetDataSource() DataSource {
	if o == nil {
		return DataSource{}
	}
	return o.DataSource
}

func (o *DestinationDatabricksUpdate) GetDatabase() *string {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *DestinationDatabricksUpdate) GetDatabricksHTTPPath() string {
	if o == nil {
		return ""
	}
	return o.DatabricksHTTPPath
}

func (o *DestinationDatabricksUpdate) GetDatabricksPersonalAccessToken() string {
	if o == nil {
		return ""
	}
	return o.DatabricksPersonalAccessToken
}

func (o *DestinationDatabricksUpdate) GetDatabricksPort() *string {
	if o == nil {
		return nil
	}
	return o.DatabricksPort
}

func (o *DestinationDatabricksUpdate) GetDatabricksServerHostname() string {
	if o == nil {
		return ""
	}
	return o.DatabricksServerHostname
}

func (o *DestinationDatabricksUpdate) GetEnableSchemaEvolution() *bool {
	if o == nil {
		return nil
	}
	return o.EnableSchemaEvolution
}

func (o *DestinationDatabricksUpdate) GetPurgeStagingData() *bool {
	if o == nil {
		return nil
	}
	return o.PurgeStagingData
}

func (o *DestinationDatabricksUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}
