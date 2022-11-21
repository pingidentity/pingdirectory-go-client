# HttpConnectionHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Connection Handler | 
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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewHttpConnectionHandlerResponse

`func NewHttpConnectionHandlerResponse(id string, schemas []EnumhttpConnectionHandlerSchemaUrn, listenPort int32, enabled bool, ) *HttpConnectionHandlerResponse`

NewHttpConnectionHandlerResponse instantiates a new HttpConnectionHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpConnectionHandlerResponseWithDefaults

`func NewHttpConnectionHandlerResponseWithDefaults() *HttpConnectionHandlerResponse`

NewHttpConnectionHandlerResponseWithDefaults instantiates a new HttpConnectionHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HttpConnectionHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HttpConnectionHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HttpConnectionHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *HttpConnectionHandlerResponse) GetSchemas() []EnumhttpConnectionHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *HttpConnectionHandlerResponse) GetSchemasOk() (*[]EnumhttpConnectionHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *HttpConnectionHandlerResponse) SetSchemas(v []EnumhttpConnectionHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetListenAddress

`func (o *HttpConnectionHandlerResponse) GetListenAddress() string`

GetListenAddress returns the ListenAddress field if non-nil, zero value otherwise.

### GetListenAddressOk

`func (o *HttpConnectionHandlerResponse) GetListenAddressOk() (*string, bool)`

GetListenAddressOk returns a tuple with the ListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddress

`func (o *HttpConnectionHandlerResponse) SetListenAddress(v string)`

SetListenAddress sets ListenAddress field to given value.

### HasListenAddress

`func (o *HttpConnectionHandlerResponse) HasListenAddress() bool`

HasListenAddress returns a boolean if a field has been set.

### GetListenPort

`func (o *HttpConnectionHandlerResponse) GetListenPort() int32`

GetListenPort returns the ListenPort field if non-nil, zero value otherwise.

### GetListenPortOk

`func (o *HttpConnectionHandlerResponse) GetListenPortOk() (*int32, bool)`

GetListenPortOk returns a tuple with the ListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPort

`func (o *HttpConnectionHandlerResponse) SetListenPort(v int32)`

SetListenPort sets ListenPort field to given value.


### GetUseSSL

`func (o *HttpConnectionHandlerResponse) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *HttpConnectionHandlerResponse) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *HttpConnectionHandlerResponse) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *HttpConnectionHandlerResponse) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *HttpConnectionHandlerResponse) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *HttpConnectionHandlerResponse) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *HttpConnectionHandlerResponse) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.

### HasSslCertNickname

`func (o *HttpConnectionHandlerResponse) HasSslCertNickname() bool`

HasSslCertNickname returns a boolean if a field has been set.

### GetHttpServletExtension

`func (o *HttpConnectionHandlerResponse) GetHttpServletExtension() []string`

GetHttpServletExtension returns the HttpServletExtension field if non-nil, zero value otherwise.

### GetHttpServletExtensionOk

`func (o *HttpConnectionHandlerResponse) GetHttpServletExtensionOk() (*[]string, bool)`

GetHttpServletExtensionOk returns a tuple with the HttpServletExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServletExtension

`func (o *HttpConnectionHandlerResponse) SetHttpServletExtension(v []string)`

SetHttpServletExtension sets HttpServletExtension field to given value.

### HasHttpServletExtension

`func (o *HttpConnectionHandlerResponse) HasHttpServletExtension() bool`

HasHttpServletExtension returns a boolean if a field has been set.

### GetWebApplicationExtension

`func (o *HttpConnectionHandlerResponse) GetWebApplicationExtension() []string`

GetWebApplicationExtension returns the WebApplicationExtension field if non-nil, zero value otherwise.

### GetWebApplicationExtensionOk

`func (o *HttpConnectionHandlerResponse) GetWebApplicationExtensionOk() (*[]string, bool)`

GetWebApplicationExtensionOk returns a tuple with the WebApplicationExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebApplicationExtension

`func (o *HttpConnectionHandlerResponse) SetWebApplicationExtension(v []string)`

SetWebApplicationExtension sets WebApplicationExtension field to given value.

### HasWebApplicationExtension

`func (o *HttpConnectionHandlerResponse) HasWebApplicationExtension() bool`

HasWebApplicationExtension returns a boolean if a field has been set.

### GetHttpOperationLogPublisher

`func (o *HttpConnectionHandlerResponse) GetHttpOperationLogPublisher() []string`

GetHttpOperationLogPublisher returns the HttpOperationLogPublisher field if non-nil, zero value otherwise.

### GetHttpOperationLogPublisherOk

`func (o *HttpConnectionHandlerResponse) GetHttpOperationLogPublisherOk() (*[]string, bool)`

GetHttpOperationLogPublisherOk returns a tuple with the HttpOperationLogPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpOperationLogPublisher

`func (o *HttpConnectionHandlerResponse) SetHttpOperationLogPublisher(v []string)`

SetHttpOperationLogPublisher sets HttpOperationLogPublisher field to given value.

### HasHttpOperationLogPublisher

`func (o *HttpConnectionHandlerResponse) HasHttpOperationLogPublisher() bool`

HasHttpOperationLogPublisher returns a boolean if a field has been set.

### GetSslProtocol

`func (o *HttpConnectionHandlerResponse) GetSslProtocol() []string`

GetSslProtocol returns the SslProtocol field if non-nil, zero value otherwise.

### GetSslProtocolOk

`func (o *HttpConnectionHandlerResponse) GetSslProtocolOk() (*[]string, bool)`

GetSslProtocolOk returns a tuple with the SslProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslProtocol

`func (o *HttpConnectionHandlerResponse) SetSslProtocol(v []string)`

SetSslProtocol sets SslProtocol field to given value.

### HasSslProtocol

`func (o *HttpConnectionHandlerResponse) HasSslProtocol() bool`

HasSslProtocol returns a boolean if a field has been set.

### GetSslCipherSuite

`func (o *HttpConnectionHandlerResponse) GetSslCipherSuite() []string`

GetSslCipherSuite returns the SslCipherSuite field if non-nil, zero value otherwise.

### GetSslCipherSuiteOk

`func (o *HttpConnectionHandlerResponse) GetSslCipherSuiteOk() (*[]string, bool)`

GetSslCipherSuiteOk returns a tuple with the SslCipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCipherSuite

`func (o *HttpConnectionHandlerResponse) SetSslCipherSuite(v []string)`

SetSslCipherSuite sets SslCipherSuite field to given value.

### HasSslCipherSuite

`func (o *HttpConnectionHandlerResponse) HasSslCipherSuite() bool`

HasSslCipherSuite returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *HttpConnectionHandlerResponse) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *HttpConnectionHandlerResponse) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *HttpConnectionHandlerResponse) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *HttpConnectionHandlerResponse) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *HttpConnectionHandlerResponse) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *HttpConnectionHandlerResponse) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *HttpConnectionHandlerResponse) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *HttpConnectionHandlerResponse) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetNumRequestHandlers

`func (o *HttpConnectionHandlerResponse) GetNumRequestHandlers() int32`

GetNumRequestHandlers returns the NumRequestHandlers field if non-nil, zero value otherwise.

### GetNumRequestHandlersOk

`func (o *HttpConnectionHandlerResponse) GetNumRequestHandlersOk() (*int32, bool)`

GetNumRequestHandlersOk returns a tuple with the NumRequestHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRequestHandlers

`func (o *HttpConnectionHandlerResponse) SetNumRequestHandlers(v int32)`

SetNumRequestHandlers sets NumRequestHandlers field to given value.

### HasNumRequestHandlers

`func (o *HttpConnectionHandlerResponse) HasNumRequestHandlers() bool`

HasNumRequestHandlers returns a boolean if a field has been set.

### GetKeepStats

`func (o *HttpConnectionHandlerResponse) GetKeepStats() bool`

GetKeepStats returns the KeepStats field if non-nil, zero value otherwise.

### GetKeepStatsOk

`func (o *HttpConnectionHandlerResponse) GetKeepStatsOk() (*bool, bool)`

GetKeepStatsOk returns a tuple with the KeepStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepStats

`func (o *HttpConnectionHandlerResponse) SetKeepStats(v bool)`

SetKeepStats sets KeepStats field to given value.

### HasKeepStats

`func (o *HttpConnectionHandlerResponse) HasKeepStats() bool`

HasKeepStats returns a boolean if a field has been set.

### GetAcceptBacklog

`func (o *HttpConnectionHandlerResponse) GetAcceptBacklog() int32`

GetAcceptBacklog returns the AcceptBacklog field if non-nil, zero value otherwise.

### GetAcceptBacklogOk

`func (o *HttpConnectionHandlerResponse) GetAcceptBacklogOk() (*int32, bool)`

GetAcceptBacklogOk returns a tuple with the AcceptBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptBacklog

`func (o *HttpConnectionHandlerResponse) SetAcceptBacklog(v int32)`

SetAcceptBacklog sets AcceptBacklog field to given value.

### HasAcceptBacklog

`func (o *HttpConnectionHandlerResponse) HasAcceptBacklog() bool`

HasAcceptBacklog returns a boolean if a field has been set.

### GetAllowTCPReuseAddress

`func (o *HttpConnectionHandlerResponse) GetAllowTCPReuseAddress() bool`

GetAllowTCPReuseAddress returns the AllowTCPReuseAddress field if non-nil, zero value otherwise.

### GetAllowTCPReuseAddressOk

`func (o *HttpConnectionHandlerResponse) GetAllowTCPReuseAddressOk() (*bool, bool)`

GetAllowTCPReuseAddressOk returns a tuple with the AllowTCPReuseAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTCPReuseAddress

`func (o *HttpConnectionHandlerResponse) SetAllowTCPReuseAddress(v bool)`

SetAllowTCPReuseAddress sets AllowTCPReuseAddress field to given value.

### HasAllowTCPReuseAddress

`func (o *HttpConnectionHandlerResponse) HasAllowTCPReuseAddress() bool`

HasAllowTCPReuseAddress returns a boolean if a field has been set.

### GetIdleTimeLimit

`func (o *HttpConnectionHandlerResponse) GetIdleTimeLimit() string`

GetIdleTimeLimit returns the IdleTimeLimit field if non-nil, zero value otherwise.

### GetIdleTimeLimitOk

`func (o *HttpConnectionHandlerResponse) GetIdleTimeLimitOk() (*string, bool)`

GetIdleTimeLimitOk returns a tuple with the IdleTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeLimit

`func (o *HttpConnectionHandlerResponse) SetIdleTimeLimit(v string)`

SetIdleTimeLimit sets IdleTimeLimit field to given value.

### HasIdleTimeLimit

`func (o *HttpConnectionHandlerResponse) HasIdleTimeLimit() bool`

HasIdleTimeLimit returns a boolean if a field has been set.

### GetLowResourcesConnectionThreshold

`func (o *HttpConnectionHandlerResponse) GetLowResourcesConnectionThreshold() int32`

GetLowResourcesConnectionThreshold returns the LowResourcesConnectionThreshold field if non-nil, zero value otherwise.

### GetLowResourcesConnectionThresholdOk

`func (o *HttpConnectionHandlerResponse) GetLowResourcesConnectionThresholdOk() (*int32, bool)`

GetLowResourcesConnectionThresholdOk returns a tuple with the LowResourcesConnectionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowResourcesConnectionThreshold

`func (o *HttpConnectionHandlerResponse) SetLowResourcesConnectionThreshold(v int32)`

SetLowResourcesConnectionThreshold sets LowResourcesConnectionThreshold field to given value.

### HasLowResourcesConnectionThreshold

`func (o *HttpConnectionHandlerResponse) HasLowResourcesConnectionThreshold() bool`

HasLowResourcesConnectionThreshold returns a boolean if a field has been set.

### GetLowResourcesIdleTimeLimit

`func (o *HttpConnectionHandlerResponse) GetLowResourcesIdleTimeLimit() string`

GetLowResourcesIdleTimeLimit returns the LowResourcesIdleTimeLimit field if non-nil, zero value otherwise.

### GetLowResourcesIdleTimeLimitOk

`func (o *HttpConnectionHandlerResponse) GetLowResourcesIdleTimeLimitOk() (*string, bool)`

GetLowResourcesIdleTimeLimitOk returns a tuple with the LowResourcesIdleTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowResourcesIdleTimeLimit

`func (o *HttpConnectionHandlerResponse) SetLowResourcesIdleTimeLimit(v string)`

SetLowResourcesIdleTimeLimit sets LowResourcesIdleTimeLimit field to given value.

### HasLowResourcesIdleTimeLimit

`func (o *HttpConnectionHandlerResponse) HasLowResourcesIdleTimeLimit() bool`

HasLowResourcesIdleTimeLimit returns a boolean if a field has been set.

### GetEnableMultipartMIMEParameters

`func (o *HttpConnectionHandlerResponse) GetEnableMultipartMIMEParameters() bool`

GetEnableMultipartMIMEParameters returns the EnableMultipartMIMEParameters field if non-nil, zero value otherwise.

### GetEnableMultipartMIMEParametersOk

`func (o *HttpConnectionHandlerResponse) GetEnableMultipartMIMEParametersOk() (*bool, bool)`

GetEnableMultipartMIMEParametersOk returns a tuple with the EnableMultipartMIMEParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultipartMIMEParameters

`func (o *HttpConnectionHandlerResponse) SetEnableMultipartMIMEParameters(v bool)`

SetEnableMultipartMIMEParameters sets EnableMultipartMIMEParameters field to given value.

### HasEnableMultipartMIMEParameters

`func (o *HttpConnectionHandlerResponse) HasEnableMultipartMIMEParameters() bool`

HasEnableMultipartMIMEParameters returns a boolean if a field has been set.

### GetUseForwardedHeaders

`func (o *HttpConnectionHandlerResponse) GetUseForwardedHeaders() bool`

GetUseForwardedHeaders returns the UseForwardedHeaders field if non-nil, zero value otherwise.

### GetUseForwardedHeadersOk

`func (o *HttpConnectionHandlerResponse) GetUseForwardedHeadersOk() (*bool, bool)`

GetUseForwardedHeadersOk returns a tuple with the UseForwardedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardedHeaders

`func (o *HttpConnectionHandlerResponse) SetUseForwardedHeaders(v bool)`

SetUseForwardedHeaders sets UseForwardedHeaders field to given value.

### HasUseForwardedHeaders

`func (o *HttpConnectionHandlerResponse) HasUseForwardedHeaders() bool`

HasUseForwardedHeaders returns a boolean if a field has been set.

### GetHttpRequestHeaderSize

`func (o *HttpConnectionHandlerResponse) GetHttpRequestHeaderSize() int32`

GetHttpRequestHeaderSize returns the HttpRequestHeaderSize field if non-nil, zero value otherwise.

### GetHttpRequestHeaderSizeOk

`func (o *HttpConnectionHandlerResponse) GetHttpRequestHeaderSizeOk() (*int32, bool)`

GetHttpRequestHeaderSizeOk returns a tuple with the HttpRequestHeaderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestHeaderSize

`func (o *HttpConnectionHandlerResponse) SetHttpRequestHeaderSize(v int32)`

SetHttpRequestHeaderSize sets HttpRequestHeaderSize field to given value.

### HasHttpRequestHeaderSize

`func (o *HttpConnectionHandlerResponse) HasHttpRequestHeaderSize() bool`

HasHttpRequestHeaderSize returns a boolean if a field has been set.

### GetResponseHeader

`func (o *HttpConnectionHandlerResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *HttpConnectionHandlerResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *HttpConnectionHandlerResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *HttpConnectionHandlerResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetUseCorrelationIDHeader

`func (o *HttpConnectionHandlerResponse) GetUseCorrelationIDHeader() bool`

GetUseCorrelationIDHeader returns the UseCorrelationIDHeader field if non-nil, zero value otherwise.

### GetUseCorrelationIDHeaderOk

`func (o *HttpConnectionHandlerResponse) GetUseCorrelationIDHeaderOk() (*bool, bool)`

GetUseCorrelationIDHeaderOk returns a tuple with the UseCorrelationIDHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCorrelationIDHeader

`func (o *HttpConnectionHandlerResponse) SetUseCorrelationIDHeader(v bool)`

SetUseCorrelationIDHeader sets UseCorrelationIDHeader field to given value.

### HasUseCorrelationIDHeader

`func (o *HttpConnectionHandlerResponse) HasUseCorrelationIDHeader() bool`

HasUseCorrelationIDHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *HttpConnectionHandlerResponse) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *HttpConnectionHandlerResponse) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *HttpConnectionHandlerResponse) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *HttpConnectionHandlerResponse) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDRequestHeader

`func (o *HttpConnectionHandlerResponse) GetCorrelationIDRequestHeader() []string`

GetCorrelationIDRequestHeader returns the CorrelationIDRequestHeader field if non-nil, zero value otherwise.

### GetCorrelationIDRequestHeaderOk

`func (o *HttpConnectionHandlerResponse) GetCorrelationIDRequestHeaderOk() (*[]string, bool)`

GetCorrelationIDRequestHeaderOk returns a tuple with the CorrelationIDRequestHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDRequestHeader

`func (o *HttpConnectionHandlerResponse) SetCorrelationIDRequestHeader(v []string)`

SetCorrelationIDRequestHeader sets CorrelationIDRequestHeader field to given value.

### HasCorrelationIDRequestHeader

`func (o *HttpConnectionHandlerResponse) HasCorrelationIDRequestHeader() bool`

HasCorrelationIDRequestHeader returns a boolean if a field has been set.

### GetSslClientAuthPolicy

`func (o *HttpConnectionHandlerResponse) GetSslClientAuthPolicy() EnumconnectionHandlerSslClientAuthPolicyProp`

GetSslClientAuthPolicy returns the SslClientAuthPolicy field if non-nil, zero value otherwise.

### GetSslClientAuthPolicyOk

`func (o *HttpConnectionHandlerResponse) GetSslClientAuthPolicyOk() (*EnumconnectionHandlerSslClientAuthPolicyProp, bool)`

