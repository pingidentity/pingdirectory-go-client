# FifoEntryCacheShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumfifoEntryCacheSchemaUrn**](EnumfifoEntryCacheSchemaUrn.md) |  | [optional] 
**MaxMemoryPercent** | Pointer to **int32** | Specifies the maximum amount of memory, as a percentage of the total maximum JVM heap size, that this cache should occupy when full. If the amount of memory the cache is using is greater than this amount, then an attempt to put a new entry in the cache will be ignored and will cause the oldest entry to be purged. | [optional] 
**MaxEntries** | Pointer to **int32** | Specifies the maximum number of entries that will be allowed in the cache. Once the cache reaches this size, then adding new entries will cause existing entries to be purged, starting with the oldest. | [optional] 
**OnlyCacheFrequentlyAccessed** | Pointer to **bool** | Specifies that the cache should only store entries which are accessed much more frequently than the average entry. The cache will observe attempts to place entries in the cache and compare an entry&#39;s accesses to the average entry&#39;s. | [optional] 
**IncludeFilter** | Pointer to **[]string** |  | [optional] 
**ExcludeFilter** | Pointer to **[]string** |  | [optional] 
**MinCacheEntryValueCount** | Pointer to **int32** | Specifies the minimum number of attribute values (optionally across a specified subset of attributes as defined in the min-cache-entry-attributes property) for entries that should be held in the cache. Entries with fewer than this number of attribute values will be excluded from the cache. | [optional] 
**MinCacheEntryAttribute** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Entry Cache | [optional] 
**Enabled** | **bool** | Indicates whether the Entry Cache is enabled. | 
**CacheLevel** | **int32** | Specifies the cache level in the cache order if more than one instance of the cache is configured. | 
**CacheUnindexedSearchResults** | Pointer to **bool** | Indicates whether the entry cache should be updated with entries that have been returned to the client during the course of processing an unindexed search. | [optional] 

## Methods

### NewFifoEntryCacheShared

`func NewFifoEntryCacheShared(enabled bool, cacheLevel int32, ) *FifoEntryCacheShared`

NewFifoEntryCacheShared instantiates a new FifoEntryCacheShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFifoEntryCacheSharedWithDefaults

`func NewFifoEntryCacheSharedWithDefaults() *FifoEntryCacheShared`

NewFifoEntryCacheSharedWithDefaults instantiates a new FifoEntryCacheShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *FifoEntryCacheShared) GetSchemas() []EnumfifoEntryCacheSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FifoEntryCacheShared) GetSchemasOk() (*[]EnumfifoEntryCacheSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FifoEntryCacheShared) SetSchemas(v []EnumfifoEntryCacheSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *FifoEntryCacheShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetMaxMemoryPercent

`func (o *FifoEntryCacheShared) GetMaxMemoryPercent() int32`

GetMaxMemoryPercent returns the MaxMemoryPercent field if non-nil, zero value otherwise.

### GetMaxMemoryPercentOk

`func (o *FifoEntryCacheShared) GetMaxMemoryPercentOk() (*int32, bool)`

GetMaxMemoryPercentOk returns a tuple with the MaxMemoryPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryPercent

`func (o *FifoEntryCacheShared) SetMaxMemoryPercent(v int32)`

SetMaxMemoryPercent sets MaxMemoryPercent field to given value.

### HasMaxMemoryPercent

`func (o *FifoEntryCacheShared) HasMaxMemoryPercent() bool`

HasMaxMemoryPercent returns a boolean if a field has been set.

### GetMaxEntries

`func (o *FifoEntryCacheShared) GetMaxEntries() int32`

GetMaxEntries returns the MaxEntries field if non-nil, zero value otherwise.

### GetMaxEntriesOk

`func (o *FifoEntryCacheShared) GetMaxEntriesOk() (*int32, bool)`

GetMaxEntriesOk returns a tuple with the MaxEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEntries

`func (o *FifoEntryCacheShared) SetMaxEntries(v int32)`

SetMaxEntries sets MaxEntries field to given value.

### HasMaxEntries

`func (o *FifoEntryCacheShared) HasMaxEntries() bool`

HasMaxEntries returns a boolean if a field has been set.

### GetOnlyCacheFrequentlyAccessed

`func (o *FifoEntryCacheShared) GetOnlyCacheFrequentlyAccessed() bool`

GetOnlyCacheFrequentlyAccessed returns the OnlyCacheFrequentlyAccessed field if non-nil, zero value otherwise.

### GetOnlyCacheFrequentlyAccessedOk

`func (o *FifoEntryCacheShared) GetOnlyCacheFrequentlyAccessedOk() (*bool, bool)`

GetOnlyCacheFrequentlyAccessedOk returns a tuple with the OnlyCacheFrequentlyAccessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyCacheFrequentlyAccessed

`func (o *FifoEntryCacheShared) SetOnlyCacheFrequentlyAccessed(v bool)`

SetOnlyCacheFrequentlyAccessed sets OnlyCacheFrequentlyAccessed field to given value.

### HasOnlyCacheFrequentlyAccessed

`func (o *FifoEntryCacheShared) HasOnlyCacheFrequentlyAccessed() bool`

HasOnlyCacheFrequentlyAccessed returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *FifoEntryCacheShared) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *FifoEntryCacheShared) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *FifoEntryCacheShared) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *FifoEntryCacheShared) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetExcludeFilter

