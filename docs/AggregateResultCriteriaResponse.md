# AggregateResultCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Result Criteria | 
**Schemas** | [**[]EnumaggregateResultCriteriaSchemaUrn**](EnumaggregateResultCriteriaSchemaUrn.md) |  | 
**AllIncludedResultCriteria** | Pointer to **[]string** |  | [optional] 
**AnyIncludedResultCriteria** | Pointer to **[]string** |  | [optional] 
**NotAllIncludedResultCriteria** | Pointer to **[]string** |  | [optional] 
**NoneIncludedResultCriteria** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Result Criteria | [optional] 

## Methods

### NewAggregateResultCriteriaResponse

`func NewAggregateResultCriteriaResponse(id string, schemas []EnumaggregateResultCriteriaSchemaUrn, ) *AggregateResultCriteriaResponse`

NewAggregateResultCriteriaResponse instantiates a new AggregateResultCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateResultCriteriaResponseWithDefaults

`func NewAggregateResultCriteriaResponseWithDefaults() *AggregateResultCriteriaResponse`

NewAggregateResultCriteriaResponseWithDefaults instantiates a new AggregateResultCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AggregateResultCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AggregateResultCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AggregateResultCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AggregateResultCriteriaResponse) GetSchemas() []EnumaggregateResultCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AggregateResultCriteriaResponse) GetSchemasOk() (*[]EnumaggregateResultCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AggregateResultCriteriaResponse) SetSchemas(v []EnumaggregateResultCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllIncludedResultCriteria

`func (o *AggregateResultCriteriaResponse) GetAllIncludedResultCriteria() []string`

GetAllIncludedResultCriteria returns the AllIncludedResultCriteria field if non-nil, zero value otherwise.

### GetAllIncludedResultCriteriaOk

`func (o *AggregateResultCriteriaResponse) GetAllIncludedResultCriteriaOk() (*[]string, bool)`

GetAllIncludedResultCriteriaOk returns a tuple with the AllIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedResultCriteria

`func (o *AggregateResultCriteriaResponse) SetAllIncludedResultCriteria(v []string)`

SetAllIncludedResultCriteria sets AllIncludedResultCriteria field to given value.

### HasAllIncludedResultCriteria

`func (o *AggregateResultCriteriaResponse) HasAllIncludedResultCriteria() bool`

HasAllIncludedResultCriteria returns a boolean if a field has been set.

### GetAnyIncludedResultCriteria

`func (o *AggregateResultCriteriaResponse) GetAnyIncludedResultCriteria() []string`

GetAnyIncludedResultCriteria returns the AnyIncludedResultCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedResultCriteriaOk

`func (o *AggregateResultCriteriaResponse) GetAnyIncludedResultCriteriaOk() (*[]string, bool)`

GetAnyIncludedResultCriteriaOk returns a tuple with the AnyIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedResultCriteria

`func (o *AggregateResultCriteriaResponse) SetAnyIncludedResultCriteria(v []string)`

SetAnyIncludedResultCriteria sets AnyIncludedResultCriteria field to given value.

### HasAnyIncludedResultCriteria

`func (o *AggregateResultCriteriaResponse) HasAnyIncludedResultCriteria() bool`

HasAnyIncludedResultCriteria returns a boolean if a field has been set.

### GetNotAllIncludedResultCriteria

`func (o *AggregateResultCriteriaResponse) GetNotAllIncludedResultCriteria() []string`

GetNotAllIncludedResultCriteria returns the NotAllIncludedResultCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedResultCriteriaOk

`func (o *AggregateResultCriteriaResponse) GetNotAllIncludedResultCriteriaOk() (*[]string, bool)`

GetNotAllIncludedResultCriteriaOk returns a tuple with the NotAllIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedResultCriteria

`func (o *AggregateResultCriteriaResponse) SetNotAllIncludedResultCriteria(v []string)`

SetNotAllIncludedResultCriteria sets NotAllIncludedResultCriteria field to given value.

### HasNotAllIncludedResultCriteria

`func (o *AggregateResultCriteriaResponse) HasNotAllIncludedResultCriteria() bool`

HasNotAllIncludedResultCriteria returns a boolean if a field has been set.

### GetNoneIncludedResultCriteria

`func (o *AggregateResultCriteriaResponse) GetNoneIncludedResultCriteria() []string`

GetNoneIncludedResultCriteria returns the NoneIncludedResultCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedResultCriteriaOk

`func (o *AggregateResultCriteriaResponse) GetNoneIncludedResultCriteriaOk() (*[]string, bool)`

GetNoneIncludedResultCriteriaOk returns a tuple with the NoneIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedResultCriteria

`func (o *AggregateResultCriteriaResponse) SetNoneIncludedResultCriteria(v []string)`

SetNoneIncludedResultCriteria sets NoneIncludedResultCriteria field to given value.

### HasNoneIncludedResultCriteria

`func (o *AggregateResultCriteriaResponse) HasNoneIncludedResultCriteria() bool`

HasNoneIncludedResultCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *AggregateResultCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AggregateResultCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AggregateResultCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AggregateResultCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


