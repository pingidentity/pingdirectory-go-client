# AddLockAccountFailureLockoutActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumlockAccountFailureLockoutActionSchemaUrn**](EnumlockAccountFailureLockoutActionSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Failure Lockout Action | [optional] 
**ActionName** | **string** | Name of the new Failure Lockout Action | 

## Methods

### NewAddLockAccountFailureLockoutActionRequest

`func NewAddLockAccountFailureLockoutActionRequest(schemas []EnumlockAccountFailureLockoutActionSchemaUrn, actionName string, ) *AddLockAccountFailureLockoutActionRequest`

NewAddLockAccountFailureLockoutActionRequest instantiates a new AddLockAccountFailureLockoutActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLockAccountFailureLockoutActionRequestWithDefaults

`func NewAddLockAccountFailureLockoutActionRequestWithDefaults() *AddLockAccountFailureLockoutActionRequest`

NewAddLockAccountFailureLockoutActionRequestWithDefaults instantiates a new AddLockAccountFailureLockoutActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddLockAccountFailureLockoutActionRequest) GetSchemas() []EnumlockAccountFailureLockoutActionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLockAccountFailureLockoutActionRequest) GetSchemasOk() (*[]EnumlockAccountFailureLockoutActionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLockAccountFailureLockoutActionRequest) SetSchemas(v []EnumlockAccountFailureLockoutActionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddLockAccountFailureLockoutActionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLockAccountFailureLockoutActionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLockAccountFailureLockoutActionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLockAccountFailureLockoutActionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionName

`func (o *AddLockAccountFailureLockoutActionRequest) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *AddLockAccountFailureLockoutActionRequest) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *AddLockAccountFailureLockoutActionRequest) SetActionName(v string)`

SetActionName sets ActionName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


