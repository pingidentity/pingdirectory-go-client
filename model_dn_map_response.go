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

// DnMapResponse struct for DnMapResponse
type DnMapResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the DN Map
	Id string `json:"id"`
	Schemas []EnumdnMapSchemaUrn `json:"schemas,omitempty"`
	// A description for this DN Map
	Description *string `json:"description,omitempty"`
	// Specifies the DN pattern to match when determining whether this map applies to a specific source DN. If the provided bind DN matches this pattern, then the to-dn-pattern will be used to perform the mapping. If the provided bind DN does not match this pattern, then no mapping will be performed.
	FromDNPattern string `json:"fromDNPattern"`
	// Specifies a pattern for constructing the DN value using fixed text, DN components matching wild-card values in from-dn-pattern, and attribute values from the source entry.
	ToDNPattern string `json:"toDNPattern"`
}

// NewDnMapResponse instantiates a new DnMapResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnMapResponse(id string, fromDNPattern string, toDNPattern string) *DnMapResponse {
	this := DnMapResponse{}
	this.Id = id
	this.FromDNPattern = fromDNPattern
	this.ToDNPattern = toDNPattern
	return &this
}

// NewDnMapResponseWithDefaults instantiates a new DnMapResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnMapResponseWithDefaults() *DnMapResponse {
	this := DnMapResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DnMapResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnMapResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DnMapResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *DnMapResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *DnMapResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnMapResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *DnMapResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *DnMapResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *DnMapResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DnMapResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DnMapResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *DnMapResponse) GetSchemas() []EnumdnMapSchemaUrn {
	if o == nil || isNil(o.Schemas) {
		var ret []EnumdnMapSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnMapResponse) GetSchemasOk() ([]EnumdnMapSchemaUrn, bool) {
	if o == nil || isNil(o.Schemas) {
    return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *DnMapResponse) HasSchemas() bool {
	if o != nil && !isNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumdnMapSchemaUrn and assigns it to the Schemas field.
func (o *DnMapResponse) SetSchemas(v []EnumdnMapSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DnMapResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnMapResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DnMapResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DnMapResponse) SetDescription(v string) {
	o.Description = &v
}

// GetFromDNPattern returns the FromDNPattern field value
func (o *DnMapResponse) GetFromDNPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromDNPattern
}

// GetFromDNPatternOk returns a tuple with the FromDNPattern field value
// and a boolean to check if the value has been set.
func (o *DnMapResponse) GetFromDNPatternOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FromDNPattern, true
}

// SetFromDNPattern sets field value
func (o *DnMapResponse) SetFromDNPattern(v string) {
	o.FromDNPattern = v
}

// GetToDNPattern returns the ToDNPattern field value
func (o *DnMapResponse) GetToDNPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToDNPattern
}

// GetToDNPatternOk returns a tuple with the ToDNPattern field value
// and a boolean to check if the value has been set.
func (o *DnMapResponse) GetToDNPatternOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ToDNPattern, true
}

// SetToDNPattern sets field value
func (o *DnMapResponse) SetToDNPattern(v string) {
	o.ToDNPattern = v
}

func (o DnMapResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["fromDNPattern"] = o.FromDNPattern
	}
	if true {
		toSerialize["toDNPattern"] = o.ToDNPattern
	}
	return json.Marshal(toSerialize)
}

type NullableDnMapResponse struct {
	value *DnMapResponse
	isSet bool
}

func (v NullableDnMapResponse) Get() *DnMapResponse {
	return v.value
}

func (v *NullableDnMapResponse) Set(val *DnMapResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDnMapResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDnMapResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnMapResponse(val *DnMapResponse) *NullableDnMapResponse {
	return &NullableDnMapResponse{value: val, isSet: true}
}

func (v NullableDnMapResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnMapResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


