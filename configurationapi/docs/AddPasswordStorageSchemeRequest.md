# AddPasswordStorageSchemeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemeName** | **string** | Name of the new Password Storage Scheme | 
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

### NewAddPasswordStorageSchemeRequest

`func NewAddPasswordStorageSchemeRequest(schemeName string, schemas []EnumscryptPasswordStorageSchemeSchemaUrn, iterationCount int32, parallelismFactor int32, memoryUsageKb int32, saltLengthBytes int32, derivedKeyLengthBytes int32, enabled bool, vaultExternalServer string, extensionClass string, awsExternalServer string, keyVaultURI string, azureAuthenticationMethod string, conjurExternalServer string, ) *AddPasswordStorageSchemeRequest`

NewAddPasswordStorageSchemeRequest instantiates a new AddPasswordStorageSchemeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPasswordStorageSchemeRequestWithDefaults

`func NewAddPasswordStorageSchemeRequestWithDefaults() *AddPasswordStorageSchemeRequest`

NewAddPasswordStorageSchemeRequestWithDefaults instantiates a new AddPasswordStorageSchemeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemeName

`func (o *AddPasswordStorageSchemeRequest) GetSchemeName() string`

GetSchemeName returns the SchemeName field if non-nil, zero value otherwise.

### GetSchemeNameOk

`func (o *AddPasswordStorageSchemeRequest) GetSchemeNameOk() (*string, bool)`

GetSchemeNameOk returns a tuple with the SchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeName

`func (o *AddPasswordStorageSchemeRequest) SetSchemeName(v string)`

SetSchemeName sets SchemeName field to given value.


### GetSchemas

`func (o *AddPasswordStorageSchemeRequest) GetSchemas() []EnumscryptPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPasswordStorageSchemeRequest) GetSchemasOk() (*[]EnumscryptPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPasswordStorageSchemeRequest) SetSchemas(v []EnumscryptPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIterationCount

`func (o *AddPasswordStorageSchemeRequest) GetIterationCount() int32`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *AddPasswordStorageSchemeRequest) GetIterationCountOk() (*int32, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *AddPasswordStorageSchemeRequest) SetIterationCount(v int32)`

SetIterationCount sets IterationCount field to given value.


### GetParallelismFactor

`func (o *AddPasswordStorageSchemeRequest) GetParallelismFactor() int32`

GetParallelismFactor returns the ParallelismFactor field if non-nil, zero value otherwise.

### GetParallelismFactorOk

`func (o *AddPasswordStorageSchemeRequest) GetParallelismFactorOk() (*int32, bool)`

GetParallelismFactorOk returns a tuple with the ParallelismFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelismFactor

`func (o *AddPasswordStorageSchemeRequest) SetParallelismFactor(v int32)`

SetParallelismFactor sets ParallelismFactor field to given value.


### GetMemoryUsageKb

`func (o *AddPasswordStorageSchemeRequest) GetMemoryUsageKb() int32`

GetMemoryUsageKb returns the MemoryUsageKb field if non-nil, zero value otherwise.

### GetMemoryUsageKbOk

`func (o *AddPasswordStorageSchemeRequest) GetMemoryUsageKbOk() (*int32, bool)`

GetMemoryUsageKbOk returns a tuple with the MemoryUsageKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageKb

`func (o *AddPasswordStorageSchemeRequest) SetMemoryUsageKb(v int32)`

SetMemoryUsageKb sets MemoryUsageKb field to given value.


### GetSaltLengthBytes

`func (o *AddPasswordStorageSchemeRequest) GetSaltLengthBytes() int32`

GetSaltLengthBytes returns the SaltLengthBytes field if non-nil, zero value otherwise.

### GetSaltLengthBytesOk

`func (o *AddPasswordStorageSchemeRequest) GetSaltLengthBytesOk() (*int32, bool)`

GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLengthBytes

`func (o *AddPasswordStorageSchemeRequest) SetSaltLengthBytes(v int32)`

SetSaltLengthBytes sets SaltLengthBytes field to given value.


### GetDerivedKeyLengthBytes

`func (o *AddPasswordStorageSchemeRequest) GetDerivedKeyLengthBytes() int32`

