// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/hooks"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/errors"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/operations"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"io"
	"net/http"
	"net/url"
)

type Workspaces struct {
	sdkConfiguration sdkConfiguration
}

func newWorkspaces(sdkConfig sdkConfiguration) *Workspaces {
	return &Workspaces{
		sdkConfiguration: sdkConfig,
	}
}

// CreateOrUpdateWorkspaceOAuthCredentials - Create OAuth override credentials for a workspace and source type.
// Create/update a set of OAuth credentials to override the Airbyte-provided OAuth credentials used for source/destination OAuth.
// In order to determine what the credential configuration needs to be, please see the connector specification of the relevant  source/destination.
func (s *Workspaces) CreateOrUpdateWorkspaceOAuthCredentials(ctx context.Context, request operations.CreateOrUpdateWorkspaceOAuthCredentialsRequest) (*operations.CreateOrUpdateWorkspaceOAuthCredentialsResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "createOrUpdateWorkspaceOAuthCredentials",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/workspaces/{workspaceId}/oauthCredentials", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "WorkspaceOAuthCredentialsRequest", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{}, httpRes.StatusCode) {
		_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		} else if _httpRes != nil {
			httpRes = _httpRes
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.CreateOrUpdateWorkspaceOAuthCredentialsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		fallthrough
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 403:
	default:
		return nil, errors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// CreateWorkspace - Create a workspace
func (s *Workspaces) CreateWorkspace(ctx context.Context, request shared.WorkspaceCreateRequest) (*operations.CreateWorkspaceResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "createWorkspace",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := url.JoinPath(baseURL, "/workspaces")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "Request", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{}, httpRes.StatusCode) {
		_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		} else if _httpRes != nil {
			httpRes = _httpRes
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.CreateWorkspaceResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out shared.WorkspaceResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.WorkspaceResponse = &out
		default:
			return nil, errors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 403:
	default:
		return nil, errors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// DeleteWorkspace - Delete a Workspace
func (s *Workspaces) DeleteWorkspace(ctx context.Context, request operations.DeleteWorkspaceRequest) (*operations.DeleteWorkspaceResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "deleteWorkspace",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/workspaces/{workspaceId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{}, httpRes.StatusCode) {
		_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		} else if _httpRes != nil {
			httpRes = _httpRes
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.DeleteWorkspaceResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode >= 200 && httpRes.StatusCode < 300:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
	default:
		return nil, errors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// GetWorkspace - Get Workspace details
func (s *Workspaces) GetWorkspace(ctx context.Context, request operations.GetWorkspaceRequest) (*operations.GetWorkspaceResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "getWorkspace",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/workspaces/{workspaceId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{}, httpRes.StatusCode) {
		_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		} else if _httpRes != nil {
			httpRes = _httpRes
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.GetWorkspaceResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out shared.WorkspaceResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.WorkspaceResponse = &out
		default:
			return nil, errors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
	default:
		return nil, errors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// ListWorkspaces - List workspaces
func (s *Workspaces) ListWorkspaces(ctx context.Context, request operations.ListWorkspacesRequest) (*operations.ListWorkspacesResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "listWorkspaces",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := url.JoinPath(baseURL, "/workspaces")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{}, httpRes.StatusCode) {
		_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		} else if _httpRes != nil {
			httpRes = _httpRes
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.ListWorkspacesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out shared.WorkspacesResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.WorkspacesResponse = &out
		default:
			return nil, errors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
	default:
		return nil, errors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// UpdateWorkspace - Update a workspace
func (s *Workspaces) UpdateWorkspace(ctx context.Context, request operations.UpdateWorkspaceRequest) (*operations.UpdateWorkspaceResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "updateWorkspace",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/workspaces/{workspaceId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "WorkspaceUpdateRequest", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{}, httpRes.StatusCode) {
		_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		} else if _httpRes != nil {
			httpRes = _httpRes
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.UpdateWorkspaceResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out shared.WorkspaceResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.WorkspaceResponse = &out
		default:
			return nil, errors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 403:
	default:
		return nil, errors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}
