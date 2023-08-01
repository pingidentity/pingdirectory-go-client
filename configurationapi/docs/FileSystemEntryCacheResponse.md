# FileSystemEntryCacheResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumfileSystemEntryCacheSchemaUrn**](EnumfileSystemEntryCacheSchemaUrn.md) |  | 
**Id** | **string** | Name of the Entry Cache | 
**MaxMemorySize** | Pointer to **string** | The maximum size of the entry cache in bytes. | [optional] 
**MaxEntries** | Pointer to **int64** | The maximum number of entries allowed in the cache. | [optional] 
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

## Methods

### NewFileSystemEntryCacheResponse

`func NewFileSystemEntryCacheResponse(schemas []EnumfileSystemEntryCacheSchemaUrn, id string, enabled bool, cacheLevel int64, ) *FileSystemEntryCacheResponse`

NewFileSystemEntryCacheResponse instantiates a new FileSystemEntryCacheResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSystemEntryCacheResponseWithDefaults

`func NewFileSystemEntryCacheResponseWithDefaults() *FileSystemEntryCacheResponse`

NewFileSystemEntryCacheResponseWithDefaults instantiates a new FileSystemEntryCacheResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *FileSystemEntryCacheResponse) GetSchemas() []EnumfileSystemEntryCacheSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FileSystemEntryCacheResponse) GetSchemasOk() (*[]EnumfileSystemEntryCacheSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FileSystemEntryCacheResponse) SetSchemas(v []EnumfileSystemEntryCacheSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *FileSystemEntryCacheResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileSystemEntryCacheResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileSystemEntryCacheResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMaxMemorySize

`func (o *FileSystemEntryCacheResponse) GetMaxMemorySize() string`

GetMaxMemorySize returns the MaxMemorySize field if non-nil, zero value otherwise.

### GetMaxMemorySizeOk

`func (o *FileSystemEntryCacheResponse) GetMaxMemorySizeOk() (*string, bool)`

GetMaxMemorySizeOk returns a tuple with the MaxMemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemorySize

`func (o *FileSystemEntryCacheResponse) SetMaxMemorySize(v string)`

SetMaxMemorySize sets MaxMemorySize field to given value.

### HasMaxMemorySize

`func (o *FileSystemEntryCacheResponse) HasMaxMemorySize() bool`

HasMaxMemorySize returns a boolean if a field has been set.

### GetMaxEntries

`func (o *FileSystemEntryCacheResponse) GetMaxEntries() int64`

GetMaxEntries returns the MaxEntries field if non-nil, zero value otherwise.

### GetMaxEntriesOk

`func (o *FileSystemEntryCacheResponse) GetMaxEntriesOk() (*int64, bool)`

GetMaxEntriesOk returns a tuple with the MaxEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEntries

`func (o *FileSystemEntryCacheResponse) SetMaxEntries(v int64)`

SetMaxEntries sets MaxEntries field to given value.

### HasMaxEntries

`func (o *FileSystemEntryCacheResponse) HasMaxEntries() bool`

HasMaxEntries returns a boolean if a field has been set.

### GetCacheType

`func (o *FileSystemEntryCacheResponse) GetCacheType() EnumentryCacheCacheTypeProp`

GetCacheType returns the CacheType field if non-nil, zero value otherwise.

### GetCacheTypeOk

`func (o *FileSystemEntryCacheResponse) GetCacheTypeOk() (*EnumentryCacheCacheTypeProp, bool)`

GetCacheTypeOk returns a tuple with the CacheType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheType

`func (o *FileSystemEntryCacheResponse) SetCacheType(v EnumentryCacheCacheTypeProp)`

SetCacheType sets CacheType field to given value.

### HasCacheType

`func (o *FileSystemEntryCacheResponse) HasCacheType() bool`

HasCacheType returns a boolean if a field has been set.

### GetCacheDirectory

`func (o *FileSystemEntryCacheResponse) GetCacheDirectory() string`

GetCacheDirectory returns the CacheDirectory field if non-nil, zero value otherwise.

### GetCacheDirectoryOk

`func (o *FileSystemEntryCacheResponse) GetCacheDirectoryOk() (*string, bool)`

GetCacheDirectoryOk returns a tuple with the CacheDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheDirectory

`func (o *FileSystemEntryCacheResponse) SetCacheDirectory(v string)`

SetCacheDirectory sets CacheDirectory field to given value.

### HasCacheDirectory

`func (o *FileSystemEntryCacheResponse) HasCacheDirectory() bool`

HasCacheDirectory returns a boolean if a field has been set.

### GetPersistentCache

`func (o *FileSystemEntryCacheResponse) GetPersistentCache() bool`

GetPersistentCache returns the PersistentCache field if non-nil, zero value otherwise.

### GetPersistentCacheOk

`func (o *FileSystemEntryCacheResponse) GetPersistentCacheOk() (*bool, bool)`

GetPersistentCacheOk returns a tuple with the PersistentCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentCache

`func (o *FileSystemEntryCacheResponse) SetPersistentCache(v bool)`

SetPersistentCache sets PersistentCache field to given value.

### HasPersistentCache

`func (o *FileSystemEntryCacheResponse) HasPersistentCache() bool`

HasPersistentCache returns a boolean if a field has been set.

### GetCompactEncoding

`func (o *FileSystemEntryCacheResponse) GetCompactEncoding() bool`

GetCompactEncoding returns the CompactEncoding field if non-nil, zero value otherwise.

### GetCompactEncodingOk

`func (o *FileSystemEntryCacheResponse) GetCompactEncodingOk() (*bool, bool)`

GetCompactEncodingOk returns a tuple with the CompactEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactEncoding

`func (o *FileSystemEntryCacheResponse) SetCompactEncoding(v bool)`

SetCompactEncoding sets CompactEncoding field to given value.

### HasCompactEncoding

`func (o *FileSystemEntryCacheResponse) HasCompactEncoding() bool`

HasCompactEncoding returns a boolean if a field has been set.

### GetDbCachePercent

`func (o *FileSystemEntryCacheResponse) GetDbCachePercent() int64`

GetDbCachePercent returns the DbCachePercent field if non-nil, zero value otherwise.

### GetDbCachePercentOk

`func (o *FileSystemEntryCacheResponse) GetDbCachePercentOk() (*int64, bool)`

GetDbCachePercentOk returns a tuple with the DbCachePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCachePercent

`func (o *FileSystemEntryCacheResponse) SetDbCachePercent(v int64)`

SetDbCachePercent sets DbCachePercent field to given value.

### HasDbCachePercent

`func (o *FileSystemEntryCacheResponse) HasDbCachePercent() bool`

HasDbCachePercent returns a boolean if a field has been set.

### GetDbCacheSize

`func (o *FileSystemEntryCacheResponse) GetDbCacheSize() string`

GetDbCacheSize returns the DbCacheSize field if non-nil, zero value otherwise.

### GetDbCacheSizeOk

`func (o *FileSystemEntryCacheResponse) GetDbCacheSizeOk() (*string, bool)`

GetDbCacheSizeOk returns a tuple with the DbCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCacheSize

`func (o *FileSystemEntryCacheResponse) SetDbCacheSize(v string)`

SetDbCacheSize sets DbCacheSize field to given value.

### HasDbCacheSize

`func (o *FileSystemEntryCacheResponse) HasDbCacheSize() bool`

HasDbCacheSize returns a boolean if a field has been set.

### GetJeProperty

`func (o *FileSystemEntryCacheResponse) GetJeProperty() []string`

GetJeProperty returns the JeProperty field if non-nil, zero value otherwise.

### GetJePropertyOk

`func (o *FileSystemEntryCacheResponse) GetJePropertyOk() (*[]string, bool)`

GetJePropertyOk returns a tuple with the JeProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJeProperty

`func (o *FileSystemEntryCacheResponse) SetJeProperty(v []string)`

SetJeProperty sets JeProperty field to given value.

### HasJeProperty

`func (o *FileSystemEntryCacheResponse) HasJeProperty() bool`

HasJeProperty returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *FileSystemEntryCacheResponse) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *FileSystemEntryCacheResponse) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *FileSystemEntryCacheResponse) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *FileSystemEntryCacheResponse) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetExcludeFilter

