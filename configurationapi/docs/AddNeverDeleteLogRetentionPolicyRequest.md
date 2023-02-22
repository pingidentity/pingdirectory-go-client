# AddNeverDeleteLogRetentionPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | **string** | Name of the new Log Retention Policy | 
**Schemas** | [**[]EnumneverDeleteLogRetentionPolicySchemaUrn**](EnumneverDeleteLogRetentionPolicySchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Log Retention Policy | [optional] 

## Methods

### NewAddNeverDeleteLogRetentionPolicyRequest

`func NewAddNeverDeleteLogRetentionPolicyRequest(policyName string, schemas []EnumneverDeleteLogRetentionPolicySchemaUrn, ) *AddNeverDeleteLogRetentionPolicyRequest`

NewAddNeverDeleteLogRetentionPolicyRequest instantiates a new AddNeverDeleteLogRetentionPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddNeverDeleteLogRetentionPolicyRequestWithDefaults

`func NewAddNeverDeleteLogRetentionPolicyRequestWithDefaults() *AddNeverDeleteLogRetentionPolicyRequest`

NewAddNeverDeleteLogRetentionPolicyRequestWithDefaults instantiates a new AddNeverDeleteLogRetentionPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyName

`func (o *AddNeverDeleteLogRetentionPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AddNeverDeleteLogRetentionPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AddNeverDeleteLogRetentionPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetSchemas

`func (o *AddNeverDeleteLogRetentionPolicyRequest) GetSchemas() []EnumneverDeleteLogRetentionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddNeverDeleteLogRetentionPolicyRequest) GetSchemasOk() (*[]EnumneverDeleteLogRetentionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddNeverDeleteLogRetentionPolicyRequest) SetSchemas(v []EnumneverDeleteLogRetentionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddNeverDeleteLogRetentionPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddNeverDeleteLogRetentionPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddNeverDeleteLogRetentionPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddNeverDeleteLogRetentionPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


