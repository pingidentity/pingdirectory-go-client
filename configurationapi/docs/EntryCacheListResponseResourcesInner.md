# EntryCacheListResponseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumfifoEntryCacheSchemaUrn**](EnumfifoEntryCacheSchemaUrn.md) |  | 
**Id** | **string** | Name of the Entry Cache | 
**MaxMemorySize** | Pointer to **string** | The maximum size of the entry cache in bytes. | [optional] 
**MaxEntries** | Pointer to **int64** | Specifies the maximum number of entries that will be allowed in the cache. Once the cache reaches this size, then adding new entries will cause existing entries to be purged, starting with the oldest. | [optional] 
**CacheType** | Pointer to [**EnumentryCacheCacheTypeProp**](EnumentryCacheCacheTypeProp.md) |  | [optional] 
**CacheDirectory** | Pointer to **string** | Specifies the directory in which the cache database files should be stored. | [optional] 
**PersistentCache** | Pointer to **bool** | Specifies whether the cache should persist across restarts. | [optional] 
**CompactEncoding** | Pointer to **bool** | Indicates whether the cache should use a compact form when encoding cache entries by compressing the attribute descriptions and object class sets. | [optional] 
**DbCachePercent** | Pointer to **int64** | Specifies the maximum memory usage for the internal JE cache as a percentage of the total JVM memory. | [optional] 
**DbCacheSize** | Pointer to **string** | Specifies the maximum JVM memory usage in bytes for the internal JE cache. | [optional] 
**JeProperty** | Pointer to **[]string** | Specifies the environment properties for the Berkeley DB Java Edition database providing the backend for this entry cache. | [optional] 
**IncludeFilter** | Pointer to **[]string** | The set of filters that define the entries that should be included in the cache. | [optional] 
**ExcludeFilter** | Pointer to **[]string** | The set of filters that define the entries that should be excluded from the cache. | [optional] 
**Description** | Pointer to **string** | A description for this Entry Cache | [optional] 
**Enabled** | **bool** | Indicates whether the Entry Cache is enabled. | 
**CacheLevel** | **int64** | Specifies the cache level in the cache order if more than one instance of the cache is configured. | 
**CacheUnindexedSearchResults** | Pointer to **bool** | Indicates whether the entry cache should be updated with entries that have been returned to the client during the course of processing an unindexed search. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**MaxMemoryPercent** | Pointer to **int64** | Specifies the maximum amount of memory, as a percentage of the total maximum JVM heap size, that this cache should occupy when full. If the amount of memory the cache is using is greater than this amount, then an attempt to put a new entry in the cache will be ignored and will cause the oldest entry to be purged. | [optional] 
**OnlyCacheFrequentlyAccessed** | Pointer to **bool** | Specifies that the cache should only store entries which are accessed much more frequently than the average entry. The cache will observe attempts to place entries in the cache and compare an entry&#39;s accesses to the average entry&#39;s. | [optional] 
**MinCacheEntryValueCount** | Pointer to **int64** | Specifies the minimum number of attribute values (optionally across a specified subset of attributes as defined in the min-cache-entry-attributes property) for entries that should be held in the cache. Entries with fewer than this number of attribute values will be excluded from the cache. | [optional] 
**MinCacheEntryAttribute** | Pointer to **[]string** | Specifies the names of the attribute types for which the min-cache-entry-value-count property should apply. If no attribute types are specified, then all user attributes will be examined. | [optional] 

## Methods

### NewEntryCacheListResponseResourcesInner

`func NewEntryCacheListResponseResourcesInner(schemas []EnumfifoEntryCacheSchemaUrn, id string, enabled bool, cacheLevel int64, ) *EntryCacheListResponseResourcesInner`

NewEntryCacheListResponseResourcesInner instantiates a new EntryCacheListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryCacheListResponseResourcesInnerWithDefaults

`func NewEntryCacheListResponseResourcesInnerWithDefaults() *EntryCacheListResponseResourcesInner`

