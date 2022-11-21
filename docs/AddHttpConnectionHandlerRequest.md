# AddHttpConnectionHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new Connection Handler | 
**Schemas** | [**[]EnumhttpConnectionHandlerSchemaUrn**](EnumhttpConnectionHandlerSchemaUrn.md) |  | 
**ListenAddress** | Pointer to **string** | Specifies the address on which to listen for connections from HTTP clients. If no value is defined, the server will listen on all addresses on all interfaces. | [optional] 
**ListenPort** | **int32** | Specifies the port number on which the HTTP Connection Handler will listen for connections from clients. | 
**UseSSL** | Pointer to **bool** | Indicates whether the HTTP Connection Handler should use SSL. | [optional] 
**SslCertNickname** | Pointer to **string** | Specifies the nickname (also called the alias) of the certificate that the HTTP Connection Handler should use when performing SSL communication. | [optional] 
**HttpServletExtension** | Pointer to **[]string** | Specifies information about servlets that will be provided via this connection handler. | [optional] 
**WebApplicationExtension** | Pointer to **[]string** | Specifies information about web applications that will be provided via this connection handler. | [optional] 
**HttpOperationLogPublisher** | Pointer to **[]string** | Specifies the set of HTTP operation loggers that should be used to log information about requests and responses for operations processed through this HTTP Connection Handler. | [optional] 
**SslProtocol** | Pointer to **[]string** | Specifies the names of the SSL protocols that are allowed for use in SSL communication. The set of supported ssl protocols can be viewed via the ssl context monitor entry. | [optional] 
**SslCipherSuite** | Pointer to **[]string** | Specifies the names of the SSL cipher suites that are allowed for use in SSL communication. The set of supported cipher suites can be viewed via the ssl context monitor entry. | [optional] 
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
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**UseCorrelationIDHeader** | Pointer to **bool** | If enabled, a correlation ID header will be added to outgoing HTTP responses. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 
**CorrelationIDRequestHeader** | Pointer to **[]string** | Specifies the set of HTTP request headers that may contain a value to be used as the correlation ID. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 
**SslClientAuthPolicy** | Pointer to [**EnumconnectionHandlerSslClientAuthPolicyProp**](EnumconnectionHandlerSslClientAuthPolicyProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Connection Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Connection Handler is enabled. | 
**AllowedClient** | Pointer to **[]string** | Specifies a set of address masks that determines the addresses of the clients that are allowed to establish connections to this connection handler. | [optional] 
**DeniedClient** | Pointer to **[]string** | Specifies a set of address masks that determines the addresses of the clients that are not allowed to establish connections to this connection handler. | [optional] 

## Methods

### NewAddHttpConnectionHandlerRequest

`func NewAddHttpConnectionHandlerRequest(handlerName string, schemas []EnumhttpConnectionHandlerSchemaUrn, listenPort int32, enabled bool, ) *AddHttpConnectionHandlerRequest`

NewAddHttpConnectionHandlerRequest instantiates a new AddHttpConnectionHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddHttpConnectionHandlerRequestWithDefaults

`func NewAddHttpConnectionHandlerRequestWithDefaults() *AddHttpConnectionHandlerRequest`

NewAddHttpConnectionHandlerRequestWithDefaults instantiates a new AddHttpConnectionHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddHttpConnectionHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddHttpConnectionHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddHttpConnectionHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddHttpConnectionHandlerRequest) GetSchemas() []EnumhttpConnectionHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddHttpConnectionHandlerRequest) GetSchemasOk() (*[]EnumhttpConnectionHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddHttpConnectionHandlerRequest) SetSchemas(v []EnumhttpConnectionHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetListenAddress

`func (o *AddHttpConnectionHandlerRequest) GetListenAddress() string`

GetListenAddress returns the ListenAddress field if non-nil, zero value otherwise.

### GetListenAddressOk

`func (o *AddHttpConnectionHandlerRequest) GetListenAddressOk() (*string, bool)`

GetListenAddressOk returns a tuple with the ListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddress

`func (o *AddHttpConnectionHandlerRequest) SetListenAddress(v string)`

SetListenAddress sets ListenAddress field to given value.

### HasListenAddress

`func (o *AddHttpConnectionHandlerRequest) HasListenAddress() bool`

HasListenAddress returns a boolean if a field has been set.

### GetListenPort

`func (o *AddHttpConnectionHandlerRequest) GetListenPort() int32`

GetListenPort returns the ListenPort field if non-nil, zero value otherwise.

### GetListenPortOk

`func (o *AddHttpConnectionHandlerRequest) GetListenPortOk() (*int32, bool)`

GetListenPortOk returns a tuple with the ListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPort

`func (o *AddHttpConnectionHandlerRequest) SetListenPort(v int32)`

SetListenPort sets ListenPort field to given value.


### GetUseSSL

`func (o *AddHttpConnectionHandlerRequest) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *AddHttpConnectionHandlerRequest) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *AddHttpConnectionHandlerRequest) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *AddHttpConnectionHandlerRequest) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *AddHttpConnectionHandlerRequest) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *AddHttpConnectionHandlerRequest) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *AddHttpConnectionHandlerRequest) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.

### HasSslCertNickname

`func (o *AddHttpConnectionHandlerRequest) HasSslCertNickname() bool`

HasSslCertNickname returns a boolean if a field has been set.

### GetHttpServletExtension

`func (o *AddHttpConnectionHandlerRequest) GetHttpServletExtension() []string`

GetHttpServletExtension returns the HttpServletExtension field if non-nil, zero value otherwise.

### GetHttpServletExtensionOk

`func (o *AddHttpConnectionHandlerRequest) GetHttpServletExtensionOk() (*[]string, bool)`

GetHttpServletExtensionOk returns a tuple with the HttpServletExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServletExtension

`func (o *AddHttpConnectionHandlerRequest) SetHttpServletExtension(v []string)`

SetHttpServletExtension sets HttpServletExtension field to given value.

### HasHttpServletExtension

`func (o *AddHttpConnectionHandlerRequest) HasHttpServletExtension() bool`

HasHttpServletExtension returns a boolean if a field has been set.

### GetWebApplicationExtension

`func (o *AddHttpConnectionHandlerRequest) GetWebApplicationExtension() []string`

GetWebApplicationExtension returns the WebApplicationExtension field if non-nil, zero value otherwise.

### GetWebApplicationExtensionOk

`func (o *AddHttpConnectionHandlerRequest) GetWebApplicationExtensionOk() (*[]string, bool)`

GetWebApplicationExtensionOk returns a tuple with the WebApplicationExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebApplicationExtension

`func (o *AddHttpConnectionHandlerRequest) SetWebApplicationExtension(v []string)`

SetWebApplicationExtension sets WebApplicationExtension field to given value.

### HasWebApplicationExtension

`func (o *AddHttpConnectionHandlerRequest) HasWebApplicationExtension() bool`

HasWebApplicationExtension returns a boolean if a field has been set.

### GetHttpOperationLogPublisher

`func (o *AddHttpConnectionHandlerRequest) GetHttpOperationLogPublisher() []string`

GetHttpOperationLogPublisher returns the HttpOperationLogPublisher field if non-nil, zero value otherwise.

### GetHttpOperationLogPublisherOk

`func (o *AddHttpConnectionHandlerRequest) GetHttpOperationLogPublisherOk() (*[]string, bool)`

GetHttpOperationLogPublisherOk returns a tuple with the HttpOperationLogPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpOperationLogPublisher

`func (o *AddHttpConnectionHandlerRequest) SetHttpOperationLogPublisher(v []string)`

SetHttpOperationLogPublisher sets HttpOperationLogPublisher field to given value.

### HasHttpOperationLogPublisher

`func (o *AddHttpConnectionHandlerRequest) HasHttpOperationLogPublisher() bool`

HasHttpOperationLogPublisher returns a boolean if a field has been set.

### GetSslProtocol

`func (o *AddHttpConnectionHandlerRequest) GetSslProtocol() []string`

GetSslProtocol returns the SslProtocol field if non-nil, zero value otherwise.

### GetSslProtocolOk

`func (o *AddHttpConnectionHandlerRequest) GetSslProtocolOk() (*[]string, bool)`

GetSslProtocolOk returns a tuple with the SslProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslProtocol

`func (o *AddHttpConnectionHandlerRequest) SetSslProtocol(v []string)`

SetSslProtocol sets SslProtocol field to given value.

### HasSslProtocol

`func (o *AddHttpConnectionHandlerRequest) HasSslProtocol() bool`

HasSslProtocol returns a boolean if a field has been set.

### GetSslCipherSuite

`func (o *AddHttpConnectionHandlerRequest) GetSslCipherSuite() []string`

GetSslCipherSuite returns the SslCipherSuite field if non-nil, zero value otherwise.

### GetSslCipherSuiteOk

`func (o *AddHttpConnectionHandlerRequest) GetSslCipherSuiteOk() (*[]string, bool)`

GetSslCipherSuiteOk returns a tuple with the SslCipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCipherSuite

`func (o *AddHttpConnectionHandlerRequest) SetSslCipherSuite(v []string)`

SetSslCipherSuite sets SslCipherSuite field to given value.

### HasSslCipherSuite

`func (o *AddHttpConnectionHandlerRequest) HasSslCipherSuite() bool`

HasSslCipherSuite returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *AddHttpConnectionHandlerRequest) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *AddHttpConnectionHandlerRequest) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *AddHttpConnectionHandlerRequest) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *AddHttpConnectionHandlerRequest) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *AddHttpConnectionHandlerRequest) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *AddHttpConnectionHandlerRequest) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *AddHttpConnectionHandlerRequest) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *AddHttpConnectionHandlerRequest) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetNumRequestHandlers

`func (o *AddHttpConnectionHandlerRequest) GetNumRequestHandlers() int32`

GetNumRequestHandlers returns the NumRequestHandlers field if non-nil, zero value otherwise.

### GetNumRequestHandlersOk

`func (o *AddHttpConnectionHandlerRequest) GetNumRequestHandlersOk() (*int32, bool)`

GetNumRequestHandlersOk returns a tuple with the NumRequestHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRequestHandlers

`func (o *AddHttpConnectionHandlerRequest) SetNumRequestHandlers(v int32)`

SetNumRequestHandlers sets NumRequestHandlers field to given value.

### HasNumRequestHandlers

`func (o *AddHttpConnectionHandlerRequest) HasNumRequestHandlers() bool`

HasNumRequestHandlers returns a boolean if a field has been set.

### GetKeepStats

`func (o *AddHttpConnectionHandlerRequest) GetKeepStats() bool`

GetKeepStats returns the KeepStats field if non-nil, zero value otherwise.

### GetKeepStatsOk

`func (o *AddHttpConnectionHandlerRequest) GetKeepStatsOk() (*bool, bool)`

GetKeepStatsOk returns a tuple with the KeepStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepStats

`func (o *AddHttpConnectionHandlerRequest) SetKeepStats(v bool)`

SetKeepStats sets KeepStats field to given value.

### HasKeepStats

`func (o *AddHttpConnectionHandlerRequest) HasKeepStats() bool`

HasKeepStats returns a boolean if a field has been set.

### GetAcceptBacklog

`func (o *AddHttpConnectionHandlerRequest) GetAcceptBacklog() int32`

GetAcceptBacklog returns the AcceptBacklog field if non-nil, zero value otherwise.

### GetAcceptBacklogOk

`func (o *AddHttpConnectionHandlerRequest) GetAcceptBacklogOk() (*int32, bool)`

GetAcceptBacklogOk returns a tuple with the AcceptBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptBacklog

`func (o *AddHttpConnectionHandlerRequest) SetAcceptBacklog(v int32)`

SetAcceptBacklog sets AcceptBacklog field to given value.

### HasAcceptBacklog

`func (o *AddHttpConnectionHandlerRequest) HasAcceptBacklog() bool`

HasAcceptBacklog returns a boolean if a field has been set.

### GetAllowTCPReuseAddress

`func (o *AddHttpConnectionHandlerRequest) GetAllowTCPReuseAddress() bool`

GetAllowTCPReuseAddress returns the AllowTCPReuseAddress field if non-nil, zero value otherwise.

### GetAllowTCPReuseAddressOk

`func (o *AddHttpConnectionHandlerRequest) GetAllowTCPReuseAddressOk() (*bool, bool)`

GetAllowTCPReuseAddressOk returns a tuple with the AllowTCPReuseAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTCPReuseAddress

`func (o *AddHttpConnectionHandlerRequest) SetAllowTCPReuseAddress(v bool)`

SetAllowTCPReuseAddress sets AllowTCPReuseAddress field to given value.

### HasAllowTCPReuseAddress

`func (o *AddHttpConnectionHandlerRequest) HasAllowTCPReuseAddress() bool`

HasAllowTCPReuseAddress returns a boolean if a field has been set.

### GetIdleTimeLimit

`func (o *AddHttpConnectionHandlerRequest) GetIdleTimeLimit() string`

GetIdleTimeLimit returns the IdleTimeLimit field if non-nil, zero value otherwise.

### GetIdleTimeLimitOk

`func (o *AddHttpConnectionHandlerRequest) GetIdleTimeLimitOk() (*string, bool)`

GetIdleTimeLimitOk returns a tuple with the IdleTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeLimit

`func (o *AddHttpConnectionHandlerRequest) SetIdleTimeLimit(v string)`

SetIdleTimeLimit sets IdleTimeLimit field to given value.

### HasIdleTimeLimit

`func (o *AddHttpConnectionHandlerRequest) HasIdleTimeLimit() bool`

HasIdleTimeLimit returns a boolean if a field has been set.

### GetLowResourcesConnectionThreshold

`func (o *AddHttpConnectionHandlerRequest) GetLowResourcesConnectionThreshold() int32`

GetLowResourcesConnectionThreshold returns the LowResourcesConnectionThreshold field if non-nil, zero value otherwise.

### GetLowResourcesConnectionThresholdOk

`func (o *AddHttpConnectionHandlerRequest) GetLowResourcesConnectionThresholdOk() (*int32, bool)`

GetLowResourcesConnectionThresholdOk returns a tuple with the LowResourcesConnectionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowResourcesConnectionThreshold

`func (o *AddHttpConnectionHandlerRequest) SetLowResourcesConnectionThreshold(v int32)`

SetLowResourcesConnectionThreshold sets LowResourcesConnectionThreshold field to given value.

### HasLowResourcesConnectionThreshold

`func (o *AddHttpConnectionHandlerRequest) HasLowResourcesConnectionThreshold() bool`

HasLowResourcesConnectionThreshold returns a boolean if a field has been set.

### GetLowResourcesIdleTimeLimit

`func (o *AddHttpConnectionHandlerRequest) GetLowResourcesIdleTimeLimit() string`

GetLowResourcesIdleTimeLimit returns the LowResourcesIdleTimeLimit field if non-nil, zero value otherwise.

### GetLowResourcesIdleTimeLimitOk

`func (o *AddHttpConnectionHandlerRequest) GetLowResourcesIdleTimeLimitOk() (*string, bool)`

GetLowResourcesIdleTimeLimitOk returns a tuple with the LowResourcesIdleTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowResourcesIdleTimeLimit

`func (o *AddHttpConnectionHandlerRequest) SetLowResourcesIdleTimeLimit(v string)`

SetLowResourcesIdleTimeLimit sets LowResourcesIdleTimeLimit field to given value.

### HasLowResourcesIdleTimeLimit

`func (o *AddHttpConnectionHandlerRequest) HasLowResourcesIdleTimeLimit() bool`

HasLowResourcesIdleTimeLimit returns a boolean if a field has been set.

### GetEnableMultipartMIMEParameters

`func (o *AddHttpConnectionHandlerRequest) GetEnableMultipartMIMEParameters() bool`

GetEnableMultipartMIMEParameters returns the EnableMultipartMIMEParameters field if non-nil, zero value otherwise.

### GetEnableMultipartMIMEParametersOk

`func (o *AddHttpConnectionHandlerRequest) GetEnableMultipartMIMEParametersOk() (*bool, bool)`

GetEnableMultipartMIMEParametersOk returns a tuple with the EnableMultipartMIMEParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultipartMIMEParameters

`func (o *AddHttpConnectionHandlerRequest) SetEnableMultipartMIMEParameters(v bool)`

SetEnableMultipartMIMEParameters sets EnableMultipartMIMEParameters field to given value.

### HasEnableMultipartMIMEParameters

`func (o *AddHttpConnectionHandlerRequest) HasEnableMultipartMIMEParameters() bool`

HasEnableMultipartMIMEParameters returns a boolean if a field has been set.

### GetUseForwardedHeaders

`func (o *AddHttpConnectionHandlerRequest) GetUseForwardedHeaders() bool`

GetUseForwardedHeaders returns the UseForwardedHeaders field if non-nil, zero value otherwise.

### GetUseForwardedHeadersOk

`func (o *AddHttpConnectionHandlerRequest) GetUseForwardedHeadersOk() (*bool, bool)`

GetUseForwardedHeadersOk returns a tuple with the UseForwardedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardedHeaders

`func (o *AddHttpConnectionHandlerRequest) SetUseForwardedHeaders(v bool)`

SetUseForwardedHeaders sets UseForwardedHeaders field to given value.

### HasUseForwardedHeaders

`func (o *AddHttpConnectionHandlerRequest) HasUseForwardedHeaders() bool`

HasUseForwardedHeaders returns a boolean if a field has been set.

### GetHttpRequestHeaderSize

`func (o *AddHttpConnectionHandlerRequest) GetHttpRequestHeaderSize() int32`

GetHttpRequestHeaderSize returns the HttpRequestHeaderSize field if non-nil, zero value otherwise.

### GetHttpRequestHeaderSizeOk

`func (o *AddHttpConnectionHandlerRequest) GetHttpRequestHeaderSizeOk() (*int32, bool)`

GetHttpRequestHeaderSizeOk returns a tuple with the HttpRequestHeaderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestHeaderSize

`func (o *AddHttpConnectionHandlerRequest) SetHttpRequestHeaderSize(v int32)`

SetHttpRequestHeaderSize sets HttpRequestHeaderSize field to given value.

### HasHttpRequestHeaderSize

`func (o *AddHttpConnectionHandlerRequest) HasHttpRequestHeaderSize() bool`

HasHttpRequestHeaderSize returns a boolean if a field has been set.

### GetResponseHeader

`func (o *AddHttpConnectionHandlerRequest) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *AddHttpConnectionHandlerRequest) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *AddHttpConnectionHandlerRequest) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *AddHttpConnectionHandlerRequest) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetUseCorrelationIDHeader

`func (o *AddHttpConnectionHandlerRequest) GetUseCorrelationIDHeader() bool`

GetUseCorrelationIDHeader returns the UseCorrelationIDHeader field if non-nil, zero value otherwise.

### GetUseCorrelationIDHeaderOk

`func (o *AddHttpConnectionHandlerRequest) GetUseCorrelationIDHeaderOk() (*bool, bool)`

GetUseCorrelationIDHeaderOk returns a tuple with the UseCorrelationIDHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCorrelationIDHeader

`func (o *AddHttpConnectionHandlerRequest) SetUseCorrelationIDHeader(v bool)`

SetUseCorrelationIDHeader sets UseCorrelationIDHeader field to given value.

### HasUseCorrelationIDHeader

`func (o *AddHttpConnectionHandlerRequest) HasUseCorrelationIDHeader() bool`

HasUseCorrelationIDHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *AddHttpConnectionHandlerRequest) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *AddHttpConnectionHandlerRequest) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *AddHttpConnectionHandlerRequest) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *AddHttpConnectionHandlerRequest) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDRequestHeader

`func (o *AddHttpConnectionHandlerRequest) GetCorrelationIDRequestHeader() []string`

GetCorrelationIDRequestHeader returns the CorrelationIDRequestHeader field if non-nil, zero value otherwise.

