// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PostInstancesInstanceIDVolumeAttachmentsReader is a Reader for the PostInstancesInstanceIDVolumeAttachments structure.
type PostInstancesInstanceIDVolumeAttachmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInstancesInstanceIDVolumeAttachmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostInstancesInstanceIDVolumeAttachmentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostInstancesInstanceIDVolumeAttachmentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostInstancesInstanceIDVolumeAttachmentsCreated creates a PostInstancesInstanceIDVolumeAttachmentsCreated with default headers values
func NewPostInstancesInstanceIDVolumeAttachmentsCreated() *PostInstancesInstanceIDVolumeAttachmentsCreated {
	return &PostInstancesInstanceIDVolumeAttachmentsCreated{}
}

/*PostInstancesInstanceIDVolumeAttachmentsCreated handles this case with default header values.

dummy
*/
type PostInstancesInstanceIDVolumeAttachmentsCreated struct {
	Payload *models.InstanceVolumeAttachment
}

func (o *PostInstancesInstanceIDVolumeAttachmentsCreated) Error() string {
	return fmt.Sprintf("[POST /instances/{instance_id}/volume_attachments][%d] postInstancesInstanceIdVolumeAttachmentsCreated  %+v", 201, o.Payload)
}

func (o *PostInstancesInstanceIDVolumeAttachmentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstanceVolumeAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInstancesInstanceIDVolumeAttachmentsBadRequest creates a PostInstancesInstanceIDVolumeAttachmentsBadRequest with default headers values
func NewPostInstancesInstanceIDVolumeAttachmentsBadRequest() *PostInstancesInstanceIDVolumeAttachmentsBadRequest {
	return &PostInstancesInstanceIDVolumeAttachmentsBadRequest{}
}

/*PostInstancesInstanceIDVolumeAttachmentsBadRequest handles this case with default header values.

error
*/
type PostInstancesInstanceIDVolumeAttachmentsBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PostInstancesInstanceIDVolumeAttachmentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /instances/{instance_id}/volume_attachments][%d] postInstancesInstanceIdVolumeAttachmentsBadRequest  %+v", 400, o.Payload)
}

func (o *PostInstancesInstanceIDVolumeAttachmentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostInstancesInstanceIDVolumeAttachmentsBody VolumeAttachmentTemplate
swagger:model PostInstancesInstanceIDVolumeAttachmentsBody
*/
type PostInstancesInstanceIDVolumeAttachmentsBody struct {

	// If set to true, when deleting the instance the volume will also be deleted
	DeleteVolumeOnInstanceDelete bool `json:"delete_volume_on_instance_delete,omitempty"`

	// The user-defined name for this subnet
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// volume
	Volume *PostInstancesInstanceIDVolumeAttachmentsParamsBodyVolume `json:"volume,omitempty"`
}

// Validate validates this post instances instance ID volume attachments body
func (o *PostInstancesInstanceIDVolumeAttachmentsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostInstancesInstanceIDVolumeAttachmentsBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"name", "body", string(o.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (o *PostInstancesInstanceIDVolumeAttachmentsBody) validateVolume(formats strfmt.Registry) error {

	if swag.IsZero(o.Volume) { // not required
		return nil
	}

	if o.Volume != nil {
		if err := o.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostInstancesInstanceIDVolumeAttachmentsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostInstancesInstanceIDVolumeAttachmentsBody) UnmarshalBinary(b []byte) error {
	var res PostInstancesInstanceIDVolumeAttachmentsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostInstancesInstanceIDVolumeAttachmentsParamsBodyVolume VolumeIdentity
//
// The identity of the volume to attach to the instance
swagger:model PostInstancesInstanceIDVolumeAttachmentsParamsBodyVolume
*/
type PostInstancesInstanceIDVolumeAttachmentsParamsBodyVolume struct {

	// The CRN for this volume
	Crn string `json:"crn,omitempty"`

	// The unique identifier for this volume
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The user-defined name for this volume
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this post instances instance ID volume attachments params body volume
func (o *PostInstancesInstanceIDVolumeAttachmentsParamsBodyVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostInstancesInstanceIDVolumeAttachmentsParamsBodyVolume) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"volume"+"."+"id", "body", "uuid", o.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PostInstancesInstanceIDVolumeAttachmentsParamsBodyVolume) validateName(formats strfmt.Registry) error {

	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"volume"+"."+"name", "body", string(o.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostInstancesInstanceIDVolumeAttachmentsParamsBodyVolume) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostInstancesInstanceIDVolumeAttachmentsParamsBodyVolume) UnmarshalBinary(b []byte) error {
	var res PostInstancesInstanceIDVolumeAttachmentsParamsBodyVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
