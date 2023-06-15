# LocalDbCompositeIndexResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Local DB Composite Index | 
**Schemas** | Pointer to [**[]EnumlocalDbCompositeIndexSchemaUrn**](EnumlocalDbCompositeIndexSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Local DB Composite Index | [optional] 
**IndexFilterPattern** | **string** | A filter pattern that identifies which entries to include in the index. | 
**IndexBaseDNPattern** | Pointer to **string** | An optional base DN pattern that identifies portions of the DIT in which entries to index may exist. | [optional] 
**IndexEntryLimit** | Pointer to **int64** | The maximum number of entries that any single index key will be allowed to match before the server stops maintaining the ID set for that index key. | [optional] 
**PrimeIndex** | Pointer to **bool** | Indicates whether the server should load the contents of this index into memory when the backend is being opened. | [optional] 
**PrimeInternalNodesOnly** | Pointer to **bool** | Indicates whether to only prime the internal nodes of the index database, rather than priming both internal and leaf nodes. | [optional] 
**CacheMode** | Pointer to [**EnumlocalDbCompositeIndexCacheModeProp**](EnumlocalDbCompositeIndexCacheModeProp.md) |  | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewLocalDbCompositeIndexResponse

`func NewLocalDbCompositeIndexResponse(id string, indexFilterPattern string, ) *LocalDbCompositeIndexResponse`

NewLocalDbCompositeIndexResponse instantiates a new LocalDbCompositeIndexResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalDbCompositeIndexResponseWithDefaults

`func NewLocalDbCompositeIndexResponseWithDefaults() *LocalDbCompositeIndexResponse`

NewLocalDbCompositeIndexResponseWithDefaults instantiates a new LocalDbCompositeIndexResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocalDbCompositeIndexResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalDbCompositeIndexResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalDbCompositeIndexResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *LocalDbCompositeIndexResponse) GetSchemas() []EnumlocalDbCompositeIndexSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LocalDbCompositeIndexResponse) GetSchemasOk() (*[]EnumlocalDbCompositeIndexSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LocalDbCompositeIndexResponse) SetSchemas(v []EnumlocalDbCompositeIndexSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *LocalDbCompositeIndexResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *LocalDbCompositeIndexResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LocalDbCompositeIndexResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LocalDbCompositeIndexResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LocalDbCompositeIndexResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIndexFilterPattern

`func (o *LocalDbCompositeIndexResponse) GetIndexFilterPattern() string`

GetIndexFilterPattern returns the IndexFilterPattern field if non-nil, zero value otherwise.

### GetIndexFilterPatternOk

`func (o *LocalDbCompositeIndexResponse) GetIndexFilterPatternOk() (*string, bool)`

GetIndexFilterPatternOk returns a tuple with the IndexFilterPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexFilterPattern

`func (o *LocalDbCompositeIndexResponse) SetIndexFilterPattern(v string)`

SetIndexFilterPattern sets IndexFilterPattern field to given value.


### GetIndexBaseDNPattern

`func (o *LocalDbCompositeIndexResponse) GetIndexBaseDNPattern() string`

GetIndexBaseDNPattern returns the IndexBaseDNPattern field if non-nil, zero value otherwise.

### GetIndexBaseDNPatternOk

`func (o *LocalDbCompositeIndexResponse) GetIndexBaseDNPatternOk() (*string, bool)`

GetIndexBaseDNPatternOk returns a tuple with the IndexBaseDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexBaseDNPattern

`func (o *LocalDbCompositeIndexResponse) SetIndexBaseDNPattern(v string)`

SetIndexBaseDNPattern sets IndexBaseDNPattern field to given value.

### HasIndexBaseDNPattern

`func (o *LocalDbCompositeIndexResponse) HasIndexBaseDNPattern() bool`

HasIndexBaseDNPattern returns a boolean if a field has been set.

### GetIndexEntryLimit

`func (o *LocalDbCompositeIndexResponse) GetIndexEntryLimit() int64`

GetIndexEntryLimit returns the IndexEntryLimit field if non-nil, zero value otherwise.

### GetIndexEntryLimitOk

`func (o *LocalDbCompositeIndexResponse) GetIndexEntryLimitOk() (*int64, bool)`

GetIndexEntryLimitOk returns a tuple with the IndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexEntryLimit

`func (o *LocalDbCompositeIndexResponse) SetIndexEntryLimit(v int64)`

SetIndexEntryLimit sets IndexEntryLimit field to given value.

### HasIndexEntryLimit

`func (o *LocalDbCompositeIndexResponse) HasIndexEntryLimit() bool`

HasIndexEntryLimit returns a boolean if a field has been set.

### GetPrimeIndex

`func (o *LocalDbCompositeIndexResponse) GetPrimeIndex() bool`

GetPrimeIndex returns the PrimeIndex field if non-nil, zero value otherwise.

### GetPrimeIndexOk

`func (o *LocalDbCompositeIndexResponse) GetPrimeIndexOk() (*bool, bool)`

GetPrimeIndexOk returns a tuple with the PrimeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeIndex

`func (o *LocalDbCompositeIndexResponse) SetPrimeIndex(v bool)`

SetPrimeIndex sets PrimeIndex field to given value.

### HasPrimeIndex

`func (o *LocalDbCompositeIndexResponse) HasPrimeIndex() bool`

HasPrimeIndex returns a boolean if a field has been set.

### GetPrimeInternalNodesOnly

`func (o *LocalDbCompositeIndexResponse) GetPrimeInternalNodesOnly() bool`

GetPrimeInternalNodesOnly returns the PrimeInternalNodesOnly field if non-nil, zero value otherwise.

### GetPrimeInternalNodesOnlyOk

`func (o *LocalDbCompositeIndexResponse) GetPrimeInternalNodesOnlyOk() (*bool, bool)`

GetPrimeInternalNodesOnlyOk returns a tuple with the PrimeInternalNodesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeInternalNodesOnly

`func (o *LocalDbCompositeIndexResponse) SetPrimeInternalNodesOnly(v bool)`

SetPrimeInternalNodesOnly sets PrimeInternalNodesOnly field to given value.

### HasPrimeInternalNodesOnly

`func (o *LocalDbCompositeIndexResponse) HasPrimeInternalNodesOnly() bool`

HasPrimeInternalNodesOnly returns a boolean if a field has been set.

### GetCacheMode

`func (o *LocalDbCompositeIndexResponse) GetCacheMode() EnumlocalDbCompositeIndexCacheModeProp`

GetCacheMode returns the CacheMode field if non-nil, zero value otherwise.

### GetCacheModeOk

`func (o *LocalDbCompositeIndexResponse) GetCacheModeOk() (*EnumlocalDbCompositeIndexCacheModeProp, bool)`

GetCacheModeOk returns a tuple with the CacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMode

`func (o *LocalDbCompositeIndexResponse) SetCacheMode(v EnumlocalDbCompositeIndexCacheModeProp)`

SetCacheMode sets CacheMode field to given value.

### HasCacheMode

`func (o *LocalDbCompositeIndexResponse) HasCacheMode() bool`

HasCacheMode returns a boolean if a field has been set.

### GetMeta

`func (o *LocalDbCompositeIndexResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LocalDbCompositeIndexResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LocalDbCompositeIndexResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LocalDbCompositeIndexResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *LocalDbCompositeIndexResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *LocalDbCompositeIndexResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *LocalDbCompositeIndexResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *LocalDbCompositeIndexResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


