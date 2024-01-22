# AddFileCountLogRetentionPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumfileCountLogRetentionPolicySchemaUrn**](EnumfileCountLogRetentionPolicySchemaUrn.md) |  | 
**NumberOfFiles** | **int64** | Specifies the number of archived log files to retain before the oldest ones are cleaned. | 
**Description** | Pointer to **string** | A description for this Log Retention Policy | [optional] 
**PolicyName** | **string** | Name of the new Log Retention Policy | 

## Methods

### NewAddFileCountLogRetentionPolicyRequest

`func NewAddFileCountLogRetentionPolicyRequest(schemas []EnumfileCountLogRetentionPolicySchemaUrn, numberOfFiles int64, policyName string, ) *AddFileCountLogRetentionPolicyRequest`

NewAddFileCountLogRetentionPolicyRequest instantiates a new AddFileCountLogRetentionPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFileCountLogRetentionPolicyRequestWithDefaults

`func NewAddFileCountLogRetentionPolicyRequestWithDefaults() *AddFileCountLogRetentionPolicyRequest`

NewAddFileCountLogRetentionPolicyRequestWithDefaults instantiates a new AddFileCountLogRetentionPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddFileCountLogRetentionPolicyRequest) GetSchemas() []EnumfileCountLogRetentionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddFileCountLogRetentionPolicyRequest) GetSchemasOk() (*[]EnumfileCountLogRetentionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddFileCountLogRetentionPolicyRequest) SetSchemas(v []EnumfileCountLogRetentionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetNumberOfFiles

`func (o *AddFileCountLogRetentionPolicyRequest) GetNumberOfFiles() int64`

GetNumberOfFiles returns the NumberOfFiles field if non-nil, zero value otherwise.

### GetNumberOfFilesOk

`func (o *AddFileCountLogRetentionPolicyRequest) GetNumberOfFilesOk() (*int64, bool)`

GetNumberOfFilesOk returns a tuple with the NumberOfFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFiles

`func (o *AddFileCountLogRetentionPolicyRequest) SetNumberOfFiles(v int64)`

SetNumberOfFiles sets NumberOfFiles field to given value.


### GetDescription

`func (o *AddFileCountLogRetentionPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddFileCountLogRetentionPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddFileCountLogRetentionPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddFileCountLogRetentionPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyName

`func (o *AddFileCountLogRetentionPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AddFileCountLogRetentionPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AddFileCountLogRetentionPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


