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

// AddSimpleSearchEntryCriteriaRequest struct for AddSimpleSearchEntryCriteriaRequest
type AddSimpleSearchEntryCriteriaRequest struct {
	// Name of the new Search Entry Criteria
	CriteriaName string                                   `json:"criteriaName"`
	Schemas      []EnumsimpleSearchEntryCriteriaSchemaUrn `json:"schemas"`
	// Specifies a request criteria object that must match the associated request for entries included in this Simple Search Entry Criteria. of them.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
	// Specifies the OID of a control that must be present in search result entries included in this Simple Search Entry Criteria. If any control OIDs are provided, then the entry must contain all of those controls.
	AllIncludedEntryControl []string `json:"allIncludedEntryControl,omitempty"`
	// Specifies the OID of a control that may be present in search result entries included in this Simple Search Entry Criteria. If any control OIDs are provided, then the entry must contain at least one of those controls.
	AnyIncludedEntryControl []string `json:"anyIncludedEntryControl,omitempty"`
	// Specifies the OID of a control that should not be present in search result entries included in this Simple Search Entry Criteria. If any control OIDs are provided, then the entry must not contain at least one of those controls (that is, it may contain zero or more of those controls, but not all of them).
	NotAllIncludedEntryControl []string `json:"notAllIncludedEntryControl,omitempty"`
	// Specifies the OID of a control that must not be present in search result entries included in this Simple Search Entry Criteria. If any control OIDs are provided, then the entry must not contain any of those controls.
	NoneIncludedEntryControl []string `json:"noneIncludedEntryControl,omitempty"`
	// Specifies a base DN below which entries included in this Simple Search Entry Criteria may exist.
	IncludedEntryBaseDN []string `json:"includedEntryBaseDN,omitempty"`
	// Specifies a base DN below which entries included in this Simple Search Entry Criteria may not exist.
	ExcludedEntryBaseDN []string `json:"excludedEntryBaseDN,omitempty"`
	// Specifies a search filter that must match search result entries included in this Simple Search Entry Criteria. Note that this matching will be performed against the entry that is actually returned to the client and may not reflect the complete entry stored in the server. If any filters are provided, then the returned entry must match all of those filters.
	AllIncludedEntryFilter []string `json:"allIncludedEntryFilter,omitempty"`
	// Specifies a search filter that may match search result entries included in this Simple Search Entry Criteria. Note that this matching will be performed against the entry that is actually returned to the client and may not reflect the complete entry stored in the server. If any filters are provided, then the entry must match at least one of those filters.
	AnyIncludedEntryFilter []string `json:"anyIncludedEntryFilter,omitempty"`
	// Specifies a search filter that should not match search result entries included in this Simple Search Entry Criteria. Note that this matching will be performed against the entry that is actually returned to the client and may not reflect the complete entry stored in the server. If any filters are provided, then the entry must not match at least one of those filters (that is, the entry may match zero or more of those filters, but not of all of them).
	NotAllIncludedEntryFilter []string `json:"notAllIncludedEntryFilter,omitempty"`
	// Specifies a search filter that must not match search result entries included in this Simple Search Entry Criteria. Note that this matching will be performed against the entry that is actually returned to the client and may not reflect the complete entry stored in the server. If any filters are provided, then the entry must not match any of those filters.
	NoneIncludedEntryFilter []string `json:"noneIncludedEntryFilter,omitempty"`
	// Specifies the DN of a group in which the user associated with the entry must be a member to be included in this Simple Search Entry Criteria. If any group DNs are provided, then the entry must be a member of all of them.
	AllIncludedEntryGroupDN []string `json:"allIncludedEntryGroupDN,omitempty"`
	// Specifies the DN of a group in which the user associated with the entry may be a member to be included in this Simple Search Entry Criteria. If any group DNs are provided, then the entry must be a member of at least one of them.
	AnyIncludedEntryGroupDN []string `json:"anyIncludedEntryGroupDN,omitempty"`
	// Specifies the DN of a group in which the user associated with the entry should not be a member to be included in this Simple Search Entry Criteria. If any group DNs are provided, then the entry must not be a member of at least one of them (that is, the entry may be a member of zero or more of the specified groups, but not of all of them).
	NotAllIncludedEntryGroupDN []string `json:"notAllIncludedEntryGroupDN,omitempty"`
	// Specifies the DN of a group in which the user associated with the entry must not be a member to be included in this Simple Search Entry Criteria. If any group DNs are provided, then the entry must not be a member of any of them.
	NoneIncludedEntryGroupDN []string `json:"noneIncludedEntryGroupDN,omitempty"`
	// A description for this Search Entry Criteria
	Description *string `json:"description,omitempty"`
}

