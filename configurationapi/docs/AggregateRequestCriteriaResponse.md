# AggregateRequestCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Request Criteria | 
**Schemas** | [**[]EnumaggregateRequestCriteriaSchemaUrn**](EnumaggregateRequestCriteriaSchemaUrn.md) |  | 
**AllIncludedRequestCriteria** | Pointer to **[]string** | Specifies a request criteria object that must match the associated operation request in order to match the aggregate request criteria. If one or more all-included request criteria objects are provided, then an operation request must match all of them in order to match the aggregate request criteria. | [optional] 
**AnyIncludedRequestCriteria** | Pointer to **[]string** | Specifies a request criteria object that may match the associated operation request in order to the this aggregate request criteria. If one or more any-included request criteria objects are provided, then an operation request must match at least one of them in order to match the aggregate request criteria. | [optional] 
**NotAllIncludedRequestCriteria** | Pointer to **[]string** | Specifies a request criteria object that should not match the associated operation request in order to match the aggregate request criteria. If one or more not-all-included request criteria objects are provided, then an operation request must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate request criteria. | [optional] 
**NoneIncludedRequestCriteria** | Pointer to **[]string** | Specifies a request criteria object that must not match the associated operation request in order to match the aggregate request criteria. If one or more none-included request criteria objects are provided, then an operation request must not match any of them in order to match the aggregate request criteria. | [optional] 
**Description** | Pointer to **string** | A description for this Request Criteria | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewAggregateRequestCriteriaResponse

`func NewAggregateRequestCriteriaResponse(id string, schemas []EnumaggregateRequestCriteriaSchemaUrn, ) *AggregateRequestCriteriaResponse`

NewAggregateRequestCriteriaResponse instantiates a new AggregateRequestCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateRequestCriteriaResponseWithDefaults

`func NewAggregateRequestCriteriaResponseWithDefaults() *AggregateRequestCriteriaResponse`

NewAggregateRequestCriteriaResponseWithDefaults instantiates a new AggregateRequestCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AggregateRequestCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AggregateRequestCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AggregateRequestCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AggregateRequestCriteriaResponse) GetSchemas() []EnumaggregateRequestCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AggregateRequestCriteriaResponse) GetSchemasOk() (*[]EnumaggregateRequestCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AggregateRequestCriteriaResponse) SetSchemas(v []EnumaggregateRequestCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllIncludedRequestCriteria

`func (o *AggregateRequestCriteriaResponse) GetAllIncludedRequestCriteria() []string`

GetAllIncludedRequestCriteria returns the AllIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetAllIncludedRequestCriteriaOk

`func (o *AggregateRequestCriteriaResponse) GetAllIncludedRequestCriteriaOk() (*[]string, bool)`

GetAllIncludedRequestCriteriaOk returns a tuple with the AllIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedRequestCriteria

`func (o *AggregateRequestCriteriaResponse) SetAllIncludedRequestCriteria(v []string)`

SetAllIncludedRequestCriteria sets AllIncludedRequestCriteria field to given value.

### HasAllIncludedRequestCriteria

`func (o *AggregateRequestCriteriaResponse) HasAllIncludedRequestCriteria() bool`

HasAllIncludedRequestCriteria returns a boolean if a field has been set.

### GetAnyIncludedRequestCriteria

`func (o *AggregateRequestCriteriaResponse) GetAnyIncludedRequestCriteria() []string`

GetAnyIncludedRequestCriteria returns the AnyIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedRequestCriteriaOk

`func (o *AggregateRequestCriteriaResponse) GetAnyIncludedRequestCriteriaOk() (*[]string, bool)`

GetAnyIncludedRequestCriteriaOk returns a tuple with the AnyIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedRequestCriteria

`func (o *AggregateRequestCriteriaResponse) SetAnyIncludedRequestCriteria(v []string)`

SetAnyIncludedRequestCriteria sets AnyIncludedRequestCriteria field to given value.

### HasAnyIncludedRequestCriteria

`func (o *AggregateRequestCriteriaResponse) HasAnyIncludedRequestCriteria() bool`

HasAnyIncludedRequestCriteria returns a boolean if a field has been set.

### GetNotAllIncludedRequestCriteria

`func (o *AggregateRequestCriteriaResponse) GetNotAllIncludedRequestCriteria() []string`

GetNotAllIncludedRequestCriteria returns the NotAllIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedRequestCriteriaOk

`func (o *AggregateRequestCriteriaResponse) GetNotAllIncludedRequestCriteriaOk() (*[]string, bool)`

GetNotAllIncludedRequestCriteriaOk returns a tuple with the NotAllIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedRequestCriteria

`func (o *AggregateRequestCriteriaResponse) SetNotAllIncludedRequestCriteria(v []string)`

SetNotAllIncludedRequestCriteria sets NotAllIncludedRequestCriteria field to given value.

### HasNotAllIncludedRequestCriteria

`func (o *AggregateRequestCriteriaResponse) HasNotAllIncludedRequestCriteria() bool`

HasNotAllIncludedRequestCriteria returns a boolean if a field has been set.

### GetNoneIncludedRequestCriteria

`func (o *AggregateRequestCriteriaResponse) GetNoneIncludedRequestCriteria() []string`

GetNoneIncludedRequestCriteria returns the NoneIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedRequestCriteriaOk

`func (o *AggregateRequestCriteriaResponse) GetNoneIncludedRequestCriteriaOk() (*[]string, bool)`

GetNoneIncludedRequestCriteriaOk returns a tuple with the NoneIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedRequestCriteria

`func (o *AggregateRequestCriteriaResponse) SetNoneIncludedRequestCriteria(v []string)`

SetNoneIncludedRequestCriteria sets NoneIncludedRequestCriteria field to given value.

### HasNoneIncludedRequestCriteria

`func (o *AggregateRequestCriteriaResponse) HasNoneIncludedRequestCriteria() bool`

HasNoneIncludedRequestCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *AggregateRequestCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AggregateRequestCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AggregateRequestCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AggregateRequestCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *AggregateRequestCriteriaResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AggregateRequestCriteriaResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AggregateRequestCriteriaResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AggregateRequestCriteriaResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AggregateRequestCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AggregateRequestCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AggregateRequestCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AggregateRequestCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


