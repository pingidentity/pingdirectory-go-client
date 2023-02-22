# PingIdentityDsExternalServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the External Server | 
**Schemas** | [**[]EnumpingIdentityDsExternalServerSchemaUrn**](EnumpingIdentityDsExternalServerSchemaUrn.md) |  | 
**VerifyCredentialsMethod** | [**EnumexternalServerVerifyCredentialsMethodProp**](EnumexternalServerVerifyCredentialsMethodProp.md) |  | 
**UseAdministrativeOperationControl** | Pointer to **bool** | Indicates whether to include the administrative operation request control in requests sent to this server which are intended for administrative operations (e.g., health checking) rather than requests directly from clients. | [optional] 
**ServerHostName** | **string** | The host name or IP address of the target LDAP server. | 
**ServerPort** | **int32** | The port number on which the server listens for requests. | 
**Location** | Pointer to **string** | Specifies the location for the LDAP External Server. | [optional] 
**BindDN** | Pointer to **string** | The DN to use to bind to the target LDAP server if simple authentication is required. | [optional] 
**Password** | Pointer to **string** | The login password for the specified user. | [optional] 
**PassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the login password for the specified user. | [optional] 
**ConnectionSecurity** | [**EnumexternalServerConnectionSecurityProp**](EnumexternalServerConnectionSecurityProp.md) |  | 
**AuthenticationMethod** | [**EnumexternalServerAuthenticationMethodProp**](EnumexternalServerAuthenticationMethodProp.md) |  | 
**HealthCheckConnectTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for a connection to be established for the purpose of performing a health check. If the connection cannot be established within this length of time, the server will be classified as unavailable. | [optional] 
**MaxConnectionAge** | **string** | Specifies the maximum length of time that connections to this server should be allowed to remain established before being closed and replaced with newly-established connections. | 
**MinExpiredConnectionDisconnectInterval** | Pointer to **string** | Specifies the minimum length of time that should pass between connection closures as a result of the connections being established for longer than the maximum connection age. This may help avoid cases in which a large number of connections are closed and re-established in a short period of time because of the maximum connection age. | [optional] 
**ConnectTimeout** | **string** | Specifies the maximum length of time to wait for a connection to be established before giving up and considering the server unavailable. | 
**MaxResponseSize** | **string** | Specifies the maximum response size that should be supported for messages received from the LDAP external server. | 
**KeyManagerProvider** | Pointer to **string** | The key manager provider to use if SSL or StartTLS is to be used for connection-level security. When specifying a value for this property (except when using the Null key manager provider) you must ensure that the external server trusts this server&#39;s public certificate by adding this server&#39;s public certificate to the external server&#39;s trust store. | [optional] 
**TrustManagerProvider** | Pointer to **string** | The trust manager provider to use if SSL or StartTLS is to be used for connection-level security. | [optional] 
**InitialConnections** | Pointer to **int32** | The number of connections to initially establish to the LDAP external server. A value of zero indicates that the number of connections should be dynamically based on the number of available worker threads. This will be ignored when using a thread-local connection pool. | [optional] 
**MaxConnections** | Pointer to **int32** | The maximum number of concurrent connections to maintain for the LDAP external server. A value of zero indicates that the number of connections should be dynamically based on the number of available worker threads. This will be ignored when using a thread-local connection pool. | [optional] 
**DefunctConnectionResultCode** | Pointer to [**[]EnumexternalServerDefunctConnectionResultCodeProp**](EnumexternalServerDefunctConnectionResultCodeProp.md) |  | [optional] 
**AbandonOnTimeout** | Pointer to **bool** | Indicates whether to send an abandon request for an operation for which a response timeout is encountered. A request which has timed out on one server may be retried on another server regardless of whether an abandon request is sent, but if the initial attempt is not abandoned then a long-running operation may unnecessarily continue to consume processing resources on the initial server. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 

## Methods

### NewPingIdentityDsExternalServerResponse

`func NewPingIdentityDsExternalServerResponse(id string, schemas []EnumpingIdentityDsExternalServerSchemaUrn, verifyCredentialsMethod EnumexternalServerVerifyCredentialsMethodProp, serverHostName string, serverPort int32, connectionSecurity EnumexternalServerConnectionSecurityProp, authenticationMethod EnumexternalServerAuthenticationMethodProp, maxConnectionAge string, connectTimeout string, maxResponseSize string, ) *PingIdentityDsExternalServerResponse`

