// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceMicrosoftOnedriveCreateRequest struct {
	// SourceMicrosoftOneDriveSpec class for Microsoft OneDrive Source Specification.
	// This class combines the authentication details with additional configuration for the OneDrive API.
	Configuration SourceMicrosoftOnedrive `json:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the source e.g. dev-mysql-instance.
	Name string `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceMicrosoftOnedriveCreateRequest) GetConfiguration() SourceMicrosoftOnedrive {
	if o == nil {
		return SourceMicrosoftOnedrive{}
	}
	return o.Configuration
}

func (o *SourceMicrosoftOnedriveCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *SourceMicrosoftOnedriveCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceMicrosoftOnedriveCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceMicrosoftOnedriveCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
