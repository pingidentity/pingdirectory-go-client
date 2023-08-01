# ServerInstanceListenerListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** |  | [optional] 
**TotalResults** | Pointer to **float64** |  | [optional] 
**Resources** | Pointer to [**[]GetServerInstanceListener200Response**](GetServerInstanceListener200Response.md) |  | [optional] 

## Methods

### NewServerInstanceListenerListResponse

`func NewServerInstanceListenerListResponse() *ServerInstanceListenerListResponse`

NewServerInstanceListenerListResponse instantiates a new ServerInstanceListenerListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceListenerListResponseWithDefaults

`func NewServerInstanceListenerListResponseWithDefaults() *ServerInstanceListenerListResponse`

NewServerInstanceListenerListResponseWithDefaults instantiates a new ServerInstanceListenerListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ServerInstanceListenerListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ServerInstanceListenerListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ServerInstanceListenerListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ServerInstanceListenerListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetTotalResults

`func (o *ServerInstanceListenerListResponse) GetTotalResults() float64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ServerInstanceListenerListResponse) GetTotalResultsOk() (*float64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ServerInstanceListenerListResponse) SetTotalResults(v float64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *ServerInstanceListenerListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *ServerInstanceListenerListResponse) GetResources() []GetServerInstanceListener200Response`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ServerInstanceListenerListResponse) GetResourcesOk() (*[]GetServerInstanceListener200Response, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ServerInstanceListenerListResponse) SetResources(v []GetServerInstanceListener200Response)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ServerInstanceListenerListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


