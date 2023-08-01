# CipherStreamProviderListResponseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Cipher Stream Provider | 
**Schemas** | [**[]EnumthirdPartyCipherStreamProviderSchemaUrn**](EnumthirdPartyCipherStreamProviderSchemaUrn.md) |  | 
**EncryptedPassphraseFile** | **string** | The path to a file that will hold the encrypted passphrase used by this cipher stream provider. | 
**AwsExternalServer** | **string** | The external server with information to use when interacting with the AWS Secrets Manager. | 
**AwsAccessKeyID** | Pointer to **string** | The access key ID that will be used if this cipher stream provider will authenticate to the Amazon Key Management Service using an access key rather than an IAM role associated with an EC2 instance. | [optional] 
**AwsSecretAccessKey** | Pointer to **string** | The secret access key that will be used if this cipher stream provider will authenticate to the Amazon Key Management Service using an access key rather than an IAM role associated with an EC2 instance. | [optional] 
**AwsRegionName** | Pointer to **string** | The name of the Amazon Web Services region that holds the encryption key. This is optional, and if it is not provided, then the server will attempt to determine the region from the key ARN. | [optional] 
**KmsEncryptionKeyArn** | **string** | The Amazon resource name (ARN) for the KMS key that will be used to encrypt the contents of the passphrase file. This key must exist, and the AWS client must have access to encrypt and decrypt data using this key. | 
**IterationCount** | Pointer to **int64** | The PBKDF2 iteration count that will be used when deriving the encryption key used to protect the encryption settings database. | [optional] 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**SecretID** | **string** | The Amazon Resource Name (ARN) or the user-friendly name of the secret to be retrieved. | 
**SecretFieldName** | **string** | The name of the JSON field whose value is the passphrase that will be used to generate the encryption key for protecting the contents of the encryption settings database. | 
**SecretVersionID** | Pointer to **string** | The unique identifier for the version of the secret to be retrieved. | [optional] 
**SecretVersionStage** | Pointer to **string** | The staging label for the version of the secret to be retrieved. | [optional] 
**EncryptionMetadataFile** | **string** | The path to a file that will hold metadata about the encryption performed by this PKCS11 Cipher Stream Provider. | 
**KeyVaultURI** | **string** | The URI that identifies the Azure Key Vault from which the secret is to be retrieved. | 
**AzureAuthenticationMethod** | **string** | The mechanism used to authenticate to the Azure service. | 
**HttpProxyExternalServer** | Pointer to **string** | A reference to an HTTP proxy server that should be used for requests sent to the Azure service. | [optional] 
**SecretName** | **string** | The name of the secret to retrieve. | 
**PasswordFile** | **string** | The path to the file containing the password to use when generating ciphers. | 
**WaitForPasswordFile** | Pointer to **bool** | Indicates whether the server should wait for the password file to become available if it does not exist. | [optional] 
**ConjurExternalServer** | **string** | An external server definition with information needed to connect and authenticate to the Conjur server. | 
**ConjurSecretRelativePath** | **string** | The portion of the path that follows the account name in the URI needed to obtain the secret passphrase to use to generate the encryption key. Any special characters in the path must be URL-encoded. | 
**Pkcs11ProviderClass** | Pointer to **string** | The fully-qualified name of the Java security provider class that implements support for interacting with PKCS #11 tokens. | [optional] 
**Pkcs11ProviderConfigurationFile** | Pointer to **string** | The path to the file to use to configure the security provider that implements support for interacting with PKCS #11 tokens. | [optional] 
**KeyStorePin** | Pointer to **string** | The clear-text user PIN needed to interact with the PKCS #11 token. | [optional] 
**KeyStorePinFile** | Pointer to **string** | The path to a file containing the user PIN needed to interact with the PKCS #11 token. The file must exist and must contain exactly one line with a clear-text representation of the PIN. | [optional] 
**KeyStorePinEnvironmentVariable** | Pointer to **string** | The name of an environment variable whose value is the user PIN needed to interact with the PKCS #11 token. The environment variable must be defined and must contain a clear-text representation of the PIN. | [optional] 
**Pkcs11KeyStoreType** | Pointer to **string** | The key store type to use when obtaining an instance of a key store for interacting with a PKCS #11 token. | [optional] 
**SslCertNickname** | **string** | The alias for the certificate in the PKCS #11 token that will be used to wrap the encryption key. The target certificate must exist in the PKCS #11 token, and it must have an RSA key pair because the JVM does not currently provide adequate key wrapping support for elliptic curve key pairs.  If you have also configured the server to use a PKCS #11 token for accessing listener certificates, we strongly recommend that you use a different certificate to protect the contents of the encryption settings database than you use for negotiating TLS sessions with clients. It is imperative that the certificate used by this PKCS11 Cipher Stream Provider remain constant for the life of the provider because if the certificate were to be replaced, then the contents of the encryption settings database could become inaccessible. Unlike with listener certificates used for TLS negotiation that need to be replaced on a regular basis, this PKCS11 Cipher Stream Provider does not consider the validity period for the associated certificate, and it will continue to function even after the certificate has expired.  If you need to rotate the certificate used to protect the server&#39;s encryption settings database, you should first install the desired new certificate in the PKCS #11 token under a different alias. Then, you should create a new instance of this PKCS11 Cipher Stream Provider that is configured to use that certificate, and that also uses a different value for the encryption-metadata-file because the information in that file is tied to the certificate used to generate it. Finally, you will need to update the global configuration so that the encryption-settings-cipher-stream-provider property references the new cipher stream provider rather than this one. The update to the global configuration must be done with the server online so that it can properly re-encrypt the contents of the encryption settings database with the correct key tied to the new certificate. | 
**VaultExternalServer** | Pointer to **string** | An external server definition with information needed to connect and authenticate to the Vault server. | [optional] 
**VaultServerBaseURI** | Pointer to **[]string** | The base URL needed to access the Vault server. The base URL should consist of the protocol (\&quot;http\&quot; or \&quot;https\&quot;), the server address (resolvable name or IP address), and the port number. For example, \&quot;https://vault.example.com:8200/\&quot;. | [optional] 
**VaultAuthenticationMethod** | Pointer to **string** | The mechanism used to authenticate to the Vault server. | [optional] 
**VaultSecretPath** | **string** | The path to the desired secret in the Vault service. This will be appended to the value of the base-url property for the associated Vault external server. | 
**VaultSecretFieldName** | **string** | The name of the field in the Vault secret record that contains the passphrase to use to generate the encryption key. | 
**VaultEncryptionMetadataFile** | **string** | The path to a file that will hold metadata about the encryption performed by this Vault Cipher Stream Provider. | 
**TrustStoreFile** | Pointer to **string** | The path to a file containing the information needed to trust the certificate presented by the Vault servers. | [optional] 
**TrustStorePin** | Pointer to **string** | The passphrase needed to access the contents of the trust store. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents. | [optional] 
**TrustStoreType** | Pointer to **string** | The store type for the specified trust store file. The value should likely be one of \&quot;JKS\&quot; or \&quot;PKCS12\&quot;. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Cipher Stream Provider. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Cipher Stream Provider. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewCipherStreamProviderListResponseResourcesInner

