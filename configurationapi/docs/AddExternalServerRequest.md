# AddExternalServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerName** | **string** | Name of the new External Server | 
**Schemas** | [**[]EnumvaultExternalServerSchemaUrn**](EnumvaultExternalServerSchemaUrn.md) |  | 
**ServerHostName** | **string** | The host name or IP address of the target LDAP server. | 
**ServerPort** | **int64** | The port number on which the server listens for requests. | 
**SmtpSecurity** | Pointer to [**EnumexternalServerSmtpSecurityProp**](EnumexternalServerSmtpSecurityProp.md) |  | [optional] 
**UserName** | Pointer to **string** | The name of the login account to use when connecting to the database server. | [optional] 
**Password** | Pointer to **string** | The login password for the specified user. | [optional] 
**PassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the login password for the specified user. | [optional] 
**SmtpTimeout** | Pointer to **string** | Specifies the maximum length of time that a connection or attempted connection to a SMTP server may take. | [optional] 
**SmtpConnectionProperties** | Pointer to **[]string** | Specifies the connection properties for the smtp server. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 
**VerifyCredentialsMethod** | Pointer to [**EnumexternalServerVerifyCredentialsMethodProp**](EnumexternalServerVerifyCredentialsMethodProp.md) |  | [optional] 
**UseAdministrativeOperationControl** | Pointer to **bool** | Indicates whether to include the administrative operation request control in requests sent to this server which are intended for administrative operations (e.g., health checking) rather than requests directly from clients. | [optional] 
**Location** | Pointer to **string** | Specifies the location for the LDAP External Server. | [optional] 
**BindDN** | Pointer to **string** | The DN to use to bind to the target LDAP server if simple authentication is required. | [optional] 
**ConnectionSecurity** | Pointer to [**EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp**](EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp.md) |  | [optional] 
**AuthenticationMethod** | Pointer to [**EnumexternalServerAmazonAwsAuthenticationMethodProp**](EnumexternalServerAmazonAwsAuthenticationMethodProp.md) |  | [optional] 
**HealthCheckConnectTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for a connection to be established for the purpose of performing a health check. If the connection cannot be established within this length of time, the server will be classified as unavailable. | [optional] 
**MaxConnectionAge** | Pointer to **string** | Specifies the maximum length of time that connections to this server should be allowed to remain established before being closed and replaced with newly-established connections. | [optional] 
**MinExpiredConnectionDisconnectInterval** | Pointer to **string** | Specifies the minimum length of time that should pass between connection closures as a result of the connections being established for longer than the maximum connection age. This may help avoid cases in which a large number of connections are closed and re-established in a short period of time because of the maximum connection age. | [optional] 
**ConnectTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for a connection to be established before giving up and considering the server unavailable. | [optional] 
**MaxResponseSize** | Pointer to **string** | Specifies the maximum response size that should be supported for messages received from the LDAP external server. | [optional] 
**KeyManagerProvider** | Pointer to **string** | The key manager provider to use if SSL or StartTLS is to be used for connection-level security. When specifying a value for this property (except when using the Null key manager provider) you must ensure that the external server trusts this server&#39;s public certificate by adding this server&#39;s public certificate to the external server&#39;s trust store. | [optional] 
**TrustManagerProvider** | Pointer to **string** | The trust manager provider to use if SSL or StartTLS is to be used for connection-level security. | [optional] 
**InitialConnections** | Pointer to **int64** | The number of connections to initially establish to the LDAP external server. A value of zero indicates that the number of connections should be dynamically based on the number of available worker threads. This will be ignored when using a thread-local connection pool. | [optional] 
**MaxConnections** | Pointer to **int64** | The maximum number of concurrent connections to maintain for the LDAP external server. A value of zero indicates that the number of connections should be dynamically based on the number of available worker threads. This will be ignored when using a thread-local connection pool. | [optional] 
**DefunctConnectionResultCode** | Pointer to [**[]EnumexternalServerDefunctConnectionResultCodeProp**](EnumexternalServerDefunctConnectionResultCodeProp.md) |  | [optional] 
**AbandonOnTimeout** | Pointer to **bool** | Indicates whether to send an abandon request for an operation for which a response timeout is encountered. A request which has timed out on one server may be retried on another server regardless of whether an abandon request is sent, but if the initial attempt is not abandoned then a long-running operation may unnecessarily continue to consume processing resources on the initial server. | [optional] 
**JdbcDriverType** | [**EnumexternalServerJdbcDriverTypeProp**](EnumexternalServerJdbcDriverTypeProp.md) |  | 
**JdbcDriverURL** | Pointer to **string** | Specify the complete JDBC URL which will be used instead of the automatic URL format. You must select type &#39;other&#39; for the jdbc-driver-type. | [optional] 
**DatabaseName** | Pointer to **string** | Specifies which database to connect to. This is ignored if jdbc-driver-url is specified. | [optional] 
**ValidationQuery** | Pointer to **string** | The SQL query that will be used to validate connections to the database before making them available to the Directory Server. | [optional] 
**ValidationQueryTimeout** | Pointer to **string** | Specifies the amount of time to wait for a response from the database when executing the validation query, if one is set. If the timeout is exceeded, the Directory Server will drop the connection and obtain a new one. A value of zero indicates no timeout. | [optional] 
**JdbcConnectionProperties** | Pointer to **[]string** | Specifies the connection properties for the JDBC datasource. | [optional] 
**TransactionIsolationLevel** | Pointer to [**EnumexternalServerTransactionIsolationLevelProp**](EnumexternalServerTransactionIsolationLevelProp.md) |  | [optional] 
**TransportMechanism** | [**EnumexternalServerTransportMechanismProp**](EnumexternalServerTransportMechanismProp.md) |  | 
**BasicAuthenticationUsername** | Pointer to **string** | The username to use to authenticate to the HTTP Proxy External Server. | [optional] 
**BasicAuthenticationPassphraseProvider** | Pointer to **string** | A passphrase provider that provides access to the password to use to authenticate to the HTTP Proxy External Server. | [optional] 
**HostnameVerificationMethod** | Pointer to [**EnumexternalServerHttpHostnameVerificationMethodProp**](EnumexternalServerHttpHostnameVerificationMethodProp.md) |  | [optional] 
**ResponseTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for response data to be read from an established connection before aborting a request to the server. | [optional] 
**BaseURL** | **string** | The base URL of the external server, optionally including port number, for example \&quot;https://externalService:9031\&quot;. | 
**SslCertNickname** | Pointer to **string** | The certificate alias within the keystore to use if SSL (HTTPS) is to be used for connection-level security. When specifying a value for this property you must ensure that the external server trusts this server&#39;s public certificate by adding this server&#39;s public certificate to the external server&#39;s trust store. | [optional] 
**ConjurServerBaseURI** | **[]string** | The base URL needed to access the CyberArk Conjur server. The base URL should consist of the protocol (\&quot;http\&quot; or \&quot;https\&quot;), the server address (resolvable name or IP address), and the port number. For example, \&quot;https://conjur.example.com:8443/\&quot;. | 
**ConjurAuthenticationMethod** | **string** | The mechanism used to authenticate to the Conjur server. | 
**ConjurAccountName** | **string** | The name of the account with which the desired secrets are associated. | 
**HttpConnectTimeout** | Pointer to **string** | The maximum length of time to wait to obtain an HTTP connection. | [optional] 
**HttpResponseTimeout** | Pointer to **string** | The maximum length of time to wait for a response to an HTTP request. | [optional] 
**TrustStoreFile** | Pointer to **string** | The path to a file containing the information needed to trust the certificate presented by the Vault servers. | [optional] 
**TrustStorePin** | Pointer to **string** | The passphrase needed to access the contents of the trust store. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents. | [optional] 
**TrustStoreType** | Pointer to **string** | The store type for the specified trust store file. The value should likely be one of \&quot;JKS\&quot;, \&quot;PKCS12\&quot;, or \&quot;BCFKS\&quot;. | [optional] 
**HttpProxyExternalServer** | Pointer to **string** | A reference to an HTTP proxy server that should be used for requests sent to the AWS service. | [optional] 
**AwsAccessKeyID** | Pointer to **string** | The access key ID that will be used if authentication should use an access key. If this is provided, then an aws-secret-access-key must also be provided. | [optional] 
**AwsSecretAccessKey** | Pointer to **string** | The secret access key that will be used if authentication should use an access key. If this is provided, then an aws-access-key-id must also be provided. | [optional] 
**AwsRegionName** | **string** | The name of the AWS region containing the resources that will be accessed. | 
**VaultServerBaseURI** | **[]string** | The base URL needed to access the Vault server. The base URL should consist of the protocol (\&quot;http\&quot; or \&quot;https\&quot;), the server address (resolvable name or IP address), and the port number. For example, \&quot;https://vault.example.com:8200/\&quot;. | 
**VaultAuthenticationMethod** | **string** | The mechanism used to authenticate to the Vault server. | 

