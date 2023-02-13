# ComposedAttributePluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Plugin | 
**Schemas** | [**[]EnumcomposedAttributePluginSchemaUrn**](EnumcomposedAttributePluginSchemaUrn.md) |  | 
**PluginType** | [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) |  | 
**AttributeType** | **string** | The name or OID of the attribute type for which values are to be generated. | 
**ValuePattern** | **[]string** | Specifies a pattern for constructing the values to use for the target attribute type. | 
**MultipleValuePatternBehavior** | Pointer to [**EnumpluginMultipleValuePatternBehaviorProp**](EnumpluginMultipleValuePatternBehaviorProp.md) |  | [optional] 
**MultiValuedAttributeBehavior** | Pointer to [**EnumpluginMultiValuedAttributeBehaviorProp**](EnumpluginMultiValuedAttributeBehaviorProp.md) |  | [optional] 
**TargetAttributeExistsDuringInitialPopulationBehavior** | Pointer to [**EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp**](EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp.md) |  | [optional] 
**UpdateSourceAttributeBehavior** | Pointer to [**EnumpluginUpdateSourceAttributeBehaviorProp**](EnumpluginUpdateSourceAttributeBehaviorProp.md) |  | [optional] 
**SourceAttributeRemovalBehavior** | Pointer to [**EnumpluginSourceAttributeRemovalBehaviorProp**](EnumpluginSourceAttributeRemovalBehaviorProp.md) |  | [optional] 
**UpdateTargetAttributeBehavior** | Pointer to [**EnumpluginUpdateTargetAttributeBehaviorProp**](EnumpluginUpdateTargetAttributeBehaviorProp.md) |  | [optional] 
**IncludeBaseDN** | Pointer to **[]string** | The set of base DNs below which composed values may be generated. | [optional] 
**ExcludeBaseDN** | Pointer to **[]string** | The set of base DNs below which composed values will not be generated. | [optional] 
**IncludeFilter** | Pointer to **[]string** | The set of search filters that identify entries for which composed values may be generated. | [optional] 
**ExcludeFilter** | Pointer to **[]string** | The set of search filters that identify entries for which composed values will not be generated. | [optional] 
**UpdatedEntryNewlyMatchesCriteriaBehavior** | Pointer to [**EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp**](EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp.md) |  | [optional] 
**UpdatedEntryNoLongerMatchesCriteriaBehavior** | Pointer to [**EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp**](EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 

## Methods

### NewComposedAttributePluginResponse

`func NewComposedAttributePluginResponse(id string, schemas []EnumcomposedAttributePluginSchemaUrn, pluginType []EnumpluginPluginTypeProp, attributeType string, valuePattern []string, enabled bool, ) *ComposedAttributePluginResponse`

NewComposedAttributePluginResponse instantiates a new ComposedAttributePluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComposedAttributePluginResponseWithDefaults

`func NewComposedAttributePluginResponseWithDefaults() *ComposedAttributePluginResponse`

NewComposedAttributePluginResponseWithDefaults instantiates a new ComposedAttributePluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ComposedAttributePluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ComposedAttributePluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ComposedAttributePluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ComposedAttributePluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ComposedAttributePluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ComposedAttributePluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ComposedAttributePluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ComposedAttributePluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *ComposedAttributePluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComposedAttributePluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComposedAttributePluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ComposedAttributePluginResponse) GetSchemas() []EnumcomposedAttributePluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ComposedAttributePluginResponse) GetSchemasOk() (*[]EnumcomposedAttributePluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ComposedAttributePluginResponse) SetSchemas(v []EnumcomposedAttributePluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPluginType

`func (o *ComposedAttributePluginResponse) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *ComposedAttributePluginResponse) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *ComposedAttributePluginResponse) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.


### GetAttributeType

`func (o *ComposedAttributePluginResponse) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *ComposedAttributePluginResponse) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *ComposedAttributePluginResponse) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetValuePattern

`func (o *ComposedAttributePluginResponse) GetValuePattern() []string`

GetValuePattern returns the ValuePattern field if non-nil, zero value otherwise.

### GetValuePatternOk

`func (o *ComposedAttributePluginResponse) GetValuePatternOk() (*[]string, bool)`

GetValuePatternOk returns a tuple with the ValuePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuePattern

`func (o *ComposedAttributePluginResponse) SetValuePattern(v []string)`

SetValuePattern sets ValuePattern field to given value.


### GetMultipleValuePatternBehavior

`func (o *ComposedAttributePluginResponse) GetMultipleValuePatternBehavior() EnumpluginMultipleValuePatternBehaviorProp`

GetMultipleValuePatternBehavior returns the MultipleValuePatternBehavior field if non-nil, zero value otherwise.

### GetMultipleValuePatternBehaviorOk

`func (o *ComposedAttributePluginResponse) GetMultipleValuePatternBehaviorOk() (*EnumpluginMultipleValuePatternBehaviorProp, bool)`

GetMultipleValuePatternBehaviorOk returns a tuple with the MultipleValuePatternBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleValuePatternBehavior

`func (o *ComposedAttributePluginResponse) SetMultipleValuePatternBehavior(v EnumpluginMultipleValuePatternBehaviorProp)`

SetMultipleValuePatternBehavior sets MultipleValuePatternBehavior field to given value.

### HasMultipleValuePatternBehavior

`func (o *ComposedAttributePluginResponse) HasMultipleValuePatternBehavior() bool`

HasMultipleValuePatternBehavior returns a boolean if a field has been set.

### GetMultiValuedAttributeBehavior

`func (o *ComposedAttributePluginResponse) GetMultiValuedAttributeBehavior() EnumpluginMultiValuedAttributeBehaviorProp`

GetMultiValuedAttributeBehavior returns the MultiValuedAttributeBehavior field if non-nil, zero value otherwise.

### GetMultiValuedAttributeBehaviorOk

`func (o *ComposedAttributePluginResponse) GetMultiValuedAttributeBehaviorOk() (*EnumpluginMultiValuedAttributeBehaviorProp, bool)`

GetMultiValuedAttributeBehaviorOk returns a tuple with the MultiValuedAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValuedAttributeBehavior

`func (o *ComposedAttributePluginResponse) SetMultiValuedAttributeBehavior(v EnumpluginMultiValuedAttributeBehaviorProp)`

SetMultiValuedAttributeBehavior sets MultiValuedAttributeBehavior field to given value.

### HasMultiValuedAttributeBehavior

`func (o *ComposedAttributePluginResponse) HasMultiValuedAttributeBehavior() bool`

HasMultiValuedAttributeBehavior returns a boolean if a field has been set.

### GetTargetAttributeExistsDuringInitialPopulationBehavior

`func (o *ComposedAttributePluginResponse) GetTargetAttributeExistsDuringInitialPopulationBehavior() EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp`

GetTargetAttributeExistsDuringInitialPopulationBehavior returns the TargetAttributeExistsDuringInitialPopulationBehavior field if non-nil, zero value otherwise.

### GetTargetAttributeExistsDuringInitialPopulationBehaviorOk

`func (o *ComposedAttributePluginResponse) GetTargetAttributeExistsDuringInitialPopulationBehaviorOk() (*EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp, bool)`

GetTargetAttributeExistsDuringInitialPopulationBehaviorOk returns a tuple with the TargetAttributeExistsDuringInitialPopulationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttributeExistsDuringInitialPopulationBehavior

`func (o *ComposedAttributePluginResponse) SetTargetAttributeExistsDuringInitialPopulationBehavior(v EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp)`

SetTargetAttributeExistsDuringInitialPopulationBehavior sets TargetAttributeExistsDuringInitialPopulationBehavior field to given value.

### HasTargetAttributeExistsDuringInitialPopulationBehavior

`func (o *ComposedAttributePluginResponse) HasTargetAttributeExistsDuringInitialPopulationBehavior() bool`

HasTargetAttributeExistsDuringInitialPopulationBehavior returns a boolean if a field has been set.

### GetUpdateSourceAttributeBehavior

`func (o *ComposedAttributePluginResponse) GetUpdateSourceAttributeBehavior() EnumpluginUpdateSourceAttributeBehaviorProp`

GetUpdateSourceAttributeBehavior returns the UpdateSourceAttributeBehavior field if non-nil, zero value otherwise.

### GetUpdateSourceAttributeBehaviorOk

`func (o *ComposedAttributePluginResponse) GetUpdateSourceAttributeBehaviorOk() (*EnumpluginUpdateSourceAttributeBehaviorProp, bool)`

