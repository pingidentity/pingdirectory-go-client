# AvailabilityStateHttpServletExtensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the HTTP Servlet Extension | 
**Schemas** | [**[]EnumavailabilityStateHttpServletExtensionSchemaUrn**](EnumavailabilityStateHttpServletExtensionSchemaUrn.md) |  | 
**BaseContextPath** | **string** | Specifies the base context path that HTTP clients should use to access this servlet. The value must start with a forward slash and must represent a valid HTTP context path. | 
**AvailableStatusCode** | **int64** | Specifies the HTTP status code that the servlet should return if the server considers itself to be available. | 
**DegradedStatusCode** | **int64** | Specifies the HTTP status code that the servlet should return if the server considers itself to be degraded. | 
**UnavailableStatusCode** | **int64** | Specifies the HTTP status code that the servlet should return if the server considers itself to be unavailable. | 
**OverrideStatusCode** | Pointer to **int64** | Specifies a HTTP status code that the servlet should always return, regardless of the server&#39;s availability. If this value is defined, it will override the availability-based return codes. | [optional] 
**IncludeResponseBody** | Pointer to **bool** | Indicates whether the response should include a body that is a JSON object. | [optional] 
**AdditionalResponseContents** | Pointer to **string** | A JSON-formatted string containing additional fields to be returned in the response body. For example, an additional-response-contents value of &#39;{ \&quot;key\&quot;: \&quot;value\&quot; }&#39; would add the key and value to the root of the JSON response body. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewAvailabilityStateHttpServletExtensionResponse

`func NewAvailabilityStateHttpServletExtensionResponse(id string, schemas []EnumavailabilityStateHttpServletExtensionSchemaUrn, baseContextPath string, availableStatusCode int64, degradedStatusCode int64, unavailableStatusCode int64, ) *AvailabilityStateHttpServletExtensionResponse`

NewAvailabilityStateHttpServletExtensionResponse instantiates a new AvailabilityStateHttpServletExtensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityStateHttpServletExtensionResponseWithDefaults

`func NewAvailabilityStateHttpServletExtensionResponseWithDefaults() *AvailabilityStateHttpServletExtensionResponse`

