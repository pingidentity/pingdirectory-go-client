# FailureLockoutActionListResponseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Failure Lockout Action | 
**Schemas** | [**[]EnumlockAccountFailureLockoutActionSchemaUrn**](EnumlockAccountFailureLockoutActionSchemaUrn.md) |  | 
**GenerateAccountStatusNotification** | Pointer to **bool** | Indicates whether to generate an account status notification for cases in which a bind response is delayed because of failure lockout. | [optional] 
**Description** | Pointer to **string** | A description for this Failure Lockout Action | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Delay** | **string** | The length of time to delay the bind response for accounts with too many failed authentication attempts. | 
**AllowBlockingDelay** | Pointer to **bool** | Indicates whether to delay the response for authentication attempts even if that delay may block the thread being used to process the attempt. | [optional] 

## Methods

### NewFailureLockoutActionListResponseResourcesInner

`func NewFailureLockoutActionListResponseResourcesInner(id string, schemas []EnumlockAccountFailureLockoutActionSchemaUrn, delay string, ) *FailureLockoutActionListResponseResourcesInner`

NewFailureLockoutActionListResponseResourcesInner instantiates a new FailureLockoutActionListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureLockoutActionListResponseResourcesInnerWithDefaults

`func NewFailureLockoutActionListResponseResourcesInnerWithDefaults() *FailureLockoutActionListResponseResourcesInner`

NewFailureLockoutActionListResponseResourcesInnerWithDefaults instantiates a new FailureLockoutActionListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FailureLockoutActionListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FailureLockoutActionListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FailureLockoutActionListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *FailureLockoutActionListResponseResourcesInner) GetSchemas() []EnumlockAccountFailureLockoutActionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FailureLockoutActionListResponseResourcesInner) GetSchemasOk() (*[]EnumlockAccountFailureLockoutActionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FailureLockoutActionListResponseResourcesInner) SetSchemas(v []EnumlockAccountFailureLockoutActionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetGenerateAccountStatusNotification

`func (o *FailureLockoutActionListResponseResourcesInner) GetGenerateAccountStatusNotification() bool`

GetGenerateAccountStatusNotification returns the GenerateAccountStatusNotification field if non-nil, zero value otherwise.

### GetGenerateAccountStatusNotificationOk

`func (o *FailureLockoutActionListResponseResourcesInner) GetGenerateAccountStatusNotificationOk() (*bool, bool)`

GetGenerateAccountStatusNotificationOk returns a tuple with the GenerateAccountStatusNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateAccountStatusNotification

`func (o *FailureLockoutActionListResponseResourcesInner) SetGenerateAccountStatusNotification(v bool)`

SetGenerateAccountStatusNotification sets GenerateAccountStatusNotification field to given value.

### HasGenerateAccountStatusNotification

`func (o *FailureLockoutActionListResponseResourcesInner) HasGenerateAccountStatusNotification() bool`

HasGenerateAccountStatusNotification returns a boolean if a field has been set.

### GetDescription

`func (o *FailureLockoutActionListResponseResourcesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FailureLockoutActionListResponseResourcesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FailureLockoutActionListResponseResourcesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FailureLockoutActionListResponseResourcesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *FailureLockoutActionListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FailureLockoutActionListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FailureLockoutActionListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FailureLockoutActionListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *FailureLockoutActionListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *FailureLockoutActionListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *FailureLockoutActionListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *FailureLockoutActionListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetDelay

`func (o *FailureLockoutActionListResponseResourcesInner) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *FailureLockoutActionListResponseResourcesInner) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *FailureLockoutActionListResponseResourcesInner) SetDelay(v string)`

SetDelay sets Delay field to given value.


### GetAllowBlockingDelay

`func (o *FailureLockoutActionListResponseResourcesInner) GetAllowBlockingDelay() bool`

GetAllowBlockingDelay returns the AllowBlockingDelay field if non-nil, zero value otherwise.

### GetAllowBlockingDelayOk

`func (o *FailureLockoutActionListResponseResourcesInner) GetAllowBlockingDelayOk() (*bool, bool)`

GetAllowBlockingDelayOk returns a tuple with the AllowBlockingDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBlockingDelay

`func (o *FailureLockoutActionListResponseResourcesInner) SetAllowBlockingDelay(v bool)`

SetAllowBlockingDelay sets AllowBlockingDelay field to given value.

### HasAllowBlockingDelay

`func (o *FailureLockoutActionListResponseResourcesInner) HasAllowBlockingDelay() bool`

HasAllowBlockingDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


