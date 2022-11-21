# AddFailureLockoutAction200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Failure Lockout Action | 
**Schemas** | [**[]EnumlockAccountFailureLockoutActionSchemaUrn**](EnumlockAccountFailureLockoutActionSchemaUrn.md) |  | 
**Delay** | **string** | The length of time to delay the bind response for accounts with too many failed authentication attempts. | 
**AllowBlockingDelay** | Pointer to **bool** | Indicates whether to delay the response for authentication attempts even if that delay may block the thread being used to process the attempt. | [optional] 
**GenerateAccountStatusNotification** | Pointer to **bool** | Indicates whether to generate an account status notification for cases in which this failure lockout action is invoked for a bind attempt with too many outstanding authentication failures. | [optional] 
**Description** | Pointer to **string** | A description for this Failure Lockout Action | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewAddFailureLockoutAction200Response

`func NewAddFailureLockoutAction200Response(id string, schemas []EnumlockAccountFailureLockoutActionSchemaUrn, delay string, ) *AddFailureLockoutAction200Response`

NewAddFailureLockoutAction200Response instantiates a new AddFailureLockoutAction200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFailureLockoutAction200ResponseWithDefaults

`func NewAddFailureLockoutAction200ResponseWithDefaults() *AddFailureLockoutAction200Response`

NewAddFailureLockoutAction200ResponseWithDefaults instantiates a new AddFailureLockoutAction200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddFailureLockoutAction200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddFailureLockoutAction200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddFailureLockoutAction200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddFailureLockoutAction200Response) GetSchemas() []EnumlockAccountFailureLockoutActionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddFailureLockoutAction200Response) GetSchemasOk() (*[]EnumlockAccountFailureLockoutActionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddFailureLockoutAction200Response) SetSchemas(v []EnumlockAccountFailureLockoutActionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDelay

`func (o *AddFailureLockoutAction200Response) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *AddFailureLockoutAction200Response) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *AddFailureLockoutAction200Response) SetDelay(v string)`

SetDelay sets Delay field to given value.


### GetAllowBlockingDelay

`func (o *AddFailureLockoutAction200Response) GetAllowBlockingDelay() bool`

GetAllowBlockingDelay returns the AllowBlockingDelay field if non-nil, zero value otherwise.

### GetAllowBlockingDelayOk

`func (o *AddFailureLockoutAction200Response) GetAllowBlockingDelayOk() (*bool, bool)`

GetAllowBlockingDelayOk returns a tuple with the AllowBlockingDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBlockingDelay

`func (o *AddFailureLockoutAction200Response) SetAllowBlockingDelay(v bool)`

SetAllowBlockingDelay sets AllowBlockingDelay field to given value.

### HasAllowBlockingDelay

`func (o *AddFailureLockoutAction200Response) HasAllowBlockingDelay() bool`

HasAllowBlockingDelay returns a boolean if a field has been set.

### GetGenerateAccountStatusNotification

`func (o *AddFailureLockoutAction200Response) GetGenerateAccountStatusNotification() bool`

GetGenerateAccountStatusNotification returns the GenerateAccountStatusNotification field if non-nil, zero value otherwise.

### GetGenerateAccountStatusNotificationOk

`func (o *AddFailureLockoutAction200Response) GetGenerateAccountStatusNotificationOk() (*bool, bool)`

GetGenerateAccountStatusNotificationOk returns a tuple with the GenerateAccountStatusNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateAccountStatusNotification

`func (o *AddFailureLockoutAction200Response) SetGenerateAccountStatusNotification(v bool)`

SetGenerateAccountStatusNotification sets GenerateAccountStatusNotification field to given value.

### HasGenerateAccountStatusNotification

`func (o *AddFailureLockoutAction200Response) HasGenerateAccountStatusNotification() bool`

HasGenerateAccountStatusNotification returns a boolean if a field has been set.

### GetDescription

`func (o *AddFailureLockoutAction200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddFailureLockoutAction200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddFailureLockoutAction200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddFailureLockoutAction200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *AddFailureLockoutAction200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddFailureLockoutAction200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddFailureLockoutAction200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddFailureLockoutAction200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


