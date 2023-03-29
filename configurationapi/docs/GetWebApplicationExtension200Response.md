# GetWebApplicationExtension200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumgenericWebApplicationExtensionSchemaUrn**](EnumgenericWebApplicationExtensionSchemaUrn.md) |  | 
**Id** | **string** | Name of the Web Application Extension | 
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
**InitParameter** | Pointer to **[]string** | Specifies an initialization parameter to pass into the web application during startup. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewGetWebApplicationExtension200Response

`func NewGetWebApplicationExtension200Response(schemas []EnumgenericWebApplicationExtensionSchemaUrn, id string, baseContextPath string, ) *GetWebApplicationExtension200Response`

NewGetWebApplicationExtension200Response instantiates a new GetWebApplicationExtension200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWebApplicationExtension200ResponseWithDefaults

`func NewGetWebApplicationExtension200ResponseWithDefaults() *GetWebApplicationExtension200Response`

NewGetWebApplicationExtension200ResponseWithDefaults instantiates a new GetWebApplicationExtension200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GetWebApplicationExtension200Response) GetSchemas() []EnumgenericWebApplicationExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetWebApplicationExtension200Response) GetSchemasOk() (*[]EnumgenericWebApplicationExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetWebApplicationExtension200Response) SetSchemas(v []EnumgenericWebApplicationExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GetWebApplicationExtension200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetWebApplicationExtension200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetWebApplicationExtension200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSsoEnabled

`func (o *GetWebApplicationExtension200Response) GetSsoEnabled() bool`

GetSsoEnabled returns the SsoEnabled field if non-nil, zero value otherwise.

### GetSsoEnabledOk

`func (o *GetWebApplicationExtension200Response) GetSsoEnabledOk() (*bool, bool)`

GetSsoEnabledOk returns a tuple with the SsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEnabled

`func (o *GetWebApplicationExtension200Response) SetSsoEnabled(v bool)`

SetSsoEnabled sets SsoEnabled field to given value.

### HasSsoEnabled

`func (o *GetWebApplicationExtension200Response) HasSsoEnabled() bool`

HasSsoEnabled returns a boolean if a field has been set.

### GetOidcClientID

`func (o *GetWebApplicationExtension200Response) GetOidcClientID() string`

GetOidcClientID returns the OidcClientID field if non-nil, zero value otherwise.

### GetOidcClientIDOk

`func (o *GetWebApplicationExtension200Response) GetOidcClientIDOk() (*string, bool)`

GetOidcClientIDOk returns a tuple with the OidcClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientID

`func (o *GetWebApplicationExtension200Response) SetOidcClientID(v string)`

SetOidcClientID sets OidcClientID field to given value.

### HasOidcClientID

`func (o *GetWebApplicationExtension200Response) HasOidcClientID() bool`

HasOidcClientID returns a boolean if a field has been set.

### GetOidcClientSecret

`func (o *GetWebApplicationExtension200Response) GetOidcClientSecret() string`

GetOidcClientSecret returns the OidcClientSecret field if non-nil, zero value otherwise.

### GetOidcClientSecretOk

`func (o *GetWebApplicationExtension200Response) GetOidcClientSecretOk() (*string, bool)`

GetOidcClientSecretOk returns a tuple with the OidcClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientSecret

`func (o *GetWebApplicationExtension200Response) SetOidcClientSecret(v string)`

SetOidcClientSecret sets OidcClientSecret field to given value.

### HasOidcClientSecret

`func (o *GetWebApplicationExtension200Response) HasOidcClientSecret() bool`

HasOidcClientSecret returns a boolean if a field has been set.

### GetOidcClientSecretPassphraseProvider

`func (o *GetWebApplicationExtension200Response) GetOidcClientSecretPassphraseProvider() string`

GetOidcClientSecretPassphraseProvider returns the OidcClientSecretPassphraseProvider field if non-nil, zero value otherwise.

### GetOidcClientSecretPassphraseProviderOk

`func (o *GetWebApplicationExtension200Response) GetOidcClientSecretPassphraseProviderOk() (*string, bool)`

GetOidcClientSecretPassphraseProviderOk returns a tuple with the OidcClientSecretPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientSecretPassphraseProvider

`func (o *GetWebApplicationExtension200Response) SetOidcClientSecretPassphraseProvider(v string)`

SetOidcClientSecretPassphraseProvider sets OidcClientSecretPassphraseProvider field to given value.

### HasOidcClientSecretPassphraseProvider

`func (o *GetWebApplicationExtension200Response) HasOidcClientSecretPassphraseProvider() bool`

HasOidcClientSecretPassphraseProvider returns a boolean if a field has been set.

### GetOidcIssuerURL

`func (o *GetWebApplicationExtension200Response) GetOidcIssuerURL() string`

GetOidcIssuerURL returns the OidcIssuerURL field if non-nil, zero value otherwise.

### GetOidcIssuerURLOk

`func (o *GetWebApplicationExtension200Response) GetOidcIssuerURLOk() (*string, bool)`

GetOidcIssuerURLOk returns a tuple with the OidcIssuerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcIssuerURL

`func (o *GetWebApplicationExtension200Response) SetOidcIssuerURL(v string)`

SetOidcIssuerURL sets OidcIssuerURL field to given value.

### HasOidcIssuerURL

`func (o *GetWebApplicationExtension200Response) HasOidcIssuerURL() bool`

HasOidcIssuerURL returns a boolean if a field has been set.

### GetOidcTrustStoreFile

`func (o *GetWebApplicationExtension200Response) GetOidcTrustStoreFile() string`

GetOidcTrustStoreFile returns the OidcTrustStoreFile field if non-nil, zero value otherwise.

### GetOidcTrustStoreFileOk

`func (o *GetWebApplicationExtension200Response) GetOidcTrustStoreFileOk() (*string, bool)`

GetOidcTrustStoreFileOk returns a tuple with the OidcTrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcTrustStoreFile

`func (o *GetWebApplicationExtension200Response) SetOidcTrustStoreFile(v string)`

SetOidcTrustStoreFile sets OidcTrustStoreFile field to given value.

### HasOidcTrustStoreFile

`func (o *GetWebApplicationExtension200Response) HasOidcTrustStoreFile() bool`

HasOidcTrustStoreFile returns a boolean if a field has been set.

### GetOidcTrustStoreType

`func (o *GetWebApplicationExtension200Response) GetOidcTrustStoreType() string`

GetOidcTrustStoreType returns the OidcTrustStoreType field if non-nil, zero value otherwise.

### GetOidcTrustStoreTypeOk

`func (o *GetWebApplicationExtension200Response) GetOidcTrustStoreTypeOk() (*string, bool)`

GetOidcTrustStoreTypeOk returns a tuple with the OidcTrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcTrustStoreType

`func (o *GetWebApplicationExtension200Response) SetOidcTrustStoreType(v string)`

SetOidcTrustStoreType sets OidcTrustStoreType field to given value.

### HasOidcTrustStoreType

`func (o *GetWebApplicationExtension200Response) HasOidcTrustStoreType() bool`

HasOidcTrustStoreType returns a boolean if a field has been set.

### GetOidcTrustStorePinPassphraseProvider

`func (o *GetWebApplicationExtension200Response) GetOidcTrustStorePinPassphraseProvider() string`

GetOidcTrustStorePinPassphraseProvider returns the OidcTrustStorePinPassphraseProvider field if non-nil, zero value otherwise.

### GetOidcTrustStorePinPassphraseProviderOk

`func (o *GetWebApplicationExtension200Response) GetOidcTrustStorePinPassphraseProviderOk() (*string, bool)`

GetOidcTrustStorePinPassphraseProviderOk returns a tuple with the OidcTrustStorePinPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcTrustStorePinPassphraseProvider

`func (o *GetWebApplicationExtension200Response) SetOidcTrustStorePinPassphraseProvider(v string)`

SetOidcTrustStorePinPassphraseProvider sets OidcTrustStorePinPassphraseProvider field to given value.

### HasOidcTrustStorePinPassphraseProvider

`func (o *GetWebApplicationExtension200Response) HasOidcTrustStorePinPassphraseProvider() bool`

HasOidcTrustStorePinPassphraseProvider returns a boolean if a field has been set.

### GetOidcStrictHostnameVerification

`func (o *GetWebApplicationExtension200Response) GetOidcStrictHostnameVerification() bool`

GetOidcStrictHostnameVerification returns the OidcStrictHostnameVerification field if non-nil, zero value otherwise.

### GetOidcStrictHostnameVerificationOk

`func (o *GetWebApplicationExtension200Response) GetOidcStrictHostnameVerificationOk() (*bool, bool)`

GetOidcStrictHostnameVerificationOk returns a tuple with the OidcStrictHostnameVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcStrictHostnameVerification

`func (o *GetWebApplicationExtension200Response) SetOidcStrictHostnameVerification(v bool)`

SetOidcStrictHostnameVerification sets OidcStrictHostnameVerification field to given value.

### HasOidcStrictHostnameVerification

`func (o *GetWebApplicationExtension200Response) HasOidcStrictHostnameVerification() bool`

HasOidcStrictHostnameVerification returns a boolean if a field has been set.

### GetOidcTrustAll

`func (o *GetWebApplicationExtension200Response) GetOidcTrustAll() bool`

GetOidcTrustAll returns the OidcTrustAll field if non-nil, zero value otherwise.

### GetOidcTrustAllOk

`func (o *GetWebApplicationExtension200Response) GetOidcTrustAllOk() (*bool, bool)`

GetOidcTrustAllOk returns a tuple with the OidcTrustAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcTrustAll

`func (o *GetWebApplicationExtension200Response) SetOidcTrustAll(v bool)`

SetOidcTrustAll sets OidcTrustAll field to given value.

### HasOidcTrustAll

`func (o *GetWebApplicationExtension200Response) HasOidcTrustAll() bool`

HasOidcTrustAll returns a boolean if a field has been set.

### GetLdapServer

`func (o *GetWebApplicationExtension200Response) GetLdapServer() string`

GetLdapServer returns the LdapServer field if non-nil, zero value otherwise.

### GetLdapServerOk

`func (o *GetWebApplicationExtension200Response) GetLdapServerOk() (*string, bool)`

GetLdapServerOk returns a tuple with the LdapServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServer

`func (o *GetWebApplicationExtension200Response) SetLdapServer(v string)`

SetLdapServer sets LdapServer field to given value.

### HasLdapServer

`func (o *GetWebApplicationExtension200Response) HasLdapServer() bool`

HasLdapServer returns a boolean if a field has been set.

### GetTrustStoreFile

`func (o *GetWebApplicationExtension200Response) GetTrustStoreFile() string`

GetTrustStoreFile returns the TrustStoreFile field if non-nil, zero value otherwise.

### GetTrustStoreFileOk

`func (o *GetWebApplicationExtension200Response) GetTrustStoreFileOk() (*string, bool)`

GetTrustStoreFileOk returns a tuple with the TrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreFile

`func (o *GetWebApplicationExtension200Response) SetTrustStoreFile(v string)`

SetTrustStoreFile sets TrustStoreFile field to given value.

### HasTrustStoreFile

`func (o *GetWebApplicationExtension200Response) HasTrustStoreFile() bool`

HasTrustStoreFile returns a boolean if a field has been set.

### GetTrustStoreType

`func (o *GetWebApplicationExtension200Response) GetTrustStoreType() string`

GetTrustStoreType returns the TrustStoreType field if non-nil, zero value otherwise.

### GetTrustStoreTypeOk

`func (o *GetWebApplicationExtension200Response) GetTrustStoreTypeOk() (*string, bool)`

GetTrustStoreTypeOk returns a tuple with the TrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreType

`func (o *GetWebApplicationExtension200Response) SetTrustStoreType(v string)`

SetTrustStoreType sets TrustStoreType field to given value.

### HasTrustStoreType

`func (o *GetWebApplicationExtension200Response) HasTrustStoreType() bool`

HasTrustStoreType returns a boolean if a field has been set.

### GetTrustStorePinPassphraseProvider

`func (o *GetWebApplicationExtension200Response) GetTrustStorePinPassphraseProvider() string`

GetTrustStorePinPassphraseProvider returns the TrustStorePinPassphraseProvider field if non-nil, zero value otherwise.

### GetTrustStorePinPassphraseProviderOk

`func (o *GetWebApplicationExtension200Response) GetTrustStorePinPassphraseProviderOk() (*string, bool)`

GetTrustStorePinPassphraseProviderOk returns a tuple with the TrustStorePinPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePinPassphraseProvider

`func (o *GetWebApplicationExtension200Response) SetTrustStorePinPassphraseProvider(v string)`

SetTrustStorePinPassphraseProvider sets TrustStorePinPassphraseProvider field to given value.

### HasTrustStorePinPassphraseProvider

`func (o *GetWebApplicationExtension200Response) HasTrustStorePinPassphraseProvider() bool`

HasTrustStorePinPassphraseProvider returns a boolean if a field has been set.

### GetLogFile

`func (o *GetWebApplicationExtension200Response) GetLogFile() string`

GetLogFile returns the LogFile field if non-nil, zero value otherwise.

### GetLogFileOk

`func (o *GetWebApplicationExtension200Response) GetLogFileOk() (*string, bool)`

GetLogFileOk returns a tuple with the LogFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFile

`func (o *GetWebApplicationExtension200Response) SetLogFile(v string)`

SetLogFile sets LogFile field to given value.

### HasLogFile

`func (o *GetWebApplicationExtension200Response) HasLogFile() bool`

HasLogFile returns a boolean if a field has been set.

### GetComplexity

`func (o *GetWebApplicationExtension200Response) GetComplexity() EnumwebApplicationExtensionComplexityProp`

GetComplexity returns the Complexity field if non-nil, zero value otherwise.

### GetComplexityOk

`func (o *GetWebApplicationExtension200Response) GetComplexityOk() (*EnumwebApplicationExtensionComplexityProp, bool)`

GetComplexityOk returns a tuple with the Complexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexity

`func (o *GetWebApplicationExtension200Response) SetComplexity(v EnumwebApplicationExtensionComplexityProp)`

SetComplexity sets Complexity field to given value.

### HasComplexity

`func (o *GetWebApplicationExtension200Response) HasComplexity() bool`

HasComplexity returns a boolean if a field has been set.

### GetDescription

`func (o *GetWebApplicationExtension200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetWebApplicationExtension200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetWebApplicationExtension200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetWebApplicationExtension200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBaseContextPath

`func (o *GetWebApplicationExtension200Response) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *GetWebApplicationExtension200Response) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *GetWebApplicationExtension200Response) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.


