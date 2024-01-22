# FifoEntryCacheResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Entry Cache | 

## Methods

### NewFifoEntryCacheResponse

`func NewFifoEntryCacheResponse(schemas []EnumfifoEntryCacheSchemaUrn, enabled bool, cacheLevel int64, id string, ) *FifoEntryCacheResponse`

NewFifoEntryCacheResponse instantiates a new FifoEntryCacheResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFifoEntryCacheResponseWithDefaults

`func NewFifoEntryCacheResponseWithDefaults() *FifoEntryCacheResponse`

NewFifoEntryCacheResponseWithDefaults instantiates a new FifoEntryCacheResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *FifoEntryCacheResponse) GetSchemas() []EnumfifoEntryCacheSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FifoEntryCacheResponse) GetSchemasOk() (*[]EnumfifoEntryCacheSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FifoEntryCacheResponse) SetSchemas(v []EnumfifoEntryCacheSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetMaxMemoryPercent

`func (o *FifoEntryCacheResponse) GetMaxMemoryPercent() int64`

GetMaxMemoryPercent returns the MaxMemoryPercent field if non-nil, zero value otherwise.

### GetMaxMemoryPercentOk

`func (o *FifoEntryCacheResponse) GetMaxMemoryPercentOk() (*int64, bool)`

GetMaxMemoryPercentOk returns a tuple with the MaxMemoryPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryPercent

`func (o *FifoEntryCacheResponse) SetMaxMemoryPercent(v int64)`

SetMaxMemoryPercent sets MaxMemoryPercent field to given value.

### HasMaxMemoryPercent

`func (o *FifoEntryCacheResponse) HasMaxMemoryPercent() bool`

HasMaxMemoryPercent returns a boolean if a field has been set.

### GetMaxEntries

`func (o *FifoEntryCacheResponse) GetMaxEntries() int64`

GetMaxEntries returns the MaxEntries field if non-nil, zero value otherwise.

### GetMaxEntriesOk

`func (o *FifoEntryCacheResponse) GetMaxEntriesOk() (*int64, bool)`

GetMaxEntriesOk returns a tuple with the MaxEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEntries

`func (o *FifoEntryCacheResponse) SetMaxEntries(v int64)`

SetMaxEntries sets MaxEntries field to given value.

### HasMaxEntries

`func (o *FifoEntryCacheResponse) HasMaxEntries() bool`

HasMaxEntries returns a boolean if a field has been set.

### GetOnlyCacheFrequentlyAccessed

`func (o *FifoEntryCacheResponse) GetOnlyCacheFrequentlyAccessed() bool`

GetOnlyCacheFrequentlyAccessed returns the OnlyCacheFrequentlyAccessed field if non-nil, zero value otherwise.

### GetOnlyCacheFrequentlyAccessedOk

`func (o *FifoEntryCacheResponse) GetOnlyCacheFrequentlyAccessedOk() (*bool, bool)`

GetOnlyCacheFrequentlyAccessedOk returns a tuple with the OnlyCacheFrequentlyAccessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyCacheFrequentlyAccessed

`func (o *FifoEntryCacheResponse) SetOnlyCacheFrequentlyAccessed(v bool)`

SetOnlyCacheFrequentlyAccessed sets OnlyCacheFrequentlyAccessed field to given value.

### HasOnlyCacheFrequentlyAccessed

`func (o *FifoEntryCacheResponse) HasOnlyCacheFrequentlyAccessed() bool`

HasOnlyCacheFrequentlyAccessed returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *FifoEntryCacheResponse) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *FifoEntryCacheResponse) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *FifoEntryCacheResponse) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *FifoEntryCacheResponse) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetExcludeFilter

`func (o *FifoEntryCacheResponse) GetExcludeFilter() []string`

GetExcludeFilter returns the ExcludeFilter field if non-nil, zero value otherwise.

### GetExcludeFilterOk

`func (o *FifoEntryCacheResponse) GetExcludeFilterOk() (*[]string, bool)`

GetExcludeFilterOk returns a tuple with the ExcludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilter

`func (o *FifoEntryCacheResponse) SetExcludeFilter(v []string)`

