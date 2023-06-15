# PingOnePassThroughAuthenticationPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Plugin | 
**Schemas** | [**[]EnumpingOnePassThroughAuthenticationPluginSchemaUrn**](EnumpingOnePassThroughAuthenticationPluginSchemaUrn.md) |  | 
**ApiURL** | **string** | Specifies the API endpoint for the PingOne web service. | 
**AuthURL** | **string** | Specifies the API endpoint for the PingOne authentication service. | 
**OAuthClientID** | **string** | Specifies the OAuth Client ID used to authenticate connections to the PingOne API. | 
**OAuthClientSecret** | Pointer to **string** | Specifies the OAuth Client Secret used to authenticate connections to the PingOne API. | [optional] 
**OAuthClientSecretPassphraseProvider** | Pointer to **string** | Specifies a passphrase provider that can be used to obtain the OAuth Client Secret used to authenticate connections to the PingOne API. | [optional] 
**EnvironmentID** | **string** | Specifies the PingOne Environment that will be associated with this PingOne Pass Through Authentication Plugin. | 
**HttpProxyExternalServer** | Pointer to **string** | A reference to an HTTP proxy server that should be used for requests sent to the PingOne service. | [optional] 
**IncludedLocalEntryBaseDN** | Pointer to **[]string** | The base DNs for the local users whose authentication attempts may be passed through to the PingOne service. | [optional] 
**ConnectionCriteria** | Pointer to **string** | A reference to connection criteria that will be used to indicate which bind requests should be passed through to the PingOne service. | [optional] 
**RequestCriteria** | Pointer to **string** | A reference to request criteria that will be used to indicate which bind requests should be passed through to the PingOne service. | [optional] 
**TryLocalBind** | Pointer to **bool** | Indicates whether to attempt the bind in the local server first, or to only send it to the PingOne service. | [optional] 
**OverrideLocalPassword** | Pointer to **bool** | Indicates whether to attempt the authentication in the PingOne service if the local user entry includes a password. This property will only be used if try-local-bind is true. | [optional] 
**UpdateLocalPassword** | Pointer to **bool** | Indicates whether to overwrite the user&#39;s local password if the local bind fails but the authentication attempt succeeds when attempted in the PingOne service. | [optional] 
**UpdateLocalPasswordDN** | Pointer to **string** | This is the DN of the user that will be used to overwrite the user&#39;s local password if update-local-password is set. The DN put here should be added to &#39;ignore-changes-by-dn&#39; in the appropriate Sync Source. | [optional] 
**AllowLaxPassThroughAuthenticationPasswords** | Pointer to **bool** | Indicates whether to overwrite the user&#39;s local password even if the password used to authenticate to the PingOne service would have failed validation if the user attempted to set it directly. | [optional] 
**IgnoredPasswordPolicyStateErrorCondition** | Pointer to [**[]EnumpluginIgnoredPasswordPolicyStateErrorConditionProp**](EnumpluginIgnoredPasswordPolicyStateErrorConditionProp.md) |  | [optional] 
**UserMappingLocalAttribute** | **[]string** | The names of the attributes in the local user entry whose values must match the values of the corresponding fields in the PingOne service. | 
**UserMappingRemoteJSONField** | **[]string** | The names of the fields in the PingOne service whose values must match the values of the corresponding attributes in the local user entry, as specified in the user-mapping-local-attribute property. | 
**AdditionalUserMappingSCIMFilter** | Pointer to **string** | An optional SCIM filter that will be ANDed with the filter created to identify the account in the PingOne service that corresponds to the local entry. Only the \&quot;eq\&quot;, \&quot;sw\&quot;, \&quot;and\&quot;, and \&quot;or\&quot; filter types may be used. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewPingOnePassThroughAuthenticationPluginResponse

`func NewPingOnePassThroughAuthenticationPluginResponse(id string, schemas []EnumpingOnePassThroughAuthenticationPluginSchemaUrn, apiURL string, authURL string, oAuthClientID string, environmentID string, userMappingLocalAttribute []string, userMappingRemoteJSONField []string, enabled bool, ) *PingOnePassThroughAuthenticationPluginResponse`

NewPingOnePassThroughAuthenticationPluginResponse instantiates a new PingOnePassThroughAuthenticationPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingOnePassThroughAuthenticationPluginResponseWithDefaults

`func NewPingOnePassThroughAuthenticationPluginResponseWithDefaults() *PingOnePassThroughAuthenticationPluginResponse`

