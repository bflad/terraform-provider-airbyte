// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateDestinationKeenResponse struct {
	ContentType string
	// Successful operation
	DestinationKeenGetResponse *shared.DestinationKeenGetResponse
	StatusCode                 int
	RawResponse                *http.Response
}
