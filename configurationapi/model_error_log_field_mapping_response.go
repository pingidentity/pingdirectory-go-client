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

// checks if the ErrorLogFieldMappingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorLogFieldMappingResponse{}

// ErrorLogFieldMappingResponse struct for ErrorLogFieldMappingResponse
type ErrorLogFieldMappingResponse struct {
	Schemas []EnumerrorLogFieldMappingSchemaUrn `json:"schemas"`
	// The time that the log message was generated.
	LogFieldTimestamp *string `json:"logFieldTimestamp,omitempty"`
	// The name for this Directory Server product, which may be used to identify which product was used to log the message if multiple products log to the same database table.
	LogFieldProductName *string `json:"logFieldProductName,omitempty"`
	// A name that uniquely identifies this Directory Server instance, which may be used to identify which instance was used to log the message if multiple server instances log to the same database table.
	LogFieldInstanceName *string `json:"logFieldInstanceName,omitempty"`
	// The startup ID for the Directory Server. A different value will be generated each time the server is started.
	LogFieldStartupid *string `json:"logFieldStartupid,omitempty"`
	// The category for the log message.
	LogFieldCategory *string `json:"logFieldCategory,omitempty"`
	// The severity for the log message.
	LogFieldSeverity *string `json:"logFieldSeverity,omitempty"`
	// The numeric value which uniquely identifies the type of message.
	LogFieldMessageID *string `json:"logFieldMessageID,omitempty"`
	// The text of the log message.
	LogFieldMessage *string `json:"logFieldMessage,omitempty"`
	// A description for this Log Field Mapping
	Description                                   *string                                            `json:"description,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Log Field Mapping
	Id string `json:"id"`
}

// NewErrorLogFieldMappingResponse instantiates a new ErrorLogFieldMappingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorLogFieldMappingResponse(schemas []EnumerrorLogFieldMappingSchemaUrn, id string) *ErrorLogFieldMappingResponse {
	this := ErrorLogFieldMappingResponse{}
	this.Schemas = schemas
	this.Id = id
	return &this
}

// NewErrorLogFieldMappingResponseWithDefaults instantiates a new ErrorLogFieldMappingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorLogFieldMappingResponseWithDefaults() *ErrorLogFieldMappingResponse {
	this := ErrorLogFieldMappingResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *ErrorLogFieldMappingResponse) GetSchemas() []EnumerrorLogFieldMappingSchemaUrn {
	if o == nil {
		var ret []EnumerrorLogFieldMappingSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ErrorLogFieldMappingResponse) GetSchemasOk() ([]EnumerrorLogFieldMappingSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ErrorLogFieldMappingResponse) SetSchemas(v []EnumerrorLogFieldMappingSchemaUrn) {
	o.Schemas = v
}

// GetLogFieldTimestamp returns the LogFieldTimestamp field value if set, zero value otherwise.
func (o *ErrorLogFieldMappingResponse) GetLogFieldTimestamp() string {
	if o == nil || IsNil(o.LogFieldTimestamp) {
		var ret string
		return ret
	}
	return *o.LogFieldTimestamp
}

// GetLogFieldTimestampOk returns a tuple with the LogFieldTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogFieldMappingResponse) GetLogFieldTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.LogFieldTimestamp) {
		return nil, false
	}
	return o.LogFieldTimestamp, true
}

// HasLogFieldTimestamp returns a boolean if a field has been set.
func (o *ErrorLogFieldMappingResponse) HasLogFieldTimestamp() bool {
	if o != nil && !IsNil(o.LogFieldTimestamp) {
		return true
	}

	return false
}

// SetLogFieldTimestamp gets a reference to the given string and assigns it to the LogFieldTimestamp field.
func (o *ErrorLogFieldMappingResponse) SetLogFieldTimestamp(v string) {
	o.LogFieldTimestamp = &v
}

// GetLogFieldProductName returns the LogFieldProductName field value if set, zero value otherwise.
func (o *ErrorLogFieldMappingResponse) GetLogFieldProductName() string {
	if o == nil || IsNil(o.LogFieldProductName) {
		var ret string
		return ret
	}
	return *o.LogFieldProductName
}

// GetLogFieldProductNameOk returns a tuple with the LogFieldProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogFieldMappingResponse) GetLogFieldProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.LogFieldProductName) {
		return nil, false
	}
	return o.LogFieldProductName, true
}

// HasLogFieldProductName returns a boolean if a field has been set.
func (o *ErrorLogFieldMappingResponse) HasLogFieldProductName() bool {
	if o != nil && !IsNil(o.LogFieldProductName) {
		return true
	}

	return false
}

// SetLogFieldProductName gets a reference to the given string and assigns it to the LogFieldProductName field.
func (o *ErrorLogFieldMappingResponse) SetLogFieldProductName(v string) {
	o.LogFieldProductName = &v
}

// GetLogFieldInstanceName returns the LogFieldInstanceName field value if set, zero value otherwise.
func (o *ErrorLogFieldMappingResponse) GetLogFieldInstanceName() string {
	if o == nil || IsNil(o.LogFieldInstanceName) {
		var ret string
		return ret
	}
	return *o.LogFieldInstanceName
}

// GetLogFieldInstanceNameOk returns a tuple with the LogFieldInstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogFieldMappingResponse) GetLogFieldInstanceNameOk() (*string, bool) {
	if o == nil || IsNil(o.LogFieldInstanceName) {
		return nil, false
	}
	return o.LogFieldInstanceName, true
}

// HasLogFieldInstanceName returns a boolean if a field has been set.
func (o *ErrorLogFieldMappingResponse) HasLogFieldInstanceName() bool {
	if o != nil && !IsNil(o.LogFieldInstanceName) {
		return true
	}

	return false
}

// SetLogFieldInstanceName gets a reference to the given string and assigns it to the LogFieldInstanceName field.
func (o *ErrorLogFieldMappingResponse) SetLogFieldInstanceName(v string) {
	o.LogFieldInstanceName = &v
}

// GetLogFieldStartupid returns the LogFieldStartupid field value if set, zero value otherwise.
func (o *ErrorLogFieldMappingResponse) GetLogFieldStartupid() string {
	if o == nil || IsNil(o.LogFieldStartupid) {
		var ret string
		return ret
	}
	return *o.LogFieldStartupid
}

// GetLogFieldStartupidOk returns a tuple with the LogFieldStartupid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogFieldMappingResponse) GetLogFieldStartupidOk() (*string, bool) {
	if o == nil || IsNil(o.LogFieldStartupid) {
		return nil, false
	}
	return o.LogFieldStartupid, true
}

// HasLogFieldStartupid returns a boolean if a field has been set.
func (o *ErrorLogFieldMappingResponse) HasLogFieldStartupid() bool {
	if o != nil && !IsNil(o.LogFieldStartupid) {
		return true
	}

	return false
}

// SetLogFieldStartupid gets a reference to the given string and assigns it to the LogFieldStartupid field.
func (o *ErrorLogFieldMappingResponse) SetLogFieldStartupid(v string) {
	o.LogFieldStartupid = &v
}

// GetLogFieldCategory returns the LogFieldCategory field value if set, zero value otherwise.
func (o *ErrorLogFieldMappingResponse) GetLogFieldCategory() string {
	if o == nil || IsNil(o.LogFieldCategory) {
		var ret string
		return ret
	}
	return *o.LogFieldCategory
}

// GetLogFieldCategoryOk returns a tuple with the LogFieldCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogFieldMappingResponse) GetLogFieldCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.LogFieldCategory) {
		return nil, false
	}
	return o.LogFieldCategory, true
}

// HasLogFieldCategory returns a boolean if a field has been set.
func (o *ErrorLogFieldMappingResponse) HasLogFieldCategory() bool {
	if o != nil && !IsNil(o.LogFieldCategory) {
		return true
	}

	return false
}

// SetLogFieldCategory gets a reference to the given string and assigns it to the LogFieldCategory field.
func (o *ErrorLogFieldMappingResponse) SetLogFieldCategory(v string) {
	o.LogFieldCategory = &v
}

// GetLogFieldSeverity returns the LogFieldSeverity field value if set, zero value otherwise.
func (o *ErrorLogFieldMappingResponse) GetLogFieldSeverity() string {
	if o == nil || IsNil(o.LogFieldSeverity) {
		var ret string
		return ret
	}
	return *o.LogFieldSeverity
}

// GetLogFieldSeverityOk returns a tuple with the LogFieldSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogFieldMappingResponse) GetLogFieldSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.LogFieldSeverity) {
		return nil, false
	}
	return o.LogFieldSeverity, true
}

// HasLogFieldSeverity returns a boolean if a field has been set.
func (o *ErrorLogFieldMappingResponse) HasLogFieldSeverity() bool {
	if o != nil && !IsNil(o.LogFieldSeverity) {
		return true
	}

	return false
}

// SetLogFieldSeverity gets a reference to the given string and assigns it to the LogFieldSeverity field.
func (o *ErrorLogFieldMappingResponse) SetLogFieldSeverity(v string) {
	o.LogFieldSeverity = &v
}

// GetLogFieldMessageID returns the LogFieldMessageID field value if set, zero value otherwise.
func (o *ErrorLogFieldMappingResponse) GetLogFieldMessageID() string {
	if o == nil || IsNil(o.LogFieldMessageID) {
		var ret string
		return ret
	}
	return *o.LogFieldMessageID
}

// GetLogFieldMessageIDOk returns a tuple with the LogFieldMessageID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogFieldMappingResponse) GetLogFieldMessageIDOk() (*string, bool) {
	if o == nil || IsNil(o.LogFieldMessageID) {
		return nil, false
	}
	return o.LogFieldMessageID, true
}

// HasLogFieldMessageID returns a boolean if a field has been set.
func (o *ErrorLogFieldMappingResponse) HasLogFieldMessageID() bool {
	if o != nil && !IsNil(o.LogFieldMessageID) {
		return true
	}

	return false
}

// SetLogFieldMessageID gets a reference to the given string and assigns it to the LogFieldMessageID field.
func (o *ErrorLogFieldMappingResponse) SetLogFieldMessageID(v string) {
	o.LogFieldMessageID = &v
}

// GetLogFieldMessage returns the LogFieldMessage field value if set, zero value otherwise.
func (o *ErrorLogFieldMappingResponse) GetLogFieldMessage() string {
	if o == nil || IsNil(o.LogFieldMessage) {
		var ret string
		return ret
	}
	return *o.LogFieldMessage
}

// GetLogFieldMessageOk returns a tuple with the LogFieldMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogFieldMappingResponse) GetLogFieldMessageOk() (*string, bool) {
	if o == nil || IsNil(o.LogFieldMessage) {
		return nil, false
	}
	return o.LogFieldMessage, true
}

// HasLogFieldMessage returns a boolean if a field has been set.
func (o *ErrorLogFieldMappingResponse) HasLogFieldMessage() bool {
	if o != nil && !IsNil(o.LogFieldMessage) {
		return true
	}

	return false
}

// SetLogFieldMessage gets a reference to the given string and assigns it to the LogFieldMessage field.
func (o *ErrorLogFieldMappingResponse) SetLogFieldMessage(v string) {
	o.LogFieldMessage = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ErrorLogFieldMappingResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogFieldMappingResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ErrorLogFieldMappingResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ErrorLogFieldMappingResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ErrorLogFieldMappingResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogFieldMappingResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ErrorLogFieldMappingResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ErrorLogFieldMappingResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ErrorLogFieldMappingResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogFieldMappingResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ErrorLogFieldMappingResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ErrorLogFieldMappingResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *ErrorLogFieldMappingResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ErrorLogFieldMappingResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ErrorLogFieldMappingResponse) SetId(v string) {
	o.Id = v
}

func (o ErrorLogFieldMappingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorLogFieldMappingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.LogFieldTimestamp) {
		toSerialize["logFieldTimestamp"] = o.LogFieldTimestamp
	}
	if !IsNil(o.LogFieldProductName) {
		toSerialize["logFieldProductName"] = o.LogFieldProductName
	}
	if !IsNil(o.LogFieldInstanceName) {
		toSerialize["logFieldInstanceName"] = o.LogFieldInstanceName
	}
	if !IsNil(o.LogFieldStartupid) {
		toSerialize["logFieldStartupid"] = o.LogFieldStartupid
	}
	if !IsNil(o.LogFieldCategory) {
		toSerialize["logFieldCategory"] = o.LogFieldCategory
	}
	if !IsNil(o.LogFieldSeverity) {
		toSerialize["logFieldSeverity"] = o.LogFieldSeverity
	}
	if !IsNil(o.LogFieldMessageID) {
		toSerialize["logFieldMessageID"] = o.LogFieldMessageID
	}
	if !IsNil(o.LogFieldMessage) {
		toSerialize["logFieldMessage"] = o.LogFieldMessage
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableErrorLogFieldMappingResponse struct {
	value *ErrorLogFieldMappingResponse
	isSet bool
}

func (v NullableErrorLogFieldMappingResponse) Get() *ErrorLogFieldMappingResponse {
	return v.value
}

func (v *NullableErrorLogFieldMappingResponse) Set(val *ErrorLogFieldMappingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorLogFieldMappingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorLogFieldMappingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorLogFieldMappingResponse(val *ErrorLogFieldMappingResponse) *NullableErrorLogFieldMappingResponse {
	return &NullableErrorLogFieldMappingResponse{value: val, isSet: true}
}

func (v NullableErrorLogFieldMappingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorLogFieldMappingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
