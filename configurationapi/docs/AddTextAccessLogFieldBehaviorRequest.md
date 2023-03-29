# AddTextAccessLogFieldBehaviorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BehaviorName** | **string** | Name of the new Log Field Behavior | 
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

### NewAddTextAccessLogFieldBehaviorRequest

`func NewAddTextAccessLogFieldBehaviorRequest(behaviorName string, schemas []EnumtextAccessLogFieldBehaviorSchemaUrn, ) *AddTextAccessLogFieldBehaviorRequest`

NewAddTextAccessLogFieldBehaviorRequest instantiates a new AddTextAccessLogFieldBehaviorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTextAccessLogFieldBehaviorRequestWithDefaults

`func NewAddTextAccessLogFieldBehaviorRequestWithDefaults() *AddTextAccessLogFieldBehaviorRequest`

NewAddTextAccessLogFieldBehaviorRequestWithDefaults instantiates a new AddTextAccessLogFieldBehaviorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehaviorName

`func (o *AddTextAccessLogFieldBehaviorRequest) GetBehaviorName() string`

GetBehaviorName returns the BehaviorName field if non-nil, zero value otherwise.

### GetBehaviorNameOk

`func (o *AddTextAccessLogFieldBehaviorRequest) GetBehaviorNameOk() (*string, bool)`

GetBehaviorNameOk returns a tuple with the BehaviorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviorName

`func (o *AddTextAccessLogFieldBehaviorRequest) SetBehaviorName(v string)`

SetBehaviorName sets BehaviorName field to given value.


### GetSchemas

`func (o *AddTextAccessLogFieldBehaviorRequest) GetSchemas() []EnumtextAccessLogFieldBehaviorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddTextAccessLogFieldBehaviorRequest) GetSchemasOk() (*[]EnumtextAccessLogFieldBehaviorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddTextAccessLogFieldBehaviorRequest) SetSchemas(v []EnumtextAccessLogFieldBehaviorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPreserveField

`func (o *AddTextAccessLogFieldBehaviorRequest) GetPreserveField() []EnumlogFieldBehaviorTextAccessPreserveFieldProp`

GetPreserveField returns the PreserveField field if non-nil, zero value otherwise.

### GetPreserveFieldOk

`func (o *AddTextAccessLogFieldBehaviorRequest) GetPreserveFieldOk() (*[]EnumlogFieldBehaviorTextAccessPreserveFieldProp, bool)`

GetPreserveFieldOk returns a tuple with the PreserveField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveField

`func (o *AddTextAccessLogFieldBehaviorRequest) SetPreserveField(v []EnumlogFieldBehaviorTextAccessPreserveFieldProp)`

SetPreserveField sets PreserveField field to given value.

### HasPreserveField

`func (o *AddTextAccessLogFieldBehaviorRequest) HasPreserveField() bool`

HasPreserveField returns a boolean if a field has been set.

### GetPreserveFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) GetPreserveFieldName() []string`

GetPreserveFieldName returns the PreserveFieldName field if non-nil, zero value otherwise.

### GetPreserveFieldNameOk

`func (o *AddTextAccessLogFieldBehaviorRequest) GetPreserveFieldNameOk() (*[]string, bool)`

GetPreserveFieldNameOk returns a tuple with the PreserveFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) SetPreserveFieldName(v []string)`

SetPreserveFieldName sets PreserveFieldName field to given value.

### HasPreserveFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) HasPreserveFieldName() bool`

HasPreserveFieldName returns a boolean if a field has been set.

### GetOmitField

`func (o *AddTextAccessLogFieldBehaviorRequest) GetOmitField() []EnumlogFieldBehaviorTextAccessOmitFieldProp`

GetOmitField returns the OmitField field if non-nil, zero value otherwise.

### GetOmitFieldOk

`func (o *AddTextAccessLogFieldBehaviorRequest) GetOmitFieldOk() (*[]EnumlogFieldBehaviorTextAccessOmitFieldProp, bool)`

GetOmitFieldOk returns a tuple with the OmitField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitField

`func (o *AddTextAccessLogFieldBehaviorRequest) SetOmitField(v []EnumlogFieldBehaviorTextAccessOmitFieldProp)`

SetOmitField sets OmitField field to given value.

### HasOmitField

`func (o *AddTextAccessLogFieldBehaviorRequest) HasOmitField() bool`

HasOmitField returns a boolean if a field has been set.

### GetOmitFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) GetOmitFieldName() []string`

GetOmitFieldName returns the OmitFieldName field if non-nil, zero value otherwise.

### GetOmitFieldNameOk

`func (o *AddTextAccessLogFieldBehaviorRequest) GetOmitFieldNameOk() (*[]string, bool)`

