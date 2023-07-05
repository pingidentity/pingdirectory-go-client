# AddPassThroughAuthenticationHandler200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Pass Through Authentication Handler | 
**Schemas** | [**[]EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn**](EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn.md) |  | 
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
**Server** | **[]string** | Specifies the LDAP external server(s) to which authentication attempts should be forwarded. | 
**ServerAccessMode** | [**EnumpassThroughAuthenticationHandlerServerAccessModeProp**](EnumpassThroughAuthenticationHandlerServerAccessModeProp.md) |  | 
**DnMap** | Pointer to **[]string** | Specifies one or more DN mappings that may be used to transform bind DNs before attempting to bind to the external servers. | [optional] 
**BindDNPattern** | Pointer to **string** | A pattern to use to construct the bind DN for the simple bind request to send to the remote server. This may consist of a combination of static text and attribute values and other directives enclosed in curly braces.  For example, the value \&quot;cn&#x3D;{cn},ou&#x3D;People,dc&#x3D;example,dc&#x3D;com\&quot; indicates that the remote bind DN should be constructed from the text \&quot;cn&#x3D;\&quot; followed by the value of the local entry&#39;s cn attribute followed by the text \&quot;ou&#x3D;People,dc&#x3D;example,dc&#x3D;com\&quot;. If an attribute contains the value to use as the bind DN for pass-through authentication, then the pattern may simply be the name of that attribute in curly braces (e.g., if the seeAlso attribute contains the bind DN for the target user, then a bind DN pattern of \&quot;{seeAlso}\&quot; would be appropriate).  Note that a bind DN pattern can be used to construct a bind DN that is not actually a valid LDAP distinguished name. For example, if authentication is being passed through to a Microsoft Active Directory server, then a bind DN pattern could be used to construct a user principal name (UPN) as an alternative to a distinguished name. | [optional] 
**SearchBaseDN** | Pointer to **string** | The base DN to use when searching for the user entry using a filter constructed from the pattern defined in the search-filter-pattern property. If no base DN is specified, the null DN will be used as the search base DN. | [optional] 
**SearchFilterPattern** | Pointer to **string** | A pattern to use to construct a filter to use when searching an external server for the entry of the user as whom to bind. For example, \&quot;(mail&#x3D;{uid:ldapFilterEscape}@example.com)\&quot; would construct a search filter to search for a user whose entry in the local server contains a uid attribute whose value appears before \&quot;@example.com\&quot; in the mail attribute in the external server. Note that the \&quot;ldapFilterEscape\&quot; modifier should almost always be used with attributes specified in the pattern. | [optional] 
**InitialConnections** | **int64** | Specifies the initial number of connections to establish to each external server against which authentication may be attempted. | 
**MaxConnections** | **int64** | Specifies the maximum number of connections to maintain to each external server against which authentication may be attempted. This value must be greater than or equal to the value for the initial-connections property. | 
**UseLocation** | Pointer to **bool** | Indicates whether to take server locations into account when prioritizing the servers to use for pass-through authentication attempts. | [optional] 
**MaximumAllowedLocalResponseTime** | Pointer to **string** | The maximum length of time to wait for a response from an external server in the same location as this Directory Server before considering it unavailable. | [optional] 
**MaximumAllowedNonlocalResponseTime** | Pointer to **string** | The maximum length of time to wait for a response from an external server in a different location from this Directory Server before considering it unavailable. | [optional] 
**UsePasswordPolicyControl** | Pointer to **bool** | Indicates whether to include the password policy request control (as defined in draft-behera-ldap-password-policy-10) in bind requests sent to the external server. | [optional] 
**SubordinatePassThroughAuthenticationHandler** | **[]string** | The set of subordinate pass-through authentication handlers that may be used to perform the authentication processing. Handlers will be invoked in order until one is found for which the bind operation matches the associated criteria and either succeeds or fails in a manner that should not be ignored. | 
**ContinueOnFailureType** | Pointer to [**[]EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp**](EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp.md) |  | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Pass Through Authentication Handler. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Pass Through Authentication Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddPassThroughAuthenticationHandler200Response

`func NewAddPassThroughAuthenticationHandler200Response(id string, schemas []EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn, apiURL string, authURL string, oAuthClientID string, environmentID string, userMappingLocalAttribute []string, userMappingRemoteJSONField []string, server []string, serverAccessMode EnumpassThroughAuthenticationHandlerServerAccessModeProp, initialConnections int64, maxConnections int64, subordinatePassThroughAuthenticationHandler []string, extensionClass string, ) *AddPassThroughAuthenticationHandler200Response`

NewAddPassThroughAuthenticationHandler200Response instantiates a new AddPassThroughAuthenticationHandler200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPassThroughAuthenticationHandler200ResponseWithDefaults

`func NewAddPassThroughAuthenticationHandler200ResponseWithDefaults() *AddPassThroughAuthenticationHandler200Response`

NewAddPassThroughAuthenticationHandler200ResponseWithDefaults instantiates a new AddPassThroughAuthenticationHandler200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddPassThroughAuthenticationHandler200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddPassThroughAuthenticationHandler200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddPassThroughAuthenticationHandler200Response) GetSchemas() []EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetSchemasOk() (*[]EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPassThroughAuthenticationHandler200Response) SetSchemas(v []EnumthirdPartyPassThroughAuthenticationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetApiURL

`func (o *AddPassThroughAuthenticationHandler200Response) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *AddPassThroughAuthenticationHandler200Response) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.


### GetAuthURL

`func (o *AddPassThroughAuthenticationHandler200Response) GetAuthURL() string`

GetAuthURL returns the AuthURL field if non-nil, zero value otherwise.

### GetAuthURLOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetAuthURLOk() (*string, bool)`

GetAuthURLOk returns a tuple with the AuthURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthURL

`func (o *AddPassThroughAuthenticationHandler200Response) SetAuthURL(v string)`

SetAuthURL sets AuthURL field to given value.


### GetOAuthClientID

`func (o *AddPassThroughAuthenticationHandler200Response) GetOAuthClientID() string`

GetOAuthClientID returns the OAuthClientID field if non-nil, zero value otherwise.

### GetOAuthClientIDOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetOAuthClientIDOk() (*string, bool)`

GetOAuthClientIDOk returns a tuple with the OAuthClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientID

`func (o *AddPassThroughAuthenticationHandler200Response) SetOAuthClientID(v string)`

SetOAuthClientID sets OAuthClientID field to given value.


### GetOAuthClientSecret

`func (o *AddPassThroughAuthenticationHandler200Response) GetOAuthClientSecret() string`

GetOAuthClientSecret returns the OAuthClientSecret field if non-nil, zero value otherwise.

### GetOAuthClientSecretOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetOAuthClientSecretOk() (*string, bool)`

GetOAuthClientSecretOk returns a tuple with the OAuthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientSecret

`func (o *AddPassThroughAuthenticationHandler200Response) SetOAuthClientSecret(v string)`

SetOAuthClientSecret sets OAuthClientSecret field to given value.

### HasOAuthClientSecret

`func (o *AddPassThroughAuthenticationHandler200Response) HasOAuthClientSecret() bool`

HasOAuthClientSecret returns a boolean if a field has been set.

### GetOAuthClientSecretPassphraseProvider

`func (o *AddPassThroughAuthenticationHandler200Response) GetOAuthClientSecretPassphraseProvider() string`

GetOAuthClientSecretPassphraseProvider returns the OAuthClientSecretPassphraseProvider field if non-nil, zero value otherwise.

### GetOAuthClientSecretPassphraseProviderOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetOAuthClientSecretPassphraseProviderOk() (*string, bool)`

GetOAuthClientSecretPassphraseProviderOk returns a tuple with the OAuthClientSecretPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientSecretPassphraseProvider

`func (o *AddPassThroughAuthenticationHandler200Response) SetOAuthClientSecretPassphraseProvider(v string)`

SetOAuthClientSecretPassphraseProvider sets OAuthClientSecretPassphraseProvider field to given value.

### HasOAuthClientSecretPassphraseProvider

`func (o *AddPassThroughAuthenticationHandler200Response) HasOAuthClientSecretPassphraseProvider() bool`

HasOAuthClientSecretPassphraseProvider returns a boolean if a field has been set.

### GetEnvironmentID

`func (o *AddPassThroughAuthenticationHandler200Response) GetEnvironmentID() string`

GetEnvironmentID returns the EnvironmentID field if non-nil, zero value otherwise.

### GetEnvironmentIDOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetEnvironmentIDOk() (*string, bool)`

GetEnvironmentIDOk returns a tuple with the EnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentID

`func (o *AddPassThroughAuthenticationHandler200Response) SetEnvironmentID(v string)`

SetEnvironmentID sets EnvironmentID field to given value.


### GetHttpProxyExternalServer

`func (o *AddPassThroughAuthenticationHandler200Response) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *AddPassThroughAuthenticationHandler200Response) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *AddPassThroughAuthenticationHandler200Response) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetUserMappingLocalAttribute

`func (o *AddPassThroughAuthenticationHandler200Response) GetUserMappingLocalAttribute() []string`

GetUserMappingLocalAttribute returns the UserMappingLocalAttribute field if non-nil, zero value otherwise.

### GetUserMappingLocalAttributeOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetUserMappingLocalAttributeOk() (*[]string, bool)`

GetUserMappingLocalAttributeOk returns a tuple with the UserMappingLocalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMappingLocalAttribute

`func (o *AddPassThroughAuthenticationHandler200Response) SetUserMappingLocalAttribute(v []string)`

SetUserMappingLocalAttribute sets UserMappingLocalAttribute field to given value.


### GetUserMappingRemoteJSONField

`func (o *AddPassThroughAuthenticationHandler200Response) GetUserMappingRemoteJSONField() []string`

GetUserMappingRemoteJSONField returns the UserMappingRemoteJSONField field if non-nil, zero value otherwise.

### GetUserMappingRemoteJSONFieldOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetUserMappingRemoteJSONFieldOk() (*[]string, bool)`

GetUserMappingRemoteJSONFieldOk returns a tuple with the UserMappingRemoteJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMappingRemoteJSONField

`func (o *AddPassThroughAuthenticationHandler200Response) SetUserMappingRemoteJSONField(v []string)`

SetUserMappingRemoteJSONField sets UserMappingRemoteJSONField field to given value.


### GetAdditionalUserMappingSCIMFilter

`func (o *AddPassThroughAuthenticationHandler200Response) GetAdditionalUserMappingSCIMFilter() string`

GetAdditionalUserMappingSCIMFilter returns the AdditionalUserMappingSCIMFilter field if non-nil, zero value otherwise.

### GetAdditionalUserMappingSCIMFilterOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetAdditionalUserMappingSCIMFilterOk() (*string, bool)`

GetAdditionalUserMappingSCIMFilterOk returns a tuple with the AdditionalUserMappingSCIMFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUserMappingSCIMFilter

`func (o *AddPassThroughAuthenticationHandler200Response) SetAdditionalUserMappingSCIMFilter(v string)`

SetAdditionalUserMappingSCIMFilter sets AdditionalUserMappingSCIMFilter field to given value.

### HasAdditionalUserMappingSCIMFilter

`func (o *AddPassThroughAuthenticationHandler200Response) HasAdditionalUserMappingSCIMFilter() bool`

HasAdditionalUserMappingSCIMFilter returns a boolean if a field has been set.

### GetDescription

`func (o *AddPassThroughAuthenticationHandler200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPassThroughAuthenticationHandler200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPassThroughAuthenticationHandler200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncludedLocalEntryBaseDN

`func (o *AddPassThroughAuthenticationHandler200Response) GetIncludedLocalEntryBaseDN() []string`

GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field if non-nil, zero value otherwise.

### GetIncludedLocalEntryBaseDNOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetIncludedLocalEntryBaseDNOk() (*[]string, bool)`

GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLocalEntryBaseDN

`func (o *AddPassThroughAuthenticationHandler200Response) SetIncludedLocalEntryBaseDN(v []string)`

SetIncludedLocalEntryBaseDN sets IncludedLocalEntryBaseDN field to given value.

### HasIncludedLocalEntryBaseDN

`func (o *AddPassThroughAuthenticationHandler200Response) HasIncludedLocalEntryBaseDN() bool`

HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *AddPassThroughAuthenticationHandler200Response) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *AddPassThroughAuthenticationHandler200Response) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *AddPassThroughAuthenticationHandler200Response) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *AddPassThroughAuthenticationHandler200Response) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddPassThroughAuthenticationHandler200Response) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddPassThroughAuthenticationHandler200Response) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetMeta

