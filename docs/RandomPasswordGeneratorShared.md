# RandomPasswordGeneratorShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumrandomPasswordGeneratorSchemaUrn**](EnumrandomPasswordGeneratorSchemaUrn.md) |  | 
**PasswordCharacterSet** | **[]string** |  | 
**PasswordFormat** | **string** | Specifies the format to use for the generated password. | 
**Description** | Pointer to **string** | A description for this Password Generator | [optional] 
**Enabled** | **bool** | Indicates whether the Password Generator is enabled for use. | 

## Methods

### NewRandomPasswordGeneratorShared

`func NewRandomPasswordGeneratorShared(schemas []EnumrandomPasswordGeneratorSchemaUrn, passwordCharacterSet []string, passwordFormat string, enabled bool, ) *RandomPasswordGeneratorShared`

NewRandomPasswordGeneratorShared instantiates a new RandomPasswordGeneratorShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRandomPasswordGeneratorSharedWithDefaults

`func NewRandomPasswordGeneratorSharedWithDefaults() *RandomPasswordGeneratorShared`

NewRandomPasswordGeneratorSharedWithDefaults instantiates a new RandomPasswordGeneratorShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *RandomPasswordGeneratorShared) GetSchemas() []EnumrandomPasswordGeneratorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *RandomPasswordGeneratorShared) GetSchemasOk() (*[]EnumrandomPasswordGeneratorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *RandomPasswordGeneratorShared) SetSchemas(v []EnumrandomPasswordGeneratorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordCharacterSet

`func (o *RandomPasswordGeneratorShared) GetPasswordCharacterSet() []string`

GetPasswordCharacterSet returns the PasswordCharacterSet field if non-nil, zero value otherwise.

### GetPasswordCharacterSetOk

`func (o *RandomPasswordGeneratorShared) GetPasswordCharacterSetOk() (*[]string, bool)`

GetPasswordCharacterSetOk returns a tuple with the PasswordCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCharacterSet

`func (o *RandomPasswordGeneratorShared) SetPasswordCharacterSet(v []string)`

SetPasswordCharacterSet sets PasswordCharacterSet field to given value.


### GetPasswordFormat

`func (o *RandomPasswordGeneratorShared) GetPasswordFormat() string`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *RandomPasswordGeneratorShared) GetPasswordFormatOk() (*string, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *RandomPasswordGeneratorShared) SetPasswordFormat(v string)`

SetPasswordFormat sets PasswordFormat field to given value.


### GetDescription

`func (o *RandomPasswordGeneratorShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RandomPasswordGeneratorShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RandomPasswordGeneratorShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RandomPasswordGeneratorShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *RandomPasswordGeneratorShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RandomPasswordGeneratorShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RandomPasswordGeneratorShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


