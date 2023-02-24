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

// checks if the AddExecAlertHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddExecAlertHandlerRequest{}

// AddExecAlertHandlerRequest struct for AddExecAlertHandlerRequest
type AddExecAlertHandlerRequest struct {
	// Name of the new Alert Handler
	HandlerName string                          `json:"handlerName"`
	Schemas     []EnumexecAlertHandlerSchemaUrn `json:"schemas"`
	// Specifies the path of the command to execute, without any arguments. It must be an absolute path for reasons of security and reliability.
	Command string `json:"command"`
	// Indicates whether the server should attempt to invoke this Exec Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated.
	Asynchronous *bool `json:"asynchronous,omitempty"`
	// A description for this Alert Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Alert Handler is enabled.
	Enabled              bool                                       `json:"enabled"`
	EnabledAlertSeverity []EnumalertHandlerEnabledAlertSeverityProp `json:"enabledAlertSeverity,omitempty"`
	EnabledAlertType     []EnumalertHandlerEnabledAlertTypeProp     `json:"enabledAlertType,omitempty"`
	DisabledAlertType    []EnumalertHandlerDisabledAlertTypeProp    `json:"disabledAlertType,omitempty"`
}

// NewAddExecAlertHandlerRequest instantiates a new AddExecAlertHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddExecAlertHandlerRequest(handlerName string, schemas []EnumexecAlertHandlerSchemaUrn, command string, enabled bool) *AddExecAlertHandlerRequest {
	this := AddExecAlertHandlerRequest{}
	this.HandlerName = handlerName
	this.Schemas = schemas
	this.Command = command
	this.Enabled = enabled
	return &this
}

// NewAddExecAlertHandlerRequestWithDefaults instantiates a new AddExecAlertHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddExecAlertHandlerRequestWithDefaults() *AddExecAlertHandlerRequest {
	this := AddExecAlertHandlerRequest{}
	return &this
}

// GetHandlerName returns the HandlerName field value
func (o *AddExecAlertHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddExecAlertHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddExecAlertHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddExecAlertHandlerRequest) GetSchemas() []EnumexecAlertHandlerSchemaUrn {
	if o == nil {
		var ret []EnumexecAlertHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddExecAlertHandlerRequest) GetSchemasOk() ([]EnumexecAlertHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddExecAlertHandlerRequest) SetSchemas(v []EnumexecAlertHandlerSchemaUrn) {
	o.Schemas = v
}

// GetCommand returns the Command field value
func (o *AddExecAlertHandlerRequest) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *AddExecAlertHandlerRequest) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *AddExecAlertHandlerRequest) SetCommand(v string) {
	o.Command = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *AddExecAlertHandlerRequest) GetAsynchronous() bool {
	if o == nil || IsNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddExecAlertHandlerRequest) GetAsynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Asynchronous) {
		return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *AddExecAlertHandlerRequest) HasAsynchronous() bool {
	if o != nil && !IsNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *AddExecAlertHandlerRequest) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddExecAlertHandlerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddExecAlertHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddExecAlertHandlerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddExecAlertHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddExecAlertHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddExecAlertHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddExecAlertHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnabledAlertSeverity returns the EnabledAlertSeverity field value if set, zero value otherwise.
func (o *AddExecAlertHandlerRequest) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp {
	if o == nil || IsNil(o.EnabledAlertSeverity) {
		var ret []EnumalertHandlerEnabledAlertSeverityProp
		return ret
	}
	return o.EnabledAlertSeverity
}

// GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddExecAlertHandlerRequest) GetEnabledAlertSeverityOk() ([]EnumalertHandlerEnabledAlertSeverityProp, bool) {
	if o == nil || IsNil(o.EnabledAlertSeverity) {
		return nil, false
	}
	return o.EnabledAlertSeverity, true
}

// HasEnabledAlertSeverity returns a boolean if a field has been set.
func (o *AddExecAlertHandlerRequest) HasEnabledAlertSeverity() bool {
	if o != nil && !IsNil(o.EnabledAlertSeverity) {
		return true
	}

	return false
}

// SetEnabledAlertSeverity gets a reference to the given []EnumalertHandlerEnabledAlertSeverityProp and assigns it to the EnabledAlertSeverity field.
func (o *AddExecAlertHandlerRequest) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp) {
	o.EnabledAlertSeverity = v
}

// GetEnabledAlertType returns the EnabledAlertType field value if set, zero value otherwise.
func (o *AddExecAlertHandlerRequest) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp {
	if o == nil || IsNil(o.EnabledAlertType) {
		var ret []EnumalertHandlerEnabledAlertTypeProp
		return ret
	}
	return o.EnabledAlertType
}

// GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddExecAlertHandlerRequest) GetEnabledAlertTypeOk() ([]EnumalertHandlerEnabledAlertTypeProp, bool) {
	if o == nil || IsNil(o.EnabledAlertType) {
		return nil, false
	}
	return o.EnabledAlertType, true
}

// HasEnabledAlertType returns a boolean if a field has been set.
func (o *AddExecAlertHandlerRequest) HasEnabledAlertType() bool {
	if o != nil && !IsNil(o.EnabledAlertType) {
		return true
	}

	return false
}

// SetEnabledAlertType gets a reference to the given []EnumalertHandlerEnabledAlertTypeProp and assigns it to the EnabledAlertType field.
func (o *AddExecAlertHandlerRequest) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp) {
	o.EnabledAlertType = v
}

// GetDisabledAlertType returns the DisabledAlertType field value if set, zero value otherwise.
func (o *AddExecAlertHandlerRequest) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp {
	if o == nil || IsNil(o.DisabledAlertType) {
		var ret []EnumalertHandlerDisabledAlertTypeProp
		return ret
	}
	return o.DisabledAlertType
}

// GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddExecAlertHandlerRequest) GetDisabledAlertTypeOk() ([]EnumalertHandlerDisabledAlertTypeProp, bool) {
	if o == nil || IsNil(o.DisabledAlertType) {
		return nil, false
	}
	return o.DisabledAlertType, true
}

// HasDisabledAlertType returns a boolean if a field has been set.
func (o *AddExecAlertHandlerRequest) HasDisabledAlertType() bool {
	if o != nil && !IsNil(o.DisabledAlertType) {
		return true
	}

	return false
}

// SetDisabledAlertType gets a reference to the given []EnumalertHandlerDisabledAlertTypeProp and assigns it to the DisabledAlertType field.
func (o *AddExecAlertHandlerRequest) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp) {
	o.DisabledAlertType = v
}

func (o AddExecAlertHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddExecAlertHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["handlerName"] = o.HandlerName
	toSerialize["schemas"] = o.Schemas
	toSerialize["command"] = o.Command
	if !IsNil(o.Asynchronous) {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
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

type NullableAddExecAlertHandlerRequest struct {
	value *AddExecAlertHandlerRequest
	isSet bool
}

func (v NullableAddExecAlertHandlerRequest) Get() *AddExecAlertHandlerRequest {
	return v.value
}

func (v *NullableAddExecAlertHandlerRequest) Set(val *AddExecAlertHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddExecAlertHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddExecAlertHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddExecAlertHandlerRequest(val *AddExecAlertHandlerRequest) *NullableAddExecAlertHandlerRequest {
	return &NullableAddExecAlertHandlerRequest{value: val, isSet: true}
}

func (v NullableAddExecAlertHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddExecAlertHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
