// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceSurveySparrowResponse struct {
	ContentType string
	// Successful operation
	SourceSurveySparrowGetResponse *shared.SourceSurveySparrowGetResponse
	StatusCode                     int
	RawResponse                    *http.Response
}
