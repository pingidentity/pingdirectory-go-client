# AddPassphraseProvider200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Passphrase Provider | 
**Schemas** | [**[]EnumthirdPartyPassphraseProviderSchemaUrn**](EnumthirdPartyPassphraseProviderSchemaUrn.md) |  | 
**EnvironmentVariable** | **string** | The name of the environment variable that is expected to hold the passphrase. | 
**Description** | Pointer to **string** | A description for this Passphrase Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Passphrase Provider is enabled for use in the server. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**AwsExternalServer** | **string** | The external server with information to use when interacting with the AWS Secrets Manager. | 
**SecretID** | **string** | The Amazon Resource Name (ARN) or the user-friendly name of the secret to be retrieved. | 
**SecretFieldName** | **string** | The name of the JSON field whose value is the passphrase that will be retrieved. | 
**SecretVersionID** | Pointer to **string** | The unique identifier for the version of the secret to be retrieved. | [optional] 
**SecretVersionStage** | Pointer to **string** | The staging label for the version of the secret to be retrieved. | [optional] 
**MaxCacheDuration** | Pointer to **string** | The maximum length of time that the passphrase provider may cache the passphrase that has been read from Vault. A value of zero seconds indicates that the provider should always attempt to read the passphrase from Vault. | [optional] 
**ObscuredValue** | **string** | The value to be stored in an obscured form. | 
**KeyVaultURI** | **string** | The URI that identifies the Azure Key Vault from which the secret is to be retrieved. | 
**AzureAuthenticationMethod** | **string** | The mechanism used to authenticate to the Azure service. | 
**SecretName** | **string** | The name of the secret to retrieve. | 
**PasswordFile** | **string** | The path to the file containing the passphrase. | 
**ConjurExternalServer** | **string** | An external server definition with information needed to connect and authenticate to the Conjur instance containing the passphrase. | 
**ConjurSecretRelativePath** | **string** | The portion of the path that follows the account name in the URI needed to obtain the desired secret. Any special characters in the path must be URL-encoded. | 
**VaultExternalServer** | **string** | An external server definition with information needed to connect and authenticate to the Vault instance containing the passphrase. | 
**VaultSecretPath** | **string** | The path to the desired secret in the Vault service. This will be appended to the value of the base-url property for the associated Vault external server. | 
**VaultSecretFieldName** | **string** | The name of the field in the Vault secret record that contains the passphrase to use to generate the encryption key. | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Passphrase Provider. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Passphrase Provider. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddPassphraseProvider200Response

`func NewAddPassphraseProvider200Response(id string, schemas []EnumthirdPartyPassphraseProviderSchemaUrn, environmentVariable string, enabled bool, awsExternalServer string, secretID string, secretFieldName string, obscuredValue string, keyVaultURI string, azureAuthenticationMethod string, secretName string, passwordFile string, conjurExternalServer string, conjurSecretRelativePath string, vaultExternalServer string, vaultSecretPath string, vaultSecretFieldName string, extensionClass string, ) *AddPassphraseProvider200Response`

NewAddPassphraseProvider200Response instantiates a new AddPassphraseProvider200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPassphraseProvider200ResponseWithDefaults

`func NewAddPassphraseProvider200ResponseWithDefaults() *AddPassphraseProvider200Response`

NewAddPassphraseProvider200ResponseWithDefaults instantiates a new AddPassphraseProvider200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddPassphraseProvider200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddPassphraseProvider200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddPassphraseProvider200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddPassphraseProvider200Response) GetSchemas() []EnumthirdPartyPassphraseProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPassphraseProvider200Response) GetSchemasOk() (*[]EnumthirdPartyPassphraseProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPassphraseProvider200Response) SetSchemas(v []EnumthirdPartyPassphraseProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnvironmentVariable

`func (o *AddPassphraseProvider200Response) GetEnvironmentVariable() string`

GetEnvironmentVariable returns the EnvironmentVariable field if non-nil, zero value otherwise.

### GetEnvironmentVariableOk

`func (o *AddPassphraseProvider200Response) GetEnvironmentVariableOk() (*string, bool)`

GetEnvironmentVariableOk returns a tuple with the EnvironmentVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariable

`func (o *AddPassphraseProvider200Response) SetEnvironmentVariable(v string)`

SetEnvironmentVariable sets EnvironmentVariable field to given value.


### GetDescription

`func (o *AddPassphraseProvider200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPassphraseProvider200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPassphraseProvider200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPassphraseProvider200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPassphraseProvider200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPassphraseProvider200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPassphraseProvider200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *AddPassphraseProvider200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddPassphraseProvider200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddPassphraseProvider200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddPassphraseProvider200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetAwsExternalServer

`func (o *AddPassphraseProvider200Response) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *AddPassphraseProvider200Response) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *AddPassphraseProvider200Response) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.


