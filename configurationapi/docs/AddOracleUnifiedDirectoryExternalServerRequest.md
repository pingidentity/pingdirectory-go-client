# AddOracleUnifiedDirectoryExternalServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerName** | **string** | Name of the new External Server | 
**Schemas** | [**[]EnumoracleUnifiedDirectoryExternalServerSchemaUrn**](EnumoracleUnifiedDirectoryExternalServerSchemaUrn.md) |  | 
**ServerHostName** | **string** | The host name or IP address of the target LDAP server. | 
**ServerPort** | Pointer to **int32** | The port number on which the server listens for requests. | [optional] 
**Location** | Pointer to **string** | Specifies the location for the LDAP External Server. | [optional] 
**BindDN** | Pointer to **string** | The DN to use to bind to the target LDAP server if simple authentication is required. | [optional] 
**Password** | Pointer to **string** | The login password for the specified user. | [optional] 
**PassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the login password for the specified user. | [optional] 
**ConnectionSecurity** | Pointer to [**EnumexternalServerConnectionSecurityProp**](EnumexternalServerConnectionSecurityProp.md) |  | [optional] 
**AuthenticationMethod** | Pointer to [**EnumexternalServerOracleUnifiedDirectoryAuthenticationMethodProp**](EnumexternalServerOracleUnifiedDirectoryAuthenticationMethodProp.md) |  | [optional] 
**VerifyCredentialsMethod** | Pointer to [**EnumexternalServerVerifyCredentialsMethodProp**](EnumexternalServerVerifyCredentialsMethodProp.md) |  | [optional] 
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

### NewAddOracleUnifiedDirectoryExternalServerRequest

`func NewAddOracleUnifiedDirectoryExternalServerRequest(serverName string, schemas []EnumoracleUnifiedDirectoryExternalServerSchemaUrn, serverHostName string, ) *AddOracleUnifiedDirectoryExternalServerRequest`

NewAddOracleUnifiedDirectoryExternalServerRequest instantiates a new AddOracleUnifiedDirectoryExternalServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOracleUnifiedDirectoryExternalServerRequestWithDefaults

`func NewAddOracleUnifiedDirectoryExternalServerRequestWithDefaults() *AddOracleUnifiedDirectoryExternalServerRequest`

