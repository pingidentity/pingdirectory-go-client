# ChangeSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Change Subscription | 
**Schemas** | Pointer to [**[]EnumchangeSubscriptionSchemaUrn**](EnumchangeSubscriptionSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Change Subscription | [optional] 
**ConnectionCriteria** | Pointer to **string** | Specifies a set of connection criteria that must match the client connection associated with an operation in order for that operation to be processed by a change subscription handler. | [optional] 
**RequestCriteria** | Pointer to **string** | Specifies a set of request criteria that must match the request associated with an operation in order for that operation to be processed by a change subscription handler. | [optional] 
**ResultCriteria** | Pointer to **string** | Specifies a set of result criteria that must match the result associated with an operation in order for that operation to be processed by a change subscription handler. | [optional] 
**ExpirationTime** | Pointer to **string** | Specifies a timestamp that provides an expiration time for this change subscription. If an expiration time is provided, then the change subscription will not be active after that time has passed. | [optional] 

## Methods

### NewChangeSubscriptionResponse

`func NewChangeSubscriptionResponse(id string, ) *ChangeSubscriptionResponse`

NewChangeSubscriptionResponse instantiates a new ChangeSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeSubscriptionResponseWithDefaults

`func NewChangeSubscriptionResponseWithDefaults() *ChangeSubscriptionResponse`

NewChangeSubscriptionResponseWithDefaults instantiates a new ChangeSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChangeSubscriptionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangeSubscriptionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangeSubscriptionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ChangeSubscriptionResponse) GetSchemas() []EnumchangeSubscriptionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ChangeSubscriptionResponse) GetSchemasOk() (*[]EnumchangeSubscriptionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ChangeSubscriptionResponse) SetSchemas(v []EnumchangeSubscriptionSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ChangeSubscriptionResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *ChangeSubscriptionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChangeSubscriptionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChangeSubscriptionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChangeSubscriptionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *ChangeSubscriptionResponse) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *ChangeSubscriptionResponse) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *ChangeSubscriptionResponse) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *ChangeSubscriptionResponse) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *ChangeSubscriptionResponse) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *ChangeSubscriptionResponse) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *ChangeSubscriptionResponse) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *ChangeSubscriptionResponse) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetResultCriteria

`func (o *ChangeSubscriptionResponse) GetResultCriteria() string`

GetResultCriteria returns the ResultCriteria field if non-nil, zero value otherwise.

### GetResultCriteriaOk

`func (o *ChangeSubscriptionResponse) GetResultCriteriaOk() (*string, bool)`

GetResultCriteriaOk returns a tuple with the ResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCriteria

`func (o *ChangeSubscriptionResponse) SetResultCriteria(v string)`

SetResultCriteria sets ResultCriteria field to given value.

### HasResultCriteria

`func (o *ChangeSubscriptionResponse) HasResultCriteria() bool`

HasResultCriteria returns a boolean if a field has been set.

### GetExpirationTime

`func (o *ChangeSubscriptionResponse) GetExpirationTime() string`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *ChangeSubscriptionResponse) GetExpirationTimeOk() (*string, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *ChangeSubscriptionResponse) SetExpirationTime(v string)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *ChangeSubscriptionResponse) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


