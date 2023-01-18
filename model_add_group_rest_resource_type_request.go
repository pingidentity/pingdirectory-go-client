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

// AddGroupRestResourceTypeRequest struct for AddGroupRestResourceTypeRequest
type AddGroupRestResourceTypeRequest struct {
	// Name of the new REST Resource Type
	TypeName string `json:"typeName"`
	Schemas []EnumgroupRestResourceTypeSchemaUrn `json:"schemas"`
	// Specifies the name of the group member column that will be displayed in the Delegated Admin UI
	MembersColumnName *string `json:"membersColumnName,omitempty"`
	// Specifies the name of the group nonmember column that will be displayed in the Delegated Admin UI
	NonmembersColumnName *string `json:"nonmembersColumnName,omitempty"`
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
}

// NewAddGroupRestResourceTypeRequest instantiates a new AddGroupRestResourceTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddGroupRestResourceTypeRequest(typeName string, schemas []EnumgroupRestResourceTypeSchemaUrn, enabled bool, resourceEndpoint string, structuralLDAPObjectclass string, searchBaseDN string) *AddGroupRestResourceTypeRequest {
	this := AddGroupRestResourceTypeRequest{}
	this.TypeName = typeName
	this.Schemas = schemas
	this.Enabled = enabled
	this.ResourceEndpoint = resourceEndpoint
	this.StructuralLDAPObjectclass = structuralLDAPObjectclass
	this.SearchBaseDN = searchBaseDN
	return &this
}

// NewAddGroupRestResourceTypeRequestWithDefaults instantiates a new AddGroupRestResourceTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddGroupRestResourceTypeRequestWithDefaults() *AddGroupRestResourceTypeRequest {
	this := AddGroupRestResourceTypeRequest{}
	return &this
}

// GetTypeName returns the TypeName field value
func (o *AddGroupRestResourceTypeRequest) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetTypeNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *AddGroupRestResourceTypeRequest) SetTypeName(v string) {
	o.TypeName = v
}

// GetSchemas returns the Schemas field value
func (o *AddGroupRestResourceTypeRequest) GetSchemas() []EnumgroupRestResourceTypeSchemaUrn {
	if o == nil {
		var ret []EnumgroupRestResourceTypeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetSchemasOk() ([]EnumgroupRestResourceTypeSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddGroupRestResourceTypeRequest) SetSchemas(v []EnumgroupRestResourceTypeSchemaUrn) {
	o.Schemas = v
}

// GetMembersColumnName returns the MembersColumnName field value if set, zero value otherwise.
func (o *AddGroupRestResourceTypeRequest) GetMembersColumnName() string {
	if o == nil || isNil(o.MembersColumnName) {
		var ret string
		return ret
	}
	return *o.MembersColumnName
}

// GetMembersColumnNameOk returns a tuple with the MembersColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetMembersColumnNameOk() (*string, bool) {
	if o == nil || isNil(o.MembersColumnName) {
    return nil, false
	}
	return o.MembersColumnName, true
}

// HasMembersColumnName returns a boolean if a field has been set.
func (o *AddGroupRestResourceTypeRequest) HasMembersColumnName() bool {
	if o != nil && !isNil(o.MembersColumnName) {
		return true
	}

	return false
}

// SetMembersColumnName gets a reference to the given string and assigns it to the MembersColumnName field.
func (o *AddGroupRestResourceTypeRequest) SetMembersColumnName(v string) {
	o.MembersColumnName = &v
}

// GetNonmembersColumnName returns the NonmembersColumnName field value if set, zero value otherwise.
func (o *AddGroupRestResourceTypeRequest) GetNonmembersColumnName() string {
	if o == nil || isNil(o.NonmembersColumnName) {
		var ret string
		return ret
	}
	return *o.NonmembersColumnName
}

// GetNonmembersColumnNameOk returns a tuple with the NonmembersColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetNonmembersColumnNameOk() (*string, bool) {
	if o == nil || isNil(o.NonmembersColumnName) {
    return nil, false
	}
	return o.NonmembersColumnName, true
}

// HasNonmembersColumnName returns a boolean if a field has been set.
func (o *AddGroupRestResourceTypeRequest) HasNonmembersColumnName() bool {
	if o != nil && !isNil(o.NonmembersColumnName) {
		return true
	}

	return false
}

