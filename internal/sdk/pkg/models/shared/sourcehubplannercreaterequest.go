// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceHubplannerCreateRequest struct {
	Configuration SourceHubplanner `json:"configuration"`
	Name          string           `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceHubplannerCreateRequest) GetConfiguration() SourceHubplanner {
	if o == nil {
		return SourceHubplanner{}
	}
	return o.Configuration
}

func (o *SourceHubplannerCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceHubplannerCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceHubplannerCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
