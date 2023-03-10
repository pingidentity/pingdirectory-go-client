# AddPasswordStorageScheme200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Password Storage Scheme | 
**Schemas** | [**[]EnumscryptPasswordStorageSchemeSchemaUrn**](EnumscryptPasswordStorageSchemeSchemaUrn.md) |  | 
**IterationCount** | **int32** | Specifies the number of iterations to use when encoding passwords. The value must be greater than or equal to 1000. | 
**ParallelismFactor** | **int32** | The number of concurrent threads that will be used in the course of encoding each password. | 
**MemoryUsageKb** | **int32** | The number of kilobytes of memory that must be used in the course of encoding each password. | 
**SaltLengthBytes** | **int32** | Specifies the number of bytes to use for the generated salt. The value must be greater than or equal to 8. | 
**DerivedKeyLengthBytes** | **int32** | Specifies the number of bytes to use for the derived key. The value must be greater than or equal to 8. | 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 
**PasswordEncodingMechanism** | Pointer to [**EnumpasswordStorageSchemePasswordEncodingMechanismProp**](EnumpasswordStorageSchemePasswordEncodingMechanismProp.md) |  | [optional] 
**NumDigestRounds** | Pointer to **int32** | Specifies the number of digest rounds to use for the SHA-2 encodings. This will not be used for the legacy or MD5-based encodings. | [optional] 
**MaxPasswordLength** | Pointer to **int32** | Specifies the maximum allowed length, in bytes, for passwords encoded with this scheme, which can help mitigate denial of service attacks from clients that attempt to bind with very long passwords. | [optional] 
**VaultExternalServer** | **string** | An external server definition with information needed to connect and authenticate to the Vault instance containing the passphrase. | 
**DefaultField** | Pointer to **string** | The default name of the field in JSON objects contained in the AWS Secrets Manager service that contains the password for the target user. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Enhanced Password Storage Scheme. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Enhanced Password Storage Scheme. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**DigestAlgorithm** | Pointer to [**EnumpasswordStorageSchemeDigestAlgorithmProp**](EnumpasswordStorageSchemeDigestAlgorithmProp.md) |  | [optional] 
**EncryptionSettingsDefinitionID** | Pointer to **string** | The identifier for the encryption settings definition that should be used to derive the encryption key to use when encrypting new passwords. If this is not provided, the server&#39;s preferred encryption settings definition will be used. | [optional] 
**BcryptCostFactor** | Pointer to **int32** | Specifies the cost factor to use when encoding passwords with Bcrypt. A higher cost factor requires more processing to generate a password, which makes attacks against the password more expensive. | [optional] 
**AwsExternalServer** | **string** | The external server with information to use when interacting with the AWS Secrets Manager service. | 
**KeyVaultURI** | **string** | The URI that identifies the Azure Key Vault from which the secret is to be retrieved. | 
**AzureAuthenticationMethod** | **string** | The mechanism used to authenticate to the Azure service. | 
**HttpProxyExternalServer** | Pointer to **string** | A reference to an HTTP proxy server that should be used for requests sent to the Azure service. | [optional] 
**ConjurExternalServer** | **string** | An external server definition with information needed to connect and authenticate to the Conjur instance containing user passwords. | 
**ScryptCpuMemoryCostFactorExponent** | Pointer to **int32** | Specifies the exponent that should be used for the CPU/memory cost factor. The cost factor must be a power of two, so the value of this property represents the power to which two is raised. The CPU/memory cost factor specifies the number of iterations required for encoding the password, and also affects the amount of memory required during processing. A higher cost factor requires more processing and more memory to generate a password, which makes attacks against the password more expensive. | [optional] 
**ScryptBlockSize** | Pointer to **int32** | Specifies the block size for the digest that will be used in the course of encoding passwords. Increasing the block size while keeping the CPU/memory cost factor constant will increase the amount of memory required to encode a password, but it also increases the ratio of sequential memory access to random memory access (and sequential memory access is generally faster than random memory access). | [optional] 
**ScryptParallelizationParameter** | Pointer to **int32** | Specifies the number of times that scrypt has to perform the entire encoding process to produce the final result. | [optional] 