// SetNonmembersColumnName gets a reference to the given string and assigns it to the NonmembersColumnName field.
func (o *AddGroupRestResourceTypeRequest) SetNonmembersColumnName(v string) {
	o.NonmembersColumnName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddGroupRestResourceTypeRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddGroupRestResourceTypeRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddGroupRestResourceTypeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddGroupRestResourceTypeRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddGroupRestResourceTypeRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetResourceEndpoint returns the ResourceEndpoint field value
func (o *AddGroupRestResourceTypeRequest) GetResourceEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceEndpoint
}

// GetResourceEndpointOk returns a tuple with the ResourceEndpoint field value
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetResourceEndpointOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ResourceEndpoint, true
}

// SetResourceEndpoint sets field value
func (o *AddGroupRestResourceTypeRequest) SetResourceEndpoint(v string) {
	o.ResourceEndpoint = v
}

// GetStructuralLDAPObjectclass returns the StructuralLDAPObjectclass field value
func (o *AddGroupRestResourceTypeRequest) GetStructuralLDAPObjectclass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StructuralLDAPObjectclass
}

// GetStructuralLDAPObjectclassOk returns a tuple with the StructuralLDAPObjectclass field value
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetStructuralLDAPObjectclassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StructuralLDAPObjectclass, true
}

// SetStructuralLDAPObjectclass sets field value
func (o *AddGroupRestResourceTypeRequest) SetStructuralLDAPObjectclass(v string) {
	o.StructuralLDAPObjectclass = v
}

// GetAuxiliaryLDAPObjectclass returns the AuxiliaryLDAPObjectclass field value if set, zero value otherwise.
func (o *AddGroupRestResourceTypeRequest) GetAuxiliaryLDAPObjectclass() []string {
	if o == nil || isNil(o.AuxiliaryLDAPObjectclass) {
		var ret []string
		return ret
	}
	return o.AuxiliaryLDAPObjectclass
}

// GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetAuxiliaryLDAPObjectclassOk() ([]string, bool) {
	if o == nil || isNil(o.AuxiliaryLDAPObjectclass) {
    return nil, false
	}
	return o.AuxiliaryLDAPObjectclass, true
}

// HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.
func (o *AddGroupRestResourceTypeRequest) HasAuxiliaryLDAPObjectclass() bool {
	if o != nil && !isNil(o.AuxiliaryLDAPObjectclass) {
		return true
	}

	return false
}

// SetAuxiliaryLDAPObjectclass gets a reference to the given []string and assigns it to the AuxiliaryLDAPObjectclass field.
func (o *AddGroupRestResourceTypeRequest) SetAuxiliaryLDAPObjectclass(v []string) {
	o.AuxiliaryLDAPObjectclass = v
}

// GetSearchBaseDN returns the SearchBaseDN field value
func (o *AddGroupRestResourceTypeRequest) GetSearchBaseDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SearchBaseDN
}

// GetSearchBaseDNOk returns a tuple with the SearchBaseDN field value
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetSearchBaseDNOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SearchBaseDN, true
}

// SetSearchBaseDN sets field value
func (o *AddGroupRestResourceTypeRequest) SetSearchBaseDN(v string) {
	o.SearchBaseDN = v
}

// GetIncludeFilter returns the IncludeFilter field value if set, zero value otherwise.
func (o *AddGroupRestResourceTypeRequest) GetIncludeFilter() []string {
	if o == nil || isNil(o.IncludeFilter) {
		var ret []string
		return ret
	}
	return o.IncludeFilter
}

// GetIncludeFilterOk returns a tuple with the IncludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetIncludeFilterOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeFilter) {
    return nil, false
	}
	return o.IncludeFilter, true
}

// HasIncludeFilter returns a boolean if a field has been set.
func (o *AddGroupRestResourceTypeRequest) HasIncludeFilter() bool {
	if o != nil && !isNil(o.IncludeFilter) {
		return true
	}

	return false
}

// SetIncludeFilter gets a reference to the given []string and assigns it to the IncludeFilter field.
func (o *AddGroupRestResourceTypeRequest) SetIncludeFilter(v []string) {
	o.IncludeFilter = v
}

// GetParentDN returns the ParentDN field value if set, zero value otherwise.
func (o *AddGroupRestResourceTypeRequest) GetParentDN() string {
	if o == nil || isNil(o.ParentDN) {
		var ret string
		return ret
	}
	return *o.ParentDN
}