// NewAddSimpleSearchEntryCriteriaRequest instantiates a new AddSimpleSearchEntryCriteriaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSimpleSearchEntryCriteriaRequest(criteriaName string, schemas []EnumsimpleSearchEntryCriteriaSchemaUrn) *AddSimpleSearchEntryCriteriaRequest {
	this := AddSimpleSearchEntryCriteriaRequest{}
	this.CriteriaName = criteriaName
	this.Schemas = schemas
	return &this
}

// NewAddSimpleSearchEntryCriteriaRequestWithDefaults instantiates a new AddSimpleSearchEntryCriteriaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSimpleSearchEntryCriteriaRequestWithDefaults() *AddSimpleSearchEntryCriteriaRequest {
	this := AddSimpleSearchEntryCriteriaRequest{}
	return &this
}

// GetCriteriaName returns the CriteriaName field value
func (o *AddSimpleSearchEntryCriteriaRequest) GetCriteriaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CriteriaName
}

// GetCriteriaNameOk returns a tuple with the CriteriaName field value
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetCriteriaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CriteriaName, true
}

// SetCriteriaName sets field value
func (o *AddSimpleSearchEntryCriteriaRequest) SetCriteriaName(v string) {
	o.CriteriaName = v
}

// GetSchemas returns the Schemas field value
func (o *AddSimpleSearchEntryCriteriaRequest) GetSchemas() []EnumsimpleSearchEntryCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumsimpleSearchEntryCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetSchemasOk() ([]EnumsimpleSearchEntryCriteriaSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddSimpleSearchEntryCriteriaRequest) SetSchemas(v []EnumsimpleSearchEntryCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *AddSimpleSearchEntryCriteriaRequest) GetRequestCriteria() string {
	if o == nil || isNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) HasRequestCriteria() bool {
	if o != nil && !isNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *AddSimpleSearchEntryCriteriaRequest) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetAllIncludedEntryControl returns the AllIncludedEntryControl field value if set, zero value otherwise.
func (o *AddSimpleSearchEntryCriteriaRequest) GetAllIncludedEntryControl() []string {
	if o == nil || isNil(o.AllIncludedEntryControl) {
		var ret []string
		return ret
	}
	return o.AllIncludedEntryControl
}

// GetAllIncludedEntryControlOk returns a tuple with the AllIncludedEntryControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetAllIncludedEntryControlOk() ([]string, bool) {
	if o == nil || isNil(o.AllIncludedEntryControl) {
		return nil, false
	}
	return o.AllIncludedEntryControl, true
}

// HasAllIncludedEntryControl returns a boolean if a field has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) HasAllIncludedEntryControl() bool {
	if o != nil && !isNil(o.AllIncludedEntryControl) {
		return true
	}

	return false
}

// SetAllIncludedEntryControl gets a reference to the given []string and assigns it to the AllIncludedEntryControl field.
func (o *AddSimpleSearchEntryCriteriaRequest) SetAllIncludedEntryControl(v []string) {
	o.AllIncludedEntryControl = v
}

