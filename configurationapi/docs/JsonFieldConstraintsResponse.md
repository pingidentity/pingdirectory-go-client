# JsonFieldConstraintsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumjsonFieldConstraintsSchemaUrn**](EnumjsonFieldConstraintsSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this JSON Field Constraints | [optional] 
**JsonField** | **string** | The full name of the JSON field to which these constraints apply. | 
**ValueType** | [**EnumjsonFieldConstraintsValueTypeProp**](EnumjsonFieldConstraintsValueTypeProp.md) |  | 
**IsRequired** | Pointer to **bool** | Indicates whether the target field must be present in JSON objects stored as values of the associated attribute type. | [optional] 
**IsArray** | Pointer to [**EnumjsonFieldConstraintsIsArrayProp**](EnumjsonFieldConstraintsIsArrayProp.md) |  | [optional] 
**AllowNullValue** | Pointer to **bool** | Indicates whether the target field may have a value that is the JSON null value as an alternative to a value (or array of values) of the specified value-type. | [optional] 
**AllowEmptyObject** | Pointer to **bool** | Indicates whether the target field may have a value that is an empty JSON object (i.e., a JSON object with zero fields). This may only be set to true if value-type property is set to object. | [optional] 
**IndexValues** | Pointer to **bool** | Indicates whether backends that support JSON indexing should maintain an index for values of the target field. | [optional] 
**IndexEntryLimit** | Pointer to **int64** | The maximum number of entries that may contain a particular value for the target field before the server will stop maintaining the index for that value. | [optional] 
**PrimeIndex** | Pointer to **bool** | Indicates whether backends that support database priming should load the contents of the associated JSON index into memory whenever the backend is opened. | [optional] 
**CacheMode** | Pointer to [**EnumjsonFieldConstraintsCacheModeProp**](EnumjsonFieldConstraintsCacheModeProp.md) |  | [optional] 
**TokenizeValues** | Pointer to **bool** | Indicates whether the backend should attempt to assign a compact token for each distinct value for the target field in an attempt to reduce the encoded size of the field in JSON objects. These tokens would be assigned prior to using any from the token set used for automatic compaction of some JSON string values. | [optional] 
**AllowedValue** | Pointer to **[]string** | Specifies an explicit set of string values that will be the only values permitted for the target field. If a set of allowed values is defined, then the server will reject any attempt to store a JSON object with a value for the target field that is not included in that set. | [optional] 
**AllowedValueRegularExpression** | Pointer to **[]string** | Specifies an explicit set of regular expressions that may be used to restrict the set of values that may be used for the target field. If a set of allowed value regular expressions is defined, then the server will reject any attempt to store a JSON object with a value for the target field that does not match at least one of those regular expressions. | [optional] 
**MinimumNumericValue** | Pointer to **string** | Specifies the smallest numeric value that may be used as the value for the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is less than that minimum numeric value. | [optional] 
**MaximumNumericValue** | Pointer to **string** | Specifies the largest numeric value that may be used as the value for the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is greater than that maximum numeric value. | [optional] 
**MinimumValueLength** | Pointer to **int64** | Specifies the smallest number of characters that may be present in string values of the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is shorter than that minimum value length. | [optional] 
**MaximumValueLength** | Pointer to **int64** | Specifies the largest number of characters that may be present in string values of the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is longer than that maximum value length. | [optional] 
**MinimumValueCount** | Pointer to **int64** | Specifies the smallest number of elements that may be present in an array of values for the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is an array with fewer than this number of elements. | [optional] 
**MaximumValueCount** | Pointer to **int64** | Specifies the largest number of elements that may be present in an array of values for the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is an array with more than this number of elements. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the JSON Field Constraints | 

## Methods

### NewJsonFieldConstraintsResponse

`func NewJsonFieldConstraintsResponse(jsonField string, valueType EnumjsonFieldConstraintsValueTypeProp, id string, ) *JsonFieldConstraintsResponse`

NewJsonFieldConstraintsResponse instantiates a new JsonFieldConstraintsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonFieldConstraintsResponseWithDefaults

`func NewJsonFieldConstraintsResponseWithDefaults() *JsonFieldConstraintsResponse`