`func (o *FifoEntryCacheShared) GetExcludeFilter() []string`

GetExcludeFilter returns the ExcludeFilter field if non-nil, zero value otherwise.

### GetExcludeFilterOk

`func (o *FifoEntryCacheShared) GetExcludeFilterOk() (*[]string, bool)`

GetExcludeFilterOk returns a tuple with the ExcludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilter

`func (o *FifoEntryCacheShared) SetExcludeFilter(v []string)`

SetExcludeFilter sets ExcludeFilter field to given value.

### HasExcludeFilter

`func (o *FifoEntryCacheShared) HasExcludeFilter() bool`

HasExcludeFilter returns a boolean if a field has been set.

### GetMinCacheEntryValueCount

`func (o *FifoEntryCacheShared) GetMinCacheEntryValueCount() int32`

GetMinCacheEntryValueCount returns the MinCacheEntryValueCount field if non-nil, zero value otherwise.

### GetMinCacheEntryValueCountOk

`func (o *FifoEntryCacheShared) GetMinCacheEntryValueCountOk() (*int32, bool)`

GetMinCacheEntryValueCountOk returns a tuple with the MinCacheEntryValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCacheEntryValueCount

`func (o *FifoEntryCacheShared) SetMinCacheEntryValueCount(v int32)`

SetMinCacheEntryValueCount sets MinCacheEntryValueCount field to given value.

### HasMinCacheEntryValueCount

`func (o *FifoEntryCacheShared) HasMinCacheEntryValueCount() bool`

HasMinCacheEntryValueCount returns a boolean if a field has been set.

### GetMinCacheEntryAttribute

`func (o *FifoEntryCacheShared) GetMinCacheEntryAttribute() []string`

GetMinCacheEntryAttribute returns the MinCacheEntryAttribute field if non-nil, zero value otherwise.

### GetMinCacheEntryAttributeOk

`func (o *FifoEntryCacheShared) GetMinCacheEntryAttributeOk() (*[]string, bool)`

GetMinCacheEntryAttributeOk returns a tuple with the MinCacheEntryAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCacheEntryAttribute

`func (o *FifoEntryCacheShared) SetMinCacheEntryAttribute(v []string)`

SetMinCacheEntryAttribute sets MinCacheEntryAttribute field to given value.

### HasMinCacheEntryAttribute

`func (o *FifoEntryCacheShared) HasMinCacheEntryAttribute() bool`

HasMinCacheEntryAttribute returns a boolean if a field has been set.

### GetDescription

`func (o *FifoEntryCacheShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FifoEntryCacheShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FifoEntryCacheShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FifoEntryCacheShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *FifoEntryCacheShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FifoEntryCacheShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FifoEntryCacheShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCacheLevel

`func (o *FifoEntryCacheShared) GetCacheLevel() int32`

GetCacheLevel returns the CacheLevel field if non-nil, zero value otherwise.

### GetCacheLevelOk

`func (o *FifoEntryCacheShared) GetCacheLevelOk() (*int32, bool)`

GetCacheLevelOk returns a tuple with the CacheLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLevel

`func (o *FifoEntryCacheShared) SetCacheLevel(v int32)`

SetCacheLevel sets CacheLevel field to given value.


### GetCacheUnindexedSearchResults

`func (o *FifoEntryCacheShared) GetCacheUnindexedSearchResults() bool`

GetCacheUnindexedSearchResults returns the CacheUnindexedSearchResults field if non-nil, zero value otherwise.

### GetCacheUnindexedSearchResultsOk

`func (o *FifoEntryCacheShared) GetCacheUnindexedSearchResultsOk() (*bool, bool)`

GetCacheUnindexedSearchResultsOk returns a tuple with the CacheUnindexedSearchResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheUnindexedSearchResults

`func (o *FifoEntryCacheShared) SetCacheUnindexedSearchResults(v bool)`

SetCacheUnindexedSearchResults sets CacheUnindexedSearchResults field to given value.

### HasCacheUnindexedSearchResults

`func (o *FifoEntryCacheShared) HasCacheUnindexedSearchResults() bool`

HasCacheUnindexedSearchResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


