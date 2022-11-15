# HttpConnectionHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumhttpConnectionHandlerSchemaUrn**](EnumhttpConnectionHandlerSchemaUrn.md) |  | 
**ListenAddress** | Pointer to **string** | Specifies the address on which to listen for connections from HTTP clients. If no value is defined, the server will listen on all addresses on all interfaces. | [optional] 
**ListenPort** | **int32** | Specifies the port number on which the HTTP Connection Handler will listen for connections from clients. | 
**UseSSL** | Pointer to **bool** | Indicates whether the HTTP Connection Handler should use SSL. | [optional] 
**SslCertNickname** | Pointer to **string** | Specifies the nickname (also called the alias) of the certificate that the HTTP Connection Handler should use when performing SSL communication. | [optional] 
**HttpServletExtension** | Pointer to **[]string** |  | [optional] 
**WebApplicationExtension** | Pointer to **[]string** |  | [optional] 
**HttpOperationLogPublisher** | Pointer to **[]string** |  | [optional] 
**SslProtocol** | Pointer to **[]string** |  | [optional] 
**SslCipherSuite** | Pointer to **[]string** |  | [optional] 
**KeyManagerProvider** | Pointer to **string** | Specifies the key manager provider that will be used to obtain the certificate to present to HTTPS clients. | [optional] 
**TrustManagerProvider** | Pointer to **string** | Specifies the trust manager provider that will be used to validate any certificates presented by HTTPS clients. | [optional] 
**NumRequestHandlers** | Pointer to **int32** | Specifies the number of threads that will be used for accepting connections and reading requests from clients. | [optional] 
**KeepStats** | Pointer to **bool** | Indicates whether to enable statistics collection for this connection handler. | [optional] 
**AcceptBacklog** | Pointer to **int32** | Specifies the number of concurrent outstanding connection attempts that the connection handler should allow. The default value should be acceptable in most cases, but it may need to be increased in environments that may attempt to establish large numbers of connections simultaneously. | [optional] 
**AllowTCPReuseAddress** | Pointer to **bool** | Indicates whether the server should attempt to reuse socket descriptors. This may be useful in environments with a high rate of connection establishment and termination. | [optional] 
**IdleTimeLimit** | Pointer to **string** | Specifies the maximum idle time for a connection. The max idle time is applied when waiting for a new request to be received on a connection, when reading the headers and content of a request, or when writing the headers and content of a response. | [optional] 
**LowResourcesConnectionThreshold** | Pointer to **int32** | Specifies the number of connections, which if exceeded, places this handler in a low resource state where a different idle time limit is applied on the connections. | [optional] 
**LowResourcesIdleTimeLimit** | Pointer to **string** | Specifies the maximum idle time for a connection when this handler is in a low resource state as defined by low-resource-connections. The max idle time is applied when waiting for a new request to be received on a connection, when reading the headers and content of a request, or when writing the headers and content of a response. | [optional] 
**EnableMultipartMIMEParameters** | Pointer to **bool** | Determines whether request form parameters submitted in multipart/ form-data (RFC 2388) format should be processed as request parameters. | [optional] 
**UseForwardedHeaders** | Pointer to **bool** | Indicates whether to use \&quot;Forwarded\&quot; and \&quot;X-Forwarded-*\&quot; request headers to override corresponding HTTP request information available during request processing. | [optional] 
**HttpRequestHeaderSize** | Pointer to **int32** | Specifies the maximum buffer size of an http request including the request uri and all of the request headers. | [optional] 
**ResponseHeader** | Pointer to **[]string** |  | [optional] 
**UseCorrelationIDHeader** | Pointer to **bool** | If enabled, a correlation ID header will be added to outgoing HTTP responses. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 
**CorrelationIDRequestHeader** | Pointer to **[]string** |  | [optional] 
**SslClientAuthPolicy** | Pointer to [**EnumconnectionHandlerSslClientAuthPolicyProp**](EnumconnectionHandlerSslClientAuthPolicyProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Connection Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Connection Handler is enabled. | 
**AllowedClient** | Pointer to **[]string** |  | [optional] 
**DeniedClient** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHttpConnectionHandlerShared

`func NewHttpConnectionHandlerShared(schemas []EnumhttpConnectionHandlerSchemaUrn, listenPort int32, enabled bool, ) *HttpConnectionHandlerShared`

NewHttpConnectionHandlerShared instantiates a new HttpConnectionHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpConnectionHandlerSharedWithDefaults

`func NewHttpConnectionHandlerSharedWithDefaults() *HttpConnectionHandlerShared`

NewHttpConnectionHandlerSharedWithDefaults instantiates a new HttpConnectionHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *HttpConnectionHandlerShared) GetSchemas() []EnumhttpConnectionHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *HttpConnectionHandlerShared) GetSchemasOk() (*[]EnumhttpConnectionHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *HttpConnectionHandlerShared) SetSchemas(v []EnumhttpConnectionHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetListenAddress

`func (o *HttpConnectionHandlerShared) GetListenAddress() string`

GetListenAddress returns the ListenAddress field if non-nil, zero value otherwise.

### GetListenAddressOk

`func (o *HttpConnectionHandlerShared) GetListenAddressOk() (*string, bool)`

GetListenAddressOk returns a tuple with the ListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddress

`func (o *HttpConnectionHandlerShared) SetListenAddress(v string)`

SetListenAddress sets ListenAddress field to given value.

### HasListenAddress

`func (o *HttpConnectionHandlerShared) HasListenAddress() bool`

HasListenAddress returns a boolean if a field has been set.

### GetListenPort

`func (o *HttpConnectionHandlerShared) GetListenPort() int32`

GetListenPort returns the ListenPort field if non-nil, zero value otherwise.

### GetListenPortOk

`func (o *HttpConnectionHandlerShared) GetListenPortOk() (*int32, bool)`

GetListenPortOk returns a tuple with the ListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPort

`func (o *HttpConnectionHandlerShared) SetListenPort(v int32)`

SetListenPort sets ListenPort field to given value.


### GetUseSSL

`func (o *HttpConnectionHandlerShared) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *HttpConnectionHandlerShared) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *HttpConnectionHandlerShared) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *HttpConnectionHandlerShared) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *HttpConnectionHandlerShared) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *HttpConnectionHandlerShared) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *HttpConnectionHandlerShared) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.

### HasSslCertNickname

`func (o *HttpConnectionHandlerShared) HasSslCertNickname() bool`

HasSslCertNickname returns a boolean if a field has been set.

### GetHttpServletExtension

`func (o *HttpConnectionHandlerShared) GetHttpServletExtension() []string`

GetHttpServletExtension returns the HttpServletExtension field if non-nil, zero value otherwise.

### GetHttpServletExtensionOk

`func (o *HttpConnectionHandlerShared) GetHttpServletExtensionOk() (*[]string, bool)`

GetHttpServletExtensionOk returns a tuple with the HttpServletExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServletExtension

`func (o *HttpConnectionHandlerShared) SetHttpServletExtension(v []string)`

SetHttpServletExtension sets HttpServletExtension field to given value.

### HasHttpServletExtension

`func (o *HttpConnectionHandlerShared) HasHttpServletExtension() bool`

HasHttpServletExtension returns a boolean if a field has been set.

### GetWebApplicationExtension

`func (o *HttpConnectionHandlerShared) GetWebApplicationExtension() []string`

GetWebApplicationExtension returns the WebApplicationExtension field if non-nil, zero value otherwise.

### GetWebApplicationExtensionOk

`func (o *HttpConnectionHandlerShared) GetWebApplicationExtensionOk() (*[]string, bool)`

GetWebApplicationExtensionOk returns a tuple with the WebApplicationExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebApplicationExtension

`func (o *HttpConnectionHandlerShared) SetWebApplicationExtension(v []string)`

SetWebApplicationExtension sets WebApplicationExtension field to given value.

### HasWebApplicationExtension

`func (o *HttpConnectionHandlerShared) HasWebApplicationExtension() bool`

HasWebApplicationExtension returns a boolean if a field has been set.

### GetHttpOperationLogPublisher

`func (o *HttpConnectionHandlerShared) GetHttpOperationLogPublisher() []string`

GetHttpOperationLogPublisher returns the HttpOperationLogPublisher field if non-nil, zero value otherwise.

### GetHttpOperationLogPublisherOk

`func (o *HttpConnectionHandlerShared) GetHttpOperationLogPublisherOk() (*[]string, bool)`

GetHttpOperationLogPublisherOk returns a tuple with the HttpOperationLogPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpOperationLogPublisher

`func (o *HttpConnectionHandlerShared) SetHttpOperationLogPublisher(v []string)`

SetHttpOperationLogPublisher sets HttpOperationLogPublisher field to given value.

### HasHttpOperationLogPublisher

`func (o *HttpConnectionHandlerShared) HasHttpOperationLogPublisher() bool`

HasHttpOperationLogPublisher returns a boolean if a field has been set.

### GetSslProtocol

`func (o *HttpConnectionHandlerShared) GetSslProtocol() []string`

GetSslProtocol returns the SslProtocol field if non-nil, zero value otherwise.

### GetSslProtocolOk

`func (o *HttpConnectionHandlerShared) GetSslProtocolOk() (*[]string, bool)`

GetSslProtocolOk returns a tuple with the SslProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslProtocol

`func (o *HttpConnectionHandlerShared) SetSslProtocol(v []string)`

SetSslProtocol sets SslProtocol field to given value.

### HasSslProtocol

`func (o *HttpConnectionHandlerShared) HasSslProtocol() bool`

HasSslProtocol returns a boolean if a field has been set.

### GetSslCipherSuite

`func (o *HttpConnectionHandlerShared) GetSslCipherSuite() []string`

GetSslCipherSuite returns the SslCipherSuite field if non-nil, zero value otherwise.

### GetSslCipherSuiteOk

`func (o *HttpConnectionHandlerShared) GetSslCipherSuiteOk() (*[]string, bool)`

GetSslCipherSuiteOk returns a tuple with the SslCipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCipherSuite

`func (o *HttpConnectionHandlerShared) SetSslCipherSuite(v []string)`

SetSslCipherSuite sets SslCipherSuite field to given value.

### HasSslCipherSuite

`func (o *HttpConnectionHandlerShared) HasSslCipherSuite() bool`

HasSslCipherSuite returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *HttpConnectionHandlerShared) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *HttpConnectionHandlerShared) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *HttpConnectionHandlerShared) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *HttpConnectionHandlerShared) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *HttpConnectionHandlerShared) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *HttpConnectionHandlerShared) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *HttpConnectionHandlerShared) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *HttpConnectionHandlerShared) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetNumRequestHandlers

`func (o *HttpConnectionHandlerShared) GetNumRequestHandlers() int32`

GetNumRequestHandlers returns the NumRequestHandlers field if non-nil, zero value otherwise.

### GetNumRequestHandlersOk

`func (o *HttpConnectionHandlerShared) GetNumRequestHandlersOk() (*int32, bool)`

GetNumRequestHandlersOk returns a tuple with the NumRequestHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRequestHandlers

`func (o *HttpConnectionHandlerShared) SetNumRequestHandlers(v int32)`

SetNumRequestHandlers sets NumRequestHandlers field to given value.

### HasNumRequestHandlers

`func (o *HttpConnectionHandlerShared) HasNumRequestHandlers() bool`

HasNumRequestHandlers returns a boolean if a field has been set.

### GetKeepStats

`func (o *HttpConnectionHandlerShared) GetKeepStats() bool`

GetKeepStats returns the KeepStats field if non-nil, zero value otherwise.

### GetKeepStatsOk

`func (o *HttpConnectionHandlerShared) GetKeepStatsOk() (*bool, bool)`

GetKeepStatsOk returns a tuple with the KeepStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepStats

`func (o *HttpConnectionHandlerShared) SetKeepStats(v bool)`

SetKeepStats sets KeepStats field to given value.

### HasKeepStats

`func (o *HttpConnectionHandlerShared) HasKeepStats() bool`

HasKeepStats returns a boolean if a field has been set.

### GetAcceptBacklog

`func (o *HttpConnectionHandlerShared) GetAcceptBacklog() int32`

GetAcceptBacklog returns the AcceptBacklog field if non-nil, zero value otherwise.

### GetAcceptBacklogOk

`func (o *HttpConnectionHandlerShared) GetAcceptBacklogOk() (*int32, bool)`

GetAcceptBacklogOk returns a tuple with the AcceptBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptBacklog

`func (o *HttpConnectionHandlerShared) SetAcceptBacklog(v int32)`

SetAcceptBacklog sets AcceptBacklog field to given value.

### HasAcceptBacklog

`func (o *HttpConnectionHandlerShared) HasAcceptBacklog() bool`

HasAcceptBacklog returns a boolean if a field has been set.

### GetAllowTCPReuseAddress

`func (o *HttpConnectionHandlerShared) GetAllowTCPReuseAddress() bool`

GetAllowTCPReuseAddress returns the AllowTCPReuseAddress field if non-nil, zero value otherwise.

### GetAllowTCPReuseAddressOk

`func (o *HttpConnectionHandlerShared) GetAllowTCPReuseAddressOk() (*bool, bool)`

GetAllowTCPReuseAddressOk returns a tuple with the AllowTCPReuseAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTCPReuseAddress

`func (o *HttpConnectionHandlerShared) SetAllowTCPReuseAddress(v bool)`

SetAllowTCPReuseAddress sets AllowTCPReuseAddress field to given value.

### HasAllowTCPReuseAddress

`func (o *HttpConnectionHandlerShared) HasAllowTCPReuseAddress() bool`

HasAllowTCPReuseAddress returns a boolean if a field has been set.

### GetIdleTimeLimit

`func (o *HttpConnectionHandlerShared) GetIdleTimeLimit() string`

GetIdleTimeLimit returns the IdleTimeLimit field if non-nil, zero value otherwise.

### GetIdleTimeLimitOk

`func (o *HttpConnectionHandlerShared) GetIdleTimeLimitOk() (*string, bool)`

GetIdleTimeLimitOk returns a tuple with the IdleTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeLimit

`func (o *HttpConnectionHandlerShared) SetIdleTimeLimit(v string)`

SetIdleTimeLimit sets IdleTimeLimit field to given value.

### HasIdleTimeLimit

`func (o *HttpConnectionHandlerShared) HasIdleTimeLimit() bool`

HasIdleTimeLimit returns a boolean if a field has been set.

### GetLowResourcesConnectionThreshold

`func (o *HttpConnectionHandlerShared) GetLowResourcesConnectionThreshold() int32`

GetLowResourcesConnectionThreshold returns the LowResourcesConnectionThreshold field if non-nil, zero value otherwise.

### GetLowResourcesConnectionThresholdOk

`func (o *HttpConnectionHandlerShared) GetLowResourcesConnectionThresholdOk() (*int32, bool)`

GetLowResourcesConnectionThresholdOk returns a tuple with the LowResourcesConnectionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowResourcesConnectionThreshold

`func (o *HttpConnectionHandlerShared) SetLowResourcesConnectionThreshold(v int32)`

SetLowResourcesConnectionThreshold sets LowResourcesConnectionThreshold field to given value.

### HasLowResourcesConnectionThreshold

`func (o *HttpConnectionHandlerShared) HasLowResourcesConnectionThreshold() bool`

HasLowResourcesConnectionThreshold returns a boolean if a field has been set.

### GetLowResourcesIdleTimeLimit

`func (o *HttpConnectionHandlerShared) GetLowResourcesIdleTimeLimit() string`

GetLowResourcesIdleTimeLimit returns the LowResourcesIdleTimeLimit field if non-nil, zero value otherwise.

### GetLowResourcesIdleTimeLimitOk

`func (o *HttpConnectionHandlerShared) GetLowResourcesIdleTimeLimitOk() (*string, bool)`

GetLowResourcesIdleTimeLimitOk returns a tuple with the LowResourcesIdleTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowResourcesIdleTimeLimit

`func (o *HttpConnectionHandlerShared) SetLowResourcesIdleTimeLimit(v string)`

SetLowResourcesIdleTimeLimit sets LowResourcesIdleTimeLimit field to given value.

### HasLowResourcesIdleTimeLimit

`func (o *HttpConnectionHandlerShared) HasLowResourcesIdleTimeLimit() bool`

HasLowResourcesIdleTimeLimit returns a boolean if a field has been set.

### GetEnableMultipartMIMEParameters

`func (o *HttpConnectionHandlerShared) GetEnableMultipartMIMEParameters() bool`

GetEnableMultipartMIMEParameters returns the EnableMultipartMIMEParameters field if non-nil, zero value otherwise.

### GetEnableMultipartMIMEParametersOk

`func (o *HttpConnectionHandlerShared) GetEnableMultipartMIMEParametersOk() (*bool, bool)`

GetEnableMultipartMIMEParametersOk returns a tuple with the EnableMultipartMIMEParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultipartMIMEParameters

`func (o *HttpConnectionHandlerShared) SetEnableMultipartMIMEParameters(v bool)`

SetEnableMultipartMIMEParameters sets EnableMultipartMIMEParameters field to given value.

### HasEnableMultipartMIMEParameters

`func (o *HttpConnectionHandlerShared) HasEnableMultipartMIMEParameters() bool`

HasEnableMultipartMIMEParameters returns a boolean if a field has been set.

### GetUseForwardedHeaders

`func (o *HttpConnectionHandlerShared) GetUseForwardedHeaders() bool`

GetUseForwardedHeaders returns the UseForwardedHeaders field if non-nil, zero value otherwise.

### GetUseForwardedHeadersOk

`func (o *HttpConnectionHandlerShared) GetUseForwardedHeadersOk() (*bool, bool)`

GetUseForwardedHeadersOk returns a tuple with the UseForwardedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardedHeaders

`func (o *HttpConnectionHandlerShared) SetUseForwardedHeaders(v bool)`

SetUseForwardedHeaders sets UseForwardedHeaders field to given value.

### HasUseForwardedHeaders

`func (o *HttpConnectionHandlerShared) HasUseForwardedHeaders() bool`

HasUseForwardedHeaders returns a boolean if a field has been set.

### GetHttpRequestHeaderSize

`func (o *HttpConnectionHandlerShared) GetHttpRequestHeaderSize() int32`

GetHttpRequestHeaderSize returns the HttpRequestHeaderSize field if non-nil, zero value otherwise.

### GetHttpRequestHeaderSizeOk

`func (o *HttpConnectionHandlerShared) GetHttpRequestHeaderSizeOk() (*int32, bool)`

GetHttpRequestHeaderSizeOk returns a tuple with the HttpRequestHeaderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestHeaderSize

`func (o *HttpConnectionHandlerShared) SetHttpRequestHeaderSize(v int32)`

SetHttpRequestHeaderSize sets HttpRequestHeaderSize field to given value.

### HasHttpRequestHeaderSize

`func (o *HttpConnectionHandlerShared) HasHttpRequestHeaderSize() bool`

HasHttpRequestHeaderSize returns a boolean if a field has been set.

### GetResponseHeader

`func (o *HttpConnectionHandlerShared) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *HttpConnectionHandlerShared) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *HttpConnectionHandlerShared) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *HttpConnectionHandlerShared) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetUseCorrelationIDHeader

`func (o *HttpConnectionHandlerShared) GetUseCorrelationIDHeader() bool`

GetUseCorrelationIDHeader returns the UseCorrelationIDHeader field if non-nil, zero value otherwise.

### GetUseCorrelationIDHeaderOk

`func (o *HttpConnectionHandlerShared) GetUseCorrelationIDHeaderOk() (*bool, bool)`

GetUseCorrelationIDHeaderOk returns a tuple with the UseCorrelationIDHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCorrelationIDHeader

`func (o *HttpConnectionHandlerShared) SetUseCorrelationIDHeader(v bool)`

SetUseCorrelationIDHeader sets UseCorrelationIDHeader field to given value.

### HasUseCorrelationIDHeader

`func (o *HttpConnectionHandlerShared) HasUseCorrelationIDHeader() bool`

HasUseCorrelationIDHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *HttpConnectionHandlerShared) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *HttpConnectionHandlerShared) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *HttpConnectionHandlerShared) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *HttpConnectionHandlerShared) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDRequestHeader

`func (o *HttpConnectionHandlerShared) GetCorrelationIDRequestHeader() []string`

GetCorrelationIDRequestHeader returns the CorrelationIDRequestHeader field if non-nil, zero value otherwise.

### GetCorrelationIDRequestHeaderOk