GetDerivedKeyLengthBytes returns the DerivedKeyLengthBytes field if non-nil, zero value otherwise.

### GetDerivedKeyLengthBytesOk

`func (o *AddPasswordStorageSchemeRequest) GetDerivedKeyLengthBytesOk() (*int32, bool)`

GetDerivedKeyLengthBytesOk returns a tuple with the DerivedKeyLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedKeyLengthBytes

`func (o *AddPasswordStorageSchemeRequest) SetDerivedKeyLengthBytes(v int32)`

SetDerivedKeyLengthBytes sets DerivedKeyLengthBytes field to given value.


### GetDescription

`func (o *AddPasswordStorageSchemeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPasswordStorageSchemeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPasswordStorageSchemeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPasswordStorageSchemeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPasswordStorageSchemeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPasswordStorageSchemeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPasswordStorageSchemeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPasswordEncodingMechanism

`func (o *AddPasswordStorageSchemeRequest) GetPasswordEncodingMechanism() EnumpasswordStorageSchemePasswordEncodingMechanismProp`

GetPasswordEncodingMechanism returns the PasswordEncodingMechanism field if non-nil, zero value otherwise.

### GetPasswordEncodingMechanismOk

`func (o *AddPasswordStorageSchemeRequest) GetPasswordEncodingMechanismOk() (*EnumpasswordStorageSchemePasswordEncodingMechanismProp, bool)`

GetPasswordEncodingMechanismOk returns a tuple with the PasswordEncodingMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordEncodingMechanism

`func (o *AddPasswordStorageSchemeRequest) SetPasswordEncodingMechanism(v EnumpasswordStorageSchemePasswordEncodingMechanismProp)`

SetPasswordEncodingMechanism sets PasswordEncodingMechanism field to given value.

### HasPasswordEncodingMechanism

`func (o *AddPasswordStorageSchemeRequest) HasPasswordEncodingMechanism() bool`

HasPasswordEncodingMechanism returns a boolean if a field has been set.

### GetNumDigestRounds

`func (o *AddPasswordStorageSchemeRequest) GetNumDigestRounds() int32`

GetNumDigestRounds returns the NumDigestRounds field if non-nil, zero value otherwise.

### GetNumDigestRoundsOk

`func (o *AddPasswordStorageSchemeRequest) GetNumDigestRoundsOk() (*int32, bool)`

GetNumDigestRoundsOk returns a tuple with the NumDigestRounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDigestRounds

`func (o *AddPasswordStorageSchemeRequest) SetNumDigestRounds(v int32)`

SetNumDigestRounds sets NumDigestRounds field to given value.

### HasNumDigestRounds

`func (o *AddPasswordStorageSchemeRequest) HasNumDigestRounds() bool`

HasNumDigestRounds returns a boolean if a field has been set.

### GetMaxPasswordLength

`func (o *AddPasswordStorageSchemeRequest) GetMaxPasswordLength() int32`

GetMaxPasswordLength returns the MaxPasswordLength field if non-nil, zero value otherwise.

### GetMaxPasswordLengthOk

`func (o *AddPasswordStorageSchemeRequest) GetMaxPasswordLengthOk() (*int32, bool)`

GetMaxPasswordLengthOk returns a tuple with the MaxPasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPasswordLength

`func (o *AddPasswordStorageSchemeRequest) SetMaxPasswordLength(v int32)`

SetMaxPasswordLength sets MaxPasswordLength field to given value.

### HasMaxPasswordLength

`func (o *AddPasswordStorageSchemeRequest) HasMaxPasswordLength() bool`

HasMaxPasswordLength returns a boolean if a field has been set.

### GetVaultExternalServer

`func (o *AddPasswordStorageSchemeRequest) GetVaultExternalServer() string`

GetVaultExternalServer returns the VaultExternalServer field if non-nil, zero value otherwise.

### GetVaultExternalServerOk

`func (o *AddPasswordStorageSchemeRequest) GetVaultExternalServerOk() (*string, bool)`

GetVaultExternalServerOk returns a tuple with the VaultExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultExternalServer

`func (o *AddPasswordStorageSchemeRequest) SetVaultExternalServer(v string)`

SetVaultExternalServer sets VaultExternalServer field to given value.


### GetDefaultField

`func (o *AddPasswordStorageSchemeRequest) GetDefaultField() string`

GetDefaultField returns the DefaultField field if non-nil, zero value otherwise.

### GetDefaultFieldOk

`func (o *AddPasswordStorageSchemeRequest) GetDefaultFieldOk() (*string, bool)`

GetDefaultFieldOk returns a tuple with the DefaultField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultField

`func (o *AddPasswordStorageSchemeRequest) SetDefaultField(v string)`

SetDefaultField sets DefaultField field to given value.

### HasDefaultField

`func (o *AddPasswordStorageSchemeRequest) HasDefaultField() bool`

HasDefaultField returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddPasswordStorageSchemeRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddPasswordStorageSchemeRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddPasswordStorageSchemeRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddPasswordStorageSchemeRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddPasswordStorageSchemeRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddPasswordStorageSchemeRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddPasswordStorageSchemeRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDigestAlgorithm

`func (o *AddPasswordStorageSchemeRequest) GetDigestAlgorithm() EnumpasswordStorageSchemeDigestAlgorithmProp`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *AddPasswordStorageSchemeRequest) GetDigestAlgorithmOk() (*EnumpasswordStorageSchemeDigestAlgorithmProp, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *AddPasswordStorageSchemeRequest) SetDigestAlgorithm(v EnumpasswordStorageSchemeDigestAlgorithmProp)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *AddPasswordStorageSchemeRequest) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetEncryptionSettingsDefinitionID

`func (o *AddPasswordStorageSchemeRequest) GetEncryptionSettingsDefinitionID() string`

GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field if non-nil, zero value otherwise.

### GetEncryptionSettingsDefinitionIDOk

`func (o *AddPasswordStorageSchemeRequest) GetEncryptionSettingsDefinitionIDOk() (*string, bool)`

GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSettingsDefinitionID

`func (o *AddPasswordStorageSchemeRequest) SetEncryptionSettingsDefinitionID(v string)`

SetEncryptionSettingsDefinitionID sets EncryptionSettingsDefinitionID field to given value.

### HasEncryptionSettingsDefinitionID

`func (o *AddPasswordStorageSchemeRequest) HasEncryptionSettingsDefinitionID() bool`

HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.

### GetBcryptCostFactor

`func (o *AddPasswordStorageSchemeRequest) GetBcryptCostFactor() int32`

GetBcryptCostFactor returns the BcryptCostFactor field if non-nil, zero value otherwise.

### GetBcryptCostFactorOk

`func (o *AddPasswordStorageSchemeRequest) GetBcryptCostFactorOk() (*int32, bool)`

GetBcryptCostFactorOk returns a tuple with the BcryptCostFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcryptCostFactor

`func (o *AddPasswordStorageSchemeRequest) SetBcryptCostFactor(v int32)`

SetBcryptCostFactor sets BcryptCostFactor field to given value.

### HasBcryptCostFactor

`func (o *AddPasswordStorageSchemeRequest) HasBcryptCostFactor() bool`

HasBcryptCostFactor returns a boolean if a field has been set.

### GetAwsExternalServer

`func (o *AddPasswordStorageSchemeRequest) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *AddPasswordStorageSchemeRequest) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *AddPasswordStorageSchemeRequest) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.


