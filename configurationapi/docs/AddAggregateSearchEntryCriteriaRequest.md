# AddAggregateSearchEntryCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumaggregateSearchEntryCriteriaSchemaUrn**](EnumaggregateSearchEntryCriteriaSchemaUrn.md) |  | 
**AllIncludedSearchEntryCriteria** | Pointer to **[]string** | Specifies a search entry criteria object that must match the associated search result entry in order to match the aggregate search entry criteria. If one or more all-included search entry criteria objects are provided, then a search result entry must match all of them in order to match the aggregate search entry criteria. | [optional] 
**AnyIncludedSearchEntryCriteria** | Pointer to **[]string** | Specifies a search entry criteria object that may match the associated search result entry in order to match the aggregate search entry criteria. If one or more any-included search entry criteria objects are provided, then a search result entry must match at least one of them in order to match the aggregate search entry criteria. | [optional] 
**NotAllIncludedSearchEntryCriteria** | Pointer to **[]string** | Specifies a search entry criteria object that should not match the associated search result entry in order to match the aggregate search entry criteria. If one or more not-all-included search entry criteria objects are provided, then a search result entry must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate search entry criteria. | [optional] 
**NoneIncludedSearchEntryCriteria** | Pointer to **[]string** | Specifies a search entry criteria object that must not match the associated search result entry in order to match the aggregate search entry criteria. If one or more none-included search entry criteria objects are provided, then a search result entry must not match any of them in order to match the aggregate search entry criteria. | [optional] 
**Description** | Pointer to **string** | A description for this Search Entry Criteria | [optional] 
**CriteriaName** | **string** | Name of the new Search Entry Criteria | 

## Methods

### NewAddAggregateSearchEntryCriteriaRequest

`func NewAddAggregateSearchEntryCriteriaRequest(schemas []EnumaggregateSearchEntryCriteriaSchemaUrn, criteriaName string, ) *AddAggregateSearchEntryCriteriaRequest`

NewAddAggregateSearchEntryCriteriaRequest instantiates a new AddAggregateSearchEntryCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAggregateSearchEntryCriteriaRequestWithDefaults

`func NewAddAggregateSearchEntryCriteriaRequestWithDefaults() *AddAggregateSearchEntryCriteriaRequest`

NewAddAggregateSearchEntryCriteriaRequestWithDefaults instantiates a new AddAggregateSearchEntryCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddAggregateSearchEntryCriteriaRequest) GetSchemas() []EnumaggregateSearchEntryCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAggregateSearchEntryCriteriaRequest) GetSchemasOk() (*[]EnumaggregateSearchEntryCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAggregateSearchEntryCriteriaRequest) SetSchemas(v []EnumaggregateSearchEntryCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllIncludedSearchEntryCriteria

`func (o *AddAggregateSearchEntryCriteriaRequest) GetAllIncludedSearchEntryCriteria() []string`

GetAllIncludedSearchEntryCriteria returns the AllIncludedSearchEntryCriteria field if non-nil, zero value otherwise.

### GetAllIncludedSearchEntryCriteriaOk

`func (o *AddAggregateSearchEntryCriteriaRequest) GetAllIncludedSearchEntryCriteriaOk() (*[]string, bool)`

GetAllIncludedSearchEntryCriteriaOk returns a tuple with the AllIncludedSearchEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedSearchEntryCriteria

`func (o *AddAggregateSearchEntryCriteriaRequest) SetAllIncludedSearchEntryCriteria(v []string)`

SetAllIncludedSearchEntryCriteria sets AllIncludedSearchEntryCriteria field to given value.

### HasAllIncludedSearchEntryCriteria

`func (o *AddAggregateSearchEntryCriteriaRequest) HasAllIncludedSearchEntryCriteria() bool`

HasAllIncludedSearchEntryCriteria returns a boolean if a field has been set.

### GetAnyIncludedSearchEntryCriteria

`func (o *AddAggregateSearchEntryCriteriaRequest) GetAnyIncludedSearchEntryCriteria() []string`

GetAnyIncludedSearchEntryCriteria returns the AnyIncludedSearchEntryCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedSearchEntryCriteriaOk

`func (o *AddAggregateSearchEntryCriteriaRequest) GetAnyIncludedSearchEntryCriteriaOk() (*[]string, bool)`

GetAnyIncludedSearchEntryCriteriaOk returns a tuple with the AnyIncludedSearchEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedSearchEntryCriteria

`func (o *AddAggregateSearchEntryCriteriaRequest) SetAnyIncludedSearchEntryCriteria(v []string)`

SetAnyIncludedSearchEntryCriteria sets AnyIncludedSearchEntryCriteria field to given value.

### HasAnyIncludedSearchEntryCriteria

`func (o *AddAggregateSearchEntryCriteriaRequest) HasAnyIncludedSearchEntryCriteria() bool`

HasAnyIncludedSearchEntryCriteria returns a boolean if a field has been set.

### GetNotAllIncludedSearchEntryCriteria

`func (o *AddAggregateSearchEntryCriteriaRequest) GetNotAllIncludedSearchEntryCriteria() []string`

GetNotAllIncludedSearchEntryCriteria returns the NotAllIncludedSearchEntryCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedSearchEntryCriteriaOk

`func (o *AddAggregateSearchEntryCriteriaRequest) GetNotAllIncludedSearchEntryCriteriaOk() (*[]string, bool)`

GetNotAllIncludedSearchEntryCriteriaOk returns a tuple with the NotAllIncludedSearchEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedSearchEntryCriteria

`func (o *AddAggregateSearchEntryCriteriaRequest) SetNotAllIncludedSearchEntryCriteria(v []string)`

SetNotAllIncludedSearchEntryCriteria sets NotAllIncludedSearchEntryCriteria field to given value.

### HasNotAllIncludedSearchEntryCriteria

`func (o *AddAggregateSearchEntryCriteriaRequest) HasNotAllIncludedSearchEntryCriteria() bool`

HasNotAllIncludedSearchEntryCriteria returns a boolean if a field has been set.

### GetNoneIncludedSearchEntryCriteria

`func (o *AddAggregateSearchEntryCriteriaRequest) GetNoneIncludedSearchEntryCriteria() []string`

GetNoneIncludedSearchEntryCriteria returns the NoneIncludedSearchEntryCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedSearchEntryCriteriaOk

`func (o *AddAggregateSearchEntryCriteriaRequest) GetNoneIncludedSearchEntryCriteriaOk() (*[]string, bool)`

GetNoneIncludedSearchEntryCriteriaOk returns a tuple with the NoneIncludedSearchEntryCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedSearchEntryCriteria

`func (o *AddAggregateSearchEntryCriteriaRequest) SetNoneIncludedSearchEntryCriteria(v []string)`

SetNoneIncludedSearchEntryCriteria sets NoneIncludedSearchEntryCriteria field to given value.

### HasNoneIncludedSearchEntryCriteria

`func (o *AddAggregateSearchEntryCriteriaRequest) HasNoneIncludedSearchEntryCriteria() bool`

HasNoneIncludedSearchEntryCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *AddAggregateSearchEntryCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAggregateSearchEntryCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAggregateSearchEntryCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAggregateSearchEntryCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCriteriaName

`func (o *AddAggregateSearchEntryCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddAggregateSearchEntryCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddAggregateSearchEntryCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


