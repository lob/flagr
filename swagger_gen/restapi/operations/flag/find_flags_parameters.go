// Code generated by go-swagger; DO NOT EDIT.

package flag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindFlagsParams creates a new FindFlagsParams object
// no default values defined in spec.
func NewFindFlagsParams() FindFlagsParams {

	return FindFlagsParams{}
}

// FindFlagsParams contains all the bound params for the find flags operation
// typically these are obtained from a http.Request
//
// swagger:parameters findFlags
type FindFlagsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*return flags exactly matching given description
	  In: query
	*/
	Description *string
	/*return flags partially matching given description
	  In: query
	*/
	DescriptionLike *string
	/*return flags having given enabled status
	  In: query
	*/
	Enabled *bool
	/*return flags matching given key
	  In: query
	*/
	Key *string
	/*the numbers of flags to return
	  In: query
	*/
	Limit *int64
	/*return flags given the offset, it should usually set together with limit
	  In: query
	*/
	Offset *int64
	/*return flags with preloaded segments and variants
	  In: query
	*/
	Preload *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewFindFlagsParams() beforehand.
func (o *FindFlagsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDescription, qhkDescription, _ := qs.GetOK("description")
	if err := o.bindDescription(qDescription, qhkDescription, route.Formats); err != nil {
		res = append(res, err)
	}

	qDescriptionLike, qhkDescriptionLike, _ := qs.GetOK("description_like")
	if err := o.bindDescriptionLike(qDescriptionLike, qhkDescriptionLike, route.Formats); err != nil {
		res = append(res, err)
	}

	qEnabled, qhkEnabled, _ := qs.GetOK("enabled")
	if err := o.bindEnabled(qEnabled, qhkEnabled, route.Formats); err != nil {
		res = append(res, err)
	}

	qKey, qhkKey, _ := qs.GetOK("key")
	if err := o.bindKey(qKey, qhkKey, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qOffset, qhkOffset, _ := qs.GetOK("offset")
	if err := o.bindOffset(qOffset, qhkOffset, route.Formats); err != nil {
		res = append(res, err)
	}

	qPreload, qhkPreload, _ := qs.GetOK("preload")
	if err := o.bindPreload(qPreload, qhkPreload, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDescription binds and validates parameter Description from query.
func (o *FindFlagsParams) bindDescription(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Description = &raw

	return nil
}

// bindDescriptionLike binds and validates parameter DescriptionLike from query.
func (o *FindFlagsParams) bindDescriptionLike(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.DescriptionLike = &raw

	return nil
}

// bindEnabled binds and validates parameter Enabled from query.
func (o *FindFlagsParams) bindEnabled(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("enabled", "query", "bool", raw)
	}
	o.Enabled = &value

	return nil
}

// bindKey binds and validates parameter Key from query.
func (o *FindFlagsParams) bindKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Key = &raw

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *FindFlagsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}

// bindOffset binds and validates parameter Offset from query.
func (o *FindFlagsParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("offset", "query", "int64", raw)
	}
	o.Offset = &value

	return nil
}

// bindPreload binds and validates parameter Preload from query.
func (o *FindFlagsParams) bindPreload(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("preload", "query", "bool", raw)
	}
	o.Preload = &value

	return nil
}
