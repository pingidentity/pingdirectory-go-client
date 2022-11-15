# PasswordModifyExtendedOperationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumpasswordModifyExtendedOperationHandlerSchemaUrn**](EnumpasswordModifyExtendedOperationHandlerSchemaUrn.md) |  | 
**IdentityMapper** | **string** | Specifies the name of the identity mapper that should be used in conjunction with the password modify extended operation. | 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewPasswordModifyExtendedOperationHandlerResponse

`func NewPasswordModifyExtendedOperationHandlerResponse(schemas []EnumpasswordModifyExtendedOperationHandlerSchemaUrn, identityMapper string, enabled bool, ) *PasswordModifyExtendedOperationHandlerResponse`

NewPasswordModifyExtendedOperationHandlerResponse instantiates a new PasswordModifyExtendedOperationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordModifyExtendedOperationHandlerResponseWithDefaults

`func NewPasswordModifyExtendedOperationHandlerResponseWithDefaults() *PasswordModifyExtendedOperationHandlerResponse`

NewPasswordModifyExtendedOperationHandlerResponseWithDefaults instantiates a new PasswordModifyExtendedOperationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *PasswordModifyExtendedOperationHandlerResponse) GetSchemas() []EnumpasswordModifyExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PasswordModifyExtendedOperationHandlerResponse) GetSchemasOk() (*[]EnumpasswordModifyExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PasswordModifyExtendedOperationHandlerResponse) SetSchemas(v []EnumpasswordModifyExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIdentityMapper

`func (o *PasswordModifyExtendedOperationHandlerResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *PasswordModifyExtendedOperationHandlerResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *PasswordModifyExtendedOperationHandlerResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetDescription

`func (o *PasswordModifyExtendedOperationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PasswordModifyExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PasswordModifyExtendedOperationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PasswordModifyExtendedOperationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PasswordModifyExtendedOperationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PasswordModifyExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PasswordModifyExtendedOperationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


