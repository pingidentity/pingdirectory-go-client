/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UserRestResourceTypeResponse struct for UserRestResourceTypeResponse
type UserRestResourceTypeResponse struct {
	// Name of the REST Resource Type
	Id string `json:"id"`
	Schemas []EnumuserRestResourceTypeSchemaUrn `json:"schemas"`
	// Specifies which attribute category the password belongs to.
	PasswordAttributeCategory *string `json:"passwordAttributeCategory,omitempty"`
	// This property determines the display order for the password within its attribute category. Attributes are ordered within their category based on this index from least to greatest.
	PasswordDisplayOrderIndex *int32 `json:"passwordDisplayOrderIndex,omitempty"`
	// A description for this REST Resource Type
	Description *string `json:"description,omitempty"`
	// Indicates whether the REST Resource Type is enabled.
	Enabled bool `json:"enabled"`
	// The HTTP addressable endpoint of this REST Resource Type relative to a REST API base URL. Do not include a leading '/'.
	ResourceEndpoint string `json:"resourceEndpoint"`
	// Specifies the LDAP structural object class that should be exposed by this REST Resource Type.
	StructuralLDAPObjectclass string `json:"structuralLDAPObjectclass"`
	// Specifies an auxiliary LDAP object class that should be exposed by this REST Resource Type.
	AuxiliaryLDAPObjectclass []string `json:"auxiliaryLDAPObjectclass,omitempty"`
	// Specifies the base DN of the branch of the LDAP directory where resources of this type are located.
	SearchBaseDN string `json:"searchBaseDN"`
	// The set of LDAP filters that define the LDAP entries that should be included in this REST Resource Type.
	IncludeFilter []string `json:"includeFilter,omitempty"`
	// Specifies the DN of the parent entry for new resources of this type, when a parent resource is not provided by the app. The parent DN must be at or below the search base of this resource type.
	ParentDN *string `json:"parentDN,omitempty"`
	// Specifies the name of another resource type which may be a parent of new resources of this type. The search base DN of the parent resource type must be at or above the search base DN of this resource type.
	ParentResourceType *string `json:"parentResourceType,omitempty"`
	// Specifies a template for a relative DN from the parent resource which identifies the parent entry for a new resource of this type. If this property is not specified then new resources are created immediately below the parent resource or parent DN.
	RelativeDNFromParentResource *string `json:"relativeDNFromParentResource,omitempty"`
	// Specifies the name or OID of the LDAP attribute type to be used as the RDN of new resources.
	CreateRDNAttributeType *string `json:"createRDNAttributeType,omitempty"`
	// Specifies an attribute whose values are to be constructed when a new resource is created. The values are only set at creation time. Subsequent modifications to attributes in the constructed attribute value-pattern are not propagated here.
	PostCreateConstructedAttribute []string `json:"postCreateConstructedAttribute,omitempty"`
	// Specifies an attribute whose values are to be constructed when a resource is updated. The constructed values replace any existing values of the attribute.
	UpdateConstructedAttribute []string `json:"updateConstructedAttribute,omitempty"`
	// A human readable display name for this REST Resource Type.
	DisplayName *string `json:"displayName,omitempty"`
	// Specifies the LDAP filter that should be used when searching for resources matching provided search text. All attribute types in the filter pattern referencing the search text must have a Delegated Admin Attribute definition.
	SearchFilterPattern *string `json:"searchFilterPattern,omitempty"`
	// Specifies the name or OID of the LDAP attribute type which is the primary display attribute. This attribute type must be in the search filter pattern and must have a Delegated Admin Attribute definition.
	PrimaryDisplayAttributeType *string `json:"primaryDisplayAttributeType,omitempty"`
	// The maximum number of resources that may be returned from a search request.
	DelegatedAdminSearchSizeLimit *int32 `json:"delegatedAdminSearchSizeLimit,omitempty"`
	// The maximum number of resources that may be included in a report.
	DelegatedAdminReportSizeLimit *int32 `json:"delegatedAdminReportSizeLimit,omitempty"`
	// Specifies the name of the group member column that will be displayed in the Delegated Admin UI
	MembersColumnName *string `json:"membersColumnName,omitempty"`
	// Specifies the name of the group nonmember column that will be displayed in the Delegated Admin UI
	NonmembersColumnName *string `json:"nonmembersColumnName,omitempty"`
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewUserRestResourceTypeResponse instantiates a new UserRestResourceTypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRestResourceTypeResponse(id string, schemas []EnumuserRestResourceTypeSchemaUrn, enabled bool, resourceEndpoint string, structuralLDAPObjectclass string, searchBaseDN string) *UserRestResourceTypeResponse {
	this := UserRestResourceTypeResponse{}
	this.Id = id
	this.Schemas = schemas
	this.Enabled = enabled
	this.ResourceEndpoint = resourceEndpoint
	this.StructuralLDAPObjectclass = structuralLDAPObjectclass
	this.SearchBaseDN = searchBaseDN
	return &this
}

// NewUserRestResourceTypeResponseWithDefaults instantiates a new UserRestResourceTypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRestResourceTypeResponseWithDefaults() *UserRestResourceTypeResponse {
	this := UserRestResourceTypeResponse{}
	return &this
}