GetUpdateSourceAttributeBehaviorOk returns a tuple with the UpdateSourceAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSourceAttributeBehavior

`func (o *ComposedAttributePluginResponse) SetUpdateSourceAttributeBehavior(v EnumpluginUpdateSourceAttributeBehaviorProp)`

SetUpdateSourceAttributeBehavior sets UpdateSourceAttributeBehavior field to given value.

### HasUpdateSourceAttributeBehavior

`func (o *ComposedAttributePluginResponse) HasUpdateSourceAttributeBehavior() bool`

HasUpdateSourceAttributeBehavior returns a boolean if a field has been set.

### GetSourceAttributeRemovalBehavior

`func (o *ComposedAttributePluginResponse) GetSourceAttributeRemovalBehavior() EnumpluginSourceAttributeRemovalBehaviorProp`

GetSourceAttributeRemovalBehavior returns the SourceAttributeRemovalBehavior field if non-nil, zero value otherwise.

### GetSourceAttributeRemovalBehaviorOk

`func (o *ComposedAttributePluginResponse) GetSourceAttributeRemovalBehaviorOk() (*EnumpluginSourceAttributeRemovalBehaviorProp, bool)`

GetSourceAttributeRemovalBehaviorOk returns a tuple with the SourceAttributeRemovalBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttributeRemovalBehavior

`func (o *ComposedAttributePluginResponse) SetSourceAttributeRemovalBehavior(v EnumpluginSourceAttributeRemovalBehaviorProp)`

SetSourceAttributeRemovalBehavior sets SourceAttributeRemovalBehavior field to given value.

### HasSourceAttributeRemovalBehavior

`func (o *ComposedAttributePluginResponse) HasSourceAttributeRemovalBehavior() bool`

HasSourceAttributeRemovalBehavior returns a boolean if a field has been set.

### GetUpdateTargetAttributeBehavior

`func (o *ComposedAttributePluginResponse) GetUpdateTargetAttributeBehavior() EnumpluginUpdateTargetAttributeBehaviorProp`

GetUpdateTargetAttributeBehavior returns the UpdateTargetAttributeBehavior field if non-nil, zero value otherwise.

### GetUpdateTargetAttributeBehaviorOk

`func (o *ComposedAttributePluginResponse) GetUpdateTargetAttributeBehaviorOk() (*EnumpluginUpdateTargetAttributeBehaviorProp, bool)`

GetUpdateTargetAttributeBehaviorOk returns a tuple with the UpdateTargetAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTargetAttributeBehavior

`func (o *ComposedAttributePluginResponse) SetUpdateTargetAttributeBehavior(v EnumpluginUpdateTargetAttributeBehaviorProp)`

SetUpdateTargetAttributeBehavior sets UpdateTargetAttributeBehavior field to given value.

### HasUpdateTargetAttributeBehavior

`func (o *ComposedAttributePluginResponse) HasUpdateTargetAttributeBehavior() bool`

HasUpdateTargetAttributeBehavior returns a boolean if a field has been set.

### GetIncludeBaseDN

`func (o *ComposedAttributePluginResponse) GetIncludeBaseDN() []string`

GetIncludeBaseDN returns the IncludeBaseDN field if non-nil, zero value otherwise.

### GetIncludeBaseDNOk

`func (o *ComposedAttributePluginResponse) GetIncludeBaseDNOk() (*[]string, bool)`

GetIncludeBaseDNOk returns a tuple with the IncludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBaseDN

`func (o *ComposedAttributePluginResponse) SetIncludeBaseDN(v []string)`

SetIncludeBaseDN sets IncludeBaseDN field to given value.

### HasIncludeBaseDN

`func (o *ComposedAttributePluginResponse) HasIncludeBaseDN() bool`

HasIncludeBaseDN returns a boolean if a field has been set.

### GetExcludeBaseDN

`func (o *ComposedAttributePluginResponse) GetExcludeBaseDN() []string`

GetExcludeBaseDN returns the ExcludeBaseDN field if non-nil, zero value otherwise.

### GetExcludeBaseDNOk

`func (o *ComposedAttributePluginResponse) GetExcludeBaseDNOk() (*[]string, bool)`

GetExcludeBaseDNOk returns a tuple with the ExcludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeBaseDN

`func (o *ComposedAttributePluginResponse) SetExcludeBaseDN(v []string)`

