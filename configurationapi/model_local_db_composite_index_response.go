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

// checks if the LocalDbCompositeIndexResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalDbCompositeIndexResponse{}

// LocalDbCompositeIndexResponse struct for LocalDbCompositeIndexResponse
type LocalDbCompositeIndexResponse struct {
	Schemas []EnumlocalDbCompositeIndexSchemaUrn `json:"schemas,omitempty"`
	// A description for this Local DB Composite Index
	Description *string `json:"description,omitempty"`
	// A filter pattern that identifies which entries to include in the index.
	IndexFilterPattern string `json:"indexFilterPattern"`
	// An optional base DN pattern that identifies portions of the DIT in which entries to index may exist.
	IndexBaseDNPattern *string `json:"indexBaseDNPattern,omitempty"`
	// The maximum number of entries that any single index key will be allowed to match before the server stops maintaining the ID set for that index key.
	IndexEntryLimit *int64 `json:"indexEntryLimit,omitempty"`
	// Indicates whether the server should load the contents of this index into memory when the backend is being opened.
	PrimeIndex *bool `json:"primeIndex,omitempty"`
	// Indicates whether to only prime the internal nodes of the index database, rather than priming both internal and leaf nodes.
	PrimeInternalNodesOnly                        *bool                                              `json:"primeInternalNodesOnly,omitempty"`
	CacheMode                                     *EnumlocalDbCompositeIndexCacheModeProp            `json:"cacheMode,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Local DB Composite Index
	Id string `json:"id"`
}

// NewLocalDbCompositeIndexResponse instantiates a new LocalDbCompositeIndexResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalDbCompositeIndexResponse(indexFilterPattern string, id string) *LocalDbCompositeIndexResponse {
	this := LocalDbCompositeIndexResponse{}
	this.IndexFilterPattern = indexFilterPattern
	this.Id = id
	return &this
}

// NewLocalDbCompositeIndexResponseWithDefaults instantiates a new LocalDbCompositeIndexResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalDbCompositeIndexResponseWithDefaults() *LocalDbCompositeIndexResponse {
	this := LocalDbCompositeIndexResponse{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *LocalDbCompositeIndexResponse) GetSchemas() []EnumlocalDbCompositeIndexSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumlocalDbCompositeIndexSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbCompositeIndexResponse) GetSchemasOk() ([]EnumlocalDbCompositeIndexSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *LocalDbCompositeIndexResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumlocalDbCompositeIndexSchemaUrn and assigns it to the Schemas field.
func (o *LocalDbCompositeIndexResponse) SetSchemas(v []EnumlocalDbCompositeIndexSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LocalDbCompositeIndexResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbCompositeIndexResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LocalDbCompositeIndexResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LocalDbCompositeIndexResponse) SetDescription(v string) {
	o.Description = &v
}

// GetIndexFilterPattern returns the IndexFilterPattern field value
func (o *LocalDbCompositeIndexResponse) GetIndexFilterPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexFilterPattern
}

// GetIndexFilterPatternOk returns a tuple with the IndexFilterPattern field value
// and a boolean to check if the value has been set.
func (o *LocalDbCompositeIndexResponse) GetIndexFilterPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexFilterPattern, true
}

// SetIndexFilterPattern sets field value
func (o *LocalDbCompositeIndexResponse) SetIndexFilterPattern(v string) {
	o.IndexFilterPattern = v
}

// GetIndexBaseDNPattern returns the IndexBaseDNPattern field value if set, zero value otherwise.
func (o *LocalDbCompositeIndexResponse) GetIndexBaseDNPattern() string {
	if o == nil || IsNil(o.IndexBaseDNPattern) {
		var ret string
		return ret
	}
	return *o.IndexBaseDNPattern
}

// GetIndexBaseDNPatternOk returns a tuple with the IndexBaseDNPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbCompositeIndexResponse) GetIndexBaseDNPatternOk() (*string, bool) {
	if o == nil || IsNil(o.IndexBaseDNPattern) {
		return nil, false
	}
	return o.IndexBaseDNPattern, true
}

// HasIndexBaseDNPattern returns a boolean if a field has been set.
func (o *LocalDbCompositeIndexResponse) HasIndexBaseDNPattern() bool {
	if o != nil && !IsNil(o.IndexBaseDNPattern) {
		return true
	}

	return false
}

// SetIndexBaseDNPattern gets a reference to the given string and assigns it to the IndexBaseDNPattern field.
func (o *LocalDbCompositeIndexResponse) SetIndexBaseDNPattern(v string) {
	o.IndexBaseDNPattern = &v
}

// GetIndexEntryLimit returns the IndexEntryLimit field value if set, zero value otherwise.
func (o *LocalDbCompositeIndexResponse) GetIndexEntryLimit() int64 {
	if o == nil || IsNil(o.IndexEntryLimit) {
		var ret int64
		return ret
	}
	return *o.IndexEntryLimit
}

// GetIndexEntryLimitOk returns a tuple with the IndexEntryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbCompositeIndexResponse) GetIndexEntryLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.IndexEntryLimit) {
		return nil, false
	}
	return o.IndexEntryLimit, true
}

// HasIndexEntryLimit returns a boolean if a field has been set.
func (o *LocalDbCompositeIndexResponse) HasIndexEntryLimit() bool {
	if o != nil && !IsNil(o.IndexEntryLimit) {
		return true
	}

	return false
}

// SetIndexEntryLimit gets a reference to the given int64 and assigns it to the IndexEntryLimit field.
func (o *LocalDbCompositeIndexResponse) SetIndexEntryLimit(v int64) {
	o.IndexEntryLimit = &v
}

// GetPrimeIndex returns the PrimeIndex field value if set, zero value otherwise.
func (o *LocalDbCompositeIndexResponse) GetPrimeIndex() bool {
	if o == nil || IsNil(o.PrimeIndex) {
		var ret bool
		return ret
	}
	return *o.PrimeIndex
}

// GetPrimeIndexOk returns a tuple with the PrimeIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbCompositeIndexResponse) GetPrimeIndexOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimeIndex) {
		return nil, false
	}
	return o.PrimeIndex, true
}

// HasPrimeIndex returns a boolean if a field has been set.
func (o *LocalDbCompositeIndexResponse) HasPrimeIndex() bool {
	if o != nil && !IsNil(o.PrimeIndex) {
		return true
	}

	return false
}

// SetPrimeIndex gets a reference to the given bool and assigns it to the PrimeIndex field.
func (o *LocalDbCompositeIndexResponse) SetPrimeIndex(v bool) {
	o.PrimeIndex = &v
}

// GetPrimeInternalNodesOnly returns the PrimeInternalNodesOnly field value if set, zero value otherwise.
func (o *LocalDbCompositeIndexResponse) GetPrimeInternalNodesOnly() bool {
	if o == nil || IsNil(o.PrimeInternalNodesOnly) {
		var ret bool
		return ret
	}
	return *o.PrimeInternalNodesOnly
}

// GetPrimeInternalNodesOnlyOk returns a tuple with the PrimeInternalNodesOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbCompositeIndexResponse) GetPrimeInternalNodesOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimeInternalNodesOnly) {
		return nil, false
	}
	return o.PrimeInternalNodesOnly, true
}

// HasPrimeInternalNodesOnly returns a boolean if a field has been set.
func (o *LocalDbCompositeIndexResponse) HasPrimeInternalNodesOnly() bool {
	if o != nil && !IsNil(o.PrimeInternalNodesOnly) {
		return true
	}

	return false
}

// SetPrimeInternalNodesOnly gets a reference to the given bool and assigns it to the PrimeInternalNodesOnly field.
func (o *LocalDbCompositeIndexResponse) SetPrimeInternalNodesOnly(v bool) {
	o.PrimeInternalNodesOnly = &v
}

// GetCacheMode returns the CacheMode field value if set, zero value otherwise.
func (o *LocalDbCompositeIndexResponse) GetCacheMode() EnumlocalDbCompositeIndexCacheModeProp {
	if o == nil || IsNil(o.CacheMode) {
		var ret EnumlocalDbCompositeIndexCacheModeProp
		return ret
	}
	return *o.CacheMode
}

// GetCacheModeOk returns a tuple with the CacheMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbCompositeIndexResponse) GetCacheModeOk() (*EnumlocalDbCompositeIndexCacheModeProp, bool) {
	if o == nil || IsNil(o.CacheMode) {
		return nil, false
	}
	return o.CacheMode, true
}

// HasCacheMode returns a boolean if a field has been set.
func (o *LocalDbCompositeIndexResponse) HasCacheMode() bool {
	if o != nil && !IsNil(o.CacheMode) {
		return true
	}

	return false
}

// SetCacheMode gets a reference to the given EnumlocalDbCompositeIndexCacheModeProp and assigns it to the CacheMode field.
func (o *LocalDbCompositeIndexResponse) SetCacheMode(v EnumlocalDbCompositeIndexCacheModeProp) {
	o.CacheMode = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LocalDbCompositeIndexResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbCompositeIndexResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LocalDbCompositeIndexResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *LocalDbCompositeIndexResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *LocalDbCompositeIndexResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbCompositeIndexResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *LocalDbCompositeIndexResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *LocalDbCompositeIndexResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *LocalDbCompositeIndexResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LocalDbCompositeIndexResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LocalDbCompositeIndexResponse) SetId(v string) {
	o.Id = v
}

func (o LocalDbCompositeIndexResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalDbCompositeIndexResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["indexFilterPattern"] = o.IndexFilterPattern
	if !IsNil(o.IndexBaseDNPattern) {
		toSerialize["indexBaseDNPattern"] = o.IndexBaseDNPattern
	}
	if !IsNil(o.IndexEntryLimit) {
		toSerialize["indexEntryLimit"] = o.IndexEntryLimit
	}
	if !IsNil(o.PrimeIndex) {
		toSerialize["primeIndex"] = o.PrimeIndex
	}
	if !IsNil(o.PrimeInternalNodesOnly) {
		toSerialize["primeInternalNodesOnly"] = o.PrimeInternalNodesOnly
	}
	if !IsNil(o.CacheMode) {
		toSerialize["cacheMode"] = o.CacheMode
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

type NullableLocalDbCompositeIndexResponse struct {
	value *LocalDbCompositeIndexResponse
	isSet bool
}

func (v NullableLocalDbCompositeIndexResponse) Get() *LocalDbCompositeIndexResponse {
	return v.value
}

func (v *NullableLocalDbCompositeIndexResponse) Set(val *LocalDbCompositeIndexResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalDbCompositeIndexResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalDbCompositeIndexResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalDbCompositeIndexResponse(val *LocalDbCompositeIndexResponse) *NullableLocalDbCompositeIndexResponse {
	return &NullableLocalDbCompositeIndexResponse{value: val, isSet: true}
}

func (v NullableLocalDbCompositeIndexResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalDbCompositeIndexResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
