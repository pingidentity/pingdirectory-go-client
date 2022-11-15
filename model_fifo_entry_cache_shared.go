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

// FifoEntryCacheShared struct for FifoEntryCacheShared
type FifoEntryCacheShared struct {
	Schemas []EnumfifoEntryCacheSchemaUrn `json:"schemas,omitempty"`
	// Specifies the maximum amount of memory, as a percentage of the total maximum JVM heap size, that this cache should occupy when full. If the amount of memory the cache is using is greater than this amount, then an attempt to put a new entry in the cache will be ignored and will cause the oldest entry to be purged.
	MaxMemoryPercent *int32 `json:"maxMemoryPercent,omitempty"`
	// Specifies the maximum number of entries that will be allowed in the cache. Once the cache reaches this size, then adding new entries will cause existing entries to be purged, starting with the oldest.
	MaxEntries *int32 `json:"maxEntries,omitempty"`
	// Specifies that the cache should only store entries which are accessed much more frequently than the average entry. The cache will observe attempts to place entries in the cache and compare an entry's accesses to the average entry's.
	OnlyCacheFrequentlyAccessed *bool `json:"onlyCacheFrequentlyAccessed,omitempty"`
	IncludeFilter []string `json:"includeFilter,omitempty"`
	ExcludeFilter []string `json:"excludeFilter,omitempty"`
	// Specifies the minimum number of attribute values (optionally across a specified subset of attributes as defined in the min-cache-entry-attributes property) for entries that should be held in the cache. Entries with fewer than this number of attribute values will be excluded from the cache.
	MinCacheEntryValueCount *int32 `json:"minCacheEntryValueCount,omitempty"`
	MinCacheEntryAttribute []string `json:"minCacheEntryAttribute,omitempty"`
	// A description for this Entry Cache
	Description *string `json:"description,omitempty"`
	// Indicates whether the Entry Cache is enabled.
	Enabled bool `json:"enabled"`
	// Specifies the cache level in the cache order if more than one instance of the cache is configured.
	CacheLevel int32 `json:"cacheLevel"`
	// Indicates whether the entry cache should be updated with entries that have been returned to the client during the course of processing an unindexed search.
	CacheUnindexedSearchResults *bool `json:"cacheUnindexedSearchResults,omitempty"`
}

// NewFifoEntryCacheShared instantiates a new FifoEntryCacheShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFifoEntryCacheShared(enabled bool, cacheLevel int32) *FifoEntryCacheShared {
	this := FifoEntryCacheShared{}
	this.Enabled = enabled
	this.CacheLevel = cacheLevel
	return &this
}

// NewFifoEntryCacheSharedWithDefaults instantiates a new FifoEntryCacheShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFifoEntryCacheSharedWithDefaults() *FifoEntryCacheShared {
	this := FifoEntryCacheShared{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *FifoEntryCacheShared) GetSchemas() []EnumfifoEntryCacheSchemaUrn {
	if o == nil || isNil(o.Schemas) {
		var ret []EnumfifoEntryCacheSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FifoEntryCacheShared) GetSchemasOk() ([]EnumfifoEntryCacheSchemaUrn, bool) {
	if o == nil || isNil(o.Schemas) {
    return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *FifoEntryCacheShared) HasSchemas() bool {
	if o != nil && !isNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumfifoEntryCacheSchemaUrn and assigns it to the Schemas field.
func (o *FifoEntryCacheShared) SetSchemas(v []EnumfifoEntryCacheSchemaUrn) {
	o.Schemas = v
}

// GetMaxMemoryPercent returns the MaxMemoryPercent field value if set, zero value otherwise.
func (o *FifoEntryCacheShared) GetMaxMemoryPercent() int32 {
	if o == nil || isNil(o.MaxMemoryPercent) {
		var ret int32
		return ret
	}
	return *o.MaxMemoryPercent
}

// GetMaxMemoryPercentOk returns a tuple with the MaxMemoryPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FifoEntryCacheShared) GetMaxMemoryPercentOk() (*int32, bool) {
	if o == nil || isNil(o.MaxMemoryPercent) {
    return nil, false
	}
	return o.MaxMemoryPercent, true
}

// HasMaxMemoryPercent returns a boolean if a field has been set.
func (o *FifoEntryCacheShared) HasMaxMemoryPercent() bool {
	if o != nil && !isNil(o.MaxMemoryPercent) {
		return true
	}

	return false
}

// SetMaxMemoryPercent gets a reference to the given int32 and assigns it to the MaxMemoryPercent field.
func (o *FifoEntryCacheShared) SetMaxMemoryPercent(v int32) {
	o.MaxMemoryPercent = &v
}

// GetMaxEntries returns the MaxEntries field value if set, zero value otherwise.
func (o *FifoEntryCacheShared) GetMaxEntries() int32 {
	if o == nil || isNil(o.MaxEntries) {
		var ret int32
		return ret
	}
	return *o.MaxEntries
}

// GetMaxEntriesOk returns a tuple with the MaxEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FifoEntryCacheShared) GetMaxEntriesOk() (*int32, bool) {
	if o == nil || isNil(o.MaxEntries) {
    return nil, false
	}
	return o.MaxEntries, true
}

// HasMaxEntries returns a boolean if a field has been set.
func (o *FifoEntryCacheShared) HasMaxEntries() bool {
	if o != nil && !isNil(o.MaxEntries) {
		return true
	}

	return false
}

// SetMaxEntries gets a reference to the given int32 and assigns it to the MaxEntries field.
func (o *FifoEntryCacheShared) SetMaxEntries(v int32) {
	o.MaxEntries = &v
}

// GetOnlyCacheFrequentlyAccessed returns the OnlyCacheFrequentlyAccessed field value if set, zero value otherwise.
func (o *FifoEntryCacheShared) GetOnlyCacheFrequentlyAccessed() bool {
	if o == nil || isNil(o.OnlyCacheFrequentlyAccessed) {
		var ret bool
		return ret
	}
	return *o.OnlyCacheFrequentlyAccessed
}

// GetOnlyCacheFrequentlyAccessedOk returns a tuple with the OnlyCacheFrequentlyAccessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FifoEntryCacheShared) GetOnlyCacheFrequentlyAccessedOk() (*bool, bool) {
	if o == nil || isNil(o.OnlyCacheFrequentlyAccessed) {
    return nil, false
	}
	return o.OnlyCacheFrequentlyAccessed, true
}

