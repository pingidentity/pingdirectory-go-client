# NokiaProxyServerExternalServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the External Server | 
**Schemas** | [**[]EnumnokiaProxyServerExternalServerSchemaUrn**](EnumnokiaProxyServerExternalServerSchemaUrn.md) |  | 
**VerifyCredentialsMethod** | [**EnumexternalServerVerifyCredentialsMethodProp**](EnumexternalServerVerifyCredentialsMethodProp.md) |  | 
**UseAdministrativeOperationControl** | Pointer to **bool** | Indicates whether to include the administrative operation request control in requests sent to this server which are intended for administrative operations (e.g., health checking) rather than requests directly from clients. | [optional] 
**ServerHostName** | **string** | The host name or IP address of the target LDAP server. | 
**ServerPort** | **int64** | The port number on which the server listens for requests. | 
**Location** | Pointer to **string** | Specifies the location for the LDAP External Server. | [optional] 
**BindDN** | Pointer to **string** | The DN to use to bind to the target LDAP server if simple authentication is required. | [optional] 
**Password** | Pointer to **string** | The login password for the specified user. | [optional] 
**PassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the login password for the specified user. | [optional] 
**ConnectionSecurity** | [**EnumexternalServerConnectionSecurityProp**](EnumexternalServerConnectionSecurityProp.md) |  | 
**AuthenticationMethod** | [**EnumexternalServerNokiaProxyServerAuthenticationMethodProp**](EnumexternalServerNokiaProxyServerAuthenticationMethodProp.md) |  | 
**HealthCheckConnectTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for a connection to be established for the purpose of performing a health check. If the connection cannot be established within this length of time, the server will be classified as unavailable. | [optional] 
**MaxConnectionAge** | **string** | Specifies the maximum length of time that connections to this server should be allowed to remain established before being closed and replaced with newly-established connections. | 
**MinExpiredConnectionDisconnectInterval** | Pointer to **string** | Specifies the minimum length of time that should pass between connection closures as a result of the connections being established for longer than the maximum connection age. This may help avoid cases in which a large number of connections are closed and re-established in a short period of time because of the maximum connection age. | [optional] 
**ConnectTimeout** | **string** | Specifies the maximum length of time to wait for a connection to be established before giving up and considering the server unavailable. | 
**MaxResponseSize** | **string** | Specifies the maximum response size that should be supported for messages received from the LDAP external server. | 
**KeyManagerProvider** | Pointer to **string** | The key manager provider to use if SSL or StartTLS is to be used for connection-level security. When specifying a value for this property (except when using the Null key manager provider) you must ensure that the external server trusts this server&#39;s public certificate by adding this server&#39;s public certificate to the external server&#39;s trust store. | [optional] 
**TrustManagerProvider** | Pointer to **string** | The trust manager provider to use if SSL or StartTLS is to be used for connection-level security. | [optional] 
**InitialConnections** | Pointer to **int64** | The number of connections to initially establish to the LDAP external server. A value of zero indicates that the number of connections should be dynamically based on the number of available worker threads. This will be ignored when using a thread-local connection pool. | [optional] 
**MaxConnections** | Pointer to **int64** | The maximum number of concurrent connections to maintain for the LDAP external server. A value of zero indicates that the number of connections should be dynamically based on the number of available worker threads. This will be ignored when using a thread-local connection pool. | [optional] 
**DefunctConnectionResultCode** | Pointer to [**[]EnumexternalServerDefunctConnectionResultCodeProp**](EnumexternalServerDefunctConnectionResultCodeProp.md) |  | [optional] 
**AbandonOnTimeout** | Pointer to **bool** | Indicates whether to send an abandon request for an operation for which a response timeout is encountered. A request which has timed out on one server may be retried on another server regardless of whether an abandon request is sent, but if the initial attempt is not abandoned then a long-running operation may unnecessarily continue to consume processing resources on the initial server. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewNokiaProxyServerExternalServerResponse

`func NewNokiaProxyServerExternalServerResponse(id string, schemas []EnumnokiaProxyServerExternalServerSchemaUrn, verifyCredentialsMethod EnumexternalServerVerifyCredentialsMethodProp, serverHostName string, serverPort int64, connectionSecurity EnumexternalServerConnectionSecurityProp, authenticationMethod EnumexternalServerNokiaProxyServerAuthenticationMethodProp, maxConnectionAge string, connectTimeout string, maxResponseSize string, ) *NokiaProxyServerExternalServerResponse`

