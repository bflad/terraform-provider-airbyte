// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceIp2whoisResourceModel) ToSharedSourceIp2whoisCreateRequest() *shared.SourceIp2whoisCreateRequest {
	apiKey := new(string)
	if !r.Configuration.APIKey.IsUnknown() && !r.Configuration.APIKey.IsNull() {
		*apiKey = r.Configuration.APIKey.ValueString()
	} else {
		apiKey = nil
	}
	domain := new(string)
	if !r.Configuration.Domain.IsUnknown() && !r.Configuration.Domain.IsNull() {
		*domain = r.Configuration.Domain.ValueString()
	} else {
		domain = nil
	}
	configuration := shared.SourceIp2whois{
		APIKey: apiKey,
		Domain: domain,
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
	out := shared.SourceIp2whoisCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceIp2whoisResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceIp2whoisResourceModel) ToSharedSourceIp2whoisPutRequest() *shared.SourceIp2whoisPutRequest {
	apiKey := new(string)
	if !r.Configuration.APIKey.IsUnknown() && !r.Configuration.APIKey.IsNull() {
		*apiKey = r.Configuration.APIKey.ValueString()
	} else {
		apiKey = nil
	}
	domain := new(string)
	if !r.Configuration.Domain.IsUnknown() && !r.Configuration.Domain.IsNull() {
		*domain = r.Configuration.Domain.ValueString()
	} else {
		domain = nil
	}
	configuration := shared.SourceIp2whoisUpdate{
		APIKey: apiKey,
		Domain: domain,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceIp2whoisPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
