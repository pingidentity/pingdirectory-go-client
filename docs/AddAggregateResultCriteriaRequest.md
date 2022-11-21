# AddAggregateResultCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriteriaName** | **string** | Name of the new Result Criteria | 
**Schemas** | [**[]EnumaggregateResultCriteriaSchemaUrn**](EnumaggregateResultCriteriaSchemaUrn.md) |  | 
**AllIncludedResultCriteria** | Pointer to **[]string** | Specifies a result criteria object that must match the associated operation result in order to match the aggregate result criteria. If one or more all-included result criteria objects are provided, then an operation result must match all of them in order to match the aggregate result criteria. | [optional] 
**AnyIncludedResultCriteria** | Pointer to **[]string** | Specifies a result criteria object that may match the associated operation result in order to match the aggregate result criteria. If one or more any-included result criteria objects are provided, then an operation result must match at least one of them in order to match the aggregate result criteria. | [optional] 
**NotAllIncludedResultCriteria** | Pointer to **[]string** | Specifies a result criteria object that should not match the associated operation result in order to match the aggregate result criteria. If one or more not-all-included result criteria objects are provided, then an operation result must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate result criteria. | [optional] 
**NoneIncludedResultCriteria** | Pointer to **[]string** | Specifies a result criteria object that must not match the associated operation result in order to match the aggregate result criteria. If one or more none-included result criteria objects are provided, then an operation result must not match any of them in order to match the aggregate result criteria. | [optional] 
**Description** | Pointer to **string** | A description for this Result Criteria | [optional] 

## Methods

### NewAddAggregateResultCriteriaRequest

`func NewAddAggregateResultCriteriaRequest(criteriaName string, schemas []EnumaggregateResultCriteriaSchemaUrn, ) *AddAggregateResultCriteriaRequest`

NewAddAggregateResultCriteriaRequest instantiates a new AddAggregateResultCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAggregateResultCriteriaRequestWithDefaults

`func NewAddAggregateResultCriteriaRequestWithDefaults() *AddAggregateResultCriteriaRequest`

NewAddAggregateResultCriteriaRequestWithDefaults instantiates a new AddAggregateResultCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteriaName

`func (o *AddAggregateResultCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddAggregateResultCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddAggregateResultCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.


### GetSchemas

`func (o *AddAggregateResultCriteriaRequest) GetSchemas() []EnumaggregateResultCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAggregateResultCriteriaRequest) GetSchemasOk() (*[]EnumaggregateResultCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAggregateResultCriteriaRequest) SetSchemas(v []EnumaggregateResultCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllIncludedResultCriteria

`func (o *AddAggregateResultCriteriaRequest) GetAllIncludedResultCriteria() []string`

GetAllIncludedResultCriteria returns the AllIncludedResultCriteria field if non-nil, zero value otherwise.

### GetAllIncludedResultCriteriaOk

`func (o *AddAggregateResultCriteriaRequest) GetAllIncludedResultCriteriaOk() (*[]string, bool)`

GetAllIncludedResultCriteriaOk returns a tuple with the AllIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedResultCriteria

`func (o *AddAggregateResultCriteriaRequest) SetAllIncludedResultCriteria(v []string)`

SetAllIncludedResultCriteria sets AllIncludedResultCriteria field to given value.

### HasAllIncludedResultCriteria

`func (o *AddAggregateResultCriteriaRequest) HasAllIncludedResultCriteria() bool`

HasAllIncludedResultCriteria returns a boolean if a field has been set.

### GetAnyIncludedResultCriteria

`func (o *AddAggregateResultCriteriaRequest) GetAnyIncludedResultCriteria() []string`

GetAnyIncludedResultCriteria returns the AnyIncludedResultCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedResultCriteriaOk

`func (o *AddAggregateResultCriteriaRequest) GetAnyIncludedResultCriteriaOk() (*[]string, bool)`

GetAnyIncludedResultCriteriaOk returns a tuple with the AnyIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedResultCriteria

`func (o *AddAggregateResultCriteriaRequest) SetAnyIncludedResultCriteria(v []string)`

SetAnyIncludedResultCriteria sets AnyIncludedResultCriteria field to given value.

### HasAnyIncludedResultCriteria

`func (o *AddAggregateResultCriteriaRequest) HasAnyIncludedResultCriteria() bool`

HasAnyIncludedResultCriteria returns a boolean if a field has been set.

### GetNotAllIncludedResultCriteria

`func (o *AddAggregateResultCriteriaRequest) GetNotAllIncludedResultCriteria() []string`

GetNotAllIncludedResultCriteria returns the NotAllIncludedResultCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedResultCriteriaOk

`func (o *AddAggregateResultCriteriaRequest) GetNotAllIncludedResultCriteriaOk() (*[]string, bool)`

GetNotAllIncludedResultCriteriaOk returns a tuple with the NotAllIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedResultCriteria

`func (o *AddAggregateResultCriteriaRequest) SetNotAllIncludedResultCriteria(v []string)`

SetNotAllIncludedResultCriteria sets NotAllIncludedResultCriteria field to given value.

### HasNotAllIncludedResultCriteria

`func (o *AddAggregateResultCriteriaRequest) HasNotAllIncludedResultCriteria() bool`

HasNotAllIncludedResultCriteria returns a boolean if a field has been set.

### GetNoneIncludedResultCriteria

`func (o *AddAggregateResultCriteriaRequest) GetNoneIncludedResultCriteria() []string`

GetNoneIncludedResultCriteria returns the NoneIncludedResultCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedResultCriteriaOk

`func (o *AddAggregateResultCriteriaRequest) GetNoneIncludedResultCriteriaOk() (*[]string, bool)`

GetNoneIncludedResultCriteriaOk returns a tuple with the NoneIncludedResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedResultCriteria

`func (o *AddAggregateResultCriteriaRequest) SetNoneIncludedResultCriteria(v []string)`

SetNoneIncludedResultCriteria sets NoneIncludedResultCriteria field to given value.

### HasNoneIncludedResultCriteria

`func (o *AddAggregateResultCriteriaRequest) HasNoneIncludedResultCriteria() bool`

HasNoneIncludedResultCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *AddAggregateResultCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAggregateResultCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAggregateResultCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAggregateResultCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


