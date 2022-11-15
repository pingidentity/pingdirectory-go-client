# ApproximateMatchingRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumapproximateMatchingRuleSchemaUrn**](EnumapproximateMatchingRuleSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicates whether the Matching Rule is enabled for use. | 

## Methods

### NewApproximateMatchingRuleResponse

`func NewApproximateMatchingRuleResponse(schemas []EnumapproximateMatchingRuleSchemaUrn, enabled bool, ) *ApproximateMatchingRuleResponse`

NewApproximateMatchingRuleResponse instantiates a new ApproximateMatchingRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproximateMatchingRuleResponseWithDefaults

`func NewApproximateMatchingRuleResponseWithDefaults() *ApproximateMatchingRuleResponse`

NewApproximateMatchingRuleResponseWithDefaults instantiates a new ApproximateMatchingRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ApproximateMatchingRuleResponse) GetSchemas() []EnumapproximateMatchingRuleSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ApproximateMatchingRuleResponse) GetSchemasOk() (*[]EnumapproximateMatchingRuleSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ApproximateMatchingRuleResponse) SetSchemas(v []EnumapproximateMatchingRuleSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *ApproximateMatchingRuleResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApproximateMatchingRuleResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApproximateMatchingRuleResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


