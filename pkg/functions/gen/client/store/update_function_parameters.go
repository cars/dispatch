///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// NewUpdateFunctionParams creates a new UpdateFunctionParams object
// with the default values initialized.
func NewUpdateFunctionParams() *UpdateFunctionParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &UpdateFunctionParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFunctionParamsWithTimeout creates a new UpdateFunctionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateFunctionParamsWithTimeout(timeout time.Duration) *UpdateFunctionParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &UpdateFunctionParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		timeout: timeout,
	}
}

// NewUpdateFunctionParamsWithContext creates a new UpdateFunctionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateFunctionParamsWithContext(ctx context.Context) *UpdateFunctionParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &UpdateFunctionParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		Context: ctx,
	}
}

// NewUpdateFunctionParamsWithHTTPClient creates a new UpdateFunctionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateFunctionParamsWithHTTPClient(client *http.Client) *UpdateFunctionParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &UpdateFunctionParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,
		HTTPClient:       client,
	}
}

/*UpdateFunctionParams contains all the parameters to send to the API endpoint
for the update function operation typically these are written to a http.Request
*/
type UpdateFunctionParams struct {

	/*XDispatchOrg*/
	XDispatchOrg *string
	/*XDispatchProject*/
	XDispatchProject *string
	/*Body
	  function object

	*/
	Body *v1.Function
	/*FunctionName
	  Name of function to work on

	*/
	FunctionName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update function params
func (o *UpdateFunctionParams) WithTimeout(timeout time.Duration) *UpdateFunctionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update function params
func (o *UpdateFunctionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update function params
func (o *UpdateFunctionParams) WithContext(ctx context.Context) *UpdateFunctionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update function params
func (o *UpdateFunctionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update function params
func (o *UpdateFunctionParams) WithHTTPClient(client *http.Client) *UpdateFunctionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update function params
func (o *UpdateFunctionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDispatchOrg adds the xDispatchOrg to the update function params
func (o *UpdateFunctionParams) WithXDispatchOrg(xDispatchOrg *string) *UpdateFunctionParams {
	o.SetXDispatchOrg(xDispatchOrg)
	return o
}

// SetXDispatchOrg adds the xDispatchOrg to the update function params
func (o *UpdateFunctionParams) SetXDispatchOrg(xDispatchOrg *string) {
	o.XDispatchOrg = xDispatchOrg
}

// WithXDispatchProject adds the xDispatchProject to the update function params
func (o *UpdateFunctionParams) WithXDispatchProject(xDispatchProject *string) *UpdateFunctionParams {
	o.SetXDispatchProject(xDispatchProject)
	return o
}

// SetXDispatchProject adds the xDispatchProject to the update function params
func (o *UpdateFunctionParams) SetXDispatchProject(xDispatchProject *string) {
	o.XDispatchProject = xDispatchProject
}

// WithBody adds the body to the update function params
func (o *UpdateFunctionParams) WithBody(body *v1.Function) *UpdateFunctionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update function params
func (o *UpdateFunctionParams) SetBody(body *v1.Function) {
	o.Body = body
}

// WithFunctionName adds the functionName to the update function params
func (o *UpdateFunctionParams) WithFunctionName(functionName string) *UpdateFunctionParams {
	o.SetFunctionName(functionName)
	return o
}

// SetFunctionName adds the functionName to the update function params
func (o *UpdateFunctionParams) SetFunctionName(functionName string) {
	o.FunctionName = functionName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFunctionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XDispatchOrg != nil {

		// header param X-Dispatch-Org
		if err := r.SetHeaderParam("X-Dispatch-Org", *o.XDispatchOrg); err != nil {
			return err
		}

	}

	if o.XDispatchProject != nil {

		// header param X-Dispatch-Project
		if err := r.SetHeaderParam("X-Dispatch-Project", *o.XDispatchProject); err != nil {
			return err
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param functionName
	if err := r.SetPathParam("functionName", o.FunctionName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
