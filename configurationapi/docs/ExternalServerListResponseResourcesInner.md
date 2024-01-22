# ExternalServerListResponseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]Enumscim2ExternalServerSchemaUrn**](Enumscim2ExternalServerSchemaUrn.md) |  | 
**Id** | **string** | Name of the External Server | 
**BaseURL** | **string** | The base URL of the external server, optionally including port number, for example \&quot;https://externalService:9031\&quot;. | 
**UserName** | **string** | The name of the login account to use when connecting to the smtp server. Both username and password must be supplied if this attribute is set. | 
**Password** | **string** | The login password for the specified user. | 
**HostnameVerificationMethod** | Pointer to [**EnumexternalServerScim2HostnameVerificationMethodProp**](EnumexternalServerScim2HostnameVerificationMethodProp.md) |  | [optional] 
**KeyManagerProvider** | Pointer to **string** | The key manager provider to use if it is necessary to present a client certificate to the SCIMv2 server. | [optional] 
**TrustManagerProvider** | **string** | The trust manager provider to use to determine whether to trust the certificate presented by the SCIMv2 server during TLS negotiation. | 
**SslCertNickname** | Pointer to **string** | The nickname (alias) of the entry in the associated key store that holds the client certificate chain to present to the SCIMv2 server during TLS negotiation. | [optional] 
**ConnectTimeout** | **string** | Specifies the maximum length of time to wait for a connection to be established before giving up and considering the server unavailable. | 
**ResponseTimeout** | Pointer to **string** | The maximum length of time to wait for a response from the SCIMv2 server when processing operations. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**ScimServiceURL** | **string** | The base URL for the SCIMv2 service. It must include the hostname, port, and base path to use to make SCIMv2 calls. | 
**PassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the login password for the specified user. | [optional] 
**Location** | Pointer to **string** | Specifies the location for the LDAP External Server. | [optional] 
**ConnectionSecurity** | [**EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp**](EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp.md) |  | 
**AuthenticationMethod** | [**EnumexternalServerAmazonAwsAuthenticationMethodProp**](EnumexternalServerAmazonAwsAuthenticationMethodProp.md) |  | 
**HttpProxyExternalServer** | Pointer to **string** | A reference to an HTTP proxy server that should be used for requests sent to the SCIMv2 service. | [optional] 
**OAuthTokenType** | Pointer to [**EnumexternalServerOAuthTokenTypeProp**](EnumexternalServerOAuthTokenTypeProp.md) |  | [optional] 
**OAuthToken** | Pointer to **string** | The token to use in conjunction with the OAuth authentication-method and the chosen oauth-token-type. | [optional] 
**VerifyCredentialsMethod** | [**EnumexternalServerVerifyCredentialsMethodProp**](EnumexternalServerVerifyCredentialsMethodProp.md) |  | 
**UseAdministrativeOperationControl** | Pointer to **bool** | Indicates whether to include the administrative operation request control in requests sent to this server which are intended for administrative operations (e.g., health checking) rather than requests directly from clients. | [optional] 
**ServerHostName** | **string** | The host name or IP address of the target LDAP server. | 
**ServerPort** | **int64** | The port number on which the server listens for requests. | 
**BindDN** | Pointer to **string** | The DN to use to bind to the target LDAP server if simple authentication is required. | [optional] 
**HealthCheckConnectTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for a connection to be established for the purpose of performing a health check. If the connection cannot be established within this length of time, the server will be classified as unavailable. | [optional] 
**MaxConnectionAge** | **string** | Specifies the maximum length of time that connections to this server should be allowed to remain established before being closed and replaced with newly-established connections. | 
**MinExpiredConnectionDisconnectInterval** | Pointer to **string** | Specifies the minimum length of time that should pass between connection closures as a result of the connections being established for longer than the maximum connection age. This may help avoid cases in which a large number of connections are closed and re-established in a short period of time because of the maximum connection age. | [optional] 
**MaxResponseSize** | **string** | Specifies the maximum response size that should be supported for messages received from the LDAP external server. | 
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
**AllowedHeader** | Pointer to **[]string** | A list of HTTP headers that will be forwarded by the PingAuthorize Server to the downstream API server. | [optional] 
**SyncServerPriorityIndex** | **int64** | The relative failover priority of this server. Lower numbers have a higher priority. | 
**VaultServerBaseURI** | **[]string** | The base URL needed to access the Vault server. The base URL should consist of the protocol (\&quot;http\&quot; or \&quot;https\&quot;), the server address (resolvable name or IP address), and the port number. For example, \&quot;https://vault.example.com:8200/\&quot;. | 
**VaultAuthenticationMethod** | **string** | The mechanism used to authenticate to the Vault server. | 
**HttpConnectTimeout** | Pointer to **string** | The maximum length of time to wait to obtain an HTTP connection. | [optional] 
**HttpResponseTimeout** | Pointer to **string** | The maximum length of time to wait for a response to an HTTP request. | [optional] 
**TrustStoreFile** | Pointer to **string** | The path to a file containing the information needed to trust the certificate presented by the Conjur servers. | [optional] 
**TrustStorePin** | Pointer to **string** | The PIN needed to access the contents of the trust store. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents. | [optional] 
**TrustStoreType** | Pointer to **string** | The store type for the specified trust store file. The value should likely be one of \&quot;JKS\&quot;, \&quot;PKCS12\&quot;, or \&quot;BCFKS\&quot;. | [optional] 
**UserID** | **string** | Specifies the user ID to authenticate calls to the policy server&#39;s governance engine API. | 
**SharedSecret** | **string** | Specifies the shared secret to authenticate calls to the policy server&#39;s governance engine API. | 
**DecisionNode** | Pointer to **string** | Specifies the ID of the policy tree node that will act as the root node for policy evaluation. | [optional] 
**Branch** | Pointer to **string** | Specifies the name of the policy branch to use for policy evaluation. | [optional] 
**Snapshot** | Pointer to **string** | Specifies the ID of a specific commit to use for policy evaluation. | [optional] 
**SmtpSecurity** | Pointer to [**EnumexternalServerSmtpSecurityProp**](EnumexternalServerSmtpSecurityProp.md) |  | [optional] 
**SmtpTimeout** | Pointer to **string** | Specifies the maximum length of time that a connection or attempted connection to a SMTP server may take. | [optional] 
**SmtpConnectionProperties** | Pointer to **[]string** | Specifies the connection properties for the smtp server. | [optional] 
**ServerHTTPPort** | **int64** | The port number on which the server listens for HTTP requests. | 
**UseSSL** | Pointer to **bool** | If enabled, the Kafka Cluster External Server will use SSL to encrypt communication with the Kafka brokers. | [optional] 
**BasicAuthenticationUsername** | Pointer to **string** | The username to use to authenticate to the HTTP Proxy External Server. | [optional] 
**BasicAuthenticationPassphraseProvider** | Pointer to **string** | A passphrase provider that provides access to the password to use to authenticate to the HTTP Proxy External Server. | [optional] 
**BootstrapServer** | **[]string** | List of Kafka brokers to use for this Kafka Cluster External Server, following the host:port format. | 
**ProducerProperty** | Pointer to **[]string** | Specifies extra properties to use when constructing the KafkaProducer for sending messages. | [optional] 
**BaseDN** | **[]string** | Specifies the base DN stored in this mock resource. | 
**ConjurServerBaseURI** | **[]string** | The base URL needed to access the CyberArk Conjur server. The base URL should consist of the protocol (\&quot;http\&quot; or \&quot;https\&quot;), the server address (resolvable name or IP address), and the port number. For example, \&quot;https://conjur.example.com:8443/\&quot;. | 
**ConjurAuthenticationMethod** | **string** | The mechanism used to authenticate to the Conjur server. | 
**ConjurAccountName** | **string** | The name of the account with which the desired secrets are associated. | 
**AwsAccessKeyID** | Pointer to **string** | The access key ID that will be used if authentication should use an access key. If this is provided, then an aws-secret-access-key must also be provided. | [optional] 
**AwsSecretAccessKey** | Pointer to **string** | The secret access key that will be used if authentication should use an access key. If this is provided, then an aws-access-key-id must also be provided. | [optional] 
**AwsRegionName** | **string** | The name of the AWS region containing the resources that will be accessed. | 
**HttpAuthorizationMethod** | **string** | The method to use to authorize requests sent to the SCIMv2 server. | 
**ClientReconnectInterval** | Pointer to **string** | The maximum length of time that a client instance should remain active before being recreated. | [optional] 

