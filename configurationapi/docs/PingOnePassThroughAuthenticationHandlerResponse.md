# PingOnePassThroughAuthenticationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumpingOnePassThroughAuthenticationHandlerSchemaUrn**](EnumpingOnePassThroughAuthenticationHandlerSchemaUrn.md) |  | 
**ApiURL** | **string** | Specifies the API endpoint for the PingOne web service. | 
**AuthURL** | **string** | Specifies the API endpoint for the PingOne authentication service. | 
**OAuthClientID** | **string** | Specifies the OAuth Client ID used to authenticate connections to the PingOne API. | 
**OAuthClientSecret** | Pointer to **string** | Specifies the OAuth Client Secret used to authenticate connections to the PingOne API. | [optional] 
**OAuthClientSecretPassphraseProvider** | Pointer to **string** | Specifies a passphrase provider that can be used to obtain the OAuth Client Secret used to authenticate connections to the PingOne API. | [optional] 
**EnvironmentID** | **string** | Specifies the PingOne Environment that will be associated with this PingOne Pass Through Authentication Handler. | 
**HttpProxyExternalServer** | Pointer to **string** | A reference to an HTTP proxy server that should be used for requests sent to the PingOne service. | [optional] 
**UserMappingLocalAttribute** | **[]string** | The names of the attributes in the local user entry whose values must match the values of the corresponding fields in the PingOne service. | 
**UserMappingRemoteJSONField** | **[]string** | The names of the fields in the PingOne service whose values must match the values of the corresponding attributes in the local user entry, as specified in the user-mapping-local-attribute property. | 
**AdditionalUserMappingSCIMFilter** | Pointer to **string** | An optional SCIM filter that will be ANDed with the filter created to identify the account in the PingOne service that corresponds to the local entry. Only the \&quot;eq\&quot;, \&quot;sw\&quot;, \&quot;and\&quot;, and \&quot;or\&quot; filter types may be used. | [optional] 
**Description** | Pointer to **string** | A description for this Pass Through Authentication Handler | [optional] 
**IncludedLocalEntryBaseDN** | Pointer to **[]string** | The base DNs for the local users whose authentication attempts may be passed through to the external authentication service. | [optional] 
**ConnectionCriteria** | Pointer to **string** | A reference to connection criteria that will be used to indicate which bind requests should be passed through to the external authentication service. | [optional] 
**RequestCriteria** | Pointer to **string** | A reference to request criteria that will be used to indicate which bind requests should be passed through to the external authentication service. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Pass Through Authentication Handler | 

## Methods

### NewPingOnePassThroughAuthenticationHandlerResponse

`func NewPingOnePassThroughAuthenticationHandlerResponse(schemas []EnumpingOnePassThroughAuthenticationHandlerSchemaUrn, apiURL string, authURL string, oAuthClientID string, environmentID string, userMappingLocalAttribute []string, userMappingRemoteJSONField []string, id string, ) *PingOnePassThroughAuthenticationHandlerResponse`

NewPingOnePassThroughAuthenticationHandlerResponse instantiates a new PingOnePassThroughAuthenticationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingOnePassThroughAuthenticationHandlerResponseWithDefaults

`func NewPingOnePassThroughAuthenticationHandlerResponseWithDefaults() *PingOnePassThroughAuthenticationHandlerResponse`

