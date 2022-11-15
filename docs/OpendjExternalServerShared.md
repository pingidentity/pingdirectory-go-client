# OpendjExternalServerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumopendjExternalServerSchemaUrn**](EnumopendjExternalServerSchemaUrn.md) |  | 
**ServerHostName** | **string** | The host name or IP address of the target LDAP server. | 
**ServerPort** | **int32** | The port number on which the server listens for requests. | 
**Location** | Pointer to **string** | Specifies the location for the LDAP External Server. | [optional] 
**BindDN** | Pointer to **string** | The DN to use to bind to the target LDAP server if simple authentication is required. | [optional] 
**Password** | Pointer to **string** | The login password for the specified user. | [optional] 
**PassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the login password for the specified user. | [optional] 
**ConnectionSecurity** | [**EnumexternalServerConnectionSecurityProp**](EnumexternalServerConnectionSecurityProp.md) |  | 
**AuthenticationMethod** | [**EnumexternalServerAuthenticationMethodProp**](EnumexternalServerAuthenticationMethodProp.md) |  | 
**VerifyCredentialsMethod** | [**EnumexternalServerVerifyCredentialsMethodProp**](EnumexternalServerVerifyCredentialsMethodProp.md) |  | 
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

### NewOpendjExternalServerShared

`func NewOpendjExternalServerShared(schemas []EnumopendjExternalServerSchemaUrn, serverHostName string, serverPort int32, connectionSecurity EnumexternalServerConnectionSecurityProp, authenticationMethod EnumexternalServerAuthenticationMethodProp, verifyCredentialsMethod EnumexternalServerVerifyCredentialsMethodProp, maxConnectionAge string, connectTimeout string, maxResponseSize string, ) *OpendjExternalServerShared`

NewOpendjExternalServerShared instantiates a new OpendjExternalServerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpendjExternalServerSharedWithDefaults

`func NewOpendjExternalServerSharedWithDefaults() *OpendjExternalServerShared`

NewOpendjExternalServerSharedWithDefaults instantiates a new OpendjExternalServerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *OpendjExternalServerShared) GetSchemas() []EnumopendjExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *OpendjExternalServerShared) GetSchemasOk() (*[]EnumopendjExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *OpendjExternalServerShared) SetSchemas(v []EnumopendjExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServerHostName

`func (o *OpendjExternalServerShared) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *OpendjExternalServerShared) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *OpendjExternalServerShared) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *OpendjExternalServerShared) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *OpendjExternalServerShared) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *OpendjExternalServerShared) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.


### GetLocation

`func (o *OpendjExternalServerShared) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OpendjExternalServerShared) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OpendjExternalServerShared) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OpendjExternalServerShared) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetBindDN

`func (o *OpendjExternalServerShared) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *OpendjExternalServerShared) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *OpendjExternalServerShared) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.

### HasBindDN

`func (o *OpendjExternalServerShared) HasBindDN() bool`

HasBindDN returns a boolean if a field has been set.

### GetPassword

`func (o *OpendjExternalServerShared) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OpendjExternalServerShared) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OpendjExternalServerShared) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OpendjExternalServerShared) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPassphraseProvider

`func (o *OpendjExternalServerShared) GetPassphraseProvider() string`

GetPassphraseProvider returns the PassphraseProvider field if non-nil, zero value otherwise.

### GetPassphraseProviderOk

`func (o *OpendjExternalServerShared) GetPassphraseProviderOk() (*string, bool)`

GetPassphraseProviderOk returns a tuple with the PassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseProvider

`func (o *OpendjExternalServerShared) SetPassphraseProvider(v string)`

SetPassphraseProvider sets PassphraseProvider field to given value.

### HasPassphraseProvider

`func (o *OpendjExternalServerShared) HasPassphraseProvider() bool`

HasPassphraseProvider returns a boolean if a field has been set.

### GetConnectionSecurity

