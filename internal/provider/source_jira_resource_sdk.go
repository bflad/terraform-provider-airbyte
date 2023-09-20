// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceJiraResourceModel) ToCreateSDKType() *shared.SourceJiraCreateRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	domain := r.Configuration.Domain.ValueString()
	email := r.Configuration.Email.ValueString()
	enableExperimentalStreams := new(bool)
	if !r.Configuration.EnableExperimentalStreams.IsUnknown() && !r.Configuration.EnableExperimentalStreams.IsNull() {
		*enableExperimentalStreams = r.Configuration.EnableExperimentalStreams.ValueBool()
	} else {
		enableExperimentalStreams = nil
	}
	expandIssueChangelog := new(bool)
	if !r.Configuration.ExpandIssueChangelog.IsUnknown() && !r.Configuration.ExpandIssueChangelog.IsNull() {
		*expandIssueChangelog = r.Configuration.ExpandIssueChangelog.ValueBool()
	} else {
		expandIssueChangelog = nil
	}
	var projects []string = nil
	for _, projectsItem := range r.Configuration.Projects {
		projects = append(projects, projectsItem.ValueString())
	}
	renderFields := new(bool)
	if !r.Configuration.RenderFields.IsUnknown() && !r.Configuration.RenderFields.IsNull() {
		*renderFields = r.Configuration.RenderFields.ValueBool()
	} else {
		renderFields = nil
	}
	sourceType := shared.SourceJiraJira(r.Configuration.SourceType.ValueString())
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceJira{
		APIToken:                  apiToken,
		Domain:                    domain,
		Email:                     email,
		EnableExperimentalStreams: enableExperimentalStreams,
		ExpandIssueChangelog:      expandIssueChangelog,
		Projects:                  projects,
		RenderFields:              renderFields,
		SourceType:                sourceType,
		StartDate:                 startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceJiraCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceJiraResourceModel) ToGetSDKType() *shared.SourceJiraCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceJiraResourceModel) ToUpdateSDKType() *shared.SourceJiraPutRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	domain := r.Configuration.Domain.ValueString()
	email := r.Configuration.Email.ValueString()
	enableExperimentalStreams := new(bool)
	if !r.Configuration.EnableExperimentalStreams.IsUnknown() && !r.Configuration.EnableExperimentalStreams.IsNull() {
		*enableExperimentalStreams = r.Configuration.EnableExperimentalStreams.ValueBool()
	} else {
		enableExperimentalStreams = nil
	}
	expandIssueChangelog := new(bool)
	if !r.Configuration.ExpandIssueChangelog.IsUnknown() && !r.Configuration.ExpandIssueChangelog.IsNull() {
		*expandIssueChangelog = r.Configuration.ExpandIssueChangelog.ValueBool()
	} else {
		expandIssueChangelog = nil
	}
	var projects []string = nil
	for _, projectsItem := range r.Configuration.Projects {
		projects = append(projects, projectsItem.ValueString())
	}
	renderFields := new(bool)
	if !r.Configuration.RenderFields.IsUnknown() && !r.Configuration.RenderFields.IsNull() {
		*renderFields = r.Configuration.RenderFields.ValueBool()
	} else {
		renderFields = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceJiraUpdate{
		APIToken:                  apiToken,
		Domain:                    domain,
		Email:                     email,
		EnableExperimentalStreams: enableExperimentalStreams,
		ExpandIssueChangelog:      expandIssueChangelog,
		Projects:                  projects,
		RenderFields:              renderFields,
		StartDate:                 startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceJiraPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceJiraResourceModel) ToDeleteSDKType() *shared.SourceJiraCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceJiraResourceModel) RefreshFromGetResponse(resp *shared.SourceJiraGetResponse) {
	r.Configuration.APIToken = types.StringValue(resp.Configuration.APIToken)
	r.Configuration.Domain = types.StringValue(resp.Configuration.Domain)
	r.Configuration.Email = types.StringValue(resp.Configuration.Email)
	if resp.Configuration.EnableExperimentalStreams != nil {
		r.Configuration.EnableExperimentalStreams = types.BoolValue(*resp.Configuration.EnableExperimentalStreams)
	} else {
		r.Configuration.EnableExperimentalStreams = types.BoolNull()
	}
	if resp.Configuration.ExpandIssueChangelog != nil {
		r.Configuration.ExpandIssueChangelog = types.BoolValue(*resp.Configuration.ExpandIssueChangelog)
	} else {
		r.Configuration.ExpandIssueChangelog = types.BoolNull()
	}
	r.Configuration.Projects = nil
	for _, v := range resp.Configuration.Projects {
		r.Configuration.Projects = append(r.Configuration.Projects, types.StringValue(v))
	}
	if resp.Configuration.RenderFields != nil {
		r.Configuration.RenderFields = types.BoolValue(*resp.Configuration.RenderFields)
	} else {
		r.Configuration.RenderFields = types.BoolNull()
	}
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	if resp.Configuration.StartDate != nil {
		r.Configuration.StartDate = types.StringValue(resp.Configuration.StartDate.Format(time.RFC3339))
	} else {
		r.Configuration.StartDate = types.StringNull()
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

func (r *SourceJiraResourceModel) RefreshFromCreateResponse(resp *shared.SourceJiraGetResponse) {
	r.RefreshFromGetResponse(resp)
}
