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

// checks if the LocalDbIndexResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalDbIndexResponse{}

// LocalDbIndexResponse struct for LocalDbIndexResponse
type LocalDbIndexResponse struct {
	// Name of the Local DB Index
	Id      string                      `json:"id"`
	Schemas []EnumlocalDbIndexSchemaUrn `json:"schemas,omitempty"`
	// Specifies the name of the attribute for which the index is to be maintained.
	Attribute string `json:"attribute"`
	// Specifies the maximum number of entries that are allowed to match a given index key before that particular index key is no longer maintained.
	IndexEntryLimit *int64 `json:"indexEntryLimit,omitempty"`
	// Specifies, for substring indexes, the maximum number of entries that are allowed to match a given index key before that particular index key is no longer maintained. Setting a large limit can dramatically increase the database size on disk and have a big impact on server performance if the indexed attribute is modified frequently. When a very large limit is required, creating a dedicated composite index with an index-filter-pattern of (attr=*?*) will give the best balance between search and update performance.
	SubstringIndexEntryLimit *int64 `json:"substringIndexEntryLimit,omitempty"`
	// Indicates whether to continue to maintain a count of the number of matching entries for an index key even after that count exceeds the index entry limit.
	MaintainMatchCountForKeysExceedingEntryLimit *bool                           `json:"maintainMatchCountForKeysExceedingEntryLimit,omitempty"`
	IndexType                                    []EnumlocalDbIndexIndexTypeProp `json:"indexType"`
	// The length of substrings in a substring index.
	SubstringLength *int64 `json:"substringLength,omitempty"`
	// If this option is enabled and this index's backend is configured to prime indexes, then this index will be loaded at startup.
	PrimeIndex *bool `json:"primeIndex,omitempty"`
	// If this option is enabled and this index's backend is configured to prime indexes using the preload method, then only the internal database nodes (i.e., the database keys but not values) should be primed when the backend is initialized.
	PrimeInternalNodesOnly *bool `json:"primeInternalNodesOnly,omitempty"`
	// A search filter that may be used in conjunction with an equality component for the associated attribute type. If an equality index filter is defined, then an additional equality index will be maintained for the associated attribute, but only for entries which match the provided filter. Further, the index will be used only for searches containing an equality component with the associated attribute type ANDed with this filter.
	EqualityIndexFilter []string `json:"equalityIndexFilter,omitempty"`
	// Indicates whether to maintain a separate equality index for the associated attribute without any filter, in addition to maintaining an index for each equality index filter that is defined. If this is false, then the attribute will not be indexed for equality by itself but only in conjunction with the defined equality index filters.
	MaintainEqualityIndexWithoutFilter            *bool                                              `json:"maintainEqualityIndexWithoutFilter,omitempty"`
	CacheMode                                     *EnumlocalDbIndexCacheModeProp                     `json:"cacheMode,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewLocalDbIndexResponse instantiates a new LocalDbIndexResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalDbIndexResponse(id string, attribute string, indexType []EnumlocalDbIndexIndexTypeProp) *LocalDbIndexResponse {
	this := LocalDbIndexResponse{}
	this.Id = id
	this.Attribute = attribute
	this.IndexType = indexType
	return &this
}

// NewLocalDbIndexResponseWithDefaults instantiates a new LocalDbIndexResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalDbIndexResponseWithDefaults() *LocalDbIndexResponse {
	this := LocalDbIndexResponse{}
	return &this
}

// GetId returns the Id field value
func (o *LocalDbIndexResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LocalDbIndexResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LocalDbIndexResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *LocalDbIndexResponse) GetSchemas() []EnumlocalDbIndexSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumlocalDbIndexSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbIndexResponse) GetSchemasOk() ([]EnumlocalDbIndexSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *LocalDbIndexResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumlocalDbIndexSchemaUrn and assigns it to the Schemas field.
func (o *LocalDbIndexResponse) SetSchemas(v []EnumlocalDbIndexSchemaUrn) {
	o.Schemas = v
}

// GetAttribute returns the Attribute field value
func (o *LocalDbIndexResponse) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *LocalDbIndexResponse) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *LocalDbIndexResponse) SetAttribute(v string) {
	o.Attribute = v
}

// GetIndexEntryLimit returns the IndexEntryLimit field value if set, zero value otherwise.
func (o *LocalDbIndexResponse) GetIndexEntryLimit() int64 {
	if o == nil || IsNil(o.IndexEntryLimit) {
		var ret int64
		return ret
	}
	return *o.IndexEntryLimit
}

// GetIndexEntryLimitOk returns a tuple with the IndexEntryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbIndexResponse) GetIndexEntryLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.IndexEntryLimit) {
		return nil, false
	}
	return o.IndexEntryLimit, true
}

// HasIndexEntryLimit returns a boolean if a field has been set.
func (o *LocalDbIndexResponse) HasIndexEntryLimit() bool {
	if o != nil && !IsNil(o.IndexEntryLimit) {
		return true
	}

	return false
}

// SetIndexEntryLimit gets a reference to the given int64 and assigns it to the IndexEntryLimit field.
func (o *LocalDbIndexResponse) SetIndexEntryLimit(v int64) {
	o.IndexEntryLimit = &v
}

// GetSubstringIndexEntryLimit returns the SubstringIndexEntryLimit field value if set, zero value otherwise.
func (o *LocalDbIndexResponse) GetSubstringIndexEntryLimit() int64 {
	if o == nil || IsNil(o.SubstringIndexEntryLimit) {
		var ret int64
		return ret
	}
	return *o.SubstringIndexEntryLimit
}

// GetSubstringIndexEntryLimitOk returns a tuple with the SubstringIndexEntryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbIndexResponse) GetSubstringIndexEntryLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.SubstringIndexEntryLimit) {
		return nil, false
	}
	return o.SubstringIndexEntryLimit, true
}

// HasSubstringIndexEntryLimit returns a boolean if a field has been set.
func (o *LocalDbIndexResponse) HasSubstringIndexEntryLimit() bool {
	if o != nil && !IsNil(o.SubstringIndexEntryLimit) {
		return true
	}

	return false
}

// SetSubstringIndexEntryLimit gets a reference to the given int64 and assigns it to the SubstringIndexEntryLimit field.
func (o *LocalDbIndexResponse) SetSubstringIndexEntryLimit(v int64) {
	o.SubstringIndexEntryLimit = &v
}

// GetMaintainMatchCountForKeysExceedingEntryLimit returns the MaintainMatchCountForKeysExceedingEntryLimit field value if set, zero value otherwise.
func (o *LocalDbIndexResponse) GetMaintainMatchCountForKeysExceedingEntryLimit() bool {
	if o == nil || IsNil(o.MaintainMatchCountForKeysExceedingEntryLimit) {
		var ret bool
		return ret
	}
	return *o.MaintainMatchCountForKeysExceedingEntryLimit
}

// GetMaintainMatchCountForKeysExceedingEntryLimitOk returns a tuple with the MaintainMatchCountForKeysExceedingEntryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbIndexResponse) GetMaintainMatchCountForKeysExceedingEntryLimitOk() (*bool, bool) {
	if o == nil || IsNil(o.MaintainMatchCountForKeysExceedingEntryLimit) {
		return nil, false
	}
	return o.MaintainMatchCountForKeysExceedingEntryLimit, true
}

// HasMaintainMatchCountForKeysExceedingEntryLimit returns a boolean if a field has been set.
func (o *LocalDbIndexResponse) HasMaintainMatchCountForKeysExceedingEntryLimit() bool {
	if o != nil && !IsNil(o.MaintainMatchCountForKeysExceedingEntryLimit) {
		return true
	}

	return false
}

// SetMaintainMatchCountForKeysExceedingEntryLimit gets a reference to the given bool and assigns it to the MaintainMatchCountForKeysExceedingEntryLimit field.
func (o *LocalDbIndexResponse) SetMaintainMatchCountForKeysExceedingEntryLimit(v bool) {
	o.MaintainMatchCountForKeysExceedingEntryLimit = &v
}

// GetIndexType returns the IndexType field value
func (o *LocalDbIndexResponse) GetIndexType() []EnumlocalDbIndexIndexTypeProp {
	if o == nil {
		var ret []EnumlocalDbIndexIndexTypeProp
		return ret
	}

	return o.IndexType
}

// GetIndexTypeOk returns a tuple with the IndexType field value
// and a boolean to check if the value has been set.
func (o *LocalDbIndexResponse) GetIndexTypeOk() ([]EnumlocalDbIndexIndexTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexType, true
}

// SetIndexType sets field value
func (o *LocalDbIndexResponse) SetIndexType(v []EnumlocalDbIndexIndexTypeProp) {
	o.IndexType = v
}

// GetSubstringLength returns the SubstringLength field value if set, zero value otherwise.
func (o *LocalDbIndexResponse) GetSubstringLength() int64 {
	if o == nil || IsNil(o.SubstringLength) {
		var ret int64
		return ret
	}
	return *o.SubstringLength
}

// GetSubstringLengthOk returns a tuple with the SubstringLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbIndexResponse) GetSubstringLengthOk() (*int64, bool) {
	if o == nil || IsNil(o.SubstringLength) {
		return nil, false
	}
	return o.SubstringLength, true
}

// HasSubstringLength returns a boolean if a field has been set.
func (o *LocalDbIndexResponse) HasSubstringLength() bool {
	if o != nil && !IsNil(o.SubstringLength) {
		return true
	}

	return false
}

// SetSubstringLength gets a reference to the given int64 and assigns it to the SubstringLength field.
func (o *LocalDbIndexResponse) SetSubstringLength(v int64) {
	o.SubstringLength = &v
}

// GetPrimeIndex returns the PrimeIndex field value if set, zero value otherwise.
func (o *LocalDbIndexResponse) GetPrimeIndex() bool {
	if o == nil || IsNil(o.PrimeIndex) {
		var ret bool
		return ret
	}
	return *o.PrimeIndex
}

// GetPrimeIndexOk returns a tuple with the PrimeIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbIndexResponse) GetPrimeIndexOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimeIndex) {
		return nil, false
	}
	return o.PrimeIndex, true
}

// HasPrimeIndex returns a boolean if a field has been set.
func (o *LocalDbIndexResponse) HasPrimeIndex() bool {
	if o != nil && !IsNil(o.PrimeIndex) {
		return true
	}

	return false
}

// SetPrimeIndex gets a reference to the given bool and assigns it to the PrimeIndex field.
func (o *LocalDbIndexResponse) SetPrimeIndex(v bool) {
	o.PrimeIndex = &v
}

// GetPrimeInternalNodesOnly returns the PrimeInternalNodesOnly field value if set, zero value otherwise.
func (o *LocalDbIndexResponse) GetPrimeInternalNodesOnly() bool {
	if o == nil || IsNil(o.PrimeInternalNodesOnly) {
		var ret bool
		return ret
	}
	return *o.PrimeInternalNodesOnly
}

// GetPrimeInternalNodesOnlyOk returns a tuple with the PrimeInternalNodesOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbIndexResponse) GetPrimeInternalNodesOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimeInternalNodesOnly) {
		return nil, false
	}
	return o.PrimeInternalNodesOnly, true
}

// HasPrimeInternalNodesOnly returns a boolean if a field has been set.
func (o *LocalDbIndexResponse) HasPrimeInternalNodesOnly() bool {
	if o != nil && !IsNil(o.PrimeInternalNodesOnly) {
		return true
	}

	return false
}

// SetPrimeInternalNodesOnly gets a reference to the given bool and assigns it to the PrimeInternalNodesOnly field.
func (o *LocalDbIndexResponse) SetPrimeInternalNodesOnly(v bool) {
	o.PrimeInternalNodesOnly = &v
}

// GetEqualityIndexFilter returns the EqualityIndexFilter field value if set, zero value otherwise.
func (o *LocalDbIndexResponse) GetEqualityIndexFilter() []string {
	if o == nil || IsNil(o.EqualityIndexFilter) {
		var ret []string
		return ret
	}
	return o.EqualityIndexFilter
}

// GetEqualityIndexFilterOk returns a tuple with the EqualityIndexFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbIndexResponse) GetEqualityIndexFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.EqualityIndexFilter) {
		return nil, false
	}
	return o.EqualityIndexFilter, true
}

// HasEqualityIndexFilter returns a boolean if a field has been set.
func (o *LocalDbIndexResponse) HasEqualityIndexFilter() bool {
	if o != nil && !IsNil(o.EqualityIndexFilter) {
		return true
	}

	return false
}

// SetEqualityIndexFilter gets a reference to the given []string and assigns it to the EqualityIndexFilter field.
func (o *LocalDbIndexResponse) SetEqualityIndexFilter(v []string) {
	o.EqualityIndexFilter = v
}

// GetMaintainEqualityIndexWithoutFilter returns the MaintainEqualityIndexWithoutFilter field value if set, zero value otherwise.
func (o *LocalDbIndexResponse) GetMaintainEqualityIndexWithoutFilter() bool {
	if o == nil || IsNil(o.MaintainEqualityIndexWithoutFilter) {
		var ret bool
		return ret
	}
	return *o.MaintainEqualityIndexWithoutFilter
}

// GetMaintainEqualityIndexWithoutFilterOk returns a tuple with the MaintainEqualityIndexWithoutFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbIndexResponse) GetMaintainEqualityIndexWithoutFilterOk() (*bool, bool) {
	if o == nil || IsNil(o.MaintainEqualityIndexWithoutFilter) {
		return nil, false
	}
	return o.MaintainEqualityIndexWithoutFilter, true
}

// HasMaintainEqualityIndexWithoutFilter returns a boolean if a field has been set.
func (o *LocalDbIndexResponse) HasMaintainEqualityIndexWithoutFilter() bool {
	if o != nil && !IsNil(o.MaintainEqualityIndexWithoutFilter) {
		return true
	}

	return false
}

// SetMaintainEqualityIndexWithoutFilter gets a reference to the given bool and assigns it to the MaintainEqualityIndexWithoutFilter field.
func (o *LocalDbIndexResponse) SetMaintainEqualityIndexWithoutFilter(v bool) {
	o.MaintainEqualityIndexWithoutFilter = &v
}

// GetCacheMode returns the CacheMode field value if set, zero value otherwise.
func (o *LocalDbIndexResponse) GetCacheMode() EnumlocalDbIndexCacheModeProp {
	if o == nil || IsNil(o.CacheMode) {
		var ret EnumlocalDbIndexCacheModeProp
		return ret
	}
	return *o.CacheMode
}

// GetCacheModeOk returns a tuple with the CacheMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbIndexResponse) GetCacheModeOk() (*EnumlocalDbIndexCacheModeProp, bool) {
	if o == nil || IsNil(o.CacheMode) {
		return nil, false
	}
	return o.CacheMode, true
}

// HasCacheMode returns a boolean if a field has been set.
func (o *LocalDbIndexResponse) HasCacheMode() bool {
	if o != nil && !IsNil(o.CacheMode) {
		return true
	}

	return false
}

// SetCacheMode gets a reference to the given EnumlocalDbIndexCacheModeProp and assigns it to the CacheMode field.
func (o *LocalDbIndexResponse) SetCacheMode(v EnumlocalDbIndexCacheModeProp) {
	o.CacheMode = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LocalDbIndexResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbIndexResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LocalDbIndexResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *LocalDbIndexResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *LocalDbIndexResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDbIndexResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *LocalDbIndexResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *LocalDbIndexResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o LocalDbIndexResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalDbIndexResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	toSerialize["attribute"] = o.Attribute
	if !IsNil(o.IndexEntryLimit) {
		toSerialize["indexEntryLimit"] = o.IndexEntryLimit
	}
	if !IsNil(o.SubstringIndexEntryLimit) {
		toSerialize["substringIndexEntryLimit"] = o.SubstringIndexEntryLimit
	}
	if !IsNil(o.MaintainMatchCountForKeysExceedingEntryLimit) {
		toSerialize["maintainMatchCountForKeysExceedingEntryLimit"] = o.MaintainMatchCountForKeysExceedingEntryLimit
	}
	toSerialize["indexType"] = o.IndexType
	if !IsNil(o.SubstringLength) {
		toSerialize["substringLength"] = o.SubstringLength
	}
	if !IsNil(o.PrimeIndex) {
		toSerialize["primeIndex"] = o.PrimeIndex
	}
	if !IsNil(o.PrimeInternalNodesOnly) {
		toSerialize["primeInternalNodesOnly"] = o.PrimeInternalNodesOnly
	}
	if !IsNil(o.EqualityIndexFilter) {
		toSerialize["equalityIndexFilter"] = o.EqualityIndexFilter
	}
	if !IsNil(o.MaintainEqualityIndexWithoutFilter) {
		toSerialize["maintainEqualityIndexWithoutFilter"] = o.MaintainEqualityIndexWithoutFilter
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
	return toSerialize, nil
}

type NullableLocalDbIndexResponse struct {
	value *LocalDbIndexResponse
	isSet bool
}

func (v NullableLocalDbIndexResponse) Get() *LocalDbIndexResponse {
	return v.value
}

func (v *NullableLocalDbIndexResponse) Set(val *LocalDbIndexResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalDbIndexResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalDbIndexResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalDbIndexResponse(val *LocalDbIndexResponse) *NullableLocalDbIndexResponse {
	return &NullableLocalDbIndexResponse{value: val, isSet: true}
}

func (v NullableLocalDbIndexResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalDbIndexResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
