# LdapConnectionHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumldapConnectionHandlerSchemaUrn**](EnumldapConnectionHandlerSchemaUrn.md) |  | 
**ListenAddress** | Pointer to **[]string** |  | [optional] 
**ListenPort** | **int32** | Specifies the port number on which the LDAP Connection Handler will listen for connections from clients. | 
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
**MaxCancelHandlers** | Pointer to **int32** | Specifies the maximum number of threads that are used to process cancel and abandon requests from clients. | [optional] 
**NumAcceptHandlers** | Pointer to **int32** | Specifies the number of threads that are used to accept new client connections, and to perform any initial preparation on those connections that may be needed before the connection can be used to read requests and send responses. | [optional] 
**NumRequestHandlers** | Pointer to **int32** | Specifies the number of request handlers that are used to read requests from clients. | [optional] 
**SslClientAuthPolicy** | Pointer to [**EnumconnectionHandlerSslClientAuthPolicyProp**](EnumconnectionHandlerSslClientAuthPolicyProp.md) |  | [optional] 
**AcceptBacklog** | Pointer to **int32** | Specifies the maximum number of pending connection attempts that are allowed to queue up in the accept backlog before the server starts rejecting new connection attempts. | [optional] 
**SslProtocol** | Pointer to **[]string** |  | [optional] 
**SslCipherSuite** | Pointer to **[]string** |  | [optional] 
**MaxBlockedWriteTimeLimit** | Pointer to **string** | Specifies the maximum length of time that attempts to write data to LDAP clients should be allowed to block. | [optional] 
**AutoAuthenticateUsingClientCertificate** | Pointer to **bool** | Indicates whether to attempt to automatically authenticate a client connection that has established a secure communication channel (using either SSL or StartTLS) and presented its own client certificate. Generally, clients should use the SASL EXTERNAL mechanism to authenticate using a client certificate, but some clients may not support that capability and/or may expect automatic authentication. | [optional] 
**CloseConnectionsWhenUnavailable** | Pointer to **bool** | Indicates whether all connections associated with this LDAP Connection Handler should be closed and no new connections accepted when the server has determined that it is \&quot;unavailable.\&quot; This allows clients (or a network load balancer) to route requests to another server. | [optional] 
**CloseConnectionsOnExplicitGC** | Pointer to **bool** | Indicates whether all connections associated with this LDAP Connection Handler should be closed before an explicit garbage collection is performed to allow clients to route requests to another server. | [optional] 
**Description** | Pointer to **string** | A description for this Connection Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Connection Handler is enabled. | 
**AllowedClient** | Pointer to **[]string** |  | [optional] 
**DeniedClient** | Pointer to **[]string** |  | [optional] 

## Methods

### NewLdapConnectionHandlerShared

`func NewLdapConnectionHandlerShared(schemas []EnumldapConnectionHandlerSchemaUrn, listenPort int32, enabled bool, ) *LdapConnectionHandlerShared`

NewLdapConnectionHandlerShared instantiates a new LdapConnectionHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapConnectionHandlerSharedWithDefaults

`func NewLdapConnectionHandlerSharedWithDefaults() *LdapConnectionHandlerShared`

