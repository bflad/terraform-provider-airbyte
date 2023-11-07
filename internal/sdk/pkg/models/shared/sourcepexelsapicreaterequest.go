// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourcePexelsAPICreateRequest struct {
	Configuration SourcePexelsAPI `json:"configuration"`
	Name          string          `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourcePexelsAPICreateRequest) GetConfiguration() SourcePexelsAPI {
	if o == nil {
		return SourcePexelsAPI{}
	}
	return o.Configuration
}

func (o *SourcePexelsAPICreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourcePexelsAPICreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourcePexelsAPICreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