## Methods

### NewExternalServerListResponseResourcesInner

`func NewExternalServerListResponseResourcesInner(schemas []Enumscim2ExternalServerSchemaUrn, id string, baseURL string, userName string, password string, trustManagerProvider string, connectTimeout string, scimServiceURL string, connectionSecurity EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp, authenticationMethod EnumexternalServerAmazonAwsAuthenticationMethodProp, verifyCredentialsMethod EnumexternalServerVerifyCredentialsMethodProp, serverHostName string, serverPort int64, maxConnectionAge string, maxResponseSize string, jdbcDriverType EnumexternalServerJdbcDriverTypeProp, transportMechanism EnumexternalServerTransportMechanismProp, syncServerPriorityIndex int64, vaultServerBaseURI []string, vaultAuthenticationMethod string, userID string, sharedSecret string, serverHTTPPort int64, bootstrapServer []string, baseDN []string, conjurServerBaseURI []string, conjurAuthenticationMethod string, conjurAccountName string, awsRegionName string, httpAuthorizationMethod string, ) *ExternalServerListResponseResourcesInner`

NewExternalServerListResponseResourcesInner instantiates a new ExternalServerListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalServerListResponseResourcesInnerWithDefaults

`func NewExternalServerListResponseResourcesInnerWithDefaults() *ExternalServerListResponseResourcesInner`

NewExternalServerListResponseResourcesInnerWithDefaults instantiates a new ExternalServerListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ExternalServerListResponseResourcesInner) GetSchemas() []Enumscim2ExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ExternalServerListResponseResourcesInner) GetSchemasOk() (*[]Enumscim2ExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ExternalServerListResponseResourcesInner) SetSchemas(v []Enumscim2ExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ExternalServerListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalServerListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalServerListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetBaseURL

`func (o *ExternalServerListResponseResourcesInner) GetBaseURL() string`

GetBaseURL returns the BaseURL field if non-nil, zero value otherwise.

### GetBaseURLOk

`func (o *ExternalServerListResponseResourcesInner) GetBaseURLOk() (*string, bool)`

GetBaseURLOk returns a tuple with the BaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseURL

`func (o *ExternalServerListResponseResourcesInner) SetBaseURL(v string)`

SetBaseURL sets BaseURL field to given value.


### GetUserName

`func (o *ExternalServerListResponseResourcesInner) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ExternalServerListResponseResourcesInner) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ExternalServerListResponseResourcesInner) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetPassword

`func (o *ExternalServerListResponseResourcesInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ExternalServerListResponseResourcesInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ExternalServerListResponseResourcesInner) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetHostnameVerificationMethod

`func (o *ExternalServerListResponseResourcesInner) GetHostnameVerificationMethod() EnumexternalServerScim2HostnameVerificationMethodProp`

GetHostnameVerificationMethod returns the HostnameVerificationMethod field if non-nil, zero value otherwise.

### GetHostnameVerificationMethodOk

`func (o *ExternalServerListResponseResourcesInner) GetHostnameVerificationMethodOk() (*EnumexternalServerScim2HostnameVerificationMethodProp, bool)`

GetHostnameVerificationMethodOk returns a tuple with the HostnameVerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameVerificationMethod

`func (o *ExternalServerListResponseResourcesInner) SetHostnameVerificationMethod(v EnumexternalServerScim2HostnameVerificationMethodProp)`

SetHostnameVerificationMethod sets HostnameVerificationMethod field to given value.

### HasHostnameVerificationMethod

`func (o *ExternalServerListResponseResourcesInner) HasHostnameVerificationMethod() bool`

HasHostnameVerificationMethod returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *ExternalServerListResponseResourcesInner) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *ExternalServerListResponseResourcesInner) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *ExternalServerListResponseResourcesInner) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *ExternalServerListResponseResourcesInner) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *ExternalServerListResponseResourcesInner) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *ExternalServerListResponseResourcesInner) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *ExternalServerListResponseResourcesInner) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.


### GetSslCertNickname

