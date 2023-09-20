// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSourceOktaRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceOktaResponse struct {
	ContentType string
	// Get a Source by the id in the path.
	SourceOktaGetResponse *shared.SourceOktaGetResponse
	StatusCode            int
	RawResponse           *http.Response
}
