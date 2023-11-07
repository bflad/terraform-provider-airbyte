// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceMssqlPutRequest struct {
	Configuration SourceMssqlUpdate `json:"configuration"`
	Name          string            `json:"name"`
	WorkspaceID   string            `json:"workspaceId"`
}

func (o *SourceMssqlPutRequest) GetConfiguration() SourceMssqlUpdate {
	if o == nil {
		return SourceMssqlUpdate{}
	}
	return o.Configuration
}

func (o *SourceMssqlPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceMssqlPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
