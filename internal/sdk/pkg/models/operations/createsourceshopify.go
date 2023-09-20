// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceShopifyResponse struct {
	ContentType string
	// Successful operation
	SourceShopifyGetResponse *shared.SourceShopifyGetResponse
	StatusCode               int
	RawResponse              *http.Response
}
