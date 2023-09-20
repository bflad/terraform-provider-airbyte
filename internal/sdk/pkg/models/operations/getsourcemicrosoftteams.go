// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSourceMicrosoftTeamsRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceMicrosoftTeamsResponse struct {
	ContentType string
	// Get a Source by the id in the path.
	SourceMicrosoftTeamsGetResponse *shared.SourceMicrosoftTeamsGetResponse
	StatusCode                      int
	RawResponse                     *http.Response
}
