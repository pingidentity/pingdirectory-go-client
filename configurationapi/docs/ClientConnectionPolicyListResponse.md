# ClientConnectionPolicyListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** |  | [optional] 
**TotalResults** | Pointer to **float64** |  | [optional] 
**Resources** | Pointer to [**[]ClientConnectionPolicyResponse**](ClientConnectionPolicyResponse.md) |  | [optional] 

## Methods

### NewClientConnectionPolicyListResponse

`func NewClientConnectionPolicyListResponse() *ClientConnectionPolicyListResponse`

NewClientConnectionPolicyListResponse instantiates a new ClientConnectionPolicyListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConnectionPolicyListResponseWithDefaults

`func NewClientConnectionPolicyListResponseWithDefaults() *ClientConnectionPolicyListResponse`

NewClientConnectionPolicyListResponseWithDefaults instantiates a new ClientConnectionPolicyListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ClientConnectionPolicyListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ClientConnectionPolicyListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ClientConnectionPolicyListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ClientConnectionPolicyListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetTotalResults

`func (o *ClientConnectionPolicyListResponse) GetTotalResults() float64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ClientConnectionPolicyListResponse) GetTotalResultsOk() (*float64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ClientConnectionPolicyListResponse) SetTotalResults(v float64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *ClientConnectionPolicyListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *ClientConnectionPolicyListResponse) GetResources() []ClientConnectionPolicyResponse`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ClientConnectionPolicyListResponse) GetResourcesOk() (*[]ClientConnectionPolicyResponse, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ClientConnectionPolicyListResponse) SetResources(v []ClientConnectionPolicyResponse)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ClientConnectionPolicyListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


