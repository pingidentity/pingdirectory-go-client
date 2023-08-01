# LocalDbCompositeIndexListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** |  | [optional] 
**TotalResults** | Pointer to **float64** |  | [optional] 
**Resources** | Pointer to [**[]LocalDbCompositeIndexResponse**](LocalDbCompositeIndexResponse.md) |  | [optional] 

## Methods

### NewLocalDbCompositeIndexListResponse

`func NewLocalDbCompositeIndexListResponse() *LocalDbCompositeIndexListResponse`

NewLocalDbCompositeIndexListResponse instantiates a new LocalDbCompositeIndexListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalDbCompositeIndexListResponseWithDefaults

`func NewLocalDbCompositeIndexListResponseWithDefaults() *LocalDbCompositeIndexListResponse`

NewLocalDbCompositeIndexListResponseWithDefaults instantiates a new LocalDbCompositeIndexListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *LocalDbCompositeIndexListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LocalDbCompositeIndexListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LocalDbCompositeIndexListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *LocalDbCompositeIndexListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetTotalResults

`func (o *LocalDbCompositeIndexListResponse) GetTotalResults() float64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *LocalDbCompositeIndexListResponse) GetTotalResultsOk() (*float64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *LocalDbCompositeIndexListResponse) SetTotalResults(v float64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *LocalDbCompositeIndexListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *LocalDbCompositeIndexListResponse) GetResources() []LocalDbCompositeIndexResponse`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *LocalDbCompositeIndexListResponse) GetResourcesOk() (*[]LocalDbCompositeIndexResponse, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *LocalDbCompositeIndexListResponse) SetResources(v []LocalDbCompositeIndexResponse)`

SetResources sets Resources field to given value.

### HasResources

`func (o *LocalDbCompositeIndexListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


