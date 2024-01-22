# AddJsonFormattedAccessLogFieldBehaviorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn**](EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn.md) |  | 
**PreserveField** | Pointer to [**[]EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp**](EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp.md) |  | [optional] 
**PreserveFieldName** | Pointer to **[]string** | The names of any custom fields whose values should be preserved. This should generally only be used for fields that are not available through the preserve-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**OmitField** | Pointer to [**[]EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp**](EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp.md) |  | [optional] 
**OmitFieldName** | Pointer to **[]string** | The names of any custom fields that should be omitted from log messages. This should generally only be used for fields that are not available through the omit-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**RedactEntireValueField** | Pointer to [**[]EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp**](EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp.md) |  | [optional] 
**RedactEntireValueFieldName** | Pointer to **[]string** | The names of any custom fields whose values should be completely redacted. This should generally only be used for fields that are not available through the redact-entire-value-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**RedactValueComponentsField** | Pointer to [**[]EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp**](EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp.md) |  | [optional] 
**RedactValueComponentsFieldName** | Pointer to **[]string** | The names of any custom fields for which to redact components within the value. This should generally only be used for fields that are not available through the redact-value-components-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**TokenizeEntireValueField** | Pointer to [**[]EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp**](EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp.md) |  | [optional] 
**TokenizeEntireValueFieldName** | Pointer to **[]string** | The names of any custom fields whose values should be completely tokenized. This should generally only be used for fields that are not available through the tokenize-entire-value-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**TokenizeValueComponentsField** | Pointer to [**[]EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp**](EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp.md) |  | [optional] 
**TokenizeValueComponentsFieldName** | Pointer to **[]string** | The names of any custom fields for which to tokenize components within the value. This should generally only be used for fields that are not available through the tokenize-value-components-field property (for example, custom log fields defined in Server SDK extensions). | [optional] 
**Description** | Pointer to **string** | A description for this Log Field Behavior | [optional] 
**DefaultBehavior** | Pointer to [**EnumlogFieldBehaviorDefaultBehaviorProp**](EnumlogFieldBehaviorDefaultBehaviorProp.md) |  | [optional] 
**BehaviorName** | **string** | Name of the new Log Field Behavior | 

## Methods

### NewAddJsonFormattedAccessLogFieldBehaviorRequest

`func NewAddJsonFormattedAccessLogFieldBehaviorRequest(schemas []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn, behaviorName string, ) *AddJsonFormattedAccessLogFieldBehaviorRequest`

NewAddJsonFormattedAccessLogFieldBehaviorRequest instantiates a new AddJsonFormattedAccessLogFieldBehaviorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddJsonFormattedAccessLogFieldBehaviorRequestWithDefaults

`func NewAddJsonFormattedAccessLogFieldBehaviorRequestWithDefaults() *AddJsonFormattedAccessLogFieldBehaviorRequest`

