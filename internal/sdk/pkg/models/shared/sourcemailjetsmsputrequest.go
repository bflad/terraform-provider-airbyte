// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceMailjetSmsPutRequest struct {
	Configuration SourceMailjetSmsUpdate `json:"configuration"`
	Name          string                 `json:"name"`
	WorkspaceID   string                 `json:"workspaceId"`
}

func (o *SourceMailjetSmsPutRequest) GetConfiguration() SourceMailjetSmsUpdate {
	if o == nil {
		return SourceMailjetSmsUpdate{}
	}
	return o.Configuration
}

func (o *SourceMailjetSmsPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceMailjetSmsPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
