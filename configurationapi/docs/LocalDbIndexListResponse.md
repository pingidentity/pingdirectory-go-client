# LocalDbIndexListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** |  | [optional] 
**TotalResults** | Pointer to **float64** |  | [optional] 
**Resources** | Pointer to [**[]LocalDbIndexResponse**](LocalDbIndexResponse.md) |  | [optional] 

## Methods

### NewLocalDbIndexListResponse

`func NewLocalDbIndexListResponse() *LocalDbIndexListResponse`

NewLocalDbIndexListResponse instantiates a new LocalDbIndexListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalDbIndexListResponseWithDefaults

`func NewLocalDbIndexListResponseWithDefaults() *LocalDbIndexListResponse`

NewLocalDbIndexListResponseWithDefaults instantiates a new LocalDbIndexListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *LocalDbIndexListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LocalDbIndexListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LocalDbIndexListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *LocalDbIndexListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetTotalResults

`func (o *LocalDbIndexListResponse) GetTotalResults() float64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *LocalDbIndexListResponse) GetTotalResultsOk() (*float64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *LocalDbIndexListResponse) SetTotalResults(v float64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *LocalDbIndexListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *LocalDbIndexListResponse) GetResources() []LocalDbIndexResponse`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *LocalDbIndexListResponse) GetResourcesOk() (*[]LocalDbIndexResponse, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *LocalDbIndexListResponse) SetResources(v []LocalDbIndexResponse)`

SetResources sets Resources field to given value.

### HasResources

`func (o *LocalDbIndexListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


