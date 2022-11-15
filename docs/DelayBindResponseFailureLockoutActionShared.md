# DelayBindResponseFailureLockoutActionShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumdelayBindResponseFailureLockoutActionSchemaUrn**](EnumdelayBindResponseFailureLockoutActionSchemaUrn.md) |  | 
**Delay** | **string** | The length of time to delay the bind response for accounts with too many failed authentication attempts. | 
**AllowBlockingDelay** | Pointer to **bool** | Indicates whether to delay the response for authentication attempts even if that delay may block the thread being used to process the attempt. | [optional] 
**GenerateAccountStatusNotification** | Pointer to **bool** | Indicates whether to generate an account status notification for cases in which a bind response is delayed because of failure lockout. | [optional] 
**Description** | Pointer to **string** | A description for this Failure Lockout Action | [optional] 

## Methods

### NewDelayBindResponseFailureLockoutActionShared

`func NewDelayBindResponseFailureLockoutActionShared(schemas []EnumdelayBindResponseFailureLockoutActionSchemaUrn, delay string, ) *DelayBindResponseFailureLockoutActionShared`

NewDelayBindResponseFailureLockoutActionShared instantiates a new DelayBindResponseFailureLockoutActionShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelayBindResponseFailureLockoutActionSharedWithDefaults

`func NewDelayBindResponseFailureLockoutActionSharedWithDefaults() *DelayBindResponseFailureLockoutActionShared`

NewDelayBindResponseFailureLockoutActionSharedWithDefaults instantiates a new DelayBindResponseFailureLockoutActionShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *DelayBindResponseFailureLockoutActionShared) GetSchemas() []EnumdelayBindResponseFailureLockoutActionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DelayBindResponseFailureLockoutActionShared) GetSchemasOk() (*[]EnumdelayBindResponseFailureLockoutActionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DelayBindResponseFailureLockoutActionShared) SetSchemas(v []EnumdelayBindResponseFailureLockoutActionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDelay

`func (o *DelayBindResponseFailureLockoutActionShared) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *DelayBindResponseFailureLockoutActionShared) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *DelayBindResponseFailureLockoutActionShared) SetDelay(v string)`

SetDelay sets Delay field to given value.


### GetAllowBlockingDelay

`func (o *DelayBindResponseFailureLockoutActionShared) GetAllowBlockingDelay() bool`

GetAllowBlockingDelay returns the AllowBlockingDelay field if non-nil, zero value otherwise.

### GetAllowBlockingDelayOk

`func (o *DelayBindResponseFailureLockoutActionShared) GetAllowBlockingDelayOk() (*bool, bool)`

GetAllowBlockingDelayOk returns a tuple with the AllowBlockingDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBlockingDelay

`func (o *DelayBindResponseFailureLockoutActionShared) SetAllowBlockingDelay(v bool)`

SetAllowBlockingDelay sets AllowBlockingDelay field to given value.

### HasAllowBlockingDelay

`func (o *DelayBindResponseFailureLockoutActionShared) HasAllowBlockingDelay() bool`

HasAllowBlockingDelay returns a boolean if a field has been set.

### GetGenerateAccountStatusNotification

`func (o *DelayBindResponseFailureLockoutActionShared) GetGenerateAccountStatusNotification() bool`

GetGenerateAccountStatusNotification returns the GenerateAccountStatusNotification field if non-nil, zero value otherwise.

### GetGenerateAccountStatusNotificationOk

`func (o *DelayBindResponseFailureLockoutActionShared) GetGenerateAccountStatusNotificationOk() (*bool, bool)`

GetGenerateAccountStatusNotificationOk returns a tuple with the GenerateAccountStatusNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateAccountStatusNotification

`func (o *DelayBindResponseFailureLockoutActionShared) SetGenerateAccountStatusNotification(v bool)`

SetGenerateAccountStatusNotification sets GenerateAccountStatusNotification field to given value.

### HasGenerateAccountStatusNotification

`func (o *DelayBindResponseFailureLockoutActionShared) HasGenerateAccountStatusNotification() bool`

HasGenerateAccountStatusNotification returns a boolean if a field has been set.

### GetDescription

`func (o *DelayBindResponseFailureLockoutActionShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DelayBindResponseFailureLockoutActionShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DelayBindResponseFailureLockoutActionShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DelayBindResponseFailureLockoutActionShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


