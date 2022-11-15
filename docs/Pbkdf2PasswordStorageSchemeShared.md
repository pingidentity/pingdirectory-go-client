# Pbkdf2PasswordStorageSchemeShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]Enumpbkdf2PasswordStorageSchemeSchemaUrn**](Enumpbkdf2PasswordStorageSchemeSchemaUrn.md) |  | 
**DigestAlgorithm** | Pointer to [**EnumpasswordStorageSchemeDigestAlgorithmProp**](EnumpasswordStorageSchemeDigestAlgorithmProp.md) |  | [optional] 
**IterationCount** | **int32** | Specifies the number of iterations to use when encoding passwords. The value must be greater than or equal to 1000. | 
**SaltLengthBytes** | **int32** | Specifies the number of bytes to use for the generated salt. The value must be greater than or equal to 8. | 
**DerivedKeyLengthBytes** | **int32** | Specifies the number of bytes to use for the derived key. The value must be greater than or equal to 8. | 
**MaxPasswordLength** | Pointer to **int32** | Specifies the maximum allowed length, in bytes, for passwords encoded with this scheme, which can help mitigate denial of service attacks from clients that attempt to bind with very long passwords. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 

## Methods

### NewPbkdf2PasswordStorageSchemeShared

`func NewPbkdf2PasswordStorageSchemeShared(schemas []Enumpbkdf2PasswordStorageSchemeSchemaUrn, iterationCount int32, saltLengthBytes int32, derivedKeyLengthBytes int32, enabled bool, ) *Pbkdf2PasswordStorageSchemeShared`

NewPbkdf2PasswordStorageSchemeShared instantiates a new Pbkdf2PasswordStorageSchemeShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPbkdf2PasswordStorageSchemeSharedWithDefaults

`func NewPbkdf2PasswordStorageSchemeSharedWithDefaults() *Pbkdf2PasswordStorageSchemeShared`

NewPbkdf2PasswordStorageSchemeSharedWithDefaults instantiates a new Pbkdf2PasswordStorageSchemeShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *Pbkdf2PasswordStorageSchemeShared) GetSchemas() []Enumpbkdf2PasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Pbkdf2PasswordStorageSchemeShared) GetSchemasOk() (*[]Enumpbkdf2PasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Pbkdf2PasswordStorageSchemeShared) SetSchemas(v []Enumpbkdf2PasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDigestAlgorithm

`func (o *Pbkdf2PasswordStorageSchemeShared) GetDigestAlgorithm() EnumpasswordStorageSchemeDigestAlgorithmProp`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *Pbkdf2PasswordStorageSchemeShared) GetDigestAlgorithmOk() (*EnumpasswordStorageSchemeDigestAlgorithmProp, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *Pbkdf2PasswordStorageSchemeShared) SetDigestAlgorithm(v EnumpasswordStorageSchemeDigestAlgorithmProp)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *Pbkdf2PasswordStorageSchemeShared) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetIterationCount

`func (o *Pbkdf2PasswordStorageSchemeShared) GetIterationCount() int32`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *Pbkdf2PasswordStorageSchemeShared) GetIterationCountOk() (*int32, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *Pbkdf2PasswordStorageSchemeShared) SetIterationCount(v int32)`

SetIterationCount sets IterationCount field to given value.


### GetSaltLengthBytes

`func (o *Pbkdf2PasswordStorageSchemeShared) GetSaltLengthBytes() int32`

GetSaltLengthBytes returns the SaltLengthBytes field if non-nil, zero value otherwise.

### GetSaltLengthBytesOk

`func (o *Pbkdf2PasswordStorageSchemeShared) GetSaltLengthBytesOk() (*int32, bool)`

GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLengthBytes

`func (o *Pbkdf2PasswordStorageSchemeShared) SetSaltLengthBytes(v int32)`

SetSaltLengthBytes sets SaltLengthBytes field to given value.


### GetDerivedKeyLengthBytes

`func (o *Pbkdf2PasswordStorageSchemeShared) GetDerivedKeyLengthBytes() int32`

GetDerivedKeyLengthBytes returns the DerivedKeyLengthBytes field if non-nil, zero value otherwise.

### GetDerivedKeyLengthBytesOk

`func (o *Pbkdf2PasswordStorageSchemeShared) GetDerivedKeyLengthBytesOk() (*int32, bool)`

GetDerivedKeyLengthBytesOk returns a tuple with the DerivedKeyLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedKeyLengthBytes

`func (o *Pbkdf2PasswordStorageSchemeShared) SetDerivedKeyLengthBytes(v int32)`

SetDerivedKeyLengthBytes sets DerivedKeyLengthBytes field to given value.


### GetMaxPasswordLength

`func (o *Pbkdf2PasswordStorageSchemeShared) GetMaxPasswordLength() int32`

GetMaxPasswordLength returns the MaxPasswordLength field if non-nil, zero value otherwise.

### GetMaxPasswordLengthOk

`func (o *Pbkdf2PasswordStorageSchemeShared) GetMaxPasswordLengthOk() (*int32, bool)`

GetMaxPasswordLengthOk returns a tuple with the MaxPasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPasswordLength

`func (o *Pbkdf2PasswordStorageSchemeShared) SetMaxPasswordLength(v int32)`

SetMaxPasswordLength sets MaxPasswordLength field to given value.

### HasMaxPasswordLength

`func (o *Pbkdf2PasswordStorageSchemeShared) HasMaxPasswordLength() bool`

HasMaxPasswordLength returns a boolean if a field has been set.

### GetDescription

`func (o *Pbkdf2PasswordStorageSchemeShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Pbkdf2PasswordStorageSchemeShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Pbkdf2PasswordStorageSchemeShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Pbkdf2PasswordStorageSchemeShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Pbkdf2PasswordStorageSchemeShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Pbkdf2PasswordStorageSchemeShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Pbkdf2PasswordStorageSchemeShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


