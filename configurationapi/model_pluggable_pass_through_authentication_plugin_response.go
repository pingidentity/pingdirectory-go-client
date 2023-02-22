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

// PluggablePassThroughAuthenticationPluginResponse struct for PluggablePassThroughAuthenticationPluginResponse
type PluggablePassThroughAuthenticationPluginResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Plugin
	Id      string                                                  `json:"id"`
	Schemas []EnumpluggablePassThroughAuthenticationPluginSchemaUrn `json:"schemas"`
	// The component used to manage authentication with the external authentication service.
	PassThroughAuthenticationHandler string `json:"passThroughAuthenticationHandler"`
	// The base DNs for the local users whose authentication attempts may be passed through to the external authentication service.
	IncludedLocalEntryBaseDN []string `json:"includedLocalEntryBaseDN,omitempty"`
	// A reference to connection criteria that will be used to indicate which bind requests should be passed through to the external authentication service.
	ConnectionCriteria *string `json:"connectionCriteria,omitempty"`
	// A reference to request criteria that will be used to indicate which bind requests should be passed through to the external authentication service.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
	// Indicates whether to attempt the bind in the local server first and only send the request to the external authentication service if the local bind attempt fails, or to only attempt the bind in the external service.
	TryLocalBind *bool `json:"tryLocalBind,omitempty"`
	// Indicates whether to attempt the authentication in the external service if the local user entry includes a password. This property will be ignored if try-local-bind is false.
	OverrideLocalPassword *bool `json:"overrideLocalPassword,omitempty"`
	// Indicates whether to overwrite the user's local password if the local bind fails but the authentication attempt succeeds when attempted in the external service. This property may only be set to true if try-local-bind is also true.
	UpdateLocalPassword *bool `json:"updateLocalPassword,omitempty"`
	// The DN of the authorization identity that will be used when updating the user's local password if update-local-password is true. This is primarily intended for use if the Data Sync Server will be used to synchronize passwords between the local server and the external service, and in that case, the DN used here should also be added to the ignore-changes-by-dn property in the appropriate Sync Source object in the Data Sync Server configuration.
	UpdateLocalPasswordDN *string `json:"updateLocalPasswordDN,omitempty"`
	// Indicates whether to overwrite the user's local password even if the password used to authenticate to the external service would have failed validation if the user attempted to set it directly.
	AllowLaxPassThroughAuthenticationPasswords *bool                                                    `json:"allowLaxPassThroughAuthenticationPasswords,omitempty"`
	IgnoredPasswordPolicyStateErrorCondition   []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp `json:"ignoredPasswordPolicyStateErrorCondition,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
	// Indicates whether the plug-in should be invoked for internal operations.
	InvokeForInternalOperations *bool `json:"invokeForInternalOperations,omitempty"`
}

// NewPluggablePassThroughAuthenticationPluginResponse instantiates a new PluggablePassThroughAuthenticationPluginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluggablePassThroughAuthenticationPluginResponse(id string, schemas []EnumpluggablePassThroughAuthenticationPluginSchemaUrn, passThroughAuthenticationHandler string, enabled bool) *PluggablePassThroughAuthenticationPluginResponse {
	this := PluggablePassThroughAuthenticationPluginResponse{}
	this.Id = id
	this.Schemas = schemas
	this.PassThroughAuthenticationHandler = passThroughAuthenticationHandler
	this.Enabled = enabled
	return &this
}

// NewPluggablePassThroughAuthenticationPluginResponseWithDefaults instantiates a new PluggablePassThroughAuthenticationPluginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluggablePassThroughAuthenticationPluginResponseWithDefaults() *PluggablePassThroughAuthenticationPluginResponse {
	this := PluggablePassThroughAuthenticationPluginResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *PluggablePassThroughAuthenticationPluginResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *PluggablePassThroughAuthenticationPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *PluggablePassThroughAuthenticationPluginResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PluggablePassThroughAuthenticationPluginResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *PluggablePassThroughAuthenticationPluginResponse) GetSchemas() []EnumpluggablePassThroughAuthenticationPluginSchemaUrn {
	if o == nil {
		var ret []EnumpluggablePassThroughAuthenticationPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetSchemasOk() ([]EnumpluggablePassThroughAuthenticationPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *PluggablePassThroughAuthenticationPluginResponse) SetSchemas(v []EnumpluggablePassThroughAuthenticationPluginSchemaUrn) {
	o.Schemas = v
}

// GetPassThroughAuthenticationHandler returns the PassThroughAuthenticationHandler field value
func (o *PluggablePassThroughAuthenticationPluginResponse) GetPassThroughAuthenticationHandler() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PassThroughAuthenticationHandler
}

// GetPassThroughAuthenticationHandlerOk returns a tuple with the PassThroughAuthenticationHandler field value
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetPassThroughAuthenticationHandlerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PassThroughAuthenticationHandler, true
}

// SetPassThroughAuthenticationHandler sets field value
func (o *PluggablePassThroughAuthenticationPluginResponse) SetPassThroughAuthenticationHandler(v string) {
	o.PassThroughAuthenticationHandler = v
}

// GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field value if set, zero value otherwise.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetIncludedLocalEntryBaseDN() []string {
	if o == nil || isNil(o.IncludedLocalEntryBaseDN) {
		var ret []string
		return ret
	}
	return o.IncludedLocalEntryBaseDN
}

// GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetIncludedLocalEntryBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.IncludedLocalEntryBaseDN) {
		return nil, false
	}
	return o.IncludedLocalEntryBaseDN, true
}

// HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) HasIncludedLocalEntryBaseDN() bool {
	if o != nil && !isNil(o.IncludedLocalEntryBaseDN) {
		return true
	}

	return false
}

// SetIncludedLocalEntryBaseDN gets a reference to the given []string and assigns it to the IncludedLocalEntryBaseDN field.
func (o *PluggablePassThroughAuthenticationPluginResponse) SetIncludedLocalEntryBaseDN(v []string) {
	o.IncludedLocalEntryBaseDN = v
}

// GetConnectionCriteria returns the ConnectionCriteria field value if set, zero value otherwise.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetConnectionCriteria() string {
	if o == nil || isNil(o.ConnectionCriteria) {
		var ret string
		return ret
	}
	return *o.ConnectionCriteria
}

// GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetConnectionCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.ConnectionCriteria) {
		return nil, false
	}
	return o.ConnectionCriteria, true
}

// HasConnectionCriteria returns a boolean if a field has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) HasConnectionCriteria() bool {
	if o != nil && !isNil(o.ConnectionCriteria) {
		return true
	}

	return false
}

// SetConnectionCriteria gets a reference to the given string and assigns it to the ConnectionCriteria field.
func (o *PluggablePassThroughAuthenticationPluginResponse) SetConnectionCriteria(v string) {
	o.ConnectionCriteria = &v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetRequestCriteria() string {
	if o == nil || isNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) HasRequestCriteria() bool {
	if o != nil && !isNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *PluggablePassThroughAuthenticationPluginResponse) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetTryLocalBind returns the TryLocalBind field value if set, zero value otherwise.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetTryLocalBind() bool {
	if o == nil || isNil(o.TryLocalBind) {
		var ret bool
		return ret
	}
	return *o.TryLocalBind
}

// GetTryLocalBindOk returns a tuple with the TryLocalBind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetTryLocalBindOk() (*bool, bool) {
	if o == nil || isNil(o.TryLocalBind) {
		return nil, false
	}
	return o.TryLocalBind, true
}

// HasTryLocalBind returns a boolean if a field has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) HasTryLocalBind() bool {
	if o != nil && !isNil(o.TryLocalBind) {
		return true
	}

	return false
}

// SetTryLocalBind gets a reference to the given bool and assigns it to the TryLocalBind field.
func (o *PluggablePassThroughAuthenticationPluginResponse) SetTryLocalBind(v bool) {
	o.TryLocalBind = &v
}

// GetOverrideLocalPassword returns the OverrideLocalPassword field value if set, zero value otherwise.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetOverrideLocalPassword() bool {
	if o == nil || isNil(o.OverrideLocalPassword) {
		var ret bool
		return ret
	}
	return *o.OverrideLocalPassword
}

// GetOverrideLocalPasswordOk returns a tuple with the OverrideLocalPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetOverrideLocalPasswordOk() (*bool, bool) {
	if o == nil || isNil(o.OverrideLocalPassword) {
		return nil, false
	}
	return o.OverrideLocalPassword, true
}

// HasOverrideLocalPassword returns a boolean if a field has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) HasOverrideLocalPassword() bool {
	if o != nil && !isNil(o.OverrideLocalPassword) {
		return true
	}

	return false
}

// SetOverrideLocalPassword gets a reference to the given bool and assigns it to the OverrideLocalPassword field.
func (o *PluggablePassThroughAuthenticationPluginResponse) SetOverrideLocalPassword(v bool) {
	o.OverrideLocalPassword = &v
}

// GetUpdateLocalPassword returns the UpdateLocalPassword field value if set, zero value otherwise.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetUpdateLocalPassword() bool {
	if o == nil || isNil(o.UpdateLocalPassword) {
		var ret bool
		return ret
	}
	return *o.UpdateLocalPassword
}

// GetUpdateLocalPasswordOk returns a tuple with the UpdateLocalPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetUpdateLocalPasswordOk() (*bool, bool) {
	if o == nil || isNil(o.UpdateLocalPassword) {
		return nil, false
	}
	return o.UpdateLocalPassword, true
}

// HasUpdateLocalPassword returns a boolean if a field has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) HasUpdateLocalPassword() bool {
	if o != nil && !isNil(o.UpdateLocalPassword) {
		return true
	}

	return false
}

// SetUpdateLocalPassword gets a reference to the given bool and assigns it to the UpdateLocalPassword field.
func (o *PluggablePassThroughAuthenticationPluginResponse) SetUpdateLocalPassword(v bool) {
	o.UpdateLocalPassword = &v
}

// GetUpdateLocalPasswordDN returns the UpdateLocalPasswordDN field value if set, zero value otherwise.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetUpdateLocalPasswordDN() string {
	if o == nil || isNil(o.UpdateLocalPasswordDN) {
		var ret string
		return ret
	}
	return *o.UpdateLocalPasswordDN
}

// GetUpdateLocalPasswordDNOk returns a tuple with the UpdateLocalPasswordDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetUpdateLocalPasswordDNOk() (*string, bool) {
	if o == nil || isNil(o.UpdateLocalPasswordDN) {
		return nil, false
	}
	return o.UpdateLocalPasswordDN, true
}

// HasUpdateLocalPasswordDN returns a boolean if a field has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) HasUpdateLocalPasswordDN() bool {
	if o != nil && !isNil(o.UpdateLocalPasswordDN) {
		return true
	}

	return false
}

// SetUpdateLocalPasswordDN gets a reference to the given string and assigns it to the UpdateLocalPasswordDN field.
func (o *PluggablePassThroughAuthenticationPluginResponse) SetUpdateLocalPasswordDN(v string) {
	o.UpdateLocalPasswordDN = &v
}

// GetAllowLaxPassThroughAuthenticationPasswords returns the AllowLaxPassThroughAuthenticationPasswords field value if set, zero value otherwise.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetAllowLaxPassThroughAuthenticationPasswords() bool {
	if o == nil || isNil(o.AllowLaxPassThroughAuthenticationPasswords) {
		var ret bool
		return ret
	}
	return *o.AllowLaxPassThroughAuthenticationPasswords
}

// GetAllowLaxPassThroughAuthenticationPasswordsOk returns a tuple with the AllowLaxPassThroughAuthenticationPasswords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetAllowLaxPassThroughAuthenticationPasswordsOk() (*bool, bool) {
	if o == nil || isNil(o.AllowLaxPassThroughAuthenticationPasswords) {
		return nil, false
	}
	return o.AllowLaxPassThroughAuthenticationPasswords, true
}

// HasAllowLaxPassThroughAuthenticationPasswords returns a boolean if a field has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) HasAllowLaxPassThroughAuthenticationPasswords() bool {
	if o != nil && !isNil(o.AllowLaxPassThroughAuthenticationPasswords) {
		return true
	}

	return false
}

// SetAllowLaxPassThroughAuthenticationPasswords gets a reference to the given bool and assigns it to the AllowLaxPassThroughAuthenticationPasswords field.
func (o *PluggablePassThroughAuthenticationPluginResponse) SetAllowLaxPassThroughAuthenticationPasswords(v bool) {
	o.AllowLaxPassThroughAuthenticationPasswords = &v
}

// GetIgnoredPasswordPolicyStateErrorCondition returns the IgnoredPasswordPolicyStateErrorCondition field value if set, zero value otherwise.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetIgnoredPasswordPolicyStateErrorCondition() []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp {
	if o == nil || isNil(o.IgnoredPasswordPolicyStateErrorCondition) {
		var ret []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp
		return ret
	}
	return o.IgnoredPasswordPolicyStateErrorCondition
}

// GetIgnoredPasswordPolicyStateErrorConditionOk returns a tuple with the IgnoredPasswordPolicyStateErrorCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetIgnoredPasswordPolicyStateErrorConditionOk() ([]EnumpluginIgnoredPasswordPolicyStateErrorConditionProp, bool) {
	if o == nil || isNil(o.IgnoredPasswordPolicyStateErrorCondition) {
		return nil, false
	}
	return o.IgnoredPasswordPolicyStateErrorCondition, true
}

// HasIgnoredPasswordPolicyStateErrorCondition returns a boolean if a field has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) HasIgnoredPasswordPolicyStateErrorCondition() bool {
	if o != nil && !isNil(o.IgnoredPasswordPolicyStateErrorCondition) {
		return true
	}

	return false
}

// SetIgnoredPasswordPolicyStateErrorCondition gets a reference to the given []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp and assigns it to the IgnoredPasswordPolicyStateErrorCondition field.
func (o *PluggablePassThroughAuthenticationPluginResponse) SetIgnoredPasswordPolicyStateErrorCondition(v []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp) {
	o.IgnoredPasswordPolicyStateErrorCondition = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PluggablePassThroughAuthenticationPluginResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *PluggablePassThroughAuthenticationPluginResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PluggablePassThroughAuthenticationPluginResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetInvokeForInternalOperations() bool {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *PluggablePassThroughAuthenticationPluginResponse) HasInvokeForInternalOperations() bool {
	if o != nil && !isNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *PluggablePassThroughAuthenticationPluginResponse) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

func (o PluggablePassThroughAuthenticationPluginResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["passThroughAuthenticationHandler"] = o.PassThroughAuthenticationHandler
	}
	if !isNil(o.IncludedLocalEntryBaseDN) {
		toSerialize["includedLocalEntryBaseDN"] = o.IncludedLocalEntryBaseDN
	}
	if !isNil(o.ConnectionCriteria) {
		toSerialize["connectionCriteria"] = o.ConnectionCriteria
	}
	if !isNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	if !isNil(o.TryLocalBind) {
		toSerialize["tryLocalBind"] = o.TryLocalBind
	}
	if !isNil(o.OverrideLocalPassword) {
		toSerialize["overrideLocalPassword"] = o.OverrideLocalPassword
	}
	if !isNil(o.UpdateLocalPassword) {
		toSerialize["updateLocalPassword"] = o.UpdateLocalPassword
	}
	if !isNil(o.UpdateLocalPasswordDN) {
		toSerialize["updateLocalPasswordDN"] = o.UpdateLocalPasswordDN
	}
	if !isNil(o.AllowLaxPassThroughAuthenticationPasswords) {
		toSerialize["allowLaxPassThroughAuthenticationPasswords"] = o.AllowLaxPassThroughAuthenticationPasswords
	}
	if !isNil(o.IgnoredPasswordPolicyStateErrorCondition) {
		toSerialize["ignoredPasswordPolicyStateErrorCondition"] = o.IgnoredPasswordPolicyStateErrorCondition
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.InvokeForInternalOperations) {
		toSerialize["invokeForInternalOperations"] = o.InvokeForInternalOperations
	}
	return json.Marshal(toSerialize)
}

type NullablePluggablePassThroughAuthenticationPluginResponse struct {
	value *PluggablePassThroughAuthenticationPluginResponse
	isSet bool
}

func (v NullablePluggablePassThroughAuthenticationPluginResponse) Get() *PluggablePassThroughAuthenticationPluginResponse {
	return v.value
}

func (v *NullablePluggablePassThroughAuthenticationPluginResponse) Set(val *PluggablePassThroughAuthenticationPluginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePluggablePassThroughAuthenticationPluginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePluggablePassThroughAuthenticationPluginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluggablePassThroughAuthenticationPluginResponse(val *PluggablePassThroughAuthenticationPluginResponse) *NullablePluggablePassThroughAuthenticationPluginResponse {
	return &NullablePluggablePassThroughAuthenticationPluginResponse{value: val, isSet: true}
}

func (v NullablePluggablePassThroughAuthenticationPluginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluggablePassThroughAuthenticationPluginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}