// GetAnyIncludedEntryControl returns the AnyIncludedEntryControl field value if set, zero value otherwise.
func (o *AddSimpleSearchEntryCriteriaRequest) GetAnyIncludedEntryControl() []string {
	if o == nil || isNil(o.AnyIncludedEntryControl) {
		var ret []string
		return ret
	}
	return o.AnyIncludedEntryControl
}

// GetAnyIncludedEntryControlOk returns a tuple with the AnyIncludedEntryControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetAnyIncludedEntryControlOk() ([]string, bool) {
	if o == nil || isNil(o.AnyIncludedEntryControl) {
		return nil, false
	}
	return o.AnyIncludedEntryControl, true
}

// HasAnyIncludedEntryControl returns a boolean if a field has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) HasAnyIncludedEntryControl() bool {
	if o != nil && !isNil(o.AnyIncludedEntryControl) {
		return true
	}

	return false
}

// SetAnyIncludedEntryControl gets a reference to the given []string and assigns it to the AnyIncludedEntryControl field.
func (o *AddSimpleSearchEntryCriteriaRequest) SetAnyIncludedEntryControl(v []string) {
	o.AnyIncludedEntryControl = v
}

// GetNotAllIncludedEntryControl returns the NotAllIncludedEntryControl field value if set, zero value otherwise.
func (o *AddSimpleSearchEntryCriteriaRequest) GetNotAllIncludedEntryControl() []string {
	if o == nil || isNil(o.NotAllIncludedEntryControl) {
		var ret []string
		return ret
	}
	return o.NotAllIncludedEntryControl
}

// GetNotAllIncludedEntryControlOk returns a tuple with the NotAllIncludedEntryControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetNotAllIncludedEntryControlOk() ([]string, bool) {
	if o == nil || isNil(o.NotAllIncludedEntryControl) {
		return nil, false
	}
	return o.NotAllIncludedEntryControl, true
}

// HasNotAllIncludedEntryControl returns a boolean if a field has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) HasNotAllIncludedEntryControl() bool {
	if o != nil && !isNil(o.NotAllIncludedEntryControl) {
		return true
	}

	return false
}

// SetNotAllIncludedEntryControl gets a reference to the given []string and assigns it to the NotAllIncludedEntryControl field.
func (o *AddSimpleSearchEntryCriteriaRequest) SetNotAllIncludedEntryControl(v []string) {
	o.NotAllIncludedEntryControl = v
}

// GetNoneIncludedEntryControl returns the NoneIncludedEntryControl field value if set, zero value otherwise.
func (o *AddSimpleSearchEntryCriteriaRequest) GetNoneIncludedEntryControl() []string {
	if o == nil || isNil(o.NoneIncludedEntryControl) {
		var ret []string
		return ret
	}
	return o.NoneIncludedEntryControl
}

// GetNoneIncludedEntryControlOk returns a tuple with the NoneIncludedEntryControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetNoneIncludedEntryControlOk() ([]string, bool) {
	if o == nil || isNil(o.NoneIncludedEntryControl) {
		return nil, false
	}
	return o.NoneIncludedEntryControl, true
}

// HasNoneIncludedEntryControl returns a boolean if a field has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) HasNoneIncludedEntryControl() bool {
	if o != nil && !isNil(o.NoneIncludedEntryControl) {
		return true
	}

	return false
}

// SetNoneIncludedEntryControl gets a reference to the given []string and assigns it to the NoneIncludedEntryControl field.
func (o *AddSimpleSearchEntryCriteriaRequest) SetNoneIncludedEntryControl(v []string) {
	o.NoneIncludedEntryControl = v
}

// GetIncludedEntryBaseDN returns the IncludedEntryBaseDN field value if set, zero value otherwise.
func (o *AddSimpleSearchEntryCriteriaRequest) GetIncludedEntryBaseDN() []string {
	if o == nil || isNil(o.IncludedEntryBaseDN) {
		var ret []string
		return ret
	}
	return o.IncludedEntryBaseDN
}