### GetKeyVaultURI

`func (o *AddPasswordStorageSchemeRequest) GetKeyVaultURI() string`

GetKeyVaultURI returns the KeyVaultURI field if non-nil, zero value otherwise.

### GetKeyVaultURIOk

`func (o *AddPasswordStorageSchemeRequest) GetKeyVaultURIOk() (*string, bool)`

GetKeyVaultURIOk returns a tuple with the KeyVaultURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVaultURI

`func (o *AddPasswordStorageSchemeRequest) SetKeyVaultURI(v string)`

SetKeyVaultURI sets KeyVaultURI field to given value.


### GetAzureAuthenticationMethod

`func (o *AddPasswordStorageSchemeRequest) GetAzureAuthenticationMethod() string`

GetAzureAuthenticationMethod returns the AzureAuthenticationMethod field if non-nil, zero value otherwise.

### GetAzureAuthenticationMethodOk

`func (o *AddPasswordStorageSchemeRequest) GetAzureAuthenticationMethodOk() (*string, bool)`

GetAzureAuthenticationMethodOk returns a tuple with the AzureAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAuthenticationMethod

`func (o *AddPasswordStorageSchemeRequest) SetAzureAuthenticationMethod(v string)`

SetAzureAuthenticationMethod sets AzureAuthenticationMethod field to given value.