`func NewCipherStreamProviderListResponseResourcesInner(id string, schemas []EnumthirdPartyCipherStreamProviderSchemaUrn, encryptedPassphraseFile string, awsExternalServer string, kmsEncryptionKeyArn string, enabled bool, secretID string, secretFieldName string, encryptionMetadataFile string, keyVaultURI string, azureAuthenticationMethod string, secretName string, passwordFile string, conjurExternalServer string, conjurSecretRelativePath string, sslCertNickname string, vaultSecretPath string, vaultSecretFieldName string, vaultEncryptionMetadataFile string, extensionClass string, ) *CipherStreamProviderListResponseResourcesInner`

NewCipherStreamProviderListResponseResourcesInner instantiates a new CipherStreamProviderListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCipherStreamProviderListResponseResourcesInnerWithDefaults

`func NewCipherStreamProviderListResponseResourcesInnerWithDefaults() *CipherStreamProviderListResponseResourcesInner`

NewCipherStreamProviderListResponseResourcesInnerWithDefaults instantiates a new CipherStreamProviderListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CipherStreamProviderListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CipherStreamProviderListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *CipherStreamProviderListResponseResourcesInner) GetSchemas() []EnumthirdPartyCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetSchemasOk() (*[]EnumthirdPartyCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CipherStreamProviderListResponseResourcesInner) SetSchemas(v []EnumthirdPartyCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEncryptedPassphraseFile

`func (o *CipherStreamProviderListResponseResourcesInner) GetEncryptedPassphraseFile() string`

GetEncryptedPassphraseFile returns the EncryptedPassphraseFile field if non-nil, zero value otherwise.

### GetEncryptedPassphraseFileOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetEncryptedPassphraseFileOk() (*string, bool)`

GetEncryptedPassphraseFileOk returns a tuple with the EncryptedPassphraseFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassphraseFile

`func (o *CipherStreamProviderListResponseResourcesInner) SetEncryptedPassphraseFile(v string)`

SetEncryptedPassphraseFile sets EncryptedPassphraseFile field to given value.


### GetAwsExternalServer

`func (o *CipherStreamProviderListResponseResourcesInner) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *CipherStreamProviderListResponseResourcesInner) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.


### GetAwsAccessKeyID

`func (o *CipherStreamProviderListResponseResourcesInner) GetAwsAccessKeyID() string`

GetAwsAccessKeyID returns the AwsAccessKeyID field if non-nil, zero value otherwise.

### GetAwsAccessKeyIDOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetAwsAccessKeyIDOk() (*string, bool)`

GetAwsAccessKeyIDOk returns a tuple with the AwsAccessKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyID

`func (o *CipherStreamProviderListResponseResourcesInner) SetAwsAccessKeyID(v string)`

SetAwsAccessKeyID sets AwsAccessKeyID field to given value.

### HasAwsAccessKeyID

`func (o *CipherStreamProviderListResponseResourcesInner) HasAwsAccessKeyID() bool`

HasAwsAccessKeyID returns a boolean if a field has been set.

### GetAwsSecretAccessKey

`func (o *CipherStreamProviderListResponseResourcesInner) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *CipherStreamProviderListResponseResourcesInner) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *CipherStreamProviderListResponseResourcesInner) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### GetAwsRegionName

`func (o *CipherStreamProviderListResponseResourcesInner) GetAwsRegionName() string`

GetAwsRegionName returns the AwsRegionName field if non-nil, zero value otherwise.

### GetAwsRegionNameOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetAwsRegionNameOk() (*string, bool)`

GetAwsRegionNameOk returns a tuple with the AwsRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegionName

`func (o *CipherStreamProviderListResponseResourcesInner) SetAwsRegionName(v string)`

SetAwsRegionName sets AwsRegionName field to given value.

### HasAwsRegionName

`func (o *CipherStreamProviderListResponseResourcesInner) HasAwsRegionName() bool`

HasAwsRegionName returns a boolean if a field has been set.

### GetKmsEncryptionKeyArn

`func (o *CipherStreamProviderListResponseResourcesInner) GetKmsEncryptionKeyArn() string`

GetKmsEncryptionKeyArn returns the KmsEncryptionKeyArn field if non-nil, zero value otherwise.

### GetKmsEncryptionKeyArnOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetKmsEncryptionKeyArnOk() (*string, bool)`

GetKmsEncryptionKeyArnOk returns a tuple with the KmsEncryptionKeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsEncryptionKeyArn

`func (o *CipherStreamProviderListResponseResourcesInner) SetKmsEncryptionKeyArn(v string)`

SetKmsEncryptionKeyArn sets KmsEncryptionKeyArn field to given value.


### GetIterationCount

`func (o *CipherStreamProviderListResponseResourcesInner) GetIterationCount() int64`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetIterationCountOk() (*int64, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *CipherStreamProviderListResponseResourcesInner) SetIterationCount(v int64)`

SetIterationCount sets IterationCount field to given value.

### HasIterationCount

`func (o *CipherStreamProviderListResponseResourcesInner) HasIterationCount() bool`

HasIterationCount returns a boolean if a field has been set.

### GetDescription

`func (o *CipherStreamProviderListResponseResourcesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CipherStreamProviderListResponseResourcesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CipherStreamProviderListResponseResourcesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CipherStreamProviderListResponseResourcesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CipherStreamProviderListResponseResourcesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *CipherStreamProviderListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CipherStreamProviderListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CipherStreamProviderListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *CipherStreamProviderListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *CipherStreamProviderListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *CipherStreamProviderListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *CipherStreamProviderListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSecretID

`func (o *CipherStreamProviderListResponseResourcesInner) GetSecretID() string`

GetSecretID returns the SecretID field if non-nil, zero value otherwise.

### GetSecretIDOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetSecretIDOk() (*string, bool)`

GetSecretIDOk returns a tuple with the SecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretID

`func (o *CipherStreamProviderListResponseResourcesInner) SetSecretID(v string)`

SetSecretID sets SecretID field to given value.


### GetSecretFieldName

`func (o *CipherStreamProviderListResponseResourcesInner) GetSecretFieldName() string`

GetSecretFieldName returns the SecretFieldName field if non-nil, zero value otherwise.

### GetSecretFieldNameOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetSecretFieldNameOk() (*string, bool)`

GetSecretFieldNameOk returns a tuple with the SecretFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretFieldName

`func (o *CipherStreamProviderListResponseResourcesInner) SetSecretFieldName(v string)`

SetSecretFieldName sets SecretFieldName field to given value.


### GetSecretVersionID

`func (o *CipherStreamProviderListResponseResourcesInner) GetSecretVersionID() string`

GetSecretVersionID returns the SecretVersionID field if non-nil, zero value otherwise.

### GetSecretVersionIDOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetSecretVersionIDOk() (*string, bool)`

GetSecretVersionIDOk returns a tuple with the SecretVersionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVersionID

`func (o *CipherStreamProviderListResponseResourcesInner) SetSecretVersionID(v string)`

SetSecretVersionID sets SecretVersionID field to given value.

### HasSecretVersionID

`func (o *CipherStreamProviderListResponseResourcesInner) HasSecretVersionID() bool`

HasSecretVersionID returns a boolean if a field has been set.

### GetSecretVersionStage

`func (o *CipherStreamProviderListResponseResourcesInner) GetSecretVersionStage() string`

GetSecretVersionStage returns the SecretVersionStage field if non-nil, zero value otherwise.

### GetSecretVersionStageOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetSecretVersionStageOk() (*string, bool)`

GetSecretVersionStageOk returns a tuple with the SecretVersionStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVersionStage

`func (o *CipherStreamProviderListResponseResourcesInner) SetSecretVersionStage(v string)`

SetSecretVersionStage sets SecretVersionStage field to given value.

### HasSecretVersionStage

`func (o *CipherStreamProviderListResponseResourcesInner) HasSecretVersionStage() bool`

HasSecretVersionStage returns a boolean if a field has been set.

### GetEncryptionMetadataFile

`func (o *CipherStreamProviderListResponseResourcesInner) GetEncryptionMetadataFile() string`

GetEncryptionMetadataFile returns the EncryptionMetadataFile field if non-nil, zero value otherwise.

### GetEncryptionMetadataFileOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetEncryptionMetadataFileOk() (*string, bool)`

GetEncryptionMetadataFileOk returns a tuple with the EncryptionMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMetadataFile

`func (o *CipherStreamProviderListResponseResourcesInner) SetEncryptionMetadataFile(v string)`

SetEncryptionMetadataFile sets EncryptionMetadataFile field to given value.


### GetKeyVaultURI

`func (o *CipherStreamProviderListResponseResourcesInner) GetKeyVaultURI() string`

GetKeyVaultURI returns the KeyVaultURI field if non-nil, zero value otherwise.

### GetKeyVaultURIOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetKeyVaultURIOk() (*string, bool)`

GetKeyVaultURIOk returns a tuple with the KeyVaultURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVaultURI

`func (o *CipherStreamProviderListResponseResourcesInner) SetKeyVaultURI(v string)`

SetKeyVaultURI sets KeyVaultURI field to given value.


### GetAzureAuthenticationMethod

`func (o *CipherStreamProviderListResponseResourcesInner) GetAzureAuthenticationMethod() string`

GetAzureAuthenticationMethod returns the AzureAuthenticationMethod field if non-nil, zero value otherwise.

### GetAzureAuthenticationMethodOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetAzureAuthenticationMethodOk() (*string, bool)`

GetAzureAuthenticationMethodOk returns a tuple with the AzureAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAuthenticationMethod

`func (o *CipherStreamProviderListResponseResourcesInner) SetAzureAuthenticationMethod(v string)`

SetAzureAuthenticationMethod sets AzureAuthenticationMethod field to given value.


### GetHttpProxyExternalServer

`func (o *CipherStreamProviderListResponseResourcesInner) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *CipherStreamProviderListResponseResourcesInner) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *CipherStreamProviderListResponseResourcesInner) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetSecretName

`func (o *CipherStreamProviderListResponseResourcesInner) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *CipherStreamProviderListResponseResourcesInner) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetPasswordFile

`func (o *CipherStreamProviderListResponseResourcesInner) GetPasswordFile() string`

GetPasswordFile returns the PasswordFile field if non-nil, zero value otherwise.

### GetPasswordFileOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetPasswordFileOk() (*string, bool)`

GetPasswordFileOk returns a tuple with the PasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFile

`func (o *CipherStreamProviderListResponseResourcesInner) SetPasswordFile(v string)`

SetPasswordFile sets PasswordFile field to given value.


### GetWaitForPasswordFile

`func (o *CipherStreamProviderListResponseResourcesInner) GetWaitForPasswordFile() bool`

GetWaitForPasswordFile returns the WaitForPasswordFile field if non-nil, zero value otherwise.

### GetWaitForPasswordFileOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetWaitForPasswordFileOk() (*bool, bool)`

GetWaitForPasswordFileOk returns a tuple with the WaitForPasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForPasswordFile

`func (o *CipherStreamProviderListResponseResourcesInner) SetWaitForPasswordFile(v bool)`

SetWaitForPasswordFile sets WaitForPasswordFile field to given value.

### HasWaitForPasswordFile

`func (o *CipherStreamProviderListResponseResourcesInner) HasWaitForPasswordFile() bool`

HasWaitForPasswordFile returns a boolean if a field has been set.

### GetConjurExternalServer

`func (o *CipherStreamProviderListResponseResourcesInner) GetConjurExternalServer() string`

GetConjurExternalServer returns the ConjurExternalServer field if non-nil, zero value otherwise.

### GetConjurExternalServerOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetConjurExternalServerOk() (*string, bool)`

GetConjurExternalServerOk returns a tuple with the ConjurExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurExternalServer

`func (o *CipherStreamProviderListResponseResourcesInner) SetConjurExternalServer(v string)`

SetConjurExternalServer sets ConjurExternalServer field to given value.


### GetConjurSecretRelativePath

`func (o *CipherStreamProviderListResponseResourcesInner) GetConjurSecretRelativePath() string`

GetConjurSecretRelativePath returns the ConjurSecretRelativePath field if non-nil, zero value otherwise.

### GetConjurSecretRelativePathOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetConjurSecretRelativePathOk() (*string, bool)`

GetConjurSecretRelativePathOk returns a tuple with the ConjurSecretRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurSecretRelativePath

`func (o *CipherStreamProviderListResponseResourcesInner) SetConjurSecretRelativePath(v string)`

SetConjurSecretRelativePath sets ConjurSecretRelativePath field to given value.


### GetPkcs11ProviderClass

`func (o *CipherStreamProviderListResponseResourcesInner) GetPkcs11ProviderClass() string`

GetPkcs11ProviderClass returns the Pkcs11ProviderClass field if non-nil, zero value otherwise.

### GetPkcs11ProviderClassOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetPkcs11ProviderClassOk() (*string, bool)`

GetPkcs11ProviderClassOk returns a tuple with the Pkcs11ProviderClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11ProviderClass

`func (o *CipherStreamProviderListResponseResourcesInner) SetPkcs11ProviderClass(v string)`

SetPkcs11ProviderClass sets Pkcs11ProviderClass field to given value.

### HasPkcs11ProviderClass

`func (o *CipherStreamProviderListResponseResourcesInner) HasPkcs11ProviderClass() bool`

HasPkcs11ProviderClass returns a boolean if a field has been set.

### GetPkcs11ProviderConfigurationFile

`func (o *CipherStreamProviderListResponseResourcesInner) GetPkcs11ProviderConfigurationFile() string`

GetPkcs11ProviderConfigurationFile returns the Pkcs11ProviderConfigurationFile field if non-nil, zero value otherwise.

### GetPkcs11ProviderConfigurationFileOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetPkcs11ProviderConfigurationFileOk() (*string, bool)`

GetPkcs11ProviderConfigurationFileOk returns a tuple with the Pkcs11ProviderConfigurationFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11ProviderConfigurationFile

`func (o *CipherStreamProviderListResponseResourcesInner) SetPkcs11ProviderConfigurationFile(v string)`

SetPkcs11ProviderConfigurationFile sets Pkcs11ProviderConfigurationFile field to given value.

### HasPkcs11ProviderConfigurationFile

`func (o *CipherStreamProviderListResponseResourcesInner) HasPkcs11ProviderConfigurationFile() bool`

HasPkcs11ProviderConfigurationFile returns a boolean if a field has been set.

### GetKeyStorePin

`func (o *CipherStreamProviderListResponseResourcesInner) GetKeyStorePin() string`

GetKeyStorePin returns the KeyStorePin field if non-nil, zero value otherwise.

### GetKeyStorePinOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetKeyStorePinOk() (*string, bool)`

GetKeyStorePinOk returns a tuple with the KeyStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePin

`func (o *CipherStreamProviderListResponseResourcesInner) SetKeyStorePin(v string)`

SetKeyStorePin sets KeyStorePin field to given value.

### HasKeyStorePin

`func (o *CipherStreamProviderListResponseResourcesInner) HasKeyStorePin() bool`

HasKeyStorePin returns a boolean if a field has been set.

### GetKeyStorePinFile

`func (o *CipherStreamProviderListResponseResourcesInner) GetKeyStorePinFile() string`

GetKeyStorePinFile returns the KeyStorePinFile field if non-nil, zero value otherwise.

### GetKeyStorePinFileOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetKeyStorePinFileOk() (*string, bool)`

GetKeyStorePinFileOk returns a tuple with the KeyStorePinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePinFile

`func (o *CipherStreamProviderListResponseResourcesInner) SetKeyStorePinFile(v string)`

SetKeyStorePinFile sets KeyStorePinFile field to given value.

### HasKeyStorePinFile

`func (o *CipherStreamProviderListResponseResourcesInner) HasKeyStorePinFile() bool`

HasKeyStorePinFile returns a boolean if a field has been set.

### GetKeyStorePinEnvironmentVariable

`func (o *CipherStreamProviderListResponseResourcesInner) GetKeyStorePinEnvironmentVariable() string`

GetKeyStorePinEnvironmentVariable returns the KeyStorePinEnvironmentVariable field if non-nil, zero value otherwise.

### GetKeyStorePinEnvironmentVariableOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetKeyStorePinEnvironmentVariableOk() (*string, bool)`

GetKeyStorePinEnvironmentVariableOk returns a tuple with the KeyStorePinEnvironmentVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePinEnvironmentVariable

`func (o *CipherStreamProviderListResponseResourcesInner) SetKeyStorePinEnvironmentVariable(v string)`

SetKeyStorePinEnvironmentVariable sets KeyStorePinEnvironmentVariable field to given value.

### HasKeyStorePinEnvironmentVariable

`func (o *CipherStreamProviderListResponseResourcesInner) HasKeyStorePinEnvironmentVariable() bool`

HasKeyStorePinEnvironmentVariable returns a boolean if a field has been set.

### GetPkcs11KeyStoreType

`func (o *CipherStreamProviderListResponseResourcesInner) GetPkcs11KeyStoreType() string`

GetPkcs11KeyStoreType returns the Pkcs11KeyStoreType field if non-nil, zero value otherwise.

### GetPkcs11KeyStoreTypeOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetPkcs11KeyStoreTypeOk() (*string, bool)`

GetPkcs11KeyStoreTypeOk returns a tuple with the Pkcs11KeyStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11KeyStoreType

`func (o *CipherStreamProviderListResponseResourcesInner) SetPkcs11KeyStoreType(v string)`

SetPkcs11KeyStoreType sets Pkcs11KeyStoreType field to given value.

### HasPkcs11KeyStoreType

`func (o *CipherStreamProviderListResponseResourcesInner) HasPkcs11KeyStoreType() bool`

HasPkcs11KeyStoreType returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *CipherStreamProviderListResponseResourcesInner) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *CipherStreamProviderListResponseResourcesInner) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.


