# AggregateRequestCriteriaShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumaggregateRequestCriteriaSchemaUrn**](EnumaggregateRequestCriteriaSchemaUrn.md) |  | 
**AllIncludedRequestCriteria** | Pointer to **[]string** |  | [optional] 
**AnyIncludedRequestCriteria** | Pointer to **[]string** |  | [optional] 
**NotAllIncludedRequestCriteria** | Pointer to **[]string** |  | [optional] 
**NoneIncludedRequestCriteria** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Request Criteria | [optional] 

## Methods

### NewAggregateRequestCriteriaShared

`func NewAggregateRequestCriteriaShared(schemas []EnumaggregateRequestCriteriaSchemaUrn, ) *AggregateRequestCriteriaShared`

NewAggregateRequestCriteriaShared instantiates a new AggregateRequestCriteriaShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateRequestCriteriaSharedWithDefaults

`func NewAggregateRequestCriteriaSharedWithDefaults() *AggregateRequestCriteriaShared`

NewAggregateRequestCriteriaSharedWithDefaults instantiates a new AggregateRequestCriteriaShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AggregateRequestCriteriaShared) GetSchemas() []EnumaggregateRequestCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AggregateRequestCriteriaShared) GetSchemasOk() (*[]EnumaggregateRequestCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AggregateRequestCriteriaShared) SetSchemas(v []EnumaggregateRequestCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllIncludedRequestCriteria

`func (o *AggregateRequestCriteriaShared) GetAllIncludedRequestCriteria() []string`

GetAllIncludedRequestCriteria returns the AllIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetAllIncludedRequestCriteriaOk

`func (o *AggregateRequestCriteriaShared) GetAllIncludedRequestCriteriaOk() (*[]string, bool)`

GetAllIncludedRequestCriteriaOk returns a tuple with the AllIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedRequestCriteria

`func (o *AggregateRequestCriteriaShared) SetAllIncludedRequestCriteria(v []string)`

SetAllIncludedRequestCriteria sets AllIncludedRequestCriteria field to given value.

### HasAllIncludedRequestCriteria

`func (o *AggregateRequestCriteriaShared) HasAllIncludedRequestCriteria() bool`

HasAllIncludedRequestCriteria returns a boolean if a field has been set.

### GetAnyIncludedRequestCriteria

`func (o *AggregateRequestCriteriaShared) GetAnyIncludedRequestCriteria() []string`

GetAnyIncludedRequestCriteria returns the AnyIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedRequestCriteriaOk

`func (o *AggregateRequestCriteriaShared) GetAnyIncludedRequestCriteriaOk() (*[]string, bool)`

GetAnyIncludedRequestCriteriaOk returns a tuple with the AnyIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedRequestCriteria

`func (o *AggregateRequestCriteriaShared) SetAnyIncludedRequestCriteria(v []string)`

SetAnyIncludedRequestCriteria sets AnyIncludedRequestCriteria field to given value.

### HasAnyIncludedRequestCriteria

`func (o *AggregateRequestCriteriaShared) HasAnyIncludedRequestCriteria() bool`

HasAnyIncludedRequestCriteria returns a boolean if a field has been set.

### GetNotAllIncludedRequestCriteria

`func (o *AggregateRequestCriteriaShared) GetNotAllIncludedRequestCriteria() []string`

GetNotAllIncludedRequestCriteria returns the NotAllIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedRequestCriteriaOk

`func (o *AggregateRequestCriteriaShared) GetNotAllIncludedRequestCriteriaOk() (*[]string, bool)`

GetNotAllIncludedRequestCriteriaOk returns a tuple with the NotAllIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedRequestCriteria

`func (o *AggregateRequestCriteriaShared) SetNotAllIncludedRequestCriteria(v []string)`

SetNotAllIncludedRequestCriteria sets NotAllIncludedRequestCriteria field to given value.

### HasNotAllIncludedRequestCriteria

`func (o *AggregateRequestCriteriaShared) HasNotAllIncludedRequestCriteria() bool`

HasNotAllIncludedRequestCriteria returns a boolean if a field has been set.

### GetNoneIncludedRequestCriteria

`func (o *AggregateRequestCriteriaShared) GetNoneIncludedRequestCriteria() []string`

GetNoneIncludedRequestCriteria returns the NoneIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedRequestCriteriaOk

`func (o *AggregateRequestCriteriaShared) GetNoneIncludedRequestCriteriaOk() (*[]string, bool)`

GetNoneIncludedRequestCriteriaOk returns a tuple with the NoneIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedRequestCriteria

`func (o *AggregateRequestCriteriaShared) SetNoneIncludedRequestCriteria(v []string)`

SetNoneIncludedRequestCriteria sets NoneIncludedRequestCriteria field to given value.

### HasNoneIncludedRequestCriteria

`func (o *AggregateRequestCriteriaShared) HasNoneIncludedRequestCriteria() bool`

HasNoneIncludedRequestCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *AggregateRequestCriteriaShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AggregateRequestCriteriaShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AggregateRequestCriteriaShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AggregateRequestCriteriaShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


