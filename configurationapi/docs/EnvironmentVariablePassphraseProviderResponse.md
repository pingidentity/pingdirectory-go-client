# EnvironmentVariablePassphraseProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Passphrase Provider | 
**Schemas** | [**[]EnumenvironmentVariablePassphraseProviderSchemaUrn**](EnumenvironmentVariablePassphraseProviderSchemaUrn.md) |  | 
**EnvironmentVariable** | **string** | The name of the environment variable that is expected to hold the passphrase. | 
**Description** | Pointer to **string** | A description for this Passphrase Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Passphrase Provider is enabled for use in the server. | 

## Methods

### NewEnvironmentVariablePassphraseProviderResponse

`func NewEnvironmentVariablePassphraseProviderResponse(id string, schemas []EnumenvironmentVariablePassphraseProviderSchemaUrn, environmentVariable string, enabled bool, ) *EnvironmentVariablePassphraseProviderResponse`

NewEnvironmentVariablePassphraseProviderResponse instantiates a new EnvironmentVariablePassphraseProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentVariablePassphraseProviderResponseWithDefaults

`func NewEnvironmentVariablePassphraseProviderResponseWithDefaults() *EnvironmentVariablePassphraseProviderResponse`

NewEnvironmentVariablePassphraseProviderResponseWithDefaults instantiates a new EnvironmentVariablePassphraseProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *EnvironmentVariablePassphraseProviderResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EnvironmentVariablePassphraseProviderResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EnvironmentVariablePassphraseProviderResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EnvironmentVariablePassphraseProviderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *EnvironmentVariablePassphraseProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *EnvironmentVariablePassphraseProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *EnvironmentVariablePassphraseProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *EnvironmentVariablePassphraseProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *EnvironmentVariablePassphraseProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentVariablePassphraseProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentVariablePassphraseProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *EnvironmentVariablePassphraseProviderResponse) GetSchemas() []EnumenvironmentVariablePassphraseProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *EnvironmentVariablePassphraseProviderResponse) GetSchemasOk() (*[]EnumenvironmentVariablePassphraseProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *EnvironmentVariablePassphraseProviderResponse) SetSchemas(v []EnumenvironmentVariablePassphraseProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnvironmentVariable

`func (o *EnvironmentVariablePassphraseProviderResponse) GetEnvironmentVariable() string`

GetEnvironmentVariable returns the EnvironmentVariable field if non-nil, zero value otherwise.

### GetEnvironmentVariableOk

`func (o *EnvironmentVariablePassphraseProviderResponse) GetEnvironmentVariableOk() (*string, bool)`

GetEnvironmentVariableOk returns a tuple with the EnvironmentVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariable

`func (o *EnvironmentVariablePassphraseProviderResponse) SetEnvironmentVariable(v string)`

SetEnvironmentVariable sets EnvironmentVariable field to given value.


### GetDescription

`func (o *EnvironmentVariablePassphraseProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentVariablePassphraseProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentVariablePassphraseProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentVariablePassphraseProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *EnvironmentVariablePassphraseProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EnvironmentVariablePassphraseProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EnvironmentVariablePassphraseProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


