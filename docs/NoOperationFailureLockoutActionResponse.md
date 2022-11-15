# NoOperationFailureLockoutActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Failure Lockout Action | 
**Schemas** | [**[]EnumnoOperationFailureLockoutActionSchemaUrn**](EnumnoOperationFailureLockoutActionSchemaUrn.md) |  | 
**GenerateAccountStatusNotification** | Pointer to **bool** | Indicates whether to generate an account status notification for cases in which this failure lockout action is invoked for a bind attempt with too many outstanding authentication failures. | [optional] 
**Description** | Pointer to **string** | A description for this Failure Lockout Action | [optional] 

## Methods

### NewNoOperationFailureLockoutActionResponse

`func NewNoOperationFailureLockoutActionResponse(id string, schemas []EnumnoOperationFailureLockoutActionSchemaUrn, ) *NoOperationFailureLockoutActionResponse`

NewNoOperationFailureLockoutActionResponse instantiates a new NoOperationFailureLockoutActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoOperationFailureLockoutActionResponseWithDefaults

`func NewNoOperationFailureLockoutActionResponseWithDefaults() *NoOperationFailureLockoutActionResponse`

NewNoOperationFailureLockoutActionResponseWithDefaults instantiates a new NoOperationFailureLockoutActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NoOperationFailureLockoutActionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NoOperationFailureLockoutActionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NoOperationFailureLockoutActionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *NoOperationFailureLockoutActionResponse) GetSchemas() []EnumnoOperationFailureLockoutActionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *NoOperationFailureLockoutActionResponse) GetSchemasOk() (*[]EnumnoOperationFailureLockoutActionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *NoOperationFailureLockoutActionResponse) SetSchemas(v []EnumnoOperationFailureLockoutActionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetGenerateAccountStatusNotification

`func (o *NoOperationFailureLockoutActionResponse) GetGenerateAccountStatusNotification() bool`

GetGenerateAccountStatusNotification returns the GenerateAccountStatusNotification field if non-nil, zero value otherwise.

### GetGenerateAccountStatusNotificationOk

`func (o *NoOperationFailureLockoutActionResponse) GetGenerateAccountStatusNotificationOk() (*bool, bool)`

GetGenerateAccountStatusNotificationOk returns a tuple with the GenerateAccountStatusNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateAccountStatusNotification

`func (o *NoOperationFailureLockoutActionResponse) SetGenerateAccountStatusNotification(v bool)`

SetGenerateAccountStatusNotification sets GenerateAccountStatusNotification field to given value.

### HasGenerateAccountStatusNotification

`func (o *NoOperationFailureLockoutActionResponse) HasGenerateAccountStatusNotification() bool`

HasGenerateAccountStatusNotification returns a boolean if a field has been set.

### GetDescription

`func (o *NoOperationFailureLockoutActionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NoOperationFailureLockoutActionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NoOperationFailureLockoutActionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NoOperationFailureLockoutActionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


