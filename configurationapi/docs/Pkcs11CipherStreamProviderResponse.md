# Pkcs11CipherStreamProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]Enumpkcs11CipherStreamProviderSchemaUrn**](Enumpkcs11CipherStreamProviderSchemaUrn.md) |  | 
**Pkcs11ProviderClass** | Pointer to **string** | The fully-qualified name of the Java security provider class that implements support for interacting with PKCS #11 tokens. | [optional] 
**Pkcs11ProviderConfigurationFile** | Pointer to **string** | The path to the file to use to configure the security provider that implements support for interacting with PKCS #11 tokens. | [optional] 
**KeyStorePin** | Pointer to **string** | The clear-text user PIN needed to interact with the PKCS #11 token. | [optional] 
**KeyStorePinFile** | Pointer to **string** | The path to a file containing the user PIN needed to interact with the PKCS #11 token. The file must exist and must contain exactly one line with a clear-text representation of the PIN. | [optional] 
**KeyStorePinEnvironmentVariable** | Pointer to **string** | The name of an environment variable whose value is the user PIN needed to interact with the PKCS #11 token. The environment variable must be defined and must contain a clear-text representation of the PIN. | [optional] 
**Pkcs11KeyStoreType** | Pointer to **string** | The key store type to use when obtaining an instance of a key store for interacting with a PKCS #11 token. | [optional] 
**SslCertNickname** | **string** | The alias for the certificate in the PKCS #11 token that will be used to wrap the encryption key. The target certificate must exist in the PKCS #11 token, and it must have an RSA key pair because the JVM does not currently provide adequate key wrapping support for elliptic curve key pairs.  If you have also configured the server to use a PKCS #11 token for accessing listener certificates, we strongly recommend that you use a different certificate to protect the contents of the encryption settings database than you use for negotiating TLS sessions with clients. It is imperative that the certificate used by this PKCS11 Cipher Stream Provider remain constant for the life of the provider because if the certificate were to be replaced, then the contents of the encryption settings database could become inaccessible. Unlike with listener certificates used for TLS negotiation that need to be replaced on a regular basis, this PKCS11 Cipher Stream Provider does not consider the validity period for the associated certificate, and it will continue to function even after the certificate has expired.  If you need to rotate the certificate used to protect the server&#39;s encryption settings database, you should first install the desired new certificate in the PKCS #11 token under a different alias. Then, you should create a new instance of this PKCS11 Cipher Stream Provider that is configured to use that certificate, and that also uses a different value for the encryption-metadata-file because the information in that file is tied to the certificate used to generate it. Finally, you will need to update the global configuration so that the encryption-settings-cipher-stream-provider property references the new cipher stream provider rather than this one. The update to the global configuration must be done with the server online so that it can properly re-encrypt the contents of the encryption settings database with the correct key tied to the new certificate. | 
**EncryptionMetadataFile** | **string** | The path to a file that will hold metadata about the encryption performed by this PKCS11 Cipher Stream Provider. | 
**IterationCount** | Pointer to **int64** | The PBKDF2 iteration count that will be used when deriving the encryption key used to protect the encryption settings database. | [optional] 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Cipher Stream Provider | 

## Methods

### NewPkcs11CipherStreamProviderResponse

`func NewPkcs11CipherStreamProviderResponse(schemas []Enumpkcs11CipherStreamProviderSchemaUrn, sslCertNickname string, encryptionMetadataFile string, enabled bool, id string, ) *Pkcs11CipherStreamProviderResponse`

NewPkcs11CipherStreamProviderResponse instantiates a new Pkcs11CipherStreamProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkcs11CipherStreamProviderResponseWithDefaults

`func NewPkcs11CipherStreamProviderResponseWithDefaults() *Pkcs11CipherStreamProviderResponse`