`func (o *ExternalServerListResponseResourcesInner) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *ExternalServerListResponseResourcesInner) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *ExternalServerListResponseResourcesInner) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.

### HasSslCertNickname

`func (o *ExternalServerListResponseResourcesInner) HasSslCertNickname() bool`

HasSslCertNickname returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *ExternalServerListResponseResourcesInner) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *ExternalServerListResponseResourcesInner) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *ExternalServerListResponseResourcesInner) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.


### GetResponseTimeout

`func (o *ExternalServerListResponseResourcesInner) GetResponseTimeout() string`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *ExternalServerListResponseResourcesInner) GetResponseTimeoutOk() (*string, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *ExternalServerListResponseResourcesInner) SetResponseTimeout(v string)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *ExternalServerListResponseResourcesInner) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### GetDescription

`func (o *ExternalServerListResponseResourcesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalServerListResponseResourcesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalServerListResponseResourcesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalServerListResponseResourcesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *ExternalServerListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ExternalServerListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ExternalServerListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ExternalServerListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ExternalServerListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ExternalServerListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ExternalServerListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ExternalServerListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetScimServiceURL

`func (o *ExternalServerListResponseResourcesInner) GetScimServiceURL() string`

GetScimServiceURL returns the ScimServiceURL field if non-nil, zero value otherwise.

### GetScimServiceURLOk

`func (o *ExternalServerListResponseResourcesInner) GetScimServiceURLOk() (*string, bool)`

GetScimServiceURLOk returns a tuple with the ScimServiceURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimServiceURL

`func (o *ExternalServerListResponseResourcesInner) SetScimServiceURL(v string)`

SetScimServiceURL sets ScimServiceURL field to given value.


### GetPassphraseProvider

`func (o *ExternalServerListResponseResourcesInner) GetPassphraseProvider() string`

GetPassphraseProvider returns the PassphraseProvider field if non-nil, zero value otherwise.

### GetPassphraseProviderOk

`func (o *ExternalServerListResponseResourcesInner) GetPassphraseProviderOk() (*string, bool)`

GetPassphraseProviderOk returns a tuple with the PassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseProvider

`func (o *ExternalServerListResponseResourcesInner) SetPassphraseProvider(v string)`

SetPassphraseProvider sets PassphraseProvider field to given value.

### HasPassphraseProvider

`func (o *ExternalServerListResponseResourcesInner) HasPassphraseProvider() bool`

HasPassphraseProvider returns a boolean if a field has been set.

### GetLocation

`func (o *ExternalServerListResponseResourcesInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ExternalServerListResponseResourcesInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ExternalServerListResponseResourcesInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ExternalServerListResponseResourcesInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetConnectionSecurity

`func (o *ExternalServerListResponseResourcesInner) GetConnectionSecurity() EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *ExternalServerListResponseResourcesInner) GetConnectionSecurityOk() (*EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *ExternalServerListResponseResourcesInner) SetConnectionSecurity(v EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp)`

SetConnectionSecurity sets ConnectionSecurity field to given value.


### GetAuthenticationMethod

`func (o *ExternalServerListResponseResourcesInner) GetAuthenticationMethod() EnumexternalServerAmazonAwsAuthenticationMethodProp`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *ExternalServerListResponseResourcesInner) GetAuthenticationMethodOk() (*EnumexternalServerAmazonAwsAuthenticationMethodProp, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *ExternalServerListResponseResourcesInner) SetAuthenticationMethod(v EnumexternalServerAmazonAwsAuthenticationMethodProp)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.


### GetHttpProxyExternalServer

`func (o *ExternalServerListResponseResourcesInner) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *ExternalServerListResponseResourcesInner) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *ExternalServerListResponseResourcesInner) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *ExternalServerListResponseResourcesInner) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetOAuthTokenType

`func (o *ExternalServerListResponseResourcesInner) GetOAuthTokenType() EnumexternalServerOAuthTokenTypeProp`

GetOAuthTokenType returns the OAuthTokenType field if non-nil, zero value otherwise.

### GetOAuthTokenTypeOk

`func (o *ExternalServerListResponseResourcesInner) GetOAuthTokenTypeOk() (*EnumexternalServerOAuthTokenTypeProp, bool)`

GetOAuthTokenTypeOk returns a tuple with the OAuthTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthTokenType

`func (o *ExternalServerListResponseResourcesInner) SetOAuthTokenType(v EnumexternalServerOAuthTokenTypeProp)`

SetOAuthTokenType sets OAuthTokenType field to given value.

### HasOAuthTokenType

`func (o *ExternalServerListResponseResourcesInner) HasOAuthTokenType() bool`

HasOAuthTokenType returns a boolean if a field has been set.

### GetOAuthToken

`func (o *ExternalServerListResponseResourcesInner) GetOAuthToken() string`

GetOAuthToken returns the OAuthToken field if non-nil, zero value otherwise.

### GetOAuthTokenOk

`func (o *ExternalServerListResponseResourcesInner) GetOAuthTokenOk() (*string, bool)`

GetOAuthTokenOk returns a tuple with the OAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthToken

`func (o *ExternalServerListResponseResourcesInner) SetOAuthToken(v string)`

SetOAuthToken sets OAuthToken field to given value.

### HasOAuthToken

`func (o *ExternalServerListResponseResourcesInner) HasOAuthToken() bool`

HasOAuthToken returns a boolean if a field has been set.

### GetVerifyCredentialsMethod

`func (o *ExternalServerListResponseResourcesInner) GetVerifyCredentialsMethod() EnumexternalServerVerifyCredentialsMethodProp`

GetVerifyCredentialsMethod returns the VerifyCredentialsMethod field if non-nil, zero value otherwise.

### GetVerifyCredentialsMethodOk

`func (o *ExternalServerListResponseResourcesInner) GetVerifyCredentialsMethodOk() (*EnumexternalServerVerifyCredentialsMethodProp, bool)`

GetVerifyCredentialsMethodOk returns a tuple with the VerifyCredentialsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCredentialsMethod

`func (o *ExternalServerListResponseResourcesInner) SetVerifyCredentialsMethod(v EnumexternalServerVerifyCredentialsMethodProp)`

SetVerifyCredentialsMethod sets VerifyCredentialsMethod field to given value.


### GetUseAdministrativeOperationControl

`func (o *ExternalServerListResponseResourcesInner) GetUseAdministrativeOperationControl() bool`

GetUseAdministrativeOperationControl returns the UseAdministrativeOperationControl field if non-nil, zero value otherwise.

### GetUseAdministrativeOperationControlOk

`func (o *ExternalServerListResponseResourcesInner) GetUseAdministrativeOperationControlOk() (*bool, bool)`

GetUseAdministrativeOperationControlOk returns a tuple with the UseAdministrativeOperationControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAdministrativeOperationControl

`func (o *ExternalServerListResponseResourcesInner) SetUseAdministrativeOperationControl(v bool)`

SetUseAdministrativeOperationControl sets UseAdministrativeOperationControl field to given value.

### HasUseAdministrativeOperationControl

`func (o *ExternalServerListResponseResourcesInner) HasUseAdministrativeOperationControl() bool`

HasUseAdministrativeOperationControl returns a boolean if a field has been set.

### GetServerHostName

`func (o *ExternalServerListResponseResourcesInner) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *ExternalServerListResponseResourcesInner) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *ExternalServerListResponseResourcesInner) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *ExternalServerListResponseResourcesInner) GetServerPort() int64`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *ExternalServerListResponseResourcesInner) GetServerPortOk() (*int64, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *ExternalServerListResponseResourcesInner) SetServerPort(v int64)`

SetServerPort sets ServerPort field to given value.


### GetBindDN

`func (o *ExternalServerListResponseResourcesInner) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *ExternalServerListResponseResourcesInner) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *ExternalServerListResponseResourcesInner) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.