// GetId returns the Id field value
func (o *UserRestResourceTypeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserRestResourceTypeResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *UserRestResourceTypeResponse) GetSchemas() []EnumuserRestResourceTypeSchemaUrn {
	if o == nil {
		var ret []EnumuserRestResourceTypeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetSchemasOk() ([]EnumuserRestResourceTypeSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *UserRestResourceTypeResponse) SetSchemas(v []EnumuserRestResourceTypeSchemaUrn) {
	o.Schemas = v
}

// GetPasswordAttributeCategory returns the PasswordAttributeCategory field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetPasswordAttributeCategory() string {
	if o == nil || isNil(o.PasswordAttributeCategory) {
		var ret string
		return ret
	}
	return *o.PasswordAttributeCategory
}

// GetPasswordAttributeCategoryOk returns a tuple with the PasswordAttributeCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetPasswordAttributeCategoryOk() (*string, bool) {
	if o == nil || isNil(o.PasswordAttributeCategory) {
    return nil, false
	}
	return o.PasswordAttributeCategory, true
}

// HasPasswordAttributeCategory returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasPasswordAttributeCategory() bool {
	if o != nil && !isNil(o.PasswordAttributeCategory) {
		return true
	}

	return false
}

// SetPasswordAttributeCategory gets a reference to the given string and assigns it to the PasswordAttributeCategory field.
func (o *UserRestResourceTypeResponse) SetPasswordAttributeCategory(v string) {
	o.PasswordAttributeCategory = &v
}

// GetPasswordDisplayOrderIndex returns the PasswordDisplayOrderIndex field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetPasswordDisplayOrderIndex() int32 {
	if o == nil || isNil(o.PasswordDisplayOrderIndex) {
		var ret int32
		return ret
	}
	return *o.PasswordDisplayOrderIndex
}

// GetPasswordDisplayOrderIndexOk returns a tuple with the PasswordDisplayOrderIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetPasswordDisplayOrderIndexOk() (*int32, bool) {
	if o == nil || isNil(o.PasswordDisplayOrderIndex) {
    return nil, false
	}
	return o.PasswordDisplayOrderIndex, true
}

// HasPasswordDisplayOrderIndex returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasPasswordDisplayOrderIndex() bool {
	if o != nil && !isNil(o.PasswordDisplayOrderIndex) {
		return true
	}

	return false
}

// SetPasswordDisplayOrderIndex gets a reference to the given int32 and assigns it to the PasswordDisplayOrderIndex field.
func (o *UserRestResourceTypeResponse) SetPasswordDisplayOrderIndex(v int32) {
	o.PasswordDisplayOrderIndex = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UserRestResourceTypeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *UserRestResourceTypeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *UserRestResourceTypeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetResourceEndpoint returns the ResourceEndpoint field value
func (o *UserRestResourceTypeResponse) GetResourceEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceEndpoint
}

// GetResourceEndpointOk returns a tuple with the ResourceEndpoint field value
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetResourceEndpointOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ResourceEndpoint, true
}

// SetResourceEndpoint sets field value
func (o *UserRestResourceTypeResponse) SetResourceEndpoint(v string) {
	o.ResourceEndpoint = v
}

// GetStructuralLDAPObjectclass returns the StructuralLDAPObjectclass field value
func (o *UserRestResourceTypeResponse) GetStructuralLDAPObjectclass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StructuralLDAPObjectclass
}

// GetStructuralLDAPObjectclassOk returns a tuple with the StructuralLDAPObjectclass field value
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetStructuralLDAPObjectclassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StructuralLDAPObjectclass, true
}

