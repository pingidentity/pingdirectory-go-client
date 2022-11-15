# AggregateConnectionCriteriaShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumaggregateConnectionCriteriaSchemaUrn**](EnumaggregateConnectionCriteriaSchemaUrn.md) |  | 
**AllIncludedConnectionCriteria** | Pointer to **[]string** |  | [optional] 
**AnyIncludedConnectionCriteria** | Pointer to **[]string** |  | [optional] 
**NotAllIncludedConnectionCriteria** | Pointer to **[]string** |  | [optional] 
**NoneIncludedConnectionCriteria** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Connection Criteria | [optional] 

## Methods

### NewAggregateConnectionCriteriaShared

`func NewAggregateConnectionCriteriaShared(schemas []EnumaggregateConnectionCriteriaSchemaUrn, ) *AggregateConnectionCriteriaShared`

NewAggregateConnectionCriteriaShared instantiates a new AggregateConnectionCriteriaShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateConnectionCriteriaSharedWithDefaults

`func NewAggregateConnectionCriteriaSharedWithDefaults() *AggregateConnectionCriteriaShared`

NewAggregateConnectionCriteriaSharedWithDefaults instantiates a new AggregateConnectionCriteriaShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AggregateConnectionCriteriaShared) GetSchemas() []EnumaggregateConnectionCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AggregateConnectionCriteriaShared) GetSchemasOk() (*[]EnumaggregateConnectionCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AggregateConnectionCriteriaShared) SetSchemas(v []EnumaggregateConnectionCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaShared) GetAllIncludedConnectionCriteria() []string`

GetAllIncludedConnectionCriteria returns the AllIncludedConnectionCriteria field if non-nil, zero value otherwise.

### GetAllIncludedConnectionCriteriaOk

`func (o *AggregateConnectionCriteriaShared) GetAllIncludedConnectionCriteriaOk() (*[]string, bool)`

GetAllIncludedConnectionCriteriaOk returns a tuple with the AllIncludedConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaShared) SetAllIncludedConnectionCriteria(v []string)`

SetAllIncludedConnectionCriteria sets AllIncludedConnectionCriteria field to given value.

### HasAllIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaShared) HasAllIncludedConnectionCriteria() bool`

HasAllIncludedConnectionCriteria returns a boolean if a field has been set.

### GetAnyIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaShared) GetAnyIncludedConnectionCriteria() []string`

GetAnyIncludedConnectionCriteria returns the AnyIncludedConnectionCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedConnectionCriteriaOk

`func (o *AggregateConnectionCriteriaShared) GetAnyIncludedConnectionCriteriaOk() (*[]string, bool)`

GetAnyIncludedConnectionCriteriaOk returns a tuple with the AnyIncludedConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaShared) SetAnyIncludedConnectionCriteria(v []string)`

SetAnyIncludedConnectionCriteria sets AnyIncludedConnectionCriteria field to given value.

### HasAnyIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaShared) HasAnyIncludedConnectionCriteria() bool`

HasAnyIncludedConnectionCriteria returns a boolean if a field has been set.

### GetNotAllIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaShared) GetNotAllIncludedConnectionCriteria() []string`

GetNotAllIncludedConnectionCriteria returns the NotAllIncludedConnectionCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedConnectionCriteriaOk

`func (o *AggregateConnectionCriteriaShared) GetNotAllIncludedConnectionCriteriaOk() (*[]string, bool)`

GetNotAllIncludedConnectionCriteriaOk returns a tuple with the NotAllIncludedConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaShared) SetNotAllIncludedConnectionCriteria(v []string)`

SetNotAllIncludedConnectionCriteria sets NotAllIncludedConnectionCriteria field to given value.

### HasNotAllIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaShared) HasNotAllIncludedConnectionCriteria() bool`

HasNotAllIncludedConnectionCriteria returns a boolean if a field has been set.

### GetNoneIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaShared) GetNoneIncludedConnectionCriteria() []string`

GetNoneIncludedConnectionCriteria returns the NoneIncludedConnectionCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedConnectionCriteriaOk

`func (o *AggregateConnectionCriteriaShared) GetNoneIncludedConnectionCriteriaOk() (*[]string, bool)`

GetNoneIncludedConnectionCriteriaOk returns a tuple with the NoneIncludedConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaShared) SetNoneIncludedConnectionCriteria(v []string)`

SetNoneIncludedConnectionCriteria sets NoneIncludedConnectionCriteria field to given value.

### HasNoneIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaShared) HasNoneIncludedConnectionCriteria() bool`

HasNoneIncludedConnectionCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *AggregateConnectionCriteriaShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AggregateConnectionCriteriaShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AggregateConnectionCriteriaShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AggregateConnectionCriteriaShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


