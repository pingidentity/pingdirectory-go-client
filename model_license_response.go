/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LicenseResponse struct for LicenseResponse
type LicenseResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas []EnumlicenseSchemaUrn `json:"schemas,omitempty"`
	// License key enabling use of Directory Server, Directory Proxy Server, Data Sync Server, and Data Metrics Server products.
	DirectoryPlatformLicenseKey *string `json:"directoryPlatformLicenseKey,omitempty"`
}

// NewLicenseResponse instantiates a new LicenseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseResponse() *LicenseResponse {
	this := LicenseResponse{}
	return &this
}

// NewLicenseResponseWithDefaults instantiates a new LicenseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseResponseWithDefaults() *LicenseResponse {
	this := LicenseResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LicenseResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LicenseResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *LicenseResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *LicenseResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *LicenseResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *LicenseResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *LicenseResponse) GetSchemas() []EnumlicenseSchemaUrn {
	if o == nil || isNil(o.Schemas) {
		var ret []EnumlicenseSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseResponse) GetSchemasOk() ([]EnumlicenseSchemaUrn, bool) {
	if o == nil || isNil(o.Schemas) {
    return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *LicenseResponse) HasSchemas() bool {
	if o != nil && !isNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumlicenseSchemaUrn and assigns it to the Schemas field.
func (o *LicenseResponse) SetSchemas(v []EnumlicenseSchemaUrn) {
	o.Schemas = v
}

// GetDirectoryPlatformLicenseKey returns the DirectoryPlatformLicenseKey field value if set, zero value otherwise.
func (o *LicenseResponse) GetDirectoryPlatformLicenseKey() string {
	if o == nil || isNil(o.DirectoryPlatformLicenseKey) {
		var ret string
		return ret
	}
	return *o.DirectoryPlatformLicenseKey
}

// GetDirectoryPlatformLicenseKeyOk returns a tuple with the DirectoryPlatformLicenseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseResponse) GetDirectoryPlatformLicenseKeyOk() (*string, bool) {
	if o == nil || isNil(o.DirectoryPlatformLicenseKey) {
    return nil, false
	}
	return o.DirectoryPlatformLicenseKey, true
}

// HasDirectoryPlatformLicenseKey returns a boolean if a field has been set.
func (o *LicenseResponse) HasDirectoryPlatformLicenseKey() bool {
	if o != nil && !isNil(o.DirectoryPlatformLicenseKey) {
		return true
	}

	return false
}

// SetDirectoryPlatformLicenseKey gets a reference to the given string and assigns it to the DirectoryPlatformLicenseKey field.
func (o *LicenseResponse) SetDirectoryPlatformLicenseKey(v string) {
	o.DirectoryPlatformLicenseKey = &v
}

func (o LicenseResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if !isNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.DirectoryPlatformLicenseKey) {
		toSerialize["directoryPlatformLicenseKey"] = o.DirectoryPlatformLicenseKey
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseResponse struct {
	value *LicenseResponse
	isSet bool
}

func (v NullableLicenseResponse) Get() *LicenseResponse {
	return v.value
}

func (v *NullableLicenseResponse) Set(val *LicenseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseResponse(val *LicenseResponse) *NullableLicenseResponse {
	return &NullableLicenseResponse{value: val, isSet: true}
}

func (v NullableLicenseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