`func (o *AddPassThroughAuthenticationHandler200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddPassThroughAuthenticationHandler200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddPassThroughAuthenticationHandler200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddPassThroughAuthenticationHandler200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddPassThroughAuthenticationHandler200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddPassThroughAuthenticationHandler200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddPassThroughAuthenticationHandler200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetServer

`func (o *AddPassThroughAuthenticationHandler200Response) GetServer() []string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetServerOk() (*[]string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AddPassThroughAuthenticationHandler200Response) SetServer(v []string)`

SetServer sets Server field to given value.


### GetServerAccessMode

`func (o *AddPassThroughAuthenticationHandler200Response) GetServerAccessMode() EnumpassThroughAuthenticationHandlerServerAccessModeProp`

GetServerAccessMode returns the ServerAccessMode field if non-nil, zero value otherwise.

### GetServerAccessModeOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetServerAccessModeOk() (*EnumpassThroughAuthenticationHandlerServerAccessModeProp, bool)`

GetServerAccessModeOk returns a tuple with the ServerAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAccessMode

`func (o *AddPassThroughAuthenticationHandler200Response) SetServerAccessMode(v EnumpassThroughAuthenticationHandlerServerAccessModeProp)`

SetServerAccessMode sets ServerAccessMode field to given value.


