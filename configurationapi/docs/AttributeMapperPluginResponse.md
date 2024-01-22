# AttributeMapperPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumattributeMapperPluginSchemaUrn**](EnumattributeMapperPluginSchemaUrn.md) |  | 
**PluginType** | [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) |  | 
**SourceAttribute** | **string** | Specifies the source attribute type that may appear in client requests which should be remapped to the target attribute. Note that the source attribute type must be defined in the server schema and must not be equal to the target attribute type. | 
**TargetAttribute** | **string** | Specifies the target attribute type to which the source attribute type should be mapped. Note that the target attribute type must be defined in the server schema and must not be equal to the source attribute type. | 
**EnableControlMapping** | **bool** | Indicates whether mapping should be applied to attribute types that may be present in specific controls. If enabled, attribute mapping will only be applied for control types which are specifically supported by the attribute mapper plugin. | 
**AlwaysMapResponses** | **bool** | Indicates whether the target attribute in response messages should always be remapped back to the source attribute. If this is \&quot;false\&quot;, then the mapping will be performed for a response message only if one or more elements of the associated request are mapped. Otherwise, the mapping will be performed for all responses regardless of whether the mapping was applied to the request. | 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Plugin | 

## Methods

### NewAttributeMapperPluginResponse

`func NewAttributeMapperPluginResponse(schemas []EnumattributeMapperPluginSchemaUrn, pluginType []EnumpluginPluginTypeProp, sourceAttribute string, targetAttribute string, enableControlMapping bool, alwaysMapResponses bool, enabled bool, id string, ) *AttributeMapperPluginResponse`

NewAttributeMapperPluginResponse instantiates a new AttributeMapperPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeMapperPluginResponseWithDefaults

`func NewAttributeMapperPluginResponseWithDefaults() *AttributeMapperPluginResponse`

NewAttributeMapperPluginResponseWithDefaults instantiates a new AttributeMapperPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AttributeMapperPluginResponse) GetSchemas() []EnumattributeMapperPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AttributeMapperPluginResponse) GetSchemasOk() (*[]EnumattributeMapperPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AttributeMapperPluginResponse) SetSchemas(v []EnumattributeMapperPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPluginType

`func (o *AttributeMapperPluginResponse) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *AttributeMapperPluginResponse) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *AttributeMapperPluginResponse) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.


### GetSourceAttribute

`func (o *AttributeMapperPluginResponse) GetSourceAttribute() string`

GetSourceAttribute returns the SourceAttribute field if non-nil, zero value otherwise.

### GetSourceAttributeOk

`func (o *AttributeMapperPluginResponse) GetSourceAttributeOk() (*string, bool)`

GetSourceAttributeOk returns a tuple with the SourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttribute

`func (o *AttributeMapperPluginResponse) SetSourceAttribute(v string)`

SetSourceAttribute sets SourceAttribute field to given value.


### GetTargetAttribute

`func (o *AttributeMapperPluginResponse) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *AttributeMapperPluginResponse) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *AttributeMapperPluginResponse) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetEnableControlMapping

`func (o *AttributeMapperPluginResponse) GetEnableControlMapping() bool`

GetEnableControlMapping returns the EnableControlMapping field if non-nil, zero value otherwise.

### GetEnableControlMappingOk

`func (o *AttributeMapperPluginResponse) GetEnableControlMappingOk() (*bool, bool)`

GetEnableControlMappingOk returns a tuple with the EnableControlMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableControlMapping

`func (o *AttributeMapperPluginResponse) SetEnableControlMapping(v bool)`

SetEnableControlMapping sets EnableControlMapping field to given value.


### GetAlwaysMapResponses

`func (o *AttributeMapperPluginResponse) GetAlwaysMapResponses() bool`

GetAlwaysMapResponses returns the AlwaysMapResponses field if non-nil, zero value otherwise.

### GetAlwaysMapResponsesOk

`func (o *AttributeMapperPluginResponse) GetAlwaysMapResponsesOk() (*bool, bool)`

GetAlwaysMapResponsesOk returns a tuple with the AlwaysMapResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysMapResponses

`func (o *AttributeMapperPluginResponse) SetAlwaysMapResponses(v bool)`

SetAlwaysMapResponses sets AlwaysMapResponses field to given value.


### GetDescription

`func (o *AttributeMapperPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AttributeMapperPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AttributeMapperPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AttributeMapperPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AttributeMapperPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AttributeMapperPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AttributeMapperPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *AttributeMapperPluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *AttributeMapperPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *AttributeMapperPluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *AttributeMapperPluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetMeta

`func (o *AttributeMapperPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AttributeMapperPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AttributeMapperPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AttributeMapperPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AttributeMapperPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AttributeMapperPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AttributeMapperPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AttributeMapperPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *AttributeMapperPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttributeMapperPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttributeMapperPluginResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


