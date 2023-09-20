// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSourceHarvestRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceHarvestResponse struct {
	ContentType string
	// Get a Source by the id in the path.
	SourceHarvestGetResponse *shared.SourceHarvestGetResponse
	StatusCode               int
	RawResponse              *http.Response
}
