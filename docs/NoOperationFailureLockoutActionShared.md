# NoOperationFailureLockoutActionShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumnoOperationFailureLockoutActionSchemaUrn**](EnumnoOperationFailureLockoutActionSchemaUrn.md) |  | 
**GenerateAccountStatusNotification** | Pointer to **bool** | Indicates whether to generate an account status notification for cases in which this failure lockout action is invoked for a bind attempt with too many outstanding authentication failures. | [optional] 
**Description** | Pointer to **string** | A description for this Failure Lockout Action | [optional] 

## Methods

### NewNoOperationFailureLockoutActionShared

`func NewNoOperationFailureLockoutActionShared(schemas []EnumnoOperationFailureLockoutActionSchemaUrn, ) *NoOperationFailureLockoutActionShared`

NewNoOperationFailureLockoutActionShared instantiates a new NoOperationFailureLockoutActionShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoOperationFailureLockoutActionSharedWithDefaults

`func NewNoOperationFailureLockoutActionSharedWithDefaults() *NoOperationFailureLockoutActionShared`

NewNoOperationFailureLockoutActionSharedWithDefaults instantiates a new NoOperationFailureLockoutActionShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *NoOperationFailureLockoutActionShared) GetSchemas() []EnumnoOperationFailureLockoutActionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *NoOperationFailureLockoutActionShared) GetSchemasOk() (*[]EnumnoOperationFailureLockoutActionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *NoOperationFailureLockoutActionShared) SetSchemas(v []EnumnoOperationFailureLockoutActionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetGenerateAccountStatusNotification

`func (o *NoOperationFailureLockoutActionShared) GetGenerateAccountStatusNotification() bool`

GetGenerateAccountStatusNotification returns the GenerateAccountStatusNotification field if non-nil, zero value otherwise.

### GetGenerateAccountStatusNotificationOk

`func (o *NoOperationFailureLockoutActionShared) GetGenerateAccountStatusNotificationOk() (*bool, bool)`

GetGenerateAccountStatusNotificationOk returns a tuple with the GenerateAccountStatusNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateAccountStatusNotification

`func (o *NoOperationFailureLockoutActionShared) SetGenerateAccountStatusNotification(v bool)`

SetGenerateAccountStatusNotification sets GenerateAccountStatusNotification field to given value.

### HasGenerateAccountStatusNotification

`func (o *NoOperationFailureLockoutActionShared) HasGenerateAccountStatusNotification() bool`

HasGenerateAccountStatusNotification returns a boolean if a field has been set.

### GetDescription

`func (o *NoOperationFailureLockoutActionShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NoOperationFailureLockoutActionShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NoOperationFailureLockoutActionShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NoOperationFailureLockoutActionShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