NewAvailabilityStateHttpServletExtensionResponseWithDefaults instantiates a new AvailabilityStateHttpServletExtensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AvailabilityStateHttpServletExtensionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AvailabilityStateHttpServletExtensionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AvailabilityStateHttpServletExtensionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AvailabilityStateHttpServletExtensionResponse) GetSchemas() []EnumavailabilityStateHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AvailabilityStateHttpServletExtensionResponse) GetSchemasOk() (*[]EnumavailabilityStateHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AvailabilityStateHttpServletExtensionResponse) SetSchemas(v []EnumavailabilityStateHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBaseContextPath

`func (o *AvailabilityStateHttpServletExtensionResponse) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *AvailabilityStateHttpServletExtensionResponse) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *AvailabilityStateHttpServletExtensionResponse) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.


### GetAvailableStatusCode

`func (o *AvailabilityStateHttpServletExtensionResponse) GetAvailableStatusCode() int64`

GetAvailableStatusCode returns the AvailableStatusCode field if non-nil, zero value otherwise.

### GetAvailableStatusCodeOk

`func (o *AvailabilityStateHttpServletExtensionResponse) GetAvailableStatusCodeOk() (*int64, bool)`

GetAvailableStatusCodeOk returns a tuple with the AvailableStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStatusCode

`func (o *AvailabilityStateHttpServletExtensionResponse) SetAvailableStatusCode(v int64)`

SetAvailableStatusCode sets AvailableStatusCode field to given value.


### GetDegradedStatusCode

`func (o *AvailabilityStateHttpServletExtensionResponse) GetDegradedStatusCode() int64`

GetDegradedStatusCode returns the DegradedStatusCode field if non-nil, zero value otherwise.

### GetDegradedStatusCodeOk

`func (o *AvailabilityStateHttpServletExtensionResponse) GetDegradedStatusCodeOk() (*int64, bool)`

GetDegradedStatusCodeOk returns a tuple with the DegradedStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegradedStatusCode

`func (o *AvailabilityStateHttpServletExtensionResponse) SetDegradedStatusCode(v int64)`

SetDegradedStatusCode sets DegradedStatusCode field to given value.


### GetUnavailableStatusCode

`func (o *AvailabilityStateHttpServletExtensionResponse) GetUnavailableStatusCode() int64`

GetUnavailableStatusCode returns the UnavailableStatusCode field if non-nil, zero value otherwise.

### GetUnavailableStatusCodeOk

`func (o *AvailabilityStateHttpServletExtensionResponse) GetUnavailableStatusCodeOk() (*int64, bool)`

GetUnavailableStatusCodeOk returns a tuple with the UnavailableStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableStatusCode

`func (o *AvailabilityStateHttpServletExtensionResponse) SetUnavailableStatusCode(v int64)`

SetUnavailableStatusCode sets UnavailableStatusCode field to given value.


### GetOverrideStatusCode

`func (o *AvailabilityStateHttpServletExtensionResponse) GetOverrideStatusCode() int64`

GetOverrideStatusCode returns the OverrideStatusCode field if non-nil, zero value otherwise.

### GetOverrideStatusCodeOk

`func (o *AvailabilityStateHttpServletExtensionResponse) GetOverrideStatusCodeOk() (*int64, bool)`

GetOverrideStatusCodeOk returns a tuple with the OverrideStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideStatusCode

`func (o *AvailabilityStateHttpServletExtensionResponse) SetOverrideStatusCode(v int64)`

SetOverrideStatusCode sets OverrideStatusCode field to given value.

### HasOverrideStatusCode

`func (o *AvailabilityStateHttpServletExtensionResponse) HasOverrideStatusCode() bool`

HasOverrideStatusCode returns a boolean if a field has been set.

### GetIncludeResponseBody

`func (o *AvailabilityStateHttpServletExtensionResponse) GetIncludeResponseBody() bool`

GetIncludeResponseBody returns the IncludeResponseBody field if non-nil, zero value otherwise.

### GetIncludeResponseBodyOk

`func (o *AvailabilityStateHttpServletExtensionResponse) GetIncludeResponseBodyOk() (*bool, bool)`

GetIncludeResponseBodyOk returns a tuple with the IncludeResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResponseBody

`func (o *AvailabilityStateHttpServletExtensionResponse) SetIncludeResponseBody(v bool)`

SetIncludeResponseBody sets IncludeResponseBody field to given value.

### HasIncludeResponseBody

`func (o *AvailabilityStateHttpServletExtensionResponse) HasIncludeResponseBody() bool`

HasIncludeResponseBody returns a boolean if a field has been set.

### GetAdditionalResponseContents

`func (o *AvailabilityStateHttpServletExtensionResponse) GetAdditionalResponseContents() string`

GetAdditionalResponseContents returns the AdditionalResponseContents field if non-nil, zero value otherwise.

### GetAdditionalResponseContentsOk

`func (o *AvailabilityStateHttpServletExtensionResponse) GetAdditionalResponseContentsOk() (*string, bool)`

GetAdditionalResponseContentsOk returns a tuple with the AdditionalResponseContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalResponseContents

`func (o *AvailabilityStateHttpServletExtensionResponse) SetAdditionalResponseContents(v string)`

SetAdditionalResponseContents sets AdditionalResponseContents field to given value.

### HasAdditionalResponseContents

`func (o *AvailabilityStateHttpServletExtensionResponse) HasAdditionalResponseContents() bool`

HasAdditionalResponseContents returns a boolean if a field has been set.

### GetDescription

`func (o *AvailabilityStateHttpServletExtensionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AvailabilityStateHttpServletExtensionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AvailabilityStateHttpServletExtensionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AvailabilityStateHttpServletExtensionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *AvailabilityStateHttpServletExtensionResponse) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *AvailabilityStateHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *AvailabilityStateHttpServletExtensionResponse) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *AvailabilityStateHttpServletExtensionResponse) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *AvailabilityStateHttpServletExtensionResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *AvailabilityStateHttpServletExtensionResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *AvailabilityStateHttpServletExtensionResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *AvailabilityStateHttpServletExtensionResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *AvailabilityStateHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *AvailabilityStateHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *AvailabilityStateHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *AvailabilityStateHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.

### GetMeta

`func (o *AvailabilityStateHttpServletExtensionResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AvailabilityStateHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AvailabilityStateHttpServletExtensionResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AvailabilityStateHttpServletExtensionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AvailabilityStateHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AvailabilityStateHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AvailabilityStateHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AvailabilityStateHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


