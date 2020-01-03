// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetInstancesInstanceIDNetworkInterfacesReader is a Reader for the GetInstancesInstanceIDNetworkInterfaces structure.
type GetInstancesInstanceIDNetworkInterfacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesInstanceIDNetworkInterfacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstancesInstanceIDNetworkInterfacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetInstancesInstanceIDNetworkInterfacesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstancesInstanceIDNetworkInterfacesOK creates a GetInstancesInstanceIDNetworkInterfacesOK with default headers values
func NewGetInstancesInstanceIDNetworkInterfacesOK() *GetInstancesInstanceIDNetworkInterfacesOK {
	return &GetInstancesInstanceIDNetworkInterfacesOK{}
}

/*GetInstancesInstanceIDNetworkInterfacesOK handles this case with default header values.

dummy
*/
type GetInstancesInstanceIDNetworkInterfacesOK struct {
	Payload *GetInstancesInstanceIDNetworkInterfacesOKBody
}

func (o *GetInstancesInstanceIDNetworkInterfacesOK) Error() string {
	return fmt.Sprintf("[GET /instances/{instance_id}/network_interfaces][%d] getInstancesInstanceIdNetworkInterfacesOK  %+v", 200, o.Payload)
}

func (o *GetInstancesInstanceIDNetworkInterfacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetInstancesInstanceIDNetworkInterfacesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesInstanceIDNetworkInterfacesNotFound creates a GetInstancesInstanceIDNetworkInterfacesNotFound with default headers values
func NewGetInstancesInstanceIDNetworkInterfacesNotFound() *GetInstancesInstanceIDNetworkInterfacesNotFound {
	return &GetInstancesInstanceIDNetworkInterfacesNotFound{}
}

/*GetInstancesInstanceIDNetworkInterfacesNotFound handles this case with default header values.

error
*/
type GetInstancesInstanceIDNetworkInterfacesNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetInstancesInstanceIDNetworkInterfacesNotFound) Error() string {
	return fmt.Sprintf("[GET /instances/{instance_id}/network_interfaces][%d] getInstancesInstanceIdNetworkInterfacesNotFound  %+v", 404, o.Payload)
}

func (o *GetInstancesInstanceIDNetworkInterfacesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetInstancesInstanceIDNetworkInterfacesOKBody NetworkInterfaceCollection
swagger:model GetInstancesInstanceIDNetworkInterfacesOKBody
*/
type GetInstancesInstanceIDNetworkInterfacesOKBody struct {

	// Collection of network interfaces
	NetworkInterfaces []*models.InstanceNetworkInterface `json:"network_interfaces,omitempty"`
}

// Validate validates this get instances instance ID network interfaces o k body
func (o *GetInstancesInstanceIDNetworkInterfacesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNetworkInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInstancesInstanceIDNetworkInterfacesOKBody) validateNetworkInterfaces(formats strfmt.Registry) error {

	if swag.IsZero(o.NetworkInterfaces) { // not required
		return nil
	}

	for i := 0; i < len(o.NetworkInterfaces); i++ {
		if swag.IsZero(o.NetworkInterfaces[i]) { // not required
			continue
		}

		if o.NetworkInterfaces[i] != nil {
			if err := o.NetworkInterfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getInstancesInstanceIdNetworkInterfacesOK" + "." + "network_interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetInstancesInstanceIDNetworkInterfacesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstancesInstanceIDNetworkInterfacesOKBody) UnmarshalBinary(b []byte) error {
	var res GetInstancesInstanceIDNetworkInterfacesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
