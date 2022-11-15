# NokiaProxyServerExternalServerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumnokiaProxyServerExternalServerSchemaUrn**](EnumnokiaProxyServerExternalServerSchemaUrn.md) |  | 
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

### NewNokiaProxyServerExternalServerShared

`func NewNokiaProxyServerExternalServerShared(schemas []EnumnokiaProxyServerExternalServerSchemaUrn, verifyCredentialsMethod EnumexternalServerVerifyCredentialsMethodProp, serverHostName string, serverPort int32, connectionSecurity EnumexternalServerConnectionSecurityProp, authenticationMethod EnumexternalServerAuthenticationMethodProp, maxConnectionAge string, connectTimeout string, maxResponseSize string, ) *NokiaProxyServerExternalServerShared`

NewNokiaProxyServerExternalServerShared instantiates a new NokiaProxyServerExternalServerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNokiaProxyServerExternalServerSharedWithDefaults

`func NewNokiaProxyServerExternalServerSharedWithDefaults() *NokiaProxyServerExternalServerShared`

NewNokiaProxyServerExternalServerSharedWithDefaults instantiates a new NokiaProxyServerExternalServerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *NokiaProxyServerExternalServerShared) GetSchemas() []EnumnokiaProxyServerExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *NokiaProxyServerExternalServerShared) GetSchemasOk() (*[]EnumnokiaProxyServerExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *NokiaProxyServerExternalServerShared) SetSchemas(v []EnumnokiaProxyServerExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVerifyCredentialsMethod

`func (o *NokiaProxyServerExternalServerShared) GetVerifyCredentialsMethod() EnumexternalServerVerifyCredentialsMethodProp`

GetVerifyCredentialsMethod returns the VerifyCredentialsMethod field if non-nil, zero value otherwise.

### GetVerifyCredentialsMethodOk

`func (o *NokiaProxyServerExternalServerShared) GetVerifyCredentialsMethodOk() (*EnumexternalServerVerifyCredentialsMethodProp, bool)`

GetVerifyCredentialsMethodOk returns a tuple with the VerifyCredentialsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCredentialsMethod

`func (o *NokiaProxyServerExternalServerShared) SetVerifyCredentialsMethod(v EnumexternalServerVerifyCredentialsMethodProp)`

SetVerifyCredentialsMethod sets VerifyCredentialsMethod field to given value.


### GetUseAdministrativeOperationControl

`func (o *NokiaProxyServerExternalServerShared) GetUseAdministrativeOperationControl() bool`

GetUseAdministrativeOperationControl returns the UseAdministrativeOperationControl field if non-nil, zero value otherwise.

### GetUseAdministrativeOperationControlOk

`func (o *NokiaProxyServerExternalServerShared) GetUseAdministrativeOperationControlOk() (*bool, bool)`

GetUseAdministrativeOperationControlOk returns a tuple with the UseAdministrativeOperationControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAdministrativeOperationControl

`func (o *NokiaProxyServerExternalServerShared) SetUseAdministrativeOperationControl(v bool)`

SetUseAdministrativeOperationControl sets UseAdministrativeOperationControl field to given value.

### HasUseAdministrativeOperationControl

`func (o *NokiaProxyServerExternalServerShared) HasUseAdministrativeOperationControl() bool`

HasUseAdministrativeOperationControl returns a boolean if a field has been set.

### GetServerHostName

`func (o *NokiaProxyServerExternalServerShared) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *NokiaProxyServerExternalServerShared) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *NokiaProxyServerExternalServerShared) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *NokiaProxyServerExternalServerShared) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *NokiaProxyServerExternalServerShared) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *NokiaProxyServerExternalServerShared) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.


### GetLocation

`func (o *NokiaProxyServerExternalServerShared) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *NokiaProxyServerExternalServerShared) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *NokiaProxyServerExternalServerShared) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *NokiaProxyServerExternalServerShared) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetBindDN

`func (o *NokiaProxyServerExternalServerShared) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *NokiaProxyServerExternalServerShared) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *NokiaProxyServerExternalServerShared) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.

### HasBindDN

`func (o *NokiaProxyServerExternalServerShared) HasBindDN() bool`

HasBindDN returns a boolean if a field has been set.

### GetPassword

`func (o *NokiaProxyServerExternalServerShared) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NokiaProxyServerExternalServerShared) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NokiaProxyServerExternalServerShared) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NokiaProxyServerExternalServerShared) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPassphraseProvider

