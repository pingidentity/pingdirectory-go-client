# OauthTokenHandlerListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** |  | [optional] 
**TotalResults** | Pointer to **float64** |  | [optional] 
**Resources** | Pointer to [**[]AddOauthTokenHandler200Response**](AddOauthTokenHandler200Response.md) |  | [optional] 

## Methods

### NewOauthTokenHandlerListResponse

`func NewOauthTokenHandlerListResponse() *OauthTokenHandlerListResponse`

NewOauthTokenHandlerListResponse instantiates a new OauthTokenHandlerListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthTokenHandlerListResponseWithDefaults

`func NewOauthTokenHandlerListResponseWithDefaults() *OauthTokenHandlerListResponse`

NewOauthTokenHandlerListResponseWithDefaults instantiates a new OauthTokenHandlerListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *OauthTokenHandlerListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *OauthTokenHandlerListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *OauthTokenHandlerListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *OauthTokenHandlerListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetTotalResults

`func (o *OauthTokenHandlerListResponse) GetTotalResults() float64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *OauthTokenHandlerListResponse) GetTotalResultsOk() (*float64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *OauthTokenHandlerListResponse) SetTotalResults(v float64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *OauthTokenHandlerListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *OauthTokenHandlerListResponse) GetResources() []AddOauthTokenHandler200Response`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *OauthTokenHandlerListResponse) GetResourcesOk() (*[]AddOauthTokenHandler200Response, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *OauthTokenHandlerListResponse) SetResources(v []AddOauthTokenHandler200Response)`

SetResources sets Resources field to given value.

### HasResources

`func (o *OauthTokenHandlerListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


