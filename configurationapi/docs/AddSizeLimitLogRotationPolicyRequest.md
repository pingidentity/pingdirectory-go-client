# AddSizeLimitLogRotationPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | **string** | Name of the new Log Rotation Policy | 
**Schemas** | [**[]EnumsizeLimitLogRotationPolicySchemaUrn**](EnumsizeLimitLogRotationPolicySchemaUrn.md) |  | 
**FileSizeLimit** | **string** | Specifies the maximum size that a log file can reach before it is rotated. | 
**Description** | Pointer to **string** | A description for this Log Rotation Policy | [optional] 

## Methods

### NewAddSizeLimitLogRotationPolicyRequest

`func NewAddSizeLimitLogRotationPolicyRequest(policyName string, schemas []EnumsizeLimitLogRotationPolicySchemaUrn, fileSizeLimit string, ) *AddSizeLimitLogRotationPolicyRequest`

NewAddSizeLimitLogRotationPolicyRequest instantiates a new AddSizeLimitLogRotationPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSizeLimitLogRotationPolicyRequestWithDefaults

`func NewAddSizeLimitLogRotationPolicyRequestWithDefaults() *AddSizeLimitLogRotationPolicyRequest`

NewAddSizeLimitLogRotationPolicyRequestWithDefaults instantiates a new AddSizeLimitLogRotationPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyName

`func (o *AddSizeLimitLogRotationPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AddSizeLimitLogRotationPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AddSizeLimitLogRotationPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetSchemas

`func (o *AddSizeLimitLogRotationPolicyRequest) GetSchemas() []EnumsizeLimitLogRotationPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSizeLimitLogRotationPolicyRequest) GetSchemasOk() (*[]EnumsizeLimitLogRotationPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSizeLimitLogRotationPolicyRequest) SetSchemas(v []EnumsizeLimitLogRotationPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetFileSizeLimit

`func (o *AddSizeLimitLogRotationPolicyRequest) GetFileSizeLimit() string`

GetFileSizeLimit returns the FileSizeLimit field if non-nil, zero value otherwise.

### GetFileSizeLimitOk

`func (o *AddSizeLimitLogRotationPolicyRequest) GetFileSizeLimitOk() (*string, bool)`

GetFileSizeLimitOk returns a tuple with the FileSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeLimit

`func (o *AddSizeLimitLogRotationPolicyRequest) SetFileSizeLimit(v string)`

SetFileSizeLimit sets FileSizeLimit field to given value.


### GetDescription

`func (o *AddSizeLimitLogRotationPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSizeLimitLogRotationPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSizeLimitLogRotationPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSizeLimitLogRotationPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


