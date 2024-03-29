# AddThirdPartyPasswordValidatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumthirdPartyPasswordValidatorSchemaUrn**](EnumthirdPartyPasswordValidatorSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Password Validator. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Password Validator. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Password Validator | [optional] 
**Enabled** | **bool** | Indicates whether the password validator is enabled for use. | 
**ValidatorRequirementDescription** | Pointer to **string** | Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator. | [optional] 
**ValidatorFailureMessage** | Pointer to **string** | Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator. | [optional] 
**ValidatorName** | **string** | Name of the new Password Validator | 

## Methods

### NewAddThirdPartyPasswordValidatorRequest

`func NewAddThirdPartyPasswordValidatorRequest(schemas []EnumthirdPartyPasswordValidatorSchemaUrn, extensionClass string, enabled bool, validatorName string, ) *AddThirdPartyPasswordValidatorRequest`

NewAddThirdPartyPasswordValidatorRequest instantiates a new AddThirdPartyPasswordValidatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddThirdPartyPasswordValidatorRequestWithDefaults

`func NewAddThirdPartyPasswordValidatorRequestWithDefaults() *AddThirdPartyPasswordValidatorRequest`

NewAddThirdPartyPasswordValidatorRequestWithDefaults instantiates a new AddThirdPartyPasswordValidatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddThirdPartyPasswordValidatorRequest) GetSchemas() []EnumthirdPartyPasswordValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddThirdPartyPasswordValidatorRequest) GetSchemasOk() (*[]EnumthirdPartyPasswordValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddThirdPartyPasswordValidatorRequest) SetSchemas(v []EnumthirdPartyPasswordValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *AddThirdPartyPasswordValidatorRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddThirdPartyPasswordValidatorRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddThirdPartyPasswordValidatorRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddThirdPartyPasswordValidatorRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddThirdPartyPasswordValidatorRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddThirdPartyPasswordValidatorRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddThirdPartyPasswordValidatorRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddThirdPartyPasswordValidatorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddThirdPartyPasswordValidatorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddThirdPartyPasswordValidatorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddThirdPartyPasswordValidatorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddThirdPartyPasswordValidatorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddThirdPartyPasswordValidatorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddThirdPartyPasswordValidatorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetValidatorRequirementDescription

`func (o *AddThirdPartyPasswordValidatorRequest) GetValidatorRequirementDescription() string`

GetValidatorRequirementDescription returns the ValidatorRequirementDescription field if non-nil, zero value otherwise.

### GetValidatorRequirementDescriptionOk

`func (o *AddThirdPartyPasswordValidatorRequest) GetValidatorRequirementDescriptionOk() (*string, bool)`

GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorRequirementDescription

`func (o *AddThirdPartyPasswordValidatorRequest) SetValidatorRequirementDescription(v string)`

SetValidatorRequirementDescription sets ValidatorRequirementDescription field to given value.

### HasValidatorRequirementDescription

`func (o *AddThirdPartyPasswordValidatorRequest) HasValidatorRequirementDescription() bool`

HasValidatorRequirementDescription returns a boolean if a field has been set.

### GetValidatorFailureMessage

`func (o *AddThirdPartyPasswordValidatorRequest) GetValidatorFailureMessage() string`

GetValidatorFailureMessage returns the ValidatorFailureMessage field if non-nil, zero value otherwise.

### GetValidatorFailureMessageOk

`func (o *AddThirdPartyPasswordValidatorRequest) GetValidatorFailureMessageOk() (*string, bool)`

GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorFailureMessage

`func (o *AddThirdPartyPasswordValidatorRequest) SetValidatorFailureMessage(v string)`

SetValidatorFailureMessage sets ValidatorFailureMessage field to given value.

### HasValidatorFailureMessage

`func (o *AddThirdPartyPasswordValidatorRequest) HasValidatorFailureMessage() bool`

HasValidatorFailureMessage returns a boolean if a field has been set.

### GetValidatorName

`func (o *AddThirdPartyPasswordValidatorRequest) GetValidatorName() string`

GetValidatorName returns the ValidatorName field if non-nil, zero value otherwise.

### GetValidatorNameOk

`func (o *AddThirdPartyPasswordValidatorRequest) GetValidatorNameOk() (*string, bool)`

GetValidatorNameOk returns a tuple with the ValidatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorName

`func (o *AddThirdPartyPasswordValidatorRequest) SetValidatorName(v string)`

SetValidatorName sets ValidatorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