### GetDnMap

`func (o *AddPassThroughAuthenticationHandler200Response) GetDnMap() []string`

GetDnMap returns the DnMap field if non-nil, zero value otherwise.

### GetDnMapOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetDnMapOk() (*[]string, bool)`

GetDnMapOk returns a tuple with the DnMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnMap

`func (o *AddPassThroughAuthenticationHandler200Response) SetDnMap(v []string)`

SetDnMap sets DnMap field to given value.

### HasDnMap

`func (o *AddPassThroughAuthenticationHandler200Response) HasDnMap() bool`

HasDnMap returns a boolean if a field has been set.

### GetBindDNPattern

`func (o *AddPassThroughAuthenticationHandler200Response) GetBindDNPattern() string`

GetBindDNPattern returns the BindDNPattern field if non-nil, zero value otherwise.

### GetBindDNPatternOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetBindDNPatternOk() (*string, bool)`

GetBindDNPatternOk returns a tuple with the BindDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDNPattern

`func (o *AddPassThroughAuthenticationHandler200Response) SetBindDNPattern(v string)`

SetBindDNPattern sets BindDNPattern field to given value.

### HasBindDNPattern

`func (o *AddPassThroughAuthenticationHandler200Response) HasBindDNPattern() bool`

HasBindDNPattern returns a boolean if a field has been set.

