// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSourceBraintreeRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceBraintreeResponse struct {
	ContentType string
	// Get a Source by the id in the path.
	SourceBraintreeGetResponse *shared.SourceBraintreeGetResponse
	StatusCode                 int
	RawResponse                *http.Response
}
