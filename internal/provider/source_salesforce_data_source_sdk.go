// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceSalesforceDataSourceModel) RefreshFromGetResponse(resp *shared.SourceSalesforceGetResponse) {
	if resp.Configuration.AuthType != nil {
		r.Configuration.AuthType = types.StringValue(string(*resp.Configuration.AuthType))
	} else {
		r.Configuration.AuthType = types.StringNull()
	}
	r.Configuration.ClientID = types.StringValue(resp.Configuration.ClientID)
	r.Configuration.ClientSecret = types.StringValue(resp.Configuration.ClientSecret)
	if resp.Configuration.ForceUseBulkAPI != nil {
		r.Configuration.ForceUseBulkAPI = types.BoolValue(*resp.Configuration.ForceUseBulkAPI)
	} else {
		r.Configuration.ForceUseBulkAPI = types.BoolNull()
	}
	if resp.Configuration.IsSandbox != nil {
		r.Configuration.IsSandbox = types.BoolValue(*resp.Configuration.IsSandbox)
	} else {
		r.Configuration.IsSandbox = types.BoolNull()
	}
	r.Configuration.RefreshToken = types.StringValue(resp.Configuration.RefreshToken)
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	if resp.Configuration.StartDate != nil {
		r.Configuration.StartDate = types.StringValue(resp.Configuration.StartDate.Format(time.RFC3339))
	} else {
		r.Configuration.StartDate = types.StringNull()
	}
	r.Configuration.StreamsCriteria = nil
	for _, streamsCriteriaItem := range resp.Configuration.StreamsCriteria {
		var streamsCriteria1 SourceSalesforceStreamsCriteria
		streamsCriteria1.Criteria = types.StringValue(string(streamsCriteriaItem.Criteria))
		streamsCriteria1.Value = types.StringValue(streamsCriteriaItem.Value)
		r.Configuration.StreamsCriteria = append(r.Configuration.StreamsCriteria, streamsCriteria1)
	}
	r.Name = types.StringValue(resp.Name)
	if resp.SecretID != nil {
		r.SecretID = types.StringValue(*resp.SecretID)
	} else {
		r.SecretID = types.StringNull()
	}
	if resp.SourceID != nil {
		r.SourceID = types.StringValue(*resp.SourceID)
	} else {
		r.SourceID = types.StringNull()
	}
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
