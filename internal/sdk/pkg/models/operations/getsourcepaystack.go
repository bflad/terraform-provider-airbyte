// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSourcePaystackRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourcePaystackResponse struct {
	ContentType string
	// Get a Source by the id in the path.
	SourcePaystackGetResponse *shared.SourcePaystackGetResponse
	StatusCode                int
	RawResponse               *http.Response
}
