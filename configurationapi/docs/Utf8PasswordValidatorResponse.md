# Utf8PasswordValidatorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]Enumutf8PasswordValidatorSchemaUrn**](Enumutf8PasswordValidatorSchemaUrn.md) |  | 
**AllowNonAsciiCharacters** | Pointer to **bool** | Indicates whether passwords will be allowed to include characters from outside the ASCII character set. | [optional] 
**AllowUnknownCharacters** | Pointer to **bool** | Indicates whether passwords will be allowed to include characters that are not recognized by the JVM&#39;s Unicode support. | [optional] 
**AllowedCharacterType** | Pointer to [**[]EnumpasswordValidatorAllowedCharacterTypeProp**](EnumpasswordValidatorAllowedCharacterTypeProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Password Validator | [optional] 
**Enabled** | **bool** | Indicates whether the password validator is enabled for use. | 
**ValidatorRequirementDescription** | Pointer to **string** | Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator. | [optional] 
**ValidatorFailureMessage** | Pointer to **string** | Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Password Validator | 

## Methods

### NewUtf8PasswordValidatorResponse

`func NewUtf8PasswordValidatorResponse(schemas []Enumutf8PasswordValidatorSchemaUrn, enabled bool, id string, ) *Utf8PasswordValidatorResponse`

NewUtf8PasswordValidatorResponse instantiates a new Utf8PasswordValidatorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtf8PasswordValidatorResponseWithDefaults

`func NewUtf8PasswordValidatorResponseWithDefaults() *Utf8PasswordValidatorResponse`

NewUtf8PasswordValidatorResponseWithDefaults instantiates a new Utf8PasswordValidatorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *Utf8PasswordValidatorResponse) GetSchemas() []Enumutf8PasswordValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Utf8PasswordValidatorResponse) GetSchemasOk() (*[]Enumutf8PasswordValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Utf8PasswordValidatorResponse) SetSchemas(v []Enumutf8PasswordValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllowNonAsciiCharacters

`func (o *Utf8PasswordValidatorResponse) GetAllowNonAsciiCharacters() bool`

GetAllowNonAsciiCharacters returns the AllowNonAsciiCharacters field if non-nil, zero value otherwise.

### GetAllowNonAsciiCharactersOk

`func (o *Utf8PasswordValidatorResponse) GetAllowNonAsciiCharactersOk() (*bool, bool)`

GetAllowNonAsciiCharactersOk returns a tuple with the AllowNonAsciiCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNonAsciiCharacters

`func (o *Utf8PasswordValidatorResponse) SetAllowNonAsciiCharacters(v bool)`

SetAllowNonAsciiCharacters sets AllowNonAsciiCharacters field to given value.

### HasAllowNonAsciiCharacters

`func (o *Utf8PasswordValidatorResponse) HasAllowNonAsciiCharacters() bool`

HasAllowNonAsciiCharacters returns a boolean if a field has been set.

### GetAllowUnknownCharacters

`func (o *Utf8PasswordValidatorResponse) GetAllowUnknownCharacters() bool`

GetAllowUnknownCharacters returns the AllowUnknownCharacters field if non-nil, zero value otherwise.

### GetAllowUnknownCharactersOk

`func (o *Utf8PasswordValidatorResponse) GetAllowUnknownCharactersOk() (*bool, bool)`

GetAllowUnknownCharactersOk returns a tuple with the AllowUnknownCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnknownCharacters

`func (o *Utf8PasswordValidatorResponse) SetAllowUnknownCharacters(v bool)`

SetAllowUnknownCharacters sets AllowUnknownCharacters field to given value.

### HasAllowUnknownCharacters

`func (o *Utf8PasswordValidatorResponse) HasAllowUnknownCharacters() bool`

HasAllowUnknownCharacters returns a boolean if a field has been set.

### GetAllowedCharacterType

`func (o *Utf8PasswordValidatorResponse) GetAllowedCharacterType() []EnumpasswordValidatorAllowedCharacterTypeProp`

GetAllowedCharacterType returns the AllowedCharacterType field if non-nil, zero value otherwise.

### GetAllowedCharacterTypeOk

`func (o *Utf8PasswordValidatorResponse) GetAllowedCharacterTypeOk() (*[]EnumpasswordValidatorAllowedCharacterTypeProp, bool)`

GetAllowedCharacterTypeOk returns a tuple with the AllowedCharacterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCharacterType

`func (o *Utf8PasswordValidatorResponse) SetAllowedCharacterType(v []EnumpasswordValidatorAllowedCharacterTypeProp)`

SetAllowedCharacterType sets AllowedCharacterType field to given value.

### HasAllowedCharacterType

`func (o *Utf8PasswordValidatorResponse) HasAllowedCharacterType() bool`

HasAllowedCharacterType returns a boolean if a field has been set.

### GetDescription

`func (o *Utf8PasswordValidatorResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Utf8PasswordValidatorResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Utf8PasswordValidatorResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Utf8PasswordValidatorResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Utf8PasswordValidatorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Utf8PasswordValidatorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Utf8PasswordValidatorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetValidatorRequirementDescription

`func (o *Utf8PasswordValidatorResponse) GetValidatorRequirementDescription() string`

GetValidatorRequirementDescription returns the ValidatorRequirementDescription field if non-nil, zero value otherwise.

### GetValidatorRequirementDescriptionOk

`func (o *Utf8PasswordValidatorResponse) GetValidatorRequirementDescriptionOk() (*string, bool)`

GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorRequirementDescription

`func (o *Utf8PasswordValidatorResponse) SetValidatorRequirementDescription(v string)`

SetValidatorRequirementDescription sets ValidatorRequirementDescription field to given value.

### HasValidatorRequirementDescription

`func (o *Utf8PasswordValidatorResponse) HasValidatorRequirementDescription() bool`

HasValidatorRequirementDescription returns a boolean if a field has been set.

### GetValidatorFailureMessage

`func (o *Utf8PasswordValidatorResponse) GetValidatorFailureMessage() string`

GetValidatorFailureMessage returns the ValidatorFailureMessage field if non-nil, zero value otherwise.

### GetValidatorFailureMessageOk

`func (o *Utf8PasswordValidatorResponse) GetValidatorFailureMessageOk() (*string, bool)`

GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorFailureMessage

`func (o *Utf8PasswordValidatorResponse) SetValidatorFailureMessage(v string)`

SetValidatorFailureMessage sets ValidatorFailureMessage field to given value.

### HasValidatorFailureMessage

`func (o *Utf8PasswordValidatorResponse) HasValidatorFailureMessage() bool`

HasValidatorFailureMessage returns a boolean if a field has been set.

### GetMeta

`func (o *Utf8PasswordValidatorResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Utf8PasswordValidatorResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Utf8PasswordValidatorResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Utf8PasswordValidatorResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *Utf8PasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *Utf8PasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *Utf8PasswordValidatorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *Utf8PasswordValidatorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *Utf8PasswordValidatorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Utf8PasswordValidatorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Utf8PasswordValidatorResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


