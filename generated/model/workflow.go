// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: 2023 Siemens AG
//
// SPDX-License-Identifier: Apache-2.0
//

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Workflow workflow
//
// swagger:model Workflow
type Workflow struct {

	// Description of the workflow
	// Example: This is a workflow
	// Max Length: 1024
	Description string `json:"description,omitempty"`

	// groups
	// Max Items: 1024
	Groups []*Group `json:"groups"`

	// User provided unique workflow name
	// Example: wfx.workflow.dau.direct
	// Required: true
	// Max Length: 64
	// Min Length: 1
	// Pattern: ^[a-zA-Z0-9\-\.]+$
	Name string `json:"name"`

	// states
	// Required: true
	// Max Items: 4096
	// Min Items: 1
	States []*State `json:"states"`

	// transitions
	// Required: true
	// Max Items: 16384
	// Min Items: 1
	Transitions []*Transition `json:"transitions"`
}

// Validate validates this workflow
func (m *Workflow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransitions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Workflow) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 1024); err != nil {
		return err
	}

	return nil
}

func (m *Workflow) validateGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	iGroupsSize := int64(len(m.Groups))

	if err := validate.MaxItems("groups", "body", iGroupsSize, 1024); err != nil {
		return err
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Workflow) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 64); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z0-9\-\.]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Workflow) validateStates(formats strfmt.Registry) error {

	if err := validate.Required("states", "body", m.States); err != nil {
		return err
	}

	iStatesSize := int64(len(m.States))

	if err := validate.MinItems("states", "body", iStatesSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("states", "body", iStatesSize, 4096); err != nil {
		return err
	}

	for i := 0; i < len(m.States); i++ {
		if swag.IsZero(m.States[i]) { // not required
			continue
		}

		if m.States[i] != nil {
			if err := m.States[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("states" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("states" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Workflow) validateTransitions(formats strfmt.Registry) error {

	if err := validate.Required("transitions", "body", m.Transitions); err != nil {
		return err
	}

	iTransitionsSize := int64(len(m.Transitions))

	if err := validate.MinItems("transitions", "body", iTransitionsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("transitions", "body", iTransitionsSize, 16384); err != nil {
		return err
	}

	for i := 0; i < len(m.Transitions); i++ {
		if swag.IsZero(m.Transitions[i]) { // not required
			continue
		}

		if m.Transitions[i] != nil {
			if err := m.Transitions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transitions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this workflow based on the context it is used
func (m *Workflow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransitions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Workflow) contextValidateGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Groups); i++ {

		if m.Groups[i] != nil {

			if swag.IsZero(m.Groups[i]) { // not required
				return nil
			}

			if err := m.Groups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Workflow) contextValidateStates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.States); i++ {

		if m.States[i] != nil {

			if swag.IsZero(m.States[i]) { // not required
				return nil
			}

			if err := m.States[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("states" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("states" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Workflow) contextValidateTransitions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Transitions); i++ {

		if m.Transitions[i] != nil {

			if swag.IsZero(m.Transitions[i]) { // not required
				return nil
			}

			if err := m.Transitions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transitions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Workflow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Workflow) UnmarshalBinary(b []byte) error {
	var res Workflow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
