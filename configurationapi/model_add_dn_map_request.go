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

// checks if the AddDnMapRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddDnMapRequest{}

// AddDnMapRequest struct for AddDnMapRequest
type AddDnMapRequest struct {
	Schemas []EnumdnMapSchemaUrn `json:"schemas,omitempty"`
	// A description for this DN Map
	Description *string `json:"description,omitempty"`
	// Specifies the DN pattern to match when determining whether this map applies to a specific source DN. If the provided bind DN matches this pattern, then the to-dn-pattern will be used to perform the mapping. If the provided bind DN does not match this pattern, then no mapping will be performed.
	FromDNPattern string `json:"fromDNPattern"`
	// Specifies a pattern for constructing the DN value using fixed text, DN components matching wild-card values in from-dn-pattern, and attribute values from the source entry.
	ToDNPattern string `json:"toDNPattern"`
	// Name of the new DN Map
	MapName string `json:"mapName"`
}

// NewAddDnMapRequest instantiates a new AddDnMapRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDnMapRequest(fromDNPattern string, toDNPattern string, mapName string) *AddDnMapRequest {
	this := AddDnMapRequest{}
	this.FromDNPattern = fromDNPattern
	this.ToDNPattern = toDNPattern
	this.MapName = mapName
	return &this
}

// NewAddDnMapRequestWithDefaults instantiates a new AddDnMapRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDnMapRequestWithDefaults() *AddDnMapRequest {
	this := AddDnMapRequest{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddDnMapRequest) GetSchemas() []EnumdnMapSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumdnMapSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDnMapRequest) GetSchemasOk() ([]EnumdnMapSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddDnMapRequest) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumdnMapSchemaUrn and assigns it to the Schemas field.
func (o *AddDnMapRequest) SetSchemas(v []EnumdnMapSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddDnMapRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDnMapRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddDnMapRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddDnMapRequest) SetDescription(v string) {
	o.Description = &v
}

// GetFromDNPattern returns the FromDNPattern field value
func (o *AddDnMapRequest) GetFromDNPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromDNPattern
}

// GetFromDNPatternOk returns a tuple with the FromDNPattern field value
// and a boolean to check if the value has been set.
func (o *AddDnMapRequest) GetFromDNPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromDNPattern, true
}

// SetFromDNPattern sets field value
func (o *AddDnMapRequest) SetFromDNPattern(v string) {
	o.FromDNPattern = v
}

// GetToDNPattern returns the ToDNPattern field value
func (o *AddDnMapRequest) GetToDNPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToDNPattern
}

// GetToDNPatternOk returns a tuple with the ToDNPattern field value
// and a boolean to check if the value has been set.
func (o *AddDnMapRequest) GetToDNPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToDNPattern, true
}

// SetToDNPattern sets field value
func (o *AddDnMapRequest) SetToDNPattern(v string) {
	o.ToDNPattern = v
}

// GetMapName returns the MapName field value
func (o *AddDnMapRequest) GetMapName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MapName
}

// GetMapNameOk returns a tuple with the MapName field value
// and a boolean to check if the value has been set.
func (o *AddDnMapRequest) GetMapNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MapName, true
}

// SetMapName sets field value
func (o *AddDnMapRequest) SetMapName(v string) {
	o.MapName = v
}

func (o AddDnMapRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddDnMapRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["fromDNPattern"] = o.FromDNPattern
	toSerialize["toDNPattern"] = o.ToDNPattern
	toSerialize["mapName"] = o.MapName
	return toSerialize, nil
}

type NullableAddDnMapRequest struct {
	value *AddDnMapRequest
	isSet bool
}

func (v NullableAddDnMapRequest) Get() *AddDnMapRequest {
	return v.value
}

func (v *NullableAddDnMapRequest) Set(val *AddDnMapRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDnMapRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDnMapRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDnMapRequest(val *AddDnMapRequest) *NullableAddDnMapRequest {
	return &NullableAddDnMapRequest{value: val, isSet: true}
}

func (v NullableAddDnMapRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDnMapRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
