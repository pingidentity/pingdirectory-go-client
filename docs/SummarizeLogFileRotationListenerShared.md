# SummarizeLogFileRotationListenerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsummarizeLogFileRotationListenerSchemaUrn**](EnumsummarizeLogFileRotationListenerSchemaUrn.md) |  | 
**OutputDirectory** | Pointer to **string** | The path to the directory in which the summarize-access-log output should be written. If no value is provided, the output file will be written into the same directory as the rotated log file. | [optional] 
**Description** | Pointer to **string** | A description for this Log File Rotation Listener | [optional] 
**Enabled** | **bool** | Indicates whether the Log File Rotation Listener is enabled for use. | 

## Methods

### NewSummarizeLogFileRotationListenerShared

`func NewSummarizeLogFileRotationListenerShared(schemas []EnumsummarizeLogFileRotationListenerSchemaUrn, enabled bool, ) *SummarizeLogFileRotationListenerShared`

NewSummarizeLogFileRotationListenerShared instantiates a new SummarizeLogFileRotationListenerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummarizeLogFileRotationListenerSharedWithDefaults

`func NewSummarizeLogFileRotationListenerSharedWithDefaults() *SummarizeLogFileRotationListenerShared`

NewSummarizeLogFileRotationListenerSharedWithDefaults instantiates a new SummarizeLogFileRotationListenerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SummarizeLogFileRotationListenerShared) GetSchemas() []EnumsummarizeLogFileRotationListenerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SummarizeLogFileRotationListenerShared) GetSchemasOk() (*[]EnumsummarizeLogFileRotationListenerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SummarizeLogFileRotationListenerShared) SetSchemas(v []EnumsummarizeLogFileRotationListenerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetOutputDirectory

`func (o *SummarizeLogFileRotationListenerShared) GetOutputDirectory() string`

GetOutputDirectory returns the OutputDirectory field if non-nil, zero value otherwise.

### GetOutputDirectoryOk

`func (o *SummarizeLogFileRotationListenerShared) GetOutputDirectoryOk() (*string, bool)`

GetOutputDirectoryOk returns a tuple with the OutputDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDirectory

`func (o *SummarizeLogFileRotationListenerShared) SetOutputDirectory(v string)`

SetOutputDirectory sets OutputDirectory field to given value.

### HasOutputDirectory

`func (o *SummarizeLogFileRotationListenerShared) HasOutputDirectory() bool`

HasOutputDirectory returns a boolean if a field has been set.

### GetDescription

`func (o *SummarizeLogFileRotationListenerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SummarizeLogFileRotationListenerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SummarizeLogFileRotationListenerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SummarizeLogFileRotationListenerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SummarizeLogFileRotationListenerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SummarizeLogFileRotationListenerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SummarizeLogFileRotationListenerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


