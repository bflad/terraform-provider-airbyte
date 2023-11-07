// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceMondayPutRequest struct {
	Configuration SourceMondayUpdate `json:"configuration"`
	Name          string             `json:"name"`
	WorkspaceID   string             `json:"workspaceId"`
}

func (o *SourceMondayPutRequest) GetConfiguration() SourceMondayUpdate {
	if o == nil {
		return SourceMondayUpdate{}
	}
	return o.Configuration
}

func (o *SourceMondayPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceMondayPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