// GetParentDNOk returns a tuple with the ParentDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetParentDNOk() (*string, bool) {
	if o == nil || isNil(o.ParentDN) {
    return nil, false
	}
	return o.ParentDN, true
}

// HasParentDN returns a boolean if a field has been set.
func (o *AddGroupRestResourceTypeRequest) HasParentDN() bool {
	if o != nil && !isNil(o.ParentDN) {
		return true
	}

	return false
}

// SetParentDN gets a reference to the given string and assigns it to the ParentDN field.
func (o *AddGroupRestResourceTypeRequest) SetParentDN(v string) {
	o.ParentDN = &v
}

// GetParentResourceType returns the ParentResourceType field value if set, zero value otherwise.
func (o *AddGroupRestResourceTypeRequest) GetParentResourceType() string {
	if o == nil || isNil(o.ParentResourceType) {
		var ret string
		return ret
	}
	return *o.ParentResourceType
}

// GetParentResourceTypeOk returns a tuple with the ParentResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetParentResourceTypeOk() (*string, bool) {
	if o == nil || isNil(o.ParentResourceType) {
    return nil, false
	}
	return o.ParentResourceType, true
}

// HasParentResourceType returns a boolean if a field has been set.
func (o *AddGroupRestResourceTypeRequest) HasParentResourceType() bool {
	if o != nil && !isNil(o.ParentResourceType) {
		return true
	}

	return false
}

// SetParentResourceType gets a reference to the given string and assigns it to the ParentResourceType field.
func (o *AddGroupRestResourceTypeRequest) SetParentResourceType(v string) {
	o.ParentResourceType = &v
}

// GetRelativeDNFromParentResource returns the RelativeDNFromParentResource field value if set, zero value otherwise.
func (o *AddGroupRestResourceTypeRequest) GetRelativeDNFromParentResource() string {
	if o == nil || isNil(o.RelativeDNFromParentResource) {
		var ret string
		return ret
	}
	return *o.RelativeDNFromParentResource
}

// GetRelativeDNFromParentResourceOk returns a tuple with the RelativeDNFromParentResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetRelativeDNFromParentResourceOk() (*string, bool) {
	if o == nil || isNil(o.RelativeDNFromParentResource) {
    return nil, false
	}
	return o.RelativeDNFromParentResource, true
}

// HasRelativeDNFromParentResource returns a boolean if a field has been set.
func (o *AddGroupRestResourceTypeRequest) HasRelativeDNFromParentResource() bool {
	if o != nil && !isNil(o.RelativeDNFromParentResource) {
		return true
	}

	return false
}

// SetRelativeDNFromParentResource gets a reference to the given string and assigns it to the RelativeDNFromParentResource field.
func (o *AddGroupRestResourceTypeRequest) SetRelativeDNFromParentResource(v string) {
	o.RelativeDNFromParentResource = &v
}

// GetCreateRDNAttributeType returns the CreateRDNAttributeType field value if set, zero value otherwise.
func (o *AddGroupRestResourceTypeRequest) GetCreateRDNAttributeType() string {
	if o == nil || isNil(o.CreateRDNAttributeType) {
		var ret string
		return ret
	}
	return *o.CreateRDNAttributeType
}

// GetCreateRDNAttributeTypeOk returns a tuple with the CreateRDNAttributeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetCreateRDNAttributeTypeOk() (*string, bool) {
	if o == nil || isNil(o.CreateRDNAttributeType) {
    return nil, false
	}
	return o.CreateRDNAttributeType, true
}

// HasCreateRDNAttributeType returns a boolean if a field has been set.
func (o *AddGroupRestResourceTypeRequest) HasCreateRDNAttributeType() bool {
	if o != nil && !isNil(o.CreateRDNAttributeType) {
		return true
	}

	return false
}

// SetCreateRDNAttributeType gets a reference to the given string and assigns it to the CreateRDNAttributeType field.
func (o *AddGroupRestResourceTypeRequest) SetCreateRDNAttributeType(v string) {
	o.CreateRDNAttributeType = &v
}

// GetPostCreateConstructedAttribute returns the PostCreateConstructedAttribute field value if set, zero value otherwise.
func (o *AddGroupRestResourceTypeRequest) GetPostCreateConstructedAttribute() []string {
	if o == nil || isNil(o.PostCreateConstructedAttribute) {
		var ret []string
		return ret
	}
	return o.PostCreateConstructedAttribute
}

