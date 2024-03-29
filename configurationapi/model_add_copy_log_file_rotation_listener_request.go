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

// checks if the AddCopyLogFileRotationListenerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCopyLogFileRotationListenerRequest{}

// AddCopyLogFileRotationListenerRequest struct for AddCopyLogFileRotationListenerRequest
type AddCopyLogFileRotationListenerRequest struct {
	Schemas []EnumcopyLogFileRotationListenerSchemaUrn `json:"schemas"`
	// The path to the directory to which log files should be copied. It must be different from the directory to which the log file is originally written, and administrators should ensure that the filesystem has sufficient space to hold files as they are copied.
	CopyToDirectory string `json:"copyToDirectory"`
	// Indicates whether the file should be gzip-compressed as it is copied into the destination directory.
	CompressOnCopy *bool `json:"compressOnCopy,omitempty"`
	// A description for this Log File Rotation Listener
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log File Rotation Listener is enabled for use.
	Enabled bool `json:"enabled"`
	// Name of the new Log File Rotation Listener
	ListenerName string `json:"listenerName"`
}

// NewAddCopyLogFileRotationListenerRequest instantiates a new AddCopyLogFileRotationListenerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCopyLogFileRotationListenerRequest(schemas []EnumcopyLogFileRotationListenerSchemaUrn, copyToDirectory string, enabled bool, listenerName string) *AddCopyLogFileRotationListenerRequest {
	this := AddCopyLogFileRotationListenerRequest{}
	this.Schemas = schemas
	this.CopyToDirectory = copyToDirectory
	this.Enabled = enabled
	this.ListenerName = listenerName
	return &this
}

// NewAddCopyLogFileRotationListenerRequestWithDefaults instantiates a new AddCopyLogFileRotationListenerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCopyLogFileRotationListenerRequestWithDefaults() *AddCopyLogFileRotationListenerRequest {
	this := AddCopyLogFileRotationListenerRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddCopyLogFileRotationListenerRequest) GetSchemas() []EnumcopyLogFileRotationListenerSchemaUrn {
	if o == nil {
		var ret []EnumcopyLogFileRotationListenerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddCopyLogFileRotationListenerRequest) GetSchemasOk() ([]EnumcopyLogFileRotationListenerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddCopyLogFileRotationListenerRequest) SetSchemas(v []EnumcopyLogFileRotationListenerSchemaUrn) {
	o.Schemas = v
}

// GetCopyToDirectory returns the CopyToDirectory field value
func (o *AddCopyLogFileRotationListenerRequest) GetCopyToDirectory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CopyToDirectory
}

// GetCopyToDirectoryOk returns a tuple with the CopyToDirectory field value
// and a boolean to check if the value has been set.
func (o *AddCopyLogFileRotationListenerRequest) GetCopyToDirectoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CopyToDirectory, true
}

// SetCopyToDirectory sets field value
func (o *AddCopyLogFileRotationListenerRequest) SetCopyToDirectory(v string) {
	o.CopyToDirectory = v
}

// GetCompressOnCopy returns the CompressOnCopy field value if set, zero value otherwise.
func (o *AddCopyLogFileRotationListenerRequest) GetCompressOnCopy() bool {
	if o == nil || IsNil(o.CompressOnCopy) {
		var ret bool
		return ret
	}
	return *o.CompressOnCopy
}

// GetCompressOnCopyOk returns a tuple with the CompressOnCopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCopyLogFileRotationListenerRequest) GetCompressOnCopyOk() (*bool, bool) {
	if o == nil || IsNil(o.CompressOnCopy) {
		return nil, false
	}
	return o.CompressOnCopy, true
}

// HasCompressOnCopy returns a boolean if a field has been set.
func (o *AddCopyLogFileRotationListenerRequest) HasCompressOnCopy() bool {
	if o != nil && !IsNil(o.CompressOnCopy) {
		return true
	}

	return false
}

// SetCompressOnCopy gets a reference to the given bool and assigns it to the CompressOnCopy field.
func (o *AddCopyLogFileRotationListenerRequest) SetCompressOnCopy(v bool) {
	o.CompressOnCopy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddCopyLogFileRotationListenerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCopyLogFileRotationListenerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddCopyLogFileRotationListenerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddCopyLogFileRotationListenerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddCopyLogFileRotationListenerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddCopyLogFileRotationListenerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddCopyLogFileRotationListenerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetListenerName returns the ListenerName field value
func (o *AddCopyLogFileRotationListenerRequest) GetListenerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListenerName
}

// GetListenerNameOk returns a tuple with the ListenerName field value
// and a boolean to check if the value has been set.
func (o *AddCopyLogFileRotationListenerRequest) GetListenerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListenerName, true
}

// SetListenerName sets field value
func (o *AddCopyLogFileRotationListenerRequest) SetListenerName(v string) {
	o.ListenerName = v
}

func (o AddCopyLogFileRotationListenerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCopyLogFileRotationListenerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["copyToDirectory"] = o.CopyToDirectory
	if !IsNil(o.CompressOnCopy) {
		toSerialize["compressOnCopy"] = o.CompressOnCopy
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["listenerName"] = o.ListenerName
	return toSerialize, nil
}

type NullableAddCopyLogFileRotationListenerRequest struct {
	value *AddCopyLogFileRotationListenerRequest
	isSet bool
}

func (v NullableAddCopyLogFileRotationListenerRequest) Get() *AddCopyLogFileRotationListenerRequest {
	return v.value
}

func (v *NullableAddCopyLogFileRotationListenerRequest) Set(val *AddCopyLogFileRotationListenerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCopyLogFileRotationListenerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCopyLogFileRotationListenerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCopyLogFileRotationListenerRequest(val *AddCopyLogFileRotationListenerRequest) *NullableAddCopyLogFileRotationListenerRequest {
	return &NullableAddCopyLogFileRotationListenerRequest{value: val, isSet: true}
}

func (v NullableAddCopyLogFileRotationListenerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCopyLogFileRotationListenerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
