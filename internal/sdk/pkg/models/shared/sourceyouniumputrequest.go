// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceYouniumPutRequest struct {
	Configuration SourceYouniumUpdate `json:"configuration"`
	Name          string              `json:"name"`
	WorkspaceID   string              `json:"workspaceId"`
}

func (o *SourceYouniumPutRequest) GetConfiguration() SourceYouniumUpdate {
	if o == nil {
		return SourceYouniumUpdate{}
	}
	return o.Configuration
}

func (o *SourceYouniumPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceYouniumPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
