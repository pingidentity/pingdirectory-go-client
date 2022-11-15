# GroupRestResourceTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the REST Resource Type | 
**Schemas** | [**[]EnumgroupRestResourceTypeSchemaUrn**](EnumgroupRestResourceTypeSchemaUrn.md) |  | 
**MembersColumnName** | Pointer to **string** | Specifies the name of the group member column that will be displayed in the Delegated Admin UI | [optional] 
**NonmembersColumnName** | Pointer to **string** | Specifies the name of the group nonmember column that will be displayed in the Delegated Admin UI | [optional] 
**Description** | Pointer to **string** | A description for this REST Resource Type | [optional] 
**Enabled** | **bool** | Indicates whether the REST Resource Type is enabled. | 
**ResourceEndpoint** | **string** | The HTTP addressable endpoint of this REST Resource Type relative to a REST API base URL. Do not include a leading &#39;/&#39;. | 
**StructuralLDAPObjectclass** | **string** | Specifies the LDAP structural object class that should be exposed by this REST Resource Type. | 
**AuxiliaryLDAPObjectclass** | Pointer to **[]string** |  | [optional] 
**SearchBaseDN** | **string** | Specifies the base DN of the branch of the LDAP directory where resources of this type are located. | 
**IncludeFilter** | Pointer to **[]string** |  | [optional] 
**ParentDN** | Pointer to **string** | Specifies the DN of the parent entry for new resources of this type, when a parent resource is not provided by the app. The parent DN must be at or below the search base of this resource type. | [optional] 
**ParentResourceType** | Pointer to **string** | Specifies the name of another resource type which may be a parent of new resources of this type. The search base DN of the parent resource type must be at or above the search base DN of this resource type. | [optional] 
**RelativeDNFromParentResource** | Pointer to **string** | Specifies a template for a relative DN from the parent resource which identifies the parent entry for a new resource of this type. If this property is not specified then new resources are created immediately below the parent resource or parent DN. | [optional] 
**CreateRDNAttributeType** | Pointer to **string** | Specifies the name or OID of the LDAP attribute type to be used as the RDN of new resources. | [optional] 
**PostCreateConstructedAttribute** | Pointer to **[]string** |  | [optional] 
**UpdateConstructedAttribute** | Pointer to **[]string** |  | [optional] 
**DisplayName** | Pointer to **string** | A human readable display name for this REST Resource Type. | [optional] 
**SearchFilterPattern** | Pointer to **string** | Specifies the LDAP filter that should be used when searching for resources matching provided search text. All attribute types in the filter pattern referencing the search text must have a Delegated Admin Attribute definition. | [optional] 
**PrimaryDisplayAttributeType** | Pointer to **string** | Specifies the name or OID of the LDAP attribute type which is the primary display attribute. This attribute type must be in the search filter pattern and must have a Delegated Admin Attribute definition. | [optional] 
**DelegatedAdminSearchSizeLimit** | Pointer to **int32** | The maximum number of resources that may be returned from a search request. | [optional] 
**DelegatedAdminReportSizeLimit** | Pointer to **int32** | The maximum number of resources that may be included in a report. | [optional] 

## Methods

### NewGroupRestResourceTypeResponse

`func NewGroupRestResourceTypeResponse(id string, schemas []EnumgroupRestResourceTypeSchemaUrn, enabled bool, resourceEndpoint string, structuralLDAPObjectclass string, searchBaseDN string, ) *GroupRestResourceTypeResponse`

NewGroupRestResourceTypeResponse instantiates a new GroupRestResourceTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRestResourceTypeResponseWithDefaults

`func NewGroupRestResourceTypeResponseWithDefaults() *GroupRestResourceTypeResponse`

