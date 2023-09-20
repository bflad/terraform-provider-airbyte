// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetDestinationPineconeRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
}

type GetDestinationPineconeResponse struct {
	ContentType string
	// Get a Destination by the id in the path.
	DestinationPineconeGetResponse *shared.DestinationPineconeGetResponse
	StatusCode                     int
	RawResponse                    *http.Response
}
