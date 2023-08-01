# ConnectionHandlerListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** |  | [optional] 
**TotalResults** | Pointer to **float64** |  | [optional] 
**Resources** | Pointer to [**[]ConnectionHandlerListResponseResourcesInner**](ConnectionHandlerListResponseResourcesInner.md) |  | [optional] 

## Methods

### NewConnectionHandlerListResponse

`func NewConnectionHandlerListResponse() *ConnectionHandlerListResponse`

NewConnectionHandlerListResponse instantiates a new ConnectionHandlerListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionHandlerListResponseWithDefaults

`func NewConnectionHandlerListResponseWithDefaults() *ConnectionHandlerListResponse`

NewConnectionHandlerListResponseWithDefaults instantiates a new ConnectionHandlerListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ConnectionHandlerListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ConnectionHandlerListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ConnectionHandlerListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ConnectionHandlerListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetTotalResults

`func (o *ConnectionHandlerListResponse) GetTotalResults() float64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ConnectionHandlerListResponse) GetTotalResultsOk() (*float64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ConnectionHandlerListResponse) SetTotalResults(v float64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *ConnectionHandlerListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *ConnectionHandlerListResponse) GetResources() []ConnectionHandlerListResponseResourcesInner`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ConnectionHandlerListResponse) GetResourcesOk() (*[]ConnectionHandlerListResponseResourcesInner, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ConnectionHandlerListResponse) SetResources(v []ConnectionHandlerListResponseResourcesInner)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ConnectionHandlerListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


