# AddNokiaProxyServerExternalServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerName** | **string** | Name of the new External Server | 
**Schemas** | [**[]EnumnokiaProxyServerExternalServerSchemaUrn**](EnumnokiaProxyServerExternalServerSchemaUrn.md) |  | 
**VerifyCredentialsMethod** | Pointer to [**EnumexternalServerVerifyCredentialsMethodProp**](EnumexternalServerVerifyCredentialsMethodProp.md) |  | [optional] 
**UseAdministrativeOperationControl** | Pointer to **bool** | Indicates whether to include the administrative operation request control in requests sent to this server which are intended for administrative operations (e.g., health checking) rather than requests directly from clients. | [optional] 
**ServerHostName** | **string** | The host name or IP address of the target LDAP server. | 
**ServerPort** | Pointer to **int32** | The port number on which the server listens for requests. | [optional] 
**Location** | Pointer to **string** | Specifies the location for the LDAP External Server. | [optional] 
**BindDN** | Pointer to **string** | The DN to use to bind to the target LDAP server if simple authentication is required. | [optional] 
**Password** | Pointer to **string** | The login password for the specified user. | [optional] 
**PassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the login password for the specified user. | [optional] 
**ConnectionSecurity** | Pointer to [**EnumexternalServerConnectionSecurityProp**](EnumexternalServerConnectionSecurityProp.md) |  | [optional] 
**AuthenticationMethod** | Pointer to [**EnumexternalServerNokiaProxyServerAuthenticationMethodProp**](EnumexternalServerNokiaProxyServerAuthenticationMethodProp.md) |  | [optional] 
**HealthCheckConnectTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for a connection to be established for the purpose of performing a health check. If the connection cannot be established within this length of time, the server will be classified as unavailable. | [optional] 
**MaxConnectionAge** | Pointer to **string** | Specifies the maximum length of time that connections to this server should be allowed to remain established before being closed and replaced with newly-established connections. | [optional] 
**MinExpiredConnectionDisconnectInterval** | Pointer to **string** | Specifies the minimum length of time that should pass between connection closures as a result of the connections being established for longer than the maximum connection age. This may help avoid cases in which a large number of connections are closed and re-established in a short period of time because of the maximum connection age. | [optional] 
**ConnectTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for a connection to be established before giving up and considering the server unavailable. | [optional] 
**MaxResponseSize** | Pointer to **string** | Specifies the maximum response size that should be supported for messages received from the LDAP external server. | [optional] 
**KeyManagerProvider** | Pointer to **string** | The key manager provider to use if SSL or StartTLS is to be used for connection-level security. When specifying a value for this property (except when using the Null key manager provider) you must ensure that the external server trusts this server&#39;s public certificate by adding this server&#39;s public certificate to the external server&#39;s trust store. | [optional] 
**TrustManagerProvider** | Pointer to **string** | The trust manager provider to use if SSL or StartTLS is to be used for connection-level security. | [optional] 
**InitialConnections** | Pointer to **int32** | The number of connections to initially establish to the LDAP external server. A value of zero indicates that the number of connections should be dynamically based on the number of available worker threads. This will be ignored when using a thread-local connection pool. | [optional] 
**MaxConnections** | Pointer to **int32** | The maximum number of concurrent connections to maintain for the LDAP external server. A value of zero indicates that the number of connections should be dynamically based on the number of available worker threads. This will be ignored when using a thread-local connection pool. | [optional] 
**DefunctConnectionResultCode** | Pointer to [**[]EnumexternalServerDefunctConnectionResultCodeProp**](EnumexternalServerDefunctConnectionResultCodeProp.md) | Specifies the operation result code values that should cause the associated connection should be considered defunct. If an operation fails with one of these result codes, then it will be terminated and an attempt will be made to establish a new connection in its place. | [optional] 
**AbandonOnTimeout** | Pointer to **bool** | Indicates whether to send an abandon request for an operation for which a response timeout is encountered. A request which has timed out on one server may be retried on another server regardless of whether an abandon request is sent, but if the initial attempt is not abandoned then a long-running operation may unnecessarily continue to consume processing resources on the initial server. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 

## Methods

### NewAddNokiaProxyServerExternalServerRequest

