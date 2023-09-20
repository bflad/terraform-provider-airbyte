// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceGoogleWebfontsDataSourceModel) RefreshFromGetResponse(resp *shared.SourceGoogleWebfontsGetResponse) {
	if resp.Configuration.Alt != nil {
		r.Configuration.Alt = types.StringValue(*resp.Configuration.Alt)
	} else {
		r.Configuration.Alt = types.StringNull()
	}
	r.Configuration.APIKey = types.StringValue(resp.Configuration.APIKey)
	if resp.Configuration.PrettyPrint != nil {
		r.Configuration.PrettyPrint = types.StringValue(*resp.Configuration.PrettyPrint)
	} else {
		r.Configuration.PrettyPrint = types.StringNull()
	}
	if resp.Configuration.Sort != nil {
		r.Configuration.Sort = types.StringValue(*resp.Configuration.Sort)
	} else {
		r.Configuration.Sort = types.StringNull()
	}
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
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