NewEntryCacheListResponseResourcesInnerWithDefaults instantiates a new EntryCacheListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *EntryCacheListResponseResourcesInner) GetSchemas() []EnumfifoEntryCacheSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *EntryCacheListResponseResourcesInner) GetSchemasOk() (*[]EnumfifoEntryCacheSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *EntryCacheListResponseResourcesInner) SetSchemas(v []EnumfifoEntryCacheSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *EntryCacheListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntryCacheListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntryCacheListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetMaxMemorySize

`func (o *EntryCacheListResponseResourcesInner) GetMaxMemorySize() string`

GetMaxMemorySize returns the MaxMemorySize field if non-nil, zero value otherwise.

### GetMaxMemorySizeOk

`func (o *EntryCacheListResponseResourcesInner) GetMaxMemorySizeOk() (*string, bool)`

GetMaxMemorySizeOk returns a tuple with the MaxMemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemorySize

`func (o *EntryCacheListResponseResourcesInner) SetMaxMemorySize(v string)`

SetMaxMemorySize sets MaxMemorySize field to given value.

### HasMaxMemorySize

`func (o *EntryCacheListResponseResourcesInner) HasMaxMemorySize() bool`

HasMaxMemorySize returns a boolean if a field has been set.

### GetMaxEntries

`func (o *EntryCacheListResponseResourcesInner) GetMaxEntries() int64`

GetMaxEntries returns the MaxEntries field if non-nil, zero value otherwise.

### GetMaxEntriesOk

`func (o *EntryCacheListResponseResourcesInner) GetMaxEntriesOk() (*int64, bool)`

GetMaxEntriesOk returns a tuple with the MaxEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEntries

`func (o *EntryCacheListResponseResourcesInner) SetMaxEntries(v int64)`

SetMaxEntries sets MaxEntries field to given value.

### HasMaxEntries

`func (o *EntryCacheListResponseResourcesInner) HasMaxEntries() bool`

HasMaxEntries returns a boolean if a field has been set.

### GetCacheType

`func (o *EntryCacheListResponseResourcesInner) GetCacheType() EnumentryCacheCacheTypeProp`

GetCacheType returns the CacheType field if non-nil, zero value otherwise.

### GetCacheTypeOk

`func (o *EntryCacheListResponseResourcesInner) GetCacheTypeOk() (*EnumentryCacheCacheTypeProp, bool)`

GetCacheTypeOk returns a tuple with the CacheType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheType

`func (o *EntryCacheListResponseResourcesInner) SetCacheType(v EnumentryCacheCacheTypeProp)`

SetCacheType sets CacheType field to given value.

### HasCacheType

`func (o *EntryCacheListResponseResourcesInner) HasCacheType() bool`

HasCacheType returns a boolean if a field has been set.

### GetCacheDirectory

`func (o *EntryCacheListResponseResourcesInner) GetCacheDirectory() string`

GetCacheDirectory returns the CacheDirectory field if non-nil, zero value otherwise.

### GetCacheDirectoryOk

`func (o *EntryCacheListResponseResourcesInner) GetCacheDirectoryOk() (*string, bool)`

GetCacheDirectoryOk returns a tuple with the CacheDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheDirectory

`func (o *EntryCacheListResponseResourcesInner) SetCacheDirectory(v string)`

SetCacheDirectory sets CacheDirectory field to given value.

### HasCacheDirectory

`func (o *EntryCacheListResponseResourcesInner) HasCacheDirectory() bool`

HasCacheDirectory returns a boolean if a field has been set.

### GetPersistentCache

`func (o *EntryCacheListResponseResourcesInner) GetPersistentCache() bool`

GetPersistentCache returns the PersistentCache field if non-nil, zero value otherwise.

### GetPersistentCacheOk

`func (o *EntryCacheListResponseResourcesInner) GetPersistentCacheOk() (*bool, bool)`

GetPersistentCacheOk returns a tuple with the PersistentCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentCache

`func (o *EntryCacheListResponseResourcesInner) SetPersistentCache(v bool)`

SetPersistentCache sets PersistentCache field to given value.

### HasPersistentCache

`func (o *EntryCacheListResponseResourcesInner) HasPersistentCache() bool`

HasPersistentCache returns a boolean if a field has been set.

### GetCompactEncoding

`func (o *EntryCacheListResponseResourcesInner) GetCompactEncoding() bool`

GetCompactEncoding returns the CompactEncoding field if non-nil, zero value otherwise.

### GetCompactEncodingOk

`func (o *EntryCacheListResponseResourcesInner) GetCompactEncodingOk() (*bool, bool)`

GetCompactEncodingOk returns a tuple with the CompactEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactEncoding

`func (o *EntryCacheListResponseResourcesInner) SetCompactEncoding(v bool)`

SetCompactEncoding sets CompactEncoding field to given value.

### HasCompactEncoding

`func (o *EntryCacheListResponseResourcesInner) HasCompactEncoding() bool`

HasCompactEncoding returns a boolean if a field has been set.

### GetDbCachePercent

`func (o *EntryCacheListResponseResourcesInner) GetDbCachePercent() int64`

GetDbCachePercent returns the DbCachePercent field if non-nil, zero value otherwise.

### GetDbCachePercentOk

`func (o *EntryCacheListResponseResourcesInner) GetDbCachePercentOk() (*int64, bool)`

GetDbCachePercentOk returns a tuple with the DbCachePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCachePercent

`func (o *EntryCacheListResponseResourcesInner) SetDbCachePercent(v int64)`

SetDbCachePercent sets DbCachePercent field to given value.

### HasDbCachePercent

`func (o *EntryCacheListResponseResourcesInner) HasDbCachePercent() bool`

HasDbCachePercent returns a boolean if a field has been set.

### GetDbCacheSize

`func (o *EntryCacheListResponseResourcesInner) GetDbCacheSize() string`

GetDbCacheSize returns the DbCacheSize field if non-nil, zero value otherwise.

### GetDbCacheSizeOk

`func (o *EntryCacheListResponseResourcesInner) GetDbCacheSizeOk() (*string, bool)`

GetDbCacheSizeOk returns a tuple with the DbCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCacheSize

`func (o *EntryCacheListResponseResourcesInner) SetDbCacheSize(v string)`

SetDbCacheSize sets DbCacheSize field to given value.

### HasDbCacheSize

`func (o *EntryCacheListResponseResourcesInner) HasDbCacheSize() bool`

HasDbCacheSize returns a boolean if a field has been set.

### GetJeProperty

`func (o *EntryCacheListResponseResourcesInner) GetJeProperty() []string`

GetJeProperty returns the JeProperty field if non-nil, zero value otherwise.

### GetJePropertyOk

`func (o *EntryCacheListResponseResourcesInner) GetJePropertyOk() (*[]string, bool)`

GetJePropertyOk returns a tuple with the JeProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJeProperty

`func (o *EntryCacheListResponseResourcesInner) SetJeProperty(v []string)`

SetJeProperty sets JeProperty field to given value.

### HasJeProperty

`func (o *EntryCacheListResponseResourcesInner) HasJeProperty() bool`

HasJeProperty returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *EntryCacheListResponseResourcesInner) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *EntryCacheListResponseResourcesInner) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *EntryCacheListResponseResourcesInner) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *EntryCacheListResponseResourcesInner) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetExcludeFilter

