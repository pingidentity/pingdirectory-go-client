# CopyLogFileRotationListenerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumcopyLogFileRotationListenerSchemaUrn**](EnumcopyLogFileRotationListenerSchemaUrn.md) |  | 
**CopyToDirectory** | **string** | The path to the directory to which log files should be copied. It must be different from the directory to which the log file is originally written, and administrators should ensure that the filesystem has sufficient space to hold files as they are copied. | 
**CompressOnCopy** | Pointer to **bool** | Indicates whether the file should be gzip-compressed as it is copied into the destination directory. | [optional] 
**Description** | Pointer to **string** | A description for this Log File Rotation Listener | [optional] 
**Enabled** | **bool** | Indicates whether the Log File Rotation Listener is enabled for use. | 

## Methods

### NewCopyLogFileRotationListenerShared

`func NewCopyLogFileRotationListenerShared(schemas []EnumcopyLogFileRotationListenerSchemaUrn, copyToDirectory string, enabled bool, ) *CopyLogFileRotationListenerShared`

NewCopyLogFileRotationListenerShared instantiates a new CopyLogFileRotationListenerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyLogFileRotationListenerSharedWithDefaults

`func NewCopyLogFileRotationListenerSharedWithDefaults() *CopyLogFileRotationListenerShared`

NewCopyLogFileRotationListenerSharedWithDefaults instantiates a new CopyLogFileRotationListenerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *CopyLogFileRotationListenerShared) GetSchemas() []EnumcopyLogFileRotationListenerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CopyLogFileRotationListenerShared) GetSchemasOk() (*[]EnumcopyLogFileRotationListenerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CopyLogFileRotationListenerShared) SetSchemas(v []EnumcopyLogFileRotationListenerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetCopyToDirectory

`func (o *CopyLogFileRotationListenerShared) GetCopyToDirectory() string`

GetCopyToDirectory returns the CopyToDirectory field if non-nil, zero value otherwise.

### GetCopyToDirectoryOk

`func (o *CopyLogFileRotationListenerShared) GetCopyToDirectoryOk() (*string, bool)`

GetCopyToDirectoryOk returns a tuple with the CopyToDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyToDirectory

`func (o *CopyLogFileRotationListenerShared) SetCopyToDirectory(v string)`

SetCopyToDirectory sets CopyToDirectory field to given value.


### GetCompressOnCopy

`func (o *CopyLogFileRotationListenerShared) GetCompressOnCopy() bool`

GetCompressOnCopy returns the CompressOnCopy field if non-nil, zero value otherwise.

### GetCompressOnCopyOk

`func (o *CopyLogFileRotationListenerShared) GetCompressOnCopyOk() (*bool, bool)`

GetCompressOnCopyOk returns a tuple with the CompressOnCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressOnCopy

`func (o *CopyLogFileRotationListenerShared) SetCompressOnCopy(v bool)`

SetCompressOnCopy sets CompressOnCopy field to given value.

### HasCompressOnCopy

`func (o *CopyLogFileRotationListenerShared) HasCompressOnCopy() bool`

HasCompressOnCopy returns a boolean if a field has been set.

### GetDescription

`func (o *CopyLogFileRotationListenerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CopyLogFileRotationListenerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CopyLogFileRotationListenerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CopyLogFileRotationListenerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CopyLogFileRotationListenerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CopyLogFileRotationListenerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CopyLogFileRotationListenerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