### GetVaultExternalServer

`func (o *CipherStreamProviderListResponseResourcesInner) GetVaultExternalServer() string`

GetVaultExternalServer returns the VaultExternalServer field if non-nil, zero value otherwise.

### GetVaultExternalServerOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetVaultExternalServerOk() (*string, bool)`

GetVaultExternalServerOk returns a tuple with the VaultExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultExternalServer

`func (o *CipherStreamProviderListResponseResourcesInner) SetVaultExternalServer(v string)`

SetVaultExternalServer sets VaultExternalServer field to given value.

### HasVaultExternalServer

`func (o *CipherStreamProviderListResponseResourcesInner) HasVaultExternalServer() bool`

HasVaultExternalServer returns a boolean if a field has been set.

### GetVaultServerBaseURI

`func (o *CipherStreamProviderListResponseResourcesInner) GetVaultServerBaseURI() []string`

GetVaultServerBaseURI returns the VaultServerBaseURI field if non-nil, zero value otherwise.

### GetVaultServerBaseURIOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetVaultServerBaseURIOk() (*[]string, bool)`

GetVaultServerBaseURIOk returns a tuple with the VaultServerBaseURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultServerBaseURI

`func (o *CipherStreamProviderListResponseResourcesInner) SetVaultServerBaseURI(v []string)`

