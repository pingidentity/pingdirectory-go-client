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

// LocalDbVlvIndexResponse struct for LocalDbVlvIndexResponse
type LocalDbVlvIndexResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Local DB VLV Index
	Id      string                         `json:"id"`
	Schemas []EnumlocalDbVlvIndexSchemaUrn `json:"schemas,omitempty"`
	// Specifies the base DN used in the search query that is being indexed.
	BaseDN string                       `json:"baseDN"`
	Scope  EnumlocalDbVlvIndexScopeProp `json:"scope"`
	// Specifies the LDAP filter used in the query that is being indexed.
	Filter string `json:"filter"`
	// Specifies the names of the attributes that are used to sort the entries for the query being indexed.
	SortOrder string `json:"sortOrder"`
	// Specifies a unique name for this VLV index.
	Name string `json:"name"`
	// Specifies the number of entry IDs to store in a single sorted set before it must be split.
	MaxBlockSize *int32                            `json:"maxBlockSize,omitempty"`
	CacheMode    *EnumlocalDbVlvIndexCacheModeProp `json:"cacheMode,omitempty"`
}

// NewLocalDbVlvIndexResponse instantiates a new LocalDbVlvIndexResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalDbVlvIndexResponse(id string, baseDN string, scope EnumlocalDbVlvIndexScopeProp, filter string, sortOrder string, name string) *LocalDbVlvIndexResponse {
	this := LocalDbVlvIndexResponse{}
	this.Id = id
	this.BaseDN = baseDN
	this.Scope = scope
	this.Filter = filter
	this.SortOrder = sortOrder
	this.Name = name
	return &this
}

// NewLocalDbVlvIndexResponseWithDefaults instantiates a new LocalDbVlvIndexResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalDbVlvIndexResponseWithDefaults() *LocalDbVlvIndexResponse {
	this := LocalDbVlvIndexResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LocalDbVlvIndexResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbVlvIndexResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LocalDbVlvIndexResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *LocalDbVlvIndexResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *LocalDbVlvIndexResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbVlvIndexResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *LocalDbVlvIndexResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *LocalDbVlvIndexResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *LocalDbVlvIndexResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LocalDbVlvIndexResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LocalDbVlvIndexResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *LocalDbVlvIndexResponse) GetSchemas() []EnumlocalDbVlvIndexSchemaUrn {
	if o == nil || isNil(o.Schemas) {
		var ret []EnumlocalDbVlvIndexSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbVlvIndexResponse) GetSchemasOk() ([]EnumlocalDbVlvIndexSchemaUrn, bool) {
	if o == nil || isNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *LocalDbVlvIndexResponse) HasSchemas() bool {
	if o != nil && !isNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumlocalDbVlvIndexSchemaUrn and assigns it to the Schemas field.
func (o *LocalDbVlvIndexResponse) SetSchemas(v []EnumlocalDbVlvIndexSchemaUrn) {
	o.Schemas = v
}

// GetBaseDN returns the BaseDN field value
func (o *LocalDbVlvIndexResponse) GetBaseDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value
// and a boolean to check if the value has been set.
func (o *LocalDbVlvIndexResponse) GetBaseDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseDN, true
}

// SetBaseDN sets field value
func (o *LocalDbVlvIndexResponse) SetBaseDN(v string) {
	o.BaseDN = v
}

// GetScope returns the Scope field value
func (o *LocalDbVlvIndexResponse) GetScope() EnumlocalDbVlvIndexScopeProp {
	if o == nil {
		var ret EnumlocalDbVlvIndexScopeProp
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *LocalDbVlvIndexResponse) GetScopeOk() (*EnumlocalDbVlvIndexScopeProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *LocalDbVlvIndexResponse) SetScope(v EnumlocalDbVlvIndexScopeProp) {
	o.Scope = v
}

// GetFilter returns the Filter field value
func (o *LocalDbVlvIndexResponse) GetFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *LocalDbVlvIndexResponse) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *LocalDbVlvIndexResponse) SetFilter(v string) {
	o.Filter = v
}

