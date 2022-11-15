# JsonFormattedAccessLogFieldBehaviorShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn**](EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn.md) |  | 
**PreserveField** | Pointer to [**[]EnumlogFieldBehaviorPreserveFieldProp**](EnumlogFieldBehaviorPreserveFieldProp.md) |  | [optional] 
**PreserveFieldName** | Pointer to **[]string** |  | [optional] 
**OmitField** | Pointer to [**[]EnumlogFieldBehaviorOmitFieldProp**](EnumlogFieldBehaviorOmitFieldProp.md) |  | [optional] 
**OmitFieldName** | Pointer to **[]string** |  | [optional] 
**RedactEntireValueField** | Pointer to [**[]EnumlogFieldBehaviorRedactEntireValueFieldProp**](EnumlogFieldBehaviorRedactEntireValueFieldProp.md) |  | [optional] 
**RedactEntireValueFieldName** | Pointer to **[]string** |  | [optional] 
**RedactValueComponentsField** | Pointer to [**[]EnumlogFieldBehaviorRedactValueComponentsFieldProp**](EnumlogFieldBehaviorRedactValueComponentsFieldProp.md) |  | [optional] 
**RedactValueComponentsFieldName** | Pointer to **[]string** |  | [optional] 
**TokenizeEntireValueField** | Pointer to [**[]EnumlogFieldBehaviorTokenizeEntireValueFieldProp**](EnumlogFieldBehaviorTokenizeEntireValueFieldProp.md) |  | [optional] 
**TokenizeEntireValueFieldName** | Pointer to **[]string** |  | [optional] 
**TokenizeValueComponentsField** | Pointer to [**[]EnumlogFieldBehaviorTokenizeValueComponentsFieldProp**](EnumlogFieldBehaviorTokenizeValueComponentsFieldProp.md) |  | [optional] 
**TokenizeValueComponentsFieldName** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Log Field Behavior | [optional] 
**DefaultBehavior** | Pointer to [**EnumlogFieldBehaviorDefaultBehaviorProp**](EnumlogFieldBehaviorDefaultBehaviorProp.md) |  | [optional] 

## Methods

### NewJsonFormattedAccessLogFieldBehaviorShared

`func NewJsonFormattedAccessLogFieldBehaviorShared(schemas []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn, ) *JsonFormattedAccessLogFieldBehaviorShared`

NewJsonFormattedAccessLogFieldBehaviorShared instantiates a new JsonFormattedAccessLogFieldBehaviorShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonFormattedAccessLogFieldBehaviorSharedWithDefaults

`func NewJsonFormattedAccessLogFieldBehaviorSharedWithDefaults() *JsonFormattedAccessLogFieldBehaviorShared`

NewJsonFormattedAccessLogFieldBehaviorSharedWithDefaults instantiates a new JsonFormattedAccessLogFieldBehaviorShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetSchemas() []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetSchemasOk() (*[]EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *JsonFormattedAccessLogFieldBehaviorShared) SetSchemas(v []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPreserveField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetPreserveField() []EnumlogFieldBehaviorPreserveFieldProp`

GetPreserveField returns the PreserveField field if non-nil, zero value otherwise.

### GetPreserveFieldOk

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetPreserveFieldOk() (*[]EnumlogFieldBehaviorPreserveFieldProp, bool)`

GetPreserveFieldOk returns a tuple with the PreserveField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) SetPreserveField(v []EnumlogFieldBehaviorPreserveFieldProp)`

SetPreserveField sets PreserveField field to given value.

### HasPreserveField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) HasPreserveField() bool`

HasPreserveField returns a boolean if a field has been set.

### GetPreserveFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetPreserveFieldName() []string`

GetPreserveFieldName returns the PreserveFieldName field if non-nil, zero value otherwise.

### GetPreserveFieldNameOk

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetPreserveFieldNameOk() (*[]string, bool)`

GetPreserveFieldNameOk returns a tuple with the PreserveFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) SetPreserveFieldName(v []string)`

SetPreserveFieldName sets PreserveFieldName field to given value.

### HasPreserveFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) HasPreserveFieldName() bool`

HasPreserveFieldName returns a boolean if a field has been set.

### GetOmitField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetOmitField() []EnumlogFieldBehaviorOmitFieldProp`

GetOmitField returns the OmitField field if non-nil, zero value otherwise.

