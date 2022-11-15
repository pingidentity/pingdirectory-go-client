# SummarizeLogFileRotationListenerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Log File Rotation Listener | 
**Schemas** | [**[]EnumsummarizeLogFileRotationListenerSchemaUrn**](EnumsummarizeLogFileRotationListenerSchemaUrn.md) |  | 
**OutputDirectory** | Pointer to **string** | The path to the directory in which the summarize-access-log output should be written. If no value is provided, the output file will be written into the same directory as the rotated log file. | [optional] 
**Description** | Pointer to **string** | A description for this Log File Rotation Listener | [optional] 
**Enabled** | **bool** | Indicates whether the Log File Rotation Listener is enabled for use. | 

## Methods

### NewSummarizeLogFileRotationListenerResponse

`func NewSummarizeLogFileRotationListenerResponse(id string, schemas []EnumsummarizeLogFileRotationListenerSchemaUrn, enabled bool, ) *SummarizeLogFileRotationListenerResponse`

NewSummarizeLogFileRotationListenerResponse instantiates a new SummarizeLogFileRotationListenerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummarizeLogFileRotationListenerResponseWithDefaults

`func NewSummarizeLogFileRotationListenerResponseWithDefaults() *SummarizeLogFileRotationListenerResponse`

NewSummarizeLogFileRotationListenerResponseWithDefaults instantiates a new SummarizeLogFileRotationListenerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SummarizeLogFileRotationListenerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SummarizeLogFileRotationListenerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SummarizeLogFileRotationListenerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *SummarizeLogFileRotationListenerResponse) GetSchemas() []EnumsummarizeLogFileRotationListenerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SummarizeLogFileRotationListenerResponse) GetSchemasOk() (*[]EnumsummarizeLogFileRotationListenerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SummarizeLogFileRotationListenerResponse) SetSchemas(v []EnumsummarizeLogFileRotationListenerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetOutputDirectory

`func (o *SummarizeLogFileRotationListenerResponse) GetOutputDirectory() string`

GetOutputDirectory returns the OutputDirectory field if non-nil, zero value otherwise.

### GetOutputDirectoryOk

`func (o *SummarizeLogFileRotationListenerResponse) GetOutputDirectoryOk() (*string, bool)`

GetOutputDirectoryOk returns a tuple with the OutputDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDirectory

`func (o *SummarizeLogFileRotationListenerResponse) SetOutputDirectory(v string)`

SetOutputDirectory sets OutputDirectory field to given value.

### HasOutputDirectory

`func (o *SummarizeLogFileRotationListenerResponse) HasOutputDirectory() bool`

HasOutputDirectory returns a boolean if a field has been set.

### GetDescription

`func (o *SummarizeLogFileRotationListenerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SummarizeLogFileRotationListenerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SummarizeLogFileRotationListenerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SummarizeLogFileRotationListenerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SummarizeLogFileRotationListenerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SummarizeLogFileRotationListenerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SummarizeLogFileRotationListenerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


