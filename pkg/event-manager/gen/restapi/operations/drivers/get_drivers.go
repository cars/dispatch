///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDriversHandlerFunc turns a function with the right signature into a get drivers handler
type GetDriversHandlerFunc func(GetDriversParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDriversHandlerFunc) Handle(params GetDriversParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetDriversHandler interface for that can handle valid get drivers params
type GetDriversHandler interface {
	Handle(GetDriversParams, interface{}) middleware.Responder
}

// NewGetDrivers creates a new http.Handler for the get drivers operation
func NewGetDrivers(ctx *middleware.Context, handler GetDriversHandler) *GetDrivers {
	return &GetDrivers{Context: ctx, Handler: handler}
}

/*GetDrivers swagger:route GET /drivers drivers getDrivers

List all existing drivers

*/
type GetDrivers struct {
	Context *middleware.Context
	Handler GetDriversHandler
}

func (o *GetDrivers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDriversParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}