NewPkcs11CipherStreamProviderResponseWithDefaults instantiates a new Pkcs11CipherStreamProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *Pkcs11CipherStreamProviderResponse) GetSchemas() []Enumpkcs11CipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Pkcs11CipherStreamProviderResponse) GetSchemasOk() (*[]Enumpkcs11CipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Pkcs11CipherStreamProviderResponse) SetSchemas(v []Enumpkcs11CipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPkcs11ProviderClass

`func (o *Pkcs11CipherStreamProviderResponse) GetPkcs11ProviderClass() string`

GetPkcs11ProviderClass returns the Pkcs11ProviderClass field if non-nil, zero value otherwise.

### GetPkcs11ProviderClassOk

`func (o *Pkcs11CipherStreamProviderResponse) GetPkcs11ProviderClassOk() (*string, bool)`

GetPkcs11ProviderClassOk returns a tuple with the Pkcs11ProviderClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11ProviderClass

`func (o *Pkcs11CipherStreamProviderResponse) SetPkcs11ProviderClass(v string)`

SetPkcs11ProviderClass sets Pkcs11ProviderClass field to given value.

### HasPkcs11ProviderClass

`func (o *Pkcs11CipherStreamProviderResponse) HasPkcs11ProviderClass() bool`

HasPkcs11ProviderClass returns a boolean if a field has been set.

### GetPkcs11ProviderConfigurationFile

`func (o *Pkcs11CipherStreamProviderResponse) GetPkcs11ProviderConfigurationFile() string`

GetPkcs11ProviderConfigurationFile returns the Pkcs11ProviderConfigurationFile field if non-nil, zero value otherwise.

### GetPkcs11ProviderConfigurationFileOk

`func (o *Pkcs11CipherStreamProviderResponse) GetPkcs11ProviderConfigurationFileOk() (*string, bool)`

GetPkcs11ProviderConfigurationFileOk returns a tuple with the Pkcs11ProviderConfigurationFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11ProviderConfigurationFile

`func (o *Pkcs11CipherStreamProviderResponse) SetPkcs11ProviderConfigurationFile(v string)`

SetPkcs11ProviderConfigurationFile sets Pkcs11ProviderConfigurationFile field to given value.

### HasPkcs11ProviderConfigurationFile

`func (o *Pkcs11CipherStreamProviderResponse) HasPkcs11ProviderConfigurationFile() bool`

HasPkcs11ProviderConfigurationFile returns a boolean if a field has been set.

### GetKeyStorePin

`func (o *Pkcs11CipherStreamProviderResponse) GetKeyStorePin() string`

GetKeyStorePin returns the KeyStorePin field if non-nil, zero value otherwise.

### GetKeyStorePinOk

`func (o *Pkcs11CipherStreamProviderResponse) GetKeyStorePinOk() (*string, bool)`

GetKeyStorePinOk returns a tuple with the KeyStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePin

`func (o *Pkcs11CipherStreamProviderResponse) SetKeyStorePin(v string)`

SetKeyStorePin sets KeyStorePin field to given value.

### HasKeyStorePin

`func (o *Pkcs11CipherStreamProviderResponse) HasKeyStorePin() bool`

HasKeyStorePin returns a boolean if a field has been set.

### GetKeyStorePinFile

`func (o *Pkcs11CipherStreamProviderResponse) GetKeyStorePinFile() string`

GetKeyStorePinFile returns the KeyStorePinFile field if non-nil, zero value otherwise.

### GetKeyStorePinFileOk

`func (o *Pkcs11CipherStreamProviderResponse) GetKeyStorePinFileOk() (*string, bool)`

GetKeyStorePinFileOk returns a tuple with the KeyStorePinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePinFile

`func (o *Pkcs11CipherStreamProviderResponse) SetKeyStorePinFile(v string)`

SetKeyStorePinFile sets KeyStorePinFile field to given value.

### HasKeyStorePinFile

`func (o *Pkcs11CipherStreamProviderResponse) HasKeyStorePinFile() bool`

HasKeyStorePinFile returns a boolean if a field has been set.

### GetKeyStorePinEnvironmentVariable

`func (o *Pkcs11CipherStreamProviderResponse) GetKeyStorePinEnvironmentVariable() string`

GetKeyStorePinEnvironmentVariable returns the KeyStorePinEnvironmentVariable field if non-nil, zero value otherwise.

### GetKeyStorePinEnvironmentVariableOk

`func (o *Pkcs11CipherStreamProviderResponse) GetKeyStorePinEnvironmentVariableOk() (*string, bool)`

GetKeyStorePinEnvironmentVariableOk returns a tuple with the KeyStorePinEnvironmentVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePinEnvironmentVariable

`func (o *Pkcs11CipherStreamProviderResponse) SetKeyStorePinEnvironmentVariable(v string)`

SetKeyStorePinEnvironmentVariable sets KeyStorePinEnvironmentVariable field to given value.

### HasKeyStorePinEnvironmentVariable

`func (o *Pkcs11CipherStreamProviderResponse) HasKeyStorePinEnvironmentVariable() bool`

HasKeyStorePinEnvironmentVariable returns a boolean if a field has been set.

### GetPkcs11KeyStoreType

`func (o *Pkcs11CipherStreamProviderResponse) GetPkcs11KeyStoreType() string`

GetPkcs11KeyStoreType returns the Pkcs11KeyStoreType field if non-nil, zero value otherwise.

### GetPkcs11KeyStoreTypeOk

`func (o *Pkcs11CipherStreamProviderResponse) GetPkcs11KeyStoreTypeOk() (*string, bool)`

GetPkcs11KeyStoreTypeOk returns a tuple with the Pkcs11KeyStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11KeyStoreType

`func (o *Pkcs11CipherStreamProviderResponse) SetPkcs11KeyStoreType(v string)`

SetPkcs11KeyStoreType sets Pkcs11KeyStoreType field to given value.

### HasPkcs11KeyStoreType

`func (o *Pkcs11CipherStreamProviderResponse) HasPkcs11KeyStoreType() bool`

HasPkcs11KeyStoreType returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *Pkcs11CipherStreamProviderResponse) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *Pkcs11CipherStreamProviderResponse) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *Pkcs11CipherStreamProviderResponse) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.