## Methods

### NewAddExternalServerRequest

`func NewAddExternalServerRequest(serverName string, schemas []EnumvaultExternalServerSchemaUrn, serverHostName string, serverPort int64, jdbcDriverType EnumexternalServerJdbcDriverTypeProp, transportMechanism EnumexternalServerTransportMechanismProp, baseURL string, conjurServerBaseURI []string, conjurAuthenticationMethod string, conjurAccountName string, awsRegionName string, vaultServerBaseURI []string, vaultAuthenticationMethod string, ) *AddExternalServerRequest`

NewAddExternalServerRequest instantiates a new AddExternalServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddExternalServerRequestWithDefaults

`func NewAddExternalServerRequestWithDefaults() *AddExternalServerRequest`

NewAddExternalServerRequestWithDefaults instantiates a new AddExternalServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerName

`func (o *AddExternalServerRequest) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *AddExternalServerRequest) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *AddExternalServerRequest) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetSchemas

`func (o *AddExternalServerRequest) GetSchemas() []EnumvaultExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddExternalServerRequest) GetSchemasOk() (*[]EnumvaultExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddExternalServerRequest) SetSchemas(v []EnumvaultExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServerHostName

`func (o *AddExternalServerRequest) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *AddExternalServerRequest) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *AddExternalServerRequest) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *AddExternalServerRequest) GetServerPort() int64`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *AddExternalServerRequest) GetServerPortOk() (*int64, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *AddExternalServerRequest) SetServerPort(v int64)`

SetServerPort sets ServerPort field to given value.


### GetSmtpSecurity

`func (o *AddExternalServerRequest) GetSmtpSecurity() EnumexternalServerSmtpSecurityProp`

GetSmtpSecurity returns the SmtpSecurity field if non-nil, zero value otherwise.

### GetSmtpSecurityOk

`func (o *AddExternalServerRequest) GetSmtpSecurityOk() (*EnumexternalServerSmtpSecurityProp, bool)`

GetSmtpSecurityOk returns a tuple with the SmtpSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpSecurity

`func (o *AddExternalServerRequest) SetSmtpSecurity(v EnumexternalServerSmtpSecurityProp)`

SetSmtpSecurity sets SmtpSecurity field to given value.

### HasSmtpSecurity

`func (o *AddExternalServerRequest) HasSmtpSecurity() bool`

HasSmtpSecurity returns a boolean if a field has been set.

### GetUserName

`func (o *AddExternalServerRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AddExternalServerRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AddExternalServerRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AddExternalServerRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *AddExternalServerRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddExternalServerRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddExternalServerRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddExternalServerRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPassphraseProvider

`func (o *AddExternalServerRequest) GetPassphraseProvider() string`

GetPassphraseProvider returns the PassphraseProvider field if non-nil, zero value otherwise.

### GetPassphraseProviderOk

`func (o *AddExternalServerRequest) GetPassphraseProviderOk() (*string, bool)`

GetPassphraseProviderOk returns a tuple with the PassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseProvider

`func (o *AddExternalServerRequest) SetPassphraseProvider(v string)`

SetPassphraseProvider sets PassphraseProvider field to given value.

### HasPassphraseProvider

`func (o *AddExternalServerRequest) HasPassphraseProvider() bool`

HasPassphraseProvider returns a boolean if a field has been set.

### GetSmtpTimeout

`func (o *AddExternalServerRequest) GetSmtpTimeout() string`

GetSmtpTimeout returns the SmtpTimeout field if non-nil, zero value otherwise.

### GetSmtpTimeoutOk

`func (o *AddExternalServerRequest) GetSmtpTimeoutOk() (*string, bool)`

GetSmtpTimeoutOk returns a tuple with the SmtpTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpTimeout

`func (o *AddExternalServerRequest) SetSmtpTimeout(v string)`

SetSmtpTimeout sets SmtpTimeout field to given value.

### HasSmtpTimeout

`func (o *AddExternalServerRequest) HasSmtpTimeout() bool`

HasSmtpTimeout returns a boolean if a field has been set.

### GetSmtpConnectionProperties

`func (o *AddExternalServerRequest) GetSmtpConnectionProperties() []string`

GetSmtpConnectionProperties returns the SmtpConnectionProperties field if non-nil, zero value otherwise.

### GetSmtpConnectionPropertiesOk

`func (o *AddExternalServerRequest) GetSmtpConnectionPropertiesOk() (*[]string, bool)`

GetSmtpConnectionPropertiesOk returns a tuple with the SmtpConnectionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpConnectionProperties

`func (o *AddExternalServerRequest) SetSmtpConnectionProperties(v []string)`

SetSmtpConnectionProperties sets SmtpConnectionProperties field to given value.

### HasSmtpConnectionProperties

`func (o *AddExternalServerRequest) HasSmtpConnectionProperties() bool`

HasSmtpConnectionProperties returns a boolean if a field has been set.

### GetDescription

`func (o *AddExternalServerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddExternalServerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddExternalServerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddExternalServerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVerifyCredentialsMethod

`func (o *AddExternalServerRequest) GetVerifyCredentialsMethod() EnumexternalServerVerifyCredentialsMethodProp`

GetVerifyCredentialsMethod returns the VerifyCredentialsMethod field if non-nil, zero value otherwise.

### GetVerifyCredentialsMethodOk

`func (o *AddExternalServerRequest) GetVerifyCredentialsMethodOk() (*EnumexternalServerVerifyCredentialsMethodProp, bool)`

GetVerifyCredentialsMethodOk returns a tuple with the VerifyCredentialsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCredentialsMethod

`func (o *AddExternalServerRequest) SetVerifyCredentialsMethod(v EnumexternalServerVerifyCredentialsMethodProp)`

SetVerifyCredentialsMethod sets VerifyCredentialsMethod field to given value.

### HasVerifyCredentialsMethod

`func (o *AddExternalServerRequest) HasVerifyCredentialsMethod() bool`

HasVerifyCredentialsMethod returns a boolean if a field has been set.

### GetUseAdministrativeOperationControl

`func (o *AddExternalServerRequest) GetUseAdministrativeOperationControl() bool`

GetUseAdministrativeOperationControl returns the UseAdministrativeOperationControl field if non-nil, zero value otherwise.

### GetUseAdministrativeOperationControlOk

`func (o *AddExternalServerRequest) GetUseAdministrativeOperationControlOk() (*bool, bool)`

GetUseAdministrativeOperationControlOk returns a tuple with the UseAdministrativeOperationControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAdministrativeOperationControl

`func (o *AddExternalServerRequest) SetUseAdministrativeOperationControl(v bool)`

SetUseAdministrativeOperationControl sets UseAdministrativeOperationControl field to given value.

### HasUseAdministrativeOperationControl

`func (o *AddExternalServerRequest) HasUseAdministrativeOperationControl() bool`

HasUseAdministrativeOperationControl returns a boolean if a field has been set.

### GetLocation

`func (o *AddExternalServerRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AddExternalServerRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AddExternalServerRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AddExternalServerRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetBindDN

`func (o *AddExternalServerRequest) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *AddExternalServerRequest) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *AddExternalServerRequest) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.

### HasBindDN

`func (o *AddExternalServerRequest) HasBindDN() bool`

HasBindDN returns a boolean if a field has been set.

### GetConnectionSecurity

`func (o *AddExternalServerRequest) GetConnectionSecurity() EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *AddExternalServerRequest) GetConnectionSecurityOk() (*EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *AddExternalServerRequest) SetConnectionSecurity(v EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp)`

SetConnectionSecurity sets ConnectionSecurity field to given value.

### HasConnectionSecurity

`func (o *AddExternalServerRequest) HasConnectionSecurity() bool`

HasConnectionSecurity returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *AddExternalServerRequest) GetAuthenticationMethod() EnumexternalServerAmazonAwsAuthenticationMethodProp`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *AddExternalServerRequest) GetAuthenticationMethodOk() (*EnumexternalServerAmazonAwsAuthenticationMethodProp, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *AddExternalServerRequest) SetAuthenticationMethod(v EnumexternalServerAmazonAwsAuthenticationMethodProp)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *AddExternalServerRequest) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetHealthCheckConnectTimeout

`func (o *AddExternalServerRequest) GetHealthCheckConnectTimeout() string`

GetHealthCheckConnectTimeout returns the HealthCheckConnectTimeout field if non-nil, zero value otherwise.

### GetHealthCheckConnectTimeoutOk

`func (o *AddExternalServerRequest) GetHealthCheckConnectTimeoutOk() (*string, bool)`

GetHealthCheckConnectTimeoutOk returns a tuple with the HealthCheckConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckConnectTimeout

`func (o *AddExternalServerRequest) SetHealthCheckConnectTimeout(v string)`

SetHealthCheckConnectTimeout sets HealthCheckConnectTimeout field to given value.

### HasHealthCheckConnectTimeout

`func (o *AddExternalServerRequest) HasHealthCheckConnectTimeout() bool`

HasHealthCheckConnectTimeout returns a boolean if a field has been set.

### GetMaxConnectionAge

`func (o *AddExternalServerRequest) GetMaxConnectionAge() string`

GetMaxConnectionAge returns the MaxConnectionAge field if non-nil, zero value otherwise.

### GetMaxConnectionAgeOk

`func (o *AddExternalServerRequest) GetMaxConnectionAgeOk() (*string, bool)`

GetMaxConnectionAgeOk returns a tuple with the MaxConnectionAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionAge

`func (o *AddExternalServerRequest) SetMaxConnectionAge(v string)`

SetMaxConnectionAge sets MaxConnectionAge field to given value.

### HasMaxConnectionAge

`func (o *AddExternalServerRequest) HasMaxConnectionAge() bool`

HasMaxConnectionAge returns a boolean if a field has been set.

### GetMinExpiredConnectionDisconnectInterval

`func (o *AddExternalServerRequest) GetMinExpiredConnectionDisconnectInterval() string`

GetMinExpiredConnectionDisconnectInterval returns the MinExpiredConnectionDisconnectInterval field if non-nil, zero value otherwise.

### GetMinExpiredConnectionDisconnectIntervalOk

`func (o *AddExternalServerRequest) GetMinExpiredConnectionDisconnectIntervalOk() (*string, bool)`

GetMinExpiredConnectionDisconnectIntervalOk returns a tuple with the MinExpiredConnectionDisconnectInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinExpiredConnectionDisconnectInterval

`func (o *AddExternalServerRequest) SetMinExpiredConnectionDisconnectInterval(v string)`

SetMinExpiredConnectionDisconnectInterval sets MinExpiredConnectionDisconnectInterval field to given value.

### HasMinExpiredConnectionDisconnectInterval

`func (o *AddExternalServerRequest) HasMinExpiredConnectionDisconnectInterval() bool`

HasMinExpiredConnectionDisconnectInterval returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *AddExternalServerRequest) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *AddExternalServerRequest) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *AddExternalServerRequest) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *AddExternalServerRequest) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetMaxResponseSize

`func (o *AddExternalServerRequest) GetMaxResponseSize() string`

GetMaxResponseSize returns the MaxResponseSize field if non-nil, zero value otherwise.

### GetMaxResponseSizeOk

`func (o *AddExternalServerRequest) GetMaxResponseSizeOk() (*string, bool)`

GetMaxResponseSizeOk returns a tuple with the MaxResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseSize

`func (o *AddExternalServerRequest) SetMaxResponseSize(v string)`

SetMaxResponseSize sets MaxResponseSize field to given value.

### HasMaxResponseSize

`func (o *AddExternalServerRequest) HasMaxResponseSize() bool`

