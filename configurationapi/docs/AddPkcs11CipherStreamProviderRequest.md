# AddPkcs11CipherStreamProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Name of the new Cipher Stream Provider | 
**Schemas** | [**[]Enumpkcs11CipherStreamProviderSchemaUrn**](Enumpkcs11CipherStreamProviderSchemaUrn.md) |  | 
**Pkcs11ProviderClass** | Pointer to **string** | The fully-qualified name of the Java security provider class that implements support for interacting with PKCS #11 tokens. | [optional] 
**Pkcs11ProviderConfigurationFile** | Pointer to **string** | The path to the file to use to configure the security provider that implements support for interacting with PKCS #11 tokens. | [optional] 
**KeyStorePin** | Pointer to **string** | The clear-text user PIN needed to interact with the PKCS #11 token. | [optional] 
**KeyStorePinFile** | Pointer to **string** | The path to a file containing the user PIN needed to interact with the PKCS #11 token. The file must exist and must contain exactly one line with a clear-text representation of the PIN. | [optional] 
**KeyStorePinEnvironmentVariable** | Pointer to **string** | The name of an environment variable whose value is the user PIN needed to interact with the PKCS #11 token. The environment variable must be defined and must contain a clear-text representation of the PIN. | [optional] 
**Pkcs11KeyStoreType** | Pointer to **string** | The key store type to use when obtaining an instance of a key store for interacting with a PKCS #11 token. | [optional] 
**SslCertNickname** | **string** | The alias for the certificate in the PKCS #11 token that will be used to wrap the encryption key. The target certificate must exist in the PKCS #11 token, and it must have an RSA key pair because the JVM does not currently provide adequate key wrapping support for elliptic curve key pairs.  If you have also configured the server to use a PKCS #11 token for accessing listener certificates, we strongly recommend that you use a different certificate to protect the contents of the encryption settings database than you use for negotiating TLS sessions with clients. It is imperative that the certificate used by this PKCS11 Cipher Stream Provider remain constant for the life of the provider because if the certificate were to be replaced, then the contents of the encryption settings database could become inaccessible. Unlike with listener certificates used for TLS negotiation that need to be replaced on a regular basis, this PKCS11 Cipher Stream Provider does not consider the validity period for the associated certificate, and it will continue to function even after the certificate has expired.  If you need to rotate the certificate used to protect the server&#39;s encryption settings database, you should first install the desired new certificate in the PKCS #11 token under a different alias. Then, you should create a new instance of this PKCS11 Cipher Stream Provider that is configured to use that certificate, and that also uses a different value for the encryption-metadata-file because the information in that file is tied to the certificate used to generate it. Finally, you will need to update the global configuration so that the encryption-settings-cipher-stream-provider property references the new cipher stream provider rather than this one. The update to the global configuration must be done with the server online so that it can properly re-encrypt the contents of the encryption settings database with the correct key tied to the new certificate. | 
**EncryptionMetadataFile** | Pointer to **string** | The path to a file that will hold metadata about the encryption performed by this PKCS11 Cipher Stream Provider. | [optional] 
**IterationCount** | Pointer to **int64** | The PBKDF2 iteration count that will be used when deriving the encryption key used to protect the encryption settings database. | [optional] 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 

## Methods

### NewAddPkcs11CipherStreamProviderRequest

`func NewAddPkcs11CipherStreamProviderRequest(providerName string, schemas []Enumpkcs11CipherStreamProviderSchemaUrn, sslCertNickname string, enabled bool, ) *AddPkcs11CipherStreamProviderRequest`

NewAddPkcs11CipherStreamProviderRequest instantiates a new AddPkcs11CipherStreamProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPkcs11CipherStreamProviderRequestWithDefaults

`func NewAddPkcs11CipherStreamProviderRequestWithDefaults() *AddPkcs11CipherStreamProviderRequest`

NewAddPkcs11CipherStreamProviderRequestWithDefaults instantiates a new AddPkcs11CipherStreamProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *AddPkcs11CipherStreamProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddPkcs11CipherStreamProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddPkcs11CipherStreamProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetSchemas

`func (o *AddPkcs11CipherStreamProviderRequest) GetSchemas() []Enumpkcs11CipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPkcs11CipherStreamProviderRequest) GetSchemasOk() (*[]Enumpkcs11CipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPkcs11CipherStreamProviderRequest) SetSchemas(v []Enumpkcs11CipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPkcs11ProviderClass

`func (o *AddPkcs11CipherStreamProviderRequest) GetPkcs11ProviderClass() string`

GetPkcs11ProviderClass returns the Pkcs11ProviderClass field if non-nil, zero value otherwise.

### GetPkcs11ProviderClassOk

`func (o *AddPkcs11CipherStreamProviderRequest) GetPkcs11ProviderClassOk() (*string, bool)`

GetPkcs11ProviderClassOk returns a tuple with the Pkcs11ProviderClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11ProviderClass

`func (o *AddPkcs11CipherStreamProviderRequest) SetPkcs11ProviderClass(v string)`

SetPkcs11ProviderClass sets Pkcs11ProviderClass field to given value.

### HasPkcs11ProviderClass

`func (o *AddPkcs11CipherStreamProviderRequest) HasPkcs11ProviderClass() bool`

HasPkcs11ProviderClass returns a boolean if a field has been set.

### GetPkcs11ProviderConfigurationFile

`func (o *AddPkcs11CipherStreamProviderRequest) GetPkcs11ProviderConfigurationFile() string`

GetPkcs11ProviderConfigurationFile returns the Pkcs11ProviderConfigurationFile field if non-nil, zero value otherwise.

### GetPkcs11ProviderConfigurationFileOk

`func (o *AddPkcs11CipherStreamProviderRequest) GetPkcs11ProviderConfigurationFileOk() (*string, bool)`

GetPkcs11ProviderConfigurationFileOk returns a tuple with the Pkcs11ProviderConfigurationFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11ProviderConfigurationFile

`func (o *AddPkcs11CipherStreamProviderRequest) SetPkcs11ProviderConfigurationFile(v string)`

SetPkcs11ProviderConfigurationFile sets Pkcs11ProviderConfigurationFile field to given value.

### HasPkcs11ProviderConfigurationFile

`func (o *AddPkcs11CipherStreamProviderRequest) HasPkcs11ProviderConfigurationFile() bool`

HasPkcs11ProviderConfigurationFile returns a boolean if a field has been set.

### GetKeyStorePin

`func (o *AddPkcs11CipherStreamProviderRequest) GetKeyStorePin() string`

GetKeyStorePin returns the KeyStorePin field if non-nil, zero value otherwise.

### GetKeyStorePinOk

`func (o *AddPkcs11CipherStreamProviderRequest) GetKeyStorePinOk() (*string, bool)`

GetKeyStorePinOk returns a tuple with the KeyStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePin

`func (o *AddPkcs11CipherStreamProviderRequest) SetKeyStorePin(v string)`

SetKeyStorePin sets KeyStorePin field to given value.

### HasKeyStorePin

`func (o *AddPkcs11CipherStreamProviderRequest) HasKeyStorePin() bool`

HasKeyStorePin returns a boolean if a field has been set.

### GetKeyStorePinFile

`func (o *AddPkcs11CipherStreamProviderRequest) GetKeyStorePinFile() string`

GetKeyStorePinFile returns the KeyStorePinFile field if non-nil, zero value otherwise.

### GetKeyStorePinFileOk

`func (o *AddPkcs11CipherStreamProviderRequest) GetKeyStorePinFileOk() (*string, bool)`

GetKeyStorePinFileOk returns a tuple with the KeyStorePinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePinFile

`func (o *AddPkcs11CipherStreamProviderRequest) SetKeyStorePinFile(v string)`

SetKeyStorePinFile sets KeyStorePinFile field to given value.

### HasKeyStorePinFile

`func (o *AddPkcs11CipherStreamProviderRequest) HasKeyStorePinFile() bool`

HasKeyStorePinFile returns a boolean if a field has been set.

### GetKeyStorePinEnvironmentVariable

`func (o *AddPkcs11CipherStreamProviderRequest) GetKeyStorePinEnvironmentVariable() string`

GetKeyStorePinEnvironmentVariable returns the KeyStorePinEnvironmentVariable field if non-nil, zero value otherwise.

### GetKeyStorePinEnvironmentVariableOk

`func (o *AddPkcs11CipherStreamProviderRequest) GetKeyStorePinEnvironmentVariableOk() (*string, bool)`

GetKeyStorePinEnvironmentVariableOk returns a tuple with the KeyStorePinEnvironmentVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePinEnvironmentVariable

`func (o *AddPkcs11CipherStreamProviderRequest) SetKeyStorePinEnvironmentVariable(v string)`

SetKeyStorePinEnvironmentVariable sets KeyStorePinEnvironmentVariable field to given value.

### HasKeyStorePinEnvironmentVariable

`func (o *AddPkcs11CipherStreamProviderRequest) HasKeyStorePinEnvironmentVariable() bool`

HasKeyStorePinEnvironmentVariable returns a boolean if a field has been set.

### GetPkcs11KeyStoreType

`func (o *AddPkcs11CipherStreamProviderRequest) GetPkcs11KeyStoreType() string`

GetPkcs11KeyStoreType returns the Pkcs11KeyStoreType field if non-nil, zero value otherwise.

### GetPkcs11KeyStoreTypeOk

`func (o *AddPkcs11CipherStreamProviderRequest) GetPkcs11KeyStoreTypeOk() (*string, bool)`

GetPkcs11KeyStoreTypeOk returns a tuple with the Pkcs11KeyStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11KeyStoreType

`func (o *AddPkcs11CipherStreamProviderRequest) SetPkcs11KeyStoreType(v string)`

SetPkcs11KeyStoreType sets Pkcs11KeyStoreType field to given value.

### HasPkcs11KeyStoreType

`func (o *AddPkcs11CipherStreamProviderRequest) HasPkcs11KeyStoreType() bool`

HasPkcs11KeyStoreType returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *AddPkcs11CipherStreamProviderRequest) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *AddPkcs11CipherStreamProviderRequest) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *AddPkcs11CipherStreamProviderRequest) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.


### GetEncryptionMetadataFile

`func (o *AddPkcs11CipherStreamProviderRequest) GetEncryptionMetadataFile() string`

GetEncryptionMetadataFile returns the EncryptionMetadataFile field if non-nil, zero value otherwise.

### GetEncryptionMetadataFileOk

`func (o *AddPkcs11CipherStreamProviderRequest) GetEncryptionMetadataFileOk() (*string, bool)`

GetEncryptionMetadataFileOk returns a tuple with the EncryptionMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMetadataFile

`func (o *AddPkcs11CipherStreamProviderRequest) SetEncryptionMetadataFile(v string)`

SetEncryptionMetadataFile sets EncryptionMetadataFile field to given value.

### HasEncryptionMetadataFile

`func (o *AddPkcs11CipherStreamProviderRequest) HasEncryptionMetadataFile() bool`

HasEncryptionMetadataFile returns a boolean if a field has been set.

### GetIterationCount

`func (o *AddPkcs11CipherStreamProviderRequest) GetIterationCount() int64`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *AddPkcs11CipherStreamProviderRequest) GetIterationCountOk() (*int64, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *AddPkcs11CipherStreamProviderRequest) SetIterationCount(v int64)`

SetIterationCount sets IterationCount field to given value.

### HasIterationCount

`func (o *AddPkcs11CipherStreamProviderRequest) HasIterationCount() bool`

HasIterationCount returns a boolean if a field has been set.

### GetDescription

`func (o *AddPkcs11CipherStreamProviderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPkcs11CipherStreamProviderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPkcs11CipherStreamProviderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPkcs11CipherStreamProviderRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPkcs11CipherStreamProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPkcs11CipherStreamProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPkcs11CipherStreamProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