NewJsonFieldConstraintsResponseWithDefaults instantiates a new JsonFieldConstraintsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *JsonFieldConstraintsResponse) GetSchemas() []EnumjsonFieldConstraintsSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *JsonFieldConstraintsResponse) GetSchemasOk() (*[]EnumjsonFieldConstraintsSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *JsonFieldConstraintsResponse) SetSchemas(v []EnumjsonFieldConstraintsSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *JsonFieldConstraintsResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *JsonFieldConstraintsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JsonFieldConstraintsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JsonFieldConstraintsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JsonFieldConstraintsResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJsonField

`func (o *JsonFieldConstraintsResponse) GetJsonField() string`

GetJsonField returns the JsonField field if non-nil, zero value otherwise.

### GetJsonFieldOk

`func (o *JsonFieldConstraintsResponse) GetJsonFieldOk() (*string, bool)`

GetJsonFieldOk returns a tuple with the JsonField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonField

`func (o *JsonFieldConstraintsResponse) SetJsonField(v string)`

SetJsonField sets JsonField field to given value.


### GetValueType

`func (o *JsonFieldConstraintsResponse) GetValueType() EnumjsonFieldConstraintsValueTypeProp`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *JsonFieldConstraintsResponse) GetValueTypeOk() (*EnumjsonFieldConstraintsValueTypeProp, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *JsonFieldConstraintsResponse) SetValueType(v EnumjsonFieldConstraintsValueTypeProp)`

SetValueType sets ValueType field to given value.


### GetIsRequired

`func (o *JsonFieldConstraintsResponse) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *JsonFieldConstraintsResponse) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *JsonFieldConstraintsResponse) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *JsonFieldConstraintsResponse) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetIsArray

`func (o *JsonFieldConstraintsResponse) GetIsArray() EnumjsonFieldConstraintsIsArrayProp`

GetIsArray returns the IsArray field if non-nil, zero value otherwise.

### GetIsArrayOk

`func (o *JsonFieldConstraintsResponse) GetIsArrayOk() (*EnumjsonFieldConstraintsIsArrayProp, bool)`

GetIsArrayOk returns a tuple with the IsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArray

`func (o *JsonFieldConstraintsResponse) SetIsArray(v EnumjsonFieldConstraintsIsArrayProp)`

SetIsArray sets IsArray field to given value.

### HasIsArray

`func (o *JsonFieldConstraintsResponse) HasIsArray() bool`

HasIsArray returns a boolean if a field has been set.

### GetAllowNullValue

`func (o *JsonFieldConstraintsResponse) GetAllowNullValue() bool`

GetAllowNullValue returns the AllowNullValue field if non-nil, zero value otherwise.

### GetAllowNullValueOk

`func (o *JsonFieldConstraintsResponse) GetAllowNullValueOk() (*bool, bool)`

GetAllowNullValueOk returns a tuple with the AllowNullValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNullValue

`func (o *JsonFieldConstraintsResponse) SetAllowNullValue(v bool)`

SetAllowNullValue sets AllowNullValue field to given value.

### HasAllowNullValue

`func (o *JsonFieldConstraintsResponse) HasAllowNullValue() bool`

HasAllowNullValue returns a boolean if a field has been set.

### GetAllowEmptyObject

`func (o *JsonFieldConstraintsResponse) GetAllowEmptyObject() bool`

GetAllowEmptyObject returns the AllowEmptyObject field if non-nil, zero value otherwise.

### GetAllowEmptyObjectOk

`func (o *JsonFieldConstraintsResponse) GetAllowEmptyObjectOk() (*bool, bool)`

GetAllowEmptyObjectOk returns a tuple with the AllowEmptyObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmptyObject

`func (o *JsonFieldConstraintsResponse) SetAllowEmptyObject(v bool)`

SetAllowEmptyObject sets AllowEmptyObject field to given value.

### HasAllowEmptyObject

`func (o *JsonFieldConstraintsResponse) HasAllowEmptyObject() bool`

HasAllowEmptyObject returns a boolean if a field has been set.

### GetIndexValues

`func (o *JsonFieldConstraintsResponse) GetIndexValues() bool`

GetIndexValues returns the IndexValues field if non-nil, zero value otherwise.

### GetIndexValuesOk

`func (o *JsonFieldConstraintsResponse) GetIndexValuesOk() (*bool, bool)`

GetIndexValuesOk returns a tuple with the IndexValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexValues

`func (o *JsonFieldConstraintsResponse) SetIndexValues(v bool)`

SetIndexValues sets IndexValues field to given value.

### HasIndexValues

`func (o *JsonFieldConstraintsResponse) HasIndexValues() bool`

HasIndexValues returns a boolean if a field has been set.

### GetIndexEntryLimit

`func (o *JsonFieldConstraintsResponse) GetIndexEntryLimit() int64`

GetIndexEntryLimit returns the IndexEntryLimit field if non-nil, zero value otherwise.

### GetIndexEntryLimitOk

`func (o *JsonFieldConstraintsResponse) GetIndexEntryLimitOk() (*int64, bool)`

GetIndexEntryLimitOk returns a tuple with the IndexEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexEntryLimit

`func (o *JsonFieldConstraintsResponse) SetIndexEntryLimit(v int64)`