`func (o *EntryCacheListResponseResourcesInner) GetExcludeFilter() []string`

GetExcludeFilter returns the ExcludeFilter field if non-nil, zero value otherwise.

### GetExcludeFilterOk

`func (o *EntryCacheListResponseResourcesInner) GetExcludeFilterOk() (*[]string, bool)`

GetExcludeFilterOk returns a tuple with the ExcludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilter

`func (o *EntryCacheListResponseResourcesInner) SetExcludeFilter(v []string)`

SetExcludeFilter sets ExcludeFilter field to given value.

### HasExcludeFilter

`func (o *EntryCacheListResponseResourcesInner) HasExcludeFilter() bool`

HasExcludeFilter returns a boolean if a field has been set.

### GetDescription

`func (o *EntryCacheListResponseResourcesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntryCacheListResponseResourcesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntryCacheListResponseResourcesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntryCacheListResponseResourcesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *EntryCacheListResponseResourcesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EntryCacheListResponseResourcesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EntryCacheListResponseResourcesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCacheLevel

`func (o *EntryCacheListResponseResourcesInner) GetCacheLevel() int64`

GetCacheLevel returns the CacheLevel field if non-nil, zero value otherwise.

### GetCacheLevelOk

`func (o *EntryCacheListResponseResourcesInner) GetCacheLevelOk() (*int64, bool)`

GetCacheLevelOk returns a tuple with the CacheLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLevel

`func (o *EntryCacheListResponseResourcesInner) SetCacheLevel(v int64)`

SetCacheLevel sets CacheLevel field to given value.


### GetCacheUnindexedSearchResults

`func (o *EntryCacheListResponseResourcesInner) GetCacheUnindexedSearchResults() bool`

GetCacheUnindexedSearchResults returns the CacheUnindexedSearchResults field if non-nil, zero value otherwise.

### GetCacheUnindexedSearchResultsOk

`func (o *EntryCacheListResponseResourcesInner) GetCacheUnindexedSearchResultsOk() (*bool, bool)`

GetCacheUnindexedSearchResultsOk returns a tuple with the CacheUnindexedSearchResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheUnindexedSearchResults

`func (o *EntryCacheListResponseResourcesInner) SetCacheUnindexedSearchResults(v bool)`

SetCacheUnindexedSearchResults sets CacheUnindexedSearchResults field to given value.

### HasCacheUnindexedSearchResults

`func (o *EntryCacheListResponseResourcesInner) HasCacheUnindexedSearchResults() bool`

HasCacheUnindexedSearchResults returns a boolean if a field has been set.

### GetMeta

`func (o *EntryCacheListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EntryCacheListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EntryCacheListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EntryCacheListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *EntryCacheListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *EntryCacheListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *EntryCacheListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *EntryCacheListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetMaxMemoryPercent

`func (o *EntryCacheListResponseResourcesInner) GetMaxMemoryPercent() int64`

GetMaxMemoryPercent returns the MaxMemoryPercent field if non-nil, zero value otherwise.

### GetMaxMemoryPercentOk

`func (o *EntryCacheListResponseResourcesInner) GetMaxMemoryPercentOk() (*int64, bool)`

GetMaxMemoryPercentOk returns a tuple with the MaxMemoryPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryPercent

`func (o *EntryCacheListResponseResourcesInner) SetMaxMemoryPercent(v int64)`

SetMaxMemoryPercent sets MaxMemoryPercent field to given value.

### HasMaxMemoryPercent

`func (o *EntryCacheListResponseResourcesInner) HasMaxMemoryPercent() bool`

HasMaxMemoryPercent returns a boolean if a field has been set.

### GetOnlyCacheFrequentlyAccessed

`func (o *EntryCacheListResponseResourcesInner) GetOnlyCacheFrequentlyAccessed() bool`

GetOnlyCacheFrequentlyAccessed returns the OnlyCacheFrequentlyAccessed field if non-nil, zero value otherwise.

### GetOnlyCacheFrequentlyAccessedOk

`func (o *EntryCacheListResponseResourcesInner) GetOnlyCacheFrequentlyAccessedOk() (*bool, bool)`

GetOnlyCacheFrequentlyAccessedOk returns a tuple with the OnlyCacheFrequentlyAccessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyCacheFrequentlyAccessed

`func (o *EntryCacheListResponseResourcesInner) SetOnlyCacheFrequentlyAccessed(v bool)`

SetOnlyCacheFrequentlyAccessed sets OnlyCacheFrequentlyAccessed field to given value.

### HasOnlyCacheFrequentlyAccessed

`func (o *EntryCacheListResponseResourcesInner) HasOnlyCacheFrequentlyAccessed() bool`

HasOnlyCacheFrequentlyAccessed returns a boolean if a field has been set.

### GetMinCacheEntryValueCount

`func (o *EntryCacheListResponseResourcesInner) GetMinCacheEntryValueCount() int64`

GetMinCacheEntryValueCount returns the MinCacheEntryValueCount field if non-nil, zero value otherwise.

### GetMinCacheEntryValueCountOk

`func (o *EntryCacheListResponseResourcesInner) GetMinCacheEntryValueCountOk() (*int64, bool)`

GetMinCacheEntryValueCountOk returns a tuple with the MinCacheEntryValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCacheEntryValueCount

`func (o *EntryCacheListResponseResourcesInner) SetMinCacheEntryValueCount(v int64)`

SetMinCacheEntryValueCount sets MinCacheEntryValueCount field to given value.

### HasMinCacheEntryValueCount

`func (o *EntryCacheListResponseResourcesInner) HasMinCacheEntryValueCount() bool`

HasMinCacheEntryValueCount returns a boolean if a field has been set.

### GetMinCacheEntryAttribute

`func (o *EntryCacheListResponseResourcesInner) GetMinCacheEntryAttribute() []string`

GetMinCacheEntryAttribute returns the MinCacheEntryAttribute field if non-nil, zero value otherwise.

### GetMinCacheEntryAttributeOk

`func (o *EntryCacheListResponseResourcesInner) GetMinCacheEntryAttributeOk() (*[]string, bool)`

GetMinCacheEntryAttributeOk returns a tuple with the MinCacheEntryAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCacheEntryAttribute

`func (o *EntryCacheListResponseResourcesInner) SetMinCacheEntryAttribute(v []string)`

SetMinCacheEntryAttribute sets MinCacheEntryAttribute field to given value.

### HasMinCacheEntryAttribute

`func (o *EntryCacheListResponseResourcesInner) HasMinCacheEntryAttribute() bool`

HasMinCacheEntryAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


