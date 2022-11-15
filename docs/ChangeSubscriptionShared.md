# ChangeSubscriptionShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumchangeSubscriptionSchemaUrn**](EnumchangeSubscriptionSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Change Subscription | [optional] 
**ConnectionCriteria** | Pointer to **string** | Specifies a set of connection criteria that must match the client connection associated with an operation in order for that operation to be processed by a change subscription handler. | [optional] 
**RequestCriteria** | Pointer to **string** | Specifies a set of request criteria that must match the request associated with an operation in order for that operation to be processed by a change subscription handler. | [optional] 
**ResultCriteria** | Pointer to **string** | Specifies a set of result criteria that must match the result associated with an operation in order for that operation to be processed by a change subscription handler. | [optional] 
**ExpirationTime** | Pointer to **string** | Specifies a timestamp that provides an expiration time for this change subscription. If an expiration time is provided, then the change subscription will not be active after that time has passed. | [optional] 

## Methods

### NewChangeSubscriptionShared

`func NewChangeSubscriptionShared() *ChangeSubscriptionShared`

NewChangeSubscriptionShared instantiates a new ChangeSubscriptionShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeSubscriptionSharedWithDefaults

`func NewChangeSubscriptionSharedWithDefaults() *ChangeSubscriptionShared`

NewChangeSubscriptionSharedWithDefaults instantiates a new ChangeSubscriptionShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ChangeSubscriptionShared) GetSchemas() []EnumchangeSubscriptionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ChangeSubscriptionShared) GetSchemasOk() (*[]EnumchangeSubscriptionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ChangeSubscriptionShared) SetSchemas(v []EnumchangeSubscriptionSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ChangeSubscriptionShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *ChangeSubscriptionShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChangeSubscriptionShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChangeSubscriptionShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChangeSubscriptionShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *ChangeSubscriptionShared) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *ChangeSubscriptionShared) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *ChangeSubscriptionShared) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *ChangeSubscriptionShared) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *ChangeSubscriptionShared) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *ChangeSubscriptionShared) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *ChangeSubscriptionShared) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *ChangeSubscriptionShared) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetResultCriteria

`func (o *ChangeSubscriptionShared) GetResultCriteria() string`

GetResultCriteria returns the ResultCriteria field if non-nil, zero value otherwise.

### GetResultCriteriaOk

`func (o *ChangeSubscriptionShared) GetResultCriteriaOk() (*string, bool)`

GetResultCriteriaOk returns a tuple with the ResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCriteria

`func (o *ChangeSubscriptionShared) SetResultCriteria(v string)`

SetResultCriteria sets ResultCriteria field to given value.

### HasResultCriteria

`func (o *ChangeSubscriptionShared) HasResultCriteria() bool`

HasResultCriteria returns a boolean if a field has been set.

### GetExpirationTime

`func (o *ChangeSubscriptionShared) GetExpirationTime() string`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *ChangeSubscriptionShared) GetExpirationTimeOk() (*string, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *ChangeSubscriptionShared) SetExpirationTime(v string)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *ChangeSubscriptionShared) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