// SetStructuralLDAPObjectclass sets field value
func (o *UserRestResourceTypeResponse) SetStructuralLDAPObjectclass(v string) {
	o.StructuralLDAPObjectclass = v
}

// GetAuxiliaryLDAPObjectclass returns the AuxiliaryLDAPObjectclass field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetAuxiliaryLDAPObjectclass() []string {
	if o == nil || isNil(o.AuxiliaryLDAPObjectclass) {
		var ret []string
		return ret
	}
	return o.AuxiliaryLDAPObjectclass
}

// GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetAuxiliaryLDAPObjectclassOk() ([]string, bool) {
	if o == nil || isNil(o.AuxiliaryLDAPObjectclass) {
    return nil, false
	}
	return o.AuxiliaryLDAPObjectclass, true
}

// HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasAuxiliaryLDAPObjectclass() bool {
	if o != nil && !isNil(o.AuxiliaryLDAPObjectclass) {
		return true
	}

	return false
}

// SetAuxiliaryLDAPObjectclass gets a reference to the given []string and assigns it to the AuxiliaryLDAPObjectclass field.
func (o *UserRestResourceTypeResponse) SetAuxiliaryLDAPObjectclass(v []string) {
	o.AuxiliaryLDAPObjectclass = v
}

// GetSearchBaseDN returns the SearchBaseDN field value
func (o *UserRestResourceTypeResponse) GetSearchBaseDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SearchBaseDN
}

// GetSearchBaseDNOk returns a tuple with the SearchBaseDN field value
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetSearchBaseDNOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SearchBaseDN, true
}

// SetSearchBaseDN sets field value
func (o *UserRestResourceTypeResponse) SetSearchBaseDN(v string) {
	o.SearchBaseDN = v
}

// GetIncludeFilter returns the IncludeFilter field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetIncludeFilter() []string {
	if o == nil || isNil(o.IncludeFilter) {
		var ret []string
		return ret
	}
	return o.IncludeFilter
}

// GetIncludeFilterOk returns a tuple with the IncludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetIncludeFilterOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeFilter) {
    return nil, false
	}
	return o.IncludeFilter, true
}

// HasIncludeFilter returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasIncludeFilter() bool {
	if o != nil && !isNil(o.IncludeFilter) {
		return true
	}

	return false
}

// SetIncludeFilter gets a reference to the given []string and assigns it to the IncludeFilter field.
func (o *UserRestResourceTypeResponse) SetIncludeFilter(v []string) {
	o.IncludeFilter = v
}

// GetParentDN returns the ParentDN field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetParentDN() string {
	if o == nil || isNil(o.ParentDN) {
		var ret string
		return ret
	}
	return *o.ParentDN
}

// GetParentDNOk returns a tuple with the ParentDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetParentDNOk() (*string, bool) {
	if o == nil || isNil(o.ParentDN) {
    return nil, false
	}
	return o.ParentDN, true
}

// HasParentDN returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasParentDN() bool {
	if o != nil && !isNil(o.ParentDN) {
		return true
	}

	return false
}

// SetParentDN gets a reference to the given string and assigns it to the ParentDN field.
func (o *UserRestResourceTypeResponse) SetParentDN(v string) {
	o.ParentDN = &v
}

// GetParentResourceType returns the ParentResourceType field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetParentResourceType() string {
	if o == nil || isNil(o.ParentResourceType) {
		var ret string
		return ret
	}
	return *o.ParentResourceType
}

// GetParentResourceTypeOk returns a tuple with the ParentResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetParentResourceTypeOk() (*string, bool) {
	if o == nil || isNil(o.ParentResourceType) {
    return nil, false
	}
	return o.ParentResourceType, true
}

// HasParentResourceType returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasParentResourceType() bool {
	if o != nil && !isNil(o.ParentResourceType) {
		return true
	}

	return false
}

// SetParentResourceType gets a reference to the given string and assigns it to the ParentResourceType field.
func (o *UserRestResourceTypeResponse) SetParentResourceType(v string) {
	o.ParentResourceType = &v
}

// GetRelativeDNFromParentResource returns the RelativeDNFromParentResource field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetRelativeDNFromParentResource() string {
	if o == nil || isNil(o.RelativeDNFromParentResource) {
		var ret string
		return ret
	}
	return *o.RelativeDNFromParentResource
}

