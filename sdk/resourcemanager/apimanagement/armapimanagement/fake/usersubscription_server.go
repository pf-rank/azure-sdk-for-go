// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// UserSubscriptionServer is a fake server for instances of the armapimanagement.UserSubscriptionClient type.
type UserSubscriptionServer struct {
	// Get is the fake for method UserSubscriptionClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, userID string, sid string, options *armapimanagement.UserSubscriptionClientGetOptions) (resp azfake.Responder[armapimanagement.UserSubscriptionClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method UserSubscriptionClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, serviceName string, userID string, options *armapimanagement.UserSubscriptionClientListOptions) (resp azfake.PagerResponder[armapimanagement.UserSubscriptionClientListResponse])
}

// NewUserSubscriptionServerTransport creates a new instance of UserSubscriptionServerTransport with the provided implementation.
// The returned UserSubscriptionServerTransport instance is connected to an instance of armapimanagement.UserSubscriptionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewUserSubscriptionServerTransport(srv *UserSubscriptionServer) *UserSubscriptionServerTransport {
	return &UserSubscriptionServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armapimanagement.UserSubscriptionClientListResponse]](),
	}
}

// UserSubscriptionServerTransport connects instances of armapimanagement.UserSubscriptionClient to instances of UserSubscriptionServer.
// Don't use this type directly, use NewUserSubscriptionServerTransport instead.
type UserSubscriptionServerTransport struct {
	srv          *UserSubscriptionServer
	newListPager *tracker[azfake.PagerResponder[armapimanagement.UserSubscriptionClientListResponse]]
}

// Do implements the policy.Transporter interface for UserSubscriptionServerTransport.
func (u *UserSubscriptionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return u.dispatchToMethodFake(req, method)
}

func (u *UserSubscriptionServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if userSubscriptionServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = userSubscriptionServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "UserSubscriptionClient.Get":
				res.resp, res.err = u.dispatchGet(req)
			case "UserSubscriptionClient.NewListPager":
				res.resp, res.err = u.dispatchNewListPager(req)
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

func (u *UserSubscriptionServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if u.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<userId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscriptions/(?P<sid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	userIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("userId")])
	if err != nil {
		return nil, err
	}
	sidParam, err := url.PathUnescape(matches[regex.SubexpIndex("sid")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := u.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, userIDParam, sidParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SubscriptionContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (u *UserSubscriptionServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if u.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := u.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<userId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscriptions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		userIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("userId")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armapimanagement.UserSubscriptionClientListOptions
		if filterParam != nil || topParam != nil || skipParam != nil {
			options = &armapimanagement.UserSubscriptionClientListOptions{
				Filter: filterParam,
				Top:    topParam,
				Skip:   skipParam,
			}
		}
		resp := u.srv.NewListPager(resourceGroupNameParam, serviceNameParam, userIDParam, options)
		newListPager = &resp
		u.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armapimanagement.UserSubscriptionClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		u.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		u.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to UserSubscriptionServerTransport
var userSubscriptionServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
