// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// APIExportClient contains the methods for the APIExport group.
// Don't use this type directly, use NewAPIExportClient() instead.
type APIExportClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAPIExportClient creates a new instance of APIExportClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAPIExportClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*APIExportClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &APIExportClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the details of the API specified by its identifier in the format specified to the Storage Blob with SAS Key
// valid for 5 minutes.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - formatParam - Format in which to export the Api Details to the Storage Blob with Sas Key valid for 5 minutes. New formats
//     can be added in the future.
//   - export - Query parameter required to export the API details.
//   - options - APIExportClientGetOptions contains the optional parameters for the APIExportClient.Get method.
func (client *APIExportClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiID string, formatParam ExportFormat, export ExportAPI, options *APIExportClientGetOptions) (APIExportClientGetResponse, error) {
	var err error
	const operationName = "APIExportClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, apiID, formatParam, export, options)
	if err != nil {
		return APIExportClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIExportClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return APIExportClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *APIExportClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, formatParam ExportFormat, export ExportAPI, _ *APIExportClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	reqQP.Set("export", string(export))
	reqQP.Set("format", string(formatParam))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *APIExportClient) getHandleResponse(resp *http.Response) (APIExportClientGetResponse, error) {
	result := APIExportClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIExportResult); err != nil {
		return APIExportClientGetResponse{}, err
	}
	return result, nil
}
