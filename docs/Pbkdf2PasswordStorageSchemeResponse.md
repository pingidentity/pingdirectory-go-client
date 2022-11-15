# Pbkdf2PasswordStorageSchemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Password Storage Scheme | 
**Schemas** | [**[]Enumpbkdf2PasswordStorageSchemeSchemaUrn**](Enumpbkdf2PasswordStorageSchemeSchemaUrn.md) |  | 
**DigestAlgorithm** | Pointer to [**EnumpasswordStorageSchemeDigestAlgorithmProp**](EnumpasswordStorageSchemeDigestAlgorithmProp.md) |  | [optional] 
**IterationCount** | **int32** | Specifies the number of iterations to use when encoding passwords. The value must be greater than or equal to 1000. | 
**SaltLengthBytes** | **int32** | Specifies the number of bytes to use for the generated salt. The value must be greater than or equal to 8. | 
**DerivedKeyLengthBytes** | **int32** | Specifies the number of bytes to use for the derived key. The value must be greater than or equal to 8. | 
**MaxPasswordLength** | Pointer to **int32** | Specifies the maximum allowed length, in bytes, for passwords encoded with this scheme, which can help mitigate denial of service attacks from clients that attempt to bind with very long passwords. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 

## Methods

### NewPbkdf2PasswordStorageSchemeResponse

`func NewPbkdf2PasswordStorageSchemeResponse(id string, schemas []Enumpbkdf2PasswordStorageSchemeSchemaUrn, iterationCount int32, saltLengthBytes int32, derivedKeyLengthBytes int32, enabled bool, ) *Pbkdf2PasswordStorageSchemeResponse`

NewPbkdf2PasswordStorageSchemeResponse instantiates a new Pbkdf2PasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPbkdf2PasswordStorageSchemeResponseWithDefaults

`func NewPbkdf2PasswordStorageSchemeResponseWithDefaults() *Pbkdf2PasswordStorageSchemeResponse`

NewPbkdf2PasswordStorageSchemeResponseWithDefaults instantiates a new Pbkdf2PasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pbkdf2PasswordStorageSchemeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetSchemas() []Enumpbkdf2PasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetSchemasOk() (*[]Enumpbkdf2PasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Pbkdf2PasswordStorageSchemeResponse) SetSchemas(v []Enumpbkdf2PasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDigestAlgorithm

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetDigestAlgorithm() EnumpasswordStorageSchemeDigestAlgorithmProp`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetDigestAlgorithmOk() (*EnumpasswordStorageSchemeDigestAlgorithmProp, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *Pbkdf2PasswordStorageSchemeResponse) SetDigestAlgorithm(v EnumpasswordStorageSchemeDigestAlgorithmProp)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *Pbkdf2PasswordStorageSchemeResponse) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetIterationCount

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetIterationCount() int32`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetIterationCountOk() (*int32, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *Pbkdf2PasswordStorageSchemeResponse) SetIterationCount(v int32)`

SetIterationCount sets IterationCount field to given value.


### GetSaltLengthBytes

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetSaltLengthBytes() int32`

GetSaltLengthBytes returns the SaltLengthBytes field if non-nil, zero value otherwise.

### GetSaltLengthBytesOk

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetSaltLengthBytesOk() (*int32, bool)`

GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLengthBytes

`func (o *Pbkdf2PasswordStorageSchemeResponse) SetSaltLengthBytes(v int32)`

SetSaltLengthBytes sets SaltLengthBytes field to given value.


### GetDerivedKeyLengthBytes

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetDerivedKeyLengthBytes() int32`

GetDerivedKeyLengthBytes returns the DerivedKeyLengthBytes field if non-nil, zero value otherwise.

### GetDerivedKeyLengthBytesOk

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetDerivedKeyLengthBytesOk() (*int32, bool)`

GetDerivedKeyLengthBytesOk returns a tuple with the DerivedKeyLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedKeyLengthBytes

`func (o *Pbkdf2PasswordStorageSchemeResponse) SetDerivedKeyLengthBytes(v int32)`

SetDerivedKeyLengthBytes sets DerivedKeyLengthBytes field to given value.


### GetMaxPasswordLength

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetMaxPasswordLength() int32`

GetMaxPasswordLength returns the MaxPasswordLength field if non-nil, zero value otherwise.

### GetMaxPasswordLengthOk

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetMaxPasswordLengthOk() (*int32, bool)`

GetMaxPasswordLengthOk returns a tuple with the MaxPasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPasswordLength

`func (o *Pbkdf2PasswordStorageSchemeResponse) SetMaxPasswordLength(v int32)`

SetMaxPasswordLength sets MaxPasswordLength field to given value.

### HasMaxPasswordLength

`func (o *Pbkdf2PasswordStorageSchemeResponse) HasMaxPasswordLength() bool`

HasMaxPasswordLength returns a boolean if a field has been set.

### GetDescription

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Pbkdf2PasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Pbkdf2PasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Pbkdf2PasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Pbkdf2PasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