// HasOnlyCacheFrequentlyAccessed returns a boolean if a field has been set.
func (o *FifoEntryCacheShared) HasOnlyCacheFrequentlyAccessed() bool {
	if o != nil && !isNil(o.OnlyCacheFrequentlyAccessed) {
		return true
	}

	return false
}

// SetOnlyCacheFrequentlyAccessed gets a reference to the given bool and assigns it to the OnlyCacheFrequentlyAccessed field.
func (o *FifoEntryCacheShared) SetOnlyCacheFrequentlyAccessed(v bool) {
	o.OnlyCacheFrequentlyAccessed = &v
}

// GetIncludeFilter returns the IncludeFilter field value if set, zero value otherwise.
func (o *FifoEntryCacheShared) GetIncludeFilter() []string {
	if o == nil || isNil(o.IncludeFilter) {
		var ret []string
		return ret
	}
	return o.IncludeFilter
}

// GetIncludeFilterOk returns a tuple with the IncludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FifoEntryCacheShared) GetIncludeFilterOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeFilter) {
    return nil, false
	}
	return o.IncludeFilter, true
}

// HasIncludeFilter returns a boolean if a field has been set.
func (o *FifoEntryCacheShared) HasIncludeFilter() bool {
	if o != nil && !isNil(o.IncludeFilter) {
		return true
	}

	return false
}

// SetIncludeFilter gets a reference to the given []string and assigns it to the IncludeFilter field.
func (o *FifoEntryCacheShared) SetIncludeFilter(v []string) {
	o.IncludeFilter = v
}

// GetExcludeFilter returns the ExcludeFilter field value if set, zero value otherwise.
func (o *FifoEntryCacheShared) GetExcludeFilter() []string {
	if o == nil || isNil(o.ExcludeFilter) {
		var ret []string
		return ret
	}
	return o.ExcludeFilter
}

// GetExcludeFilterOk returns a tuple with the ExcludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FifoEntryCacheShared) GetExcludeFilterOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludeFilter) {
    return nil, false
	}
	return o.ExcludeFilter, true
}

// HasExcludeFilter returns a boolean if a field has been set.
func (o *FifoEntryCacheShared) HasExcludeFilter() bool {
	if o != nil && !isNil(o.ExcludeFilter) {
		return true
	}

	return false
}

// SetExcludeFilter gets a reference to the given []string and assigns it to the ExcludeFilter field.
func (o *FifoEntryCacheShared) SetExcludeFilter(v []string) {
	o.ExcludeFilter = v
}

// GetMinCacheEntryValueCount returns the MinCacheEntryValueCount field value if set, zero value otherwise.
func (o *FifoEntryCacheShared) GetMinCacheEntryValueCount() int32 {
	if o == nil || isNil(o.MinCacheEntryValueCount) {
		var ret int32
		return ret
	}
	return *o.MinCacheEntryValueCount
}

// GetMinCacheEntryValueCountOk returns a tuple with the MinCacheEntryValueCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FifoEntryCacheShared) GetMinCacheEntryValueCountOk() (*int32, bool) {
	if o == nil || isNil(o.MinCacheEntryValueCount) {
    return nil, false
	}
	return o.MinCacheEntryValueCount, true
}

// HasMinCacheEntryValueCount returns a boolean if a field has been set.
func (o *FifoEntryCacheShared) HasMinCacheEntryValueCount() bool {
	if o != nil && !isNil(o.MinCacheEntryValueCount) {
		return true
	}

	return false
}

// SetMinCacheEntryValueCount gets a reference to the given int32 and assigns it to the MinCacheEntryValueCount field.
func (o *FifoEntryCacheShared) SetMinCacheEntryValueCount(v int32) {
	o.MinCacheEntryValueCount = &v
}

