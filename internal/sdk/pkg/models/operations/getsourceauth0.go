// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSourceAuth0Request struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceAuth0Response struct {
	ContentType string
	// Get a Source by the id in the path.
	SourceAuth0GetResponse *shared.SourceAuth0GetResponse
	StatusCode             int
	RawResponse            *http.Response
}
