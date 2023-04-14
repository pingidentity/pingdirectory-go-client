# AddConnectionHandler200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Connection Handler | 
**Schemas** | [**[]EnumhttpConnectionHandlerSchemaUrn**](EnumhttpConnectionHandlerSchemaUrn.md) |  | 
**ListenPort** | **int64** | Specifies the port number on which the HTTP Connection Handler will listen for connections from clients. | 
**UseSSL** | Pointer to **bool** | Indicates whether the HTTP Connection Handler should use SSL. | [optional] 
**SslCertNickname** | Pointer to **string** | Specifies the nickname (also called the alias) of the certificate that the HTTP Connection Handler should use when performing SSL communication. | [optional] 
**KeyManagerProvider** | Pointer to **string** | Specifies the key manager provider that will be used to obtain the certificate to present to HTTPS clients. | [optional] 
**Description** | Pointer to **string** | A description for this Connection Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Connection Handler is enabled. | 
**AllowedClient** | Pointer to **[]string** | Specifies a set of address masks that determines the addresses of the clients that are allowed to establish connections to this connection handler. | [optional] 
**DeniedClient** | Pointer to **[]string** | Specifies a set of address masks that determines the addresses of the clients that are not allowed to establish connections to this connection handler. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**ListenAddress** | Pointer to **string** | Specifies the address on which to listen for connections from HTTP clients. If no value is defined, the server will listen on all addresses on all interfaces. | [optional] 
**AllowStartTLS** | Pointer to **bool** | Indicates whether clients are allowed to use StartTLS. | [optional] 
**TrustManagerProvider** | Pointer to **string** | Specifies the trust manager provider that will be used to validate any certificates presented by HTTPS clients. | [optional] 
**AllowLDAPV2** | Pointer to **bool** | Indicates whether connections from LDAPv2 clients are allowed. | [optional] 
**UseTCPKeepAlive** | Pointer to **bool** | Indicates whether the LDAP Connection Handler should use TCP keep-alive. | [optional] 
**SendRejectionNotice** | Pointer to **bool** | Indicates whether the LDAP Connection Handler should send a notice of disconnection extended response message to the client if a new connection is rejected for some reason. | [optional] 
**FailedBindResponseDelay** | Pointer to **string** | Specifies the length of time that the server should delay the response to non-successful bind operations. A value of zero milliseconds indicates that non-successful bind operations should not be delayed. | [optional] 
**MaxRequestSize** | Pointer to **string** | Specifies the size of the largest LDAP request message that will be allowed by this LDAP Connection handler. | [optional] 
**MaxCancelHandlers** | Pointer to **int64** | Specifies the maximum number of threads that are used to process cancel and abandon requests from clients. | [optional] 
**NumAcceptHandlers** | Pointer to **int64** | Specifies the number of threads that are used to accept new client connections, and to perform any initial preparation on those connections that may be needed before the connection can be used to read requests and send responses. | [optional] 
**NumRequestHandlers** | Pointer to **int64** | Specifies the number of threads that will be used for accepting connections and reading requests from clients. | [optional] 
**SslClientAuthPolicy** | Pointer to [**EnumconnectionHandlerSslClientAuthPolicyProp**](EnumconnectionHandlerSslClientAuthPolicyProp.md) |  | [optional] 
**AcceptBacklog** | Pointer to **int64** | Specifies the number of concurrent outstanding connection attempts that the connection handler should allow. The default value should be acceptable in most cases, but it may need to be increased in environments that may attempt to establish large numbers of connections simultaneously. | [optional] 
**SslProtocol** | Pointer to **[]string** | Specifies the names of the SSL protocols that are allowed for use in SSL communication. The set of supported ssl protocols can be viewed via the ssl context monitor entry. | [optional] 
**SslCipherSuite** | Pointer to **[]string** | Specifies the names of the SSL cipher suites that are allowed for use in SSL communication. The set of supported cipher suites can be viewed via the ssl context monitor entry. | [optional] 
**MaxBlockedWriteTimeLimit** | Pointer to **string** | Specifies the maximum length of time that attempts to write data to LDAP clients should be allowed to block. | [optional] 
**AutoAuthenticateUsingClientCertificate** | Pointer to **bool** | Indicates whether to attempt to automatically authenticate a client connection that has established a secure communication channel (using either SSL or StartTLS) and presented its own client certificate. Generally, clients should use the SASL EXTERNAL mechanism to authenticate using a client certificate, but some clients may not support that capability and/or may expect automatic authentication. | [optional] 
**CloseConnectionsWhenUnavailable** | Pointer to **bool** | Indicates whether all connections associated with this LDAP Connection Handler should be closed and no new connections accepted when the server has determined that it is \&quot;unavailable.\&quot; This allows clients (or a network load balancer) to route requests to another server. | [optional] 
**CloseConnectionsOnExplicitGC** | Pointer to **bool** | Indicates whether all connections associated with this LDAP Connection Handler should be closed before an explicit garbage collection is performed to allow clients to route requests to another server. | [optional] 
**LdifDirectory** | **string** | Specifies the path to the directory in which the LDIF files should be placed. | 
**PollInterval** | **string** | Specifies how frequently the LDIF connection handler should check the LDIF directory to determine whether a new LDIF file has been added. | 
**HttpServletExtension** | Pointer to **[]string** | Specifies information about servlets that will be provided via this connection handler. | [optional] 
**WebApplicationExtension** | Pointer to **[]string** | Specifies information about web applications that will be provided via this connection handler. | [optional] 
**HttpOperationLogPublisher** | Pointer to **[]string** | Specifies the set of HTTP operation loggers that should be used to log information about requests and responses for operations processed through this HTTP Connection Handler. | [optional] 
**KeepStats** | Pointer to **bool** | Indicates whether to enable statistics collection for this connection handler. | [optional] 
**AllowTCPReuseAddress** | Pointer to **bool** | Indicates whether the server should attempt to reuse socket descriptors. This may be useful in environments with a high rate of connection establishment and termination. | [optional] 
**IdleTimeLimit** | Pointer to **string** | Specifies the maximum idle time for a connection. The max idle time is applied when waiting for a new request to be received on a connection, when reading the headers and content of a request, or when writing the headers and content of a response. | [optional] 
**LowResourcesConnectionThreshold** | Pointer to **int64** | Specifies the number of connections, which if exceeded, places this handler in a low resource state where a different idle time limit is applied on the connections. | [optional] 
**LowResourcesIdleTimeLimit** | Pointer to **string** | Specifies the maximum idle time for a connection when this handler is in a low resource state as defined by low-resource-connections. The max idle time is applied when waiting for a new request to be received on a connection, when reading the headers and content of a request, or when writing the headers and content of a response. | [optional] 
**EnableMultipartMIMEParameters** | Pointer to **bool** | Determines whether request form parameters submitted in multipart/ form-data (RFC 2388) format should be processed as request parameters. | [optional] 
**UseForwardedHeaders** | Pointer to **bool** | Indicates whether to use \&quot;Forwarded\&quot; and \&quot;X-Forwarded-*\&quot; request headers to override corresponding HTTP request information available during request processing. | [optional] 
**HttpRequestHeaderSize** | Pointer to **int64** | Specifies the maximum buffer size of an http request including the request uri and all of the request headers. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**UseCorrelationIDHeader** | Pointer to **bool** | If enabled, a correlation ID header will be added to outgoing HTTP responses. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 
**CorrelationIDRequestHeader** | Pointer to **[]string** | Specifies the set of HTTP request headers that may contain a value to be used as the correlation ID. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 