`func (o *HttpConnectionHandlerShared) GetCorrelationIDRequestHeaderOk() (*[]string, bool)`

GetCorrelationIDRequestHeaderOk returns a tuple with the CorrelationIDRequestHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDRequestHeader

`func (o *HttpConnectionHandlerShared) SetCorrelationIDRequestHeader(v []string)`

SetCorrelationIDRequestHeader sets CorrelationIDRequestHeader field to given value.

### HasCorrelationIDRequestHeader

`func (o *HttpConnectionHandlerShared) HasCorrelationIDRequestHeader() bool`

HasCorrelationIDRequestHeader returns a boolean if a field has been set.

### GetSslClientAuthPolicy

`func (o *HttpConnectionHandlerShared) GetSslClientAuthPolicy() EnumconnectionHandlerSslClientAuthPolicyProp`

GetSslClientAuthPolicy returns the SslClientAuthPolicy field if non-nil, zero value otherwise.

### GetSslClientAuthPolicyOk

`func (o *HttpConnectionHandlerShared) GetSslClientAuthPolicyOk() (*EnumconnectionHandlerSslClientAuthPolicyProp, bool)`

GetSslClientAuthPolicyOk returns a tuple with the SslClientAuthPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslClientAuthPolicy

`func (o *HttpConnectionHandlerShared) SetSslClientAuthPolicy(v EnumconnectionHandlerSslClientAuthPolicyProp)`

