// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceOnesignalDataSourceModel) RefreshFromGetResponse(resp *shared.SourceOnesignalGetResponse) {
	r.Configuration.Applications = nil
	for _, applicationsItem := range resp.Configuration.Applications {
		var applications1 SourceOnesignalApplications
		applications1.AppAPIKey = types.StringValue(applicationsItem.AppAPIKey)
		applications1.AppID = types.StringValue(applicationsItem.AppID)
		if applicationsItem.AppName != nil {
			applications1.AppName = types.StringValue(*applicationsItem.AppName)
		} else {
			applications1.AppName = types.StringNull()
		}
		r.Configuration.Applications = append(r.Configuration.Applications, applications1)
	}
	r.Configuration.OutcomeNames = types.StringValue(resp.Configuration.OutcomeNames)
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	r.Configuration.StartDate = types.StringValue(resp.Configuration.StartDate.Format(time.RFC3339))
	r.Configuration.UserAuthKey = types.StringValue(resp.Configuration.UserAuthKey)
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
