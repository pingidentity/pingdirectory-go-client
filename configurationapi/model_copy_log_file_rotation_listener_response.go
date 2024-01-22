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

// checks if the CopyLogFileRotationListenerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CopyLogFileRotationListenerResponse{}

// CopyLogFileRotationListenerResponse struct for CopyLogFileRotationListenerResponse
type CopyLogFileRotationListenerResponse struct {
	Schemas []EnumcopyLogFileRotationListenerSchemaUrn `json:"schemas"`
	// The path to the directory to which log files should be copied. It must be different from the directory to which the log file is originally written, and administrators should ensure that the filesystem has sufficient space to hold files as they are copied.
	CopyToDirectory string `json:"copyToDirectory"`
	// Indicates whether the file should be gzip-compressed as it is copied into the destination directory.
	CompressOnCopy *bool `json:"compressOnCopy,omitempty"`
	// A description for this Log File Rotation Listener
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log File Rotation Listener is enabled for use.
	Enabled                                       bool                                               `json:"enabled"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Log File Rotation Listener
	Id string `json:"id"`
}

// NewCopyLogFileRotationListenerResponse instantiates a new CopyLogFileRotationListenerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCopyLogFileRotationListenerResponse(schemas []EnumcopyLogFileRotationListenerSchemaUrn, copyToDirectory string, enabled bool, id string) *CopyLogFileRotationListenerResponse {
	this := CopyLogFileRotationListenerResponse{}
	this.Schemas = schemas
	this.CopyToDirectory = copyToDirectory
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewCopyLogFileRotationListenerResponseWithDefaults instantiates a new CopyLogFileRotationListenerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCopyLogFileRotationListenerResponseWithDefaults() *CopyLogFileRotationListenerResponse {
	this := CopyLogFileRotationListenerResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *CopyLogFileRotationListenerResponse) GetSchemas() []EnumcopyLogFileRotationListenerSchemaUrn {
	if o == nil {
		var ret []EnumcopyLogFileRotationListenerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *CopyLogFileRotationListenerResponse) GetSchemasOk() ([]EnumcopyLogFileRotationListenerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *CopyLogFileRotationListenerResponse) SetSchemas(v []EnumcopyLogFileRotationListenerSchemaUrn) {
	o.Schemas = v
}

// GetCopyToDirectory returns the CopyToDirectory field value
func (o *CopyLogFileRotationListenerResponse) GetCopyToDirectory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CopyToDirectory
}

// GetCopyToDirectoryOk returns a tuple with the CopyToDirectory field value
// and a boolean to check if the value has been set.
func (o *CopyLogFileRotationListenerResponse) GetCopyToDirectoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CopyToDirectory, true
}

// SetCopyToDirectory sets field value
func (o *CopyLogFileRotationListenerResponse) SetCopyToDirectory(v string) {
	o.CopyToDirectory = v
}

// GetCompressOnCopy returns the CompressOnCopy field value if set, zero value otherwise.
func (o *CopyLogFileRotationListenerResponse) GetCompressOnCopy() bool {
	if o == nil || IsNil(o.CompressOnCopy) {
		var ret bool
		return ret
	}
	return *o.CompressOnCopy
}

// GetCompressOnCopyOk returns a tuple with the CompressOnCopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyLogFileRotationListenerResponse) GetCompressOnCopyOk() (*bool, bool) {
	if o == nil || IsNil(o.CompressOnCopy) {
		return nil, false
	}
	return o.CompressOnCopy, true
}

// HasCompressOnCopy returns a boolean if a field has been set.
func (o *CopyLogFileRotationListenerResponse) HasCompressOnCopy() bool {
	if o != nil && !IsNil(o.CompressOnCopy) {
		return true
	}

	return false
}

// SetCompressOnCopy gets a reference to the given bool and assigns it to the CompressOnCopy field.
func (o *CopyLogFileRotationListenerResponse) SetCompressOnCopy(v bool) {
	o.CompressOnCopy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CopyLogFileRotationListenerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyLogFileRotationListenerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CopyLogFileRotationListenerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CopyLogFileRotationListenerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *CopyLogFileRotationListenerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CopyLogFileRotationListenerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CopyLogFileRotationListenerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CopyLogFileRotationListenerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyLogFileRotationListenerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CopyLogFileRotationListenerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *CopyLogFileRotationListenerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *CopyLogFileRotationListenerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyLogFileRotationListenerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *CopyLogFileRotationListenerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *CopyLogFileRotationListenerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *CopyLogFileRotationListenerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CopyLogFileRotationListenerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CopyLogFileRotationListenerResponse) SetId(v string) {
	o.Id = v
}

func (o CopyLogFileRotationListenerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CopyLogFileRotationListenerResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableCopyLogFileRotationListenerResponse struct {
	value *CopyLogFileRotationListenerResponse
	isSet bool
}

func (v NullableCopyLogFileRotationListenerResponse) Get() *CopyLogFileRotationListenerResponse {
	return v.value
}

func (v *NullableCopyLogFileRotationListenerResponse) Set(val *CopyLogFileRotationListenerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCopyLogFileRotationListenerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCopyLogFileRotationListenerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCopyLogFileRotationListenerResponse(val *CopyLogFileRotationListenerResponse) *NullableCopyLogFileRotationListenerResponse {
	return &NullableCopyLogFileRotationListenerResponse{value: val, isSet: true}
}

func (v NullableCopyLogFileRotationListenerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCopyLogFileRotationListenerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
