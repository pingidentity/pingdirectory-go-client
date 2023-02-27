# TextAccessLogFieldBehaviorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
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

## Methods

### NewTextAccessLogFieldBehaviorResponse

`func NewTextAccessLogFieldBehaviorResponse(id string, schemas []EnumtextAccessLogFieldBehaviorSchemaUrn, ) *TextAccessLogFieldBehaviorResponse`

NewTextAccessLogFieldBehaviorResponse instantiates a new TextAccessLogFieldBehaviorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextAccessLogFieldBehaviorResponseWithDefaults

`func NewTextAccessLogFieldBehaviorResponseWithDefaults() *TextAccessLogFieldBehaviorResponse`

NewTextAccessLogFieldBehaviorResponseWithDefaults instantiates a new TextAccessLogFieldBehaviorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *TextAccessLogFieldBehaviorResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TextAccessLogFieldBehaviorResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TextAccessLogFieldBehaviorResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TextAccessLogFieldBehaviorResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *TextAccessLogFieldBehaviorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *TextAccessLogFieldBehaviorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *TextAccessLogFieldBehaviorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *TextAccessLogFieldBehaviorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *TextAccessLogFieldBehaviorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TextAccessLogFieldBehaviorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TextAccessLogFieldBehaviorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *TextAccessLogFieldBehaviorResponse) GetSchemas() []EnumtextAccessLogFieldBehaviorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TextAccessLogFieldBehaviorResponse) GetSchemasOk() (*[]EnumtextAccessLogFieldBehaviorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TextAccessLogFieldBehaviorResponse) SetSchemas(v []EnumtextAccessLogFieldBehaviorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPreserveField

`func (o *TextAccessLogFieldBehaviorResponse) GetPreserveField() []EnumlogFieldBehaviorTextAccessPreserveFieldProp`

GetPreserveField returns the PreserveField field if non-nil, zero value otherwise.

### GetPreserveFieldOk

`func (o *TextAccessLogFieldBehaviorResponse) GetPreserveFieldOk() (*[]EnumlogFieldBehaviorTextAccessPreserveFieldProp, bool)`

GetPreserveFieldOk returns a tuple with the PreserveField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveField

`func (o *TextAccessLogFieldBehaviorResponse) SetPreserveField(v []EnumlogFieldBehaviorTextAccessPreserveFieldProp)`

SetPreserveField sets PreserveField field to given value.

### HasPreserveField

`func (o *TextAccessLogFieldBehaviorResponse) HasPreserveField() bool`

HasPreserveField returns a boolean if a field has been set.

### GetPreserveFieldName

`func (o *TextAccessLogFieldBehaviorResponse) GetPreserveFieldName() []string`

GetPreserveFieldName returns the PreserveFieldName field if non-nil, zero value otherwise.

### GetPreserveFieldNameOk

`func (o *TextAccessLogFieldBehaviorResponse) GetPreserveFieldNameOk() (*[]string, bool)`

GetPreserveFieldNameOk returns a tuple with the PreserveFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveFieldName

`func (o *TextAccessLogFieldBehaviorResponse) SetPreserveFieldName(v []string)`

SetPreserveFieldName sets PreserveFieldName field to given value.

### HasPreserveFieldName

`func (o *TextAccessLogFieldBehaviorResponse) HasPreserveFieldName() bool`

HasPreserveFieldName returns a boolean if a field has been set.

### GetOmitField

`func (o *TextAccessLogFieldBehaviorResponse) GetOmitField() []EnumlogFieldBehaviorTextAccessOmitFieldProp`

GetOmitField returns the OmitField field if non-nil, zero value otherwise.

### GetOmitFieldOk

`func (o *TextAccessLogFieldBehaviorResponse) GetOmitFieldOk() (*[]EnumlogFieldBehaviorTextAccessOmitFieldProp, bool)`

GetOmitFieldOk returns a tuple with the OmitField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitField

`func (o *TextAccessLogFieldBehaviorResponse) SetOmitField(v []EnumlogFieldBehaviorTextAccessOmitFieldProp)`

SetOmitField sets OmitField field to given value.

### HasOmitField

`func (o *TextAccessLogFieldBehaviorResponse) HasOmitField() bool`

HasOmitField returns a boolean if a field has been set.

### GetOmitFieldName

`func (o *TextAccessLogFieldBehaviorResponse) GetOmitFieldName() []string`

GetOmitFieldName returns the OmitFieldName field if non-nil, zero value otherwise.

### GetOmitFieldNameOk

`func (o *TextAccessLogFieldBehaviorResponse) GetOmitFieldNameOk() (*[]string, bool)`

GetOmitFieldNameOk returns a tuple with the OmitFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitFieldName

`func (o *TextAccessLogFieldBehaviorResponse) SetOmitFieldName(v []string)`

SetOmitFieldName sets OmitFieldName field to given value.

### HasOmitFieldName

`func (o *TextAccessLogFieldBehaviorResponse) HasOmitFieldName() bool`

HasOmitFieldName returns a boolean if a field has been set.

### GetRedactEntireValueField

`func (o *TextAccessLogFieldBehaviorResponse) GetRedactEntireValueField() []EnumlogFieldBehaviorTextAccessRedactEntireValueFieldProp`

GetRedactEntireValueField returns the RedactEntireValueField field if non-nil, zero value otherwise.

### GetRedactEntireValueFieldOk

`func (o *TextAccessLogFieldBehaviorResponse) GetRedactEntireValueFieldOk() (*[]EnumlogFieldBehaviorTextAccessRedactEntireValueFieldProp, bool)`

GetRedactEntireValueFieldOk returns a tuple with the RedactEntireValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactEntireValueField

`func (o *TextAccessLogFieldBehaviorResponse) SetRedactEntireValueField(v []EnumlogFieldBehaviorTextAccessRedactEntireValueFieldProp)`

SetRedactEntireValueField sets RedactEntireValueField field to given value.

### HasRedactEntireValueField

`func (o *TextAccessLogFieldBehaviorResponse) HasRedactEntireValueField() bool`

HasRedactEntireValueField returns a boolean if a field has been set.

### GetRedactEntireValueFieldName

`func (o *TextAccessLogFieldBehaviorResponse) GetRedactEntireValueFieldName() []string`

GetRedactEntireValueFieldName returns the RedactEntireValueFieldName field if non-nil, zero value otherwise.

### GetRedactEntireValueFieldNameOk

`func (o *TextAccessLogFieldBehaviorResponse) GetRedactEntireValueFieldNameOk() (*[]string, bool)`

GetRedactEntireValueFieldNameOk returns a tuple with the RedactEntireValueFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactEntireValueFieldName

`func (o *TextAccessLogFieldBehaviorResponse) SetRedactEntireValueFieldName(v []string)`

SetRedactEntireValueFieldName sets RedactEntireValueFieldName field to given value.

### HasRedactEntireValueFieldName

`func (o *TextAccessLogFieldBehaviorResponse) HasRedactEntireValueFieldName() bool`

HasRedactEntireValueFieldName returns a boolean if a field has been set.

### GetRedactValueComponentsField

`func (o *TextAccessLogFieldBehaviorResponse) GetRedactValueComponentsField() []EnumlogFieldBehaviorTextAccessRedactValueComponentsFieldProp`

GetRedactValueComponentsField returns the RedactValueComponentsField field if non-nil, zero value otherwise.

### GetRedactValueComponentsFieldOk

`func (o *TextAccessLogFieldBehaviorResponse) GetRedactValueComponentsFieldOk() (*[]EnumlogFieldBehaviorTextAccessRedactValueComponentsFieldProp, bool)`

GetRedactValueComponentsFieldOk returns a tuple with the RedactValueComponentsField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactValueComponentsField

`func (o *TextAccessLogFieldBehaviorResponse) SetRedactValueComponentsField(v []EnumlogFieldBehaviorTextAccessRedactValueComponentsFieldProp)`

SetRedactValueComponentsField sets RedactValueComponentsField field to given value.

### HasRedactValueComponentsField

`func (o *TextAccessLogFieldBehaviorResponse) HasRedactValueComponentsField() bool`

HasRedactValueComponentsField returns a boolean if a field has been set.

### GetRedactValueComponentsFieldName

`func (o *TextAccessLogFieldBehaviorResponse) GetRedactValueComponentsFieldName() []string`

GetRedactValueComponentsFieldName returns the RedactValueComponentsFieldName field if non-nil, zero value otherwise.

### GetRedactValueComponentsFieldNameOk

`func (o *TextAccessLogFieldBehaviorResponse) GetRedactValueComponentsFieldNameOk() (*[]string, bool)`

GetRedactValueComponentsFieldNameOk returns a tuple with the RedactValueComponentsFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactValueComponentsFieldName

`func (o *TextAccessLogFieldBehaviorResponse) SetRedactValueComponentsFieldName(v []string)`

SetRedactValueComponentsFieldName sets RedactValueComponentsFieldName field to given value.

### HasRedactValueComponentsFieldName

`func (o *TextAccessLogFieldBehaviorResponse) HasRedactValueComponentsFieldName() bool`

HasRedactValueComponentsFieldName returns a boolean if a field has been set.

### GetTokenizeEntireValueField

`func (o *TextAccessLogFieldBehaviorResponse) GetTokenizeEntireValueField() []EnumlogFieldBehaviorTextAccessTokenizeEntireValueFieldProp`

GetTokenizeEntireValueField returns the TokenizeEntireValueField field if non-nil, zero value otherwise.

### GetTokenizeEntireValueFieldOk

`func (o *TextAccessLogFieldBehaviorResponse) GetTokenizeEntireValueFieldOk() (*[]EnumlogFieldBehaviorTextAccessTokenizeEntireValueFieldProp, bool)`

GetTokenizeEntireValueFieldOk returns a tuple with the TokenizeEntireValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeEntireValueField

`func (o *TextAccessLogFieldBehaviorResponse) SetTokenizeEntireValueField(v []EnumlogFieldBehaviorTextAccessTokenizeEntireValueFieldProp)`

SetTokenizeEntireValueField sets TokenizeEntireValueField field to given value.

### HasTokenizeEntireValueField

`func (o *TextAccessLogFieldBehaviorResponse) HasTokenizeEntireValueField() bool`

HasTokenizeEntireValueField returns a boolean if a field has been set.

### GetTokenizeEntireValueFieldName

`func (o *TextAccessLogFieldBehaviorResponse) GetTokenizeEntireValueFieldName() []string`

GetTokenizeEntireValueFieldName returns the TokenizeEntireValueFieldName field if non-nil, zero value otherwise.

### GetTokenizeEntireValueFieldNameOk

`func (o *TextAccessLogFieldBehaviorResponse) GetTokenizeEntireValueFieldNameOk() (*[]string, bool)`

GetTokenizeEntireValueFieldNameOk returns a tuple with the TokenizeEntireValueFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeEntireValueFieldName

`func (o *TextAccessLogFieldBehaviorResponse) SetTokenizeEntireValueFieldName(v []string)`

SetTokenizeEntireValueFieldName sets TokenizeEntireValueFieldName field to given value.

### HasTokenizeEntireValueFieldName

`func (o *TextAccessLogFieldBehaviorResponse) HasTokenizeEntireValueFieldName() bool`

HasTokenizeEntireValueFieldName returns a boolean if a field has been set.

### GetTokenizeValueComponentsField

`func (o *TextAccessLogFieldBehaviorResponse) GetTokenizeValueComponentsField() []EnumlogFieldBehaviorTextAccessTokenizeValueComponentsFieldProp`

GetTokenizeValueComponentsField returns the TokenizeValueComponentsField field if non-nil, zero value otherwise.

### GetTokenizeValueComponentsFieldOk

`func (o *TextAccessLogFieldBehaviorResponse) GetTokenizeValueComponentsFieldOk() (*[]EnumlogFieldBehaviorTextAccessTokenizeValueComponentsFieldProp, bool)`

GetTokenizeValueComponentsFieldOk returns a tuple with the TokenizeValueComponentsField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeValueComponentsField

`func (o *TextAccessLogFieldBehaviorResponse) SetTokenizeValueComponentsField(v []EnumlogFieldBehaviorTextAccessTokenizeValueComponentsFieldProp)`

SetTokenizeValueComponentsField sets TokenizeValueComponentsField field to given value.

### HasTokenizeValueComponentsField

`func (o *TextAccessLogFieldBehaviorResponse) HasTokenizeValueComponentsField() bool`

HasTokenizeValueComponentsField returns a boolean if a field has been set.

### GetTokenizeValueComponentsFieldName

`func (o *TextAccessLogFieldBehaviorResponse) GetTokenizeValueComponentsFieldName() []string`

GetTokenizeValueComponentsFieldName returns the TokenizeValueComponentsFieldName field if non-nil, zero value otherwise.

### GetTokenizeValueComponentsFieldNameOk

`func (o *TextAccessLogFieldBehaviorResponse) GetTokenizeValueComponentsFieldNameOk() (*[]string, bool)`

GetTokenizeValueComponentsFieldNameOk returns a tuple with the TokenizeValueComponentsFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeValueComponentsFieldName

`func (o *TextAccessLogFieldBehaviorResponse) SetTokenizeValueComponentsFieldName(v []string)`

SetTokenizeValueComponentsFieldName sets TokenizeValueComponentsFieldName field to given value.

### HasTokenizeValueComponentsFieldName

`func (o *TextAccessLogFieldBehaviorResponse) HasTokenizeValueComponentsFieldName() bool`

HasTokenizeValueComponentsFieldName returns a boolean if a field has been set.

### GetDescription

`func (o *TextAccessLogFieldBehaviorResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TextAccessLogFieldBehaviorResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TextAccessLogFieldBehaviorResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TextAccessLogFieldBehaviorResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultBehavior

`func (o *TextAccessLogFieldBehaviorResponse) GetDefaultBehavior() EnumlogFieldBehaviorDefaultBehaviorProp`

GetDefaultBehavior returns the DefaultBehavior field if non-nil, zero value otherwise.

### GetDefaultBehaviorOk

`func (o *TextAccessLogFieldBehaviorResponse) GetDefaultBehaviorOk() (*EnumlogFieldBehaviorDefaultBehaviorProp, bool)`

GetDefaultBehaviorOk returns a tuple with the DefaultBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBehavior

`func (o *TextAccessLogFieldBehaviorResponse) SetDefaultBehavior(v EnumlogFieldBehaviorDefaultBehaviorProp)`

SetDefaultBehavior sets DefaultBehavior field to given value.

### HasDefaultBehavior

`func (o *TextAccessLogFieldBehaviorResponse) HasDefaultBehavior() bool`

HasDefaultBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


