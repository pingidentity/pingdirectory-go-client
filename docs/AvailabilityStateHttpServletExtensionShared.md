# AvailabilityStateHttpServletExtensionShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumavailabilityStateHttpServletExtensionSchemaUrn**](EnumavailabilityStateHttpServletExtensionSchemaUrn.md) |  | 
**BaseContextPath** | **string** | Specifies the base context path that HTTP clients should use to access this servlet. The value must start with a forward slash and must represent a valid HTTP context path. | 
**AvailableStatusCode** | **int32** | Specifies the HTTP status code that the servlet should return if the server considers itself to be available. | 
**DegradedStatusCode** | **int32** | Specifies the HTTP status code that the servlet should return if the server considers itself to be degraded. | 
**UnavailableStatusCode** | **int32** | Specifies the HTTP status code that the servlet should return if the server considers itself to be unavailable. | 
**OverrideStatusCode** | Pointer to **int32** | Specifies a HTTP status code that the servlet should always return, regardless of the server&#39;s availability. If this value is defined, it will override the availability-based return codes. | [optional] 
**IncludeResponseBody** | Pointer to **bool** | Indicates whether the response should include a body that is a JSON object. | [optional] 
**AdditionalResponseContents** | Pointer to **string** | A JSON-formatted string containing additional fields to be returned in the response body. For example, an additional-response-contents value of &#39;{ \&quot;key\&quot;: \&quot;value\&quot; }&#39; would add the key and value to the root of the JSON response body. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** |  | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 

## Methods

### NewAvailabilityStateHttpServletExtensionShared

`func NewAvailabilityStateHttpServletExtensionShared(schemas []EnumavailabilityStateHttpServletExtensionSchemaUrn, baseContextPath string, availableStatusCode int32, degradedStatusCode int32, unavailableStatusCode int32, ) *AvailabilityStateHttpServletExtensionShared`

NewAvailabilityStateHttpServletExtensionShared instantiates a new AvailabilityStateHttpServletExtensionShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityStateHttpServletExtensionSharedWithDefaults

`func NewAvailabilityStateHttpServletExtensionSharedWithDefaults() *AvailabilityStateHttpServletExtensionShared`

