// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense/v2"
	"net/http"
	"net/url"
	"regexp"
)

// BinaryHardeningServer is a fake server for instances of the armiotfirmwaredefense.BinaryHardeningClient type.
type BinaryHardeningServer struct {
	// NewListByFirmwarePager is the fake for method BinaryHardeningClient.NewListByFirmwarePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByFirmwarePager func(resourceGroupName string, workspaceName string, firmwareID string, options *armiotfirmwaredefense.BinaryHardeningClientListByFirmwareOptions) (resp azfake.PagerResponder[armiotfirmwaredefense.BinaryHardeningClientListByFirmwareResponse])
}

// NewBinaryHardeningServerTransport creates a new instance of BinaryHardeningServerTransport with the provided implementation.
// The returned BinaryHardeningServerTransport instance is connected to an instance of armiotfirmwaredefense.BinaryHardeningClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBinaryHardeningServerTransport(srv *BinaryHardeningServer) *BinaryHardeningServerTransport {
	return &BinaryHardeningServerTransport{
		srv:                    srv,
		newListByFirmwarePager: newTracker[azfake.PagerResponder[armiotfirmwaredefense.BinaryHardeningClientListByFirmwareResponse]](),
	}
}

// BinaryHardeningServerTransport connects instances of armiotfirmwaredefense.BinaryHardeningClient to instances of BinaryHardeningServer.
// Don't use this type directly, use NewBinaryHardeningServerTransport instead.
type BinaryHardeningServerTransport struct {
	srv                    *BinaryHardeningServer
	newListByFirmwarePager *tracker[azfake.PagerResponder[armiotfirmwaredefense.BinaryHardeningClientListByFirmwareResponse]]
}

// Do implements the policy.Transporter interface for BinaryHardeningServerTransport.
func (b *BinaryHardeningServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BinaryHardeningServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if binaryHardeningServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = binaryHardeningServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "BinaryHardeningClient.NewListByFirmwarePager":
				res.resp, res.err = b.dispatchNewListByFirmwarePager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (b *BinaryHardeningServerTransport) dispatchNewListByFirmwarePager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListByFirmwarePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByFirmwarePager not implemented")}
	}
	newListByFirmwarePager := b.newListByFirmwarePager.get(req)
	if newListByFirmwarePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IoTFirmwareDefense/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/firmwares/(?P<firmwareId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/binaryHardeningResults`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		firmwareIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("firmwareId")])
		if err != nil {
			return nil, err
		}
		resp := b.srv.NewListByFirmwarePager(resourceGroupNameParam, workspaceNameParam, firmwareIDParam, nil)
		newListByFirmwarePager = &resp
		b.newListByFirmwarePager.add(req, newListByFirmwarePager)
		server.PagerResponderInjectNextLinks(newListByFirmwarePager, req, func(page *armiotfirmwaredefense.BinaryHardeningClientListByFirmwareResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByFirmwarePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListByFirmwarePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByFirmwarePager) {
		b.newListByFirmwarePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to BinaryHardeningServerTransport
var binaryHardeningServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
