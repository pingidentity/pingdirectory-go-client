# AggregateIdentityMapperShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumaggregateIdentityMapperSchemaUrn**](EnumaggregateIdentityMapperSchemaUrn.md) |  | 
**AllIncludedIdentityMapper** | Pointer to **[]string** |  | [optional] 
**AnyIncludedIdentityMapper** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Identity Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Identity Mapper is enabled for use. | 

## Methods

### NewAggregateIdentityMapperShared

`func NewAggregateIdentityMapperShared(schemas []EnumaggregateIdentityMapperSchemaUrn, enabled bool, ) *AggregateIdentityMapperShared`

NewAggregateIdentityMapperShared instantiates a new AggregateIdentityMapperShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateIdentityMapperSharedWithDefaults

`func NewAggregateIdentityMapperSharedWithDefaults() *AggregateIdentityMapperShared`

NewAggregateIdentityMapperSharedWithDefaults instantiates a new AggregateIdentityMapperShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AggregateIdentityMapperShared) GetSchemas() []EnumaggregateIdentityMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AggregateIdentityMapperShared) GetSchemasOk() (*[]EnumaggregateIdentityMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AggregateIdentityMapperShared) SetSchemas(v []EnumaggregateIdentityMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllIncludedIdentityMapper

`func (o *AggregateIdentityMapperShared) GetAllIncludedIdentityMapper() []string`

GetAllIncludedIdentityMapper returns the AllIncludedIdentityMapper field if non-nil, zero value otherwise.

### GetAllIncludedIdentityMapperOk

`func (o *AggregateIdentityMapperShared) GetAllIncludedIdentityMapperOk() (*[]string, bool)`

GetAllIncludedIdentityMapperOk returns a tuple with the AllIncludedIdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedIdentityMapper

`func (o *AggregateIdentityMapperShared) SetAllIncludedIdentityMapper(v []string)`

SetAllIncludedIdentityMapper sets AllIncludedIdentityMapper field to given value.

### HasAllIncludedIdentityMapper

`func (o *AggregateIdentityMapperShared) HasAllIncludedIdentityMapper() bool`

HasAllIncludedIdentityMapper returns a boolean if a field has been set.

### GetAnyIncludedIdentityMapper

`func (o *AggregateIdentityMapperShared) GetAnyIncludedIdentityMapper() []string`

GetAnyIncludedIdentityMapper returns the AnyIncludedIdentityMapper field if non-nil, zero value otherwise.

### GetAnyIncludedIdentityMapperOk

`func (o *AggregateIdentityMapperShared) GetAnyIncludedIdentityMapperOk() (*[]string, bool)`

GetAnyIncludedIdentityMapperOk returns a tuple with the AnyIncludedIdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedIdentityMapper

`func (o *AggregateIdentityMapperShared) SetAnyIncludedIdentityMapper(v []string)`

SetAnyIncludedIdentityMapper sets AnyIncludedIdentityMapper field to given value.

### HasAnyIncludedIdentityMapper

`func (o *AggregateIdentityMapperShared) HasAnyIncludedIdentityMapper() bool`

HasAnyIncludedIdentityMapper returns a boolean if a field has been set.

### GetDescription

`func (o *AggregateIdentityMapperShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AggregateIdentityMapperShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AggregateIdentityMapperShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AggregateIdentityMapperShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AggregateIdentityMapperShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AggregateIdentityMapperShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AggregateIdentityMapperShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


