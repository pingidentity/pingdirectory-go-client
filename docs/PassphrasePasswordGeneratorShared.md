# PassphrasePasswordGeneratorShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumpassphrasePasswordGeneratorSchemaUrn**](EnumpassphrasePasswordGeneratorSchemaUrn.md) |  | 
**DictionaryFile** | **string** | The path to the dictionary file that will be used to obtain the words for use in generated passwords. | 
**MinimumPasswordCharacters** | Pointer to **int32** | The minimum number of characters that generated passwords will be required to have. | [optional] 
**MinimumPasswordWords** | Pointer to **int32** | The minimum number of words that must be concatenated in the course of generating a password. | [optional] 
**CapitalizeWords** | Pointer to **bool** | Indicates whether to capitalize each word used in the generated password. | [optional] 
**Description** | Pointer to **string** | A description for this Password Generator | [optional] 
**Enabled** | **bool** | Indicates whether the Password Generator is enabled for use. | 

## Methods

### NewPassphrasePasswordGeneratorShared

`func NewPassphrasePasswordGeneratorShared(schemas []EnumpassphrasePasswordGeneratorSchemaUrn, dictionaryFile string, enabled bool, ) *PassphrasePasswordGeneratorShared`

NewPassphrasePasswordGeneratorShared instantiates a new PassphrasePasswordGeneratorShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassphrasePasswordGeneratorSharedWithDefaults

`func NewPassphrasePasswordGeneratorSharedWithDefaults() *PassphrasePasswordGeneratorShared`

NewPassphrasePasswordGeneratorSharedWithDefaults instantiates a new PassphrasePasswordGeneratorShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *PassphrasePasswordGeneratorShared) GetSchemas() []EnumpassphrasePasswordGeneratorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PassphrasePasswordGeneratorShared) GetSchemasOk() (*[]EnumpassphrasePasswordGeneratorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PassphrasePasswordGeneratorShared) SetSchemas(v []EnumpassphrasePasswordGeneratorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDictionaryFile

`func (o *PassphrasePasswordGeneratorShared) GetDictionaryFile() string`

GetDictionaryFile returns the DictionaryFile field if non-nil, zero value otherwise.

### GetDictionaryFileOk

`func (o *PassphrasePasswordGeneratorShared) GetDictionaryFileOk() (*string, bool)`

GetDictionaryFileOk returns a tuple with the DictionaryFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionaryFile

`func (o *PassphrasePasswordGeneratorShared) SetDictionaryFile(v string)`

SetDictionaryFile sets DictionaryFile field to given value.


### GetMinimumPasswordCharacters

`func (o *PassphrasePasswordGeneratorShared) GetMinimumPasswordCharacters() int32`

GetMinimumPasswordCharacters returns the MinimumPasswordCharacters field if non-nil, zero value otherwise.

### GetMinimumPasswordCharactersOk

`func (o *PassphrasePasswordGeneratorShared) GetMinimumPasswordCharactersOk() (*int32, bool)`

GetMinimumPasswordCharactersOk returns a tuple with the MinimumPasswordCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPasswordCharacters

`func (o *PassphrasePasswordGeneratorShared) SetMinimumPasswordCharacters(v int32)`

SetMinimumPasswordCharacters sets MinimumPasswordCharacters field to given value.

### HasMinimumPasswordCharacters

`func (o *PassphrasePasswordGeneratorShared) HasMinimumPasswordCharacters() bool`

HasMinimumPasswordCharacters returns a boolean if a field has been set.

### GetMinimumPasswordWords

`func (o *PassphrasePasswordGeneratorShared) GetMinimumPasswordWords() int32`

GetMinimumPasswordWords returns the MinimumPasswordWords field if non-nil, zero value otherwise.

### GetMinimumPasswordWordsOk

`func (o *PassphrasePasswordGeneratorShared) GetMinimumPasswordWordsOk() (*int32, bool)`

GetMinimumPasswordWordsOk returns a tuple with the MinimumPasswordWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPasswordWords

`func (o *PassphrasePasswordGeneratorShared) SetMinimumPasswordWords(v int32)`

SetMinimumPasswordWords sets MinimumPasswordWords field to given value.

### HasMinimumPasswordWords

`func (o *PassphrasePasswordGeneratorShared) HasMinimumPasswordWords() bool`

HasMinimumPasswordWords returns a boolean if a field has been set.

### GetCapitalizeWords

`func (o *PassphrasePasswordGeneratorShared) GetCapitalizeWords() bool`

GetCapitalizeWords returns the CapitalizeWords field if non-nil, zero value otherwise.

### GetCapitalizeWordsOk

`func (o *PassphrasePasswordGeneratorShared) GetCapitalizeWordsOk() (*bool, bool)`

GetCapitalizeWordsOk returns a tuple with the CapitalizeWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalizeWords

`func (o *PassphrasePasswordGeneratorShared) SetCapitalizeWords(v bool)`

SetCapitalizeWords sets CapitalizeWords field to given value.

### HasCapitalizeWords

`func (o *PassphrasePasswordGeneratorShared) HasCapitalizeWords() bool`

HasCapitalizeWords returns a boolean if a field has been set.

### GetDescription

`func (o *PassphrasePasswordGeneratorShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PassphrasePasswordGeneratorShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PassphrasePasswordGeneratorShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PassphrasePasswordGeneratorShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PassphrasePasswordGeneratorShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PassphrasePasswordGeneratorShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PassphrasePasswordGeneratorShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