### GetSecretID

`func (o *AddPassphraseProvider200Response) GetSecretID() string`

GetSecretID returns the SecretID field if non-nil, zero value otherwise.

### GetSecretIDOk

`func (o *AddPassphraseProvider200Response) GetSecretIDOk() (*string, bool)`

GetSecretIDOk returns a tuple with the SecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretID

`func (o *AddPassphraseProvider200Response) SetSecretID(v string)`

SetSecretID sets SecretID field to given value.


### GetSecretFieldName

`func (o *AddPassphraseProvider200Response) GetSecretFieldName() string`

GetSecretFieldName returns the SecretFieldName field if non-nil, zero value otherwise.

### GetSecretFieldNameOk

`func (o *AddPassphraseProvider200Response) GetSecretFieldNameOk() (*string, bool)`

GetSecretFieldNameOk returns a tuple with the SecretFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretFieldName

`func (o *AddPassphraseProvider200Response) SetSecretFieldName(v string)`

SetSecretFieldName sets SecretFieldName field to given value.


### GetSecretVersionID

`func (o *AddPassphraseProvider200Response) GetSecretVersionID() string`

GetSecretVersionID returns the SecretVersionID field if non-nil, zero value otherwise.

### GetSecretVersionIDOk

`func (o *AddPassphraseProvider200Response) GetSecretVersionIDOk() (*string, bool)`

GetSecretVersionIDOk returns a tuple with the SecretVersionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVersionID

`func (o *AddPassphraseProvider200Response) SetSecretVersionID(v string)`

SetSecretVersionID sets SecretVersionID field to given value.

### HasSecretVersionID

`func (o *AddPassphraseProvider200Response) HasSecretVersionID() bool`

HasSecretVersionID returns a boolean if a field has been set.

### GetSecretVersionStage

`func (o *AddPassphraseProvider200Response) GetSecretVersionStage() string`

GetSecretVersionStage returns the SecretVersionStage field if non-nil, zero value otherwise.

### GetSecretVersionStageOk

`func (o *AddPassphraseProvider200Response) GetSecretVersionStageOk() (*string, bool)`

GetSecretVersionStageOk returns a tuple with the SecretVersionStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVersionStage

`func (o *AddPassphraseProvider200Response) SetSecretVersionStage(v string)`

SetSecretVersionStage sets SecretVersionStage field to given value.

### HasSecretVersionStage

`func (o *AddPassphraseProvider200Response) HasSecretVersionStage() bool`

HasSecretVersionStage returns a boolean if a field has been set.

### GetMaxCacheDuration

`func (o *AddPassphraseProvider200Response) GetMaxCacheDuration() string`

GetMaxCacheDuration returns the MaxCacheDuration field if non-nil, zero value otherwise.

### GetMaxCacheDurationOk

`func (o *AddPassphraseProvider200Response) GetMaxCacheDurationOk() (*string, bool)`

GetMaxCacheDurationOk returns a tuple with the MaxCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheDuration

`func (o *AddPassphraseProvider200Response) SetMaxCacheDuration(v string)`

SetMaxCacheDuration sets MaxCacheDuration field to given value.

### HasMaxCacheDuration

`func (o *AddPassphraseProvider200Response) HasMaxCacheDuration() bool`

HasMaxCacheDuration returns a boolean if a field has been set.

### GetObscuredValue

`func (o *AddPassphraseProvider200Response) GetObscuredValue() string`

GetObscuredValue returns the ObscuredValue field if non-nil, zero value otherwise.

### GetObscuredValueOk

`func (o *AddPassphraseProvider200Response) GetObscuredValueOk() (*string, bool)`

GetObscuredValueOk returns a tuple with the ObscuredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObscuredValue

`func (o *AddPassphraseProvider200Response) SetObscuredValue(v string)`

SetObscuredValue sets ObscuredValue field to given value.


### GetKeyVaultURI

`func (o *AddPassphraseProvider200Response) GetKeyVaultURI() string`

GetKeyVaultURI returns the KeyVaultURI field if non-nil, zero value otherwise.

### GetKeyVaultURIOk

`func (o *AddPassphraseProvider200Response) GetKeyVaultURIOk() (*string, bool)`

GetKeyVaultURIOk returns a tuple with the KeyVaultURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVaultURI

`func (o *AddPassphraseProvider200Response) SetKeyVaultURI(v string)`

SetKeyVaultURI sets KeyVaultURI field to given value.