// GetPostCreateConstructedAttributeOk returns a tuple with the PostCreateConstructedAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetPostCreateConstructedAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.PostCreateConstructedAttribute) {
    return nil, false
	}
	return o.PostCreateConstructedAttribute, true
}

// HasPostCreateConstructedAttribute returns a boolean if a field has been set.
func (o *AddGroupRestResourceTypeRequest) HasPostCreateConstructedAttribute() bool {
	if o != nil && !isNil(o.PostCreateConstructedAttribute) {
		return true
	}

	return false
}

// SetPostCreateConstructedAttribute gets a reference to the given []string and assigns it to the PostCreateConstructedAttribute field.
func (o *AddGroupRestResourceTypeRequest) SetPostCreateConstructedAttribute(v []string) {
	o.PostCreateConstructedAttribute = v
}

// GetUpdateConstructedAttribute returns the UpdateConstructedAttribute field value if set, zero value otherwise.
func (o *AddGroupRestResourceTypeRequest) GetUpdateConstructedAttribute() []string {
	if o == nil || isNil(o.UpdateConstructedAttribute) {
		var ret []string
		return ret
	}
	return o.UpdateConstructedAttribute
}

// GetUpdateConstructedAttributeOk returns a tuple with the UpdateConstructedAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetUpdateConstructedAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.UpdateConstructedAttribute) {
    return nil, false
	}
	return o.UpdateConstructedAttribute, true
}

// HasUpdateConstructedAttribute returns a boolean if a field has been set.
func (o *AddGroupRestResourceTypeRequest) HasUpdateConstructedAttribute() bool {
	if o != nil && !isNil(o.UpdateConstructedAttribute) {
		return true
	}

	return false
}

// SetUpdateConstructedAttribute gets a reference to the given []string and assigns it to the UpdateConstructedAttribute field.
func (o *AddGroupRestResourceTypeRequest) SetUpdateConstructedAttribute(v []string) {
	o.UpdateConstructedAttribute = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AddGroupRestResourceTypeRequest) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AddGroupRestResourceTypeRequest) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AddGroupRestResourceTypeRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSearchFilterPattern returns the SearchFilterPattern field value if set, zero value otherwise.
func (o *AddGroupRestResourceTypeRequest) GetSearchFilterPattern() string {
	if o == nil || isNil(o.SearchFilterPattern) {
		var ret string
		return ret
	}
	return *o.SearchFilterPattern
}

// GetSearchFilterPatternOk returns a tuple with the SearchFilterPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetSearchFilterPatternOk() (*string, bool) {
	if o == nil || isNil(o.SearchFilterPattern) {
    return nil, false
	}
	return o.SearchFilterPattern, true
}

// HasSearchFilterPattern returns a boolean if a field has been set.
func (o *AddGroupRestResourceTypeRequest) HasSearchFilterPattern() bool {
	if o != nil && !isNil(o.SearchFilterPattern) {
		return true
	}

	return false
}

// SetSearchFilterPattern gets a reference to the given string and assigns it to the SearchFilterPattern field.
func (o *AddGroupRestResourceTypeRequest) SetSearchFilterPattern(v string) {
	o.SearchFilterPattern = &v
}

// GetPrimaryDisplayAttributeType returns the PrimaryDisplayAttributeType field value if set, zero value otherwise.
func (o *AddGroupRestResourceTypeRequest) GetPrimaryDisplayAttributeType() string {
	if o == nil || isNil(o.PrimaryDisplayAttributeType) {
		var ret string
		return ret
	}
	return *o.PrimaryDisplayAttributeType
}

// GetPrimaryDisplayAttributeTypeOk returns a tuple with the PrimaryDisplayAttributeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetPrimaryDisplayAttributeTypeOk() (*string, bool) {
	if o == nil || isNil(o.PrimaryDisplayAttributeType) {
    return nil, false
	}
	return o.PrimaryDisplayAttributeType, true
}

// HasPrimaryDisplayAttributeType returns a boolean if a field has been set.
func (o *AddGroupRestResourceTypeRequest) HasPrimaryDisplayAttributeType() bool {
	if o != nil && !isNil(o.PrimaryDisplayAttributeType) {
		return true
	}

	return false
}