HasMaxResponseSize returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *AddExternalServerRequest) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *AddExternalServerRequest) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *AddExternalServerRequest) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *AddExternalServerRequest) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *AddExternalServerRequest) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *AddExternalServerRequest) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *AddExternalServerRequest) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *AddExternalServerRequest) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetInitialConnections

`func (o *AddExternalServerRequest) GetInitialConnections() int64`

GetInitialConnections returns the InitialConnections field if non-nil, zero value otherwise.

### GetInitialConnectionsOk

`func (o *AddExternalServerRequest) GetInitialConnectionsOk() (*int64, bool)`

GetInitialConnectionsOk returns a tuple with the InitialConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialConnections

`func (o *AddExternalServerRequest) SetInitialConnections(v int64)`

SetInitialConnections sets InitialConnections field to given value.

### HasInitialConnections

`func (o *AddExternalServerRequest) HasInitialConnections() bool`

HasInitialConnections returns a boolean if a field has been set.

### GetMaxConnections

`func (o *AddExternalServerRequest) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *AddExternalServerRequest) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *AddExternalServerRequest) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *AddExternalServerRequest) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetDefunctConnectionResultCode

`func (o *AddExternalServerRequest) GetDefunctConnectionResultCode() []EnumexternalServerDefunctConnectionResultCodeProp`

GetDefunctConnectionResultCode returns the DefunctConnectionResultCode field if non-nil, zero value otherwise.

### GetDefunctConnectionResultCodeOk

`func (o *AddExternalServerRequest) GetDefunctConnectionResultCodeOk() (*[]EnumexternalServerDefunctConnectionResultCodeProp, bool)`

GetDefunctConnectionResultCodeOk returns a tuple with the DefunctConnectionResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefunctConnectionResultCode

`func (o *AddExternalServerRequest) SetDefunctConnectionResultCode(v []EnumexternalServerDefunctConnectionResultCodeProp)`

SetDefunctConnectionResultCode sets DefunctConnectionResultCode field to given value.

### HasDefunctConnectionResultCode

`func (o *AddExternalServerRequest) HasDefunctConnectionResultCode() bool`

HasDefunctConnectionResultCode returns a boolean if a field has been set.

### GetAbandonOnTimeout

`func (o *AddExternalServerRequest) GetAbandonOnTimeout() bool`

GetAbandonOnTimeout returns the AbandonOnTimeout field if non-nil, zero value otherwise.

### GetAbandonOnTimeoutOk

`func (o *AddExternalServerRequest) GetAbandonOnTimeoutOk() (*bool, bool)`

GetAbandonOnTimeoutOk returns a tuple with the AbandonOnTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonOnTimeout

`func (o *AddExternalServerRequest) SetAbandonOnTimeout(v bool)`

SetAbandonOnTimeout sets AbandonOnTimeout field to given value.

### HasAbandonOnTimeout

`func (o *AddExternalServerRequest) HasAbandonOnTimeout() bool`

HasAbandonOnTimeout returns a boolean if a field has been set.

### GetJdbcDriverType

`func (o *AddExternalServerRequest) GetJdbcDriverType() EnumexternalServerJdbcDriverTypeProp`

GetJdbcDriverType returns the JdbcDriverType field if non-nil, zero value otherwise.

### GetJdbcDriverTypeOk

`func (o *AddExternalServerRequest) GetJdbcDriverTypeOk() (*EnumexternalServerJdbcDriverTypeProp, bool)`

GetJdbcDriverTypeOk returns a tuple with the JdbcDriverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdbcDriverType

`func (o *AddExternalServerRequest) SetJdbcDriverType(v EnumexternalServerJdbcDriverTypeProp)`

SetJdbcDriverType sets JdbcDriverType field to given value.


### GetJdbcDriverURL

`func (o *AddExternalServerRequest) GetJdbcDriverURL() string`

GetJdbcDriverURL returns the JdbcDriverURL field if non-nil, zero value otherwise.

### GetJdbcDriverURLOk

`func (o *AddExternalServerRequest) GetJdbcDriverURLOk() (*string, bool)`

GetJdbcDriverURLOk returns a tuple with the JdbcDriverURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdbcDriverURL

`func (o *AddExternalServerRequest) SetJdbcDriverURL(v string)`

SetJdbcDriverURL sets JdbcDriverURL field to given value.

### HasJdbcDriverURL

`func (o *AddExternalServerRequest) HasJdbcDriverURL() bool`

HasJdbcDriverURL returns a boolean if a field has been set.

### GetDatabaseName

`func (o *AddExternalServerRequest) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *AddExternalServerRequest) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *AddExternalServerRequest) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *AddExternalServerRequest) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetValidationQuery

`func (o *AddExternalServerRequest) GetValidationQuery() string`

GetValidationQuery returns the ValidationQuery field if non-nil, zero value otherwise.

### GetValidationQueryOk

`func (o *AddExternalServerRequest) GetValidationQueryOk() (*string, bool)`

GetValidationQueryOk returns a tuple with the ValidationQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationQuery

`func (o *AddExternalServerRequest) SetValidationQuery(v string)`

SetValidationQuery sets ValidationQuery field to given value.

### HasValidationQuery

`func (o *AddExternalServerRequest) HasValidationQuery() bool`

HasValidationQuery returns a boolean if a field has been set.

### GetValidationQueryTimeout

