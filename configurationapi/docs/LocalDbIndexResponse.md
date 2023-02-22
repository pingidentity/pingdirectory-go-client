# LocalDbIndexResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Local DB Index | 
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

### NewLocalDbIndexResponse

`func NewLocalDbIndexResponse(id string, attribute string, indexType []EnumlocalDbIndexIndexTypeProp, ) *LocalDbIndexResponse`

NewLocalDbIndexResponse instantiates a new LocalDbIndexResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalDbIndexResponseWithDefaults

`func NewLocalDbIndexResponseWithDefaults() *LocalDbIndexResponse`

NewLocalDbIndexResponseWithDefaults instantiates a new LocalDbIndexResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *LocalDbIndexResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LocalDbIndexResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LocalDbIndexResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LocalDbIndexResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *LocalDbIndexResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *LocalDbIndexResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *LocalDbIndexResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *LocalDbIndexResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *LocalDbIndexResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalDbIndexResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalDbIndexResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *LocalDbIndexResponse) GetSchemas() []EnumlocalDbIndexSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LocalDbIndexResponse) GetSchemasOk() (*[]EnumlocalDbIndexSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LocalDbIndexResponse) SetSchemas(v []EnumlocalDbIndexSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *LocalDbIndexResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetAttribute

`func (o *LocalDbIndexResponse) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *LocalDbIndexResponse) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *LocalDbIndexResponse) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetIndexEntryLimit

`func (o *LocalDbIndexResponse) GetIndexEntryLimit() int32`

GetIndexEntryLimit returns the IndexEntryLimit field if non-nil, zero value otherwise.

### GetIndexEntryLimitOk

`func (o *LocalDbIndexResponse) GetIndexEntryLimitOk() (*int32, bool)`

GetIndexEntryLimitOk returns a tuple with the IndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexEntryLimit

`func (o *LocalDbIndexResponse) SetIndexEntryLimit(v int32)`

SetIndexEntryLimit sets IndexEntryLimit field to given value.

### HasIndexEntryLimit

`func (o *LocalDbIndexResponse) HasIndexEntryLimit() bool`

HasIndexEntryLimit returns a boolean if a field has been set.

### GetSubstringIndexEntryLimit

`func (o *LocalDbIndexResponse) GetSubstringIndexEntryLimit() int32`

GetSubstringIndexEntryLimit returns the SubstringIndexEntryLimit field if non-nil, zero value otherwise.

### GetSubstringIndexEntryLimitOk

`func (o *LocalDbIndexResponse) GetSubstringIndexEntryLimitOk() (*int32, bool)`

GetSubstringIndexEntryLimitOk returns a tuple with the SubstringIndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstringIndexEntryLimit

`func (o *LocalDbIndexResponse) SetSubstringIndexEntryLimit(v int32)`

SetSubstringIndexEntryLimit sets SubstringIndexEntryLimit field to given value.

### HasSubstringIndexEntryLimit

`func (o *LocalDbIndexResponse) HasSubstringIndexEntryLimit() bool`

HasSubstringIndexEntryLimit returns a boolean if a field has been set.

### GetMaintainMatchCountForKeysExceedingEntryLimit

`func (o *LocalDbIndexResponse) GetMaintainMatchCountForKeysExceedingEntryLimit() bool`

GetMaintainMatchCountForKeysExceedingEntryLimit returns the MaintainMatchCountForKeysExceedingEntryLimit field if non-nil, zero value otherwise.

### GetMaintainMatchCountForKeysExceedingEntryLimitOk

`func (o *LocalDbIndexResponse) GetMaintainMatchCountForKeysExceedingEntryLimitOk() (*bool, bool)`

GetMaintainMatchCountForKeysExceedingEntryLimitOk returns a tuple with the MaintainMatchCountForKeysExceedingEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainMatchCountForKeysExceedingEntryLimit

`func (o *LocalDbIndexResponse) SetMaintainMatchCountForKeysExceedingEntryLimit(v bool)`

SetMaintainMatchCountForKeysExceedingEntryLimit sets MaintainMatchCountForKeysExceedingEntryLimit field to given value.

### HasMaintainMatchCountForKeysExceedingEntryLimit

`func (o *LocalDbIndexResponse) HasMaintainMatchCountForKeysExceedingEntryLimit() bool`

HasMaintainMatchCountForKeysExceedingEntryLimit returns a boolean if a field has been set.

### GetIndexType

`func (o *LocalDbIndexResponse) GetIndexType() []EnumlocalDbIndexIndexTypeProp`

GetIndexType returns the IndexType field if non-nil, zero value otherwise.

### GetIndexTypeOk

`func (o *LocalDbIndexResponse) GetIndexTypeOk() (*[]EnumlocalDbIndexIndexTypeProp, bool)`

GetIndexTypeOk returns a tuple with the IndexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexType

`func (o *LocalDbIndexResponse) SetIndexType(v []EnumlocalDbIndexIndexTypeProp)`

SetIndexType sets IndexType field to given value.


### GetSubstringLength

`func (o *LocalDbIndexResponse) GetSubstringLength() int32`

GetSubstringLength returns the SubstringLength field if non-nil, zero value otherwise.

### GetSubstringLengthOk

`func (o *LocalDbIndexResponse) GetSubstringLengthOk() (*int32, bool)`

GetSubstringLengthOk returns a tuple with the SubstringLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstringLength

`func (o *LocalDbIndexResponse) SetSubstringLength(v int32)`

SetSubstringLength sets SubstringLength field to given value.

### HasSubstringLength

`func (o *LocalDbIndexResponse) HasSubstringLength() bool`

HasSubstringLength returns a boolean if a field has been set.

### GetPrimeIndex

`func (o *LocalDbIndexResponse) GetPrimeIndex() bool`

GetPrimeIndex returns the PrimeIndex field if non-nil, zero value otherwise.

### GetPrimeIndexOk

`func (o *LocalDbIndexResponse) GetPrimeIndexOk() (*bool, bool)`

GetPrimeIndexOk returns a tuple with the PrimeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeIndex

`func (o *LocalDbIndexResponse) SetPrimeIndex(v bool)`

SetPrimeIndex sets PrimeIndex field to given value.

### HasPrimeIndex

`func (o *LocalDbIndexResponse) HasPrimeIndex() bool`

HasPrimeIndex returns a boolean if a field has been set.

### GetPrimeInternalNodesOnly

`func (o *LocalDbIndexResponse) GetPrimeInternalNodesOnly() bool`

GetPrimeInternalNodesOnly returns the PrimeInternalNodesOnly field if non-nil, zero value otherwise.

### GetPrimeInternalNodesOnlyOk

`func (o *LocalDbIndexResponse) GetPrimeInternalNodesOnlyOk() (*bool, bool)`

GetPrimeInternalNodesOnlyOk returns a tuple with the PrimeInternalNodesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeInternalNodesOnly

`func (o *LocalDbIndexResponse) SetPrimeInternalNodesOnly(v bool)`

SetPrimeInternalNodesOnly sets PrimeInternalNodesOnly field to given value.

### HasPrimeInternalNodesOnly

`func (o *LocalDbIndexResponse) HasPrimeInternalNodesOnly() bool`

HasPrimeInternalNodesOnly returns a boolean if a field has been set.

### GetEqualityIndexFilter

`func (o *LocalDbIndexResponse) GetEqualityIndexFilter() []string`

GetEqualityIndexFilter returns the EqualityIndexFilter field if non-nil, zero value otherwise.

### GetEqualityIndexFilterOk

`func (o *LocalDbIndexResponse) GetEqualityIndexFilterOk() (*[]string, bool)`

GetEqualityIndexFilterOk returns a tuple with the EqualityIndexFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqualityIndexFilter

`func (o *LocalDbIndexResponse) SetEqualityIndexFilter(v []string)`

SetEqualityIndexFilter sets EqualityIndexFilter field to given value.

### HasEqualityIndexFilter

`func (o *LocalDbIndexResponse) HasEqualityIndexFilter() bool`

HasEqualityIndexFilter returns a boolean if a field has been set.

### GetMaintainEqualityIndexWithoutFilter

`func (o *LocalDbIndexResponse) GetMaintainEqualityIndexWithoutFilter() bool`

GetMaintainEqualityIndexWithoutFilter returns the MaintainEqualityIndexWithoutFilter field if non-nil, zero value otherwise.

### GetMaintainEqualityIndexWithoutFilterOk

`func (o *LocalDbIndexResponse) GetMaintainEqualityIndexWithoutFilterOk() (*bool, bool)`

GetMaintainEqualityIndexWithoutFilterOk returns a tuple with the MaintainEqualityIndexWithoutFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainEqualityIndexWithoutFilter

`func (o *LocalDbIndexResponse) SetMaintainEqualityIndexWithoutFilter(v bool)`

SetMaintainEqualityIndexWithoutFilter sets MaintainEqualityIndexWithoutFilter field to given value.

### HasMaintainEqualityIndexWithoutFilter

`func (o *LocalDbIndexResponse) HasMaintainEqualityIndexWithoutFilter() bool`

HasMaintainEqualityIndexWithoutFilter returns a boolean if a field has been set.

### GetCacheMode

`func (o *LocalDbIndexResponse) GetCacheMode() EnumlocalDbIndexCacheModeProp`

GetCacheMode returns the CacheMode field if non-nil, zero value otherwise.

### GetCacheModeOk

`func (o *LocalDbIndexResponse) GetCacheModeOk() (*EnumlocalDbIndexCacheModeProp, bool)`

GetCacheModeOk returns a tuple with the CacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMode

`func (o *LocalDbIndexResponse) SetCacheMode(v EnumlocalDbIndexCacheModeProp)`

SetCacheMode sets CacheMode field to given value.

### HasCacheMode

`func (o *LocalDbIndexResponse) HasCacheMode() bool`

HasCacheMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


