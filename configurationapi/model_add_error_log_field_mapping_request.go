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

// AddErrorLogFieldMappingRequest struct for AddErrorLogFieldMappingRequest
type AddErrorLogFieldMappingRequest struct {
	// Name of the new Log Field Mapping
	MappingName string                              `json:"mappingName"`
	Schemas     []EnumerrorLogFieldMappingSchemaUrn `json:"schemas"`
	// The time that the log message was generated.
	LogFieldTimestamp *string `json:"logFieldTimestamp,omitempty"`
	// The name for this Directory Server product, which may be used to identify which product was used to log the message if multiple products log to the same database table.
	LogFieldProductName *string `json:"logFieldProductName,omitempty"`
	// A name that uniquely identifies this Directory Server instance, which may be used to identify which instance was used to log the message if multiple server instances log to the same database table.
	LogFieldInstanceName *string `json:"logFieldInstanceName,omitempty"`
	// The startup ID for the Directory Server. A different value will be generated each time the server is started.
	LogFieldStartupid *string `json:"logFieldStartupid,omitempty"`
	// The category for the log message.
	LogFieldCategory *string `json:"logFieldCategory,omitempty"`
	// The severity for the log message.
	LogFieldSeverity *string `json:"logFieldSeverity,omitempty"`
	// The numeric value which uniquely identifies the type of message.
	LogFieldMessageID *string `json:"logFieldMessageID,omitempty"`
	// The text of the log message.
	LogFieldMessage *string `json:"logFieldMessage,omitempty"`
	// A description for this Log Field Mapping
	Description *string `json:"description,omitempty"`
}

// NewAddErrorLogFieldMappingRequest instantiates a new AddErrorLogFieldMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddErrorLogFieldMappingRequest(mappingName string, schemas []EnumerrorLogFieldMappingSchemaUrn) *AddErrorLogFieldMappingRequest {
	this := AddErrorLogFieldMappingRequest{}
	this.MappingName = mappingName
	this.Schemas = schemas
	return &this
}

// NewAddErrorLogFieldMappingRequestWithDefaults instantiates a new AddErrorLogFieldMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddErrorLogFieldMappingRequestWithDefaults() *AddErrorLogFieldMappingRequest {
	this := AddErrorLogFieldMappingRequest{}
	return &this
}

// GetMappingName returns the MappingName field value
func (o *AddErrorLogFieldMappingRequest) GetMappingName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MappingName
}

// GetMappingNameOk returns a tuple with the MappingName field value
// and a boolean to check if the value has been set.
func (o *AddErrorLogFieldMappingRequest) GetMappingNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MappingName, true
}

// SetMappingName sets field value
func (o *AddErrorLogFieldMappingRequest) SetMappingName(v string) {
	o.MappingName = v
}

// GetSchemas returns the Schemas field value
func (o *AddErrorLogFieldMappingRequest) GetSchemas() []EnumerrorLogFieldMappingSchemaUrn {
	if o == nil {
		var ret []EnumerrorLogFieldMappingSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddErrorLogFieldMappingRequest) GetSchemasOk() ([]EnumerrorLogFieldMappingSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddErrorLogFieldMappingRequest) SetSchemas(v []EnumerrorLogFieldMappingSchemaUrn) {
	o.Schemas = v
}

// GetLogFieldTimestamp returns the LogFieldTimestamp field value if set, zero value otherwise.
func (o *AddErrorLogFieldMappingRequest) GetLogFieldTimestamp() string {
	if o == nil || isNil(o.LogFieldTimestamp) {
		var ret string
		return ret
	}
	return *o.LogFieldTimestamp
}

// GetLogFieldTimestampOk returns a tuple with the LogFieldTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogFieldMappingRequest) GetLogFieldTimestampOk() (*string, bool) {
	if o == nil || isNil(o.LogFieldTimestamp) {
		return nil, false
	}
	return o.LogFieldTimestamp, true
}

// HasLogFieldTimestamp returns a boolean if a field has been set.
func (o *AddErrorLogFieldMappingRequest) HasLogFieldTimestamp() bool {
	if o != nil && !isNil(o.LogFieldTimestamp) {
		return true
	}

	return false
}

// SetLogFieldTimestamp gets a reference to the given string and assigns it to the LogFieldTimestamp field.
func (o *AddErrorLogFieldMappingRequest) SetLogFieldTimestamp(v string) {
	o.LogFieldTimestamp = &v
}

// GetLogFieldProductName returns the LogFieldProductName field value if set, zero value otherwise.
func (o *AddErrorLogFieldMappingRequest) GetLogFieldProductName() string {
	if o == nil || isNil(o.LogFieldProductName) {
		var ret string
		return ret
	}
	return *o.LogFieldProductName
}

