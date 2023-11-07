// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceNytimesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful operation
	SourceResponse *shared.SourceResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateSourceNytimesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSourceNytimesResponse) GetSourceResponse() *shared.SourceResponse {
	if o == nil {
		return nil
	}
	return o.SourceResponse
}

func (o *CreateSourceNytimesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSourceNytimesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