## Methods

### NewAddPasswordStorageScheme200Response

`func NewAddPasswordStorageScheme200Response(id string, schemas []EnumscryptPasswordStorageSchemeSchemaUrn, iterationCount int32, parallelismFactor int32, memoryUsageKb int32, saltLengthBytes int32, derivedKeyLengthBytes int32, enabled bool, vaultExternalServer string, extensionClass string, awsExternalServer string, keyVaultURI string, azureAuthenticationMethod string, conjurExternalServer string, ) *AddPasswordStorageScheme200Response`

NewAddPasswordStorageScheme200Response instantiates a new AddPasswordStorageScheme200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPasswordStorageScheme200ResponseWithDefaults

`func NewAddPasswordStorageScheme200ResponseWithDefaults() *AddPasswordStorageScheme200Response`

NewAddPasswordStorageScheme200ResponseWithDefaults instantiates a new AddPasswordStorageScheme200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *AddPasswordStorageScheme200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddPasswordStorageScheme200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddPasswordStorageScheme200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddPasswordStorageScheme200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddPasswordStorageScheme200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddPasswordStorageScheme200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddPasswordStorageScheme200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddPasswordStorageScheme200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *AddPasswordStorageScheme200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddPasswordStorageScheme200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddPasswordStorageScheme200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddPasswordStorageScheme200Response) GetSchemas() []EnumscryptPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPasswordStorageScheme200Response) GetSchemasOk() (*[]EnumscryptPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPasswordStorageScheme200Response) SetSchemas(v []EnumscryptPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIterationCount

`func (o *AddPasswordStorageScheme200Response) GetIterationCount() int32`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *AddPasswordStorageScheme200Response) GetIterationCountOk() (*int32, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *AddPasswordStorageScheme200Response) SetIterationCount(v int32)`

SetIterationCount sets IterationCount field to given value.


### GetParallelismFactor

`func (o *AddPasswordStorageScheme200Response) GetParallelismFactor() int32`

GetParallelismFactor returns the ParallelismFactor field if non-nil, zero value otherwise.

### GetParallelismFactorOk

`func (o *AddPasswordStorageScheme200Response) GetParallelismFactorOk() (*int32, bool)`

GetParallelismFactorOk returns a tuple with the ParallelismFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelismFactor

`func (o *AddPasswordStorageScheme200Response) SetParallelismFactor(v int32)`

SetParallelismFactor sets ParallelismFactor field to given value.


### GetMemoryUsageKb

`func (o *AddPasswordStorageScheme200Response) GetMemoryUsageKb() int32`

GetMemoryUsageKb returns the MemoryUsageKb field if non-nil, zero value otherwise.

### GetMemoryUsageKbOk

`func (o *AddPasswordStorageScheme200Response) GetMemoryUsageKbOk() (*int32, bool)`

GetMemoryUsageKbOk returns a tuple with the MemoryUsageKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageKb

`func (o *AddPasswordStorageScheme200Response) SetMemoryUsageKb(v int32)`

SetMemoryUsageKb sets MemoryUsageKb field to given value.


### GetSaltLengthBytes

`func (o *AddPasswordStorageScheme200Response) GetSaltLengthBytes() int32`

GetSaltLengthBytes returns the SaltLengthBytes field if non-nil, zero value otherwise.

### GetSaltLengthBytesOk

`func (o *AddPasswordStorageScheme200Response) GetSaltLengthBytesOk() (*int32, bool)`

GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLengthBytes

`func (o *AddPasswordStorageScheme200Response) SetSaltLengthBytes(v int32)`

SetSaltLengthBytes sets SaltLengthBytes field to given value.


### GetDerivedKeyLengthBytes

`func (o *AddPasswordStorageScheme200Response) GetDerivedKeyLengthBytes() int32`

GetDerivedKeyLengthBytes returns the DerivedKeyLengthBytes field if non-nil, zero value otherwise.

### GetDerivedKeyLengthBytesOk

`func (o *AddPasswordStorageScheme200Response) GetDerivedKeyLengthBytesOk() (*int32, bool)`

GetDerivedKeyLengthBytesOk returns a tuple with the DerivedKeyLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedKeyLengthBytes

`func (o *AddPasswordStorageScheme200Response) SetDerivedKeyLengthBytes(v int32)`

SetDerivedKeyLengthBytes sets DerivedKeyLengthBytes field to given value.


### GetDescription

`func (o *AddPasswordStorageScheme200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPasswordStorageScheme200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPasswordStorageScheme200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPasswordStorageScheme200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPasswordStorageScheme200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPasswordStorageScheme200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPasswordStorageScheme200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPasswordEncodingMechanism

`func (o *AddPasswordStorageScheme200Response) GetPasswordEncodingMechanism() EnumpasswordStorageSchemePasswordEncodingMechanismProp`

GetPasswordEncodingMechanism returns the PasswordEncodingMechanism field if non-nil, zero value otherwise.

### GetPasswordEncodingMechanismOk

`func (o *AddPasswordStorageScheme200Response) GetPasswordEncodingMechanismOk() (*EnumpasswordStorageSchemePasswordEncodingMechanismProp, bool)`

GetPasswordEncodingMechanismOk returns a tuple with the PasswordEncodingMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordEncodingMechanism

`func (o *AddPasswordStorageScheme200Response) SetPasswordEncodingMechanism(v EnumpasswordStorageSchemePasswordEncodingMechanismProp)`

SetPasswordEncodingMechanism sets PasswordEncodingMechanism field to given value.

### HasPasswordEncodingMechanism

`func (o *AddPasswordStorageScheme200Response) HasPasswordEncodingMechanism() bool`

HasPasswordEncodingMechanism returns a boolean if a field has been set.

### GetNumDigestRounds

`func (o *AddPasswordStorageScheme200Response) GetNumDigestRounds() int32`

GetNumDigestRounds returns the NumDigestRounds field if non-nil, zero value otherwise.

### GetNumDigestRoundsOk

`func (o *AddPasswordStorageScheme200Response) GetNumDigestRoundsOk() (*int32, bool)`

GetNumDigestRoundsOk returns a tuple with the NumDigestRounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDigestRounds

`func (o *AddPasswordStorageScheme200Response) SetNumDigestRounds(v int32)`

SetNumDigestRounds sets NumDigestRounds field to given value.

### HasNumDigestRounds

`func (o *AddPasswordStorageScheme200Response) HasNumDigestRounds() bool`

HasNumDigestRounds returns a boolean if a field has been set.

### GetMaxPasswordLength

`func (o *AddPasswordStorageScheme200Response) GetMaxPasswordLength() int32`

GetMaxPasswordLength returns the MaxPasswordLength field if non-nil, zero value otherwise.

### GetMaxPasswordLengthOk

`func (o *AddPasswordStorageScheme200Response) GetMaxPasswordLengthOk() (*int32, bool)`

GetMaxPasswordLengthOk returns a tuple with the MaxPasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPasswordLength

`func (o *AddPasswordStorageScheme200Response) SetMaxPasswordLength(v int32)`

SetMaxPasswordLength sets MaxPasswordLength field to given value.

### HasMaxPasswordLength

`func (o *AddPasswordStorageScheme200Response) HasMaxPasswordLength() bool`

HasMaxPasswordLength returns a boolean if a field has been set.

### GetVaultExternalServer

`func (o *AddPasswordStorageScheme200Response) GetVaultExternalServer() string`

GetVaultExternalServer returns the VaultExternalServer field if non-nil, zero value otherwise.

### GetVaultExternalServerOk

`func (o *AddPasswordStorageScheme200Response) GetVaultExternalServerOk() (*string, bool)`

GetVaultExternalServerOk returns a tuple with the VaultExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultExternalServer

`func (o *AddPasswordStorageScheme200Response) SetVaultExternalServer(v string)`

SetVaultExternalServer sets VaultExternalServer field to given value.


### GetDefaultField

`func (o *AddPasswordStorageScheme200Response) GetDefaultField() string`

GetDefaultField returns the DefaultField field if non-nil, zero value otherwise.

### GetDefaultFieldOk

`func (o *AddPasswordStorageScheme200Response) GetDefaultFieldOk() (*string, bool)`

GetDefaultFieldOk returns a tuple with the DefaultField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultField

`func (o *AddPasswordStorageScheme200Response) SetDefaultField(v string)`

SetDefaultField sets DefaultField field to given value.

### HasDefaultField

`func (o *AddPasswordStorageScheme200Response) HasDefaultField() bool`

HasDefaultField returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddPasswordStorageScheme200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddPasswordStorageScheme200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddPasswordStorageScheme200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddPasswordStorageScheme200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddPasswordStorageScheme200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddPasswordStorageScheme200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddPasswordStorageScheme200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDigestAlgorithm

`func (o *AddPasswordStorageScheme200Response) GetDigestAlgorithm() EnumpasswordStorageSchemeDigestAlgorithmProp`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *AddPasswordStorageScheme200Response) GetDigestAlgorithmOk() (*EnumpasswordStorageSchemeDigestAlgorithmProp, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *AddPasswordStorageScheme200Response) SetDigestAlgorithm(v EnumpasswordStorageSchemeDigestAlgorithmProp)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *AddPasswordStorageScheme200Response) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetEncryptionSettingsDefinitionID

`func (o *AddPasswordStorageScheme200Response) GetEncryptionSettingsDefinitionID() string`

GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field if non-nil, zero value otherwise.

### GetEncryptionSettingsDefinitionIDOk

`func (o *AddPasswordStorageScheme200Response) GetEncryptionSettingsDefinitionIDOk() (*string, bool)`

GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSettingsDefinitionID

`func (o *AddPasswordStorageScheme200Response) SetEncryptionSettingsDefinitionID(v string)`

SetEncryptionSettingsDefinitionID sets EncryptionSettingsDefinitionID field to given value.

### HasEncryptionSettingsDefinitionID

`func (o *AddPasswordStorageScheme200Response) HasEncryptionSettingsDefinitionID() bool`

HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.

### GetBcryptCostFactor

`func (o *AddPasswordStorageScheme200Response) GetBcryptCostFactor() int32`

GetBcryptCostFactor returns the BcryptCostFactor field if non-nil, zero value otherwise.

### GetBcryptCostFactorOk

`func (o *AddPasswordStorageScheme200Response) GetBcryptCostFactorOk() (*int32, bool)`

GetBcryptCostFactorOk returns a tuple with the BcryptCostFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcryptCostFactor

`func (o *AddPasswordStorageScheme200Response) SetBcryptCostFactor(v int32)`

