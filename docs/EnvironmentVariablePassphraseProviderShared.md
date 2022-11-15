# EnvironmentVariablePassphraseProviderShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumenvironmentVariablePassphraseProviderSchemaUrn**](EnumenvironmentVariablePassphraseProviderSchemaUrn.md) |  | 
**EnvironmentVariable** | **string** | The name of the environment variable that is expected to hold the passphrase. | 
**Description** | Pointer to **string** | A description for this Passphrase Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Passphrase Provider is enabled for use in the server. | 

## Methods

### NewEnvironmentVariablePassphraseProviderShared

`func NewEnvironmentVariablePassphraseProviderShared(schemas []EnumenvironmentVariablePassphraseProviderSchemaUrn, environmentVariable string, enabled bool, ) *EnvironmentVariablePassphraseProviderShared`

NewEnvironmentVariablePassphraseProviderShared instantiates a new EnvironmentVariablePassphraseProviderShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentVariablePassphraseProviderSharedWithDefaults

`func NewEnvironmentVariablePassphraseProviderSharedWithDefaults() *EnvironmentVariablePassphraseProviderShared`

NewEnvironmentVariablePassphraseProviderSharedWithDefaults instantiates a new EnvironmentVariablePassphraseProviderShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *EnvironmentVariablePassphraseProviderShared) GetSchemas() []EnumenvironmentVariablePassphraseProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *EnvironmentVariablePassphraseProviderShared) GetSchemasOk() (*[]EnumenvironmentVariablePassphraseProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *EnvironmentVariablePassphraseProviderShared) SetSchemas(v []EnumenvironmentVariablePassphraseProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnvironmentVariable

`func (o *EnvironmentVariablePassphraseProviderShared) GetEnvironmentVariable() string`

GetEnvironmentVariable returns the EnvironmentVariable field if non-nil, zero value otherwise.

### GetEnvironmentVariableOk

`func (o *EnvironmentVariablePassphraseProviderShared) GetEnvironmentVariableOk() (*string, bool)`

GetEnvironmentVariableOk returns a tuple with the EnvironmentVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariable

`func (o *EnvironmentVariablePassphraseProviderShared) SetEnvironmentVariable(v string)`

SetEnvironmentVariable sets EnvironmentVariable field to given value.


### GetDescription

`func (o *EnvironmentVariablePassphraseProviderShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentVariablePassphraseProviderShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentVariablePassphraseProviderShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentVariablePassphraseProviderShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *EnvironmentVariablePassphraseProviderShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EnvironmentVariablePassphraseProviderShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EnvironmentVariablePassphraseProviderShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


