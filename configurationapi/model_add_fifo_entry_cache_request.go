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

// checks if the AddFifoEntryCacheRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddFifoEntryCacheRequest{}

// AddFifoEntryCacheRequest struct for AddFifoEntryCacheRequest
type AddFifoEntryCacheRequest struct {
	Schemas []EnumfifoEntryCacheSchemaUrn `json:"schemas"`
	// Specifies the maximum amount of memory, as a percentage of the total maximum JVM heap size, that this cache should occupy when full. If the amount of memory the cache is using is greater than this amount, then an attempt to put a new entry in the cache will be ignored and will cause the oldest entry to be purged.
	MaxMemoryPercent *int64 `json:"maxMemoryPercent,omitempty"`
	// Specifies the maximum number of entries that will be allowed in the cache. Once the cache reaches this size, then adding new entries will cause existing entries to be purged, starting with the oldest.
	MaxEntries *int64 `json:"maxEntries,omitempty"`
	// Specifies that the cache should only store entries which are accessed much more frequently than the average entry. The cache will observe attempts to place entries in the cache and compare an entry's accesses to the average entry's.
	OnlyCacheFrequentlyAccessed *bool `json:"onlyCacheFrequentlyAccessed,omitempty"`
	// The set of filters that define the entries that should be included in the cache.
	IncludeFilter []string `json:"includeFilter,omitempty"`
	// The set of filters that define the entries that should be excluded from the cache.
	ExcludeFilter []string `json:"excludeFilter,omitempty"`
	// Specifies the minimum number of attribute values (optionally across a specified subset of attributes as defined in the min-cache-entry-attributes property) for entries that should be held in the cache. Entries with fewer than this number of attribute values will be excluded from the cache.
	MinCacheEntryValueCount *int64 `json:"minCacheEntryValueCount,omitempty"`
	// Specifies the names of the attribute types for which the min-cache-entry-value-count property should apply. If no attribute types are specified, then all user attributes will be examined.
	MinCacheEntryAttribute []string `json:"minCacheEntryAttribute,omitempty"`
	// A description for this Entry Cache
	Description *string `json:"description,omitempty"`
	// Indicates whether the Entry Cache is enabled.
	Enabled bool `json:"enabled"`
	// Specifies the cache level in the cache order if more than one instance of the cache is configured.
	CacheLevel int64 `json:"cacheLevel"`
	// Indicates whether the entry cache should be updated with entries that have been returned to the client during the course of processing an unindexed search.
	CacheUnindexedSearchResults *bool `json:"cacheUnindexedSearchResults,omitempty"`
	// Name of the new Entry Cache
	CacheName string `json:"cacheName"`
}

// NewAddFifoEntryCacheRequest instantiates a new AddFifoEntryCacheRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddFifoEntryCacheRequest(schemas []EnumfifoEntryCacheSchemaUrn, enabled bool, cacheLevel int64, cacheName string) *AddFifoEntryCacheRequest {
	this := AddFifoEntryCacheRequest{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.CacheLevel = cacheLevel
	this.CacheName = cacheName
	return &this
}

// NewAddFifoEntryCacheRequestWithDefaults instantiates a new AddFifoEntryCacheRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddFifoEntryCacheRequestWithDefaults() *AddFifoEntryCacheRequest {
	this := AddFifoEntryCacheRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddFifoEntryCacheRequest) GetSchemas() []EnumfifoEntryCacheSchemaUrn {
	if o == nil {
		var ret []EnumfifoEntryCacheSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddFifoEntryCacheRequest) GetSchemasOk() ([]EnumfifoEntryCacheSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddFifoEntryCacheRequest) SetSchemas(v []EnumfifoEntryCacheSchemaUrn) {
	o.Schemas = v
}

// GetMaxMemoryPercent returns the MaxMemoryPercent field value if set, zero value otherwise.
func (o *AddFifoEntryCacheRequest) GetMaxMemoryPercent() int64 {
	if o == nil || IsNil(o.MaxMemoryPercent) {
		var ret int64
		return ret
	}
	return *o.MaxMemoryPercent
}

// GetMaxMemoryPercentOk returns a tuple with the MaxMemoryPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFifoEntryCacheRequest) GetMaxMemoryPercentOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxMemoryPercent) {
		return nil, false
	}
	return o.MaxMemoryPercent, true
}

