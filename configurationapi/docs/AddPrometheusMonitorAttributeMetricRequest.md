# AddPrometheusMonitorAttributeMetricRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumprometheusMonitorAttributeMetricSchemaUrn**](EnumprometheusMonitorAttributeMetricSchemaUrn.md) |  | [optional] 
**MetricName** | **string** | Name of the new Prometheus Monitor Attribute Metric | 
**MonitorAttributeName** | **string** | The name of the monitor attribute that contains the numeric value to be published. | 
**MonitorObjectClassName** | **string** | The name of the object class for monitor entries that contain the monitor attribute. | 
**MetricType** | [**EnumprometheusMonitorAttributeMetricMetricTypeProp**](EnumprometheusMonitorAttributeMetricMetricTypeProp.md) |  | 
**Filter** | Pointer to **string** | A filter that may be used to restrict the set of monitor entries for which the metric should be generated. | [optional] 
**MetricDescription** | Pointer to **string** | A human-readable description that should be published as part of the metric definition. | [optional] 
**LabelNameValuePair** | Pointer to **[]string** | A set of name-value pairs for labels that should be included in the published metric for the target attribute. | [optional] 

## Methods

### NewAddPrometheusMonitorAttributeMetricRequest

`func NewAddPrometheusMonitorAttributeMetricRequest(metricName string, monitorAttributeName string, monitorObjectClassName string, metricType EnumprometheusMonitorAttributeMetricMetricTypeProp, ) *AddPrometheusMonitorAttributeMetricRequest`

NewAddPrometheusMonitorAttributeMetricRequest instantiates a new AddPrometheusMonitorAttributeMetricRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPrometheusMonitorAttributeMetricRequestWithDefaults

`func NewAddPrometheusMonitorAttributeMetricRequestWithDefaults() *AddPrometheusMonitorAttributeMetricRequest`

NewAddPrometheusMonitorAttributeMetricRequestWithDefaults instantiates a new AddPrometheusMonitorAttributeMetricRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddPrometheusMonitorAttributeMetricRequest) GetSchemas() []EnumprometheusMonitorAttributeMetricSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPrometheusMonitorAttributeMetricRequest) GetSchemasOk() (*[]EnumprometheusMonitorAttributeMetricSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPrometheusMonitorAttributeMetricRequest) SetSchemas(v []EnumprometheusMonitorAttributeMetricSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddPrometheusMonitorAttributeMetricRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetMetricName

`func (o *AddPrometheusMonitorAttributeMetricRequest) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *AddPrometheusMonitorAttributeMetricRequest) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *AddPrometheusMonitorAttributeMetricRequest) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetMonitorAttributeName

`func (o *AddPrometheusMonitorAttributeMetricRequest) GetMonitorAttributeName() string`

GetMonitorAttributeName returns the MonitorAttributeName field if non-nil, zero value otherwise.

### GetMonitorAttributeNameOk

`func (o *AddPrometheusMonitorAttributeMetricRequest) GetMonitorAttributeNameOk() (*string, bool)`

GetMonitorAttributeNameOk returns a tuple with the MonitorAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorAttributeName

`func (o *AddPrometheusMonitorAttributeMetricRequest) SetMonitorAttributeName(v string)`

SetMonitorAttributeName sets MonitorAttributeName field to given value.


### GetMonitorObjectClassName

`func (o *AddPrometheusMonitorAttributeMetricRequest) GetMonitorObjectClassName() string`

GetMonitorObjectClassName returns the MonitorObjectClassName field if non-nil, zero value otherwise.

### GetMonitorObjectClassNameOk

`func (o *AddPrometheusMonitorAttributeMetricRequest) GetMonitorObjectClassNameOk() (*string, bool)`

GetMonitorObjectClassNameOk returns a tuple with the MonitorObjectClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorObjectClassName

`func (o *AddPrometheusMonitorAttributeMetricRequest) SetMonitorObjectClassName(v string)`

SetMonitorObjectClassName sets MonitorObjectClassName field to given value.


### GetMetricType

`func (o *AddPrometheusMonitorAttributeMetricRequest) GetMetricType() EnumprometheusMonitorAttributeMetricMetricTypeProp`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *AddPrometheusMonitorAttributeMetricRequest) GetMetricTypeOk() (*EnumprometheusMonitorAttributeMetricMetricTypeProp, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *AddPrometheusMonitorAttributeMetricRequest) SetMetricType(v EnumprometheusMonitorAttributeMetricMetricTypeProp)`

SetMetricType sets MetricType field to given value.


### GetFilter

`func (o *AddPrometheusMonitorAttributeMetricRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddPrometheusMonitorAttributeMetricRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddPrometheusMonitorAttributeMetricRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AddPrometheusMonitorAttributeMetricRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetMetricDescription

`func (o *AddPrometheusMonitorAttributeMetricRequest) GetMetricDescription() string`

GetMetricDescription returns the MetricDescription field if non-nil, zero value otherwise.

### GetMetricDescriptionOk

`func (o *AddPrometheusMonitorAttributeMetricRequest) GetMetricDescriptionOk() (*string, bool)`

GetMetricDescriptionOk returns a tuple with the MetricDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescription

`func (o *AddPrometheusMonitorAttributeMetricRequest) SetMetricDescription(v string)`

SetMetricDescription sets MetricDescription field to given value.

### HasMetricDescription

`func (o *AddPrometheusMonitorAttributeMetricRequest) HasMetricDescription() bool`

HasMetricDescription returns a boolean if a field has been set.

### GetLabelNameValuePair

`func (o *AddPrometheusMonitorAttributeMetricRequest) GetLabelNameValuePair() []string`

GetLabelNameValuePair returns the LabelNameValuePair field if non-nil, zero value otherwise.

### GetLabelNameValuePairOk

`func (o *AddPrometheusMonitorAttributeMetricRequest) GetLabelNameValuePairOk() (*[]string, bool)`

GetLabelNameValuePairOk returns a tuple with the LabelNameValuePair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelNameValuePair

`func (o *AddPrometheusMonitorAttributeMetricRequest) SetLabelNameValuePair(v []string)`

SetLabelNameValuePair sets LabelNameValuePair field to given value.

### HasLabelNameValuePair

`func (o *AddPrometheusMonitorAttributeMetricRequest) HasLabelNameValuePair() bool`

HasLabelNameValuePair returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


