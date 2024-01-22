# PrometheusMonitoringHttpServletExtensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumprometheusMonitoringHttpServletExtensionSchemaUrn**](EnumprometheusMonitoringHttpServletExtensionSchemaUrn.md) |  | 
**BaseContextPath** | **string** | Specifies the base context path that HTTP clients should use to access this servlet. The value must start with a forward slash and must represent a valid HTTP context path. | 
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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the HTTP Servlet Extension | 

## Methods

### NewPrometheusMonitoringHttpServletExtensionResponse

`func NewPrometheusMonitoringHttpServletExtensionResponse(schemas []EnumprometheusMonitoringHttpServletExtensionSchemaUrn, baseContextPath string, id string, ) *PrometheusMonitoringHttpServletExtensionResponse`

NewPrometheusMonitoringHttpServletExtensionResponse instantiates a new PrometheusMonitoringHttpServletExtensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusMonitoringHttpServletExtensionResponseWithDefaults

`func NewPrometheusMonitoringHttpServletExtensionResponseWithDefaults() *PrometheusMonitoringHttpServletExtensionResponse`

NewPrometheusMonitoringHttpServletExtensionResponseWithDefaults instantiates a new PrometheusMonitoringHttpServletExtensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetSchemas() []EnumprometheusMonitoringHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetSchemasOk() (*[]EnumprometheusMonitoringHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PrometheusMonitoringHttpServletExtensionResponse) SetSchemas(v []EnumprometheusMonitoringHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBaseContextPath

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetBaseContextPath() string`

GetBaseContextPath returns the BaseContextPath field if non-nil, zero value otherwise.

### GetBaseContextPathOk

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetBaseContextPathOk() (*string, bool)`

GetBaseContextPathOk returns a tuple with the BaseContextPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContextPath

`func (o *PrometheusMonitoringHttpServletExtensionResponse) SetBaseContextPath(v string)`

SetBaseContextPath sets BaseContextPath field to given value.


### GetIncludeInstanceNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetIncludeInstanceNameLabel() bool`

GetIncludeInstanceNameLabel returns the IncludeInstanceNameLabel field if non-nil, zero value otherwise.

### GetIncludeInstanceNameLabelOk

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetIncludeInstanceNameLabelOk() (*bool, bool)`

GetIncludeInstanceNameLabelOk returns a tuple with the IncludeInstanceNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInstanceNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) SetIncludeInstanceNameLabel(v bool)`

SetIncludeInstanceNameLabel sets IncludeInstanceNameLabel field to given value.

### HasIncludeInstanceNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) HasIncludeInstanceNameLabel() bool`

HasIncludeInstanceNameLabel returns a boolean if a field has been set.

### GetIncludeProductNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetIncludeProductNameLabel() bool`

GetIncludeProductNameLabel returns the IncludeProductNameLabel field if non-nil, zero value otherwise.

### GetIncludeProductNameLabelOk

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetIncludeProductNameLabelOk() (*bool, bool)`

GetIncludeProductNameLabelOk returns a tuple with the IncludeProductNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProductNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) SetIncludeProductNameLabel(v bool)`

SetIncludeProductNameLabel sets IncludeProductNameLabel field to given value.

### HasIncludeProductNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) HasIncludeProductNameLabel() bool`

HasIncludeProductNameLabel returns a boolean if a field has been set.

### GetIncludeLocationNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetIncludeLocationNameLabel() bool`

GetIncludeLocationNameLabel returns the IncludeLocationNameLabel field if non-nil, zero value otherwise.

### GetIncludeLocationNameLabelOk

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetIncludeLocationNameLabelOk() (*bool, bool)`

GetIncludeLocationNameLabelOk returns a tuple with the IncludeLocationNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLocationNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) SetIncludeLocationNameLabel(v bool)`

SetIncludeLocationNameLabel sets IncludeLocationNameLabel field to given value.

### HasIncludeLocationNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) HasIncludeLocationNameLabel() bool`

HasIncludeLocationNameLabel returns a boolean if a field has been set.

### GetAlwaysIncludeMonitorEntryNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetAlwaysIncludeMonitorEntryNameLabel() bool`

GetAlwaysIncludeMonitorEntryNameLabel returns the AlwaysIncludeMonitorEntryNameLabel field if non-nil, zero value otherwise.

### GetAlwaysIncludeMonitorEntryNameLabelOk

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetAlwaysIncludeMonitorEntryNameLabelOk() (*bool, bool)`

GetAlwaysIncludeMonitorEntryNameLabelOk returns a tuple with the AlwaysIncludeMonitorEntryNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysIncludeMonitorEntryNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) SetAlwaysIncludeMonitorEntryNameLabel(v bool)`

SetAlwaysIncludeMonitorEntryNameLabel sets AlwaysIncludeMonitorEntryNameLabel field to given value.

### HasAlwaysIncludeMonitorEntryNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) HasAlwaysIncludeMonitorEntryNameLabel() bool`

HasAlwaysIncludeMonitorEntryNameLabel returns a boolean if a field has been set.

### GetIncludeMonitorObjectClassNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetIncludeMonitorObjectClassNameLabel() bool`

GetIncludeMonitorObjectClassNameLabel returns the IncludeMonitorObjectClassNameLabel field if non-nil, zero value otherwise.

### GetIncludeMonitorObjectClassNameLabelOk

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetIncludeMonitorObjectClassNameLabelOk() (*bool, bool)`

GetIncludeMonitorObjectClassNameLabelOk returns a tuple with the IncludeMonitorObjectClassNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMonitorObjectClassNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) SetIncludeMonitorObjectClassNameLabel(v bool)`

SetIncludeMonitorObjectClassNameLabel sets IncludeMonitorObjectClassNameLabel field to given value.

### HasIncludeMonitorObjectClassNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) HasIncludeMonitorObjectClassNameLabel() bool`

HasIncludeMonitorObjectClassNameLabel returns a boolean if a field has been set.

### GetIncludeMonitorAttributeNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetIncludeMonitorAttributeNameLabel() bool`

GetIncludeMonitorAttributeNameLabel returns the IncludeMonitorAttributeNameLabel field if non-nil, zero value otherwise.

### GetIncludeMonitorAttributeNameLabelOk

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetIncludeMonitorAttributeNameLabelOk() (*bool, bool)`

GetIncludeMonitorAttributeNameLabelOk returns a tuple with the IncludeMonitorAttributeNameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMonitorAttributeNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) SetIncludeMonitorAttributeNameLabel(v bool)`

SetIncludeMonitorAttributeNameLabel sets IncludeMonitorAttributeNameLabel field to given value.

### HasIncludeMonitorAttributeNameLabel

`func (o *PrometheusMonitoringHttpServletExtensionResponse) HasIncludeMonitorAttributeNameLabel() bool`

HasIncludeMonitorAttributeNameLabel returns a boolean if a field has been set.

### GetLabelNameValuePair

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetLabelNameValuePair() []string`

GetLabelNameValuePair returns the LabelNameValuePair field if non-nil, zero value otherwise.

### GetLabelNameValuePairOk

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetLabelNameValuePairOk() (*[]string, bool)`

GetLabelNameValuePairOk returns a tuple with the LabelNameValuePair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelNameValuePair

`func (o *PrometheusMonitoringHttpServletExtensionResponse) SetLabelNameValuePair(v []string)`

SetLabelNameValuePair sets LabelNameValuePair field to given value.

### HasLabelNameValuePair

`func (o *PrometheusMonitoringHttpServletExtensionResponse) HasLabelNameValuePair() bool`

HasLabelNameValuePair returns a boolean if a field has been set.

### GetDescription

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrometheusMonitoringHttpServletExtensionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrometheusMonitoringHttpServletExtensionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *PrometheusMonitoringHttpServletExtensionResponse) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *PrometheusMonitoringHttpServletExtensionResponse) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *PrometheusMonitoringHttpServletExtensionResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *PrometheusMonitoringHttpServletExtensionResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *PrometheusMonitoringHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *PrometheusMonitoringHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.

### GetMeta

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PrometheusMonitoringHttpServletExtensionResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PrometheusMonitoringHttpServletExtensionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *PrometheusMonitoringHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *PrometheusMonitoringHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrometheusMonitoringHttpServletExtensionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrometheusMonitoringHttpServletExtensionResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


