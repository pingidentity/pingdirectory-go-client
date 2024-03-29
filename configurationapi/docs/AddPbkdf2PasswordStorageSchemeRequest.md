# AddPbkdf2PasswordStorageSchemeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]Enumpbkdf2PasswordStorageSchemeSchemaUrn**](Enumpbkdf2PasswordStorageSchemeSchemaUrn.md) |  | 
**DigestAlgorithm** | Pointer to [**EnumpasswordStorageSchemeDigestAlgorithmProp**](EnumpasswordStorageSchemeDigestAlgorithmProp.md) |  | [optional] 
**IterationCount** | Pointer to **int64** | Specifies the number of iterations to use when encoding passwords. The value must be greater than or equal to 1000. | [optional] 
**SaltLengthBytes** | Pointer to **int64** | Specifies the number of bytes to use for the generated salt. The value must be greater than or equal to 8. | [optional] 
**DerivedKeyLengthBytes** | Pointer to **int64** | Specifies the number of bytes to use for the derived key. The value must be greater than or equal to 8. | [optional] 
**MaxPasswordLength** | Pointer to **int64** | Specifies the maximum allowed length, in bytes, for passwords encoded with this scheme, which can help mitigate denial of service attacks from clients that attempt to bind with very long passwords. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 
**SchemeName** | **string** | Name of the new Password Storage Scheme | 

## Methods

### NewAddPbkdf2PasswordStorageSchemeRequest

`func NewAddPbkdf2PasswordStorageSchemeRequest(schemas []Enumpbkdf2PasswordStorageSchemeSchemaUrn, enabled bool, schemeName string, ) *AddPbkdf2PasswordStorageSchemeRequest`

NewAddPbkdf2PasswordStorageSchemeRequest instantiates a new AddPbkdf2PasswordStorageSchemeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPbkdf2PasswordStorageSchemeRequestWithDefaults

`func NewAddPbkdf2PasswordStorageSchemeRequestWithDefaults() *AddPbkdf2PasswordStorageSchemeRequest`

NewAddPbkdf2PasswordStorageSchemeRequestWithDefaults instantiates a new AddPbkdf2PasswordStorageSchemeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetSchemas() []Enumpbkdf2PasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetSchemasOk() (*[]Enumpbkdf2PasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPbkdf2PasswordStorageSchemeRequest) SetSchemas(v []Enumpbkdf2PasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDigestAlgorithm

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetDigestAlgorithm() EnumpasswordStorageSchemeDigestAlgorithmProp`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetDigestAlgorithmOk() (*EnumpasswordStorageSchemeDigestAlgorithmProp, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *AddPbkdf2PasswordStorageSchemeRequest) SetDigestAlgorithm(v EnumpasswordStorageSchemeDigestAlgorithmProp)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *AddPbkdf2PasswordStorageSchemeRequest) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetIterationCount

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetIterationCount() int64`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetIterationCountOk() (*int64, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *AddPbkdf2PasswordStorageSchemeRequest) SetIterationCount(v int64)`

SetIterationCount sets IterationCount field to given value.

### HasIterationCount

`func (o *AddPbkdf2PasswordStorageSchemeRequest) HasIterationCount() bool`

HasIterationCount returns a boolean if a field has been set.

### GetSaltLengthBytes

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetSaltLengthBytes() int64`

GetSaltLengthBytes returns the SaltLengthBytes field if non-nil, zero value otherwise.

### GetSaltLengthBytesOk

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetSaltLengthBytesOk() (*int64, bool)`

GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLengthBytes

`func (o *AddPbkdf2PasswordStorageSchemeRequest) SetSaltLengthBytes(v int64)`

SetSaltLengthBytes sets SaltLengthBytes field to given value.

### HasSaltLengthBytes

`func (o *AddPbkdf2PasswordStorageSchemeRequest) HasSaltLengthBytes() bool`

HasSaltLengthBytes returns a boolean if a field has been set.

### GetDerivedKeyLengthBytes

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetDerivedKeyLengthBytes() int64`

GetDerivedKeyLengthBytes returns the DerivedKeyLengthBytes field if non-nil, zero value otherwise.

### GetDerivedKeyLengthBytesOk

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetDerivedKeyLengthBytesOk() (*int64, bool)`

GetDerivedKeyLengthBytesOk returns a tuple with the DerivedKeyLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedKeyLengthBytes

`func (o *AddPbkdf2PasswordStorageSchemeRequest) SetDerivedKeyLengthBytes(v int64)`

SetDerivedKeyLengthBytes sets DerivedKeyLengthBytes field to given value.

### HasDerivedKeyLengthBytes

`func (o *AddPbkdf2PasswordStorageSchemeRequest) HasDerivedKeyLengthBytes() bool`

HasDerivedKeyLengthBytes returns a boolean if a field has been set.

### GetMaxPasswordLength

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetMaxPasswordLength() int64`

GetMaxPasswordLength returns the MaxPasswordLength field if non-nil, zero value otherwise.

### GetMaxPasswordLengthOk

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetMaxPasswordLengthOk() (*int64, bool)`

GetMaxPasswordLengthOk returns a tuple with the MaxPasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPasswordLength

`func (o *AddPbkdf2PasswordStorageSchemeRequest) SetMaxPasswordLength(v int64)`

SetMaxPasswordLength sets MaxPasswordLength field to given value.

### HasMaxPasswordLength

`func (o *AddPbkdf2PasswordStorageSchemeRequest) HasMaxPasswordLength() bool`

HasMaxPasswordLength returns a boolean if a field has been set.

### GetDescription

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPbkdf2PasswordStorageSchemeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPbkdf2PasswordStorageSchemeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPbkdf2PasswordStorageSchemeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSchemeName

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetSchemeName() string`

GetSchemeName returns the SchemeName field if non-nil, zero value otherwise.

### GetSchemeNameOk

`func (o *AddPbkdf2PasswordStorageSchemeRequest) GetSchemeNameOk() (*string, bool)`

GetSchemeNameOk returns a tuple with the SchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeName

`func (o *AddPbkdf2PasswordStorageSchemeRequest) SetSchemeName(v string)`

SetSchemeName sets SchemeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