// GetLogFieldProductNameOk returns a tuple with the LogFieldProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogFieldMappingRequest) GetLogFieldProductNameOk() (*string, bool) {
	if o == nil || isNil(o.LogFieldProductName) {
		return nil, false
	}
	return o.LogFieldProductName, true
}

// HasLogFieldProductName returns a boolean if a field has been set.
func (o *AddErrorLogFieldMappingRequest) HasLogFieldProductName() bool {
	if o != nil && !isNil(o.LogFieldProductName) {
		return true
	}

	return false
}

// SetLogFieldProductName gets a reference to the given string and assigns it to the LogFieldProductName field.
func (o *AddErrorLogFieldMappingRequest) SetLogFieldProductName(v string) {
	o.LogFieldProductName = &v
}

// GetLogFieldInstanceName returns the LogFieldInstanceName field value if set, zero value otherwise.
func (o *AddErrorLogFieldMappingRequest) GetLogFieldInstanceName() string {
	if o == nil || isNil(o.LogFieldInstanceName) {
		var ret string
		return ret
	}
	return *o.LogFieldInstanceName
}

// GetLogFieldInstanceNameOk returns a tuple with the LogFieldInstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogFieldMappingRequest) GetLogFieldInstanceNameOk() (*string, bool) {
	if o == nil || isNil(o.LogFieldInstanceName) {
		return nil, false
	}
	return o.LogFieldInstanceName, true
}

// HasLogFieldInstanceName returns a boolean if a field has been set.
func (o *AddErrorLogFieldMappingRequest) HasLogFieldInstanceName() bool {
	if o != nil && !isNil(o.LogFieldInstanceName) {
		return true
	}

	return false
}

// SetLogFieldInstanceName gets a reference to the given string and assigns it to the LogFieldInstanceName field.
func (o *AddErrorLogFieldMappingRequest) SetLogFieldInstanceName(v string) {
	o.LogFieldInstanceName = &v
}

// GetLogFieldStartupid returns the LogFieldStartupid field value if set, zero value otherwise.
func (o *AddErrorLogFieldMappingRequest) GetLogFieldStartupid() string {
	if o == nil || isNil(o.LogFieldStartupid) {
		var ret string
		return ret
	}
	return *o.LogFieldStartupid
}

// GetLogFieldStartupidOk returns a tuple with the LogFieldStartupid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogFieldMappingRequest) GetLogFieldStartupidOk() (*string, bool) {
	if o == nil || isNil(o.LogFieldStartupid) {
		return nil, false
	}
	return o.LogFieldStartupid, true
}

// HasLogFieldStartupid returns a boolean if a field has been set.
func (o *AddErrorLogFieldMappingRequest) HasLogFieldStartupid() bool {
	if o != nil && !isNil(o.LogFieldStartupid) {
		return true
	}

	return false
}

// SetLogFieldStartupid gets a reference to the given string and assigns it to the LogFieldStartupid field.
func (o *AddErrorLogFieldMappingRequest) SetLogFieldStartupid(v string) {
	o.LogFieldStartupid = &v
}

// GetLogFieldCategory returns the LogFieldCategory field value if set, zero value otherwise.
func (o *AddErrorLogFieldMappingRequest) GetLogFieldCategory() string {
	if o == nil || isNil(o.LogFieldCategory) {
		var ret string
		return ret
	}
	return *o.LogFieldCategory
}

// GetLogFieldCategoryOk returns a tuple with the LogFieldCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogFieldMappingRequest) GetLogFieldCategoryOk() (*string, bool) {
	if o == nil || isNil(o.LogFieldCategory) {
		return nil, false
	}
	return o.LogFieldCategory, true
}

// HasLogFieldCategory returns a boolean if a field has been set.
func (o *AddErrorLogFieldMappingRequest) HasLogFieldCategory() bool {
	if o != nil && !isNil(o.LogFieldCategory) {
		return true
	}

	return false
}

// SetLogFieldCategory gets a reference to the given string and assigns it to the LogFieldCategory field.
func (o *AddErrorLogFieldMappingRequest) SetLogFieldCategory(v string) {
	o.LogFieldCategory = &v
}

// GetLogFieldSeverity returns the LogFieldSeverity field value if set, zero value otherwise.
func (o *AddErrorLogFieldMappingRequest) GetLogFieldSeverity() string {
	if o == nil || isNil(o.LogFieldSeverity) {
		var ret string
		return ret
	}
	return *o.LogFieldSeverity
}

// GetLogFieldSeverityOk returns a tuple with the LogFieldSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogFieldMappingRequest) GetLogFieldSeverityOk() (*string, bool) {
	if o == nil || isNil(o.LogFieldSeverity) {
		return nil, false
	}
	return o.LogFieldSeverity, true
}

