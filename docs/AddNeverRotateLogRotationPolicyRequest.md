# AddNeverRotateLogRotationPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | **string** | Name of the new Log Rotation Policy | 
**Schemas** | [**[]EnumneverRotateLogRotationPolicySchemaUrn**](EnumneverRotateLogRotationPolicySchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Log Rotation Policy | [optional] 

## Methods

### NewAddNeverRotateLogRotationPolicyRequest

`func NewAddNeverRotateLogRotationPolicyRequest(policyName string, schemas []EnumneverRotateLogRotationPolicySchemaUrn, ) *AddNeverRotateLogRotationPolicyRequest`

NewAddNeverRotateLogRotationPolicyRequest instantiates a new AddNeverRotateLogRotationPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddNeverRotateLogRotationPolicyRequestWithDefaults

`func NewAddNeverRotateLogRotationPolicyRequestWithDefaults() *AddNeverRotateLogRotationPolicyRequest`

NewAddNeverRotateLogRotationPolicyRequestWithDefaults instantiates a new AddNeverRotateLogRotationPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyName

`func (o *AddNeverRotateLogRotationPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AddNeverRotateLogRotationPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AddNeverRotateLogRotationPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetSchemas

`func (o *AddNeverRotateLogRotationPolicyRequest) GetSchemas() []EnumneverRotateLogRotationPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddNeverRotateLogRotationPolicyRequest) GetSchemasOk() (*[]EnumneverRotateLogRotationPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddNeverRotateLogRotationPolicyRequest) SetSchemas(v []EnumneverRotateLogRotationPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddNeverRotateLogRotationPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddNeverRotateLogRotationPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddNeverRotateLogRotationPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddNeverRotateLogRotationPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