SetIndexEntryLimit sets IndexEntryLimit field to given value.

### HasIndexEntryLimit

`func (o *JsonFieldConstraintsResponse) HasIndexEntryLimit() bool`

HasIndexEntryLimit returns a boolean if a field has been set.

### GetPrimeIndex

`func (o *JsonFieldConstraintsResponse) GetPrimeIndex() bool`

GetPrimeIndex returns the PrimeIndex field if non-nil, zero value otherwise.

### GetPrimeIndexOk

`func (o *JsonFieldConstraintsResponse) GetPrimeIndexOk() (*bool, bool)`

GetPrimeIndexOk returns a tuple with the PrimeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeIndex

`func (o *JsonFieldConstraintsResponse) SetPrimeIndex(v bool)`

SetPrimeIndex sets PrimeIndex field to given value.

### HasPrimeIndex

`func (o *JsonFieldConstraintsResponse) HasPrimeIndex() bool`

HasPrimeIndex returns a boolean if a field has been set.

### GetCacheMode

`func (o *JsonFieldConstraintsResponse) GetCacheMode() EnumjsonFieldConstraintsCacheModeProp`

GetCacheMode returns the CacheMode field if non-nil, zero value otherwise.

### GetCacheModeOk

`func (o *JsonFieldConstraintsResponse) GetCacheModeOk() (*EnumjsonFieldConstraintsCacheModeProp, bool)`

GetCacheModeOk returns a tuple with the CacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMode

`func (o *JsonFieldConstraintsResponse) SetCacheMode(v EnumjsonFieldConstraintsCacheModeProp)`

SetCacheMode sets CacheMode field to given value.

### HasCacheMode

`func (o *JsonFieldConstraintsResponse) HasCacheMode() bool`

HasCacheMode returns a boolean if a field has been set.

### GetTokenizeValues

`func (o *JsonFieldConstraintsResponse) GetTokenizeValues() bool`

GetTokenizeValues returns the TokenizeValues field if non-nil, zero value otherwise.

### GetTokenizeValuesOk

`func (o *JsonFieldConstraintsResponse) GetTokenizeValuesOk() (*bool, bool)`

GetTokenizeValuesOk returns a tuple with the TokenizeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizeValues

`func (o *JsonFieldConstraintsResponse) SetTokenizeValues(v bool)`

SetTokenizeValues sets TokenizeValues field to given value.

### HasTokenizeValues

`func (o *JsonFieldConstraintsResponse) HasTokenizeValues() bool`

HasTokenizeValues returns a boolean if a field has been set.

### GetAllowedValue

`func (o *JsonFieldConstraintsResponse) GetAllowedValue() []string`

GetAllowedValue returns the AllowedValue field if non-nil, zero value otherwise.

### GetAllowedValueOk

`func (o *JsonFieldConstraintsResponse) GetAllowedValueOk() (*[]string, bool)`

GetAllowedValueOk returns a tuple with the AllowedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValue

`func (o *JsonFieldConstraintsResponse) SetAllowedValue(v []string)`

SetAllowedValue sets AllowedValue field to given value.

### HasAllowedValue

`func (o *JsonFieldConstraintsResponse) HasAllowedValue() bool`

HasAllowedValue returns a boolean if a field has been set.

### GetAllowedValueRegularExpression

`func (o *JsonFieldConstraintsResponse) GetAllowedValueRegularExpression() []string`

GetAllowedValueRegularExpression returns the AllowedValueRegularExpression field if non-nil, zero value otherwise.

### GetAllowedValueRegularExpressionOk

`func (o *JsonFieldConstraintsResponse) GetAllowedValueRegularExpressionOk() (*[]string, bool)`

GetAllowedValueRegularExpressionOk returns a tuple with the AllowedValueRegularExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValueRegularExpression

`func (o *JsonFieldConstraintsResponse) SetAllowedValueRegularExpression(v []string)`

SetAllowedValueRegularExpression sets AllowedValueRegularExpression field to given value.

### HasAllowedValueRegularExpression

`func (o *JsonFieldConstraintsResponse) HasAllowedValueRegularExpression() bool`

HasAllowedValueRegularExpression returns a boolean if a field has been set.

### GetMinimumNumericValue

`func (o *JsonFieldConstraintsResponse) GetMinimumNumericValue() string`

GetMinimumNumericValue returns the MinimumNumericValue field if non-nil, zero value otherwise.

### GetMinimumNumericValueOk

`func (o *JsonFieldConstraintsResponse) GetMinimumNumericValueOk() (*string, bool)`

GetMinimumNumericValueOk returns a tuple with the MinimumNumericValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNumericValue

`func (o *JsonFieldConstraintsResponse) SetMinimumNumericValue(v string)`

SetMinimumNumericValue sets MinimumNumericValue field to given value.

### HasMinimumNumericValue

`func (o *JsonFieldConstraintsResponse) HasMinimumNumericValue() bool`

HasMinimumNumericValue returns a boolean if a field has been set.

### GetMaximumNumericValue

`func (o *JsonFieldConstraintsResponse) GetMaximumNumericValue() string`

GetMaximumNumericValue returns the MaximumNumericValue field if non-nil, zero value otherwise.

### GetMaximumNumericValueOk

`func (o *JsonFieldConstraintsResponse) GetMaximumNumericValueOk() (*string, bool)`

GetMaximumNumericValueOk returns a tuple with the MaximumNumericValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNumericValue

`func (o *JsonFieldConstraintsResponse) SetMaximumNumericValue(v string)`

SetMaximumNumericValue sets MaximumNumericValue field to given value.

### HasMaximumNumericValue

`func (o *JsonFieldConstraintsResponse) HasMaximumNumericValue() bool`

HasMaximumNumericValue returns a boolean if a field has been set.

### GetMinimumValueLength

`func (o *JsonFieldConstraintsResponse) GetMinimumValueLength() int64`

GetMinimumValueLength returns the MinimumValueLength field if non-nil, zero value otherwise.

### GetMinimumValueLengthOk

`func (o *JsonFieldConstraintsResponse) GetMinimumValueLengthOk() (*int64, bool)`

GetMinimumValueLengthOk returns a tuple with the MinimumValueLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumValueLength

`func (o *JsonFieldConstraintsResponse) SetMinimumValueLength(v int64)`

SetMinimumValueLength sets MinimumValueLength field to given value.

### HasMinimumValueLength

`func (o *JsonFieldConstraintsResponse) HasMinimumValueLength() bool`

HasMinimumValueLength returns a boolean if a field has been set.

### GetMaximumValueLength

`func (o *JsonFieldConstraintsResponse) GetMaximumValueLength() int64`

GetMaximumValueLength returns the MaximumValueLength field if non-nil, zero value otherwise.

### GetMaximumValueLengthOk

`func (o *JsonFieldConstraintsResponse) GetMaximumValueLengthOk() (*int64, bool)`

GetMaximumValueLengthOk returns a tuple with the MaximumValueLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValueLength

`func (o *JsonFieldConstraintsResponse) SetMaximumValueLength(v int64)`

SetMaximumValueLength sets MaximumValueLength field to given value.

### HasMaximumValueLength

`func (o *JsonFieldConstraintsResponse) HasMaximumValueLength() bool`

HasMaximumValueLength returns a boolean if a field has been set.

### GetMinimumValueCount

`func (o *JsonFieldConstraintsResponse) GetMinimumValueCount() int64`

GetMinimumValueCount returns the MinimumValueCount field if non-nil, zero value otherwise.

### GetMinimumValueCountOk

`func (o *JsonFieldConstraintsResponse) GetMinimumValueCountOk() (*int64, bool)`

GetMinimumValueCountOk returns a tuple with the MinimumValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumValueCount

`func (o *JsonFieldConstraintsResponse) SetMinimumValueCount(v int64)`

SetMinimumValueCount sets MinimumValueCount field to given value.

### HasMinimumValueCount

`func (o *JsonFieldConstraintsResponse) HasMinimumValueCount() bool`

HasMinimumValueCount returns a boolean if a field has been set.

### GetMaximumValueCount

`func (o *JsonFieldConstraintsResponse) GetMaximumValueCount() int64`

GetMaximumValueCount returns the MaximumValueCount field if non-nil, zero value otherwise.

### GetMaximumValueCountOk

`func (o *JsonFieldConstraintsResponse) GetMaximumValueCountOk() (*int64, bool)`

GetMaximumValueCountOk returns a tuple with the MaximumValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValueCount

`func (o *JsonFieldConstraintsResponse) SetMaximumValueCount(v int64)`

SetMaximumValueCount sets MaximumValueCount field to given value.

### HasMaximumValueCount

`func (o *JsonFieldConstraintsResponse) HasMaximumValueCount() bool`

HasMaximumValueCount returns a boolean if a field has been set.

### GetMeta

`func (o *JsonFieldConstraintsResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *JsonFieldConstraintsResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *JsonFieldConstraintsResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *JsonFieldConstraintsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *JsonFieldConstraintsResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *JsonFieldConstraintsResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *JsonFieldConstraintsResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *JsonFieldConstraintsResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *JsonFieldConstraintsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JsonFieldConstraintsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JsonFieldConstraintsResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


