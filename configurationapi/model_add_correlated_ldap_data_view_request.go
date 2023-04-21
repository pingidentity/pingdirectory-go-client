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

// checks if the AddCorrelatedLdapDataViewRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCorrelatedLdapDataViewRequest{}

// AddCorrelatedLdapDataViewRequest struct for AddCorrelatedLdapDataViewRequest
type AddCorrelatedLdapDataViewRequest struct {
	// Name of the new Correlated LDAP Data View
	ViewName string                                `json:"viewName"`
	Schemas  []EnumcorrelatedLdapDataViewSchemaUrn `json:"schemas,omitempty"`
	// Specifies the LDAP structural object class that should be exposed by this Correlated LDAP Data View.
	StructuralLDAPObjectclass string `json:"structuralLDAPObjectclass"`
	// Specifies an auxiliary LDAP object class that should be exposed by this Correlated LDAP Data View.
	AuxiliaryLDAPObjectclass []string `json:"auxiliaryLDAPObjectclass,omitempty"`
	// Specifies the base DN of the branch of the LDAP directory that can be accessed by this Correlated LDAP Data View.
	IncludeBaseDN string `json:"includeBaseDN"`
	// The set of LDAP filters that define the LDAP entries that should be included in this Correlated LDAP Data View.
	IncludeFilter []string `json:"includeFilter,omitempty"`
	// Specifies the set of operational LDAP attributes to be provided by this Correlated LDAP Data View.
	IncludeOperationalAttribute []string `json:"includeOperationalAttribute,omitempty"`
	// Specifies the template to use for the DN when creating new entries.
	CreateDNPattern *string `json:"createDNPattern,omitempty"`
	// The LDAP attribute from the parent SCIM Resource Type whose value will be used to match objects in the Correlated LDAP Data View. If multiple correlation attributes are required they may be created using additional correlation-attribute-pairs.
	PrimaryCorrelationAttribute string `json:"primaryCorrelationAttribute"`
	// The LDAP attribute from the Correlated LDAP Data View whose value will be matched with the primary-correlation-attribute. If multiple correlation attributes are required they may be specified by creating additional correlation-attribute-pairs.
	SecondaryCorrelationAttribute string `json:"secondaryCorrelationAttribute"`
}

// NewAddCorrelatedLdapDataViewRequest instantiates a new AddCorrelatedLdapDataViewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCorrelatedLdapDataViewRequest(viewName string, structuralLDAPObjectclass string, includeBaseDN string, primaryCorrelationAttribute string, secondaryCorrelationAttribute string) *AddCorrelatedLdapDataViewRequest {
	this := AddCorrelatedLdapDataViewRequest{}
	this.ViewName = viewName
	this.StructuralLDAPObjectclass = structuralLDAPObjectclass
	this.IncludeBaseDN = includeBaseDN
	this.PrimaryCorrelationAttribute = primaryCorrelationAttribute
	this.SecondaryCorrelationAttribute = secondaryCorrelationAttribute
	return &this
}

// NewAddCorrelatedLdapDataViewRequestWithDefaults instantiates a new AddCorrelatedLdapDataViewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCorrelatedLdapDataViewRequestWithDefaults() *AddCorrelatedLdapDataViewRequest {
	this := AddCorrelatedLdapDataViewRequest{}
	return &this
}

// GetViewName returns the ViewName field value
func (o *AddCorrelatedLdapDataViewRequest) GetViewName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value
// and a boolean to check if the value has been set.
func (o *AddCorrelatedLdapDataViewRequest) GetViewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewName, true
}

// SetViewName sets field value
func (o *AddCorrelatedLdapDataViewRequest) SetViewName(v string) {
	o.ViewName = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddCorrelatedLdapDataViewRequest) GetSchemas() []EnumcorrelatedLdapDataViewSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumcorrelatedLdapDataViewSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorrelatedLdapDataViewRequest) GetSchemasOk() ([]EnumcorrelatedLdapDataViewSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddCorrelatedLdapDataViewRequest) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumcorrelatedLdapDataViewSchemaUrn and assigns it to the Schemas field.
func (o *AddCorrelatedLdapDataViewRequest) SetSchemas(v []EnumcorrelatedLdapDataViewSchemaUrn) {
	o.Schemas = v
}

// GetStructuralLDAPObjectclass returns the StructuralLDAPObjectclass field value
func (o *AddCorrelatedLdapDataViewRequest) GetStructuralLDAPObjectclass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StructuralLDAPObjectclass
}