SetBcryptCostFactor sets BcryptCostFactor field to given value.

### HasBcryptCostFactor

`func (o *AddPasswordStorageScheme200Response) HasBcryptCostFactor() bool`

HasBcryptCostFactor returns a boolean if a field has been set.

### GetAwsExternalServer

`func (o *AddPasswordStorageScheme200Response) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *AddPasswordStorageScheme200Response) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *AddPasswordStorageScheme200Response) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.


### GetKeyVaultURI

`func (o *AddPasswordStorageScheme200Response) GetKeyVaultURI() string`

GetKeyVaultURI returns the KeyVaultURI field if non-nil, zero value otherwise.

### GetKeyVaultURIOk

`func (o *AddPasswordStorageScheme200Response) GetKeyVaultURIOk() (*string, bool)`

GetKeyVaultURIOk returns a tuple with the KeyVaultURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVaultURI

`func (o *AddPasswordStorageScheme200Response) SetKeyVaultURI(v string)`

SetKeyVaultURI sets KeyVaultURI field to given value.


### GetAzureAuthenticationMethod

`func (o *AddPasswordStorageScheme200Response) GetAzureAuthenticationMethod() string`

GetAzureAuthenticationMethod returns the AzureAuthenticationMethod field if non-nil, zero value otherwise.

### GetAzureAuthenticationMethodOk

`func (o *AddPasswordStorageScheme200Response) GetAzureAuthenticationMethodOk() (*string, bool)`

GetAzureAuthenticationMethodOk returns a tuple with the AzureAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAuthenticationMethod

`func (o *AddPasswordStorageScheme200Response) SetAzureAuthenticationMethod(v string)`

SetAzureAuthenticationMethod sets AzureAuthenticationMethod field to given value.


### GetHttpProxyExternalServer

`func (o *AddPasswordStorageScheme200Response) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *AddPasswordStorageScheme200Response) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *AddPasswordStorageScheme200Response) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *AddPasswordStorageScheme200Response) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetConjurExternalServer

`func (o *AddPasswordStorageScheme200Response) GetConjurExternalServer() string`

GetConjurExternalServer returns the ConjurExternalServer field if non-nil, zero value otherwise.

### GetConjurExternalServerOk

`func (o *AddPasswordStorageScheme200Response) GetConjurExternalServerOk() (*string, bool)`

GetConjurExternalServerOk returns a tuple with the ConjurExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurExternalServer

`func (o *AddPasswordStorageScheme200Response) SetConjurExternalServer(v string)`

SetConjurExternalServer sets ConjurExternalServer field to given value.


### GetScryptCpuMemoryCostFactorExponent

`func (o *AddPasswordStorageScheme200Response) GetScryptCpuMemoryCostFactorExponent() int32`

GetScryptCpuMemoryCostFactorExponent returns the ScryptCpuMemoryCostFactorExponent field if non-nil, zero value otherwise.

### GetScryptCpuMemoryCostFactorExponentOk

`func (o *AddPasswordStorageScheme200Response) GetScryptCpuMemoryCostFactorExponentOk() (*int32, bool)`

GetScryptCpuMemoryCostFactorExponentOk returns a tuple with the ScryptCpuMemoryCostFactorExponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScryptCpuMemoryCostFactorExponent

`func (o *AddPasswordStorageScheme200Response) SetScryptCpuMemoryCostFactorExponent(v int32)`

SetScryptCpuMemoryCostFactorExponent sets ScryptCpuMemoryCostFactorExponent field to given value.

### HasScryptCpuMemoryCostFactorExponent

`func (o *AddPasswordStorageScheme200Response) HasScryptCpuMemoryCostFactorExponent() bool`

HasScryptCpuMemoryCostFactorExponent returns a boolean if a field has been set.

### GetScryptBlockSize

`func (o *AddPasswordStorageScheme200Response) GetScryptBlockSize() int32`

GetScryptBlockSize returns the ScryptBlockSize field if non-nil, zero value otherwise.

### GetScryptBlockSizeOk

`func (o *AddPasswordStorageScheme200Response) GetScryptBlockSizeOk() (*int32, bool)`

GetScryptBlockSizeOk returns a tuple with the ScryptBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScryptBlockSize

`func (o *AddPasswordStorageScheme200Response) SetScryptBlockSize(v int32)`

SetScryptBlockSize sets ScryptBlockSize field to given value.

### HasScryptBlockSize

`func (o *AddPasswordStorageScheme200Response) HasScryptBlockSize() bool`

HasScryptBlockSize returns a boolean if a field has been set.

### GetScryptParallelizationParameter

`func (o *AddPasswordStorageScheme200Response) GetScryptParallelizationParameter() int32`

GetScryptParallelizationParameter returns the ScryptParallelizationParameter field if non-nil, zero value otherwise.

### GetScryptParallelizationParameterOk

`func (o *AddPasswordStorageScheme200Response) GetScryptParallelizationParameterOk() (*int32, bool)`

GetScryptParallelizationParameterOk returns a tuple with the ScryptParallelizationParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScryptParallelizationParameter

`func (o *AddPasswordStorageScheme200Response) SetScryptParallelizationParameter(v int32)`

SetScryptParallelizationParameter sets ScryptParallelizationParameter field to given value.

### HasScryptParallelizationParameter

`func (o *AddPasswordStorageScheme200Response) HasScryptParallelizationParameter() bool`

HasScryptParallelizationParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