`func (o *OpendjExternalServerShared) GetConnectionSecurity() EnumexternalServerConnectionSecurityProp`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *OpendjExternalServerShared) GetConnectionSecurityOk() (*EnumexternalServerConnectionSecurityProp, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *OpendjExternalServerShared) SetConnectionSecurity(v EnumexternalServerConnectionSecurityProp)`

SetConnectionSecurity sets ConnectionSecurity field to given value.


### GetAuthenticationMethod

`func (o *OpendjExternalServerShared) GetAuthenticationMethod() EnumexternalServerAuthenticationMethodProp`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *OpendjExternalServerShared) GetAuthenticationMethodOk() (*EnumexternalServerAuthenticationMethodProp, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *OpendjExternalServerShared) SetAuthenticationMethod(v EnumexternalServerAuthenticationMethodProp)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.


### GetVerifyCredentialsMethod

`func (o *OpendjExternalServerShared) GetVerifyCredentialsMethod() EnumexternalServerVerifyCredentialsMethodProp`

GetVerifyCredentialsMethod returns the VerifyCredentialsMethod field if non-nil, zero value otherwise.

### GetVerifyCredentialsMethodOk

`func (o *OpendjExternalServerShared) GetVerifyCredentialsMethodOk() (*EnumexternalServerVerifyCredentialsMethodProp, bool)`

GetVerifyCredentialsMethodOk returns a tuple with the VerifyCredentialsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCredentialsMethod

`func (o *OpendjExternalServerShared) SetVerifyCredentialsMethod(v EnumexternalServerVerifyCredentialsMethodProp)`

SetVerifyCredentialsMethod sets VerifyCredentialsMethod field to given value.


### GetHealthCheckConnectTimeout

`func (o *OpendjExternalServerShared) GetHealthCheckConnectTimeout() string`

GetHealthCheckConnectTimeout returns the HealthCheckConnectTimeout field if non-nil, zero value otherwise.

### GetHealthCheckConnectTimeoutOk

`func (o *OpendjExternalServerShared) GetHealthCheckConnectTimeoutOk() (*string, bool)`

GetHealthCheckConnectTimeoutOk returns a tuple with the HealthCheckConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckConnectTimeout

`func (o *OpendjExternalServerShared) SetHealthCheckConnectTimeout(v string)`

SetHealthCheckConnectTimeout sets HealthCheckConnectTimeout field to given value.

### HasHealthCheckConnectTimeout

`func (o *OpendjExternalServerShared) HasHealthCheckConnectTimeout() bool`

HasHealthCheckConnectTimeout returns a boolean if a field has been set.

### GetMaxConnectionAge

`func (o *OpendjExternalServerShared) GetMaxConnectionAge() string`

GetMaxConnectionAge returns the MaxConnectionAge field if non-nil, zero value otherwise.

### GetMaxConnectionAgeOk

`func (o *OpendjExternalServerShared) GetMaxConnectionAgeOk() (*string, bool)`

GetMaxConnectionAgeOk returns a tuple with the MaxConnectionAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionAge

`func (o *OpendjExternalServerShared) SetMaxConnectionAge(v string)`

SetMaxConnectionAge sets MaxConnectionAge field to given value.


### GetMinExpiredConnectionDisconnectInterval

`func (o *OpendjExternalServerShared) GetMinExpiredConnectionDisconnectInterval() string`

GetMinExpiredConnectionDisconnectInterval returns the MinExpiredConnectionDisconnectInterval field if non-nil, zero value otherwise.

### GetMinExpiredConnectionDisconnectIntervalOk

`func (o *OpendjExternalServerShared) GetMinExpiredConnectionDisconnectIntervalOk() (*string, bool)`

GetMinExpiredConnectionDisconnectIntervalOk returns a tuple with the MinExpiredConnectionDisconnectInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinExpiredConnectionDisconnectInterval

`func (o *OpendjExternalServerShared) SetMinExpiredConnectionDisconnectInterval(v string)`

SetMinExpiredConnectionDisconnectInterval sets MinExpiredConnectionDisconnectInterval field to given value.

### HasMinExpiredConnectionDisconnectInterval

`func (o *OpendjExternalServerShared) HasMinExpiredConnectionDisconnectInterval() bool`

HasMinExpiredConnectionDisconnectInterval returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *OpendjExternalServerShared) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *OpendjExternalServerShared) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *OpendjExternalServerShared) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.


### GetMaxResponseSize

`func (o *OpendjExternalServerShared) GetMaxResponseSize() string`

GetMaxResponseSize returns the MaxResponseSize field if non-nil, zero value otherwise.

### GetMaxResponseSizeOk

`func (o *OpendjExternalServerShared) GetMaxResponseSizeOk() (*string, bool)`

GetMaxResponseSizeOk returns a tuple with the MaxResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseSize

`func (o *OpendjExternalServerShared) SetMaxResponseSize(v string)`

SetMaxResponseSize sets MaxResponseSize field to given value.


### GetKeyManagerProvider

`func (o *OpendjExternalServerShared) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *OpendjExternalServerShared) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *OpendjExternalServerShared) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *OpendjExternalServerShared) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *OpendjExternalServerShared) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *OpendjExternalServerShared) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *OpendjExternalServerShared) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *OpendjExternalServerShared) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetInitialConnections

`func (o *OpendjExternalServerShared) GetInitialConnections() int32`

GetInitialConnections returns the InitialConnections field if non-nil, zero value otherwise.

### GetInitialConnectionsOk

`func (o *OpendjExternalServerShared) GetInitialConnectionsOk() (*int32, bool)`

GetInitialConnectionsOk returns a tuple with the InitialConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialConnections

`func (o *OpendjExternalServerShared) SetInitialConnections(v int32)`

SetInitialConnections sets InitialConnections field to given value.

### HasInitialConnections

`func (o *OpendjExternalServerShared) HasInitialConnections() bool`

HasInitialConnections returns a boolean if a field has been set.

### GetMaxConnections

`func (o *OpendjExternalServerShared) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *OpendjExternalServerShared) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *OpendjExternalServerShared) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *OpendjExternalServerShared) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetDefunctConnectionResultCode

`func (o *OpendjExternalServerShared) GetDefunctConnectionResultCode() []EnumexternalServerDefunctConnectionResultCodeProp`

GetDefunctConnectionResultCode returns the DefunctConnectionResultCode field if non-nil, zero value otherwise.

### GetDefunctConnectionResultCodeOk

`func (o *OpendjExternalServerShared) GetDefunctConnectionResultCodeOk() (*[]EnumexternalServerDefunctConnectionResultCodeProp, bool)`

GetDefunctConnectionResultCodeOk returns a tuple with the DefunctConnectionResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefunctConnectionResultCode

`func (o *OpendjExternalServerShared) SetDefunctConnectionResultCode(v []EnumexternalServerDefunctConnectionResultCodeProp)`

SetDefunctConnectionResultCode sets DefunctConnectionResultCode field to given value.

### HasDefunctConnectionResultCode

`func (o *OpendjExternalServerShared) HasDefunctConnectionResultCode() bool`

HasDefunctConnectionResultCode returns a boolean if a field has been set.

### GetAbandonOnTimeout

`func (o *OpendjExternalServerShared) GetAbandonOnTimeout() bool`

GetAbandonOnTimeout returns the AbandonOnTimeout field if non-nil, zero value otherwise.

### GetAbandonOnTimeoutOk

`func (o *OpendjExternalServerShared) GetAbandonOnTimeoutOk() (*bool, bool)`

GetAbandonOnTimeoutOk returns a tuple with the AbandonOnTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonOnTimeout

`func (o *OpendjExternalServerShared) SetAbandonOnTimeout(v bool)`

SetAbandonOnTimeout sets AbandonOnTimeout field to given value.

### HasAbandonOnTimeout

`func (o *OpendjExternalServerShared) HasAbandonOnTimeout() bool`

HasAbandonOnTimeout returns a boolean if a field has been set.

### GetDescription

`func (o *OpendjExternalServerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpendjExternalServerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpendjExternalServerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpendjExternalServerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


