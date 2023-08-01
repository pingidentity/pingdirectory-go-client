# GlobalReferentialIntegrityPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumglobalReferentialIntegrityPluginSchemaUrn**](EnumglobalReferentialIntegrityPluginSchemaUrn.md) |  | 
**Id** | **string** | Name of the Plugin | 
**AttributeType** | **[]string** | The attribute type(s) for which to maintain referential integrity. The attribute must have a distinguished name or a name and optional UID syntax and must be indexed for equality searches in all subtree views for which referential integrity is to be maintained. | 
**SubtreeView** | **[]string** | The subtree view(s) for which to maintain referential integrity. | 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewGlobalReferentialIntegrityPluginResponse

`func NewGlobalReferentialIntegrityPluginResponse(schemas []EnumglobalReferentialIntegrityPluginSchemaUrn, id string, attributeType []string, subtreeView []string, enabled bool, ) *GlobalReferentialIntegrityPluginResponse`

NewGlobalReferentialIntegrityPluginResponse instantiates a new GlobalReferentialIntegrityPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalReferentialIntegrityPluginResponseWithDefaults

`func NewGlobalReferentialIntegrityPluginResponseWithDefaults() *GlobalReferentialIntegrityPluginResponse`

NewGlobalReferentialIntegrityPluginResponseWithDefaults instantiates a new GlobalReferentialIntegrityPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GlobalReferentialIntegrityPluginResponse) GetSchemas() []EnumglobalReferentialIntegrityPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GlobalReferentialIntegrityPluginResponse) GetSchemasOk() (*[]EnumglobalReferentialIntegrityPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GlobalReferentialIntegrityPluginResponse) SetSchemas(v []EnumglobalReferentialIntegrityPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GlobalReferentialIntegrityPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalReferentialIntegrityPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalReferentialIntegrityPluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAttributeType

`func (o *GlobalReferentialIntegrityPluginResponse) GetAttributeType() []string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *GlobalReferentialIntegrityPluginResponse) GetAttributeTypeOk() (*[]string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *GlobalReferentialIntegrityPluginResponse) SetAttributeType(v []string)`

SetAttributeType sets AttributeType field to given value.


### GetSubtreeView

`func (o *GlobalReferentialIntegrityPluginResponse) GetSubtreeView() []string`

GetSubtreeView returns the SubtreeView field if non-nil, zero value otherwise.

### GetSubtreeViewOk

`func (o *GlobalReferentialIntegrityPluginResponse) GetSubtreeViewOk() (*[]string, bool)`

GetSubtreeViewOk returns a tuple with the SubtreeView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtreeView

`func (o *GlobalReferentialIntegrityPluginResponse) SetSubtreeView(v []string)`

SetSubtreeView sets SubtreeView field to given value.


### GetDescription

`func (o *GlobalReferentialIntegrityPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GlobalReferentialIntegrityPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GlobalReferentialIntegrityPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GlobalReferentialIntegrityPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GlobalReferentialIntegrityPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GlobalReferentialIntegrityPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GlobalReferentialIntegrityPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *GlobalReferentialIntegrityPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GlobalReferentialIntegrityPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GlobalReferentialIntegrityPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GlobalReferentialIntegrityPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GlobalReferentialIntegrityPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GlobalReferentialIntegrityPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GlobalReferentialIntegrityPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GlobalReferentialIntegrityPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


