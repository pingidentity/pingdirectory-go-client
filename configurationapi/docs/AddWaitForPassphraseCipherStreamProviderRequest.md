# AddWaitForPassphraseCipherStreamProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Name of the new Cipher Stream Provider | 
**Schemas** | [**[]EnumwaitForPassphraseCipherStreamProviderSchemaUrn**](EnumwaitForPassphraseCipherStreamProviderSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 

## Methods

### NewAddWaitForPassphraseCipherStreamProviderRequest

`func NewAddWaitForPassphraseCipherStreamProviderRequest(providerName string, schemas []EnumwaitForPassphraseCipherStreamProviderSchemaUrn, enabled bool, ) *AddWaitForPassphraseCipherStreamProviderRequest`

NewAddWaitForPassphraseCipherStreamProviderRequest instantiates a new AddWaitForPassphraseCipherStreamProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddWaitForPassphraseCipherStreamProviderRequestWithDefaults

`func NewAddWaitForPassphraseCipherStreamProviderRequestWithDefaults() *AddWaitForPassphraseCipherStreamProviderRequest`

NewAddWaitForPassphraseCipherStreamProviderRequestWithDefaults instantiates a new AddWaitForPassphraseCipherStreamProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *AddWaitForPassphraseCipherStreamProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddWaitForPassphraseCipherStreamProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddWaitForPassphraseCipherStreamProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetSchemas

`func (o *AddWaitForPassphraseCipherStreamProviderRequest) GetSchemas() []EnumwaitForPassphraseCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddWaitForPassphraseCipherStreamProviderRequest) GetSchemasOk() (*[]EnumwaitForPassphraseCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddWaitForPassphraseCipherStreamProviderRequest) SetSchemas(v []EnumwaitForPassphraseCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddWaitForPassphraseCipherStreamProviderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddWaitForPassphraseCipherStreamProviderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddWaitForPassphraseCipherStreamProviderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddWaitForPassphraseCipherStreamProviderRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddWaitForPassphraseCipherStreamProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddWaitForPassphraseCipherStreamProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddWaitForPassphraseCipherStreamProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


