# AggregateIdentityMapperResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Identity Mapper | 
**Schemas** | [**[]EnumaggregateIdentityMapperSchemaUrn**](EnumaggregateIdentityMapperSchemaUrn.md) |  | 
**AllIncludedIdentityMapper** | Pointer to **[]string** | The set of identity mappers that must all match the target entry. Each identity mapper must uniquely match the same target entry. If any of the identity mappers match multiple entries, if any of them match zero entries, or if any of them match different entries, then the mapping will fail. | [optional] 
**AnyIncludedIdentityMapper** | Pointer to **[]string** | The set of identity mappers that will be used to identify the target entry. At least one identity mapper must uniquely match an entry. If multiple identity mappers match entries, then they must all uniquely match the same entry. If none of the identity mappers match any entries, if any of them match multiple entries, or if any of them match different entries, then the mapping will fail. | [optional] 
**Description** | Pointer to **string** | A description for this Identity Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Identity Mapper is enabled for use. | 

## Methods

### NewAggregateIdentityMapperResponse

`func NewAggregateIdentityMapperResponse(id string, schemas []EnumaggregateIdentityMapperSchemaUrn, enabled bool, ) *AggregateIdentityMapperResponse`

NewAggregateIdentityMapperResponse instantiates a new AggregateIdentityMapperResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateIdentityMapperResponseWithDefaults

`func NewAggregateIdentityMapperResponseWithDefaults() *AggregateIdentityMapperResponse`

NewAggregateIdentityMapperResponseWithDefaults instantiates a new AggregateIdentityMapperResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *AggregateIdentityMapperResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AggregateIdentityMapperResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AggregateIdentityMapperResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AggregateIdentityMapperResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AggregateIdentityMapperResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AggregateIdentityMapperResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AggregateIdentityMapperResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AggregateIdentityMapperResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *AggregateIdentityMapperResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AggregateIdentityMapperResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AggregateIdentityMapperResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AggregateIdentityMapperResponse) GetSchemas() []EnumaggregateIdentityMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AggregateIdentityMapperResponse) GetSchemasOk() (*[]EnumaggregateIdentityMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AggregateIdentityMapperResponse) SetSchemas(v []EnumaggregateIdentityMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllIncludedIdentityMapper

`func (o *AggregateIdentityMapperResponse) GetAllIncludedIdentityMapper() []string`

GetAllIncludedIdentityMapper returns the AllIncludedIdentityMapper field if non-nil, zero value otherwise.

### GetAllIncludedIdentityMapperOk

`func (o *AggregateIdentityMapperResponse) GetAllIncludedIdentityMapperOk() (*[]string, bool)`

GetAllIncludedIdentityMapperOk returns a tuple with the AllIncludedIdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedIdentityMapper

`func (o *AggregateIdentityMapperResponse) SetAllIncludedIdentityMapper(v []string)`

SetAllIncludedIdentityMapper sets AllIncludedIdentityMapper field to given value.

### HasAllIncludedIdentityMapper

`func (o *AggregateIdentityMapperResponse) HasAllIncludedIdentityMapper() bool`

HasAllIncludedIdentityMapper returns a boolean if a field has been set.

### GetAnyIncludedIdentityMapper

`func (o *AggregateIdentityMapperResponse) GetAnyIncludedIdentityMapper() []string`

GetAnyIncludedIdentityMapper returns the AnyIncludedIdentityMapper field if non-nil, zero value otherwise.

### GetAnyIncludedIdentityMapperOk

`func (o *AggregateIdentityMapperResponse) GetAnyIncludedIdentityMapperOk() (*[]string, bool)`

GetAnyIncludedIdentityMapperOk returns a tuple with the AnyIncludedIdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedIdentityMapper

`func (o *AggregateIdentityMapperResponse) SetAnyIncludedIdentityMapper(v []string)`

SetAnyIncludedIdentityMapper sets AnyIncludedIdentityMapper field to given value.

### HasAnyIncludedIdentityMapper

`func (o *AggregateIdentityMapperResponse) HasAnyIncludedIdentityMapper() bool`

HasAnyIncludedIdentityMapper returns a boolean if a field has been set.

### GetDescription

`func (o *AggregateIdentityMapperResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AggregateIdentityMapperResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AggregateIdentityMapperResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AggregateIdentityMapperResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AggregateIdentityMapperResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AggregateIdentityMapperResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AggregateIdentityMapperResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


