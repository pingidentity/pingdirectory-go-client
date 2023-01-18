# AddGroovyScriptedPasswordValidatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidatorName** | **string** | Name of the new Password Validator | 
**Schemas** | [**[]EnumgroovyScriptedPasswordValidatorSchemaUrn**](EnumgroovyScriptedPasswordValidatorSchemaUrn.md) |  | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Password Validator. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted Password Validator. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Password Validator | [optional] 
**Enabled** | **bool** | Indicates whether the password validator is enabled for use. | 
**ValidatorRequirementDescription** | Pointer to **string** | Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator. | [optional] 
**ValidatorFailureMessage** | Pointer to **string** | Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator. | [optional] 

## Methods

### NewAddGroovyScriptedPasswordValidatorRequest

`func NewAddGroovyScriptedPasswordValidatorRequest(validatorName string, schemas []EnumgroovyScriptedPasswordValidatorSchemaUrn, scriptClass string, enabled bool, ) *AddGroovyScriptedPasswordValidatorRequest`

NewAddGroovyScriptedPasswordValidatorRequest instantiates a new AddGroovyScriptedPasswordValidatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGroovyScriptedPasswordValidatorRequestWithDefaults

`func NewAddGroovyScriptedPasswordValidatorRequestWithDefaults() *AddGroovyScriptedPasswordValidatorRequest`

NewAddGroovyScriptedPasswordValidatorRequestWithDefaults instantiates a new AddGroovyScriptedPasswordValidatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidatorName

`func (o *AddGroovyScriptedPasswordValidatorRequest) GetValidatorName() string`

GetValidatorName returns the ValidatorName field if non-nil, zero value otherwise.

### GetValidatorNameOk

`func (o *AddGroovyScriptedPasswordValidatorRequest) GetValidatorNameOk() (*string, bool)`

GetValidatorNameOk returns a tuple with the ValidatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorName

`func (o *AddGroovyScriptedPasswordValidatorRequest) SetValidatorName(v string)`

SetValidatorName sets ValidatorName field to given value.


### GetSchemas

`func (o *AddGroovyScriptedPasswordValidatorRequest) GetSchemas() []EnumgroovyScriptedPasswordValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddGroovyScriptedPasswordValidatorRequest) GetSchemasOk() (*[]EnumgroovyScriptedPasswordValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddGroovyScriptedPasswordValidatorRequest) SetSchemas(v []EnumgroovyScriptedPasswordValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScriptClass

`func (o *AddGroovyScriptedPasswordValidatorRequest) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddGroovyScriptedPasswordValidatorRequest) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddGroovyScriptedPasswordValidatorRequest) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddGroovyScriptedPasswordValidatorRequest) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddGroovyScriptedPasswordValidatorRequest) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddGroovyScriptedPasswordValidatorRequest) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddGroovyScriptedPasswordValidatorRequest) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddGroovyScriptedPasswordValidatorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddGroovyScriptedPasswordValidatorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddGroovyScriptedPasswordValidatorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddGroovyScriptedPasswordValidatorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddGroovyScriptedPasswordValidatorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddGroovyScriptedPasswordValidatorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddGroovyScriptedPasswordValidatorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetValidatorRequirementDescription

`func (o *AddGroovyScriptedPasswordValidatorRequest) GetValidatorRequirementDescription() string`

GetValidatorRequirementDescription returns the ValidatorRequirementDescription field if non-nil, zero value otherwise.

### GetValidatorRequirementDescriptionOk

`func (o *AddGroovyScriptedPasswordValidatorRequest) GetValidatorRequirementDescriptionOk() (*string, bool)`

GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorRequirementDescription

`func (o *AddGroovyScriptedPasswordValidatorRequest) SetValidatorRequirementDescription(v string)`

SetValidatorRequirementDescription sets ValidatorRequirementDescription field to given value.

### HasValidatorRequirementDescription

`func (o *AddGroovyScriptedPasswordValidatorRequest) HasValidatorRequirementDescription() bool`

HasValidatorRequirementDescription returns a boolean if a field has been set.

### GetValidatorFailureMessage

`func (o *AddGroovyScriptedPasswordValidatorRequest) GetValidatorFailureMessage() string`

GetValidatorFailureMessage returns the ValidatorFailureMessage field if non-nil, zero value otherwise.

### GetValidatorFailureMessageOk

`func (o *AddGroovyScriptedPasswordValidatorRequest) GetValidatorFailureMessageOk() (*string, bool)`

GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorFailureMessage

`func (o *AddGroovyScriptedPasswordValidatorRequest) SetValidatorFailureMessage(v string)`

SetValidatorFailureMessage sets ValidatorFailureMessage field to given value.

### HasValidatorFailureMessage

`func (o *AddGroovyScriptedPasswordValidatorRequest) HasValidatorFailureMessage() bool`

HasValidatorFailureMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

