# AddLocalDbCompositeIndexRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumlocalDbCompositeIndexSchemaUrn**](EnumlocalDbCompositeIndexSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Local DB Composite Index | [optional] 
**IndexFilterPattern** | **string** | A filter pattern that identifies which entries to include in the index. | 
**IndexBaseDNPattern** | Pointer to **string** | An optional base DN pattern that identifies portions of the DIT in which entries to index may exist. | [optional] 
**IndexEntryLimit** | Pointer to **int64** | The maximum number of entries that any single index key will be allowed to match before the server stops maintaining the ID set for that index key. | [optional] 
**PrimeIndex** | Pointer to **bool** | Indicates whether the server should load the contents of this index into memory when the backend is being opened. | [optional] 
**PrimeInternalNodesOnly** | Pointer to **bool** | Indicates whether to only prime the internal nodes of the index database, rather than priming both internal and leaf nodes. | [optional] 
**CacheMode** | Pointer to [**EnumlocalDbCompositeIndexCacheModeProp**](EnumlocalDbCompositeIndexCacheModeProp.md) |  | [optional] 
**IndexName** | **string** | Name of the new Local DB Composite Index | 

## Methods

### NewAddLocalDbCompositeIndexRequest

`func NewAddLocalDbCompositeIndexRequest(indexFilterPattern string, indexName string, ) *AddLocalDbCompositeIndexRequest`

NewAddLocalDbCompositeIndexRequest instantiates a new AddLocalDbCompositeIndexRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLocalDbCompositeIndexRequestWithDefaults

`func NewAddLocalDbCompositeIndexRequestWithDefaults() *AddLocalDbCompositeIndexRequest`

NewAddLocalDbCompositeIndexRequestWithDefaults instantiates a new AddLocalDbCompositeIndexRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddLocalDbCompositeIndexRequest) GetSchemas() []EnumlocalDbCompositeIndexSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLocalDbCompositeIndexRequest) GetSchemasOk() (*[]EnumlocalDbCompositeIndexSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLocalDbCompositeIndexRequest) SetSchemas(v []EnumlocalDbCompositeIndexSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddLocalDbCompositeIndexRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddLocalDbCompositeIndexRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLocalDbCompositeIndexRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLocalDbCompositeIndexRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLocalDbCompositeIndexRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIndexFilterPattern

`func (o *AddLocalDbCompositeIndexRequest) GetIndexFilterPattern() string`

GetIndexFilterPattern returns the IndexFilterPattern field if non-nil, zero value otherwise.

### GetIndexFilterPatternOk

`func (o *AddLocalDbCompositeIndexRequest) GetIndexFilterPatternOk() (*string, bool)`

GetIndexFilterPatternOk returns a tuple with the IndexFilterPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexFilterPattern

`func (o *AddLocalDbCompositeIndexRequest) SetIndexFilterPattern(v string)`

SetIndexFilterPattern sets IndexFilterPattern field to given value.


### GetIndexBaseDNPattern

`func (o *AddLocalDbCompositeIndexRequest) GetIndexBaseDNPattern() string`

GetIndexBaseDNPattern returns the IndexBaseDNPattern field if non-nil, zero value otherwise.

### GetIndexBaseDNPatternOk

`func (o *AddLocalDbCompositeIndexRequest) GetIndexBaseDNPatternOk() (*string, bool)`

GetIndexBaseDNPatternOk returns a tuple with the IndexBaseDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexBaseDNPattern

`func (o *AddLocalDbCompositeIndexRequest) SetIndexBaseDNPattern(v string)`

SetIndexBaseDNPattern sets IndexBaseDNPattern field to given value.

### HasIndexBaseDNPattern

`func (o *AddLocalDbCompositeIndexRequest) HasIndexBaseDNPattern() bool`

HasIndexBaseDNPattern returns a boolean if a field has been set.

### GetIndexEntryLimit

`func (o *AddLocalDbCompositeIndexRequest) GetIndexEntryLimit() int64`

GetIndexEntryLimit returns the IndexEntryLimit field if non-nil, zero value otherwise.

### GetIndexEntryLimitOk

`func (o *AddLocalDbCompositeIndexRequest) GetIndexEntryLimitOk() (*int64, bool)`

GetIndexEntryLimitOk returns a tuple with the IndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexEntryLimit

`func (o *AddLocalDbCompositeIndexRequest) SetIndexEntryLimit(v int64)`

SetIndexEntryLimit sets IndexEntryLimit field to given value.

### HasIndexEntryLimit

`func (o *AddLocalDbCompositeIndexRequest) HasIndexEntryLimit() bool`

HasIndexEntryLimit returns a boolean if a field has been set.

### GetPrimeIndex

`func (o *AddLocalDbCompositeIndexRequest) GetPrimeIndex() bool`

GetPrimeIndex returns the PrimeIndex field if non-nil, zero value otherwise.

### GetPrimeIndexOk

`func (o *AddLocalDbCompositeIndexRequest) GetPrimeIndexOk() (*bool, bool)`

GetPrimeIndexOk returns a tuple with the PrimeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeIndex

`func (o *AddLocalDbCompositeIndexRequest) SetPrimeIndex(v bool)`

SetPrimeIndex sets PrimeIndex field to given value.

### HasPrimeIndex

`func (o *AddLocalDbCompositeIndexRequest) HasPrimeIndex() bool`

HasPrimeIndex returns a boolean if a field has been set.

### GetPrimeInternalNodesOnly

`func (o *AddLocalDbCompositeIndexRequest) GetPrimeInternalNodesOnly() bool`

GetPrimeInternalNodesOnly returns the PrimeInternalNodesOnly field if non-nil, zero value otherwise.

### GetPrimeInternalNodesOnlyOk

`func (o *AddLocalDbCompositeIndexRequest) GetPrimeInternalNodesOnlyOk() (*bool, bool)`

GetPrimeInternalNodesOnlyOk returns a tuple with the PrimeInternalNodesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeInternalNodesOnly

`func (o *AddLocalDbCompositeIndexRequest) SetPrimeInternalNodesOnly(v bool)`

SetPrimeInternalNodesOnly sets PrimeInternalNodesOnly field to given value.

### HasPrimeInternalNodesOnly

`func (o *AddLocalDbCompositeIndexRequest) HasPrimeInternalNodesOnly() bool`

HasPrimeInternalNodesOnly returns a boolean if a field has been set.

### GetCacheMode

`func (o *AddLocalDbCompositeIndexRequest) GetCacheMode() EnumlocalDbCompositeIndexCacheModeProp`

GetCacheMode returns the CacheMode field if non-nil, zero value otherwise.

### GetCacheModeOk

`func (o *AddLocalDbCompositeIndexRequest) GetCacheModeOk() (*EnumlocalDbCompositeIndexCacheModeProp, bool)`

GetCacheModeOk returns a tuple with the CacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMode

`func (o *AddLocalDbCompositeIndexRequest) SetCacheMode(v EnumlocalDbCompositeIndexCacheModeProp)`

SetCacheMode sets CacheMode field to given value.

### HasCacheMode

`func (o *AddLocalDbCompositeIndexRequest) HasCacheMode() bool`

HasCacheMode returns a boolean if a field has been set.

### GetIndexName

`func (o *AddLocalDbCompositeIndexRequest) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *AddLocalDbCompositeIndexRequest) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *AddLocalDbCompositeIndexRequest) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