`func (o *NokiaProxyServerExternalServerShared) GetPassphraseProvider() string`

GetPassphraseProvider returns the PassphraseProvider field if non-nil, zero value otherwise.

### GetPassphraseProviderOk

`func (o *NokiaProxyServerExternalServerShared) GetPassphraseProviderOk() (*string, bool)`

GetPassphraseProviderOk returns a tuple with the PassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseProvider

`func (o *NokiaProxyServerExternalServerShared) SetPassphraseProvider(v string)`

SetPassphraseProvider sets PassphraseProvider field to given value.

### HasPassphraseProvider

`func (o *NokiaProxyServerExternalServerShared) HasPassphraseProvider() bool`

HasPassphraseProvider returns a boolean if a field has been set.

### GetConnectionSecurity

`func (o *NokiaProxyServerExternalServerShared) GetConnectionSecurity() EnumexternalServerConnectionSecurityProp`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *NokiaProxyServerExternalServerShared) GetConnectionSecurityOk() (*EnumexternalServerConnectionSecurityProp, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *NokiaProxyServerExternalServerShared) SetConnectionSecurity(v EnumexternalServerConnectionSecurityProp)`

SetConnectionSecurity sets ConnectionSecurity field to given value.


### GetAuthenticationMethod

`func (o *NokiaProxyServerExternalServerShared) GetAuthenticationMethod() EnumexternalServerAuthenticationMethodProp`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *NokiaProxyServerExternalServerShared) GetAuthenticationMethodOk() (*EnumexternalServerAuthenticationMethodProp, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *NokiaProxyServerExternalServerShared) SetAuthenticationMethod(v EnumexternalServerAuthenticationMethodProp)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.


### GetHealthCheckConnectTimeout

`func (o *NokiaProxyServerExternalServerShared) GetHealthCheckConnectTimeout() string`

GetHealthCheckConnectTimeout returns the HealthCheckConnectTimeout field if non-nil, zero value otherwise.

### GetHealthCheckConnectTimeoutOk

`func (o *NokiaProxyServerExternalServerShared) GetHealthCheckConnectTimeoutOk() (*string, bool)`

GetHealthCheckConnectTimeoutOk returns a tuple with the HealthCheckConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckConnectTimeout

`func (o *NokiaProxyServerExternalServerShared) SetHealthCheckConnectTimeout(v string)`

SetHealthCheckConnectTimeout sets HealthCheckConnectTimeout field to given value.

### HasHealthCheckConnectTimeout

`func (o *NokiaProxyServerExternalServerShared) HasHealthCheckConnectTimeout() bool`

HasHealthCheckConnectTimeout returns a boolean if a field has been set.

### GetMaxConnectionAge

`func (o *NokiaProxyServerExternalServerShared) GetMaxConnectionAge() string`

GetMaxConnectionAge returns the MaxConnectionAge field if non-nil, zero value otherwise.

### GetMaxConnectionAgeOk

`func (o *NokiaProxyServerExternalServerShared) GetMaxConnectionAgeOk() (*string, bool)`

GetMaxConnectionAgeOk returns a tuple with the MaxConnectionAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionAge

`func (o *NokiaProxyServerExternalServerShared) SetMaxConnectionAge(v string)`

SetMaxConnectionAge sets MaxConnectionAge field to given value.


### GetMinExpiredConnectionDisconnectInterval

`func (o *NokiaProxyServerExternalServerShared) GetMinExpiredConnectionDisconnectInterval() string`

GetMinExpiredConnectionDisconnectInterval returns the MinExpiredConnectionDisconnectInterval field if non-nil, zero value otherwise.

### GetMinExpiredConnectionDisconnectIntervalOk

`func (o *NokiaProxyServerExternalServerShared) GetMinExpiredConnectionDisconnectIntervalOk() (*string, bool)`

GetMinExpiredConnectionDisconnectIntervalOk returns a tuple with the MinExpiredConnectionDisconnectInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinExpiredConnectionDisconnectInterval

`func (o *NokiaProxyServerExternalServerShared) SetMinExpiredConnectionDisconnectInterval(v string)`

SetMinExpiredConnectionDisconnectInterval sets MinExpiredConnectionDisconnectInterval field to given value.

### HasMinExpiredConnectionDisconnectInterval

`func (o *NokiaProxyServerExternalServerShared) HasMinExpiredConnectionDisconnectInterval() bool`

HasMinExpiredConnectionDisconnectInterval returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *NokiaProxyServerExternalServerShared) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *NokiaProxyServerExternalServerShared) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *NokiaProxyServerExternalServerShared) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.


### GetMaxResponseSize

`func (o *NokiaProxyServerExternalServerShared) GetMaxResponseSize() string`

GetMaxResponseSize returns the MaxResponseSize field if non-nil, zero value otherwise.

### GetMaxResponseSizeOk

`func (o *NokiaProxyServerExternalServerShared) GetMaxResponseSizeOk() (*string, bool)`

GetMaxResponseSizeOk returns a tuple with the MaxResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseSize

`func (o *NokiaProxyServerExternalServerShared) SetMaxResponseSize(v string)`

SetMaxResponseSize sets MaxResponseSize field to given value.


### GetKeyManagerProvider

`func (o *NokiaProxyServerExternalServerShared) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *NokiaProxyServerExternalServerShared) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *NokiaProxyServerExternalServerShared) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *NokiaProxyServerExternalServerShared) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *NokiaProxyServerExternalServerShared) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *NokiaProxyServerExternalServerShared) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *NokiaProxyServerExternalServerShared) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *NokiaProxyServerExternalServerShared) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetInitialConnections

