// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceRecreationCreateRequest struct {
	Configuration SourceRecreation `json:"configuration"`
	Name          string           `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceRecreationCreateRequest) GetConfiguration() SourceRecreation {
	if o == nil {
		return SourceRecreation{}
	}
	return o.Configuration
}

func (o *SourceRecreationCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceRecreationCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceRecreationCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