NewPingOnePassThroughAuthenticationPluginResponseWithDefaults instantiates a new PingOnePassThroughAuthenticationPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetSchemas() []EnumpingOnePassThroughAuthenticationPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetSchemasOk() (*[]EnumpingOnePassThroughAuthenticationPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetSchemas(v []EnumpingOnePassThroughAuthenticationPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetApiURL

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.


### GetAuthURL

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetAuthURL() string`

GetAuthURL returns the AuthURL field if non-nil, zero value otherwise.

### GetAuthURLOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetAuthURLOk() (*string, bool)`

GetAuthURLOk returns a tuple with the AuthURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthURL

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetAuthURL(v string)`

SetAuthURL sets AuthURL field to given value.


### GetOAuthClientID

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetOAuthClientID() string`

GetOAuthClientID returns the OAuthClientID field if non-nil, zero value otherwise.

### GetOAuthClientIDOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetOAuthClientIDOk() (*string, bool)`

GetOAuthClientIDOk returns a tuple with the OAuthClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientID

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetOAuthClientID(v string)`

SetOAuthClientID sets OAuthClientID field to given value.


### GetOAuthClientSecret

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetOAuthClientSecret() string`

GetOAuthClientSecret returns the OAuthClientSecret field if non-nil, zero value otherwise.

### GetOAuthClientSecretOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetOAuthClientSecretOk() (*string, bool)`

GetOAuthClientSecretOk returns a tuple with the OAuthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientSecret

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetOAuthClientSecret(v string)`

SetOAuthClientSecret sets OAuthClientSecret field to given value.

### HasOAuthClientSecret

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasOAuthClientSecret() bool`

HasOAuthClientSecret returns a boolean if a field has been set.

### GetOAuthClientSecretPassphraseProvider

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetOAuthClientSecretPassphraseProvider() string`

GetOAuthClientSecretPassphraseProvider returns the OAuthClientSecretPassphraseProvider field if non-nil, zero value otherwise.

### GetOAuthClientSecretPassphraseProviderOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetOAuthClientSecretPassphraseProviderOk() (*string, bool)`

GetOAuthClientSecretPassphraseProviderOk returns a tuple with the OAuthClientSecretPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientSecretPassphraseProvider

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetOAuthClientSecretPassphraseProvider(v string)`

SetOAuthClientSecretPassphraseProvider sets OAuthClientSecretPassphraseProvider field to given value.

### HasOAuthClientSecretPassphraseProvider

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasOAuthClientSecretPassphraseProvider() bool`

HasOAuthClientSecretPassphraseProvider returns a boolean if a field has been set.

### GetEnvironmentID

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetEnvironmentID() string`

GetEnvironmentID returns the EnvironmentID field if non-nil, zero value otherwise.

### GetEnvironmentIDOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetEnvironmentIDOk() (*string, bool)`

GetEnvironmentIDOk returns a tuple with the EnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentID

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetEnvironmentID(v string)`

SetEnvironmentID sets EnvironmentID field to given value.


### GetHttpProxyExternalServer

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetIncludedLocalEntryBaseDN

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetIncludedLocalEntryBaseDN() []string`

GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field if non-nil, zero value otherwise.

### GetIncludedLocalEntryBaseDNOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetIncludedLocalEntryBaseDNOk() (*[]string, bool)`

GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLocalEntryBaseDN

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetIncludedLocalEntryBaseDN(v []string)`

SetIncludedLocalEntryBaseDN sets IncludedLocalEntryBaseDN field to given value.

### HasIncludedLocalEntryBaseDN

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasIncludedLocalEntryBaseDN() bool`

HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetTryLocalBind

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetTryLocalBind() bool`

GetTryLocalBind returns the TryLocalBind field if non-nil, zero value otherwise.

### GetTryLocalBindOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetTryLocalBindOk() (*bool, bool)`

GetTryLocalBindOk returns a tuple with the TryLocalBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryLocalBind

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetTryLocalBind(v bool)`

SetTryLocalBind sets TryLocalBind field to given value.

### HasTryLocalBind

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasTryLocalBind() bool`

HasTryLocalBind returns a boolean if a field has been set.

### GetOverrideLocalPassword

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetOverrideLocalPassword() bool`

GetOverrideLocalPassword returns the OverrideLocalPassword field if non-nil, zero value otherwise.

### GetOverrideLocalPasswordOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetOverrideLocalPasswordOk() (*bool, bool)`

GetOverrideLocalPasswordOk returns a tuple with the OverrideLocalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideLocalPassword

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetOverrideLocalPassword(v bool)`

SetOverrideLocalPassword sets OverrideLocalPassword field to given value.

### HasOverrideLocalPassword

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasOverrideLocalPassword() bool`

HasOverrideLocalPassword returns a boolean if a field has been set.

### GetUpdateLocalPassword

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetUpdateLocalPassword() bool`

GetUpdateLocalPassword returns the UpdateLocalPassword field if non-nil, zero value otherwise.

### GetUpdateLocalPasswordOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetUpdateLocalPasswordOk() (*bool, bool)`

GetUpdateLocalPasswordOk returns a tuple with the UpdateLocalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateLocalPassword

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetUpdateLocalPassword(v bool)`

SetUpdateLocalPassword sets UpdateLocalPassword field to given value.

### HasUpdateLocalPassword

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasUpdateLocalPassword() bool`

HasUpdateLocalPassword returns a boolean if a field has been set.

### GetUpdateLocalPasswordDN

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetUpdateLocalPasswordDN() string`

GetUpdateLocalPasswordDN returns the UpdateLocalPasswordDN field if non-nil, zero value otherwise.

### GetUpdateLocalPasswordDNOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetUpdateLocalPasswordDNOk() (*string, bool)`

GetUpdateLocalPasswordDNOk returns a tuple with the UpdateLocalPasswordDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateLocalPasswordDN

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetUpdateLocalPasswordDN(v string)`

SetUpdateLocalPasswordDN sets UpdateLocalPasswordDN field to given value.

### HasUpdateLocalPasswordDN

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasUpdateLocalPasswordDN() bool`

HasUpdateLocalPasswordDN returns a boolean if a field has been set.

### GetAllowLaxPassThroughAuthenticationPasswords

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetAllowLaxPassThroughAuthenticationPasswords() bool`

GetAllowLaxPassThroughAuthenticationPasswords returns the AllowLaxPassThroughAuthenticationPasswords field if non-nil, zero value otherwise.

### GetAllowLaxPassThroughAuthenticationPasswordsOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetAllowLaxPassThroughAuthenticationPasswordsOk() (*bool, bool)`

GetAllowLaxPassThroughAuthenticationPasswordsOk returns a tuple with the AllowLaxPassThroughAuthenticationPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLaxPassThroughAuthenticationPasswords

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetAllowLaxPassThroughAuthenticationPasswords(v bool)`

SetAllowLaxPassThroughAuthenticationPasswords sets AllowLaxPassThroughAuthenticationPasswords field to given value.

### HasAllowLaxPassThroughAuthenticationPasswords

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasAllowLaxPassThroughAuthenticationPasswords() bool`

HasAllowLaxPassThroughAuthenticationPasswords returns a boolean if a field has been set.

### GetIgnoredPasswordPolicyStateErrorCondition

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetIgnoredPasswordPolicyStateErrorCondition() []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp`

GetIgnoredPasswordPolicyStateErrorCondition returns the IgnoredPasswordPolicyStateErrorCondition field if non-nil, zero value otherwise.

### GetIgnoredPasswordPolicyStateErrorConditionOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetIgnoredPasswordPolicyStateErrorConditionOk() (*[]EnumpluginIgnoredPasswordPolicyStateErrorConditionProp, bool)`

GetIgnoredPasswordPolicyStateErrorConditionOk returns a tuple with the IgnoredPasswordPolicyStateErrorCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredPasswordPolicyStateErrorCondition

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetIgnoredPasswordPolicyStateErrorCondition(v []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp)`

SetIgnoredPasswordPolicyStateErrorCondition sets IgnoredPasswordPolicyStateErrorCondition field to given value.

### HasIgnoredPasswordPolicyStateErrorCondition

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasIgnoredPasswordPolicyStateErrorCondition() bool`

HasIgnoredPasswordPolicyStateErrorCondition returns a boolean if a field has been set.

### GetUserMappingLocalAttribute

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetUserMappingLocalAttribute() []string`

GetUserMappingLocalAttribute returns the UserMappingLocalAttribute field if non-nil, zero value otherwise.

### GetUserMappingLocalAttributeOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetUserMappingLocalAttributeOk() (*[]string, bool)`

GetUserMappingLocalAttributeOk returns a tuple with the UserMappingLocalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMappingLocalAttribute

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetUserMappingLocalAttribute(v []string)`

SetUserMappingLocalAttribute sets UserMappingLocalAttribute field to given value.


### GetUserMappingRemoteJSONField

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetUserMappingRemoteJSONField() []string`

GetUserMappingRemoteJSONField returns the UserMappingRemoteJSONField field if non-nil, zero value otherwise.

### GetUserMappingRemoteJSONFieldOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetUserMappingRemoteJSONFieldOk() (*[]string, bool)`

GetUserMappingRemoteJSONFieldOk returns a tuple with the UserMappingRemoteJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMappingRemoteJSONField

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetUserMappingRemoteJSONField(v []string)`

SetUserMappingRemoteJSONField sets UserMappingRemoteJSONField field to given value.


### GetAdditionalUserMappingSCIMFilter

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetAdditionalUserMappingSCIMFilter() string`

GetAdditionalUserMappingSCIMFilter returns the AdditionalUserMappingSCIMFilter field if non-nil, zero value otherwise.

### GetAdditionalUserMappingSCIMFilterOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetAdditionalUserMappingSCIMFilterOk() (*string, bool)`

GetAdditionalUserMappingSCIMFilterOk returns a tuple with the AdditionalUserMappingSCIMFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUserMappingSCIMFilter

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetAdditionalUserMappingSCIMFilter(v string)`

SetAdditionalUserMappingSCIMFilter sets AdditionalUserMappingSCIMFilter field to given value.

### HasAdditionalUserMappingSCIMFilter

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasAdditionalUserMappingSCIMFilter() bool`

HasAdditionalUserMappingSCIMFilter returns a boolean if a field has been set.

### GetDescription

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetMeta

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *PingOnePassThroughAuthenticationPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *PingOnePassThroughAuthenticationPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *PingOnePassThroughAuthenticationPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


