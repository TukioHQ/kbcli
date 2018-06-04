// Code generated by go-swagger; DO NOT EDIT.

package payment_transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// GetTransactionTagsReader is a Reader for the GetTransactionTags structure.
type GetTransactionTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTransactionTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetTransactionTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewGetTransactionTagsOK creates a GetTransactionTagsOK with default headers values
func NewGetTransactionTagsOK() *GetTransactionTagsOK {
	return &GetTransactionTagsOK{}
}

/*GetTransactionTagsOK handles this case with default header values.

successful operation
*/
type GetTransactionTagsOK struct {
	Payload []*kbmodel.Tag
}

func (o *GetTransactionTagsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}/tags][%d] getTransactionTagsOK  %+v", 200, o.Payload)
}

func (o *GetTransactionTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionTagsBadRequest creates a GetTransactionTagsBadRequest with default headers values
func NewGetTransactionTagsBadRequest() *GetTransactionTagsBadRequest {
	return &GetTransactionTagsBadRequest{}
}

/*GetTransactionTagsBadRequest handles this case with default header values.

Invalid transaction id supplied
*/
type GetTransactionTagsBadRequest struct {
}

func (o *GetTransactionTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}/tags][%d] getTransactionTagsBadRequest ", 400)
}

func (o *GetTransactionTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTransactionTagsNotFound creates a GetTransactionTagsNotFound with default headers values
func NewGetTransactionTagsNotFound() *GetTransactionTagsNotFound {
	return &GetTransactionTagsNotFound{}
}

/*GetTransactionTagsNotFound handles this case with default header values.

Invoice not found
*/
type GetTransactionTagsNotFound struct {
}

func (o *GetTransactionTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}/tags][%d] getTransactionTagsNotFound ", 404)
}

func (o *GetTransactionTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}