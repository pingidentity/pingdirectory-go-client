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

// SizeLimitLogRetentionPolicyResponse struct for SizeLimitLogRetentionPolicyResponse
type SizeLimitLogRetentionPolicyResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Log Retention Policy
	Id      string                                     `json:"id"`
	Schemas []EnumsizeLimitLogRetentionPolicySchemaUrn `json:"schemas"`
	// Specifies the maximum total disk space used by the log files.
	DiskSpaceUsed string `json:"diskSpaceUsed"`
	// A description for this Log Retention Policy
	Description *string `json:"description,omitempty"`
}

// NewSizeLimitLogRetentionPolicyResponse instantiates a new SizeLimitLogRetentionPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSizeLimitLogRetentionPolicyResponse(id string, schemas []EnumsizeLimitLogRetentionPolicySchemaUrn, diskSpaceUsed string) *SizeLimitLogRetentionPolicyResponse {
	this := SizeLimitLogRetentionPolicyResponse{}
	this.Id = id
	this.Schemas = schemas
	this.DiskSpaceUsed = diskSpaceUsed
	return &this
}

// NewSizeLimitLogRetentionPolicyResponseWithDefaults instantiates a new SizeLimitLogRetentionPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSizeLimitLogRetentionPolicyResponseWithDefaults() *SizeLimitLogRetentionPolicyResponse {
	this := SizeLimitLogRetentionPolicyResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SizeLimitLogRetentionPolicyResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SizeLimitLogRetentionPolicyResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SizeLimitLogRetentionPolicyResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SizeLimitLogRetentionPolicyResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SizeLimitLogRetentionPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SizeLimitLogRetentionPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SizeLimitLogRetentionPolicyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SizeLimitLogRetentionPolicyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *SizeLimitLogRetentionPolicyResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SizeLimitLogRetentionPolicyResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SizeLimitLogRetentionPolicyResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *SizeLimitLogRetentionPolicyResponse) GetSchemas() []EnumsizeLimitLogRetentionPolicySchemaUrn {
	if o == nil {
		var ret []EnumsizeLimitLogRetentionPolicySchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SizeLimitLogRetentionPolicyResponse) GetSchemasOk() ([]EnumsizeLimitLogRetentionPolicySchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SizeLimitLogRetentionPolicyResponse) SetSchemas(v []EnumsizeLimitLogRetentionPolicySchemaUrn) {
	o.Schemas = v
}

// GetDiskSpaceUsed returns the DiskSpaceUsed field value
func (o *SizeLimitLogRetentionPolicyResponse) GetDiskSpaceUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiskSpaceUsed
}

// GetDiskSpaceUsedOk returns a tuple with the DiskSpaceUsed field value
// and a boolean to check if the value has been set.
func (o *SizeLimitLogRetentionPolicyResponse) GetDiskSpaceUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskSpaceUsed, true
}

// SetDiskSpaceUsed sets field value
func (o *SizeLimitLogRetentionPolicyResponse) SetDiskSpaceUsed(v string) {
	o.DiskSpaceUsed = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SizeLimitLogRetentionPolicyResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SizeLimitLogRetentionPolicyResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SizeLimitLogRetentionPolicyResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SizeLimitLogRetentionPolicyResponse) SetDescription(v string) {
	o.Description = &v
}

func (o SizeLimitLogRetentionPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["diskSpaceUsed"] = o.DiskSpaceUsed
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableSizeLimitLogRetentionPolicyResponse struct {
	value *SizeLimitLogRetentionPolicyResponse
	isSet bool
}

func (v NullableSizeLimitLogRetentionPolicyResponse) Get() *SizeLimitLogRetentionPolicyResponse {
	return v.value
}

func (v *NullableSizeLimitLogRetentionPolicyResponse) Set(val *SizeLimitLogRetentionPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSizeLimitLogRetentionPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSizeLimitLogRetentionPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSizeLimitLogRetentionPolicyResponse(val *SizeLimitLogRetentionPolicyResponse) *NullableSizeLimitLogRetentionPolicyResponse {
	return &NullableSizeLimitLogRetentionPolicyResponse{value: val, isSet: true}
}

func (v NullableSizeLimitLogRetentionPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSizeLimitLogRetentionPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}