### GetWarFile

`func (o *GetWebApplicationExtension200Response) GetWarFile() string`

GetWarFile returns the WarFile field if non-nil, zero value otherwise.

### GetWarFileOk

`func (o *GetWebApplicationExtension200Response) GetWarFileOk() (*string, bool)`

GetWarFileOk returns a tuple with the WarFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarFile

`func (o *GetWebApplicationExtension200Response) SetWarFile(v string)`

SetWarFile sets WarFile field to given value.

### HasWarFile

`func (o *GetWebApplicationExtension200Response) HasWarFile() bool`

HasWarFile returns a boolean if a field has been set.

### GetDocumentRootDirectory

`func (o *GetWebApplicationExtension200Response) GetDocumentRootDirectory() string`

GetDocumentRootDirectory returns the DocumentRootDirectory field if non-nil, zero value otherwise.

### GetDocumentRootDirectoryOk

`func (o *GetWebApplicationExtension200Response) GetDocumentRootDirectoryOk() (*string, bool)`

GetDocumentRootDirectoryOk returns a tuple with the DocumentRootDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentRootDirectory

`func (o *GetWebApplicationExtension200Response) SetDocumentRootDirectory(v string)`

SetDocumentRootDirectory sets DocumentRootDirectory field to given value.

### HasDocumentRootDirectory