`func (o *FileSystemEntryCacheResponse) GetExcludeFilter() []string`

GetExcludeFilter returns the ExcludeFilter field if non-nil, zero value otherwise.

### GetExcludeFilterOk

`func (o *FileSystemEntryCacheResponse) GetExcludeFilterOk() (*[]string, bool)`

GetExcludeFilterOk returns a tuple with the ExcludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilter

`func (o *FileSystemEntryCacheResponse) SetExcludeFilter(v []string)`

SetExcludeFilter sets ExcludeFilter field to given value.

### HasExcludeFilter

`func (o *FileSystemEntryCacheResponse) HasExcludeFilter() bool`

HasExcludeFilter returns a boolean if a field has been set.

### GetDescription

`func (o *FileSystemEntryCacheResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FileSystemEntryCacheResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FileSystemEntryCacheResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FileSystemEntryCacheResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *FileSystemEntryCacheResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FileSystemEntryCacheResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FileSystemEntryCacheResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCacheLevel

`func (o *FileSystemEntryCacheResponse) GetCacheLevel() int64`

GetCacheLevel returns the CacheLevel field if non-nil, zero value otherwise.

### GetCacheLevelOk

`func (o *FileSystemEntryCacheResponse) GetCacheLevelOk() (*int64, bool)`

GetCacheLevelOk returns a tuple with the CacheLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLevel

`func (o *FileSystemEntryCacheResponse) SetCacheLevel(v int64)`

SetCacheLevel sets CacheLevel field to given value.


### GetCacheUnindexedSearchResults

`func (o *FileSystemEntryCacheResponse) GetCacheUnindexedSearchResults() bool`

GetCacheUnindexedSearchResults returns the CacheUnindexedSearchResults field if non-nil, zero value otherwise.

### GetCacheUnindexedSearchResultsOk

`func (o *FileSystemEntryCacheResponse) GetCacheUnindexedSearchResultsOk() (*bool, bool)`

GetCacheUnindexedSearchResultsOk returns a tuple with the CacheUnindexedSearchResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheUnindexedSearchResults

`func (o *FileSystemEntryCacheResponse) SetCacheUnindexedSearchResults(v bool)`

SetCacheUnindexedSearchResults sets CacheUnindexedSearchResults field to given value.

### HasCacheUnindexedSearchResults

`func (o *FileSystemEntryCacheResponse) HasCacheUnindexedSearchResults() bool`

HasCacheUnindexedSearchResults returns a boolean if a field has been set.

### GetMeta

`func (o *FileSystemEntryCacheResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FileSystemEntryCacheResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FileSystemEntryCacheResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FileSystemEntryCacheResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *FileSystemEntryCacheResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *FileSystemEntryCacheResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *FileSystemEntryCacheResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *FileSystemEntryCacheResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


