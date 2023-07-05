# LdapConnectionHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Connection Handler | 
**Schemas** | [**[]EnumldapConnectionHandlerSchemaUrn**](EnumldapConnectionHandlerSchemaUrn.md) |  | 
**ListenAddress** | Pointer to **[]string** | Specifies the address or set of addresses on which this LDAP Connection Handler should listen for connections from LDAP clients. | [optional] 
**ListenPort** | **int64** | Specifies the port number on which the LDAP Connection Handler will listen for connections from clients. | 
**UseSSL** | Pointer to **bool** | Indicates whether the LDAP Connection Handler should use SSL. | [optional] 
**AllowStartTLS** | Pointer to **bool** | Indicates whether clients are allowed to use StartTLS. | [optional] 
**SslCertNickname** | Pointer to **string** | Specifies the nickname (also called the alias) of the certificate that the LDAP Connection Handler should use when performing SSL communication. | [optional] 
**KeyManagerProvider** | Pointer to **string** | Specifies the name of the key manager that should be used with this LDAP Connection Handler . | [optional] 
**TrustManagerProvider** | Pointer to **string** | Specifies the name of the trust manager that should be used with the LDAP Connection Handler . | [optional] 
**AllowLDAPV2** | Pointer to **bool** | Indicates whether connections from LDAPv2 clients are allowed. | [optional] 
**UseTCPKeepAlive** | Pointer to **bool** | Indicates whether the LDAP Connection Handler should use TCP keep-alive. | [optional] 
**SendRejectionNotice** | Pointer to **bool** | Indicates whether the LDAP Connection Handler should send a notice of disconnection extended response message to the client if a new connection is rejected for some reason. | [optional] 
**FailedBindResponseDelay** | Pointer to **string** | Specifies the length of time that the server should delay the response to non-successful bind operations. A value of zero milliseconds indicates that non-successful bind operations should not be delayed. | [optional] 
**MaxRequestSize** | Pointer to **string** | Specifies the size of the largest LDAP request message that will be allowed by this LDAP Connection handler. | [optional] 
**MaxCancelHandlers** | Pointer to **int64** | Specifies the maximum number of threads that are used to process cancel and abandon requests from clients. | [optional] 
**NumAcceptHandlers** | Pointer to **int64** | Specifies the number of threads that are used to accept new client connections, and to perform any initial preparation on those connections that may be needed before the connection can be used to read requests and send responses. | [optional] 
**NumRequestHandlers** | Pointer to **int64** | Specifies the number of request handlers that are used to read requests from clients. | [optional] 
**SslClientAuthPolicy** | Pointer to [**EnumconnectionHandlerSslClientAuthPolicyProp**](EnumconnectionHandlerSslClientAuthPolicyProp.md) |  | [optional] 
**AcceptBacklog** | Pointer to **int64** | Specifies the maximum number of pending connection attempts that are allowed to queue up in the accept backlog before the server starts rejecting new connection attempts. | [optional] 
**SslProtocol** | Pointer to **[]string** | Specifies the names of the TLS protocols that are allowed for use in SSL or StartTLS communication. The set of supported ssl protocols can be viewed via the ssl context monitor entry. | [optional] 
**SslCipherSuite** | Pointer to **[]string** | Specifies the names of the TLS cipher suites that are allowed for use in SSL or StartTLS communication. The set of supported cipher suites can be viewed via the ssl context monitor entry. | [optional] 
**MaxBlockedWriteTimeLimit** | Pointer to **string** | Specifies the maximum length of time that attempts to write data to LDAP clients should be allowed to block. | [optional] 
**AutoAuthenticateUsingClientCertificate** | Pointer to **bool** | Indicates whether to attempt to automatically authenticate a client connection that has established a secure communication channel (using either SSL or StartTLS) and presented its own client certificate. Generally, clients should use the SASL EXTERNAL mechanism to authenticate using a client certificate, but some clients may not support that capability and/or may expect automatic authentication. | [optional] 
**CloseConnectionsWhenUnavailable** | Pointer to **bool** | Indicates whether all connections associated with this LDAP Connection Handler should be closed and no new connections accepted when the server has determined that it is \&quot;unavailable.\&quot; This allows clients (or a network load balancer) to route requests to another server. | [optional] 
**CloseConnectionsOnExplicitGC** | Pointer to **bool** | Indicates whether all connections associated with this LDAP Connection Handler should be closed before an explicit garbage collection is performed to allow clients to route requests to another server. | [optional] 
**Description** | Pointer to **string** | A description for this Connection Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Connection Handler is enabled. | 
**AllowedClient** | Pointer to **[]string** | Specifies a set of address masks that determines the addresses of the clients that are allowed to establish connections to this connection handler. | [optional] 
**DeniedClient** | Pointer to **[]string** | Specifies a set of address masks that determines the addresses of the clients that are not allowed to establish connections to this connection handler. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewLdapConnectionHandlerResponse

`func NewLdapConnectionHandlerResponse(id string, schemas []EnumldapConnectionHandlerSchemaUrn, listenPort int64, enabled bool, ) *LdapConnectionHandlerResponse`

NewLdapConnectionHandlerResponse instantiates a new LdapConnectionHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapConnectionHandlerResponseWithDefaults

`func NewLdapConnectionHandlerResponseWithDefaults() *LdapConnectionHandlerResponse`

NewLdapConnectionHandlerResponseWithDefaults instantiates a new LdapConnectionHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LdapConnectionHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LdapConnectionHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LdapConnectionHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *LdapConnectionHandlerResponse) GetSchemas() []EnumldapConnectionHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LdapConnectionHandlerResponse) GetSchemasOk() (*[]EnumldapConnectionHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LdapConnectionHandlerResponse) SetSchemas(v []EnumldapConnectionHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetListenAddress

`func (o *LdapConnectionHandlerResponse) GetListenAddress() []string`

GetListenAddress returns the ListenAddress field if non-nil, zero value otherwise.

### GetListenAddressOk

`func (o *LdapConnectionHandlerResponse) GetListenAddressOk() (*[]string, bool)`

GetListenAddressOk returns a tuple with the ListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddress

`func (o *LdapConnectionHandlerResponse) SetListenAddress(v []string)`

SetListenAddress sets ListenAddress field to given value.

### HasListenAddress

`func (o *LdapConnectionHandlerResponse) HasListenAddress() bool`

HasListenAddress returns a boolean if a field has been set.

### GetListenPort

`func (o *LdapConnectionHandlerResponse) GetListenPort() int64`

GetListenPort returns the ListenPort field if non-nil, zero value otherwise.

### GetListenPortOk

`func (o *LdapConnectionHandlerResponse) GetListenPortOk() (*int64, bool)`

GetListenPortOk returns a tuple with the ListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPort

`func (o *LdapConnectionHandlerResponse) SetListenPort(v int64)`

SetListenPort sets ListenPort field to given value.


### GetUseSSL

`func (o *LdapConnectionHandlerResponse) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *LdapConnectionHandlerResponse) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *LdapConnectionHandlerResponse) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *LdapConnectionHandlerResponse) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetAllowStartTLS

`func (o *LdapConnectionHandlerResponse) GetAllowStartTLS() bool`

GetAllowStartTLS returns the AllowStartTLS field if non-nil, zero value otherwise.

### GetAllowStartTLSOk

`func (o *LdapConnectionHandlerResponse) GetAllowStartTLSOk() (*bool, bool)`

GetAllowStartTLSOk returns a tuple with the AllowStartTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStartTLS

`func (o *LdapConnectionHandlerResponse) SetAllowStartTLS(v bool)`

SetAllowStartTLS sets AllowStartTLS field to given value.

### HasAllowStartTLS

`func (o *LdapConnectionHandlerResponse) HasAllowStartTLS() bool`

HasAllowStartTLS returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *LdapConnectionHandlerResponse) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *LdapConnectionHandlerResponse) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *LdapConnectionHandlerResponse) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.

### HasSslCertNickname

`func (o *LdapConnectionHandlerResponse) HasSslCertNickname() bool`

HasSslCertNickname returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *LdapConnectionHandlerResponse) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *LdapConnectionHandlerResponse) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *LdapConnectionHandlerResponse) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *LdapConnectionHandlerResponse) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *LdapConnectionHandlerResponse) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *LdapConnectionHandlerResponse) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *LdapConnectionHandlerResponse) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *LdapConnectionHandlerResponse) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetAllowLDAPV2

`func (o *LdapConnectionHandlerResponse) GetAllowLDAPV2() bool`

GetAllowLDAPV2 returns the AllowLDAPV2 field if non-nil, zero value otherwise.

### GetAllowLDAPV2Ok

`func (o *LdapConnectionHandlerResponse) GetAllowLDAPV2Ok() (*bool, bool)`

GetAllowLDAPV2Ok returns a tuple with the AllowLDAPV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLDAPV2

`func (o *LdapConnectionHandlerResponse) SetAllowLDAPV2(v bool)`

SetAllowLDAPV2 sets AllowLDAPV2 field to given value.

### HasAllowLDAPV2

`func (o *LdapConnectionHandlerResponse) HasAllowLDAPV2() bool`

HasAllowLDAPV2 returns a boolean if a field has been set.

### GetUseTCPKeepAlive

`func (o *LdapConnectionHandlerResponse) GetUseTCPKeepAlive() bool`

GetUseTCPKeepAlive returns the UseTCPKeepAlive field if non-nil, zero value otherwise.

### GetUseTCPKeepAliveOk

`func (o *LdapConnectionHandlerResponse) GetUseTCPKeepAliveOk() (*bool, bool)`

GetUseTCPKeepAliveOk returns a tuple with the UseTCPKeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTCPKeepAlive

`func (o *LdapConnectionHandlerResponse) SetUseTCPKeepAlive(v bool)`

SetUseTCPKeepAlive sets UseTCPKeepAlive field to given value.

### HasUseTCPKeepAlive

`func (o *LdapConnectionHandlerResponse) HasUseTCPKeepAlive() bool`

HasUseTCPKeepAlive returns a boolean if a field has been set.

### GetSendRejectionNotice

`func (o *LdapConnectionHandlerResponse) GetSendRejectionNotice() bool`

GetSendRejectionNotice returns the SendRejectionNotice field if non-nil, zero value otherwise.

### GetSendRejectionNoticeOk

`func (o *LdapConnectionHandlerResponse) GetSendRejectionNoticeOk() (*bool, bool)`

GetSendRejectionNoticeOk returns a tuple with the SendRejectionNotice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendRejectionNotice

`func (o *LdapConnectionHandlerResponse) SetSendRejectionNotice(v bool)`

SetSendRejectionNotice sets SendRejectionNotice field to given value.

### HasSendRejectionNotice

`func (o *LdapConnectionHandlerResponse) HasSendRejectionNotice() bool`

HasSendRejectionNotice returns a boolean if a field has been set.

### GetFailedBindResponseDelay

`func (o *LdapConnectionHandlerResponse) GetFailedBindResponseDelay() string`

GetFailedBindResponseDelay returns the FailedBindResponseDelay field if non-nil, zero value otherwise.

### GetFailedBindResponseDelayOk

`func (o *LdapConnectionHandlerResponse) GetFailedBindResponseDelayOk() (*string, bool)`

GetFailedBindResponseDelayOk returns a tuple with the FailedBindResponseDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedBindResponseDelay

`func (o *LdapConnectionHandlerResponse) SetFailedBindResponseDelay(v string)`

SetFailedBindResponseDelay sets FailedBindResponseDelay field to given value.

### HasFailedBindResponseDelay

`func (o *LdapConnectionHandlerResponse) HasFailedBindResponseDelay() bool`

HasFailedBindResponseDelay returns a boolean if a field has been set.

### GetMaxRequestSize

`func (o *LdapConnectionHandlerResponse) GetMaxRequestSize() string`

GetMaxRequestSize returns the MaxRequestSize field if non-nil, zero value otherwise.

### GetMaxRequestSizeOk

`func (o *LdapConnectionHandlerResponse) GetMaxRequestSizeOk() (*string, bool)`

GetMaxRequestSizeOk returns a tuple with the MaxRequestSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequestSize

`func (o *LdapConnectionHandlerResponse) SetMaxRequestSize(v string)`

SetMaxRequestSize sets MaxRequestSize field to given value.

### HasMaxRequestSize

`func (o *LdapConnectionHandlerResponse) HasMaxRequestSize() bool`

HasMaxRequestSize returns a boolean if a field has been set.

### GetMaxCancelHandlers

`func (o *LdapConnectionHandlerResponse) GetMaxCancelHandlers() int64`

GetMaxCancelHandlers returns the MaxCancelHandlers field if non-nil, zero value otherwise.

### GetMaxCancelHandlersOk

`func (o *LdapConnectionHandlerResponse) GetMaxCancelHandlersOk() (*int64, bool)`

GetMaxCancelHandlersOk returns a tuple with the MaxCancelHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCancelHandlers

`func (o *LdapConnectionHandlerResponse) SetMaxCancelHandlers(v int64)`

SetMaxCancelHandlers sets MaxCancelHandlers field to given value.

### HasMaxCancelHandlers

`func (o *LdapConnectionHandlerResponse) HasMaxCancelHandlers() bool`

HasMaxCancelHandlers returns a boolean if a field has been set.

### GetNumAcceptHandlers

`func (o *LdapConnectionHandlerResponse) GetNumAcceptHandlers() int64`

GetNumAcceptHandlers returns the NumAcceptHandlers field if non-nil, zero value otherwise.

### GetNumAcceptHandlersOk

`func (o *LdapConnectionHandlerResponse) GetNumAcceptHandlersOk() (*int64, bool)`

GetNumAcceptHandlersOk returns a tuple with the NumAcceptHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAcceptHandlers

`func (o *LdapConnectionHandlerResponse) SetNumAcceptHandlers(v int64)`

SetNumAcceptHandlers sets NumAcceptHandlers field to given value.

### HasNumAcceptHandlers

`func (o *LdapConnectionHandlerResponse) HasNumAcceptHandlers() bool`

HasNumAcceptHandlers returns a boolean if a field has been set.

### GetNumRequestHandlers

`func (o *LdapConnectionHandlerResponse) GetNumRequestHandlers() int64`

GetNumRequestHandlers returns the NumRequestHandlers field if non-nil, zero value otherwise.

### GetNumRequestHandlersOk

`func (o *LdapConnectionHandlerResponse) GetNumRequestHandlersOk() (*int64, bool)`

GetNumRequestHandlersOk returns a tuple with the NumRequestHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRequestHandlers

`func (o *LdapConnectionHandlerResponse) SetNumRequestHandlers(v int64)`

SetNumRequestHandlers sets NumRequestHandlers field to given value.

### HasNumRequestHandlers

`func (o *LdapConnectionHandlerResponse) HasNumRequestHandlers() bool`

HasNumRequestHandlers returns a boolean if a field has been set.

### GetSslClientAuthPolicy

`func (o *LdapConnectionHandlerResponse) GetSslClientAuthPolicy() EnumconnectionHandlerSslClientAuthPolicyProp`

GetSslClientAuthPolicy returns the SslClientAuthPolicy field if non-nil, zero value otherwise.

### GetSslClientAuthPolicyOk

`func (o *LdapConnectionHandlerResponse) GetSslClientAuthPolicyOk() (*EnumconnectionHandlerSslClientAuthPolicyProp, bool)`

GetSslClientAuthPolicyOk returns a tuple with the SslClientAuthPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslClientAuthPolicy

`func (o *LdapConnectionHandlerResponse) SetSslClientAuthPolicy(v EnumconnectionHandlerSslClientAuthPolicyProp)`

SetSslClientAuthPolicy sets SslClientAuthPolicy field to given value.

### HasSslClientAuthPolicy

`func (o *LdapConnectionHandlerResponse) HasSslClientAuthPolicy() bool`

HasSslClientAuthPolicy returns a boolean if a field has been set.

### GetAcceptBacklog

`func (o *LdapConnectionHandlerResponse) GetAcceptBacklog() int64`

GetAcceptBacklog returns the AcceptBacklog field if non-nil, zero value otherwise.

### GetAcceptBacklogOk

`func (o *LdapConnectionHandlerResponse) GetAcceptBacklogOk() (*int64, bool)`

GetAcceptBacklogOk returns a tuple with the AcceptBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptBacklog

`func (o *LdapConnectionHandlerResponse) SetAcceptBacklog(v int64)`

SetAcceptBacklog sets AcceptBacklog field to given value.

### HasAcceptBacklog

`func (o *LdapConnectionHandlerResponse) HasAcceptBacklog() bool`

HasAcceptBacklog returns a boolean if a field has been set.

### GetSslProtocol

`func (o *LdapConnectionHandlerResponse) GetSslProtocol() []string`

GetSslProtocol returns the SslProtocol field if non-nil, zero value otherwise.

### GetSslProtocolOk

`func (o *LdapConnectionHandlerResponse) GetSslProtocolOk() (*[]string, bool)`

GetSslProtocolOk returns a tuple with the SslProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslProtocol

`func (o *LdapConnectionHandlerResponse) SetSslProtocol(v []string)`

SetSslProtocol sets SslProtocol field to given value.

### HasSslProtocol

`func (o *LdapConnectionHandlerResponse) HasSslProtocol() bool`

HasSslProtocol returns a boolean if a field has been set.

### GetSslCipherSuite

`func (o *LdapConnectionHandlerResponse) GetSslCipherSuite() []string`

GetSslCipherSuite returns the SslCipherSuite field if non-nil, zero value otherwise.

### GetSslCipherSuiteOk

`func (o *LdapConnectionHandlerResponse) GetSslCipherSuiteOk() (*[]string, bool)`

GetSslCipherSuiteOk returns a tuple with the SslCipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCipherSuite

`func (o *LdapConnectionHandlerResponse) SetSslCipherSuite(v []string)`

SetSslCipherSuite sets SslCipherSuite field to given value.

### HasSslCipherSuite

`func (o *LdapConnectionHandlerResponse) HasSslCipherSuite() bool`

HasSslCipherSuite returns a boolean if a field has been set.

### GetMaxBlockedWriteTimeLimit

`func (o *LdapConnectionHandlerResponse) GetMaxBlockedWriteTimeLimit() string`

GetMaxBlockedWriteTimeLimit returns the MaxBlockedWriteTimeLimit field if non-nil, zero value otherwise.

### GetMaxBlockedWriteTimeLimitOk

`func (o *LdapConnectionHandlerResponse) GetMaxBlockedWriteTimeLimitOk() (*string, bool)`

GetMaxBlockedWriteTimeLimitOk returns a tuple with the MaxBlockedWriteTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBlockedWriteTimeLimit

`func (o *LdapConnectionHandlerResponse) SetMaxBlockedWriteTimeLimit(v string)`

SetMaxBlockedWriteTimeLimit sets MaxBlockedWriteTimeLimit field to given value.

### HasMaxBlockedWriteTimeLimit

`func (o *LdapConnectionHandlerResponse) HasMaxBlockedWriteTimeLimit() bool`

HasMaxBlockedWriteTimeLimit returns a boolean if a field has been set.

### GetAutoAuthenticateUsingClientCertificate

`func (o *LdapConnectionHandlerResponse) GetAutoAuthenticateUsingClientCertificate() bool`

GetAutoAuthenticateUsingClientCertificate returns the AutoAuthenticateUsingClientCertificate field if non-nil, zero value otherwise.

### GetAutoAuthenticateUsingClientCertificateOk

`func (o *LdapConnectionHandlerResponse) GetAutoAuthenticateUsingClientCertificateOk() (*bool, bool)`

GetAutoAuthenticateUsingClientCertificateOk returns a tuple with the AutoAuthenticateUsingClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAuthenticateUsingClientCertificate

`func (o *LdapConnectionHandlerResponse) SetAutoAuthenticateUsingClientCertificate(v bool)`

SetAutoAuthenticateUsingClientCertificate sets AutoAuthenticateUsingClientCertificate field to given value.

### HasAutoAuthenticateUsingClientCertificate

`func (o *LdapConnectionHandlerResponse) HasAutoAuthenticateUsingClientCertificate() bool`

HasAutoAuthenticateUsingClientCertificate returns a boolean if a field has been set.

### GetCloseConnectionsWhenUnavailable

`func (o *LdapConnectionHandlerResponse) GetCloseConnectionsWhenUnavailable() bool`

GetCloseConnectionsWhenUnavailable returns the CloseConnectionsWhenUnavailable field if non-nil, zero value otherwise.

### GetCloseConnectionsWhenUnavailableOk

`func (o *LdapConnectionHandlerResponse) GetCloseConnectionsWhenUnavailableOk() (*bool, bool)`

GetCloseConnectionsWhenUnavailableOk returns a tuple with the CloseConnectionsWhenUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseConnectionsWhenUnavailable

`func (o *LdapConnectionHandlerResponse) SetCloseConnectionsWhenUnavailable(v bool)`

SetCloseConnectionsWhenUnavailable sets CloseConnectionsWhenUnavailable field to given value.

### HasCloseConnectionsWhenUnavailable

`func (o *LdapConnectionHandlerResponse) HasCloseConnectionsWhenUnavailable() bool`

HasCloseConnectionsWhenUnavailable returns a boolean if a field has been set.

### GetCloseConnectionsOnExplicitGC

`func (o *LdapConnectionHandlerResponse) GetCloseConnectionsOnExplicitGC() bool`

GetCloseConnectionsOnExplicitGC returns the CloseConnectionsOnExplicitGC field if non-nil, zero value otherwise.

### GetCloseConnectionsOnExplicitGCOk

`func (o *LdapConnectionHandlerResponse) GetCloseConnectionsOnExplicitGCOk() (*bool, bool)`

GetCloseConnectionsOnExplicitGCOk returns a tuple with the CloseConnectionsOnExplicitGC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseConnectionsOnExplicitGC

`func (o *LdapConnectionHandlerResponse) SetCloseConnectionsOnExplicitGC(v bool)`

SetCloseConnectionsOnExplicitGC sets CloseConnectionsOnExplicitGC field to given value.

### HasCloseConnectionsOnExplicitGC

`func (o *LdapConnectionHandlerResponse) HasCloseConnectionsOnExplicitGC() bool`

HasCloseConnectionsOnExplicitGC returns a boolean if a field has been set.

### GetDescription

`func (o *LdapConnectionHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LdapConnectionHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LdapConnectionHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LdapConnectionHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *LdapConnectionHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LdapConnectionHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LdapConnectionHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowedClient

`func (o *LdapConnectionHandlerResponse) GetAllowedClient() []string`

GetAllowedClient returns the AllowedClient field if non-nil, zero value otherwise.

### GetAllowedClientOk

`func (o *LdapConnectionHandlerResponse) GetAllowedClientOk() (*[]string, bool)`

GetAllowedClientOk returns a tuple with the AllowedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClient

`func (o *LdapConnectionHandlerResponse) SetAllowedClient(v []string)`

SetAllowedClient sets AllowedClient field to given value.

### HasAllowedClient

`func (o *LdapConnectionHandlerResponse) HasAllowedClient() bool`

HasAllowedClient returns a boolean if a field has been set.

### GetDeniedClient

`func (o *LdapConnectionHandlerResponse) GetDeniedClient() []string`

GetDeniedClient returns the DeniedClient field if non-nil, zero value otherwise.

### GetDeniedClientOk

`func (o *LdapConnectionHandlerResponse) GetDeniedClientOk() (*[]string, bool)`

GetDeniedClientOk returns a tuple with the DeniedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedClient

`func (o *LdapConnectionHandlerResponse) SetDeniedClient(v []string)`

SetDeniedClient sets DeniedClient field to given value.

### HasDeniedClient

`func (o *LdapConnectionHandlerResponse) HasDeniedClient() bool`

HasDeniedClient returns a boolean if a field has been set.

### GetMeta

`func (o *LdapConnectionHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LdapConnectionHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LdapConnectionHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LdapConnectionHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *LdapConnectionHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *LdapConnectionHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *LdapConnectionHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *LdapConnectionHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


