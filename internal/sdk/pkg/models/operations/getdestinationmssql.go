// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetDestinationMssqlRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
}

type GetDestinationMssqlResponse struct {
	ContentType string
	// Get a Destination by the id in the path.
	DestinationMssqlGetResponse *shared.DestinationMssqlGetResponse
	StatusCode                  int
	RawResponse                 *http.Response
}
