///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// DeleteSecretNoContentCode is the HTTP code returned for type DeleteSecretNoContent
const DeleteSecretNoContentCode int = 204

/*DeleteSecretNoContent Successful deletion

swagger:response deleteSecretNoContent
*/
type DeleteSecretNoContent struct {
}

// NewDeleteSecretNoContent creates DeleteSecretNoContent with default headers values
func NewDeleteSecretNoContent() *DeleteSecretNoContent {

	return &DeleteSecretNoContent{}
}

// WriteResponse to the client
func (o *DeleteSecretNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteSecretBadRequestCode is the HTTP code returned for type DeleteSecretBadRequest
const DeleteSecretBadRequestCode int = 400

/*DeleteSecretBadRequest Bad Request

swagger:response deleteSecretBadRequest
*/
type DeleteSecretBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeleteSecretBadRequest creates DeleteSecretBadRequest with default headers values
func NewDeleteSecretBadRequest() *DeleteSecretBadRequest {

	return &DeleteSecretBadRequest{}
}

// WithPayload adds the payload to the delete secret bad request response
func (o *DeleteSecretBadRequest) WithPayload(payload *v1.Error) *DeleteSecretBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete secret bad request response
func (o *DeleteSecretBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSecretBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteSecretUnauthorizedCode is the HTTP code returned for type DeleteSecretUnauthorized
const DeleteSecretUnauthorizedCode int = 401

/*DeleteSecretUnauthorized Unauthorized Request

swagger:response deleteSecretUnauthorized
*/
type DeleteSecretUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeleteSecretUnauthorized creates DeleteSecretUnauthorized with default headers values
func NewDeleteSecretUnauthorized() *DeleteSecretUnauthorized {

	return &DeleteSecretUnauthorized{}
}

// WithPayload adds the payload to the delete secret unauthorized response
func (o *DeleteSecretUnauthorized) WithPayload(payload *v1.Error) *DeleteSecretUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete secret unauthorized response
func (o *DeleteSecretUnauthorized) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSecretUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteSecretForbiddenCode is the HTTP code returned for type DeleteSecretForbidden
const DeleteSecretForbiddenCode int = 403

/*DeleteSecretForbidden access to this resource is forbidden

swagger:response deleteSecretForbidden
*/
type DeleteSecretForbidden struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeleteSecretForbidden creates DeleteSecretForbidden with default headers values
func NewDeleteSecretForbidden() *DeleteSecretForbidden {

	return &DeleteSecretForbidden{}
}

// WithPayload adds the payload to the delete secret forbidden response
func (o *DeleteSecretForbidden) WithPayload(payload *v1.Error) *DeleteSecretForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete secret forbidden response
func (o *DeleteSecretForbidden) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSecretForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteSecretNotFoundCode is the HTTP code returned for type DeleteSecretNotFound
const DeleteSecretNotFoundCode int = 404

/*DeleteSecretNotFound Resource Not Found if no secret exists with the given name

swagger:response deleteSecretNotFound
*/
type DeleteSecretNotFound struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeleteSecretNotFound creates DeleteSecretNotFound with default headers values
func NewDeleteSecretNotFound() *DeleteSecretNotFound {

	return &DeleteSecretNotFound{}
}

// WithPayload adds the payload to the delete secret not found response
func (o *DeleteSecretNotFound) WithPayload(payload *v1.Error) *DeleteSecretNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete secret not found response
func (o *DeleteSecretNotFound) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSecretNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteSecretDefault generic error

swagger:response deleteSecretDefault
*/
type DeleteSecretDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeleteSecretDefault creates DeleteSecretDefault with default headers values
func NewDeleteSecretDefault(code int) *DeleteSecretDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteSecretDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete secret default response
func (o *DeleteSecretDefault) WithStatusCode(code int) *DeleteSecretDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete secret default response
func (o *DeleteSecretDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete secret default response
func (o *DeleteSecretDefault) WithPayload(payload *v1.Error) *DeleteSecretDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete secret default response
func (o *DeleteSecretDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSecretDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