### HasBindDN

`func (o *ExternalServerListResponseResourcesInner) HasBindDN() bool`

HasBindDN returns a boolean if a field has been set.

### GetHealthCheckConnectTimeout

`func (o *ExternalServerListResponseResourcesInner) GetHealthCheckConnectTimeout() string`

GetHealthCheckConnectTimeout returns the HealthCheckConnectTimeout field if non-nil, zero value otherwise.

### GetHealthCheckConnectTimeoutOk

`func (o *ExternalServerListResponseResourcesInner) GetHealthCheckConnectTimeoutOk() (*string, bool)`

GetHealthCheckConnectTimeoutOk returns a tuple with the HealthCheckConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckConnectTimeout

`func (o *ExternalServerListResponseResourcesInner) SetHealthCheckConnectTimeout(v string)`

SetHealthCheckConnectTimeout sets HealthCheckConnectTimeout field to given value.

### HasHealthCheckConnectTimeout

`func (o *ExternalServerListResponseResourcesInner) HasHealthCheckConnectTimeout() bool`

HasHealthCheckConnectTimeout returns a boolean if a field has been set.

### GetMaxConnectionAge

`func (o *ExternalServerListResponseResourcesInner) GetMaxConnectionAge() string`

GetMaxConnectionAge returns the MaxConnectionAge field if non-nil, zero value otherwise.

### GetMaxConnectionAgeOk

`func (o *ExternalServerListResponseResourcesInner) GetMaxConnectionAgeOk() (*string, bool)`

GetMaxConnectionAgeOk returns a tuple with the MaxConnectionAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionAge

`func (o *ExternalServerListResponseResourcesInner) SetMaxConnectionAge(v string)`

SetMaxConnectionAge sets MaxConnectionAge field to given value.


### GetMinExpiredConnectionDisconnectInterval

`func (o *ExternalServerListResponseResourcesInner) GetMinExpiredConnectionDisconnectInterval() string`

GetMinExpiredConnectionDisconnectInterval returns the MinExpiredConnectionDisconnectInterval field if non-nil, zero value otherwise.

### GetMinExpiredConnectionDisconnectIntervalOk

`func (o *ExternalServerListResponseResourcesInner) GetMinExpiredConnectionDisconnectIntervalOk() (*string, bool)`

GetMinExpiredConnectionDisconnectIntervalOk returns a tuple with the MinExpiredConnectionDisconnectInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinExpiredConnectionDisconnectInterval

`func (o *ExternalServerListResponseResourcesInner) SetMinExpiredConnectionDisconnectInterval(v string)`

SetMinExpiredConnectionDisconnectInterval sets MinExpiredConnectionDisconnectInterval field to given value.

### HasMinExpiredConnectionDisconnectInterval

`func (o *ExternalServerListResponseResourcesInner) HasMinExpiredConnectionDisconnectInterval() bool`

HasMinExpiredConnectionDisconnectInterval returns a boolean if a field has been set.

### GetMaxResponseSize

`func (o *ExternalServerListResponseResourcesInner) GetMaxResponseSize() string`

GetMaxResponseSize returns the MaxResponseSize field if non-nil, zero value otherwise.

### GetMaxResponseSizeOk

`func (o *ExternalServerListResponseResourcesInner) GetMaxResponseSizeOk() (*string, bool)`

GetMaxResponseSizeOk returns a tuple with the MaxResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseSize

`func (o *ExternalServerListResponseResourcesInner) SetMaxResponseSize(v string)`

SetMaxResponseSize sets MaxResponseSize field to given value.


### GetInitialConnections

`func (o *ExternalServerListResponseResourcesInner) GetInitialConnections() int64`

GetInitialConnections returns the InitialConnections field if non-nil, zero value otherwise.

### GetInitialConnectionsOk

`func (o *ExternalServerListResponseResourcesInner) GetInitialConnectionsOk() (*int64, bool)`

GetInitialConnectionsOk returns a tuple with the InitialConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialConnections

`func (o *ExternalServerListResponseResourcesInner) SetInitialConnections(v int64)`

SetInitialConnections sets InitialConnections field to given value.

### HasInitialConnections

`func (o *ExternalServerListResponseResourcesInner) HasInitialConnections() bool`

HasInitialConnections returns a boolean if a field has been set.

### GetMaxConnections

`func (o *ExternalServerListResponseResourcesInner) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *ExternalServerListResponseResourcesInner) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *ExternalServerListResponseResourcesInner) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *ExternalServerListResponseResourcesInner) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetDefunctConnectionResultCode

`func (o *ExternalServerListResponseResourcesInner) GetDefunctConnectionResultCode() []EnumexternalServerDefunctConnectionResultCodeProp`

GetDefunctConnectionResultCode returns the DefunctConnectionResultCode field if non-nil, zero value otherwise.

### GetDefunctConnectionResultCodeOk

`func (o *ExternalServerListResponseResourcesInner) GetDefunctConnectionResultCodeOk() (*[]EnumexternalServerDefunctConnectionResultCodeProp, bool)`

GetDefunctConnectionResultCodeOk returns a tuple with the DefunctConnectionResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefunctConnectionResultCode

`func (o *ExternalServerListResponseResourcesInner) SetDefunctConnectionResultCode(v []EnumexternalServerDefunctConnectionResultCodeProp)`

SetDefunctConnectionResultCode sets DefunctConnectionResultCode field to given value.

### HasDefunctConnectionResultCode

`func (o *ExternalServerListResponseResourcesInner) HasDefunctConnectionResultCode() bool`

HasDefunctConnectionResultCode returns a boolean if a field has been set.

### GetAbandonOnTimeout

`func (o *ExternalServerListResponseResourcesInner) GetAbandonOnTimeout() bool`

GetAbandonOnTimeout returns the AbandonOnTimeout field if non-nil, zero value otherwise.

### GetAbandonOnTimeoutOk

`func (o *ExternalServerListResponseResourcesInner) GetAbandonOnTimeoutOk() (*bool, bool)`

GetAbandonOnTimeoutOk returns a tuple with the AbandonOnTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonOnTimeout

`func (o *ExternalServerListResponseResourcesInner) SetAbandonOnTimeout(v bool)`

SetAbandonOnTimeout sets AbandonOnTimeout field to given value.

### HasAbandonOnTimeout

`func (o *ExternalServerListResponseResourcesInner) HasAbandonOnTimeout() bool`

HasAbandonOnTimeout returns a boolean if a field has been set.

### GetJdbcDriverType

`func (o *ExternalServerListResponseResourcesInner) GetJdbcDriverType() EnumexternalServerJdbcDriverTypeProp`

GetJdbcDriverType returns the JdbcDriverType field if non-nil, zero value otherwise.

### GetJdbcDriverTypeOk

`func (o *ExternalServerListResponseResourcesInner) GetJdbcDriverTypeOk() (*EnumexternalServerJdbcDriverTypeProp, bool)`

GetJdbcDriverTypeOk returns a tuple with the JdbcDriverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdbcDriverType

`func (o *ExternalServerListResponseResourcesInner) SetJdbcDriverType(v EnumexternalServerJdbcDriverTypeProp)`

SetJdbcDriverType sets JdbcDriverType field to given value.


### GetJdbcDriverURL

`func (o *ExternalServerListResponseResourcesInner) GetJdbcDriverURL() string`

GetJdbcDriverURL returns the JdbcDriverURL field if non-nil, zero value otherwise.

### GetJdbcDriverURLOk

`func (o *ExternalServerListResponseResourcesInner) GetJdbcDriverURLOk() (*string, bool)`

GetJdbcDriverURLOk returns a tuple with the JdbcDriverURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdbcDriverURL

`func (o *ExternalServerListResponseResourcesInner) SetJdbcDriverURL(v string)`

SetJdbcDriverURL sets JdbcDriverURL field to given value.

### HasJdbcDriverURL

`func (o *ExternalServerListResponseResourcesInner) HasJdbcDriverURL() bool`

HasJdbcDriverURL returns a boolean if a field has been set.

### GetDatabaseName

`func (o *ExternalServerListResponseResourcesInner) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ExternalServerListResponseResourcesInner) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ExternalServerListResponseResourcesInner) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *ExternalServerListResponseResourcesInner) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetValidationQuery

`func (o *ExternalServerListResponseResourcesInner) GetValidationQuery() string`

GetValidationQuery returns the ValidationQuery field if non-nil, zero value otherwise.

### GetValidationQueryOk

`func (o *ExternalServerListResponseResourcesInner) GetValidationQueryOk() (*string, bool)`

GetValidationQueryOk returns a tuple with the ValidationQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationQuery

`func (o *ExternalServerListResponseResourcesInner) SetValidationQuery(v string)`

SetValidationQuery sets ValidationQuery field to given value.

### HasValidationQuery

`func (o *ExternalServerListResponseResourcesInner) HasValidationQuery() bool`

HasValidationQuery returns a boolean if a field has been set.

### GetValidationQueryTimeout

`func (o *ExternalServerListResponseResourcesInner) GetValidationQueryTimeout() string`

GetValidationQueryTimeout returns the ValidationQueryTimeout field if non-nil, zero value otherwise.

### GetValidationQueryTimeoutOk

`func (o *ExternalServerListResponseResourcesInner) GetValidationQueryTimeoutOk() (*string, bool)`

GetValidationQueryTimeoutOk returns a tuple with the ValidationQueryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationQueryTimeout

`func (o *ExternalServerListResponseResourcesInner) SetValidationQueryTimeout(v string)`

SetValidationQueryTimeout sets ValidationQueryTimeout field to given value.

### HasValidationQueryTimeout

`func (o *ExternalServerListResponseResourcesInner) HasValidationQueryTimeout() bool`

HasValidationQueryTimeout returns a boolean if a field has been set.

### GetJdbcConnectionProperties

`func (o *ExternalServerListResponseResourcesInner) GetJdbcConnectionProperties() []string`

GetJdbcConnectionProperties returns the JdbcConnectionProperties field if non-nil, zero value otherwise.

### GetJdbcConnectionPropertiesOk

`func (o *ExternalServerListResponseResourcesInner) GetJdbcConnectionPropertiesOk() (*[]string, bool)`

GetJdbcConnectionPropertiesOk returns a tuple with the JdbcConnectionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdbcConnectionProperties

`func (o *ExternalServerListResponseResourcesInner) SetJdbcConnectionProperties(v []string)`

SetJdbcConnectionProperties sets JdbcConnectionProperties field to given value.

### HasJdbcConnectionProperties

`func (o *ExternalServerListResponseResourcesInner) HasJdbcConnectionProperties() bool`

HasJdbcConnectionProperties returns a boolean if a field has been set.

### GetTransactionIsolationLevel

`func (o *ExternalServerListResponseResourcesInner) GetTransactionIsolationLevel() EnumexternalServerTransactionIsolationLevelProp`

GetTransactionIsolationLevel returns the TransactionIsolationLevel field if non-nil, zero value otherwise.

### GetTransactionIsolationLevelOk

`func (o *ExternalServerListResponseResourcesInner) GetTransactionIsolationLevelOk() (*EnumexternalServerTransactionIsolationLevelProp, bool)`

GetTransactionIsolationLevelOk returns a tuple with the TransactionIsolationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIsolationLevel

`func (o *ExternalServerListResponseResourcesInner) SetTransactionIsolationLevel(v EnumexternalServerTransactionIsolationLevelProp)`

SetTransactionIsolationLevel sets TransactionIsolationLevel field to given value.

### HasTransactionIsolationLevel

`func (o *ExternalServerListResponseResourcesInner) HasTransactionIsolationLevel() bool`

HasTransactionIsolationLevel returns a boolean if a field has been set.

### GetTransportMechanism

`func (o *ExternalServerListResponseResourcesInner) GetTransportMechanism() EnumexternalServerTransportMechanismProp`

GetTransportMechanism returns the TransportMechanism field if non-nil, zero value otherwise.

### GetTransportMechanismOk

`func (o *ExternalServerListResponseResourcesInner) GetTransportMechanismOk() (*EnumexternalServerTransportMechanismProp, bool)`

GetTransportMechanismOk returns a tuple with the TransportMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportMechanism

`func (o *ExternalServerListResponseResourcesInner) SetTransportMechanism(v EnumexternalServerTransportMechanismProp)`

SetTransportMechanism sets TransportMechanism field to given value.


### GetAllowedHeader

`func (o *ExternalServerListResponseResourcesInner) GetAllowedHeader() []string`

GetAllowedHeader returns the AllowedHeader field if non-nil, zero value otherwise.

### GetAllowedHeaderOk

`func (o *ExternalServerListResponseResourcesInner) GetAllowedHeaderOk() (*[]string, bool)`

GetAllowedHeaderOk returns a tuple with the AllowedHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHeader

`func (o *ExternalServerListResponseResourcesInner) SetAllowedHeader(v []string)`

SetAllowedHeader sets AllowedHeader field to given value.