SetVaultServerBaseURI sets VaultServerBaseURI field to given value.

### HasVaultServerBaseURI

`func (o *CipherStreamProviderListResponseResourcesInner) HasVaultServerBaseURI() bool`

HasVaultServerBaseURI returns a boolean if a field has been set.

### GetVaultAuthenticationMethod

`func (o *CipherStreamProviderListResponseResourcesInner) GetVaultAuthenticationMethod() string`

GetVaultAuthenticationMethod returns the VaultAuthenticationMethod field if non-nil, zero value otherwise.

### GetVaultAuthenticationMethodOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetVaultAuthenticationMethodOk() (*string, bool)`

GetVaultAuthenticationMethodOk returns a tuple with the VaultAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAuthenticationMethod

`func (o *CipherStreamProviderListResponseResourcesInner) SetVaultAuthenticationMethod(v string)`

SetVaultAuthenticationMethod sets VaultAuthenticationMethod field to given value.

### HasVaultAuthenticationMethod

`func (o *CipherStreamProviderListResponseResourcesInner) HasVaultAuthenticationMethod() bool`

HasVaultAuthenticationMethod returns a boolean if a field has been set.

### GetVaultSecretPath

`func (o *CipherStreamProviderListResponseResourcesInner) GetVaultSecretPath() string`

