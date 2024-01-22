# AddObscuredValuePassphraseProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumobscuredValuePassphraseProviderSchemaUrn**](EnumobscuredValuePassphraseProviderSchemaUrn.md) |  | 
**ObscuredValue** | **string** | The value to be stored in an obscured form. | 
**Description** | Pointer to **string** | A description for this Passphrase Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Passphrase Provider is enabled for use in the server. | 
**ProviderName** | **string** | Name of the new Passphrase Provider | 

## Methods

### NewAddObscuredValuePassphraseProviderRequest

`func NewAddObscuredValuePassphraseProviderRequest(schemas []EnumobscuredValuePassphraseProviderSchemaUrn, obscuredValue string, enabled bool, providerName string, ) *AddObscuredValuePassphraseProviderRequest`

NewAddObscuredValuePassphraseProviderRequest instantiates a new AddObscuredValuePassphraseProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddObscuredValuePassphraseProviderRequestWithDefaults

`func NewAddObscuredValuePassphraseProviderRequestWithDefaults() *AddObscuredValuePassphraseProviderRequest`

NewAddObscuredValuePassphraseProviderRequestWithDefaults instantiates a new AddObscuredValuePassphraseProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddObscuredValuePassphraseProviderRequest) GetSchemas() []EnumobscuredValuePassphraseProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddObscuredValuePassphraseProviderRequest) GetSchemasOk() (*[]EnumobscuredValuePassphraseProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddObscuredValuePassphraseProviderRequest) SetSchemas(v []EnumobscuredValuePassphraseProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetObscuredValue

`func (o *AddObscuredValuePassphraseProviderRequest) GetObscuredValue() string`

GetObscuredValue returns the ObscuredValue field if non-nil, zero value otherwise.

### GetObscuredValueOk

`func (o *AddObscuredValuePassphraseProviderRequest) GetObscuredValueOk() (*string, bool)`

GetObscuredValueOk returns a tuple with the ObscuredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObscuredValue

`func (o *AddObscuredValuePassphraseProviderRequest) SetObscuredValue(v string)`

SetObscuredValue sets ObscuredValue field to given value.


### GetDescription

`func (o *AddObscuredValuePassphraseProviderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddObscuredValuePassphraseProviderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddObscuredValuePassphraseProviderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddObscuredValuePassphraseProviderRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddObscuredValuePassphraseProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddObscuredValuePassphraseProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddObscuredValuePassphraseProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProviderName

`func (o *AddObscuredValuePassphraseProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddObscuredValuePassphraseProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddObscuredValuePassphraseProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