### GetHttpProxyExternalServer

`func (o *AddPasswordStorageSchemeRequest) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *AddPasswordStorageSchemeRequest) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *AddPasswordStorageSchemeRequest) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *AddPasswordStorageSchemeRequest) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetConjurExternalServer

`func (o *AddPasswordStorageSchemeRequest) GetConjurExternalServer() string`

GetConjurExternalServer returns the ConjurExternalServer field if non-nil, zero value otherwise.

### GetConjurExternalServerOk

`func (o *AddPasswordStorageSchemeRequest) GetConjurExternalServerOk() (*string, bool)`

GetConjurExternalServerOk returns a tuple with the ConjurExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurExternalServer

`func (o *AddPasswordStorageSchemeRequest) SetConjurExternalServer(v string)`

SetConjurExternalServer sets ConjurExternalServer field to given value.


### GetScryptCpuMemoryCostFactorExponent

`func (o *AddPasswordStorageSchemeRequest) GetScryptCpuMemoryCostFactorExponent() int32`

GetScryptCpuMemoryCostFactorExponent returns the ScryptCpuMemoryCostFactorExponent field if non-nil, zero value otherwise.

### GetScryptCpuMemoryCostFactorExponentOk

`func (o *AddPasswordStorageSchemeRequest) GetScryptCpuMemoryCostFactorExponentOk() (*int32, bool)`

GetScryptCpuMemoryCostFactorExponentOk returns a tuple with the ScryptCpuMemoryCostFactorExponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScryptCpuMemoryCostFactorExponent

`func (o *AddPasswordStorageSchemeRequest) SetScryptCpuMemoryCostFactorExponent(v int32)`

SetScryptCpuMemoryCostFactorExponent sets ScryptCpuMemoryCostFactorExponent field to given value.

### HasScryptCpuMemoryCostFactorExponent

`func (o *AddPasswordStorageSchemeRequest) HasScryptCpuMemoryCostFactorExponent() bool`

HasScryptCpuMemoryCostFactorExponent returns a boolean if a field has been set.

### GetScryptBlockSize

`func (o *AddPasswordStorageSchemeRequest) GetScryptBlockSize() int32`

GetScryptBlockSize returns the ScryptBlockSize field if non-nil, zero value otherwise.

### GetScryptBlockSizeOk

`func (o *AddPasswordStorageSchemeRequest) GetScryptBlockSizeOk() (*int32, bool)`

GetScryptBlockSizeOk returns a tuple with the ScryptBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScryptBlockSize

`func (o *AddPasswordStorageSchemeRequest) SetScryptBlockSize(v int32)`

SetScryptBlockSize sets ScryptBlockSize field to given value.

### HasScryptBlockSize

`func (o *AddPasswordStorageSchemeRequest) HasScryptBlockSize() bool`

HasScryptBlockSize returns a boolean if a field has been set.

### GetScryptParallelizationParameter

`func (o *AddPasswordStorageSchemeRequest) GetScryptParallelizationParameter() int32`

GetScryptParallelizationParameter returns the ScryptParallelizationParameter field if non-nil, zero value otherwise.

### GetScryptParallelizationParameterOk

`func (o *AddPasswordStorageSchemeRequest) GetScryptParallelizationParameterOk() (*int32, bool)`

GetScryptParallelizationParameterOk returns a tuple with the ScryptParallelizationParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScryptParallelizationParameter

`func (o *AddPasswordStorageSchemeRequest) SetScryptParallelizationParameter(v int32)`

SetScryptParallelizationParameter sets ScryptParallelizationParameter field to given value.

### HasScryptParallelizationParameter

`func (o *AddPasswordStorageSchemeRequest) HasScryptParallelizationParameter() bool`

HasScryptParallelizationParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


