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

// checks if the ErrorLogAccountStatusNotificationHandlerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorLogAccountStatusNotificationHandlerResponse{}

// ErrorLogAccountStatusNotificationHandlerResponse struct for ErrorLogAccountStatusNotificationHandlerResponse
type ErrorLogAccountStatusNotificationHandlerResponse struct {
	Schemas                       []EnumerrorLogAccountStatusNotificationHandlerSchemaUrn                 `json:"schemas"`
	AccountStatusNotificationType []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp `json:"accountStatusNotificationType"`
	// A description for this Account Status Notification Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Account Status Notification Handler is enabled. Only enabled handlers are invoked whenever a related event occurs in the server.
	Enabled bool `json:"enabled"`
	// Indicates whether the server should attempt to invoke this Account Status Notification Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver a message) will not delay processing for the operation that triggered the notification.
	Asynchronous *bool `json:"asynchronous,omitempty"`
	// A result criteria object that identifies which successful bind operations should result in account authentication notifications for this handler.
	AccountAuthenticationNotificationResultCriteria *string `json:"accountAuthenticationNotificationResultCriteria,omitempty"`
	// A request criteria object that identifies which add requests should result in account creation notifications for this handler.
	AccountCreationNotificationRequestCriteria *string `json:"accountCreationNotificationRequestCriteria,omitempty"`
	// A request criteria object that identifies which delete requests should result in account deletion notifications for this handler.
	AccountDeletionNotificationRequestCriteria *string `json:"accountDeletionNotificationRequestCriteria,omitempty"`
	// A request criteria object that identifies which modify and modify DN requests should result in account update notifications for this handler.
	AccountUpdateNotificationRequestCriteria      *string                                            `json:"accountUpdateNotificationRequestCriteria,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Account Status Notification Handler
	Id string `json:"id"`
}

// NewErrorLogAccountStatusNotificationHandlerResponse instantiates a new ErrorLogAccountStatusNotificationHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorLogAccountStatusNotificationHandlerResponse(schemas []EnumerrorLogAccountStatusNotificationHandlerSchemaUrn, accountStatusNotificationType []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp, enabled bool, id string) *ErrorLogAccountStatusNotificationHandlerResponse {
	this := ErrorLogAccountStatusNotificationHandlerResponse{}
	this.Schemas = schemas
	this.AccountStatusNotificationType = accountStatusNotificationType
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewErrorLogAccountStatusNotificationHandlerResponseWithDefaults instantiates a new ErrorLogAccountStatusNotificationHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorLogAccountStatusNotificationHandlerResponseWithDefaults() *ErrorLogAccountStatusNotificationHandlerResponse {
	this := ErrorLogAccountStatusNotificationHandlerResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetSchemas() []EnumerrorLogAccountStatusNotificationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumerrorLogAccountStatusNotificationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetSchemasOk() ([]EnumerrorLogAccountStatusNotificationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ErrorLogAccountStatusNotificationHandlerResponse) SetSchemas(v []EnumerrorLogAccountStatusNotificationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetAccountStatusNotificationType returns the AccountStatusNotificationType field value
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetAccountStatusNotificationType() []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp {
	if o == nil {
		var ret []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp
		return ret
	}

	return o.AccountStatusNotificationType
}

// GetAccountStatusNotificationTypeOk returns a tuple with the AccountStatusNotificationType field value
// and a boolean to check if the value has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetAccountStatusNotificationTypeOk() ([]EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountStatusNotificationType, true
}

// SetAccountStatusNotificationType sets field value
func (o *ErrorLogAccountStatusNotificationHandlerResponse) SetAccountStatusNotificationType(v []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp) {
	o.AccountStatusNotificationType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ErrorLogAccountStatusNotificationHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetAsynchronous() bool {
	if o == nil || IsNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetAsynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Asynchronous) {
		return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) HasAsynchronous() bool {
	if o != nil && !IsNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetAccountAuthenticationNotificationResultCriteria returns the AccountAuthenticationNotificationResultCriteria field value if set, zero value otherwise.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetAccountAuthenticationNotificationResultCriteria() string {
	if o == nil || IsNil(o.AccountAuthenticationNotificationResultCriteria) {
		var ret string
		return ret
	}
	return *o.AccountAuthenticationNotificationResultCriteria
}

// GetAccountAuthenticationNotificationResultCriteriaOk returns a tuple with the AccountAuthenticationNotificationResultCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetAccountAuthenticationNotificationResultCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AccountAuthenticationNotificationResultCriteria) {
		return nil, false
	}
	return o.AccountAuthenticationNotificationResultCriteria, true
}

// HasAccountAuthenticationNotificationResultCriteria returns a boolean if a field has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) HasAccountAuthenticationNotificationResultCriteria() bool {
	if o != nil && !IsNil(o.AccountAuthenticationNotificationResultCriteria) {
		return true
	}

	return false
}

// SetAccountAuthenticationNotificationResultCriteria gets a reference to the given string and assigns it to the AccountAuthenticationNotificationResultCriteria field.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) SetAccountAuthenticationNotificationResultCriteria(v string) {
	o.AccountAuthenticationNotificationResultCriteria = &v
}

// GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field value if set, zero value otherwise.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetAccountCreationNotificationRequestCriteria() string {
	if o == nil || IsNil(o.AccountCreationNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountCreationNotificationRequestCriteria
}

// GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AccountCreationNotificationRequestCriteria) {
		return nil, false
	}
	return o.AccountCreationNotificationRequestCriteria, true
}

// HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) HasAccountCreationNotificationRequestCriteria() bool {
	if o != nil && !IsNil(o.AccountCreationNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountCreationNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountCreationNotificationRequestCriteria field.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) SetAccountCreationNotificationRequestCriteria(v string) {
	o.AccountCreationNotificationRequestCriteria = &v
}

// GetAccountDeletionNotificationRequestCriteria returns the AccountDeletionNotificationRequestCriteria field value if set, zero value otherwise.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetAccountDeletionNotificationRequestCriteria() string {
	if o == nil || IsNil(o.AccountDeletionNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountDeletionNotificationRequestCriteria
}

// GetAccountDeletionNotificationRequestCriteriaOk returns a tuple with the AccountDeletionNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetAccountDeletionNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AccountDeletionNotificationRequestCriteria) {
		return nil, false
	}
	return o.AccountDeletionNotificationRequestCriteria, true
}

// HasAccountDeletionNotificationRequestCriteria returns a boolean if a field has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) HasAccountDeletionNotificationRequestCriteria() bool {
	if o != nil && !IsNil(o.AccountDeletionNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountDeletionNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountDeletionNotificationRequestCriteria field.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) SetAccountDeletionNotificationRequestCriteria(v string) {
	o.AccountDeletionNotificationRequestCriteria = &v
}

// GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field value if set, zero value otherwise.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetAccountUpdateNotificationRequestCriteria() string {
	if o == nil || IsNil(o.AccountUpdateNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountUpdateNotificationRequestCriteria
}

// GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AccountUpdateNotificationRequestCriteria) {
		return nil, false
	}
	return o.AccountUpdateNotificationRequestCriteria, true
}

// HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) HasAccountUpdateNotificationRequestCriteria() bool {
	if o != nil && !IsNil(o.AccountUpdateNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountUpdateNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountUpdateNotificationRequestCriteria field.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) SetAccountUpdateNotificationRequestCriteria(v string) {
	o.AccountUpdateNotificationRequestCriteria = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ErrorLogAccountStatusNotificationHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ErrorLogAccountStatusNotificationHandlerResponse) SetId(v string) {
	o.Id = v
}

func (o ErrorLogAccountStatusNotificationHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorLogAccountStatusNotificationHandlerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["accountStatusNotificationType"] = o.AccountStatusNotificationType
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Asynchronous) {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	if !IsNil(o.AccountAuthenticationNotificationResultCriteria) {
		toSerialize["accountAuthenticationNotificationResultCriteria"] = o.AccountAuthenticationNotificationResultCriteria
	}
	if !IsNil(o.AccountCreationNotificationRequestCriteria) {
		toSerialize["accountCreationNotificationRequestCriteria"] = o.AccountCreationNotificationRequestCriteria
	}
	if !IsNil(o.AccountDeletionNotificationRequestCriteria) {
		toSerialize["accountDeletionNotificationRequestCriteria"] = o.AccountDeletionNotificationRequestCriteria
	}
	if !IsNil(o.AccountUpdateNotificationRequestCriteria) {
		toSerialize["accountUpdateNotificationRequestCriteria"] = o.AccountUpdateNotificationRequestCriteria
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

type NullableErrorLogAccountStatusNotificationHandlerResponse struct {
	value *ErrorLogAccountStatusNotificationHandlerResponse
	isSet bool
}

func (v NullableErrorLogAccountStatusNotificationHandlerResponse) Get() *ErrorLogAccountStatusNotificationHandlerResponse {
	return v.value
}

func (v *NullableErrorLogAccountStatusNotificationHandlerResponse) Set(val *ErrorLogAccountStatusNotificationHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorLogAccountStatusNotificationHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorLogAccountStatusNotificationHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorLogAccountStatusNotificationHandlerResponse(val *ErrorLogAccountStatusNotificationHandlerResponse) *NullableErrorLogAccountStatusNotificationHandlerResponse {
	return &NullableErrorLogAccountStatusNotificationHandlerResponse{value: val, isSet: true}
}

func (v NullableErrorLogAccountStatusNotificationHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorLogAccountStatusNotificationHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
