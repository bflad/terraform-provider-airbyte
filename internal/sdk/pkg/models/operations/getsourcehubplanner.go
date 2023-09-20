// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSourceHubplannerRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceHubplannerResponse struct {
	ContentType string
	// Get a Source by the id in the path.
	SourceHubplannerGetResponse *shared.SourceHubplannerGetResponse
	StatusCode                  int
	RawResponse                 *http.Response
}
