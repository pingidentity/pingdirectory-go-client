# CryptPasswordStorageSchemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Password Storage Scheme | 
**Schemas** | [**[]EnumcryptPasswordStorageSchemeSchemaUrn**](EnumcryptPasswordStorageSchemeSchemaUrn.md) |  | 
**PasswordEncodingMechanism** | Pointer to [**EnumpasswordStorageSchemePasswordEncodingMechanismProp**](EnumpasswordStorageSchemePasswordEncodingMechanismProp.md) |  | [optional] 
**NumDigestRounds** | Pointer to **int32** | Specifies the number of digest rounds to use for the SHA-2 encodings. This will not be used for the legacy or MD5-based encodings. | [optional] 
**MaxPasswordLength** | Pointer to **int32** | Specifies the maximum allowed length, in bytes, for passwords encoded with this scheme, which can help mitigate denial of service attacks from clients that attempt to bind with very long passwords. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewCryptPasswordStorageSchemeResponse

`func NewCryptPasswordStorageSchemeResponse(id string, schemas []EnumcryptPasswordStorageSchemeSchemaUrn, enabled bool, ) *CryptPasswordStorageSchemeResponse`

NewCryptPasswordStorageSchemeResponse instantiates a new CryptPasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptPasswordStorageSchemeResponseWithDefaults

`func NewCryptPasswordStorageSchemeResponseWithDefaults() *CryptPasswordStorageSchemeResponse`

NewCryptPasswordStorageSchemeResponseWithDefaults instantiates a new CryptPasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CryptPasswordStorageSchemeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CryptPasswordStorageSchemeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CryptPasswordStorageSchemeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *CryptPasswordStorageSchemeResponse) GetSchemas() []EnumcryptPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CryptPasswordStorageSchemeResponse) GetSchemasOk() (*[]EnumcryptPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CryptPasswordStorageSchemeResponse) SetSchemas(v []EnumcryptPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordEncodingMechanism

`func (o *CryptPasswordStorageSchemeResponse) GetPasswordEncodingMechanism() EnumpasswordStorageSchemePasswordEncodingMechanismProp`

GetPasswordEncodingMechanism returns the PasswordEncodingMechanism field if non-nil, zero value otherwise.

### GetPasswordEncodingMechanismOk

`func (o *CryptPasswordStorageSchemeResponse) GetPasswordEncodingMechanismOk() (*EnumpasswordStorageSchemePasswordEncodingMechanismProp, bool)`

GetPasswordEncodingMechanismOk returns a tuple with the PasswordEncodingMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordEncodingMechanism

`func (o *CryptPasswordStorageSchemeResponse) SetPasswordEncodingMechanism(v EnumpasswordStorageSchemePasswordEncodingMechanismProp)`

SetPasswordEncodingMechanism sets PasswordEncodingMechanism field to given value.

### HasPasswordEncodingMechanism

`func (o *CryptPasswordStorageSchemeResponse) HasPasswordEncodingMechanism() bool`

HasPasswordEncodingMechanism returns a boolean if a field has been set.

### GetNumDigestRounds

`func (o *CryptPasswordStorageSchemeResponse) GetNumDigestRounds() int32`

GetNumDigestRounds returns the NumDigestRounds field if non-nil, zero value otherwise.

### GetNumDigestRoundsOk

`func (o *CryptPasswordStorageSchemeResponse) GetNumDigestRoundsOk() (*int32, bool)`

GetNumDigestRoundsOk returns a tuple with the NumDigestRounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDigestRounds

`func (o *CryptPasswordStorageSchemeResponse) SetNumDigestRounds(v int32)`

SetNumDigestRounds sets NumDigestRounds field to given value.

### HasNumDigestRounds

`func (o *CryptPasswordStorageSchemeResponse) HasNumDigestRounds() bool`

HasNumDigestRounds returns a boolean if a field has been set.

### GetMaxPasswordLength

`func (o *CryptPasswordStorageSchemeResponse) GetMaxPasswordLength() int32`

GetMaxPasswordLength returns the MaxPasswordLength field if non-nil, zero value otherwise.

### GetMaxPasswordLengthOk

`func (o *CryptPasswordStorageSchemeResponse) GetMaxPasswordLengthOk() (*int32, bool)`

GetMaxPasswordLengthOk returns a tuple with the MaxPasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPasswordLength

`func (o *CryptPasswordStorageSchemeResponse) SetMaxPasswordLength(v int32)`

SetMaxPasswordLength sets MaxPasswordLength field to given value.

### HasMaxPasswordLength

`func (o *CryptPasswordStorageSchemeResponse) HasMaxPasswordLength() bool`

HasMaxPasswordLength returns a boolean if a field has been set.

### GetDescription

`func (o *CryptPasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CryptPasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CryptPasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CryptPasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CryptPasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CryptPasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CryptPasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *CryptPasswordStorageSchemeResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CryptPasswordStorageSchemeResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CryptPasswordStorageSchemeResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CryptPasswordStorageSchemeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *CryptPasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *CryptPasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *CryptPasswordStorageSchemeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *CryptPasswordStorageSchemeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


