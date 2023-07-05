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

// checks if the AdminAlertAccountStatusNotificationHandlerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminAlertAccountStatusNotificationHandlerResponse{}

// AdminAlertAccountStatusNotificationHandlerResponse struct for AdminAlertAccountStatusNotificationHandlerResponse
type AdminAlertAccountStatusNotificationHandlerResponse struct {
	// Name of the Account Status Notification Handler
	Id                            string                                                                  `json:"id"`
	Schemas                       []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn               `json:"schemas"`
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
}

// NewAdminAlertAccountStatusNotificationHandlerResponse instantiates a new AdminAlertAccountStatusNotificationHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminAlertAccountStatusNotificationHandlerResponse(id string, schemas []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn, accountStatusNotificationType []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp, enabled bool) *AdminAlertAccountStatusNotificationHandlerResponse {
	this := AdminAlertAccountStatusNotificationHandlerResponse{}
	this.Id = id
	this.Schemas = schemas
	this.AccountStatusNotificationType = accountStatusNotificationType
	this.Enabled = enabled
	return &this
}

// NewAdminAlertAccountStatusNotificationHandlerResponseWithDefaults instantiates a new AdminAlertAccountStatusNotificationHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminAlertAccountStatusNotificationHandlerResponseWithDefaults() *AdminAlertAccountStatusNotificationHandlerResponse {
	this := AdminAlertAccountStatusNotificationHandlerResponse{}
	return &this
}

// GetId returns the Id field value
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetSchemas() []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetSchemasOk() ([]EnumadminAlertAccountStatusNotificationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetSchemas(v []EnumadminAlertAccountStatusNotificationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetAccountStatusNotificationType returns the AccountStatusNotificationType field value
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountStatusNotificationType() []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp {
	if o == nil {
		var ret []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp
		return ret
	}

	return o.AccountStatusNotificationType
}

// GetAccountStatusNotificationTypeOk returns a tuple with the AccountStatusNotificationType field value
// and a boolean to check if the value has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountStatusNotificationTypeOk() ([]EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountStatusNotificationType, true
}

// SetAccountStatusNotificationType sets field value
func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetAccountStatusNotificationType(v []EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp) {
	o.AccountStatusNotificationType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAsynchronous() bool {
	if o == nil || IsNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAsynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Asynchronous) {
		return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) HasAsynchronous() bool {
	if o != nil && !IsNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetAccountAuthenticationNotificationResultCriteria returns the AccountAuthenticationNotificationResultCriteria field value if set, zero value otherwise.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountAuthenticationNotificationResultCriteria() string {
	if o == nil || IsNil(o.AccountAuthenticationNotificationResultCriteria) {
		var ret string
		return ret
	}
	return *o.AccountAuthenticationNotificationResultCriteria
}

// GetAccountAuthenticationNotificationResultCriteriaOk returns a tuple with the AccountAuthenticationNotificationResultCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountAuthenticationNotificationResultCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AccountAuthenticationNotificationResultCriteria) {
		return nil, false
	}
	return o.AccountAuthenticationNotificationResultCriteria, true
}

// HasAccountAuthenticationNotificationResultCriteria returns a boolean if a field has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) HasAccountAuthenticationNotificationResultCriteria() bool {
	if o != nil && !IsNil(o.AccountAuthenticationNotificationResultCriteria) {
		return true
	}

	return false
}

// SetAccountAuthenticationNotificationResultCriteria gets a reference to the given string and assigns it to the AccountAuthenticationNotificationResultCriteria field.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetAccountAuthenticationNotificationResultCriteria(v string) {
	o.AccountAuthenticationNotificationResultCriteria = &v
}

// GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field value if set, zero value otherwise.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountCreationNotificationRequestCriteria() string {
	if o == nil || IsNil(o.AccountCreationNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountCreationNotificationRequestCriteria
}

// GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AccountCreationNotificationRequestCriteria) {
		return nil, false
	}
	return o.AccountCreationNotificationRequestCriteria, true
}

// HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) HasAccountCreationNotificationRequestCriteria() bool {
	if o != nil && !IsNil(o.AccountCreationNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountCreationNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountCreationNotificationRequestCriteria field.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetAccountCreationNotificationRequestCriteria(v string) {
	o.AccountCreationNotificationRequestCriteria = &v
}

// GetAccountDeletionNotificationRequestCriteria returns the AccountDeletionNotificationRequestCriteria field value if set, zero value otherwise.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountDeletionNotificationRequestCriteria() string {
	if o == nil || IsNil(o.AccountDeletionNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountDeletionNotificationRequestCriteria
}

// GetAccountDeletionNotificationRequestCriteriaOk returns a tuple with the AccountDeletionNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountDeletionNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AccountDeletionNotificationRequestCriteria) {
		return nil, false
	}
	return o.AccountDeletionNotificationRequestCriteria, true
}

// HasAccountDeletionNotificationRequestCriteria returns a boolean if a field has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) HasAccountDeletionNotificationRequestCriteria() bool {
	if o != nil && !IsNil(o.AccountDeletionNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountDeletionNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountDeletionNotificationRequestCriteria field.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetAccountDeletionNotificationRequestCriteria(v string) {
	o.AccountDeletionNotificationRequestCriteria = &v
}

// GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field value if set, zero value otherwise.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountUpdateNotificationRequestCriteria() string {
	if o == nil || IsNil(o.AccountUpdateNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountUpdateNotificationRequestCriteria
}

// GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AccountUpdateNotificationRequestCriteria) {
		return nil, false
	}
	return o.AccountUpdateNotificationRequestCriteria, true
}

// HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) HasAccountUpdateNotificationRequestCriteria() bool {
	if o != nil && !IsNil(o.AccountUpdateNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountUpdateNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountUpdateNotificationRequestCriteria field.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetAccountUpdateNotificationRequestCriteria(v string) {
	o.AccountUpdateNotificationRequestCriteria = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *AdminAlertAccountStatusNotificationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o AdminAlertAccountStatusNotificationHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminAlertAccountStatusNotificationHandlerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
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
	return toSerialize, nil
}

type NullableAdminAlertAccountStatusNotificationHandlerResponse struct {
	value *AdminAlertAccountStatusNotificationHandlerResponse
	isSet bool
}

func (v NullableAdminAlertAccountStatusNotificationHandlerResponse) Get() *AdminAlertAccountStatusNotificationHandlerResponse {
	return v.value
}

func (v *NullableAdminAlertAccountStatusNotificationHandlerResponse) Set(val *AdminAlertAccountStatusNotificationHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminAlertAccountStatusNotificationHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminAlertAccountStatusNotificationHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminAlertAccountStatusNotificationHandlerResponse(val *AdminAlertAccountStatusNotificationHandlerResponse) *NullableAdminAlertAccountStatusNotificationHandlerResponse {
	return &NullableAdminAlertAccountStatusNotificationHandlerResponse{value: val, isSet: true}
}

func (v NullableAdminAlertAccountStatusNotificationHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminAlertAccountStatusNotificationHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
