/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AddErrorLogAlertHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddErrorLogAlertHandlerRequest{}

// AddErrorLogAlertHandlerRequest struct for AddErrorLogAlertHandlerRequest
type AddErrorLogAlertHandlerRequest struct {
	// Name of the new Alert Handler
	HandlerName string                              `json:"handlerName"`
	Schemas     []EnumerrorLogAlertHandlerSchemaUrn `json:"schemas"`
	// A description for this Alert Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Alert Handler is enabled.
	Enabled bool `json:"enabled"`
	// Indicates whether the server should attempt to invoke this Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated.
	Asynchronous *bool `json:"asynchronous,omitempty"`
	// Specifies the alert severities for which this alert handler should be used. If no values are provided, then this alert handler will be enabled for alerts with any severity.
	EnabledAlertSeverity []EnumalertHandlerEnabledAlertSeverityProp `json:"enabledAlertSeverity,omitempty"`
	// Specifies the names of the alert types that are enabled for this alert handler.
	EnabledAlertType []EnumalertHandlerEnabledAlertTypeProp `json:"enabledAlertType,omitempty"`
	// Specifies the names of the alert types that are disabled for this alert handler.
	DisabledAlertType []EnumalertHandlerDisabledAlertTypeProp `json:"disabledAlertType,omitempty"`
}

// NewAddErrorLogAlertHandlerRequest instantiates a new AddErrorLogAlertHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddErrorLogAlertHandlerRequest(handlerName string, schemas []EnumerrorLogAlertHandlerSchemaUrn, enabled bool) *AddErrorLogAlertHandlerRequest {
	this := AddErrorLogAlertHandlerRequest{}
	this.HandlerName = handlerName
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewAddErrorLogAlertHandlerRequestWithDefaults instantiates a new AddErrorLogAlertHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddErrorLogAlertHandlerRequestWithDefaults() *AddErrorLogAlertHandlerRequest {
	this := AddErrorLogAlertHandlerRequest{}
	return &this
}

// GetHandlerName returns the HandlerName field value
func (o *AddErrorLogAlertHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddErrorLogAlertHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddErrorLogAlertHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddErrorLogAlertHandlerRequest) GetSchemas() []EnumerrorLogAlertHandlerSchemaUrn {
	if o == nil {
		var ret []EnumerrorLogAlertHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddErrorLogAlertHandlerRequest) GetSchemasOk() ([]EnumerrorLogAlertHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddErrorLogAlertHandlerRequest) SetSchemas(v []EnumerrorLogAlertHandlerSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddErrorLogAlertHandlerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogAlertHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddErrorLogAlertHandlerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddErrorLogAlertHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddErrorLogAlertHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddErrorLogAlertHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddErrorLogAlertHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *AddErrorLogAlertHandlerRequest) GetAsynchronous() bool {
	if o == nil || IsNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogAlertHandlerRequest) GetAsynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Asynchronous) {
		return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *AddErrorLogAlertHandlerRequest) HasAsynchronous() bool {
	if o != nil && !IsNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *AddErrorLogAlertHandlerRequest) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetEnabledAlertSeverity returns the EnabledAlertSeverity field value if set, zero value otherwise.
func (o *AddErrorLogAlertHandlerRequest) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp {
	if o == nil || IsNil(o.EnabledAlertSeverity) {
		var ret []EnumalertHandlerEnabledAlertSeverityProp
		return ret
	}
	return o.EnabledAlertSeverity
}

// GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogAlertHandlerRequest) GetEnabledAlertSeverityOk() ([]EnumalertHandlerEnabledAlertSeverityProp, bool) {
	if o == nil || IsNil(o.EnabledAlertSeverity) {
		return nil, false
	}
	return o.EnabledAlertSeverity, true
}

// HasEnabledAlertSeverity returns a boolean if a field has been set.
func (o *AddErrorLogAlertHandlerRequest) HasEnabledAlertSeverity() bool {
	if o != nil && !IsNil(o.EnabledAlertSeverity) {
		return true
	}

	return false
}

