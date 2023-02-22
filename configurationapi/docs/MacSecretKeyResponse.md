# MacSecretKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | Pointer to [**[]EnummacSecretKeySchemaUrn**](EnummacSecretKeySchemaUrn.md) |  | [optional] 
**MacAlgorithmName** | Pointer to **string** | The algorithm name used to generate this MAC key, e.g. HmacMD5, HmacSHA1, HMacSHA256, etc. | [optional] 
**KeyID** | **string** | The unique system-generated identifier for the Secret Key. | 
**IsCompromised** | Pointer to **bool** | If the key is compromised, an administrator may set this flag to immediately trigger the creation of a new secret key. After the new key is generated, the value of this property will be reset to false. | [optional] 
**SymmetricKey** | Pointer to **[]string** | The symmetric key that is used for both encryption of plain text and decryption of cipher text. This stores the secret key for each server instance encrypted with that server&#39;s inter-server certificate. | [optional] 
**KeyLengthBits** | **int32** | The length of the key in bits. | 

## Methods

### NewMacSecretKeyResponse

`func NewMacSecretKeyResponse(keyID string, keyLengthBits int32, ) *MacSecretKeyResponse`

NewMacSecretKeyResponse instantiates a new MacSecretKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacSecretKeyResponseWithDefaults

`func NewMacSecretKeyResponseWithDefaults() *MacSecretKeyResponse`

NewMacSecretKeyResponseWithDefaults instantiates a new MacSecretKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *MacSecretKeyResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MacSecretKeyResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MacSecretKeyResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MacSecretKeyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *MacSecretKeyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *MacSecretKeyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *MacSecretKeyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *MacSecretKeyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *MacSecretKeyResponse) GetSchemas() []EnummacSecretKeySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *MacSecretKeyResponse) GetSchemasOk() (*[]EnummacSecretKeySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *MacSecretKeyResponse) SetSchemas(v []EnummacSecretKeySchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *MacSecretKeyResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetMacAlgorithmName

`func (o *MacSecretKeyResponse) GetMacAlgorithmName() string`

GetMacAlgorithmName returns the MacAlgorithmName field if non-nil, zero value otherwise.

### GetMacAlgorithmNameOk

`func (o *MacSecretKeyResponse) GetMacAlgorithmNameOk() (*string, bool)`

GetMacAlgorithmNameOk returns a tuple with the MacAlgorithmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAlgorithmName

`func (o *MacSecretKeyResponse) SetMacAlgorithmName(v string)`

SetMacAlgorithmName sets MacAlgorithmName field to given value.

### HasMacAlgorithmName

`func (o *MacSecretKeyResponse) HasMacAlgorithmName() bool`

HasMacAlgorithmName returns a boolean if a field has been set.

### GetKeyID

`func (o *MacSecretKeyResponse) GetKeyID() string`

GetKeyID returns the KeyID field if non-nil, zero value otherwise.

### GetKeyIDOk

`func (o *MacSecretKeyResponse) GetKeyIDOk() (*string, bool)`

GetKeyIDOk returns a tuple with the KeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyID

`func (o *MacSecretKeyResponse) SetKeyID(v string)`

SetKeyID sets KeyID field to given value.


### GetIsCompromised

`func (o *MacSecretKeyResponse) GetIsCompromised() bool`

GetIsCompromised returns the IsCompromised field if non-nil, zero value otherwise.

### GetIsCompromisedOk

`func (o *MacSecretKeyResponse) GetIsCompromisedOk() (*bool, bool)`

GetIsCompromisedOk returns a tuple with the IsCompromised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompromised

`func (o *MacSecretKeyResponse) SetIsCompromised(v bool)`

SetIsCompromised sets IsCompromised field to given value.

### HasIsCompromised

`func (o *MacSecretKeyResponse) HasIsCompromised() bool`

HasIsCompromised returns a boolean if a field has been set.

### GetSymmetricKey

`func (o *MacSecretKeyResponse) GetSymmetricKey() []string`

GetSymmetricKey returns the SymmetricKey field if non-nil, zero value otherwise.

### GetSymmetricKeyOk

`func (o *MacSecretKeyResponse) GetSymmetricKeyOk() (*[]string, bool)`

GetSymmetricKeyOk returns a tuple with the SymmetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymmetricKey

`func (o *MacSecretKeyResponse) SetSymmetricKey(v []string)`

SetSymmetricKey sets SymmetricKey field to given value.

### HasSymmetricKey

`func (o *MacSecretKeyResponse) HasSymmetricKey() bool`

HasSymmetricKey returns a boolean if a field has been set.

### GetKeyLengthBits

`func (o *MacSecretKeyResponse) GetKeyLengthBits() int32`

GetKeyLengthBits returns the KeyLengthBits field if non-nil, zero value otherwise.

### GetKeyLengthBitsOk

`func (o *MacSecretKeyResponse) GetKeyLengthBitsOk() (*int32, bool)`

GetKeyLengthBitsOk returns a tuple with the KeyLengthBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLengthBits

`func (o *MacSecretKeyResponse) SetKeyLengthBits(v int32)`

SetKeyLengthBits sets KeyLengthBits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


