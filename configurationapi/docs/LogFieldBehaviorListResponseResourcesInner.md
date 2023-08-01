# LogFieldBehaviorListResponseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Log Field Behavior | 
**Schemas** | [**[]EnumtextAccessLogFieldBehaviorSchemaUrn**](EnumtextAccessLogFieldBehaviorSchemaUrn.md) |  | 
**PreserveField** | Pointer to [**[]EnumlogFieldBehaviorTextAccessPreserveFieldProp**](EnumlogFieldBehaviorTextAccessPreserveFieldProp.md) |  | [optional] 
**PreserveFieldName** | Pointer to **[]string** | The names of any custom fields whose values should be preserved. This should generally only be used for fields that are not available through the preserve-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**OmitField** | Pointer to [**[]EnumlogFieldBehaviorTextAccessOmitFieldProp**](EnumlogFieldBehaviorTextAccessOmitFieldProp.md) |  | [optional] 
**OmitFieldName** | Pointer to **[]string** | The names of any custom fields that should be omitted from log messages. This should generally only be used for fields that are not available through the omit-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**RedactEntireValueField** | Pointer to [**[]EnumlogFieldBehaviorTextAccessRedactEntireValueFieldProp**](EnumlogFieldBehaviorTextAccessRedactEntireValueFieldProp.md) |  | [optional] 
**RedactEntireValueFieldName** | Pointer to **[]string** | The names of any custom fields whose values should be completely redacted. This should generally only be used for fields that are not available through the redact-entire-value-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**RedactValueComponentsField** | Pointer to [**[]EnumlogFieldBehaviorTextAccessRedactValueComponentsFieldProp**](EnumlogFieldBehaviorTextAccessRedactValueComponentsFieldProp.md) |  | [optional] 
**RedactValueComponentsFieldName** | Pointer to **[]string** | The names of any custom fields for which to redact components within the value. This should generally only be used for fields that are not available through the redact-value-components-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**TokenizeEntireValueField** | Pointer to [**[]EnumlogFieldBehaviorTextAccessTokenizeEntireValueFieldProp**](EnumlogFieldBehaviorTextAccessTokenizeEntireValueFieldProp.md) |  | [optional] 
**TokenizeEntireValueFieldName** | Pointer to **[]string** | The names of any custom fields whose values should be completely tokenized. This should generally only be used for fields that are not available through the tokenize-entire-value-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**TokenizeValueComponentsField** | Pointer to [**[]EnumlogFieldBehaviorTextAccessTokenizeValueComponentsFieldProp**](EnumlogFieldBehaviorTextAccessTokenizeValueComponentsFieldProp.md) |  | [optional] 
**TokenizeValueComponentsFieldName** | Pointer to **[]string** | The names of any custom fields for which to tokenize components within the value. This should generally only be used for fields that are not available through the tokenize-value-components-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**Description** | Pointer to **string** | A description for this Log Field Behavior | [optional] 
**DefaultBehavior** | Pointer to [**EnumlogFieldBehaviorDefaultBehaviorProp**](EnumlogFieldBehaviorDefaultBehaviorProp.md) |  | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewLogFieldBehaviorListResponseResourcesInner

`func NewLogFieldBehaviorListResponseResourcesInner(id string, schemas []EnumtextAccessLogFieldBehaviorSchemaUrn, ) *LogFieldBehaviorListResponseResourcesInner`

NewLogFieldBehaviorListResponseResourcesInner instantiates a new LogFieldBehaviorListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogFieldBehaviorListResponseResourcesInnerWithDefaults

`func NewLogFieldBehaviorListResponseResourcesInnerWithDefaults() *LogFieldBehaviorListResponseResourcesInner`