// GetRelativeDNFromParentResourceOk returns a tuple with the RelativeDNFromParentResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetRelativeDNFromParentResourceOk() (*string, bool) {
	if o == nil || isNil(o.RelativeDNFromParentResource) {
    return nil, false
	}
	return o.RelativeDNFromParentResource, true
}

// HasRelativeDNFromParentResource returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasRelativeDNFromParentResource() bool {
	if o != nil && !isNil(o.RelativeDNFromParentResource) {
		return true
	}

	return false
}

// SetRelativeDNFromParentResource gets a reference to the given string and assigns it to the RelativeDNFromParentResource field.
func (o *UserRestResourceTypeResponse) SetRelativeDNFromParentResource(v string) {
	o.RelativeDNFromParentResource = &v
}

// GetCreateRDNAttributeType returns the CreateRDNAttributeType field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetCreateRDNAttributeType() string {
	if o == nil || isNil(o.CreateRDNAttributeType) {
		var ret string
		return ret
	}
	return *o.CreateRDNAttributeType
}

// GetCreateRDNAttributeTypeOk returns a tuple with the CreateRDNAttributeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetCreateRDNAttributeTypeOk() (*string, bool) {
	if o == nil || isNil(o.CreateRDNAttributeType) {
    return nil, false
	}
	return o.CreateRDNAttributeType, true
}

// HasCreateRDNAttributeType returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasCreateRDNAttributeType() bool {
	if o != nil && !isNil(o.CreateRDNAttributeType) {
		return true
	}

	return false
}

// SetCreateRDNAttributeType gets a reference to the given string and assigns it to the CreateRDNAttributeType field.
func (o *UserRestResourceTypeResponse) SetCreateRDNAttributeType(v string) {
	o.CreateRDNAttributeType = &v
}

// GetPostCreateConstructedAttribute returns the PostCreateConstructedAttribute field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetPostCreateConstructedAttribute() []string {
	if o == nil || isNil(o.PostCreateConstructedAttribute) {
		var ret []string
		return ret
	}
	return o.PostCreateConstructedAttribute
}

// GetPostCreateConstructedAttributeOk returns a tuple with the PostCreateConstructedAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetPostCreateConstructedAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.PostCreateConstructedAttribute) {
    return nil, false
	}
	return o.PostCreateConstructedAttribute, true
}

// HasPostCreateConstructedAttribute returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasPostCreateConstructedAttribute() bool {
	if o != nil && !isNil(o.PostCreateConstructedAttribute) {
		return true
	}

	return false
}

// SetPostCreateConstructedAttribute gets a reference to the given []string and assigns it to the PostCreateConstructedAttribute field.
func (o *UserRestResourceTypeResponse) SetPostCreateConstructedAttribute(v []string) {
	o.PostCreateConstructedAttribute = v
}

// GetUpdateConstructedAttribute returns the UpdateConstructedAttribute field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetUpdateConstructedAttribute() []string {
	if o == nil || isNil(o.UpdateConstructedAttribute) {
		var ret []string
		return ret
	}
	return o.UpdateConstructedAttribute
}

// GetUpdateConstructedAttributeOk returns a tuple with the UpdateConstructedAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetUpdateConstructedAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.UpdateConstructedAttribute) {
    return nil, false
	}
	return o.UpdateConstructedAttribute, true
}

// HasUpdateConstructedAttribute returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasUpdateConstructedAttribute() bool {
	if o != nil && !isNil(o.UpdateConstructedAttribute) {
		return true
	}

	return false
}

// SetUpdateConstructedAttribute gets a reference to the given []string and assigns it to the UpdateConstructedAttribute field.
func (o *UserRestResourceTypeResponse) SetUpdateConstructedAttribute(v []string) {
	o.UpdateConstructedAttribute = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UserRestResourceTypeResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSearchFilterPattern returns the SearchFilterPattern field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetSearchFilterPattern() string {
	if o == nil || isNil(o.SearchFilterPattern) {
		var ret string
		return ret
	}
	return *o.SearchFilterPattern
}

// GetSearchFilterPatternOk returns a tuple with the SearchFilterPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetSearchFilterPatternOk() (*string, bool) {
	if o == nil || isNil(o.SearchFilterPattern) {
    return nil, false
	}
	return o.SearchFilterPattern, true
}

// HasSearchFilterPattern returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasSearchFilterPattern() bool {
	if o != nil && !isNil(o.SearchFilterPattern) {
		return true
	}

	return false
}

