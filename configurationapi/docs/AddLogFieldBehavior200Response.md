# AddLogFieldBehavior200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Log Field Behavior | 
**Schemas** | [**[]EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn**](EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn.md) |  | 
**PreserveField** | Pointer to [**[]EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp**](EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp.md) | The log fields whose values should be logged with the intended value. The values for these fields will be preserved, although they may be sanitized for parsability or safety purposes (for example, to escape special characters in the value), and values that are too long may be truncated. | [optional] 
**PreserveFieldName** | Pointer to **[]string** | The names of any custom fields whose values should be preserved. This should generally only be used for fields that are not available through the preserve-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**OmitField** | Pointer to [**[]EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp**](EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp.md) | The log fields that should be omitted entirely from log messages. Neither the field name nor value will be included. | [optional] 
**OmitFieldName** | Pointer to **[]string** | The names of any custom fields that should be omitted from log messages. This should generally only be used for fields that are not available through the omit-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**RedactEntireValueField** | Pointer to [**[]EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp**](EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp.md) | The log fields whose values should be completely redacted in log messages. The field name will be included, but with a fixed value that does not reflect the actual value for the field. | [optional] 
**RedactEntireValueFieldName** | Pointer to **[]string** | The names of any custom fields whose values should be completely redacted. This should generally only be used for fields that are not available through the redact-entire-value-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**RedactValueComponentsField** | Pointer to [**[]EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp**](EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp.md) | The log fields whose values will include redacted components. | [optional] 
**RedactValueComponentsFieldName** | Pointer to **[]string** | The names of any custom fields for which to redact components within the value. This should generally only be used for fields that are not available through the redact-value-components-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**TokenizeEntireValueField** | Pointer to [**[]EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp**](EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp.md) | The log fields whose values should be completely tokenized in log messages. The field name will be included, but the value will be replaced with a token that does not reveal the actual value, but that is generated from the value. | [optional] 
**TokenizeEntireValueFieldName** | Pointer to **[]string** | The names of any custom fields whose values should be completely tokenized. This should generally only be used for fields that are not available through the tokenize-entire-value-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**TokenizeValueComponentsField** | Pointer to [**[]EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp**](EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp.md) | The log fields whose values will include tokenized components. | [optional] 
**TokenizeValueComponentsFieldName** | Pointer to **[]string** | The names of any custom fields for which to tokenize components within the value. This should generally only be used for fields that are not available through the tokenize-value-components-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**Description** | Pointer to **string** | A description for this Log Field Behavior | [optional] 
**DefaultBehavior** | Pointer to [**EnumlogFieldBehaviorDefaultBehaviorProp**](EnumlogFieldBehaviorDefaultBehaviorProp.md) |  | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewAddLogFieldBehavior200Response

`func NewAddLogFieldBehavior200Response(id string, schemas []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn, ) *AddLogFieldBehavior200Response`

NewAddLogFieldBehavior200Response instantiates a new AddLogFieldBehavior200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLogFieldBehavior200ResponseWithDefaults

`func NewAddLogFieldBehavior200ResponseWithDefaults() *AddLogFieldBehavior200Response`

