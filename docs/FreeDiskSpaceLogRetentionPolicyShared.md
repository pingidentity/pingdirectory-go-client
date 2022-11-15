# FreeDiskSpaceLogRetentionPolicyShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumfreeDiskSpaceLogRetentionPolicySchemaUrn**](EnumfreeDiskSpaceLogRetentionPolicySchemaUrn.md) |  | 
**FreeDiskSpace** | **string** | Specifies the minimum amount of free disk space that should be available on the file system on which the archived log files are stored. | 
**Description** | Pointer to **string** | A description for this Log Retention Policy | [optional] 

## Methods

### NewFreeDiskSpaceLogRetentionPolicyShared

`func NewFreeDiskSpaceLogRetentionPolicyShared(schemas []EnumfreeDiskSpaceLogRetentionPolicySchemaUrn, freeDiskSpace string, ) *FreeDiskSpaceLogRetentionPolicyShared`

NewFreeDiskSpaceLogRetentionPolicyShared instantiates a new FreeDiskSpaceLogRetentionPolicyShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeDiskSpaceLogRetentionPolicySharedWithDefaults

`func NewFreeDiskSpaceLogRetentionPolicySharedWithDefaults() *FreeDiskSpaceLogRetentionPolicyShared`

NewFreeDiskSpaceLogRetentionPolicySharedWithDefaults instantiates a new FreeDiskSpaceLogRetentionPolicyShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *FreeDiskSpaceLogRetentionPolicyShared) GetSchemas() []EnumfreeDiskSpaceLogRetentionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FreeDiskSpaceLogRetentionPolicyShared) GetSchemasOk() (*[]EnumfreeDiskSpaceLogRetentionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FreeDiskSpaceLogRetentionPolicyShared) SetSchemas(v []EnumfreeDiskSpaceLogRetentionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetFreeDiskSpace

`func (o *FreeDiskSpaceLogRetentionPolicyShared) GetFreeDiskSpace() string`

GetFreeDiskSpace returns the FreeDiskSpace field if non-nil, zero value otherwise.

### GetFreeDiskSpaceOk

`func (o *FreeDiskSpaceLogRetentionPolicyShared) GetFreeDiskSpaceOk() (*string, bool)`

GetFreeDiskSpaceOk returns a tuple with the FreeDiskSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeDiskSpace

`func (o *FreeDiskSpaceLogRetentionPolicyShared) SetFreeDiskSpace(v string)`

SetFreeDiskSpace sets FreeDiskSpace field to given value.


### GetDescription

`func (o *FreeDiskSpaceLogRetentionPolicyShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FreeDiskSpaceLogRetentionPolicyShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FreeDiskSpaceLogRetentionPolicyShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FreeDiskSpaceLogRetentionPolicyShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