// GetStructuralLDAPObjectclassOk returns a tuple with the StructuralLDAPObjectclass field value
// and a boolean to check if the value has been set.
func (o *AddCorrelatedLdapDataViewRequest) GetStructuralLDAPObjectclassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StructuralLDAPObjectclass, true
}

// SetStructuralLDAPObjectclass sets field value
func (o *AddCorrelatedLdapDataViewRequest) SetStructuralLDAPObjectclass(v string) {
	o.StructuralLDAPObjectclass = v
}

// GetAuxiliaryLDAPObjectclass returns the AuxiliaryLDAPObjectclass field value if set, zero value otherwise.
func (o *AddCorrelatedLdapDataViewRequest) GetAuxiliaryLDAPObjectclass() []string {
	if o == nil || IsNil(o.AuxiliaryLDAPObjectclass) {
		var ret []string
		return ret
	}
	return o.AuxiliaryLDAPObjectclass
}

// GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorrelatedLdapDataViewRequest) GetAuxiliaryLDAPObjectclassOk() ([]string, bool) {
	if o == nil || IsNil(o.AuxiliaryLDAPObjectclass) {
		return nil, false
	}
	return o.AuxiliaryLDAPObjectclass, true
}

// HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.
func (o *AddCorrelatedLdapDataViewRequest) HasAuxiliaryLDAPObjectclass() bool {
	if o != nil && !IsNil(o.AuxiliaryLDAPObjectclass) {
		return true
	}

	return false
}

// SetAuxiliaryLDAPObjectclass gets a reference to the given []string and assigns it to the AuxiliaryLDAPObjectclass field.
func (o *AddCorrelatedLdapDataViewRequest) SetAuxiliaryLDAPObjectclass(v []string) {
	o.AuxiliaryLDAPObjectclass = v
}

// GetIncludeBaseDN returns the IncludeBaseDN field value
func (o *AddCorrelatedLdapDataViewRequest) GetIncludeBaseDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IncludeBaseDN
}

// GetIncludeBaseDNOk returns a tuple with the IncludeBaseDN field value
// and a boolean to check if the value has been set.
func (o *AddCorrelatedLdapDataViewRequest) GetIncludeBaseDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeBaseDN, true
}

// SetIncludeBaseDN sets field value
func (o *AddCorrelatedLdapDataViewRequest) SetIncludeBaseDN(v string) {
	o.IncludeBaseDN = v
}

// GetIncludeFilter returns the IncludeFilter field value if set, zero value otherwise.
func (o *AddCorrelatedLdapDataViewRequest) GetIncludeFilter() []string {
	if o == nil || IsNil(o.IncludeFilter) {
		var ret []string
		return ret
	}
	return o.IncludeFilter
}

// GetIncludeFilterOk returns a tuple with the IncludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorrelatedLdapDataViewRequest) GetIncludeFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeFilter) {
		return nil, false
	}
	return o.IncludeFilter, true
}

// HasIncludeFilter returns a boolean if a field has been set.
func (o *AddCorrelatedLdapDataViewRequest) HasIncludeFilter() bool {
	if o != nil && !IsNil(o.IncludeFilter) {
		return true
	}

	return false
}

// SetIncludeFilter gets a reference to the given []string and assigns it to the IncludeFilter field.
func (o *AddCorrelatedLdapDataViewRequest) SetIncludeFilter(v []string) {
	o.IncludeFilter = v
}

// GetIncludeOperationalAttribute returns the IncludeOperationalAttribute field value if set, zero value otherwise.
func (o *AddCorrelatedLdapDataViewRequest) GetIncludeOperationalAttribute() []string {
	if o == nil || IsNil(o.IncludeOperationalAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeOperationalAttribute
}

// GetIncludeOperationalAttributeOk returns a tuple with the IncludeOperationalAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorrelatedLdapDataViewRequest) GetIncludeOperationalAttributeOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeOperationalAttribute) {
		return nil, false
	}
	return o.IncludeOperationalAttribute, true
}

// HasIncludeOperationalAttribute returns a boolean if a field has been set.
func (o *AddCorrelatedLdapDataViewRequest) HasIncludeOperationalAttribute() bool {
	if o != nil && !IsNil(o.IncludeOperationalAttribute) {
		return true
	}

	return false
}

// SetIncludeOperationalAttribute gets a reference to the given []string and assigns it to the IncludeOperationalAttribute field.
func (o *AddCorrelatedLdapDataViewRequest) SetIncludeOperationalAttribute(v []string) {
	o.IncludeOperationalAttribute = v
}

