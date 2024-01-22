# DelayBindResponseFailureLockoutActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumdelayBindResponseFailureLockoutActionSchemaUrn**](EnumdelayBindResponseFailureLockoutActionSchemaUrn.md) |  | 
**Delay** | **string** | The length of time to delay the bind response for accounts with too many failed authentication attempts. | 
**AllowBlockingDelay** | Pointer to **bool** | Indicates whether to delay the response for authentication attempts even if that delay may block the thread being used to process the attempt. | [optional] 
**GenerateAccountStatusNotification** | Pointer to **bool** | Indicates whether to generate an account status notification for cases in which a bind response is delayed because of failure lockout. | [optional] 
**Description** | Pointer to **string** | A description for this Failure Lockout Action | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Failure Lockout Action | 

## Methods

### NewDelayBindResponseFailureLockoutActionResponse

`func NewDelayBindResponseFailureLockoutActionResponse(schemas []EnumdelayBindResponseFailureLockoutActionSchemaUrn, delay string, id string, ) *DelayBindResponseFailureLockoutActionResponse`

NewDelayBindResponseFailureLockoutActionResponse instantiates a new DelayBindResponseFailureLockoutActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelayBindResponseFailureLockoutActionResponseWithDefaults

`func NewDelayBindResponseFailureLockoutActionResponseWithDefaults() *DelayBindResponseFailureLockoutActionResponse`

NewDelayBindResponseFailureLockoutActionResponseWithDefaults instantiates a new DelayBindResponseFailureLockoutActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *DelayBindResponseFailureLockoutActionResponse) GetSchemas() []EnumdelayBindResponseFailureLockoutActionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DelayBindResponseFailureLockoutActionResponse) GetSchemasOk() (*[]EnumdelayBindResponseFailureLockoutActionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DelayBindResponseFailureLockoutActionResponse) SetSchemas(v []EnumdelayBindResponseFailureLockoutActionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDelay

`func (o *DelayBindResponseFailureLockoutActionResponse) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *DelayBindResponseFailureLockoutActionResponse) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *DelayBindResponseFailureLockoutActionResponse) SetDelay(v string)`

SetDelay sets Delay field to given value.


### GetAllowBlockingDelay

`func (o *DelayBindResponseFailureLockoutActionResponse) GetAllowBlockingDelay() bool`

GetAllowBlockingDelay returns the AllowBlockingDelay field if non-nil, zero value otherwise.

### GetAllowBlockingDelayOk

`func (o *DelayBindResponseFailureLockoutActionResponse) GetAllowBlockingDelayOk() (*bool, bool)`

GetAllowBlockingDelayOk returns a tuple with the AllowBlockingDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBlockingDelay

`func (o *DelayBindResponseFailureLockoutActionResponse) SetAllowBlockingDelay(v bool)`

SetAllowBlockingDelay sets AllowBlockingDelay field to given value.

### HasAllowBlockingDelay

`func (o *DelayBindResponseFailureLockoutActionResponse) HasAllowBlockingDelay() bool`

HasAllowBlockingDelay returns a boolean if a field has been set.

### GetGenerateAccountStatusNotification

`func (o *DelayBindResponseFailureLockoutActionResponse) GetGenerateAccountStatusNotification() bool`

GetGenerateAccountStatusNotification returns the GenerateAccountStatusNotification field if non-nil, zero value otherwise.

### GetGenerateAccountStatusNotificationOk

`func (o *DelayBindResponseFailureLockoutActionResponse) GetGenerateAccountStatusNotificationOk() (*bool, bool)`

GetGenerateAccountStatusNotificationOk returns a tuple with the GenerateAccountStatusNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateAccountStatusNotification

`func (o *DelayBindResponseFailureLockoutActionResponse) SetGenerateAccountStatusNotification(v bool)`

SetGenerateAccountStatusNotification sets GenerateAccountStatusNotification field to given value.

### HasGenerateAccountStatusNotification

`func (o *DelayBindResponseFailureLockoutActionResponse) HasGenerateAccountStatusNotification() bool`

HasGenerateAccountStatusNotification returns a boolean if a field has been set.

### GetDescription

`func (o *DelayBindResponseFailureLockoutActionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DelayBindResponseFailureLockoutActionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DelayBindResponseFailureLockoutActionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DelayBindResponseFailureLockoutActionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *DelayBindResponseFailureLockoutActionResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DelayBindResponseFailureLockoutActionResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DelayBindResponseFailureLockoutActionResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DelayBindResponseFailureLockoutActionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *DelayBindResponseFailureLockoutActionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *DelayBindResponseFailureLockoutActionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *DelayBindResponseFailureLockoutActionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *DelayBindResponseFailureLockoutActionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *DelayBindResponseFailureLockoutActionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DelayBindResponseFailureLockoutActionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DelayBindResponseFailureLockoutActionResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


