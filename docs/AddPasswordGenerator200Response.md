# AddPasswordGenerator200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Password Generator | 
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

### NewAddPasswordGenerator200Response

`func NewAddPasswordGenerator200Response(id string, schemas []EnumthirdPartyPasswordGeneratorSchemaUrn, passwordCharacterSet []string, passwordFormat string, enabled bool, scriptClass string, dictionaryFile string, extensionClass string, ) *AddPasswordGenerator200Response`

NewAddPasswordGenerator200Response instantiates a new AddPasswordGenerator200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPasswordGenerator200ResponseWithDefaults

`func NewAddPasswordGenerator200ResponseWithDefaults() *AddPasswordGenerator200Response`

NewAddPasswordGenerator200ResponseWithDefaults instantiates a new AddPasswordGenerator200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddPasswordGenerator200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddPasswordGenerator200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddPasswordGenerator200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddPasswordGenerator200Response) GetSchemas() []EnumthirdPartyPasswordGeneratorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPasswordGenerator200Response) GetSchemasOk() (*[]EnumthirdPartyPasswordGeneratorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPasswordGenerator200Response) SetSchemas(v []EnumthirdPartyPasswordGeneratorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordCharacterSet

`func (o *AddPasswordGenerator200Response) GetPasswordCharacterSet() []string`

GetPasswordCharacterSet returns the PasswordCharacterSet field if non-nil, zero value otherwise.

### GetPasswordCharacterSetOk

`func (o *AddPasswordGenerator200Response) GetPasswordCharacterSetOk() (*[]string, bool)`

GetPasswordCharacterSetOk returns a tuple with the PasswordCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCharacterSet

`func (o *AddPasswordGenerator200Response) SetPasswordCharacterSet(v []string)`

SetPasswordCharacterSet sets PasswordCharacterSet field to given value.


### GetPasswordFormat

`func (o *AddPasswordGenerator200Response) GetPasswordFormat() string`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *AddPasswordGenerator200Response) GetPasswordFormatOk() (*string, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *AddPasswordGenerator200Response) SetPasswordFormat(v string)`

SetPasswordFormat sets PasswordFormat field to given value.


### GetDescription

`func (o *AddPasswordGenerator200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPasswordGenerator200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPasswordGenerator200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPasswordGenerator200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPasswordGenerator200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPasswordGenerator200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPasswordGenerator200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetScriptClass

`func (o *AddPasswordGenerator200Response) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddPasswordGenerator200Response) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddPasswordGenerator200Response) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddPasswordGenerator200Response) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddPasswordGenerator200Response) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddPasswordGenerator200Response) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddPasswordGenerator200Response) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetDictionaryFile

`func (o *AddPasswordGenerator200Response) GetDictionaryFile() string`

GetDictionaryFile returns the DictionaryFile field if non-nil, zero value otherwise.

### GetDictionaryFileOk

`func (o *AddPasswordGenerator200Response) GetDictionaryFileOk() (*string, bool)`

GetDictionaryFileOk returns a tuple with the DictionaryFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionaryFile

`func (o *AddPasswordGenerator200Response) SetDictionaryFile(v string)`

SetDictionaryFile sets DictionaryFile field to given value.


### GetMinimumPasswordCharacters

`func (o *AddPasswordGenerator200Response) GetMinimumPasswordCharacters() int32`

GetMinimumPasswordCharacters returns the MinimumPasswordCharacters field if non-nil, zero value otherwise.

### GetMinimumPasswordCharactersOk

`func (o *AddPasswordGenerator200Response) GetMinimumPasswordCharactersOk() (*int32, bool)`

GetMinimumPasswordCharactersOk returns a tuple with the MinimumPasswordCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPasswordCharacters

`func (o *AddPasswordGenerator200Response) SetMinimumPasswordCharacters(v int32)`

SetMinimumPasswordCharacters sets MinimumPasswordCharacters field to given value.

### HasMinimumPasswordCharacters

`func (o *AddPasswordGenerator200Response) HasMinimumPasswordCharacters() bool`

HasMinimumPasswordCharacters returns a boolean if a field has been set.

### GetMinimumPasswordWords

`func (o *AddPasswordGenerator200Response) GetMinimumPasswordWords() int32`

GetMinimumPasswordWords returns the MinimumPasswordWords field if non-nil, zero value otherwise.

### GetMinimumPasswordWordsOk

`func (o *AddPasswordGenerator200Response) GetMinimumPasswordWordsOk() (*int32, bool)`

GetMinimumPasswordWordsOk returns a tuple with the MinimumPasswordWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPasswordWords

`func (o *AddPasswordGenerator200Response) SetMinimumPasswordWords(v int32)`

SetMinimumPasswordWords sets MinimumPasswordWords field to given value.

### HasMinimumPasswordWords

`func (o *AddPasswordGenerator200Response) HasMinimumPasswordWords() bool`

HasMinimumPasswordWords returns a boolean if a field has been set.

### GetCapitalizeWords

`func (o *AddPasswordGenerator200Response) GetCapitalizeWords() bool`

GetCapitalizeWords returns the CapitalizeWords field if non-nil, zero value otherwise.

### GetCapitalizeWordsOk

`func (o *AddPasswordGenerator200Response) GetCapitalizeWordsOk() (*bool, bool)`

GetCapitalizeWordsOk returns a tuple with the CapitalizeWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalizeWords

`func (o *AddPasswordGenerator200Response) SetCapitalizeWords(v bool)`

SetCapitalizeWords sets CapitalizeWords field to given value.

### HasCapitalizeWords

`func (o *AddPasswordGenerator200Response) HasCapitalizeWords() bool`

HasCapitalizeWords returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddPasswordGenerator200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddPasswordGenerator200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddPasswordGenerator200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddPasswordGenerator200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddPasswordGenerator200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddPasswordGenerator200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddPasswordGenerator200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