### GetOmitFieldOk

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetOmitFieldOk() (*[]EnumlogFieldBehaviorOmitFieldProp, bool)`

GetOmitFieldOk returns a tuple with the OmitField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) SetOmitField(v []EnumlogFieldBehaviorOmitFieldProp)`

SetOmitField sets OmitField field to given value.

### HasOmitField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) HasOmitField() bool`

HasOmitField returns a boolean if a field has been set.

### GetOmitFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetOmitFieldName() []string`

GetOmitFieldName returns the OmitFieldName field if non-nil, zero value otherwise.

### GetOmitFieldNameOk

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetOmitFieldNameOk() (*[]string, bool)`

GetOmitFieldNameOk returns a tuple with the OmitFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) SetOmitFieldName(v []string)`

SetOmitFieldName sets OmitFieldName field to given value.

### HasOmitFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) HasOmitFieldName() bool`

HasOmitFieldName returns a boolean if a field has been set.

### GetRedactEntireValueField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetRedactEntireValueField() []EnumlogFieldBehaviorRedactEntireValueFieldProp`

GetRedactEntireValueField returns the RedactEntireValueField field if non-nil, zero value otherwise.

### GetRedactEntireValueFieldOk

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetRedactEntireValueFieldOk() (*[]EnumlogFieldBehaviorRedactEntireValueFieldProp, bool)`

GetRedactEntireValueFieldOk returns a tuple with the RedactEntireValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactEntireValueField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) SetRedactEntireValueField(v []EnumlogFieldBehaviorRedactEntireValueFieldProp)`

SetRedactEntireValueField sets RedactEntireValueField field to given value.

### HasRedactEntireValueField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) HasRedactEntireValueField() bool`

HasRedactEntireValueField returns a boolean if a field has been set.

### GetRedactEntireValueFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetRedactEntireValueFieldName() []string`

GetRedactEntireValueFieldName returns the RedactEntireValueFieldName field if non-nil, zero value otherwise.

### GetRedactEntireValueFieldNameOk

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetRedactEntireValueFieldNameOk() (*[]string, bool)`

GetRedactEntireValueFieldNameOk returns a tuple with the RedactEntireValueFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactEntireValueFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) SetRedactEntireValueFieldName(v []string)`

SetRedactEntireValueFieldName sets RedactEntireValueFieldName field to given value.

### HasRedactEntireValueFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) HasRedactEntireValueFieldName() bool`

HasRedactEntireValueFieldName returns a boolean if a field has been set.

### GetRedactValueComponentsField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetRedactValueComponentsField() []EnumlogFieldBehaviorRedactValueComponentsFieldProp`

GetRedactValueComponentsField returns the RedactValueComponentsField field if non-nil, zero value otherwise.

### GetRedactValueComponentsFieldOk

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetRedactValueComponentsFieldOk() (*[]EnumlogFieldBehaviorRedactValueComponentsFieldProp, bool)`

GetRedactValueComponentsFieldOk returns a tuple with the RedactValueComponentsField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactValueComponentsField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) SetRedactValueComponentsField(v []EnumlogFieldBehaviorRedactValueComponentsFieldProp)`

SetRedactValueComponentsField sets RedactValueComponentsField field to given value.

### HasRedactValueComponentsField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) HasRedactValueComponentsField() bool`

HasRedactValueComponentsField returns a boolean if a field has been set.

### GetRedactValueComponentsFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetRedactValueComponentsFieldName() []string`

GetRedactValueComponentsFieldName returns the RedactValueComponentsFieldName field if non-nil, zero value otherwise.

### GetRedactValueComponentsFieldNameOk

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetRedactValueComponentsFieldNameOk() (*[]string, bool)`

GetRedactValueComponentsFieldNameOk returns a tuple with the RedactValueComponentsFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactValueComponentsFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) SetRedactValueComponentsFieldName(v []string)`

SetRedactValueComponentsFieldName sets RedactValueComponentsFieldName field to given value.

### HasRedactValueComponentsFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) HasRedactValueComponentsFieldName() bool`

HasRedactValueComponentsFieldName returns a boolean if a field has been set.

### GetTokenizeEntireValueField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetTokenizeEntireValueField() []EnumlogFieldBehaviorTokenizeEntireValueFieldProp`

GetTokenizeEntireValueField returns the TokenizeEntireValueField field if non-nil, zero value otherwise.