// SetSearchFilterPattern gets a reference to the given string and assigns it to the SearchFilterPattern field.
func (o *UserRestResourceTypeResponse) SetSearchFilterPattern(v string) {
	o.SearchFilterPattern = &v
}

// GetPrimaryDisplayAttributeType returns the PrimaryDisplayAttributeType field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetPrimaryDisplayAttributeType() string {
	if o == nil || isNil(o.PrimaryDisplayAttributeType) {
		var ret string
		return ret
	}
	return *o.PrimaryDisplayAttributeType
}

// GetPrimaryDisplayAttributeTypeOk returns a tuple with the PrimaryDisplayAttributeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetPrimaryDisplayAttributeTypeOk() (*string, bool) {
	if o == nil || isNil(o.PrimaryDisplayAttributeType) {
    return nil, false
	}
	return o.PrimaryDisplayAttributeType, true
}

// HasPrimaryDisplayAttributeType returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasPrimaryDisplayAttributeType() bool {
	if o != nil && !isNil(o.PrimaryDisplayAttributeType) {
		return true
	}

	return false
}

// SetPrimaryDisplayAttributeType gets a reference to the given string and assigns it to the PrimaryDisplayAttributeType field.
func (o *UserRestResourceTypeResponse) SetPrimaryDisplayAttributeType(v string) {
	o.PrimaryDisplayAttributeType = &v
}

// GetDelegatedAdminSearchSizeLimit returns the DelegatedAdminSearchSizeLimit field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetDelegatedAdminSearchSizeLimit() int32 {
	if o == nil || isNil(o.DelegatedAdminSearchSizeLimit) {
		var ret int32
		return ret
	}
	return *o.DelegatedAdminSearchSizeLimit
}

// GetDelegatedAdminSearchSizeLimitOk returns a tuple with the DelegatedAdminSearchSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetDelegatedAdminSearchSizeLimitOk() (*int32, bool) {
	if o == nil || isNil(o.DelegatedAdminSearchSizeLimit) {
    return nil, false
	}
	return o.DelegatedAdminSearchSizeLimit, true
}

// HasDelegatedAdminSearchSizeLimit returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasDelegatedAdminSearchSizeLimit() bool {
	if o != nil && !isNil(o.DelegatedAdminSearchSizeLimit) {
		return true
	}

	return false
}

// SetDelegatedAdminSearchSizeLimit gets a reference to the given int32 and assigns it to the DelegatedAdminSearchSizeLimit field.
func (o *UserRestResourceTypeResponse) SetDelegatedAdminSearchSizeLimit(v int32) {
	o.DelegatedAdminSearchSizeLimit = &v
}

// GetDelegatedAdminReportSizeLimit returns the DelegatedAdminReportSizeLimit field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetDelegatedAdminReportSizeLimit() int32 {
	if o == nil || isNil(o.DelegatedAdminReportSizeLimit) {
		var ret int32
		return ret
	}
	return *o.DelegatedAdminReportSizeLimit
}

// GetDelegatedAdminReportSizeLimitOk returns a tuple with the DelegatedAdminReportSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetDelegatedAdminReportSizeLimitOk() (*int32, bool) {
	if o == nil || isNil(o.DelegatedAdminReportSizeLimit) {
    return nil, false
	}
	return o.DelegatedAdminReportSizeLimit, true
}

// HasDelegatedAdminReportSizeLimit returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasDelegatedAdminReportSizeLimit() bool {
	if o != nil && !isNil(o.DelegatedAdminReportSizeLimit) {
		return true
	}

	return false
}

// SetDelegatedAdminReportSizeLimit gets a reference to the given int32 and assigns it to the DelegatedAdminReportSizeLimit field.
func (o *UserRestResourceTypeResponse) SetDelegatedAdminReportSizeLimit(v int32) {
	o.DelegatedAdminReportSizeLimit = &v
}

// GetMembersColumnName returns the MembersColumnName field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetMembersColumnName() string {
	if o == nil || isNil(o.MembersColumnName) {
		var ret string
		return ret
	}
	return *o.MembersColumnName
}

// GetMembersColumnNameOk returns a tuple with the MembersColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetMembersColumnNameOk() (*string, bool) {
	if o == nil || isNil(o.MembersColumnName) {
    return nil, false
	}
	return o.MembersColumnName, true
}

// HasMembersColumnName returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasMembersColumnName() bool {
	if o != nil && !isNil(o.MembersColumnName) {
		return true
	}

	return false
}