NewAddJsonFormattedAccessLogFieldBehaviorRequestWithDefaults instantiates a new AddJsonFormattedAccessLogFieldBehaviorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetSchemas() []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetSchemasOk() (*[]EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetSchemas(v []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPreserveField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetPreserveField() []EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp`

GetPreserveField returns the PreserveField field if non-nil, zero value otherwise.

### GetPreserveFieldOk

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetPreserveFieldOk() (*[]EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp, bool)`

GetPreserveFieldOk returns a tuple with the PreserveField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetPreserveField(v []EnumlogFieldBehaviorJsonFormattedAccessPreserveFieldProp)`

SetPreserveField sets PreserveField field to given value.

### HasPreserveField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasPreserveField() bool`

HasPreserveField returns a boolean if a field has been set.

### GetPreserveFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetPreserveFieldName() []string`

GetPreserveFieldName returns the PreserveFieldName field if non-nil, zero value otherwise.

### GetPreserveFieldNameOk

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetPreserveFieldNameOk() (*[]string, bool)`

GetPreserveFieldNameOk returns a tuple with the PreserveFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetPreserveFieldName(v []string)`

SetPreserveFieldName sets PreserveFieldName field to given value.

### HasPreserveFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasPreserveFieldName() bool`

HasPreserveFieldName returns a boolean if a field has been set.

### GetOmitField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetOmitField() []EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp`

GetOmitField returns the OmitField field if non-nil, zero value otherwise.

### GetOmitFieldOk

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetOmitFieldOk() (*[]EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp, bool)`

GetOmitFieldOk returns a tuple with the OmitField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetOmitField(v []EnumlogFieldBehaviorJsonFormattedAccessOmitFieldProp)`

SetOmitField sets OmitField field to given value.

### HasOmitField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasOmitField() bool`

HasOmitField returns a boolean if a field has been set.

### GetOmitFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetOmitFieldName() []string`

GetOmitFieldName returns the OmitFieldName field if non-nil, zero value otherwise.

### GetOmitFieldNameOk

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetOmitFieldNameOk() (*[]string, bool)`

GetOmitFieldNameOk returns a tuple with the OmitFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetOmitFieldName(v []string)`

SetOmitFieldName sets OmitFieldName field to given value.

### HasOmitFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasOmitFieldName() bool`

HasOmitFieldName returns a boolean if a field has been set.

### GetRedactEntireValueField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetRedactEntireValueField() []EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp`

GetRedactEntireValueField returns the RedactEntireValueField field if non-nil, zero value otherwise.

### GetRedactEntireValueFieldOk

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetRedactEntireValueFieldOk() (*[]EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp, bool)`

GetRedactEntireValueFieldOk returns a tuple with the RedactEntireValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactEntireValueField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetRedactEntireValueField(v []EnumlogFieldBehaviorJsonFormattedAccessRedactEntireValueFieldProp)`

SetRedactEntireValueField sets RedactEntireValueField field to given value.

### HasRedactEntireValueField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasRedactEntireValueField() bool`

HasRedactEntireValueField returns a boolean if a field has been set.

### GetRedactEntireValueFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetRedactEntireValueFieldName() []string`

GetRedactEntireValueFieldName returns the RedactEntireValueFieldName field if non-nil, zero value otherwise.

### GetRedactEntireValueFieldNameOk

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetRedactEntireValueFieldNameOk() (*[]string, bool)`

GetRedactEntireValueFieldNameOk returns a tuple with the RedactEntireValueFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactEntireValueFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetRedactEntireValueFieldName(v []string)`

SetRedactEntireValueFieldName sets RedactEntireValueFieldName field to given value.

### HasRedactEntireValueFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasRedactEntireValueFieldName() bool`

HasRedactEntireValueFieldName returns a boolean if a field has been set.

### GetRedactValueComponentsField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetRedactValueComponentsField() []EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp`

GetRedactValueComponentsField returns the RedactValueComponentsField field if non-nil, zero value otherwise.

### GetRedactValueComponentsFieldOk

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetRedactValueComponentsFieldOk() (*[]EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp, bool)`

GetRedactValueComponentsFieldOk returns a tuple with the RedactValueComponentsField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactValueComponentsField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetRedactValueComponentsField(v []EnumlogFieldBehaviorJsonFormattedAccessRedactValueComponentsFieldProp)`

SetRedactValueComponentsField sets RedactValueComponentsField field to given value.

### HasRedactValueComponentsField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasRedactValueComponentsField() bool`

HasRedactValueComponentsField returns a boolean if a field has been set.

### GetRedactValueComponentsFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetRedactValueComponentsFieldName() []string`

GetRedactValueComponentsFieldName returns the RedactValueComponentsFieldName field if non-nil, zero value otherwise.

### GetRedactValueComponentsFieldNameOk

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetRedactValueComponentsFieldNameOk() (*[]string, bool)`

GetRedactValueComponentsFieldNameOk returns a tuple with the RedactValueComponentsFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactValueComponentsFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetRedactValueComponentsFieldName(v []string)`

SetRedactValueComponentsFieldName sets RedactValueComponentsFieldName field to given value.

### HasRedactValueComponentsFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasRedactValueComponentsFieldName() bool`

HasRedactValueComponentsFieldName returns a boolean if a field has been set.

### GetTokenizeEntireValueField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetTokenizeEntireValueField() []EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp`

GetTokenizeEntireValueField returns the TokenizeEntireValueField field if non-nil, zero value otherwise.

### GetTokenizeEntireValueFieldOk

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetTokenizeEntireValueFieldOk() (*[]EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp, bool)`

GetTokenizeEntireValueFieldOk returns a tuple with the TokenizeEntireValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeEntireValueField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetTokenizeEntireValueField(v []EnumlogFieldBehaviorJsonFormattedAccessTokenizeEntireValueFieldProp)`

SetTokenizeEntireValueField sets TokenizeEntireValueField field to given value.

### HasTokenizeEntireValueField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasTokenizeEntireValueField() bool`

HasTokenizeEntireValueField returns a boolean if a field has been set.

### GetTokenizeEntireValueFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetTokenizeEntireValueFieldName() []string`

GetTokenizeEntireValueFieldName returns the TokenizeEntireValueFieldName field if non-nil, zero value otherwise.

### GetTokenizeEntireValueFieldNameOk

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetTokenizeEntireValueFieldNameOk() (*[]string, bool)`

GetTokenizeEntireValueFieldNameOk returns a tuple with the TokenizeEntireValueFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeEntireValueFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetTokenizeEntireValueFieldName(v []string)`

SetTokenizeEntireValueFieldName sets TokenizeEntireValueFieldName field to given value.

### HasTokenizeEntireValueFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasTokenizeEntireValueFieldName() bool`

HasTokenizeEntireValueFieldName returns a boolean if a field has been set.

### GetTokenizeValueComponentsField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetTokenizeValueComponentsField() []EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp`

GetTokenizeValueComponentsField returns the TokenizeValueComponentsField field if non-nil, zero value otherwise.

### GetTokenizeValueComponentsFieldOk

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetTokenizeValueComponentsFieldOk() (*[]EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp, bool)`

GetTokenizeValueComponentsFieldOk returns a tuple with the TokenizeValueComponentsField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeValueComponentsField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetTokenizeValueComponentsField(v []EnumlogFieldBehaviorJsonFormattedAccessTokenizeValueComponentsFieldProp)`

SetTokenizeValueComponentsField sets TokenizeValueComponentsField field to given value.

### HasTokenizeValueComponentsField

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasTokenizeValueComponentsField() bool`

HasTokenizeValueComponentsField returns a boolean if a field has been set.

### GetTokenizeValueComponentsFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetTokenizeValueComponentsFieldName() []string`

GetTokenizeValueComponentsFieldName returns the TokenizeValueComponentsFieldName field if non-nil, zero value otherwise.

### GetTokenizeValueComponentsFieldNameOk

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetTokenizeValueComponentsFieldNameOk() (*[]string, bool)`

GetTokenizeValueComponentsFieldNameOk returns a tuple with the TokenizeValueComponentsFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeValueComponentsFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetTokenizeValueComponentsFieldName(v []string)`

SetTokenizeValueComponentsFieldName sets TokenizeValueComponentsFieldName field to given value.

### HasTokenizeValueComponentsFieldName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasTokenizeValueComponentsFieldName() bool`

HasTokenizeValueComponentsFieldName returns a boolean if a field has been set.

### GetDescription

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultBehavior

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetDefaultBehavior() EnumlogFieldBehaviorDefaultBehaviorProp`

GetDefaultBehavior returns the DefaultBehavior field if non-nil, zero value otherwise.

### GetDefaultBehaviorOk

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetDefaultBehaviorOk() (*EnumlogFieldBehaviorDefaultBehaviorProp, bool)`

GetDefaultBehaviorOk returns a tuple with the DefaultBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBehavior

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetDefaultBehavior(v EnumlogFieldBehaviorDefaultBehaviorProp)`

SetDefaultBehavior sets DefaultBehavior field to given value.

### HasDefaultBehavior

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) HasDefaultBehavior() bool`

HasDefaultBehavior returns a boolean if a field has been set.

### GetBehaviorName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetBehaviorName() string`

GetBehaviorName returns the BehaviorName field if non-nil, zero value otherwise.

### GetBehaviorNameOk

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) GetBehaviorNameOk() (*string, bool)`

GetBehaviorNameOk returns a tuple with the BehaviorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviorName

`func (o *AddJsonFormattedAccessLogFieldBehaviorRequest) SetBehaviorName(v string)`

SetBehaviorName sets BehaviorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


