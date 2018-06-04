// Code generated by go-swagger; DO NOT EDIT.

package account

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

// GetAccountAuditLogsWithHistoryReader is a Reader for the GetAccountAuditLogsWithHistory structure.
type GetAccountAuditLogsWithHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountAuditLogsWithHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountAuditLogsWithHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetAccountAuditLogsWithHistoryNotFound()
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

// NewGetAccountAuditLogsWithHistoryOK creates a GetAccountAuditLogsWithHistoryOK with default headers values
func NewGetAccountAuditLogsWithHistoryOK() *GetAccountAuditLogsWithHistoryOK {
	return &GetAccountAuditLogsWithHistoryOK{}
}

/*GetAccountAuditLogsWithHistoryOK handles this case with default header values.

successful operation
*/
type GetAccountAuditLogsWithHistoryOK struct {
	Payload []*kbmodel.AuditLog
}

func (o *GetAccountAuditLogsWithHistoryOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/auditLogsWithHistory][%d] getAccountAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetAccountAuditLogsWithHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAuditLogsWithHistoryNotFound creates a GetAccountAuditLogsWithHistoryNotFound with default headers values
func NewGetAccountAuditLogsWithHistoryNotFound() *GetAccountAuditLogsWithHistoryNotFound {
	return &GetAccountAuditLogsWithHistoryNotFound{}
}

/*GetAccountAuditLogsWithHistoryNotFound handles this case with default header values.

Account not found
*/
type GetAccountAuditLogsWithHistoryNotFound struct {
}

func (o *GetAccountAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/auditLogsWithHistory][%d] getAccountAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetAccountAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}