NewAvailabilityStateHttpServletExtensionSharedWithDefaults instantiates a new AvailabilityStateHttpServletExtensionShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AvailabilityStateHttpServletExtensionShared) GetSchemas() []EnumavailabilityStateHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AvailabilityStateHttpServletExtensionShared) GetSchemasOk() (*[]EnumavailabilityStateHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AvailabilityStateHttpServletExtensionShared) SetSchemas(v []EnumavailabilityStateHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBaseContextPath

`func (o *AvailabilityStateHttpServletExtensionShared) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *AvailabilityStateHttpServletExtensionShared) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *AvailabilityStateHttpServletExtensionShared) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.


### GetAvailableStatusCode

`func (o *AvailabilityStateHttpServletExtensionShared) GetAvailableStatusCode() int32`

GetAvailableStatusCode returns the AvailableStatusCode field if non-nil, zero value otherwise.

### GetAvailableStatusCodeOk

`func (o *AvailabilityStateHttpServletExtensionShared) GetAvailableStatusCodeOk() (*int32, bool)`

GetAvailableStatusCodeOk returns a tuple with the AvailableStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStatusCode

`func (o *AvailabilityStateHttpServletExtensionShared) SetAvailableStatusCode(v int32)`

SetAvailableStatusCode sets AvailableStatusCode field to given value.


### GetDegradedStatusCode

`func (o *AvailabilityStateHttpServletExtensionShared) GetDegradedStatusCode() int32`

GetDegradedStatusCode returns the DegradedStatusCode field if non-nil, zero value otherwise.

### GetDegradedStatusCodeOk

`func (o *AvailabilityStateHttpServletExtensionShared) GetDegradedStatusCodeOk() (*int32, bool)`

GetDegradedStatusCodeOk returns a tuple with the DegradedStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegradedStatusCode

`func (o *AvailabilityStateHttpServletExtensionShared) SetDegradedStatusCode(v int32)`

SetDegradedStatusCode sets DegradedStatusCode field to given value.


### GetUnavailableStatusCode

`func (o *AvailabilityStateHttpServletExtensionShared) GetUnavailableStatusCode() int32`

GetUnavailableStatusCode returns the UnavailableStatusCode field if non-nil, zero value otherwise.

### GetUnavailableStatusCodeOk

`func (o *AvailabilityStateHttpServletExtensionShared) GetUnavailableStatusCodeOk() (*int32, bool)`

GetUnavailableStatusCodeOk returns a tuple with the UnavailableStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableStatusCode

`func (o *AvailabilityStateHttpServletExtensionShared) SetUnavailableStatusCode(v int32)`

SetUnavailableStatusCode sets UnavailableStatusCode field to given value.


### GetOverrideStatusCode

`func (o *AvailabilityStateHttpServletExtensionShared) GetOverrideStatusCode() int32`

GetOverrideStatusCode returns the OverrideStatusCode field if non-nil, zero value otherwise.

### GetOverrideStatusCodeOk

`func (o *AvailabilityStateHttpServletExtensionShared) GetOverrideStatusCodeOk() (*int32, bool)`

GetOverrideStatusCodeOk returns a tuple with the OverrideStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideStatusCode

`func (o *AvailabilityStateHttpServletExtensionShared) SetOverrideStatusCode(v int32)`

SetOverrideStatusCode sets OverrideStatusCode field to given value.

### HasOverrideStatusCode

`func (o *AvailabilityStateHttpServletExtensionShared) HasOverrideStatusCode() bool`

HasOverrideStatusCode returns a boolean if a field has been set.

### GetIncludeResponseBody

`func (o *AvailabilityStateHttpServletExtensionShared) GetIncludeResponseBody() bool`

GetIncludeResponseBody returns the IncludeResponseBody field if non-nil, zero value otherwise.

### GetIncludeResponseBodyOk

`func (o *AvailabilityStateHttpServletExtensionShared) GetIncludeResponseBodyOk() (*bool, bool)`

GetIncludeResponseBodyOk returns a tuple with the IncludeResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResponseBody

`func (o *AvailabilityStateHttpServletExtensionShared) SetIncludeResponseBody(v bool)`

SetIncludeResponseBody sets IncludeResponseBody field to given value.

### HasIncludeResponseBody

`func (o *AvailabilityStateHttpServletExtensionShared) HasIncludeResponseBody() bool`

HasIncludeResponseBody returns a boolean if a field has been set.

### GetAdditionalResponseContents

`func (o *AvailabilityStateHttpServletExtensionShared) GetAdditionalResponseContents() string`

GetAdditionalResponseContents returns the AdditionalResponseContents field if non-nil, zero value otherwise.

### GetAdditionalResponseContentsOk

`func (o *AvailabilityStateHttpServletExtensionShared) GetAdditionalResponseContentsOk() (*string, bool)`

GetAdditionalResponseContentsOk returns a tuple with the AdditionalResponseContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalResponseContents

`func (o *AvailabilityStateHttpServletExtensionShared) SetAdditionalResponseContents(v string)`

SetAdditionalResponseContents sets AdditionalResponseContents field to given value.

### HasAdditionalResponseContents

`func (o *AvailabilityStateHttpServletExtensionShared) HasAdditionalResponseContents() bool`

HasAdditionalResponseContents returns a boolean if a field has been set.

### GetDescription

`func (o *AvailabilityStateHttpServletExtensionShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AvailabilityStateHttpServletExtensionShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AvailabilityStateHttpServletExtensionShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AvailabilityStateHttpServletExtensionShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *AvailabilityStateHttpServletExtensionShared) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *AvailabilityStateHttpServletExtensionShared) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *AvailabilityStateHttpServletExtensionShared) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *AvailabilityStateHttpServletExtensionShared) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *AvailabilityStateHttpServletExtensionShared) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *AvailabilityStateHttpServletExtensionShared) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *AvailabilityStateHttpServletExtensionShared) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *AvailabilityStateHttpServletExtensionShared) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *AvailabilityStateHttpServletExtensionShared) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *AvailabilityStateHttpServletExtensionShared) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *AvailabilityStateHttpServletExtensionShared) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *AvailabilityStateHttpServletExtensionShared) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


