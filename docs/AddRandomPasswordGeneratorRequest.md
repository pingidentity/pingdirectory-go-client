# AddRandomPasswordGeneratorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneratorName** | **string** | Name of the new Password Generator | 
**Schemas** | [**[]EnumrandomPasswordGeneratorSchemaUrn**](EnumrandomPasswordGeneratorSchemaUrn.md) |  | 
**PasswordCharacterSet** | **[]string** |  | 
**PasswordFormat** | **string** | Specifies the format to use for the generated password. | 
**Description** | Pointer to **string** | A description for this Password Generator | [optional] 
**Enabled** | **bool** | Indicates whether the Password Generator is enabled for use. | 

## Methods

### NewAddRandomPasswordGeneratorRequest

`func NewAddRandomPasswordGeneratorRequest(generatorName string, schemas []EnumrandomPasswordGeneratorSchemaUrn, passwordCharacterSet []string, passwordFormat string, enabled bool, ) *AddRandomPasswordGeneratorRequest`

NewAddRandomPasswordGeneratorRequest instantiates a new AddRandomPasswordGeneratorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRandomPasswordGeneratorRequestWithDefaults

`func NewAddRandomPasswordGeneratorRequestWithDefaults() *AddRandomPasswordGeneratorRequest`

NewAddRandomPasswordGeneratorRequestWithDefaults instantiates a new AddRandomPasswordGeneratorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneratorName

`func (o *AddRandomPasswordGeneratorRequest) GetGeneratorName() string`

GetGeneratorName returns the GeneratorName field if non-nil, zero value otherwise.

### GetGeneratorNameOk

`func (o *AddRandomPasswordGeneratorRequest) GetGeneratorNameOk() (*string, bool)`

GetGeneratorNameOk returns a tuple with the GeneratorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorName

`func (o *AddRandomPasswordGeneratorRequest) SetGeneratorName(v string)`

SetGeneratorName sets GeneratorName field to given value.


### GetSchemas

`func (o *AddRandomPasswordGeneratorRequest) GetSchemas() []EnumrandomPasswordGeneratorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddRandomPasswordGeneratorRequest) GetSchemasOk() (*[]EnumrandomPasswordGeneratorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddRandomPasswordGeneratorRequest) SetSchemas(v []EnumrandomPasswordGeneratorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordCharacterSet

`func (o *AddRandomPasswordGeneratorRequest) GetPasswordCharacterSet() []string`

GetPasswordCharacterSet returns the PasswordCharacterSet field if non-nil, zero value otherwise.

### GetPasswordCharacterSetOk

`func (o *AddRandomPasswordGeneratorRequest) GetPasswordCharacterSetOk() (*[]string, bool)`

GetPasswordCharacterSetOk returns a tuple with the PasswordCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCharacterSet

`func (o *AddRandomPasswordGeneratorRequest) SetPasswordCharacterSet(v []string)`

SetPasswordCharacterSet sets PasswordCharacterSet field to given value.


### GetPasswordFormat

`func (o *AddRandomPasswordGeneratorRequest) GetPasswordFormat() string`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *AddRandomPasswordGeneratorRequest) GetPasswordFormatOk() (*string, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *AddRandomPasswordGeneratorRequest) SetPasswordFormat(v string)`

SetPasswordFormat sets PasswordFormat field to given value.


### GetDescription

`func (o *AddRandomPasswordGeneratorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddRandomPasswordGeneratorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddRandomPasswordGeneratorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddRandomPasswordGeneratorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddRandomPasswordGeneratorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddRandomPasswordGeneratorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddRandomPasswordGeneratorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