## Methods

### NewAddConnectionHandler200Response

`func NewAddConnectionHandler200Response(id string, schemas []EnumhttpConnectionHandlerSchemaUrn, listenPort int64, enabled bool, ldifDirectory string, pollInterval string, ) *AddConnectionHandler200Response`

NewAddConnectionHandler200Response instantiates a new AddConnectionHandler200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddConnectionHandler200ResponseWithDefaults

`func NewAddConnectionHandler200ResponseWithDefaults() *AddConnectionHandler200Response`

NewAddConnectionHandler200ResponseWithDefaults instantiates a new AddConnectionHandler200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddConnectionHandler200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddConnectionHandler200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddConnectionHandler200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddConnectionHandler200Response) GetSchemas() []EnumhttpConnectionHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddConnectionHandler200Response) GetSchemasOk() (*[]EnumhttpConnectionHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddConnectionHandler200Response) SetSchemas(v []EnumhttpConnectionHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetListenPort

`func (o *AddConnectionHandler200Response) GetListenPort() int64`

GetListenPort returns the ListenPort field if non-nil, zero value otherwise.

### GetListenPortOk

`func (o *AddConnectionHandler200Response) GetListenPortOk() (*int64, bool)`

GetListenPortOk returns a tuple with the ListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPort

`func (o *AddConnectionHandler200Response) SetListenPort(v int64)`

SetListenPort sets ListenPort field to given value.


### GetUseSSL

`func (o *AddConnectionHandler200Response) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *AddConnectionHandler200Response) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *AddConnectionHandler200Response) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *AddConnectionHandler200Response) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *AddConnectionHandler200Response) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *AddConnectionHandler200Response) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *AddConnectionHandler200Response) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.

### HasSslCertNickname

`func (o *AddConnectionHandler200Response) HasSslCertNickname() bool`

HasSslCertNickname returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *AddConnectionHandler200Response) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *AddConnectionHandler200Response) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *AddConnectionHandler200Response) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *AddConnectionHandler200Response) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetDescription

`func (o *AddConnectionHandler200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddConnectionHandler200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddConnectionHandler200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddConnectionHandler200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddConnectionHandler200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddConnectionHandler200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddConnectionHandler200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowedClient

`func (o *AddConnectionHandler200Response) GetAllowedClient() []string`

GetAllowedClient returns the AllowedClient field if non-nil, zero value otherwise.

### GetAllowedClientOk

`func (o *AddConnectionHandler200Response) GetAllowedClientOk() (*[]string, bool)`

GetAllowedClientOk returns a tuple with the AllowedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClient

`func (o *AddConnectionHandler200Response) SetAllowedClient(v []string)`

SetAllowedClient sets AllowedClient field to given value.

### HasAllowedClient

`func (o *AddConnectionHandler200Response) HasAllowedClient() bool`

HasAllowedClient returns a boolean if a field has been set.

### GetDeniedClient

`func (o *AddConnectionHandler200Response) GetDeniedClient() []string`

GetDeniedClient returns the DeniedClient field if non-nil, zero value otherwise.

### GetDeniedClientOk

`func (o *AddConnectionHandler200Response) GetDeniedClientOk() (*[]string, bool)`

GetDeniedClientOk returns a tuple with the DeniedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedClient

`func (o *AddConnectionHandler200Response) SetDeniedClient(v []string)`

SetDeniedClient sets DeniedClient field to given value.

### HasDeniedClient

`func (o *AddConnectionHandler200Response) HasDeniedClient() bool`

HasDeniedClient returns a boolean if a field has been set.

### GetMeta

`func (o *AddConnectionHandler200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddConnectionHandler200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddConnectionHandler200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddConnectionHandler200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddConnectionHandler200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddConnectionHandler200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddConnectionHandler200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddConnectionHandler200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetListenAddress

`func (o *AddConnectionHandler200Response) GetListenAddress() string`

GetListenAddress returns the ListenAddress field if non-nil, zero value otherwise.

### GetListenAddressOk

`func (o *AddConnectionHandler200Response) GetListenAddressOk() (*string, bool)`

GetListenAddressOk returns a tuple with the ListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddress

`func (o *AddConnectionHandler200Response) SetListenAddress(v string)`

SetListenAddress sets ListenAddress field to given value.

### HasListenAddress

`func (o *AddConnectionHandler200Response) HasListenAddress() bool`

HasListenAddress returns a boolean if a field has been set.

### GetAllowStartTLS

`func (o *AddConnectionHandler200Response) GetAllowStartTLS() bool`

GetAllowStartTLS returns the AllowStartTLS field if non-nil, zero value otherwise.

### GetAllowStartTLSOk

`func (o *AddConnectionHandler200Response) GetAllowStartTLSOk() (*bool, bool)`

GetAllowStartTLSOk returns a tuple with the AllowStartTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStartTLS

`func (o *AddConnectionHandler200Response) SetAllowStartTLS(v bool)`

SetAllowStartTLS sets AllowStartTLS field to given value.

### HasAllowStartTLS

`func (o *AddConnectionHandler200Response) HasAllowStartTLS() bool`

HasAllowStartTLS returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *AddConnectionHandler200Response) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *AddConnectionHandler200Response) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *AddConnectionHandler200Response) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *AddConnectionHandler200Response) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetAllowLDAPV2

`func (o *AddConnectionHandler200Response) GetAllowLDAPV2() bool`

GetAllowLDAPV2 returns the AllowLDAPV2 field if non-nil, zero value otherwise.

### GetAllowLDAPV2Ok

`func (o *AddConnectionHandler200Response) GetAllowLDAPV2Ok() (*bool, bool)`

GetAllowLDAPV2Ok returns a tuple with the AllowLDAPV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLDAPV2

`func (o *AddConnectionHandler200Response) SetAllowLDAPV2(v bool)`

SetAllowLDAPV2 sets AllowLDAPV2 field to given value.

### HasAllowLDAPV2

`func (o *AddConnectionHandler200Response) HasAllowLDAPV2() bool`

HasAllowLDAPV2 returns a boolean if a field has been set.

### GetUseTCPKeepAlive

`func (o *AddConnectionHandler200Response) GetUseTCPKeepAlive() bool`

GetUseTCPKeepAlive returns the UseTCPKeepAlive field if non-nil, zero value otherwise.

### GetUseTCPKeepAliveOk

`func (o *AddConnectionHandler200Response) GetUseTCPKeepAliveOk() (*bool, bool)`

GetUseTCPKeepAliveOk returns a tuple with the UseTCPKeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTCPKeepAlive

`func (o *AddConnectionHandler200Response) SetUseTCPKeepAlive(v bool)`

SetUseTCPKeepAlive sets UseTCPKeepAlive field to given value.

### HasUseTCPKeepAlive

`func (o *AddConnectionHandler200Response) HasUseTCPKeepAlive() bool`

HasUseTCPKeepAlive returns a boolean if a field has been set.

### GetSendRejectionNotice

`func (o *AddConnectionHandler200Response) GetSendRejectionNotice() bool`

GetSendRejectionNotice returns the SendRejectionNotice field if non-nil, zero value otherwise.

### GetSendRejectionNoticeOk

`func (o *AddConnectionHandler200Response) GetSendRejectionNoticeOk() (*bool, bool)`

GetSendRejectionNoticeOk returns a tuple with the SendRejectionNotice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendRejectionNotice

`func (o *AddConnectionHandler200Response) SetSendRejectionNotice(v bool)`

SetSendRejectionNotice sets SendRejectionNotice field to given value.

### HasSendRejectionNotice

`func (o *AddConnectionHandler200Response) HasSendRejectionNotice() bool`

HasSendRejectionNotice returns a boolean if a field has been set.

### GetFailedBindResponseDelay

`func (o *AddConnectionHandler200Response) GetFailedBindResponseDelay() string`

GetFailedBindResponseDelay returns the FailedBindResponseDelay field if non-nil, zero value otherwise.

### GetFailedBindResponseDelayOk

`func (o *AddConnectionHandler200Response) GetFailedBindResponseDelayOk() (*string, bool)`

GetFailedBindResponseDelayOk returns a tuple with the FailedBindResponseDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedBindResponseDelay

`func (o *AddConnectionHandler200Response) SetFailedBindResponseDelay(v string)`

SetFailedBindResponseDelay sets FailedBindResponseDelay field to given value.

### HasFailedBindResponseDelay

`func (o *AddConnectionHandler200Response) HasFailedBindResponseDelay() bool`

HasFailedBindResponseDelay returns a boolean if a field has been set.

### GetMaxRequestSize

`func (o *AddConnectionHandler200Response) GetMaxRequestSize() string`

GetMaxRequestSize returns the MaxRequestSize field if non-nil, zero value otherwise.

### GetMaxRequestSizeOk

`func (o *AddConnectionHandler200Response) GetMaxRequestSizeOk() (*string, bool)`

GetMaxRequestSizeOk returns a tuple with the MaxRequestSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequestSize

`func (o *AddConnectionHandler200Response) SetMaxRequestSize(v string)`

SetMaxRequestSize sets MaxRequestSize field to given value.

### HasMaxRequestSize

`func (o *AddConnectionHandler200Response) HasMaxRequestSize() bool`

HasMaxRequestSize returns a boolean if a field has been set.

### GetMaxCancelHandlers

`func (o *AddConnectionHandler200Response) GetMaxCancelHandlers() int64`

GetMaxCancelHandlers returns the MaxCancelHandlers field if non-nil, zero value otherwise.

### GetMaxCancelHandlersOk

`func (o *AddConnectionHandler200Response) GetMaxCancelHandlersOk() (*int64, bool)`

GetMaxCancelHandlersOk returns a tuple with the MaxCancelHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCancelHandlers

`func (o *AddConnectionHandler200Response) SetMaxCancelHandlers(v int64)`

SetMaxCancelHandlers sets MaxCancelHandlers field to given value.

### HasMaxCancelHandlers

`func (o *AddConnectionHandler200Response) HasMaxCancelHandlers() bool`

HasMaxCancelHandlers returns a boolean if a field has been set.

### GetNumAcceptHandlers

`func (o *AddConnectionHandler200Response) GetNumAcceptHandlers() int64`

GetNumAcceptHandlers returns the NumAcceptHandlers field if non-nil, zero value otherwise.

### GetNumAcceptHandlersOk

`func (o *AddConnectionHandler200Response) GetNumAcceptHandlersOk() (*int64, bool)`

GetNumAcceptHandlersOk returns a tuple with the NumAcceptHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAcceptHandlers

`func (o *AddConnectionHandler200Response) SetNumAcceptHandlers(v int64)`

SetNumAcceptHandlers sets NumAcceptHandlers field to given value.

### HasNumAcceptHandlers

`func (o *AddConnectionHandler200Response) HasNumAcceptHandlers() bool`

HasNumAcceptHandlers returns a boolean if a field has been set.

### GetNumRequestHandlers

`func (o *AddConnectionHandler200Response) GetNumRequestHandlers() int64`

GetNumRequestHandlers returns the NumRequestHandlers field if non-nil, zero value otherwise.

### GetNumRequestHandlersOk

`func (o *AddConnectionHandler200Response) GetNumRequestHandlersOk() (*int64, bool)`

GetNumRequestHandlersOk returns a tuple with the NumRequestHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRequestHandlers

`func (o *AddConnectionHandler200Response) SetNumRequestHandlers(v int64)`

SetNumRequestHandlers sets NumRequestHandlers field to given value.

### HasNumRequestHandlers

`func (o *AddConnectionHandler200Response) HasNumRequestHandlers() bool`

HasNumRequestHandlers returns a boolean if a field has been set.

### GetSslClientAuthPolicy

`func (o *AddConnectionHandler200Response) GetSslClientAuthPolicy() EnumconnectionHandlerSslClientAuthPolicyProp`

GetSslClientAuthPolicy returns the SslClientAuthPolicy field if non-nil, zero value otherwise.

### GetSslClientAuthPolicyOk

`func (o *AddConnectionHandler200Response) GetSslClientAuthPolicyOk() (*EnumconnectionHandlerSslClientAuthPolicyProp, bool)`

GetSslClientAuthPolicyOk returns a tuple with the SslClientAuthPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslClientAuthPolicy

`func (o *AddConnectionHandler200Response) SetSslClientAuthPolicy(v EnumconnectionHandlerSslClientAuthPolicyProp)`

SetSslClientAuthPolicy sets SslClientAuthPolicy field to given value.

### HasSslClientAuthPolicy

`func (o *AddConnectionHandler200Response) HasSslClientAuthPolicy() bool`

HasSslClientAuthPolicy returns a boolean if a field has been set.

### GetAcceptBacklog

`func (o *AddConnectionHandler200Response) GetAcceptBacklog() int64`

GetAcceptBacklog returns the AcceptBacklog field if non-nil, zero value otherwise.

### GetAcceptBacklogOk

`func (o *AddConnectionHandler200Response) GetAcceptBacklogOk() (*int64, bool)`

GetAcceptBacklogOk returns a tuple with the AcceptBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptBacklog

`func (o *AddConnectionHandler200Response) SetAcceptBacklog(v int64)`

SetAcceptBacklog sets AcceptBacklog field to given value.

### HasAcceptBacklog

`func (o *AddConnectionHandler200Response) HasAcceptBacklog() bool`

HasAcceptBacklog returns a boolean if a field has been set.

### GetSslProtocol

`func (o *AddConnectionHandler200Response) GetSslProtocol() []string`

GetSslProtocol returns the SslProtocol field if non-nil, zero value otherwise.

### GetSslProtocolOk

`func (o *AddConnectionHandler200Response) GetSslProtocolOk() (*[]string, bool)`

GetSslProtocolOk returns a tuple with the SslProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslProtocol

`func (o *AddConnectionHandler200Response) SetSslProtocol(v []string)`

SetSslProtocol sets SslProtocol field to given value.

### HasSslProtocol

`func (o *AddConnectionHandler200Response) HasSslProtocol() bool`

HasSslProtocol returns a boolean if a field has been set.

### GetSslCipherSuite

`func (o *AddConnectionHandler200Response) GetSslCipherSuite() []string`

GetSslCipherSuite returns the SslCipherSuite field if non-nil, zero value otherwise.

### GetSslCipherSuiteOk

`func (o *AddConnectionHandler200Response) GetSslCipherSuiteOk() (*[]string, bool)`

GetSslCipherSuiteOk returns a tuple with the SslCipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCipherSuite

`func (o *AddConnectionHandler200Response) SetSslCipherSuite(v []string)`

SetSslCipherSuite sets SslCipherSuite field to given value.

### HasSslCipherSuite

`func (o *AddConnectionHandler200Response) HasSslCipherSuite() bool`

HasSslCipherSuite returns a boolean if a field has been set.

### GetMaxBlockedWriteTimeLimit

`func (o *AddConnectionHandler200Response) GetMaxBlockedWriteTimeLimit() string`

GetMaxBlockedWriteTimeLimit returns the MaxBlockedWriteTimeLimit field if non-nil, zero value otherwise.

### GetMaxBlockedWriteTimeLimitOk

`func (o *AddConnectionHandler200Response) GetMaxBlockedWriteTimeLimitOk() (*string, bool)`

GetMaxBlockedWriteTimeLimitOk returns a tuple with the MaxBlockedWriteTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBlockedWriteTimeLimit

`func (o *AddConnectionHandler200Response) SetMaxBlockedWriteTimeLimit(v string)`

SetMaxBlockedWriteTimeLimit sets MaxBlockedWriteTimeLimit field to given value.

### HasMaxBlockedWriteTimeLimit

`func (o *AddConnectionHandler200Response) HasMaxBlockedWriteTimeLimit() bool`

HasMaxBlockedWriteTimeLimit returns a boolean if a field has been set.

### GetAutoAuthenticateUsingClientCertificate

`func (o *AddConnectionHandler200Response) GetAutoAuthenticateUsingClientCertificate() bool`

GetAutoAuthenticateUsingClientCertificate returns the AutoAuthenticateUsingClientCertificate field if non-nil, zero value otherwise.

### GetAutoAuthenticateUsingClientCertificateOk

`func (o *AddConnectionHandler200Response) GetAutoAuthenticateUsingClientCertificateOk() (*bool, bool)`

GetAutoAuthenticateUsingClientCertificateOk returns a tuple with the AutoAuthenticateUsingClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAuthenticateUsingClientCertificate

`func (o *AddConnectionHandler200Response) SetAutoAuthenticateUsingClientCertificate(v bool)`

SetAutoAuthenticateUsingClientCertificate sets AutoAuthenticateUsingClientCertificate field to given value.

### HasAutoAuthenticateUsingClientCertificate

`func (o *AddConnectionHandler200Response) HasAutoAuthenticateUsingClientCertificate() bool`

HasAutoAuthenticateUsingClientCertificate returns a boolean if a field has been set.

### GetCloseConnectionsWhenUnavailable

`func (o *AddConnectionHandler200Response) GetCloseConnectionsWhenUnavailable() bool`

GetCloseConnectionsWhenUnavailable returns the CloseConnectionsWhenUnavailable field if non-nil, zero value otherwise.

### GetCloseConnectionsWhenUnavailableOk

`func (o *AddConnectionHandler200Response) GetCloseConnectionsWhenUnavailableOk() (*bool, bool)`

GetCloseConnectionsWhenUnavailableOk returns a tuple with the CloseConnectionsWhenUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseConnectionsWhenUnavailable

`func (o *AddConnectionHandler200Response) SetCloseConnectionsWhenUnavailable(v bool)`

SetCloseConnectionsWhenUnavailable sets CloseConnectionsWhenUnavailable field to given value.

### HasCloseConnectionsWhenUnavailable

`func (o *AddConnectionHandler200Response) HasCloseConnectionsWhenUnavailable() bool`

HasCloseConnectionsWhenUnavailable returns a boolean if a field has been set.

### GetCloseConnectionsOnExplicitGC

`func (o *AddConnectionHandler200Response) GetCloseConnectionsOnExplicitGC() bool`

GetCloseConnectionsOnExplicitGC returns the CloseConnectionsOnExplicitGC field if non-nil, zero value otherwise.

### GetCloseConnectionsOnExplicitGCOk

`func (o *AddConnectionHandler200Response) GetCloseConnectionsOnExplicitGCOk() (*bool, bool)`

GetCloseConnectionsOnExplicitGCOk returns a tuple with the CloseConnectionsOnExplicitGC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseConnectionsOnExplicitGC

`func (o *AddConnectionHandler200Response) SetCloseConnectionsOnExplicitGC(v bool)`

SetCloseConnectionsOnExplicitGC sets CloseConnectionsOnExplicitGC field to given value.

### HasCloseConnectionsOnExplicitGC

`func (o *AddConnectionHandler200Response) HasCloseConnectionsOnExplicitGC() bool`

HasCloseConnectionsOnExplicitGC returns a boolean if a field has been set.

### GetLdifDirectory

`func (o *AddConnectionHandler200Response) GetLdifDirectory() string`

GetLdifDirectory returns the LdifDirectory field if non-nil, zero value otherwise.

### GetLdifDirectoryOk

`func (o *AddConnectionHandler200Response) GetLdifDirectoryOk() (*string, bool)`

GetLdifDirectoryOk returns a tuple with the LdifDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdifDirectory

`func (o *AddConnectionHandler200Response) SetLdifDirectory(v string)`

SetLdifDirectory sets LdifDirectory field to given value.


### GetPollInterval

`func (o *AddConnectionHandler200Response) GetPollInterval() string`

GetPollInterval returns the PollInterval field if non-nil, zero value otherwise.

### GetPollIntervalOk

`func (o *AddConnectionHandler200Response) GetPollIntervalOk() (*string, bool)`

GetPollIntervalOk returns a tuple with the PollInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollInterval

`func (o *AddConnectionHandler200Response) SetPollInterval(v string)`

SetPollInterval sets PollInterval field to given value.


### GetHttpServletExtension

`func (o *AddConnectionHandler200Response) GetHttpServletExtension() []string`

GetHttpServletExtension returns the HttpServletExtension field if non-nil, zero value otherwise.

### GetHttpServletExtensionOk

`func (o *AddConnectionHandler200Response) GetHttpServletExtensionOk() (*[]string, bool)`

GetHttpServletExtensionOk returns a tuple with the HttpServletExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServletExtension

`func (o *AddConnectionHandler200Response) SetHttpServletExtension(v []string)`

SetHttpServletExtension sets HttpServletExtension field to given value.

### HasHttpServletExtension

`func (o *AddConnectionHandler200Response) HasHttpServletExtension() bool`

HasHttpServletExtension returns a boolean if a field has been set.

### GetWebApplicationExtension

`func (o *AddConnectionHandler200Response) GetWebApplicationExtension() []string`

GetWebApplicationExtension returns the WebApplicationExtension field if non-nil, zero value otherwise.

### GetWebApplicationExtensionOk

`func (o *AddConnectionHandler200Response) GetWebApplicationExtensionOk() (*[]string, bool)`

GetWebApplicationExtensionOk returns a tuple with the WebApplicationExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebApplicationExtension

`func (o *AddConnectionHandler200Response) SetWebApplicationExtension(v []string)`

SetWebApplicationExtension sets WebApplicationExtension field to given value.

### HasWebApplicationExtension

`func (o *AddConnectionHandler200Response) HasWebApplicationExtension() bool`

HasWebApplicationExtension returns a boolean if a field has been set.

### GetHttpOperationLogPublisher

`func (o *AddConnectionHandler200Response) GetHttpOperationLogPublisher() []string`

GetHttpOperationLogPublisher returns the HttpOperationLogPublisher field if non-nil, zero value otherwise.

### GetHttpOperationLogPublisherOk

`func (o *AddConnectionHandler200Response) GetHttpOperationLogPublisherOk() (*[]string, bool)`

GetHttpOperationLogPublisherOk returns a tuple with the HttpOperationLogPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpOperationLogPublisher

`func (o *AddConnectionHandler200Response) SetHttpOperationLogPublisher(v []string)`

SetHttpOperationLogPublisher sets HttpOperationLogPublisher field to given value.

### HasHttpOperationLogPublisher

`func (o *AddConnectionHandler200Response) HasHttpOperationLogPublisher() bool`

HasHttpOperationLogPublisher returns a boolean if a field has been set.

### GetKeepStats

`func (o *AddConnectionHandler200Response) GetKeepStats() bool`

GetKeepStats returns the KeepStats field if non-nil, zero value otherwise.

### GetKeepStatsOk

`func (o *AddConnectionHandler200Response) GetKeepStatsOk() (*bool, bool)`

GetKeepStatsOk returns a tuple with the KeepStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepStats

`func (o *AddConnectionHandler200Response) SetKeepStats(v bool)`

SetKeepStats sets KeepStats field to given value.

### HasKeepStats

`func (o *AddConnectionHandler200Response) HasKeepStats() bool`

HasKeepStats returns a boolean if a field has been set.

### GetAllowTCPReuseAddress

`func (o *AddConnectionHandler200Response) GetAllowTCPReuseAddress() bool`

GetAllowTCPReuseAddress returns the AllowTCPReuseAddress field if non-nil, zero value otherwise.

### GetAllowTCPReuseAddressOk

`func (o *AddConnectionHandler200Response) GetAllowTCPReuseAddressOk() (*bool, bool)`

GetAllowTCPReuseAddressOk returns a tuple with the AllowTCPReuseAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTCPReuseAddress

`func (o *AddConnectionHandler200Response) SetAllowTCPReuseAddress(v bool)`

SetAllowTCPReuseAddress sets AllowTCPReuseAddress field to given value.

### HasAllowTCPReuseAddress

`func (o *AddConnectionHandler200Response) HasAllowTCPReuseAddress() bool`

HasAllowTCPReuseAddress returns a boolean if a field has been set.

### GetIdleTimeLimit

`func (o *AddConnectionHandler200Response) GetIdleTimeLimit() string`

GetIdleTimeLimit returns the IdleTimeLimit field if non-nil, zero value otherwise.

### GetIdleTimeLimitOk

`func (o *AddConnectionHandler200Response) GetIdleTimeLimitOk() (*string, bool)`

GetIdleTimeLimitOk returns a tuple with the IdleTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeLimit

`func (o *AddConnectionHandler200Response) SetIdleTimeLimit(v string)`

SetIdleTimeLimit sets IdleTimeLimit field to given value.

### HasIdleTimeLimit

`func (o *AddConnectionHandler200Response) HasIdleTimeLimit() bool`

HasIdleTimeLimit returns a boolean if a field has been set.

### GetLowResourcesConnectionThreshold

`func (o *AddConnectionHandler200Response) GetLowResourcesConnectionThreshold() int64`

GetLowResourcesConnectionThreshold returns the LowResourcesConnectionThreshold field if non-nil, zero value otherwise.

### GetLowResourcesConnectionThresholdOk

`func (o *AddConnectionHandler200Response) GetLowResourcesConnectionThresholdOk() (*int64, bool)`

GetLowResourcesConnectionThresholdOk returns a tuple with the LowResourcesConnectionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowResourcesConnectionThreshold

`func (o *AddConnectionHandler200Response) SetLowResourcesConnectionThreshold(v int64)`

SetLowResourcesConnectionThreshold sets LowResourcesConnectionThreshold field to given value.

### HasLowResourcesConnectionThreshold

`func (o *AddConnectionHandler200Response) HasLowResourcesConnectionThreshold() bool`

HasLowResourcesConnectionThreshold returns a boolean if a field has been set.

### GetLowResourcesIdleTimeLimit

`func (o *AddConnectionHandler200Response) GetLowResourcesIdleTimeLimit() string`

GetLowResourcesIdleTimeLimit returns the LowResourcesIdleTimeLimit field if non-nil, zero value otherwise.

### GetLowResourcesIdleTimeLimitOk

`func (o *AddConnectionHandler200Response) GetLowResourcesIdleTimeLimitOk() (*string, bool)`

GetLowResourcesIdleTimeLimitOk returns a tuple with the LowResourcesIdleTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowResourcesIdleTimeLimit

`func (o *AddConnectionHandler200Response) SetLowResourcesIdleTimeLimit(v string)`

SetLowResourcesIdleTimeLimit sets LowResourcesIdleTimeLimit field to given value.

### HasLowResourcesIdleTimeLimit

`func (o *AddConnectionHandler200Response) HasLowResourcesIdleTimeLimit() bool`

HasLowResourcesIdleTimeLimit returns a boolean if a field has been set.

### GetEnableMultipartMIMEParameters

`func (o *AddConnectionHandler200Response) GetEnableMultipartMIMEParameters() bool`

GetEnableMultipartMIMEParameters returns the EnableMultipartMIMEParameters field if non-nil, zero value otherwise.

### GetEnableMultipartMIMEParametersOk

`func (o *AddConnectionHandler200Response) GetEnableMultipartMIMEParametersOk() (*bool, bool)`

GetEnableMultipartMIMEParametersOk returns a tuple with the EnableMultipartMIMEParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultipartMIMEParameters

`func (o *AddConnectionHandler200Response) SetEnableMultipartMIMEParameters(v bool)`

SetEnableMultipartMIMEParameters sets EnableMultipartMIMEParameters field to given value.

### HasEnableMultipartMIMEParameters

`func (o *AddConnectionHandler200Response) HasEnableMultipartMIMEParameters() bool`

HasEnableMultipartMIMEParameters returns a boolean if a field has been set.

### GetUseForwardedHeaders

`func (o *AddConnectionHandler200Response) GetUseForwardedHeaders() bool`

GetUseForwardedHeaders returns the UseForwardedHeaders field if non-nil, zero value otherwise.

### GetUseForwardedHeadersOk

`func (o *AddConnectionHandler200Response) GetUseForwardedHeadersOk() (*bool, bool)`

GetUseForwardedHeadersOk returns a tuple with the UseForwardedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardedHeaders

`func (o *AddConnectionHandler200Response) SetUseForwardedHeaders(v bool)`

SetUseForwardedHeaders sets UseForwardedHeaders field to given value.

### HasUseForwardedHeaders

`func (o *AddConnectionHandler200Response) HasUseForwardedHeaders() bool`

HasUseForwardedHeaders returns a boolean if a field has been set.

### GetHttpRequestHeaderSize

`func (o *AddConnectionHandler200Response) GetHttpRequestHeaderSize() int64`

GetHttpRequestHeaderSize returns the HttpRequestHeaderSize field if non-nil, zero value otherwise.

### GetHttpRequestHeaderSizeOk

`func (o *AddConnectionHandler200Response) GetHttpRequestHeaderSizeOk() (*int64, bool)`

GetHttpRequestHeaderSizeOk returns a tuple with the HttpRequestHeaderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestHeaderSize

`func (o *AddConnectionHandler200Response) SetHttpRequestHeaderSize(v int64)`

SetHttpRequestHeaderSize sets HttpRequestHeaderSize field to given value.

### HasHttpRequestHeaderSize

`func (o *AddConnectionHandler200Response) HasHttpRequestHeaderSize() bool`

HasHttpRequestHeaderSize returns a boolean if a field has been set.

### GetResponseHeader

`func (o *AddConnectionHandler200Response) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *AddConnectionHandler200Response) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *AddConnectionHandler200Response) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *AddConnectionHandler200Response) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetUseCorrelationIDHeader

`func (o *AddConnectionHandler200Response) GetUseCorrelationIDHeader() bool`

GetUseCorrelationIDHeader returns the UseCorrelationIDHeader field if non-nil, zero value otherwise.

### GetUseCorrelationIDHeaderOk

`func (o *AddConnectionHandler200Response) GetUseCorrelationIDHeaderOk() (*bool, bool)`

GetUseCorrelationIDHeaderOk returns a tuple with the UseCorrelationIDHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCorrelationIDHeader

`func (o *AddConnectionHandler200Response) SetUseCorrelationIDHeader(v bool)`

SetUseCorrelationIDHeader sets UseCorrelationIDHeader field to given value.

### HasUseCorrelationIDHeader

`func (o *AddConnectionHandler200Response) HasUseCorrelationIDHeader() bool`

HasUseCorrelationIDHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *AddConnectionHandler200Response) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *AddConnectionHandler200Response) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *AddConnectionHandler200Response) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *AddConnectionHandler200Response) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDRequestHeader

`func (o *AddConnectionHandler200Response) GetCorrelationIDRequestHeader() []string`

GetCorrelationIDRequestHeader returns the CorrelationIDRequestHeader field if non-nil, zero value otherwise.

### GetCorrelationIDRequestHeaderOk

`func (o *AddConnectionHandler200Response) GetCorrelationIDRequestHeaderOk() (*[]string, bool)`

GetCorrelationIDRequestHeaderOk returns a tuple with the CorrelationIDRequestHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDRequestHeader

`func (o *AddConnectionHandler200Response) SetCorrelationIDRequestHeader(v []string)`

SetCorrelationIDRequestHeader sets CorrelationIDRequestHeader field to given value.

### HasCorrelationIDRequestHeader

`func (o *AddConnectionHandler200Response) HasCorrelationIDRequestHeader() bool`

HasCorrelationIDRequestHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


