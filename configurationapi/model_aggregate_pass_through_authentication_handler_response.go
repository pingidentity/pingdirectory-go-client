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

// checks if the AggregatePassThroughAuthenticationHandlerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggregatePassThroughAuthenticationHandlerResponse{}

// AggregatePassThroughAuthenticationHandlerResponse struct for AggregatePassThroughAuthenticationHandlerResponse
type AggregatePassThroughAuthenticationHandlerResponse struct {
	Schemas []EnumaggregatePassThroughAuthenticationHandlerSchemaUrn `json:"schemas"`
	// The set of subordinate pass-through authentication handlers that may be used to perform the authentication processing. Handlers will be invoked in order until one is found for which the bind operation matches the associated criteria and either succeeds or fails in a manner that should not be ignored.
	SubordinatePassThroughAuthenticationHandler []string                                                        `json:"subordinatePassThroughAuthenticationHandler"`
	ContinueOnFailureType                       []EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp `json:"continueOnFailureType,omitempty"`
	// A description for this Pass Through Authentication Handler
	Description *string `json:"description,omitempty"`
	// The base DNs for the local users whose authentication attempts may be passed through to the external authentication service.
	IncludedLocalEntryBaseDN []string `json:"includedLocalEntryBaseDN,omitempty"`
	// A reference to connection criteria that will be used to indicate which bind requests should be passed through to the external authentication service.
	ConnectionCriteria *string `json:"connectionCriteria,omitempty"`
	// A reference to request criteria that will be used to indicate which bind requests should be passed through to the external authentication service.
	RequestCriteria                               *string                                            `json:"requestCriteria,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Pass Through Authentication Handler
	Id string `json:"id"`
}

// NewAggregatePassThroughAuthenticationHandlerResponse instantiates a new AggregatePassThroughAuthenticationHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregatePassThroughAuthenticationHandlerResponse(schemas []EnumaggregatePassThroughAuthenticationHandlerSchemaUrn, subordinatePassThroughAuthenticationHandler []string, id string) *AggregatePassThroughAuthenticationHandlerResponse {
	this := AggregatePassThroughAuthenticationHandlerResponse{}
	this.Schemas = schemas
	this.SubordinatePassThroughAuthenticationHandler = subordinatePassThroughAuthenticationHandler
	this.Id = id
	return &this
}

// NewAggregatePassThroughAuthenticationHandlerResponseWithDefaults instantiates a new AggregatePassThroughAuthenticationHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregatePassThroughAuthenticationHandlerResponseWithDefaults() *AggregatePassThroughAuthenticationHandlerResponse {
	this := AggregatePassThroughAuthenticationHandlerResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetSchemas() []EnumaggregatePassThroughAuthenticationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumaggregatePassThroughAuthenticationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetSchemasOk() ([]EnumaggregatePassThroughAuthenticationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AggregatePassThroughAuthenticationHandlerResponse) SetSchemas(v []EnumaggregatePassThroughAuthenticationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetSubordinatePassThroughAuthenticationHandler returns the SubordinatePassThroughAuthenticationHandler field value
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetSubordinatePassThroughAuthenticationHandler() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SubordinatePassThroughAuthenticationHandler
}

// GetSubordinatePassThroughAuthenticationHandlerOk returns a tuple with the SubordinatePassThroughAuthenticationHandler field value
// and a boolean to check if the value has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetSubordinatePassThroughAuthenticationHandlerOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubordinatePassThroughAuthenticationHandler, true
}

// SetSubordinatePassThroughAuthenticationHandler sets field value
func (o *AggregatePassThroughAuthenticationHandlerResponse) SetSubordinatePassThroughAuthenticationHandler(v []string) {
	o.SubordinatePassThroughAuthenticationHandler = v
}

// GetContinueOnFailureType returns the ContinueOnFailureType field value if set, zero value otherwise.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetContinueOnFailureType() []EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp {
	if o == nil || IsNil(o.ContinueOnFailureType) {
		var ret []EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp
		return ret
	}
	return o.ContinueOnFailureType
}

// GetContinueOnFailureTypeOk returns a tuple with the ContinueOnFailureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetContinueOnFailureTypeOk() ([]EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp, bool) {
	if o == nil || IsNil(o.ContinueOnFailureType) {
		return nil, false
	}
	return o.ContinueOnFailureType, true
}

// HasContinueOnFailureType returns a boolean if a field has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) HasContinueOnFailureType() bool {
	if o != nil && !IsNil(o.ContinueOnFailureType) {
		return true
	}

	return false
}

// SetContinueOnFailureType gets a reference to the given []EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp and assigns it to the ContinueOnFailureType field.
func (o *AggregatePassThroughAuthenticationHandlerResponse) SetContinueOnFailureType(v []EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp) {
	o.ContinueOnFailureType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AggregatePassThroughAuthenticationHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field value if set, zero value otherwise.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetIncludedLocalEntryBaseDN() []string {
	if o == nil || IsNil(o.IncludedLocalEntryBaseDN) {
		var ret []string
		return ret
	}
	return o.IncludedLocalEntryBaseDN
}

// GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetIncludedLocalEntryBaseDNOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedLocalEntryBaseDN) {
		return nil, false
	}
	return o.IncludedLocalEntryBaseDN, true
}

// HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) HasIncludedLocalEntryBaseDN() bool {
	if o != nil && !IsNil(o.IncludedLocalEntryBaseDN) {
		return true
	}

	return false
}

// SetIncludedLocalEntryBaseDN gets a reference to the given []string and assigns it to the IncludedLocalEntryBaseDN field.
func (o *AggregatePassThroughAuthenticationHandlerResponse) SetIncludedLocalEntryBaseDN(v []string) {
	o.IncludedLocalEntryBaseDN = v
}

// GetConnectionCriteria returns the ConnectionCriteria field value if set, zero value otherwise.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetConnectionCriteria() string {
	if o == nil || IsNil(o.ConnectionCriteria) {
		var ret string
		return ret
	}
	return *o.ConnectionCriteria
}

// GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetConnectionCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionCriteria) {
		return nil, false
	}
	return o.ConnectionCriteria, true
}

// HasConnectionCriteria returns a boolean if a field has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) HasConnectionCriteria() bool {
	if o != nil && !IsNil(o.ConnectionCriteria) {
		return true
	}

	return false
}

// SetConnectionCriteria gets a reference to the given string and assigns it to the ConnectionCriteria field.
func (o *AggregatePassThroughAuthenticationHandlerResponse) SetConnectionCriteria(v string) {
	o.ConnectionCriteria = &v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetRequestCriteria() string {
	if o == nil || IsNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) HasRequestCriteria() bool {
	if o != nil && !IsNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *AggregatePassThroughAuthenticationHandlerResponse) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AggregatePassThroughAuthenticationHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *AggregatePassThroughAuthenticationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AggregatePassThroughAuthenticationHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AggregatePassThroughAuthenticationHandlerResponse) SetId(v string) {
	o.Id = v
}

func (o AggregatePassThroughAuthenticationHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregatePassThroughAuthenticationHandlerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["subordinatePassThroughAuthenticationHandler"] = o.SubordinatePassThroughAuthenticationHandler
	if !IsNil(o.ContinueOnFailureType) {
		toSerialize["continueOnFailureType"] = o.ContinueOnFailureType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IncludedLocalEntryBaseDN) {
		toSerialize["includedLocalEntryBaseDN"] = o.IncludedLocalEntryBaseDN
	}
	if !IsNil(o.ConnectionCriteria) {
		toSerialize["connectionCriteria"] = o.ConnectionCriteria
	}
	if !IsNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
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

type NullableAggregatePassThroughAuthenticationHandlerResponse struct {
	value *AggregatePassThroughAuthenticationHandlerResponse
	isSet bool
}

func (v NullableAggregatePassThroughAuthenticationHandlerResponse) Get() *AggregatePassThroughAuthenticationHandlerResponse {
	return v.value
}

func (v *NullableAggregatePassThroughAuthenticationHandlerResponse) Set(val *AggregatePassThroughAuthenticationHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregatePassThroughAuthenticationHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregatePassThroughAuthenticationHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregatePassThroughAuthenticationHandlerResponse(val *AggregatePassThroughAuthenticationHandlerResponse) *NullableAggregatePassThroughAuthenticationHandlerResponse {
	return &NullableAggregatePassThroughAuthenticationHandlerResponse{value: val, isSet: true}
}

func (v NullableAggregatePassThroughAuthenticationHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregatePassThroughAuthenticationHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