`func (o *GetWebApplicationExtension200Response) HasDocumentRootDirectory() bool`

HasDocumentRootDirectory returns a boolean if a field has been set.

### GetDeploymentDescriptorFile

`func (o *GetWebApplicationExtension200Response) GetDeploymentDescriptorFile() string`

GetDeploymentDescriptorFile returns the DeploymentDescriptorFile field if non-nil, zero value otherwise.

### GetDeploymentDescriptorFileOk

`func (o *GetWebApplicationExtension200Response) GetDeploymentDescriptorFileOk() (*string, bool)`

GetDeploymentDescriptorFileOk returns a tuple with the DeploymentDescriptorFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentDescriptorFile

`func (o *GetWebApplicationExtension200Response) SetDeploymentDescriptorFile(v string)`

SetDeploymentDescriptorFile sets DeploymentDescriptorFile field to given value.

### HasDeploymentDescriptorFile

`func (o *GetWebApplicationExtension200Response) HasDeploymentDescriptorFile() bool`

HasDeploymentDescriptorFile returns a boolean if a field has been set.

### GetTemporaryDirectory

`func (o *GetWebApplicationExtension200Response) GetTemporaryDirectory() string`

GetTemporaryDirectory returns the TemporaryDirectory field if non-nil, zero value otherwise.