NewAddOracleUnifiedDirectoryExternalServerRequestWithDefaults instantiates a new AddOracleUnifiedDirectoryExternalServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerName

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetSchemas

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetSchemas() []EnumoracleUnifiedDirectoryExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetSchemasOk() (*[]EnumoracleUnifiedDirectoryExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetSchemas(v []EnumoracleUnifiedDirectoryExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServerHostName

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetLocation

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetBindDN

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.

### HasBindDN

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasBindDN() bool`

HasBindDN returns a boolean if a field has been set.

### GetPassword

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPassphraseProvider

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetPassphraseProvider() string`

GetPassphraseProvider returns the PassphraseProvider field if non-nil, zero value otherwise.

### GetPassphraseProviderOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetPassphraseProviderOk() (*string, bool)`

GetPassphraseProviderOk returns a tuple with the PassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseProvider

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetPassphraseProvider(v string)`

SetPassphraseProvider sets PassphraseProvider field to given value.

### HasPassphraseProvider

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasPassphraseProvider() bool`

HasPassphraseProvider returns a boolean if a field has been set.

### GetConnectionSecurity

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetConnectionSecurity() EnumexternalServerConnectionSecurityProp`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetConnectionSecurityOk() (*EnumexternalServerConnectionSecurityProp, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetConnectionSecurity(v EnumexternalServerConnectionSecurityProp)`

SetConnectionSecurity sets ConnectionSecurity field to given value.

### HasConnectionSecurity

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasConnectionSecurity() bool`

HasConnectionSecurity returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetAuthenticationMethod() EnumexternalServerOracleUnifiedDirectoryAuthenticationMethodProp`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetAuthenticationMethodOk() (*EnumexternalServerOracleUnifiedDirectoryAuthenticationMethodProp, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetAuthenticationMethod(v EnumexternalServerOracleUnifiedDirectoryAuthenticationMethodProp)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetVerifyCredentialsMethod

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetVerifyCredentialsMethod() EnumexternalServerVerifyCredentialsMethodProp`

GetVerifyCredentialsMethod returns the VerifyCredentialsMethod field if non-nil, zero value otherwise.

### GetVerifyCredentialsMethodOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetVerifyCredentialsMethodOk() (*EnumexternalServerVerifyCredentialsMethodProp, bool)`

GetVerifyCredentialsMethodOk returns a tuple with the VerifyCredentialsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCredentialsMethod

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetVerifyCredentialsMethod(v EnumexternalServerVerifyCredentialsMethodProp)`

SetVerifyCredentialsMethod sets VerifyCredentialsMethod field to given value.

### HasVerifyCredentialsMethod

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasVerifyCredentialsMethod() bool`

HasVerifyCredentialsMethod returns a boolean if a field has been set.

### GetHealthCheckConnectTimeout

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetHealthCheckConnectTimeout() string`

GetHealthCheckConnectTimeout returns the HealthCheckConnectTimeout field if non-nil, zero value otherwise.

### GetHealthCheckConnectTimeoutOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetHealthCheckConnectTimeoutOk() (*string, bool)`

GetHealthCheckConnectTimeoutOk returns a tuple with the HealthCheckConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckConnectTimeout

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetHealthCheckConnectTimeout(v string)`

SetHealthCheckConnectTimeout sets HealthCheckConnectTimeout field to given value.

### HasHealthCheckConnectTimeout

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasHealthCheckConnectTimeout() bool`

HasHealthCheckConnectTimeout returns a boolean if a field has been set.

### GetMaxConnectionAge

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetMaxConnectionAge() string`

GetMaxConnectionAge returns the MaxConnectionAge field if non-nil, zero value otherwise.

### GetMaxConnectionAgeOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetMaxConnectionAgeOk() (*string, bool)`

GetMaxConnectionAgeOk returns a tuple with the MaxConnectionAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionAge

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetMaxConnectionAge(v string)`

SetMaxConnectionAge sets MaxConnectionAge field to given value.

### HasMaxConnectionAge

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasMaxConnectionAge() bool`

HasMaxConnectionAge returns a boolean if a field has been set.

### GetMinExpiredConnectionDisconnectInterval

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetMinExpiredConnectionDisconnectInterval() string`

GetMinExpiredConnectionDisconnectInterval returns the MinExpiredConnectionDisconnectInterval field if non-nil, zero value otherwise.

### GetMinExpiredConnectionDisconnectIntervalOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetMinExpiredConnectionDisconnectIntervalOk() (*string, bool)`

GetMinExpiredConnectionDisconnectIntervalOk returns a tuple with the MinExpiredConnectionDisconnectInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinExpiredConnectionDisconnectInterval

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetMinExpiredConnectionDisconnectInterval(v string)`

SetMinExpiredConnectionDisconnectInterval sets MinExpiredConnectionDisconnectInterval field to given value.

### HasMinExpiredConnectionDisconnectInterval

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasMinExpiredConnectionDisconnectInterval() bool`

HasMinExpiredConnectionDisconnectInterval returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetMaxResponseSize

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetMaxResponseSize() string`

GetMaxResponseSize returns the MaxResponseSize field if non-nil, zero value otherwise.

### GetMaxResponseSizeOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetMaxResponseSizeOk() (*string, bool)`

GetMaxResponseSizeOk returns a tuple with the MaxResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseSize

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetMaxResponseSize(v string)`

SetMaxResponseSize sets MaxResponseSize field to given value.

### HasMaxResponseSize

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasMaxResponseSize() bool`

HasMaxResponseSize returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetInitialConnections

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetInitialConnections() int32`

GetInitialConnections returns the InitialConnections field if non-nil, zero value otherwise.

### GetInitialConnectionsOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetInitialConnectionsOk() (*int32, bool)`

GetInitialConnectionsOk returns a tuple with the InitialConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialConnections

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetInitialConnections(v int32)`

SetInitialConnections sets InitialConnections field to given value.

### HasInitialConnections

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasInitialConnections() bool`

HasInitialConnections returns a boolean if a field has been set.

### GetMaxConnections

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetDefunctConnectionResultCode

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetDefunctConnectionResultCode() []EnumexternalServerDefunctConnectionResultCodeProp`

GetDefunctConnectionResultCode returns the DefunctConnectionResultCode field if non-nil, zero value otherwise.

### GetDefunctConnectionResultCodeOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetDefunctConnectionResultCodeOk() (*[]EnumexternalServerDefunctConnectionResultCodeProp, bool)`

GetDefunctConnectionResultCodeOk returns a tuple with the DefunctConnectionResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefunctConnectionResultCode

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetDefunctConnectionResultCode(v []EnumexternalServerDefunctConnectionResultCodeProp)`

SetDefunctConnectionResultCode sets DefunctConnectionResultCode field to given value.

### HasDefunctConnectionResultCode

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasDefunctConnectionResultCode() bool`

HasDefunctConnectionResultCode returns a boolean if a field has been set.

### GetAbandonOnTimeout

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetAbandonOnTimeout() bool`

GetAbandonOnTimeout returns the AbandonOnTimeout field if non-nil, zero value otherwise.

### GetAbandonOnTimeoutOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetAbandonOnTimeoutOk() (*bool, bool)`

GetAbandonOnTimeoutOk returns a tuple with the AbandonOnTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonOnTimeout

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetAbandonOnTimeout(v bool)`

SetAbandonOnTimeout sets AbandonOnTimeout field to given value.

### HasAbandonOnTimeout

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasAbandonOnTimeout() bool`

HasAbandonOnTimeout returns a boolean if a field has been set.

### GetDescription

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddOracleUnifiedDirectoryExternalServerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


