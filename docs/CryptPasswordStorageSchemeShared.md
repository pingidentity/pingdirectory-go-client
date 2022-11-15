# CryptPasswordStorageSchemeShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumcryptPasswordStorageSchemeSchemaUrn**](EnumcryptPasswordStorageSchemeSchemaUrn.md) |  | 
**PasswordEncodingMechanism** | Pointer to [**EnumpasswordStorageSchemePasswordEncodingMechanismProp**](EnumpasswordStorageSchemePasswordEncodingMechanismProp.md) |  | [optional] 
**NumDigestRounds** | Pointer to **int32** | Specifies the number of digest rounds to use for the SHA-2 encodings. This will not be used for the legacy or MD5-based encodings. | [optional] 
**MaxPasswordLength** | Pointer to **int32** | Specifies the maximum allowed length, in bytes, for passwords encoded with this scheme, which can help mitigate denial of service attacks from clients that attempt to bind with very long passwords. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 

## Methods

### NewCryptPasswordStorageSchemeShared

`func NewCryptPasswordStorageSchemeShared(schemas []EnumcryptPasswordStorageSchemeSchemaUrn, enabled bool, ) *CryptPasswordStorageSchemeShared`

NewCryptPasswordStorageSchemeShared instantiates a new CryptPasswordStorageSchemeShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptPasswordStorageSchemeSharedWithDefaults

`func NewCryptPasswordStorageSchemeSharedWithDefaults() *CryptPasswordStorageSchemeShared`

NewCryptPasswordStorageSchemeSharedWithDefaults instantiates a new CryptPasswordStorageSchemeShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *CryptPasswordStorageSchemeShared) GetSchemas() []EnumcryptPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CryptPasswordStorageSchemeShared) GetSchemasOk() (*[]EnumcryptPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CryptPasswordStorageSchemeShared) SetSchemas(v []EnumcryptPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordEncodingMechanism

`func (o *CryptPasswordStorageSchemeShared) GetPasswordEncodingMechanism() EnumpasswordStorageSchemePasswordEncodingMechanismProp`

GetPasswordEncodingMechanism returns the PasswordEncodingMechanism field if non-nil, zero value otherwise.

### GetPasswordEncodingMechanismOk

`func (o *CryptPasswordStorageSchemeShared) GetPasswordEncodingMechanismOk() (*EnumpasswordStorageSchemePasswordEncodingMechanismProp, bool)`

GetPasswordEncodingMechanismOk returns a tuple with the PasswordEncodingMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordEncodingMechanism

`func (o *CryptPasswordStorageSchemeShared) SetPasswordEncodingMechanism(v EnumpasswordStorageSchemePasswordEncodingMechanismProp)`

SetPasswordEncodingMechanism sets PasswordEncodingMechanism field to given value.

### HasPasswordEncodingMechanism

`func (o *CryptPasswordStorageSchemeShared) HasPasswordEncodingMechanism() bool`

HasPasswordEncodingMechanism returns a boolean if a field has been set.

### GetNumDigestRounds

`func (o *CryptPasswordStorageSchemeShared) GetNumDigestRounds() int32`

GetNumDigestRounds returns the NumDigestRounds field if non-nil, zero value otherwise.

### GetNumDigestRoundsOk

`func (o *CryptPasswordStorageSchemeShared) GetNumDigestRoundsOk() (*int32, bool)`

GetNumDigestRoundsOk returns a tuple with the NumDigestRounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDigestRounds

`func (o *CryptPasswordStorageSchemeShared) SetNumDigestRounds(v int32)`

SetNumDigestRounds sets NumDigestRounds field to given value.

### HasNumDigestRounds

`func (o *CryptPasswordStorageSchemeShared) HasNumDigestRounds() bool`

HasNumDigestRounds returns a boolean if a field has been set.

### GetMaxPasswordLength

`func (o *CryptPasswordStorageSchemeShared) GetMaxPasswordLength() int32`

GetMaxPasswordLength returns the MaxPasswordLength field if non-nil, zero value otherwise.

### GetMaxPasswordLengthOk

`func (o *CryptPasswordStorageSchemeShared) GetMaxPasswordLengthOk() (*int32, bool)`

GetMaxPasswordLengthOk returns a tuple with the MaxPasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPasswordLength

`func (o *CryptPasswordStorageSchemeShared) SetMaxPasswordLength(v int32)`

SetMaxPasswordLength sets MaxPasswordLength field to given value.

### HasMaxPasswordLength

`func (o *CryptPasswordStorageSchemeShared) HasMaxPasswordLength() bool`

HasMaxPasswordLength returns a boolean if a field has been set.

### GetDescription

`func (o *CryptPasswordStorageSchemeShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CryptPasswordStorageSchemeShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CryptPasswordStorageSchemeShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CryptPasswordStorageSchemeShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CryptPasswordStorageSchemeShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CryptPasswordStorageSchemeShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CryptPasswordStorageSchemeShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