// SetMembersColumnName gets a reference to the given string and assigns it to the MembersColumnName field.
func (o *UserRestResourceTypeResponse) SetMembersColumnName(v string) {
	o.MembersColumnName = &v
}

// GetNonmembersColumnName returns the NonmembersColumnName field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetNonmembersColumnName() string {
	if o == nil || isNil(o.NonmembersColumnName) {
		var ret string
		return ret
	}
	return *o.NonmembersColumnName
}

// GetNonmembersColumnNameOk returns a tuple with the NonmembersColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetNonmembersColumnNameOk() (*string, bool) {
	if o == nil || isNil(o.NonmembersColumnName) {
    return nil, false
	}
	return o.NonmembersColumnName, true
}

// HasNonmembersColumnName returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasNonmembersColumnName() bool {
	if o != nil && !isNil(o.NonmembersColumnName) {
		return true
	}

	return false
}

// SetNonmembersColumnName gets a reference to the given string and assigns it to the NonmembersColumnName field.
func (o *UserRestResourceTypeResponse) SetNonmembersColumnName(v string) {
	o.NonmembersColumnName = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *UserRestResourceTypeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o UserRestResourceTypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.PasswordAttributeCategory) {
		toSerialize["passwordAttributeCategory"] = o.PasswordAttributeCategory
	}
	if !isNil(o.PasswordDisplayOrderIndex) {
		toSerialize["passwordDisplayOrderIndex"] = o.PasswordDisplayOrderIndex
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["resourceEndpoint"] = o.ResourceEndpoint
	}
	if true {
		toSerialize["structuralLDAPObjectclass"] = o.StructuralLDAPObjectclass
	}
	if !isNil(o.AuxiliaryLDAPObjectclass) {
		toSerialize["auxiliaryLDAPObjectclass"] = o.AuxiliaryLDAPObjectclass
	}
	if true {
		toSerialize["searchBaseDN"] = o.SearchBaseDN
	}
	if !isNil(o.IncludeFilter) {
		toSerialize["includeFilter"] = o.IncludeFilter
	}
	if !isNil(o.ParentDN) {
		toSerialize["parentDN"] = o.ParentDN
	}
	if !isNil(o.ParentResourceType) {
		toSerialize["parentResourceType"] = o.ParentResourceType
	}
	if !isNil(o.RelativeDNFromParentResource) {
		toSerialize["relativeDNFromParentResource"] = o.RelativeDNFromParentResource
	}
	if !isNil(o.CreateRDNAttributeType) {
		toSerialize["createRDNAttributeType"] = o.CreateRDNAttributeType
	}
	if !isNil(o.PostCreateConstructedAttribute) {
		toSerialize["postCreateConstructedAttribute"] = o.PostCreateConstructedAttribute
	}
	if !isNil(o.UpdateConstructedAttribute) {
		toSerialize["updateConstructedAttribute"] = o.UpdateConstructedAttribute
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.SearchFilterPattern) {
		toSerialize["searchFilterPattern"] = o.SearchFilterPattern
	}
	if !isNil(o.PrimaryDisplayAttributeType) {
		toSerialize["primaryDisplayAttributeType"] = o.PrimaryDisplayAttributeType
	}
	if !isNil(o.DelegatedAdminSearchSizeLimit) {
		toSerialize["delegatedAdminSearchSizeLimit"] = o.DelegatedAdminSearchSizeLimit
	}
	if !isNil(o.DelegatedAdminReportSizeLimit) {
		toSerialize["delegatedAdminReportSizeLimit"] = o.DelegatedAdminReportSizeLimit
	}
	if !isNil(o.MembersColumnName) {
		toSerialize["membersColumnName"] = o.MembersColumnName
	}
	if !isNil(o.NonmembersColumnName) {
		toSerialize["nonmembersColumnName"] = o.NonmembersColumnName
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableUserRestResourceTypeResponse struct {
	value *UserRestResourceTypeResponse
	isSet bool
}

func (v NullableUserRestResourceTypeResponse) Get() *UserRestResourceTypeResponse {
	return v.value
}

func (v *NullableUserRestResourceTypeResponse) Set(val *UserRestResourceTypeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRestResourceTypeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRestResourceTypeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRestResourceTypeResponse(val *UserRestResourceTypeResponse) *NullableUserRestResourceTypeResponse {
	return &NullableUserRestResourceTypeResponse{value: val, isSet: true}
}

func (v NullableUserRestResourceTypeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRestResourceTypeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


