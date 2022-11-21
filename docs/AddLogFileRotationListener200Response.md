# AddLogFileRotationListener200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Log File Rotation Listener | 
**Schemas** | [**[]EnumthirdPartyLogFileRotationListenerSchemaUrn**](EnumthirdPartyLogFileRotationListenerSchemaUrn.md) |  | 
**OutputDirectory** | Pointer to **string** | The path to the directory in which the summarize-access-log output should be written. If no value is provided, the output file will be written into the same directory as the rotated log file. | [optional] 
**Description** | Pointer to **string** | A description for this Log File Rotation Listener | [optional] 
**Enabled** | **bool** | Indicates whether the Log File Rotation Listener is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**CopyToDirectory** | **string** | The path to the directory to which log files should be copied. It must be different from the directory to which the log file is originally written, and administrators should ensure that the filesystem has sufficient space to hold files as they are copied. | 
**CompressOnCopy** | Pointer to **bool** | Indicates whether the file should be gzip-compressed as it is copied into the destination directory. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Log File Rotation Listener. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Log File Rotation Listener. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddLogFileRotationListener200Response

`func NewAddLogFileRotationListener200Response(id string, schemas []EnumthirdPartyLogFileRotationListenerSchemaUrn, enabled bool, copyToDirectory string, extensionClass string, ) *AddLogFileRotationListener200Response`

NewAddLogFileRotationListener200Response instantiates a new AddLogFileRotationListener200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLogFileRotationListener200ResponseWithDefaults

`func NewAddLogFileRotationListener200ResponseWithDefaults() *AddLogFileRotationListener200Response`

NewAddLogFileRotationListener200ResponseWithDefaults instantiates a new AddLogFileRotationListener200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddLogFileRotationListener200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddLogFileRotationListener200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddLogFileRotationListener200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddLogFileRotationListener200Response) GetSchemas() []EnumthirdPartyLogFileRotationListenerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLogFileRotationListener200Response) GetSchemasOk() (*[]EnumthirdPartyLogFileRotationListenerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLogFileRotationListener200Response) SetSchemas(v []EnumthirdPartyLogFileRotationListenerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetOutputDirectory

`func (o *AddLogFileRotationListener200Response) GetOutputDirectory() string`

GetOutputDirectory returns the OutputDirectory field if non-nil, zero value otherwise.

### GetOutputDirectoryOk

`func (o *AddLogFileRotationListener200Response) GetOutputDirectoryOk() (*string, bool)`

GetOutputDirectoryOk returns a tuple with the OutputDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDirectory

`func (o *AddLogFileRotationListener200Response) SetOutputDirectory(v string)`

SetOutputDirectory sets OutputDirectory field to given value.

### HasOutputDirectory

`func (o *AddLogFileRotationListener200Response) HasOutputDirectory() bool`

HasOutputDirectory returns a boolean if a field has been set.

### GetDescription

`func (o *AddLogFileRotationListener200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLogFileRotationListener200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLogFileRotationListener200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLogFileRotationListener200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddLogFileRotationListener200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddLogFileRotationListener200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddLogFileRotationListener200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *AddLogFileRotationListener200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddLogFileRotationListener200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddLogFileRotationListener200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddLogFileRotationListener200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetCopyToDirectory

`func (o *AddLogFileRotationListener200Response) GetCopyToDirectory() string`

GetCopyToDirectory returns the CopyToDirectory field if non-nil, zero value otherwise.

### GetCopyToDirectoryOk

`func (o *AddLogFileRotationListener200Response) GetCopyToDirectoryOk() (*string, bool)`

GetCopyToDirectoryOk returns a tuple with the CopyToDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyToDirectory

`func (o *AddLogFileRotationListener200Response) SetCopyToDirectory(v string)`

SetCopyToDirectory sets CopyToDirectory field to given value.


### GetCompressOnCopy

`func (o *AddLogFileRotationListener200Response) GetCompressOnCopy() bool`

GetCompressOnCopy returns the CompressOnCopy field if non-nil, zero value otherwise.

### GetCompressOnCopyOk

`func (o *AddLogFileRotationListener200Response) GetCompressOnCopyOk() (*bool, bool)`

GetCompressOnCopyOk returns a tuple with the CompressOnCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressOnCopy

`func (o *AddLogFileRotationListener200Response) SetCompressOnCopy(v bool)`

SetCompressOnCopy sets CompressOnCopy field to given value.

### HasCompressOnCopy

`func (o *AddLogFileRotationListener200Response) HasCompressOnCopy() bool`

HasCompressOnCopy returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddLogFileRotationListener200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddLogFileRotationListener200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddLogFileRotationListener200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddLogFileRotationListener200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddLogFileRotationListener200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddLogFileRotationListener200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddLogFileRotationListener200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


