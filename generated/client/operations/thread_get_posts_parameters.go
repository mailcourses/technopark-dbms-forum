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

// NewThreadGetPostsParams creates a new ThreadGetPostsParams object
// with the default values initialized.
func NewThreadGetPostsParams() *ThreadGetPostsParams {
	var (
		limitDefault = int32(100)
		sortDefault  = string("flat")
	)
	return &ThreadGetPostsParams{
		Limit: &limitDefault,
		Sort:  &sortDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewThreadGetPostsParamsWithTimeout creates a new ThreadGetPostsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewThreadGetPostsParamsWithTimeout(timeout time.Duration) *ThreadGetPostsParams {
	var (
		limitDefault = int32(100)
		sortDefault  = string("flat")
	)
	return &ThreadGetPostsParams{
		Limit: &limitDefault,
		Sort:  &sortDefault,

		timeout: timeout,
	}
}

// NewThreadGetPostsParamsWithContext creates a new ThreadGetPostsParams object
// with the default values initialized, and the ability to set a context for a request
func NewThreadGetPostsParamsWithContext(ctx context.Context) *ThreadGetPostsParams {
	var (
		limitDefault = int32(100)
		sortDefault  = string("flat")
	)
	return &ThreadGetPostsParams{
		Limit: &limitDefault,
		Sort:  &sortDefault,

		Context: ctx,
	}
}

// NewThreadGetPostsParamsWithHTTPClient creates a new ThreadGetPostsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewThreadGetPostsParamsWithHTTPClient(client *http.Client) *ThreadGetPostsParams {
	var (
		limitDefault = int32(100)
		sortDefault  = string("flat")
	)
	return &ThreadGetPostsParams{
		Limit:      &limitDefault,
		Sort:       &sortDefault,
		HTTPClient: client,
	}
}

/*ThreadGetPostsParams contains all the parameters to send to the API endpoint
for the thread get posts operation typically these are written to a http.Request
*/
type ThreadGetPostsParams struct {

	/*Desc
	  Флаг сортировки по убыванию.


	*/
	Desc *bool
	/*Limit
	  Максимальное кол-во возвращаемых записей.

	*/
	Limit *int32
	/*Since
	  Идентификатор поста, после которого будут выводиться записи
	(пост с данным идентификатором в результат не попадает).


	*/
	Since *int64
	/*SlugOrID
	  Идентификатор ветки обсуждения.

	*/
	SlugOrID string
	/*Sort
	  Вид сортировки:

	 * flat - по дате, комментарии выводятся простым списком в порядке создания;
	 * tree - древовидный, комментарии выводятся отсортированные в дереве
	   по N штук;
	 * parent_tree - древовидные с пагинацией по родительским (parent_tree),
	   на странице N родительских комментов и все комментарии прикрепленные
	   к ним, в древвидном отображение.

	Подробности: https://park.mail.ru/blog/topic/view/1191/


	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the thread get posts params
func (o *ThreadGetPostsParams) WithTimeout(timeout time.Duration) *ThreadGetPostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thread get posts params
func (o *ThreadGetPostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thread get posts params
func (o *ThreadGetPostsParams) WithContext(ctx context.Context) *ThreadGetPostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thread get posts params
func (o *ThreadGetPostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thread get posts params
func (o *ThreadGetPostsParams) WithHTTPClient(client *http.Client) *ThreadGetPostsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thread get posts params
func (o *ThreadGetPostsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDesc adds the desc to the thread get posts params
func (o *ThreadGetPostsParams) WithDesc(desc *bool) *ThreadGetPostsParams {
	o.SetDesc(desc)
	return o
}

// SetDesc adds the desc to the thread get posts params
func (o *ThreadGetPostsParams) SetDesc(desc *bool) {
	o.Desc = desc
}

// WithLimit adds the limit to the thread get posts params
func (o *ThreadGetPostsParams) WithLimit(limit *int32) *ThreadGetPostsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the thread get posts params
func (o *ThreadGetPostsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithSince adds the since to the thread get posts params
func (o *ThreadGetPostsParams) WithSince(since *int64) *ThreadGetPostsParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the thread get posts params
func (o *ThreadGetPostsParams) SetSince(since *int64) {
	o.Since = since
}

// WithSlugOrID adds the slugOrID to the thread get posts params
func (o *ThreadGetPostsParams) WithSlugOrID(slugOrID string) *ThreadGetPostsParams {
	o.SetSlugOrID(slugOrID)
	return o
}

// SetSlugOrID adds the slugOrId to the thread get posts params
func (o *ThreadGetPostsParams) SetSlugOrID(slugOrID string) {
	o.SlugOrID = slugOrID
}

// WithSort adds the sort to the thread get posts params
func (o *ThreadGetPostsParams) WithSort(sort *string) *ThreadGetPostsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the thread get posts params
func (o *ThreadGetPostsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ThreadGetPostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
		var qrSince int64
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := swag.FormatInt64(qrSince)
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	// path param slug_or_id
	if err := r.SetPathParam("slug_or_id", o.SlugOrID); err != nil {
		return err
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