GetVaultSecretPath returns the VaultSecretPath field if non-nil, zero value otherwise.

### GetVaultSecretPathOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetVaultSecretPathOk() (*string, bool)`

GetVaultSecretPathOk returns a tuple with the VaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretPath

`func (o *CipherStreamProviderListResponseResourcesInner) SetVaultSecretPath(v string)`

SetVaultSecretPath sets VaultSecretPath field to given value.


### GetVaultSecretFieldName

`func (o *CipherStreamProviderListResponseResourcesInner) GetVaultSecretFieldName() string`

GetVaultSecretFieldName returns the VaultSecretFieldName field if non-nil, zero value otherwise.

### GetVaultSecretFieldNameOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetVaultSecretFieldNameOk() (*string, bool)`

GetVaultSecretFieldNameOk returns a tuple with the VaultSecretFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretFieldName

`func (o *CipherStreamProviderListResponseResourcesInner) SetVaultSecretFieldName(v string)`

SetVaultSecretFieldName sets VaultSecretFieldName field to given value.


### GetVaultEncryptionMetadataFile

`func (o *CipherStreamProviderListResponseResourcesInner) GetVaultEncryptionMetadataFile() string`

GetVaultEncryptionMetadataFile returns the VaultEncryptionMetadataFile field if non-nil, zero value otherwise.

### GetVaultEncryptionMetadataFileOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetVaultEncryptionMetadataFileOk() (*string, bool)`

GetVaultEncryptionMetadataFileOk returns a tuple with the VaultEncryptionMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultEncryptionMetadataFile

`func (o *CipherStreamProviderListResponseResourcesInner) SetVaultEncryptionMetadataFile(v string)`

SetVaultEncryptionMetadataFile sets VaultEncryptionMetadataFile field to given value.


### GetTrustStoreFile

`func (o *CipherStreamProviderListResponseResourcesInner) GetTrustStoreFile() string`

GetTrustStoreFile returns the TrustStoreFile field if non-nil, zero value otherwise.

### GetTrustStoreFileOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetTrustStoreFileOk() (*string, bool)`

GetTrustStoreFileOk returns a tuple with the TrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreFile

`func (o *CipherStreamProviderListResponseResourcesInner) SetTrustStoreFile(v string)`

SetTrustStoreFile sets TrustStoreFile field to given value.

### HasTrustStoreFile

`func (o *CipherStreamProviderListResponseResourcesInner) HasTrustStoreFile() bool`

HasTrustStoreFile returns a boolean if a field has been set.

### GetTrustStorePin

`func (o *CipherStreamProviderListResponseResourcesInner) GetTrustStorePin() string`