SetSslClientAuthPolicy sets SslClientAuthPolicy field to given value.

### HasSslClientAuthPolicy

`func (o *HttpConnectionHandlerShared) HasSslClientAuthPolicy() bool`

HasSslClientAuthPolicy returns a boolean if a field has been set.

### GetDescription

`func (o *HttpConnectionHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HttpConnectionHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HttpConnectionHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HttpConnectionHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *HttpConnectionHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HttpConnectionHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HttpConnectionHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowedClient

`func (o *HttpConnectionHandlerShared) GetAllowedClient() []string`

GetAllowedClient returns the AllowedClient field if non-nil, zero value otherwise.

### GetAllowedClientOk

`func (o *HttpConnectionHandlerShared) GetAllowedClientOk() (*[]string, bool)`

GetAllowedClientOk returns a tuple with the AllowedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClient

`func (o *HttpConnectionHandlerShared) SetAllowedClient(v []string)`

SetAllowedClient sets AllowedClient field to given value.

### HasAllowedClient

`func (o *HttpConnectionHandlerShared) HasAllowedClient() bool`

HasAllowedClient returns a boolean if a field has been set.

### GetDeniedClient

`func (o *HttpConnectionHandlerShared) GetDeniedClient() []string`

GetDeniedClient returns the DeniedClient field if non-nil, zero value otherwise.

### GetDeniedClientOk

`func (o *HttpConnectionHandlerShared) GetDeniedClientOk() (*[]string, bool)`

GetDeniedClientOk returns a tuple with the DeniedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedClient

`func (o *HttpConnectionHandlerShared) SetDeniedClient(v []string)`

SetDeniedClient sets DeniedClient field to given value.

### HasDeniedClient

`func (o *HttpConnectionHandlerShared) HasDeniedClient() bool`

HasDeniedClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