NewNokiaProxyServerExternalServerResponse instantiates a new NokiaProxyServerExternalServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNokiaProxyServerExternalServerResponseWithDefaults

`func NewNokiaProxyServerExternalServerResponseWithDefaults() *NokiaProxyServerExternalServerResponse`

NewNokiaProxyServerExternalServerResponseWithDefaults instantiates a new NokiaProxyServerExternalServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NokiaProxyServerExternalServerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NokiaProxyServerExternalServerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NokiaProxyServerExternalServerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *NokiaProxyServerExternalServerResponse) GetSchemas() []EnumnokiaProxyServerExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *NokiaProxyServerExternalServerResponse) GetSchemasOk() (*[]EnumnokiaProxyServerExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *NokiaProxyServerExternalServerResponse) SetSchemas(v []EnumnokiaProxyServerExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVerifyCredentialsMethod

`func (o *NokiaProxyServerExternalServerResponse) GetVerifyCredentialsMethod() EnumexternalServerVerifyCredentialsMethodProp`

GetVerifyCredentialsMethod returns the VerifyCredentialsMethod field if non-nil, zero value otherwise.

### GetVerifyCredentialsMethodOk

`func (o *NokiaProxyServerExternalServerResponse) GetVerifyCredentialsMethodOk() (*EnumexternalServerVerifyCredentialsMethodProp, bool)`

GetVerifyCredentialsMethodOk returns a tuple with the VerifyCredentialsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCredentialsMethod

`func (o *NokiaProxyServerExternalServerResponse) SetVerifyCredentialsMethod(v EnumexternalServerVerifyCredentialsMethodProp)`

SetVerifyCredentialsMethod sets VerifyCredentialsMethod field to given value.


### GetUseAdministrativeOperationControl

`func (o *NokiaProxyServerExternalServerResponse) GetUseAdministrativeOperationControl() bool`

GetUseAdministrativeOperationControl returns the UseAdministrativeOperationControl field if non-nil, zero value otherwise.

### GetUseAdministrativeOperationControlOk

`func (o *NokiaProxyServerExternalServerResponse) GetUseAdministrativeOperationControlOk() (*bool, bool)`

GetUseAdministrativeOperationControlOk returns a tuple with the UseAdministrativeOperationControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAdministrativeOperationControl

`func (o *NokiaProxyServerExternalServerResponse) SetUseAdministrativeOperationControl(v bool)`

SetUseAdministrativeOperationControl sets UseAdministrativeOperationControl field to given value.

### HasUseAdministrativeOperationControl

`func (o *NokiaProxyServerExternalServerResponse) HasUseAdministrativeOperationControl() bool`

HasUseAdministrativeOperationControl returns a boolean if a field has been set.

### GetServerHostName

`func (o *NokiaProxyServerExternalServerResponse) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *NokiaProxyServerExternalServerResponse) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *NokiaProxyServerExternalServerResponse) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *NokiaProxyServerExternalServerResponse) GetServerPort() int64`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *NokiaProxyServerExternalServerResponse) GetServerPortOk() (*int64, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *NokiaProxyServerExternalServerResponse) SetServerPort(v int64)`

SetServerPort sets ServerPort field to given value.


### GetLocation

`func (o *NokiaProxyServerExternalServerResponse) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *NokiaProxyServerExternalServerResponse) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *NokiaProxyServerExternalServerResponse) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *NokiaProxyServerExternalServerResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetBindDN

`func (o *NokiaProxyServerExternalServerResponse) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *NokiaProxyServerExternalServerResponse) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *NokiaProxyServerExternalServerResponse) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.

### HasBindDN

`func (o *NokiaProxyServerExternalServerResponse) HasBindDN() bool`

HasBindDN returns a boolean if a field has been set.

### GetPassword

`func (o *NokiaProxyServerExternalServerResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NokiaProxyServerExternalServerResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NokiaProxyServerExternalServerResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NokiaProxyServerExternalServerResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPassphraseProvider

`func (o *NokiaProxyServerExternalServerResponse) GetPassphraseProvider() string`

GetPassphraseProvider returns the PassphraseProvider field if non-nil, zero value otherwise.

### GetPassphraseProviderOk

`func (o *NokiaProxyServerExternalServerResponse) GetPassphraseProviderOk() (*string, bool)`

GetPassphraseProviderOk returns a tuple with the PassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseProvider

`func (o *NokiaProxyServerExternalServerResponse) SetPassphraseProvider(v string)`

SetPassphraseProvider sets PassphraseProvider field to given value.

### HasPassphraseProvider

`func (o *NokiaProxyServerExternalServerResponse) HasPassphraseProvider() bool`

HasPassphraseProvider returns a boolean if a field has been set.

### GetConnectionSecurity

`func (o *NokiaProxyServerExternalServerResponse) GetConnectionSecurity() EnumexternalServerConnectionSecurityProp`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *NokiaProxyServerExternalServerResponse) GetConnectionSecurityOk() (*EnumexternalServerConnectionSecurityProp, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *NokiaProxyServerExternalServerResponse) SetConnectionSecurity(v EnumexternalServerConnectionSecurityProp)`

SetConnectionSecurity sets ConnectionSecurity field to given value.


### GetAuthenticationMethod

`func (o *NokiaProxyServerExternalServerResponse) GetAuthenticationMethod() EnumexternalServerNokiaProxyServerAuthenticationMethodProp`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *NokiaProxyServerExternalServerResponse) GetAuthenticationMethodOk() (*EnumexternalServerNokiaProxyServerAuthenticationMethodProp, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *NokiaProxyServerExternalServerResponse) SetAuthenticationMethod(v EnumexternalServerNokiaProxyServerAuthenticationMethodProp)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.


### GetHealthCheckConnectTimeout

`func (o *NokiaProxyServerExternalServerResponse) GetHealthCheckConnectTimeout() string`

GetHealthCheckConnectTimeout returns the HealthCheckConnectTimeout field if non-nil, zero value otherwise.

### GetHealthCheckConnectTimeoutOk

`func (o *NokiaProxyServerExternalServerResponse) GetHealthCheckConnectTimeoutOk() (*string, bool)`

GetHealthCheckConnectTimeoutOk returns a tuple with the HealthCheckConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckConnectTimeout

`func (o *NokiaProxyServerExternalServerResponse) SetHealthCheckConnectTimeout(v string)`

SetHealthCheckConnectTimeout sets HealthCheckConnectTimeout field to given value.

### HasHealthCheckConnectTimeout

`func (o *NokiaProxyServerExternalServerResponse) HasHealthCheckConnectTimeout() bool`

HasHealthCheckConnectTimeout returns a boolean if a field has been set.

### GetMaxConnectionAge

`func (o *NokiaProxyServerExternalServerResponse) GetMaxConnectionAge() string`

GetMaxConnectionAge returns the MaxConnectionAge field if non-nil, zero value otherwise.

### GetMaxConnectionAgeOk

`func (o *NokiaProxyServerExternalServerResponse) GetMaxConnectionAgeOk() (*string, bool)`

GetMaxConnectionAgeOk returns a tuple with the MaxConnectionAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionAge

`func (o *NokiaProxyServerExternalServerResponse) SetMaxConnectionAge(v string)`

SetMaxConnectionAge sets MaxConnectionAge field to given value.


### GetMinExpiredConnectionDisconnectInterval

`func (o *NokiaProxyServerExternalServerResponse) GetMinExpiredConnectionDisconnectInterval() string`

GetMinExpiredConnectionDisconnectInterval returns the MinExpiredConnectionDisconnectInterval field if non-nil, zero value otherwise.

### GetMinExpiredConnectionDisconnectIntervalOk

`func (o *NokiaProxyServerExternalServerResponse) GetMinExpiredConnectionDisconnectIntervalOk() (*string, bool)`

GetMinExpiredConnectionDisconnectIntervalOk returns a tuple with the MinExpiredConnectionDisconnectInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinExpiredConnectionDisconnectInterval

`func (o *NokiaProxyServerExternalServerResponse) SetMinExpiredConnectionDisconnectInterval(v string)`

SetMinExpiredConnectionDisconnectInterval sets MinExpiredConnectionDisconnectInterval field to given value.

### HasMinExpiredConnectionDisconnectInterval

`func (o *NokiaProxyServerExternalServerResponse) HasMinExpiredConnectionDisconnectInterval() bool`

HasMinExpiredConnectionDisconnectInterval returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *NokiaProxyServerExternalServerResponse) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *NokiaProxyServerExternalServerResponse) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *NokiaProxyServerExternalServerResponse) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.


### GetMaxResponseSize

`func (o *NokiaProxyServerExternalServerResponse) GetMaxResponseSize() string`

GetMaxResponseSize returns the MaxResponseSize field if non-nil, zero value otherwise.

### GetMaxResponseSizeOk

`func (o *NokiaProxyServerExternalServerResponse) GetMaxResponseSizeOk() (*string, bool)`

GetMaxResponseSizeOk returns a tuple with the MaxResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseSize

`func (o *NokiaProxyServerExternalServerResponse) SetMaxResponseSize(v string)`

SetMaxResponseSize sets MaxResponseSize field to given value.


### GetKeyManagerProvider

`func (o *NokiaProxyServerExternalServerResponse) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *NokiaProxyServerExternalServerResponse) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *NokiaProxyServerExternalServerResponse) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *NokiaProxyServerExternalServerResponse) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *NokiaProxyServerExternalServerResponse) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *NokiaProxyServerExternalServerResponse) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *NokiaProxyServerExternalServerResponse) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *NokiaProxyServerExternalServerResponse) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetInitialConnections

`func (o *NokiaProxyServerExternalServerResponse) GetInitialConnections() int64`

GetInitialConnections returns the InitialConnections field if non-nil, zero value otherwise.

### GetInitialConnectionsOk

`func (o *NokiaProxyServerExternalServerResponse) GetInitialConnectionsOk() (*int64, bool)`

GetInitialConnectionsOk returns a tuple with the InitialConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialConnections

`func (o *NokiaProxyServerExternalServerResponse) SetInitialConnections(v int64)`

SetInitialConnections sets InitialConnections field to given value.

### HasInitialConnections

`func (o *NokiaProxyServerExternalServerResponse) HasInitialConnections() bool`

HasInitialConnections returns a boolean if a field has been set.

### GetMaxConnections

`func (o *NokiaProxyServerExternalServerResponse) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *NokiaProxyServerExternalServerResponse) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *NokiaProxyServerExternalServerResponse) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *NokiaProxyServerExternalServerResponse) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetDefunctConnectionResultCode

`func (o *NokiaProxyServerExternalServerResponse) GetDefunctConnectionResultCode() []EnumexternalServerDefunctConnectionResultCodeProp`

GetDefunctConnectionResultCode returns the DefunctConnectionResultCode field if non-nil, zero value otherwise.

### GetDefunctConnectionResultCodeOk

`func (o *NokiaProxyServerExternalServerResponse) GetDefunctConnectionResultCodeOk() (*[]EnumexternalServerDefunctConnectionResultCodeProp, bool)`

GetDefunctConnectionResultCodeOk returns a tuple with the DefunctConnectionResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefunctConnectionResultCode

`func (o *NokiaProxyServerExternalServerResponse) SetDefunctConnectionResultCode(v []EnumexternalServerDefunctConnectionResultCodeProp)`

SetDefunctConnectionResultCode sets DefunctConnectionResultCode field to given value.

### HasDefunctConnectionResultCode

`func (o *NokiaProxyServerExternalServerResponse) HasDefunctConnectionResultCode() bool`

HasDefunctConnectionResultCode returns a boolean if a field has been set.

### GetAbandonOnTimeout

`func (o *NokiaProxyServerExternalServerResponse) GetAbandonOnTimeout() bool`

GetAbandonOnTimeout returns the AbandonOnTimeout field if non-nil, zero value otherwise.

### GetAbandonOnTimeoutOk

`func (o *NokiaProxyServerExternalServerResponse) GetAbandonOnTimeoutOk() (*bool, bool)`

GetAbandonOnTimeoutOk returns a tuple with the AbandonOnTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonOnTimeout

`func (o *NokiaProxyServerExternalServerResponse) SetAbandonOnTimeout(v bool)`

SetAbandonOnTimeout sets AbandonOnTimeout field to given value.

### HasAbandonOnTimeout

`func (o *NokiaProxyServerExternalServerResponse) HasAbandonOnTimeout() bool`

HasAbandonOnTimeout returns a boolean if a field has been set.

### GetDescription

`func (o *NokiaProxyServerExternalServerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NokiaProxyServerExternalServerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NokiaProxyServerExternalServerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NokiaProxyServerExternalServerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *NokiaProxyServerExternalServerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NokiaProxyServerExternalServerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NokiaProxyServerExternalServerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *NokiaProxyServerExternalServerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *NokiaProxyServerExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *NokiaProxyServerExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *NokiaProxyServerExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *NokiaProxyServerExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


