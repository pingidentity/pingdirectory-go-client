# AddReferentialIntegrityPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginName** | **string** | Name of the new Plugin | 
**Schemas** | [**[]EnumreferentialIntegrityPluginSchemaUrn**](EnumreferentialIntegrityPluginSchemaUrn.md) |  | 
**PluginType** | [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) |  | 
**AttributeType** | **[]string** | Specifies the attribute types for which referential integrity is to be maintained. | 
**BaseDN** | Pointer to **[]string** | Specifies the base DN that limits the scope within which referential integrity is maintained. | [optional] 
**LogFile** | Pointer to **string** | Specifies the log file location where the update records are written when the plug-in is in background-mode processing. | [optional] 
**UpdateInterval** | Pointer to **string** | Specifies the interval in seconds when referential integrity updates are made. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 

## Methods

### NewAddReferentialIntegrityPluginRequest

`func NewAddReferentialIntegrityPluginRequest(pluginName string, schemas []EnumreferentialIntegrityPluginSchemaUrn, pluginType []EnumpluginPluginTypeProp, attributeType []string, enabled bool, ) *AddReferentialIntegrityPluginRequest`

NewAddReferentialIntegrityPluginRequest instantiates a new AddReferentialIntegrityPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddReferentialIntegrityPluginRequestWithDefaults

`func NewAddReferentialIntegrityPluginRequestWithDefaults() *AddReferentialIntegrityPluginRequest`

NewAddReferentialIntegrityPluginRequestWithDefaults instantiates a new AddReferentialIntegrityPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginName

`func (o *AddReferentialIntegrityPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddReferentialIntegrityPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddReferentialIntegrityPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### GetSchemas

`func (o *AddReferentialIntegrityPluginRequest) GetSchemas() []EnumreferentialIntegrityPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddReferentialIntegrityPluginRequest) GetSchemasOk() (*[]EnumreferentialIntegrityPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddReferentialIntegrityPluginRequest) SetSchemas(v []EnumreferentialIntegrityPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPluginType

`func (o *AddReferentialIntegrityPluginRequest) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *AddReferentialIntegrityPluginRequest) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *AddReferentialIntegrityPluginRequest) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.


### GetAttributeType

`func (o *AddReferentialIntegrityPluginRequest) GetAttributeType() []string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddReferentialIntegrityPluginRequest) GetAttributeTypeOk() (*[]string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddReferentialIntegrityPluginRequest) SetAttributeType(v []string)`

SetAttributeType sets AttributeType field to given value.


### GetBaseDN

`func (o *AddReferentialIntegrityPluginRequest) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddReferentialIntegrityPluginRequest) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddReferentialIntegrityPluginRequest) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *AddReferentialIntegrityPluginRequest) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetLogFile

`func (o *AddReferentialIntegrityPluginRequest) GetLogFile() string`

GetLogFile returns the LogFile field if non-nil, zero value otherwise.

### GetLogFileOk

`func (o *AddReferentialIntegrityPluginRequest) GetLogFileOk() (*string, bool)`

GetLogFileOk returns a tuple with the LogFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFile

`func (o *AddReferentialIntegrityPluginRequest) SetLogFile(v string)`

SetLogFile sets LogFile field to given value.

### HasLogFile

`func (o *AddReferentialIntegrityPluginRequest) HasLogFile() bool`

HasLogFile returns a boolean if a field has been set.

### GetUpdateInterval

`func (o *AddReferentialIntegrityPluginRequest) GetUpdateInterval() string`

GetUpdateInterval returns the UpdateInterval field if non-nil, zero value otherwise.

### GetUpdateIntervalOk

`func (o *AddReferentialIntegrityPluginRequest) GetUpdateIntervalOk() (*string, bool)`

GetUpdateIntervalOk returns a tuple with the UpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInterval

`func (o *AddReferentialIntegrityPluginRequest) SetUpdateInterval(v string)`

SetUpdateInterval sets UpdateInterval field to given value.

### HasUpdateInterval

`func (o *AddReferentialIntegrityPluginRequest) HasUpdateInterval() bool`

HasUpdateInterval returns a boolean if a field has been set.

### GetDescription

`func (o *AddReferentialIntegrityPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddReferentialIntegrityPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddReferentialIntegrityPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddReferentialIntegrityPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddReferentialIntegrityPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddReferentialIntegrityPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddReferentialIntegrityPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *AddReferentialIntegrityPluginRequest) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *AddReferentialIntegrityPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *AddReferentialIntegrityPluginRequest) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *AddReferentialIntegrityPluginRequest) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


