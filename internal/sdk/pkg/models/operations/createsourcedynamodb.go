// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceDynamodbResponse struct {
	ContentType string
	// Successful operation
	SourceDynamodbGetResponse *shared.SourceDynamodbGetResponse
	StatusCode                int
	RawResponse               *http.Response
}
