# WaitForPassphraseCipherStreamProviderShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumwaitForPassphraseCipherStreamProviderSchemaUrn**](EnumwaitForPassphraseCipherStreamProviderSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 

## Methods

### NewWaitForPassphraseCipherStreamProviderShared

`func NewWaitForPassphraseCipherStreamProviderShared(schemas []EnumwaitForPassphraseCipherStreamProviderSchemaUrn, enabled bool, ) *WaitForPassphraseCipherStreamProviderShared`

NewWaitForPassphraseCipherStreamProviderShared instantiates a new WaitForPassphraseCipherStreamProviderShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaitForPassphraseCipherStreamProviderSharedWithDefaults

`func NewWaitForPassphraseCipherStreamProviderSharedWithDefaults() *WaitForPassphraseCipherStreamProviderShared`

NewWaitForPassphraseCipherStreamProviderSharedWithDefaults instantiates a new WaitForPassphraseCipherStreamProviderShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *WaitForPassphraseCipherStreamProviderShared) GetSchemas() []EnumwaitForPassphraseCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *WaitForPassphraseCipherStreamProviderShared) GetSchemasOk() (*[]EnumwaitForPassphraseCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *WaitForPassphraseCipherStreamProviderShared) SetSchemas(v []EnumwaitForPassphraseCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *WaitForPassphraseCipherStreamProviderShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WaitForPassphraseCipherStreamProviderShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WaitForPassphraseCipherStreamProviderShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WaitForPassphraseCipherStreamProviderShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *WaitForPassphraseCipherStreamProviderShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WaitForPassphraseCipherStreamProviderShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WaitForPassphraseCipherStreamProviderShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


