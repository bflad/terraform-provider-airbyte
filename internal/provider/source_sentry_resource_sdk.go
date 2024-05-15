// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSentryResourceModel) ToSharedSourceSentryCreateRequest() *shared.SourceSentryCreateRequest {
	authToken := r.Configuration.AuthToken.ValueString()
	var discoverFields []interface{} = []interface{}{}
	for _, discoverFieldsItem := range r.Configuration.DiscoverFields {
		var discoverFieldsTmp interface{}
		_ = json.Unmarshal([]byte(discoverFieldsItem.ValueString()), &discoverFieldsTmp)
		discoverFields = append(discoverFields, discoverFieldsTmp)
	}
	hostname := new(string)
	if !r.Configuration.Hostname.IsUnknown() && !r.Configuration.Hostname.IsNull() {
		*hostname = r.Configuration.Hostname.ValueString()
	} else {
		hostname = nil
	}
	organization := r.Configuration.Organization.ValueString()
	project := r.Configuration.Project.ValueString()
	configuration := shared.SourceSentry{
		AuthToken:      authToken,
		DiscoverFields: discoverFields,
		Hostname:       hostname,
		Organization:   organization,
		Project:        project,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSentryCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSentryResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceSentryResourceModel) ToSharedSourceSentryPutRequest() *shared.SourceSentryPutRequest {
	authToken := r.Configuration.AuthToken.ValueString()
	var discoverFields []interface{} = []interface{}{}
	for _, discoverFieldsItem := range r.Configuration.DiscoverFields {
		var discoverFieldsTmp interface{}
		_ = json.Unmarshal([]byte(discoverFieldsItem.ValueString()), &discoverFieldsTmp)
		discoverFields = append(discoverFields, discoverFieldsTmp)
	}
	hostname := new(string)
	if !r.Configuration.Hostname.IsUnknown() && !r.Configuration.Hostname.IsNull() {
		*hostname = r.Configuration.Hostname.ValueString()
	} else {
		hostname = nil
	}
	organization := r.Configuration.Organization.ValueString()
	project := r.Configuration.Project.ValueString()
	configuration := shared.SourceSentryUpdate{
		AuthToken:      authToken,
		DiscoverFields: discoverFields,
		Hostname:       hostname,
		Organization:   organization,
		Project:        project,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSentryPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
