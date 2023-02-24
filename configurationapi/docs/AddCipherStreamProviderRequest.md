# AddCipherStreamProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Name of the new Cipher Stream Provider | 
**Schemas** | [**[]EnumthirdPartyCipherStreamProviderSchemaUrn**](EnumthirdPartyCipherStreamProviderSchemaUrn.md) |  | 
**EncryptedPassphraseFile** | Pointer to **string** | The path to a file that will hold the encrypted passphrase used by this cipher stream provider. | [optional] 
**AwsExternalServer** | **string** | The external server with information to use when interacting with the AWS Secrets Manager. | 
**AwsAccessKeyID** | Pointer to **string** | The access key ID that will be used if this cipher stream provider will authenticate to the Amazon Key Management Service using an access key rather than an IAM role associated with an EC2 instance. | [optional] 
**AwsSecretAccessKey** | Pointer to **string** | The secret access key that will be used if this cipher stream provider will authenticate to the Amazon Key Management Service using an access key rather than an IAM role associated with an EC2 instance. | [optional] 
**AwsRegionName** | Pointer to **string** | The name of the Amazon Web Services region that holds the encryption key. This is optional, and if it is not provided, then the server will attempt to determine the region from the key ARN. | [optional] 
**KmsEncryptionKeyArn** | **string** | The Amazon resource name (ARN) for the KMS key that will be used to encrypt the contents of the passphrase file. This key must exist, and the AWS client must have access to encrypt and decrypt data using this key. | 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 
**SecretID** | **string** | The Amazon Resource Name (ARN) or the user-friendly name of the secret to be retrieved. | 
**SecretFieldName** | **string** | The name of the JSON field whose value is the passphrase that will be used to generate the encryption key for protecting the contents of the encryption settings database. | 
**SecretVersionID** | Pointer to **string** | The unique identifier for the version of the secret to be retrieved. | [optional] 
**SecretVersionStage** | Pointer to **string** | The staging label for the version of the secret to be retrieved. | [optional] 
**EncryptionMetadataFile** | Pointer to **string** | The path to a file that will hold metadata about the encryption performed by this PKCS11 Cipher Stream Provider. | [optional] 
**KeyVaultURI** | **string** | The URI that identifies the Azure Key Vault from which the secret is to be retrieved. | 
**AzureAuthenticationMethod** | **string** | The mechanism used to authenticate to the Azure service. | 
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
**VaultEncryptionMetadataFile** | Pointer to **string** | The path to a file that will hold metadata about the encryption performed by this Vault Cipher Stream Provider. | [optional] 
**TrustStoreFile** | Pointer to **string** | The path to a file containing the information needed to trust the certificate presented by the Vault servers. | [optional] 
**TrustStorePin** | Pointer to **string** | The passphrase needed to access the contents of the trust store. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents. | [optional] 
**TrustStoreType** | Pointer to **string** | The store type for the specified trust store file. The value should likely be one of \&quot;JKS\&quot; or \&quot;PKCS12\&quot;. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Cipher Stream Provider. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Cipher Stream Provider. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddCipherStreamProviderRequest

`func NewAddCipherStreamProviderRequest(providerName string, schemas []EnumthirdPartyCipherStreamProviderSchemaUrn, awsExternalServer string, kmsEncryptionKeyArn string, enabled bool, secretID string, secretFieldName string, keyVaultURI string, azureAuthenticationMethod string, secretName string, passwordFile string, conjurExternalServer string, conjurSecretRelativePath string, sslCertNickname string, vaultSecretPath string, vaultSecretFieldName string, extensionClass string, ) *AddCipherStreamProviderRequest`

NewAddCipherStreamProviderRequest instantiates a new AddCipherStreamProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCipherStreamProviderRequestWithDefaults

`func NewAddCipherStreamProviderRequestWithDefaults() *AddCipherStreamProviderRequest`

NewAddCipherStreamProviderRequestWithDefaults instantiates a new AddCipherStreamProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *AddCipherStreamProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddCipherStreamProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddCipherStreamProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetSchemas

`func (o *AddCipherStreamProviderRequest) GetSchemas() []EnumthirdPartyCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddCipherStreamProviderRequest) GetSchemasOk() (*[]EnumthirdPartyCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddCipherStreamProviderRequest) SetSchemas(v []EnumthirdPartyCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEncryptedPassphraseFile

`func (o *AddCipherStreamProviderRequest) GetEncryptedPassphraseFile() string`

GetEncryptedPassphraseFile returns the EncryptedPassphraseFile field if non-nil, zero value otherwise.

### GetEncryptedPassphraseFileOk

`func (o *AddCipherStreamProviderRequest) GetEncryptedPassphraseFileOk() (*string, bool)`

GetEncryptedPassphraseFileOk returns a tuple with the EncryptedPassphraseFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassphraseFile

`func (o *AddCipherStreamProviderRequest) SetEncryptedPassphraseFile(v string)`

SetEncryptedPassphraseFile sets EncryptedPassphraseFile field to given value.

### HasEncryptedPassphraseFile

`func (o *AddCipherStreamProviderRequest) HasEncryptedPassphraseFile() bool`

HasEncryptedPassphraseFile returns a boolean if a field has been set.

### GetAwsExternalServer

`func (o *AddCipherStreamProviderRequest) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *AddCipherStreamProviderRequest) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *AddCipherStreamProviderRequest) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.


### GetAwsAccessKeyID

`func (o *AddCipherStreamProviderRequest) GetAwsAccessKeyID() string`

GetAwsAccessKeyID returns the AwsAccessKeyID field if non-nil, zero value otherwise.

### GetAwsAccessKeyIDOk

`func (o *AddCipherStreamProviderRequest) GetAwsAccessKeyIDOk() (*string, bool)`

GetAwsAccessKeyIDOk returns a tuple with the AwsAccessKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyID

`func (o *AddCipherStreamProviderRequest) SetAwsAccessKeyID(v string)`

SetAwsAccessKeyID sets AwsAccessKeyID field to given value.

### HasAwsAccessKeyID

`func (o *AddCipherStreamProviderRequest) HasAwsAccessKeyID() bool`

HasAwsAccessKeyID returns a boolean if a field has been set.

### GetAwsSecretAccessKey