GetTrustStorePin returns the TrustStorePin field if non-nil, zero value otherwise.

### GetTrustStorePinOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetTrustStorePinOk() (*string, bool)`

GetTrustStorePinOk returns a tuple with the TrustStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePin

`func (o *CipherStreamProviderListResponseResourcesInner) SetTrustStorePin(v string)`

SetTrustStorePin sets TrustStorePin field to given value.

### HasTrustStorePin

`func (o *CipherStreamProviderListResponseResourcesInner) HasTrustStorePin() bool`

HasTrustStorePin returns a boolean if a field has been set.

### GetTrustStoreType

`func (o *CipherStreamProviderListResponseResourcesInner) GetTrustStoreType() string`

GetTrustStoreType returns the TrustStoreType field if non-nil, zero value otherwise.

### GetTrustStoreTypeOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetTrustStoreTypeOk() (*string, bool)`

GetTrustStoreTypeOk returns a tuple with the TrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreType

`func (o *CipherStreamProviderListResponseResourcesInner) SetTrustStoreType(v string)`

SetTrustStoreType sets TrustStoreType field to given value.

### HasTrustStoreType

`func (o *CipherStreamProviderListResponseResourcesInner) HasTrustStoreType() bool`

HasTrustStoreType returns a boolean if a field has been set.

### GetExtensionClass

`func (o *CipherStreamProviderListResponseResourcesInner) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *CipherStreamProviderListResponseResourcesInner) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *CipherStreamProviderListResponseResourcesInner) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *CipherStreamProviderListResponseResourcesInner) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *CipherStreamProviderListResponseResourcesInner) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *CipherStreamProviderListResponseResourcesInner) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


