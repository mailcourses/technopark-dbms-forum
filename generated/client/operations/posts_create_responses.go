// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mailcourses/technopark-dbms-forum/generated/models"
)

// PostsCreateReader is a Reader for the PostsCreate structure.
type PostsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPostsCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostsCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostsCreateCreated creates a PostsCreateCreated with default headers values
func NewPostsCreateCreated() *PostsCreateCreated {
	return &PostsCreateCreated{}
}

/*PostsCreateCreated handles this case with default header values.

Посты успешно созданы.
Возвращает данные созданных постов в том же порядке, в котором их передали на вход метода.

*/
type PostsCreateCreated struct {
	Payload models.Posts
}

func (o *PostsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /thread/{slug_or_id}/create][%d] postsCreateCreated  %+v", 201, o.Payload)
}

func (o *PostsCreateCreated) GetPayload() models.Posts {
	return o.Payload
}

func (o *PostsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostsCreateNotFound creates a PostsCreateNotFound with default headers values
func NewPostsCreateNotFound() *PostsCreateNotFound {
	return &PostsCreateNotFound{}
}

/*PostsCreateNotFound handles this case with default header values.

Ветка обсуждения отсутствует в базе данных.

*/
type PostsCreateNotFound struct {
	Payload *models.Error
}

func (o *PostsCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /thread/{slug_or_id}/create][%d] postsCreateNotFound  %+v", 404, o.Payload)
}

func (o *PostsCreateNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostsCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostsCreateConflict creates a PostsCreateConflict with default headers values
func NewPostsCreateConflict() *PostsCreateConflict {
	return &PostsCreateConflict{}
}

/*PostsCreateConflict handles this case with default header values.

Хотя бы один родительский пост отсутсвует в текущей ветке обсуждения.

*/
type PostsCreateConflict struct {
	Payload *models.Error
}

func (o *PostsCreateConflict) Error() string {
	return fmt.Sprintf("[POST /thread/{slug_or_id}/create][%d] postsCreateConflict  %+v", 409, o.Payload)
}

func (o *PostsCreateConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostsCreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