// HasMaxMemoryPercent returns a boolean if a field has been set.
func (o *AddFifoEntryCacheRequest) HasMaxMemoryPercent() bool {
	if o != nil && !IsNil(o.MaxMemoryPercent) {
		return true
	}

	return false
}

// SetMaxMemoryPercent gets a reference to the given int64 and assigns it to the MaxMemoryPercent field.
func (o *AddFifoEntryCacheRequest) SetMaxMemoryPercent(v int64) {
	o.MaxMemoryPercent = &v
}

// GetMaxEntries returns the MaxEntries field value if set, zero value otherwise.
func (o *AddFifoEntryCacheRequest) GetMaxEntries() int64 {
	if o == nil || IsNil(o.MaxEntries) {
		var ret int64
		return ret
	}
	return *o.MaxEntries
}

// GetMaxEntriesOk returns a tuple with the MaxEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFifoEntryCacheRequest) GetMaxEntriesOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxEntries) {
		return nil, false
	}
	return o.MaxEntries, true
}

// HasMaxEntries returns a boolean if a field has been set.
func (o *AddFifoEntryCacheRequest) HasMaxEntries() bool {
	if o != nil && !IsNil(o.MaxEntries) {
		return true
	}

	return false
}

// SetMaxEntries gets a reference to the given int64 and assigns it to the MaxEntries field.
func (o *AddFifoEntryCacheRequest) SetMaxEntries(v int64) {
	o.MaxEntries = &v
}

// GetOnlyCacheFrequentlyAccessed returns the OnlyCacheFrequentlyAccessed field value if set, zero value otherwise.
func (o *AddFifoEntryCacheRequest) GetOnlyCacheFrequentlyAccessed() bool {
	if o == nil || IsNil(o.OnlyCacheFrequentlyAccessed) {
		var ret bool
		return ret
	}
	return *o.OnlyCacheFrequentlyAccessed
}

// GetOnlyCacheFrequentlyAccessedOk returns a tuple with the OnlyCacheFrequentlyAccessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFifoEntryCacheRequest) GetOnlyCacheFrequentlyAccessedOk() (*bool, bool) {
	if o == nil || IsNil(o.OnlyCacheFrequentlyAccessed) {
		return nil, false
	}
	return o.OnlyCacheFrequentlyAccessed, true
}

// HasOnlyCacheFrequentlyAccessed returns a boolean if a field has been set.
func (o *AddFifoEntryCacheRequest) HasOnlyCacheFrequentlyAccessed() bool {
	if o != nil && !IsNil(o.OnlyCacheFrequentlyAccessed) {
		return true
	}

	return false
}

// SetOnlyCacheFrequentlyAccessed gets a reference to the given bool and assigns it to the OnlyCacheFrequentlyAccessed field.
func (o *AddFifoEntryCacheRequest) SetOnlyCacheFrequentlyAccessed(v bool) {
	o.OnlyCacheFrequentlyAccessed = &v
}

// GetIncludeFilter returns the IncludeFilter field value if set, zero value otherwise.
func (o *AddFifoEntryCacheRequest) GetIncludeFilter() []string {
	if o == nil || IsNil(o.IncludeFilter) {
		var ret []string
		return ret
	}
	return o.IncludeFilter
}

// GetIncludeFilterOk returns a tuple with the IncludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFifoEntryCacheRequest) GetIncludeFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeFilter) {
		return nil, false
	}
	return o.IncludeFilter, true
}

// HasIncludeFilter returns a boolean if a field has been set.
func (o *AddFifoEntryCacheRequest) HasIncludeFilter() bool {
	if o != nil && !IsNil(o.IncludeFilter) {
		return true
	}

	return false
}

// SetIncludeFilter gets a reference to the given []string and assigns it to the IncludeFilter field.
func (o *AddFifoEntryCacheRequest) SetIncludeFilter(v []string) {
	o.IncludeFilter = v
}