NewGroupRestResourceTypeResponseWithDefaults instantiates a new GroupRestResourceTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupRestResourceTypeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupRestResourceTypeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupRestResourceTypeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *GroupRestResourceTypeResponse) GetSchemas() []EnumgroupRestResourceTypeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GroupRestResourceTypeResponse) GetSchemasOk() (*[]EnumgroupRestResourceTypeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GroupRestResourceTypeResponse) SetSchemas(v []EnumgroupRestResourceTypeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetMembersColumnName

`func (o *GroupRestResourceTypeResponse) GetMembersColumnName() string`

GetMembersColumnName returns the MembersColumnName field if non-nil, zero value otherwise.

### GetMembersColumnNameOk

`func (o *GroupRestResourceTypeResponse) GetMembersColumnNameOk() (*string, bool)`

GetMembersColumnNameOk returns a tuple with the MembersColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersColumnName

`func (o *GroupRestResourceTypeResponse) SetMembersColumnName(v string)`

SetMembersColumnName sets MembersColumnName field to given value.

### HasMembersColumnName

`func (o *GroupRestResourceTypeResponse) HasMembersColumnName() bool`

HasMembersColumnName returns a boolean if a field has been set.

### GetNonmembersColumnName

`func (o *GroupRestResourceTypeResponse) GetNonmembersColumnName() string`

GetNonmembersColumnName returns the NonmembersColumnName field if non-nil, zero value otherwise.

### GetNonmembersColumnNameOk

`func (o *GroupRestResourceTypeResponse) GetNonmembersColumnNameOk() (*string, bool)`

GetNonmembersColumnNameOk returns a tuple with the NonmembersColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonmembersColumnName

`func (o *GroupRestResourceTypeResponse) SetNonmembersColumnName(v string)`

SetNonmembersColumnName sets NonmembersColumnName field to given value.

### HasNonmembersColumnName

`func (o *GroupRestResourceTypeResponse) HasNonmembersColumnName() bool`

HasNonmembersColumnName returns a boolean if a field has been set.

### GetDescription

`func (o *GroupRestResourceTypeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupRestResourceTypeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupRestResourceTypeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupRestResourceTypeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GroupRestResourceTypeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GroupRestResourceTypeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GroupRestResourceTypeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetResourceEndpoint

`func (o *GroupRestResourceTypeResponse) GetResourceEndpoint() string`

GetResourceEndpoint returns the ResourceEndpoint field if non-nil, zero value otherwise.

### GetResourceEndpointOk

`func (o *GroupRestResourceTypeResponse) GetResourceEndpointOk() (*string, bool)`

GetResourceEndpointOk returns a tuple with the ResourceEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceEndpoint

`func (o *GroupRestResourceTypeResponse) SetResourceEndpoint(v string)`

SetResourceEndpoint sets ResourceEndpoint field to given value.


### GetStructuralLDAPObjectclass

`func (o *GroupRestResourceTypeResponse) GetStructuralLDAPObjectclass() string`

GetStructuralLDAPObjectclass returns the StructuralLDAPObjectclass field if non-nil, zero value otherwise.

### GetStructuralLDAPObjectclassOk

`func (o *GroupRestResourceTypeResponse) GetStructuralLDAPObjectclassOk() (*string, bool)`

GetStructuralLDAPObjectclassOk returns a tuple with the StructuralLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuralLDAPObjectclass

`func (o *GroupRestResourceTypeResponse) SetStructuralLDAPObjectclass(v string)`

SetStructuralLDAPObjectclass sets StructuralLDAPObjectclass field to given value.


### GetAuxiliaryLDAPObjectclass

`func (o *GroupRestResourceTypeResponse) GetAuxiliaryLDAPObjectclass() []string`

GetAuxiliaryLDAPObjectclass returns the AuxiliaryLDAPObjectclass field if non-nil, zero value otherwise.

### GetAuxiliaryLDAPObjectclassOk

`func (o *GroupRestResourceTypeResponse) GetAuxiliaryLDAPObjectclassOk() (*[]string, bool)`

GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryLDAPObjectclass

`func (o *GroupRestResourceTypeResponse) SetAuxiliaryLDAPObjectclass(v []string)`

SetAuxiliaryLDAPObjectclass sets AuxiliaryLDAPObjectclass field to given value.

### HasAuxiliaryLDAPObjectclass

`func (o *GroupRestResourceTypeResponse) HasAuxiliaryLDAPObjectclass() bool`

HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.

### GetSearchBaseDN

`func (o *GroupRestResourceTypeResponse) GetSearchBaseDN() string`

GetSearchBaseDN returns the SearchBaseDN field if non-nil, zero value otherwise.

### GetSearchBaseDNOk

`func (o *GroupRestResourceTypeResponse) GetSearchBaseDNOk() (*string, bool)`

GetSearchBaseDNOk returns a tuple with the SearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBaseDN

`func (o *GroupRestResourceTypeResponse) SetSearchBaseDN(v string)`

SetSearchBaseDN sets SearchBaseDN field to given value.


### GetIncludeFilter

`func (o *GroupRestResourceTypeResponse) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *GroupRestResourceTypeResponse) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *GroupRestResourceTypeResponse) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *GroupRestResourceTypeResponse) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetParentDN

`func (o *GroupRestResourceTypeResponse) GetParentDN() string`

GetParentDN returns the ParentDN field if non-nil, zero value otherwise.