### GetSearchBaseDN

`func (o *AddPassThroughAuthenticationHandler200Response) GetSearchBaseDN() string`

GetSearchBaseDN returns the SearchBaseDN field if non-nil, zero value otherwise.

### GetSearchBaseDNOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetSearchBaseDNOk() (*string, bool)`

GetSearchBaseDNOk returns a tuple with the SearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBaseDN

`func (o *AddPassThroughAuthenticationHandler200Response) SetSearchBaseDN(v string)`

SetSearchBaseDN sets SearchBaseDN field to given value.

### HasSearchBaseDN

`func (o *AddPassThroughAuthenticationHandler200Response) HasSearchBaseDN() bool`

HasSearchBaseDN returns a boolean if a field has been set.

### GetSearchFilterPattern

`func (o *AddPassThroughAuthenticationHandler200Response) GetSearchFilterPattern() string`

GetSearchFilterPattern returns the SearchFilterPattern field if non-nil, zero value otherwise.

### GetSearchFilterPatternOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetSearchFilterPatternOk() (*string, bool)`

GetSearchFilterPatternOk returns a tuple with the SearchFilterPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterPattern

`func (o *AddPassThroughAuthenticationHandler200Response) SetSearchFilterPattern(v string)`

SetSearchFilterPattern sets SearchFilterPattern field to given value.

