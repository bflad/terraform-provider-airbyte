// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceConvexResponse struct {
	ContentType string
	// Successful operation
	SourceConvexGetResponse *shared.SourceConvexGetResponse
	StatusCode              int
	RawResponse             *http.Response
}
