# MonitoringEndpointListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** |  | [optional] 
**TotalResults** | Pointer to **float64** |  | [optional] 
**Resources** | Pointer to [**[]StatsdMonitoringEndpointResponse**](StatsdMonitoringEndpointResponse.md) |  | [optional] 

## Methods

### NewMonitoringEndpointListResponse

`func NewMonitoringEndpointListResponse() *MonitoringEndpointListResponse`

NewMonitoringEndpointListResponse instantiates a new MonitoringEndpointListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringEndpointListResponseWithDefaults

`func NewMonitoringEndpointListResponseWithDefaults() *MonitoringEndpointListResponse`

NewMonitoringEndpointListResponseWithDefaults instantiates a new MonitoringEndpointListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *MonitoringEndpointListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *MonitoringEndpointListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *MonitoringEndpointListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *MonitoringEndpointListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetTotalResults

`func (o *MonitoringEndpointListResponse) GetTotalResults() float64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *MonitoringEndpointListResponse) GetTotalResultsOk() (*float64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *MonitoringEndpointListResponse) SetTotalResults(v float64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *MonitoringEndpointListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *MonitoringEndpointListResponse) GetResources() []StatsdMonitoringEndpointResponse`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *MonitoringEndpointListResponse) GetResourcesOk() (*[]StatsdMonitoringEndpointResponse, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *MonitoringEndpointListResponse) SetResources(v []StatsdMonitoringEndpointResponse)`

SetResources sets Resources field to given value.

### HasResources

`func (o *MonitoringEndpointListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


