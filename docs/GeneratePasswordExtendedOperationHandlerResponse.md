# GeneratePasswordExtendedOperationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumgeneratePasswordExtendedOperationHandlerSchemaUrn**](EnumgeneratePasswordExtendedOperationHandlerSchemaUrn.md) |  | 
**Id** | Pointer to **string** | Name of the Extended Operation Handler | [optional] 
**DefaultPasswordPolicy** | Pointer to **string** | The default password policy that should be used when generating and validating passwords if the request does not specify an alternate policy. If this is not provided, then this Generate Password Extended Operation Handler will use the default password policy defined in the global configuration. | [optional] 
**DefaultPasswordGenerator** | **string** | The default password generator that will be used if the selected password policy is not configured with a password generator. | 
**MaximumPasswordsPerRequest** | Pointer to **int32** | The maximum number of passwords that may be generated and returned to the client for a single request. | [optional] 
**MaximumValidationAttemptsPerPassword** | Pointer to **int32** | The maximum number of attempts that the server may use to generate a password that passes validation. | [optional] 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewGeneratePasswordExtendedOperationHandlerResponse

`func NewGeneratePasswordExtendedOperationHandlerResponse(schemas []EnumgeneratePasswordExtendedOperationHandlerSchemaUrn, defaultPasswordGenerator string, enabled bool, ) *GeneratePasswordExtendedOperationHandlerResponse`

NewGeneratePasswordExtendedOperationHandlerResponse instantiates a new GeneratePasswordExtendedOperationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneratePasswordExtendedOperationHandlerResponseWithDefaults

`func NewGeneratePasswordExtendedOperationHandlerResponseWithDefaults() *GeneratePasswordExtendedOperationHandlerResponse`

NewGeneratePasswordExtendedOperationHandlerResponseWithDefaults instantiates a new GeneratePasswordExtendedOperationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GeneratePasswordExtendedOperationHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GeneratePasswordExtendedOperationHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GeneratePasswordExtendedOperationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GeneratePasswordExtendedOperationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetSchemas() []EnumgeneratePasswordExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetSchemasOk() (*[]EnumgeneratePasswordExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GeneratePasswordExtendedOperationHandlerResponse) SetSchemas(v []EnumgeneratePasswordExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GeneratePasswordExtendedOperationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GeneratePasswordExtendedOperationHandlerResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDefaultPasswordPolicy

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetDefaultPasswordPolicy() string`

GetDefaultPasswordPolicy returns the DefaultPasswordPolicy field if non-nil, zero value otherwise.

### GetDefaultPasswordPolicyOk

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetDefaultPasswordPolicyOk() (*string, bool)`

GetDefaultPasswordPolicyOk returns a tuple with the DefaultPasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPasswordPolicy

`func (o *GeneratePasswordExtendedOperationHandlerResponse) SetDefaultPasswordPolicy(v string)`

SetDefaultPasswordPolicy sets DefaultPasswordPolicy field to given value.

### HasDefaultPasswordPolicy

`func (o *GeneratePasswordExtendedOperationHandlerResponse) HasDefaultPasswordPolicy() bool`

HasDefaultPasswordPolicy returns a boolean if a field has been set.

### GetDefaultPasswordGenerator

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetDefaultPasswordGenerator() string`

GetDefaultPasswordGenerator returns the DefaultPasswordGenerator field if non-nil, zero value otherwise.

### GetDefaultPasswordGeneratorOk

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetDefaultPasswordGeneratorOk() (*string, bool)`

GetDefaultPasswordGeneratorOk returns a tuple with the DefaultPasswordGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPasswordGenerator

`func (o *GeneratePasswordExtendedOperationHandlerResponse) SetDefaultPasswordGenerator(v string)`

SetDefaultPasswordGenerator sets DefaultPasswordGenerator field to given value.


### GetMaximumPasswordsPerRequest

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetMaximumPasswordsPerRequest() int32`

GetMaximumPasswordsPerRequest returns the MaximumPasswordsPerRequest field if non-nil, zero value otherwise.

### GetMaximumPasswordsPerRequestOk

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetMaximumPasswordsPerRequestOk() (*int32, bool)`

GetMaximumPasswordsPerRequestOk returns a tuple with the MaximumPasswordsPerRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPasswordsPerRequest

`func (o *GeneratePasswordExtendedOperationHandlerResponse) SetMaximumPasswordsPerRequest(v int32)`

SetMaximumPasswordsPerRequest sets MaximumPasswordsPerRequest field to given value.

### HasMaximumPasswordsPerRequest

`func (o *GeneratePasswordExtendedOperationHandlerResponse) HasMaximumPasswordsPerRequest() bool`

HasMaximumPasswordsPerRequest returns a boolean if a field has been set.

### GetMaximumValidationAttemptsPerPassword

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetMaximumValidationAttemptsPerPassword() int32`

GetMaximumValidationAttemptsPerPassword returns the MaximumValidationAttemptsPerPassword field if non-nil, zero value otherwise.

### GetMaximumValidationAttemptsPerPasswordOk

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetMaximumValidationAttemptsPerPasswordOk() (*int32, bool)`

GetMaximumValidationAttemptsPerPasswordOk returns a tuple with the MaximumValidationAttemptsPerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValidationAttemptsPerPassword

`func (o *GeneratePasswordExtendedOperationHandlerResponse) SetMaximumValidationAttemptsPerPassword(v int32)`

SetMaximumValidationAttemptsPerPassword sets MaximumValidationAttemptsPerPassword field to given value.

### HasMaximumValidationAttemptsPerPassword

`func (o *GeneratePasswordExtendedOperationHandlerResponse) HasMaximumValidationAttemptsPerPassword() bool`

HasMaximumValidationAttemptsPerPassword returns a boolean if a field has been set.

### GetDescription

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GeneratePasswordExtendedOperationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GeneratePasswordExtendedOperationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GeneratePasswordExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GeneratePasswordExtendedOperationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


