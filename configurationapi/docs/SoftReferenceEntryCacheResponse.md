# SoftReferenceEntryCacheResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsoftReferenceEntryCacheSchemaUrn**](EnumsoftReferenceEntryCacheSchemaUrn.md) |  | 
**Id** | **string** | Name of the Entry Cache | 
**IncludeFilter** | Pointer to **[]string** | The set of filters that define the entries that should be included in the cache. | [optional] 
**ExcludeFilter** | Pointer to **[]string** | The set of filters that define the entries that should be excluded from the cache. | [optional] 
**Description** | Pointer to **string** | A description for this Entry Cache | [optional] 
**Enabled** | **bool** | Indicates whether the Entry Cache is enabled. | 
**CacheLevel** | **int64** | Specifies the cache level in the cache order if more than one instance of the cache is configured. | 
**CacheUnindexedSearchResults** | Pointer to **bool** | Indicates whether the entry cache should be updated with entries that have been returned to the client during the course of processing an unindexed search. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewSoftReferenceEntryCacheResponse

`func NewSoftReferenceEntryCacheResponse(schemas []EnumsoftReferenceEntryCacheSchemaUrn, id string, enabled bool, cacheLevel int64, ) *SoftReferenceEntryCacheResponse`

NewSoftReferenceEntryCacheResponse instantiates a new SoftReferenceEntryCacheResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftReferenceEntryCacheResponseWithDefaults

`func NewSoftReferenceEntryCacheResponseWithDefaults() *SoftReferenceEntryCacheResponse`

NewSoftReferenceEntryCacheResponseWithDefaults instantiates a new SoftReferenceEntryCacheResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SoftReferenceEntryCacheResponse) GetSchemas() []EnumsoftReferenceEntryCacheSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SoftReferenceEntryCacheResponse) GetSchemasOk() (*[]EnumsoftReferenceEntryCacheSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SoftReferenceEntryCacheResponse) SetSchemas(v []EnumsoftReferenceEntryCacheSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *SoftReferenceEntryCacheResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SoftReferenceEntryCacheResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SoftReferenceEntryCacheResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIncludeFilter

`func (o *SoftReferenceEntryCacheResponse) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *SoftReferenceEntryCacheResponse) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *SoftReferenceEntryCacheResponse) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *SoftReferenceEntryCacheResponse) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetExcludeFilter

`func (o *SoftReferenceEntryCacheResponse) GetExcludeFilter() []string`

GetExcludeFilter returns the ExcludeFilter field if non-nil, zero value otherwise.

### GetExcludeFilterOk

`func (o *SoftReferenceEntryCacheResponse) GetExcludeFilterOk() (*[]string, bool)`

GetExcludeFilterOk returns a tuple with the ExcludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilter

`func (o *SoftReferenceEntryCacheResponse) SetExcludeFilter(v []string)`

SetExcludeFilter sets ExcludeFilter field to given value.

### HasExcludeFilter

`func (o *SoftReferenceEntryCacheResponse) HasExcludeFilter() bool`

HasExcludeFilter returns a boolean if a field has been set.

### GetDescription

`func (o *SoftReferenceEntryCacheResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SoftReferenceEntryCacheResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SoftReferenceEntryCacheResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SoftReferenceEntryCacheResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SoftReferenceEntryCacheResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SoftReferenceEntryCacheResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SoftReferenceEntryCacheResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCacheLevel

`func (o *SoftReferenceEntryCacheResponse) GetCacheLevel() int64`

GetCacheLevel returns the CacheLevel field if non-nil, zero value otherwise.

### GetCacheLevelOk

`func (o *SoftReferenceEntryCacheResponse) GetCacheLevelOk() (*int64, bool)`

GetCacheLevelOk returns a tuple with the CacheLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLevel

`func (o *SoftReferenceEntryCacheResponse) SetCacheLevel(v int64)`

SetCacheLevel sets CacheLevel field to given value.


### GetCacheUnindexedSearchResults

`func (o *SoftReferenceEntryCacheResponse) GetCacheUnindexedSearchResults() bool`

GetCacheUnindexedSearchResults returns the CacheUnindexedSearchResults field if non-nil, zero value otherwise.

### GetCacheUnindexedSearchResultsOk

`func (o *SoftReferenceEntryCacheResponse) GetCacheUnindexedSearchResultsOk() (*bool, bool)`

GetCacheUnindexedSearchResultsOk returns a tuple with the CacheUnindexedSearchResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheUnindexedSearchResults

`func (o *SoftReferenceEntryCacheResponse) SetCacheUnindexedSearchResults(v bool)`

SetCacheUnindexedSearchResults sets CacheUnindexedSearchResults field to given value.

### HasCacheUnindexedSearchResults

`func (o *SoftReferenceEntryCacheResponse) HasCacheUnindexedSearchResults() bool`

HasCacheUnindexedSearchResults returns a boolean if a field has been set.

### GetMeta

`func (o *SoftReferenceEntryCacheResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SoftReferenceEntryCacheResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SoftReferenceEntryCacheResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SoftReferenceEntryCacheResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SoftReferenceEntryCacheResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SoftReferenceEntryCacheResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SoftReferenceEntryCacheResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SoftReferenceEntryCacheResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