// GetIncludedEntryBaseDNOk returns a tuple with the IncludedEntryBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetIncludedEntryBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.IncludedEntryBaseDN) {
		return nil, false
	}
	return o.IncludedEntryBaseDN, true
}

// HasIncludedEntryBaseDN returns a boolean if a field has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) HasIncludedEntryBaseDN() bool {
	if o != nil && !isNil(o.IncludedEntryBaseDN) {
		return true
	}

	return false
}

// SetIncludedEntryBaseDN gets a reference to the given []string and assigns it to the IncludedEntryBaseDN field.
func (o *AddSimpleSearchEntryCriteriaRequest) SetIncludedEntryBaseDN(v []string) {
	o.IncludedEntryBaseDN = v
}

// GetExcludedEntryBaseDN returns the ExcludedEntryBaseDN field value if set, zero value otherwise.
func (o *AddSimpleSearchEntryCriteriaRequest) GetExcludedEntryBaseDN() []string {
	if o == nil || isNil(o.ExcludedEntryBaseDN) {
		var ret []string
		return ret
	}
	return o.ExcludedEntryBaseDN
}

// GetExcludedEntryBaseDNOk returns a tuple with the ExcludedEntryBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetExcludedEntryBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludedEntryBaseDN) {
		return nil, false
	}
	return o.ExcludedEntryBaseDN, true
}

// HasExcludedEntryBaseDN returns a boolean if a field has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) HasExcludedEntryBaseDN() bool {
	if o != nil && !isNil(o.ExcludedEntryBaseDN) {
		return true
	}

	return false
}

// SetExcludedEntryBaseDN gets a reference to the given []string and assigns it to the ExcludedEntryBaseDN field.
func (o *AddSimpleSearchEntryCriteriaRequest) SetExcludedEntryBaseDN(v []string) {
	o.ExcludedEntryBaseDN = v
}

// GetAllIncludedEntryFilter returns the AllIncludedEntryFilter field value if set, zero value otherwise.
func (o *AddSimpleSearchEntryCriteriaRequest) GetAllIncludedEntryFilter() []string {
	if o == nil || isNil(o.AllIncludedEntryFilter) {
		var ret []string
		return ret
	}
	return o.AllIncludedEntryFilter
}

// GetAllIncludedEntryFilterOk returns a tuple with the AllIncludedEntryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetAllIncludedEntryFilterOk() ([]string, bool) {
	if o == nil || isNil(o.AllIncludedEntryFilter) {
		return nil, false
	}
	return o.AllIncludedEntryFilter, true
}

// HasAllIncludedEntryFilter returns a boolean if a field has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) HasAllIncludedEntryFilter() bool {
	if o != nil && !isNil(o.AllIncludedEntryFilter) {
		return true
	}

	return false
}

// SetAllIncludedEntryFilter gets a reference to the given []string and assigns it to the AllIncludedEntryFilter field.
func (o *AddSimpleSearchEntryCriteriaRequest) SetAllIncludedEntryFilter(v []string) {
	o.AllIncludedEntryFilter = v
}

// GetAnyIncludedEntryFilter returns the AnyIncludedEntryFilter field value if set, zero value otherwise.
func (o *AddSimpleSearchEntryCriteriaRequest) GetAnyIncludedEntryFilter() []string {
	if o == nil || isNil(o.AnyIncludedEntryFilter) {
		var ret []string
		return ret
	}
	return o.AnyIncludedEntryFilter
}

// GetAnyIncludedEntryFilterOk returns a tuple with the AnyIncludedEntryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetAnyIncludedEntryFilterOk() ([]string, bool) {
	if o == nil || isNil(o.AnyIncludedEntryFilter) {
		return nil, false
	}
	return o.AnyIncludedEntryFilter, true
}