### HasAllowedHeader

`func (o *ExternalServerListResponseResourcesInner) HasAllowedHeader() bool`

HasAllowedHeader returns a boolean if a field has been set.

### GetSyncServerPriorityIndex

`func (o *ExternalServerListResponseResourcesInner) GetSyncServerPriorityIndex() int64`

GetSyncServerPriorityIndex returns the SyncServerPriorityIndex field if non-nil, zero value otherwise.

### GetSyncServerPriorityIndexOk

`func (o *ExternalServerListResponseResourcesInner) GetSyncServerPriorityIndexOk() (*int64, bool)`

GetSyncServerPriorityIndexOk returns a tuple with the SyncServerPriorityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncServerPriorityIndex

`func (o *ExternalServerListResponseResourcesInner) SetSyncServerPriorityIndex(v int64)`

SetSyncServerPriorityIndex sets SyncServerPriorityIndex field to given value.


### GetVaultServerBaseURI

`func (o *ExternalServerListResponseResourcesInner) GetVaultServerBaseURI() []string`

GetVaultServerBaseURI returns the VaultServerBaseURI field if non-nil, zero value otherwise.

### GetVaultServerBaseURIOk

`func (o *ExternalServerListResponseResourcesInner) GetVaultServerBaseURIOk() (*[]string, bool)`

GetVaultServerBaseURIOk returns a tuple with the VaultServerBaseURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultServerBaseURI

`func (o *ExternalServerListResponseResourcesInner) SetVaultServerBaseURI(v []string)`

SetVaultServerBaseURI sets VaultServerBaseURI field to given value.


### GetVaultAuthenticationMethod

`func (o *ExternalServerListResponseResourcesInner) GetVaultAuthenticationMethod() string`

GetVaultAuthenticationMethod returns the VaultAuthenticationMethod field if non-nil, zero value otherwise.

### GetVaultAuthenticationMethodOk

`func (o *ExternalServerListResponseResourcesInner) GetVaultAuthenticationMethodOk() (*string, bool)`

GetVaultAuthenticationMethodOk returns a tuple with the VaultAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAuthenticationMethod

`func (o *ExternalServerListResponseResourcesInner) SetVaultAuthenticationMethod(v string)`

SetVaultAuthenticationMethod sets VaultAuthenticationMethod field to given value.


### GetHttpConnectTimeout

`func (o *ExternalServerListResponseResourcesInner) GetHttpConnectTimeout() string`

GetHttpConnectTimeout returns the HttpConnectTimeout field if non-nil, zero value otherwise.

### GetHttpConnectTimeoutOk

`func (o *ExternalServerListResponseResourcesInner) GetHttpConnectTimeoutOk() (*string, bool)`

GetHttpConnectTimeoutOk returns a tuple with the HttpConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConnectTimeout

`func (o *ExternalServerListResponseResourcesInner) SetHttpConnectTimeout(v string)`

SetHttpConnectTimeout sets HttpConnectTimeout field to given value.

### HasHttpConnectTimeout

`func (o *ExternalServerListResponseResourcesInner) HasHttpConnectTimeout() bool`

HasHttpConnectTimeout returns a boolean if a field has been set.

### GetHttpResponseTimeout

`func (o *ExternalServerListResponseResourcesInner) GetHttpResponseTimeout() string`

GetHttpResponseTimeout returns the HttpResponseTimeout field if non-nil, zero value otherwise.

### GetHttpResponseTimeoutOk

`func (o *ExternalServerListResponseResourcesInner) GetHttpResponseTimeoutOk() (*string, bool)`

GetHttpResponseTimeoutOk returns a tuple with the HttpResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponseTimeout

`func (o *ExternalServerListResponseResourcesInner) SetHttpResponseTimeout(v string)`

SetHttpResponseTimeout sets HttpResponseTimeout field to given value.

### HasHttpResponseTimeout

`func (o *ExternalServerListResponseResourcesInner) HasHttpResponseTimeout() bool`

HasHttpResponseTimeout returns a boolean if a field has been set.

### GetTrustStoreFile

`func (o *ExternalServerListResponseResourcesInner) GetTrustStoreFile() string`

GetTrustStoreFile returns the TrustStoreFile field if non-nil, zero value otherwise.

### GetTrustStoreFileOk

`func (o *ExternalServerListResponseResourcesInner) GetTrustStoreFileOk() (*string, bool)`

GetTrustStoreFileOk returns a tuple with the TrustStoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreFile

`func (o *ExternalServerListResponseResourcesInner) SetTrustStoreFile(v string)`

SetTrustStoreFile sets TrustStoreFile field to given value.

### HasTrustStoreFile

`func (o *ExternalServerListResponseResourcesInner) HasTrustStoreFile() bool`

HasTrustStoreFile returns a boolean if a field has been set.

### GetTrustStorePin

`func (o *ExternalServerListResponseResourcesInner) GetTrustStorePin() string`

GetTrustStorePin returns the TrustStorePin field if non-nil, zero value otherwise.

### GetTrustStorePinOk

`func (o *ExternalServerListResponseResourcesInner) GetTrustStorePinOk() (*string, bool)`

GetTrustStorePinOk returns a tuple with the TrustStorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStorePin

`func (o *ExternalServerListResponseResourcesInner) SetTrustStorePin(v string)`

SetTrustStorePin sets TrustStorePin field to given value.

### HasTrustStorePin

`func (o *ExternalServerListResponseResourcesInner) HasTrustStorePin() bool`

HasTrustStorePin returns a boolean if a field has been set.

### GetTrustStoreType

`func (o *ExternalServerListResponseResourcesInner) GetTrustStoreType() string`

GetTrustStoreType returns the TrustStoreType field if non-nil, zero value otherwise.

### GetTrustStoreTypeOk

`func (o *ExternalServerListResponseResourcesInner) GetTrustStoreTypeOk() (*string, bool)`

GetTrustStoreTypeOk returns a tuple with the TrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStoreType

`func (o *ExternalServerListResponseResourcesInner) SetTrustStoreType(v string)`

SetTrustStoreType sets TrustStoreType field to given value.

### HasTrustStoreType

`func (o *ExternalServerListResponseResourcesInner) HasTrustStoreType() bool`

HasTrustStoreType returns a boolean if a field has been set.

### GetUserID

`func (o *ExternalServerListResponseResourcesInner) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *ExternalServerListResponseResourcesInner) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *ExternalServerListResponseResourcesInner) SetUserID(v string)`

SetUserID sets UserID field to given value.


### GetSharedSecret

