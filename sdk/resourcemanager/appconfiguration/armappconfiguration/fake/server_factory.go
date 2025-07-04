// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armappconfiguration.ClientFactory type.
type ServerFactory struct {
	// ConfigurationStoresServer contains the fakes for client ConfigurationStoresClient
	ConfigurationStoresServer ConfigurationStoresServer

	// KeyValuesServer contains the fakes for client KeyValuesClient
	KeyValuesServer KeyValuesServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// PrivateEndpointConnectionsServer contains the fakes for client PrivateEndpointConnectionsClient
	PrivateEndpointConnectionsServer PrivateEndpointConnectionsServer

	// PrivateLinkResourcesServer contains the fakes for client PrivateLinkResourcesClient
	PrivateLinkResourcesServer PrivateLinkResourcesServer

	// ReplicasServer contains the fakes for client ReplicasClient
	ReplicasServer ReplicasServer

	// SnapshotsServer contains the fakes for client SnapshotsClient
	SnapshotsServer SnapshotsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armappconfiguration.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armappconfiguration.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                *ServerFactory
	trMu                               sync.Mutex
	trConfigurationStoresServer        *ConfigurationStoresServerTransport
	trKeyValuesServer                  *KeyValuesServerTransport
	trOperationsServer                 *OperationsServerTransport
	trPrivateEndpointConnectionsServer *PrivateEndpointConnectionsServerTransport
	trPrivateLinkResourcesServer       *PrivateLinkResourcesServerTransport
	trReplicasServer                   *ReplicasServerTransport
	trSnapshotsServer                  *SnapshotsServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "ConfigurationStoresClient":
		initServer(s, &s.trConfigurationStoresServer, func() *ConfigurationStoresServerTransport {
			return NewConfigurationStoresServerTransport(&s.srv.ConfigurationStoresServer)
		})
		resp, err = s.trConfigurationStoresServer.Do(req)
	case "KeyValuesClient":
		initServer(s, &s.trKeyValuesServer, func() *KeyValuesServerTransport { return NewKeyValuesServerTransport(&s.srv.KeyValuesServer) })
		resp, err = s.trKeyValuesServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PrivateEndpointConnectionsClient":
		initServer(s, &s.trPrivateEndpointConnectionsServer, func() *PrivateEndpointConnectionsServerTransport {
			return NewPrivateEndpointConnectionsServerTransport(&s.srv.PrivateEndpointConnectionsServer)
		})
		resp, err = s.trPrivateEndpointConnectionsServer.Do(req)
	case "PrivateLinkResourcesClient":
		initServer(s, &s.trPrivateLinkResourcesServer, func() *PrivateLinkResourcesServerTransport {
			return NewPrivateLinkResourcesServerTransport(&s.srv.PrivateLinkResourcesServer)
		})
		resp, err = s.trPrivateLinkResourcesServer.Do(req)
	case "ReplicasClient":
		initServer(s, &s.trReplicasServer, func() *ReplicasServerTransport { return NewReplicasServerTransport(&s.srv.ReplicasServer) })
		resp, err = s.trReplicasServer.Do(req)
	case "SnapshotsClient":
		initServer(s, &s.trSnapshotsServer, func() *SnapshotsServerTransport { return NewSnapshotsServerTransport(&s.srv.SnapshotsServer) })
		resp, err = s.trSnapshotsServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