// HasAnyIncludedEntryFilter returns a boolean if a field has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) HasAnyIncludedEntryFilter() bool {
	if o != nil && !isNil(o.AnyIncludedEntryFilter) {
		return true
	}

	return false
}

// SetAnyIncludedEntryFilter gets a reference to the given []string and assigns it to the AnyIncludedEntryFilter field.
func (o *AddSimpleSearchEntryCriteriaRequest) SetAnyIncludedEntryFilter(v []string) {
	o.AnyIncludedEntryFilter = v
}

// GetNotAllIncludedEntryFilter returns the NotAllIncludedEntryFilter field value if set, zero value otherwise.
func (o *AddSimpleSearchEntryCriteriaRequest) GetNotAllIncludedEntryFilter() []string {
	if o == nil || isNil(o.NotAllIncludedEntryFilter) {
		var ret []string
		return ret
	}
	return o.NotAllIncludedEntryFilter
}

// GetNotAllIncludedEntryFilterOk returns a tuple with the NotAllIncludedEntryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetNotAllIncludedEntryFilterOk() ([]string, bool) {
	if o == nil || isNil(o.NotAllIncludedEntryFilter) {
		return nil, false
	}
	return o.NotAllIncludedEntryFilter, true
}

// HasNotAllIncludedEntryFilter returns a boolean if a field has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) HasNotAllIncludedEntryFilter() bool {
	if o != nil && !isNil(o.NotAllIncludedEntryFilter) {
		return true
	}

	return false
}

// SetNotAllIncludedEntryFilter gets a reference to the given []string and assigns it to the NotAllIncludedEntryFilter field.
func (o *AddSimpleSearchEntryCriteriaRequest) SetNotAllIncludedEntryFilter(v []string) {
	o.NotAllIncludedEntryFilter = v
}

// GetNoneIncludedEntryFilter returns the NoneIncludedEntryFilter field value if set, zero value otherwise.
func (o *AddSimpleSearchEntryCriteriaRequest) GetNoneIncludedEntryFilter() []string {
	if o == nil || isNil(o.NoneIncludedEntryFilter) {
		var ret []string
		return ret
	}
	return o.NoneIncludedEntryFilter
}

// GetNoneIncludedEntryFilterOk returns a tuple with the NoneIncludedEntryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetNoneIncludedEntryFilterOk() ([]string, bool) {
	if o == nil || isNil(o.NoneIncludedEntryFilter) {
		return nil, false
	}
	return o.NoneIncludedEntryFilter, true
}

// HasNoneIncludedEntryFilter returns a boolean if a field has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) HasNoneIncludedEntryFilter() bool {
	if o != nil && !isNil(o.NoneIncludedEntryFilter) {
		return true
	}

	return false
}

// SetNoneIncludedEntryFilter gets a reference to the given []string and assigns it to the NoneIncludedEntryFilter field.
func (o *AddSimpleSearchEntryCriteriaRequest) SetNoneIncludedEntryFilter(v []string) {
	o.NoneIncludedEntryFilter = v
}

// GetAllIncludedEntryGroupDN returns the AllIncludedEntryGroupDN field value if set, zero value otherwise.
func (o *AddSimpleSearchEntryCriteriaRequest) GetAllIncludedEntryGroupDN() []string {
	if o == nil || isNil(o.AllIncludedEntryGroupDN) {
		var ret []string
		return ret
	}
	return o.AllIncludedEntryGroupDN
}

// GetAllIncludedEntryGroupDNOk returns a tuple with the AllIncludedEntryGroupDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetAllIncludedEntryGroupDNOk() ([]string, bool) {
	if o == nil || isNil(o.AllIncludedEntryGroupDN) {
		return nil, false
	}
	return o.AllIncludedEntryGroupDN, true
}

// HasAllIncludedEntryGroupDN returns a boolean if a field has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) HasAllIncludedEntryGroupDN() bool {
	if o != nil && !isNil(o.AllIncludedEntryGroupDN) {
		return true
	}

	return false
}

