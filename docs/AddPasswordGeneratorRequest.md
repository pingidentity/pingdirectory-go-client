# AddPasswordGeneratorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneratorName** | **string** | Name of the new Password Generator | 
**Schemas** | [**[]EnumthirdPartyPasswordGeneratorSchemaUrn**](EnumthirdPartyPasswordGeneratorSchemaUrn.md) |  | 
**PasswordCharacterSet** | **[]string** |  | 
**PasswordFormat** | **string** | Specifies the format to use for the generated password. | 
**Description** | Pointer to **string** | A description for this Password Generator | [optional] 
**Enabled** | **bool** | Indicates whether the Password Generator is enabled for use. | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Password Generator. | 
**ScriptArgument** | Pointer to **[]string** |  | [optional] 
**DictionaryFile** | **string** | The path to the dictionary file that will be used to obtain the words for use in generated passwords. | 
**MinimumPasswordCharacters** | Pointer to **int32** | The minimum number of characters that generated passwords will be required to have. | [optional] 
**MinimumPasswordWords** | Pointer to **int32** | The minimum number of words that must be concatenated in the course of generating a password. | [optional] 
**CapitalizeWords** | Pointer to **bool** | Indicates whether to capitalize each word used in the generated password. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Password Generator. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAddPasswordGeneratorRequest

`func NewAddPasswordGeneratorRequest(generatorName string, schemas []EnumthirdPartyPasswordGeneratorSchemaUrn, passwordCharacterSet []string, passwordFormat string, enabled bool, scriptClass string, dictionaryFile string, extensionClass string, ) *AddPasswordGeneratorRequest`

NewAddPasswordGeneratorRequest instantiates a new AddPasswordGeneratorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPasswordGeneratorRequestWithDefaults

`func NewAddPasswordGeneratorRequestWithDefaults() *AddPasswordGeneratorRequest`

NewAddPasswordGeneratorRequestWithDefaults instantiates a new AddPasswordGeneratorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneratorName

`func (o *AddPasswordGeneratorRequest) GetGeneratorName() string`

GetGeneratorName returns the GeneratorName field if non-nil, zero value otherwise.

### GetGeneratorNameOk

`func (o *AddPasswordGeneratorRequest) GetGeneratorNameOk() (*string, bool)`

GetGeneratorNameOk returns a tuple with the GeneratorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorName

`func (o *AddPasswordGeneratorRequest) SetGeneratorName(v string)`

SetGeneratorName sets GeneratorName field to given value.


### GetSchemas

`func (o *AddPasswordGeneratorRequest) GetSchemas() []EnumthirdPartyPasswordGeneratorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPasswordGeneratorRequest) GetSchemasOk() (*[]EnumthirdPartyPasswordGeneratorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPasswordGeneratorRequest) SetSchemas(v []EnumthirdPartyPasswordGeneratorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordCharacterSet

`func (o *AddPasswordGeneratorRequest) GetPasswordCharacterSet() []string`

GetPasswordCharacterSet returns the PasswordCharacterSet field if non-nil, zero value otherwise.

### GetPasswordCharacterSetOk

`func (o *AddPasswordGeneratorRequest) GetPasswordCharacterSetOk() (*[]string, bool)`

GetPasswordCharacterSetOk returns a tuple with the PasswordCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCharacterSet

`func (o *AddPasswordGeneratorRequest) SetPasswordCharacterSet(v []string)`

SetPasswordCharacterSet sets PasswordCharacterSet field to given value.


### GetPasswordFormat

`func (o *AddPasswordGeneratorRequest) GetPasswordFormat() string`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *AddPasswordGeneratorRequest) GetPasswordFormatOk() (*string, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *AddPasswordGeneratorRequest) SetPasswordFormat(v string)`

SetPasswordFormat sets PasswordFormat field to given value.


### GetDescription

`func (o *AddPasswordGeneratorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPasswordGeneratorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPasswordGeneratorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPasswordGeneratorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPasswordGeneratorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPasswordGeneratorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPasswordGeneratorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetScriptClass

`func (o *AddPasswordGeneratorRequest) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddPasswordGeneratorRequest) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddPasswordGeneratorRequest) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddPasswordGeneratorRequest) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddPasswordGeneratorRequest) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddPasswordGeneratorRequest) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddPasswordGeneratorRequest) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDictionaryFile

`func (o *AddPasswordGeneratorRequest) GetDictionaryFile() string`

GetDictionaryFile returns the DictionaryFile field if non-nil, zero value otherwise.

### GetDictionaryFileOk

`func (o *AddPasswordGeneratorRequest) GetDictionaryFileOk() (*string, bool)`

GetDictionaryFileOk returns a tuple with the DictionaryFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionaryFile

`func (o *AddPasswordGeneratorRequest) SetDictionaryFile(v string)`

SetDictionaryFile sets DictionaryFile field to given value.


### GetMinimumPasswordCharacters

`func (o *AddPasswordGeneratorRequest) GetMinimumPasswordCharacters() int32`

GetMinimumPasswordCharacters returns the MinimumPasswordCharacters field if non-nil, zero value otherwise.

### GetMinimumPasswordCharactersOk

`func (o *AddPasswordGeneratorRequest) GetMinimumPasswordCharactersOk() (*int32, bool)`

GetMinimumPasswordCharactersOk returns a tuple with the MinimumPasswordCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPasswordCharacters

`func (o *AddPasswordGeneratorRequest) SetMinimumPasswordCharacters(v int32)`

SetMinimumPasswordCharacters sets MinimumPasswordCharacters field to given value.

### HasMinimumPasswordCharacters

`func (o *AddPasswordGeneratorRequest) HasMinimumPasswordCharacters() bool`

HasMinimumPasswordCharacters returns a boolean if a field has been set.

### GetMinimumPasswordWords

`func (o *AddPasswordGeneratorRequest) GetMinimumPasswordWords() int32`

GetMinimumPasswordWords returns the MinimumPasswordWords field if non-nil, zero value otherwise.

### GetMinimumPasswordWordsOk

`func (o *AddPasswordGeneratorRequest) GetMinimumPasswordWordsOk() (*int32, bool)`

GetMinimumPasswordWordsOk returns a tuple with the MinimumPasswordWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPasswordWords

`func (o *AddPasswordGeneratorRequest) SetMinimumPasswordWords(v int32)`

SetMinimumPasswordWords sets MinimumPasswordWords field to given value.

### HasMinimumPasswordWords

`func (o *AddPasswordGeneratorRequest) HasMinimumPasswordWords() bool`

HasMinimumPasswordWords returns a boolean if a field has been set.

### GetCapitalizeWords

`func (o *AddPasswordGeneratorRequest) GetCapitalizeWords() bool`

GetCapitalizeWords returns the CapitalizeWords field if non-nil, zero value otherwise.

### GetCapitalizeWordsOk

`func (o *AddPasswordGeneratorRequest) GetCapitalizeWordsOk() (*bool, bool)`

GetCapitalizeWordsOk returns a tuple with the CapitalizeWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalizeWords

`func (o *AddPasswordGeneratorRequest) SetCapitalizeWords(v bool)`

SetCapitalizeWords sets CapitalizeWords field to given value.

### HasCapitalizeWords

`func (o *AddPasswordGeneratorRequest) HasCapitalizeWords() bool`

HasCapitalizeWords returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddPasswordGeneratorRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddPasswordGeneratorRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddPasswordGeneratorRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddPasswordGeneratorRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddPasswordGeneratorRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddPasswordGeneratorRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddPasswordGeneratorRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