// GetSortOrder returns the SortOrder field value
func (o *LocalDbVlvIndexResponse) GetSortOrder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *LocalDbVlvIndexResponse) GetSortOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *LocalDbVlvIndexResponse) SetSortOrder(v string) {
	o.SortOrder = v
}

// GetName returns the Name field value
func (o *LocalDbVlvIndexResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LocalDbVlvIndexResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LocalDbVlvIndexResponse) SetName(v string) {
	o.Name = v
}

// GetMaxBlockSize returns the MaxBlockSize field value if set, zero value otherwise.
func (o *LocalDbVlvIndexResponse) GetMaxBlockSize() int32 {
	if o == nil || isNil(o.MaxBlockSize) {
		var ret int32
		return ret
	}
	return *o.MaxBlockSize
}

// GetMaxBlockSizeOk returns a tuple with the MaxBlockSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbVlvIndexResponse) GetMaxBlockSizeOk() (*int32, bool) {
	if o == nil || isNil(o.MaxBlockSize) {
		return nil, false
	}
	return o.MaxBlockSize, true
}

// HasMaxBlockSize returns a boolean if a field has been set.
func (o *LocalDbVlvIndexResponse) HasMaxBlockSize() bool {
	if o != nil && !isNil(o.MaxBlockSize) {
		return true
	}

	return false
}

// SetMaxBlockSize gets a reference to the given int32 and assigns it to the MaxBlockSize field.
func (o *LocalDbVlvIndexResponse) SetMaxBlockSize(v int32) {
	o.MaxBlockSize = &v
}

// GetCacheMode returns the CacheMode field value if set, zero value otherwise.
func (o *LocalDbVlvIndexResponse) GetCacheMode() EnumlocalDbVlvIndexCacheModeProp {
	if o == nil || isNil(o.CacheMode) {
		var ret EnumlocalDbVlvIndexCacheModeProp
		return ret
	}
	return *o.CacheMode
}

// GetCacheModeOk returns a tuple with the CacheMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbVlvIndexResponse) GetCacheModeOk() (*EnumlocalDbVlvIndexCacheModeProp, bool) {
	if o == nil || isNil(o.CacheMode) {
		return nil, false
	}
	return o.CacheMode, true
}

// HasCacheMode returns a boolean if a field has been set.
func (o *LocalDbVlvIndexResponse) HasCacheMode() bool {
	if o != nil && !isNil(o.CacheMode) {
		return true
	}

	return false
}

// SetCacheMode gets a reference to the given EnumlocalDbVlvIndexCacheModeProp and assigns it to the CacheMode field.
func (o *LocalDbVlvIndexResponse) SetCacheMode(v EnumlocalDbVlvIndexCacheModeProp) {
	o.CacheMode = &v
}

func (o LocalDbVlvIndexResponse) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["baseDN"] = o.BaseDN
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["filter"] = o.Filter
	}
	if true {
		toSerialize["sortOrder"] = o.SortOrder
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.MaxBlockSize) {
		toSerialize["maxBlockSize"] = o.MaxBlockSize
	}
	if !isNil(o.CacheMode) {
		toSerialize["cacheMode"] = o.CacheMode
	}
	return json.Marshal(toSerialize)
}

type NullableLocalDbVlvIndexResponse struct {
	value *LocalDbVlvIndexResponse
	isSet bool
}

func (v NullableLocalDbVlvIndexResponse) Get() *LocalDbVlvIndexResponse {
	return v.value
}

func (v *NullableLocalDbVlvIndexResponse) Set(val *LocalDbVlvIndexResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalDbVlvIndexResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalDbVlvIndexResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalDbVlvIndexResponse(val *LocalDbVlvIndexResponse) *NullableLocalDbVlvIndexResponse {
	return &NullableLocalDbVlvIndexResponse{value: val, isSet: true}
}

func (v NullableLocalDbVlvIndexResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalDbVlvIndexResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}