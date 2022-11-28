# AttributeBasedLogFieldSyntaxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumattributeBasedLogFieldSyntaxSchemaUrn**](EnumattributeBasedLogFieldSyntaxSchemaUrn.md) |  | 
**Id** | Pointer to **string** | Name of the Log Field Syntax | [optional] 
**IncludedSensitiveAttribute** | Pointer to **[]string** | The set of attribute types that will be considered sensitive. | [optional] 
**ExcludedSensitiveAttribute** | Pointer to **[]string** | The set of attribute types that will not be considered sensitive. | [optional] 
**Description** | Pointer to **string** | A description for this Log Field Syntax | [optional] 
**DefaultBehavior** | Pointer to [**EnumlogFieldSyntaxDefaultBehaviorProp**](EnumlogFieldSyntaxDefaultBehaviorProp.md) |  | [optional] 

## Methods

### NewAttributeBasedLogFieldSyntaxResponse

`func NewAttributeBasedLogFieldSyntaxResponse(schemas []EnumattributeBasedLogFieldSyntaxSchemaUrn, ) *AttributeBasedLogFieldSyntaxResponse`

NewAttributeBasedLogFieldSyntaxResponse instantiates a new AttributeBasedLogFieldSyntaxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeBasedLogFieldSyntaxResponseWithDefaults

`func NewAttributeBasedLogFieldSyntaxResponseWithDefaults() *AttributeBasedLogFieldSyntaxResponse`

NewAttributeBasedLogFieldSyntaxResponseWithDefaults instantiates a new AttributeBasedLogFieldSyntaxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *AttributeBasedLogFieldSyntaxResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AttributeBasedLogFieldSyntaxResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AttributeBasedLogFieldSyntaxResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AttributeBasedLogFieldSyntaxResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AttributeBasedLogFieldSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AttributeBasedLogFieldSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AttributeBasedLogFieldSyntaxResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AttributeBasedLogFieldSyntaxResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *AttributeBasedLogFieldSyntaxResponse) GetSchemas() []EnumattributeBasedLogFieldSyntaxSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AttributeBasedLogFieldSyntaxResponse) GetSchemasOk() (*[]EnumattributeBasedLogFieldSyntaxSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AttributeBasedLogFieldSyntaxResponse) SetSchemas(v []EnumattributeBasedLogFieldSyntaxSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *AttributeBasedLogFieldSyntaxResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttributeBasedLogFieldSyntaxResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttributeBasedLogFieldSyntaxResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AttributeBasedLogFieldSyntaxResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncludedSensitiveAttribute

`func (o *AttributeBasedLogFieldSyntaxResponse) GetIncludedSensitiveAttribute() []string`

GetIncludedSensitiveAttribute returns the IncludedSensitiveAttribute field if non-nil, zero value otherwise.

### GetIncludedSensitiveAttributeOk

`func (o *AttributeBasedLogFieldSyntaxResponse) GetIncludedSensitiveAttributeOk() (*[]string, bool)`

GetIncludedSensitiveAttributeOk returns a tuple with the IncludedSensitiveAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSensitiveAttribute

`func (o *AttributeBasedLogFieldSyntaxResponse) SetIncludedSensitiveAttribute(v []string)`

SetIncludedSensitiveAttribute sets IncludedSensitiveAttribute field to given value.

### HasIncludedSensitiveAttribute

`func (o *AttributeBasedLogFieldSyntaxResponse) HasIncludedSensitiveAttribute() bool`

HasIncludedSensitiveAttribute returns a boolean if a field has been set.

### GetExcludedSensitiveAttribute

`func (o *AttributeBasedLogFieldSyntaxResponse) GetExcludedSensitiveAttribute() []string`

GetExcludedSensitiveAttribute returns the ExcludedSensitiveAttribute field if non-nil, zero value otherwise.

### GetExcludedSensitiveAttributeOk

`func (o *AttributeBasedLogFieldSyntaxResponse) GetExcludedSensitiveAttributeOk() (*[]string, bool)`

GetExcludedSensitiveAttributeOk returns a tuple with the ExcludedSensitiveAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSensitiveAttribute

`func (o *AttributeBasedLogFieldSyntaxResponse) SetExcludedSensitiveAttribute(v []string)`

SetExcludedSensitiveAttribute sets ExcludedSensitiveAttribute field to given value.

### HasExcludedSensitiveAttribute

`func (o *AttributeBasedLogFieldSyntaxResponse) HasExcludedSensitiveAttribute() bool`

HasExcludedSensitiveAttribute returns a boolean if a field has been set.

### GetDescription

`func (o *AttributeBasedLogFieldSyntaxResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AttributeBasedLogFieldSyntaxResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AttributeBasedLogFieldSyntaxResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AttributeBasedLogFieldSyntaxResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultBehavior

`func (o *AttributeBasedLogFieldSyntaxResponse) GetDefaultBehavior() EnumlogFieldSyntaxDefaultBehaviorProp`

GetDefaultBehavior returns the DefaultBehavior field if non-nil, zero value otherwise.

### GetDefaultBehaviorOk

`func (o *AttributeBasedLogFieldSyntaxResponse) GetDefaultBehaviorOk() (*EnumlogFieldSyntaxDefaultBehaviorProp, bool)`

GetDefaultBehaviorOk returns a tuple with the DefaultBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBehavior

`func (o *AttributeBasedLogFieldSyntaxResponse) SetDefaultBehavior(v EnumlogFieldSyntaxDefaultBehaviorProp)`

SetDefaultBehavior sets DefaultBehavior field to given value.

### HasDefaultBehavior

`func (o *AttributeBasedLogFieldSyntaxResponse) HasDefaultBehavior() bool`

HasDefaultBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


