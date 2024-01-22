# AddSevenBitCleanPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsevenBitCleanPluginSchemaUrn**](EnumsevenBitCleanPluginSchemaUrn.md) |  | 
**PluginType** | Pointer to [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) |  | [optional] 
**AttributeType** | Pointer to **[]string** | Specifies the name or OID of an attribute type for which values should be checked to ensure that they are 7-bit clean. | [optional] 
**BaseDN** | Pointer to **[]string** | Specifies the base DN below which the checking is performed. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**PluginName** | **string** | Name of the new Plugin | 

## Methods

### NewAddSevenBitCleanPluginRequest

`func NewAddSevenBitCleanPluginRequest(schemas []EnumsevenBitCleanPluginSchemaUrn, enabled bool, pluginName string, ) *AddSevenBitCleanPluginRequest`

NewAddSevenBitCleanPluginRequest instantiates a new AddSevenBitCleanPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSevenBitCleanPluginRequestWithDefaults

`func NewAddSevenBitCleanPluginRequestWithDefaults() *AddSevenBitCleanPluginRequest`

NewAddSevenBitCleanPluginRequestWithDefaults instantiates a new AddSevenBitCleanPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddSevenBitCleanPluginRequest) GetSchemas() []EnumsevenBitCleanPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSevenBitCleanPluginRequest) GetSchemasOk() (*[]EnumsevenBitCleanPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSevenBitCleanPluginRequest) SetSchemas(v []EnumsevenBitCleanPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPluginType

`func (o *AddSevenBitCleanPluginRequest) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *AddSevenBitCleanPluginRequest) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *AddSevenBitCleanPluginRequest) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.

### HasPluginType

`func (o *AddSevenBitCleanPluginRequest) HasPluginType() bool`

HasPluginType returns a boolean if a field has been set.

### GetAttributeType

`func (o *AddSevenBitCleanPluginRequest) GetAttributeType() []string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddSevenBitCleanPluginRequest) GetAttributeTypeOk() (*[]string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddSevenBitCleanPluginRequest) SetAttributeType(v []string)`

SetAttributeType sets AttributeType field to given value.

### HasAttributeType

`func (o *AddSevenBitCleanPluginRequest) HasAttributeType() bool`

HasAttributeType returns a boolean if a field has been set.

### GetBaseDN

`func (o *AddSevenBitCleanPluginRequest) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddSevenBitCleanPluginRequest) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddSevenBitCleanPluginRequest) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *AddSevenBitCleanPluginRequest) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetDescription

`func (o *AddSevenBitCleanPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSevenBitCleanPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSevenBitCleanPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSevenBitCleanPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddSevenBitCleanPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddSevenBitCleanPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddSevenBitCleanPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *AddSevenBitCleanPluginRequest) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *AddSevenBitCleanPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *AddSevenBitCleanPluginRequest) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *AddSevenBitCleanPluginRequest) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetPluginName

`func (o *AddSevenBitCleanPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddSevenBitCleanPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddSevenBitCleanPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


