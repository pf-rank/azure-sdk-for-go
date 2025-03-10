// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdatabasewatcher

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// TargetsClient contains the methods for the Targets group.
// Don't use this type directly, use NewTargetsClient() instead.
type TargetsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTargetsClient creates a new instance of TargetsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTargetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TargetsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TargetsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create a Target
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-02
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - watcherName - The database watcher name.
//   - targetName - The target resource name.
//   - resource - Resource create parameters.
//   - options - TargetsClientCreateOrUpdateOptions contains the optional parameters for the TargetsClient.CreateOrUpdate method.
func (client *TargetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, watcherName string, targetName string, resource Target, options *TargetsClientCreateOrUpdateOptions) (TargetsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "TargetsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, watcherName, targetName, resource, options)
	if err != nil {
		return TargetsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TargetsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return TargetsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TargetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, watcherName string, targetName string, resource Target, _ *TargetsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DatabaseWatcher/watchers/{watcherName}/targets/{targetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	if targetName == "" {
		return nil, errors.New("parameter targetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetName}", url.PathEscape(targetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TargetsClient) createOrUpdateHandleResponse(resp *http.Response) (TargetsClientCreateOrUpdateResponse, error) {
	result := TargetsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Target); err != nil {
		return TargetsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Target
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-02
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - watcherName - The database watcher name.
//   - targetName - The target resource name.
//   - options - TargetsClientDeleteOptions contains the optional parameters for the TargetsClient.Delete method.
func (client *TargetsClient) Delete(ctx context.Context, resourceGroupName string, watcherName string, targetName string, options *TargetsClientDeleteOptions) (TargetsClientDeleteResponse, error) {
	var err error
	const operationName = "TargetsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, watcherName, targetName, options)
	if err != nil {
		return TargetsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TargetsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return TargetsClientDeleteResponse{}, err
	}
	return TargetsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TargetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, watcherName string, targetName string, _ *TargetsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DatabaseWatcher/watchers/{watcherName}/targets/{targetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	if targetName == "" {
		return nil, errors.New("parameter targetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetName}", url.PathEscape(targetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Target
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-02
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - watcherName - The database watcher name.
//   - targetName - The target resource name.
//   - options - TargetsClientGetOptions contains the optional parameters for the TargetsClient.Get method.
func (client *TargetsClient) Get(ctx context.Context, resourceGroupName string, watcherName string, targetName string, options *TargetsClientGetOptions) (TargetsClientGetResponse, error) {
	var err error
	const operationName = "TargetsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, watcherName, targetName, options)
	if err != nil {
		return TargetsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TargetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TargetsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TargetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, watcherName string, targetName string, _ *TargetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DatabaseWatcher/watchers/{watcherName}/targets/{targetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	if targetName == "" {
		return nil, errors.New("parameter targetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetName}", url.PathEscape(targetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TargetsClient) getHandleResponse(resp *http.Response) (TargetsClientGetResponse, error) {
	result := TargetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Target); err != nil {
		return TargetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWatcherPager - List Target resources by Watcher
//
// Generated from API version 2025-01-02
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - watcherName - The database watcher name.
//   - options - TargetsClientListByWatcherOptions contains the optional parameters for the TargetsClient.NewListByWatcherPager
//     method.
func (client *TargetsClient) NewListByWatcherPager(resourceGroupName string, watcherName string, options *TargetsClientListByWatcherOptions) *runtime.Pager[TargetsClientListByWatcherResponse] {
	return runtime.NewPager(runtime.PagingHandler[TargetsClientListByWatcherResponse]{
		More: func(page TargetsClientListByWatcherResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TargetsClientListByWatcherResponse) (TargetsClientListByWatcherResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TargetsClient.NewListByWatcherPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByWatcherCreateRequest(ctx, resourceGroupName, watcherName, options)
			}, nil)
			if err != nil {
				return TargetsClientListByWatcherResponse{}, err
			}
			return client.listByWatcherHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWatcherCreateRequest creates the ListByWatcher request.
func (client *TargetsClient) listByWatcherCreateRequest(ctx context.Context, resourceGroupName string, watcherName string, _ *TargetsClientListByWatcherOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DatabaseWatcher/watchers/{watcherName}/targets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWatcherHandleResponse handles the ListByWatcher response.
func (client *TargetsClient) listByWatcherHandleResponse(resp *http.Response) (TargetsClientListByWatcherResponse, error) {
	result := TargetsClientListByWatcherResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TargetListResult); err != nil {
		return TargetsClientListByWatcherResponse{}, err
	}
	return result, nil
}
