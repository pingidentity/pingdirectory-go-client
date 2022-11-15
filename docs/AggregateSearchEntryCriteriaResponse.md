# AggregateSearchEntryCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Search Entry Criteria | 
**Schemas** | [**[]EnumaggregateSearchEntryCriteriaSchemaUrn**](EnumaggregateSearchEntryCriteriaSchemaUrn.md) |  | 
**AllIncludedSearchEntryCriteria** | Pointer to **[]string** |  | [optional] 
**AnyIncludedSearchEntryCriteria** | Pointer to **[]string** |  | [optional] 
**NotAllIncludedSearchEntryCriteria** | Pointer to **[]string** |  | [optional] 
**NoneIncludedSearchEntryCriteria** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Search Entry Criteria | [optional] 

## Methods

### NewAggregateSearchEntryCriteriaResponse

`func NewAggregateSearchEntryCriteriaResponse(id string, schemas []EnumaggregateSearchEntryCriteriaSchemaUrn, ) *AggregateSearchEntryCriteriaResponse`

NewAggregateSearchEntryCriteriaResponse instantiates a new AggregateSearchEntryCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateSearchEntryCriteriaResponseWithDefaults

`func NewAggregateSearchEntryCriteriaResponseWithDefaults() *AggregateSearchEntryCriteriaResponse`

NewAggregateSearchEntryCriteriaResponseWithDefaults instantiates a new AggregateSearchEntryCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AggregateSearchEntryCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AggregateSearchEntryCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AggregateSearchEntryCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AggregateSearchEntryCriteriaResponse) GetSchemas() []EnumaggregateSearchEntryCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AggregateSearchEntryCriteriaResponse) GetSchemasOk() (*[]EnumaggregateSearchEntryCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AggregateSearchEntryCriteriaResponse) SetSchemas(v []EnumaggregateSearchEntryCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaResponse) GetAllIncludedSearchEntryCriteria() []string`

GetAllIncludedSearchEntryCriteria returns the AllIncludedSearchEntryCriteria field if non-nil, zero value otherwise.

### GetAllIncludedSearchEntryCriteriaOk

`func (o *AggregateSearchEntryCriteriaResponse) GetAllIncludedSearchEntryCriteriaOk() (*[]string, bool)`

GetAllIncludedSearchEntryCriteriaOk returns a tuple with the AllIncludedSearchEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaResponse) SetAllIncludedSearchEntryCriteria(v []string)`

SetAllIncludedSearchEntryCriteria sets AllIncludedSearchEntryCriteria field to given value.

### HasAllIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaResponse) HasAllIncludedSearchEntryCriteria() bool`

HasAllIncludedSearchEntryCriteria returns a boolean if a field has been set.

### GetAnyIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaResponse) GetAnyIncludedSearchEntryCriteria() []string`

GetAnyIncludedSearchEntryCriteria returns the AnyIncludedSearchEntryCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedSearchEntryCriteriaOk

`func (o *AggregateSearchEntryCriteriaResponse) GetAnyIncludedSearchEntryCriteriaOk() (*[]string, bool)`

GetAnyIncludedSearchEntryCriteriaOk returns a tuple with the AnyIncludedSearchEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaResponse) SetAnyIncludedSearchEntryCriteria(v []string)`

SetAnyIncludedSearchEntryCriteria sets AnyIncludedSearchEntryCriteria field to given value.

### HasAnyIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaResponse) HasAnyIncludedSearchEntryCriteria() bool`

HasAnyIncludedSearchEntryCriteria returns a boolean if a field has been set.

### GetNotAllIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaResponse) GetNotAllIncludedSearchEntryCriteria() []string`

GetNotAllIncludedSearchEntryCriteria returns the NotAllIncludedSearchEntryCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedSearchEntryCriteriaOk

`func (o *AggregateSearchEntryCriteriaResponse) GetNotAllIncludedSearchEntryCriteriaOk() (*[]string, bool)`

GetNotAllIncludedSearchEntryCriteriaOk returns a tuple with the NotAllIncludedSearchEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaResponse) SetNotAllIncludedSearchEntryCriteria(v []string)`

SetNotAllIncludedSearchEntryCriteria sets NotAllIncludedSearchEntryCriteria field to given value.

### HasNotAllIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaResponse) HasNotAllIncludedSearchEntryCriteria() bool`

HasNotAllIncludedSearchEntryCriteria returns a boolean if a field has been set.

### GetNoneIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaResponse) GetNoneIncludedSearchEntryCriteria() []string`

GetNoneIncludedSearchEntryCriteria returns the NoneIncludedSearchEntryCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedSearchEntryCriteriaOk

`func (o *AggregateSearchEntryCriteriaResponse) GetNoneIncludedSearchEntryCriteriaOk() (*[]string, bool)`

GetNoneIncludedSearchEntryCriteriaOk returns a tuple with the NoneIncludedSearchEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaResponse) SetNoneIncludedSearchEntryCriteria(v []string)`

SetNoneIncludedSearchEntryCriteria sets NoneIncludedSearchEntryCriteria field to given value.

### HasNoneIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaResponse) HasNoneIncludedSearchEntryCriteria() bool`

HasNoneIncludedSearchEntryCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *AggregateSearchEntryCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AggregateSearchEntryCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AggregateSearchEntryCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AggregateSearchEntryCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


