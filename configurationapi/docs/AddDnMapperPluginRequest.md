# AddDnMapperPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumdnMapperPluginSchemaUrn**](EnumdnMapperPluginSchemaUrn.md) |  | 
**PluginType** | Pointer to [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) |  | [optional] 
**SourceDN** | **string** | Specifies the source DN that may appear in client requests which should be remapped to the target DN. Note that the source DN must not be equal to the target DN. | 
**TargetDN** | **string** | Specifies the DN to which the source DN should be mapped. Note that the target DN must not be equal to the source DN. | 
**EnableAttributeMapping** | Pointer to **bool** | Indicates whether DN mapping should be applied to the values of attributes with appropriate syntaxes. | [optional] 
**MapAttribute** | Pointer to **[]string** | Specifies a set of specific attributes for which DN mapping should be applied. This will only be applicable if the enable-attribute-mapping property has a value of \&quot;true\&quot;. Any attributes listed must be defined in the server schema with either the distinguished name syntax or the name and optional UID syntax. | [optional] 
**EnableControlMapping** | Pointer to **bool** | Indicates whether DN mapping should be applied to DNs that may be present in specific controls. DN mapping will only be applied for control types which are specifically supported by the DN mapper plugin. | [optional] 
**AlwaysMapResponses** | Pointer to **bool** | Indicates whether DNs in response messages containing the target DN should always be remapped back to the source DN. If this is \&quot;false\&quot;, then mapping will be performed for a response message only if one or more elements of the associated request are mapped. Otherwise, the mapping will be performed for all responses regardless of whether the mapping was applied to the request. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**PluginName** | **string** | Name of the new Plugin | 

## Methods

### NewAddDnMapperPluginRequest

`func NewAddDnMapperPluginRequest(schemas []EnumdnMapperPluginSchemaUrn, sourceDN string, targetDN string, enabled bool, pluginName string, ) *AddDnMapperPluginRequest`

NewAddDnMapperPluginRequest instantiates a new AddDnMapperPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDnMapperPluginRequestWithDefaults

`func NewAddDnMapperPluginRequestWithDefaults() *AddDnMapperPluginRequest`

NewAddDnMapperPluginRequestWithDefaults instantiates a new AddDnMapperPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddDnMapperPluginRequest) GetSchemas() []EnumdnMapperPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDnMapperPluginRequest) GetSchemasOk() (*[]EnumdnMapperPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDnMapperPluginRequest) SetSchemas(v []EnumdnMapperPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPluginType

`func (o *AddDnMapperPluginRequest) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *AddDnMapperPluginRequest) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *AddDnMapperPluginRequest) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.

### HasPluginType

`func (o *AddDnMapperPluginRequest) HasPluginType() bool`

HasPluginType returns a boolean if a field has been set.

### GetSourceDN

`func (o *AddDnMapperPluginRequest) GetSourceDN() string`

GetSourceDN returns the SourceDN field if non-nil, zero value otherwise.

### GetSourceDNOk

`func (o *AddDnMapperPluginRequest) GetSourceDNOk() (*string, bool)`

GetSourceDNOk returns a tuple with the SourceDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDN

`func (o *AddDnMapperPluginRequest) SetSourceDN(v string)`

SetSourceDN sets SourceDN field to given value.


### GetTargetDN

`func (o *AddDnMapperPluginRequest) GetTargetDN() string`

GetTargetDN returns the TargetDN field if non-nil, zero value otherwise.

### GetTargetDNOk

`func (o *AddDnMapperPluginRequest) GetTargetDNOk() (*string, bool)`

GetTargetDNOk returns a tuple with the TargetDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDN

`func (o *AddDnMapperPluginRequest) SetTargetDN(v string)`

SetTargetDN sets TargetDN field to given value.


### GetEnableAttributeMapping

`func (o *AddDnMapperPluginRequest) GetEnableAttributeMapping() bool`

GetEnableAttributeMapping returns the EnableAttributeMapping field if non-nil, zero value otherwise.

### GetEnableAttributeMappingOk

`func (o *AddDnMapperPluginRequest) GetEnableAttributeMappingOk() (*bool, bool)`

GetEnableAttributeMappingOk returns a tuple with the EnableAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAttributeMapping

`func (o *AddDnMapperPluginRequest) SetEnableAttributeMapping(v bool)`

SetEnableAttributeMapping sets EnableAttributeMapping field to given value.

### HasEnableAttributeMapping

`func (o *AddDnMapperPluginRequest) HasEnableAttributeMapping() bool`

HasEnableAttributeMapping returns a boolean if a field has been set.

### GetMapAttribute

`func (o *AddDnMapperPluginRequest) GetMapAttribute() []string`

GetMapAttribute returns the MapAttribute field if non-nil, zero value otherwise.

### GetMapAttributeOk

`func (o *AddDnMapperPluginRequest) GetMapAttributeOk() (*[]string, bool)`

GetMapAttributeOk returns a tuple with the MapAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAttribute

`func (o *AddDnMapperPluginRequest) SetMapAttribute(v []string)`

SetMapAttribute sets MapAttribute field to given value.

### HasMapAttribute

`func (o *AddDnMapperPluginRequest) HasMapAttribute() bool`

HasMapAttribute returns a boolean if a field has been set.

### GetEnableControlMapping

`func (o *AddDnMapperPluginRequest) GetEnableControlMapping() bool`

GetEnableControlMapping returns the EnableControlMapping field if non-nil, zero value otherwise.

### GetEnableControlMappingOk

`func (o *AddDnMapperPluginRequest) GetEnableControlMappingOk() (*bool, bool)`

GetEnableControlMappingOk returns a tuple with the EnableControlMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableControlMapping

`func (o *AddDnMapperPluginRequest) SetEnableControlMapping(v bool)`

SetEnableControlMapping sets EnableControlMapping field to given value.

### HasEnableControlMapping

`func (o *AddDnMapperPluginRequest) HasEnableControlMapping() bool`

HasEnableControlMapping returns a boolean if a field has been set.

### GetAlwaysMapResponses

`func (o *AddDnMapperPluginRequest) GetAlwaysMapResponses() bool`

GetAlwaysMapResponses returns the AlwaysMapResponses field if non-nil, zero value otherwise.

### GetAlwaysMapResponsesOk

`func (o *AddDnMapperPluginRequest) GetAlwaysMapResponsesOk() (*bool, bool)`

GetAlwaysMapResponsesOk returns a tuple with the AlwaysMapResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysMapResponses

`func (o *AddDnMapperPluginRequest) SetAlwaysMapResponses(v bool)`

SetAlwaysMapResponses sets AlwaysMapResponses field to given value.

### HasAlwaysMapResponses

`func (o *AddDnMapperPluginRequest) HasAlwaysMapResponses() bool`

HasAlwaysMapResponses returns a boolean if a field has been set.

### GetDescription

`func (o *AddDnMapperPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddDnMapperPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddDnMapperPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddDnMapperPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddDnMapperPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddDnMapperPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddDnMapperPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *AddDnMapperPluginRequest) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *AddDnMapperPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *AddDnMapperPluginRequest) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *AddDnMapperPluginRequest) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetPluginName

`func (o *AddDnMapperPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddDnMapperPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddDnMapperPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


