# ConjurAuthenticationMethodListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** |  | [optional] 
**TotalResults** | Pointer to **float64** |  | [optional] 
**Resources** | Pointer to [**[]ApiKeyConjurAuthenticationMethodResponse**](ApiKeyConjurAuthenticationMethodResponse.md) |  | [optional] 

## Methods

### NewConjurAuthenticationMethodListResponse

`func NewConjurAuthenticationMethodListResponse() *ConjurAuthenticationMethodListResponse`

NewConjurAuthenticationMethodListResponse instantiates a new ConjurAuthenticationMethodListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConjurAuthenticationMethodListResponseWithDefaults

`func NewConjurAuthenticationMethodListResponseWithDefaults() *ConjurAuthenticationMethodListResponse`

NewConjurAuthenticationMethodListResponseWithDefaults instantiates a new ConjurAuthenticationMethodListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ConjurAuthenticationMethodListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ConjurAuthenticationMethodListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ConjurAuthenticationMethodListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ConjurAuthenticationMethodListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetTotalResults

`func (o *ConjurAuthenticationMethodListResponse) GetTotalResults() float64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ConjurAuthenticationMethodListResponse) GetTotalResultsOk() (*float64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ConjurAuthenticationMethodListResponse) SetTotalResults(v float64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *ConjurAuthenticationMethodListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *ConjurAuthenticationMethodListResponse) GetResources() []ApiKeyConjurAuthenticationMethodResponse`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ConjurAuthenticationMethodListResponse) GetResourcesOk() (*[]ApiKeyConjurAuthenticationMethodResponse, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ConjurAuthenticationMethodListResponse) SetResources(v []ApiKeyConjurAuthenticationMethodResponse)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ConjurAuthenticationMethodListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