// GetCreateDNPattern returns the CreateDNPattern field value if set, zero value otherwise.
func (o *AddCorrelatedLdapDataViewRequest) GetCreateDNPattern() string {
	if o == nil || IsNil(o.CreateDNPattern) {
		var ret string
		return ret
	}
	return *o.CreateDNPattern
}

// GetCreateDNPatternOk returns a tuple with the CreateDNPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCorrelatedLdapDataViewRequest) GetCreateDNPatternOk() (*string, bool) {
	if o == nil || IsNil(o.CreateDNPattern) {
		return nil, false
	}
	return o.CreateDNPattern, true
}

// HasCreateDNPattern returns a boolean if a field has been set.
func (o *AddCorrelatedLdapDataViewRequest) HasCreateDNPattern() bool {
	if o != nil && !IsNil(o.CreateDNPattern) {
		return true
	}

	return false
}

// SetCreateDNPattern gets a reference to the given string and assigns it to the CreateDNPattern field.
func (o *AddCorrelatedLdapDataViewRequest) SetCreateDNPattern(v string) {
	o.CreateDNPattern = &v
}

// GetPrimaryCorrelationAttribute returns the PrimaryCorrelationAttribute field value
func (o *AddCorrelatedLdapDataViewRequest) GetPrimaryCorrelationAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryCorrelationAttribute
}

// GetPrimaryCorrelationAttributeOk returns a tuple with the PrimaryCorrelationAttribute field value
// and a boolean to check if the value has been set.
func (o *AddCorrelatedLdapDataViewRequest) GetPrimaryCorrelationAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryCorrelationAttribute, true
}

// SetPrimaryCorrelationAttribute sets field value
func (o *AddCorrelatedLdapDataViewRequest) SetPrimaryCorrelationAttribute(v string) {
	o.PrimaryCorrelationAttribute = v
}

// GetSecondaryCorrelationAttribute returns the SecondaryCorrelationAttribute field value
func (o *AddCorrelatedLdapDataViewRequest) GetSecondaryCorrelationAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecondaryCorrelationAttribute
}

// GetSecondaryCorrelationAttributeOk returns a tuple with the SecondaryCorrelationAttribute field value
// and a boolean to check if the value has been set.
func (o *AddCorrelatedLdapDataViewRequest) GetSecondaryCorrelationAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecondaryCorrelationAttribute, true
}

// SetSecondaryCorrelationAttribute sets field value
func (o *AddCorrelatedLdapDataViewRequest) SetSecondaryCorrelationAttribute(v string) {
	o.SecondaryCorrelationAttribute = v
}

func (o AddCorrelatedLdapDataViewRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCorrelatedLdapDataViewRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["viewName"] = o.ViewName
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	toSerialize["structuralLDAPObjectclass"] = o.StructuralLDAPObjectclass
	if !IsNil(o.AuxiliaryLDAPObjectclass) {
		toSerialize["auxiliaryLDAPObjectclass"] = o.AuxiliaryLDAPObjectclass
	}
	toSerialize["includeBaseDN"] = o.IncludeBaseDN
	if !IsNil(o.IncludeFilter) {
		toSerialize["includeFilter"] = o.IncludeFilter
	}
	if !IsNil(o.IncludeOperationalAttribute) {
		toSerialize["includeOperationalAttribute"] = o.IncludeOperationalAttribute
	}
	if !IsNil(o.CreateDNPattern) {
		toSerialize["createDNPattern"] = o.CreateDNPattern
	}
	toSerialize["primaryCorrelationAttribute"] = o.PrimaryCorrelationAttribute
	toSerialize["secondaryCorrelationAttribute"] = o.SecondaryCorrelationAttribute
	return toSerialize, nil
}

type NullableAddCorrelatedLdapDataViewRequest struct {
	value *AddCorrelatedLdapDataViewRequest
	isSet bool
}

func (v NullableAddCorrelatedLdapDataViewRequest) Get() *AddCorrelatedLdapDataViewRequest {
	return v.value
}

func (v *NullableAddCorrelatedLdapDataViewRequest) Set(val *AddCorrelatedLdapDataViewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCorrelatedLdapDataViewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCorrelatedLdapDataViewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCorrelatedLdapDataViewRequest(val *AddCorrelatedLdapDataViewRequest) *NullableAddCorrelatedLdapDataViewRequest {
	return &NullableAddCorrelatedLdapDataViewRequest{value: val, isSet: true}
}

func (v NullableAddCorrelatedLdapDataViewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCorrelatedLdapDataViewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}