// HasLogFieldSeverity returns a boolean if a field has been set.
func (o *AddErrorLogFieldMappingRequest) HasLogFieldSeverity() bool {
	if o != nil && !isNil(o.LogFieldSeverity) {
		return true
	}

	return false
}

// SetLogFieldSeverity gets a reference to the given string and assigns it to the LogFieldSeverity field.
func (o *AddErrorLogFieldMappingRequest) SetLogFieldSeverity(v string) {
	o.LogFieldSeverity = &v
}

// GetLogFieldMessageID returns the LogFieldMessageID field value if set, zero value otherwise.
func (o *AddErrorLogFieldMappingRequest) GetLogFieldMessageID() string {
	if o == nil || isNil(o.LogFieldMessageID) {
		var ret string
		return ret
	}
	return *o.LogFieldMessageID
}

// GetLogFieldMessageIDOk returns a tuple with the LogFieldMessageID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogFieldMappingRequest) GetLogFieldMessageIDOk() (*string, bool) {
	if o == nil || isNil(o.LogFieldMessageID) {
		return nil, false
	}
	return o.LogFieldMessageID, true
}

// HasLogFieldMessageID returns a boolean if a field has been set.
func (o *AddErrorLogFieldMappingRequest) HasLogFieldMessageID() bool {
	if o != nil && !isNil(o.LogFieldMessageID) {
		return true
	}

	return false
}

// SetLogFieldMessageID gets a reference to the given string and assigns it to the LogFieldMessageID field.
func (o *AddErrorLogFieldMappingRequest) SetLogFieldMessageID(v string) {
	o.LogFieldMessageID = &v
}

// GetLogFieldMessage returns the LogFieldMessage field value if set, zero value otherwise.
func (o *AddErrorLogFieldMappingRequest) GetLogFieldMessage() string {
	if o == nil || isNil(o.LogFieldMessage) {
		var ret string
		return ret
	}
	return *o.LogFieldMessage
}

// GetLogFieldMessageOk returns a tuple with the LogFieldMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogFieldMappingRequest) GetLogFieldMessageOk() (*string, bool) {
	if o == nil || isNil(o.LogFieldMessage) {
		return nil, false
	}
	return o.LogFieldMessage, true
}

// HasLogFieldMessage returns a boolean if a field has been set.
func (o *AddErrorLogFieldMappingRequest) HasLogFieldMessage() bool {
	if o != nil && !isNil(o.LogFieldMessage) {
		return true
	}

	return false
}

// SetLogFieldMessage gets a reference to the given string and assigns it to the LogFieldMessage field.
func (o *AddErrorLogFieldMappingRequest) SetLogFieldMessage(v string) {
	o.LogFieldMessage = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddErrorLogFieldMappingRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddErrorLogFieldMappingRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddErrorLogFieldMappingRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddErrorLogFieldMappingRequest) SetDescription(v string) {
	o.Description = &v
}

func (o AddErrorLogFieldMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mappingName"] = o.MappingName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.LogFieldTimestamp) {
		toSerialize["logFieldTimestamp"] = o.LogFieldTimestamp
	}
	if !isNil(o.LogFieldProductName) {
		toSerialize["logFieldProductName"] = o.LogFieldProductName
	}
	if !isNil(o.LogFieldInstanceName) {
		toSerialize["logFieldInstanceName"] = o.LogFieldInstanceName
	}
	if !isNil(o.LogFieldStartupid) {
		toSerialize["logFieldStartupid"] = o.LogFieldStartupid
	}
	if !isNil(o.LogFieldCategory) {
		toSerialize["logFieldCategory"] = o.LogFieldCategory
	}
	if !isNil(o.LogFieldSeverity) {
		toSerialize["logFieldSeverity"] = o.LogFieldSeverity
	}
	if !isNil(o.LogFieldMessageID) {
		toSerialize["logFieldMessageID"] = o.LogFieldMessageID
	}
	if !isNil(o.LogFieldMessage) {
		toSerialize["logFieldMessage"] = o.LogFieldMessage
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAddErrorLogFieldMappingRequest struct {
	value *AddErrorLogFieldMappingRequest
	isSet bool
}

func (v NullableAddErrorLogFieldMappingRequest) Get() *AddErrorLogFieldMappingRequest {
	return v.value
}

func (v *NullableAddErrorLogFieldMappingRequest) Set(val *AddErrorLogFieldMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddErrorLogFieldMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddErrorLogFieldMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddErrorLogFieldMappingRequest(val *AddErrorLogFieldMappingRequest) *NullableAddErrorLogFieldMappingRequest {
	return &NullableAddErrorLogFieldMappingRequest{value: val, isSet: true}
}

func (v NullableAddErrorLogFieldMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddErrorLogFieldMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}