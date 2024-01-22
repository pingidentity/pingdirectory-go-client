# CipherSecretKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | Pointer to [**[]EnumcipherSecretKeySchemaUrn**](EnumcipherSecretKeySchemaUrn.md) |  | [optional] 
**Id** | **string** | Name of the Cipher Secret Key | 
**CipherTransformationName** | Pointer to **string** | The algorithm name used to produce this cipher, e.g. AES/CBC/PKCS5Padding. | [optional] 
**InitializationVectorLengthBits** | Pointer to **int64** | The initialization vector length of the cipher in bits. | [optional] 
**KeyID** | **string** | The unique system-generated identifier for the Secret Key. | 
**IsCompromised** | Pointer to **bool** | If the key is compromised, an administrator may set this flag to immediately trigger the creation of a new secret key. After the new key is generated, the value of this property will be reset to false. | [optional] 
**SymmetricKey** | Pointer to **[]string** | The symmetric key that is used for both encryption of plain text and decryption of cipher text. This stores the secret key for each server instance encrypted with that server&#39;s inter-server certificate. | [optional] 
**KeyLengthBits** | **int64** | The length of the key in bits. | 

## Methods

### NewCipherSecretKeyResponse

`func NewCipherSecretKeyResponse(id string, keyID string, keyLengthBits int64, ) *CipherSecretKeyResponse`

NewCipherSecretKeyResponse instantiates a new CipherSecretKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCipherSecretKeyResponseWithDefaults

`func NewCipherSecretKeyResponseWithDefaults() *CipherSecretKeyResponse`

NewCipherSecretKeyResponseWithDefaults instantiates a new CipherSecretKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *CipherSecretKeyResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CipherSecretKeyResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CipherSecretKeyResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CipherSecretKeyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *CipherSecretKeyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *CipherSecretKeyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *CipherSecretKeyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *CipherSecretKeyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *CipherSecretKeyResponse) GetSchemas() []EnumcipherSecretKeySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CipherSecretKeyResponse) GetSchemasOk() (*[]EnumcipherSecretKeySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CipherSecretKeyResponse) SetSchemas(v []EnumcipherSecretKeySchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *CipherSecretKeyResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetId

`func (o *CipherSecretKeyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CipherSecretKeyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CipherSecretKeyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCipherTransformationName

`func (o *CipherSecretKeyResponse) GetCipherTransformationName() string`

GetCipherTransformationName returns the CipherTransformationName field if non-nil, zero value otherwise.

### GetCipherTransformationNameOk

`func (o *CipherSecretKeyResponse) GetCipherTransformationNameOk() (*string, bool)`

GetCipherTransformationNameOk returns a tuple with the CipherTransformationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherTransformationName

`func (o *CipherSecretKeyResponse) SetCipherTransformationName(v string)`

SetCipherTransformationName sets CipherTransformationName field to given value.

### HasCipherTransformationName

`func (o *CipherSecretKeyResponse) HasCipherTransformationName() bool`

HasCipherTransformationName returns a boolean if a field has been set.

### GetInitializationVectorLengthBits

`func (o *CipherSecretKeyResponse) GetInitializationVectorLengthBits() int64`

GetInitializationVectorLengthBits returns the InitializationVectorLengthBits field if non-nil, zero value otherwise.

### GetInitializationVectorLengthBitsOk

`func (o *CipherSecretKeyResponse) GetInitializationVectorLengthBitsOk() (*int64, bool)`

GetInitializationVectorLengthBitsOk returns a tuple with the InitializationVectorLengthBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationVectorLengthBits

`func (o *CipherSecretKeyResponse) SetInitializationVectorLengthBits(v int64)`

SetInitializationVectorLengthBits sets InitializationVectorLengthBits field to given value.

### HasInitializationVectorLengthBits

`func (o *CipherSecretKeyResponse) HasInitializationVectorLengthBits() bool`

HasInitializationVectorLengthBits returns a boolean if a field has been set.

### GetKeyID

`func (o *CipherSecretKeyResponse) GetKeyID() string`

GetKeyID returns the KeyID field if non-nil, zero value otherwise.

### GetKeyIDOk

`func (o *CipherSecretKeyResponse) GetKeyIDOk() (*string, bool)`

GetKeyIDOk returns a tuple with the KeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyID

`func (o *CipherSecretKeyResponse) SetKeyID(v string)`

SetKeyID sets KeyID field to given value.


### GetIsCompromised

`func (o *CipherSecretKeyResponse) GetIsCompromised() bool`

GetIsCompromised returns the IsCompromised field if non-nil, zero value otherwise.

### GetIsCompromisedOk

`func (o *CipherSecretKeyResponse) GetIsCompromisedOk() (*bool, bool)`

GetIsCompromisedOk returns a tuple with the IsCompromised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompromised

`func (o *CipherSecretKeyResponse) SetIsCompromised(v bool)`

SetIsCompromised sets IsCompromised field to given value.

### HasIsCompromised

`func (o *CipherSecretKeyResponse) HasIsCompromised() bool`

HasIsCompromised returns a boolean if a field has been set.

### GetSymmetricKey

`func (o *CipherSecretKeyResponse) GetSymmetricKey() []string`

GetSymmetricKey returns the SymmetricKey field if non-nil, zero value otherwise.

### GetSymmetricKeyOk

`func (o *CipherSecretKeyResponse) GetSymmetricKeyOk() (*[]string, bool)`

GetSymmetricKeyOk returns a tuple with the SymmetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymmetricKey

`func (o *CipherSecretKeyResponse) SetSymmetricKey(v []string)`

SetSymmetricKey sets SymmetricKey field to given value.

### HasSymmetricKey

`func (o *CipherSecretKeyResponse) HasSymmetricKey() bool`

HasSymmetricKey returns a boolean if a field has been set.

### GetKeyLengthBits

`func (o *CipherSecretKeyResponse) GetKeyLengthBits() int64`

GetKeyLengthBits returns the KeyLengthBits field if non-nil, zero value otherwise.

### GetKeyLengthBitsOk

`func (o *CipherSecretKeyResponse) GetKeyLengthBitsOk() (*int64, bool)`

GetKeyLengthBitsOk returns a tuple with the KeyLengthBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLengthBits

`func (o *CipherSecretKeyResponse) SetKeyLengthBits(v int64)`

SetKeyLengthBits sets KeyLengthBits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


