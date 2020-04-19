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

	"github.com/bozaro/tech-db-forum/generated/models"
)

// NewPostsCreateParams creates a new PostsCreateParams object
// with the default values initialized.
func NewPostsCreateParams() *PostsCreateParams {
	var ()
	return &PostsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostsCreateParamsWithTimeout creates a new PostsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostsCreateParamsWithTimeout(timeout time.Duration) *PostsCreateParams {
	var ()
	return &PostsCreateParams{

		timeout: timeout,
	}
}

// NewPostsCreateParamsWithContext creates a new PostsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostsCreateParamsWithContext(ctx context.Context) *PostsCreateParams {
	var ()
	return &PostsCreateParams{

		Context: ctx,
	}
}

// NewPostsCreateParamsWithHTTPClient creates a new PostsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostsCreateParamsWithHTTPClient(client *http.Client) *PostsCreateParams {
	var ()
	return &PostsCreateParams{
		HTTPClient: client,
	}
}

/*PostsCreateParams contains all the parameters to send to the API endpoint
for the posts create operation typically these are written to a http.Request
*/
type PostsCreateParams struct {

	/*Posts
	  Список создаваемых постов.

	*/
	Posts models.Posts
	/*SlugOrID
	  Идентификатор ветки обсуждения.

	*/
	SlugOrID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the posts create params
func (o *PostsCreateParams) WithTimeout(timeout time.Duration) *PostsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the posts create params
func (o *PostsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the posts create params
func (o *PostsCreateParams) WithContext(ctx context.Context) *PostsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the posts create params
func (o *PostsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the posts create params
func (o *PostsCreateParams) WithHTTPClient(client *http.Client) *PostsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the posts create params
func (o *PostsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPosts adds the posts to the posts create params
func (o *PostsCreateParams) WithPosts(posts models.Posts) *PostsCreateParams {
	o.SetPosts(posts)
	return o
}

// SetPosts adds the posts to the posts create params
func (o *PostsCreateParams) SetPosts(posts models.Posts) {
	o.Posts = posts
}

// WithSlugOrID adds the slugOrID to the posts create params
func (o *PostsCreateParams) WithSlugOrID(slugOrID string) *PostsCreateParams {
	o.SetSlugOrID(slugOrID)
	return o
}

// SetSlugOrID adds the slugOrId to the posts create params
func (o *PostsCreateParams) SetSlugOrID(slugOrID string) {
	o.SlugOrID = slugOrID
}

// WriteToRequest writes these params to a swagger request
func (o *PostsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Posts != nil {
		if err := r.SetBodyParam(o.Posts); err != nil {
			return err
		}
	}

	// path param slug_or_id
	if err := r.SetPathParam("slug_or_id", o.SlugOrID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
