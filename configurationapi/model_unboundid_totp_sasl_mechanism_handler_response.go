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

// checks if the UnboundidTotpSaslMechanismHandlerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnboundidTotpSaslMechanismHandlerResponse{}

// UnboundidTotpSaslMechanismHandlerResponse struct for UnboundidTotpSaslMechanismHandlerResponse
type UnboundidTotpSaslMechanismHandlerResponse struct {
	Schemas []EnumunboundidTotpSaslMechanismHandlerSchemaUrn `json:"schemas"`
	// Name of the SASL Mechanism Handler
	Id string `json:"id"`
	// The identity mapper that should be used to identify the user(s) targeted in the authentication and/or authorization identities contained in the bind request. This will only be used for \"u:\"-style identities.
	IdentityMapper string `json:"identityMapper"`
	// The name or OID of the attribute that will be used to hold the shared secret key used during TOTP processing.
	SharedSecretAttributeType *string `json:"sharedSecretAttributeType,omitempty"`
	// The duration of the time interval used for TOTP processing.
	TimeIntervalDuration *string `json:"timeIntervalDuration,omitempty"`
	// The number of adjacent time intervals (both before and after the current time) that should be checked when performing authentication.
	AdjacentIntervalsToCheck *int64 `json:"adjacentIntervalsToCheck,omitempty"`
	// Indicates whether to require a static password (as might be held in the userPassword attribute, or whatever password attribute is defined in the password policy governing the user) in addition to the one-time password.
	RequireStaticPassword *bool `json:"requireStaticPassword,omitempty"`
	// Indicates whether to prevent clients from re-using TOTP passwords.
	PreventTOTPReuse *bool `json:"preventTOTPReuse,omitempty"`
	// A description for this SASL Mechanism Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the SASL mechanism handler is enabled for use.
	Enabled                                       bool                                               `json:"enabled"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewUnboundidTotpSaslMechanismHandlerResponse instantiates a new UnboundidTotpSaslMechanismHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnboundidTotpSaslMechanismHandlerResponse(schemas []EnumunboundidTotpSaslMechanismHandlerSchemaUrn, id string, identityMapper string, enabled bool) *UnboundidTotpSaslMechanismHandlerResponse {
	this := UnboundidTotpSaslMechanismHandlerResponse{}
	this.Schemas = schemas
	this.Id = id
	this.IdentityMapper = identityMapper
	this.Enabled = enabled
	return &this
}

// NewUnboundidTotpSaslMechanismHandlerResponseWithDefaults instantiates a new UnboundidTotpSaslMechanismHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnboundidTotpSaslMechanismHandlerResponseWithDefaults() *UnboundidTotpSaslMechanismHandlerResponse {
	this := UnboundidTotpSaslMechanismHandlerResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetSchemas() []EnumunboundidTotpSaslMechanismHandlerSchemaUrn {
	if o == nil {
		var ret []EnumunboundidTotpSaslMechanismHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetSchemasOk() ([]EnumunboundidTotpSaslMechanismHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *UnboundidTotpSaslMechanismHandlerResponse) SetSchemas(v []EnumunboundidTotpSaslMechanismHandlerSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UnboundidTotpSaslMechanismHandlerResponse) SetId(v string) {
	o.Id = v
}

// GetIdentityMapper returns the IdentityMapper field value
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetIdentityMapper() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value
// and a boolean to check if the value has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetIdentityMapperOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityMapper, true
}

// SetIdentityMapper sets field value
func (o *UnboundidTotpSaslMechanismHandlerResponse) SetIdentityMapper(v string) {
	o.IdentityMapper = v
}

// GetSharedSecretAttributeType returns the SharedSecretAttributeType field value if set, zero value otherwise.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetSharedSecretAttributeType() string {
	if o == nil || IsNil(o.SharedSecretAttributeType) {
		var ret string
		return ret
	}
	return *o.SharedSecretAttributeType
}

// GetSharedSecretAttributeTypeOk returns a tuple with the SharedSecretAttributeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetSharedSecretAttributeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SharedSecretAttributeType) {
		return nil, false
	}
	return o.SharedSecretAttributeType, true
}

// HasSharedSecretAttributeType returns a boolean if a field has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) HasSharedSecretAttributeType() bool {
	if o != nil && !IsNil(o.SharedSecretAttributeType) {
		return true
	}

	return false
}

// SetSharedSecretAttributeType gets a reference to the given string and assigns it to the SharedSecretAttributeType field.
func (o *UnboundidTotpSaslMechanismHandlerResponse) SetSharedSecretAttributeType(v string) {
	o.SharedSecretAttributeType = &v
}

// GetTimeIntervalDuration returns the TimeIntervalDuration field value if set, zero value otherwise.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetTimeIntervalDuration() string {
	if o == nil || IsNil(o.TimeIntervalDuration) {
		var ret string
		return ret
	}
	return *o.TimeIntervalDuration
}

// GetTimeIntervalDurationOk returns a tuple with the TimeIntervalDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetTimeIntervalDurationOk() (*string, bool) {
	if o == nil || IsNil(o.TimeIntervalDuration) {
		return nil, false
	}
	return o.TimeIntervalDuration, true
}

// HasTimeIntervalDuration returns a boolean if a field has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) HasTimeIntervalDuration() bool {
	if o != nil && !IsNil(o.TimeIntervalDuration) {
		return true
	}

	return false
}

// SetTimeIntervalDuration gets a reference to the given string and assigns it to the TimeIntervalDuration field.
func (o *UnboundidTotpSaslMechanismHandlerResponse) SetTimeIntervalDuration(v string) {
	o.TimeIntervalDuration = &v
}

// GetAdjacentIntervalsToCheck returns the AdjacentIntervalsToCheck field value if set, zero value otherwise.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetAdjacentIntervalsToCheck() int64 {
	if o == nil || IsNil(o.AdjacentIntervalsToCheck) {
		var ret int64
		return ret
	}
	return *o.AdjacentIntervalsToCheck
}

// GetAdjacentIntervalsToCheckOk returns a tuple with the AdjacentIntervalsToCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetAdjacentIntervalsToCheckOk() (*int64, bool) {
	if o == nil || IsNil(o.AdjacentIntervalsToCheck) {
		return nil, false
	}
	return o.AdjacentIntervalsToCheck, true
}

// HasAdjacentIntervalsToCheck returns a boolean if a field has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) HasAdjacentIntervalsToCheck() bool {
	if o != nil && !IsNil(o.AdjacentIntervalsToCheck) {
		return true
	}

	return false
}

// SetAdjacentIntervalsToCheck gets a reference to the given int64 and assigns it to the AdjacentIntervalsToCheck field.
func (o *UnboundidTotpSaslMechanismHandlerResponse) SetAdjacentIntervalsToCheck(v int64) {
	o.AdjacentIntervalsToCheck = &v
}

// GetRequireStaticPassword returns the RequireStaticPassword field value if set, zero value otherwise.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetRequireStaticPassword() bool {
	if o == nil || IsNil(o.RequireStaticPassword) {
		var ret bool
		return ret
	}
	return *o.RequireStaticPassword
}

// GetRequireStaticPasswordOk returns a tuple with the RequireStaticPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetRequireStaticPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireStaticPassword) {
		return nil, false
	}
	return o.RequireStaticPassword, true
}

// HasRequireStaticPassword returns a boolean if a field has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) HasRequireStaticPassword() bool {
	if o != nil && !IsNil(o.RequireStaticPassword) {
		return true
	}

	return false
}

// SetRequireStaticPassword gets a reference to the given bool and assigns it to the RequireStaticPassword field.
func (o *UnboundidTotpSaslMechanismHandlerResponse) SetRequireStaticPassword(v bool) {
	o.RequireStaticPassword = &v
}

// GetPreventTOTPReuse returns the PreventTOTPReuse field value if set, zero value otherwise.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetPreventTOTPReuse() bool {
	if o == nil || IsNil(o.PreventTOTPReuse) {
		var ret bool
		return ret
	}
	return *o.PreventTOTPReuse
}

// GetPreventTOTPReuseOk returns a tuple with the PreventTOTPReuse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetPreventTOTPReuseOk() (*bool, bool) {
	if o == nil || IsNil(o.PreventTOTPReuse) {
		return nil, false
	}
	return o.PreventTOTPReuse, true
}

// HasPreventTOTPReuse returns a boolean if a field has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) HasPreventTOTPReuse() bool {
	if o != nil && !IsNil(o.PreventTOTPReuse) {
		return true
	}

	return false
}

// SetPreventTOTPReuse gets a reference to the given bool and assigns it to the PreventTOTPReuse field.
func (o *UnboundidTotpSaslMechanismHandlerResponse) SetPreventTOTPReuse(v bool) {
	o.PreventTOTPReuse = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnboundidTotpSaslMechanismHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *UnboundidTotpSaslMechanismHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *UnboundidTotpSaslMechanismHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *UnboundidTotpSaslMechanismHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *UnboundidTotpSaslMechanismHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o UnboundidTotpSaslMechanismHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnboundidTotpSaslMechanismHandlerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	toSerialize["identityMapper"] = o.IdentityMapper
	if !IsNil(o.SharedSecretAttributeType) {
		toSerialize["sharedSecretAttributeType"] = o.SharedSecretAttributeType
	}
	if !IsNil(o.TimeIntervalDuration) {
		toSerialize["timeIntervalDuration"] = o.TimeIntervalDuration
	}
	if !IsNil(o.AdjacentIntervalsToCheck) {
		toSerialize["adjacentIntervalsToCheck"] = o.AdjacentIntervalsToCheck
	}
	if !IsNil(o.RequireStaticPassword) {
		toSerialize["requireStaticPassword"] = o.RequireStaticPassword
	}
	if !IsNil(o.PreventTOTPReuse) {
		toSerialize["preventTOTPReuse"] = o.PreventTOTPReuse
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableUnboundidTotpSaslMechanismHandlerResponse struct {
	value *UnboundidTotpSaslMechanismHandlerResponse
	isSet bool
}

func (v NullableUnboundidTotpSaslMechanismHandlerResponse) Get() *UnboundidTotpSaslMechanismHandlerResponse {
	return v.value
}

func (v *NullableUnboundidTotpSaslMechanismHandlerResponse) Set(val *UnboundidTotpSaslMechanismHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUnboundidTotpSaslMechanismHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUnboundidTotpSaslMechanismHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnboundidTotpSaslMechanismHandlerResponse(val *UnboundidTotpSaslMechanismHandlerResponse) *NullableUnboundidTotpSaslMechanismHandlerResponse {
	return &NullableUnboundidTotpSaslMechanismHandlerResponse{value: val, isSet: true}
}

func (v NullableUnboundidTotpSaslMechanismHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnboundidTotpSaslMechanismHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