// SetAllIncludedEntryGroupDN gets a reference to the given []string and assigns it to the AllIncludedEntryGroupDN field.
func (o *AddSimpleSearchEntryCriteriaRequest) SetAllIncludedEntryGroupDN(v []string) {
	o.AllIncludedEntryGroupDN = v
}

// GetAnyIncludedEntryGroupDN returns the AnyIncludedEntryGroupDN field value if set, zero value otherwise.
func (o *AddSimpleSearchEntryCriteriaRequest) GetAnyIncludedEntryGroupDN() []string {
	if o == nil || isNil(o.AnyIncludedEntryGroupDN) {
		var ret []string
		return ret
	}
	return o.AnyIncludedEntryGroupDN
}

// GetAnyIncludedEntryGroupDNOk returns a tuple with the AnyIncludedEntryGroupDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetAnyIncludedEntryGroupDNOk() ([]string, bool) {
	if o == nil || isNil(o.AnyIncludedEntryGroupDN) {
		return nil, false
	}
	return o.AnyIncludedEntryGroupDN, true
}

// HasAnyIncludedEntryGroupDN returns a boolean if a field has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) HasAnyIncludedEntryGroupDN() bool {
	if o != nil && !isNil(o.AnyIncludedEntryGroupDN) {
		return true
	}

	return false
}

// SetAnyIncludedEntryGroupDN gets a reference to the given []string and assigns it to the AnyIncludedEntryGroupDN field.
func (o *AddSimpleSearchEntryCriteriaRequest) SetAnyIncludedEntryGroupDN(v []string) {
	o.AnyIncludedEntryGroupDN = v
}

// GetNotAllIncludedEntryGroupDN returns the NotAllIncludedEntryGroupDN field value if set, zero value otherwise.
func (o *AddSimpleSearchEntryCriteriaRequest) GetNotAllIncludedEntryGroupDN() []string {
	if o == nil || isNil(o.NotAllIncludedEntryGroupDN) {
		var ret []string
		return ret
	}
	return o.NotAllIncludedEntryGroupDN
}

// GetNotAllIncludedEntryGroupDNOk returns a tuple with the NotAllIncludedEntryGroupDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetNotAllIncludedEntryGroupDNOk() ([]string, bool) {
	if o == nil || isNil(o.NotAllIncludedEntryGroupDN) {
		return nil, false
	}
	return o.NotAllIncludedEntryGroupDN, true
}

// HasNotAllIncludedEntryGroupDN returns a boolean if a field has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) HasNotAllIncludedEntryGroupDN() bool {
	if o != nil && !isNil(o.NotAllIncludedEntryGroupDN) {
		return true
	}

	return false
}

// SetNotAllIncludedEntryGroupDN gets a reference to the given []string and assigns it to the NotAllIncludedEntryGroupDN field.
func (o *AddSimpleSearchEntryCriteriaRequest) SetNotAllIncludedEntryGroupDN(v []string) {
	o.NotAllIncludedEntryGroupDN = v
}

// GetNoneIncludedEntryGroupDN returns the NoneIncludedEntryGroupDN field value if set, zero value otherwise.
func (o *AddSimpleSearchEntryCriteriaRequest) GetNoneIncludedEntryGroupDN() []string {
	if o == nil || isNil(o.NoneIncludedEntryGroupDN) {
		var ret []string
		return ret
	}
	return o.NoneIncludedEntryGroupDN
}

// GetNoneIncludedEntryGroupDNOk returns a tuple with the NoneIncludedEntryGroupDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetNoneIncludedEntryGroupDNOk() ([]string, bool) {
	if o == nil || isNil(o.NoneIncludedEntryGroupDN) {
		return nil, false
	}
	return o.NoneIncludedEntryGroupDN, true
}

// HasNoneIncludedEntryGroupDN returns a boolean if a field has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) HasNoneIncludedEntryGroupDN() bool {
	if o != nil && !isNil(o.NoneIncludedEntryGroupDN) {
		return true
	}

	return false
}