// SetPrimaryDisplayAttributeType gets a reference to the given string and assigns it to the PrimaryDisplayAttributeType field.
func (o *AddGroupRestResourceTypeRequest) SetPrimaryDisplayAttributeType(v string) {
	o.PrimaryDisplayAttributeType = &v
}

// GetDelegatedAdminSearchSizeLimit returns the DelegatedAdminSearchSizeLimit field value if set, zero value otherwise.
func (o *AddGroupRestResourceTypeRequest) GetDelegatedAdminSearchSizeLimit() int32 {
	if o == nil || isNil(o.DelegatedAdminSearchSizeLimit) {
		var ret int32
		return ret
	}
	return *o.DelegatedAdminSearchSizeLimit
}

// GetDelegatedAdminSearchSizeLimitOk returns a tuple with the DelegatedAdminSearchSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetDelegatedAdminSearchSizeLimitOk() (*int32, bool) {
	if o == nil || isNil(o.DelegatedAdminSearchSizeLimit) {
    return nil, false
	}
	return o.DelegatedAdminSearchSizeLimit, true
}

// HasDelegatedAdminSearchSizeLimit returns a boolean if a field has been set.
func (o *AddGroupRestResourceTypeRequest) HasDelegatedAdminSearchSizeLimit() bool {
	if o != nil && !isNil(o.DelegatedAdminSearchSizeLimit) {
		return true
	}

	return false
}

// SetDelegatedAdminSearchSizeLimit gets a reference to the given int32 and assigns it to the DelegatedAdminSearchSizeLimit field.
func (o *AddGroupRestResourceTypeRequest) SetDelegatedAdminSearchSizeLimit(v int32) {
	o.DelegatedAdminSearchSizeLimit = &v
}

// GetDelegatedAdminReportSizeLimit returns the DelegatedAdminReportSizeLimit field value if set, zero value otherwise.
func (o *AddGroupRestResourceTypeRequest) GetDelegatedAdminReportSizeLimit() int32 {
	if o == nil || isNil(o.DelegatedAdminReportSizeLimit) {
		var ret int32
		return ret
	}
	return *o.DelegatedAdminReportSizeLimit
}

// GetDelegatedAdminReportSizeLimitOk returns a tuple with the DelegatedAdminReportSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRestResourceTypeRequest) GetDelegatedAdminReportSizeLimitOk() (*int32, bool) {
	if o == nil || isNil(o.DelegatedAdminReportSizeLimit) {
    return nil, false
	}
	return o.DelegatedAdminReportSizeLimit, true
}

// HasDelegatedAdminReportSizeLimit returns a boolean if a field has been set.
func (o *AddGroupRestResourceTypeRequest) HasDelegatedAdminReportSizeLimit() bool {
	if o != nil && !isNil(o.DelegatedAdminReportSizeLimit) {
		return true
	}

	return false
}

// SetDelegatedAdminReportSizeLimit gets a reference to the given int32 and assigns it to the DelegatedAdminReportSizeLimit field.
func (o *AddGroupRestResourceTypeRequest) SetDelegatedAdminReportSizeLimit(v int32) {
	o.DelegatedAdminReportSizeLimit = &v
}

func (o AddGroupRestResourceTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["typeName"] = o.TypeName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.MembersColumnName) {
		toSerialize["membersColumnName"] = o.MembersColumnName
	}
	if !isNil(o.NonmembersColumnName) {
		toSerialize["nonmembersColumnName"] = o.NonmembersColumnName
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
	return json.Marshal(toSerialize)
}

type NullableAddGroupRestResourceTypeRequest struct {
	value *AddGroupRestResourceTypeRequest
	isSet bool
}

func (v NullableAddGroupRestResourceTypeRequest) Get() *AddGroupRestResourceTypeRequest {
	return v.value
}

func (v *NullableAddGroupRestResourceTypeRequest) Set(val *AddGroupRestResourceTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddGroupRestResourceTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddGroupRestResourceTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddGroupRestResourceTypeRequest(val *AddGroupRestResourceTypeRequest) *NullableAddGroupRestResourceTypeRequest {
	return &NullableAddGroupRestResourceTypeRequest{value: val, isSet: true}
}

func (v NullableAddGroupRestResourceTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddGroupRestResourceTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