### GetTokenizeEntireValueFieldOk

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetTokenizeEntireValueFieldOk() (*[]EnumlogFieldBehaviorTokenizeEntireValueFieldProp, bool)`

GetTokenizeEntireValueFieldOk returns a tuple with the TokenizeEntireValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeEntireValueField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) SetTokenizeEntireValueField(v []EnumlogFieldBehaviorTokenizeEntireValueFieldProp)`

SetTokenizeEntireValueField sets TokenizeEntireValueField field to given value.

### HasTokenizeEntireValueField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) HasTokenizeEntireValueField() bool`

HasTokenizeEntireValueField returns a boolean if a field has been set.

### GetTokenizeEntireValueFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetTokenizeEntireValueFieldName() []string`

GetTokenizeEntireValueFieldName returns the TokenizeEntireValueFieldName field if non-nil, zero value otherwise.

### GetTokenizeEntireValueFieldNameOk

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetTokenizeEntireValueFieldNameOk() (*[]string, bool)`

GetTokenizeEntireValueFieldNameOk returns a tuple with the TokenizeEntireValueFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeEntireValueFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) SetTokenizeEntireValueFieldName(v []string)`

SetTokenizeEntireValueFieldName sets TokenizeEntireValueFieldName field to given value.

### HasTokenizeEntireValueFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) HasTokenizeEntireValueFieldName() bool`

HasTokenizeEntireValueFieldName returns a boolean if a field has been set.

### GetTokenizeValueComponentsField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetTokenizeValueComponentsField() []EnumlogFieldBehaviorTokenizeValueComponentsFieldProp`

GetTokenizeValueComponentsField returns the TokenizeValueComponentsField field if non-nil, zero value otherwise.

### GetTokenizeValueComponentsFieldOk

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetTokenizeValueComponentsFieldOk() (*[]EnumlogFieldBehaviorTokenizeValueComponentsFieldProp, bool)`

GetTokenizeValueComponentsFieldOk returns a tuple with the TokenizeValueComponentsField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeValueComponentsField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) SetTokenizeValueComponentsField(v []EnumlogFieldBehaviorTokenizeValueComponentsFieldProp)`

SetTokenizeValueComponentsField sets TokenizeValueComponentsField field to given value.

### HasTokenizeValueComponentsField

`func (o *JsonFormattedAccessLogFieldBehaviorShared) HasTokenizeValueComponentsField() bool`

HasTokenizeValueComponentsField returns a boolean if a field has been set.

### GetTokenizeValueComponentsFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetTokenizeValueComponentsFieldName() []string`

GetTokenizeValueComponentsFieldName returns the TokenizeValueComponentsFieldName field if non-nil, zero value otherwise.

### GetTokenizeValueComponentsFieldNameOk

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetTokenizeValueComponentsFieldNameOk() (*[]string, bool)`

GetTokenizeValueComponentsFieldNameOk returns a tuple with the TokenizeValueComponentsFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeValueComponentsFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) SetTokenizeValueComponentsFieldName(v []string)`

SetTokenizeValueComponentsFieldName sets TokenizeValueComponentsFieldName field to given value.

### HasTokenizeValueComponentsFieldName

`func (o *JsonFormattedAccessLogFieldBehaviorShared) HasTokenizeValueComponentsFieldName() bool`

HasTokenizeValueComponentsFieldName returns a boolean if a field has been set.

### GetDescription

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JsonFormattedAccessLogFieldBehaviorShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JsonFormattedAccessLogFieldBehaviorShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultBehavior

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetDefaultBehavior() EnumlogFieldBehaviorDefaultBehaviorProp`

GetDefaultBehavior returns the DefaultBehavior field if non-nil, zero value otherwise.

### GetDefaultBehaviorOk

`func (o *JsonFormattedAccessLogFieldBehaviorShared) GetDefaultBehaviorOk() (*EnumlogFieldBehaviorDefaultBehaviorProp, bool)`

GetDefaultBehaviorOk returns a tuple with the DefaultBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBehavior

`func (o *JsonFormattedAccessLogFieldBehaviorShared) SetDefaultBehavior(v EnumlogFieldBehaviorDefaultBehaviorProp)`

SetDefaultBehavior sets DefaultBehavior field to given value.

### HasDefaultBehavior

`func (o *JsonFormattedAccessLogFieldBehaviorShared) HasDefaultBehavior() bool`

HasDefaultBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


