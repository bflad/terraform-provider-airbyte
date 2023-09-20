// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceRedshiftResponse struct {
	ContentType string
	// Successful operation
	SourceRedshiftGetResponse *shared.SourceRedshiftGetResponse
	StatusCode                int
	RawResponse               *http.Response
}
