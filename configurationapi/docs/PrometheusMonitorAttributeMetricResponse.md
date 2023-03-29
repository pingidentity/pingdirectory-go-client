# PrometheusMonitorAttributeMetricResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the HTTP Servlet Extension | 
**Schemas** | Pointer to [**[]EnumprometheusMonitorAttributeMetricSchemaUrn**](EnumprometheusMonitorAttributeMetricSchemaUrn.md) |  | [optional] 
**MetricName** | **string** | The name that will be used in the metric to be consumed by Prometheus. | 
**MonitorAttributeName** | **string** | The name of the monitor attribute that contains the numeric value to be published. | 
**MonitorObjectClassName** | **string** | The name of the object class for monitor entries that contain the monitor attribute. | 
**MetricType** | [**EnumprometheusMonitorAttributeMetricMetricTypeProp**](EnumprometheusMonitorAttributeMetricMetricTypeProp.md) |  | 
**Filter** | Pointer to **string** | A filter that may be used to restrict the set of monitor entries for which the metric should be generated. | [optional] 
**MetricDescription** | Pointer to **string** | A human-readable description that should be published as part of the metric definition. | [optional] 
**LabelNameValuePair** | Pointer to **[]string** | A set of name-value pairs for labels that should be included in the published metric for the target attribute. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewPrometheusMonitorAttributeMetricResponse

`func NewPrometheusMonitorAttributeMetricResponse(id string, metricName string, monitorAttributeName string, monitorObjectClassName string, metricType EnumprometheusMonitorAttributeMetricMetricTypeProp, ) *PrometheusMonitorAttributeMetricResponse`

NewPrometheusMonitorAttributeMetricResponse instantiates a new PrometheusMonitorAttributeMetricResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusMonitorAttributeMetricResponseWithDefaults

`func NewPrometheusMonitorAttributeMetricResponseWithDefaults() *PrometheusMonitorAttributeMetricResponse`

NewPrometheusMonitorAttributeMetricResponseWithDefaults instantiates a new PrometheusMonitorAttributeMetricResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrometheusMonitorAttributeMetricResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrometheusMonitorAttributeMetricResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrometheusMonitorAttributeMetricResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *PrometheusMonitorAttributeMetricResponse) GetSchemas() []EnumprometheusMonitorAttributeMetricSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PrometheusMonitorAttributeMetricResponse) GetSchemasOk() (*[]EnumprometheusMonitorAttributeMetricSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PrometheusMonitorAttributeMetricResponse) SetSchemas(v []EnumprometheusMonitorAttributeMetricSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *PrometheusMonitorAttributeMetricResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetMetricName

`func (o *PrometheusMonitorAttributeMetricResponse) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *PrometheusMonitorAttributeMetricResponse) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *PrometheusMonitorAttributeMetricResponse) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetMonitorAttributeName

`func (o *PrometheusMonitorAttributeMetricResponse) GetMonitorAttributeName() string`

GetMonitorAttributeName returns the MonitorAttributeName field if non-nil, zero value otherwise.

### GetMonitorAttributeNameOk

`func (o *PrometheusMonitorAttributeMetricResponse) GetMonitorAttributeNameOk() (*string, bool)`

GetMonitorAttributeNameOk returns a tuple with the MonitorAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorAttributeName

`func (o *PrometheusMonitorAttributeMetricResponse) SetMonitorAttributeName(v string)`

SetMonitorAttributeName sets MonitorAttributeName field to given value.


### GetMonitorObjectClassName

`func (o *PrometheusMonitorAttributeMetricResponse) GetMonitorObjectClassName() string`

GetMonitorObjectClassName returns the MonitorObjectClassName field if non-nil, zero value otherwise.

### GetMonitorObjectClassNameOk

`func (o *PrometheusMonitorAttributeMetricResponse) GetMonitorObjectClassNameOk() (*string, bool)`

GetMonitorObjectClassNameOk returns a tuple with the MonitorObjectClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorObjectClassName

`func (o *PrometheusMonitorAttributeMetricResponse) SetMonitorObjectClassName(v string)`

SetMonitorObjectClassName sets MonitorObjectClassName field to given value.


### GetMetricType

`func (o *PrometheusMonitorAttributeMetricResponse) GetMetricType() EnumprometheusMonitorAttributeMetricMetricTypeProp`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *PrometheusMonitorAttributeMetricResponse) GetMetricTypeOk() (*EnumprometheusMonitorAttributeMetricMetricTypeProp, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *PrometheusMonitorAttributeMetricResponse) SetMetricType(v EnumprometheusMonitorAttributeMetricMetricTypeProp)`

SetMetricType sets MetricType field to given value.


### GetFilter

`func (o *PrometheusMonitorAttributeMetricResponse) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PrometheusMonitorAttributeMetricResponse) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PrometheusMonitorAttributeMetricResponse) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PrometheusMonitorAttributeMetricResponse) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetMetricDescription

`func (o *PrometheusMonitorAttributeMetricResponse) GetMetricDescription() string`

GetMetricDescription returns the MetricDescription field if non-nil, zero value otherwise.

### GetMetricDescriptionOk

`func (o *PrometheusMonitorAttributeMetricResponse) GetMetricDescriptionOk() (*string, bool)`

GetMetricDescriptionOk returns a tuple with the MetricDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescription

`func (o *PrometheusMonitorAttributeMetricResponse) SetMetricDescription(v string)`

SetMetricDescription sets MetricDescription field to given value.

### HasMetricDescription

`func (o *PrometheusMonitorAttributeMetricResponse) HasMetricDescription() bool`

HasMetricDescription returns a boolean if a field has been set.

### GetLabelNameValuePair

`func (o *PrometheusMonitorAttributeMetricResponse) GetLabelNameValuePair() []string`

GetLabelNameValuePair returns the LabelNameValuePair field if non-nil, zero value otherwise.

### GetLabelNameValuePairOk

`func (o *PrometheusMonitorAttributeMetricResponse) GetLabelNameValuePairOk() (*[]string, bool)`

GetLabelNameValuePairOk returns a tuple with the LabelNameValuePair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelNameValuePair

`func (o *PrometheusMonitorAttributeMetricResponse) SetLabelNameValuePair(v []string)`

SetLabelNameValuePair sets LabelNameValuePair field to given value.

### HasLabelNameValuePair

`func (o *PrometheusMonitorAttributeMetricResponse) HasLabelNameValuePair() bool`

HasLabelNameValuePair returns a boolean if a field has been set.

### GetMeta

`func (o *PrometheusMonitorAttributeMetricResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PrometheusMonitorAttributeMetricResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PrometheusMonitorAttributeMetricResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PrometheusMonitorAttributeMetricResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *PrometheusMonitorAttributeMetricResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *PrometheusMonitorAttributeMetricResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *PrometheusMonitorAttributeMetricResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *PrometheusMonitorAttributeMetricResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


