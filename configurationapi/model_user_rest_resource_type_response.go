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

// checks if the UserRestResourceTypeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRestResourceTypeResponse{}

// UserRestResourceTypeResponse struct for UserRestResourceTypeResponse
type UserRestResourceTypeResponse struct {
	Schemas []EnumuserRestResourceTypeSchemaUrn `json:"schemas"`
	// Specifies which attribute category the password belongs to.
	PasswordAttributeCategory *string `json:"passwordAttributeCategory,omitempty"`
	// This property determines the display order for the password within its attribute category. Attributes are ordered within their category based on this index from least to greatest.
	PasswordDisplayOrderIndex *int64 `json:"passwordDisplayOrderIndex,omitempty"`
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
	DelegatedAdminSearchSizeLimit *int64 `json:"delegatedAdminSearchSizeLimit,omitempty"`
	// The maximum number of resources that may be included in a report.
	DelegatedAdminReportSizeLimit *int64 `json:"delegatedAdminReportSizeLimit,omitempty"`
	// Specifies the name of the group member column that will be displayed in the Delegated Admin UI
	MembersColumnName *string `json:"membersColumnName,omitempty"`
	// Specifies the name of the group nonmember column that will be displayed in the Delegated Admin UI
	NonmembersColumnName                          *string                                            `json:"nonmembersColumnName,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the REST Resource Type
	Id string `json:"id"`
}

// NewUserRestResourceTypeResponse instantiates a new UserRestResourceTypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRestResourceTypeResponse(schemas []EnumuserRestResourceTypeSchemaUrn, enabled bool, resourceEndpoint string, structuralLDAPObjectclass string, searchBaseDN string, id string) *UserRestResourceTypeResponse {
	this := UserRestResourceTypeResponse{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.ResourceEndpoint = resourceEndpoint
	this.StructuralLDAPObjectclass = structuralLDAPObjectclass
	this.SearchBaseDN = searchBaseDN
	this.Id = id
	return &this
}

// NewUserRestResourceTypeResponseWithDefaults instantiates a new UserRestResourceTypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRestResourceTypeResponseWithDefaults() *UserRestResourceTypeResponse {
	this := UserRestResourceTypeResponse{}
	return &this
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
	if o == nil || IsNil(o.PasswordAttributeCategory) {
		var ret string
		return ret
	}
	return *o.PasswordAttributeCategory
}

// GetPasswordAttributeCategoryOk returns a tuple with the PasswordAttributeCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetPasswordAttributeCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.PasswordAttributeCategory) {
		return nil, false
	}
	return o.PasswordAttributeCategory, true
}

// HasPasswordAttributeCategory returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasPasswordAttributeCategory() bool {
	if o != nil && !IsNil(o.PasswordAttributeCategory) {
		return true
	}

	return false
}

// SetPasswordAttributeCategory gets a reference to the given string and assigns it to the PasswordAttributeCategory field.
func (o *UserRestResourceTypeResponse) SetPasswordAttributeCategory(v string) {
	o.PasswordAttributeCategory = &v
}

// GetPasswordDisplayOrderIndex returns the PasswordDisplayOrderIndex field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetPasswordDisplayOrderIndex() int64 {
	if o == nil || IsNil(o.PasswordDisplayOrderIndex) {
		var ret int64
		return ret
	}
	return *o.PasswordDisplayOrderIndex
}

// GetPasswordDisplayOrderIndexOk returns a tuple with the PasswordDisplayOrderIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetPasswordDisplayOrderIndexOk() (*int64, bool) {
	if o == nil || IsNil(o.PasswordDisplayOrderIndex) {
		return nil, false
	}
	return o.PasswordDisplayOrderIndex, true
}

// HasPasswordDisplayOrderIndex returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasPasswordDisplayOrderIndex() bool {
	if o != nil && !IsNil(o.PasswordDisplayOrderIndex) {
		return true
	}

	return false
}

// SetPasswordDisplayOrderIndex gets a reference to the given int64 and assigns it to the PasswordDisplayOrderIndex field.
func (o *UserRestResourceTypeResponse) SetPasswordDisplayOrderIndex(v int64) {
	o.PasswordDisplayOrderIndex = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
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
	if o == nil || IsNil(o.AuxiliaryLDAPObjectclass) {
		var ret []string
		return ret
	}
	return o.AuxiliaryLDAPObjectclass
}

// GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetAuxiliaryLDAPObjectclassOk() ([]string, bool) {
	if o == nil || IsNil(o.AuxiliaryLDAPObjectclass) {
		return nil, false
	}
	return o.AuxiliaryLDAPObjectclass, true
}

// HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasAuxiliaryLDAPObjectclass() bool {
	if o != nil && !IsNil(o.AuxiliaryLDAPObjectclass) {
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
	if o == nil || IsNil(o.IncludeFilter) {
		var ret []string
		return ret
	}
	return o.IncludeFilter
}

// GetIncludeFilterOk returns a tuple with the IncludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetIncludeFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeFilter) {
		return nil, false
	}
	return o.IncludeFilter, true
}

// HasIncludeFilter returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasIncludeFilter() bool {
	if o != nil && !IsNil(o.IncludeFilter) {
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
	if o == nil || IsNil(o.ParentDN) {
		var ret string
		return ret
	}
	return *o.ParentDN
}

// GetParentDNOk returns a tuple with the ParentDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetParentDNOk() (*string, bool) {
	if o == nil || IsNil(o.ParentDN) {
		return nil, false
	}
	return o.ParentDN, true
}

// HasParentDN returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasParentDN() bool {
	if o != nil && !IsNil(o.ParentDN) {
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
	if o == nil || IsNil(o.ParentResourceType) {
		var ret string
		return ret
	}
	return *o.ParentResourceType
}

// GetParentResourceTypeOk returns a tuple with the ParentResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetParentResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParentResourceType) {
		return nil, false
	}
	return o.ParentResourceType, true
}

// HasParentResourceType returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasParentResourceType() bool {
	if o != nil && !IsNil(o.ParentResourceType) {
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
	if o == nil || IsNil(o.RelativeDNFromParentResource) {
		var ret string
		return ret
	}
	return *o.RelativeDNFromParentResource
}

// GetRelativeDNFromParentResourceOk returns a tuple with the RelativeDNFromParentResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetRelativeDNFromParentResourceOk() (*string, bool) {
	if o == nil || IsNil(o.RelativeDNFromParentResource) {
		return nil, false
	}
	return o.RelativeDNFromParentResource, true
}

// HasRelativeDNFromParentResource returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasRelativeDNFromParentResource() bool {
	if o != nil && !IsNil(o.RelativeDNFromParentResource) {
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
	if o == nil || IsNil(o.CreateRDNAttributeType) {
		var ret string
		return ret
	}
	return *o.CreateRDNAttributeType
}

// GetCreateRDNAttributeTypeOk returns a tuple with the CreateRDNAttributeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetCreateRDNAttributeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateRDNAttributeType) {
		return nil, false
	}
	return o.CreateRDNAttributeType, true
}

// HasCreateRDNAttributeType returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasCreateRDNAttributeType() bool {
	if o != nil && !IsNil(o.CreateRDNAttributeType) {
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
	if o == nil || IsNil(o.PostCreateConstructedAttribute) {
		var ret []string
		return ret
	}
	return o.PostCreateConstructedAttribute
}

// GetPostCreateConstructedAttributeOk returns a tuple with the PostCreateConstructedAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetPostCreateConstructedAttributeOk() ([]string, bool) {
	if o == nil || IsNil(o.PostCreateConstructedAttribute) {
		return nil, false
	}
	return o.PostCreateConstructedAttribute, true
}

// HasPostCreateConstructedAttribute returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasPostCreateConstructedAttribute() bool {
	if o != nil && !IsNil(o.PostCreateConstructedAttribute) {
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
	if o == nil || IsNil(o.UpdateConstructedAttribute) {
		var ret []string
		return ret
	}
	return o.UpdateConstructedAttribute
}

// GetUpdateConstructedAttributeOk returns a tuple with the UpdateConstructedAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetUpdateConstructedAttributeOk() ([]string, bool) {
	if o == nil || IsNil(o.UpdateConstructedAttribute) {
		return nil, false
	}
	return o.UpdateConstructedAttribute, true
}

// HasUpdateConstructedAttribute returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasUpdateConstructedAttribute() bool {
	if o != nil && !IsNil(o.UpdateConstructedAttribute) {
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
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
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
	if o == nil || IsNil(o.SearchFilterPattern) {
		var ret string
		return ret
	}
	return *o.SearchFilterPattern
}

// GetSearchFilterPatternOk returns a tuple with the SearchFilterPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetSearchFilterPatternOk() (*string, bool) {
	if o == nil || IsNil(o.SearchFilterPattern) {
		return nil, false
	}
	return o.SearchFilterPattern, true
}

// HasSearchFilterPattern returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasSearchFilterPattern() bool {
	if o != nil && !IsNil(o.SearchFilterPattern) {
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
	if o == nil || IsNil(o.PrimaryDisplayAttributeType) {
		var ret string
		return ret
	}
	return *o.PrimaryDisplayAttributeType
}

// GetPrimaryDisplayAttributeTypeOk returns a tuple with the PrimaryDisplayAttributeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetPrimaryDisplayAttributeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryDisplayAttributeType) {
		return nil, false
	}
	return o.PrimaryDisplayAttributeType, true
}

// HasPrimaryDisplayAttributeType returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasPrimaryDisplayAttributeType() bool {
	if o != nil && !IsNil(o.PrimaryDisplayAttributeType) {
		return true
	}

	return false
}

// SetPrimaryDisplayAttributeType gets a reference to the given string and assigns it to the PrimaryDisplayAttributeType field.
func (o *UserRestResourceTypeResponse) SetPrimaryDisplayAttributeType(v string) {
	o.PrimaryDisplayAttributeType = &v
}

// GetDelegatedAdminSearchSizeLimit returns the DelegatedAdminSearchSizeLimit field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetDelegatedAdminSearchSizeLimit() int64 {
	if o == nil || IsNil(o.DelegatedAdminSearchSizeLimit) {
		var ret int64
		return ret
	}
	return *o.DelegatedAdminSearchSizeLimit
}

// GetDelegatedAdminSearchSizeLimitOk returns a tuple with the DelegatedAdminSearchSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetDelegatedAdminSearchSizeLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.DelegatedAdminSearchSizeLimit) {
		return nil, false
	}
	return o.DelegatedAdminSearchSizeLimit, true
}

// HasDelegatedAdminSearchSizeLimit returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasDelegatedAdminSearchSizeLimit() bool {
	if o != nil && !IsNil(o.DelegatedAdminSearchSizeLimit) {
		return true
	}

	return false
}

// SetDelegatedAdminSearchSizeLimit gets a reference to the given int64 and assigns it to the DelegatedAdminSearchSizeLimit field.
func (o *UserRestResourceTypeResponse) SetDelegatedAdminSearchSizeLimit(v int64) {
	o.DelegatedAdminSearchSizeLimit = &v
}

// GetDelegatedAdminReportSizeLimit returns the DelegatedAdminReportSizeLimit field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetDelegatedAdminReportSizeLimit() int64 {
	if o == nil || IsNil(o.DelegatedAdminReportSizeLimit) {
		var ret int64
		return ret
	}
	return *o.DelegatedAdminReportSizeLimit
}

// GetDelegatedAdminReportSizeLimitOk returns a tuple with the DelegatedAdminReportSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetDelegatedAdminReportSizeLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.DelegatedAdminReportSizeLimit) {
		return nil, false
	}
	return o.DelegatedAdminReportSizeLimit, true
}

// HasDelegatedAdminReportSizeLimit returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasDelegatedAdminReportSizeLimit() bool {
	if o != nil && !IsNil(o.DelegatedAdminReportSizeLimit) {
		return true
	}

	return false
}

// SetDelegatedAdminReportSizeLimit gets a reference to the given int64 and assigns it to the DelegatedAdminReportSizeLimit field.
func (o *UserRestResourceTypeResponse) SetDelegatedAdminReportSizeLimit(v int64) {
	o.DelegatedAdminReportSizeLimit = &v
}

// GetMembersColumnName returns the MembersColumnName field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetMembersColumnName() string {
	if o == nil || IsNil(o.MembersColumnName) {
		var ret string
		return ret
	}
	return *o.MembersColumnName
}

// GetMembersColumnNameOk returns a tuple with the MembersColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetMembersColumnNameOk() (*string, bool) {
	if o == nil || IsNil(o.MembersColumnName) {
		return nil, false
	}
	return o.MembersColumnName, true
}

// HasMembersColumnName returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasMembersColumnName() bool {
	if o != nil && !IsNil(o.MembersColumnName) {
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
	if o == nil || IsNil(o.NonmembersColumnName) {
		var ret string
		return ret
	}
	return *o.NonmembersColumnName
}

// GetNonmembersColumnNameOk returns a tuple with the NonmembersColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetNonmembersColumnNameOk() (*string, bool) {
	if o == nil || IsNil(o.NonmembersColumnName) {
		return nil, false
	}
	return o.NonmembersColumnName, true
}

// HasNonmembersColumnName returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasNonmembersColumnName() bool {
	if o != nil && !IsNil(o.NonmembersColumnName) {
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
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *UserRestResourceTypeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *UserRestResourceTypeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRestResourceTypeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *UserRestResourceTypeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *UserRestResourceTypeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
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

func (o UserRestResourceTypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRestResourceTypeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.PasswordAttributeCategory) {
		toSerialize["passwordAttributeCategory"] = o.PasswordAttributeCategory
	}
	if !IsNil(o.PasswordDisplayOrderIndex) {
		toSerialize["passwordDisplayOrderIndex"] = o.PasswordDisplayOrderIndex
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["resourceEndpoint"] = o.ResourceEndpoint
	toSerialize["structuralLDAPObjectclass"] = o.StructuralLDAPObjectclass
	if !IsNil(o.AuxiliaryLDAPObjectclass) {
		toSerialize["auxiliaryLDAPObjectclass"] = o.AuxiliaryLDAPObjectclass
	}
	toSerialize["searchBaseDN"] = o.SearchBaseDN
	if !IsNil(o.IncludeFilter) {
		toSerialize["includeFilter"] = o.IncludeFilter
	}
	if !IsNil(o.ParentDN) {
		toSerialize["parentDN"] = o.ParentDN
	}
	if !IsNil(o.ParentResourceType) {
		toSerialize["parentResourceType"] = o.ParentResourceType
	}
	if !IsNil(o.RelativeDNFromParentResource) {
		toSerialize["relativeDNFromParentResource"] = o.RelativeDNFromParentResource
	}
	if !IsNil(o.CreateRDNAttributeType) {
		toSerialize["createRDNAttributeType"] = o.CreateRDNAttributeType
	}
	if !IsNil(o.PostCreateConstructedAttribute) {
		toSerialize["postCreateConstructedAttribute"] = o.PostCreateConstructedAttribute
	}
	if !IsNil(o.UpdateConstructedAttribute) {
		toSerialize["updateConstructedAttribute"] = o.UpdateConstructedAttribute
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.SearchFilterPattern) {
		toSerialize["searchFilterPattern"] = o.SearchFilterPattern
	}
	if !IsNil(o.PrimaryDisplayAttributeType) {
		toSerialize["primaryDisplayAttributeType"] = o.PrimaryDisplayAttributeType
	}
	if !IsNil(o.DelegatedAdminSearchSizeLimit) {
		toSerialize["delegatedAdminSearchSizeLimit"] = o.DelegatedAdminSearchSizeLimit
	}
	if !IsNil(o.DelegatedAdminReportSizeLimit) {
		toSerialize["delegatedAdminReportSizeLimit"] = o.DelegatedAdminReportSizeLimit
	}
	if !IsNil(o.MembersColumnName) {
		toSerialize["membersColumnName"] = o.MembersColumnName
	}
	if !IsNil(o.NonmembersColumnName) {
		toSerialize["nonmembersColumnName"] = o.NonmembersColumnName
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
