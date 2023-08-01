# SynchronizationProviderListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** |  | [optional] 
**TotalResults** | Pointer to **float64** |  | [optional] 
**Resources** | Pointer to [**[]GetSynchronizationProvider200Response**](GetSynchronizationProvider200Response.md) |  | [optional] 

## Methods

### NewSynchronizationProviderListResponse

`func NewSynchronizationProviderListResponse() *SynchronizationProviderListResponse`

NewSynchronizationProviderListResponse instantiates a new SynchronizationProviderListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSynchronizationProviderListResponseWithDefaults

`func NewSynchronizationProviderListResponseWithDefaults() *SynchronizationProviderListResponse`

NewSynchronizationProviderListResponseWithDefaults instantiates a new SynchronizationProviderListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SynchronizationProviderListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SynchronizationProviderListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SynchronizationProviderListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *SynchronizationProviderListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetTotalResults

`func (o *SynchronizationProviderListResponse) GetTotalResults() float64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *SynchronizationProviderListResponse) GetTotalResultsOk() (*float64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *SynchronizationProviderListResponse) SetTotalResults(v float64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *SynchronizationProviderListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *SynchronizationProviderListResponse) GetResources() []GetSynchronizationProvider200Response`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *SynchronizationProviderListResponse) GetResourcesOk() (*[]GetSynchronizationProvider200Response, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *SynchronizationProviderListResponse) SetResources(v []GetSynchronizationProvider200Response)`

SetResources sets Resources field to given value.

### HasResources

`func (o *SynchronizationProviderListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


