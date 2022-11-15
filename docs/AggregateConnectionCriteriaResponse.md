# AggregateConnectionCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Connection Criteria | 
**Schemas** | [**[]EnumaggregateConnectionCriteriaSchemaUrn**](EnumaggregateConnectionCriteriaSchemaUrn.md) |  | 
**AllIncludedConnectionCriteria** | Pointer to **[]string** |  | [optional] 
**AnyIncludedConnectionCriteria** | Pointer to **[]string** |  | [optional] 
**NotAllIncludedConnectionCriteria** | Pointer to **[]string** |  | [optional] 
**NoneIncludedConnectionCriteria** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Connection Criteria | [optional] 

## Methods

### NewAggregateConnectionCriteriaResponse

`func NewAggregateConnectionCriteriaResponse(id string, schemas []EnumaggregateConnectionCriteriaSchemaUrn, ) *AggregateConnectionCriteriaResponse`

NewAggregateConnectionCriteriaResponse instantiates a new AggregateConnectionCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateConnectionCriteriaResponseWithDefaults

`func NewAggregateConnectionCriteriaResponseWithDefaults() *AggregateConnectionCriteriaResponse`

NewAggregateConnectionCriteriaResponseWithDefaults instantiates a new AggregateConnectionCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AggregateConnectionCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AggregateConnectionCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AggregateConnectionCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AggregateConnectionCriteriaResponse) GetSchemas() []EnumaggregateConnectionCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AggregateConnectionCriteriaResponse) GetSchemasOk() (*[]EnumaggregateConnectionCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AggregateConnectionCriteriaResponse) SetSchemas(v []EnumaggregateConnectionCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaResponse) GetAllIncludedConnectionCriteria() []string`

GetAllIncludedConnectionCriteria returns the AllIncludedConnectionCriteria field if non-nil, zero value otherwise.

### GetAllIncludedConnectionCriteriaOk

`func (o *AggregateConnectionCriteriaResponse) GetAllIncludedConnectionCriteriaOk() (*[]string, bool)`

GetAllIncludedConnectionCriteriaOk returns a tuple with the AllIncludedConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaResponse) SetAllIncludedConnectionCriteria(v []string)`

SetAllIncludedConnectionCriteria sets AllIncludedConnectionCriteria field to given value.

### HasAllIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaResponse) HasAllIncludedConnectionCriteria() bool`

HasAllIncludedConnectionCriteria returns a boolean if a field has been set.

### GetAnyIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaResponse) GetAnyIncludedConnectionCriteria() []string`

GetAnyIncludedConnectionCriteria returns the AnyIncludedConnectionCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedConnectionCriteriaOk

`func (o *AggregateConnectionCriteriaResponse) GetAnyIncludedConnectionCriteriaOk() (*[]string, bool)`

GetAnyIncludedConnectionCriteriaOk returns a tuple with the AnyIncludedConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaResponse) SetAnyIncludedConnectionCriteria(v []string)`

SetAnyIncludedConnectionCriteria sets AnyIncludedConnectionCriteria field to given value.

### HasAnyIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaResponse) HasAnyIncludedConnectionCriteria() bool`

HasAnyIncludedConnectionCriteria returns a boolean if a field has been set.

### GetNotAllIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaResponse) GetNotAllIncludedConnectionCriteria() []string`

GetNotAllIncludedConnectionCriteria returns the NotAllIncludedConnectionCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedConnectionCriteriaOk

`func (o *AggregateConnectionCriteriaResponse) GetNotAllIncludedConnectionCriteriaOk() (*[]string, bool)`

GetNotAllIncludedConnectionCriteriaOk returns a tuple with the NotAllIncludedConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaResponse) SetNotAllIncludedConnectionCriteria(v []string)`

SetNotAllIncludedConnectionCriteria sets NotAllIncludedConnectionCriteria field to given value.

### HasNotAllIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaResponse) HasNotAllIncludedConnectionCriteria() bool`

HasNotAllIncludedConnectionCriteria returns a boolean if a field has been set.

### GetNoneIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaResponse) GetNoneIncludedConnectionCriteria() []string`

GetNoneIncludedConnectionCriteria returns the NoneIncludedConnectionCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedConnectionCriteriaOk

`func (o *AggregateConnectionCriteriaResponse) GetNoneIncludedConnectionCriteriaOk() (*[]string, bool)`

GetNoneIncludedConnectionCriteriaOk returns a tuple with the NoneIncludedConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaResponse) SetNoneIncludedConnectionCriteria(v []string)`

SetNoneIncludedConnectionCriteria sets NoneIncludedConnectionCriteria field to given value.

### HasNoneIncludedConnectionCriteria

`func (o *AggregateConnectionCriteriaResponse) HasNoneIncludedConnectionCriteria() bool`

HasNoneIncludedConnectionCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *AggregateConnectionCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AggregateConnectionCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AggregateConnectionCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AggregateConnectionCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


