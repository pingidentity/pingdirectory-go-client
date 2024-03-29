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

// checks if the PingOneHttpExternalServerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PingOneHttpExternalServerResponse{}

// PingOneHttpExternalServerResponse struct for PingOneHttpExternalServerResponse
type PingOneHttpExternalServerResponse struct {
	Schemas                    []EnumpingOneHttpExternalServerSchemaUrn                     `json:"schemas"`
	HostnameVerificationMethod *EnumexternalServerPingOneHttpHostnameVerificationMethodProp `json:"hostnameVerificationMethod,omitempty"`
	// The trust manager provider to use for HTTPS connection-level security.
	TrustManagerProvider *string `json:"trustManagerProvider,omitempty"`
	// Specifies the maximum length of time to wait for a connection to be established before aborting a request to PingOne.
	ConnectTimeout *string `json:"connectTimeout,omitempty"`
	// Specifies the maximum length of time to wait for response data to be read from an established connection before aborting a request to PingOne.
	ResponseTimeout *string `json:"responseTimeout,omitempty"`
	// A description for this External Server
	Description                                   *string                                            `json:"description,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the External Server
	Id string `json:"id"`
}

// NewPingOneHttpExternalServerResponse instantiates a new PingOneHttpExternalServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingOneHttpExternalServerResponse(schemas []EnumpingOneHttpExternalServerSchemaUrn, id string) *PingOneHttpExternalServerResponse {
	this := PingOneHttpExternalServerResponse{}
	this.Schemas = schemas
	this.Id = id
	return &this
}

// NewPingOneHttpExternalServerResponseWithDefaults instantiates a new PingOneHttpExternalServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingOneHttpExternalServerResponseWithDefaults() *PingOneHttpExternalServerResponse {
	this := PingOneHttpExternalServerResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *PingOneHttpExternalServerResponse) GetSchemas() []EnumpingOneHttpExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumpingOneHttpExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *PingOneHttpExternalServerResponse) GetSchemasOk() ([]EnumpingOneHttpExternalServerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *PingOneHttpExternalServerResponse) SetSchemas(v []EnumpingOneHttpExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetHostnameVerificationMethod returns the HostnameVerificationMethod field value if set, zero value otherwise.
func (o *PingOneHttpExternalServerResponse) GetHostnameVerificationMethod() EnumexternalServerPingOneHttpHostnameVerificationMethodProp {
	if o == nil || IsNil(o.HostnameVerificationMethod) {
		var ret EnumexternalServerPingOneHttpHostnameVerificationMethodProp
		return ret
	}
	return *o.HostnameVerificationMethod
}

// GetHostnameVerificationMethodOk returns a tuple with the HostnameVerificationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneHttpExternalServerResponse) GetHostnameVerificationMethodOk() (*EnumexternalServerPingOneHttpHostnameVerificationMethodProp, bool) {
	if o == nil || IsNil(o.HostnameVerificationMethod) {
		return nil, false
	}
	return o.HostnameVerificationMethod, true
}

// HasHostnameVerificationMethod returns a boolean if a field has been set.
func (o *PingOneHttpExternalServerResponse) HasHostnameVerificationMethod() bool {
	if o != nil && !IsNil(o.HostnameVerificationMethod) {
		return true
	}

	return false
}

// SetHostnameVerificationMethod gets a reference to the given EnumexternalServerPingOneHttpHostnameVerificationMethodProp and assigns it to the HostnameVerificationMethod field.
func (o *PingOneHttpExternalServerResponse) SetHostnameVerificationMethod(v EnumexternalServerPingOneHttpHostnameVerificationMethodProp) {
	o.HostnameVerificationMethod = &v
}

// GetTrustManagerProvider returns the TrustManagerProvider field value if set, zero value otherwise.
func (o *PingOneHttpExternalServerResponse) GetTrustManagerProvider() string {
	if o == nil || IsNil(o.TrustManagerProvider) {
		var ret string
		return ret
	}
	return *o.TrustManagerProvider
}

// GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneHttpExternalServerResponse) GetTrustManagerProviderOk() (*string, bool) {
	if o == nil || IsNil(o.TrustManagerProvider) {
		return nil, false
	}
	return o.TrustManagerProvider, true
}

// HasTrustManagerProvider returns a boolean if a field has been set.
func (o *PingOneHttpExternalServerResponse) HasTrustManagerProvider() bool {
	if o != nil && !IsNil(o.TrustManagerProvider) {
		return true
	}

	return false
}

// SetTrustManagerProvider gets a reference to the given string and assigns it to the TrustManagerProvider field.
func (o *PingOneHttpExternalServerResponse) SetTrustManagerProvider(v string) {
	o.TrustManagerProvider = &v
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise.
func (o *PingOneHttpExternalServerResponse) GetConnectTimeout() string {
	if o == nil || IsNil(o.ConnectTimeout) {
		var ret string
		return ret
	}
	return *o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneHttpExternalServerResponse) GetConnectTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectTimeout) {
		return nil, false
	}
	return o.ConnectTimeout, true
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *PingOneHttpExternalServerResponse) HasConnectTimeout() bool {
	if o != nil && !IsNil(o.ConnectTimeout) {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given string and assigns it to the ConnectTimeout field.
func (o *PingOneHttpExternalServerResponse) SetConnectTimeout(v string) {
	o.ConnectTimeout = &v
}

// GetResponseTimeout returns the ResponseTimeout field value if set, zero value otherwise.
func (o *PingOneHttpExternalServerResponse) GetResponseTimeout() string {
	if o == nil || IsNil(o.ResponseTimeout) {
		var ret string
		return ret
	}
	return *o.ResponseTimeout
}

// GetResponseTimeoutOk returns a tuple with the ResponseTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneHttpExternalServerResponse) GetResponseTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseTimeout) {
		return nil, false
	}
	return o.ResponseTimeout, true
}

// HasResponseTimeout returns a boolean if a field has been set.
func (o *PingOneHttpExternalServerResponse) HasResponseTimeout() bool {
	if o != nil && !IsNil(o.ResponseTimeout) {
		return true
	}

	return false
}

// SetResponseTimeout gets a reference to the given string and assigns it to the ResponseTimeout field.
func (o *PingOneHttpExternalServerResponse) SetResponseTimeout(v string) {
	o.ResponseTimeout = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PingOneHttpExternalServerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneHttpExternalServerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PingOneHttpExternalServerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PingOneHttpExternalServerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PingOneHttpExternalServerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneHttpExternalServerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PingOneHttpExternalServerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *PingOneHttpExternalServerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *PingOneHttpExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneHttpExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *PingOneHttpExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *PingOneHttpExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *PingOneHttpExternalServerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PingOneHttpExternalServerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PingOneHttpExternalServerResponse) SetId(v string) {
	o.Id = v
}

func (o PingOneHttpExternalServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PingOneHttpExternalServerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.HostnameVerificationMethod) {
		toSerialize["hostnameVerificationMethod"] = o.HostnameVerificationMethod
	}
	if !IsNil(o.TrustManagerProvider) {
		toSerialize["trustManagerProvider"] = o.TrustManagerProvider
	}
	if !IsNil(o.ConnectTimeout) {
		toSerialize["connectTimeout"] = o.ConnectTimeout
	}
	if !IsNil(o.ResponseTimeout) {
		toSerialize["responseTimeout"] = o.ResponseTimeout
	}
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

type NullablePingOneHttpExternalServerResponse struct {
	value *PingOneHttpExternalServerResponse
	isSet bool
}

func (v NullablePingOneHttpExternalServerResponse) Get() *PingOneHttpExternalServerResponse {
	return v.value
}

func (v *NullablePingOneHttpExternalServerResponse) Set(val *PingOneHttpExternalServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePingOneHttpExternalServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePingOneHttpExternalServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingOneHttpExternalServerResponse(val *PingOneHttpExternalServerResponse) *NullablePingOneHttpExternalServerResponse {
	return &NullablePingOneHttpExternalServerResponse{value: val, isSet: true}
}

func (v NullablePingOneHttpExternalServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingOneHttpExternalServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
