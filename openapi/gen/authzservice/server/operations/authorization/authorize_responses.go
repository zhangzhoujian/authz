// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "authz/openapi/gen/authzservice/models"
)

// AuthorizeOKCode is the HTTP code returned for type AuthorizeOK
const AuthorizeOKCode int = 200

/*AuthorizeOK Status 200

swagger:response authorizeOK
*/
type AuthorizeOK struct {

	/*
	  In: Body
	*/
	Payload *AuthorizeOKBody `json:"body,omitempty"`
}

// NewAuthorizeOK creates AuthorizeOK with default headers values
func NewAuthorizeOK() *AuthorizeOK {

	return &AuthorizeOK{}
}

// WithPayload adds the payload to the authorize o k response
func (o *AuthorizeOK) WithPayload(payload *AuthorizeOKBody) *AuthorizeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the authorize o k response
func (o *AuthorizeOK) SetPayload(payload *AuthorizeOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthorizeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AuthorizeBadRequestCode is the HTTP code returned for type AuthorizeBadRequest
const AuthorizeBadRequestCode int = 400

/*AuthorizeBadRequest Status 400

swagger:response authorizeBadRequest
*/
type AuthorizeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorCode `json:"body,omitempty"`
}

// NewAuthorizeBadRequest creates AuthorizeBadRequest with default headers values
func NewAuthorizeBadRequest() *AuthorizeBadRequest {

	return &AuthorizeBadRequest{}
}

// WithPayload adds the payload to the authorize bad request response
func (o *AuthorizeBadRequest) WithPayload(payload *models.ErrorCode) *AuthorizeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the authorize bad request response
func (o *AuthorizeBadRequest) SetPayload(payload *models.ErrorCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthorizeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AuthorizeUnauthorizedCode is the HTTP code returned for type AuthorizeUnauthorized
const AuthorizeUnauthorizedCode int = 401

/*AuthorizeUnauthorized Status 401

swagger:response authorizeUnauthorized
*/
type AuthorizeUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorCode `json:"body,omitempty"`
}

// NewAuthorizeUnauthorized creates AuthorizeUnauthorized with default headers values
func NewAuthorizeUnauthorized() *AuthorizeUnauthorized {

	return &AuthorizeUnauthorized{}
}

// WithPayload adds the payload to the authorize unauthorized response
func (o *AuthorizeUnauthorized) WithPayload(payload *models.ErrorCode) *AuthorizeUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the authorize unauthorized response
func (o *AuthorizeUnauthorized) SetPayload(payload *models.ErrorCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthorizeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AuthorizeForbiddenCode is the HTTP code returned for type AuthorizeForbidden
const AuthorizeForbiddenCode int = 403

/*AuthorizeForbidden Status 403

swagger:response authorizeForbidden
*/
type AuthorizeForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorCode `json:"body,omitempty"`
}

// NewAuthorizeForbidden creates AuthorizeForbidden with default headers values
func NewAuthorizeForbidden() *AuthorizeForbidden {

	return &AuthorizeForbidden{}
}

// WithPayload adds the payload to the authorize forbidden response
func (o *AuthorizeForbidden) WithPayload(payload *models.ErrorCode) *AuthorizeForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the authorize forbidden response
func (o *AuthorizeForbidden) SetPayload(payload *models.ErrorCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthorizeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AuthorizeInternalServerErrorCode is the HTTP code returned for type AuthorizeInternalServerError
const AuthorizeInternalServerErrorCode int = 500

/*AuthorizeInternalServerError Status 500

swagger:response authorizeInternalServerError
*/
type AuthorizeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorCode `json:"body,omitempty"`
}

// NewAuthorizeInternalServerError creates AuthorizeInternalServerError with default headers values
func NewAuthorizeInternalServerError() *AuthorizeInternalServerError {

	return &AuthorizeInternalServerError{}
}

// WithPayload adds the payload to the authorize internal server error response
func (o *AuthorizeInternalServerError) WithPayload(payload *models.ErrorCode) *AuthorizeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the authorize internal server error response
func (o *AuthorizeInternalServerError) SetPayload(payload *models.ErrorCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthorizeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