### GetParentDNOk

`func (o *GroupRestResourceTypeResponse) GetParentDNOk() (*string, bool)`

GetParentDNOk returns a tuple with the ParentDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDN

`func (o *GroupRestResourceTypeResponse) SetParentDN(v string)`

SetParentDN sets ParentDN field to given value.

### HasParentDN

`func (o *GroupRestResourceTypeResponse) HasParentDN() bool`

HasParentDN returns a boolean if a field has been set.

### GetParentResourceType

`func (o *GroupRestResourceTypeResponse) GetParentResourceType() string`

GetParentResourceType returns the ParentResourceType field if non-nil, zero value otherwise.

### GetParentResourceTypeOk

`func (o *GroupRestResourceTypeResponse) GetParentResourceTypeOk() (*string, bool)`

GetParentResourceTypeOk returns a tuple with the ParentResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceType

`func (o *GroupRestResourceTypeResponse) SetParentResourceType(v string)`

SetParentResourceType sets ParentResourceType field to given value.

### HasParentResourceType

`func (o *GroupRestResourceTypeResponse) HasParentResourceType() bool`

HasParentResourceType returns a boolean if a field has been set.

### GetRelativeDNFromParentResource

`func (o *GroupRestResourceTypeResponse) GetRelativeDNFromParentResource() string`

GetRelativeDNFromParentResource returns the RelativeDNFromParentResource field if non-nil, zero value otherwise.

### GetRelativeDNFromParentResourceOk

`func (o *GroupRestResourceTypeResponse) GetRelativeDNFromParentResourceOk() (*string, bool)`

GetRelativeDNFromParentResourceOk returns a tuple with the RelativeDNFromParentResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeDNFromParentResource

`func (o *GroupRestResourceTypeResponse) SetRelativeDNFromParentResource(v string)`

SetRelativeDNFromParentResource sets RelativeDNFromParentResource field to given value.

### HasRelativeDNFromParentResource

`func (o *GroupRestResourceTypeResponse) HasRelativeDNFromParentResource() bool`

HasRelativeDNFromParentResource returns a boolean if a field has been set.

### GetCreateRDNAttributeType

`func (o *GroupRestResourceTypeResponse) GetCreateRDNAttributeType() string`

GetCreateRDNAttributeType returns the CreateRDNAttributeType field if non-nil, zero value otherwise.

### GetCreateRDNAttributeTypeOk

`func (o *GroupRestResourceTypeResponse) GetCreateRDNAttributeTypeOk() (*string, bool)`

GetCreateRDNAttributeTypeOk returns a tuple with the CreateRDNAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateRDNAttributeType

`func (o *GroupRestResourceTypeResponse) SetCreateRDNAttributeType(v string)`

SetCreateRDNAttributeType sets CreateRDNAttributeType field to given value.

### HasCreateRDNAttributeType

`func (o *GroupRestResourceTypeResponse) HasCreateRDNAttributeType() bool`

HasCreateRDNAttributeType returns a boolean if a field has been set.

### GetPostCreateConstructedAttribute

`func (o *GroupRestResourceTypeResponse) GetPostCreateConstructedAttribute() []string`

GetPostCreateConstructedAttribute returns the PostCreateConstructedAttribute field if non-nil, zero value otherwise.

### GetPostCreateConstructedAttributeOk

`func (o *GroupRestResourceTypeResponse) GetPostCreateConstructedAttributeOk() (*[]string, bool)`

GetPostCreateConstructedAttributeOk returns a tuple with the PostCreateConstructedAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCreateConstructedAttribute

`func (o *GroupRestResourceTypeResponse) SetPostCreateConstructedAttribute(v []string)`

SetPostCreateConstructedAttribute sets PostCreateConstructedAttribute field to given value.

### HasPostCreateConstructedAttribute

`func (o *GroupRestResourceTypeResponse) HasPostCreateConstructedAttribute() bool`

HasPostCreateConstructedAttribute returns a boolean if a field has been set.

### GetUpdateConstructedAttribute

`func (o *GroupRestResourceTypeResponse) GetUpdateConstructedAttribute() []string`

GetUpdateConstructedAttribute returns the UpdateConstructedAttribute field if non-nil, zero value otherwise.

### GetUpdateConstructedAttributeOk

`func (o *GroupRestResourceTypeResponse) GetUpdateConstructedAttributeOk() (*[]string, bool)`

GetUpdateConstructedAttributeOk returns a tuple with the UpdateConstructedAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateConstructedAttribute

`func (o *GroupRestResourceTypeResponse) SetUpdateConstructedAttribute(v []string)`

SetUpdateConstructedAttribute sets UpdateConstructedAttribute field to given value.

### HasUpdateConstructedAttribute

`func (o *GroupRestResourceTypeResponse) HasUpdateConstructedAttribute() bool`

HasUpdateConstructedAttribute returns a boolean if a field has been set.

### GetDisplayName

`func (o *GroupRestResourceTypeResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GroupRestResourceTypeResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GroupRestResourceTypeResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GroupRestResourceTypeResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSearchFilterPattern

`func (o *GroupRestResourceTypeResponse) GetSearchFilterPattern() string`

GetSearchFilterPattern returns the SearchFilterPattern field if non-nil, zero value otherwise.

### GetSearchFilterPatternOk

`func (o *GroupRestResourceTypeResponse) GetSearchFilterPatternOk() (*string, bool)`

GetSearchFilterPatternOk returns a tuple with the SearchFilterPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterPattern

`func (o *GroupRestResourceTypeResponse) SetSearchFilterPattern(v string)`

SetSearchFilterPattern sets SearchFilterPattern field to given value.

### HasSearchFilterPattern

`func (o *GroupRestResourceTypeResponse) HasSearchFilterPattern() bool`

HasSearchFilterPattern returns a boolean if a field has been set.

### GetPrimaryDisplayAttributeType

`func (o *GroupRestResourceTypeResponse) GetPrimaryDisplayAttributeType() string`

GetPrimaryDisplayAttributeType returns the PrimaryDisplayAttributeType field if non-nil, zero value otherwise.

### GetPrimaryDisplayAttributeTypeOk

`func (o *GroupRestResourceTypeResponse) GetPrimaryDisplayAttributeTypeOk() (*string, bool)`

GetPrimaryDisplayAttributeTypeOk returns a tuple with the PrimaryDisplayAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDisplayAttributeType

`func (o *GroupRestResourceTypeResponse) SetPrimaryDisplayAttributeType(v string)`

SetPrimaryDisplayAttributeType sets PrimaryDisplayAttributeType field to given value.

### HasPrimaryDisplayAttributeType

`func (o *GroupRestResourceTypeResponse) HasPrimaryDisplayAttributeType() bool`

HasPrimaryDisplayAttributeType returns a boolean if a field has been set.

### GetDelegatedAdminSearchSizeLimit

`func (o *GroupRestResourceTypeResponse) GetDelegatedAdminSearchSizeLimit() int32`

GetDelegatedAdminSearchSizeLimit returns the DelegatedAdminSearchSizeLimit field if non-nil, zero value otherwise.

### GetDelegatedAdminSearchSizeLimitOk

`func (o *GroupRestResourceTypeResponse) GetDelegatedAdminSearchSizeLimitOk() (*int32, bool)`

GetDelegatedAdminSearchSizeLimitOk returns a tuple with the DelegatedAdminSearchSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAdminSearchSizeLimit

`func (o *GroupRestResourceTypeResponse) SetDelegatedAdminSearchSizeLimit(v int32)`

SetDelegatedAdminSearchSizeLimit sets DelegatedAdminSearchSizeLimit field to given value.

### HasDelegatedAdminSearchSizeLimit

`func (o *GroupRestResourceTypeResponse) HasDelegatedAdminSearchSizeLimit() bool`

HasDelegatedAdminSearchSizeLimit returns a boolean if a field has been set.

### GetDelegatedAdminReportSizeLimit

`func (o *GroupRestResourceTypeResponse) GetDelegatedAdminReportSizeLimit() int32`

GetDelegatedAdminReportSizeLimit returns the DelegatedAdminReportSizeLimit field if non-nil, zero value otherwise.

### GetDelegatedAdminReportSizeLimitOk

`func (o *GroupRestResourceTypeResponse) GetDelegatedAdminReportSizeLimitOk() (*int32, bool)`

GetDelegatedAdminReportSizeLimitOk returns a tuple with the DelegatedAdminReportSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAdminReportSizeLimit

`func (o *GroupRestResourceTypeResponse) SetDelegatedAdminReportSizeLimit(v int32)`

SetDelegatedAdminReportSizeLimit sets DelegatedAdminReportSizeLimit field to given value.

### HasDelegatedAdminReportSizeLimit

`func (o *GroupRestResourceTypeResponse) HasDelegatedAdminReportSizeLimit() bool`

HasDelegatedAdminReportSizeLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


