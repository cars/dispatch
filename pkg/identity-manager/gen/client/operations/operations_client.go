///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Home ans placehold home page
*/
func (a *Client) Home(params *HomeParams, authInfo runtime.ClientAuthInfoWriter) (*HomeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHomeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "home",
		Method:             "GET",
		PathPattern:        "/v1/iam/home",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HomeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HomeOK), nil

}

/*
Redirect redirects to localhost for vs cli login testing
*/
func (a *Client) Redirect(params *RedirectParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRedirectParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "redirect",
		Method:             "GET",
		PathPattern:        "/v1/iam/redirect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RedirectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
Root ans placehold root page no authentication is required at this point
*/
func (a *Client) Root(params *RootParams) (*RootOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRootParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "root",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RootReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RootOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