`func (o *AddExternalServerRequest) GetValidationQueryTimeout() string`

GetValidationQueryTimeout returns the ValidationQueryTimeout field if non-nil, zero value otherwise.

### GetValidationQueryTimeoutOk

`func (o *AddExternalServerRequest) GetValidationQueryTimeoutOk() (*string, bool)`

GetValidationQueryTimeoutOk returns a tuple with the ValidationQueryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationQueryTimeout

`func (o *AddExternalServerRequest) SetValidationQueryTimeout(v string)`

SetValidationQueryTimeout sets ValidationQueryTimeout field to given value.

### HasValidationQueryTimeout

`func (o *AddExternalServerRequest) HasValidationQueryTimeout() bool`

HasValidationQueryTimeout returns a boolean if a field has been set.

### GetJdbcConnectionProperties

`func (o *AddExternalServerRequest) GetJdbcConnectionProperties() []string`

GetJdbcConnectionProperties returns the JdbcConnectionProperties field if non-nil, zero value otherwise.

### GetJdbcConnectionPropertiesOk

`func (o *AddExternalServerRequest) GetJdbcConnectionPropertiesOk() (*[]string, bool)`

GetJdbcConnectionPropertiesOk returns a tuple with the JdbcConnectionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdbcConnectionProperties

`func (o *AddExternalServerRequest) SetJdbcConnectionProperties(v []string)`

SetJdbcConnectionProperties sets JdbcConnectionProperties field to given value.

### HasJdbcConnectionProperties

`func (o *AddExternalServerRequest) HasJdbcConnectionProperties() bool`

HasJdbcConnectionProperties returns a boolean if a field has been set.

### GetTransactionIsolationLevel

`func (o *AddExternalServerRequest) GetTransactionIsolationLevel() EnumexternalServerTransactionIsolationLevelProp`

GetTransactionIsolationLevel returns the TransactionIsolationLevel field if non-nil, zero value otherwise.

### GetTransactionIsolationLevelOk

`func (o *AddExternalServerRequest) GetTransactionIsolationLevelOk() (*EnumexternalServerTransactionIsolationLevelProp, bool)`

GetTransactionIsolationLevelOk returns a tuple with the TransactionIsolationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIsolationLevel

`func (o *AddExternalServerRequest) SetTransactionIsolationLevel(v EnumexternalServerTransactionIsolationLevelProp)`

SetTransactionIsolationLevel sets TransactionIsolationLevel field to given value.

### HasTransactionIsolationLevel

`func (o *AddExternalServerRequest) HasTransactionIsolationLevel() bool`

HasTransactionIsolationLevel returns a boolean if a field has been set.

### GetTransportMechanism

`func (o *AddExternalServerRequest) GetTransportMechanism() EnumexternalServerTransportMechanismProp`

GetTransportMechanism returns the TransportMechanism field if non-nil, zero value otherwise.

### GetTransportMechanismOk

`func (o *AddExternalServerRequest) GetTransportMechanismOk() (*EnumexternalServerTransportMechanismProp, bool)`

GetTransportMechanismOk returns a tuple with the TransportMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportMechanism

`func (o *AddExternalServerRequest) SetTransportMechanism(v EnumexternalServerTransportMechanismProp)`

SetTransportMechanism sets TransportMechanism field to given value.


### GetBasicAuthenticationUsername

`func (o *AddExternalServerRequest) GetBasicAuthenticationUsername() string`

GetBasicAuthenticationUsername returns the BasicAuthenticationUsername field if non-nil, zero value otherwise.

### GetBasicAuthenticationUsernameOk

`func (o *AddExternalServerRequest) GetBasicAuthenticationUsernameOk() (*string, bool)`

GetBasicAuthenticationUsernameOk returns a tuple with the BasicAuthenticationUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthenticationUsername

`func (o *AddExternalServerRequest) SetBasicAuthenticationUsername(v string)`

SetBasicAuthenticationUsername sets BasicAuthenticationUsername field to given value.

### HasBasicAuthenticationUsername

`func (o *AddExternalServerRequest) HasBasicAuthenticationUsername() bool`

HasBasicAuthenticationUsername returns a boolean if a field has been set.

### GetBasicAuthenticationPassphraseProvider

`func (o *AddExternalServerRequest) GetBasicAuthenticationPassphraseProvider() string`

GetBasicAuthenticationPassphraseProvider returns the BasicAuthenticationPassphraseProvider field if non-nil, zero value otherwise.

### GetBasicAuthenticationPassphraseProviderOk

`func (o *AddExternalServerRequest) GetBasicAuthenticationPassphraseProviderOk() (*string, bool)`

GetBasicAuthenticationPassphraseProviderOk returns a tuple with the BasicAuthenticationPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthenticationPassphraseProvider

`func (o *AddExternalServerRequest) SetBasicAuthenticationPassphraseProvider(v string)`

SetBasicAuthenticationPassphraseProvider sets BasicAuthenticationPassphraseProvider field to given value.

### HasBasicAuthenticationPassphraseProvider

`func (o *AddExternalServerRequest) HasBasicAuthenticationPassphraseProvider() bool`

HasBasicAuthenticationPassphraseProvider returns a boolean if a field has been set.

### GetHostnameVerificationMethod

