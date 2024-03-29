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

// checks if the StaticallyDefinedRecurringTaskResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticallyDefinedRecurringTaskResponse{}

// StaticallyDefinedRecurringTaskResponse struct for StaticallyDefinedRecurringTaskResponse
type StaticallyDefinedRecurringTaskResponse struct {
	Schemas []EnumstaticallyDefinedRecurringTaskSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class that provides the logic for the task to be invoked.
	TaskJavaClass string `json:"taskJavaClass"`
	// The names or OIDs of the object classes to include in the tasks that are scheduled from this Statically Defined Recurring Task. All object classes must be defined in the server schema, and the combination of object classes must be valid for a task entry.
	TaskObjectClass []string `json:"taskObjectClass"`
	// The set of attribute values that should be included in the tasks that are scheduled from this Statically Defined Recurring Task. Each value must be in the form {attribute-type}={value}, where {attribute-type} is the name or OID of an attribute type that is defined in the schema and permitted with the configured set of object classes, and {value} is a value to assign to an attribute with that type. A multivalued attribute can be created by providing multiple name-value pairs with the same name and different values.
	TaskAttributeValue []string `json:"taskAttributeValue,omitempty"`
	// A description for this Recurring Task
	Description *string `json:"description,omitempty"`
	// Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running).
	CancelOnTaskDependencyFailure *bool `json:"cancelOnTaskDependencyFailure,omitempty"`
	// The email addresses to which a message should be sent whenever an instance of this Recurring Task starts running. If this option is used, then at least one smtp-server must be configured in the global configuration.
	EmailOnStart []string `json:"emailOnStart,omitempty"`
	// The email addresses to which a message should be sent whenever an instance of this Recurring Task completes successfully. If this option is used, then at least one smtp-server must be configured in the global configuration.
	EmailOnSuccess []string `json:"emailOnSuccess,omitempty"`
	// The email addresses to which a message should be sent if an instance of this Recurring Task fails to complete successfully. If this option is used, then at least one smtp-server must be configured in the global configuration.
	EmailOnFailure []string `json:"emailOnFailure,omitempty"`
	// Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running.
	AlertOnStart *bool `json:"alertOnStart,omitempty"`
	// Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully.
	AlertOnSuccess *bool `json:"alertOnSuccess,omitempty"`
	// Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully.
	AlertOnFailure                                *bool                                              `json:"alertOnFailure,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Recurring Task
	Id string `json:"id"`
}

// NewStaticallyDefinedRecurringTaskResponse instantiates a new StaticallyDefinedRecurringTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticallyDefinedRecurringTaskResponse(schemas []EnumstaticallyDefinedRecurringTaskSchemaUrn, taskJavaClass string, taskObjectClass []string, id string) *StaticallyDefinedRecurringTaskResponse {
	this := StaticallyDefinedRecurringTaskResponse{}
	this.Schemas = schemas
	this.TaskJavaClass = taskJavaClass
	this.TaskObjectClass = taskObjectClass
	this.Id = id
	return &this
}

// NewStaticallyDefinedRecurringTaskResponseWithDefaults instantiates a new StaticallyDefinedRecurringTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticallyDefinedRecurringTaskResponseWithDefaults() *StaticallyDefinedRecurringTaskResponse {
	this := StaticallyDefinedRecurringTaskResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *StaticallyDefinedRecurringTaskResponse) GetSchemas() []EnumstaticallyDefinedRecurringTaskSchemaUrn {
	if o == nil {
		var ret []EnumstaticallyDefinedRecurringTaskSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *StaticallyDefinedRecurringTaskResponse) GetSchemasOk() ([]EnumstaticallyDefinedRecurringTaskSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *StaticallyDefinedRecurringTaskResponse) SetSchemas(v []EnumstaticallyDefinedRecurringTaskSchemaUrn) {
	o.Schemas = v
}

// GetTaskJavaClass returns the TaskJavaClass field value
func (o *StaticallyDefinedRecurringTaskResponse) GetTaskJavaClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskJavaClass
}

// GetTaskJavaClassOk returns a tuple with the TaskJavaClass field value
// and a boolean to check if the value has been set.
func (o *StaticallyDefinedRecurringTaskResponse) GetTaskJavaClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskJavaClass, true
}

// SetTaskJavaClass sets field value
func (o *StaticallyDefinedRecurringTaskResponse) SetTaskJavaClass(v string) {
	o.TaskJavaClass = v
}

// GetTaskObjectClass returns the TaskObjectClass field value
func (o *StaticallyDefinedRecurringTaskResponse) GetTaskObjectClass() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TaskObjectClass
}

// GetTaskObjectClassOk returns a tuple with the TaskObjectClass field value
// and a boolean to check if the value has been set.
func (o *StaticallyDefinedRecurringTaskResponse) GetTaskObjectClassOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskObjectClass, true
}

// SetTaskObjectClass sets field value
func (o *StaticallyDefinedRecurringTaskResponse) SetTaskObjectClass(v []string) {
	o.TaskObjectClass = v
}

// GetTaskAttributeValue returns the TaskAttributeValue field value if set, zero value otherwise.
func (o *StaticallyDefinedRecurringTaskResponse) GetTaskAttributeValue() []string {
	if o == nil || IsNil(o.TaskAttributeValue) {
		var ret []string
		return ret
	}
	return o.TaskAttributeValue
}

// GetTaskAttributeValueOk returns a tuple with the TaskAttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticallyDefinedRecurringTaskResponse) GetTaskAttributeValueOk() ([]string, bool) {
	if o == nil || IsNil(o.TaskAttributeValue) {
		return nil, false
	}
	return o.TaskAttributeValue, true
}

// HasTaskAttributeValue returns a boolean if a field has been set.
func (o *StaticallyDefinedRecurringTaskResponse) HasTaskAttributeValue() bool {
	if o != nil && !IsNil(o.TaskAttributeValue) {
		return true
	}

	return false
}

// SetTaskAttributeValue gets a reference to the given []string and assigns it to the TaskAttributeValue field.
func (o *StaticallyDefinedRecurringTaskResponse) SetTaskAttributeValue(v []string) {
	o.TaskAttributeValue = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StaticallyDefinedRecurringTaskResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticallyDefinedRecurringTaskResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StaticallyDefinedRecurringTaskResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StaticallyDefinedRecurringTaskResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field value if set, zero value otherwise.
func (o *StaticallyDefinedRecurringTaskResponse) GetCancelOnTaskDependencyFailure() bool {
	if o == nil || IsNil(o.CancelOnTaskDependencyFailure) {
		var ret bool
		return ret
	}
	return *o.CancelOnTaskDependencyFailure
}

// GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticallyDefinedRecurringTaskResponse) GetCancelOnTaskDependencyFailureOk() (*bool, bool) {
	if o == nil || IsNil(o.CancelOnTaskDependencyFailure) {
		return nil, false
	}
	return o.CancelOnTaskDependencyFailure, true
}

// HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.
func (o *StaticallyDefinedRecurringTaskResponse) HasCancelOnTaskDependencyFailure() bool {
	if o != nil && !IsNil(o.CancelOnTaskDependencyFailure) {
		return true
	}

	return false
}

// SetCancelOnTaskDependencyFailure gets a reference to the given bool and assigns it to the CancelOnTaskDependencyFailure field.
func (o *StaticallyDefinedRecurringTaskResponse) SetCancelOnTaskDependencyFailure(v bool) {
	o.CancelOnTaskDependencyFailure = &v
}

// GetEmailOnStart returns the EmailOnStart field value if set, zero value otherwise.
func (o *StaticallyDefinedRecurringTaskResponse) GetEmailOnStart() []string {
	if o == nil || IsNil(o.EmailOnStart) {
		var ret []string
		return ret
	}
	return o.EmailOnStart
}

// GetEmailOnStartOk returns a tuple with the EmailOnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticallyDefinedRecurringTaskResponse) GetEmailOnStartOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailOnStart) {
		return nil, false
	}
	return o.EmailOnStart, true
}

// HasEmailOnStart returns a boolean if a field has been set.
func (o *StaticallyDefinedRecurringTaskResponse) HasEmailOnStart() bool {
	if o != nil && !IsNil(o.EmailOnStart) {
		return true
	}

	return false
}

// SetEmailOnStart gets a reference to the given []string and assigns it to the EmailOnStart field.
func (o *StaticallyDefinedRecurringTaskResponse) SetEmailOnStart(v []string) {
	o.EmailOnStart = v
}

// GetEmailOnSuccess returns the EmailOnSuccess field value if set, zero value otherwise.
func (o *StaticallyDefinedRecurringTaskResponse) GetEmailOnSuccess() []string {
	if o == nil || IsNil(o.EmailOnSuccess) {
		var ret []string
		return ret
	}
	return o.EmailOnSuccess
}

// GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticallyDefinedRecurringTaskResponse) GetEmailOnSuccessOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailOnSuccess) {
		return nil, false
	}
	return o.EmailOnSuccess, true
}

// HasEmailOnSuccess returns a boolean if a field has been set.
func (o *StaticallyDefinedRecurringTaskResponse) HasEmailOnSuccess() bool {
	if o != nil && !IsNil(o.EmailOnSuccess) {
		return true
	}

	return false
}

// SetEmailOnSuccess gets a reference to the given []string and assigns it to the EmailOnSuccess field.
func (o *StaticallyDefinedRecurringTaskResponse) SetEmailOnSuccess(v []string) {
	o.EmailOnSuccess = v
}

// GetEmailOnFailure returns the EmailOnFailure field value if set, zero value otherwise.
func (o *StaticallyDefinedRecurringTaskResponse) GetEmailOnFailure() []string {
	if o == nil || IsNil(o.EmailOnFailure) {
		var ret []string
		return ret
	}
	return o.EmailOnFailure
}

// GetEmailOnFailureOk returns a tuple with the EmailOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticallyDefinedRecurringTaskResponse) GetEmailOnFailureOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailOnFailure) {
		return nil, false
	}
	return o.EmailOnFailure, true
}

// HasEmailOnFailure returns a boolean if a field has been set.
func (o *StaticallyDefinedRecurringTaskResponse) HasEmailOnFailure() bool {
	if o != nil && !IsNil(o.EmailOnFailure) {
		return true
	}

	return false
}

// SetEmailOnFailure gets a reference to the given []string and assigns it to the EmailOnFailure field.
func (o *StaticallyDefinedRecurringTaskResponse) SetEmailOnFailure(v []string) {
	o.EmailOnFailure = v
}

// GetAlertOnStart returns the AlertOnStart field value if set, zero value otherwise.
func (o *StaticallyDefinedRecurringTaskResponse) GetAlertOnStart() bool {
	if o == nil || IsNil(o.AlertOnStart) {
		var ret bool
		return ret
	}
	return *o.AlertOnStart
}

// GetAlertOnStartOk returns a tuple with the AlertOnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticallyDefinedRecurringTaskResponse) GetAlertOnStartOk() (*bool, bool) {
	if o == nil || IsNil(o.AlertOnStart) {
		return nil, false
	}
	return o.AlertOnStart, true
}

// HasAlertOnStart returns a boolean if a field has been set.
func (o *StaticallyDefinedRecurringTaskResponse) HasAlertOnStart() bool {
	if o != nil && !IsNil(o.AlertOnStart) {
		return true
	}

	return false
}

// SetAlertOnStart gets a reference to the given bool and assigns it to the AlertOnStart field.
func (o *StaticallyDefinedRecurringTaskResponse) SetAlertOnStart(v bool) {
	o.AlertOnStart = &v
}

// GetAlertOnSuccess returns the AlertOnSuccess field value if set, zero value otherwise.
func (o *StaticallyDefinedRecurringTaskResponse) GetAlertOnSuccess() bool {
	if o == nil || IsNil(o.AlertOnSuccess) {
		var ret bool
		return ret
	}
	return *o.AlertOnSuccess
}

// GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticallyDefinedRecurringTaskResponse) GetAlertOnSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.AlertOnSuccess) {
		return nil, false
	}
	return o.AlertOnSuccess, true
}

// HasAlertOnSuccess returns a boolean if a field has been set.
func (o *StaticallyDefinedRecurringTaskResponse) HasAlertOnSuccess() bool {
	if o != nil && !IsNil(o.AlertOnSuccess) {
		return true
	}

	return false
}

// SetAlertOnSuccess gets a reference to the given bool and assigns it to the AlertOnSuccess field.
func (o *StaticallyDefinedRecurringTaskResponse) SetAlertOnSuccess(v bool) {
	o.AlertOnSuccess = &v
}

// GetAlertOnFailure returns the AlertOnFailure field value if set, zero value otherwise.
func (o *StaticallyDefinedRecurringTaskResponse) GetAlertOnFailure() bool {
	if o == nil || IsNil(o.AlertOnFailure) {
		var ret bool
		return ret
	}
	return *o.AlertOnFailure
}

// GetAlertOnFailureOk returns a tuple with the AlertOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticallyDefinedRecurringTaskResponse) GetAlertOnFailureOk() (*bool, bool) {
	if o == nil || IsNil(o.AlertOnFailure) {
		return nil, false
	}
	return o.AlertOnFailure, true
}

// HasAlertOnFailure returns a boolean if a field has been set.
func (o *StaticallyDefinedRecurringTaskResponse) HasAlertOnFailure() bool {
	if o != nil && !IsNil(o.AlertOnFailure) {
		return true
	}

	return false
}

// SetAlertOnFailure gets a reference to the given bool and assigns it to the AlertOnFailure field.
func (o *StaticallyDefinedRecurringTaskResponse) SetAlertOnFailure(v bool) {
	o.AlertOnFailure = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *StaticallyDefinedRecurringTaskResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticallyDefinedRecurringTaskResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *StaticallyDefinedRecurringTaskResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *StaticallyDefinedRecurringTaskResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *StaticallyDefinedRecurringTaskResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticallyDefinedRecurringTaskResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *StaticallyDefinedRecurringTaskResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *StaticallyDefinedRecurringTaskResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *StaticallyDefinedRecurringTaskResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StaticallyDefinedRecurringTaskResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StaticallyDefinedRecurringTaskResponse) SetId(v string) {
	o.Id = v
}

func (o StaticallyDefinedRecurringTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticallyDefinedRecurringTaskResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["taskJavaClass"] = o.TaskJavaClass
	toSerialize["taskObjectClass"] = o.TaskObjectClass
	if !IsNil(o.TaskAttributeValue) {
		toSerialize["taskAttributeValue"] = o.TaskAttributeValue
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CancelOnTaskDependencyFailure) {
		toSerialize["cancelOnTaskDependencyFailure"] = o.CancelOnTaskDependencyFailure
	}
	if !IsNil(o.EmailOnStart) {
		toSerialize["emailOnStart"] = o.EmailOnStart
	}
	if !IsNil(o.EmailOnSuccess) {
		toSerialize["emailOnSuccess"] = o.EmailOnSuccess
	}
	if !IsNil(o.EmailOnFailure) {
		toSerialize["emailOnFailure"] = o.EmailOnFailure
	}
	if !IsNil(o.AlertOnStart) {
		toSerialize["alertOnStart"] = o.AlertOnStart
	}
	if !IsNil(o.AlertOnSuccess) {
		toSerialize["alertOnSuccess"] = o.AlertOnSuccess
	}
	if !IsNil(o.AlertOnFailure) {
		toSerialize["alertOnFailure"] = o.AlertOnFailure
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

type NullableStaticallyDefinedRecurringTaskResponse struct {
	value *StaticallyDefinedRecurringTaskResponse
	isSet bool
}

func (v NullableStaticallyDefinedRecurringTaskResponse) Get() *StaticallyDefinedRecurringTaskResponse {
	return v.value
}

func (v *NullableStaticallyDefinedRecurringTaskResponse) Set(val *StaticallyDefinedRecurringTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticallyDefinedRecurringTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticallyDefinedRecurringTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticallyDefinedRecurringTaskResponse(val *StaticallyDefinedRecurringTaskResponse) *NullableStaticallyDefinedRecurringTaskResponse {
	return &NullableStaticallyDefinedRecurringTaskResponse{value: val, isSet: true}
}

func (v NullableStaticallyDefinedRecurringTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticallyDefinedRecurringTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
