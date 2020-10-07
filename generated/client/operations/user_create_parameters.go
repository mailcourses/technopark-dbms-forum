// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/mailcourses/technopark-dbms-forum/generated/models"
)

// NewUserCreateParams creates a new UserCreateParams object
// with the default values initialized.
func NewUserCreateParams() *UserCreateParams {
	var ()
	return &UserCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserCreateParamsWithTimeout creates a new UserCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserCreateParamsWithTimeout(timeout time.Duration) *UserCreateParams {
	var ()
	return &UserCreateParams{

		timeout: timeout,
	}
}

// NewUserCreateParamsWithContext creates a new UserCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserCreateParamsWithContext(ctx context.Context) *UserCreateParams {
	var ()
	return &UserCreateParams{

		Context: ctx,
	}
}

// NewUserCreateParamsWithHTTPClient creates a new UserCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserCreateParamsWithHTTPClient(client *http.Client) *UserCreateParams {
	var ()
	return &UserCreateParams{
		HTTPClient: client,
	}
}

/*UserCreateParams contains all the parameters to send to the API endpoint
for the user create operation typically these are written to a http.Request
*/
type UserCreateParams struct {

	/*Nickname
	  Идентификатор пользователя.

	*/
	Nickname string
	/*Profile
	  Данные пользовательского профиля.

	*/
	Profile *models.User

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user create params
func (o *UserCreateParams) WithTimeout(timeout time.Duration) *UserCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user create params
func (o *UserCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user create params
func (o *UserCreateParams) WithContext(ctx context.Context) *UserCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user create params
func (o *UserCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user create params
func (o *UserCreateParams) WithHTTPClient(client *http.Client) *UserCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user create params
func (o *UserCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNickname adds the nickname to the user create params
func (o *UserCreateParams) WithNickname(nickname string) *UserCreateParams {
	o.SetNickname(nickname)
	return o
}

// SetNickname adds the nickname to the user create params
func (o *UserCreateParams) SetNickname(nickname string) {
	o.Nickname = nickname
}

// WithProfile adds the profile to the user create params
func (o *UserCreateParams) WithProfile(profile *models.User) *UserCreateParams {
	o.SetProfile(profile)
	return o
}

// SetProfile adds the profile to the user create params
func (o *UserCreateParams) SetProfile(profile *models.User) {
	o.Profile = profile
}

// WriteToRequest writes these params to a swagger request
func (o *UserCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param nickname
	if err := r.SetPathParam("nickname", o.Nickname); err != nil {
		return err
	}

	if o.Profile != nil {
		if err := r.SetBodyParam(o.Profile); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
