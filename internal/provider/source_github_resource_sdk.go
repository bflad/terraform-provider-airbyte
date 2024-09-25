// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceGithubResourceModel) ToSharedSourceGithubCreateRequest() *shared.SourceGithubCreateRequest {
	apiURL := new(string)
	if !r.Configuration.APIURL.IsUnknown() && !r.Configuration.APIURL.IsNull() {
		*apiURL = r.Configuration.APIURL.ValueString()
	} else {
		apiURL = nil
	}
	var branches []string = []string{}
	for _, branchesItem := range r.Configuration.Branches {
		branches = append(branches, branchesItem.ValueString())
	}
	var credentials shared.SourceGithubAuthentication
	var sourceGithubOAuth *shared.SourceGithubOAuth
	if r.Configuration.Credentials.OAuth != nil {
		accessToken := r.Configuration.Credentials.OAuth.AccessToken.ValueString()
		clientID := new(string)
		if !r.Configuration.Credentials.OAuth.ClientID.IsUnknown() && !r.Configuration.Credentials.OAuth.ClientID.IsNull() {
			*clientID = r.Configuration.Credentials.OAuth.ClientID.ValueString()
		} else {
			clientID = nil
		}
		clientSecret := new(string)
		if !r.Configuration.Credentials.OAuth.ClientSecret.IsUnknown() && !r.Configuration.Credentials.OAuth.ClientSecret.IsNull() {
			*clientSecret = r.Configuration.Credentials.OAuth.ClientSecret.ValueString()
		} else {
			clientSecret = nil
		}
		sourceGithubOAuth = &shared.SourceGithubOAuth{
			AccessToken:  accessToken,
			ClientID:     clientID,
			ClientSecret: clientSecret,
		}
	}
	if sourceGithubOAuth != nil {
		credentials = shared.SourceGithubAuthentication{
			SourceGithubOAuth: sourceGithubOAuth,
		}
	}
	var sourceGithubPersonalAccessToken *shared.SourceGithubPersonalAccessToken
	if r.Configuration.Credentials.PersonalAccessToken != nil {
		personalAccessToken := r.Configuration.Credentials.PersonalAccessToken.PersonalAccessToken.ValueString()
		sourceGithubPersonalAccessToken = &shared.SourceGithubPersonalAccessToken{
			PersonalAccessToken: personalAccessToken,
		}
	}
	if sourceGithubPersonalAccessToken != nil {
		credentials = shared.SourceGithubAuthentication{
			SourceGithubPersonalAccessToken: sourceGithubPersonalAccessToken,
		}
	}
	maxWaitingTime := new(int64)
	if !r.Configuration.MaxWaitingTime.IsUnknown() && !r.Configuration.MaxWaitingTime.IsNull() {
		*maxWaitingTime = r.Configuration.MaxWaitingTime.ValueInt64()
	} else {
		maxWaitingTime = nil
	}
	var repositories []string = []string{}
	for _, repositoriesItem := range r.Configuration.Repositories {
		repositories = append(repositories, repositoriesItem.ValueString())
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceGithub{
		APIURL:         apiURL,
		Branches:       branches,
		Credentials:    credentials,
		MaxWaitingTime: maxWaitingTime,
		Repositories:   repositories,
		StartDate:      startDate,
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
	out := shared.SourceGithubCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGithubResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceGithubResourceModel) ToSharedSourceGithubPutRequest() *shared.SourceGithubPutRequest {
	apiURL := new(string)
	if !r.Configuration.APIURL.IsUnknown() && !r.Configuration.APIURL.IsNull() {
		*apiURL = r.Configuration.APIURL.ValueString()
	} else {
		apiURL = nil
	}
	var branches []string = []string{}
	for _, branchesItem := range r.Configuration.Branches {
		branches = append(branches, branchesItem.ValueString())
	}
	var credentials shared.SourceGithubUpdateAuthentication
	var oAuth *shared.OAuth
	if r.Configuration.Credentials.OAuth != nil {
		accessToken := r.Configuration.Credentials.OAuth.AccessToken.ValueString()
		clientID := new(string)
		if !r.Configuration.Credentials.OAuth.ClientID.IsUnknown() && !r.Configuration.Credentials.OAuth.ClientID.IsNull() {
			*clientID = r.Configuration.Credentials.OAuth.ClientID.ValueString()
		} else {
			clientID = nil
		}
		clientSecret := new(string)
		if !r.Configuration.Credentials.OAuth.ClientSecret.IsUnknown() && !r.Configuration.Credentials.OAuth.ClientSecret.IsNull() {
			*clientSecret = r.Configuration.Credentials.OAuth.ClientSecret.ValueString()
		} else {
			clientSecret = nil
		}
		oAuth = &shared.OAuth{
			AccessToken:  accessToken,
			ClientID:     clientID,
			ClientSecret: clientSecret,
		}
	}
	if oAuth != nil {
		credentials = shared.SourceGithubUpdateAuthentication{
			OAuth: oAuth,
		}
	}
	var sourceGithubUpdatePersonalAccessToken *shared.SourceGithubUpdatePersonalAccessToken
	if r.Configuration.Credentials.PersonalAccessToken != nil {
		personalAccessToken := r.Configuration.Credentials.PersonalAccessToken.PersonalAccessToken.ValueString()
		sourceGithubUpdatePersonalAccessToken = &shared.SourceGithubUpdatePersonalAccessToken{
			PersonalAccessToken: personalAccessToken,
		}
	}
	if sourceGithubUpdatePersonalAccessToken != nil {
		credentials = shared.SourceGithubUpdateAuthentication{
			SourceGithubUpdatePersonalAccessToken: sourceGithubUpdatePersonalAccessToken,
		}
	}
	maxWaitingTime := new(int64)
	if !r.Configuration.MaxWaitingTime.IsUnknown() && !r.Configuration.MaxWaitingTime.IsNull() {
		*maxWaitingTime = r.Configuration.MaxWaitingTime.ValueInt64()
	} else {
		maxWaitingTime = nil
	}
	var repositories []string = []string{}
	for _, repositoriesItem := range r.Configuration.Repositories {
		repositories = append(repositories, repositoriesItem.ValueString())
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceGithubUpdate{
		APIURL:         apiURL,
		Branches:       branches,
		Credentials:    credentials,
		MaxWaitingTime: maxWaitingTime,
		Repositories:   repositories,
		StartDate:      startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGithubPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