// GetExcludeFilter returns the ExcludeFilter field value if set, zero value otherwise.
func (o *AddFifoEntryCacheRequest) GetExcludeFilter() []string {
	if o == nil || IsNil(o.ExcludeFilter) {
		var ret []string
		return ret
	}
	return o.ExcludeFilter
}

// GetExcludeFilterOk returns a tuple with the ExcludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFifoEntryCacheRequest) GetExcludeFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeFilter) {
		return nil, false
	}
	return o.ExcludeFilter, true
}

// HasExcludeFilter returns a boolean if a field has been set.
func (o *AddFifoEntryCacheRequest) HasExcludeFilter() bool {
	if o != nil && !IsNil(o.ExcludeFilter) {
		return true
	}

	return false
}

// SetExcludeFilter gets a reference to the given []string and assigns it to the ExcludeFilter field.
func (o *AddFifoEntryCacheRequest) SetExcludeFilter(v []string) {
	o.ExcludeFilter = v
}

// GetMinCacheEntryValueCount returns the MinCacheEntryValueCount field value if set, zero value otherwise.
func (o *AddFifoEntryCacheRequest) GetMinCacheEntryValueCount() int64 {
	if o == nil || IsNil(o.MinCacheEntryValueCount) {
		var ret int64
		return ret
	}
	return *o.MinCacheEntryValueCount
}

// GetMinCacheEntryValueCountOk returns a tuple with the MinCacheEntryValueCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFifoEntryCacheRequest) GetMinCacheEntryValueCountOk() (*int64, bool) {
	if o == nil || IsNil(o.MinCacheEntryValueCount) {
		return nil, false
	}
	return o.MinCacheEntryValueCount, true
}

// HasMinCacheEntryValueCount returns a boolean if a field has been set.
func (o *AddFifoEntryCacheRequest) HasMinCacheEntryValueCount() bool {
	if o != nil && !IsNil(o.MinCacheEntryValueCount) {
		return true
	}

	return false
}

// SetMinCacheEntryValueCount gets a reference to the given int64 and assigns it to the MinCacheEntryValueCount field.
func (o *AddFifoEntryCacheRequest) SetMinCacheEntryValueCount(v int64) {
	o.MinCacheEntryValueCount = &v
}

// GetMinCacheEntryAttribute returns the MinCacheEntryAttribute field value if set, zero value otherwise.
func (o *AddFifoEntryCacheRequest) GetMinCacheEntryAttribute() []string {
	if o == nil || IsNil(o.MinCacheEntryAttribute) {
		var ret []string
		return ret
	}
	return o.MinCacheEntryAttribute
}

// GetMinCacheEntryAttributeOk returns a tuple with the MinCacheEntryAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFifoEntryCacheRequest) GetMinCacheEntryAttributeOk() ([]string, bool) {
	if o == nil || IsNil(o.MinCacheEntryAttribute) {
		return nil, false
	}
	return o.MinCacheEntryAttribute, true
}

// HasMinCacheEntryAttribute returns a boolean if a field has been set.
func (o *AddFifoEntryCacheRequest) HasMinCacheEntryAttribute() bool {
	if o != nil && !IsNil(o.MinCacheEntryAttribute) {
		return true
	}

	return false
}

// SetMinCacheEntryAttribute gets a reference to the given []string and assigns it to the MinCacheEntryAttribute field.
func (o *AddFifoEntryCacheRequest) SetMinCacheEntryAttribute(v []string) {
	o.MinCacheEntryAttribute = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddFifoEntryCacheRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFifoEntryCacheRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddFifoEntryCacheRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddFifoEntryCacheRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddFifoEntryCacheRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddFifoEntryCacheRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddFifoEntryCacheRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCacheLevel returns the CacheLevel field value
func (o *AddFifoEntryCacheRequest) GetCacheLevel() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CacheLevel
}

// GetCacheLevelOk returns a tuple with the CacheLevel field value
// and a boolean to check if the value has been set.
func (o *AddFifoEntryCacheRequest) GetCacheLevelOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheLevel, true
}

