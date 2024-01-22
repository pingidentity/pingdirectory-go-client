# DnMapperPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumdnMapperPluginSchemaUrn**](EnumdnMapperPluginSchemaUrn.md) |  | 
**PluginType** | [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) |  | 
**SourceDN** | **string** | Specifies the source DN that may appear in client requests which should be remapped to the target DN. Note that the source DN must not be equal to the target DN. | 
**TargetDN** | **string** | Specifies the DN to which the source DN should be mapped. Note that the target DN must not be equal to the source DN. | 
**EnableAttributeMapping** | **bool** | Indicates whether DN mapping should be applied to the values of attributes with appropriate syntaxes. | 
**MapAttribute** | Pointer to **[]string** | Specifies a set of specific attributes for which DN mapping should be applied. This will only be applicable if the enable-attribute-mapping property has a value of \&quot;true\&quot;. Any attributes listed must be defined in the server schema with either the distinguished name syntax or the name and optional UID syntax. | [optional] 
**EnableControlMapping** | **bool** | Indicates whether DN mapping should be applied to DNs that may be present in specific controls. DN mapping will only be applied for control types which are specifically supported by the DN mapper plugin. | 
**AlwaysMapResponses** | **bool** | Indicates whether DNs in response messages containing the target DN should always be remapped back to the source DN. If this is \&quot;false\&quot;, then mapping will be performed for a response message only if one or more elements of the associated request are mapped. Otherwise, the mapping will be performed for all responses regardless of whether the mapping was applied to the request. | 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Plugin | 

## Methods

### NewDnMapperPluginResponse

`func NewDnMapperPluginResponse(schemas []EnumdnMapperPluginSchemaUrn, pluginType []EnumpluginPluginTypeProp, sourceDN string, targetDN string, enableAttributeMapping bool, enableControlMapping bool, alwaysMapResponses bool, enabled bool, id string, ) *DnMapperPluginResponse`

NewDnMapperPluginResponse instantiates a new DnMapperPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnMapperPluginResponseWithDefaults

`func NewDnMapperPluginResponseWithDefaults() *DnMapperPluginResponse`

NewDnMapperPluginResponseWithDefaults instantiates a new DnMapperPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *DnMapperPluginResponse) GetSchemas() []EnumdnMapperPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DnMapperPluginResponse) GetSchemasOk() (*[]EnumdnMapperPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DnMapperPluginResponse) SetSchemas(v []EnumdnMapperPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPluginType

`func (o *DnMapperPluginResponse) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *DnMapperPluginResponse) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *DnMapperPluginResponse) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.


### GetSourceDN

`func (o *DnMapperPluginResponse) GetSourceDN() string`

GetSourceDN returns the SourceDN field if non-nil, zero value otherwise.

### GetSourceDNOk

`func (o *DnMapperPluginResponse) GetSourceDNOk() (*string, bool)`

GetSourceDNOk returns a tuple with the SourceDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDN

`func (o *DnMapperPluginResponse) SetSourceDN(v string)`

SetSourceDN sets SourceDN field to given value.


### GetTargetDN

`func (o *DnMapperPluginResponse) GetTargetDN() string`

GetTargetDN returns the TargetDN field if non-nil, zero value otherwise.

### GetTargetDNOk

`func (o *DnMapperPluginResponse) GetTargetDNOk() (*string, bool)`

GetTargetDNOk returns a tuple with the TargetDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDN

`func (o *DnMapperPluginResponse) SetTargetDN(v string)`

SetTargetDN sets TargetDN field to given value.


### GetEnableAttributeMapping

`func (o *DnMapperPluginResponse) GetEnableAttributeMapping() bool`

GetEnableAttributeMapping returns the EnableAttributeMapping field if non-nil, zero value otherwise.

### GetEnableAttributeMappingOk

`func (o *DnMapperPluginResponse) GetEnableAttributeMappingOk() (*bool, bool)`

GetEnableAttributeMappingOk returns a tuple with the EnableAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAttributeMapping

`func (o *DnMapperPluginResponse) SetEnableAttributeMapping(v bool)`

SetEnableAttributeMapping sets EnableAttributeMapping field to given value.


### GetMapAttribute

`func (o *DnMapperPluginResponse) GetMapAttribute() []string`

GetMapAttribute returns the MapAttribute field if non-nil, zero value otherwise.

### GetMapAttributeOk

`func (o *DnMapperPluginResponse) GetMapAttributeOk() (*[]string, bool)`

GetMapAttributeOk returns a tuple with the MapAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAttribute

`func (o *DnMapperPluginResponse) SetMapAttribute(v []string)`

SetMapAttribute sets MapAttribute field to given value.

### HasMapAttribute

`func (o *DnMapperPluginResponse) HasMapAttribute() bool`

HasMapAttribute returns a boolean if a field has been set.

### GetEnableControlMapping

`func (o *DnMapperPluginResponse) GetEnableControlMapping() bool`

GetEnableControlMapping returns the EnableControlMapping field if non-nil, zero value otherwise.

### GetEnableControlMappingOk

`func (o *DnMapperPluginResponse) GetEnableControlMappingOk() (*bool, bool)`

GetEnableControlMappingOk returns a tuple with the EnableControlMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableControlMapping

`func (o *DnMapperPluginResponse) SetEnableControlMapping(v bool)`

SetEnableControlMapping sets EnableControlMapping field to given value.


### GetAlwaysMapResponses

`func (o *DnMapperPluginResponse) GetAlwaysMapResponses() bool`

GetAlwaysMapResponses returns the AlwaysMapResponses field if non-nil, zero value otherwise.

### GetAlwaysMapResponsesOk

`func (o *DnMapperPluginResponse) GetAlwaysMapResponsesOk() (*bool, bool)`

GetAlwaysMapResponsesOk returns a tuple with the AlwaysMapResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysMapResponses

`func (o *DnMapperPluginResponse) SetAlwaysMapResponses(v bool)`

SetAlwaysMapResponses sets AlwaysMapResponses field to given value.


### GetDescription

`func (o *DnMapperPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DnMapperPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DnMapperPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DnMapperPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DnMapperPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DnMapperPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DnMapperPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *DnMapperPluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *DnMapperPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *DnMapperPluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *DnMapperPluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetMeta

`func (o *DnMapperPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DnMapperPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DnMapperPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DnMapperPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *DnMapperPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *DnMapperPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *DnMapperPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *DnMapperPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *DnMapperPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnMapperPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnMapperPluginResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