`func NewAddNokiaProxyServerExternalServerRequest(serverName string, schemas []EnumnokiaProxyServerExternalServerSchemaUrn, serverHostName string, ) *AddNokiaProxyServerExternalServerRequest`

NewAddNokiaProxyServerExternalServerRequest instantiates a new AddNokiaProxyServerExternalServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddNokiaProxyServerExternalServerRequestWithDefaults

`func NewAddNokiaProxyServerExternalServerRequestWithDefaults() *AddNokiaProxyServerExternalServerRequest`

NewAddNokiaProxyServerExternalServerRequestWithDefaults instantiates a new AddNokiaProxyServerExternalServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerName

`func (o *AddNokiaProxyServerExternalServerRequest) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *AddNokiaProxyServerExternalServerRequest) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetSchemas

`func (o *AddNokiaProxyServerExternalServerRequest) GetSchemas() []EnumnokiaProxyServerExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetSchemasOk() (*[]EnumnokiaProxyServerExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddNokiaProxyServerExternalServerRequest) SetSchemas(v []EnumnokiaProxyServerExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetVerifyCredentialsMethod

`func (o *AddNokiaProxyServerExternalServerRequest) GetVerifyCredentialsMethod() EnumexternalServerVerifyCredentialsMethodProp`

GetVerifyCredentialsMethod returns the VerifyCredentialsMethod field if non-nil, zero value otherwise.

### GetVerifyCredentialsMethodOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetVerifyCredentialsMethodOk() (*EnumexternalServerVerifyCredentialsMethodProp, bool)`

GetVerifyCredentialsMethodOk returns a tuple with the VerifyCredentialsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCredentialsMethod

`func (o *AddNokiaProxyServerExternalServerRequest) SetVerifyCredentialsMethod(v EnumexternalServerVerifyCredentialsMethodProp)`

SetVerifyCredentialsMethod sets VerifyCredentialsMethod field to given value.

### HasVerifyCredentialsMethod

`func (o *AddNokiaProxyServerExternalServerRequest) HasVerifyCredentialsMethod() bool`

HasVerifyCredentialsMethod returns a boolean if a field has been set.

### GetUseAdministrativeOperationControl

`func (o *AddNokiaProxyServerExternalServerRequest) GetUseAdministrativeOperationControl() bool`

GetUseAdministrativeOperationControl returns the UseAdministrativeOperationControl field if non-nil, zero value otherwise.

### GetUseAdministrativeOperationControlOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetUseAdministrativeOperationControlOk() (*bool, bool)`

GetUseAdministrativeOperationControlOk returns a tuple with the UseAdministrativeOperationControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAdministrativeOperationControl

`func (o *AddNokiaProxyServerExternalServerRequest) SetUseAdministrativeOperationControl(v bool)`

SetUseAdministrativeOperationControl sets UseAdministrativeOperationControl field to given value.

### HasUseAdministrativeOperationControl

`func (o *AddNokiaProxyServerExternalServerRequest) HasUseAdministrativeOperationControl() bool`

HasUseAdministrativeOperationControl returns a boolean if a field has been set.

### GetServerHostName

`func (o *AddNokiaProxyServerExternalServerRequest) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *AddNokiaProxyServerExternalServerRequest) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *AddNokiaProxyServerExternalServerRequest) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *AddNokiaProxyServerExternalServerRequest) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *AddNokiaProxyServerExternalServerRequest) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetLocation

`func (o *AddNokiaProxyServerExternalServerRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AddNokiaProxyServerExternalServerRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AddNokiaProxyServerExternalServerRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetBindDN

`func (o *AddNokiaProxyServerExternalServerRequest) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *AddNokiaProxyServerExternalServerRequest) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.

### HasBindDN

`func (o *AddNokiaProxyServerExternalServerRequest) HasBindDN() bool`

HasBindDN returns a boolean if a field has been set.

### GetPassword

`func (o *AddNokiaProxyServerExternalServerRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddNokiaProxyServerExternalServerRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddNokiaProxyServerExternalServerRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPassphraseProvider

`func (o *AddNokiaProxyServerExternalServerRequest) GetPassphraseProvider() string`

GetPassphraseProvider returns the PassphraseProvider field if non-nil, zero value otherwise.

### GetPassphraseProviderOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetPassphraseProviderOk() (*string, bool)`

GetPassphraseProviderOk returns a tuple with the PassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseProvider

`func (o *AddNokiaProxyServerExternalServerRequest) SetPassphraseProvider(v string)`

SetPassphraseProvider sets PassphraseProvider field to given value.

### HasPassphraseProvider

`func (o *AddNokiaProxyServerExternalServerRequest) HasPassphraseProvider() bool`

HasPassphraseProvider returns a boolean if a field has been set.

### GetConnectionSecurity

`func (o *AddNokiaProxyServerExternalServerRequest) GetConnectionSecurity() EnumexternalServerConnectionSecurityProp`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetConnectionSecurityOk() (*EnumexternalServerConnectionSecurityProp, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *AddNokiaProxyServerExternalServerRequest) SetConnectionSecurity(v EnumexternalServerConnectionSecurityProp)`

SetConnectionSecurity sets ConnectionSecurity field to given value.

### HasConnectionSecurity

`func (o *AddNokiaProxyServerExternalServerRequest) HasConnectionSecurity() bool`

HasConnectionSecurity returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *AddNokiaProxyServerExternalServerRequest) GetAuthenticationMethod() EnumexternalServerNokiaProxyServerAuthenticationMethodProp`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetAuthenticationMethodOk() (*EnumexternalServerNokiaProxyServerAuthenticationMethodProp, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *AddNokiaProxyServerExternalServerRequest) SetAuthenticationMethod(v EnumexternalServerNokiaProxyServerAuthenticationMethodProp)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *AddNokiaProxyServerExternalServerRequest) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetHealthCheckConnectTimeout

`func (o *AddNokiaProxyServerExternalServerRequest) GetHealthCheckConnectTimeout() string`

GetHealthCheckConnectTimeout returns the HealthCheckConnectTimeout field if non-nil, zero value otherwise.

### GetHealthCheckConnectTimeoutOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetHealthCheckConnectTimeoutOk() (*string, bool)`

GetHealthCheckConnectTimeoutOk returns a tuple with the HealthCheckConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckConnectTimeout

`func (o *AddNokiaProxyServerExternalServerRequest) SetHealthCheckConnectTimeout(v string)`

SetHealthCheckConnectTimeout sets HealthCheckConnectTimeout field to given value.

### HasHealthCheckConnectTimeout

`func (o *AddNokiaProxyServerExternalServerRequest) HasHealthCheckConnectTimeout() bool`

HasHealthCheckConnectTimeout returns a boolean if a field has been set.

### GetMaxConnectionAge

`func (o *AddNokiaProxyServerExternalServerRequest) GetMaxConnectionAge() string`

GetMaxConnectionAge returns the MaxConnectionAge field if non-nil, zero value otherwise.

### GetMaxConnectionAgeOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetMaxConnectionAgeOk() (*string, bool)`

GetMaxConnectionAgeOk returns a tuple with the MaxConnectionAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionAge

`func (o *AddNokiaProxyServerExternalServerRequest) SetMaxConnectionAge(v string)`

SetMaxConnectionAge sets MaxConnectionAge field to given value.

### HasMaxConnectionAge

`func (o *AddNokiaProxyServerExternalServerRequest) HasMaxConnectionAge() bool`

HasMaxConnectionAge returns a boolean if a field has been set.

### GetMinExpiredConnectionDisconnectInterval

`func (o *AddNokiaProxyServerExternalServerRequest) GetMinExpiredConnectionDisconnectInterval() string`

GetMinExpiredConnectionDisconnectInterval returns the MinExpiredConnectionDisconnectInterval field if non-nil, zero value otherwise.

### GetMinExpiredConnectionDisconnectIntervalOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetMinExpiredConnectionDisconnectIntervalOk() (*string, bool)`

GetMinExpiredConnectionDisconnectIntervalOk returns a tuple with the MinExpiredConnectionDisconnectInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinExpiredConnectionDisconnectInterval

`func (o *AddNokiaProxyServerExternalServerRequest) SetMinExpiredConnectionDisconnectInterval(v string)`

SetMinExpiredConnectionDisconnectInterval sets MinExpiredConnectionDisconnectInterval field to given value.

### HasMinExpiredConnectionDisconnectInterval

`func (o *AddNokiaProxyServerExternalServerRequest) HasMinExpiredConnectionDisconnectInterval() bool`

HasMinExpiredConnectionDisconnectInterval returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *AddNokiaProxyServerExternalServerRequest) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *AddNokiaProxyServerExternalServerRequest) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *AddNokiaProxyServerExternalServerRequest) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetMaxResponseSize

`func (o *AddNokiaProxyServerExternalServerRequest) GetMaxResponseSize() string`

GetMaxResponseSize returns the MaxResponseSize field if non-nil, zero value otherwise.

### GetMaxResponseSizeOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetMaxResponseSizeOk() (*string, bool)`

GetMaxResponseSizeOk returns a tuple with the MaxResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseSize

`func (o *AddNokiaProxyServerExternalServerRequest) SetMaxResponseSize(v string)`

SetMaxResponseSize sets MaxResponseSize field to given value.

### HasMaxResponseSize

`func (o *AddNokiaProxyServerExternalServerRequest) HasMaxResponseSize() bool`

HasMaxResponseSize returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *AddNokiaProxyServerExternalServerRequest) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *AddNokiaProxyServerExternalServerRequest) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *AddNokiaProxyServerExternalServerRequest) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *AddNokiaProxyServerExternalServerRequest) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *AddNokiaProxyServerExternalServerRequest) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *AddNokiaProxyServerExternalServerRequest) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetInitialConnections

`func (o *AddNokiaProxyServerExternalServerRequest) GetInitialConnections() int32`

GetInitialConnections returns the InitialConnections field if non-nil, zero value otherwise.

### GetInitialConnectionsOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetInitialConnectionsOk() (*int32, bool)`

