# GetLogFieldSyntax200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumattributeBasedLogFieldSyntaxSchemaUrn**](EnumattributeBasedLogFieldSyntaxSchemaUrn.md) |  | 
**Id** | Pointer to **string** | Name of the Log Field Syntax | [optional] 
**IncludedSensitiveField** | Pointer to **[]string** | The names of the JSON fields that will be considered sensitive. | [optional] 
**ExcludedSensitiveField** | Pointer to **[]string** | The names of the JSON fields that will not be considered sensitive. | [optional] 
**Description** | Pointer to **string** | A description for this Log Field Syntax | [optional] 
**DefaultBehavior** | Pointer to [**EnumlogFieldSyntaxDefaultBehaviorProp**](EnumlogFieldSyntaxDefaultBehaviorProp.md) |  | [optional] 
**IncludedSensitiveAttribute** | Pointer to **[]string** | The set of attribute types that will be considered sensitive. | [optional] 
**ExcludedSensitiveAttribute** | Pointer to **[]string** | The set of attribute types that will not be considered sensitive. | [optional] 

## Methods

### NewGetLogFieldSyntax200Response

`func NewGetLogFieldSyntax200Response(schemas []EnumattributeBasedLogFieldSyntaxSchemaUrn, ) *GetLogFieldSyntax200Response`

NewGetLogFieldSyntax200Response instantiates a new GetLogFieldSyntax200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLogFieldSyntax200ResponseWithDefaults

`func NewGetLogFieldSyntax200ResponseWithDefaults() *GetLogFieldSyntax200Response`

NewGetLogFieldSyntax200ResponseWithDefaults instantiates a new GetLogFieldSyntax200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetLogFieldSyntax200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetLogFieldSyntax200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetLogFieldSyntax200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetLogFieldSyntax200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GetLogFieldSyntax200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GetLogFieldSyntax200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GetLogFieldSyntax200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GetLogFieldSyntax200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *GetLogFieldSyntax200Response) GetSchemas() []EnumattributeBasedLogFieldSyntaxSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetLogFieldSyntax200Response) GetSchemasOk() (*[]EnumattributeBasedLogFieldSyntaxSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetLogFieldSyntax200Response) SetSchemas(v []EnumattributeBasedLogFieldSyntaxSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GetLogFieldSyntax200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLogFieldSyntax200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLogFieldSyntax200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetLogFieldSyntax200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncludedSensitiveField

`func (o *GetLogFieldSyntax200Response) GetIncludedSensitiveField() []string`

GetIncludedSensitiveField returns the IncludedSensitiveField field if non-nil, zero value otherwise.

### GetIncludedSensitiveFieldOk

`func (o *GetLogFieldSyntax200Response) GetIncludedSensitiveFieldOk() (*[]string, bool)`

GetIncludedSensitiveFieldOk returns a tuple with the IncludedSensitiveField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSensitiveField

`func (o *GetLogFieldSyntax200Response) SetIncludedSensitiveField(v []string)`

SetIncludedSensitiveField sets IncludedSensitiveField field to given value.

### HasIncludedSensitiveField

`func (o *GetLogFieldSyntax200Response) HasIncludedSensitiveField() bool`

HasIncludedSensitiveField returns a boolean if a field has been set.

### GetExcludedSensitiveField

`func (o *GetLogFieldSyntax200Response) GetExcludedSensitiveField() []string`

GetExcludedSensitiveField returns the ExcludedSensitiveField field if non-nil, zero value otherwise.

### GetExcludedSensitiveFieldOk

`func (o *GetLogFieldSyntax200Response) GetExcludedSensitiveFieldOk() (*[]string, bool)`

GetExcludedSensitiveFieldOk returns a tuple with the ExcludedSensitiveField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSensitiveField

`func (o *GetLogFieldSyntax200Response) SetExcludedSensitiveField(v []string)`

SetExcludedSensitiveField sets ExcludedSensitiveField field to given value.

### HasExcludedSensitiveField

`func (o *GetLogFieldSyntax200Response) HasExcludedSensitiveField() bool`

HasExcludedSensitiveField returns a boolean if a field has been set.

### GetDescription

`func (o *GetLogFieldSyntax200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetLogFieldSyntax200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetLogFieldSyntax200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetLogFieldSyntax200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultBehavior

`func (o *GetLogFieldSyntax200Response) GetDefaultBehavior() EnumlogFieldSyntaxDefaultBehaviorProp`

GetDefaultBehavior returns the DefaultBehavior field if non-nil, zero value otherwise.

### GetDefaultBehaviorOk

`func (o *GetLogFieldSyntax200Response) GetDefaultBehaviorOk() (*EnumlogFieldSyntaxDefaultBehaviorProp, bool)`

GetDefaultBehaviorOk returns a tuple with the DefaultBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBehavior

`func (o *GetLogFieldSyntax200Response) SetDefaultBehavior(v EnumlogFieldSyntaxDefaultBehaviorProp)`

SetDefaultBehavior sets DefaultBehavior field to given value.

### HasDefaultBehavior

`func (o *GetLogFieldSyntax200Response) HasDefaultBehavior() bool`

HasDefaultBehavior returns a boolean if a field has been set.

### GetIncludedSensitiveAttribute

`func (o *GetLogFieldSyntax200Response) GetIncludedSensitiveAttribute() []string`

GetIncludedSensitiveAttribute returns the IncludedSensitiveAttribute field if non-nil, zero value otherwise.

### GetIncludedSensitiveAttributeOk

`func (o *GetLogFieldSyntax200Response) GetIncludedSensitiveAttributeOk() (*[]string, bool)`

GetIncludedSensitiveAttributeOk returns a tuple with the IncludedSensitiveAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSensitiveAttribute

`func (o *GetLogFieldSyntax200Response) SetIncludedSensitiveAttribute(v []string)`

SetIncludedSensitiveAttribute sets IncludedSensitiveAttribute field to given value.

### HasIncludedSensitiveAttribute

`func (o *GetLogFieldSyntax200Response) HasIncludedSensitiveAttribute() bool`

HasIncludedSensitiveAttribute returns a boolean if a field has been set.

### GetExcludedSensitiveAttribute

`func (o *GetLogFieldSyntax200Response) GetExcludedSensitiveAttribute() []string`

GetExcludedSensitiveAttribute returns the ExcludedSensitiveAttribute field if non-nil, zero value otherwise.

### GetExcludedSensitiveAttributeOk

`func (o *GetLogFieldSyntax200Response) GetExcludedSensitiveAttributeOk() (*[]string, bool)`

GetExcludedSensitiveAttributeOk returns a tuple with the ExcludedSensitiveAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSensitiveAttribute

`func (o *GetLogFieldSyntax200Response) SetExcludedSensitiveAttribute(v []string)`

SetExcludedSensitiveAttribute sets ExcludedSensitiveAttribute field to given value.

### HasExcludedSensitiveAttribute

`func (o *GetLogFieldSyntax200Response) HasExcludedSensitiveAttribute() bool`

HasExcludedSensitiveAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


