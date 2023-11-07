// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceIntercomCreateRequest struct {
	Configuration SourceIntercom `json:"configuration"`
	Name          string         `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceIntercomCreateRequest) GetConfiguration() SourceIntercom {
	if o == nil {
		return SourceIntercom{}
	}
	return o.Configuration
}

func (o *SourceIntercomCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceIntercomCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceIntercomCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