### HasSearchFilterPattern

`func (o *AddPassThroughAuthenticationHandler200Response) HasSearchFilterPattern() bool`

HasSearchFilterPattern returns a boolean if a field has been set.

### GetInitialConnections

`func (o *AddPassThroughAuthenticationHandler200Response) GetInitialConnections() int64`

GetInitialConnections returns the InitialConnections field if non-nil, zero value otherwise.

### GetInitialConnectionsOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetInitialConnectionsOk() (*int64, bool)`

GetInitialConnectionsOk returns a tuple with the InitialConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialConnections

`func (o *AddPassThroughAuthenticationHandler200Response) SetInitialConnections(v int64)`

SetInitialConnections sets InitialConnections field to given value.


### GetMaxConnections

`func (o *AddPassThroughAuthenticationHandler200Response) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *AddPassThroughAuthenticationHandler200Response) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.


### GetUseLocation

`func (o *AddPassThroughAuthenticationHandler200Response) GetUseLocation() bool`

GetUseLocation returns the UseLocation field if non-nil, zero value otherwise.

### GetUseLocationOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetUseLocationOk() (*bool, bool)`

GetUseLocationOk returns a tuple with the UseLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocation

`func (o *AddPassThroughAuthenticationHandler200Response) SetUseLocation(v bool)`

SetUseLocation sets UseLocation field to given value.

### HasUseLocation

`func (o *AddPassThroughAuthenticationHandler200Response) HasUseLocation() bool`

HasUseLocation returns a boolean if a field has been set.

### GetMaximumAllowedLocalResponseTime

`func (o *AddPassThroughAuthenticationHandler200Response) GetMaximumAllowedLocalResponseTime() string`

GetMaximumAllowedLocalResponseTime returns the MaximumAllowedLocalResponseTime field if non-nil, zero value otherwise.

### GetMaximumAllowedLocalResponseTimeOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetMaximumAllowedLocalResponseTimeOk() (*string, bool)`

GetMaximumAllowedLocalResponseTimeOk returns a tuple with the MaximumAllowedLocalResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAllowedLocalResponseTime

`func (o *AddPassThroughAuthenticationHandler200Response) SetMaximumAllowedLocalResponseTime(v string)`

SetMaximumAllowedLocalResponseTime sets MaximumAllowedLocalResponseTime field to given value.

### HasMaximumAllowedLocalResponseTime

`func (o *AddPassThroughAuthenticationHandler200Response) HasMaximumAllowedLocalResponseTime() bool`

HasMaximumAllowedLocalResponseTime returns a boolean if a field has been set.

### GetMaximumAllowedNonlocalResponseTime

`func (o *AddPassThroughAuthenticationHandler200Response) GetMaximumAllowedNonlocalResponseTime() string`

