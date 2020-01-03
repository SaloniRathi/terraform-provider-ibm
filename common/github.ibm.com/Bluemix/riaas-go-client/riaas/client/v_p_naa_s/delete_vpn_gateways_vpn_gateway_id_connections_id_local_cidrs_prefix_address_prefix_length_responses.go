// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthReader is a Reader for the DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLength structure.
type DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNoContent creates a DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNoContent with default headers values
func NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNoContent() *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNoContent {
	return &DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNoContent{}
}

/*DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNoContent handles this case with default header values.

The CIDR was successfully removed from the resource.
*/
type DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNoContent struct {
}

func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNoContent) Error() string {
	return fmt.Sprintf("[DELETE /vpn_gateways/{vpn_gateway_id}/connections/{id}/local_cidrs/{prefix_address}/{prefix_length}][%d] deleteVpnGatewaysVpnGatewayIdConnectionsIdLocalCidrsPrefixAddressPrefixLengthNoContent ", 204)
}

func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNotFound creates a DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNotFound with default headers values
func NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNotFound() *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNotFound {
	return &DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNotFound{}
}

/*DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNotFound handles this case with default header values.

The the specified CIDR does not exist on the specified resource or the specified resource could not be found.
*/
type DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNotFound) Error() string {
	return fmt.Sprintf("[DELETE /vpn_gateways/{vpn_gateway_id}/connections/{id}/local_cidrs/{prefix_address}/{prefix_length}][%d] deleteVpnGatewaysVpnGatewayIdConnectionsIdLocalCidrsPrefixAddressPrefixLengthNotFound  %+v", 404, o.Payload)
}

func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