SetExcludeFilter sets ExcludeFilter field to given value.

### HasExcludeFilter

`func (o *FifoEntryCacheResponse) HasExcludeFilter() bool`

HasExcludeFilter returns a boolean if a field has been set.

### GetMinCacheEntryValueCount

`func (o *FifoEntryCacheResponse) GetMinCacheEntryValueCount() int64`

GetMinCacheEntryValueCount returns the MinCacheEntryValueCount field if non-nil, zero value otherwise.

### GetMinCacheEntryValueCountOk

`func (o *FifoEntryCacheResponse) GetMinCacheEntryValueCountOk() (*int64, bool)`

GetMinCacheEntryValueCountOk returns a tuple with the MinCacheEntryValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCacheEntryValueCount

`func (o *FifoEntryCacheResponse) SetMinCacheEntryValueCount(v int64)`

SetMinCacheEntryValueCount sets MinCacheEntryValueCount field to given value.

### HasMinCacheEntryValueCount

`func (o *FifoEntryCacheResponse) HasMinCacheEntryValueCount() bool`

HasMinCacheEntryValueCount returns a boolean if a field has been set.

### GetMinCacheEntryAttribute

`func (o *FifoEntryCacheResponse) GetMinCacheEntryAttribute() []string`

GetMinCacheEntryAttribute returns the MinCacheEntryAttribute field if non-nil, zero value otherwise.

### GetMinCacheEntryAttributeOk

`func (o *FifoEntryCacheResponse) GetMinCacheEntryAttributeOk() (*[]string, bool)`

GetMinCacheEntryAttributeOk returns a tuple with the MinCacheEntryAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCacheEntryAttribute

`func (o *FifoEntryCacheResponse) SetMinCacheEntryAttribute(v []string)`

SetMinCacheEntryAttribute sets MinCacheEntryAttribute field to given value.

### HasMinCacheEntryAttribute

`func (o *FifoEntryCacheResponse) HasMinCacheEntryAttribute() bool`

HasMinCacheEntryAttribute returns a boolean if a field has been set.

### GetDescription

`func (o *FifoEntryCacheResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FifoEntryCacheResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FifoEntryCacheResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FifoEntryCacheResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *FifoEntryCacheResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FifoEntryCacheResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FifoEntryCacheResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCacheLevel

`func (o *FifoEntryCacheResponse) GetCacheLevel() int64`

GetCacheLevel returns the CacheLevel field if non-nil, zero value otherwise.

### GetCacheLevelOk

`func (o *FifoEntryCacheResponse) GetCacheLevelOk() (*int64, bool)`

GetCacheLevelOk returns a tuple with the CacheLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLevel

`func (o *FifoEntryCacheResponse) SetCacheLevel(v int64)`

SetCacheLevel sets CacheLevel field to given value.


### GetCacheUnindexedSearchResults

`func (o *FifoEntryCacheResponse) GetCacheUnindexedSearchResults() bool`

GetCacheUnindexedSearchResults returns the CacheUnindexedSearchResults field if non-nil, zero value otherwise.

### GetCacheUnindexedSearchResultsOk

`func (o *FifoEntryCacheResponse) GetCacheUnindexedSearchResultsOk() (*bool, bool)`

GetCacheUnindexedSearchResultsOk returns a tuple with the CacheUnindexedSearchResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheUnindexedSearchResults

`func (o *FifoEntryCacheResponse) SetCacheUnindexedSearchResults(v bool)`

SetCacheUnindexedSearchResults sets CacheUnindexedSearchResults field to given value.

### HasCacheUnindexedSearchResults

`func (o *FifoEntryCacheResponse) HasCacheUnindexedSearchResults() bool`

HasCacheUnindexedSearchResults returns a boolean if a field has been set.

### GetMeta

`func (o *FifoEntryCacheResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FifoEntryCacheResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FifoEntryCacheResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FifoEntryCacheResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *FifoEntryCacheResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *FifoEntryCacheResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *FifoEntryCacheResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *FifoEntryCacheResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *FifoEntryCacheResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FifoEntryCacheResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FifoEntryCacheResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


