# LocalDbVlvIndexResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Backend | 
**Schemas** | Pointer to [**[]EnumlocalDbVlvIndexSchemaUrn**](EnumlocalDbVlvIndexSchemaUrn.md) |  | [optional] 
**BaseDN** | **string** | Specifies the base DN used in the search query that is being indexed. | 
**Scope** | [**EnumlocalDbVlvIndexScopeProp**](EnumlocalDbVlvIndexScopeProp.md) |  | 
**Filter** | **string** | Specifies the LDAP filter used in the query that is being indexed. | 
**SortOrder** | **string** | Specifies the names of the attributes that are used to sort the entries for the query being indexed. | 
**Name** | **string** | Specifies a unique name for this VLV index. | 
**MaxBlockSize** | Pointer to **int64** | Specifies the number of entry IDs to store in a single sorted set before it must be split. | [optional] 
**CacheMode** | Pointer to [**EnumlocalDbVlvIndexCacheModeProp**](EnumlocalDbVlvIndexCacheModeProp.md) |  | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewLocalDbVlvIndexResponse

`func NewLocalDbVlvIndexResponse(id string, baseDN string, scope EnumlocalDbVlvIndexScopeProp, filter string, sortOrder string, name string, ) *LocalDbVlvIndexResponse`

NewLocalDbVlvIndexResponse instantiates a new LocalDbVlvIndexResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalDbVlvIndexResponseWithDefaults

`func NewLocalDbVlvIndexResponseWithDefaults() *LocalDbVlvIndexResponse`

NewLocalDbVlvIndexResponseWithDefaults instantiates a new LocalDbVlvIndexResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocalDbVlvIndexResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalDbVlvIndexResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalDbVlvIndexResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *LocalDbVlvIndexResponse) GetSchemas() []EnumlocalDbVlvIndexSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LocalDbVlvIndexResponse) GetSchemasOk() (*[]EnumlocalDbVlvIndexSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LocalDbVlvIndexResponse) SetSchemas(v []EnumlocalDbVlvIndexSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *LocalDbVlvIndexResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetBaseDN

`func (o *LocalDbVlvIndexResponse) GetBaseDN() string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *LocalDbVlvIndexResponse) GetBaseDNOk() (*string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *LocalDbVlvIndexResponse) SetBaseDN(v string)`

SetBaseDN sets BaseDN field to given value.


### GetScope

`func (o *LocalDbVlvIndexResponse) GetScope() EnumlocalDbVlvIndexScopeProp`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *LocalDbVlvIndexResponse) GetScopeOk() (*EnumlocalDbVlvIndexScopeProp, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *LocalDbVlvIndexResponse) SetScope(v EnumlocalDbVlvIndexScopeProp)`

SetScope sets Scope field to given value.


### GetFilter

`func (o *LocalDbVlvIndexResponse) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LocalDbVlvIndexResponse) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LocalDbVlvIndexResponse) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetSortOrder

`func (o *LocalDbVlvIndexResponse) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *LocalDbVlvIndexResponse) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *LocalDbVlvIndexResponse) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.


### GetName

`func (o *LocalDbVlvIndexResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocalDbVlvIndexResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocalDbVlvIndexResponse) SetName(v string)`

SetName sets Name field to given value.


### GetMaxBlockSize

`func (o *LocalDbVlvIndexResponse) GetMaxBlockSize() int64`

GetMaxBlockSize returns the MaxBlockSize field if non-nil, zero value otherwise.

### GetMaxBlockSizeOk

`func (o *LocalDbVlvIndexResponse) GetMaxBlockSizeOk() (*int64, bool)`

GetMaxBlockSizeOk returns a tuple with the MaxBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBlockSize

`func (o *LocalDbVlvIndexResponse) SetMaxBlockSize(v int64)`

SetMaxBlockSize sets MaxBlockSize field to given value.

### HasMaxBlockSize

`func (o *LocalDbVlvIndexResponse) HasMaxBlockSize() bool`

HasMaxBlockSize returns a boolean if a field has been set.

### GetCacheMode

`func (o *LocalDbVlvIndexResponse) GetCacheMode() EnumlocalDbVlvIndexCacheModeProp`

GetCacheMode returns the CacheMode field if non-nil, zero value otherwise.

### GetCacheModeOk

`func (o *LocalDbVlvIndexResponse) GetCacheModeOk() (*EnumlocalDbVlvIndexCacheModeProp, bool)`

GetCacheModeOk returns a tuple with the CacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMode

`func (o *LocalDbVlvIndexResponse) SetCacheMode(v EnumlocalDbVlvIndexCacheModeProp)`

SetCacheMode sets CacheMode field to given value.

### HasCacheMode

`func (o *LocalDbVlvIndexResponse) HasCacheMode() bool`

HasCacheMode returns a boolean if a field has been set.

### GetMeta

`func (o *LocalDbVlvIndexResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LocalDbVlvIndexResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LocalDbVlvIndexResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LocalDbVlvIndexResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *LocalDbVlvIndexResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *LocalDbVlvIndexResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *LocalDbVlvIndexResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *LocalDbVlvIndexResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