### GetTemporaryDirectoryOk

`func (o *GetWebApplicationExtension200Response) GetTemporaryDirectoryOk() (*string, bool)`

GetTemporaryDirectoryOk returns a tuple with the TemporaryDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDirectory

`func (o *GetWebApplicationExtension200Response) SetTemporaryDirectory(v string)`

SetTemporaryDirectory sets TemporaryDirectory field to given value.

### HasTemporaryDirectory

`func (o *GetWebApplicationExtension200Response) HasTemporaryDirectory() bool`

HasTemporaryDirectory returns a boolean if a field has been set.

### GetInitParameter

`func (o *GetWebApplicationExtension200Response) GetInitParameter() []string`

GetInitParameter returns the InitParameter field if non-nil, zero value otherwise.

### GetInitParameterOk

`func (o *GetWebApplicationExtension200Response) GetInitParameterOk() (*[]string, bool)`

GetInitParameterOk returns a tuple with the InitParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitParameter

`func (o *GetWebApplicationExtension200Response) SetInitParameter(v []string)`

SetInitParameter sets InitParameter field to given value.

### HasInitParameter

`func (o *GetWebApplicationExtension200Response) HasInitParameter() bool`

HasInitParameter returns a boolean if a field has been set.

### GetMeta

`func (o *GetWebApplicationExtension200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetWebApplicationExtension200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetWebApplicationExtension200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetWebApplicationExtension200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GetWebApplicationExtension200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GetWebApplicationExtension200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GetWebApplicationExtension200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GetWebApplicationExtension200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