SetExcludeBaseDN sets ExcludeBaseDN field to given value.

### HasExcludeBaseDN

`func (o *ComposedAttributePluginResponse) HasExcludeBaseDN() bool`

HasExcludeBaseDN returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *ComposedAttributePluginResponse) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *ComposedAttributePluginResponse) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *ComposedAttributePluginResponse) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *ComposedAttributePluginResponse) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetExcludeFilter

`func (o *ComposedAttributePluginResponse) GetExcludeFilter() []string`

GetExcludeFilter returns the ExcludeFilter field if non-nil, zero value otherwise.

### GetExcludeFilterOk

`func (o *ComposedAttributePluginResponse) GetExcludeFilterOk() (*[]string, bool)`

GetExcludeFilterOk returns a tuple with the ExcludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilter

`func (o *ComposedAttributePluginResponse) SetExcludeFilter(v []string)`

SetExcludeFilter sets ExcludeFilter field to given value.

### HasExcludeFilter

`func (o *ComposedAttributePluginResponse) HasExcludeFilter() bool`

HasExcludeFilter returns a boolean if a field has been set.

### GetUpdatedEntryNewlyMatchesCriteriaBehavior

`func (o *ComposedAttributePluginResponse) GetUpdatedEntryNewlyMatchesCriteriaBehavior() EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp`

GetUpdatedEntryNewlyMatchesCriteriaBehavior returns the UpdatedEntryNewlyMatchesCriteriaBehavior field if non-nil, zero value otherwise.

### GetUpdatedEntryNewlyMatchesCriteriaBehaviorOk

`func (o *ComposedAttributePluginResponse) GetUpdatedEntryNewlyMatchesCriteriaBehaviorOk() (*EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp, bool)`

GetUpdatedEntryNewlyMatchesCriteriaBehaviorOk returns a tuple with the UpdatedEntryNewlyMatchesCriteriaBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedEntryNewlyMatchesCriteriaBehavior

`func (o *ComposedAttributePluginResponse) SetUpdatedEntryNewlyMatchesCriteriaBehavior(v EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp)`

SetUpdatedEntryNewlyMatchesCriteriaBehavior sets UpdatedEntryNewlyMatchesCriteriaBehavior field to given value.

### HasUpdatedEntryNewlyMatchesCriteriaBehavior

`func (o *ComposedAttributePluginResponse) HasUpdatedEntryNewlyMatchesCriteriaBehavior() bool`

HasUpdatedEntryNewlyMatchesCriteriaBehavior returns a boolean if a field has been set.

### GetUpdatedEntryNoLongerMatchesCriteriaBehavior

`func (o *ComposedAttributePluginResponse) GetUpdatedEntryNoLongerMatchesCriteriaBehavior() EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp`

GetUpdatedEntryNoLongerMatchesCriteriaBehavior returns the UpdatedEntryNoLongerMatchesCriteriaBehavior field if non-nil, zero value otherwise.

### GetUpdatedEntryNoLongerMatchesCriteriaBehaviorOk

`func (o *ComposedAttributePluginResponse) GetUpdatedEntryNoLongerMatchesCriteriaBehaviorOk() (*EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp, bool)`

GetUpdatedEntryNoLongerMatchesCriteriaBehaviorOk returns a tuple with the UpdatedEntryNoLongerMatchesCriteriaBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedEntryNoLongerMatchesCriteriaBehavior

`func (o *ComposedAttributePluginResponse) SetUpdatedEntryNoLongerMatchesCriteriaBehavior(v EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp)`

SetUpdatedEntryNoLongerMatchesCriteriaBehavior sets UpdatedEntryNoLongerMatchesCriteriaBehavior field to given value.

### HasUpdatedEntryNoLongerMatchesCriteriaBehavior

`func (o *ComposedAttributePluginResponse) HasUpdatedEntryNoLongerMatchesCriteriaBehavior() bool`

HasUpdatedEntryNoLongerMatchesCriteriaBehavior returns a boolean if a field has been set.

### GetDescription

`func (o *ComposedAttributePluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComposedAttributePluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComposedAttributePluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ComposedAttributePluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ComposedAttributePluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ComposedAttributePluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ComposedAttributePluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *ComposedAttributePluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *ComposedAttributePluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *ComposedAttributePluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *ComposedAttributePluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


