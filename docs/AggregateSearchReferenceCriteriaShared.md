# AggregateSearchReferenceCriteriaShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumaggregateSearchReferenceCriteriaSchemaUrn**](EnumaggregateSearchReferenceCriteriaSchemaUrn.md) |  | 
**AllIncludedSearchReferenceCriteria** | Pointer to **[]string** |  | [optional] 
**AnyIncludedSearchReferenceCriteria** | Pointer to **[]string** |  | [optional] 
**NotAllIncludedSearchReferenceCriteria** | Pointer to **[]string** |  | [optional] 
**NoneIncludedSearchReferenceCriteria** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Search Reference Criteria | [optional] 

## Methods

### NewAggregateSearchReferenceCriteriaShared

`func NewAggregateSearchReferenceCriteriaShared(schemas []EnumaggregateSearchReferenceCriteriaSchemaUrn, ) *AggregateSearchReferenceCriteriaShared`

NewAggregateSearchReferenceCriteriaShared instantiates a new AggregateSearchReferenceCriteriaShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateSearchReferenceCriteriaSharedWithDefaults

`func NewAggregateSearchReferenceCriteriaSharedWithDefaults() *AggregateSearchReferenceCriteriaShared`

NewAggregateSearchReferenceCriteriaSharedWithDefaults instantiates a new AggregateSearchReferenceCriteriaShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AggregateSearchReferenceCriteriaShared) GetSchemas() []EnumaggregateSearchReferenceCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AggregateSearchReferenceCriteriaShared) GetSchemasOk() (*[]EnumaggregateSearchReferenceCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AggregateSearchReferenceCriteriaShared) SetSchemas(v []EnumaggregateSearchReferenceCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaShared) GetAllIncludedSearchReferenceCriteria() []string`

GetAllIncludedSearchReferenceCriteria returns the AllIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetAllIncludedSearchReferenceCriteriaOk

`func (o *AggregateSearchReferenceCriteriaShared) GetAllIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetAllIncludedSearchReferenceCriteriaOk returns a tuple with the AllIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaShared) SetAllIncludedSearchReferenceCriteria(v []string)`

SetAllIncludedSearchReferenceCriteria sets AllIncludedSearchReferenceCriteria field to given value.

### HasAllIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaShared) HasAllIncludedSearchReferenceCriteria() bool`

HasAllIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetAnyIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaShared) GetAnyIncludedSearchReferenceCriteria() []string`

GetAnyIncludedSearchReferenceCriteria returns the AnyIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedSearchReferenceCriteriaOk

`func (o *AggregateSearchReferenceCriteriaShared) GetAnyIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetAnyIncludedSearchReferenceCriteriaOk returns a tuple with the AnyIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaShared) SetAnyIncludedSearchReferenceCriteria(v []string)`

SetAnyIncludedSearchReferenceCriteria sets AnyIncludedSearchReferenceCriteria field to given value.

### HasAnyIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaShared) HasAnyIncludedSearchReferenceCriteria() bool`

HasAnyIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetNotAllIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaShared) GetNotAllIncludedSearchReferenceCriteria() []string`

GetNotAllIncludedSearchReferenceCriteria returns the NotAllIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedSearchReferenceCriteriaOk

`func (o *AggregateSearchReferenceCriteriaShared) GetNotAllIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetNotAllIncludedSearchReferenceCriteriaOk returns a tuple with the NotAllIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaShared) SetNotAllIncludedSearchReferenceCriteria(v []string)`

SetNotAllIncludedSearchReferenceCriteria sets NotAllIncludedSearchReferenceCriteria field to given value.

### HasNotAllIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaShared) HasNotAllIncludedSearchReferenceCriteria() bool`

HasNotAllIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetNoneIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaShared) GetNoneIncludedSearchReferenceCriteria() []string`

GetNoneIncludedSearchReferenceCriteria returns the NoneIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedSearchReferenceCriteriaOk

`func (o *AggregateSearchReferenceCriteriaShared) GetNoneIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetNoneIncludedSearchReferenceCriteriaOk returns a tuple with the NoneIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaShared) SetNoneIncludedSearchReferenceCriteria(v []string)`

SetNoneIncludedSearchReferenceCriteria sets NoneIncludedSearchReferenceCriteria field to given value.

### HasNoneIncludedSearchReferenceCriteria

`func (o *AggregateSearchReferenceCriteriaShared) HasNoneIncludedSearchReferenceCriteria() bool`

HasNoneIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *AggregateSearchReferenceCriteriaShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AggregateSearchReferenceCriteriaShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AggregateSearchReferenceCriteriaShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AggregateSearchReferenceCriteriaShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