// SetCacheLevel sets field value
func (o *AddFifoEntryCacheRequest) SetCacheLevel(v int64) {
	o.CacheLevel = v
}

// GetCacheUnindexedSearchResults returns the CacheUnindexedSearchResults field value if set, zero value otherwise.
func (o *AddFifoEntryCacheRequest) GetCacheUnindexedSearchResults() bool {
	if o == nil || IsNil(o.CacheUnindexedSearchResults) {
		var ret bool
		return ret
	}
	return *o.CacheUnindexedSearchResults
}

// GetCacheUnindexedSearchResultsOk returns a tuple with the CacheUnindexedSearchResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFifoEntryCacheRequest) GetCacheUnindexedSearchResultsOk() (*bool, bool) {
	if o == nil || IsNil(o.CacheUnindexedSearchResults) {
		return nil, false
	}
	return o.CacheUnindexedSearchResults, true
}

// HasCacheUnindexedSearchResults returns a boolean if a field has been set.
func (o *AddFifoEntryCacheRequest) HasCacheUnindexedSearchResults() bool {
	if o != nil && !IsNil(o.CacheUnindexedSearchResults) {
		return true
	}

	return false
}

// SetCacheUnindexedSearchResults gets a reference to the given bool and assigns it to the CacheUnindexedSearchResults field.
func (o *AddFifoEntryCacheRequest) SetCacheUnindexedSearchResults(v bool) {
	o.CacheUnindexedSearchResults = &v
}

// GetCacheName returns the CacheName field value
func (o *AddFifoEntryCacheRequest) GetCacheName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CacheName
}

// GetCacheNameOk returns a tuple with the CacheName field value
// and a boolean to check if the value has been set.
func (o *AddFifoEntryCacheRequest) GetCacheNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheName, true
}

// SetCacheName sets field value
func (o *AddFifoEntryCacheRequest) SetCacheName(v string) {
	o.CacheName = v
}

func (o AddFifoEntryCacheRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddFifoEntryCacheRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.MaxMemoryPercent) {
		toSerialize["maxMemoryPercent"] = o.MaxMemoryPercent
	}
	if !IsNil(o.MaxEntries) {
		toSerialize["maxEntries"] = o.MaxEntries
	}
	if !IsNil(o.OnlyCacheFrequentlyAccessed) {
		toSerialize["onlyCacheFrequentlyAccessed"] = o.OnlyCacheFrequentlyAccessed
	}
	if !IsNil(o.IncludeFilter) {
		toSerialize["includeFilter"] = o.IncludeFilter
	}
	if !IsNil(o.ExcludeFilter) {
		toSerialize["excludeFilter"] = o.ExcludeFilter
	}
	if !IsNil(o.MinCacheEntryValueCount) {
		toSerialize["minCacheEntryValueCount"] = o.MinCacheEntryValueCount
	}
	if !IsNil(o.MinCacheEntryAttribute) {
		toSerialize["minCacheEntryAttribute"] = o.MinCacheEntryAttribute
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["cacheLevel"] = o.CacheLevel
	if !IsNil(o.CacheUnindexedSearchResults) {
		toSerialize["cacheUnindexedSearchResults"] = o.CacheUnindexedSearchResults
	}
	toSerialize["cacheName"] = o.CacheName
	return toSerialize, nil
}

type NullableAddFifoEntryCacheRequest struct {
	value *AddFifoEntryCacheRequest
	isSet bool
}

func (v NullableAddFifoEntryCacheRequest) Get() *AddFifoEntryCacheRequest {
	return v.value
}

func (v *NullableAddFifoEntryCacheRequest) Set(val *AddFifoEntryCacheRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddFifoEntryCacheRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddFifoEntryCacheRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddFifoEntryCacheRequest(val *AddFifoEntryCacheRequest) *NullableAddFifoEntryCacheRequest {
	return &NullableAddFifoEntryCacheRequest{value: val, isSet: true}
}

func (v NullableAddFifoEntryCacheRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddFifoEntryCacheRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
