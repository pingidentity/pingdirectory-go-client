# PassphrasePasswordGeneratorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Password Generator | 
**Schemas** | [**[]EnumpassphrasePasswordGeneratorSchemaUrn**](EnumpassphrasePasswordGeneratorSchemaUrn.md) |  | 
**DictionaryFile** | **string** | The path to the dictionary file that will be used to obtain the words for use in generated passwords. | 
**MinimumPasswordCharacters** | Pointer to **int32** | The minimum number of characters that generated passwords will be required to have. | [optional] 
**MinimumPasswordWords** | Pointer to **int32** | The minimum number of words that must be concatenated in the course of generating a password. | [optional] 
**CapitalizeWords** | Pointer to **bool** | Indicates whether to capitalize each word used in the generated password. | [optional] 
**Description** | Pointer to **string** | A description for this Password Generator | [optional] 
**Enabled** | **bool** | Indicates whether the Password Generator is enabled for use. | 

## Methods

### NewPassphrasePasswordGeneratorResponse

`func NewPassphrasePasswordGeneratorResponse(id string, schemas []EnumpassphrasePasswordGeneratorSchemaUrn, dictionaryFile string, enabled bool, ) *PassphrasePasswordGeneratorResponse`

NewPassphrasePasswordGeneratorResponse instantiates a new PassphrasePasswordGeneratorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassphrasePasswordGeneratorResponseWithDefaults

`func NewPassphrasePasswordGeneratorResponseWithDefaults() *PassphrasePasswordGeneratorResponse`

NewPassphrasePasswordGeneratorResponseWithDefaults instantiates a new PassphrasePasswordGeneratorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PassphrasePasswordGeneratorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PassphrasePasswordGeneratorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PassphrasePasswordGeneratorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *PassphrasePasswordGeneratorResponse) GetSchemas() []EnumpassphrasePasswordGeneratorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PassphrasePasswordGeneratorResponse) GetSchemasOk() (*[]EnumpassphrasePasswordGeneratorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PassphrasePasswordGeneratorResponse) SetSchemas(v []EnumpassphrasePasswordGeneratorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDictionaryFile

`func (o *PassphrasePasswordGeneratorResponse) GetDictionaryFile() string`

GetDictionaryFile returns the DictionaryFile field if non-nil, zero value otherwise.

### GetDictionaryFileOk

`func (o *PassphrasePasswordGeneratorResponse) GetDictionaryFileOk() (*string, bool)`

GetDictionaryFileOk returns a tuple with the DictionaryFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionaryFile

`func (o *PassphrasePasswordGeneratorResponse) SetDictionaryFile(v string)`

SetDictionaryFile sets DictionaryFile field to given value.


### GetMinimumPasswordCharacters

`func (o *PassphrasePasswordGeneratorResponse) GetMinimumPasswordCharacters() int32`

GetMinimumPasswordCharacters returns the MinimumPasswordCharacters field if non-nil, zero value otherwise.

### GetMinimumPasswordCharactersOk

`func (o *PassphrasePasswordGeneratorResponse) GetMinimumPasswordCharactersOk() (*int32, bool)`

GetMinimumPasswordCharactersOk returns a tuple with the MinimumPasswordCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPasswordCharacters

`func (o *PassphrasePasswordGeneratorResponse) SetMinimumPasswordCharacters(v int32)`

SetMinimumPasswordCharacters sets MinimumPasswordCharacters field to given value.

### HasMinimumPasswordCharacters

`func (o *PassphrasePasswordGeneratorResponse) HasMinimumPasswordCharacters() bool`

HasMinimumPasswordCharacters returns a boolean if a field has been set.

### GetMinimumPasswordWords

`func (o *PassphrasePasswordGeneratorResponse) GetMinimumPasswordWords() int32`

GetMinimumPasswordWords returns the MinimumPasswordWords field if non-nil, zero value otherwise.

### GetMinimumPasswordWordsOk

`func (o *PassphrasePasswordGeneratorResponse) GetMinimumPasswordWordsOk() (*int32, bool)`

GetMinimumPasswordWordsOk returns a tuple with the MinimumPasswordWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPasswordWords

`func (o *PassphrasePasswordGeneratorResponse) SetMinimumPasswordWords(v int32)`

SetMinimumPasswordWords sets MinimumPasswordWords field to given value.

### HasMinimumPasswordWords

`func (o *PassphrasePasswordGeneratorResponse) HasMinimumPasswordWords() bool`

HasMinimumPasswordWords returns a boolean if a field has been set.

### GetCapitalizeWords

`func (o *PassphrasePasswordGeneratorResponse) GetCapitalizeWords() bool`

GetCapitalizeWords returns the CapitalizeWords field if non-nil, zero value otherwise.

### GetCapitalizeWordsOk

`func (o *PassphrasePasswordGeneratorResponse) GetCapitalizeWordsOk() (*bool, bool)`

GetCapitalizeWordsOk returns a tuple with the CapitalizeWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalizeWords

`func (o *PassphrasePasswordGeneratorResponse) SetCapitalizeWords(v bool)`

SetCapitalizeWords sets CapitalizeWords field to given value.

### HasCapitalizeWords

`func (o *PassphrasePasswordGeneratorResponse) HasCapitalizeWords() bool`

HasCapitalizeWords returns a boolean if a field has been set.

### GetDescription

`func (o *PassphrasePasswordGeneratorResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PassphrasePasswordGeneratorResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PassphrasePasswordGeneratorResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PassphrasePasswordGeneratorResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PassphrasePasswordGeneratorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PassphrasePasswordGeneratorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PassphrasePasswordGeneratorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