NewLogFieldBehaviorListResponseResourcesInnerWithDefaults instantiates a new LogFieldBehaviorListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogFieldBehaviorListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogFieldBehaviorListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *LogFieldBehaviorListResponseResourcesInner) GetSchemas() []EnumtextAccessLogFieldBehaviorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetSchemasOk() (*[]EnumtextAccessLogFieldBehaviorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LogFieldBehaviorListResponseResourcesInner) SetSchemas(v []EnumtextAccessLogFieldBehaviorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPreserveField

`func (o *LogFieldBehaviorListResponseResourcesInner) GetPreserveField() []EnumlogFieldBehaviorTextAccessPreserveFieldProp`

GetPreserveField returns the PreserveField field if non-nil, zero value otherwise.

### GetPreserveFieldOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetPreserveFieldOk() (*[]EnumlogFieldBehaviorTextAccessPreserveFieldProp, bool)`

GetPreserveFieldOk returns a tuple with the PreserveField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveField

`func (o *LogFieldBehaviorListResponseResourcesInner) SetPreserveField(v []EnumlogFieldBehaviorTextAccessPreserveFieldProp)`

SetPreserveField sets PreserveField field to given value.

### HasPreserveField

`func (o *LogFieldBehaviorListResponseResourcesInner) HasPreserveField() bool`

HasPreserveField returns a boolean if a field has been set.

### GetPreserveFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) GetPreserveFieldName() []string`

GetPreserveFieldName returns the PreserveFieldName field if non-nil, zero value otherwise.

### GetPreserveFieldNameOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetPreserveFieldNameOk() (*[]string, bool)`

GetPreserveFieldNameOk returns a tuple with the PreserveFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) SetPreserveFieldName(v []string)`

SetPreserveFieldName sets PreserveFieldName field to given value.

### HasPreserveFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) HasPreserveFieldName() bool`

HasPreserveFieldName returns a boolean if a field has been set.

### GetOmitField

`func (o *LogFieldBehaviorListResponseResourcesInner) GetOmitField() []EnumlogFieldBehaviorTextAccessOmitFieldProp`

GetOmitField returns the OmitField field if non-nil, zero value otherwise.

### GetOmitFieldOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetOmitFieldOk() (*[]EnumlogFieldBehaviorTextAccessOmitFieldProp, bool)`

GetOmitFieldOk returns a tuple with the OmitField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitField

`func (o *LogFieldBehaviorListResponseResourcesInner) SetOmitField(v []EnumlogFieldBehaviorTextAccessOmitFieldProp)`

SetOmitField sets OmitField field to given value.

### HasOmitField

`func (o *LogFieldBehaviorListResponseResourcesInner) HasOmitField() bool`

HasOmitField returns a boolean if a field has been set.

### GetOmitFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) GetOmitFieldName() []string`

GetOmitFieldName returns the OmitFieldName field if non-nil, zero value otherwise.

### GetOmitFieldNameOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetOmitFieldNameOk() (*[]string, bool)`

GetOmitFieldNameOk returns a tuple with the OmitFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) SetOmitFieldName(v []string)`

SetOmitFieldName sets OmitFieldName field to given value.

### HasOmitFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) HasOmitFieldName() bool`

HasOmitFieldName returns a boolean if a field has been set.

### GetRedactEntireValueField

`func (o *LogFieldBehaviorListResponseResourcesInner) GetRedactEntireValueField() []EnumlogFieldBehaviorTextAccessRedactEntireValueFieldProp`

GetRedactEntireValueField returns the RedactEntireValueField field if non-nil, zero value otherwise.

### GetRedactEntireValueFieldOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetRedactEntireValueFieldOk() (*[]EnumlogFieldBehaviorTextAccessRedactEntireValueFieldProp, bool)`

GetRedactEntireValueFieldOk returns a tuple with the RedactEntireValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactEntireValueField

`func (o *LogFieldBehaviorListResponseResourcesInner) SetRedactEntireValueField(v []EnumlogFieldBehaviorTextAccessRedactEntireValueFieldProp)`

SetRedactEntireValueField sets RedactEntireValueField field to given value.

### HasRedactEntireValueField

`func (o *LogFieldBehaviorListResponseResourcesInner) HasRedactEntireValueField() bool`

HasRedactEntireValueField returns a boolean if a field has been set.

### GetRedactEntireValueFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) GetRedactEntireValueFieldName() []string`

GetRedactEntireValueFieldName returns the RedactEntireValueFieldName field if non-nil, zero value otherwise.

### GetRedactEntireValueFieldNameOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetRedactEntireValueFieldNameOk() (*[]string, bool)`

GetRedactEntireValueFieldNameOk returns a tuple with the RedactEntireValueFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactEntireValueFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) SetRedactEntireValueFieldName(v []string)`

SetRedactEntireValueFieldName sets RedactEntireValueFieldName field to given value.

### HasRedactEntireValueFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) HasRedactEntireValueFieldName() bool`

HasRedactEntireValueFieldName returns a boolean if a field has been set.

### GetRedactValueComponentsField

`func (o *LogFieldBehaviorListResponseResourcesInner) GetRedactValueComponentsField() []EnumlogFieldBehaviorTextAccessRedactValueComponentsFieldProp`

GetRedactValueComponentsField returns the RedactValueComponentsField field if non-nil, zero value otherwise.

### GetRedactValueComponentsFieldOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetRedactValueComponentsFieldOk() (*[]EnumlogFieldBehaviorTextAccessRedactValueComponentsFieldProp, bool)`

GetRedactValueComponentsFieldOk returns a tuple with the RedactValueComponentsField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactValueComponentsField

`func (o *LogFieldBehaviorListResponseResourcesInner) SetRedactValueComponentsField(v []EnumlogFieldBehaviorTextAccessRedactValueComponentsFieldProp)`

SetRedactValueComponentsField sets RedactValueComponentsField field to given value.

### HasRedactValueComponentsField

`func (o *LogFieldBehaviorListResponseResourcesInner) HasRedactValueComponentsField() bool`

HasRedactValueComponentsField returns a boolean if a field has been set.

### GetRedactValueComponentsFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) GetRedactValueComponentsFieldName() []string`

GetRedactValueComponentsFieldName returns the RedactValueComponentsFieldName field if non-nil, zero value otherwise.

### GetRedactValueComponentsFieldNameOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetRedactValueComponentsFieldNameOk() (*[]string, bool)`

GetRedactValueComponentsFieldNameOk returns a tuple with the RedactValueComponentsFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactValueComponentsFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) SetRedactValueComponentsFieldName(v []string)`

SetRedactValueComponentsFieldName sets RedactValueComponentsFieldName field to given value.

### HasRedactValueComponentsFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) HasRedactValueComponentsFieldName() bool`

HasRedactValueComponentsFieldName returns a boolean if a field has been set.

### GetTokenizeEntireValueField

`func (o *LogFieldBehaviorListResponseResourcesInner) GetTokenizeEntireValueField() []EnumlogFieldBehaviorTextAccessTokenizeEntireValueFieldProp`

GetTokenizeEntireValueField returns the TokenizeEntireValueField field if non-nil, zero value otherwise.

### GetTokenizeEntireValueFieldOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetTokenizeEntireValueFieldOk() (*[]EnumlogFieldBehaviorTextAccessTokenizeEntireValueFieldProp, bool)`

GetTokenizeEntireValueFieldOk returns a tuple with the TokenizeEntireValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeEntireValueField

`func (o *LogFieldBehaviorListResponseResourcesInner) SetTokenizeEntireValueField(v []EnumlogFieldBehaviorTextAccessTokenizeEntireValueFieldProp)`

SetTokenizeEntireValueField sets TokenizeEntireValueField field to given value.

### HasTokenizeEntireValueField

`func (o *LogFieldBehaviorListResponseResourcesInner) HasTokenizeEntireValueField() bool`

HasTokenizeEntireValueField returns a boolean if a field has been set.

### GetTokenizeEntireValueFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) GetTokenizeEntireValueFieldName() []string`

GetTokenizeEntireValueFieldName returns the TokenizeEntireValueFieldName field if non-nil, zero value otherwise.

### GetTokenizeEntireValueFieldNameOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetTokenizeEntireValueFieldNameOk() (*[]string, bool)`

GetTokenizeEntireValueFieldNameOk returns a tuple with the TokenizeEntireValueFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeEntireValueFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) SetTokenizeEntireValueFieldName(v []string)`

SetTokenizeEntireValueFieldName sets TokenizeEntireValueFieldName field to given value.

### HasTokenizeEntireValueFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) HasTokenizeEntireValueFieldName() bool`

HasTokenizeEntireValueFieldName returns a boolean if a field has been set.

### GetTokenizeValueComponentsField

`func (o *LogFieldBehaviorListResponseResourcesInner) GetTokenizeValueComponentsField() []EnumlogFieldBehaviorTextAccessTokenizeValueComponentsFieldProp`

GetTokenizeValueComponentsField returns the TokenizeValueComponentsField field if non-nil, zero value otherwise.

### GetTokenizeValueComponentsFieldOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetTokenizeValueComponentsFieldOk() (*[]EnumlogFieldBehaviorTextAccessTokenizeValueComponentsFieldProp, bool)`

GetTokenizeValueComponentsFieldOk returns a tuple with the TokenizeValueComponentsField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeValueComponentsField

`func (o *LogFieldBehaviorListResponseResourcesInner) SetTokenizeValueComponentsField(v []EnumlogFieldBehaviorTextAccessTokenizeValueComponentsFieldProp)`

SetTokenizeValueComponentsField sets TokenizeValueComponentsField field to given value.

### HasTokenizeValueComponentsField

`func (o *LogFieldBehaviorListResponseResourcesInner) HasTokenizeValueComponentsField() bool`

HasTokenizeValueComponentsField returns a boolean if a field has been set.

### GetTokenizeValueComponentsFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) GetTokenizeValueComponentsFieldName() []string`

GetTokenizeValueComponentsFieldName returns the TokenizeValueComponentsFieldName field if non-nil, zero value otherwise.

### GetTokenizeValueComponentsFieldNameOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetTokenizeValueComponentsFieldNameOk() (*[]string, bool)`

GetTokenizeValueComponentsFieldNameOk returns a tuple with the TokenizeValueComponentsFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeValueComponentsFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) SetTokenizeValueComponentsFieldName(v []string)`

SetTokenizeValueComponentsFieldName sets TokenizeValueComponentsFieldName field to given value.

### HasTokenizeValueComponentsFieldName

`func (o *LogFieldBehaviorListResponseResourcesInner) HasTokenizeValueComponentsFieldName() bool`

HasTokenizeValueComponentsFieldName returns a boolean if a field has been set.

### GetDescription

`func (o *LogFieldBehaviorListResponseResourcesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LogFieldBehaviorListResponseResourcesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LogFieldBehaviorListResponseResourcesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultBehavior

`func (o *LogFieldBehaviorListResponseResourcesInner) GetDefaultBehavior() EnumlogFieldBehaviorDefaultBehaviorProp`

GetDefaultBehavior returns the DefaultBehavior field if non-nil, zero value otherwise.

### GetDefaultBehaviorOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetDefaultBehaviorOk() (*EnumlogFieldBehaviorDefaultBehaviorProp, bool)`

GetDefaultBehaviorOk returns a tuple with the DefaultBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBehavior

`func (o *LogFieldBehaviorListResponseResourcesInner) SetDefaultBehavior(v EnumlogFieldBehaviorDefaultBehaviorProp)`

SetDefaultBehavior sets DefaultBehavior field to given value.

### HasDefaultBehavior

`func (o *LogFieldBehaviorListResponseResourcesInner) HasDefaultBehavior() bool`

HasDefaultBehavior returns a boolean if a field has been set.

### GetMeta

`func (o *LogFieldBehaviorListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LogFieldBehaviorListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LogFieldBehaviorListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LogFieldBehaviorListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *LogFieldBehaviorListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *LogFieldBehaviorListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *LogFieldBehaviorListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *LogFieldBehaviorListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


