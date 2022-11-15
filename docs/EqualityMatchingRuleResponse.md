# EqualityMatchingRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumequalityMatchingRuleSchemaUrn**](EnumequalityMatchingRuleSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicates whether the Matching Rule is enabled for use. | 

## Methods

### NewEqualityMatchingRuleResponse

`func NewEqualityMatchingRuleResponse(schemas []EnumequalityMatchingRuleSchemaUrn, enabled bool, ) *EqualityMatchingRuleResponse`

NewEqualityMatchingRuleResponse instantiates a new EqualityMatchingRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEqualityMatchingRuleResponseWithDefaults

`func NewEqualityMatchingRuleResponseWithDefaults() *EqualityMatchingRuleResponse`

NewEqualityMatchingRuleResponseWithDefaults instantiates a new EqualityMatchingRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *EqualityMatchingRuleResponse) GetSchemas() []EnumequalityMatchingRuleSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *EqualityMatchingRuleResponse) GetSchemasOk() (*[]EnumequalityMatchingRuleSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *EqualityMatchingRuleResponse) SetSchemas(v []EnumequalityMatchingRuleSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *EqualityMatchingRuleResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EqualityMatchingRuleResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EqualityMatchingRuleResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