### GetAzureAuthenticationMethod

`func (o *AddPassphraseProvider200Response) GetAzureAuthenticationMethod() string`

GetAzureAuthenticationMethod returns the AzureAuthenticationMethod field if non-nil, zero value otherwise.

### GetAzureAuthenticationMethodOk

`func (o *AddPassphraseProvider200Response) GetAzureAuthenticationMethodOk() (*string, bool)`

GetAzureAuthenticationMethodOk returns a tuple with the AzureAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAuthenticationMethod

`func (o *AddPassphraseProvider200Response) SetAzureAuthenticationMethod(v string)`

SetAzureAuthenticationMethod sets AzureAuthenticationMethod field to given value.


### GetSecretName

`func (o *AddPassphraseProvider200Response) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *AddPassphraseProvider200Response) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *AddPassphraseProvider200Response) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetPasswordFile

`func (o *AddPassphraseProvider200Response) GetPasswordFile() string`

GetPasswordFile returns the PasswordFile field if non-nil, zero value otherwise.

### GetPasswordFileOk

`func (o *AddPassphraseProvider200Response) GetPasswordFileOk() (*string, bool)`

GetPasswordFileOk returns a tuple with the PasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFile

`func (o *AddPassphraseProvider200Response) SetPasswordFile(v string)`

SetPasswordFile sets PasswordFile field to given value.


### GetConjurExternalServer

`func (o *AddPassphraseProvider200Response) GetConjurExternalServer() string`

GetConjurExternalServer returns the ConjurExternalServer field if non-nil, zero value otherwise.

### GetConjurExternalServerOk

`func (o *AddPassphraseProvider200Response) GetConjurExternalServerOk() (*string, bool)`

GetConjurExternalServerOk returns a tuple with the ConjurExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurExternalServer

`func (o *AddPassphraseProvider200Response) SetConjurExternalServer(v string)`

SetConjurExternalServer sets ConjurExternalServer field to given value.


### GetConjurSecretRelativePath

`func (o *AddPassphraseProvider200Response) GetConjurSecretRelativePath() string`

GetConjurSecretRelativePath returns the ConjurSecretRelativePath field if non-nil, zero value otherwise.

### GetConjurSecretRelativePathOk

`func (o *AddPassphraseProvider200Response) GetConjurSecretRelativePathOk() (*string, bool)`

GetConjurSecretRelativePathOk returns a tuple with the ConjurSecretRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurSecretRelativePath

`func (o *AddPassphraseProvider200Response) SetConjurSecretRelativePath(v string)`

SetConjurSecretRelativePath sets ConjurSecretRelativePath field to given value.


### GetVaultExternalServer

`func (o *AddPassphraseProvider200Response) GetVaultExternalServer() string`

GetVaultExternalServer returns the VaultExternalServer field if non-nil, zero value otherwise.

### GetVaultExternalServerOk

`func (o *AddPassphraseProvider200Response) GetVaultExternalServerOk() (*string, bool)`

GetVaultExternalServerOk returns a tuple with the VaultExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultExternalServer

`func (o *AddPassphraseProvider200Response) SetVaultExternalServer(v string)`

SetVaultExternalServer sets VaultExternalServer field to given value.


### GetVaultSecretPath

`func (o *AddPassphraseProvider200Response) GetVaultSecretPath() string`

GetVaultSecretPath returns the VaultSecretPath field if non-nil, zero value otherwise.

### GetVaultSecretPathOk

`func (o *AddPassphraseProvider200Response) GetVaultSecretPathOk() (*string, bool)`

GetVaultSecretPathOk returns a tuple with the VaultSecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretPath

`func (o *AddPassphraseProvider200Response) SetVaultSecretPath(v string)`

SetVaultSecretPath sets VaultSecretPath field to given value.


### GetVaultSecretFieldName

`func (o *AddPassphraseProvider200Response) GetVaultSecretFieldName() string`

GetVaultSecretFieldName returns the VaultSecretFieldName field if non-nil, zero value otherwise.

### GetVaultSecretFieldNameOk

`func (o *AddPassphraseProvider200Response) GetVaultSecretFieldNameOk() (*string, bool)`

GetVaultSecretFieldNameOk returns a tuple with the VaultSecretFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretFieldName

`func (o *AddPassphraseProvider200Response) SetVaultSecretFieldName(v string)`

SetVaultSecretFieldName sets VaultSecretFieldName field to given value.


### GetExtensionClass

`func (o *AddPassphraseProvider200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddPassphraseProvider200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddPassphraseProvider200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddPassphraseProvider200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddPassphraseProvider200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddPassphraseProvider200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddPassphraseProvider200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