// SetEnabledAlertSeverity gets a reference to the given []EnumalertHandlerEnabledAlertSeverityProp and assigns it to the EnabledAlertSeverity field.
func (o *AddErrorLogAlertHandlerRequest) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp) {
	o.EnabledAlertSeverity = v
}

// GetEnabledAlertType returns the EnabledAlertType field value if set, zero value otherwise.
func (o *AddErrorLogAlertHandlerRequest) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp {
	if o == nil || IsNil(o.EnabledAlertType) {
		var ret []EnumalertHandlerEnabledAlertTypeProp
		return ret
	}
	return o.EnabledAlertType
}

// GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogAlertHandlerRequest) GetEnabledAlertTypeOk() ([]EnumalertHandlerEnabledAlertTypeProp, bool) {
	if o == nil || IsNil(o.EnabledAlertType) {
		return nil, false
	}
	return o.EnabledAlertType, true
}

// HasEnabledAlertType returns a boolean if a field has been set.
func (o *AddErrorLogAlertHandlerRequest) HasEnabledAlertType() bool {
	if o != nil && !IsNil(o.EnabledAlertType) {
		return true
	}

	return false
}

// SetEnabledAlertType gets a reference to the given []EnumalertHandlerEnabledAlertTypeProp and assigns it to the EnabledAlertType field.
func (o *AddErrorLogAlertHandlerRequest) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp) {
	o.EnabledAlertType = v
}

// GetDisabledAlertType returns the DisabledAlertType field value if set, zero value otherwise.
func (o *AddErrorLogAlertHandlerRequest) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp {
	if o == nil || IsNil(o.DisabledAlertType) {
		var ret []EnumalertHandlerDisabledAlertTypeProp
		return ret
	}
	return o.DisabledAlertType
}

// GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogAlertHandlerRequest) GetDisabledAlertTypeOk() ([]EnumalertHandlerDisabledAlertTypeProp, bool) {
	if o == nil || IsNil(o.DisabledAlertType) {
		return nil, false
	}
	return o.DisabledAlertType, true
}

// HasDisabledAlertType returns a boolean if a field has been set.
func (o *AddErrorLogAlertHandlerRequest) HasDisabledAlertType() bool {
	if o != nil && !IsNil(o.DisabledAlertType) {
		return true
	}

	return false
}

// SetDisabledAlertType gets a reference to the given []EnumalertHandlerDisabledAlertTypeProp and assigns it to the DisabledAlertType field.
func (o *AddErrorLogAlertHandlerRequest) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp) {
	o.DisabledAlertType = v
}

func (o AddErrorLogAlertHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddErrorLogAlertHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["handlerName"] = o.HandlerName
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Asynchronous) {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	if !IsNil(o.EnabledAlertSeverity) {
		toSerialize["enabledAlertSeverity"] = o.EnabledAlertSeverity
	}
	if !IsNil(o.EnabledAlertType) {
		toSerialize["enabledAlertType"] = o.EnabledAlertType
	}
	if !IsNil(o.DisabledAlertType) {
		toSerialize["disabledAlertType"] = o.DisabledAlertType
	}
	return toSerialize, nil
}

type NullableAddErrorLogAlertHandlerRequest struct {
	value *AddErrorLogAlertHandlerRequest
	isSet bool
}

func (v NullableAddErrorLogAlertHandlerRequest) Get() *AddErrorLogAlertHandlerRequest {
	return v.value
}

func (v *NullableAddErrorLogAlertHandlerRequest) Set(val *AddErrorLogAlertHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddErrorLogAlertHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddErrorLogAlertHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddErrorLogAlertHandlerRequest(val *AddErrorLogAlertHandlerRequest) *NullableAddErrorLogAlertHandlerRequest {
	return &NullableAddErrorLogAlertHandlerRequest{value: val, isSet: true}
}

func (v NullableAddErrorLogAlertHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddErrorLogAlertHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
