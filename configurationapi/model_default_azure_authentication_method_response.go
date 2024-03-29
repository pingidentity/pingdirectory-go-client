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

// checks if the DefaultAzureAuthenticationMethodResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultAzureAuthenticationMethodResponse{}

// DefaultAzureAuthenticationMethodResponse struct for DefaultAzureAuthenticationMethodResponse
type DefaultAzureAuthenticationMethodResponse struct {
	Schemas []EnumdefaultAzureAuthenticationMethodSchemaUrn `json:"schemas"`
	// The tenant ID to use to authenticate. If this is not provided, then it will be obtained from the AZURE_TENANT_ID environment variable.
	TenantID *string `json:"tenantID,omitempty"`
	// The client ID to use to authenticate. If this is not provided, then it will be obtained from the AZURE_CLIENT_ID
	ClientID *string `json:"clientID,omitempty"`
	// A description for this Azure Authentication Method
	Description                                   *string                                            `json:"description,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Azure Authentication Method
	Id string `json:"id"`
}

// NewDefaultAzureAuthenticationMethodResponse instantiates a new DefaultAzureAuthenticationMethodResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultAzureAuthenticationMethodResponse(schemas []EnumdefaultAzureAuthenticationMethodSchemaUrn, id string) *DefaultAzureAuthenticationMethodResponse {
	this := DefaultAzureAuthenticationMethodResponse{}
	this.Schemas = schemas
	this.Id = id
	return &this
}

// NewDefaultAzureAuthenticationMethodResponseWithDefaults instantiates a new DefaultAzureAuthenticationMethodResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultAzureAuthenticationMethodResponseWithDefaults() *DefaultAzureAuthenticationMethodResponse {
	this := DefaultAzureAuthenticationMethodResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *DefaultAzureAuthenticationMethodResponse) GetSchemas() []EnumdefaultAzureAuthenticationMethodSchemaUrn {
	if o == nil {
		var ret []EnumdefaultAzureAuthenticationMethodSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *DefaultAzureAuthenticationMethodResponse) GetSchemasOk() ([]EnumdefaultAzureAuthenticationMethodSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *DefaultAzureAuthenticationMethodResponse) SetSchemas(v []EnumdefaultAzureAuthenticationMethodSchemaUrn) {
	o.Schemas = v
}

// GetTenantID returns the TenantID field value if set, zero value otherwise.
func (o *DefaultAzureAuthenticationMethodResponse) GetTenantID() string {
	if o == nil || IsNil(o.TenantID) {
		var ret string
		return ret
	}
	return *o.TenantID
}

// GetTenantIDOk returns a tuple with the TenantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultAzureAuthenticationMethodResponse) GetTenantIDOk() (*string, bool) {
	if o == nil || IsNil(o.TenantID) {
		return nil, false
	}
	return o.TenantID, true
}

// HasTenantID returns a boolean if a field has been set.
func (o *DefaultAzureAuthenticationMethodResponse) HasTenantID() bool {
	if o != nil && !IsNil(o.TenantID) {
		return true
	}

	return false
}

// SetTenantID gets a reference to the given string and assigns it to the TenantID field.
func (o *DefaultAzureAuthenticationMethodResponse) SetTenantID(v string) {
	o.TenantID = &v
}

// GetClientID returns the ClientID field value if set, zero value otherwise.
func (o *DefaultAzureAuthenticationMethodResponse) GetClientID() string {
	if o == nil || IsNil(o.ClientID) {
		var ret string
		return ret
	}
	return *o.ClientID
}

// GetClientIDOk returns a tuple with the ClientID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultAzureAuthenticationMethodResponse) GetClientIDOk() (*string, bool) {
	if o == nil || IsNil(o.ClientID) {
		return nil, false
	}
	return o.ClientID, true
}

// HasClientID returns a boolean if a field has been set.
func (o *DefaultAzureAuthenticationMethodResponse) HasClientID() bool {
	if o != nil && !IsNil(o.ClientID) {
		return true
	}

	return false
}

// SetClientID gets a reference to the given string and assigns it to the ClientID field.
func (o *DefaultAzureAuthenticationMethodResponse) SetClientID(v string) {
	o.ClientID = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DefaultAzureAuthenticationMethodResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultAzureAuthenticationMethodResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DefaultAzureAuthenticationMethodResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DefaultAzureAuthenticationMethodResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DefaultAzureAuthenticationMethodResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultAzureAuthenticationMethodResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DefaultAzureAuthenticationMethodResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *DefaultAzureAuthenticationMethodResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *DefaultAzureAuthenticationMethodResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultAzureAuthenticationMethodResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *DefaultAzureAuthenticationMethodResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *DefaultAzureAuthenticationMethodResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *DefaultAzureAuthenticationMethodResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DefaultAzureAuthenticationMethodResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DefaultAzureAuthenticationMethodResponse) SetId(v string) {
	o.Id = v
}

func (o DefaultAzureAuthenticationMethodResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultAzureAuthenticationMethodResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.TenantID) {
		toSerialize["tenantID"] = o.TenantID
	}
	if !IsNil(o.ClientID) {
		toSerialize["clientID"] = o.ClientID
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

type NullableDefaultAzureAuthenticationMethodResponse struct {
	value *DefaultAzureAuthenticationMethodResponse
	isSet bool
}

func (v NullableDefaultAzureAuthenticationMethodResponse) Get() *DefaultAzureAuthenticationMethodResponse {
	return v.value
}

func (v *NullableDefaultAzureAuthenticationMethodResponse) Set(val *DefaultAzureAuthenticationMethodResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultAzureAuthenticationMethodResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultAzureAuthenticationMethodResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultAzureAuthenticationMethodResponse(val *DefaultAzureAuthenticationMethodResponse) *NullableDefaultAzureAuthenticationMethodResponse {
	return &NullableDefaultAzureAuthenticationMethodResponse{value: val, isSet: true}
}

func (v NullableDefaultAzureAuthenticationMethodResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultAzureAuthenticationMethodResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
