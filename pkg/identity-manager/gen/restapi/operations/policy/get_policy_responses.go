///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/identity-manager/gen/models"
)

// GetPolicyOKCode is the HTTP code returned for type GetPolicyOK
const GetPolicyOKCode int = 200

/*GetPolicyOK Successful operation

swagger:response getPolicyOK
*/
type GetPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *models.Policy `json:"body,omitempty"`
}

// NewGetPolicyOK creates GetPolicyOK with default headers values
func NewGetPolicyOK() *GetPolicyOK {
	return &GetPolicyOK{}
}

// WithPayload adds the payload to the get policy o k response
func (o *GetPolicyOK) WithPayload(payload *models.Policy) *GetPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get policy o k response
func (o *GetPolicyOK) SetPayload(payload *models.Policy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPolicyBadRequestCode is the HTTP code returned for type GetPolicyBadRequest
const GetPolicyBadRequestCode int = 400

/*GetPolicyBadRequest Invalid Name supplied

swagger:response getPolicyBadRequest
*/
type GetPolicyBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPolicyBadRequest creates GetPolicyBadRequest with default headers values
func NewGetPolicyBadRequest() *GetPolicyBadRequest {
	return &GetPolicyBadRequest{}
}

// WithPayload adds the payload to the get policy bad request response
func (o *GetPolicyBadRequest) WithPayload(payload *models.Error) *GetPolicyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get policy bad request response
func (o *GetPolicyBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPolicyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPolicyNotFoundCode is the HTTP code returned for type GetPolicyNotFound
const GetPolicyNotFoundCode int = 404

/*GetPolicyNotFound Policy not found

swagger:response getPolicyNotFound
*/
type GetPolicyNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPolicyNotFound creates GetPolicyNotFound with default headers values
func NewGetPolicyNotFound() *GetPolicyNotFound {
	return &GetPolicyNotFound{}
}

// WithPayload adds the payload to the get policy not found response
func (o *GetPolicyNotFound) WithPayload(payload *models.Error) *GetPolicyNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get policy not found response
func (o *GetPolicyNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPolicyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPolicyInternalServerErrorCode is the HTTP code returned for type GetPolicyInternalServerError
const GetPolicyInternalServerErrorCode int = 500

/*GetPolicyInternalServerError Internal error

swagger:response getPolicyInternalServerError
*/
type GetPolicyInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPolicyInternalServerError creates GetPolicyInternalServerError with default headers values
func NewGetPolicyInternalServerError() *GetPolicyInternalServerError {
	return &GetPolicyInternalServerError{}
}

// WithPayload adds the payload to the get policy internal server error response
func (o *GetPolicyInternalServerError) WithPayload(payload *models.Error) *GetPolicyInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get policy internal server error response
func (o *GetPolicyInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPolicyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