GetInitialConnectionsOk returns a tuple with the InitialConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialConnections

`func (o *AddNokiaProxyServerExternalServerRequest) SetInitialConnections(v int32)`

SetInitialConnections sets InitialConnections field to given value.

### HasInitialConnections

`func (o *AddNokiaProxyServerExternalServerRequest) HasInitialConnections() bool`

HasInitialConnections returns a boolean if a field has been set.

### GetMaxConnections

`func (o *AddNokiaProxyServerExternalServerRequest) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *AddNokiaProxyServerExternalServerRequest) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *AddNokiaProxyServerExternalServerRequest) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetDefunctConnectionResultCode

`func (o *AddNokiaProxyServerExternalServerRequest) GetDefunctConnectionResultCode() []EnumexternalServerDefunctConnectionResultCodeProp`

GetDefunctConnectionResultCode returns the DefunctConnectionResultCode field if non-nil, zero value otherwise.

### GetDefunctConnectionResultCodeOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetDefunctConnectionResultCodeOk() (*[]EnumexternalServerDefunctConnectionResultCodeProp, bool)`

GetDefunctConnectionResultCodeOk returns a tuple with the DefunctConnectionResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefunctConnectionResultCode

`func (o *AddNokiaProxyServerExternalServerRequest) SetDefunctConnectionResultCode(v []EnumexternalServerDefunctConnectionResultCodeProp)`

SetDefunctConnectionResultCode sets DefunctConnectionResultCode field to given value.

### HasDefunctConnectionResultCode

`func (o *AddNokiaProxyServerExternalServerRequest) HasDefunctConnectionResultCode() bool`

HasDefunctConnectionResultCode returns a boolean if a field has been set.

### GetAbandonOnTimeout

`func (o *AddNokiaProxyServerExternalServerRequest) GetAbandonOnTimeout() bool`

GetAbandonOnTimeout returns the AbandonOnTimeout field if non-nil, zero value otherwise.

### GetAbandonOnTimeoutOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetAbandonOnTimeoutOk() (*bool, bool)`

GetAbandonOnTimeoutOk returns a tuple with the AbandonOnTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonOnTimeout

`func (o *AddNokiaProxyServerExternalServerRequest) SetAbandonOnTimeout(v bool)`

SetAbandonOnTimeout sets AbandonOnTimeout field to given value.

### HasAbandonOnTimeout

`func (o *AddNokiaProxyServerExternalServerRequest) HasAbandonOnTimeout() bool`

HasAbandonOnTimeout returns a boolean if a field has been set.

### GetDescription

`func (o *AddNokiaProxyServerExternalServerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddNokiaProxyServerExternalServerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddNokiaProxyServerExternalServerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddNokiaProxyServerExternalServerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