`func (o *AddCipherStreamProviderRequest) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *AddCipherStreamProviderRequest) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *AddCipherStreamProviderRequest) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *AddCipherStreamProviderRequest) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### GetAwsRegionName

`func (o *AddCipherStreamProviderRequest) GetAwsRegionName() string`

GetAwsRegionName returns the AwsRegionName field if non-nil, zero value otherwise.

### GetAwsRegionNameOk

`func (o *AddCipherStreamProviderRequest) GetAwsRegionNameOk() (*string, bool)`

GetAwsRegionNameOk returns a tuple with the AwsRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegionName

`func (o *AddCipherStreamProviderRequest) SetAwsRegionName(v string)`

SetAwsRegionName sets AwsRegionName field to given value.

### HasAwsRegionName

`func (o *AddCipherStreamProviderRequest) HasAwsRegionName() bool`

HasAwsRegionName returns a boolean if a field has been set.

### GetKmsEncryptionKeyArn

`func (o *AddCipherStreamProviderRequest) GetKmsEncryptionKeyArn() string`

GetKmsEncryptionKeyArn returns the KmsEncryptionKeyArn field if non-nil, zero value otherwise.

### GetKmsEncryptionKeyArnOk

`func (o *AddCipherStreamProviderRequest) GetKmsEncryptionKeyArnOk() (*string, bool)`

GetKmsEncryptionKeyArnOk returns a tuple with the KmsEncryptionKeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsEncryptionKeyArn

`func (o *AddCipherStreamProviderRequest) SetKmsEncryptionKeyArn(v string)`

SetKmsEncryptionKeyArn sets KmsEncryptionKeyArn field to given value.


### GetDescription

`func (o *AddCipherStreamProviderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCipherStreamProviderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCipherStreamProviderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCipherStreamProviderRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddCipherStreamProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCipherStreamProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCipherStreamProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSecretID

`func (o *AddCipherStreamProviderRequest) GetSecretID() string`

GetSecretID returns the SecretID field if non-nil, zero value otherwise.

### GetSecretIDOk

`func (o *AddCipherStreamProviderRequest) GetSecretIDOk() (*string, bool)`

GetSecretIDOk returns a tuple with the SecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretID

`func (o *AddCipherStreamProviderRequest) SetSecretID(v string)`

SetSecretID sets SecretID field to given value.


### GetSecretFieldName

`func (o *AddCipherStreamProviderRequest) GetSecretFieldName() string`

GetSecretFieldName returns the SecretFieldName field if non-nil, zero value otherwise.

### GetSecretFieldNameOk

`func (o *AddCipherStreamProviderRequest) GetSecretFieldNameOk() (*string, bool)`

GetSecretFieldNameOk returns a tuple with the SecretFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretFieldName

`func (o *AddCipherStreamProviderRequest) SetSecretFieldName(v string)`

SetSecretFieldName sets SecretFieldName field to given value.


### GetSecretVersionID

`func (o *AddCipherStreamProviderRequest) GetSecretVersionID() string`

GetSecretVersionID returns the SecretVersionID field if non-nil, zero value otherwise.

### GetSecretVersionIDOk

`func (o *AddCipherStreamProviderRequest) GetSecretVersionIDOk() (*string, bool)`

GetSecretVersionIDOk returns a tuple with the SecretVersionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVersionID

`func (o *AddCipherStreamProviderRequest) SetSecretVersionID(v string)`

SetSecretVersionID sets SecretVersionID field to given value.

### HasSecretVersionID

`func (o *AddCipherStreamProviderRequest) HasSecretVersionID() bool`

HasSecretVersionID returns a boolean if a field has been set.

### GetSecretVersionStage

`func (o *AddCipherStreamProviderRequest) GetSecretVersionStage() string`

GetSecretVersionStage returns the SecretVersionStage field if non-nil, zero value otherwise.

### GetSecretVersionStageOk

`func (o *AddCipherStreamProviderRequest) GetSecretVersionStageOk() (*string, bool)`

GetSecretVersionStageOk returns a tuple with the SecretVersionStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVersionStage

`func (o *AddCipherStreamProviderRequest) SetSecretVersionStage(v string)`

SetSecretVersionStage sets SecretVersionStage field to given value.

### HasSecretVersionStage

`func (o *AddCipherStreamProviderRequest) HasSecretVersionStage() bool`

HasSecretVersionStage returns a boolean if a field has been set.

### GetEncryptionMetadataFile

`func (o *AddCipherStreamProviderRequest) GetEncryptionMetadataFile() string`

GetEncryptionMetadataFile returns the EncryptionMetadataFile field if non-nil, zero value otherwise.

### GetEncryptionMetadataFileOk

`func (o *AddCipherStreamProviderRequest) GetEncryptionMetadataFileOk() (*string, bool)`

GetEncryptionMetadataFileOk returns a tuple with the EncryptionMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMetadataFile

`func (o *AddCipherStreamProviderRequest) SetEncryptionMetadataFile(v string)`

SetEncryptionMetadataFile sets EncryptionMetadataFile field to given value.

### HasEncryptionMetadataFile

`func (o *AddCipherStreamProviderRequest) HasEncryptionMetadataFile() bool`

HasEncryptionMetadataFile returns a boolean if a field has been set.

### GetKeyVaultURI

`func (o *AddCipherStreamProviderRequest) GetKeyVaultURI() string`

GetKeyVaultURI returns the KeyVaultURI field if non-nil, zero value otherwise.

### GetKeyVaultURIOk

`func (o *AddCipherStreamProviderRequest) GetKeyVaultURIOk() (*string, bool)`

GetKeyVaultURIOk returns a tuple with the KeyVaultURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVaultURI

`func (o *AddCipherStreamProviderRequest) SetKeyVaultURI(v string)`

SetKeyVaultURI sets KeyVaultURI field to given value.


### GetAzureAuthenticationMethod

`func (o *AddCipherStreamProviderRequest) GetAzureAuthenticationMethod() string`

GetAzureAuthenticationMethod returns the AzureAuthenticationMethod field if non-nil, zero value otherwise.

### GetAzureAuthenticationMethodOk

`func (o *AddCipherStreamProviderRequest) GetAzureAuthenticationMethodOk() (*string, bool)`

GetAzureAuthenticationMethodOk returns a tuple with the AzureAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAuthenticationMethod

`func (o *AddCipherStreamProviderRequest) SetAzureAuthenticationMethod(v string)`

SetAzureAuthenticationMethod sets AzureAuthenticationMethod field to given value.


### GetSecretName

`func (o *AddCipherStreamProviderRequest) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *AddCipherStreamProviderRequest) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *AddCipherStreamProviderRequest) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetPasswordFile

`func (o *AddCipherStreamProviderRequest) GetPasswordFile() string`

GetPasswordFile returns the PasswordFile field if non-nil, zero value otherwise.

### GetPasswordFileOk

`func (o *AddCipherStreamProviderRequest) GetPasswordFileOk() (*string, bool)`

GetPasswordFileOk returns a tuple with the PasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFile

`func (o *AddCipherStreamProviderRequest) SetPasswordFile(v string)`

SetPasswordFile sets PasswordFile field to given value.


### GetWaitForPasswordFile

`func (o *AddCipherStreamProviderRequest) GetWaitForPasswordFile() bool`

GetWaitForPasswordFile returns the WaitForPasswordFile field if non-nil, zero value otherwise.

### GetWaitForPasswordFileOk

`func (o *AddCipherStreamProviderRequest) GetWaitForPasswordFileOk() (*bool, bool)`

GetWaitForPasswordFileOk returns a tuple with the WaitForPasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForPasswordFile

`func (o *AddCipherStreamProviderRequest) SetWaitForPasswordFile(v bool)`

SetWaitForPasswordFile sets WaitForPasswordFile field to given value.

### HasWaitForPasswordFile

`func (o *AddCipherStreamProviderRequest) HasWaitForPasswordFile() bool`

HasWaitForPasswordFile returns a boolean if a field has been set.

### GetConjurExternalServer

`func (o *AddCipherStreamProviderRequest) GetConjurExternalServer() string`

GetConjurExternalServer returns the ConjurExternalServer field if non-nil, zero value otherwise.

### GetConjurExternalServerOk

`func (o *AddCipherStreamProviderRequest) GetConjurExternalServerOk() (*string, bool)`

GetConjurExternalServerOk returns a tuple with the ConjurExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurExternalServer

`func (o *AddCipherStreamProviderRequest) SetConjurExternalServer(v string)`

SetConjurExternalServer sets ConjurExternalServer field to given value.


### GetConjurSecretRelativePath

`func (o *AddCipherStreamProviderRequest) GetConjurSecretRelativePath() string`

GetConjurSecretRelativePath returns the ConjurSecretRelativePath field if non-nil, zero value otherwise.

### GetConjurSecretRelativePathOk

`func (o *AddCipherStreamProviderRequest) GetConjurSecretRelativePathOk() (*string, bool)`

GetConjurSecretRelativePathOk returns a tuple with the ConjurSecretRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurSecretRelativePath

`func (o *AddCipherStreamProviderRequest) SetConjurSecretRelativePath(v string)`

SetConjurSecretRelativePath sets ConjurSecretRelativePath field to given value.


### GetPkcs11ProviderClass

`func (o *AddCipherStreamProviderRequest) GetPkcs11ProviderClass() string`

GetPkcs11ProviderClass returns the Pkcs11ProviderClass field if non-nil, zero value otherwise.

### GetPkcs11ProviderClassOk

`func (o *AddCipherStreamProviderRequest) GetPkcs11ProviderClassOk() (*string, bool)`

GetPkcs11ProviderClassOk returns a tuple with the Pkcs11ProviderClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11ProviderClass

`func (o *AddCipherStreamProviderRequest) SetPkcs11ProviderClass(v string)`

SetPkcs11ProviderClass sets Pkcs11ProviderClass field to given value.

### HasPkcs11ProviderClass

`func (o *AddCipherStreamProviderRequest) HasPkcs11ProviderClass() bool`

HasPkcs11ProviderClass returns a boolean if a field has been set.

### GetPkcs11ProviderConfigurationFile

`func (o *AddCipherStreamProviderRequest) GetPkcs11ProviderConfigurationFile() string`

GetPkcs11ProviderConfigurationFile returns the Pkcs11ProviderConfigurationFile field if non-nil, zero value otherwise.

### GetPkcs11ProviderConfigurationFileOk

`func (o *AddCipherStreamProviderRequest) GetPkcs11ProviderConfigurationFileOk() (*string, bool)`

GetPkcs11ProviderConfigurationFileOk returns a tuple with the Pkcs11ProviderConfigurationFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11ProviderConfigurationFile

`func (o *AddCipherStreamProviderRequest) SetPkcs11ProviderConfigurationFile(v string)`

SetPkcs11ProviderConfigurationFile sets Pkcs11ProviderConfigurationFile field to given value.

### HasPkcs11ProviderConfigurationFile

`func (o *AddCipherStreamProviderRequest) HasPkcs11ProviderConfigurationFile() bool`

HasPkcs11ProviderConfigurationFile returns a boolean if a field has been set.

### GetKeyStorePin

`func (o *AddCipherStreamProviderRequest) GetKeyStorePin() string`

GetKeyStorePin returns the KeyStorePin field if non-nil, zero value otherwise.

### GetKeyStorePinOk

`func (o *AddCipherStreamProviderRequest) GetKeyStorePinOk() (*string, bool)`

GetKeyStorePinOk returns a tuple with the KeyStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePin

`func (o *AddCipherStreamProviderRequest) SetKeyStorePin(v string)`

SetKeyStorePin sets KeyStorePin field to given value.

### HasKeyStorePin

`func (o *AddCipherStreamProviderRequest) HasKeyStorePin() bool`

HasKeyStorePin returns a boolean if a field has been set.

### GetKeyStorePinFile

`func (o *AddCipherStreamProviderRequest) GetKeyStorePinFile() string`

GetKeyStorePinFile returns the KeyStorePinFile field if non-nil, zero value otherwise.

### GetKeyStorePinFileOk

`func (o *AddCipherStreamProviderRequest) GetKeyStorePinFileOk() (*string, bool)`

GetKeyStorePinFileOk returns a tuple with the KeyStorePinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePinFile

`func (o *AddCipherStreamProviderRequest) SetKeyStorePinFile(v string)`

SetKeyStorePinFile sets KeyStorePinFile field to given value.

### HasKeyStorePinFile

`func (o *AddCipherStreamProviderRequest) HasKeyStorePinFile() bool`

HasKeyStorePinFile returns a boolean if a field has been set.

### GetKeyStorePinEnvironmentVariable

`func (o *AddCipherStreamProviderRequest) GetKeyStorePinEnvironmentVariable() string`

GetKeyStorePinEnvironmentVariable returns the KeyStorePinEnvironmentVariable field if non-nil, zero value otherwise.

### GetKeyStorePinEnvironmentVariableOk

`func (o *AddCipherStreamProviderRequest) GetKeyStorePinEnvironmentVariableOk() (*string, bool)`

GetKeyStorePinEnvironmentVariableOk returns a tuple with the KeyStorePinEnvironmentVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStorePinEnvironmentVariable

`func (o *AddCipherStreamProviderRequest) SetKeyStorePinEnvironmentVariable(v string)`

SetKeyStorePinEnvironmentVariable sets KeyStorePinEnvironmentVariable field to given value.

### HasKeyStorePinEnvironmentVariable

`func (o *AddCipherStreamProviderRequest) HasKeyStorePinEnvironmentVariable() bool`

HasKeyStorePinEnvironmentVariable returns a boolean if a field has been set.

### GetPkcs11KeyStoreType

`func (o *AddCipherStreamProviderRequest) GetPkcs11KeyStoreType() string`

GetPkcs11KeyStoreType returns the Pkcs11KeyStoreType field if non-nil, zero value otherwise.

### GetPkcs11KeyStoreTypeOk

`func (o *AddCipherStreamProviderRequest) GetPkcs11KeyStoreTypeOk() (*string, bool)`

GetPkcs11KeyStoreTypeOk returns a tuple with the Pkcs11KeyStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs11KeyStoreType

`func (o *AddCipherStreamProviderRequest) SetPkcs11KeyStoreType(v string)`

SetPkcs11KeyStoreType sets Pkcs11KeyStoreType field to given value.

### HasPkcs11KeyStoreType

`func (o *AddCipherStreamProviderRequest) HasPkcs11KeyStoreType() bool`

HasPkcs11KeyStoreType returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *AddCipherStreamProviderRequest) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *AddCipherStreamProviderRequest) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *AddCipherStreamProviderRequest) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.


### GetVaultExternalServer

`func (o *AddCipherStreamProviderRequest) GetVaultExternalServer() string`

GetVaultExternalServer returns the VaultExternalServer field if non-nil, zero value otherwise.

### GetVaultExternalServerOk

`func (o *AddCipherStreamProviderRequest) GetVaultExternalServerOk() (*string, bool)`

GetVaultExternalServerOk returns a tuple with the VaultExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultExternalServer

`func (o *AddCipherStreamProviderRequest) SetVaultExternalServer(v string)`

SetVaultExternalServer sets VaultExternalServer field to given value.

### HasVaultExternalServer

`func (o *AddCipherStreamProviderRequest) HasVaultExternalServer() bool`

HasVaultExternalServer returns a boolean if a field has been set.

### GetVaultServerBaseURI

`func (o *AddCipherStreamProviderRequest) GetVaultServerBaseURI() []string`

GetVaultServerBaseURI returns the VaultServerBaseURI field if non-nil, zero value otherwise.

### GetVaultServerBaseURIOk

`func (o *AddCipherStreamProviderRequest) GetVaultServerBaseURIOk() (*[]string, bool)`

GetVaultServerBaseURIOk returns a tuple with the VaultServerBaseURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultServerBaseURI

`func (o *AddCipherStreamProviderRequest) SetVaultServerBaseURI(v []string)`

SetVaultServerBaseURI sets VaultServerBaseURI field to given value.

### HasVaultServerBaseURI

`func (o *AddCipherStreamProviderRequest) HasVaultServerBaseURI() bool`

HasVaultServerBaseURI returns a boolean if a field has been set.

### GetVaultAuthenticationMethod

`func (o *AddCipherStreamProviderRequest) GetVaultAuthenticationMethod() string`

GetVaultAuthenticationMethod returns the VaultAuthenticationMethod field if non-nil, zero value otherwise.

### GetVaultAuthenticationMethodOk

`func (o *AddCipherStreamProviderRequest) GetVaultAuthenticationMethodOk() (*string, bool)`

GetVaultAuthenticationMethodOk returns a tuple with the VaultAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAuthenticationMethod

`func (o *AddCipherStreamProviderRequest) SetVaultAuthenticationMethod(v string)`

SetVaultAuthenticationMethod sets VaultAuthenticationMethod field to given value.

### HasVaultAuthenticationMethod

`func (o *AddCipherStreamProviderRequest) HasVaultAuthenticationMethod() bool`

HasVaultAuthenticationMethod returns a boolean if a field has been set.

### GetVaultSecretPath

`func (o *AddCipherStreamProviderRequest) GetVaultSecretPath() string`

GetVaultSecretPath returns the VaultSecretPath field if non-nil, zero value otherwise.

### GetVaultSecretPathOk

`func (o *AddCipherStreamProviderRequest) GetVaultSecretPathOk() (*string, bool)`

GetVaultSecretPathOk returns a tuple with the VaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretPath

`func (o *AddCipherStreamProviderRequest) SetVaultSecretPath(v string)`

SetVaultSecretPath sets VaultSecretPath field to given value.


### GetVaultSecretFieldName

`func (o *AddCipherStreamProviderRequest) GetVaultSecretFieldName() string`

GetVaultSecretFieldName returns the VaultSecretFieldName field if non-nil, zero value otherwise.

### GetVaultSecretFieldNameOk

`func (o *AddCipherStreamProviderRequest) GetVaultSecretFieldNameOk() (*string, bool)`

GetVaultSecretFieldNameOk returns a tuple with the VaultSecretFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretFieldName

`func (o *AddCipherStreamProviderRequest) SetVaultSecretFieldName(v string)`

SetVaultSecretFieldName sets VaultSecretFieldName field to given value.


### GetVaultEncryptionMetadataFile

`func (o *AddCipherStreamProviderRequest) GetVaultEncryptionMetadataFile() string`

GetVaultEncryptionMetadataFile returns the VaultEncryptionMetadataFile field if non-nil, zero value otherwise.

### GetVaultEncryptionMetadataFileOk

`func (o *AddCipherStreamProviderRequest) GetVaultEncryptionMetadataFileOk() (*string, bool)`

GetVaultEncryptionMetadataFileOk returns a tuple with the VaultEncryptionMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultEncryptionMetadataFile

`func (o *AddCipherStreamProviderRequest) SetVaultEncryptionMetadataFile(v string)`

SetVaultEncryptionMetadataFile sets VaultEncryptionMetadataFile field to given value.

### HasVaultEncryptionMetadataFile

`func (o *AddCipherStreamProviderRequest) HasVaultEncryptionMetadataFile() bool`

HasVaultEncryptionMetadataFile returns a boolean if a field has been set.

### GetTrustStoreFile

`func (o *AddCipherStreamProviderRequest) GetTrustStoreFile() string`

GetTrustStoreFile returns the TrustStoreFile field if non-nil, zero value otherwise.

### GetTrustStoreFileOk

`func (o *AddCipherStreamProviderRequest) GetTrustStoreFileOk() (*string, bool)`

GetTrustStoreFileOk returns a tuple with the TrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreFile

`func (o *AddCipherStreamProviderRequest) SetTrustStoreFile(v string)`

SetTrustStoreFile sets TrustStoreFile field to given value.

### HasTrustStoreFile

`func (o *AddCipherStreamProviderRequest) HasTrustStoreFile() bool`

HasTrustStoreFile returns a boolean if a field has been set.

### GetTrustStorePin

`func (o *AddCipherStreamProviderRequest) GetTrustStorePin() string`

GetTrustStorePin returns the TrustStorePin field if non-nil, zero value otherwise.

### GetTrustStorePinOk

`func (o *AddCipherStreamProviderRequest) GetTrustStorePinOk() (*string, bool)`

GetTrustStorePinOk returns a tuple with the TrustStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePin

`func (o *AddCipherStreamProviderRequest) SetTrustStorePin(v string)`

SetTrustStorePin sets TrustStorePin field to given value.

### HasTrustStorePin

`func (o *AddCipherStreamProviderRequest) HasTrustStorePin() bool`

HasTrustStorePin returns a boolean if a field has been set.

### GetTrustStoreType

`func (o *AddCipherStreamProviderRequest) GetTrustStoreType() string`

GetTrustStoreType returns the TrustStoreType field if non-nil, zero value otherwise.

### GetTrustStoreTypeOk

`func (o *AddCipherStreamProviderRequest) GetTrustStoreTypeOk() (*string, bool)`

GetTrustStoreTypeOk returns a tuple with the TrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreType

`func (o *AddCipherStreamProviderRequest) SetTrustStoreType(v string)`

SetTrustStoreType sets TrustStoreType field to given value.

### HasTrustStoreType

`func (o *AddCipherStreamProviderRequest) HasTrustStoreType() bool`

HasTrustStoreType returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddCipherStreamProviderRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddCipherStreamProviderRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddCipherStreamProviderRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddCipherStreamProviderRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddCipherStreamProviderRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddCipherStreamProviderRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddCipherStreamProviderRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