`func (o *ExternalServerListResponseResourcesInner) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *ExternalServerListResponseResourcesInner) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *ExternalServerListResponseResourcesInner) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.


### GetDecisionNode

`func (o *ExternalServerListResponseResourcesInner) GetDecisionNode() string`

GetDecisionNode returns the DecisionNode field if non-nil, zero value otherwise.

### GetDecisionNodeOk

`func (o *ExternalServerListResponseResourcesInner) GetDecisionNodeOk() (*string, bool)`

GetDecisionNodeOk returns a tuple with the DecisionNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionNode

`func (o *ExternalServerListResponseResourcesInner) SetDecisionNode(v string)`

SetDecisionNode sets DecisionNode field to given value.

### HasDecisionNode

`func (o *ExternalServerListResponseResourcesInner) HasDecisionNode() bool`

HasDecisionNode returns a boolean if a field has been set.

### GetBranch

`func (o *ExternalServerListResponseResourcesInner) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ExternalServerListResponseResourcesInner) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ExternalServerListResponseResourcesInner) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ExternalServerListResponseResourcesInner) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetSnapshot

`func (o *ExternalServerListResponseResourcesInner) GetSnapshot() string`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *ExternalServerListResponseResourcesInner) GetSnapshotOk() (*string, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *ExternalServerListResponseResourcesInner) SetSnapshot(v string)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *ExternalServerListResponseResourcesInner) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetSmtpSecurity

`func (o *ExternalServerListResponseResourcesInner) GetSmtpSecurity() EnumexternalServerSmtpSecurityProp`

GetSmtpSecurity returns the SmtpSecurity field if non-nil, zero value otherwise.

### GetSmtpSecurityOk

`func (o *ExternalServerListResponseResourcesInner) GetSmtpSecurityOk() (*EnumexternalServerSmtpSecurityProp, bool)`

GetSmtpSecurityOk returns a tuple with the SmtpSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpSecurity

`func (o *ExternalServerListResponseResourcesInner) SetSmtpSecurity(v EnumexternalServerSmtpSecurityProp)`

SetSmtpSecurity sets SmtpSecurity field to given value.

### HasSmtpSecurity

`func (o *ExternalServerListResponseResourcesInner) HasSmtpSecurity() bool`

HasSmtpSecurity returns a boolean if a field has been set.

### GetSmtpTimeout

`func (o *ExternalServerListResponseResourcesInner) GetSmtpTimeout() string`

GetSmtpTimeout returns the SmtpTimeout field if non-nil, zero value otherwise.

### GetSmtpTimeoutOk

`func (o *ExternalServerListResponseResourcesInner) GetSmtpTimeoutOk() (*string, bool)`

GetSmtpTimeoutOk returns a tuple with the SmtpTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpTimeout

`func (o *ExternalServerListResponseResourcesInner) SetSmtpTimeout(v string)`

SetSmtpTimeout sets SmtpTimeout field to given value.

### HasSmtpTimeout

`func (o *ExternalServerListResponseResourcesInner) HasSmtpTimeout() bool`

HasSmtpTimeout returns a boolean if a field has been set.

### GetSmtpConnectionProperties

`func (o *ExternalServerListResponseResourcesInner) GetSmtpConnectionProperties() []string`

GetSmtpConnectionProperties returns the SmtpConnectionProperties field if non-nil, zero value otherwise.

### GetSmtpConnectionPropertiesOk

`func (o *ExternalServerListResponseResourcesInner) GetSmtpConnectionPropertiesOk() (*[]string, bool)`

GetSmtpConnectionPropertiesOk returns a tuple with the SmtpConnectionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpConnectionProperties

`func (o *ExternalServerListResponseResourcesInner) SetSmtpConnectionProperties(v []string)`

SetSmtpConnectionProperties sets SmtpConnectionProperties field to given value.

### HasSmtpConnectionProperties

`func (o *ExternalServerListResponseResourcesInner) HasSmtpConnectionProperties() bool`

HasSmtpConnectionProperties returns a boolean if a field has been set.

### GetServerHTTPPort

`func (o *ExternalServerListResponseResourcesInner) GetServerHTTPPort() int64`

GetServerHTTPPort returns the ServerHTTPPort field if non-nil, zero value otherwise.

### GetServerHTTPPortOk

`func (o *ExternalServerListResponseResourcesInner) GetServerHTTPPortOk() (*int64, bool)`

GetServerHTTPPortOk returns a tuple with the ServerHTTPPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHTTPPort

`func (o *ExternalServerListResponseResourcesInner) SetServerHTTPPort(v int64)`

SetServerHTTPPort sets ServerHTTPPort field to given value.


### GetUseSSL

`func (o *ExternalServerListResponseResourcesInner) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *ExternalServerListResponseResourcesInner) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *ExternalServerListResponseResourcesInner) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *ExternalServerListResponseResourcesInner) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetBasicAuthenticationUsername

`func (o *ExternalServerListResponseResourcesInner) GetBasicAuthenticationUsername() string`

GetBasicAuthenticationUsername returns the BasicAuthenticationUsername field if non-nil, zero value otherwise.

### GetBasicAuthenticationUsernameOk

`func (o *ExternalServerListResponseResourcesInner) GetBasicAuthenticationUsernameOk() (*string, bool)`

GetBasicAuthenticationUsernameOk returns a tuple with the BasicAuthenticationUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthenticationUsername

`func (o *ExternalServerListResponseResourcesInner) SetBasicAuthenticationUsername(v string)`

SetBasicAuthenticationUsername sets BasicAuthenticationUsername field to given value.

### HasBasicAuthenticationUsername

`func (o *ExternalServerListResponseResourcesInner) HasBasicAuthenticationUsername() bool`

HasBasicAuthenticationUsername returns a boolean if a field has been set.

### GetBasicAuthenticationPassphraseProvider

`func (o *ExternalServerListResponseResourcesInner) GetBasicAuthenticationPassphraseProvider() string`

GetBasicAuthenticationPassphraseProvider returns the BasicAuthenticationPassphraseProvider field if non-nil, zero value otherwise.

### GetBasicAuthenticationPassphraseProviderOk

`func (o *ExternalServerListResponseResourcesInner) GetBasicAuthenticationPassphraseProviderOk() (*string, bool)`

GetBasicAuthenticationPassphraseProviderOk returns a tuple with the BasicAuthenticationPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthenticationPassphraseProvider

`func (o *ExternalServerListResponseResourcesInner) SetBasicAuthenticationPassphraseProvider(v string)`

SetBasicAuthenticationPassphraseProvider sets BasicAuthenticationPassphraseProvider field to given value.

### HasBasicAuthenticationPassphraseProvider

`func (o *ExternalServerListResponseResourcesInner) HasBasicAuthenticationPassphraseProvider() bool`

HasBasicAuthenticationPassphraseProvider returns a boolean if a field has been set.

### GetBootstrapServer

`func (o *ExternalServerListResponseResourcesInner) GetBootstrapServer() []string`

GetBootstrapServer returns the BootstrapServer field if non-nil, zero value otherwise.

### GetBootstrapServerOk

`func (o *ExternalServerListResponseResourcesInner) GetBootstrapServerOk() (*[]string, bool)`

GetBootstrapServerOk returns a tuple with the BootstrapServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapServer

`func (o *ExternalServerListResponseResourcesInner) SetBootstrapServer(v []string)`

SetBootstrapServer sets BootstrapServer field to given value.


### GetProducerProperty

`func (o *ExternalServerListResponseResourcesInner) GetProducerProperty() []string`

GetProducerProperty returns the ProducerProperty field if non-nil, zero value otherwise.

### GetProducerPropertyOk

`func (o *ExternalServerListResponseResourcesInner) GetProducerPropertyOk() (*[]string, bool)`

GetProducerPropertyOk returns a tuple with the ProducerProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerProperty

`func (o *ExternalServerListResponseResourcesInner) SetProducerProperty(v []string)`

SetProducerProperty sets ProducerProperty field to given value.

### HasProducerProperty

`func (o *ExternalServerListResponseResourcesInner) HasProducerProperty() bool`

HasProducerProperty returns a boolean if a field has been set.

### GetBaseDN

`func (o *ExternalServerListResponseResourcesInner) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *ExternalServerListResponseResourcesInner) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *ExternalServerListResponseResourcesInner) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.


