// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceAmazonSqsResourceModel) ToCreateSDKType() *shared.SourceAmazonSqsCreateRequest {
	accessKey := new(string)
	if !r.Configuration.AccessKey.IsUnknown() && !r.Configuration.AccessKey.IsNull() {
		*accessKey = r.Configuration.AccessKey.ValueString()
	} else {
		accessKey = nil
	}
	attributesToReturn := new(string)
	if !r.Configuration.AttributesToReturn.IsUnknown() && !r.Configuration.AttributesToReturn.IsNull() {
		*attributesToReturn = r.Configuration.AttributesToReturn.ValueString()
	} else {
		attributesToReturn = nil
	}
	deleteMessages := new(bool)
	if !r.Configuration.DeleteMessages.IsUnknown() && !r.Configuration.DeleteMessages.IsNull() {
		*deleteMessages = r.Configuration.DeleteMessages.ValueBool()
	} else {
		deleteMessages = nil
	}
	maxBatchSize := new(int64)
	if !r.Configuration.MaxBatchSize.IsUnknown() && !r.Configuration.MaxBatchSize.IsNull() {
		*maxBatchSize = r.Configuration.MaxBatchSize.ValueInt64()
	} else {
		maxBatchSize = nil
	}
	maxWaitTime := new(int64)
	if !r.Configuration.MaxWaitTime.IsUnknown() && !r.Configuration.MaxWaitTime.IsNull() {
		*maxWaitTime = r.Configuration.MaxWaitTime.ValueInt64()
	} else {
		maxWaitTime = nil
	}
	queueURL := r.Configuration.QueueURL.ValueString()
	region := shared.SourceAmazonSqsAWSRegion(r.Configuration.Region.ValueString())
	secretKey := new(string)
	if !r.Configuration.SecretKey.IsUnknown() && !r.Configuration.SecretKey.IsNull() {
		*secretKey = r.Configuration.SecretKey.ValueString()
	} else {
		secretKey = nil
	}
	visibilityTimeout := new(int64)
	if !r.Configuration.VisibilityTimeout.IsUnknown() && !r.Configuration.VisibilityTimeout.IsNull() {
		*visibilityTimeout = r.Configuration.VisibilityTimeout.ValueInt64()
	} else {
		visibilityTimeout = nil
	}
	configuration := shared.SourceAmazonSqs{
		AccessKey:          accessKey,
		AttributesToReturn: attributesToReturn,
		DeleteMessages:     deleteMessages,
		MaxBatchSize:       maxBatchSize,
		MaxWaitTime:        maxWaitTime,
		QueueURL:           queueURL,
		Region:             region,
		SecretKey:          secretKey,
		VisibilityTimeout:  visibilityTimeout,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAmazonSqsCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAmazonSqsResourceModel) ToGetSDKType() *shared.SourceAmazonSqsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAmazonSqsResourceModel) ToUpdateSDKType() *shared.SourceAmazonSqsPutRequest {
	accessKey := new(string)
	if !r.Configuration.AccessKey.IsUnknown() && !r.Configuration.AccessKey.IsNull() {
		*accessKey = r.Configuration.AccessKey.ValueString()
	} else {
		accessKey = nil
	}
	attributesToReturn := new(string)
	if !r.Configuration.AttributesToReturn.IsUnknown() && !r.Configuration.AttributesToReturn.IsNull() {
		*attributesToReturn = r.Configuration.AttributesToReturn.ValueString()
	} else {
		attributesToReturn = nil
	}
	deleteMessages := new(bool)
	if !r.Configuration.DeleteMessages.IsUnknown() && !r.Configuration.DeleteMessages.IsNull() {
		*deleteMessages = r.Configuration.DeleteMessages.ValueBool()
	} else {
		deleteMessages = nil
	}
	maxBatchSize := new(int64)
	if !r.Configuration.MaxBatchSize.IsUnknown() && !r.Configuration.MaxBatchSize.IsNull() {
		*maxBatchSize = r.Configuration.MaxBatchSize.ValueInt64()
	} else {
		maxBatchSize = nil
	}
	maxWaitTime := new(int64)
	if !r.Configuration.MaxWaitTime.IsUnknown() && !r.Configuration.MaxWaitTime.IsNull() {
		*maxWaitTime = r.Configuration.MaxWaitTime.ValueInt64()
	} else {
		maxWaitTime = nil
	}
	queueURL := r.Configuration.QueueURL.ValueString()
	region := shared.SourceAmazonSqsUpdateAWSRegion(r.Configuration.Region.ValueString())
	secretKey := new(string)
	if !r.Configuration.SecretKey.IsUnknown() && !r.Configuration.SecretKey.IsNull() {
		*secretKey = r.Configuration.SecretKey.ValueString()
	} else {
		secretKey = nil
	}
	visibilityTimeout := new(int64)
	if !r.Configuration.VisibilityTimeout.IsUnknown() && !r.Configuration.VisibilityTimeout.IsNull() {
		*visibilityTimeout = r.Configuration.VisibilityTimeout.ValueInt64()
	} else {
		visibilityTimeout = nil
	}
	configuration := shared.SourceAmazonSqsUpdate{
		AccessKey:          accessKey,
		AttributesToReturn: attributesToReturn,
		DeleteMessages:     deleteMessages,
		MaxBatchSize:       maxBatchSize,
		MaxWaitTime:        maxWaitTime,
		QueueURL:           queueURL,
		Region:             region,
		SecretKey:          secretKey,
		VisibilityTimeout:  visibilityTimeout,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAmazonSqsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAmazonSqsResourceModel) ToDeleteSDKType() *shared.SourceAmazonSqsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAmazonSqsResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceAmazonSqsResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
