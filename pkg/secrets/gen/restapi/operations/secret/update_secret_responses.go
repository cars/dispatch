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

// UpdateSecretCreatedCode is the HTTP code returned for type UpdateSecretCreated
const UpdateSecretCreatedCode int = 201

/*UpdateSecretCreated The updated secret

swagger:response updateSecretCreated
*/
type UpdateSecretCreated struct {

	/*
	  In: Body
	*/
	Payload *v1.Secret `json:"body,omitempty"`
}

// NewUpdateSecretCreated creates UpdateSecretCreated with default headers values
func NewUpdateSecretCreated() *UpdateSecretCreated {

	return &UpdateSecretCreated{}
}

// WithPayload adds the payload to the update secret created response
func (o *UpdateSecretCreated) WithPayload(payload *v1.Secret) *UpdateSecretCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update secret created response
func (o *UpdateSecretCreated) SetPayload(payload *v1.Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSecretCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateSecretBadRequestCode is the HTTP code returned for type UpdateSecretBadRequest
const UpdateSecretBadRequestCode int = 400

/*UpdateSecretBadRequest Bad Request

swagger:response updateSecretBadRequest
*/
type UpdateSecretBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewUpdateSecretBadRequest creates UpdateSecretBadRequest with default headers values
func NewUpdateSecretBadRequest() *UpdateSecretBadRequest {

	return &UpdateSecretBadRequest{}
}

// WithPayload adds the payload to the update secret bad request response
func (o *UpdateSecretBadRequest) WithPayload(payload *v1.Error) *UpdateSecretBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update secret bad request response
func (o *UpdateSecretBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSecretBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateSecretUnauthorizedCode is the HTTP code returned for type UpdateSecretUnauthorized
const UpdateSecretUnauthorizedCode int = 401

/*UpdateSecretUnauthorized Unauthorized Request

swagger:response updateSecretUnauthorized
*/
type UpdateSecretUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewUpdateSecretUnauthorized creates UpdateSecretUnauthorized with default headers values
func NewUpdateSecretUnauthorized() *UpdateSecretUnauthorized {

	return &UpdateSecretUnauthorized{}
}

// WithPayload adds the payload to the update secret unauthorized response
func (o *UpdateSecretUnauthorized) WithPayload(payload *v1.Error) *UpdateSecretUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update secret unauthorized response
func (o *UpdateSecretUnauthorized) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSecretUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateSecretForbiddenCode is the HTTP code returned for type UpdateSecretForbidden
const UpdateSecretForbiddenCode int = 403

/*UpdateSecretForbidden access to this resource is forbidden

swagger:response updateSecretForbidden
*/
type UpdateSecretForbidden struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewUpdateSecretForbidden creates UpdateSecretForbidden with default headers values
func NewUpdateSecretForbidden() *UpdateSecretForbidden {

	return &UpdateSecretForbidden{}
}

// WithPayload adds the payload to the update secret forbidden response
func (o *UpdateSecretForbidden) WithPayload(payload *v1.Error) *UpdateSecretForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update secret forbidden response
func (o *UpdateSecretForbidden) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSecretForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateSecretNotFoundCode is the HTTP code returned for type UpdateSecretNotFound
const UpdateSecretNotFoundCode int = 404

/*UpdateSecretNotFound Resource Not Found if no secret exists with the given name

swagger:response updateSecretNotFound
*/
type UpdateSecretNotFound struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewUpdateSecretNotFound creates UpdateSecretNotFound with default headers values
func NewUpdateSecretNotFound() *UpdateSecretNotFound {

	return &UpdateSecretNotFound{}
}

// WithPayload adds the payload to the update secret not found response
func (o *UpdateSecretNotFound) WithPayload(payload *v1.Error) *UpdateSecretNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update secret not found response
func (o *UpdateSecretNotFound) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSecretNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateSecretDefault generic error

swagger:response updateSecretDefault
*/
type UpdateSecretDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewUpdateSecretDefault creates UpdateSecretDefault with default headers values
func NewUpdateSecretDefault(code int) *UpdateSecretDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateSecretDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update secret default response
func (o *UpdateSecretDefault) WithStatusCode(code int) *UpdateSecretDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update secret default response
func (o *UpdateSecretDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update secret default response
func (o *UpdateSecretDefault) WithPayload(payload *v1.Error) *UpdateSecretDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update secret default response
func (o *UpdateSecretDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSecretDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
