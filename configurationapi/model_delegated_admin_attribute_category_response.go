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

// checks if the DelegatedAdminAttributeCategoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DelegatedAdminAttributeCategoryResponse{}

// DelegatedAdminAttributeCategoryResponse struct for DelegatedAdminAttributeCategoryResponse
type DelegatedAdminAttributeCategoryResponse struct {
	// Name of the Delegated Admin Attribute Category
	Id      string                                         `json:"id"`
	Schemas []EnumdelegatedAdminAttributeCategorySchemaUrn `json:"schemas,omitempty"`
	// A description for this Delegated Admin Attribute Category
	Description *string `json:"description,omitempty"`
	// A human readable display name for this Delegated Admin Attribute Category.
	DisplayName string `json:"displayName"`
	// Delegated Admin Attribute Categories are ordered for display based on this index from least to greatest.
	DisplayOrderIndex                             int32                                              `json:"displayOrderIndex"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewDelegatedAdminAttributeCategoryResponse instantiates a new DelegatedAdminAttributeCategoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelegatedAdminAttributeCategoryResponse(id string, displayName string, displayOrderIndex int32) *DelegatedAdminAttributeCategoryResponse {
	this := DelegatedAdminAttributeCategoryResponse{}
	this.Id = id
	this.DisplayName = displayName
	this.DisplayOrderIndex = displayOrderIndex
	return &this
}

// NewDelegatedAdminAttributeCategoryResponseWithDefaults instantiates a new DelegatedAdminAttributeCategoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelegatedAdminAttributeCategoryResponseWithDefaults() *DelegatedAdminAttributeCategoryResponse {
	this := DelegatedAdminAttributeCategoryResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DelegatedAdminAttributeCategoryResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DelegatedAdminAttributeCategoryResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DelegatedAdminAttributeCategoryResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *DelegatedAdminAttributeCategoryResponse) GetSchemas() []EnumdelegatedAdminAttributeCategorySchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumdelegatedAdminAttributeCategorySchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminAttributeCategoryResponse) GetSchemasOk() ([]EnumdelegatedAdminAttributeCategorySchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *DelegatedAdminAttributeCategoryResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumdelegatedAdminAttributeCategorySchemaUrn and assigns it to the Schemas field.
func (o *DelegatedAdminAttributeCategoryResponse) SetSchemas(v []EnumdelegatedAdminAttributeCategorySchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DelegatedAdminAttributeCategoryResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminAttributeCategoryResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DelegatedAdminAttributeCategoryResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DelegatedAdminAttributeCategoryResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value
func (o *DelegatedAdminAttributeCategoryResponse) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *DelegatedAdminAttributeCategoryResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *DelegatedAdminAttributeCategoryResponse) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDisplayOrderIndex returns the DisplayOrderIndex field value
func (o *DelegatedAdminAttributeCategoryResponse) GetDisplayOrderIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DisplayOrderIndex
}

// GetDisplayOrderIndexOk returns a tuple with the DisplayOrderIndex field value
// and a boolean to check if the value has been set.
func (o *DelegatedAdminAttributeCategoryResponse) GetDisplayOrderIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayOrderIndex, true
}

// SetDisplayOrderIndex sets field value
func (o *DelegatedAdminAttributeCategoryResponse) SetDisplayOrderIndex(v int32) {
	o.DisplayOrderIndex = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DelegatedAdminAttributeCategoryResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminAttributeCategoryResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DelegatedAdminAttributeCategoryResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *DelegatedAdminAttributeCategoryResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *DelegatedAdminAttributeCategoryResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminAttributeCategoryResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *DelegatedAdminAttributeCategoryResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *DelegatedAdminAttributeCategoryResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o DelegatedAdminAttributeCategoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DelegatedAdminAttributeCategoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["displayOrderIndex"] = o.DisplayOrderIndex
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableDelegatedAdminAttributeCategoryResponse struct {
	value *DelegatedAdminAttributeCategoryResponse
	isSet bool
}

func (v NullableDelegatedAdminAttributeCategoryResponse) Get() *DelegatedAdminAttributeCategoryResponse {
	return v.value
}

func (v *NullableDelegatedAdminAttributeCategoryResponse) Set(val *DelegatedAdminAttributeCategoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDelegatedAdminAttributeCategoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDelegatedAdminAttributeCategoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelegatedAdminAttributeCategoryResponse(val *DelegatedAdminAttributeCategoryResponse) *NullableDelegatedAdminAttributeCategoryResponse {
	return &NullableDelegatedAdminAttributeCategoryResponse{value: val, isSet: true}
}

func (v NullableDelegatedAdminAttributeCategoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelegatedAdminAttributeCategoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
