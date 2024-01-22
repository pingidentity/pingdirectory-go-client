# AddThirdPartyPostLdifExportTaskProcessorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn**](EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Post LDIF Export Task Processor. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Post LDIF Export Task Processor. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Post LDIF Export Task Processor | [optional] 
**Enabled** | **bool** | Indicates whether the Post LDIF Export Task Processor is enabled for use. | 
**ProcessorName** | **string** | Name of the new Post LDIF Export Task Processor | 

## Methods

### NewAddThirdPartyPostLdifExportTaskProcessorRequest

`func NewAddThirdPartyPostLdifExportTaskProcessorRequest(schemas []EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn, extensionClass string, enabled bool, processorName string, ) *AddThirdPartyPostLdifExportTaskProcessorRequest`

NewAddThirdPartyPostLdifExportTaskProcessorRequest instantiates a new AddThirdPartyPostLdifExportTaskProcessorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddThirdPartyPostLdifExportTaskProcessorRequestWithDefaults

`func NewAddThirdPartyPostLdifExportTaskProcessorRequestWithDefaults() *AddThirdPartyPostLdifExportTaskProcessorRequest`

NewAddThirdPartyPostLdifExportTaskProcessorRequestWithDefaults instantiates a new AddThirdPartyPostLdifExportTaskProcessorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) GetSchemas() []EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) GetSchemasOk() (*[]EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) SetSchemas(v []EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProcessorName

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) GetProcessorName() string`

GetProcessorName returns the ProcessorName field if non-nil, zero value otherwise.

### GetProcessorNameOk

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) GetProcessorNameOk() (*string, bool)`

GetProcessorNameOk returns a tuple with the ProcessorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorName

`func (o *AddThirdPartyPostLdifExportTaskProcessorRequest) SetProcessorName(v string)`

SetProcessorName sets ProcessorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


