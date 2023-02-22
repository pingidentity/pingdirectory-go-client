# AddDelayBindResponseFailureLockoutActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionName** | **string** | Name of the new Failure Lockout Action | 
**Schemas** | [**[]EnumdelayBindResponseFailureLockoutActionSchemaUrn**](EnumdelayBindResponseFailureLockoutActionSchemaUrn.md) |  | 
**Delay** | **string** | The length of time to delay the bind response for accounts with too many failed authentication attempts. | 
**AllowBlockingDelay** | Pointer to **bool** | Indicates whether to delay the response for authentication attempts even if that delay may block the thread being used to process the attempt. | [optional] 
**GenerateAccountStatusNotification** | Pointer to **bool** | Indicates whether to generate an account status notification for cases in which a bind response is delayed because of failure lockout. | [optional] 
**Description** | Pointer to **string** | A description for this Failure Lockout Action | [optional] 

## Methods

### NewAddDelayBindResponseFailureLockoutActionRequest

`func NewAddDelayBindResponseFailureLockoutActionRequest(actionName string, schemas []EnumdelayBindResponseFailureLockoutActionSchemaUrn, delay string, ) *AddDelayBindResponseFailureLockoutActionRequest`

NewAddDelayBindResponseFailureLockoutActionRequest instantiates a new AddDelayBindResponseFailureLockoutActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDelayBindResponseFailureLockoutActionRequestWithDefaults

`func NewAddDelayBindResponseFailureLockoutActionRequestWithDefaults() *AddDelayBindResponseFailureLockoutActionRequest`

NewAddDelayBindResponseFailureLockoutActionRequestWithDefaults instantiates a new AddDelayBindResponseFailureLockoutActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionName

`func (o *AddDelayBindResponseFailureLockoutActionRequest) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *AddDelayBindResponseFailureLockoutActionRequest) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *AddDelayBindResponseFailureLockoutActionRequest) SetActionName(v string)`

SetActionName sets ActionName field to given value.


### GetSchemas

`func (o *AddDelayBindResponseFailureLockoutActionRequest) GetSchemas() []EnumdelayBindResponseFailureLockoutActionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDelayBindResponseFailureLockoutActionRequest) GetSchemasOk() (*[]EnumdelayBindResponseFailureLockoutActionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDelayBindResponseFailureLockoutActionRequest) SetSchemas(v []EnumdelayBindResponseFailureLockoutActionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDelay

`func (o *AddDelayBindResponseFailureLockoutActionRequest) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *AddDelayBindResponseFailureLockoutActionRequest) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *AddDelayBindResponseFailureLockoutActionRequest) SetDelay(v string)`

SetDelay sets Delay field to given value.


### GetAllowBlockingDelay

`func (o *AddDelayBindResponseFailureLockoutActionRequest) GetAllowBlockingDelay() bool`

GetAllowBlockingDelay returns the AllowBlockingDelay field if non-nil, zero value otherwise.

### GetAllowBlockingDelayOk

`func (o *AddDelayBindResponseFailureLockoutActionRequest) GetAllowBlockingDelayOk() (*bool, bool)`

GetAllowBlockingDelayOk returns a tuple with the AllowBlockingDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBlockingDelay

`func (o *AddDelayBindResponseFailureLockoutActionRequest) SetAllowBlockingDelay(v bool)`

SetAllowBlockingDelay sets AllowBlockingDelay field to given value.

### HasAllowBlockingDelay

`func (o *AddDelayBindResponseFailureLockoutActionRequest) HasAllowBlockingDelay() bool`

HasAllowBlockingDelay returns a boolean if a field has been set.

### GetGenerateAccountStatusNotification

`func (o *AddDelayBindResponseFailureLockoutActionRequest) GetGenerateAccountStatusNotification() bool`

GetGenerateAccountStatusNotification returns the GenerateAccountStatusNotification field if non-nil, zero value otherwise.

### GetGenerateAccountStatusNotificationOk

`func (o *AddDelayBindResponseFailureLockoutActionRequest) GetGenerateAccountStatusNotificationOk() (*bool, bool)`

GetGenerateAccountStatusNotificationOk returns a tuple with the GenerateAccountStatusNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateAccountStatusNotification

`func (o *AddDelayBindResponseFailureLockoutActionRequest) SetGenerateAccountStatusNotification(v bool)`

SetGenerateAccountStatusNotification sets GenerateAccountStatusNotification field to given value.

### HasGenerateAccountStatusNotification

`func (o *AddDelayBindResponseFailureLockoutActionRequest) HasGenerateAccountStatusNotification() bool`

HasGenerateAccountStatusNotification returns a boolean if a field has been set.

### GetDescription

`func (o *AddDelayBindResponseFailureLockoutActionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddDelayBindResponseFailureLockoutActionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddDelayBindResponseFailureLockoutActionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddDelayBindResponseFailureLockoutActionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


