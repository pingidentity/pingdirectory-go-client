# ConsoleJsonHttpOperationLogPublisherResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Log Publisher | 
**Schemas** | [**[]EnumconsoleJsonHttpOperationLogPublisherSchemaUrn**](EnumconsoleJsonHttpOperationLogPublisherSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicates whether the Console JSON HTTP Operation Log Publisher is enabled for use. | 
**OutputLocation** | Pointer to [**EnumlogPublisherOutputLocationProp**](EnumlogPublisherOutputLocationProp.md) |  | [optional] 
**LogRequests** | Pointer to **bool** | Indicates whether to record a log message with information about requests received from the client. | [optional] 
**LogResults** | Pointer to **bool** | Indicates whether to record a log message with information about the result of processing a requested HTTP operation. | [optional] 
**IncludeProductName** | Pointer to **bool** | Indicates whether log messages should include the product name for the Directory Server. | [optional] 
**IncludeInstanceName** | Pointer to **bool** | Indicates whether log messages should include the instance name for the Directory Server. | [optional] 
**IncludeStartupID** | Pointer to **bool** | Indicates whether log messages should include the startup ID for the Directory Server, which is a value assigned to the server instance at startup and may be used to identify when the server has been restarted. | [optional] 
**IncludeThreadID** | Pointer to **bool** | Indicates whether log messages should include the thread ID for the Directory Server in each log message. This ID can be used to correlate log messages from the same thread within a single log as well as generated by the same thread across different types of log files. More information about the thread with a specific ID can be obtained using the cn&#x3D;JVM Stack Trace,cn&#x3D;monitor entry. | [optional] 
**IncludeRequestDetailsInResultMessages** | Pointer to **bool** | Indicates whether result log messages should include all of the elements of request log messages. This may be used to record a single message per operation with details about both the request and response. | [optional] 
**LogRequestHeaders** | Pointer to [**EnumlogPublisherLogRequestHeadersProp**](EnumlogPublisherLogRequestHeadersProp.md) |  | [optional] 
**SuppressedRequestHeaderName** | Pointer to **[]string** | Specifies the case-insensitive names of request headers that should be omitted from log messages (e.g., for the purpose of brevity or security). This will only be used if the log-request-headers property has a value of true. | [optional] 
**LogResponseHeaders** | Pointer to [**EnumlogPublisherLogResponseHeadersProp**](EnumlogPublisherLogResponseHeadersProp.md) |  | [optional] 
**SuppressedResponseHeaderName** | Pointer to **[]string** | Specifies the case-insensitive names of response headers that should be omitted from log messages (e.g., for the purpose of brevity or security). This will only be used if the log-response-headers property has a value of true. | [optional] 
**LogRequestAuthorizationType** | Pointer to **bool** | Indicates whether to log the type of credentials given if an \&quot;Authorization\&quot; header was included in the request. Logging the authorization type may be useful, and is much more secure than logging the entire value of the \&quot;Authorization\&quot; header. | [optional] 
**LogRequestCookieNames** | Pointer to **bool** | Indicates whether to log the names of any cookies included in an HTTP request. Logging cookie names may be useful and is much more secure than logging the entire content of the cookies (which may include sensitive information). | [optional] 
**LogResponseCookieNames** | Pointer to **bool** | Indicates whether to log the names of any cookies set in an HTTP response. Logging cookie names may be useful and is much more secure than logging the entire content of the cookies (which may include sensitive information). | [optional] 
**LogRequestParameters** | Pointer to [**EnumlogPublisherLogRequestParametersProp**](EnumlogPublisherLogRequestParametersProp.md) |  | [optional] 
**SuppressedRequestParameterName** | Pointer to **[]string** | Specifies the case-insensitive names of request parameters that should be omitted from log messages (e.g., for the purpose of brevity or security). This will only be used if the log-request-parameters property has a value of parameter-names or parameter-names-and-values. | [optional] 
**LogRequestProtocol** | Pointer to **bool** | Indicates whether request log messages should include information about the HTTP version specified in the request. | [optional] 
**LogRedirectURI** | Pointer to **bool** | Indicates whether the redirect URI (i.e., the value of the \&quot;Location\&quot; header from responses) should be included in response log messages. | [optional] 
**WriteMultiLineMessages** | Pointer to **bool** | Indicates whether the JSON objects should use a multi-line representation (with each object field and array value on its own line) that may be easier for administrators to read, but each message will be larger (because of additional spaces and end-of-line markers), and it may be more difficult to consume and parse through some text-oriented tools. | [optional] 
**Description** | Pointer to **string** | A description for this Log Publisher | [optional] 
**LoggingErrorBehavior** | Pointer to [**EnumlogPublisherLoggingErrorBehaviorProp**](EnumlogPublisherLoggingErrorBehaviorProp.md) |  | [optional] 

## Methods

### NewConsoleJsonHttpOperationLogPublisherResponse

`func NewConsoleJsonHttpOperationLogPublisherResponse(id string, schemas []EnumconsoleJsonHttpOperationLogPublisherSchemaUrn, enabled bool, ) *ConsoleJsonHttpOperationLogPublisherResponse`

NewConsoleJsonHttpOperationLogPublisherResponse instantiates a new ConsoleJsonHttpOperationLogPublisherResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleJsonHttpOperationLogPublisherResponseWithDefaults

`func NewConsoleJsonHttpOperationLogPublisherResponseWithDefaults() *ConsoleJsonHttpOperationLogPublisherResponse`

NewConsoleJsonHttpOperationLogPublisherResponseWithDefaults instantiates a new ConsoleJsonHttpOperationLogPublisherResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetSchemas() []EnumconsoleJsonHttpOperationLogPublisherSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetSchemasOk() (*[]EnumconsoleJsonHttpOperationLogPublisherSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetSchemas(v []EnumconsoleJsonHttpOperationLogPublisherSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetOutputLocation

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetOutputLocation() EnumlogPublisherOutputLocationProp`

GetOutputLocation returns the OutputLocation field if non-nil, zero value otherwise.

### GetOutputLocationOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetOutputLocationOk() (*EnumlogPublisherOutputLocationProp, bool)`

GetOutputLocationOk returns a tuple with the OutputLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputLocation

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetOutputLocation(v EnumlogPublisherOutputLocationProp)`

SetOutputLocation sets OutputLocation field to given value.

### HasOutputLocation

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasOutputLocation() bool`

HasOutputLocation returns a boolean if a field has been set.

### GetLogRequests

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogRequests() bool`

GetLogRequests returns the LogRequests field if non-nil, zero value otherwise.

### GetLogRequestsOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogRequestsOk() (*bool, bool)`

GetLogRequestsOk returns a tuple with the LogRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRequests

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetLogRequests(v bool)`

SetLogRequests sets LogRequests field to given value.

### HasLogRequests

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasLogRequests() bool`

HasLogRequests returns a boolean if a field has been set.

### GetLogResults

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogResults() bool`

GetLogResults returns the LogResults field if non-nil, zero value otherwise.

### GetLogResultsOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogResultsOk() (*bool, bool)`

GetLogResultsOk returns a tuple with the LogResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogResults

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetLogResults(v bool)`

SetLogResults sets LogResults field to given value.

### HasLogResults

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasLogResults() bool`

HasLogResults returns a boolean if a field has been set.

### GetIncludeProductName

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetIncludeProductName() bool`

GetIncludeProductName returns the IncludeProductName field if non-nil, zero value otherwise.

### GetIncludeProductNameOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetIncludeProductNameOk() (*bool, bool)`

GetIncludeProductNameOk returns a tuple with the IncludeProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProductName

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetIncludeProductName(v bool)`

SetIncludeProductName sets IncludeProductName field to given value.

### HasIncludeProductName

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasIncludeProductName() bool`

HasIncludeProductName returns a boolean if a field has been set.

### GetIncludeInstanceName

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetIncludeInstanceName() bool`

GetIncludeInstanceName returns the IncludeInstanceName field if non-nil, zero value otherwise.

### GetIncludeInstanceNameOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetIncludeInstanceNameOk() (*bool, bool)`

GetIncludeInstanceNameOk returns a tuple with the IncludeInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInstanceName

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetIncludeInstanceName(v bool)`

SetIncludeInstanceName sets IncludeInstanceName field to given value.

### HasIncludeInstanceName

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasIncludeInstanceName() bool`

HasIncludeInstanceName returns a boolean if a field has been set.

### GetIncludeStartupID

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetIncludeStartupID() bool`

GetIncludeStartupID returns the IncludeStartupID field if non-nil, zero value otherwise.

### GetIncludeStartupIDOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetIncludeStartupIDOk() (*bool, bool)`

GetIncludeStartupIDOk returns a tuple with the IncludeStartupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStartupID

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetIncludeStartupID(v bool)`

SetIncludeStartupID sets IncludeStartupID field to given value.

### HasIncludeStartupID

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasIncludeStartupID() bool`

HasIncludeStartupID returns a boolean if a field has been set.

### GetIncludeThreadID

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetIncludeThreadID() bool`

GetIncludeThreadID returns the IncludeThreadID field if non-nil, zero value otherwise.

### GetIncludeThreadIDOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetIncludeThreadIDOk() (*bool, bool)`

GetIncludeThreadIDOk returns a tuple with the IncludeThreadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeThreadID

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetIncludeThreadID(v bool)`

SetIncludeThreadID sets IncludeThreadID field to given value.

### HasIncludeThreadID

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasIncludeThreadID() bool`

HasIncludeThreadID returns a boolean if a field has been set.

### GetIncludeRequestDetailsInResultMessages

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetIncludeRequestDetailsInResultMessages() bool`

GetIncludeRequestDetailsInResultMessages returns the IncludeRequestDetailsInResultMessages field if non-nil, zero value otherwise.

### GetIncludeRequestDetailsInResultMessagesOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetIncludeRequestDetailsInResultMessagesOk() (*bool, bool)`

GetIncludeRequestDetailsInResultMessagesOk returns a tuple with the IncludeRequestDetailsInResultMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRequestDetailsInResultMessages

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetIncludeRequestDetailsInResultMessages(v bool)`

SetIncludeRequestDetailsInResultMessages sets IncludeRequestDetailsInResultMessages field to given value.

### HasIncludeRequestDetailsInResultMessages

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasIncludeRequestDetailsInResultMessages() bool`

HasIncludeRequestDetailsInResultMessages returns a boolean if a field has been set.

### GetLogRequestHeaders

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogRequestHeaders() EnumlogPublisherLogRequestHeadersProp`

GetLogRequestHeaders returns the LogRequestHeaders field if non-nil, zero value otherwise.

### GetLogRequestHeadersOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogRequestHeadersOk() (*EnumlogPublisherLogRequestHeadersProp, bool)`

GetLogRequestHeadersOk returns a tuple with the LogRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRequestHeaders

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetLogRequestHeaders(v EnumlogPublisherLogRequestHeadersProp)`

SetLogRequestHeaders sets LogRequestHeaders field to given value.

### HasLogRequestHeaders

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasLogRequestHeaders() bool`

HasLogRequestHeaders returns a boolean if a field has been set.

### GetSuppressedRequestHeaderName

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetSuppressedRequestHeaderName() []string`

GetSuppressedRequestHeaderName returns the SuppressedRequestHeaderName field if non-nil, zero value otherwise.

### GetSuppressedRequestHeaderNameOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetSuppressedRequestHeaderNameOk() (*[]string, bool)`

GetSuppressedRequestHeaderNameOk returns a tuple with the SuppressedRequestHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedRequestHeaderName

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetSuppressedRequestHeaderName(v []string)`

SetSuppressedRequestHeaderName sets SuppressedRequestHeaderName field to given value.

### HasSuppressedRequestHeaderName

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasSuppressedRequestHeaderName() bool`

HasSuppressedRequestHeaderName returns a boolean if a field has been set.

### GetLogResponseHeaders

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogResponseHeaders() EnumlogPublisherLogResponseHeadersProp`

GetLogResponseHeaders returns the LogResponseHeaders field if non-nil, zero value otherwise.

### GetLogResponseHeadersOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogResponseHeadersOk() (*EnumlogPublisherLogResponseHeadersProp, bool)`

GetLogResponseHeadersOk returns a tuple with the LogResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogResponseHeaders

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetLogResponseHeaders(v EnumlogPublisherLogResponseHeadersProp)`

SetLogResponseHeaders sets LogResponseHeaders field to given value.

### HasLogResponseHeaders

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasLogResponseHeaders() bool`

HasLogResponseHeaders returns a boolean if a field has been set.

### GetSuppressedResponseHeaderName

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetSuppressedResponseHeaderName() []string`

GetSuppressedResponseHeaderName returns the SuppressedResponseHeaderName field if non-nil, zero value otherwise.

### GetSuppressedResponseHeaderNameOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetSuppressedResponseHeaderNameOk() (*[]string, bool)`

GetSuppressedResponseHeaderNameOk returns a tuple with the SuppressedResponseHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedResponseHeaderName

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetSuppressedResponseHeaderName(v []string)`

SetSuppressedResponseHeaderName sets SuppressedResponseHeaderName field to given value.

### HasSuppressedResponseHeaderName

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasSuppressedResponseHeaderName() bool`

HasSuppressedResponseHeaderName returns a boolean if a field has been set.

### GetLogRequestAuthorizationType

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogRequestAuthorizationType() bool`

GetLogRequestAuthorizationType returns the LogRequestAuthorizationType field if non-nil, zero value otherwise.

### GetLogRequestAuthorizationTypeOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogRequestAuthorizationTypeOk() (*bool, bool)`

GetLogRequestAuthorizationTypeOk returns a tuple with the LogRequestAuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRequestAuthorizationType

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetLogRequestAuthorizationType(v bool)`

SetLogRequestAuthorizationType sets LogRequestAuthorizationType field to given value.

### HasLogRequestAuthorizationType

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasLogRequestAuthorizationType() bool`

HasLogRequestAuthorizationType returns a boolean if a field has been set.

### GetLogRequestCookieNames

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogRequestCookieNames() bool`

GetLogRequestCookieNames returns the LogRequestCookieNames field if non-nil, zero value otherwise.

### GetLogRequestCookieNamesOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogRequestCookieNamesOk() (*bool, bool)`

GetLogRequestCookieNamesOk returns a tuple with the LogRequestCookieNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRequestCookieNames

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetLogRequestCookieNames(v bool)`

SetLogRequestCookieNames sets LogRequestCookieNames field to given value.

### HasLogRequestCookieNames

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasLogRequestCookieNames() bool`

HasLogRequestCookieNames returns a boolean if a field has been set.

### GetLogResponseCookieNames

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogResponseCookieNames() bool`

GetLogResponseCookieNames returns the LogResponseCookieNames field if non-nil, zero value otherwise.

### GetLogResponseCookieNamesOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogResponseCookieNamesOk() (*bool, bool)`

GetLogResponseCookieNamesOk returns a tuple with the LogResponseCookieNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogResponseCookieNames

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetLogResponseCookieNames(v bool)`

SetLogResponseCookieNames sets LogResponseCookieNames field to given value.

### HasLogResponseCookieNames

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasLogResponseCookieNames() bool`

HasLogResponseCookieNames returns a boolean if a field has been set.

### GetLogRequestParameters

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogRequestParameters() EnumlogPublisherLogRequestParametersProp`

GetLogRequestParameters returns the LogRequestParameters field if non-nil, zero value otherwise.

### GetLogRequestParametersOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogRequestParametersOk() (*EnumlogPublisherLogRequestParametersProp, bool)`

GetLogRequestParametersOk returns a tuple with the LogRequestParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRequestParameters

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetLogRequestParameters(v EnumlogPublisherLogRequestParametersProp)`

SetLogRequestParameters sets LogRequestParameters field to given value.

### HasLogRequestParameters

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasLogRequestParameters() bool`

HasLogRequestParameters returns a boolean if a field has been set.

### GetSuppressedRequestParameterName

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetSuppressedRequestParameterName() []string`

GetSuppressedRequestParameterName returns the SuppressedRequestParameterName field if non-nil, zero value otherwise.

### GetSuppressedRequestParameterNameOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetSuppressedRequestParameterNameOk() (*[]string, bool)`

GetSuppressedRequestParameterNameOk returns a tuple with the SuppressedRequestParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedRequestParameterName

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetSuppressedRequestParameterName(v []string)`

SetSuppressedRequestParameterName sets SuppressedRequestParameterName field to given value.

### HasSuppressedRequestParameterName

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasSuppressedRequestParameterName() bool`

HasSuppressedRequestParameterName returns a boolean if a field has been set.

### GetLogRequestProtocol

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogRequestProtocol() bool`

GetLogRequestProtocol returns the LogRequestProtocol field if non-nil, zero value otherwise.

### GetLogRequestProtocolOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogRequestProtocolOk() (*bool, bool)`

GetLogRequestProtocolOk returns a tuple with the LogRequestProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRequestProtocol

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetLogRequestProtocol(v bool)`

SetLogRequestProtocol sets LogRequestProtocol field to given value.

### HasLogRequestProtocol

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasLogRequestProtocol() bool`

HasLogRequestProtocol returns a boolean if a field has been set.

### GetLogRedirectURI

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogRedirectURI() bool`

GetLogRedirectURI returns the LogRedirectURI field if non-nil, zero value otherwise.

### GetLogRedirectURIOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLogRedirectURIOk() (*bool, bool)`

GetLogRedirectURIOk returns a tuple with the LogRedirectURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRedirectURI

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetLogRedirectURI(v bool)`

SetLogRedirectURI sets LogRedirectURI field to given value.

### HasLogRedirectURI

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasLogRedirectURI() bool`

HasLogRedirectURI returns a boolean if a field has been set.

### GetWriteMultiLineMessages

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetWriteMultiLineMessages() bool`

GetWriteMultiLineMessages returns the WriteMultiLineMessages field if non-nil, zero value otherwise.

### GetWriteMultiLineMessagesOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetWriteMultiLineMessagesOk() (*bool, bool)`

GetWriteMultiLineMessagesOk returns a tuple with the WriteMultiLineMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteMultiLineMessages

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetWriteMultiLineMessages(v bool)`

SetWriteMultiLineMessages sets WriteMultiLineMessages field to given value.

### HasWriteMultiLineMessages

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasWriteMultiLineMessages() bool`

HasWriteMultiLineMessages returns a boolean if a field has been set.

### GetDescription

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLoggingErrorBehavior

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *ConsoleJsonHttpOperationLogPublisherResponse) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

