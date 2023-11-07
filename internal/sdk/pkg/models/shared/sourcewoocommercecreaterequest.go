// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceWoocommerceCreateRequest struct {
	Configuration SourceWoocommerce `json:"configuration"`
	Name          string            `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceWoocommerceCreateRequest) GetConfiguration() SourceWoocommerce {
	if o == nil {
		return SourceWoocommerce{}
	}
	return o.Configuration
}

func (o *SourceWoocommerceCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceWoocommerceCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceWoocommerceCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