NewLdapConnectionHandlerSharedWithDefaults instantiates a new LdapConnectionHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *LdapConnectionHandlerShared) GetSchemas() []EnumldapConnectionHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LdapConnectionHandlerShared) GetSchemasOk() (*[]EnumldapConnectionHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LdapConnectionHandlerShared) SetSchemas(v []EnumldapConnectionHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetListenAddress

`func (o *LdapConnectionHandlerShared) GetListenAddress() []string`

GetListenAddress returns the ListenAddress field if non-nil, zero value otherwise.

### GetListenAddressOk

`func (o *LdapConnectionHandlerShared) GetListenAddressOk() (*[]string, bool)`

GetListenAddressOk returns a tuple with the ListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddress

`func (o *LdapConnectionHandlerShared) SetListenAddress(v []string)`

SetListenAddress sets ListenAddress field to given value.

### HasListenAddress

`func (o *LdapConnectionHandlerShared) HasListenAddress() bool`

HasListenAddress returns a boolean if a field has been set.

### GetListenPort

`func (o *LdapConnectionHandlerShared) GetListenPort() int32`

GetListenPort returns the ListenPort field if non-nil, zero value otherwise.

### GetListenPortOk

`func (o *LdapConnectionHandlerShared) GetListenPortOk() (*int32, bool)`

GetListenPortOk returns a tuple with the ListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPort

`func (o *LdapConnectionHandlerShared) SetListenPort(v int32)`

SetListenPort sets ListenPort field to given value.


### GetUseSSL

`func (o *LdapConnectionHandlerShared) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *LdapConnectionHandlerShared) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *LdapConnectionHandlerShared) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *LdapConnectionHandlerShared) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetAllowStartTLS

`func (o *LdapConnectionHandlerShared) GetAllowStartTLS() bool`

GetAllowStartTLS returns the AllowStartTLS field if non-nil, zero value otherwise.

### GetAllowStartTLSOk

`func (o *LdapConnectionHandlerShared) GetAllowStartTLSOk() (*bool, bool)`

GetAllowStartTLSOk returns a tuple with the AllowStartTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStartTLS

`func (o *LdapConnectionHandlerShared) SetAllowStartTLS(v bool)`

SetAllowStartTLS sets AllowStartTLS field to given value.

### HasAllowStartTLS

`func (o *LdapConnectionHandlerShared) HasAllowStartTLS() bool`

HasAllowStartTLS returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *LdapConnectionHandlerShared) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *LdapConnectionHandlerShared) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *LdapConnectionHandlerShared) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.

### HasSslCertNickname

`func (o *LdapConnectionHandlerShared) HasSslCertNickname() bool`

HasSslCertNickname returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *LdapConnectionHandlerShared) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *LdapConnectionHandlerShared) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *LdapConnectionHandlerShared) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *LdapConnectionHandlerShared) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *LdapConnectionHandlerShared) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *LdapConnectionHandlerShared) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *LdapConnectionHandlerShared) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *LdapConnectionHandlerShared) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetAllowLDAPV2

`func (o *LdapConnectionHandlerShared) GetAllowLDAPV2() bool`

GetAllowLDAPV2 returns the AllowLDAPV2 field if non-nil, zero value otherwise.

### GetAllowLDAPV2Ok

`func (o *LdapConnectionHandlerShared) GetAllowLDAPV2Ok() (*bool, bool)`

GetAllowLDAPV2Ok returns a tuple with the AllowLDAPV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLDAPV2

`func (o *LdapConnectionHandlerShared) SetAllowLDAPV2(v bool)`

SetAllowLDAPV2 sets AllowLDAPV2 field to given value.

### HasAllowLDAPV2

`func (o *LdapConnectionHandlerShared) HasAllowLDAPV2() bool`

HasAllowLDAPV2 returns a boolean if a field has been set.

### GetUseTCPKeepAlive

`func (o *LdapConnectionHandlerShared) GetUseTCPKeepAlive() bool`

GetUseTCPKeepAlive returns the UseTCPKeepAlive field if non-nil, zero value otherwise.

### GetUseTCPKeepAliveOk

`func (o *LdapConnectionHandlerShared) GetUseTCPKeepAliveOk() (*bool, bool)`

GetUseTCPKeepAliveOk returns a tuple with the UseTCPKeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTCPKeepAlive

`func (o *LdapConnectionHandlerShared) SetUseTCPKeepAlive(v bool)`

SetUseTCPKeepAlive sets UseTCPKeepAlive field to given value.

### HasUseTCPKeepAlive

`func (o *LdapConnectionHandlerShared) HasUseTCPKeepAlive() bool`

HasUseTCPKeepAlive returns a boolean if a field has been set.

### GetSendRejectionNotice

`func (o *LdapConnectionHandlerShared) GetSendRejectionNotice() bool`

GetSendRejectionNotice returns the SendRejectionNotice field if non-nil, zero value otherwise.

### GetSendRejectionNoticeOk

`func (o *LdapConnectionHandlerShared) GetSendRejectionNoticeOk() (*bool, bool)`

GetSendRejectionNoticeOk returns a tuple with the SendRejectionNotice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendRejectionNotice

`func (o *LdapConnectionHandlerShared) SetSendRejectionNotice(v bool)`

SetSendRejectionNotice sets SendRejectionNotice field to given value.

### HasSendRejectionNotice

`func (o *LdapConnectionHandlerShared) HasSendRejectionNotice() bool`

HasSendRejectionNotice returns a boolean if a field has been set.

### GetFailedBindResponseDelay

`func (o *LdapConnectionHandlerShared) GetFailedBindResponseDelay() string`

GetFailedBindResponseDelay returns the FailedBindResponseDelay field if non-nil, zero value otherwise.

### GetFailedBindResponseDelayOk

`func (o *LdapConnectionHandlerShared) GetFailedBindResponseDelayOk() (*string, bool)`

GetFailedBindResponseDelayOk returns a tuple with the FailedBindResponseDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedBindResponseDelay

`func (o *LdapConnectionHandlerShared) SetFailedBindResponseDelay(v string)`

SetFailedBindResponseDelay sets FailedBindResponseDelay field to given value.

### HasFailedBindResponseDelay

`func (o *LdapConnectionHandlerShared) HasFailedBindResponseDelay() bool`

HasFailedBindResponseDelay returns a boolean if a field has been set.

### GetMaxRequestSize

`func (o *LdapConnectionHandlerShared) GetMaxRequestSize() string`

GetMaxRequestSize returns the MaxRequestSize field if non-nil, zero value otherwise.

### GetMaxRequestSizeOk

`func (o *LdapConnectionHandlerShared) GetMaxRequestSizeOk() (*string, bool)`

GetMaxRequestSizeOk returns a tuple with the MaxRequestSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequestSize

`func (o *LdapConnectionHandlerShared) SetMaxRequestSize(v string)`

SetMaxRequestSize sets MaxRequestSize field to given value.

### HasMaxRequestSize

`func (o *LdapConnectionHandlerShared) HasMaxRequestSize() bool`

HasMaxRequestSize returns a boolean if a field has been set.

### GetMaxCancelHandlers

`func (o *LdapConnectionHandlerShared) GetMaxCancelHandlers() int32`

GetMaxCancelHandlers returns the MaxCancelHandlers field if non-nil, zero value otherwise.

### GetMaxCancelHandlersOk

`func (o *LdapConnectionHandlerShared) GetMaxCancelHandlersOk() (*int32, bool)`

GetMaxCancelHandlersOk returns a tuple with the MaxCancelHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCancelHandlers

`func (o *LdapConnectionHandlerShared) SetMaxCancelHandlers(v int32)`

SetMaxCancelHandlers sets MaxCancelHandlers field to given value.

### HasMaxCancelHandlers

`func (o *LdapConnectionHandlerShared) HasMaxCancelHandlers() bool`

HasMaxCancelHandlers returns a boolean if a field has been set.

### GetNumAcceptHandlers

`func (o *LdapConnectionHandlerShared) GetNumAcceptHandlers() int32`

GetNumAcceptHandlers returns the NumAcceptHandlers field if non-nil, zero value otherwise.

### GetNumAcceptHandlersOk

`func (o *LdapConnectionHandlerShared) GetNumAcceptHandlersOk() (*int32, bool)`

GetNumAcceptHandlersOk returns a tuple with the NumAcceptHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAcceptHandlers

`func (o *LdapConnectionHandlerShared) SetNumAcceptHandlers(v int32)`

SetNumAcceptHandlers sets NumAcceptHandlers field to given value.

### HasNumAcceptHandlers

`func (o *LdapConnectionHandlerShared) HasNumAcceptHandlers() bool`

HasNumAcceptHandlers returns a boolean if a field has been set.

### GetNumRequestHandlers

`func (o *LdapConnectionHandlerShared) GetNumRequestHandlers() int32`

GetNumRequestHandlers returns the NumRequestHandlers field if non-nil, zero value otherwise.

### GetNumRequestHandlersOk

`func (o *LdapConnectionHandlerShared) GetNumRequestHandlersOk() (*int32, bool)`

GetNumRequestHandlersOk returns a tuple with the NumRequestHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRequestHandlers

`func (o *LdapConnectionHandlerShared) SetNumRequestHandlers(v int32)`

SetNumRequestHandlers sets NumRequestHandlers field to given value.

### HasNumRequestHandlers

`func (o *LdapConnectionHandlerShared) HasNumRequestHandlers() bool`

HasNumRequestHandlers returns a boolean if a field has been set.

### GetSslClientAuthPolicy

`func (o *LdapConnectionHandlerShared) GetSslClientAuthPolicy() EnumconnectionHandlerSslClientAuthPolicyProp`

GetSslClientAuthPolicy returns the SslClientAuthPolicy field if non-nil, zero value otherwise.

### GetSslClientAuthPolicyOk

`func (o *LdapConnectionHandlerShared) GetSslClientAuthPolicyOk() (*EnumconnectionHandlerSslClientAuthPolicyProp, bool)`

GetSslClientAuthPolicyOk returns a tuple with the SslClientAuthPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslClientAuthPolicy

`func (o *LdapConnectionHandlerShared) SetSslClientAuthPolicy(v EnumconnectionHandlerSslClientAuthPolicyProp)`

SetSslClientAuthPolicy sets SslClientAuthPolicy field to given value.

### HasSslClientAuthPolicy

`func (o *LdapConnectionHandlerShared) HasSslClientAuthPolicy() bool`

HasSslClientAuthPolicy returns a boolean if a field has been set.

### GetAcceptBacklog

`func (o *LdapConnectionHandlerShared) GetAcceptBacklog() int32`

GetAcceptBacklog returns the AcceptBacklog field if non-nil, zero value otherwise.

### GetAcceptBacklogOk

`func (o *LdapConnectionHandlerShared) GetAcceptBacklogOk() (*int32, bool)`

GetAcceptBacklogOk returns a tuple with the AcceptBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptBacklog

`func (o *LdapConnectionHandlerShared) SetAcceptBacklog(v int32)`

SetAcceptBacklog sets AcceptBacklog field to given value.

### HasAcceptBacklog

`func (o *LdapConnectionHandlerShared) HasAcceptBacklog() bool`

HasAcceptBacklog returns a boolean if a field has been set.

### GetSslProtocol

`func (o *LdapConnectionHandlerShared) GetSslProtocol() []string`

GetSslProtocol returns the SslProtocol field if non-nil, zero value otherwise.

### GetSslProtocolOk

`func (o *LdapConnectionHandlerShared) GetSslProtocolOk() (*[]string, bool)`

GetSslProtocolOk returns a tuple with the SslProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslProtocol

`func (o *LdapConnectionHandlerShared) SetSslProtocol(v []string)`

SetSslProtocol sets SslProtocol field to given value.

### HasSslProtocol

`func (o *LdapConnectionHandlerShared) HasSslProtocol() bool`

HasSslProtocol returns a boolean if a field has been set.

### GetSslCipherSuite

`func (o *LdapConnectionHandlerShared) GetSslCipherSuite() []string`

GetSslCipherSuite returns the SslCipherSuite field if non-nil, zero value otherwise.

### GetSslCipherSuiteOk

`func (o *LdapConnectionHandlerShared) GetSslCipherSuiteOk() (*[]string, bool)`

GetSslCipherSuiteOk returns a tuple with the SslCipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCipherSuite

`func (o *LdapConnectionHandlerShared) SetSslCipherSuite(v []string)`

SetSslCipherSuite sets SslCipherSuite field to given value.

### HasSslCipherSuite

`func (o *LdapConnectionHandlerShared) HasSslCipherSuite() bool`

HasSslCipherSuite returns a boolean if a field has been set.

### GetMaxBlockedWriteTimeLimit

`func (o *LdapConnectionHandlerShared) GetMaxBlockedWriteTimeLimit() string`

GetMaxBlockedWriteTimeLimit returns the MaxBlockedWriteTimeLimit field if non-nil, zero value otherwise.

### GetMaxBlockedWriteTimeLimitOk

`func (o *LdapConnectionHandlerShared) GetMaxBlockedWriteTimeLimitOk() (*string, bool)`

GetMaxBlockedWriteTimeLimitOk returns a tuple with the MaxBlockedWriteTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBlockedWriteTimeLimit

`func (o *LdapConnectionHandlerShared) SetMaxBlockedWriteTimeLimit(v string)`

SetMaxBlockedWriteTimeLimit sets MaxBlockedWriteTimeLimit field to given value.

### HasMaxBlockedWriteTimeLimit

`func (o *LdapConnectionHandlerShared) HasMaxBlockedWriteTimeLimit() bool`

HasMaxBlockedWriteTimeLimit returns a boolean if a field has been set.

### GetAutoAuthenticateUsingClientCertificate

`func (o *LdapConnectionHandlerShared) GetAutoAuthenticateUsingClientCertificate() bool`

GetAutoAuthenticateUsingClientCertificate returns the AutoAuthenticateUsingClientCertificate field if non-nil, zero value otherwise.

### GetAutoAuthenticateUsingClientCertificateOk

`func (o *LdapConnectionHandlerShared) GetAutoAuthenticateUsingClientCertificateOk() (*bool, bool)`

GetAutoAuthenticateUsingClientCertificateOk returns a tuple with the AutoAuthenticateUsingClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAuthenticateUsingClientCertificate

`func (o *LdapConnectionHandlerShared) SetAutoAuthenticateUsingClientCertificate(v bool)`

SetAutoAuthenticateUsingClientCertificate sets AutoAuthenticateUsingClientCertificate field to given value.

### HasAutoAuthenticateUsingClientCertificate

`func (o *LdapConnectionHandlerShared) HasAutoAuthenticateUsingClientCertificate() bool`

HasAutoAuthenticateUsingClientCertificate returns a boolean if a field has been set.

### GetCloseConnectionsWhenUnavailable

`func (o *LdapConnectionHandlerShared) GetCloseConnectionsWhenUnavailable() bool`

GetCloseConnectionsWhenUnavailable returns the CloseConnectionsWhenUnavailable field if non-nil, zero value otherwise.

### GetCloseConnectionsWhenUnavailableOk

`func (o *LdapConnectionHandlerShared) GetCloseConnectionsWhenUnavailableOk() (*bool, bool)`

GetCloseConnectionsWhenUnavailableOk returns a tuple with the CloseConnectionsWhenUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseConnectionsWhenUnavailable

`func (o *LdapConnectionHandlerShared) SetCloseConnectionsWhenUnavailable(v bool)`

SetCloseConnectionsWhenUnavailable sets CloseConnectionsWhenUnavailable field to given value.

### HasCloseConnectionsWhenUnavailable

`func (o *LdapConnectionHandlerShared) HasCloseConnectionsWhenUnavailable() bool`

HasCloseConnectionsWhenUnavailable returns a boolean if a field has been set.

### GetCloseConnectionsOnExplicitGC

`func (o *LdapConnectionHandlerShared) GetCloseConnectionsOnExplicitGC() bool`

GetCloseConnectionsOnExplicitGC returns the CloseConnectionsOnExplicitGC field if non-nil, zero value otherwise.

### GetCloseConnectionsOnExplicitGCOk

`func (o *LdapConnectionHandlerShared) GetCloseConnectionsOnExplicitGCOk() (*bool, bool)`

GetCloseConnectionsOnExplicitGCOk returns a tuple with the CloseConnectionsOnExplicitGC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseConnectionsOnExplicitGC

`func (o *LdapConnectionHandlerShared) SetCloseConnectionsOnExplicitGC(v bool)`

SetCloseConnectionsOnExplicitGC sets CloseConnectionsOnExplicitGC field to given value.

### HasCloseConnectionsOnExplicitGC

`func (o *LdapConnectionHandlerShared) HasCloseConnectionsOnExplicitGC() bool`

HasCloseConnectionsOnExplicitGC returns a boolean if a field has been set.

### GetDescription

`func (o *LdapConnectionHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LdapConnectionHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LdapConnectionHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LdapConnectionHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *LdapConnectionHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LdapConnectionHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LdapConnectionHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowedClient

`func (o *LdapConnectionHandlerShared) GetAllowedClient() []string`

GetAllowedClient returns the AllowedClient field if non-nil, zero value otherwise.

### GetAllowedClientOk

`func (o *LdapConnectionHandlerShared) GetAllowedClientOk() (*[]string, bool)`

GetAllowedClientOk returns a tuple with the AllowedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClient

`func (o *LdapConnectionHandlerShared) SetAllowedClient(v []string)`

SetAllowedClient sets AllowedClient field to given value.

### HasAllowedClient

`func (o *LdapConnectionHandlerShared) HasAllowedClient() bool`

HasAllowedClient returns a boolean if a field has been set.

### GetDeniedClient

`func (o *LdapConnectionHandlerShared) GetDeniedClient() []string`

GetDeniedClient returns the DeniedClient field if non-nil, zero value otherwise.

### GetDeniedClientOk

`func (o *LdapConnectionHandlerShared) GetDeniedClientOk() (*[]string, bool)`

GetDeniedClientOk returns a tuple with the DeniedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedClient

`func (o *LdapConnectionHandlerShared) SetDeniedClient(v []string)`

SetDeniedClient sets DeniedClient field to given value.

### HasDeniedClient

`func (o *LdapConnectionHandlerShared) HasDeniedClient() bool`

HasDeniedClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


