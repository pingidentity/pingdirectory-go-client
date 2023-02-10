# AddLocalDbIndexRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexName** | **string** | Name of the new Local DB Index | 
**Schemas** | Pointer to [**[]EnumlocalDbIndexSchemaUrn**](EnumlocalDbIndexSchemaUrn.md) |  | [optional] 
**Attribute** | **string** | Specifies the name of the attribute for which the index is to be maintained. | 
**IndexEntryLimit** | Pointer to **int32** | Specifies the maximum number of entries that are allowed to match a given index key before that particular index key is no longer maintained. | [optional] 
**SubstringIndexEntryLimit** | Pointer to **int32** | Specifies, for substring indexes, the maximum number of entries that are allowed to match a given index key before that particular index key is no longer maintained. Setting a large limit can dramatically increase the database size on disk and have a big impact on server performance if the indexed attribute is modified frequently. When a very large limit is required, creating a dedicated composite index with an index-filter-pattern of (attr&#x3D;*?*) will give the best balance between search and update performance. | [optional] 
**MaintainMatchCountForKeysExceedingEntryLimit** | Pointer to **bool** | Indicates whether to continue to maintain a count of the number of matching entries for an index key even after that count exceeds the index entry limit. | [optional] 
**IndexType** | [**[]EnumlocalDbIndexIndexTypeProp**](EnumlocalDbIndexIndexTypeProp.md) |  | 
**SubstringLength** | Pointer to **int32** | The length of substrings in a substring index. | [optional] 
**PrimeIndex** | Pointer to **bool** | If this option is enabled and this index&#39;s backend is configured to prime indexes, then this index will be loaded at startup. | [optional] 
**PrimeInternalNodesOnly** | Pointer to **bool** | If this option is enabled and this index&#39;s backend is configured to prime indexes using the preload method, then only the internal database nodes (i.e., the database keys but not values) should be primed when the backend is initialized. | [optional] 
**EqualityIndexFilter** | Pointer to **[]string** | A search filter that may be used in conjunction with an equality component for the associated attribute type. If an equality index filter is defined, then an additional equality index will be maintained for the associated attribute, but only for entries which match the provided filter. Further, the index will be used only for searches containing an equality component with the associated attribute type ANDed with this filter. | [optional] 
**MaintainEqualityIndexWithoutFilter** | Pointer to **bool** | Indicates whether to maintain a separate equality index for the associated attribute without any filter, in addition to maintaining an index for each equality index filter that is defined. If this is false, then the attribute will not be indexed for equality by itself but only in conjunction with the defined equality index filters. | [optional] 
**CacheMode** | Pointer to [**EnumlocalDbIndexCacheModeProp**](EnumlocalDbIndexCacheModeProp.md) |  | [optional] 

## Methods

### NewAddLocalDbIndexRequest

`func NewAddLocalDbIndexRequest(indexName string, attribute string, indexType []EnumlocalDbIndexIndexTypeProp, ) *AddLocalDbIndexRequest`

NewAddLocalDbIndexRequest instantiates a new AddLocalDbIndexRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLocalDbIndexRequestWithDefaults

`func NewAddLocalDbIndexRequestWithDefaults() *AddLocalDbIndexRequest`

NewAddLocalDbIndexRequestWithDefaults instantiates a new AddLocalDbIndexRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexName

`func (o *AddLocalDbIndexRequest) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *AddLocalDbIndexRequest) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *AddLocalDbIndexRequest) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.


### GetSchemas

`func (o *AddLocalDbIndexRequest) GetSchemas() []EnumlocalDbIndexSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLocalDbIndexRequest) GetSchemasOk() (*[]EnumlocalDbIndexSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLocalDbIndexRequest) SetSchemas(v []EnumlocalDbIndexSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddLocalDbIndexRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetAttribute

`func (o *AddLocalDbIndexRequest) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *AddLocalDbIndexRequest) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *AddLocalDbIndexRequest) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetIndexEntryLimit

`func (o *AddLocalDbIndexRequest) GetIndexEntryLimit() int32`

GetIndexEntryLimit returns the IndexEntryLimit field if non-nil, zero value otherwise.

### GetIndexEntryLimitOk

`func (o *AddLocalDbIndexRequest) GetIndexEntryLimitOk() (*int32, bool)`

GetIndexEntryLimitOk returns a tuple with the IndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexEntryLimit

`func (o *AddLocalDbIndexRequest) SetIndexEntryLimit(v int32)`

SetIndexEntryLimit sets IndexEntryLimit field to given value.

### HasIndexEntryLimit

`func (o *AddLocalDbIndexRequest) HasIndexEntryLimit() bool`

HasIndexEntryLimit returns a boolean if a field has been set.

### GetSubstringIndexEntryLimit

`func (o *AddLocalDbIndexRequest) GetSubstringIndexEntryLimit() int32`

GetSubstringIndexEntryLimit returns the SubstringIndexEntryLimit field if non-nil, zero value otherwise.

### GetSubstringIndexEntryLimitOk

`func (o *AddLocalDbIndexRequest) GetSubstringIndexEntryLimitOk() (*int32, bool)`

GetSubstringIndexEntryLimitOk returns a tuple with the SubstringIndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstringIndexEntryLimit

`func (o *AddLocalDbIndexRequest) SetSubstringIndexEntryLimit(v int32)`

SetSubstringIndexEntryLimit sets SubstringIndexEntryLimit field to given value.

### HasSubstringIndexEntryLimit

`func (o *AddLocalDbIndexRequest) HasSubstringIndexEntryLimit() bool`

HasSubstringIndexEntryLimit returns a boolean if a field has been set.

### GetMaintainMatchCountForKeysExceedingEntryLimit

`func (o *AddLocalDbIndexRequest) GetMaintainMatchCountForKeysExceedingEntryLimit() bool`

GetMaintainMatchCountForKeysExceedingEntryLimit returns the MaintainMatchCountForKeysExceedingEntryLimit field if non-nil, zero value otherwise.

### GetMaintainMatchCountForKeysExceedingEntryLimitOk

`func (o *AddLocalDbIndexRequest) GetMaintainMatchCountForKeysExceedingEntryLimitOk() (*bool, bool)`

GetMaintainMatchCountForKeysExceedingEntryLimitOk returns a tuple with the MaintainMatchCountForKeysExceedingEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainMatchCountForKeysExceedingEntryLimit

`func (o *AddLocalDbIndexRequest) SetMaintainMatchCountForKeysExceedingEntryLimit(v bool)`

SetMaintainMatchCountForKeysExceedingEntryLimit sets MaintainMatchCountForKeysExceedingEntryLimit field to given value.

### HasMaintainMatchCountForKeysExceedingEntryLimit

`func (o *AddLocalDbIndexRequest) HasMaintainMatchCountForKeysExceedingEntryLimit() bool`

HasMaintainMatchCountForKeysExceedingEntryLimit returns a boolean if a field has been set.

### GetIndexType

`func (o *AddLocalDbIndexRequest) GetIndexType() []EnumlocalDbIndexIndexTypeProp`

GetIndexType returns the IndexType field if non-nil, zero value otherwise.

### GetIndexTypeOk

`func (o *AddLocalDbIndexRequest) GetIndexTypeOk() (*[]EnumlocalDbIndexIndexTypeProp, bool)`

GetIndexTypeOk returns a tuple with the IndexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexType

`func (o *AddLocalDbIndexRequest) SetIndexType(v []EnumlocalDbIndexIndexTypeProp)`

SetIndexType sets IndexType field to given value.


### GetSubstringLength

`func (o *AddLocalDbIndexRequest) GetSubstringLength() int32`

GetSubstringLength returns the SubstringLength field if non-nil, zero value otherwise.

### GetSubstringLengthOk

`func (o *AddLocalDbIndexRequest) GetSubstringLengthOk() (*int32, bool)`

GetSubstringLengthOk returns a tuple with the SubstringLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstringLength

`func (o *AddLocalDbIndexRequest) SetSubstringLength(v int32)`

SetSubstringLength sets SubstringLength field to given value.

### HasSubstringLength

`func (o *AddLocalDbIndexRequest) HasSubstringLength() bool`

HasSubstringLength returns a boolean if a field has been set.

### GetPrimeIndex

`func (o *AddLocalDbIndexRequest) GetPrimeIndex() bool`

GetPrimeIndex returns the PrimeIndex field if non-nil, zero value otherwise.

### GetPrimeIndexOk

`func (o *AddLocalDbIndexRequest) GetPrimeIndexOk() (*bool, bool)`

GetPrimeIndexOk returns a tuple with the PrimeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeIndex

`func (o *AddLocalDbIndexRequest) SetPrimeIndex(v bool)`

SetPrimeIndex sets PrimeIndex field to given value.

### HasPrimeIndex

`func (o *AddLocalDbIndexRequest) HasPrimeIndex() bool`

HasPrimeIndex returns a boolean if a field has been set.

### GetPrimeInternalNodesOnly

`func (o *AddLocalDbIndexRequest) GetPrimeInternalNodesOnly() bool`

GetPrimeInternalNodesOnly returns the PrimeInternalNodesOnly field if non-nil, zero value otherwise.

### GetPrimeInternalNodesOnlyOk

`func (o *AddLocalDbIndexRequest) GetPrimeInternalNodesOnlyOk() (*bool, bool)`

GetPrimeInternalNodesOnlyOk returns a tuple with the PrimeInternalNodesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeInternalNodesOnly

`func (o *AddLocalDbIndexRequest) SetPrimeInternalNodesOnly(v bool)`

SetPrimeInternalNodesOnly sets PrimeInternalNodesOnly field to given value.

### HasPrimeInternalNodesOnly

`func (o *AddLocalDbIndexRequest) HasPrimeInternalNodesOnly() bool`

HasPrimeInternalNodesOnly returns a boolean if a field has been set.

### GetEqualityIndexFilter

`func (o *AddLocalDbIndexRequest) GetEqualityIndexFilter() []string`

GetEqualityIndexFilter returns the EqualityIndexFilter field if non-nil, zero value otherwise.

### GetEqualityIndexFilterOk

`func (o *AddLocalDbIndexRequest) GetEqualityIndexFilterOk() (*[]string, bool)`

GetEqualityIndexFilterOk returns a tuple with the EqualityIndexFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqualityIndexFilter

`func (o *AddLocalDbIndexRequest) SetEqualityIndexFilter(v []string)`

SetEqualityIndexFilter sets EqualityIndexFilter field to given value.

### HasEqualityIndexFilter

`func (o *AddLocalDbIndexRequest) HasEqualityIndexFilter() bool`

HasEqualityIndexFilter returns a boolean if a field has been set.

### GetMaintainEqualityIndexWithoutFilter

`func (o *AddLocalDbIndexRequest) GetMaintainEqualityIndexWithoutFilter() bool`

GetMaintainEqualityIndexWithoutFilter returns the MaintainEqualityIndexWithoutFilter field if non-nil, zero value otherwise.

### GetMaintainEqualityIndexWithoutFilterOk

`func (o *AddLocalDbIndexRequest) GetMaintainEqualityIndexWithoutFilterOk() (*bool, bool)`

GetMaintainEqualityIndexWithoutFilterOk returns a tuple with the MaintainEqualityIndexWithoutFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainEqualityIndexWithoutFilter

`func (o *AddLocalDbIndexRequest) SetMaintainEqualityIndexWithoutFilter(v bool)`

SetMaintainEqualityIndexWithoutFilter sets MaintainEqualityIndexWithoutFilter field to given value.

### HasMaintainEqualityIndexWithoutFilter

`func (o *AddLocalDbIndexRequest) HasMaintainEqualityIndexWithoutFilter() bool`

HasMaintainEqualityIndexWithoutFilter returns a boolean if a field has been set.

### GetCacheMode

`func (o *AddLocalDbIndexRequest) GetCacheMode() EnumlocalDbIndexCacheModeProp`

GetCacheMode returns the CacheMode field if non-nil, zero value otherwise.

### GetCacheModeOk

`func (o *AddLocalDbIndexRequest) GetCacheModeOk() (*EnumlocalDbIndexCacheModeProp, bool)`

GetCacheModeOk returns a tuple with the CacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMode

`func (o *AddLocalDbIndexRequest) SetCacheMode(v EnumlocalDbIndexCacheModeProp)`

SetCacheMode sets CacheMode field to given value.

### HasCacheMode

`func (o *AddLocalDbIndexRequest) HasCacheMode() bool`

HasCacheMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


