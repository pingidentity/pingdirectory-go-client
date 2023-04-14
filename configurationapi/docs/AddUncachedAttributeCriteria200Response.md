# AddUncachedAttributeCriteria200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Uncached Attribute Criteria | 
**Schemas** | [**[]EnumthirdPartyUncachedAttributeCriteriaSchemaUrn**](EnumthirdPartyUncachedAttributeCriteriaSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Uncached Attribute Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Attribute Criteria is enabled for use in the server. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Uncached Attribute Criteria. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted Uncached Attribute Criteria. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**AttributeType** | **[]string** | Specifies the attribute types for attributes that may be written to the uncached-id2entry database. | 
**MinValueCount** | Pointer to **int64** | Specifies the minimum number of values that an attribute must have before it will be written into the uncached-id2entry database. | [optional] 
**MinTotalValueSize** | Pointer to **string** | Specifies the minimum total value size (i.e., the sum of the sizes of all values) that an attribute must have before it will be written into the uncached-id2entry database. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Uncached Attribute Criteria. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Uncached Attribute Criteria. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddUncachedAttributeCriteria200Response

`func NewAddUncachedAttributeCriteria200Response(id string, schemas []EnumthirdPartyUncachedAttributeCriteriaSchemaUrn, enabled bool, scriptClass string, attributeType []string, extensionClass string, ) *AddUncachedAttributeCriteria200Response`

NewAddUncachedAttributeCriteria200Response instantiates a new AddUncachedAttributeCriteria200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUncachedAttributeCriteria200ResponseWithDefaults

`func NewAddUncachedAttributeCriteria200ResponseWithDefaults() *AddUncachedAttributeCriteria200Response`

NewAddUncachedAttributeCriteria200ResponseWithDefaults instantiates a new AddUncachedAttributeCriteria200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddUncachedAttributeCriteria200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddUncachedAttributeCriteria200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddUncachedAttributeCriteria200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddUncachedAttributeCriteria200Response) GetSchemas() []EnumthirdPartyUncachedAttributeCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddUncachedAttributeCriteria200Response) GetSchemasOk() (*[]EnumthirdPartyUncachedAttributeCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddUncachedAttributeCriteria200Response) SetSchemas(v []EnumthirdPartyUncachedAttributeCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddUncachedAttributeCriteria200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddUncachedAttributeCriteria200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddUncachedAttributeCriteria200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddUncachedAttributeCriteria200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddUncachedAttributeCriteria200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddUncachedAttributeCriteria200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddUncachedAttributeCriteria200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *AddUncachedAttributeCriteria200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddUncachedAttributeCriteria200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddUncachedAttributeCriteria200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddUncachedAttributeCriteria200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddUncachedAttributeCriteria200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddUncachedAttributeCriteria200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddUncachedAttributeCriteria200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddUncachedAttributeCriteria200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetScriptClass

`func (o *AddUncachedAttributeCriteria200Response) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddUncachedAttributeCriteria200Response) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddUncachedAttributeCriteria200Response) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddUncachedAttributeCriteria200Response) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddUncachedAttributeCriteria200Response) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddUncachedAttributeCriteria200Response) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddUncachedAttributeCriteria200Response) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetAttributeType

`func (o *AddUncachedAttributeCriteria200Response) GetAttributeType() []string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddUncachedAttributeCriteria200Response) GetAttributeTypeOk() (*[]string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddUncachedAttributeCriteria200Response) SetAttributeType(v []string)`

SetAttributeType sets AttributeType field to given value.


### GetMinValueCount

`func (o *AddUncachedAttributeCriteria200Response) GetMinValueCount() int64`

GetMinValueCount returns the MinValueCount field if non-nil, zero value otherwise.

### GetMinValueCountOk

`func (o *AddUncachedAttributeCriteria200Response) GetMinValueCountOk() (*int64, bool)`

GetMinValueCountOk returns a tuple with the MinValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValueCount

`func (o *AddUncachedAttributeCriteria200Response) SetMinValueCount(v int64)`

SetMinValueCount sets MinValueCount field to given value.

### HasMinValueCount

`func (o *AddUncachedAttributeCriteria200Response) HasMinValueCount() bool`

HasMinValueCount returns a boolean if a field has been set.

### GetMinTotalValueSize

`func (o *AddUncachedAttributeCriteria200Response) GetMinTotalValueSize() string`

GetMinTotalValueSize returns the MinTotalValueSize field if non-nil, zero value otherwise.

### GetMinTotalValueSizeOk

`func (o *AddUncachedAttributeCriteria200Response) GetMinTotalValueSizeOk() (*string, bool)`

GetMinTotalValueSizeOk returns a tuple with the MinTotalValueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTotalValueSize

`func (o *AddUncachedAttributeCriteria200Response) SetMinTotalValueSize(v string)`

SetMinTotalValueSize sets MinTotalValueSize field to given value.

### HasMinTotalValueSize

`func (o *AddUncachedAttributeCriteria200Response) HasMinTotalValueSize() bool`

HasMinTotalValueSize returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddUncachedAttributeCriteria200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddUncachedAttributeCriteria200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddUncachedAttributeCriteria200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddUncachedAttributeCriteria200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddUncachedAttributeCriteria200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddUncachedAttributeCriteria200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddUncachedAttributeCriteria200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


