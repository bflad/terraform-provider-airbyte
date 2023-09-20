// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceIntercomResponse struct {
	ContentType string
	// Successful operation
	SourceIntercomGetResponse *shared.SourceIntercomGetResponse
	StatusCode                int
	RawResponse               *http.Response
}
