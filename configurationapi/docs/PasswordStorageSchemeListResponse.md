# PasswordStorageSchemeListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** |  | [optional] 
**TotalResults** | Pointer to **float64** |  | [optional] 
**Resources** | Pointer to [**[]GetPasswordStorageScheme200Response**](GetPasswordStorageScheme200Response.md) |  | [optional] 

## Methods

### NewPasswordStorageSchemeListResponse

`func NewPasswordStorageSchemeListResponse() *PasswordStorageSchemeListResponse`

NewPasswordStorageSchemeListResponse instantiates a new PasswordStorageSchemeListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordStorageSchemeListResponseWithDefaults

`func NewPasswordStorageSchemeListResponseWithDefaults() *PasswordStorageSchemeListResponse`

NewPasswordStorageSchemeListResponseWithDefaults instantiates a new PasswordStorageSchemeListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *PasswordStorageSchemeListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PasswordStorageSchemeListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PasswordStorageSchemeListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *PasswordStorageSchemeListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetTotalResults

`func (o *PasswordStorageSchemeListResponse) GetTotalResults() float64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *PasswordStorageSchemeListResponse) GetTotalResultsOk() (*float64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *PasswordStorageSchemeListResponse) SetTotalResults(v float64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *PasswordStorageSchemeListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *PasswordStorageSchemeListResponse) GetResources() []GetPasswordStorageScheme200Response`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PasswordStorageSchemeListResponse) GetResourcesOk() (*[]GetPasswordStorageScheme200Response, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PasswordStorageSchemeListResponse) SetResources(v []GetPasswordStorageScheme200Response)`

SetResources sets Resources field to given value.

### HasResources

`func (o *PasswordStorageSchemeListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


