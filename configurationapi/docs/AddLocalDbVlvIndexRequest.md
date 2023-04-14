# AddLocalDbVlvIndexRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexName** | **string** | Name of the new Local DB VLV Index | 
**Schemas** | Pointer to [**[]EnumlocalDbVlvIndexSchemaUrn**](EnumlocalDbVlvIndexSchemaUrn.md) |  | [optional] 
**BaseDN** | **string** | Specifies the base DN used in the search query that is being indexed. | 
**Scope** | [**EnumlocalDbVlvIndexScopeProp**](EnumlocalDbVlvIndexScopeProp.md) |  | 
**Filter** | **string** | Specifies the LDAP filter used in the query that is being indexed. | 
**SortOrder** | **string** | Specifies the names of the attributes that are used to sort the entries for the query being indexed. | 
**Name** | **string** | Specifies a unique name for this VLV index. | 
**MaxBlockSize** | Pointer to **int64** | Specifies the number of entry IDs to store in a single sorted set before it must be split. | [optional] 
**CacheMode** | Pointer to [**EnumlocalDbVlvIndexCacheModeProp**](EnumlocalDbVlvIndexCacheModeProp.md) |  | [optional] 

## Methods

### NewAddLocalDbVlvIndexRequest

`func NewAddLocalDbVlvIndexRequest(indexName string, baseDN string, scope EnumlocalDbVlvIndexScopeProp, filter string, sortOrder string, name string, ) *AddLocalDbVlvIndexRequest`

NewAddLocalDbVlvIndexRequest instantiates a new AddLocalDbVlvIndexRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLocalDbVlvIndexRequestWithDefaults

`func NewAddLocalDbVlvIndexRequestWithDefaults() *AddLocalDbVlvIndexRequest`

NewAddLocalDbVlvIndexRequestWithDefaults instantiates a new AddLocalDbVlvIndexRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexName

`func (o *AddLocalDbVlvIndexRequest) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *AddLocalDbVlvIndexRequest) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *AddLocalDbVlvIndexRequest) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.


### GetSchemas

`func (o *AddLocalDbVlvIndexRequest) GetSchemas() []EnumlocalDbVlvIndexSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLocalDbVlvIndexRequest) GetSchemasOk() (*[]EnumlocalDbVlvIndexSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLocalDbVlvIndexRequest) SetSchemas(v []EnumlocalDbVlvIndexSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddLocalDbVlvIndexRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetBaseDN

`func (o *AddLocalDbVlvIndexRequest) GetBaseDN() string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddLocalDbVlvIndexRequest) GetBaseDNOk() (*string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddLocalDbVlvIndexRequest) SetBaseDN(v string)`

SetBaseDN sets BaseDN field to given value.


### GetScope

`func (o *AddLocalDbVlvIndexRequest) GetScope() EnumlocalDbVlvIndexScopeProp`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AddLocalDbVlvIndexRequest) GetScopeOk() (*EnumlocalDbVlvIndexScopeProp, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AddLocalDbVlvIndexRequest) SetScope(v EnumlocalDbVlvIndexScopeProp)`

SetScope sets Scope field to given value.


### GetFilter

`func (o *AddLocalDbVlvIndexRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddLocalDbVlvIndexRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddLocalDbVlvIndexRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetSortOrder

`func (o *AddLocalDbVlvIndexRequest) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AddLocalDbVlvIndexRequest) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AddLocalDbVlvIndexRequest) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.


### GetName

`func (o *AddLocalDbVlvIndexRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddLocalDbVlvIndexRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddLocalDbVlvIndexRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMaxBlockSize

`func (o *AddLocalDbVlvIndexRequest) GetMaxBlockSize() int64`

GetMaxBlockSize returns the MaxBlockSize field if non-nil, zero value otherwise.

### GetMaxBlockSizeOk

`func (o *AddLocalDbVlvIndexRequest) GetMaxBlockSizeOk() (*int64, bool)`

GetMaxBlockSizeOk returns a tuple with the MaxBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBlockSize

`func (o *AddLocalDbVlvIndexRequest) SetMaxBlockSize(v int64)`

SetMaxBlockSize sets MaxBlockSize field to given value.

### HasMaxBlockSize

`func (o *AddLocalDbVlvIndexRequest) HasMaxBlockSize() bool`

HasMaxBlockSize returns a boolean if a field has been set.

### GetCacheMode

`func (o *AddLocalDbVlvIndexRequest) GetCacheMode() EnumlocalDbVlvIndexCacheModeProp`

GetCacheMode returns the CacheMode field if non-nil, zero value otherwise.

### GetCacheModeOk

`func (o *AddLocalDbVlvIndexRequest) GetCacheModeOk() (*EnumlocalDbVlvIndexCacheModeProp, bool)`

GetCacheModeOk returns a tuple with the CacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMode

`func (o *AddLocalDbVlvIndexRequest) SetCacheMode(v EnumlocalDbVlvIndexCacheModeProp)`

SetCacheMode sets CacheMode field to given value.

### HasCacheMode

`func (o *AddLocalDbVlvIndexRequest) HasCacheMode() bool`

HasCacheMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


