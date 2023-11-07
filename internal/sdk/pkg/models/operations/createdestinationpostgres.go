// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateDestinationPostgresResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful operation
	DestinationResponse *shared.DestinationResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateDestinationPostgresResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateDestinationPostgresResponse) GetDestinationResponse() *shared.DestinationResponse {
	if o == nil {
		return nil
	}
	return o.DestinationResponse
}

func (o *CreateDestinationPostgresResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateDestinationPostgresResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
