# JsonLogFieldSyntaxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumjsonLogFieldSyntaxSchemaUrn**](EnumjsonLogFieldSyntaxSchemaUrn.md) |  | 
**Id** | **string** | Name of the Log Field Syntax | 
**IncludedSensitiveField** | Pointer to **[]string** | The names of the JSON fields that will be considered sensitive. | [optional] 
**ExcludedSensitiveField** | Pointer to **[]string** | The names of the JSON fields that will not be considered sensitive. | [optional] 
**Description** | Pointer to **string** | A description for this Log Field Syntax | [optional] 
**DefaultBehavior** | Pointer to [**EnumlogFieldSyntaxDefaultBehaviorProp**](EnumlogFieldSyntaxDefaultBehaviorProp.md) |  | [optional] 

## Methods

### NewJsonLogFieldSyntaxResponse

`func NewJsonLogFieldSyntaxResponse(schemas []EnumjsonLogFieldSyntaxSchemaUrn, id string, ) *JsonLogFieldSyntaxResponse`

NewJsonLogFieldSyntaxResponse instantiates a new JsonLogFieldSyntaxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonLogFieldSyntaxResponseWithDefaults

`func NewJsonLogFieldSyntaxResponseWithDefaults() *JsonLogFieldSyntaxResponse`

NewJsonLogFieldSyntaxResponseWithDefaults instantiates a new JsonLogFieldSyntaxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *JsonLogFieldSyntaxResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *JsonLogFieldSyntaxResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *JsonLogFieldSyntaxResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *JsonLogFieldSyntaxResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *JsonLogFieldSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *JsonLogFieldSyntaxResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *JsonLogFieldSyntaxResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *JsonLogFieldSyntaxResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *JsonLogFieldSyntaxResponse) GetSchemas() []EnumjsonLogFieldSyntaxSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *JsonLogFieldSyntaxResponse) GetSchemasOk() (*[]EnumjsonLogFieldSyntaxSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *JsonLogFieldSyntaxResponse) SetSchemas(v []EnumjsonLogFieldSyntaxSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *JsonLogFieldSyntaxResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JsonLogFieldSyntaxResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JsonLogFieldSyntaxResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIncludedSensitiveField

`func (o *JsonLogFieldSyntaxResponse) GetIncludedSensitiveField() []string`

GetIncludedSensitiveField returns the IncludedSensitiveField field if non-nil, zero value otherwise.

### GetIncludedSensitiveFieldOk

`func (o *JsonLogFieldSyntaxResponse) GetIncludedSensitiveFieldOk() (*[]string, bool)`

GetIncludedSensitiveFieldOk returns a tuple with the IncludedSensitiveField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSensitiveField

`func (o *JsonLogFieldSyntaxResponse) SetIncludedSensitiveField(v []string)`

SetIncludedSensitiveField sets IncludedSensitiveField field to given value.

### HasIncludedSensitiveField

`func (o *JsonLogFieldSyntaxResponse) HasIncludedSensitiveField() bool`

HasIncludedSensitiveField returns a boolean if a field has been set.

### GetExcludedSensitiveField

`func (o *JsonLogFieldSyntaxResponse) GetExcludedSensitiveField() []string`

GetExcludedSensitiveField returns the ExcludedSensitiveField field if non-nil, zero value otherwise.

### GetExcludedSensitiveFieldOk

`func (o *JsonLogFieldSyntaxResponse) GetExcludedSensitiveFieldOk() (*[]string, bool)`

GetExcludedSensitiveFieldOk returns a tuple with the ExcludedSensitiveField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSensitiveField

`func (o *JsonLogFieldSyntaxResponse) SetExcludedSensitiveField(v []string)`

SetExcludedSensitiveField sets ExcludedSensitiveField field to given value.

### HasExcludedSensitiveField

`func (o *JsonLogFieldSyntaxResponse) HasExcludedSensitiveField() bool`

HasExcludedSensitiveField returns a boolean if a field has been set.

### GetDescription

`func (o *JsonLogFieldSyntaxResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JsonLogFieldSyntaxResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JsonLogFieldSyntaxResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JsonLogFieldSyntaxResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultBehavior

`func (o *JsonLogFieldSyntaxResponse) GetDefaultBehavior() EnumlogFieldSyntaxDefaultBehaviorProp`

GetDefaultBehavior returns the DefaultBehavior field if non-nil, zero value otherwise.

### GetDefaultBehaviorOk

`func (o *JsonLogFieldSyntaxResponse) GetDefaultBehaviorOk() (*EnumlogFieldSyntaxDefaultBehaviorProp, bool)`

GetDefaultBehaviorOk returns a tuple with the DefaultBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBehavior

`func (o *JsonLogFieldSyntaxResponse) SetDefaultBehavior(v EnumlogFieldSyntaxDefaultBehaviorProp)`

SetDefaultBehavior sets DefaultBehavior field to given value.

### HasDefaultBehavior

`func (o *JsonLogFieldSyntaxResponse) HasDefaultBehavior() bool`

HasDefaultBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


