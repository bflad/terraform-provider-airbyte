// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationDatabricksResourceModel) ToCreateSDKType() *shared.DestinationDatabricksCreateRequest {
	acceptTerms := new(bool)
	if !r.Configuration.AcceptTerms.IsUnknown() && !r.Configuration.AcceptTerms.IsNull() {
		*acceptTerms = r.Configuration.AcceptTerms.ValueBool()
	} else {
		acceptTerms = nil
	}
	var dataSource shared.DestinationDatabricksDataSource
	var destinationDatabricksRecommendedManagedTables *shared.DestinationDatabricksRecommendedManagedTables
	if r.Configuration.DataSource.RecommendedManagedTables != nil {
		destinationDatabricksRecommendedManagedTables = &shared.DestinationDatabricksRecommendedManagedTables{}
	}
	if destinationDatabricksRecommendedManagedTables != nil {
		dataSource = shared.DestinationDatabricksDataSource{
			RecommendedManagedTables: destinationDatabricksRecommendedManagedTables,
		}
	}
	var destinationDatabricksAmazonS3 *shared.DestinationDatabricksAmazonS3
	if r.Configuration.DataSource.AmazonS3 != nil {
		fileNamePattern := new(string)
		if !r.Configuration.DataSource.AmazonS3.FileNamePattern.IsUnknown() && !r.Configuration.DataSource.AmazonS3.FileNamePattern.IsNull() {
			*fileNamePattern = r.Configuration.DataSource.AmazonS3.FileNamePattern.ValueString()
		} else {
			fileNamePattern = nil
		}
		s3AccessKeyID := r.Configuration.DataSource.AmazonS3.S3AccessKeyID.ValueString()
		s3BucketName := r.Configuration.DataSource.AmazonS3.S3BucketName.ValueString()
		s3BucketPath := r.Configuration.DataSource.AmazonS3.S3BucketPath.ValueString()
		s3BucketRegion := new(shared.DestinationDatabricksS3BucketRegion)
		if !r.Configuration.DataSource.AmazonS3.S3BucketRegion.IsUnknown() && !r.Configuration.DataSource.AmazonS3.S3BucketRegion.IsNull() {
			*s3BucketRegion = shared.DestinationDatabricksS3BucketRegion(r.Configuration.DataSource.AmazonS3.S3BucketRegion.ValueString())
		} else {
			s3BucketRegion = nil
		}
		s3SecretAccessKey := r.Configuration.DataSource.AmazonS3.S3SecretAccessKey.ValueString()
		destinationDatabricksAmazonS3 = &shared.DestinationDatabricksAmazonS3{
			FileNamePattern:   fileNamePattern,
			S3AccessKeyID:     s3AccessKeyID,
			S3BucketName:      s3BucketName,
			S3BucketPath:      s3BucketPath,
			S3BucketRegion:    s3BucketRegion,
			S3SecretAccessKey: s3SecretAccessKey,
		}
	}
	if destinationDatabricksAmazonS3 != nil {
		dataSource = shared.DestinationDatabricksDataSource{
			AmazonS3: destinationDatabricksAmazonS3,
		}
	}
	var destinationDatabricksAzureBlobStorage *shared.DestinationDatabricksAzureBlobStorage
	if r.Configuration.DataSource.AzureBlobStorage != nil {
		azureBlobStorageAccountName := r.Configuration.DataSource.AzureBlobStorage.AzureBlobStorageAccountName.ValueString()
		azureBlobStorageContainerName := r.Configuration.DataSource.AzureBlobStorage.AzureBlobStorageContainerName.ValueString()
		azureBlobStorageEndpointDomainName := new(string)
		if !r.Configuration.DataSource.AzureBlobStorage.AzureBlobStorageEndpointDomainName.IsUnknown() && !r.Configuration.DataSource.AzureBlobStorage.AzureBlobStorageEndpointDomainName.IsNull() {
			*azureBlobStorageEndpointDomainName = r.Configuration.DataSource.AzureBlobStorage.AzureBlobStorageEndpointDomainName.ValueString()
		} else {
			azureBlobStorageEndpointDomainName = nil
		}
		azureBlobStorageSasToken := r.Configuration.DataSource.AzureBlobStorage.AzureBlobStorageSasToken.ValueString()
		destinationDatabricksAzureBlobStorage = &shared.DestinationDatabricksAzureBlobStorage{
			AzureBlobStorageAccountName:        azureBlobStorageAccountName,
			AzureBlobStorageContainerName:      azureBlobStorageContainerName,
			AzureBlobStorageEndpointDomainName: azureBlobStorageEndpointDomainName,
			AzureBlobStorageSasToken:           azureBlobStorageSasToken,
		}
	}
	if destinationDatabricksAzureBlobStorage != nil {
		dataSource = shared.DestinationDatabricksDataSource{
			AzureBlobStorage: destinationDatabricksAzureBlobStorage,
		}
	}
	database := new(string)
	if !r.Configuration.Database.IsUnknown() && !r.Configuration.Database.IsNull() {
		*database = r.Configuration.Database.ValueString()
	} else {
		database = nil
	}
	databricksHTTPPath := r.Configuration.DatabricksHTTPPath.ValueString()
	databricksPersonalAccessToken := r.Configuration.DatabricksPersonalAccessToken.ValueString()
	databricksPort := new(string)
	if !r.Configuration.DatabricksPort.IsUnknown() && !r.Configuration.DatabricksPort.IsNull() {
		*databricksPort = r.Configuration.DatabricksPort.ValueString()
	} else {
		databricksPort = nil
	}
	databricksServerHostname := r.Configuration.DatabricksServerHostname.ValueString()
	enableSchemaEvolution := new(bool)
	if !r.Configuration.EnableSchemaEvolution.IsUnknown() && !r.Configuration.EnableSchemaEvolution.IsNull() {
		*enableSchemaEvolution = r.Configuration.EnableSchemaEvolution.ValueBool()
	} else {
		enableSchemaEvolution = nil
	}
	purgeStagingData := new(bool)
	if !r.Configuration.PurgeStagingData.IsUnknown() && !r.Configuration.PurgeStagingData.IsNull() {
		*purgeStagingData = r.Configuration.PurgeStagingData.ValueBool()
	} else {
		purgeStagingData = nil
	}
	schema := new(string)
	if !r.Configuration.Schema.IsUnknown() && !r.Configuration.Schema.IsNull() {
		*schema = r.Configuration.Schema.ValueString()
	} else {
		schema = nil
	}
	configuration := shared.DestinationDatabricks{
		AcceptTerms:                   acceptTerms,
		DataSource:                    dataSource,
		Database:                      database,
		DatabricksHTTPPath:            databricksHTTPPath,
		DatabricksPersonalAccessToken: databricksPersonalAccessToken,
		DatabricksPort:                databricksPort,
		DatabricksServerHostname:      databricksServerHostname,
		EnableSchemaEvolution:         enableSchemaEvolution,
		PurgeStagingData:              purgeStagingData,
		Schema:                        schema,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDatabricksCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDatabricksResourceModel) ToGetSDKType() *shared.DestinationDatabricksCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationDatabricksResourceModel) ToUpdateSDKType() *shared.DestinationDatabricksPutRequest {
	acceptTerms := new(bool)
	if !r.Configuration.AcceptTerms.IsUnknown() && !r.Configuration.AcceptTerms.IsNull() {
		*acceptTerms = r.Configuration.AcceptTerms.ValueBool()
	} else {
		acceptTerms = nil
	}
	var dataSource shared.DataSource
	var recommendedManagedTables *shared.RecommendedManagedTables
	if r.Configuration.DataSource.RecommendedManagedTables != nil {
		recommendedManagedTables = &shared.RecommendedManagedTables{}
	}
	if recommendedManagedTables != nil {
		dataSource = shared.DataSource{
			RecommendedManagedTables: recommendedManagedTables,
		}
	}
	var amazonS3 *shared.AmazonS3
	if r.Configuration.DataSource.AmazonS3 != nil {
		fileNamePattern := new(string)
		if !r.Configuration.DataSource.AmazonS3.FileNamePattern.IsUnknown() && !r.Configuration.DataSource.AmazonS3.FileNamePattern.IsNull() {
			*fileNamePattern = r.Configuration.DataSource.AmazonS3.FileNamePattern.ValueString()
		} else {
			fileNamePattern = nil
		}
		s3AccessKeyID := r.Configuration.DataSource.AmazonS3.S3AccessKeyID.ValueString()
		s3BucketName := r.Configuration.DataSource.AmazonS3.S3BucketName.ValueString()
		s3BucketPath := r.Configuration.DataSource.AmazonS3.S3BucketPath.ValueString()
		s3BucketRegion := new(shared.DestinationDatabricksUpdateS3BucketRegion)
		if !r.Configuration.DataSource.AmazonS3.S3BucketRegion.IsUnknown() && !r.Configuration.DataSource.AmazonS3.S3BucketRegion.IsNull() {
			*s3BucketRegion = shared.DestinationDatabricksUpdateS3BucketRegion(r.Configuration.DataSource.AmazonS3.S3BucketRegion.ValueString())
		} else {
			s3BucketRegion = nil
		}
		s3SecretAccessKey := r.Configuration.DataSource.AmazonS3.S3SecretAccessKey.ValueString()
		amazonS3 = &shared.AmazonS3{
			FileNamePattern:   fileNamePattern,
			S3AccessKeyID:     s3AccessKeyID,
			S3BucketName:      s3BucketName,
			S3BucketPath:      s3BucketPath,
			S3BucketRegion:    s3BucketRegion,
			S3SecretAccessKey: s3SecretAccessKey,
		}
	}
	if amazonS3 != nil {
		dataSource = shared.DataSource{
			AmazonS3: amazonS3,
		}
	}
	var destinationDatabricksUpdateAzureBlobStorage *shared.DestinationDatabricksUpdateAzureBlobStorage
	if r.Configuration.DataSource.AzureBlobStorage != nil {
		azureBlobStorageAccountName := r.Configuration.DataSource.AzureBlobStorage.AzureBlobStorageAccountName.ValueString()
		azureBlobStorageContainerName := r.Configuration.DataSource.AzureBlobStorage.AzureBlobStorageContainerName.ValueString()
		azureBlobStorageEndpointDomainName := new(string)
		if !r.Configuration.DataSource.AzureBlobStorage.AzureBlobStorageEndpointDomainName.IsUnknown() && !r.Configuration.DataSource.AzureBlobStorage.AzureBlobStorageEndpointDomainName.IsNull() {
			*azureBlobStorageEndpointDomainName = r.Configuration.DataSource.AzureBlobStorage.AzureBlobStorageEndpointDomainName.ValueString()
		} else {
			azureBlobStorageEndpointDomainName = nil
		}
		azureBlobStorageSasToken := r.Configuration.DataSource.AzureBlobStorage.AzureBlobStorageSasToken.ValueString()
		destinationDatabricksUpdateAzureBlobStorage = &shared.DestinationDatabricksUpdateAzureBlobStorage{
			AzureBlobStorageAccountName:        azureBlobStorageAccountName,
			AzureBlobStorageContainerName:      azureBlobStorageContainerName,
			AzureBlobStorageEndpointDomainName: azureBlobStorageEndpointDomainName,
			AzureBlobStorageSasToken:           azureBlobStorageSasToken,
		}
	}
	if destinationDatabricksUpdateAzureBlobStorage != nil {
		dataSource = shared.DataSource{
			AzureBlobStorage: destinationDatabricksUpdateAzureBlobStorage,
		}
	}
	database := new(string)
	if !r.Configuration.Database.IsUnknown() && !r.Configuration.Database.IsNull() {
		*database = r.Configuration.Database.ValueString()
	} else {
		database = nil
	}
	databricksHTTPPath := r.Configuration.DatabricksHTTPPath.ValueString()
	databricksPersonalAccessToken := r.Configuration.DatabricksPersonalAccessToken.ValueString()
	databricksPort := new(string)
	if !r.Configuration.DatabricksPort.IsUnknown() && !r.Configuration.DatabricksPort.IsNull() {
		*databricksPort = r.Configuration.DatabricksPort.ValueString()
	} else {
		databricksPort = nil
	}
	databricksServerHostname := r.Configuration.DatabricksServerHostname.ValueString()
	enableSchemaEvolution := new(bool)
	if !r.Configuration.EnableSchemaEvolution.IsUnknown() && !r.Configuration.EnableSchemaEvolution.IsNull() {
		*enableSchemaEvolution = r.Configuration.EnableSchemaEvolution.ValueBool()
	} else {
		enableSchemaEvolution = nil
	}
	purgeStagingData := new(bool)
	if !r.Configuration.PurgeStagingData.IsUnknown() && !r.Configuration.PurgeStagingData.IsNull() {
		*purgeStagingData = r.Configuration.PurgeStagingData.ValueBool()
	} else {
		purgeStagingData = nil
	}
	schema := new(string)
	if !r.Configuration.Schema.IsUnknown() && !r.Configuration.Schema.IsNull() {
		*schema = r.Configuration.Schema.ValueString()
	} else {
		schema = nil
	}
	configuration := shared.DestinationDatabricksUpdate{
		AcceptTerms:                   acceptTerms,
		DataSource:                    dataSource,
		Database:                      database,
		DatabricksHTTPPath:            databricksHTTPPath,
		DatabricksPersonalAccessToken: databricksPersonalAccessToken,
		DatabricksPort:                databricksPort,
		DatabricksServerHostname:      databricksServerHostname,
		EnableSchemaEvolution:         enableSchemaEvolution,
		PurgeStagingData:              purgeStagingData,
		Schema:                        schema,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDatabricksPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDatabricksResourceModel) ToDeleteSDKType() *shared.DestinationDatabricksCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationDatabricksResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationDatabricksResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
