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
	"github.com/go-openapi/swag"
)

// NewForumGetThreadsParams creates a new ForumGetThreadsParams object
// with the default values initialized.
func NewForumGetThreadsParams() *ForumGetThreadsParams {
	var (
		limitDefault = int32(100)
	)
	return &ForumGetThreadsParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewForumGetThreadsParamsWithTimeout creates a new ForumGetThreadsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewForumGetThreadsParamsWithTimeout(timeout time.Duration) *ForumGetThreadsParams {
	var (
		limitDefault = int32(100)
	)
	return &ForumGetThreadsParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewForumGetThreadsParamsWithContext creates a new ForumGetThreadsParams object
// with the default values initialized, and the ability to set a context for a request
func NewForumGetThreadsParamsWithContext(ctx context.Context) *ForumGetThreadsParams {
	var (
		limitDefault = int32(100)
	)
	return &ForumGetThreadsParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewForumGetThreadsParamsWithHTTPClient creates a new ForumGetThreadsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewForumGetThreadsParamsWithHTTPClient(client *http.Client) *ForumGetThreadsParams {
	var (
		limitDefault = int32(100)
	)
	return &ForumGetThreadsParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*ForumGetThreadsParams contains all the parameters to send to the API endpoint
for the forum get threads operation typically these are written to a http.Request
*/
type ForumGetThreadsParams struct {

	/*Desc
	  Флаг сортировки по убыванию.


	*/
	Desc *bool
	/*Limit
	  Максимальное кол-во возвращаемых записей.

	*/
	Limit *int32
	/*Since
	  Дата создания ветви обсуждения, с которой будут выводиться записи
	(ветвь обсуждения с указанной датой попадает в результат выборки).


	*/
	Since *strfmt.DateTime
	/*Slug
	  Идентификатор форума.

	*/
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the forum get threads params
func (o *ForumGetThreadsParams) WithTimeout(timeout time.Duration) *ForumGetThreadsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the forum get threads params
func (o *ForumGetThreadsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the forum get threads params
func (o *ForumGetThreadsParams) WithContext(ctx context.Context) *ForumGetThreadsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the forum get threads params
func (o *ForumGetThreadsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the forum get threads params
func (o *ForumGetThreadsParams) WithHTTPClient(client *http.Client) *ForumGetThreadsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the forum get threads params
func (o *ForumGetThreadsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDesc adds the desc to the forum get threads params
func (o *ForumGetThreadsParams) WithDesc(desc *bool) *ForumGetThreadsParams {
	o.SetDesc(desc)
	return o
}

// SetDesc adds the desc to the forum get threads params
func (o *ForumGetThreadsParams) SetDesc(desc *bool) {
	o.Desc = desc
}

// WithLimit adds the limit to the forum get threads params
func (o *ForumGetThreadsParams) WithLimit(limit *int32) *ForumGetThreadsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the forum get threads params
func (o *ForumGetThreadsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithSince adds the since to the forum get threads params
func (o *ForumGetThreadsParams) WithSince(since *strfmt.DateTime) *ForumGetThreadsParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the forum get threads params
func (o *ForumGetThreadsParams) SetSince(since *strfmt.DateTime) {
	o.Since = since
}

// WithSlug adds the slug to the forum get threads params
func (o *ForumGetThreadsParams) WithSlug(slug string) *ForumGetThreadsParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the forum get threads params
func (o *ForumGetThreadsParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ForumGetThreadsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Desc != nil {

		// query param desc
		var qrDesc bool
		if o.Desc != nil {
			qrDesc = *o.Desc
		}
		qDesc := swag.FormatBool(qrDesc)
		if qDesc != "" {
			if err := r.SetQueryParam("desc", qDesc); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Since != nil {

		// query param since
		var qrSince strfmt.DateTime
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := qrSince.String()
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	// path param slug
	if err := r.SetPathParam("slug", o.Slug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