### GetConjurServerBaseURI

`func (o *ExternalServerListResponseResourcesInner) GetConjurServerBaseURI() []string`

GetConjurServerBaseURI returns the ConjurServerBaseURI field if non-nil, zero value otherwise.

### GetConjurServerBaseURIOk

`func (o *ExternalServerListResponseResourcesInner) GetConjurServerBaseURIOk() (*[]string, bool)`

GetConjurServerBaseURIOk returns a tuple with the ConjurServerBaseURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurServerBaseURI

`func (o *ExternalServerListResponseResourcesInner) SetConjurServerBaseURI(v []string)`

SetConjurServerBaseURI sets ConjurServerBaseURI field to given value.


### GetConjurAuthenticationMethod

`func (o *ExternalServerListResponseResourcesInner) GetConjurAuthenticationMethod() string`

GetConjurAuthenticationMethod returns the ConjurAuthenticationMethod field if non-nil, zero value otherwise.

### GetConjurAuthenticationMethodOk

`func (o *ExternalServerListResponseResourcesInner) GetConjurAuthenticationMethodOk() (*string, bool)`

GetConjurAuthenticationMethodOk returns a tuple with the ConjurAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurAuthenticationMethod

`func (o *ExternalServerListResponseResourcesInner) SetConjurAuthenticationMethod(v string)`

SetConjurAuthenticationMethod sets ConjurAuthenticationMethod field to given value.


### GetConjurAccountName

`func (o *ExternalServerListResponseResourcesInner) GetConjurAccountName() string`

GetConjurAccountName returns the ConjurAccountName field if non-nil, zero value otherwise.

### GetConjurAccountNameOk

`func (o *ExternalServerListResponseResourcesInner) GetConjurAccountNameOk() (*string, bool)`

GetConjurAccountNameOk returns a tuple with the ConjurAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjurAccountName

`func (o *ExternalServerListResponseResourcesInner) SetConjurAccountName(v string)`

SetConjurAccountName sets ConjurAccountName field to given value.


### GetAwsAccessKeyID

`func (o *ExternalServerListResponseResourcesInner) GetAwsAccessKeyID() string`

GetAwsAccessKeyID returns the AwsAccessKeyID field if non-nil, zero value otherwise.

### GetAwsAccessKeyIDOk

`func (o *ExternalServerListResponseResourcesInner) GetAwsAccessKeyIDOk() (*string, bool)`

GetAwsAccessKeyIDOk returns a tuple with the AwsAccessKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyID

`func (o *ExternalServerListResponseResourcesInner) SetAwsAccessKeyID(v string)`

SetAwsAccessKeyID sets AwsAccessKeyID field to given value.

### HasAwsAccessKeyID

`func (o *ExternalServerListResponseResourcesInner) HasAwsAccessKeyID() bool`

HasAwsAccessKeyID returns a boolean if a field has been set.

### GetAwsSecretAccessKey

`func (o *ExternalServerListResponseResourcesInner) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *ExternalServerListResponseResourcesInner) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *ExternalServerListResponseResourcesInner) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *ExternalServerListResponseResourcesInner) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### GetAwsRegionName

`func (o *ExternalServerListResponseResourcesInner) GetAwsRegionName() string`

GetAwsRegionName returns the AwsRegionName field if non-nil, zero value otherwise.

### GetAwsRegionNameOk

`func (o *ExternalServerListResponseResourcesInner) GetAwsRegionNameOk() (*string, bool)`

GetAwsRegionNameOk returns a tuple with the AwsRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegionName

`func (o *ExternalServerListResponseResourcesInner) SetAwsRegionName(v string)`

SetAwsRegionName sets AwsRegionName field to given value.


### GetHttpAuthorizationMethod

`func (o *ExternalServerListResponseResourcesInner) GetHttpAuthorizationMethod() string`

GetHttpAuthorizationMethod returns the HttpAuthorizationMethod field if non-nil, zero value otherwise.

### GetHttpAuthorizationMethodOk

`func (o *ExternalServerListResponseResourcesInner) GetHttpAuthorizationMethodOk() (*string, bool)`

GetHttpAuthorizationMethodOk returns a tuple with the HttpAuthorizationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAuthorizationMethod

`func (o *ExternalServerListResponseResourcesInner) SetHttpAuthorizationMethod(v string)`

SetHttpAuthorizationMethod sets HttpAuthorizationMethod field to given value.


### GetClientReconnectInterval

`func (o *ExternalServerListResponseResourcesInner) GetClientReconnectInterval() string`

GetClientReconnectInterval returns the ClientReconnectInterval field if non-nil, zero value otherwise.

### GetClientReconnectIntervalOk

`func (o *ExternalServerListResponseResourcesInner) GetClientReconnectIntervalOk() (*string, bool)`

GetClientReconnectIntervalOk returns a tuple with the ClientReconnectInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReconnectInterval

`func (o *ExternalServerListResponseResourcesInner) SetClientReconnectInterval(v string)`

SetClientReconnectInterval sets ClientReconnectInterval field to given value.

### HasClientReconnectInterval

`func (o *ExternalServerListResponseResourcesInner) HasClientReconnectInterval() bool`

HasClientReconnectInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


