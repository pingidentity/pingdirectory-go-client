# AddPassphrasePasswordGeneratorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneratorName** | **string** | Name of the new Password Generator | 
**Schemas** | [**[]EnumpassphrasePasswordGeneratorSchemaUrn**](EnumpassphrasePasswordGeneratorSchemaUrn.md) |  | 
**DictionaryFile** | **string** | The path to the dictionary file that will be used to obtain the words for use in generated passwords. | 
**MinimumPasswordCharacters** | Pointer to **int64** | The minimum number of characters that generated passwords will be required to have. | [optional] 
**MinimumPasswordWords** | Pointer to **int64** | The minimum number of words that must be concatenated in the course of generating a password. | [optional] 
**CapitalizeWords** | Pointer to **bool** | Indicates whether to capitalize each word used in the generated password. | [optional] 
**Description** | Pointer to **string** | A description for this Password Generator | [optional] 
**Enabled** | **bool** | Indicates whether the Password Generator is enabled for use. | 

## Methods

### NewAddPassphrasePasswordGeneratorRequest

`func NewAddPassphrasePasswordGeneratorRequest(generatorName string, schemas []EnumpassphrasePasswordGeneratorSchemaUrn, dictionaryFile string, enabled bool, ) *AddPassphrasePasswordGeneratorRequest`

NewAddPassphrasePasswordGeneratorRequest instantiates a new AddPassphrasePasswordGeneratorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPassphrasePasswordGeneratorRequestWithDefaults

`func NewAddPassphrasePasswordGeneratorRequestWithDefaults() *AddPassphrasePasswordGeneratorRequest`

NewAddPassphrasePasswordGeneratorRequestWithDefaults instantiates a new AddPassphrasePasswordGeneratorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneratorName

`func (o *AddPassphrasePasswordGeneratorRequest) GetGeneratorName() string`

GetGeneratorName returns the GeneratorName field if non-nil, zero value otherwise.

### GetGeneratorNameOk

`func (o *AddPassphrasePasswordGeneratorRequest) GetGeneratorNameOk() (*string, bool)`

GetGeneratorNameOk returns a tuple with the GeneratorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorName

`func (o *AddPassphrasePasswordGeneratorRequest) SetGeneratorName(v string)`

SetGeneratorName sets GeneratorName field to given value.


### GetSchemas

`func (o *AddPassphrasePasswordGeneratorRequest) GetSchemas() []EnumpassphrasePasswordGeneratorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPassphrasePasswordGeneratorRequest) GetSchemasOk() (*[]EnumpassphrasePasswordGeneratorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPassphrasePasswordGeneratorRequest) SetSchemas(v []EnumpassphrasePasswordGeneratorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDictionaryFile

`func (o *AddPassphrasePasswordGeneratorRequest) GetDictionaryFile() string`

GetDictionaryFile returns the DictionaryFile field if non-nil, zero value otherwise.

### GetDictionaryFileOk

`func (o *AddPassphrasePasswordGeneratorRequest) GetDictionaryFileOk() (*string, bool)`

GetDictionaryFileOk returns a tuple with the DictionaryFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionaryFile

`func (o *AddPassphrasePasswordGeneratorRequest) SetDictionaryFile(v string)`

SetDictionaryFile sets DictionaryFile field to given value.


### GetMinimumPasswordCharacters

`func (o *AddPassphrasePasswordGeneratorRequest) GetMinimumPasswordCharacters() int64`

GetMinimumPasswordCharacters returns the MinimumPasswordCharacters field if non-nil, zero value otherwise.

### GetMinimumPasswordCharactersOk

`func (o *AddPassphrasePasswordGeneratorRequest) GetMinimumPasswordCharactersOk() (*int64, bool)`

GetMinimumPasswordCharactersOk returns a tuple with the MinimumPasswordCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPasswordCharacters

`func (o *AddPassphrasePasswordGeneratorRequest) SetMinimumPasswordCharacters(v int64)`

SetMinimumPasswordCharacters sets MinimumPasswordCharacters field to given value.

### HasMinimumPasswordCharacters

`func (o *AddPassphrasePasswordGeneratorRequest) HasMinimumPasswordCharacters() bool`

HasMinimumPasswordCharacters returns a boolean if a field has been set.

### GetMinimumPasswordWords

`func (o *AddPassphrasePasswordGeneratorRequest) GetMinimumPasswordWords() int64`

GetMinimumPasswordWords returns the MinimumPasswordWords field if non-nil, zero value otherwise.

### GetMinimumPasswordWordsOk

`func (o *AddPassphrasePasswordGeneratorRequest) GetMinimumPasswordWordsOk() (*int64, bool)`

GetMinimumPasswordWordsOk returns a tuple with the MinimumPasswordWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPasswordWords

`func (o *AddPassphrasePasswordGeneratorRequest) SetMinimumPasswordWords(v int64)`

SetMinimumPasswordWords sets MinimumPasswordWords field to given value.

### HasMinimumPasswordWords

`func (o *AddPassphrasePasswordGeneratorRequest) HasMinimumPasswordWords() bool`

HasMinimumPasswordWords returns a boolean if a field has been set.

### GetCapitalizeWords

`func (o *AddPassphrasePasswordGeneratorRequest) GetCapitalizeWords() bool`

GetCapitalizeWords returns the CapitalizeWords field if non-nil, zero value otherwise.

### GetCapitalizeWordsOk

`func (o *AddPassphrasePasswordGeneratorRequest) GetCapitalizeWordsOk() (*bool, bool)`

GetCapitalizeWordsOk returns a tuple with the CapitalizeWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalizeWords

`func (o *AddPassphrasePasswordGeneratorRequest) SetCapitalizeWords(v bool)`

SetCapitalizeWords sets CapitalizeWords field to given value.

### HasCapitalizeWords

`func (o *AddPassphrasePasswordGeneratorRequest) HasCapitalizeWords() bool`

HasCapitalizeWords returns a boolean if a field has been set.

### GetDescription

`func (o *AddPassphrasePasswordGeneratorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPassphrasePasswordGeneratorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPassphrasePasswordGeneratorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPassphrasePasswordGeneratorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPassphrasePasswordGeneratorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPassphrasePasswordGeneratorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPassphrasePasswordGeneratorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


