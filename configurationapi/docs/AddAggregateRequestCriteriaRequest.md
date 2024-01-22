# AddAggregateRequestCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumaggregateRequestCriteriaSchemaUrn**](EnumaggregateRequestCriteriaSchemaUrn.md) |  | 
**AllIncludedRequestCriteria** | Pointer to **[]string** | Specifies a request criteria object that must match the associated operation request in order to match the aggregate request criteria. If one or more all-included request criteria objects are provided, then an operation request must match all of them in order to match the aggregate request criteria. | [optional] 
**AnyIncludedRequestCriteria** | Pointer to **[]string** | Specifies a request criteria object that may match the associated operation request in order to the this aggregate request criteria. If one or more any-included request criteria objects are provided, then an operation request must match at least one of them in order to match the aggregate request criteria. | [optional] 
**NotAllIncludedRequestCriteria** | Pointer to **[]string** | Specifies a request criteria object that should not match the associated operation request in order to match the aggregate request criteria. If one or more not-all-included request criteria objects are provided, then an operation request must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate request criteria. | [optional] 
**NoneIncludedRequestCriteria** | Pointer to **[]string** | Specifies a request criteria object that must not match the associated operation request in order to match the aggregate request criteria. If one or more none-included request criteria objects are provided, then an operation request must not match any of them in order to match the aggregate request criteria. | [optional] 
**Description** | Pointer to **string** | A description for this Request Criteria | [optional] 
**CriteriaName** | **string** | Name of the new Request Criteria | 

## Methods

### NewAddAggregateRequestCriteriaRequest

`func NewAddAggregateRequestCriteriaRequest(schemas []EnumaggregateRequestCriteriaSchemaUrn, criteriaName string, ) *AddAggregateRequestCriteriaRequest`

NewAddAggregateRequestCriteriaRequest instantiates a new AddAggregateRequestCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAggregateRequestCriteriaRequestWithDefaults

`func NewAddAggregateRequestCriteriaRequestWithDefaults() *AddAggregateRequestCriteriaRequest`

NewAddAggregateRequestCriteriaRequestWithDefaults instantiates a new AddAggregateRequestCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddAggregateRequestCriteriaRequest) GetSchemas() []EnumaggregateRequestCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAggregateRequestCriteriaRequest) GetSchemasOk() (*[]EnumaggregateRequestCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAggregateRequestCriteriaRequest) SetSchemas(v []EnumaggregateRequestCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllIncludedRequestCriteria

`func (o *AddAggregateRequestCriteriaRequest) GetAllIncludedRequestCriteria() []string`

GetAllIncludedRequestCriteria returns the AllIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetAllIncludedRequestCriteriaOk

`func (o *AddAggregateRequestCriteriaRequest) GetAllIncludedRequestCriteriaOk() (*[]string, bool)`

GetAllIncludedRequestCriteriaOk returns a tuple with the AllIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedRequestCriteria

`func (o *AddAggregateRequestCriteriaRequest) SetAllIncludedRequestCriteria(v []string)`

SetAllIncludedRequestCriteria sets AllIncludedRequestCriteria field to given value.

### HasAllIncludedRequestCriteria

`func (o *AddAggregateRequestCriteriaRequest) HasAllIncludedRequestCriteria() bool`

HasAllIncludedRequestCriteria returns a boolean if a field has been set.

### GetAnyIncludedRequestCriteria

`func (o *AddAggregateRequestCriteriaRequest) GetAnyIncludedRequestCriteria() []string`

GetAnyIncludedRequestCriteria returns the AnyIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedRequestCriteriaOk

`func (o *AddAggregateRequestCriteriaRequest) GetAnyIncludedRequestCriteriaOk() (*[]string, bool)`

GetAnyIncludedRequestCriteriaOk returns a tuple with the AnyIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedRequestCriteria

`func (o *AddAggregateRequestCriteriaRequest) SetAnyIncludedRequestCriteria(v []string)`

SetAnyIncludedRequestCriteria sets AnyIncludedRequestCriteria field to given value.

### HasAnyIncludedRequestCriteria

`func (o *AddAggregateRequestCriteriaRequest) HasAnyIncludedRequestCriteria() bool`

HasAnyIncludedRequestCriteria returns a boolean if a field has been set.

### GetNotAllIncludedRequestCriteria

`func (o *AddAggregateRequestCriteriaRequest) GetNotAllIncludedRequestCriteria() []string`

GetNotAllIncludedRequestCriteria returns the NotAllIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedRequestCriteriaOk

`func (o *AddAggregateRequestCriteriaRequest) GetNotAllIncludedRequestCriteriaOk() (*[]string, bool)`

GetNotAllIncludedRequestCriteriaOk returns a tuple with the NotAllIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedRequestCriteria

`func (o *AddAggregateRequestCriteriaRequest) SetNotAllIncludedRequestCriteria(v []string)`

SetNotAllIncludedRequestCriteria sets NotAllIncludedRequestCriteria field to given value.

### HasNotAllIncludedRequestCriteria

`func (o *AddAggregateRequestCriteriaRequest) HasNotAllIncludedRequestCriteria() bool`

HasNotAllIncludedRequestCriteria returns a boolean if a field has been set.

### GetNoneIncludedRequestCriteria

`func (o *AddAggregateRequestCriteriaRequest) GetNoneIncludedRequestCriteria() []string`

GetNoneIncludedRequestCriteria returns the NoneIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedRequestCriteriaOk

`func (o *AddAggregateRequestCriteriaRequest) GetNoneIncludedRequestCriteriaOk() (*[]string, bool)`

GetNoneIncludedRequestCriteriaOk returns a tuple with the NoneIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedRequestCriteria

`func (o *AddAggregateRequestCriteriaRequest) SetNoneIncludedRequestCriteria(v []string)`

SetNoneIncludedRequestCriteria sets NoneIncludedRequestCriteria field to given value.

### HasNoneIncludedRequestCriteria

`func (o *AddAggregateRequestCriteriaRequest) HasNoneIncludedRequestCriteria() bool`

HasNoneIncludedRequestCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *AddAggregateRequestCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAggregateRequestCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAggregateRequestCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAggregateRequestCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCriteriaName

`func (o *AddAggregateRequestCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddAggregateRequestCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddAggregateRequestCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


