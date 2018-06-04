// Code generated by go-swagger; DO NOT EDIT.

package security

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

// GetRoleDefinitionReader is a Reader for the GetRoleDefinition structure.
type GetRoleDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRoleDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewGetRoleDefinitionOK creates a GetRoleDefinitionOK with default headers values
func NewGetRoleDefinitionOK() *GetRoleDefinitionOK {
	return &GetRoleDefinitionOK{}
}

/*GetRoleDefinitionOK handles this case with default header values.

successful operation
*/
type GetRoleDefinitionOK struct {
	Payload *kbmodel.RoleDefinition
}

func (o *GetRoleDefinitionOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/security/roles/{role}][%d] getRoleDefinitionOK  %+v", 200, o.Payload)
}

func (o *GetRoleDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.RoleDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}