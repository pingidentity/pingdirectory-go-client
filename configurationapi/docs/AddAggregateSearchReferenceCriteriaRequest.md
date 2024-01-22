# AddAggregateSearchReferenceCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumaggregateSearchReferenceCriteriaSchemaUrn**](EnumaggregateSearchReferenceCriteriaSchemaUrn.md) |  | 
**AllIncludedSearchReferenceCriteria** | Pointer to **[]string** | Specifies a search reference criteria object that must match the associated search result reference in order to match the aggregate search reference criteria. If one or more all-included search reference criteria objects are provided, then a search result reference must match all of them in order to match the aggregate search reference criteria. | [optional] 
**AnyIncludedSearchReferenceCriteria** | Pointer to **[]string** | Specifies a search reference criteria object that may match the associated search result reference in order to match the aggregate search reference criteria. If one or more any-included search reference criteria objects are provided, then a search result reference must match at least one of them in order to match the aggregate search reference criteria. | [optional] 
**NotAllIncludedSearchReferenceCriteria** | Pointer to **[]string** | Specifies a search reference criteria object that should not match the associated search result reference in order to match the aggregate search reference criteria. If one or more not-all-included search reference criteria objects are provided, then a search result reference must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate search reference criteria. | [optional] 
**NoneIncludedSearchReferenceCriteria** | Pointer to **[]string** | Specifies a search reference criteria object that must not match the associated search result reference in order to match the aggregate search reference criteria. If one or more none-included search reference criteria objects are provided, then a search result reference must not match any of them in order to match the aggregate search reference criteria. | [optional] 
**Description** | Pointer to **string** | A description for this Search Reference Criteria | [optional] 
**CriteriaName** | **string** | Name of the new Search Reference Criteria | 

## Methods

### NewAddAggregateSearchReferenceCriteriaRequest

`func NewAddAggregateSearchReferenceCriteriaRequest(schemas []EnumaggregateSearchReferenceCriteriaSchemaUrn, criteriaName string, ) *AddAggregateSearchReferenceCriteriaRequest`

NewAddAggregateSearchReferenceCriteriaRequest instantiates a new AddAggregateSearchReferenceCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAggregateSearchReferenceCriteriaRequestWithDefaults

`func NewAddAggregateSearchReferenceCriteriaRequestWithDefaults() *AddAggregateSearchReferenceCriteriaRequest`

NewAddAggregateSearchReferenceCriteriaRequestWithDefaults instantiates a new AddAggregateSearchReferenceCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddAggregateSearchReferenceCriteriaRequest) GetSchemas() []EnumaggregateSearchReferenceCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAggregateSearchReferenceCriteriaRequest) GetSchemasOk() (*[]EnumaggregateSearchReferenceCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAggregateSearchReferenceCriteriaRequest) SetSchemas(v []EnumaggregateSearchReferenceCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllIncludedSearchReferenceCriteria

`func (o *AddAggregateSearchReferenceCriteriaRequest) GetAllIncludedSearchReferenceCriteria() []string`

GetAllIncludedSearchReferenceCriteria returns the AllIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetAllIncludedSearchReferenceCriteriaOk

`func (o *AddAggregateSearchReferenceCriteriaRequest) GetAllIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetAllIncludedSearchReferenceCriteriaOk returns a tuple with the AllIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedSearchReferenceCriteria

`func (o *AddAggregateSearchReferenceCriteriaRequest) SetAllIncludedSearchReferenceCriteria(v []string)`

SetAllIncludedSearchReferenceCriteria sets AllIncludedSearchReferenceCriteria field to given value.

### HasAllIncludedSearchReferenceCriteria

`func (o *AddAggregateSearchReferenceCriteriaRequest) HasAllIncludedSearchReferenceCriteria() bool`

HasAllIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetAnyIncludedSearchReferenceCriteria

`func (o *AddAggregateSearchReferenceCriteriaRequest) GetAnyIncludedSearchReferenceCriteria() []string`

GetAnyIncludedSearchReferenceCriteria returns the AnyIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedSearchReferenceCriteriaOk

`func (o *AddAggregateSearchReferenceCriteriaRequest) GetAnyIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetAnyIncludedSearchReferenceCriteriaOk returns a tuple with the AnyIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedSearchReferenceCriteria

`func (o *AddAggregateSearchReferenceCriteriaRequest) SetAnyIncludedSearchReferenceCriteria(v []string)`

SetAnyIncludedSearchReferenceCriteria sets AnyIncludedSearchReferenceCriteria field to given value.

### HasAnyIncludedSearchReferenceCriteria

`func (o *AddAggregateSearchReferenceCriteriaRequest) HasAnyIncludedSearchReferenceCriteria() bool`

HasAnyIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetNotAllIncludedSearchReferenceCriteria

`func (o *AddAggregateSearchReferenceCriteriaRequest) GetNotAllIncludedSearchReferenceCriteria() []string`

GetNotAllIncludedSearchReferenceCriteria returns the NotAllIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedSearchReferenceCriteriaOk

`func (o *AddAggregateSearchReferenceCriteriaRequest) GetNotAllIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetNotAllIncludedSearchReferenceCriteriaOk returns a tuple with the NotAllIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedSearchReferenceCriteria

`func (o *AddAggregateSearchReferenceCriteriaRequest) SetNotAllIncludedSearchReferenceCriteria(v []string)`

SetNotAllIncludedSearchReferenceCriteria sets NotAllIncludedSearchReferenceCriteria field to given value.

### HasNotAllIncludedSearchReferenceCriteria

`func (o *AddAggregateSearchReferenceCriteriaRequest) HasNotAllIncludedSearchReferenceCriteria() bool`

HasNotAllIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetNoneIncludedSearchReferenceCriteria

`func (o *AddAggregateSearchReferenceCriteriaRequest) GetNoneIncludedSearchReferenceCriteria() []string`

GetNoneIncludedSearchReferenceCriteria returns the NoneIncludedSearchReferenceCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedSearchReferenceCriteriaOk

`func (o *AddAggregateSearchReferenceCriteriaRequest) GetNoneIncludedSearchReferenceCriteriaOk() (*[]string, bool)`

GetNoneIncludedSearchReferenceCriteriaOk returns a tuple with the NoneIncludedSearchReferenceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedSearchReferenceCriteria

`func (o *AddAggregateSearchReferenceCriteriaRequest) SetNoneIncludedSearchReferenceCriteria(v []string)`

SetNoneIncludedSearchReferenceCriteria sets NoneIncludedSearchReferenceCriteria field to given value.

### HasNoneIncludedSearchReferenceCriteria

`func (o *AddAggregateSearchReferenceCriteriaRequest) HasNoneIncludedSearchReferenceCriteria() bool`

HasNoneIncludedSearchReferenceCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *AddAggregateSearchReferenceCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAggregateSearchReferenceCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAggregateSearchReferenceCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAggregateSearchReferenceCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCriteriaName

`func (o *AddAggregateSearchReferenceCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddAggregateSearchReferenceCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddAggregateSearchReferenceCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