NewPingOnePassThroughAuthenticationHandlerResponseWithDefaults instantiates a new PingOnePassThroughAuthenticationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetSchemas() []EnumpingOnePassThroughAuthenticationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetSchemasOk() (*[]EnumpingOnePassThroughAuthenticationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetSchemas(v []EnumpingOnePassThroughAuthenticationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetApiURL

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.


### GetAuthURL

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetAuthURL() string`

GetAuthURL returns the AuthURL field if non-nil, zero value otherwise.

### GetAuthURLOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetAuthURLOk() (*string, bool)`

GetAuthURLOk returns a tuple with the AuthURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthURL

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetAuthURL(v string)`

SetAuthURL sets AuthURL field to given value.


### GetOAuthClientID

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetOAuthClientID() string`

GetOAuthClientID returns the OAuthClientID field if non-nil, zero value otherwise.

### GetOAuthClientIDOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetOAuthClientIDOk() (*string, bool)`

GetOAuthClientIDOk returns a tuple with the OAuthClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientID

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetOAuthClientID(v string)`

SetOAuthClientID sets OAuthClientID field to given value.


### GetOAuthClientSecret

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetOAuthClientSecret() string`

GetOAuthClientSecret returns the OAuthClientSecret field if non-nil, zero value otherwise.

### GetOAuthClientSecretOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetOAuthClientSecretOk() (*string, bool)`

GetOAuthClientSecretOk returns a tuple with the OAuthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientSecret

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetOAuthClientSecret(v string)`

SetOAuthClientSecret sets OAuthClientSecret field to given value.

### HasOAuthClientSecret

`func (o *PingOnePassThroughAuthenticationHandlerResponse) HasOAuthClientSecret() bool`

HasOAuthClientSecret returns a boolean if a field has been set.

### GetOAuthClientSecretPassphraseProvider

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetOAuthClientSecretPassphraseProvider() string`

GetOAuthClientSecretPassphraseProvider returns the OAuthClientSecretPassphraseProvider field if non-nil, zero value otherwise.

### GetOAuthClientSecretPassphraseProviderOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetOAuthClientSecretPassphraseProviderOk() (*string, bool)`

GetOAuthClientSecretPassphraseProviderOk returns a tuple with the OAuthClientSecretPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientSecretPassphraseProvider

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetOAuthClientSecretPassphraseProvider(v string)`

SetOAuthClientSecretPassphraseProvider sets OAuthClientSecretPassphraseProvider field to given value.

### HasOAuthClientSecretPassphraseProvider

`func (o *PingOnePassThroughAuthenticationHandlerResponse) HasOAuthClientSecretPassphraseProvider() bool`

HasOAuthClientSecretPassphraseProvider returns a boolean if a field has been set.

### GetEnvironmentID

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetEnvironmentID() string`

GetEnvironmentID returns the EnvironmentID field if non-nil, zero value otherwise.

### GetEnvironmentIDOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetEnvironmentIDOk() (*string, bool)`

GetEnvironmentIDOk returns a tuple with the EnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentID

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetEnvironmentID(v string)`

SetEnvironmentID sets EnvironmentID field to given value.


### GetHttpProxyExternalServer

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *PingOnePassThroughAuthenticationHandlerResponse) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetUserMappingLocalAttribute

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetUserMappingLocalAttribute() []string`

GetUserMappingLocalAttribute returns the UserMappingLocalAttribute field if non-nil, zero value otherwise.

### GetUserMappingLocalAttributeOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetUserMappingLocalAttributeOk() (*[]string, bool)`

GetUserMappingLocalAttributeOk returns a tuple with the UserMappingLocalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMappingLocalAttribute

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetUserMappingLocalAttribute(v []string)`

SetUserMappingLocalAttribute sets UserMappingLocalAttribute field to given value.


### GetUserMappingRemoteJSONField

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetUserMappingRemoteJSONField() []string`

GetUserMappingRemoteJSONField returns the UserMappingRemoteJSONField field if non-nil, zero value otherwise.

### GetUserMappingRemoteJSONFieldOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetUserMappingRemoteJSONFieldOk() (*[]string, bool)`

GetUserMappingRemoteJSONFieldOk returns a tuple with the UserMappingRemoteJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMappingRemoteJSONField

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetUserMappingRemoteJSONField(v []string)`

SetUserMappingRemoteJSONField sets UserMappingRemoteJSONField field to given value.


### GetAdditionalUserMappingSCIMFilter

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetAdditionalUserMappingSCIMFilter() string`

GetAdditionalUserMappingSCIMFilter returns the AdditionalUserMappingSCIMFilter field if non-nil, zero value otherwise.

### GetAdditionalUserMappingSCIMFilterOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetAdditionalUserMappingSCIMFilterOk() (*string, bool)`

GetAdditionalUserMappingSCIMFilterOk returns a tuple with the AdditionalUserMappingSCIMFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUserMappingSCIMFilter

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetAdditionalUserMappingSCIMFilter(v string)`

SetAdditionalUserMappingSCIMFilter sets AdditionalUserMappingSCIMFilter field to given value.

### HasAdditionalUserMappingSCIMFilter

`func (o *PingOnePassThroughAuthenticationHandlerResponse) HasAdditionalUserMappingSCIMFilter() bool`

HasAdditionalUserMappingSCIMFilter returns a boolean if a field has been set.

### GetDescription

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PingOnePassThroughAuthenticationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncludedLocalEntryBaseDN

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetIncludedLocalEntryBaseDN() []string`

GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field if non-nil, zero value otherwise.

### GetIncludedLocalEntryBaseDNOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetIncludedLocalEntryBaseDNOk() (*[]string, bool)`

GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLocalEntryBaseDN

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetIncludedLocalEntryBaseDN(v []string)`

SetIncludedLocalEntryBaseDN sets IncludedLocalEntryBaseDN field to given value.

### HasIncludedLocalEntryBaseDN

`func (o *PingOnePassThroughAuthenticationHandlerResponse) HasIncludedLocalEntryBaseDN() bool`

HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *PingOnePassThroughAuthenticationHandlerResponse) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *PingOnePassThroughAuthenticationHandlerResponse) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetMeta

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PingOnePassThroughAuthenticationHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *PingOnePassThroughAuthenticationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PingOnePassThroughAuthenticationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PingOnePassThroughAuthenticationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