GetMaximumAllowedNonlocalResponseTime returns the MaximumAllowedNonlocalResponseTime field if non-nil, zero value otherwise.

### GetMaximumAllowedNonlocalResponseTimeOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetMaximumAllowedNonlocalResponseTimeOk() (*string, bool)`

GetMaximumAllowedNonlocalResponseTimeOk returns a tuple with the MaximumAllowedNonlocalResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAllowedNonlocalResponseTime

`func (o *AddPassThroughAuthenticationHandler200Response) SetMaximumAllowedNonlocalResponseTime(v string)`

SetMaximumAllowedNonlocalResponseTime sets MaximumAllowedNonlocalResponseTime field to given value.

### HasMaximumAllowedNonlocalResponseTime

`func (o *AddPassThroughAuthenticationHandler200Response) HasMaximumAllowedNonlocalResponseTime() bool`

HasMaximumAllowedNonlocalResponseTime returns a boolean if a field has been set.

### GetUsePasswordPolicyControl

`func (o *AddPassThroughAuthenticationHandler200Response) GetUsePasswordPolicyControl() bool`

GetUsePasswordPolicyControl returns the UsePasswordPolicyControl field if non-nil, zero value otherwise.

### GetUsePasswordPolicyControlOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetUsePasswordPolicyControlOk() (*bool, bool)`

GetUsePasswordPolicyControlOk returns a tuple with the UsePasswordPolicyControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePasswordPolicyControl

`func (o *AddPassThroughAuthenticationHandler200Response) SetUsePasswordPolicyControl(v bool)`

SetUsePasswordPolicyControl sets UsePasswordPolicyControl field to given value.

### HasUsePasswordPolicyControl

`func (o *AddPassThroughAuthenticationHandler200Response) HasUsePasswordPolicyControl() bool`

HasUsePasswordPolicyControl returns a boolean if a field has been set.

### GetSubordinatePassThroughAuthenticationHandler

`func (o *AddPassThroughAuthenticationHandler200Response) GetSubordinatePassThroughAuthenticationHandler() []string`

GetSubordinatePassThroughAuthenticationHandler returns the SubordinatePassThroughAuthenticationHandler field if non-nil, zero value otherwise.

### GetSubordinatePassThroughAuthenticationHandlerOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetSubordinatePassThroughAuthenticationHandlerOk() (*[]string, bool)`

GetSubordinatePassThroughAuthenticationHandlerOk returns a tuple with the SubordinatePassThroughAuthenticationHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubordinatePassThroughAuthenticationHandler

`func (o *AddPassThroughAuthenticationHandler200Response) SetSubordinatePassThroughAuthenticationHandler(v []string)`

SetSubordinatePassThroughAuthenticationHandler sets SubordinatePassThroughAuthenticationHandler field to given value.


### GetContinueOnFailureType

`func (o *AddPassThroughAuthenticationHandler200Response) GetContinueOnFailureType() []EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp`

GetContinueOnFailureType returns the ContinueOnFailureType field if non-nil, zero value otherwise.

### GetContinueOnFailureTypeOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetContinueOnFailureTypeOk() (*[]EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp, bool)`

GetContinueOnFailureTypeOk returns a tuple with the ContinueOnFailureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnFailureType

`func (o *AddPassThroughAuthenticationHandler200Response) SetContinueOnFailureType(v []EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp)`

SetContinueOnFailureType sets ContinueOnFailureType field to given value.

### HasContinueOnFailureType

`func (o *AddPassThroughAuthenticationHandler200Response) HasContinueOnFailureType() bool`

HasContinueOnFailureType returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddPassThroughAuthenticationHandler200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddPassThroughAuthenticationHandler200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddPassThroughAuthenticationHandler200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddPassThroughAuthenticationHandler200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddPassThroughAuthenticationHandler200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddPassThroughAuthenticationHandler200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


