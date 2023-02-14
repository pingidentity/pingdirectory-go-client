# AddJsonFieldConstraintsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonField** | **string** | The full name of the JSON field to which these constraints apply. | 
**Schemas** | Pointer to [**[]EnumjsonFieldConstraintsSchemaUrn**](EnumjsonFieldConstraintsSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this JSON Field Constraints | [optional] 
**ValueType** | [**EnumjsonFieldConstraintsValueTypeProp**](EnumjsonFieldConstraintsValueTypeProp.md) |  | 
**IsRequired** | Pointer to **bool** | Indicates whether the target field must be present in JSON objects stored as values of the associated attribute type. | [optional] 
**IsArray** | Pointer to [**EnumjsonFieldConstraintsIsArrayProp**](EnumjsonFieldConstraintsIsArrayProp.md) |  | [optional] 
**AllowNullValue** | Pointer to **bool** | Indicates whether the target field may have a value that is the JSON null value as an alternative to a value (or array of values) of the specified value-type. | [optional] 
**AllowEmptyObject** | Pointer to **bool** | Indicates whether the target field may have a value that is an empty JSON object (i.e., a JSON object with zero fields). This may only be set to true if value-type property is set to object. | [optional] 
**IndexValues** | Pointer to **bool** | Indicates whether backends that support JSON indexing should maintain an index for values of the target field. | [optional] 
**IndexEntryLimit** | Pointer to **int32** | The maximum number of entries that may contain a particular value for the target field before the server will stop maintaining the index for that value. | [optional] 
**PrimeIndex** | Pointer to **bool** | Indicates whether backends that support database priming should load the contents of the associated JSON index into memory whenever the backend is opened. | [optional] 
**CacheMode** | Pointer to [**EnumjsonFieldConstraintsCacheModeProp**](EnumjsonFieldConstraintsCacheModeProp.md) |  | [optional] 
**TokenizeValues** | Pointer to **bool** | Indicates whether the backend should attempt to assign a compact token for each distinct value for the target field in an attempt to reduce the encoded size of the field in JSON objects. These tokens would be assigned prior to using any from the token set used for automatic compaction of some JSON string values. | [optional] 
**AllowedValue** | Pointer to **[]string** | Specifies an explicit set of string values that will be the only values permitted for the target field. If a set of allowed values is defined, then the server will reject any attempt to store a JSON object with a value for the target field that is not included in that set. | [optional] 
**AllowedValueRegularExpression** | Pointer to **[]string** | Specifies an explicit set of regular expressions that may be used to restrict the set of values that may be used for the target field. If a set of allowed value regular expressions is defined, then the server will reject any attempt to store a JSON object with a value for the target field that does not match at least one of those regular expressions. | [optional] 
**MinimumNumericValue** | Pointer to **string** | Specifies the smallest numeric value that may be used as the value for the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is less than that minimum numeric value. | [optional] 
**MaximumNumericValue** | Pointer to **string** | Specifies the largest numeric value that may be used as the value for the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is greater than that maximum numeric value. | [optional] 
**MinimumValueLength** | Pointer to **int32** | Specifies the smallest number of characters that may be present in string values of the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is shorter than that minimum value length. | [optional] 
**MaximumValueLength** | Pointer to **int32** | Specifies the largest number of characters that may be present in string values of the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is longer than that maximum value length. | [optional] 
**MinimumValueCount** | Pointer to **int32** | Specifies the smallest number of elements that may be present in an array of values for the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is an array with fewer than this number of elements. | [optional] 
**MaximumValueCount** | Pointer to **int32** | Specifies the largest number of elements that may be present in an array of values for the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is an array with more than this number of elements. | [optional] 

## Methods

### NewAddJsonFieldConstraintsRequest

`func NewAddJsonFieldConstraintsRequest(jsonField string, valueType EnumjsonFieldConstraintsValueTypeProp, ) *AddJsonFieldConstraintsRequest`

NewAddJsonFieldConstraintsRequest instantiates a new AddJsonFieldConstraintsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddJsonFieldConstraintsRequestWithDefaults

`func NewAddJsonFieldConstraintsRequestWithDefaults() *AddJsonFieldConstraintsRequest`

NewAddJsonFieldConstraintsRequestWithDefaults instantiates a new AddJsonFieldConstraintsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonField

`func (o *AddJsonFieldConstraintsRequest) GetJsonField() string`

GetJsonField returns the JsonField field if non-nil, zero value otherwise.

### GetJsonFieldOk

`func (o *AddJsonFieldConstraintsRequest) GetJsonFieldOk() (*string, bool)`

GetJsonFieldOk returns a tuple with the JsonField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonField

`func (o *AddJsonFieldConstraintsRequest) SetJsonField(v string)`

SetJsonField sets JsonField field to given value.


### GetSchemas

`func (o *AddJsonFieldConstraintsRequest) GetSchemas() []EnumjsonFieldConstraintsSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddJsonFieldConstraintsRequest) GetSchemasOk() (*[]EnumjsonFieldConstraintsSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddJsonFieldConstraintsRequest) SetSchemas(v []EnumjsonFieldConstraintsSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddJsonFieldConstraintsRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddJsonFieldConstraintsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddJsonFieldConstraintsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddJsonFieldConstraintsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddJsonFieldConstraintsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValueType

`func (o *AddJsonFieldConstraintsRequest) GetValueType() EnumjsonFieldConstraintsValueTypeProp`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *AddJsonFieldConstraintsRequest) GetValueTypeOk() (*EnumjsonFieldConstraintsValueTypeProp, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *AddJsonFieldConstraintsRequest) SetValueType(v EnumjsonFieldConstraintsValueTypeProp)`

SetValueType sets ValueType field to given value.


### GetIsRequired

`func (o *AddJsonFieldConstraintsRequest) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *AddJsonFieldConstraintsRequest) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *AddJsonFieldConstraintsRequest) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *AddJsonFieldConstraintsRequest) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetIsArray

`func (o *AddJsonFieldConstraintsRequest) GetIsArray() EnumjsonFieldConstraintsIsArrayProp`

GetIsArray returns the IsArray field if non-nil, zero value otherwise.

### GetIsArrayOk

`func (o *AddJsonFieldConstraintsRequest) GetIsArrayOk() (*EnumjsonFieldConstraintsIsArrayProp, bool)`

GetIsArrayOk returns a tuple with the IsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArray

`func (o *AddJsonFieldConstraintsRequest) SetIsArray(v EnumjsonFieldConstraintsIsArrayProp)`

SetIsArray sets IsArray field to given value.

### HasIsArray

`func (o *AddJsonFieldConstraintsRequest) HasIsArray() bool`

HasIsArray returns a boolean if a field has been set.

### GetAllowNullValue

`func (o *AddJsonFieldConstraintsRequest) GetAllowNullValue() bool`

GetAllowNullValue returns the AllowNullValue field if non-nil, zero value otherwise.

### GetAllowNullValueOk

`func (o *AddJsonFieldConstraintsRequest) GetAllowNullValueOk() (*bool, bool)`

GetAllowNullValueOk returns a tuple with the AllowNullValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNullValue

`func (o *AddJsonFieldConstraintsRequest) SetAllowNullValue(v bool)`

SetAllowNullValue sets AllowNullValue field to given value.

### HasAllowNullValue

`func (o *AddJsonFieldConstraintsRequest) HasAllowNullValue() bool`

HasAllowNullValue returns a boolean if a field has been set.

### GetAllowEmptyObject

`func (o *AddJsonFieldConstraintsRequest) GetAllowEmptyObject() bool`

GetAllowEmptyObject returns the AllowEmptyObject field if non-nil, zero value otherwise.

### GetAllowEmptyObjectOk

`func (o *AddJsonFieldConstraintsRequest) GetAllowEmptyObjectOk() (*bool, bool)`

GetAllowEmptyObjectOk returns a tuple with the AllowEmptyObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmptyObject

`func (o *AddJsonFieldConstraintsRequest) SetAllowEmptyObject(v bool)`

SetAllowEmptyObject sets AllowEmptyObject field to given value.

### HasAllowEmptyObject

`func (o *AddJsonFieldConstraintsRequest) HasAllowEmptyObject() bool`

HasAllowEmptyObject returns a boolean if a field has been set.

### GetIndexValues

`func (o *AddJsonFieldConstraintsRequest) GetIndexValues() bool`

GetIndexValues returns the IndexValues field if non-nil, zero value otherwise.

### GetIndexValuesOk

`func (o *AddJsonFieldConstraintsRequest) GetIndexValuesOk() (*bool, bool)`

GetIndexValuesOk returns a tuple with the IndexValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexValues

`func (o *AddJsonFieldConstraintsRequest) SetIndexValues(v bool)`

SetIndexValues sets IndexValues field to given value.

### HasIndexValues

`func (o *AddJsonFieldConstraintsRequest) HasIndexValues() bool`

HasIndexValues returns a boolean if a field has been set.

### GetIndexEntryLimit

`func (o *AddJsonFieldConstraintsRequest) GetIndexEntryLimit() int32`

GetIndexEntryLimit returns the IndexEntryLimit field if non-nil, zero value otherwise.

### GetIndexEntryLimitOk

`func (o *AddJsonFieldConstraintsRequest) GetIndexEntryLimitOk() (*int32, bool)`

GetIndexEntryLimitOk returns a tuple with the IndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexEntryLimit

`func (o *AddJsonFieldConstraintsRequest) SetIndexEntryLimit(v int32)`

SetIndexEntryLimit sets IndexEntryLimit field to given value.

### HasIndexEntryLimit

`func (o *AddJsonFieldConstraintsRequest) HasIndexEntryLimit() bool`

HasIndexEntryLimit returns a boolean if a field has been set.

### GetPrimeIndex

`func (o *AddJsonFieldConstraintsRequest) GetPrimeIndex() bool`

GetPrimeIndex returns the PrimeIndex field if non-nil, zero value otherwise.

### GetPrimeIndexOk

`func (o *AddJsonFieldConstraintsRequest) GetPrimeIndexOk() (*bool, bool)`

GetPrimeIndexOk returns a tuple with the PrimeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeIndex

`func (o *AddJsonFieldConstraintsRequest) SetPrimeIndex(v bool)`

SetPrimeIndex sets PrimeIndex field to given value.

### HasPrimeIndex

`func (o *AddJsonFieldConstraintsRequest) HasPrimeIndex() bool`

HasPrimeIndex returns a boolean if a field has been set.

### GetCacheMode

`func (o *AddJsonFieldConstraintsRequest) GetCacheMode() EnumjsonFieldConstraintsCacheModeProp`

GetCacheMode returns the CacheMode field if non-nil, zero value otherwise.

### GetCacheModeOk

`func (o *AddJsonFieldConstraintsRequest) GetCacheModeOk() (*EnumjsonFieldConstraintsCacheModeProp, bool)`

GetCacheModeOk returns a tuple with the CacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMode

`func (o *AddJsonFieldConstraintsRequest) SetCacheMode(v EnumjsonFieldConstraintsCacheModeProp)`

SetCacheMode sets CacheMode field to given value.

### HasCacheMode

`func (o *AddJsonFieldConstraintsRequest) HasCacheMode() bool`

HasCacheMode returns a boolean if a field has been set.

### GetTokenizeValues

`func (o *AddJsonFieldConstraintsRequest) GetTokenizeValues() bool`

GetTokenizeValues returns the TokenizeValues field if non-nil, zero value otherwise.

### GetTokenizeValuesOk

`func (o *AddJsonFieldConstraintsRequest) GetTokenizeValuesOk() (*bool, bool)`

GetTokenizeValuesOk returns a tuple with the TokenizeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeValues

`func (o *AddJsonFieldConstraintsRequest) SetTokenizeValues(v bool)`

SetTokenizeValues sets TokenizeValues field to given value.

### HasTokenizeValues

`func (o *AddJsonFieldConstraintsRequest) HasTokenizeValues() bool`

HasTokenizeValues returns a boolean if a field has been set.

### GetAllowedValue

`func (o *AddJsonFieldConstraintsRequest) GetAllowedValue() []string`

GetAllowedValue returns the AllowedValue field if non-nil, zero value otherwise.

### GetAllowedValueOk

`func (o *AddJsonFieldConstraintsRequest) GetAllowedValueOk() (*[]string, bool)`

GetAllowedValueOk returns a tuple with the AllowedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValue

`func (o *AddJsonFieldConstraintsRequest) SetAllowedValue(v []string)`

SetAllowedValue sets AllowedValue field to given value.

### HasAllowedValue

`func (o *AddJsonFieldConstraintsRequest) HasAllowedValue() bool`

HasAllowedValue returns a boolean if a field has been set.

### GetAllowedValueRegularExpression

`func (o *AddJsonFieldConstraintsRequest) GetAllowedValueRegularExpression() []string`

GetAllowedValueRegularExpression returns the AllowedValueRegularExpression field if non-nil, zero value otherwise.

### GetAllowedValueRegularExpressionOk

`func (o *AddJsonFieldConstraintsRequest) GetAllowedValueRegularExpressionOk() (*[]string, bool)`

GetAllowedValueRegularExpressionOk returns a tuple with the AllowedValueRegularExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValueRegularExpression

`func (o *AddJsonFieldConstraintsRequest) SetAllowedValueRegularExpression(v []string)`

SetAllowedValueRegularExpression sets AllowedValueRegularExpression field to given value.

### HasAllowedValueRegularExpression

`func (o *AddJsonFieldConstraintsRequest) HasAllowedValueRegularExpression() bool`

HasAllowedValueRegularExpression returns a boolean if a field has been set.

### GetMinimumNumericValue

`func (o *AddJsonFieldConstraintsRequest) GetMinimumNumericValue() string`

GetMinimumNumericValue returns the MinimumNumericValue field if non-nil, zero value otherwise.

### GetMinimumNumericValueOk

`func (o *AddJsonFieldConstraintsRequest) GetMinimumNumericValueOk() (*string, bool)`

GetMinimumNumericValueOk returns a tuple with the MinimumNumericValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNumericValue

`func (o *AddJsonFieldConstraintsRequest) SetMinimumNumericValue(v string)`

SetMinimumNumericValue sets MinimumNumericValue field to given value.

### HasMinimumNumericValue

`func (o *AddJsonFieldConstraintsRequest) HasMinimumNumericValue() bool`

HasMinimumNumericValue returns a boolean if a field has been set.

### GetMaximumNumericValue

`func (o *AddJsonFieldConstraintsRequest) GetMaximumNumericValue() string`

GetMaximumNumericValue returns the MaximumNumericValue field if non-nil, zero value otherwise.

### GetMaximumNumericValueOk

`func (o *AddJsonFieldConstraintsRequest) GetMaximumNumericValueOk() (*string, bool)`

GetMaximumNumericValueOk returns a tuple with the MaximumNumericValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNumericValue

`func (o *AddJsonFieldConstraintsRequest) SetMaximumNumericValue(v string)`

SetMaximumNumericValue sets MaximumNumericValue field to given value.

### HasMaximumNumericValue

`func (o *AddJsonFieldConstraintsRequest) HasMaximumNumericValue() bool`

HasMaximumNumericValue returns a boolean if a field has been set.

### GetMinimumValueLength

`func (o *AddJsonFieldConstraintsRequest) GetMinimumValueLength() int32`

GetMinimumValueLength returns the MinimumValueLength field if non-nil, zero value otherwise.

### GetMinimumValueLengthOk

`func (o *AddJsonFieldConstraintsRequest) GetMinimumValueLengthOk() (*int32, bool)`

GetMinimumValueLengthOk returns a tuple with the MinimumValueLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumValueLength

`func (o *AddJsonFieldConstraintsRequest) SetMinimumValueLength(v int32)`

SetMinimumValueLength sets MinimumValueLength field to given value.

### HasMinimumValueLength

`func (o *AddJsonFieldConstraintsRequest) HasMinimumValueLength() bool`

HasMinimumValueLength returns a boolean if a field has been set.

### GetMaximumValueLength

`func (o *AddJsonFieldConstraintsRequest) GetMaximumValueLength() int32`

GetMaximumValueLength returns the MaximumValueLength field if non-nil, zero value otherwise.

### GetMaximumValueLengthOk

`func (o *AddJsonFieldConstraintsRequest) GetMaximumValueLengthOk() (*int32, bool)`

GetMaximumValueLengthOk returns a tuple with the MaximumValueLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValueLength

`func (o *AddJsonFieldConstraintsRequest) SetMaximumValueLength(v int32)`

SetMaximumValueLength sets MaximumValueLength field to given value.

### HasMaximumValueLength

`func (o *AddJsonFieldConstraintsRequest) HasMaximumValueLength() bool`

HasMaximumValueLength returns a boolean if a field has been set.

### GetMinimumValueCount

`func (o *AddJsonFieldConstraintsRequest) GetMinimumValueCount() int32`

GetMinimumValueCount returns the MinimumValueCount field if non-nil, zero value otherwise.

### GetMinimumValueCountOk

`func (o *AddJsonFieldConstraintsRequest) GetMinimumValueCountOk() (*int32, bool)`

GetMinimumValueCountOk returns a tuple with the MinimumValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumValueCount

`func (o *AddJsonFieldConstraintsRequest) SetMinimumValueCount(v int32)`

SetMinimumValueCount sets MinimumValueCount field to given value.

### HasMinimumValueCount

`func (o *AddJsonFieldConstraintsRequest) HasMinimumValueCount() bool`

HasMinimumValueCount returns a boolean if a field has been set.

### GetMaximumValueCount

`func (o *AddJsonFieldConstraintsRequest) GetMaximumValueCount() int32`

GetMaximumValueCount returns the MaximumValueCount field if non-nil, zero value otherwise.

### GetMaximumValueCountOk

`func (o *AddJsonFieldConstraintsRequest) GetMaximumValueCountOk() (*int32, bool)`

GetMaximumValueCountOk returns a tuple with the MaximumValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValueCount

`func (o *AddJsonFieldConstraintsRequest) SetMaximumValueCount(v int32)`

SetMaximumValueCount sets MaximumValueCount field to given value.

### HasMaximumValueCount

`func (o *AddJsonFieldConstraintsRequest) HasMaximumValueCount() bool`

HasMaximumValueCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