// SetNoneIncludedEntryGroupDN gets a reference to the given []string and assigns it to the NoneIncludedEntryGroupDN field.
func (o *AddSimpleSearchEntryCriteriaRequest) SetNoneIncludedEntryGroupDN(v []string) {
	o.NoneIncludedEntryGroupDN = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSimpleSearchEntryCriteriaRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSimpleSearchEntryCriteriaRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSimpleSearchEntryCriteriaRequest) SetDescription(v string) {
	o.Description = &v
}

func (o AddSimpleSearchEntryCriteriaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["criteriaName"] = o.CriteriaName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	if !isNil(o.AllIncludedEntryControl) {
		toSerialize["allIncludedEntryControl"] = o.AllIncludedEntryControl
	}
	if !isNil(o.AnyIncludedEntryControl) {
		toSerialize["anyIncludedEntryControl"] = o.AnyIncludedEntryControl
	}
	if !isNil(o.NotAllIncludedEntryControl) {
		toSerialize["notAllIncludedEntryControl"] = o.NotAllIncludedEntryControl
	}
	if !isNil(o.NoneIncludedEntryControl) {
		toSerialize["noneIncludedEntryControl"] = o.NoneIncludedEntryControl
	}
	if !isNil(o.IncludedEntryBaseDN) {
		toSerialize["includedEntryBaseDN"] = o.IncludedEntryBaseDN
	}
	if !isNil(o.ExcludedEntryBaseDN) {
		toSerialize["excludedEntryBaseDN"] = o.ExcludedEntryBaseDN
	}
	if !isNil(o.AllIncludedEntryFilter) {
		toSerialize["allIncludedEntryFilter"] = o.AllIncludedEntryFilter
	}
	if !isNil(o.AnyIncludedEntryFilter) {
		toSerialize["anyIncludedEntryFilter"] = o.AnyIncludedEntryFilter
	}
	if !isNil(o.NotAllIncludedEntryFilter) {
		toSerialize["notAllIncludedEntryFilter"] = o.NotAllIncludedEntryFilter
	}
	if !isNil(o.NoneIncludedEntryFilter) {
		toSerialize["noneIncludedEntryFilter"] = o.NoneIncludedEntryFilter
	}
	if !isNil(o.AllIncludedEntryGroupDN) {
		toSerialize["allIncludedEntryGroupDN"] = o.AllIncludedEntryGroupDN
	}
	if !isNil(o.AnyIncludedEntryGroupDN) {
		toSerialize["anyIncludedEntryGroupDN"] = o.AnyIncludedEntryGroupDN
	}
	if !isNil(o.NotAllIncludedEntryGroupDN) {
		toSerialize["notAllIncludedEntryGroupDN"] = o.NotAllIncludedEntryGroupDN
	}
	if !isNil(o.NoneIncludedEntryGroupDN) {
		toSerialize["noneIncludedEntryGroupDN"] = o.NoneIncludedEntryGroupDN
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAddSimpleSearchEntryCriteriaRequest struct {
	value *AddSimpleSearchEntryCriteriaRequest
	isSet bool
}

func (v NullableAddSimpleSearchEntryCriteriaRequest) Get() *AddSimpleSearchEntryCriteriaRequest {
	return v.value
}

func (v *NullableAddSimpleSearchEntryCriteriaRequest) Set(val *AddSimpleSearchEntryCriteriaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSimpleSearchEntryCriteriaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSimpleSearchEntryCriteriaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSimpleSearchEntryCriteriaRequest(val *AddSimpleSearchEntryCriteriaRequest) *NullableAddSimpleSearchEntryCriteriaRequest {
	return &NullableAddSimpleSearchEntryCriteriaRequest{value: val, isSet: true}
}

func (v NullableAddSimpleSearchEntryCriteriaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSimpleSearchEntryCriteriaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}