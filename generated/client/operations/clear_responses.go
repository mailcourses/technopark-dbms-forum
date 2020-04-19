// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ClearReader is a Reader for the Clear structure.
type ClearReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClearReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClearOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewClearOK creates a ClearOK with default headers values
func NewClearOK() *ClearOK {
	return &ClearOK{}
}

/*ClearOK handles this case with default header values.

Очистка базы успешно завершена
*/
type ClearOK struct {
}

func (o *ClearOK) Error() string {
	return fmt.Sprintf("[POST /service/clear][%d] clearOK ", 200)
}

func (o *ClearOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
