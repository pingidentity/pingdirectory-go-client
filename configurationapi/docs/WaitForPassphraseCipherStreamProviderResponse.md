# WaitForPassphraseCipherStreamProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumwaitForPassphraseCipherStreamProviderSchemaUrn**](EnumwaitForPassphraseCipherStreamProviderSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Cipher Stream Provider | 

## Methods

### NewWaitForPassphraseCipherStreamProviderResponse

`func NewWaitForPassphraseCipherStreamProviderResponse(schemas []EnumwaitForPassphraseCipherStreamProviderSchemaUrn, enabled bool, id string, ) *WaitForPassphraseCipherStreamProviderResponse`

NewWaitForPassphraseCipherStreamProviderResponse instantiates a new WaitForPassphraseCipherStreamProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaitForPassphraseCipherStreamProviderResponseWithDefaults

`func NewWaitForPassphraseCipherStreamProviderResponseWithDefaults() *WaitForPassphraseCipherStreamProviderResponse`

NewWaitForPassphraseCipherStreamProviderResponseWithDefaults instantiates a new WaitForPassphraseCipherStreamProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *WaitForPassphraseCipherStreamProviderResponse) GetSchemas() []EnumwaitForPassphraseCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *WaitForPassphraseCipherStreamProviderResponse) GetSchemasOk() (*[]EnumwaitForPassphraseCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *WaitForPassphraseCipherStreamProviderResponse) SetSchemas(v []EnumwaitForPassphraseCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *WaitForPassphraseCipherStreamProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WaitForPassphraseCipherStreamProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WaitForPassphraseCipherStreamProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WaitForPassphraseCipherStreamProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *WaitForPassphraseCipherStreamProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WaitForPassphraseCipherStreamProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WaitForPassphraseCipherStreamProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *WaitForPassphraseCipherStreamProviderResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WaitForPassphraseCipherStreamProviderResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WaitForPassphraseCipherStreamProviderResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WaitForPassphraseCipherStreamProviderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *WaitForPassphraseCipherStreamProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *WaitForPassphraseCipherStreamProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *WaitForPassphraseCipherStreamProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *WaitForPassphraseCipherStreamProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *WaitForPassphraseCipherStreamProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WaitForPassphraseCipherStreamProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WaitForPassphraseCipherStreamProviderResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


