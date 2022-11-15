# FileCountLogRetentionPolicyShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumfileCountLogRetentionPolicySchemaUrn**](EnumfileCountLogRetentionPolicySchemaUrn.md) |  | 
**NumberOfFiles** | **int32** | Specifies the number of archived log files to retain before the oldest ones are cleaned. | 
**Description** | Pointer to **string** | A description for this Log Retention Policy | [optional] 

## Methods

### NewFileCountLogRetentionPolicyShared

`func NewFileCountLogRetentionPolicyShared(schemas []EnumfileCountLogRetentionPolicySchemaUrn, numberOfFiles int32, ) *FileCountLogRetentionPolicyShared`

NewFileCountLogRetentionPolicyShared instantiates a new FileCountLogRetentionPolicyShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileCountLogRetentionPolicySharedWithDefaults

`func NewFileCountLogRetentionPolicySharedWithDefaults() *FileCountLogRetentionPolicyShared`

NewFileCountLogRetentionPolicySharedWithDefaults instantiates a new FileCountLogRetentionPolicyShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *FileCountLogRetentionPolicyShared) GetSchemas() []EnumfileCountLogRetentionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FileCountLogRetentionPolicyShared) GetSchemasOk() (*[]EnumfileCountLogRetentionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FileCountLogRetentionPolicyShared) SetSchemas(v []EnumfileCountLogRetentionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetNumberOfFiles

`func (o *FileCountLogRetentionPolicyShared) GetNumberOfFiles() int32`

GetNumberOfFiles returns the NumberOfFiles field if non-nil, zero value otherwise.

### GetNumberOfFilesOk

`func (o *FileCountLogRetentionPolicyShared) GetNumberOfFilesOk() (*int32, bool)`

GetNumberOfFilesOk returns a tuple with the NumberOfFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFiles

`func (o *FileCountLogRetentionPolicyShared) SetNumberOfFiles(v int32)`

SetNumberOfFiles sets NumberOfFiles field to given value.


### GetDescription

`func (o *FileCountLogRetentionPolicyShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FileCountLogRetentionPolicyShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FileCountLogRetentionPolicyShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FileCountLogRetentionPolicyShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


