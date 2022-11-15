# AggregateSearchReferenceCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Search Reference Criteria | 
**Schemas** | [**[]EnumaggregateSearchReferenceCriteriaSchemaUrn**](EnumaggregateSearchReferenceCriteriaSchemaUrn.md) |  | 
**AllIncludedSearchReferenceCriteria** | Pointer to **[]string** |  | [optional] 
**AnyIncludedSearchReferenceCriteria** | Pointer to **[]string** |  | [optional] 
**NotAllIncludedSearchReferenceCriteria** | Pointer to **[]string** |  | [optional] 
**NoneIncludedSearchReferenceCriteria** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Search Reference Criteria | [optional] 

## Methods

### NewAggregateSearchReferenceCriteriaResponse

`func NewAggregateSearchReferenceCriteriaResponse(id string, schemas []EnumaggregateSearchReferenceCriteriaSchemaUrn, ) *AggregateSearchReferenceCriteriaResponse`

NewAggregateSearchReferenceCriteriaResponse instantiates a new AggregateSearchReferenceCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateSearchReferenceCriteriaResponseWithDefaults

`func NewAggregateSearchReferenceCriteriaResponseWithDefaults() *AggregateSearchReferenceCriteriaResponse`

NewAggregateSearchReferenceCriteriaResponseWithDefaults instantiates a new AggregateSearchReferenceCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AggregateSearchReferenceCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AggregateSearchReferenceCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AggregateSearchReferenceCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AggregateSearchReferenceCriteriaResponse) GetSchemas() []EnumaggregateSearchReferenceCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AggregateSearchReferenceCriteriaResponse) GetSchemasOk() (*[]EnumaggregateSearchReferenceCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AggregateSearchReferenceCriteriaResponse) SetSchemas(v []EnumaggregateSearchReferenceCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaResponse) GetAllIncludedSearchReferenceCriteria() []string`

GetAllIncludedSearchReferenceCriteria returns the AllIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetAllIncludedSearchReferenceCriteriaOk

`func (o *AggregateSearchReferenceCriteriaResponse) GetAllIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetAllIncludedSearchReferenceCriteriaOk returns a tuple with the AllIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaResponse) SetAllIncludedSearchReferenceCriteria(v []string)`

SetAllIncludedSearchReferenceCriteria sets AllIncludedSearchReferenceCriteria field to given value.

### HasAllIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaResponse) HasAllIncludedSearchReferenceCriteria() bool`

HasAllIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetAnyIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaResponse) GetAnyIncludedSearchReferenceCriteria() []string`

GetAnyIncludedSearchReferenceCriteria returns the AnyIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedSearchReferenceCriteriaOk

`func (o *AggregateSearchReferenceCriteriaResponse) GetAnyIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetAnyIncludedSearchReferenceCriteriaOk returns a tuple with the AnyIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaResponse) SetAnyIncludedSearchReferenceCriteria(v []string)`

SetAnyIncludedSearchReferenceCriteria sets AnyIncludedSearchReferenceCriteria field to given value.

### HasAnyIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaResponse) HasAnyIncludedSearchReferenceCriteria() bool`

HasAnyIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetNotAllIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaResponse) GetNotAllIncludedSearchReferenceCriteria() []string`

GetNotAllIncludedSearchReferenceCriteria returns the NotAllIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedSearchReferenceCriteriaOk

`func (o *AggregateSearchReferenceCriteriaResponse) GetNotAllIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetNotAllIncludedSearchReferenceCriteriaOk returns a tuple with the NotAllIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaResponse) SetNotAllIncludedSearchReferenceCriteria(v []string)`

SetNotAllIncludedSearchReferenceCriteria sets NotAllIncludedSearchReferenceCriteria field to given value.

### HasNotAllIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaResponse) HasNotAllIncludedSearchReferenceCriteria() bool`

HasNotAllIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetNoneIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaResponse) GetNoneIncludedSearchReferenceCriteria() []string`

GetNoneIncludedSearchReferenceCriteria returns the NoneIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedSearchReferenceCriteriaOk

`func (o *AggregateSearchReferenceCriteriaResponse) GetNoneIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetNoneIncludedSearchReferenceCriteriaOk returns a tuple with the NoneIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaResponse) SetNoneIncludedSearchReferenceCriteria(v []string)`

SetNoneIncludedSearchReferenceCriteria sets NoneIncludedSearchReferenceCriteria field to given value.

### HasNoneIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaResponse) HasNoneIncludedSearchReferenceCriteria() bool`

HasNoneIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *AggregateSearchReferenceCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AggregateSearchReferenceCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AggregateSearchReferenceCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AggregateSearchReferenceCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


