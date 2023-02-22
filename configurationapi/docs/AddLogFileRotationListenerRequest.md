# AddLogFileRotationListenerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListenerName** | **string** | Name of the new Log File Rotation Listener | 
**Schemas** | [**[]EnumthirdPartyLogFileRotationListenerSchemaUrn**](EnumthirdPartyLogFileRotationListenerSchemaUrn.md) |  | 
**OutputDirectory** | Pointer to **string** | The path to the directory in which the summarize-access-log output should be written. If no value is provided, the output file will be written into the same directory as the rotated log file. | [optional] 
**Description** | Pointer to **string** | A description for this Log File Rotation Listener | [optional] 
**Enabled** | **bool** | Indicates whether the Log File Rotation Listener is enabled for use. | 
**CopyToDirectory** | **string** | The path to the directory to which log files should be copied. It must be different from the directory to which the log file is originally written, and administrators should ensure that the filesystem has sufficient space to hold files as they are copied. | 
**CompressOnCopy** | Pointer to **bool** | Indicates whether the file should be gzip-compressed as it is copied into the destination directory. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Log File Rotation Listener. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Log File Rotation Listener. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddLogFileRotationListenerRequest

`func NewAddLogFileRotationListenerRequest(listenerName string, schemas []EnumthirdPartyLogFileRotationListenerSchemaUrn, enabled bool, copyToDirectory string, extensionClass string, ) *AddLogFileRotationListenerRequest`

NewAddLogFileRotationListenerRequest instantiates a new AddLogFileRotationListenerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLogFileRotationListenerRequestWithDefaults

`func NewAddLogFileRotationListenerRequestWithDefaults() *AddLogFileRotationListenerRequest`

NewAddLogFileRotationListenerRequestWithDefaults instantiates a new AddLogFileRotationListenerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListenerName

`func (o *AddLogFileRotationListenerRequest) GetListenerName() string`

GetListenerName returns the ListenerName field if non-nil, zero value otherwise.

### GetListenerNameOk

`func (o *AddLogFileRotationListenerRequest) GetListenerNameOk() (*string, bool)`

GetListenerNameOk returns a tuple with the ListenerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerName

`func (o *AddLogFileRotationListenerRequest) SetListenerName(v string)`

SetListenerName sets ListenerName field to given value.


### GetSchemas

`func (o *AddLogFileRotationListenerRequest) GetSchemas() []EnumthirdPartyLogFileRotationListenerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLogFileRotationListenerRequest) GetSchemasOk() (*[]EnumthirdPartyLogFileRotationListenerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLogFileRotationListenerRequest) SetSchemas(v []EnumthirdPartyLogFileRotationListenerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetOutputDirectory

`func (o *AddLogFileRotationListenerRequest) GetOutputDirectory() string`

GetOutputDirectory returns the OutputDirectory field if non-nil, zero value otherwise.

### GetOutputDirectoryOk

`func (o *AddLogFileRotationListenerRequest) GetOutputDirectoryOk() (*string, bool)`

GetOutputDirectoryOk returns a tuple with the OutputDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDirectory

`func (o *AddLogFileRotationListenerRequest) SetOutputDirectory(v string)`

SetOutputDirectory sets OutputDirectory field to given value.

### HasOutputDirectory

`func (o *AddLogFileRotationListenerRequest) HasOutputDirectory() bool`

HasOutputDirectory returns a boolean if a field has been set.

### GetDescription

`func (o *AddLogFileRotationListenerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLogFileRotationListenerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLogFileRotationListenerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLogFileRotationListenerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddLogFileRotationListenerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddLogFileRotationListenerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddLogFileRotationListenerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCopyToDirectory

`func (o *AddLogFileRotationListenerRequest) GetCopyToDirectory() string`

GetCopyToDirectory returns the CopyToDirectory field if non-nil, zero value otherwise.

### GetCopyToDirectoryOk

`func (o *AddLogFileRotationListenerRequest) GetCopyToDirectoryOk() (*string, bool)`

GetCopyToDirectoryOk returns a tuple with the CopyToDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyToDirectory

`func (o *AddLogFileRotationListenerRequest) SetCopyToDirectory(v string)`

SetCopyToDirectory sets CopyToDirectory field to given value.


### GetCompressOnCopy

`func (o *AddLogFileRotationListenerRequest) GetCompressOnCopy() bool`

GetCompressOnCopy returns the CompressOnCopy field if non-nil, zero value otherwise.

### GetCompressOnCopyOk

`func (o *AddLogFileRotationListenerRequest) GetCompressOnCopyOk() (*bool, bool)`

GetCompressOnCopyOk returns a tuple with the CompressOnCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressOnCopy

`func (o *AddLogFileRotationListenerRequest) SetCompressOnCopy(v bool)`

SetCompressOnCopy sets CompressOnCopy field to given value.

### HasCompressOnCopy

`func (o *AddLogFileRotationListenerRequest) HasCompressOnCopy() bool`

HasCompressOnCopy returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddLogFileRotationListenerRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddLogFileRotationListenerRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddLogFileRotationListenerRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddLogFileRotationListenerRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddLogFileRotationListenerRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddLogFileRotationListenerRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddLogFileRotationListenerRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


