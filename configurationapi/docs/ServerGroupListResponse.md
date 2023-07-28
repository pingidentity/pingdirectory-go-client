# ServerGroupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** |  | [optional] 
**TotalResults** | Pointer to **float64** |  | [optional] 
**Resources** | Pointer to [**[]ServerGroupResponse**](ServerGroupResponse.md) |  | [optional] 

## Methods

### NewServerGroupListResponse

`func NewServerGroupListResponse() *ServerGroupListResponse`

NewServerGroupListResponse instantiates a new ServerGroupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerGroupListResponseWithDefaults

`func NewServerGroupListResponseWithDefaults() *ServerGroupListResponse`

NewServerGroupListResponseWithDefaults instantiates a new ServerGroupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ServerGroupListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ServerGroupListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ServerGroupListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ServerGroupListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetTotalResults

`func (o *ServerGroupListResponse) GetTotalResults() float64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ServerGroupListResponse) GetTotalResultsOk() (*float64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ServerGroupListResponse) SetTotalResults(v float64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *ServerGroupListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *ServerGroupListResponse) GetResources() []ServerGroupResponse`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ServerGroupListResponse) GetResourcesOk() (*[]ServerGroupResponse, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ServerGroupListResponse) SetResources(v []ServerGroupResponse)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ServerGroupListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