### GetCorrelationIDRequestHeaderOk

`func (o *AddHttpConnectionHandlerRequest) GetCorrelationIDRequestHeaderOk() (*[]string, bool)`

GetCorrelationIDRequestHeaderOk returns a tuple with the CorrelationIDRequestHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDRequestHeader

`func (o *AddHttpConnectionHandlerRequest) SetCorrelationIDRequestHeader(v []string)`

SetCorrelationIDRequestHeader sets CorrelationIDRequestHeader field to given value.

### HasCorrelationIDRequestHeader

`func (o *AddHttpConnectionHandlerRequest) HasCorrelationIDRequestHeader() bool`

HasCorrelationIDRequestHeader returns a boolean if a field has been set.

### GetSslClientAuthPolicy

`func (o *AddHttpConnectionHandlerRequest) GetSslClientAuthPolicy() EnumconnectionHandlerSslClientAuthPolicyProp`

GetSslClientAuthPolicy returns the SslClientAuthPolicy field if non-nil, zero value otherwise.

### GetSslClientAuthPolicyOk

`func (o *AddHttpConnectionHandlerRequest) GetSslClientAuthPolicyOk() (*EnumconnectionHandlerSslClientAuthPolicyProp, bool)`

GetSslClientAuthPolicyOk returns a tuple with the SslClientAuthPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslClientAuthPolicy

`func (o *AddHttpConnectionHandlerRequest) SetSslClientAuthPolicy(v EnumconnectionHandlerSslClientAuthPolicyProp)`

SetSslClientAuthPolicy sets SslClientAuthPolicy field to given value.

### HasSslClientAuthPolicy

`func (o *AddHttpConnectionHandlerRequest) HasSslClientAuthPolicy() bool`

HasSslClientAuthPolicy returns a boolean if a field has been set.

### GetDescription

`func (o *AddHttpConnectionHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddHttpConnectionHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddHttpConnectionHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddHttpConnectionHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddHttpConnectionHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddHttpConnectionHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddHttpConnectionHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowedClient

`func (o *AddHttpConnectionHandlerRequest) GetAllowedClient() []string`

GetAllowedClient returns the AllowedClient field if non-nil, zero value otherwise.

### GetAllowedClientOk

`func (o *AddHttpConnectionHandlerRequest) GetAllowedClientOk() (*[]string, bool)`

GetAllowedClientOk returns a tuple with the AllowedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClient

`func (o *AddHttpConnectionHandlerRequest) SetAllowedClient(v []string)`

SetAllowedClient sets AllowedClient field to given value.

### HasAllowedClient

`func (o *AddHttpConnectionHandlerRequest) HasAllowedClient() bool`

HasAllowedClient returns a boolean if a field has been set.

### GetDeniedClient

`func (o *AddHttpConnectionHandlerRequest) GetDeniedClient() []string`

GetDeniedClient returns the DeniedClient field if non-nil, zero value otherwise.

### GetDeniedClientOk

`func (o *AddHttpConnectionHandlerRequest) GetDeniedClientOk() (*[]string, bool)`

GetDeniedClientOk returns a tuple with the DeniedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedClient

`func (o *AddHttpConnectionHandlerRequest) SetDeniedClient(v []string)`

SetDeniedClient sets DeniedClient field to given value.

### HasDeniedClient

`func (o *AddHttpConnectionHandlerRequest) HasDeniedClient() bool`

HasDeniedClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