GetOmitFieldNameOk returns a tuple with the OmitFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) SetOmitFieldName(v []string)`

SetOmitFieldName sets OmitFieldName field to given value.

### HasOmitFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) HasOmitFieldName() bool`

HasOmitFieldName returns a boolean if a field has been set.

### GetRedactEntireValueField

`func (o *AddTextAccessLogFieldBehaviorRequest) GetRedactEntireValueField() []EnumlogFieldBehaviorTextAccessRedactEntireValueFieldProp`

GetRedactEntireValueField returns the RedactEntireValueField field if non-nil, zero value otherwise.

### GetRedactEntireValueFieldOk

`func (o *AddTextAccessLogFieldBehaviorRequest) GetRedactEntireValueFieldOk() (*[]EnumlogFieldBehaviorTextAccessRedactEntireValueFieldProp, bool)`

GetRedactEntireValueFieldOk returns a tuple with the RedactEntireValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactEntireValueField

`func (o *AddTextAccessLogFieldBehaviorRequest) SetRedactEntireValueField(v []EnumlogFieldBehaviorTextAccessRedactEntireValueFieldProp)`

SetRedactEntireValueField sets RedactEntireValueField field to given value.

### HasRedactEntireValueField

`func (o *AddTextAccessLogFieldBehaviorRequest) HasRedactEntireValueField() bool`

HasRedactEntireValueField returns a boolean if a field has been set.

### GetRedactEntireValueFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) GetRedactEntireValueFieldName() []string`

GetRedactEntireValueFieldName returns the RedactEntireValueFieldName field if non-nil, zero value otherwise.

### GetRedactEntireValueFieldNameOk

`func (o *AddTextAccessLogFieldBehaviorRequest) GetRedactEntireValueFieldNameOk() (*[]string, bool)`

GetRedactEntireValueFieldNameOk returns a tuple with the RedactEntireValueFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactEntireValueFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) SetRedactEntireValueFieldName(v []string)`

SetRedactEntireValueFieldName sets RedactEntireValueFieldName field to given value.

### HasRedactEntireValueFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) HasRedactEntireValueFieldName() bool`

HasRedactEntireValueFieldName returns a boolean if a field has been set.

### GetRedactValueComponentsField

`func (o *AddTextAccessLogFieldBehaviorRequest) GetRedactValueComponentsField() []EnumlogFieldBehaviorTextAccessRedactValueComponentsFieldProp`

GetRedactValueComponentsField returns the RedactValueComponentsField field if non-nil, zero value otherwise.

### GetRedactValueComponentsFieldOk

`func (o *AddTextAccessLogFieldBehaviorRequest) GetRedactValueComponentsFieldOk() (*[]EnumlogFieldBehaviorTextAccessRedactValueComponentsFieldProp, bool)`

GetRedactValueComponentsFieldOk returns a tuple with the RedactValueComponentsField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactValueComponentsField

`func (o *AddTextAccessLogFieldBehaviorRequest) SetRedactValueComponentsField(v []EnumlogFieldBehaviorTextAccessRedactValueComponentsFieldProp)`

SetRedactValueComponentsField sets RedactValueComponentsField field to given value.

### HasRedactValueComponentsField

`func (o *AddTextAccessLogFieldBehaviorRequest) HasRedactValueComponentsField() bool`

HasRedactValueComponentsField returns a boolean if a field has been set.

### GetRedactValueComponentsFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) GetRedactValueComponentsFieldName() []string`

GetRedactValueComponentsFieldName returns the RedactValueComponentsFieldName field if non-nil, zero value otherwise.

### GetRedactValueComponentsFieldNameOk

`func (o *AddTextAccessLogFieldBehaviorRequest) GetRedactValueComponentsFieldNameOk() (*[]string, bool)`

GetRedactValueComponentsFieldNameOk returns a tuple with the RedactValueComponentsFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactValueComponentsFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) SetRedactValueComponentsFieldName(v []string)`

SetRedactValueComponentsFieldName sets RedactValueComponentsFieldName field to given value.

### HasRedactValueComponentsFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) HasRedactValueComponentsFieldName() bool`

HasRedactValueComponentsFieldName returns a boolean if a field has been set.

### GetTokenizeEntireValueField

`func (o *AddTextAccessLogFieldBehaviorRequest) GetTokenizeEntireValueField() []EnumlogFieldBehaviorTextAccessTokenizeEntireValueFieldProp`

GetTokenizeEntireValueField returns the TokenizeEntireValueField field if non-nil, zero value otherwise.

### GetTokenizeEntireValueFieldOk

`func (o *AddTextAccessLogFieldBehaviorRequest) GetTokenizeEntireValueFieldOk() (*[]EnumlogFieldBehaviorTextAccessTokenizeEntireValueFieldProp, bool)`

GetTokenizeEntireValueFieldOk returns a tuple with the TokenizeEntireValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeEntireValueField

`func (o *AddTextAccessLogFieldBehaviorRequest) SetTokenizeEntireValueField(v []EnumlogFieldBehaviorTextAccessTokenizeEntireValueFieldProp)`

SetTokenizeEntireValueField sets TokenizeEntireValueField field to given value.

### HasTokenizeEntireValueField

`func (o *AddTextAccessLogFieldBehaviorRequest) HasTokenizeEntireValueField() bool`

HasTokenizeEntireValueField returns a boolean if a field has been set.

### GetTokenizeEntireValueFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) GetTokenizeEntireValueFieldName() []string`

GetTokenizeEntireValueFieldName returns the TokenizeEntireValueFieldName field if non-nil, zero value otherwise.

### GetTokenizeEntireValueFieldNameOk

`func (o *AddTextAccessLogFieldBehaviorRequest) GetTokenizeEntireValueFieldNameOk() (*[]string, bool)`

GetTokenizeEntireValueFieldNameOk returns a tuple with the TokenizeEntireValueFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeEntireValueFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) SetTokenizeEntireValueFieldName(v []string)`

SetTokenizeEntireValueFieldName sets TokenizeEntireValueFieldName field to given value.

### HasTokenizeEntireValueFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) HasTokenizeEntireValueFieldName() bool`

HasTokenizeEntireValueFieldName returns a boolean if a field has been set.

### GetTokenizeValueComponentsField

`func (o *AddTextAccessLogFieldBehaviorRequest) GetTokenizeValueComponentsField() []EnumlogFieldBehaviorTextAccessTokenizeValueComponentsFieldProp`

GetTokenizeValueComponentsField returns the TokenizeValueComponentsField field if non-nil, zero value otherwise.

### GetTokenizeValueComponentsFieldOk

`func (o *AddTextAccessLogFieldBehaviorRequest) GetTokenizeValueComponentsFieldOk() (*[]EnumlogFieldBehaviorTextAccessTokenizeValueComponentsFieldProp, bool)`

GetTokenizeValueComponentsFieldOk returns a tuple with the TokenizeValueComponentsField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeValueComponentsField

`func (o *AddTextAccessLogFieldBehaviorRequest) SetTokenizeValueComponentsField(v []EnumlogFieldBehaviorTextAccessTokenizeValueComponentsFieldProp)`

SetTokenizeValueComponentsField sets TokenizeValueComponentsField field to given value.

### HasTokenizeValueComponentsField

`func (o *AddTextAccessLogFieldBehaviorRequest) HasTokenizeValueComponentsField() bool`

HasTokenizeValueComponentsField returns a boolean if a field has been set.

### GetTokenizeValueComponentsFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) GetTokenizeValueComponentsFieldName() []string`

GetTokenizeValueComponentsFieldName returns the TokenizeValueComponentsFieldName field if non-nil, zero value otherwise.

### GetTokenizeValueComponentsFieldNameOk

`func (o *AddTextAccessLogFieldBehaviorRequest) GetTokenizeValueComponentsFieldNameOk() (*[]string, bool)`

GetTokenizeValueComponentsFieldNameOk returns a tuple with the TokenizeValueComponentsFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeValueComponentsFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) SetTokenizeValueComponentsFieldName(v []string)`

SetTokenizeValueComponentsFieldName sets TokenizeValueComponentsFieldName field to given value.

### HasTokenizeValueComponentsFieldName

`func (o *AddTextAccessLogFieldBehaviorRequest) HasTokenizeValueComponentsFieldName() bool`

HasTokenizeValueComponentsFieldName returns a boolean if a field has been set.

### GetDescription

`func (o *AddTextAccessLogFieldBehaviorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddTextAccessLogFieldBehaviorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddTextAccessLogFieldBehaviorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddTextAccessLogFieldBehaviorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultBehavior

`func (o *AddTextAccessLogFieldBehaviorRequest) GetDefaultBehavior() EnumlogFieldBehaviorDefaultBehaviorProp`

GetDefaultBehavior returns the DefaultBehavior field if non-nil, zero value otherwise.

### GetDefaultBehaviorOk

`func (o *AddTextAccessLogFieldBehaviorRequest) GetDefaultBehaviorOk() (*EnumlogFieldBehaviorDefaultBehaviorProp, bool)`

GetDefaultBehaviorOk returns a tuple with the DefaultBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBehavior

`func (o *AddTextAccessLogFieldBehaviorRequest) SetDefaultBehavior(v EnumlogFieldBehaviorDefaultBehaviorProp)`

SetDefaultBehavior sets DefaultBehavior field to given value.

### HasDefaultBehavior

`func (o *AddTextAccessLogFieldBehaviorRequest) HasDefaultBehavior() bool`

HasDefaultBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


