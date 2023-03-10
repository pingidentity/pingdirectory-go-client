# AddPrometheusMonitoringHttpServletExtensionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionName** | **string** | Name of the new HTTP Servlet Extension | 
**Schemas** | [**[]EnumprometheusMonitoringHttpServletExtensionSchemaUrn**](EnumprometheusMonitoringHttpServletExtensionSchemaUrn.md) |  | 
**BaseContextPath** | Pointer to **string** | Specifies the base context path that HTTP clients should use to access this servlet. The value must start with a forward slash and must represent a valid HTTP context path. | [optional] 
**IncludeInstanceNameLabel** | Pointer to **bool** | Indicates whether generated metrics should include an \&quot;instance\&quot; label whose value is the instance name for this Directory Server instance. | [optional] 
**IncludeProductNameLabel** | Pointer to **bool** | Indicates whether generated metrics should include a \&quot;product\&quot; label whose value is the product name for this Directory Server instance. | [optional] 
**IncludeLocationNameLabel** | Pointer to **bool** | Indicates whether generated metrics should include a \&quot;location\&quot; label whose value is the location name for this Directory Server instance. | [optional] 
**AlwaysIncludeMonitorEntryNameLabel** | Pointer to **bool** | Indicates whether generated metrics should always include a \&quot;monitor_entry\&quot; label whose value is the name of the monitor entry from which the metric was obtained. | [optional] 
**IncludeMonitorObjectClassNameLabel** | Pointer to **bool** | Indicates whether generated metrics should include a \&quot;monitor_object_class\&quot; label whose value is the name of the object class for the monitor entry from which the metric was obtained. | [optional] 
**IncludeMonitorAttributeNameLabel** | Pointer to **bool** | Indicates whether generated metrics should include a \&quot;monitor_attribute\&quot; label whose value is the name of the monitor attribute from which the metric was obtained. | [optional] 
**LabelNameValuePair** | Pointer to **[]string** | A set of name-value pairs for labels that should be included in all metrics exposed by this Directory Server instance. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 

## Methods

### NewAddPrometheusMonitoringHttpServletExtensionRequest

`func NewAddPrometheusMonitoringHttpServletExtensionRequest(extensionName string, schemas []EnumprometheusMonitoringHttpServletExtensionSchemaUrn, ) *AddPrometheusMonitoringHttpServletExtensionRequest`

NewAddPrometheusMonitoringHttpServletExtensionRequest instantiates a new AddPrometheusMonitoringHttpServletExtensionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPrometheusMonitoringHttpServletExtensionRequestWithDefaults

`func NewAddPrometheusMonitoringHttpServletExtensionRequestWithDefaults() *AddPrometheusMonitoringHttpServletExtensionRequest`

NewAddPrometheusMonitoringHttpServletExtensionRequestWithDefaults instantiates a new AddPrometheusMonitoringHttpServletExtensionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionName

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.


### GetSchemas

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetSchemas() []EnumprometheusMonitoringHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetSchemasOk() (*[]EnumprometheusMonitoringHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) SetSchemas(v []EnumprometheusMonitoringHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBaseContextPath

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.

### HasBaseContextPath

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) HasBaseContextPath() bool`

HasBaseContextPath returns a boolean if a field has been set.

### GetIncludeInstanceNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetIncludeInstanceNameLabel() bool`

GetIncludeInstanceNameLabel returns the IncludeInstanceNameLabel field if non-nil, zero value otherwise.

### GetIncludeInstanceNameLabelOk

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetIncludeInstanceNameLabelOk() (*bool, bool)`

GetIncludeInstanceNameLabelOk returns a tuple with the IncludeInstanceNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInstanceNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) SetIncludeInstanceNameLabel(v bool)`

SetIncludeInstanceNameLabel sets IncludeInstanceNameLabel field to given value.

### HasIncludeInstanceNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) HasIncludeInstanceNameLabel() bool`

HasIncludeInstanceNameLabel returns a boolean if a field has been set.

### GetIncludeProductNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetIncludeProductNameLabel() bool`

GetIncludeProductNameLabel returns the IncludeProductNameLabel field if non-nil, zero value otherwise.

### GetIncludeProductNameLabelOk

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetIncludeProductNameLabelOk() (*bool, bool)`

GetIncludeProductNameLabelOk returns a tuple with the IncludeProductNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProductNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) SetIncludeProductNameLabel(v bool)`

SetIncludeProductNameLabel sets IncludeProductNameLabel field to given value.

### HasIncludeProductNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) HasIncludeProductNameLabel() bool`

HasIncludeProductNameLabel returns a boolean if a field has been set.

### GetIncludeLocationNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetIncludeLocationNameLabel() bool`

GetIncludeLocationNameLabel returns the IncludeLocationNameLabel field if non-nil, zero value otherwise.

### GetIncludeLocationNameLabelOk

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetIncludeLocationNameLabelOk() (*bool, bool)`

GetIncludeLocationNameLabelOk returns a tuple with the IncludeLocationNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLocationNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) SetIncludeLocationNameLabel(v bool)`

SetIncludeLocationNameLabel sets IncludeLocationNameLabel field to given value.

### HasIncludeLocationNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) HasIncludeLocationNameLabel() bool`

HasIncludeLocationNameLabel returns a boolean if a field has been set.

### GetAlwaysIncludeMonitorEntryNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetAlwaysIncludeMonitorEntryNameLabel() bool`

GetAlwaysIncludeMonitorEntryNameLabel returns the AlwaysIncludeMonitorEntryNameLabel field if non-nil, zero value otherwise.

### GetAlwaysIncludeMonitorEntryNameLabelOk

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetAlwaysIncludeMonitorEntryNameLabelOk() (*bool, bool)`

GetAlwaysIncludeMonitorEntryNameLabelOk returns a tuple with the AlwaysIncludeMonitorEntryNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysIncludeMonitorEntryNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) SetAlwaysIncludeMonitorEntryNameLabel(v bool)`

SetAlwaysIncludeMonitorEntryNameLabel sets AlwaysIncludeMonitorEntryNameLabel field to given value.

### HasAlwaysIncludeMonitorEntryNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) HasAlwaysIncludeMonitorEntryNameLabel() bool`

HasAlwaysIncludeMonitorEntryNameLabel returns a boolean if a field has been set.

### GetIncludeMonitorObjectClassNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetIncludeMonitorObjectClassNameLabel() bool`

GetIncludeMonitorObjectClassNameLabel returns the IncludeMonitorObjectClassNameLabel field if non-nil, zero value otherwise.

### GetIncludeMonitorObjectClassNameLabelOk

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetIncludeMonitorObjectClassNameLabelOk() (*bool, bool)`

GetIncludeMonitorObjectClassNameLabelOk returns a tuple with the IncludeMonitorObjectClassNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMonitorObjectClassNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) SetIncludeMonitorObjectClassNameLabel(v bool)`

SetIncludeMonitorObjectClassNameLabel sets IncludeMonitorObjectClassNameLabel field to given value.

### HasIncludeMonitorObjectClassNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) HasIncludeMonitorObjectClassNameLabel() bool`

HasIncludeMonitorObjectClassNameLabel returns a boolean if a field has been set.

### GetIncludeMonitorAttributeNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetIncludeMonitorAttributeNameLabel() bool`

GetIncludeMonitorAttributeNameLabel returns the IncludeMonitorAttributeNameLabel field if non-nil, zero value otherwise.

### GetIncludeMonitorAttributeNameLabelOk

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetIncludeMonitorAttributeNameLabelOk() (*bool, bool)`

GetIncludeMonitorAttributeNameLabelOk returns a tuple with the IncludeMonitorAttributeNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMonitorAttributeNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) SetIncludeMonitorAttributeNameLabel(v bool)`

SetIncludeMonitorAttributeNameLabel sets IncludeMonitorAttributeNameLabel field to given value.

### HasIncludeMonitorAttributeNameLabel

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) HasIncludeMonitorAttributeNameLabel() bool`

HasIncludeMonitorAttributeNameLabel returns a boolean if a field has been set.

### GetLabelNameValuePair

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetLabelNameValuePair() []string`

GetLabelNameValuePair returns the LabelNameValuePair field if non-nil, zero value otherwise.

### GetLabelNameValuePairOk

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetLabelNameValuePairOk() (*[]string, bool)`

GetLabelNameValuePairOk returns a tuple with the LabelNameValuePair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelNameValuePair

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) SetLabelNameValuePair(v []string)`

SetLabelNameValuePair sets LabelNameValuePair field to given value.

### HasLabelNameValuePair

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) HasLabelNameValuePair() bool`

HasLabelNameValuePair returns a boolean if a field has been set.

### GetDescription

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *AddPrometheusMonitoringHttpServletExtensionRequest) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


