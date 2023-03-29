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

// checks if the ThirdPartyAccountStatusNotificationHandlerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyAccountStatusNotificationHandlerResponse{}

// ThirdPartyAccountStatusNotificationHandlerResponse struct for ThirdPartyAccountStatusNotificationHandlerResponse
type ThirdPartyAccountStatusNotificationHandlerResponse struct {
	// Name of the Account Status Notification Handler
	Id      string                                                    `json:"id"`
	Schemas []EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Account Status Notification Handler.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Account Status Notification Handler. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Account Status Notification Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Account Status Notification Handler is enabled. Only enabled handlers are invoked whenever a related event occurs in the server.
	Enabled bool `json:"enabled"`
	// Indicates whether the server should attempt to invoke this Account Status Notification Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver a message) will not delay processing for the operation that triggered the notification.
	Asynchronous *bool `json:"asynchronous,omitempty"`
	// A request criteria object that identifies which add requests should result in account creation notifications for this handler.
	AccountCreationNotificationRequestCriteria *string `json:"accountCreationNotificationRequestCriteria,omitempty"`
	// A request criteria object that identifies which modify and modify DN requests should result in account update notifications for this handler.
	AccountUpdateNotificationRequestCriteria      *string                                            `json:"accountUpdateNotificationRequestCriteria,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewThirdPartyAccountStatusNotificationHandlerResponse instantiates a new ThirdPartyAccountStatusNotificationHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyAccountStatusNotificationHandlerResponse(id string, schemas []EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn, extensionClass string, enabled bool) *ThirdPartyAccountStatusNotificationHandlerResponse {
	this := ThirdPartyAccountStatusNotificationHandlerResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewThirdPartyAccountStatusNotificationHandlerResponseWithDefaults instantiates a new ThirdPartyAccountStatusNotificationHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyAccountStatusNotificationHandlerResponseWithDefaults() *ThirdPartyAccountStatusNotificationHandlerResponse {
	this := ThirdPartyAccountStatusNotificationHandlerResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetSchemas() []EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetSchemasOk() ([]EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) SetSchemas(v []EnumthirdPartyAccountStatusNotificationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetExtensionArgument() []string {
	if o == nil || IsNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) HasExtensionArgument() bool {
	if o != nil && !IsNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAsynchronous returns the Asynchronous field value if set, zero value otherwise.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetAsynchronous() bool {
	if o == nil || IsNil(o.Asynchronous) {
		var ret bool
		return ret
	}
	return *o.Asynchronous
}

// GetAsynchronousOk returns a tuple with the Asynchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetAsynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Asynchronous) {
		return nil, false
	}
	return o.Asynchronous, true
}

// HasAsynchronous returns a boolean if a field has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) HasAsynchronous() bool {
	if o != nil && !IsNil(o.Asynchronous) {
		return true
	}

	return false
}

// SetAsynchronous gets a reference to the given bool and assigns it to the Asynchronous field.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) SetAsynchronous(v bool) {
	o.Asynchronous = &v
}

// GetAccountCreationNotificationRequestCriteria returns the AccountCreationNotificationRequestCriteria field value if set, zero value otherwise.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetAccountCreationNotificationRequestCriteria() string {
	if o == nil || IsNil(o.AccountCreationNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountCreationNotificationRequestCriteria
}

// GetAccountCreationNotificationRequestCriteriaOk returns a tuple with the AccountCreationNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetAccountCreationNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AccountCreationNotificationRequestCriteria) {
		return nil, false
	}
	return o.AccountCreationNotificationRequestCriteria, true
}

// HasAccountCreationNotificationRequestCriteria returns a boolean if a field has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) HasAccountCreationNotificationRequestCriteria() bool {
	if o != nil && !IsNil(o.AccountCreationNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountCreationNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountCreationNotificationRequestCriteria field.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) SetAccountCreationNotificationRequestCriteria(v string) {
	o.AccountCreationNotificationRequestCriteria = &v
}

// GetAccountUpdateNotificationRequestCriteria returns the AccountUpdateNotificationRequestCriteria field value if set, zero value otherwise.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetAccountUpdateNotificationRequestCriteria() string {
	if o == nil || IsNil(o.AccountUpdateNotificationRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AccountUpdateNotificationRequestCriteria
}

// GetAccountUpdateNotificationRequestCriteriaOk returns a tuple with the AccountUpdateNotificationRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetAccountUpdateNotificationRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AccountUpdateNotificationRequestCriteria) {
		return nil, false
	}
	return o.AccountUpdateNotificationRequestCriteria, true
}

// HasAccountUpdateNotificationRequestCriteria returns a boolean if a field has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) HasAccountUpdateNotificationRequestCriteria() bool {
	if o != nil && !IsNil(o.AccountUpdateNotificationRequestCriteria) {
		return true
	}

	return false
}

// SetAccountUpdateNotificationRequestCriteria gets a reference to the given string and assigns it to the AccountUpdateNotificationRequestCriteria field.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) SetAccountUpdateNotificationRequestCriteria(v string) {
	o.AccountUpdateNotificationRequestCriteria = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ThirdPartyAccountStatusNotificationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o ThirdPartyAccountStatusNotificationHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyAccountStatusNotificationHandlerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["extensionClass"] = o.ExtensionClass
	if !IsNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Asynchronous) {
		toSerialize["asynchronous"] = o.Asynchronous
	}
	if !IsNil(o.AccountCreationNotificationRequestCriteria) {
		toSerialize["accountCreationNotificationRequestCriteria"] = o.AccountCreationNotificationRequestCriteria
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

type NullableThirdPartyAccountStatusNotificationHandlerResponse struct {
	value *ThirdPartyAccountStatusNotificationHandlerResponse
	isSet bool
}

func (v NullableThirdPartyAccountStatusNotificationHandlerResponse) Get() *ThirdPartyAccountStatusNotificationHandlerResponse {
	return v.value
}

func (v *NullableThirdPartyAccountStatusNotificationHandlerResponse) Set(val *ThirdPartyAccountStatusNotificationHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyAccountStatusNotificationHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyAccountStatusNotificationHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyAccountStatusNotificationHandlerResponse(val *ThirdPartyAccountStatusNotificationHandlerResponse) *NullableThirdPartyAccountStatusNotificationHandlerResponse {
	return &NullableThirdPartyAccountStatusNotificationHandlerResponse{value: val, isSet: true}
}

func (v NullableThirdPartyAccountStatusNotificationHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyAccountStatusNotificationHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
