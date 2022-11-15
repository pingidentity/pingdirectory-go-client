# AddAggregateIdentityMapperRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MapperName** | **string** | Name of the new Identity Mapper | 
**Schemas** | [**[]EnumaggregateIdentityMapperSchemaUrn**](EnumaggregateIdentityMapperSchemaUrn.md) |  | 
**AllIncludedIdentityMapper** | Pointer to **[]string** |  | [optional] 
**AnyIncludedIdentityMapper** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Identity Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Identity Mapper is enabled for use. | 

## Methods

### NewAddAggregateIdentityMapperRequest

`func NewAddAggregateIdentityMapperRequest(mapperName string, schemas []EnumaggregateIdentityMapperSchemaUrn, enabled bool, ) *AddAggregateIdentityMapperRequest`

NewAddAggregateIdentityMapperRequest instantiates a new AddAggregateIdentityMapperRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAggregateIdentityMapperRequestWithDefaults

`func NewAddAggregateIdentityMapperRequestWithDefaults() *AddAggregateIdentityMapperRequest`

NewAddAggregateIdentityMapperRequestWithDefaults instantiates a new AddAggregateIdentityMapperRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMapperName

`func (o *AddAggregateIdentityMapperRequest) GetMapperName() string`

GetMapperName returns the MapperName field if non-nil, zero value otherwise.

### GetMapperNameOk

`func (o *AddAggregateIdentityMapperRequest) GetMapperNameOk() (*string, bool)`

GetMapperNameOk returns a tuple with the MapperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperName

`func (o *AddAggregateIdentityMapperRequest) SetMapperName(v string)`

SetMapperName sets MapperName field to given value.


### GetSchemas

`func (o *AddAggregateIdentityMapperRequest) GetSchemas() []EnumaggregateIdentityMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAggregateIdentityMapperRequest) GetSchemasOk() (*[]EnumaggregateIdentityMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAggregateIdentityMapperRequest) SetSchemas(v []EnumaggregateIdentityMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllIncludedIdentityMapper

`func (o *AddAggregateIdentityMapperRequest) GetAllIncludedIdentityMapper() []string`

GetAllIncludedIdentityMapper returns the AllIncludedIdentityMapper field if non-nil, zero value otherwise.

### GetAllIncludedIdentityMapperOk

`func (o *AddAggregateIdentityMapperRequest) GetAllIncludedIdentityMapperOk() (*[]string, bool)`

GetAllIncludedIdentityMapperOk returns a tuple with the AllIncludedIdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedIdentityMapper

`func (o *AddAggregateIdentityMapperRequest) SetAllIncludedIdentityMapper(v []string)`

SetAllIncludedIdentityMapper sets AllIncludedIdentityMapper field to given value.

### HasAllIncludedIdentityMapper

`func (o *AddAggregateIdentityMapperRequest) HasAllIncludedIdentityMapper() bool`

HasAllIncludedIdentityMapper returns a boolean if a field has been set.

### GetAnyIncludedIdentityMapper

`func (o *AddAggregateIdentityMapperRequest) GetAnyIncludedIdentityMapper() []string`

GetAnyIncludedIdentityMapper returns the AnyIncludedIdentityMapper field if non-nil, zero value otherwise.

### GetAnyIncludedIdentityMapperOk

`func (o *AddAggregateIdentityMapperRequest) GetAnyIncludedIdentityMapperOk() (*[]string, bool)`

GetAnyIncludedIdentityMapperOk returns a tuple with the AnyIncludedIdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedIdentityMapper

`func (o *AddAggregateIdentityMapperRequest) SetAnyIncludedIdentityMapper(v []string)`

SetAnyIncludedIdentityMapper sets AnyIncludedIdentityMapper field to given value.

### HasAnyIncludedIdentityMapper

`func (o *AddAggregateIdentityMapperRequest) HasAnyIncludedIdentityMapper() bool`

HasAnyIncludedIdentityMapper returns a boolean if a field has been set.

### GetDescription

`func (o *AddAggregateIdentityMapperRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAggregateIdentityMapperRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAggregateIdentityMapperRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAggregateIdentityMapperRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddAggregateIdentityMapperRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddAggregateIdentityMapperRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddAggregateIdentityMapperRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