// GetMinCacheEntryAttribute returns the MinCacheEntryAttribute field value if set, zero value otherwise.
func (o *FifoEntryCacheShared) GetMinCacheEntryAttribute() []string {
	if o == nil || isNil(o.MinCacheEntryAttribute) {
		var ret []string
		return ret
	}
	return o.MinCacheEntryAttribute
}

// GetMinCacheEntryAttributeOk returns a tuple with the MinCacheEntryAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FifoEntryCacheShared) GetMinCacheEntryAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.MinCacheEntryAttribute) {
    return nil, false
	}
	return o.MinCacheEntryAttribute, true
}

// HasMinCacheEntryAttribute returns a boolean if a field has been set.
func (o *FifoEntryCacheShared) HasMinCacheEntryAttribute() bool {
	if o != nil && !isNil(o.MinCacheEntryAttribute) {
		return true
	}

	return false
}

// SetMinCacheEntryAttribute gets a reference to the given []string and assigns it to the MinCacheEntryAttribute field.
func (o *FifoEntryCacheShared) SetMinCacheEntryAttribute(v []string) {
	o.MinCacheEntryAttribute = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FifoEntryCacheShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FifoEntryCacheShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FifoEntryCacheShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FifoEntryCacheShared) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *FifoEntryCacheShared) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *FifoEntryCacheShared) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *FifoEntryCacheShared) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCacheLevel returns the CacheLevel field value
func (o *FifoEntryCacheShared) GetCacheLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CacheLevel
}

// GetCacheLevelOk returns a tuple with the CacheLevel field value
// and a boolean to check if the value has been set.
func (o *FifoEntryCacheShared) GetCacheLevelOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CacheLevel, true
}

// SetCacheLevel sets field value
func (o *FifoEntryCacheShared) SetCacheLevel(v int32) {
	o.CacheLevel = v
}

// GetCacheUnindexedSearchResults returns the CacheUnindexedSearchResults field value if set, zero value otherwise.
func (o *FifoEntryCacheShared) GetCacheUnindexedSearchResults() bool {
	if o == nil || isNil(o.CacheUnindexedSearchResults) {
		var ret bool
		return ret
	}
	return *o.CacheUnindexedSearchResults
}

// GetCacheUnindexedSearchResultsOk returns a tuple with the CacheUnindexedSearchResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FifoEntryCacheShared) GetCacheUnindexedSearchResultsOk() (*bool, bool) {
	if o == nil || isNil(o.CacheUnindexedSearchResults) {
    return nil, false
	}
	return o.CacheUnindexedSearchResults, true
}

// HasCacheUnindexedSearchResults returns a boolean if a field has been set.
func (o *FifoEntryCacheShared) HasCacheUnindexedSearchResults() bool {
	if o != nil && !isNil(o.CacheUnindexedSearchResults) {
		return true
	}

	return false
}

// SetCacheUnindexedSearchResults gets a reference to the given bool and assigns it to the CacheUnindexedSearchResults field.
func (o *FifoEntryCacheShared) SetCacheUnindexedSearchResults(v bool) {
	o.CacheUnindexedSearchResults = &v
}

func (o FifoEntryCacheShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.MaxMemoryPercent) {
		toSerialize["maxMemoryPercent"] = o.MaxMemoryPercent
	}
	if !isNil(o.MaxEntries) {
		toSerialize["maxEntries"] = o.MaxEntries
	}
	if !isNil(o.OnlyCacheFrequentlyAccessed) {
		toSerialize["onlyCacheFrequentlyAccessed"] = o.OnlyCacheFrequentlyAccessed
	}
	if !isNil(o.IncludeFilter) {
		toSerialize["includeFilter"] = o.IncludeFilter
	}
	if !isNil(o.ExcludeFilter) {
		toSerialize["excludeFilter"] = o.ExcludeFilter
	}
	if !isNil(o.MinCacheEntryValueCount) {
		toSerialize["minCacheEntryValueCount"] = o.MinCacheEntryValueCount
	}
	if !isNil(o.MinCacheEntryAttribute) {
		toSerialize["minCacheEntryAttribute"] = o.MinCacheEntryAttribute
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["cacheLevel"] = o.CacheLevel
	}
	if !isNil(o.CacheUnindexedSearchResults) {
		toSerialize["cacheUnindexedSearchResults"] = o.CacheUnindexedSearchResults
	}
	return json.Marshal(toSerialize)
}

type NullableFifoEntryCacheShared struct {
	value *FifoEntryCacheShared
	isSet bool
}

func (v NullableFifoEntryCacheShared) Get() *FifoEntryCacheShared {
	return v.value
}

func (v *NullableFifoEntryCacheShared) Set(val *FifoEntryCacheShared) {
	v.value = val
	v.isSet = true
}

func (v NullableFifoEntryCacheShared) IsSet() bool {
	return v.isSet
}

func (v *NullableFifoEntryCacheShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFifoEntryCacheShared(val *FifoEntryCacheShared) *NullableFifoEntryCacheShared {
	return &NullableFifoEntryCacheShared{value: val, isSet: true}
}

func (v NullableFifoEntryCacheShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFifoEntryCacheShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


