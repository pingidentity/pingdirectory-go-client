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

// checks if the FileCountLogRetentionPolicyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileCountLogRetentionPolicyResponse{}

// FileCountLogRetentionPolicyResponse struct for FileCountLogRetentionPolicyResponse
type FileCountLogRetentionPolicyResponse struct {
	Schemas []EnumfileCountLogRetentionPolicySchemaUrn `json:"schemas"`
	// Specifies the number of archived log files to retain before the oldest ones are cleaned.
	NumberOfFiles int64 `json:"numberOfFiles"`
	// A description for this Log Retention Policy
	Description                                   *string                                            `json:"description,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Log Retention Policy
	Id string `json:"id"`
}

// NewFileCountLogRetentionPolicyResponse instantiates a new FileCountLogRetentionPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileCountLogRetentionPolicyResponse(schemas []EnumfileCountLogRetentionPolicySchemaUrn, numberOfFiles int64, id string) *FileCountLogRetentionPolicyResponse {
	this := FileCountLogRetentionPolicyResponse{}
	this.Schemas = schemas
	this.NumberOfFiles = numberOfFiles
	this.Id = id
	return &this
}

// NewFileCountLogRetentionPolicyResponseWithDefaults instantiates a new FileCountLogRetentionPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileCountLogRetentionPolicyResponseWithDefaults() *FileCountLogRetentionPolicyResponse {
	this := FileCountLogRetentionPolicyResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *FileCountLogRetentionPolicyResponse) GetSchemas() []EnumfileCountLogRetentionPolicySchemaUrn {
	if o == nil {
		var ret []EnumfileCountLogRetentionPolicySchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *FileCountLogRetentionPolicyResponse) GetSchemasOk() ([]EnumfileCountLogRetentionPolicySchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *FileCountLogRetentionPolicyResponse) SetSchemas(v []EnumfileCountLogRetentionPolicySchemaUrn) {
	o.Schemas = v
}

// GetNumberOfFiles returns the NumberOfFiles field value
func (o *FileCountLogRetentionPolicyResponse) GetNumberOfFiles() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumberOfFiles
}

// GetNumberOfFilesOk returns a tuple with the NumberOfFiles field value
// and a boolean to check if the value has been set.
func (o *FileCountLogRetentionPolicyResponse) GetNumberOfFilesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfFiles, true
}

// SetNumberOfFiles sets field value
func (o *FileCountLogRetentionPolicyResponse) SetNumberOfFiles(v int64) {
	o.NumberOfFiles = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FileCountLogRetentionPolicyResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCountLogRetentionPolicyResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FileCountLogRetentionPolicyResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FileCountLogRetentionPolicyResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FileCountLogRetentionPolicyResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCountLogRetentionPolicyResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FileCountLogRetentionPolicyResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *FileCountLogRetentionPolicyResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *FileCountLogRetentionPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCountLogRetentionPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *FileCountLogRetentionPolicyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *FileCountLogRetentionPolicyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *FileCountLogRetentionPolicyResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileCountLogRetentionPolicyResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileCountLogRetentionPolicyResponse) SetId(v string) {
	o.Id = v
}

func (o FileCountLogRetentionPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileCountLogRetentionPolicyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["numberOfFiles"] = o.NumberOfFiles
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableFileCountLogRetentionPolicyResponse struct {
	value *FileCountLogRetentionPolicyResponse
	isSet bool
}

func (v NullableFileCountLogRetentionPolicyResponse) Get() *FileCountLogRetentionPolicyResponse {
	return v.value
}

func (v *NullableFileCountLogRetentionPolicyResponse) Set(val *FileCountLogRetentionPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileCountLogRetentionPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileCountLogRetentionPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileCountLogRetentionPolicyResponse(val *FileCountLogRetentionPolicyResponse) *NullableFileCountLogRetentionPolicyResponse {
	return &NullableFileCountLogRetentionPolicyResponse{value: val, isSet: true}
}

func (v NullableFileCountLogRetentionPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileCountLogRetentionPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
