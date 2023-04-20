# AddFifoEntryCacheRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheName** | **string** | Name of the new Entry Cache | 
**Schemas** | [**[]EnumfifoEntryCacheSchemaUrn**](EnumfifoEntryCacheSchemaUrn.md) |  | 
**MaxMemoryPercent** | Pointer to **int64** | Specifies the maximum amount of memory, as a percentage of the total maximum JVM heap size, that this cache should occupy when full. If the amount of memory the cache is using is greater than this amount, then an attempt to put a new entry in the cache will be ignored and will cause the oldest entry to be purged. | [optional] 
**MaxEntries** | Pointer to **int64** | Specifies the maximum number of entries that will be allowed in the cache. Once the cache reaches this size, then adding new entries will cause existing entries to be purged, starting with the oldest. | [optional] 
**OnlyCacheFrequentlyAccessed** | Pointer to **bool** | Specifies that the cache should only store entries which are accessed much more frequently than the average entry. The cache will observe attempts to place entries in the cache and compare an entry&#39;s accesses to the average entry&#39;s. | [optional] 
**IncludeFilter** | Pointer to **[]string** | The set of filters that define the entries that should be included in the cache. | [optional] 
**ExcludeFilter** | Pointer to **[]string** | The set of filters that define the entries that should be excluded from the cache. | [optional] 
**MinCacheEntryValueCount** | Pointer to **int64** | Specifies the minimum number of attribute values (optionally across a specified subset of attributes as defined in the min-cache-entry-attributes property) for entries that should be held in the cache. Entries with fewer than this number of attribute values will be excluded from the cache. | [optional] 
**MinCacheEntryAttribute** | Pointer to **[]string** | Specifies the names of the attribute types for which the min-cache-entry-value-count property should apply. If no attribute types are specified, then all user attributes will be examined. | [optional] 
**Description** | Pointer to **string** | A description for this Entry Cache | [optional] 
**Enabled** | **bool** | Indicates whether the Entry Cache is enabled. | 
**CacheLevel** | **int64** | Specifies the cache level in the cache order if more than one instance of the cache is configured. | 
**CacheUnindexedSearchResults** | Pointer to **bool** | Indicates whether the entry cache should be updated with entries that have been returned to the client during the course of processing an unindexed search. | [optional] 

## Methods

### NewAddFifoEntryCacheRequest

`func NewAddFifoEntryCacheRequest(cacheName string, schemas []EnumfifoEntryCacheSchemaUrn, enabled bool, cacheLevel int64, ) *AddFifoEntryCacheRequest`

NewAddFifoEntryCacheRequest instantiates a new AddFifoEntryCacheRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFifoEntryCacheRequestWithDefaults

`func NewAddFifoEntryCacheRequestWithDefaults() *AddFifoEntryCacheRequest`

NewAddFifoEntryCacheRequestWithDefaults instantiates a new AddFifoEntryCacheRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheName

`func (o *AddFifoEntryCacheRequest) GetCacheName() string`

GetCacheName returns the CacheName field if non-nil, zero value otherwise.

### GetCacheNameOk

`func (o *AddFifoEntryCacheRequest) GetCacheNameOk() (*string, bool)`

GetCacheNameOk returns a tuple with the CacheName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheName

`func (o *AddFifoEntryCacheRequest) SetCacheName(v string)`

SetCacheName sets CacheName field to given value.


### GetSchemas

