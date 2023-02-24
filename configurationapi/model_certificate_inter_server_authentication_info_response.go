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

// checks if the CertificateInterServerAuthenticationInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateInterServerAuthenticationInfoResponse{}

// CertificateInterServerAuthenticationInfoResponse struct for CertificateInterServerAuthenticationInfoResponse
type CertificateInterServerAuthenticationInfoResponse struct {
	Meta                                          *MetaMeta                                               `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20      `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumcertificateInterServerAuthenticationInfoSchemaUrn `json:"schemas"`
	// Name of the Inter Server Authentication Info
	Id      string                                         `json:"id"`
	Purpose []EnuminterServerAuthenticationInfoPurposeProp `json:"purpose,omitempty"`
}

// NewCertificateInterServerAuthenticationInfoResponse instantiates a new CertificateInterServerAuthenticationInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateInterServerAuthenticationInfoResponse(schemas []EnumcertificateInterServerAuthenticationInfoSchemaUrn, id string) *CertificateInterServerAuthenticationInfoResponse {
	this := CertificateInterServerAuthenticationInfoResponse{}
	this.Schemas = schemas
	this.Id = id
	return &this
}

// NewCertificateInterServerAuthenticationInfoResponseWithDefaults instantiates a new CertificateInterServerAuthenticationInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateInterServerAuthenticationInfoResponseWithDefaults() *CertificateInterServerAuthenticationInfoResponse {
	this := CertificateInterServerAuthenticationInfoResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CertificateInterServerAuthenticationInfoResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateInterServerAuthenticationInfoResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CertificateInterServerAuthenticationInfoResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *CertificateInterServerAuthenticationInfoResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *CertificateInterServerAuthenticationInfoResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateInterServerAuthenticationInfoResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *CertificateInterServerAuthenticationInfoResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *CertificateInterServerAuthenticationInfoResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *CertificateInterServerAuthenticationInfoResponse) GetSchemas() []EnumcertificateInterServerAuthenticationInfoSchemaUrn {
	if o == nil {
		var ret []EnumcertificateInterServerAuthenticationInfoSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *CertificateInterServerAuthenticationInfoResponse) GetSchemasOk() ([]EnumcertificateInterServerAuthenticationInfoSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *CertificateInterServerAuthenticationInfoResponse) SetSchemas(v []EnumcertificateInterServerAuthenticationInfoSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *CertificateInterServerAuthenticationInfoResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CertificateInterServerAuthenticationInfoResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CertificateInterServerAuthenticationInfoResponse) SetId(v string) {
	o.Id = v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *CertificateInterServerAuthenticationInfoResponse) GetPurpose() []EnuminterServerAuthenticationInfoPurposeProp {
	if o == nil || IsNil(o.Purpose) {
		var ret []EnuminterServerAuthenticationInfoPurposeProp
		return ret
	}
	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateInterServerAuthenticationInfoResponse) GetPurposeOk() ([]EnuminterServerAuthenticationInfoPurposeProp, bool) {
	if o == nil || IsNil(o.Purpose) {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *CertificateInterServerAuthenticationInfoResponse) HasPurpose() bool {
	if o != nil && !IsNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given []EnuminterServerAuthenticationInfoPurposeProp and assigns it to the Purpose field.
func (o *CertificateInterServerAuthenticationInfoResponse) SetPurpose(v []EnuminterServerAuthenticationInfoPurposeProp) {
	o.Purpose = v
}

func (o CertificateInterServerAuthenticationInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateInterServerAuthenticationInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	if !IsNil(o.Purpose) {
		toSerialize["purpose"] = o.Purpose
	}
	return toSerialize, nil
}

type NullableCertificateInterServerAuthenticationInfoResponse struct {
	value *CertificateInterServerAuthenticationInfoResponse
	isSet bool
}

func (v NullableCertificateInterServerAuthenticationInfoResponse) Get() *CertificateInterServerAuthenticationInfoResponse {
	return v.value
}

func (v *NullableCertificateInterServerAuthenticationInfoResponse) Set(val *CertificateInterServerAuthenticationInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateInterServerAuthenticationInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateInterServerAuthenticationInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateInterServerAuthenticationInfoResponse(val *CertificateInterServerAuthenticationInfoResponse) *NullableCertificateInterServerAuthenticationInfoResponse {
	return &NullableCertificateInterServerAuthenticationInfoResponse{value: val, isSet: true}
}

func (v NullableCertificateInterServerAuthenticationInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateInterServerAuthenticationInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