`func (o *AddExternalServerRequest) GetHostnameVerificationMethod() EnumexternalServerHttpHostnameVerificationMethodProp`

GetHostnameVerificationMethod returns the HostnameVerificationMethod field if non-nil, zero value otherwise.

### GetHostnameVerificationMethodOk

`func (o *AddExternalServerRequest) GetHostnameVerificationMethodOk() (*EnumexternalServerHttpHostnameVerificationMethodProp, bool)`

GetHostnameVerificationMethodOk returns a tuple with the HostnameVerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameVerificationMethod

`func (o *AddExternalServerRequest) SetHostnameVerificationMethod(v EnumexternalServerHttpHostnameVerificationMethodProp)`

SetHostnameVerificationMethod sets HostnameVerificationMethod field to given value.

### HasHostnameVerificationMethod

`func (o *AddExternalServerRequest) HasHostnameVerificationMethod() bool`

HasHostnameVerificationMethod returns a boolean if a field has been set.

### GetResponseTimeout

`func (o *AddExternalServerRequest) GetResponseTimeout() string`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *AddExternalServerRequest) GetResponseTimeoutOk() (*string, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *AddExternalServerRequest) SetResponseTimeout(v string)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *AddExternalServerRequest) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### GetBaseURL

`func (o *AddExternalServerRequest) GetBaseURL() string`

GetBaseURL returns the BaseURL field if non-nil, zero value otherwise.

### GetBaseURLOk

`func (o *AddExternalServerRequest) GetBaseURLOk() (*string, bool)`

GetBaseURLOk returns a tuple with the BaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseURL

`func (o *AddExternalServerRequest) SetBaseURL(v string)`

SetBaseURL sets BaseURL field to given value.


### GetSslCertNickname

`func (o *AddExternalServerRequest) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *AddExternalServerRequest) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *AddExternalServerRequest) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.

### HasSslCertNickname

`func (o *AddExternalServerRequest) HasSslCertNickname() bool`

HasSslCertNickname returns a boolean if a field has been set.

### GetConjurServerBaseURI

`func (o *AddExternalServerRequest) GetConjurServerBaseURI() []string`

GetConjurServerBaseURI returns the ConjurServerBaseURI field if non-nil, zero value otherwise.

### GetConjurServerBaseURIOk

`func (o *AddExternalServerRequest) GetConjurServerBaseURIOk() (*[]string, bool)`

GetConjurServerBaseURIOk returns a tuple with the ConjurServerBaseURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurServerBaseURI

`func (o *AddExternalServerRequest) SetConjurServerBaseURI(v []string)`

SetConjurServerBaseURI sets ConjurServerBaseURI field to given value.


### GetConjurAuthenticationMethod

`func (o *AddExternalServerRequest) GetConjurAuthenticationMethod() string`

GetConjurAuthenticationMethod returns the ConjurAuthenticationMethod field if non-nil, zero value otherwise.

### GetConjurAuthenticationMethodOk

`func (o *AddExternalServerRequest) GetConjurAuthenticationMethodOk() (*string, bool)`

GetConjurAuthenticationMethodOk returns a tuple with the ConjurAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurAuthenticationMethod

`func (o *AddExternalServerRequest) SetConjurAuthenticationMethod(v string)`

SetConjurAuthenticationMethod sets ConjurAuthenticationMethod field to given value.


### GetConjurAccountName

`func (o *AddExternalServerRequest) GetConjurAccountName() string`

GetConjurAccountName returns the ConjurAccountName field if non-nil, zero value otherwise.

### GetConjurAccountNameOk

`func (o *AddExternalServerRequest) GetConjurAccountNameOk() (*string, bool)`

GetConjurAccountNameOk returns a tuple with the ConjurAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurAccountName

`func (o *AddExternalServerRequest) SetConjurAccountName(v string)`

SetConjurAccountName sets ConjurAccountName field to given value.


### GetHttpConnectTimeout

`func (o *AddExternalServerRequest) GetHttpConnectTimeout() string`

GetHttpConnectTimeout returns the HttpConnectTimeout field if non-nil, zero value otherwise.

### GetHttpConnectTimeoutOk

`func (o *AddExternalServerRequest) GetHttpConnectTimeoutOk() (*string, bool)`

GetHttpConnectTimeoutOk returns a tuple with the HttpConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConnectTimeout

`func (o *AddExternalServerRequest) SetHttpConnectTimeout(v string)`

SetHttpConnectTimeout sets HttpConnectTimeout field to given value.

### HasHttpConnectTimeout

`func (o *AddExternalServerRequest) HasHttpConnectTimeout() bool`

HasHttpConnectTimeout returns a boolean if a field has been set.

### GetHttpResponseTimeout

`func (o *AddExternalServerRequest) GetHttpResponseTimeout() string`

GetHttpResponseTimeout returns the HttpResponseTimeout field if non-nil, zero value otherwise.

### GetHttpResponseTimeoutOk

`func (o *AddExternalServerRequest) GetHttpResponseTimeoutOk() (*string, bool)`

GetHttpResponseTimeoutOk returns a tuple with the HttpResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponseTimeout

`func (o *AddExternalServerRequest) SetHttpResponseTimeout(v string)`

SetHttpResponseTimeout sets HttpResponseTimeout field to given value.

### HasHttpResponseTimeout

`func (o *AddExternalServerRequest) HasHttpResponseTimeout() bool`

HasHttpResponseTimeout returns a boolean if a field has been set.

### GetTrustStoreFile

`func (o *AddExternalServerRequest) GetTrustStoreFile() string`

GetTrustStoreFile returns the TrustStoreFile field if non-nil, zero value otherwise.

### GetTrustStoreFileOk

`func (o *AddExternalServerRequest) GetTrustStoreFileOk() (*string, bool)`

GetTrustStoreFileOk returns a tuple with the TrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreFile

`func (o *AddExternalServerRequest) SetTrustStoreFile(v string)`

SetTrustStoreFile sets TrustStoreFile field to given value.

### HasTrustStoreFile

`func (o *AddExternalServerRequest) HasTrustStoreFile() bool`

HasTrustStoreFile returns a boolean if a field has been set.

### GetTrustStorePin

`func (o *AddExternalServerRequest) GetTrustStorePin() string`

GetTrustStorePin returns the TrustStorePin field if non-nil, zero value otherwise.

### GetTrustStorePinOk

`func (o *AddExternalServerRequest) GetTrustStorePinOk() (*string, bool)`

GetTrustStorePinOk returns a tuple with the TrustStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePin

`func (o *AddExternalServerRequest) SetTrustStorePin(v string)`

SetTrustStorePin sets TrustStorePin field to given value.

### HasTrustStorePin

`func (o *AddExternalServerRequest) HasTrustStorePin() bool`

HasTrustStorePin returns a boolean if a field has been set.

### GetTrustStoreType

`func (o *AddExternalServerRequest) GetTrustStoreType() string`

GetTrustStoreType returns the TrustStoreType field if non-nil, zero value otherwise.

### GetTrustStoreTypeOk

`func (o *AddExternalServerRequest) GetTrustStoreTypeOk() (*string, bool)`

GetTrustStoreTypeOk returns a tuple with the TrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreType

`func (o *AddExternalServerRequest) SetTrustStoreType(v string)`

SetTrustStoreType sets TrustStoreType field to given value.

### HasTrustStoreType

`func (o *AddExternalServerRequest) HasTrustStoreType() bool`

HasTrustStoreType returns a boolean if a field has been set.

### GetHttpProxyExternalServer

`func (o *AddExternalServerRequest) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *AddExternalServerRequest) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *AddExternalServerRequest) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *AddExternalServerRequest) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetAwsAccessKeyID

`func (o *AddExternalServerRequest) GetAwsAccessKeyID() string`

GetAwsAccessKeyID returns the AwsAccessKeyID field if non-nil, zero value otherwise.

### GetAwsAccessKeyIDOk

`func (o *AddExternalServerRequest) GetAwsAccessKeyIDOk() (*string, bool)`

GetAwsAccessKeyIDOk returns a tuple with the AwsAccessKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyID

`func (o *AddExternalServerRequest) SetAwsAccessKeyID(v string)`

SetAwsAccessKeyID sets AwsAccessKeyID field to given value.

### HasAwsAccessKeyID

`func (o *AddExternalServerRequest) HasAwsAccessKeyID() bool`

HasAwsAccessKeyID returns a boolean if a field has been set.

### GetAwsSecretAccessKey

`func (o *AddExternalServerRequest) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *AddExternalServerRequest) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *AddExternalServerRequest) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *AddExternalServerRequest) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### GetAwsRegionName

`func (o *AddExternalServerRequest) GetAwsRegionName() string`

GetAwsRegionName returns the AwsRegionName field if non-nil, zero value otherwise.

### GetAwsRegionNameOk

`func (o *AddExternalServerRequest) GetAwsRegionNameOk() (*string, bool)`

GetAwsRegionNameOk returns a tuple with the AwsRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegionName

`func (o *AddExternalServerRequest) SetAwsRegionName(v string)`

SetAwsRegionName sets AwsRegionName field to given value.


### GetVaultServerBaseURI

`func (o *AddExternalServerRequest) GetVaultServerBaseURI() []string`

GetVaultServerBaseURI returns the VaultServerBaseURI field if non-nil, zero value otherwise.

### GetVaultServerBaseURIOk

`func (o *AddExternalServerRequest) GetVaultServerBaseURIOk() (*[]string, bool)`

GetVaultServerBaseURIOk returns a tuple with the VaultServerBaseURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultServerBaseURI

`func (o *AddExternalServerRequest) SetVaultServerBaseURI(v []string)`

SetVaultServerBaseURI sets VaultServerBaseURI field to given value.


### GetVaultAuthenticationMethod

`func (o *AddExternalServerRequest) GetVaultAuthenticationMethod() string`

GetVaultAuthenticationMethod returns the VaultAuthenticationMethod field if non-nil, zero value otherwise.

### GetVaultAuthenticationMethodOk

`func (o *AddExternalServerRequest) GetVaultAuthenticationMethodOk() (*string, bool)`

GetVaultAuthenticationMethodOk returns a tuple with the VaultAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAuthenticationMethod

`func (o *AddExternalServerRequest) SetVaultAuthenticationMethod(v string)`

SetVaultAuthenticationMethod sets VaultAuthenticationMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