### GetEncryptionMetadataFile

`func (o *Pkcs11CipherStreamProviderResponse) GetEncryptionMetadataFile() string`

GetEncryptionMetadataFile returns the EncryptionMetadataFile field if non-nil, zero value otherwise.

### GetEncryptionMetadataFileOk

`func (o *Pkcs11CipherStreamProviderResponse) GetEncryptionMetadataFileOk() (*string, bool)`

GetEncryptionMetadataFileOk returns a tuple with the EncryptionMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMetadataFile

`func (o *Pkcs11CipherStreamProviderResponse) SetEncryptionMetadataFile(v string)`

SetEncryptionMetadataFile sets EncryptionMetadataFile field to given value.


### GetIterationCount

`func (o *Pkcs11CipherStreamProviderResponse) GetIterationCount() int64`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *Pkcs11CipherStreamProviderResponse) GetIterationCountOk() (*int64, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *Pkcs11CipherStreamProviderResponse) SetIterationCount(v int64)`

SetIterationCount sets IterationCount field to given value.

### HasIterationCount

`func (o *Pkcs11CipherStreamProviderResponse) HasIterationCount() bool`

HasIterationCount returns a boolean if a field has been set.

### GetDescription

`func (o *Pkcs11CipherStreamProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Pkcs11CipherStreamProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Pkcs11CipherStreamProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Pkcs11CipherStreamProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Pkcs11CipherStreamProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Pkcs11CipherStreamProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Pkcs11CipherStreamProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *Pkcs11CipherStreamProviderResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Pkcs11CipherStreamProviderResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Pkcs11CipherStreamProviderResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Pkcs11CipherStreamProviderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *Pkcs11CipherStreamProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *Pkcs11CipherStreamProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *Pkcs11CipherStreamProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *Pkcs11CipherStreamProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *Pkcs11CipherStreamProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pkcs11CipherStreamProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pkcs11CipherStreamProviderResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