`func (o *AddFifoEntryCacheRequest) GetSchemas() []EnumfifoEntryCacheSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddFifoEntryCacheRequest) GetSchemasOk() (*[]EnumfifoEntryCacheSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddFifoEntryCacheRequest) SetSchemas(v []EnumfifoEntryCacheSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetMaxMemoryPercent

`func (o *AddFifoEntryCacheRequest) GetMaxMemoryPercent() int64`

GetMaxMemoryPercent returns the MaxMemoryPercent field if non-nil, zero value otherwise.

### GetMaxMemoryPercentOk

`func (o *AddFifoEntryCacheRequest) GetMaxMemoryPercentOk() (*int64, bool)`

GetMaxMemoryPercentOk returns a tuple with the MaxMemoryPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryPercent

`func (o *AddFifoEntryCacheRequest) SetMaxMemoryPercent(v int64)`

SetMaxMemoryPercent sets MaxMemoryPercent field to given value.

### HasMaxMemoryPercent

`func (o *AddFifoEntryCacheRequest) HasMaxMemoryPercent() bool`

HasMaxMemoryPercent returns a boolean if a field has been set.

### GetMaxEntries

`func (o *AddFifoEntryCacheRequest) GetMaxEntries() int64`

GetMaxEntries returns the MaxEntries field if non-nil, zero value otherwise.

### GetMaxEntriesOk

`func (o *AddFifoEntryCacheRequest) GetMaxEntriesOk() (*int64, bool)`

GetMaxEntriesOk returns a tuple with the MaxEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEntries

`func (o *AddFifoEntryCacheRequest) SetMaxEntries(v int64)`

SetMaxEntries sets MaxEntries field to given value.

### HasMaxEntries

`func (o *AddFifoEntryCacheRequest) HasMaxEntries() bool`

HasMaxEntries returns a boolean if a field has been set.

### GetOnlyCacheFrequentlyAccessed

`func (o *AddFifoEntryCacheRequest) GetOnlyCacheFrequentlyAccessed() bool`

GetOnlyCacheFrequentlyAccessed returns the OnlyCacheFrequentlyAccessed field if non-nil, zero value otherwise.

### GetOnlyCacheFrequentlyAccessedOk

`func (o *AddFifoEntryCacheRequest) GetOnlyCacheFrequentlyAccessedOk() (*bool, bool)`

GetOnlyCacheFrequentlyAccessedOk returns a tuple with the OnlyCacheFrequentlyAccessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyCacheFrequentlyAccessed

`func (o *AddFifoEntryCacheRequest) SetOnlyCacheFrequentlyAccessed(v bool)`

SetOnlyCacheFrequentlyAccessed sets OnlyCacheFrequentlyAccessed field to given value.

### HasOnlyCacheFrequentlyAccessed

`func (o *AddFifoEntryCacheRequest) HasOnlyCacheFrequentlyAccessed() bool`

HasOnlyCacheFrequentlyAccessed returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *AddFifoEntryCacheRequest) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *AddFifoEntryCacheRequest) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *AddFifoEntryCacheRequest) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *AddFifoEntryCacheRequest) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetExcludeFilter

`func (o *AddFifoEntryCacheRequest) GetExcludeFilter() []string`

GetExcludeFilter returns the ExcludeFilter field if non-nil, zero value otherwise.

### GetExcludeFilterOk

`func (o *AddFifoEntryCacheRequest) GetExcludeFilterOk() (*[]string, bool)`

GetExcludeFilterOk returns a tuple with the ExcludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilter

`func (o *AddFifoEntryCacheRequest) SetExcludeFilter(v []string)`

SetExcludeFilter sets ExcludeFilter field to given value.

### HasExcludeFilter

`func (o *AddFifoEntryCacheRequest) HasExcludeFilter() bool`

HasExcludeFilter returns a boolean if a field has been set.

### GetMinCacheEntryValueCount

`func (o *AddFifoEntryCacheRequest) GetMinCacheEntryValueCount() int64`

GetMinCacheEntryValueCount returns the MinCacheEntryValueCount field if non-nil, zero value otherwise.

### GetMinCacheEntryValueCountOk

`func (o *AddFifoEntryCacheRequest) GetMinCacheEntryValueCountOk() (*int64, bool)`

GetMinCacheEntryValueCountOk returns a tuple with the MinCacheEntryValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCacheEntryValueCount

`func (o *AddFifoEntryCacheRequest) SetMinCacheEntryValueCount(v int64)`

SetMinCacheEntryValueCount sets MinCacheEntryValueCount field to given value.

### HasMinCacheEntryValueCount

`func (o *AddFifoEntryCacheRequest) HasMinCacheEntryValueCount() bool`

HasMinCacheEntryValueCount returns a boolean if a field has been set.

### GetMinCacheEntryAttribute

`func (o *AddFifoEntryCacheRequest) GetMinCacheEntryAttribute() []string`

GetMinCacheEntryAttribute returns the MinCacheEntryAttribute field if non-nil, zero value otherwise.

### GetMinCacheEntryAttributeOk

`func (o *AddFifoEntryCacheRequest) GetMinCacheEntryAttributeOk() (*[]string, bool)`

GetMinCacheEntryAttributeOk returns a tuple with the MinCacheEntryAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCacheEntryAttribute

`func (o *AddFifoEntryCacheRequest) SetMinCacheEntryAttribute(v []string)`

SetMinCacheEntryAttribute sets MinCacheEntryAttribute field to given value.

### HasMinCacheEntryAttribute

`func (o *AddFifoEntryCacheRequest) HasMinCacheEntryAttribute() bool`

HasMinCacheEntryAttribute returns a boolean if a field has been set.

### GetDescription

`func (o *AddFifoEntryCacheRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddFifoEntryCacheRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddFifoEntryCacheRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddFifoEntryCacheRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddFifoEntryCacheRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddFifoEntryCacheRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddFifoEntryCacheRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCacheLevel

`func (o *AddFifoEntryCacheRequest) GetCacheLevel() int64`

GetCacheLevel returns the CacheLevel field if non-nil, zero value otherwise.

### GetCacheLevelOk

`func (o *AddFifoEntryCacheRequest) GetCacheLevelOk() (*int64, bool)`

GetCacheLevelOk returns a tuple with the CacheLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLevel

`func (o *AddFifoEntryCacheRequest) SetCacheLevel(v int64)`

SetCacheLevel sets CacheLevel field to given value.


### GetCacheUnindexedSearchResults

`func (o *AddFifoEntryCacheRequest) GetCacheUnindexedSearchResults() bool`

GetCacheUnindexedSearchResults returns the CacheUnindexedSearchResults field if non-nil, zero value otherwise.

### GetCacheUnindexedSearchResultsOk

`func (o *AddFifoEntryCacheRequest) GetCacheUnindexedSearchResultsOk() (*bool, bool)`

GetCacheUnindexedSearchResultsOk returns a tuple with the CacheUnindexedSearchResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheUnindexedSearchResults

`func (o *AddFifoEntryCacheRequest) SetCacheUnindexedSearchResults(v bool)`

SetCacheUnindexedSearchResults sets CacheUnindexedSearchResults field to given value.

### HasCacheUnindexedSearchResults

`func (o *AddFifoEntryCacheRequest) HasCacheUnindexedSearchResults() bool`

HasCacheUnindexedSearchResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


