// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration

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

// OperationsClient contains the methods for the Operations group.
// Don't use this type directly, use NewOperationsClient() instead.
type OperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOperationsClient creates a new instance of OperationsClient with the specified values.
//   - subscriptionID - The Microsoft Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - Checks whether the configuration store name is available for use.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01
//   - checkNameAvailabilityParameters - The object containing information for the availability request.
//   - options - OperationsClientCheckNameAvailabilityOptions contains the optional parameters for the OperationsClient.CheckNameAvailability
//     method.
func (client *OperationsClient) CheckNameAvailability(ctx context.Context, checkNameAvailabilityParameters CheckNameAvailabilityParameters, options *OperationsClientCheckNameAvailabilityOptions) (OperationsClientCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "OperationsClient.CheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkNameAvailabilityCreateRequest(ctx, checkNameAvailabilityParameters, options)
	if err != nil {
		return OperationsClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *OperationsClient) checkNameAvailabilityCreateRequest(ctx context.Context, checkNameAvailabilityParameters CheckNameAvailabilityParameters, _ *OperationsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AppConfiguration/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, checkNameAvailabilityParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *OperationsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (OperationsClientCheckNameAvailabilityResponse, error) {
	result := OperationsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailabilityStatus); err != nil {
		return OperationsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the operations available from this provider.
//
// Generated from API version 2024-06-01
//   - options - OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
func (client *OperationsClient) NewListPager(options *OperationsClientListOptions) *runtime.Pager[OperationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OperationsClientListResponse]{
		More: func(page OperationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OperationsClientListResponse) (OperationsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "OperationsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return OperationsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *OperationsClient) listCreateRequest(ctx context.Context, options *OperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppConfiguration/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2024-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OperationsClient) listHandleResponse(resp *http.Response) (OperationsClientListResponse, error) {
	result := OperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationDefinitionListResult); err != nil {
		return OperationsClientListResponse{}, err
	}
	return result, nil
}

// RegionalCheckNameAvailability - Checks whether the configuration store name is available for use.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01
//   - location - The location in which uniqueness will be verified.
//   - checkNameAvailabilityParameters - The object containing information for the availability request.
//   - options - OperationsClientRegionalCheckNameAvailabilityOptions contains the optional parameters for the OperationsClient.RegionalCheckNameAvailability
//     method.
func (client *OperationsClient) RegionalCheckNameAvailability(ctx context.Context, location string, checkNameAvailabilityParameters CheckNameAvailabilityParameters, options *OperationsClientRegionalCheckNameAvailabilityOptions) (OperationsClientRegionalCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "OperationsClient.RegionalCheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.regionalCheckNameAvailabilityCreateRequest(ctx, location, checkNameAvailabilityParameters, options)
	if err != nil {
		return OperationsClientRegionalCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientRegionalCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientRegionalCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.regionalCheckNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// regionalCheckNameAvailabilityCreateRequest creates the RegionalCheckNameAvailability request.
func (client *OperationsClient) regionalCheckNameAvailabilityCreateRequest(ctx context.Context, location string, checkNameAvailabilityParameters CheckNameAvailabilityParameters, _ *OperationsClientRegionalCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AppConfiguration/locations/{location}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, checkNameAvailabilityParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// regionalCheckNameAvailabilityHandleResponse handles the RegionalCheckNameAvailability response.
func (client *OperationsClient) regionalCheckNameAvailabilityHandleResponse(resp *http.Response) (OperationsClientRegionalCheckNameAvailabilityResponse, error) {
	result := OperationsClientRegionalCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailabilityStatus); err != nil {
		return OperationsClientRegionalCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}
