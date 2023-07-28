# LogFieldBehaviorListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** |  | [optional] 
**TotalResults** | Pointer to **float64** |  | [optional] 
**Resources** | Pointer to [**[]AddLogFieldBehavior200Response**](AddLogFieldBehavior200Response.md) |  | [optional] 

## Methods

### NewLogFieldBehaviorListResponse

`func NewLogFieldBehaviorListResponse() *LogFieldBehaviorListResponse`

NewLogFieldBehaviorListResponse instantiates a new LogFieldBehaviorListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogFieldBehaviorListResponseWithDefaults

`func NewLogFieldBehaviorListResponseWithDefaults() *LogFieldBehaviorListResponse`

NewLogFieldBehaviorListResponseWithDefaults instantiates a new LogFieldBehaviorListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *LogFieldBehaviorListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LogFieldBehaviorListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LogFieldBehaviorListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *LogFieldBehaviorListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetTotalResults

`func (o *LogFieldBehaviorListResponse) GetTotalResults() float64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *LogFieldBehaviorListResponse) GetTotalResultsOk() (*float64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *LogFieldBehaviorListResponse) SetTotalResults(v float64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *LogFieldBehaviorListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *LogFieldBehaviorListResponse) GetResources() []AddLogFieldBehavior200Response`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *LogFieldBehaviorListResponse) GetResourcesOk() (*[]AddLogFieldBehavior200Response, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *LogFieldBehaviorListResponse) SetResources(v []AddLogFieldBehavior200Response)`

SetResources sets Resources field to given value.

### HasResources

`func (o *LogFieldBehaviorListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


