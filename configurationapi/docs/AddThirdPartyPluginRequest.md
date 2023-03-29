# AddThirdPartyPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginName** | **string** | Name of the new Plugin | 
**Schemas** | [**[]EnumthirdPartyPluginSchemaUrn**](EnumthirdPartyPluginSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Plugin. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Plugin. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**RequestCriteria** | Pointer to **string** | Specifies a set of request criteria that may be used to indicate that this Third Party Plugin should only be invoked for operations in which the associated request matches this criteria. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**PluginType** | [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) | Specifies the set of plug-in types for the plug-in, which specifies the times at which the plug-in is invoked. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 

## Methods

### NewAddThirdPartyPluginRequest

`func NewAddThirdPartyPluginRequest(pluginName string, schemas []EnumthirdPartyPluginSchemaUrn, extensionClass string, enabled bool, pluginType []EnumpluginPluginTypeProp, ) *AddThirdPartyPluginRequest`

NewAddThirdPartyPluginRequest instantiates a new AddThirdPartyPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddThirdPartyPluginRequestWithDefaults

`func NewAddThirdPartyPluginRequestWithDefaults() *AddThirdPartyPluginRequest`

NewAddThirdPartyPluginRequestWithDefaults instantiates a new AddThirdPartyPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginName

`func (o *AddThirdPartyPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddThirdPartyPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddThirdPartyPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### GetSchemas

`func (o *AddThirdPartyPluginRequest) GetSchemas() []EnumthirdPartyPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddThirdPartyPluginRequest) GetSchemasOk() (*[]EnumthirdPartyPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddThirdPartyPluginRequest) SetSchemas(v []EnumthirdPartyPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *AddThirdPartyPluginRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddThirdPartyPluginRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddThirdPartyPluginRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddThirdPartyPluginRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddThirdPartyPluginRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddThirdPartyPluginRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddThirdPartyPluginRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *AddThirdPartyPluginRequest) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddThirdPartyPluginRequest) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddThirdPartyPluginRequest) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddThirdPartyPluginRequest) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *AddThirdPartyPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddThirdPartyPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddThirdPartyPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddThirdPartyPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddThirdPartyPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddThirdPartyPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddThirdPartyPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPluginType

`func (o *AddThirdPartyPluginRequest) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *AddThirdPartyPluginRequest) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *AddThirdPartyPluginRequest) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.


### GetInvokeForInternalOperations

`func (o *AddThirdPartyPluginRequest) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *AddThirdPartyPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *AddThirdPartyPluginRequest) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *AddThirdPartyPluginRequest) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