GetSslClientAuthPolicyOk returns a tuple with the SslClientAuthPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslClientAuthPolicy

`func (o *HttpConnectionHandlerResponse) SetSslClientAuthPolicy(v EnumconnectionHandlerSslClientAuthPolicyProp)`

SetSslClientAuthPolicy sets SslClientAuthPolicy field to given value.

### HasSslClientAuthPolicy

`func (o *HttpConnectionHandlerResponse) HasSslClientAuthPolicy() bool`

HasSslClientAuthPolicy returns a boolean if a field has been set.

### GetDescription

`func (o *HttpConnectionHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HttpConnectionHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HttpConnectionHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HttpConnectionHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *HttpConnectionHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HttpConnectionHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HttpConnectionHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowedClient

`func (o *HttpConnectionHandlerResponse) GetAllowedClient() []string`

GetAllowedClient returns the AllowedClient field if non-nil, zero value otherwise.

### GetAllowedClientOk

`func (o *HttpConnectionHandlerResponse) GetAllowedClientOk() (*[]string, bool)`

GetAllowedClientOk returns a tuple with the AllowedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClient

`func (o *HttpConnectionHandlerResponse) SetAllowedClient(v []string)`

SetAllowedClient sets AllowedClient field to given value.

### HasAllowedClient

`func (o *HttpConnectionHandlerResponse) HasAllowedClient() bool`

HasAllowedClient returns a boolean if a field has been set.

### GetDeniedClient

`func (o *HttpConnectionHandlerResponse) GetDeniedClient() []string`

GetDeniedClient returns the DeniedClient field if non-nil, zero value otherwise.

### GetDeniedClientOk

`func (o *HttpConnectionHandlerResponse) GetDeniedClientOk() (*[]string, bool)`

GetDeniedClientOk returns a tuple with the DeniedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedClient

`func (o *HttpConnectionHandlerResponse) SetDeniedClient(v []string)`

SetDeniedClient sets DeniedClient field to given value.

### HasDeniedClient

`func (o *HttpConnectionHandlerResponse) HasDeniedClient() bool`

HasDeniedClient returns a boolean if a field has been set.

### GetMeta

`func (o *HttpConnectionHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HttpConnectionHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HttpConnectionHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HttpConnectionHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


