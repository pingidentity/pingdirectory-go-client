# ScimExternalServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumscimExternalServerSchemaUrn**](EnumscimExternalServerSchemaUrn.md) |  | 
**Id** | **string** | Name of the External Server | 
**ScimServiceURL** | **string** | The complete URL which will be used to access the SCIM service provider. | 
**UserName** | Pointer to **string** | The name of the login account to use when connecting to the SCIM service provider. This is used in conjunction with the chosen authentication-method. | [optional] 
**Password** | Pointer to **string** | The login password for the specified user name. This is used in conjunction with the chosen authentication-method. | [optional] 
**PassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the login password for the specified user. | [optional] 
**Location** | Pointer to **string** | Specifies the location for the SCIM External Server. | [optional] 
**ConnectionSecurity** | [**EnumexternalServerScimConnectionSecurityProp**](EnumexternalServerScimConnectionSecurityProp.md) |  | 
**AuthenticationMethod** | [**EnumexternalServerScimAuthenticationMethodProp**](EnumexternalServerScimAuthenticationMethodProp.md) |  | 
**HostnameVerificationMethod** | Pointer to [**EnumexternalServerScimHostnameVerificationMethodProp**](EnumexternalServerScimHostnameVerificationMethodProp.md) |  | [optional] 
**HttpProxyExternalServer** | Pointer to **string** | A reference to an HTTP proxy server that should be used for requests sent to the SCIM service. | [optional] 
**KeyManagerProvider** | Pointer to **string** | The key manager provider to use if SSL is to be used for connection-level security. When specifying a value for this property (except when using the Null key manager provider) you must ensure that the external server trusts this server&#39;s public certificate by adding this server&#39;s public certificate to the external server&#39;s trust store. | [optional] 
**TrustManagerProvider** | Pointer to **string** | The trust manager provider to use if SSL is to be used for connection-level security. | [optional] 
**ConnectTimeout** | Pointer to **string** | Specifies the amount of time to wait for a response from the service provider when establishing a connection. If the timeout is exceeded, the Directory Server will attempt to fail over to a different server. A value of zero indicates no timeout. | [optional] 
**ResponseTimeout** | Pointer to **string** | Specifies the maximum length of time that an operation should be allowed to block while waiting for a response from the SCIM service provider. A value of zero indicates that there should be no client-side timeout. | [optional] 
**OAuthTokenType** | Pointer to [**EnumexternalServerOAuthTokenTypeProp**](EnumexternalServerOAuthTokenTypeProp.md) |  | [optional] 
**OAuthToken** | Pointer to **string** | The token to use in conjunction with the OAuth authentication-method and the chosen oauth-token-type. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 

## Methods

### NewScimExternalServerResponse

`func NewScimExternalServerResponse(schemas []EnumscimExternalServerSchemaUrn, id string, scimServiceURL string, connectionSecurity EnumexternalServerScimConnectionSecurityProp, authenticationMethod EnumexternalServerScimAuthenticationMethodProp, ) *ScimExternalServerResponse`

NewScimExternalServerResponse instantiates a new ScimExternalServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimExternalServerResponseWithDefaults

`func NewScimExternalServerResponseWithDefaults() *ScimExternalServerResponse`

NewScimExternalServerResponseWithDefaults instantiates a new ScimExternalServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ScimExternalServerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScimExternalServerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScimExternalServerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ScimExternalServerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ScimExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ScimExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ScimExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ScimExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *ScimExternalServerResponse) GetSchemas() []EnumscimExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimExternalServerResponse) GetSchemasOk() (*[]EnumscimExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimExternalServerResponse) SetSchemas(v []EnumscimExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ScimExternalServerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimExternalServerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimExternalServerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetScimServiceURL

`func (o *ScimExternalServerResponse) GetScimServiceURL() string`

GetScimServiceURL returns the ScimServiceURL field if non-nil, zero value otherwise.

### GetScimServiceURLOk

`func (o *ScimExternalServerResponse) GetScimServiceURLOk() (*string, bool)`

GetScimServiceURLOk returns a tuple with the ScimServiceURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimServiceURL

`func (o *ScimExternalServerResponse) SetScimServiceURL(v string)`

SetScimServiceURL sets ScimServiceURL field to given value.


### GetUserName

`func (o *ScimExternalServerResponse) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ScimExternalServerResponse) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ScimExternalServerResponse) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ScimExternalServerResponse) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *ScimExternalServerResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ScimExternalServerResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ScimExternalServerResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ScimExternalServerResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPassphraseProvider

`func (o *ScimExternalServerResponse) GetPassphraseProvider() string`

GetPassphraseProvider returns the PassphraseProvider field if non-nil, zero value otherwise.

### GetPassphraseProviderOk

`func (o *ScimExternalServerResponse) GetPassphraseProviderOk() (*string, bool)`

GetPassphraseProviderOk returns a tuple with the PassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseProvider

`func (o *ScimExternalServerResponse) SetPassphraseProvider(v string)`

SetPassphraseProvider sets PassphraseProvider field to given value.

### HasPassphraseProvider

`func (o *ScimExternalServerResponse) HasPassphraseProvider() bool`

HasPassphraseProvider returns a boolean if a field has been set.

### GetLocation

`func (o *ScimExternalServerResponse) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ScimExternalServerResponse) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ScimExternalServerResponse) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ScimExternalServerResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetConnectionSecurity

`func (o *ScimExternalServerResponse) GetConnectionSecurity() EnumexternalServerScimConnectionSecurityProp`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *ScimExternalServerResponse) GetConnectionSecurityOk() (*EnumexternalServerScimConnectionSecurityProp, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *ScimExternalServerResponse) SetConnectionSecurity(v EnumexternalServerScimConnectionSecurityProp)`

SetConnectionSecurity sets ConnectionSecurity field to given value.


### GetAuthenticationMethod

`func (o *ScimExternalServerResponse) GetAuthenticationMethod() EnumexternalServerScimAuthenticationMethodProp`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *ScimExternalServerResponse) GetAuthenticationMethodOk() (*EnumexternalServerScimAuthenticationMethodProp, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *ScimExternalServerResponse) SetAuthenticationMethod(v EnumexternalServerScimAuthenticationMethodProp)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.


### GetHostnameVerificationMethod

`func (o *ScimExternalServerResponse) GetHostnameVerificationMethod() EnumexternalServerScimHostnameVerificationMethodProp`

GetHostnameVerificationMethod returns the HostnameVerificationMethod field if non-nil, zero value otherwise.

### GetHostnameVerificationMethodOk

`func (o *ScimExternalServerResponse) GetHostnameVerificationMethodOk() (*EnumexternalServerScimHostnameVerificationMethodProp, bool)`

GetHostnameVerificationMethodOk returns a tuple with the HostnameVerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameVerificationMethod

`func (o *ScimExternalServerResponse) SetHostnameVerificationMethod(v EnumexternalServerScimHostnameVerificationMethodProp)`

SetHostnameVerificationMethod sets HostnameVerificationMethod field to given value.

### HasHostnameVerificationMethod

`func (o *ScimExternalServerResponse) HasHostnameVerificationMethod() bool`

HasHostnameVerificationMethod returns a boolean if a field has been set.

### GetHttpProxyExternalServer

`func (o *ScimExternalServerResponse) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *ScimExternalServerResponse) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *ScimExternalServerResponse) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *ScimExternalServerResponse) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *ScimExternalServerResponse) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *ScimExternalServerResponse) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *ScimExternalServerResponse) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *ScimExternalServerResponse) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *ScimExternalServerResponse) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *ScimExternalServerResponse) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *ScimExternalServerResponse) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *ScimExternalServerResponse) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *ScimExternalServerResponse) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *ScimExternalServerResponse) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *ScimExternalServerResponse) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *ScimExternalServerResponse) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetResponseTimeout

`func (o *ScimExternalServerResponse) GetResponseTimeout() string`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *ScimExternalServerResponse) GetResponseTimeoutOk() (*string, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *ScimExternalServerResponse) SetResponseTimeout(v string)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *ScimExternalServerResponse) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### GetOAuthTokenType

`func (o *ScimExternalServerResponse) GetOAuthTokenType() EnumexternalServerOAuthTokenTypeProp`

GetOAuthTokenType returns the OAuthTokenType field if non-nil, zero value otherwise.

### GetOAuthTokenTypeOk

`func (o *ScimExternalServerResponse) GetOAuthTokenTypeOk() (*EnumexternalServerOAuthTokenTypeProp, bool)`

GetOAuthTokenTypeOk returns a tuple with the OAuthTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthTokenType

`func (o *ScimExternalServerResponse) SetOAuthTokenType(v EnumexternalServerOAuthTokenTypeProp)`

SetOAuthTokenType sets OAuthTokenType field to given value.

### HasOAuthTokenType

`func (o *ScimExternalServerResponse) HasOAuthTokenType() bool`

HasOAuthTokenType returns a boolean if a field has been set.

### GetOAuthToken

`func (o *ScimExternalServerResponse) GetOAuthToken() string`

GetOAuthToken returns the OAuthToken field if non-nil, zero value otherwise.

### GetOAuthTokenOk

`func (o *ScimExternalServerResponse) GetOAuthTokenOk() (*string, bool)`

GetOAuthTokenOk returns a tuple with the OAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthToken

`func (o *ScimExternalServerResponse) SetOAuthToken(v string)`

SetOAuthToken sets OAuthToken field to given value.

### HasOAuthToken

`func (o *ScimExternalServerResponse) HasOAuthToken() bool`

HasOAuthToken returns a boolean if a field has been set.

### GetDescription

`func (o *ScimExternalServerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScimExternalServerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScimExternalServerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScimExternalServerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


