# ConsoleWebApplicationExtensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumconsoleWebApplicationExtensionSchemaUrn**](EnumconsoleWebApplicationExtensionSchemaUrn.md) |  | [optional] 
**SsoEnabled** | Pointer to **bool** | Indicates that SSO login into the Administrative Console is enabled. | [optional] 
**OidcClientID** | Pointer to **string** | The client ID to use when authenticating to the OpenID Connect provider. | [optional] 
**OidcClientSecret** | Pointer to **string** | The client secret to use when authenticating to the OpenID Connect provider. | [optional] 
**OidcClientSecretPassphraseProvider** | Pointer to **string** | A passphrase provider that may be used to obtain the client secret to use when authenticating to the OpenID Connect provider. | [optional] 
**OidcIssuerURL** | Pointer to **string** | The issuer URL of the OpenID Connect provider. | [optional] 
**OidcTrustStoreFile** | Pointer to **string** | Specifies the path to the truststore file used by this application to evaluate OIDC provider certificates. If this field is left blank, the default JVM trust store will be used. | [optional] 
**OidcTrustStoreType** | Pointer to **string** | Specifies the format for the data in the OIDC trust store file. | [optional] 
**OidcTrustStorePinPassphraseProvider** | Pointer to **string** | The passphrase provider that may be used to obtain the PIN for the trust store used with OIDC providers. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents. | [optional] 
**OidcStrictHostnameVerification** | Pointer to **bool** | Controls whether or not hostname verification is performed, which checks if the hostname of the OIDC provider matches the name(s) stored inside the certificate it provides. This property should only be set to false for testing purposes. | [optional] 
**OidcTrustAll** | Pointer to **bool** | Controls whether or not this application will always trust any certificate that is presented to it, regardless of its contents. This property should only be set to true for testing purposes. | [optional] 
**LdapServer** | Pointer to **string** | The LDAP URL used to connect to the managed server. | [optional] 
**TrustStoreFile** | Pointer to **string** | Specifies the path to the truststore file, which is used by this application to establish trust of managed servers. | [optional] 
**TrustStoreType** | Pointer to **string** | Specifies the format for the data in the trust store file. | [optional] 
**TrustStorePinPassphraseProvider** | Pointer to **string** | The passphrase provider that may be used to obtain the PIN for the trust store used with managed LDAP servers. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents. | [optional] 
**LogFile** | Pointer to **string** | The path to the log file for the web application. | [optional] 
**Complexity** | Pointer to [**EnumwebApplicationExtensionComplexityProp**](EnumwebApplicationExtensionComplexityProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Web Application Extension | [optional] 
**BaseContextPath** | **string** | Specifies the base context path that should be used by HTTP clients to reference content. The value must start with a forward slash and at least one additional character and must represent a valid HTTP context path. | 
**WarFile** | Pointer to **string** | Specifies the path to a standard web application archive (WAR) file. | [optional] 
**DocumentRootDirectory** | Pointer to **string** | Specifies the path to the directory on the local filesystem containing the files to be served by this Web Application Extension. The path must exist, and it must be a directory. | [optional] 
**DeploymentDescriptorFile** | Pointer to **string** | Specifies the path to the deployment descriptor file when used with document-root-directory. | [optional] 
**TemporaryDirectory** | Pointer to **string** | Specifies the path to the directory that may be used to store temporary files such as extracted WAR files and compiled JSP files. | [optional] 
**InitParameter** | Pointer to **[]string** |  | [optional] 

## Methods

### NewConsoleWebApplicationExtensionResponse

`func NewConsoleWebApplicationExtensionResponse(baseContextPath string, ) *ConsoleWebApplicationExtensionResponse`

NewConsoleWebApplicationExtensionResponse instantiates a new ConsoleWebApplicationExtensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleWebApplicationExtensionResponseWithDefaults

`func NewConsoleWebApplicationExtensionResponseWithDefaults() *ConsoleWebApplicationExtensionResponse`

NewConsoleWebApplicationExtensionResponseWithDefaults instantiates a new ConsoleWebApplicationExtensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ConsoleWebApplicationExtensionResponse) GetSchemas() []EnumconsoleWebApplicationExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ConsoleWebApplicationExtensionResponse) GetSchemasOk() (*[]EnumconsoleWebApplicationExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ConsoleWebApplicationExtensionResponse) SetSchemas(v []EnumconsoleWebApplicationExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ConsoleWebApplicationExtensionResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetSsoEnabled

`func (o *ConsoleWebApplicationExtensionResponse) GetSsoEnabled() bool`

GetSsoEnabled returns the SsoEnabled field if non-nil, zero value otherwise.

### GetSsoEnabledOk

`func (o *ConsoleWebApplicationExtensionResponse) GetSsoEnabledOk() (*bool, bool)`

GetSsoEnabledOk returns a tuple with the SsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEnabled

`func (o *ConsoleWebApplicationExtensionResponse) SetSsoEnabled(v bool)`

SetSsoEnabled sets SsoEnabled field to given value.

### HasSsoEnabled

`func (o *ConsoleWebApplicationExtensionResponse) HasSsoEnabled() bool`

HasSsoEnabled returns a boolean if a field has been set.

### GetOidcClientID

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcClientID() string`

GetOidcClientID returns the OidcClientID field if non-nil, zero value otherwise.

### GetOidcClientIDOk

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcClientIDOk() (*string, bool)`

GetOidcClientIDOk returns a tuple with the OidcClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientID

`func (o *ConsoleWebApplicationExtensionResponse) SetOidcClientID(v string)`

SetOidcClientID sets OidcClientID field to given value.

### HasOidcClientID

`func (o *ConsoleWebApplicationExtensionResponse) HasOidcClientID() bool`

HasOidcClientID returns a boolean if a field has been set.

### GetOidcClientSecret

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcClientSecret() string`

GetOidcClientSecret returns the OidcClientSecret field if non-nil, zero value otherwise.

### GetOidcClientSecretOk

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcClientSecretOk() (*string, bool)`

GetOidcClientSecretOk returns a tuple with the OidcClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientSecret

`func (o *ConsoleWebApplicationExtensionResponse) SetOidcClientSecret(v string)`

SetOidcClientSecret sets OidcClientSecret field to given value.

### HasOidcClientSecret

`func (o *ConsoleWebApplicationExtensionResponse) HasOidcClientSecret() bool`

HasOidcClientSecret returns a boolean if a field has been set.

### GetOidcClientSecretPassphraseProvider

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcClientSecretPassphraseProvider() string`

GetOidcClientSecretPassphraseProvider returns the OidcClientSecretPassphraseProvider field if non-nil, zero value otherwise.

### GetOidcClientSecretPassphraseProviderOk

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcClientSecretPassphraseProviderOk() (*string, bool)`

GetOidcClientSecretPassphraseProviderOk returns a tuple with the OidcClientSecretPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientSecretPassphraseProvider

`func (o *ConsoleWebApplicationExtensionResponse) SetOidcClientSecretPassphraseProvider(v string)`

SetOidcClientSecretPassphraseProvider sets OidcClientSecretPassphraseProvider field to given value.

### HasOidcClientSecretPassphraseProvider

`func (o *ConsoleWebApplicationExtensionResponse) HasOidcClientSecretPassphraseProvider() bool`

HasOidcClientSecretPassphraseProvider returns a boolean if a field has been set.

### GetOidcIssuerURL

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcIssuerURL() string`

GetOidcIssuerURL returns the OidcIssuerURL field if non-nil, zero value otherwise.

### GetOidcIssuerURLOk

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcIssuerURLOk() (*string, bool)`

GetOidcIssuerURLOk returns a tuple with the OidcIssuerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcIssuerURL

`func (o *ConsoleWebApplicationExtensionResponse) SetOidcIssuerURL(v string)`

SetOidcIssuerURL sets OidcIssuerURL field to given value.

### HasOidcIssuerURL

`func (o *ConsoleWebApplicationExtensionResponse) HasOidcIssuerURL() bool`

HasOidcIssuerURL returns a boolean if a field has been set.

### GetOidcTrustStoreFile

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcTrustStoreFile() string`

GetOidcTrustStoreFile returns the OidcTrustStoreFile field if non-nil, zero value otherwise.

### GetOidcTrustStoreFileOk

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcTrustStoreFileOk() (*string, bool)`

GetOidcTrustStoreFileOk returns a tuple with the OidcTrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcTrustStoreFile

`func (o *ConsoleWebApplicationExtensionResponse) SetOidcTrustStoreFile(v string)`

SetOidcTrustStoreFile sets OidcTrustStoreFile field to given value.

### HasOidcTrustStoreFile

`func (o *ConsoleWebApplicationExtensionResponse) HasOidcTrustStoreFile() bool`

HasOidcTrustStoreFile returns a boolean if a field has been set.

### GetOidcTrustStoreType

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcTrustStoreType() string`

GetOidcTrustStoreType returns the OidcTrustStoreType field if non-nil, zero value otherwise.

### GetOidcTrustStoreTypeOk

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcTrustStoreTypeOk() (*string, bool)`

GetOidcTrustStoreTypeOk returns a tuple with the OidcTrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcTrustStoreType

`func (o *ConsoleWebApplicationExtensionResponse) SetOidcTrustStoreType(v string)`

SetOidcTrustStoreType sets OidcTrustStoreType field to given value.

### HasOidcTrustStoreType

`func (o *ConsoleWebApplicationExtensionResponse) HasOidcTrustStoreType() bool`

HasOidcTrustStoreType returns a boolean if a field has been set.

### GetOidcTrustStorePinPassphraseProvider

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcTrustStorePinPassphraseProvider() string`

GetOidcTrustStorePinPassphraseProvider returns the OidcTrustStorePinPassphraseProvider field if non-nil, zero value otherwise.

### GetOidcTrustStorePinPassphraseProviderOk

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcTrustStorePinPassphraseProviderOk() (*string, bool)`

GetOidcTrustStorePinPassphraseProviderOk returns a tuple with the OidcTrustStorePinPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcTrustStorePinPassphraseProvider

`func (o *ConsoleWebApplicationExtensionResponse) SetOidcTrustStorePinPassphraseProvider(v string)`

SetOidcTrustStorePinPassphraseProvider sets OidcTrustStorePinPassphraseProvider field to given value.

### HasOidcTrustStorePinPassphraseProvider

`func (o *ConsoleWebApplicationExtensionResponse) HasOidcTrustStorePinPassphraseProvider() bool`

HasOidcTrustStorePinPassphraseProvider returns a boolean if a field has been set.

### GetOidcStrictHostnameVerification

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcStrictHostnameVerification() bool`

GetOidcStrictHostnameVerification returns the OidcStrictHostnameVerification field if non-nil, zero value otherwise.

### GetOidcStrictHostnameVerificationOk

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcStrictHostnameVerificationOk() (*bool, bool)`

GetOidcStrictHostnameVerificationOk returns a tuple with the OidcStrictHostnameVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcStrictHostnameVerification

`func (o *ConsoleWebApplicationExtensionResponse) SetOidcStrictHostnameVerification(v bool)`

SetOidcStrictHostnameVerification sets OidcStrictHostnameVerification field to given value.

### HasOidcStrictHostnameVerification

`func (o *ConsoleWebApplicationExtensionResponse) HasOidcStrictHostnameVerification() bool`

HasOidcStrictHostnameVerification returns a boolean if a field has been set.

### GetOidcTrustAll

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcTrustAll() bool`

GetOidcTrustAll returns the OidcTrustAll field if non-nil, zero value otherwise.

### GetOidcTrustAllOk

`func (o *ConsoleWebApplicationExtensionResponse) GetOidcTrustAllOk() (*bool, bool)`

GetOidcTrustAllOk returns a tuple with the OidcTrustAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcTrustAll

`func (o *ConsoleWebApplicationExtensionResponse) SetOidcTrustAll(v bool)`

SetOidcTrustAll sets OidcTrustAll field to given value.

### HasOidcTrustAll

`func (o *ConsoleWebApplicationExtensionResponse) HasOidcTrustAll() bool`

HasOidcTrustAll returns a boolean if a field has been set.

### GetLdapServer

`func (o *ConsoleWebApplicationExtensionResponse) GetLdapServer() string`

GetLdapServer returns the LdapServer field if non-nil, zero value otherwise.

### GetLdapServerOk

`func (o *ConsoleWebApplicationExtensionResponse) GetLdapServerOk() (*string, bool)`

GetLdapServerOk returns a tuple with the LdapServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServer

`func (o *ConsoleWebApplicationExtensionResponse) SetLdapServer(v string)`

SetLdapServer sets LdapServer field to given value.

### HasLdapServer

`func (o *ConsoleWebApplicationExtensionResponse) HasLdapServer() bool`

HasLdapServer returns a boolean if a field has been set.

### GetTrustStoreFile

`func (o *ConsoleWebApplicationExtensionResponse) GetTrustStoreFile() string`

GetTrustStoreFile returns the TrustStoreFile field if non-nil, zero value otherwise.

### GetTrustStoreFileOk

`func (o *ConsoleWebApplicationExtensionResponse) GetTrustStoreFileOk() (*string, bool)`

GetTrustStoreFileOk returns a tuple with the TrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreFile

`func (o *ConsoleWebApplicationExtensionResponse) SetTrustStoreFile(v string)`

SetTrustStoreFile sets TrustStoreFile field to given value.

### HasTrustStoreFile

`func (o *ConsoleWebApplicationExtensionResponse) HasTrustStoreFile() bool`

HasTrustStoreFile returns a boolean if a field has been set.

### GetTrustStoreType

`func (o *ConsoleWebApplicationExtensionResponse) GetTrustStoreType() string`

GetTrustStoreType returns the TrustStoreType field if non-nil, zero value otherwise.

### GetTrustStoreTypeOk

`func (o *ConsoleWebApplicationExtensionResponse) GetTrustStoreTypeOk() (*string, bool)`

GetTrustStoreTypeOk returns a tuple with the TrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreType

`func (o *ConsoleWebApplicationExtensionResponse) SetTrustStoreType(v string)`

SetTrustStoreType sets TrustStoreType field to given value.

### HasTrustStoreType

`func (o *ConsoleWebApplicationExtensionResponse) HasTrustStoreType() bool`

HasTrustStoreType returns a boolean if a field has been set.

### GetTrustStorePinPassphraseProvider

`func (o *ConsoleWebApplicationExtensionResponse) GetTrustStorePinPassphraseProvider() string`

GetTrustStorePinPassphraseProvider returns the TrustStorePinPassphraseProvider field if non-nil, zero value otherwise.

### GetTrustStorePinPassphraseProviderOk

`func (o *ConsoleWebApplicationExtensionResponse) GetTrustStorePinPassphraseProviderOk() (*string, bool)`

GetTrustStorePinPassphraseProviderOk returns a tuple with the TrustStorePinPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePinPassphraseProvider

`func (o *ConsoleWebApplicationExtensionResponse) SetTrustStorePinPassphraseProvider(v string)`

SetTrustStorePinPassphraseProvider sets TrustStorePinPassphraseProvider field to given value.

### HasTrustStorePinPassphraseProvider

`func (o *ConsoleWebApplicationExtensionResponse) HasTrustStorePinPassphraseProvider() bool`

HasTrustStorePinPassphraseProvider returns a boolean if a field has been set.

### GetLogFile

`func (o *ConsoleWebApplicationExtensionResponse) GetLogFile() string`

GetLogFile returns the LogFile field if non-nil, zero value otherwise.

### GetLogFileOk

`func (o *ConsoleWebApplicationExtensionResponse) GetLogFileOk() (*string, bool)`

GetLogFileOk returns a tuple with the LogFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFile

`func (o *ConsoleWebApplicationExtensionResponse) SetLogFile(v string)`

SetLogFile sets LogFile field to given value.

### HasLogFile

`func (o *ConsoleWebApplicationExtensionResponse) HasLogFile() bool`

HasLogFile returns a boolean if a field has been set.

### GetComplexity

`func (o *ConsoleWebApplicationExtensionResponse) GetComplexity() EnumwebApplicationExtensionComplexityProp`

GetComplexity returns the Complexity field if non-nil, zero value otherwise.

### GetComplexityOk

`func (o *ConsoleWebApplicationExtensionResponse) GetComplexityOk() (*EnumwebApplicationExtensionComplexityProp, bool)`

GetComplexityOk returns a tuple with the Complexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexity

`func (o *ConsoleWebApplicationExtensionResponse) SetComplexity(v EnumwebApplicationExtensionComplexityProp)`

SetComplexity sets Complexity field to given value.

### HasComplexity

`func (o *ConsoleWebApplicationExtensionResponse) HasComplexity() bool`

HasComplexity returns a boolean if a field has been set.

### GetDescription

`func (o *ConsoleWebApplicationExtensionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsoleWebApplicationExtensionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsoleWebApplicationExtensionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsoleWebApplicationExtensionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBaseContextPath

`func (o *ConsoleWebApplicationExtensionResponse) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *ConsoleWebApplicationExtensionResponse) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *ConsoleWebApplicationExtensionResponse) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.


### GetWarFile

`func (o *ConsoleWebApplicationExtensionResponse) GetWarFile() string`

GetWarFile returns the WarFile field if non-nil, zero value otherwise.

### GetWarFileOk

`func (o *ConsoleWebApplicationExtensionResponse) GetWarFileOk() (*string, bool)`

GetWarFileOk returns a tuple with the WarFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarFile

`func (o *ConsoleWebApplicationExtensionResponse) SetWarFile(v string)`

SetWarFile sets WarFile field to given value.

### HasWarFile

`func (o *ConsoleWebApplicationExtensionResponse) HasWarFile() bool`

HasWarFile returns a boolean if a field has been set.

### GetDocumentRootDirectory

`func (o *ConsoleWebApplicationExtensionResponse) GetDocumentRootDirectory() string`

GetDocumentRootDirectory returns the DocumentRootDirectory field if non-nil, zero value otherwise.

### GetDocumentRootDirectoryOk

`func (o *ConsoleWebApplicationExtensionResponse) GetDocumentRootDirectoryOk() (*string, bool)`

GetDocumentRootDirectoryOk returns a tuple with the DocumentRootDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentRootDirectory

`func (o *ConsoleWebApplicationExtensionResponse) SetDocumentRootDirectory(v string)`

SetDocumentRootDirectory sets DocumentRootDirectory field to given value.

### HasDocumentRootDirectory

`func (o *ConsoleWebApplicationExtensionResponse) HasDocumentRootDirectory() bool`

HasDocumentRootDirectory returns a boolean if a field has been set.

### GetDeploymentDescriptorFile

`func (o *ConsoleWebApplicationExtensionResponse) GetDeploymentDescriptorFile() string`

GetDeploymentDescriptorFile returns the DeploymentDescriptorFile field if non-nil, zero value otherwise.

### GetDeploymentDescriptorFileOk

`func (o *ConsoleWebApplicationExtensionResponse) GetDeploymentDescriptorFileOk() (*string, bool)`

GetDeploymentDescriptorFileOk returns a tuple with the DeploymentDescriptorFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentDescriptorFile

`func (o *ConsoleWebApplicationExtensionResponse) SetDeploymentDescriptorFile(v string)`

SetDeploymentDescriptorFile sets DeploymentDescriptorFile field to given value.

### HasDeploymentDescriptorFile

`func (o *ConsoleWebApplicationExtensionResponse) HasDeploymentDescriptorFile() bool`

HasDeploymentDescriptorFile returns a boolean if a field has been set.

### GetTemporaryDirectory

`func (o *ConsoleWebApplicationExtensionResponse) GetTemporaryDirectory() string`

GetTemporaryDirectory returns the TemporaryDirectory field if non-nil, zero value otherwise.

### GetTemporaryDirectoryOk

`func (o *ConsoleWebApplicationExtensionResponse) GetTemporaryDirectoryOk() (*string, bool)`

GetTemporaryDirectoryOk returns a tuple with the TemporaryDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDirectory

`func (o *ConsoleWebApplicationExtensionResponse) SetTemporaryDirectory(v string)`

SetTemporaryDirectory sets TemporaryDirectory field to given value.

### HasTemporaryDirectory

`func (o *ConsoleWebApplicationExtensionResponse) HasTemporaryDirectory() bool`

HasTemporaryDirectory returns a boolean if a field has been set.

### GetInitParameter

`func (o *ConsoleWebApplicationExtensionResponse) GetInitParameter() []string`

GetInitParameter returns the InitParameter field if non-nil, zero value otherwise.

### GetInitParameterOk

`func (o *ConsoleWebApplicationExtensionResponse) GetInitParameterOk() (*[]string, bool)`

GetInitParameterOk returns a tuple with the InitParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitParameter

`func (o *ConsoleWebApplicationExtensionResponse) SetInitParameter(v []string)`

SetInitParameter sets InitParameter field to given value.

### HasInitParameter

`func (o *ConsoleWebApplicationExtensionResponse) HasInitParameter() bool`

HasInitParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


