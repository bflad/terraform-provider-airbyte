// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceStravaResponse struct {
	ContentType string
	// Successful operation
	SourceStravaGetResponse *shared.SourceStravaGetResponse
	StatusCode              int
	RawResponse             *http.Response
}
