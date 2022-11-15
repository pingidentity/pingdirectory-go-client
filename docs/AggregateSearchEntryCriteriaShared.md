# AggregateSearchEntryCriteriaShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumaggregateSearchEntryCriteriaSchemaUrn**](EnumaggregateSearchEntryCriteriaSchemaUrn.md) |  | 
**AllIncludedSearchEntryCriteria** | Pointer to **[]string** |  | [optional] 
**AnyIncludedSearchEntryCriteria** | Pointer to **[]string** |  | [optional] 
**NotAllIncludedSearchEntryCriteria** | Pointer to **[]string** |  | [optional] 
**NoneIncludedSearchEntryCriteria** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Search Entry Criteria | [optional] 

## Methods

### NewAggregateSearchEntryCriteriaShared

`func NewAggregateSearchEntryCriteriaShared(schemas []EnumaggregateSearchEntryCriteriaSchemaUrn, ) *AggregateSearchEntryCriteriaShared`

NewAggregateSearchEntryCriteriaShared instantiates a new AggregateSearchEntryCriteriaShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateSearchEntryCriteriaSharedWithDefaults

`func NewAggregateSearchEntryCriteriaSharedWithDefaults() *AggregateSearchEntryCriteriaShared`

NewAggregateSearchEntryCriteriaSharedWithDefaults instantiates a new AggregateSearchEntryCriteriaShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AggregateSearchEntryCriteriaShared) GetSchemas() []EnumaggregateSearchEntryCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AggregateSearchEntryCriteriaShared) GetSchemasOk() (*[]EnumaggregateSearchEntryCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AggregateSearchEntryCriteriaShared) SetSchemas(v []EnumaggregateSearchEntryCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaShared) GetAllIncludedSearchEntryCriteria() []string`

GetAllIncludedSearchEntryCriteria returns the AllIncludedSearchEntryCriteria field if non-nil, zero value otherwise.

### GetAllIncludedSearchEntryCriteriaOk

`func (o *AggregateSearchEntryCriteriaShared) GetAllIncludedSearchEntryCriteriaOk() (*[]string, bool)`

GetAllIncludedSearchEntryCriteriaOk returns a tuple with the AllIncludedSearchEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaShared) SetAllIncludedSearchEntryCriteria(v []string)`

SetAllIncludedSearchEntryCriteria sets AllIncludedSearchEntryCriteria field to given value.

### HasAllIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaShared) HasAllIncludedSearchEntryCriteria() bool`

HasAllIncludedSearchEntryCriteria returns a boolean if a field has been set.

### GetAnyIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaShared) GetAnyIncludedSearchEntryCriteria() []string`

GetAnyIncludedSearchEntryCriteria returns the AnyIncludedSearchEntryCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedSearchEntryCriteriaOk

`func (o *AggregateSearchEntryCriteriaShared) GetAnyIncludedSearchEntryCriteriaOk() (*[]string, bool)`

GetAnyIncludedSearchEntryCriteriaOk returns a tuple with the AnyIncludedSearchEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaShared) SetAnyIncludedSearchEntryCriteria(v []string)`

SetAnyIncludedSearchEntryCriteria sets AnyIncludedSearchEntryCriteria field to given value.

### HasAnyIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaShared) HasAnyIncludedSearchEntryCriteria() bool`

HasAnyIncludedSearchEntryCriteria returns a boolean if a field has been set.

### GetNotAllIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaShared) GetNotAllIncludedSearchEntryCriteria() []string`

GetNotAllIncludedSearchEntryCriteria returns the NotAllIncludedSearchEntryCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedSearchEntryCriteriaOk

`func (o *AggregateSearchEntryCriteriaShared) GetNotAllIncludedSearchEntryCriteriaOk() (*[]string, bool)`

GetNotAllIncludedSearchEntryCriteriaOk returns a tuple with the NotAllIncludedSearchEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaShared) SetNotAllIncludedSearchEntryCriteria(v []string)`

SetNotAllIncludedSearchEntryCriteria sets NotAllIncludedSearchEntryCriteria field to given value.

### HasNotAllIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaShared) HasNotAllIncludedSearchEntryCriteria() bool`

HasNotAllIncludedSearchEntryCriteria returns a boolean if a field has been set.

### GetNoneIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaShared) GetNoneIncludedSearchEntryCriteria() []string`

GetNoneIncludedSearchEntryCriteria returns the NoneIncludedSearchEntryCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedSearchEntryCriteriaOk

`func (o *AggregateSearchEntryCriteriaShared) GetNoneIncludedSearchEntryCriteriaOk() (*[]string, bool)`

GetNoneIncludedSearchEntryCriteriaOk returns a tuple with the NoneIncludedSearchEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaShared) SetNoneIncludedSearchEntryCriteria(v []string)`

SetNoneIncludedSearchEntryCriteria sets NoneIncludedSearchEntryCriteria field to given value.

### HasNoneIncludedSearchEntryCriteria

`func (o *AggregateSearchEntryCriteriaShared) HasNoneIncludedSearchEntryCriteria() bool`

HasNoneIncludedSearchEntryCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *AggregateSearchEntryCriteriaShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AggregateSearchEntryCriteriaShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AggregateSearchEntryCriteriaShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AggregateSearchEntryCriteriaShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


