// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateDestinationDatabricksResponse struct {
	ContentType string
	// Successful operation
	DestinationDatabricksGetResponse *shared.DestinationDatabricksGetResponse
	StatusCode                       int
	RawResponse                      *http.Response
}
