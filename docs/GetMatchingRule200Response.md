# GetMatchingRule200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsubstringMatchingRuleSchemaUrn**](EnumsubstringMatchingRuleSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicates whether the Matching Rule is enabled for use. | 

## Methods

### NewGetMatchingRule200Response

`func NewGetMatchingRule200Response(schemas []EnumsubstringMatchingRuleSchemaUrn, enabled bool, ) *GetMatchingRule200Response`

NewGetMatchingRule200Response instantiates a new GetMatchingRule200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMatchingRule200ResponseWithDefaults

`func NewGetMatchingRule200ResponseWithDefaults() *GetMatchingRule200Response`

NewGetMatchingRule200ResponseWithDefaults instantiates a new GetMatchingRule200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GetMatchingRule200Response) GetSchemas() []EnumsubstringMatchingRuleSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetMatchingRule200Response) GetSchemasOk() (*[]EnumsubstringMatchingRuleSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetMatchingRule200Response) SetSchemas(v []EnumsubstringMatchingRuleSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *GetMatchingRule200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetMatchingRule200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetMatchingRule200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