`func (o *NokiaProxyServerExternalServerShared) GetInitialConnections() int32`

GetInitialConnections returns the InitialConnections field if non-nil, zero value otherwise.

### GetInitialConnectionsOk

`func (o *NokiaProxyServerExternalServerShared) GetInitialConnectionsOk() (*int32, bool)`

GetInitialConnectionsOk returns a tuple with the InitialConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialConnections

`func (o *NokiaProxyServerExternalServerShared) SetInitialConnections(v int32)`

SetInitialConnections sets InitialConnections field to given value.

### HasInitialConnections

`func (o *NokiaProxyServerExternalServerShared) HasInitialConnections() bool`

HasInitialConnections returns a boolean if a field has been set.

### GetMaxConnections

`func (o *NokiaProxyServerExternalServerShared) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *NokiaProxyServerExternalServerShared) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *NokiaProxyServerExternalServerShared) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *NokiaProxyServerExternalServerShared) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetDefunctConnectionResultCode

`func (o *NokiaProxyServerExternalServerShared) GetDefunctConnectionResultCode() []EnumexternalServerDefunctConnectionResultCodeProp`

GetDefunctConnectionResultCode returns the DefunctConnectionResultCode field if non-nil, zero value otherwise.

### GetDefunctConnectionResultCodeOk

`func (o *NokiaProxyServerExternalServerShared) GetDefunctConnectionResultCodeOk() (*[]EnumexternalServerDefunctConnectionResultCodeProp, bool)`

GetDefunctConnectionResultCodeOk returns a tuple with the DefunctConnectionResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefunctConnectionResultCode

`func (o *NokiaProxyServerExternalServerShared) SetDefunctConnectionResultCode(v []EnumexternalServerDefunctConnectionResultCodeProp)`

SetDefunctConnectionResultCode sets DefunctConnectionResultCode field to given value.

### HasDefunctConnectionResultCode

`func (o *NokiaProxyServerExternalServerShared) HasDefunctConnectionResultCode() bool`

HasDefunctConnectionResultCode returns a boolean if a field has been set.

### GetAbandonOnTimeout

`func (o *NokiaProxyServerExternalServerShared) GetAbandonOnTimeout() bool`

GetAbandonOnTimeout returns the AbandonOnTimeout field if non-nil, zero value otherwise.

### GetAbandonOnTimeoutOk

`func (o *NokiaProxyServerExternalServerShared) GetAbandonOnTimeoutOk() (*bool, bool)`

GetAbandonOnTimeoutOk returns a tuple with the AbandonOnTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonOnTimeout

`func (o *NokiaProxyServerExternalServerShared) SetAbandonOnTimeout(v bool)`

SetAbandonOnTimeout sets AbandonOnTimeout field to given value.

### HasAbandonOnTimeout

`func (o *NokiaProxyServerExternalServerShared) HasAbandonOnTimeout() bool`

HasAbandonOnTimeout returns a boolean if a field has been set.

### GetDescription

`func (o *NokiaProxyServerExternalServerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NokiaProxyServerExternalServerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NokiaProxyServerExternalServerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NokiaProxyServerExternalServerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


