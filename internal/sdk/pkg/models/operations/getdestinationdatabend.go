// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetDestinationDatabendRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
}

type GetDestinationDatabendResponse struct {
	ContentType string
	// Get a Destination by the id in the path.
	DestinationDatabendGetResponse *shared.DestinationDatabendGetResponse
	StatusCode                     int
	RawResponse                    *http.Response
}
