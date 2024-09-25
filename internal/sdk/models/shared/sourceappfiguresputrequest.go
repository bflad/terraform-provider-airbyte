// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceAppfiguresPutRequest struct {
	Configuration SourceAppfiguresUpdate `json:"configuration"`
	Name          string                 `json:"name"`
	WorkspaceID   string                 `json:"workspaceId"`
}

func (o *SourceAppfiguresPutRequest) GetConfiguration() SourceAppfiguresUpdate {
	if o == nil {
		return SourceAppfiguresUpdate{}
	}
	return o.Configuration
}

func (o *SourceAppfiguresPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceAppfiguresPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