NewPingIdentityDsExternalServerResponse instantiates a new PingIdentityDsExternalServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingIdentityDsExternalServerResponseWithDefaults

`func NewPingIdentityDsExternalServerResponseWithDefaults() *PingIdentityDsExternalServerResponse`

NewPingIdentityDsExternalServerResponseWithDefaults instantiates a new PingIdentityDsExternalServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *PingIdentityDsExternalServerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PingIdentityDsExternalServerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PingIdentityDsExternalServerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PingIdentityDsExternalServerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *PingIdentityDsExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *PingIdentityDsExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *PingIdentityDsExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *PingIdentityDsExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *PingIdentityDsExternalServerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PingIdentityDsExternalServerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PingIdentityDsExternalServerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *PingIdentityDsExternalServerResponse) GetSchemas() []EnumpingIdentityDsExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PingIdentityDsExternalServerResponse) GetSchemasOk() (*[]EnumpingIdentityDsExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PingIdentityDsExternalServerResponse) SetSchemas(v []EnumpingIdentityDsExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVerifyCredentialsMethod

`func (o *PingIdentityDsExternalServerResponse) GetVerifyCredentialsMethod() EnumexternalServerVerifyCredentialsMethodProp`

GetVerifyCredentialsMethod returns the VerifyCredentialsMethod field if non-nil, zero value otherwise.

### GetVerifyCredentialsMethodOk

`func (o *PingIdentityDsExternalServerResponse) GetVerifyCredentialsMethodOk() (*EnumexternalServerVerifyCredentialsMethodProp, bool)`

GetVerifyCredentialsMethodOk returns a tuple with the VerifyCredentialsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCredentialsMethod

`func (o *PingIdentityDsExternalServerResponse) SetVerifyCredentialsMethod(v EnumexternalServerVerifyCredentialsMethodProp)`

SetVerifyCredentialsMethod sets VerifyCredentialsMethod field to given value.


### GetUseAdministrativeOperationControl

`func (o *PingIdentityDsExternalServerResponse) GetUseAdministrativeOperationControl() bool`

GetUseAdministrativeOperationControl returns the UseAdministrativeOperationControl field if non-nil, zero value otherwise.

### GetUseAdministrativeOperationControlOk

`func (o *PingIdentityDsExternalServerResponse) GetUseAdministrativeOperationControlOk() (*bool, bool)`

GetUseAdministrativeOperationControlOk returns a tuple with the UseAdministrativeOperationControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAdministrativeOperationControl

`func (o *PingIdentityDsExternalServerResponse) SetUseAdministrativeOperationControl(v bool)`

SetUseAdministrativeOperationControl sets UseAdministrativeOperationControl field to given value.

### HasUseAdministrativeOperationControl

`func (o *PingIdentityDsExternalServerResponse) HasUseAdministrativeOperationControl() bool`

HasUseAdministrativeOperationControl returns a boolean if a field has been set.

### GetServerHostName

`func (o *PingIdentityDsExternalServerResponse) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *PingIdentityDsExternalServerResponse) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *PingIdentityDsExternalServerResponse) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *PingIdentityDsExternalServerResponse) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *PingIdentityDsExternalServerResponse) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *PingIdentityDsExternalServerResponse) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.


### GetLocation

`func (o *PingIdentityDsExternalServerResponse) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PingIdentityDsExternalServerResponse) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PingIdentityDsExternalServerResponse) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PingIdentityDsExternalServerResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetBindDN

`func (o *PingIdentityDsExternalServerResponse) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *PingIdentityDsExternalServerResponse) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *PingIdentityDsExternalServerResponse) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.

### HasBindDN

`func (o *PingIdentityDsExternalServerResponse) HasBindDN() bool`

HasBindDN returns a boolean if a field has been set.

### GetPassword

`func (o *PingIdentityDsExternalServerResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PingIdentityDsExternalServerResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PingIdentityDsExternalServerResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PingIdentityDsExternalServerResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPassphraseProvider

`func (o *PingIdentityDsExternalServerResponse) GetPassphraseProvider() string`

GetPassphraseProvider returns the PassphraseProvider field if non-nil, zero value otherwise.

### GetPassphraseProviderOk

`func (o *PingIdentityDsExternalServerResponse) GetPassphraseProviderOk() (*string, bool)`

GetPassphraseProviderOk returns a tuple with the PassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseProvider

`func (o *PingIdentityDsExternalServerResponse) SetPassphraseProvider(v string)`

SetPassphraseProvider sets PassphraseProvider field to given value.

### HasPassphraseProvider

`func (o *PingIdentityDsExternalServerResponse) HasPassphraseProvider() bool`

HasPassphraseProvider returns a boolean if a field has been set.

### GetConnectionSecurity

`func (o *PingIdentityDsExternalServerResponse) GetConnectionSecurity() EnumexternalServerConnectionSecurityProp`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *PingIdentityDsExternalServerResponse) GetConnectionSecurityOk() (*EnumexternalServerConnectionSecurityProp, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *PingIdentityDsExternalServerResponse) SetConnectionSecurity(v EnumexternalServerConnectionSecurityProp)`

SetConnectionSecurity sets ConnectionSecurity field to given value.


### GetAuthenticationMethod

`func (o *PingIdentityDsExternalServerResponse) GetAuthenticationMethod() EnumexternalServerAuthenticationMethodProp`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *PingIdentityDsExternalServerResponse) GetAuthenticationMethodOk() (*EnumexternalServerAuthenticationMethodProp, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *PingIdentityDsExternalServerResponse) SetAuthenticationMethod(v EnumexternalServerAuthenticationMethodProp)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.


### GetHealthCheckConnectTimeout

`func (o *PingIdentityDsExternalServerResponse) GetHealthCheckConnectTimeout() string`

GetHealthCheckConnectTimeout returns the HealthCheckConnectTimeout field if non-nil, zero value otherwise.

### GetHealthCheckConnectTimeoutOk

`func (o *PingIdentityDsExternalServerResponse) GetHealthCheckConnectTimeoutOk() (*string, bool)`

GetHealthCheckConnectTimeoutOk returns a tuple with the HealthCheckConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckConnectTimeout

`func (o *PingIdentityDsExternalServerResponse) SetHealthCheckConnectTimeout(v string)`

SetHealthCheckConnectTimeout sets HealthCheckConnectTimeout field to given value.

### HasHealthCheckConnectTimeout

`func (o *PingIdentityDsExternalServerResponse) HasHealthCheckConnectTimeout() bool`

HasHealthCheckConnectTimeout returns a boolean if a field has been set.

### GetMaxConnectionAge

`func (o *PingIdentityDsExternalServerResponse) GetMaxConnectionAge() string`

GetMaxConnectionAge returns the MaxConnectionAge field if non-nil, zero value otherwise.

### GetMaxConnectionAgeOk

`func (o *PingIdentityDsExternalServerResponse) GetMaxConnectionAgeOk() (*string, bool)`

GetMaxConnectionAgeOk returns a tuple with the MaxConnectionAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionAge

`func (o *PingIdentityDsExternalServerResponse) SetMaxConnectionAge(v string)`

SetMaxConnectionAge sets MaxConnectionAge field to given value.


### GetMinExpiredConnectionDisconnectInterval

`func (o *PingIdentityDsExternalServerResponse) GetMinExpiredConnectionDisconnectInterval() string`

GetMinExpiredConnectionDisconnectInterval returns the MinExpiredConnectionDisconnectInterval field if non-nil, zero value otherwise.

### GetMinExpiredConnectionDisconnectIntervalOk

`func (o *PingIdentityDsExternalServerResponse) GetMinExpiredConnectionDisconnectIntervalOk() (*string, bool)`

GetMinExpiredConnectionDisconnectIntervalOk returns a tuple with the MinExpiredConnectionDisconnectInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinExpiredConnectionDisconnectInterval

`func (o *PingIdentityDsExternalServerResponse) SetMinExpiredConnectionDisconnectInterval(v string)`

SetMinExpiredConnectionDisconnectInterval sets MinExpiredConnectionDisconnectInterval field to given value.

### HasMinExpiredConnectionDisconnectInterval

`func (o *PingIdentityDsExternalServerResponse) HasMinExpiredConnectionDisconnectInterval() bool`

HasMinExpiredConnectionDisconnectInterval returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *PingIdentityDsExternalServerResponse) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *PingIdentityDsExternalServerResponse) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *PingIdentityDsExternalServerResponse) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.


### GetMaxResponseSize

`func (o *PingIdentityDsExternalServerResponse) GetMaxResponseSize() string`

GetMaxResponseSize returns the MaxResponseSize field if non-nil, zero value otherwise.

### GetMaxResponseSizeOk

`func (o *PingIdentityDsExternalServerResponse) GetMaxResponseSizeOk() (*string, bool)`

GetMaxResponseSizeOk returns a tuple with the MaxResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseSize

`func (o *PingIdentityDsExternalServerResponse) SetMaxResponseSize(v string)`

SetMaxResponseSize sets MaxResponseSize field to given value.


### GetKeyManagerProvider

`func (o *PingIdentityDsExternalServerResponse) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *PingIdentityDsExternalServerResponse) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *PingIdentityDsExternalServerResponse) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *PingIdentityDsExternalServerResponse) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *PingIdentityDsExternalServerResponse) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *PingIdentityDsExternalServerResponse) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *PingIdentityDsExternalServerResponse) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *PingIdentityDsExternalServerResponse) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetInitialConnections

`func (o *PingIdentityDsExternalServerResponse) GetInitialConnections() int32`

GetInitialConnections returns the InitialConnections field if non-nil, zero value otherwise.

### GetInitialConnectionsOk

`func (o *PingIdentityDsExternalServerResponse) GetInitialConnectionsOk() (*int32, bool)`

GetInitialConnectionsOk returns a tuple with the InitialConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialConnections

`func (o *PingIdentityDsExternalServerResponse) SetInitialConnections(v int32)`

SetInitialConnections sets InitialConnections field to given value.

### HasInitialConnections

`func (o *PingIdentityDsExternalServerResponse) HasInitialConnections() bool`

HasInitialConnections returns a boolean if a field has been set.

### GetMaxConnections

`func (o *PingIdentityDsExternalServerResponse) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *PingIdentityDsExternalServerResponse) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *PingIdentityDsExternalServerResponse) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *PingIdentityDsExternalServerResponse) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetDefunctConnectionResultCode

`func (o *PingIdentityDsExternalServerResponse) GetDefunctConnectionResultCode() []EnumexternalServerDefunctConnectionResultCodeProp`

GetDefunctConnectionResultCode returns the DefunctConnectionResultCode field if non-nil, zero value otherwise.

### GetDefunctConnectionResultCodeOk

`func (o *PingIdentityDsExternalServerResponse) GetDefunctConnectionResultCodeOk() (*[]EnumexternalServerDefunctConnectionResultCodeProp, bool)`

GetDefunctConnectionResultCodeOk returns a tuple with the DefunctConnectionResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefunctConnectionResultCode

`func (o *PingIdentityDsExternalServerResponse) SetDefunctConnectionResultCode(v []EnumexternalServerDefunctConnectionResultCodeProp)`

SetDefunctConnectionResultCode sets DefunctConnectionResultCode field to given value.

### HasDefunctConnectionResultCode

`func (o *PingIdentityDsExternalServerResponse) HasDefunctConnectionResultCode() bool`

HasDefunctConnectionResultCode returns a boolean if a field has been set.

### GetAbandonOnTimeout

`func (o *PingIdentityDsExternalServerResponse) GetAbandonOnTimeout() bool`

GetAbandonOnTimeout returns the AbandonOnTimeout field if non-nil, zero value otherwise.

### GetAbandonOnTimeoutOk

`func (o *PingIdentityDsExternalServerResponse) GetAbandonOnTimeoutOk() (*bool, bool)`

GetAbandonOnTimeoutOk returns a tuple with the AbandonOnTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonOnTimeout

`func (o *PingIdentityDsExternalServerResponse) SetAbandonOnTimeout(v bool)`

SetAbandonOnTimeout sets AbandonOnTimeout field to given value.

### HasAbandonOnTimeout

`func (o *PingIdentityDsExternalServerResponse) HasAbandonOnTimeout() bool`

HasAbandonOnTimeout returns a boolean if a field has been set.

### GetDescription

`func (o *PingIdentityDsExternalServerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PingIdentityDsExternalServerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PingIdentityDsExternalServerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PingIdentityDsExternalServerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