NewAddLogFieldBehavior200ResponseWithDefaults instantiates a new AddLogFieldBehavior200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddLogFieldBehavior200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddLogFieldBehavior200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddLogFieldBehavior200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddLogFieldBehavior200Response) GetSchemas() []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLogFieldBehavior200Response) GetSchemasOk() (*[]EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLogFieldBehavior200Response) SetSchemas(v []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPreserveField

`func (o *AddLogFieldBehavior200Response) GetPreserveField() []EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp`

GetPreserveField returns the PreserveField field if non-nil, zero value otherwise.

### GetPreserveFieldOk

`func (o *AddLogFieldBehavior200Response) GetPreserveFieldOk() (*[]EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp, bool)`

GetPreserveFieldOk returns a tuple with the PreserveField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveField

`func (o *AddLogFieldBehavior200Response) SetPreserveField(v []EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp)`

SetPreserveField sets PreserveField field to given value.

### HasPreserveField

`func (o *AddLogFieldBehavior200Response) HasPreserveField() bool`

HasPreserveField returns a boolean if a field has been set.

### GetPreserveFieldName

`func (o *AddLogFieldBehavior200Response) GetPreserveFieldName() []string`

GetPreserveFieldName returns the PreserveFieldName field if non-nil, zero value otherwise.

### GetPreserveFieldNameOk

`func (o *AddLogFieldBehavior200Response) GetPreserveFieldNameOk() (*[]string, bool)`

GetPreserveFieldNameOk returns a tuple with the PreserveFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveFieldName

`func (o *AddLogFieldBehavior200Response) SetPreserveFieldName(v []string)`

SetPreserveFieldName sets PreserveFieldName field to given value.

### HasPreserveFieldName

`func (o *AddLogFieldBehavior200Response) HasPreserveFieldName() bool`

HasPreserveFieldName returns a boolean if a field has been set.

### GetOmitField

`func (o *AddLogFieldBehavior200Response) GetOmitField() []EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp`

GetOmitField returns the OmitField field if non-nil, zero value otherwise.

### GetOmitFieldOk

`func (o *AddLogFieldBehavior200Response) GetOmitFieldOk() (*[]EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp, bool)`

GetOmitFieldOk returns a tuple with the OmitField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitField

`func (o *AddLogFieldBehavior200Response) SetOmitField(v []EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp)`

SetOmitField sets OmitField field to given value.

### HasOmitField

`func (o *AddLogFieldBehavior200Response) HasOmitField() bool`

HasOmitField returns a boolean if a field has been set.

### GetOmitFieldName

`func (o *AddLogFieldBehavior200Response) GetOmitFieldName() []string`

GetOmitFieldName returns the OmitFieldName field if non-nil, zero value otherwise.

### GetOmitFieldNameOk

`func (o *AddLogFieldBehavior200Response) GetOmitFieldNameOk() (*[]string, bool)`

GetOmitFieldNameOk returns a tuple with the OmitFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitFieldName

`func (o *AddLogFieldBehavior200Response) SetOmitFieldName(v []string)`

SetOmitFieldName sets OmitFieldName field to given value.

### HasOmitFieldName

`func (o *AddLogFieldBehavior200Response) HasOmitFieldName() bool`

HasOmitFieldName returns a boolean if a field has been set.

### GetRedactEntireValueField

`func (o *AddLogFieldBehavior200Response) GetRedactEntireValueField() []EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp`

GetRedactEntireValueField returns the RedactEntireValueField field if non-nil, zero value otherwise.

### GetRedactEntireValueFieldOk

`func (o *AddLogFieldBehavior200Response) GetRedactEntireValueFieldOk() (*[]EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp, bool)`

GetRedactEntireValueFieldOk returns a tuple with the RedactEntireValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactEntireValueField

`func (o *AddLogFieldBehavior200Response) SetRedactEntireValueField(v []EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp)`

SetRedactEntireValueField sets RedactEntireValueField field to given value.

### HasRedactEntireValueField

`func (o *AddLogFieldBehavior200Response) HasRedactEntireValueField() bool`

HasRedactEntireValueField returns a boolean if a field has been set.

### GetRedactEntireValueFieldName

`func (o *AddLogFieldBehavior200Response) GetRedactEntireValueFieldName() []string`

GetRedactEntireValueFieldName returns the RedactEntireValueFieldName field if non-nil, zero value otherwise.

### GetRedactEntireValueFieldNameOk

`func (o *AddLogFieldBehavior200Response) GetRedactEntireValueFieldNameOk() (*[]string, bool)`

GetRedactEntireValueFieldNameOk returns a tuple with the RedactEntireValueFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactEntireValueFieldName

`func (o *AddLogFieldBehavior200Response) SetRedactEntireValueFieldName(v []string)`

SetRedactEntireValueFieldName sets RedactEntireValueFieldName field to given value.

### HasRedactEntireValueFieldName

`func (o *AddLogFieldBehavior200Response) HasRedactEntireValueFieldName() bool`

HasRedactEntireValueFieldName returns a boolean if a field has been set.

### GetRedactValueComponentsField

`func (o *AddLogFieldBehavior200Response) GetRedactValueComponentsField() []EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp`

GetRedactValueComponentsField returns the RedactValueComponentsField field if non-nil, zero value otherwise.

### GetRedactValueComponentsFieldOk

`func (o *AddLogFieldBehavior200Response) GetRedactValueComponentsFieldOk() (*[]EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp, bool)`

GetRedactValueComponentsFieldOk returns a tuple with the RedactValueComponentsField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactValueComponentsField

`func (o *AddLogFieldBehavior200Response) SetRedactValueComponentsField(v []EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp)`

SetRedactValueComponentsField sets RedactValueComponentsField field to given value.

### HasRedactValueComponentsField

`func (o *AddLogFieldBehavior200Response) HasRedactValueComponentsField() bool`

HasRedactValueComponentsField returns a boolean if a field has been set.

### GetRedactValueComponentsFieldName

`func (o *AddLogFieldBehavior200Response) GetRedactValueComponentsFieldName() []string`

GetRedactValueComponentsFieldName returns the RedactValueComponentsFieldName field if non-nil, zero value otherwise.

### GetRedactValueComponentsFieldNameOk

`func (o *AddLogFieldBehavior200Response) GetRedactValueComponentsFieldNameOk() (*[]string, bool)`

GetRedactValueComponentsFieldNameOk returns a tuple with the RedactValueComponentsFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactValueComponentsFieldName

`func (o *AddLogFieldBehavior200Response) SetRedactValueComponentsFieldName(v []string)`

SetRedactValueComponentsFieldName sets RedactValueComponentsFieldName field to given value.

### HasRedactValueComponentsFieldName

`func (o *AddLogFieldBehavior200Response) HasRedactValueComponentsFieldName() bool`

HasRedactValueComponentsFieldName returns a boolean if a field has been set.

### GetTokenizeEntireValueField

`func (o *AddLogFieldBehavior200Response) GetTokenizeEntireValueField() []EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp`

GetTokenizeEntireValueField returns the TokenizeEntireValueField field if non-nil, zero value otherwise.

### GetTokenizeEntireValueFieldOk

`func (o *AddLogFieldBehavior200Response) GetTokenizeEntireValueFieldOk() (*[]EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp, bool)`

GetTokenizeEntireValueFieldOk returns a tuple with the TokenizeEntireValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeEntireValueField

`func (o *AddLogFieldBehavior200Response) SetTokenizeEntireValueField(v []EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp)`

SetTokenizeEntireValueField sets TokenizeEntireValueField field to given value.

### HasTokenizeEntireValueField

`func (o *AddLogFieldBehavior200Response) HasTokenizeEntireValueField() bool`

HasTokenizeEntireValueField returns a boolean if a field has been set.

### GetTokenizeEntireValueFieldName

`func (o *AddLogFieldBehavior200Response) GetTokenizeEntireValueFieldName() []string`

GetTokenizeEntireValueFieldName returns the TokenizeEntireValueFieldName field if non-nil, zero value otherwise.

### GetTokenizeEntireValueFieldNameOk

`func (o *AddLogFieldBehavior200Response) GetTokenizeEntireValueFieldNameOk() (*[]string, bool)`

GetTokenizeEntireValueFieldNameOk returns a tuple with the TokenizeEntireValueFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeEntireValueFieldName

`func (o *AddLogFieldBehavior200Response) SetTokenizeEntireValueFieldName(v []string)`

SetTokenizeEntireValueFieldName sets TokenizeEntireValueFieldName field to given value.

### HasTokenizeEntireValueFieldName

`func (o *AddLogFieldBehavior200Response) HasTokenizeEntireValueFieldName() bool`

HasTokenizeEntireValueFieldName returns a boolean if a field has been set.

### GetTokenizeValueComponentsField

`func (o *AddLogFieldBehavior200Response) GetTokenizeValueComponentsField() []EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp`

GetTokenizeValueComponentsField returns the TokenizeValueComponentsField field if non-nil, zero value otherwise.

### GetTokenizeValueComponentsFieldOk

`func (o *AddLogFieldBehavior200Response) GetTokenizeValueComponentsFieldOk() (*[]EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp, bool)`

GetTokenizeValueComponentsFieldOk returns a tuple with the TokenizeValueComponentsField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeValueComponentsField

`func (o *AddLogFieldBehavior200Response) SetTokenizeValueComponentsField(v []EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp)`

SetTokenizeValueComponentsField sets TokenizeValueComponentsField field to given value.

### HasTokenizeValueComponentsField

`func (o *AddLogFieldBehavior200Response) HasTokenizeValueComponentsField() bool`

HasTokenizeValueComponentsField returns a boolean if a field has been set.

### GetTokenizeValueComponentsFieldName

`func (o *AddLogFieldBehavior200Response) GetTokenizeValueComponentsFieldName() []string`

GetTokenizeValueComponentsFieldName returns the TokenizeValueComponentsFieldName field if non-nil, zero value otherwise.

### GetTokenizeValueComponentsFieldNameOk

`func (o *AddLogFieldBehavior200Response) GetTokenizeValueComponentsFieldNameOk() (*[]string, bool)`

GetTokenizeValueComponentsFieldNameOk returns a tuple with the TokenizeValueComponentsFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeValueComponentsFieldName

`func (o *AddLogFieldBehavior200Response) SetTokenizeValueComponentsFieldName(v []string)`

SetTokenizeValueComponentsFieldName sets TokenizeValueComponentsFieldName field to given value.

### HasTokenizeValueComponentsFieldName

`func (o *AddLogFieldBehavior200Response) HasTokenizeValueComponentsFieldName() bool`

HasTokenizeValueComponentsFieldName returns a boolean if a field has been set.

### GetDescription

`func (o *AddLogFieldBehavior200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLogFieldBehavior200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLogFieldBehavior200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLogFieldBehavior200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultBehavior

`func (o *AddLogFieldBehavior200Response) GetDefaultBehavior() EnumlogFieldBehaviorDefaultBehaviorProp`

GetDefaultBehavior returns the DefaultBehavior field if non-nil, zero value otherwise.

### GetDefaultBehaviorOk

`func (o *AddLogFieldBehavior200Response) GetDefaultBehaviorOk() (*EnumlogFieldBehaviorDefaultBehaviorProp, bool)`

GetDefaultBehaviorOk returns a tuple with the DefaultBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBehavior

`func (o *AddLogFieldBehavior200Response) SetDefaultBehavior(v EnumlogFieldBehaviorDefaultBehaviorProp)`

SetDefaultBehavior sets DefaultBehavior field to given value.

### HasDefaultBehavior

`func (o *AddLogFieldBehavior200Response) HasDefaultBehavior() bool`

HasDefaultBehavior returns a boolean if a field has been set.

### GetMeta

`func (o *AddLogFieldBehavior200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddLogFieldBehavior200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddLogFieldBehavior200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddLogFieldBehavior200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddLogFieldBehavior200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddLogFieldBehavior200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddLogFieldBehavior200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddLogFieldBehavior200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


