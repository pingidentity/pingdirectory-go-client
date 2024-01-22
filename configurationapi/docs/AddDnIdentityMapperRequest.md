# AddDnIdentityMapperRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumdnIdentityMapperSchemaUrn**](EnumdnIdentityMapperSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Identity Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Identity Mapper is enabled for use. | 
**MapperName** | **string** | Name of the new Identity Mapper | 

## Methods

### NewAddDnIdentityMapperRequest

`func NewAddDnIdentityMapperRequest(schemas []EnumdnIdentityMapperSchemaUrn, enabled bool, mapperName string, ) *AddDnIdentityMapperRequest`

NewAddDnIdentityMapperRequest instantiates a new AddDnIdentityMapperRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDnIdentityMapperRequestWithDefaults

`func NewAddDnIdentityMapperRequestWithDefaults() *AddDnIdentityMapperRequest`

NewAddDnIdentityMapperRequestWithDefaults instantiates a new AddDnIdentityMapperRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddDnIdentityMapperRequest) GetSchemas() []EnumdnIdentityMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDnIdentityMapperRequest) GetSchemasOk() (*[]EnumdnIdentityMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDnIdentityMapperRequest) SetSchemas(v []EnumdnIdentityMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddDnIdentityMapperRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddDnIdentityMapperRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddDnIdentityMapperRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddDnIdentityMapperRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddDnIdentityMapperRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddDnIdentityMapperRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddDnIdentityMapperRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMapperName

`func (o *AddDnIdentityMapperRequest) GetMapperName() string`

GetMapperName returns the MapperName field if non-nil, zero value otherwise.

### GetMapperNameOk

`func (o *AddDnIdentityMapperRequest) GetMapperNameOk() (*string, bool)`

GetMapperNameOk returns a tuple with the MapperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperName

`func (o *AddDnIdentityMapperRequest) SetMapperName(v string)`

SetMapperName sets MapperName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


