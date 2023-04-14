/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the JsonFieldConstraintsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JsonFieldConstraintsResponse{}

// JsonFieldConstraintsResponse struct for JsonFieldConstraintsResponse
type JsonFieldConstraintsResponse struct {
	// Name of the JSON Attribute Constraints
	Id      string                              `json:"id"`
	Schemas []EnumjsonFieldConstraintsSchemaUrn `json:"schemas,omitempty"`
	// A description for this JSON Field Constraints
	Description *string `json:"description,omitempty"`
	// The full name of the JSON field to which these constraints apply.
	JsonField string                                `json:"jsonField"`
	ValueType EnumjsonFieldConstraintsValueTypeProp `json:"valueType"`
	// Indicates whether the target field must be present in JSON objects stored as values of the associated attribute type.
	IsRequired *bool                                `json:"isRequired,omitempty"`
	IsArray    *EnumjsonFieldConstraintsIsArrayProp `json:"isArray,omitempty"`
	// Indicates whether the target field may have a value that is the JSON null value as an alternative to a value (or array of values) of the specified value-type.
	AllowNullValue *bool `json:"allowNullValue,omitempty"`
	// Indicates whether the target field may have a value that is an empty JSON object (i.e., a JSON object with zero fields). This may only be set to true if value-type property is set to object.
	AllowEmptyObject *bool `json:"allowEmptyObject,omitempty"`
	// Indicates whether backends that support JSON indexing should maintain an index for values of the target field.
	IndexValues *bool `json:"indexValues,omitempty"`
	// The maximum number of entries that may contain a particular value for the target field before the server will stop maintaining the index for that value.
	IndexEntryLimit *int64 `json:"indexEntryLimit,omitempty"`
	// Indicates whether backends that support database priming should load the contents of the associated JSON index into memory whenever the backend is opened.
	PrimeIndex *bool                                  `json:"primeIndex,omitempty"`
	CacheMode  *EnumjsonFieldConstraintsCacheModeProp `json:"cacheMode,omitempty"`
	// Indicates whether the backend should attempt to assign a compact token for each distinct value for the target field in an attempt to reduce the encoded size of the field in JSON objects. These tokens would be assigned prior to using any from the token set used for automatic compaction of some JSON string values.
	TokenizeValues *bool `json:"tokenizeValues,omitempty"`
	// Specifies an explicit set of string values that will be the only values permitted for the target field. If a set of allowed values is defined, then the server will reject any attempt to store a JSON object with a value for the target field that is not included in that set.
	AllowedValue []string `json:"allowedValue,omitempty"`
	// Specifies an explicit set of regular expressions that may be used to restrict the set of values that may be used for the target field. If a set of allowed value regular expressions is defined, then the server will reject any attempt to store a JSON object with a value for the target field that does not match at least one of those regular expressions.
	AllowedValueRegularExpression []string `json:"allowedValueRegularExpression,omitempty"`
	// Specifies the smallest numeric value that may be used as the value for the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is less than that minimum numeric value.
	MinimumNumericValue *string `json:"minimumNumericValue,omitempty"`
	// Specifies the largest numeric value that may be used as the value for the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is greater than that maximum numeric value.
	MaximumNumericValue *string `json:"maximumNumericValue,omitempty"`
	// Specifies the smallest number of characters that may be present in string values of the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is shorter than that minimum value length.
	MinimumValueLength *int64 `json:"minimumValueLength,omitempty"`
	// Specifies the largest number of characters that may be present in string values of the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is longer than that maximum value length.
	MaximumValueLength *int64 `json:"maximumValueLength,omitempty"`
	// Specifies the smallest number of elements that may be present in an array of values for the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is an array with fewer than this number of elements.
	MinimumValueCount *int64 `json:"minimumValueCount,omitempty"`
	// Specifies the largest number of elements that may be present in an array of values for the target field. If configured, then the server will reject any attempt to store a JSON object with a value for the target field that is an array with more than this number of elements.
	MaximumValueCount                             *int64                                             `json:"maximumValueCount,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewJsonFieldConstraintsResponse instantiates a new JsonFieldConstraintsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonFieldConstraintsResponse(id string, jsonField string, valueType EnumjsonFieldConstraintsValueTypeProp) *JsonFieldConstraintsResponse {
	this := JsonFieldConstraintsResponse{}
	this.Id = id
	this.JsonField = jsonField
	this.ValueType = valueType
	return &this
}

// NewJsonFieldConstraintsResponseWithDefaults instantiates a new JsonFieldConstraintsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonFieldConstraintsResponseWithDefaults() *JsonFieldConstraintsResponse {
	this := JsonFieldConstraintsResponse{}
	return &this
}

// GetId returns the Id field value
func (o *JsonFieldConstraintsResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JsonFieldConstraintsResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetSchemas() []EnumjsonFieldConstraintsSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumjsonFieldConstraintsSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetSchemasOk() ([]EnumjsonFieldConstraintsSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumjsonFieldConstraintsSchemaUrn and assigns it to the Schemas field.
func (o *JsonFieldConstraintsResponse) SetSchemas(v []EnumjsonFieldConstraintsSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JsonFieldConstraintsResponse) SetDescription(v string) {
	o.Description = &v
}

// GetJsonField returns the JsonField field value
func (o *JsonFieldConstraintsResponse) GetJsonField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsonField
}

// GetJsonFieldOk returns a tuple with the JsonField field value
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetJsonFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonField, true
}

// SetJsonField sets field value
func (o *JsonFieldConstraintsResponse) SetJsonField(v string) {
	o.JsonField = v
}

// GetValueType returns the ValueType field value
func (o *JsonFieldConstraintsResponse) GetValueType() EnumjsonFieldConstraintsValueTypeProp {
	if o == nil {
		var ret EnumjsonFieldConstraintsValueTypeProp
		return ret
	}

	return o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetValueTypeOk() (*EnumjsonFieldConstraintsValueTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueType, true
}

// SetValueType sets field value
func (o *JsonFieldConstraintsResponse) SetValueType(v EnumjsonFieldConstraintsValueTypeProp) {
	o.ValueType = v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetIsRequired() bool {
	if o == nil || IsNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetIsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRequired) {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasIsRequired() bool {
	if o != nil && !IsNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *JsonFieldConstraintsResponse) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetIsArray returns the IsArray field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetIsArray() EnumjsonFieldConstraintsIsArrayProp {
	if o == nil || IsNil(o.IsArray) {
		var ret EnumjsonFieldConstraintsIsArrayProp
		return ret
	}
	return *o.IsArray
}

// GetIsArrayOk returns a tuple with the IsArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetIsArrayOk() (*EnumjsonFieldConstraintsIsArrayProp, bool) {
	if o == nil || IsNil(o.IsArray) {
		return nil, false
	}
	return o.IsArray, true
}

// HasIsArray returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasIsArray() bool {
	if o != nil && !IsNil(o.IsArray) {
		return true
	}

	return false
}

// SetIsArray gets a reference to the given EnumjsonFieldConstraintsIsArrayProp and assigns it to the IsArray field.
func (o *JsonFieldConstraintsResponse) SetIsArray(v EnumjsonFieldConstraintsIsArrayProp) {
	o.IsArray = &v
}

// GetAllowNullValue returns the AllowNullValue field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetAllowNullValue() bool {
	if o == nil || IsNil(o.AllowNullValue) {
		var ret bool
		return ret
	}
	return *o.AllowNullValue
}

// GetAllowNullValueOk returns a tuple with the AllowNullValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetAllowNullValueOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowNullValue) {
		return nil, false
	}
	return o.AllowNullValue, true
}

// HasAllowNullValue returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasAllowNullValue() bool {
	if o != nil && !IsNil(o.AllowNullValue) {
		return true
	}

	return false
}

// SetAllowNullValue gets a reference to the given bool and assigns it to the AllowNullValue field.
func (o *JsonFieldConstraintsResponse) SetAllowNullValue(v bool) {
	o.AllowNullValue = &v
}

// GetAllowEmptyObject returns the AllowEmptyObject field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetAllowEmptyObject() bool {
	if o == nil || IsNil(o.AllowEmptyObject) {
		var ret bool
		return ret
	}
	return *o.AllowEmptyObject
}

// GetAllowEmptyObjectOk returns a tuple with the AllowEmptyObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetAllowEmptyObjectOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowEmptyObject) {
		return nil, false
	}
	return o.AllowEmptyObject, true
}

// HasAllowEmptyObject returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasAllowEmptyObject() bool {
	if o != nil && !IsNil(o.AllowEmptyObject) {
		return true
	}

	return false
}

// SetAllowEmptyObject gets a reference to the given bool and assigns it to the AllowEmptyObject field.
func (o *JsonFieldConstraintsResponse) SetAllowEmptyObject(v bool) {
	o.AllowEmptyObject = &v
}

// GetIndexValues returns the IndexValues field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetIndexValues() bool {
	if o == nil || IsNil(o.IndexValues) {
		var ret bool
		return ret
	}
	return *o.IndexValues
}

// GetIndexValuesOk returns a tuple with the IndexValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetIndexValuesOk() (*bool, bool) {
	if o == nil || IsNil(o.IndexValues) {
		return nil, false
	}
	return o.IndexValues, true
}

// HasIndexValues returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasIndexValues() bool {
	if o != nil && !IsNil(o.IndexValues) {
		return true
	}

	return false
}

// SetIndexValues gets a reference to the given bool and assigns it to the IndexValues field.
func (o *JsonFieldConstraintsResponse) SetIndexValues(v bool) {
	o.IndexValues = &v
}

// GetIndexEntryLimit returns the IndexEntryLimit field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetIndexEntryLimit() int64 {
	if o == nil || IsNil(o.IndexEntryLimit) {
		var ret int64
		return ret
	}
	return *o.IndexEntryLimit
}

// GetIndexEntryLimitOk returns a tuple with the IndexEntryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetIndexEntryLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.IndexEntryLimit) {
		return nil, false
	}
	return o.IndexEntryLimit, true
}

// HasIndexEntryLimit returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasIndexEntryLimit() bool {
	if o != nil && !IsNil(o.IndexEntryLimit) {
		return true
	}

	return false
}

// SetIndexEntryLimit gets a reference to the given int64 and assigns it to the IndexEntryLimit field.
func (o *JsonFieldConstraintsResponse) SetIndexEntryLimit(v int64) {
	o.IndexEntryLimit = &v
}

// GetPrimeIndex returns the PrimeIndex field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetPrimeIndex() bool {
	if o == nil || IsNil(o.PrimeIndex) {
		var ret bool
		return ret
	}
	return *o.PrimeIndex
}

// GetPrimeIndexOk returns a tuple with the PrimeIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetPrimeIndexOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimeIndex) {
		return nil, false
	}
	return o.PrimeIndex, true
}

// HasPrimeIndex returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasPrimeIndex() bool {
	if o != nil && !IsNil(o.PrimeIndex) {
		return true
	}

	return false
}

// SetPrimeIndex gets a reference to the given bool and assigns it to the PrimeIndex field.
func (o *JsonFieldConstraintsResponse) SetPrimeIndex(v bool) {
	o.PrimeIndex = &v
}

// GetCacheMode returns the CacheMode field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetCacheMode() EnumjsonFieldConstraintsCacheModeProp {
	if o == nil || IsNil(o.CacheMode) {
		var ret EnumjsonFieldConstraintsCacheModeProp
		return ret
	}
	return *o.CacheMode
}

// GetCacheModeOk returns a tuple with the CacheMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetCacheModeOk() (*EnumjsonFieldConstraintsCacheModeProp, bool) {
	if o == nil || IsNil(o.CacheMode) {
		return nil, false
	}
	return o.CacheMode, true
}

// HasCacheMode returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasCacheMode() bool {
	if o != nil && !IsNil(o.CacheMode) {
		return true
	}

	return false
}

// SetCacheMode gets a reference to the given EnumjsonFieldConstraintsCacheModeProp and assigns it to the CacheMode field.
func (o *JsonFieldConstraintsResponse) SetCacheMode(v EnumjsonFieldConstraintsCacheModeProp) {
	o.CacheMode = &v
}

// GetTokenizeValues returns the TokenizeValues field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetTokenizeValues() bool {
	if o == nil || IsNil(o.TokenizeValues) {
		var ret bool
		return ret
	}
	return *o.TokenizeValues
}

// GetTokenizeValuesOk returns a tuple with the TokenizeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetTokenizeValuesOk() (*bool, bool) {
	if o == nil || IsNil(o.TokenizeValues) {
		return nil, false
	}
	return o.TokenizeValues, true
}

// HasTokenizeValues returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasTokenizeValues() bool {
	if o != nil && !IsNil(o.TokenizeValues) {
		return true
	}

	return false
}

// SetTokenizeValues gets a reference to the given bool and assigns it to the TokenizeValues field.
func (o *JsonFieldConstraintsResponse) SetTokenizeValues(v bool) {
	o.TokenizeValues = &v
}

// GetAllowedValue returns the AllowedValue field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetAllowedValue() []string {
	if o == nil || IsNil(o.AllowedValue) {
		var ret []string
		return ret
	}
	return o.AllowedValue
}

// GetAllowedValueOk returns a tuple with the AllowedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetAllowedValueOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedValue) {
		return nil, false
	}
	return o.AllowedValue, true
}

// HasAllowedValue returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasAllowedValue() bool {
	if o != nil && !IsNil(o.AllowedValue) {
		return true
	}

	return false
}

// SetAllowedValue gets a reference to the given []string and assigns it to the AllowedValue field.
func (o *JsonFieldConstraintsResponse) SetAllowedValue(v []string) {
	o.AllowedValue = v
}

// GetAllowedValueRegularExpression returns the AllowedValueRegularExpression field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetAllowedValueRegularExpression() []string {
	if o == nil || IsNil(o.AllowedValueRegularExpression) {
		var ret []string
		return ret
	}
	return o.AllowedValueRegularExpression
}

// GetAllowedValueRegularExpressionOk returns a tuple with the AllowedValueRegularExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetAllowedValueRegularExpressionOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedValueRegularExpression) {
		return nil, false
	}
	return o.AllowedValueRegularExpression, true
}

// HasAllowedValueRegularExpression returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasAllowedValueRegularExpression() bool {
	if o != nil && !IsNil(o.AllowedValueRegularExpression) {
		return true
	}

	return false
}

// SetAllowedValueRegularExpression gets a reference to the given []string and assigns it to the AllowedValueRegularExpression field.
func (o *JsonFieldConstraintsResponse) SetAllowedValueRegularExpression(v []string) {
	o.AllowedValueRegularExpression = v
}

// GetMinimumNumericValue returns the MinimumNumericValue field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetMinimumNumericValue() string {
	if o == nil || IsNil(o.MinimumNumericValue) {
		var ret string
		return ret
	}
	return *o.MinimumNumericValue
}

// GetMinimumNumericValueOk returns a tuple with the MinimumNumericValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetMinimumNumericValueOk() (*string, bool) {
	if o == nil || IsNil(o.MinimumNumericValue) {
		return nil, false
	}
	return o.MinimumNumericValue, true
}

// HasMinimumNumericValue returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasMinimumNumericValue() bool {
	if o != nil && !IsNil(o.MinimumNumericValue) {
		return true
	}

	return false
}

// SetMinimumNumericValue gets a reference to the given string and assigns it to the MinimumNumericValue field.
func (o *JsonFieldConstraintsResponse) SetMinimumNumericValue(v string) {
	o.MinimumNumericValue = &v
}

// GetMaximumNumericValue returns the MaximumNumericValue field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetMaximumNumericValue() string {
	if o == nil || IsNil(o.MaximumNumericValue) {
		var ret string
		return ret
	}
	return *o.MaximumNumericValue
}

// GetMaximumNumericValueOk returns a tuple with the MaximumNumericValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetMaximumNumericValueOk() (*string, bool) {
	if o == nil || IsNil(o.MaximumNumericValue) {
		return nil, false
	}
	return o.MaximumNumericValue, true
}

// HasMaximumNumericValue returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasMaximumNumericValue() bool {
	if o != nil && !IsNil(o.MaximumNumericValue) {
		return true
	}

	return false
}

// SetMaximumNumericValue gets a reference to the given string and assigns it to the MaximumNumericValue field.
func (o *JsonFieldConstraintsResponse) SetMaximumNumericValue(v string) {
	o.MaximumNumericValue = &v
}

// GetMinimumValueLength returns the MinimumValueLength field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetMinimumValueLength() int64 {
	if o == nil || IsNil(o.MinimumValueLength) {
		var ret int64
		return ret
	}
	return *o.MinimumValueLength
}

// GetMinimumValueLengthOk returns a tuple with the MinimumValueLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetMinimumValueLengthOk() (*int64, bool) {
	if o == nil || IsNil(o.MinimumValueLength) {
		return nil, false
	}
	return o.MinimumValueLength, true
}

// HasMinimumValueLength returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasMinimumValueLength() bool {
	if o != nil && !IsNil(o.MinimumValueLength) {
		return true
	}

	return false
}

// SetMinimumValueLength gets a reference to the given int64 and assigns it to the MinimumValueLength field.
func (o *JsonFieldConstraintsResponse) SetMinimumValueLength(v int64) {
	o.MinimumValueLength = &v
}

// GetMaximumValueLength returns the MaximumValueLength field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetMaximumValueLength() int64 {
	if o == nil || IsNil(o.MaximumValueLength) {
		var ret int64
		return ret
	}
	return *o.MaximumValueLength
}

// GetMaximumValueLengthOk returns a tuple with the MaximumValueLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetMaximumValueLengthOk() (*int64, bool) {
	if o == nil || IsNil(o.MaximumValueLength) {
		return nil, false
	}
	return o.MaximumValueLength, true
}

// HasMaximumValueLength returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasMaximumValueLength() bool {
	if o != nil && !IsNil(o.MaximumValueLength) {
		return true
	}

	return false
}

// SetMaximumValueLength gets a reference to the given int64 and assigns it to the MaximumValueLength field.
func (o *JsonFieldConstraintsResponse) SetMaximumValueLength(v int64) {
	o.MaximumValueLength = &v
}

// GetMinimumValueCount returns the MinimumValueCount field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetMinimumValueCount() int64 {
	if o == nil || IsNil(o.MinimumValueCount) {
		var ret int64
		return ret
	}
	return *o.MinimumValueCount
}

// GetMinimumValueCountOk returns a tuple with the MinimumValueCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetMinimumValueCountOk() (*int64, bool) {
	if o == nil || IsNil(o.MinimumValueCount) {
		return nil, false
	}
	return o.MinimumValueCount, true
}

// HasMinimumValueCount returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasMinimumValueCount() bool {
	if o != nil && !IsNil(o.MinimumValueCount) {
		return true
	}

	return false
}

// SetMinimumValueCount gets a reference to the given int64 and assigns it to the MinimumValueCount field.
func (o *JsonFieldConstraintsResponse) SetMinimumValueCount(v int64) {
	o.MinimumValueCount = &v
}

// GetMaximumValueCount returns the MaximumValueCount field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetMaximumValueCount() int64 {
	if o == nil || IsNil(o.MaximumValueCount) {
		var ret int64
		return ret
	}
	return *o.MaximumValueCount
}

// GetMaximumValueCountOk returns a tuple with the MaximumValueCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetMaximumValueCountOk() (*int64, bool) {
	if o == nil || IsNil(o.MaximumValueCount) {
		return nil, false
	}
	return o.MaximumValueCount, true
}

// HasMaximumValueCount returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasMaximumValueCount() bool {
	if o != nil && !IsNil(o.MaximumValueCount) {
		return true
	}

	return false
}

// SetMaximumValueCount gets a reference to the given int64 and assigns it to the MaximumValueCount field.
func (o *JsonFieldConstraintsResponse) SetMaximumValueCount(v int64) {
	o.MaximumValueCount = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *JsonFieldConstraintsResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *JsonFieldConstraintsResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonFieldConstraintsResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *JsonFieldConstraintsResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *JsonFieldConstraintsResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o JsonFieldConstraintsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JsonFieldConstraintsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["jsonField"] = o.JsonField
	toSerialize["valueType"] = o.ValueType
	if !IsNil(o.IsRequired) {
		toSerialize["isRequired"] = o.IsRequired
	}
	if !IsNil(o.IsArray) {
		toSerialize["isArray"] = o.IsArray
	}
	if !IsNil(o.AllowNullValue) {
		toSerialize["allowNullValue"] = o.AllowNullValue
	}
	if !IsNil(o.AllowEmptyObject) {
		toSerialize["allowEmptyObject"] = o.AllowEmptyObject
	}
	if !IsNil(o.IndexValues) {
		toSerialize["indexValues"] = o.IndexValues
	}
	if !IsNil(o.IndexEntryLimit) {
		toSerialize["indexEntryLimit"] = o.IndexEntryLimit
	}
	if !IsNil(o.PrimeIndex) {
		toSerialize["primeIndex"] = o.PrimeIndex
	}
	if !IsNil(o.CacheMode) {
		toSerialize["cacheMode"] = o.CacheMode
	}
	if !IsNil(o.TokenizeValues) {
		toSerialize["tokenizeValues"] = o.TokenizeValues
	}
	if !IsNil(o.AllowedValue) {
		toSerialize["allowedValue"] = o.AllowedValue
	}
	if !IsNil(o.AllowedValueRegularExpression) {
		toSerialize["allowedValueRegularExpression"] = o.AllowedValueRegularExpression
	}
	if !IsNil(o.MinimumNumericValue) {
		toSerialize["minimumNumericValue"] = o.MinimumNumericValue
	}
	if !IsNil(o.MaximumNumericValue) {
		toSerialize["maximumNumericValue"] = o.MaximumNumericValue
	}
	if !IsNil(o.MinimumValueLength) {
		toSerialize["minimumValueLength"] = o.MinimumValueLength
	}
	if !IsNil(o.MaximumValueLength) {
		toSerialize["maximumValueLength"] = o.MaximumValueLength
	}
	if !IsNil(o.MinimumValueCount) {
		toSerialize["minimumValueCount"] = o.MinimumValueCount
	}
	if !IsNil(o.MaximumValueCount) {
		toSerialize["maximumValueCount"] = o.MaximumValueCount
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableJsonFieldConstraintsResponse struct {
	value *JsonFieldConstraintsResponse
	isSet bool
}

func (v NullableJsonFieldConstraintsResponse) Get() *JsonFieldConstraintsResponse {
	return v.value
}

func (v *NullableJsonFieldConstraintsResponse) Set(val *JsonFieldConstraintsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonFieldConstraintsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonFieldConstraintsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonFieldConstraintsResponse(val *JsonFieldConstraintsResponse) *NullableJsonFieldConstraintsResponse {
	return &NullableJsonFieldConstraintsResponse{value: val, isSet: true}
}

func (v NullableJsonFieldConstraintsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonFieldConstraintsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
