// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceMyHoursResponse struct {
	ContentType string
	// Successful operation
	SourceMyHoursGetResponse *shared.SourceMyHoursGetResponse
	StatusCode               int
	RawResponse              *http.Response
}
