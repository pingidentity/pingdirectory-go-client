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

// checks if the GenericDelegatedAdminAttributeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericDelegatedAdminAttributeResponse{}

// GenericDelegatedAdminAttributeResponse struct for GenericDelegatedAdminAttributeResponse
type GenericDelegatedAdminAttributeResponse struct {
	// Name of the REST Resource Type
	Id      string                                        `json:"id"`
	Schemas []EnumgenericDelegatedAdminAttributeSchemaUrn `json:"schemas"`
	// A description for this Delegated Admin Attribute
	Description *string `json:"description,omitempty"`
	// Specifies the name or OID of the LDAP attribute type.
	AttributeType string `json:"attributeType"`
	// A human readable display name for this Delegated Admin Attribute.
	DisplayName string                                    `json:"displayName"`
	Mutability  EnumdelegatedAdminAttributeMutabilityProp `json:"mutability"`
	// Indicates whether this Delegated Admin Attribute may have multiple values.
	MultiValued bool `json:"multiValued"`
	// Indicates whether this Delegated Admin Attribute is to be included in the summary display for a resource.
	IncludeInSummary bool `json:"includeInSummary"`
	// Specifies which attribute category this attribute belongs to.
	AttributeCategory *string `json:"attributeCategory,omitempty"`
	// This property determines a display order for attributes within a given attribute category. Attributes are ordered within their category based on this index from least to greatest.
	DisplayOrderIndex int64 `json:"displayOrderIndex"`
	// For LDAP attributes with DN syntax, specifies what kind of resource is referenced.
	ReferenceResourceType *string                                               `json:"referenceResourceType,omitempty"`
	AttributePresentation *EnumdelegatedAdminAttributeAttributePresentationProp `json:"attributePresentation,omitempty"`
	// Specifies the format string that is used to present a date and/or time value to the user of the app. This property only applies to LDAP attribute types whose LDAP syntax is GeneralizedTime and is ignored if the attribute type has any other syntax.
	DateTimeFormat                                *string                                            `json:"dateTimeFormat,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewGenericDelegatedAdminAttributeResponse instantiates a new GenericDelegatedAdminAttributeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericDelegatedAdminAttributeResponse(id string, schemas []EnumgenericDelegatedAdminAttributeSchemaUrn, attributeType string, displayName string, mutability EnumdelegatedAdminAttributeMutabilityProp, multiValued bool, includeInSummary bool, displayOrderIndex int64) *GenericDelegatedAdminAttributeResponse {
	this := GenericDelegatedAdminAttributeResponse{}
	this.Id = id
	this.Schemas = schemas
	this.AttributeType = attributeType
	this.DisplayName = displayName
	this.Mutability = mutability
	this.MultiValued = multiValued
	this.IncludeInSummary = includeInSummary
	this.DisplayOrderIndex = displayOrderIndex
	return &this
}

// NewGenericDelegatedAdminAttributeResponseWithDefaults instantiates a new GenericDelegatedAdminAttributeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericDelegatedAdminAttributeResponseWithDefaults() *GenericDelegatedAdminAttributeResponse {
	this := GenericDelegatedAdminAttributeResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GenericDelegatedAdminAttributeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GenericDelegatedAdminAttributeResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GenericDelegatedAdminAttributeResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *GenericDelegatedAdminAttributeResponse) GetSchemas() []EnumgenericDelegatedAdminAttributeSchemaUrn {
	if o == nil {
		var ret []EnumgenericDelegatedAdminAttributeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GenericDelegatedAdminAttributeResponse) GetSchemasOk() ([]EnumgenericDelegatedAdminAttributeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GenericDelegatedAdminAttributeResponse) SetSchemas(v []EnumgenericDelegatedAdminAttributeSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GenericDelegatedAdminAttributeResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericDelegatedAdminAttributeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GenericDelegatedAdminAttributeResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GenericDelegatedAdminAttributeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetAttributeType returns the AttributeType field value
func (o *GenericDelegatedAdminAttributeResponse) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *GenericDelegatedAdminAttributeResponse) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *GenericDelegatedAdminAttributeResponse) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetDisplayName returns the DisplayName field value
func (o *GenericDelegatedAdminAttributeResponse) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *GenericDelegatedAdminAttributeResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *GenericDelegatedAdminAttributeResponse) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetMutability returns the Mutability field value
func (o *GenericDelegatedAdminAttributeResponse) GetMutability() EnumdelegatedAdminAttributeMutabilityProp {
	if o == nil {
		var ret EnumdelegatedAdminAttributeMutabilityProp
		return ret
	}

	return o.Mutability
}

// GetMutabilityOk returns a tuple with the Mutability field value
// and a boolean to check if the value has been set.
func (o *GenericDelegatedAdminAttributeResponse) GetMutabilityOk() (*EnumdelegatedAdminAttributeMutabilityProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mutability, true
}

// SetMutability sets field value
func (o *GenericDelegatedAdminAttributeResponse) SetMutability(v EnumdelegatedAdminAttributeMutabilityProp) {
	o.Mutability = v
}

// GetMultiValued returns the MultiValued field value
func (o *GenericDelegatedAdminAttributeResponse) GetMultiValued() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MultiValued
}

// GetMultiValuedOk returns a tuple with the MultiValued field value
// and a boolean to check if the value has been set.
func (o *GenericDelegatedAdminAttributeResponse) GetMultiValuedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MultiValued, true
}

// SetMultiValued sets field value
func (o *GenericDelegatedAdminAttributeResponse) SetMultiValued(v bool) {
	o.MultiValued = v
}

// GetIncludeInSummary returns the IncludeInSummary field value
func (o *GenericDelegatedAdminAttributeResponse) GetIncludeInSummary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeInSummary
}

// GetIncludeInSummaryOk returns a tuple with the IncludeInSummary field value
// and a boolean to check if the value has been set.
func (o *GenericDelegatedAdminAttributeResponse) GetIncludeInSummaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeInSummary, true
}

// SetIncludeInSummary sets field value
func (o *GenericDelegatedAdminAttributeResponse) SetIncludeInSummary(v bool) {
	o.IncludeInSummary = v
}

// GetAttributeCategory returns the AttributeCategory field value if set, zero value otherwise.
func (o *GenericDelegatedAdminAttributeResponse) GetAttributeCategory() string {
	if o == nil || IsNil(o.AttributeCategory) {
		var ret string
		return ret
	}
	return *o.AttributeCategory
}

// GetAttributeCategoryOk returns a tuple with the AttributeCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericDelegatedAdminAttributeResponse) GetAttributeCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeCategory) {
		return nil, false
	}
	return o.AttributeCategory, true
}

// HasAttributeCategory returns a boolean if a field has been set.
func (o *GenericDelegatedAdminAttributeResponse) HasAttributeCategory() bool {
	if o != nil && !IsNil(o.AttributeCategory) {
		return true
	}

	return false
}

// SetAttributeCategory gets a reference to the given string and assigns it to the AttributeCategory field.
func (o *GenericDelegatedAdminAttributeResponse) SetAttributeCategory(v string) {
	o.AttributeCategory = &v
}

// GetDisplayOrderIndex returns the DisplayOrderIndex field value
func (o *GenericDelegatedAdminAttributeResponse) GetDisplayOrderIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DisplayOrderIndex
}

// GetDisplayOrderIndexOk returns a tuple with the DisplayOrderIndex field value
// and a boolean to check if the value has been set.
func (o *GenericDelegatedAdminAttributeResponse) GetDisplayOrderIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayOrderIndex, true
}

// SetDisplayOrderIndex sets field value
func (o *GenericDelegatedAdminAttributeResponse) SetDisplayOrderIndex(v int64) {
	o.DisplayOrderIndex = v
}

// GetReferenceResourceType returns the ReferenceResourceType field value if set, zero value otherwise.
func (o *GenericDelegatedAdminAttributeResponse) GetReferenceResourceType() string {
	if o == nil || IsNil(o.ReferenceResourceType) {
		var ret string
		return ret
	}
	return *o.ReferenceResourceType
}

// GetReferenceResourceTypeOk returns a tuple with the ReferenceResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericDelegatedAdminAttributeResponse) GetReferenceResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceResourceType) {
		return nil, false
	}
	return o.ReferenceResourceType, true
}

// HasReferenceResourceType returns a boolean if a field has been set.
func (o *GenericDelegatedAdminAttributeResponse) HasReferenceResourceType() bool {
	if o != nil && !IsNil(o.ReferenceResourceType) {
		return true
	}

	return false
}

// SetReferenceResourceType gets a reference to the given string and assigns it to the ReferenceResourceType field.
func (o *GenericDelegatedAdminAttributeResponse) SetReferenceResourceType(v string) {
	o.ReferenceResourceType = &v
}

// GetAttributePresentation returns the AttributePresentation field value if set, zero value otherwise.
func (o *GenericDelegatedAdminAttributeResponse) GetAttributePresentation() EnumdelegatedAdminAttributeAttributePresentationProp {
	if o == nil || IsNil(o.AttributePresentation) {
		var ret EnumdelegatedAdminAttributeAttributePresentationProp
		return ret
	}
	return *o.AttributePresentation
}

// GetAttributePresentationOk returns a tuple with the AttributePresentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericDelegatedAdminAttributeResponse) GetAttributePresentationOk() (*EnumdelegatedAdminAttributeAttributePresentationProp, bool) {
	if o == nil || IsNil(o.AttributePresentation) {
		return nil, false
	}
	return o.AttributePresentation, true
}

// HasAttributePresentation returns a boolean if a field has been set.
func (o *GenericDelegatedAdminAttributeResponse) HasAttributePresentation() bool {
	if o != nil && !IsNil(o.AttributePresentation) {
		return true
	}

	return false
}

// SetAttributePresentation gets a reference to the given EnumdelegatedAdminAttributeAttributePresentationProp and assigns it to the AttributePresentation field.
func (o *GenericDelegatedAdminAttributeResponse) SetAttributePresentation(v EnumdelegatedAdminAttributeAttributePresentationProp) {
	o.AttributePresentation = &v
}

// GetDateTimeFormat returns the DateTimeFormat field value if set, zero value otherwise.
func (o *GenericDelegatedAdminAttributeResponse) GetDateTimeFormat() string {
	if o == nil || IsNil(o.DateTimeFormat) {
		var ret string
		return ret
	}
	return *o.DateTimeFormat
}

// GetDateTimeFormatOk returns a tuple with the DateTimeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericDelegatedAdminAttributeResponse) GetDateTimeFormatOk() (*string, bool) {
	if o == nil || IsNil(o.DateTimeFormat) {
		return nil, false
	}
	return o.DateTimeFormat, true
}

// HasDateTimeFormat returns a boolean if a field has been set.
func (o *GenericDelegatedAdminAttributeResponse) HasDateTimeFormat() bool {
	if o != nil && !IsNil(o.DateTimeFormat) {
		return true
	}

	return false
}

// SetDateTimeFormat gets a reference to the given string and assigns it to the DateTimeFormat field.
func (o *GenericDelegatedAdminAttributeResponse) SetDateTimeFormat(v string) {
	o.DateTimeFormat = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GenericDelegatedAdminAttributeResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericDelegatedAdminAttributeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GenericDelegatedAdminAttributeResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *GenericDelegatedAdminAttributeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *GenericDelegatedAdminAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericDelegatedAdminAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *GenericDelegatedAdminAttributeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *GenericDelegatedAdminAttributeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o GenericDelegatedAdminAttributeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericDelegatedAdminAttributeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["attributeType"] = o.AttributeType
	toSerialize["displayName"] = o.DisplayName
	toSerialize["mutability"] = o.Mutability
	toSerialize["multiValued"] = o.MultiValued
	toSerialize["includeInSummary"] = o.IncludeInSummary
	if !IsNil(o.AttributeCategory) {
		toSerialize["attributeCategory"] = o.AttributeCategory
	}
	toSerialize["displayOrderIndex"] = o.DisplayOrderIndex
	if !IsNil(o.ReferenceResourceType) {
		toSerialize["referenceResourceType"] = o.ReferenceResourceType
	}
	if !IsNil(o.AttributePresentation) {
		toSerialize["attributePresentation"] = o.AttributePresentation
	}
	if !IsNil(o.DateTimeFormat) {
		toSerialize["dateTimeFormat"] = o.DateTimeFormat
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableGenericDelegatedAdminAttributeResponse struct {
	value *GenericDelegatedAdminAttributeResponse
	isSet bool
}

func (v NullableGenericDelegatedAdminAttributeResponse) Get() *GenericDelegatedAdminAttributeResponse {
	return v.value
}

func (v *NullableGenericDelegatedAdminAttributeResponse) Set(val *GenericDelegatedAdminAttributeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericDelegatedAdminAttributeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericDelegatedAdminAttributeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericDelegatedAdminAttributeResponse(val *GenericDelegatedAdminAttributeResponse) *NullableGenericDelegatedAdminAttributeResponse {
	return &NullableGenericDelegatedAdminAttributeResponse{value: val, isSet: true}
}

func (v NullableGenericDelegatedAdminAttributeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericDelegatedAdminAttributeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
