# AddGenericRestResourceTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | **string** | Name of the new REST Resource Type | 
**Schemas** | [**[]EnumgenericRestResourceTypeSchemaUrn**](EnumgenericRestResourceTypeSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this REST Resource Type | [optional] 
**Enabled** | **bool** | Indicates whether the REST Resource Type is enabled. | 
**ResourceEndpoint** | **string** | The HTTP addressable endpoint of this REST Resource Type relative to a REST API base URL. Do not include a leading &#39;/&#39;. | 
**StructuralLDAPObjectclass** | **string** | Specifies the LDAP structural object class that should be exposed by this REST Resource Type. | 
**AuxiliaryLDAPObjectclass** | Pointer to **[]string** | Specifies an auxiliary LDAP object class that should be exposed by this REST Resource Type. | [optional] 
**SearchBaseDN** | **string** | Specifies the base DN of the branch of the LDAP directory where resources of this type are located. | 
**IncludeFilter** | Pointer to **[]string** | The set of LDAP filters that define the LDAP entries that should be included in this REST Resource Type. | [optional] 
**ParentDN** | Pointer to **string** | Specifies the DN of the parent entry for new resources of this type, when a parent resource is not provided by the app. The parent DN must be at or below the search base of this resource type. | [optional] 
**ParentResourceType** | Pointer to **string** | Specifies the name of another resource type which may be a parent of new resources of this type. The search base DN of the parent resource type must be at or above the search base DN of this resource type. | [optional] 
**RelativeDNFromParentResource** | Pointer to **string** | Specifies a template for a relative DN from the parent resource which identifies the parent entry for a new resource of this type. If this property is not specified then new resources are created immediately below the parent resource or parent DN. | [optional] 
**CreateRDNAttributeType** | Pointer to **string** | Specifies the name or OID of the LDAP attribute type to be used as the RDN of new resources. | [optional] 
**PostCreateConstructedAttribute** | Pointer to **[]string** | Specifies an attribute whose values are to be constructed when a new resource is created. The values are only set at creation time. Subsequent modifications to attributes in the constructed attribute value-pattern are not propagated here. | [optional] 
**UpdateConstructedAttribute** | Pointer to **[]string** | Specifies an attribute whose values are to be constructed when a resource is updated. The constructed values replace any existing values of the attribute. | [optional] 
**DisplayName** | Pointer to **string** | A human readable display name for this REST Resource Type. | [optional] 
**SearchFilterPattern** | Pointer to **string** | Specifies the LDAP filter that should be used when searching for resources matching provided search text. All attribute types in the filter pattern referencing the search text must have a Delegated Admin Attribute definition. | [optional] 
**PrimaryDisplayAttributeType** | Pointer to **string** | Specifies the name or OID of the LDAP attribute type which is the primary display attribute. This attribute type must be in the search filter pattern and must have a Delegated Admin Attribute definition. | [optional] 
**DelegatedAdminSearchSizeLimit** | Pointer to **int64** | The maximum number of resources that may be returned from a search request. | [optional] 
**DelegatedAdminReportSizeLimit** | Pointer to **int64** | The maximum number of resources that may be included in a report. | [optional] 
**MembersColumnName** | Pointer to **string** | Specifies the name of the group member column that will be displayed in the Delegated Admin UI | [optional] 
**NonmembersColumnName** | Pointer to **string** | Specifies the name of the group nonmember column that will be displayed in the Delegated Admin UI | [optional] 

## Methods

### NewAddGenericRestResourceTypeRequest

`func NewAddGenericRestResourceTypeRequest(typeName string, schemas []EnumgenericRestResourceTypeSchemaUrn, enabled bool, resourceEndpoint string, structuralLDAPObjectclass string, searchBaseDN string, ) *AddGenericRestResourceTypeRequest`

NewAddGenericRestResourceTypeRequest instantiates a new AddGenericRestResourceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGenericRestResourceTypeRequestWithDefaults

`func NewAddGenericRestResourceTypeRequestWithDefaults() *AddGenericRestResourceTypeRequest`

NewAddGenericRestResourceTypeRequestWithDefaults instantiates a new AddGenericRestResourceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *AddGenericRestResourceTypeRequest) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *AddGenericRestResourceTypeRequest) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *AddGenericRestResourceTypeRequest) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetSchemas

`func (o *AddGenericRestResourceTypeRequest) GetSchemas() []EnumgenericRestResourceTypeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddGenericRestResourceTypeRequest) GetSchemasOk() (*[]EnumgenericRestResourceTypeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddGenericRestResourceTypeRequest) SetSchemas(v []EnumgenericRestResourceTypeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddGenericRestResourceTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddGenericRestResourceTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddGenericRestResourceTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddGenericRestResourceTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddGenericRestResourceTypeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddGenericRestResourceTypeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddGenericRestResourceTypeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetResourceEndpoint

`func (o *AddGenericRestResourceTypeRequest) GetResourceEndpoint() string`

GetResourceEndpoint returns the ResourceEndpoint field if non-nil, zero value otherwise.

### GetResourceEndpointOk

`func (o *AddGenericRestResourceTypeRequest) GetResourceEndpointOk() (*string, bool)`

GetResourceEndpointOk returns a tuple with the ResourceEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceEndpoint

`func (o *AddGenericRestResourceTypeRequest) SetResourceEndpoint(v string)`

SetResourceEndpoint sets ResourceEndpoint field to given value.


### GetStructuralLDAPObjectclass

`func (o *AddGenericRestResourceTypeRequest) GetStructuralLDAPObjectclass() string`

GetStructuralLDAPObjectclass returns the StructuralLDAPObjectclass field if non-nil, zero value otherwise.

### GetStructuralLDAPObjectclassOk

`func (o *AddGenericRestResourceTypeRequest) GetStructuralLDAPObjectclassOk() (*string, bool)`

GetStructuralLDAPObjectclassOk returns a tuple with the StructuralLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuralLDAPObjectclass

`func (o *AddGenericRestResourceTypeRequest) SetStructuralLDAPObjectclass(v string)`

SetStructuralLDAPObjectclass sets StructuralLDAPObjectclass field to given value.


### GetAuxiliaryLDAPObjectclass

`func (o *AddGenericRestResourceTypeRequest) GetAuxiliaryLDAPObjectclass() []string`

GetAuxiliaryLDAPObjectclass returns the AuxiliaryLDAPObjectclass field if non-nil, zero value otherwise.

### GetAuxiliaryLDAPObjectclassOk

`func (o *AddGenericRestResourceTypeRequest) GetAuxiliaryLDAPObjectclassOk() (*[]string, bool)`

GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryLDAPObjectclass

`func (o *AddGenericRestResourceTypeRequest) SetAuxiliaryLDAPObjectclass(v []string)`

SetAuxiliaryLDAPObjectclass sets AuxiliaryLDAPObjectclass field to given value.

### HasAuxiliaryLDAPObjectclass

`func (o *AddGenericRestResourceTypeRequest) HasAuxiliaryLDAPObjectclass() bool`

HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.

### GetSearchBaseDN

`func (o *AddGenericRestResourceTypeRequest) GetSearchBaseDN() string`

GetSearchBaseDN returns the SearchBaseDN field if non-nil, zero value otherwise.

### GetSearchBaseDNOk

`func (o *AddGenericRestResourceTypeRequest) GetSearchBaseDNOk() (*string, bool)`

GetSearchBaseDNOk returns a tuple with the SearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBaseDN

`func (o *AddGenericRestResourceTypeRequest) SetSearchBaseDN(v string)`

SetSearchBaseDN sets SearchBaseDN field to given value.


### GetIncludeFilter

`func (o *AddGenericRestResourceTypeRequest) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *AddGenericRestResourceTypeRequest) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *AddGenericRestResourceTypeRequest) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *AddGenericRestResourceTypeRequest) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetParentDN

`func (o *AddGenericRestResourceTypeRequest) GetParentDN() string`

GetParentDN returns the ParentDN field if non-nil, zero value otherwise.

### GetParentDNOk

`func (o *AddGenericRestResourceTypeRequest) GetParentDNOk() (*string, bool)`

GetParentDNOk returns a tuple with the ParentDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDN

`func (o *AddGenericRestResourceTypeRequest) SetParentDN(v string)`

SetParentDN sets ParentDN field to given value.

### HasParentDN

`func (o *AddGenericRestResourceTypeRequest) HasParentDN() bool`

HasParentDN returns a boolean if a field has been set.

### GetParentResourceType

`func (o *AddGenericRestResourceTypeRequest) GetParentResourceType() string`

GetParentResourceType returns the ParentResourceType field if non-nil, zero value otherwise.

### GetParentResourceTypeOk

`func (o *AddGenericRestResourceTypeRequest) GetParentResourceTypeOk() (*string, bool)`

GetParentResourceTypeOk returns a tuple with the ParentResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceType

`func (o *AddGenericRestResourceTypeRequest) SetParentResourceType(v string)`

SetParentResourceType sets ParentResourceType field to given value.

### HasParentResourceType

`func (o *AddGenericRestResourceTypeRequest) HasParentResourceType() bool`

HasParentResourceType returns a boolean if a field has been set.

### GetRelativeDNFromParentResource

`func (o *AddGenericRestResourceTypeRequest) GetRelativeDNFromParentResource() string`

GetRelativeDNFromParentResource returns the RelativeDNFromParentResource field if non-nil, zero value otherwise.

### GetRelativeDNFromParentResourceOk

`func (o *AddGenericRestResourceTypeRequest) GetRelativeDNFromParentResourceOk() (*string, bool)`

GetRelativeDNFromParentResourceOk returns a tuple with the RelativeDNFromParentResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeDNFromParentResource

`func (o *AddGenericRestResourceTypeRequest) SetRelativeDNFromParentResource(v string)`

SetRelativeDNFromParentResource sets RelativeDNFromParentResource field to given value.

### HasRelativeDNFromParentResource

`func (o *AddGenericRestResourceTypeRequest) HasRelativeDNFromParentResource() bool`

HasRelativeDNFromParentResource returns a boolean if a field has been set.

### GetCreateRDNAttributeType

`func (o *AddGenericRestResourceTypeRequest) GetCreateRDNAttributeType() string`

GetCreateRDNAttributeType returns the CreateRDNAttributeType field if non-nil, zero value otherwise.

### GetCreateRDNAttributeTypeOk

`func (o *AddGenericRestResourceTypeRequest) GetCreateRDNAttributeTypeOk() (*string, bool)`

GetCreateRDNAttributeTypeOk returns a tuple with the CreateRDNAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateRDNAttributeType

`func (o *AddGenericRestResourceTypeRequest) SetCreateRDNAttributeType(v string)`

SetCreateRDNAttributeType sets CreateRDNAttributeType field to given value.

### HasCreateRDNAttributeType

`func (o *AddGenericRestResourceTypeRequest) HasCreateRDNAttributeType() bool`

HasCreateRDNAttributeType returns a boolean if a field has been set.

### GetPostCreateConstructedAttribute

`func (o *AddGenericRestResourceTypeRequest) GetPostCreateConstructedAttribute() []string`

GetPostCreateConstructedAttribute returns the PostCreateConstructedAttribute field if non-nil, zero value otherwise.

### GetPostCreateConstructedAttributeOk

`func (o *AddGenericRestResourceTypeRequest) GetPostCreateConstructedAttributeOk() (*[]string, bool)`

GetPostCreateConstructedAttributeOk returns a tuple with the PostCreateConstructedAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCreateConstructedAttribute

`func (o *AddGenericRestResourceTypeRequest) SetPostCreateConstructedAttribute(v []string)`

SetPostCreateConstructedAttribute sets PostCreateConstructedAttribute field to given value.

### HasPostCreateConstructedAttribute

`func (o *AddGenericRestResourceTypeRequest) HasPostCreateConstructedAttribute() bool`

HasPostCreateConstructedAttribute returns a boolean if a field has been set.

### GetUpdateConstructedAttribute

`func (o *AddGenericRestResourceTypeRequest) GetUpdateConstructedAttribute() []string`

GetUpdateConstructedAttribute returns the UpdateConstructedAttribute field if non-nil, zero value otherwise.

### GetUpdateConstructedAttributeOk

`func (o *AddGenericRestResourceTypeRequest) GetUpdateConstructedAttributeOk() (*[]string, bool)`

GetUpdateConstructedAttributeOk returns a tuple with the UpdateConstructedAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateConstructedAttribute

`func (o *AddGenericRestResourceTypeRequest) SetUpdateConstructedAttribute(v []string)`

SetUpdateConstructedAttribute sets UpdateConstructedAttribute field to given value.

### HasUpdateConstructedAttribute

`func (o *AddGenericRestResourceTypeRequest) HasUpdateConstructedAttribute() bool`

HasUpdateConstructedAttribute returns a boolean if a field has been set.

### GetDisplayName

`func (o *AddGenericRestResourceTypeRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddGenericRestResourceTypeRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddGenericRestResourceTypeRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AddGenericRestResourceTypeRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSearchFilterPattern

`func (o *AddGenericRestResourceTypeRequest) GetSearchFilterPattern() string`

GetSearchFilterPattern returns the SearchFilterPattern field if non-nil, zero value otherwise.

### GetSearchFilterPatternOk

`func (o *AddGenericRestResourceTypeRequest) GetSearchFilterPatternOk() (*string, bool)`

GetSearchFilterPatternOk returns a tuple with the SearchFilterPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterPattern

`func (o *AddGenericRestResourceTypeRequest) SetSearchFilterPattern(v string)`

SetSearchFilterPattern sets SearchFilterPattern field to given value.

### HasSearchFilterPattern

`func (o *AddGenericRestResourceTypeRequest) HasSearchFilterPattern() bool`

HasSearchFilterPattern returns a boolean if a field has been set.

### GetPrimaryDisplayAttributeType

`func (o *AddGenericRestResourceTypeRequest) GetPrimaryDisplayAttributeType() string`

GetPrimaryDisplayAttributeType returns the PrimaryDisplayAttributeType field if non-nil, zero value otherwise.

### GetPrimaryDisplayAttributeTypeOk

`func (o *AddGenericRestResourceTypeRequest) GetPrimaryDisplayAttributeTypeOk() (*string, bool)`

GetPrimaryDisplayAttributeTypeOk returns a tuple with the PrimaryDisplayAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDisplayAttributeType

`func (o *AddGenericRestResourceTypeRequest) SetPrimaryDisplayAttributeType(v string)`

SetPrimaryDisplayAttributeType sets PrimaryDisplayAttributeType field to given value.

### HasPrimaryDisplayAttributeType

`func (o *AddGenericRestResourceTypeRequest) HasPrimaryDisplayAttributeType() bool`

HasPrimaryDisplayAttributeType returns a boolean if a field has been set.

### GetDelegatedAdminSearchSizeLimit

`func (o *AddGenericRestResourceTypeRequest) GetDelegatedAdminSearchSizeLimit() int64`

GetDelegatedAdminSearchSizeLimit returns the DelegatedAdminSearchSizeLimit field if non-nil, zero value otherwise.

### GetDelegatedAdminSearchSizeLimitOk

`func (o *AddGenericRestResourceTypeRequest) GetDelegatedAdminSearchSizeLimitOk() (*int64, bool)`

GetDelegatedAdminSearchSizeLimitOk returns a tuple with the DelegatedAdminSearchSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAdminSearchSizeLimit

`func (o *AddGenericRestResourceTypeRequest) SetDelegatedAdminSearchSizeLimit(v int64)`

SetDelegatedAdminSearchSizeLimit sets DelegatedAdminSearchSizeLimit field to given value.

### HasDelegatedAdminSearchSizeLimit

`func (o *AddGenericRestResourceTypeRequest) HasDelegatedAdminSearchSizeLimit() bool`

HasDelegatedAdminSearchSizeLimit returns a boolean if a field has been set.

### GetDelegatedAdminReportSizeLimit

`func (o *AddGenericRestResourceTypeRequest) GetDelegatedAdminReportSizeLimit() int64`

GetDelegatedAdminReportSizeLimit returns the DelegatedAdminReportSizeLimit field if non-nil, zero value otherwise.

### GetDelegatedAdminReportSizeLimitOk

`func (o *AddGenericRestResourceTypeRequest) GetDelegatedAdminReportSizeLimitOk() (*int64, bool)`

GetDelegatedAdminReportSizeLimitOk returns a tuple with the DelegatedAdminReportSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAdminReportSizeLimit

`func (o *AddGenericRestResourceTypeRequest) SetDelegatedAdminReportSizeLimit(v int64)`

SetDelegatedAdminReportSizeLimit sets DelegatedAdminReportSizeLimit field to given value.

### HasDelegatedAdminReportSizeLimit

`func (o *AddGenericRestResourceTypeRequest) HasDelegatedAdminReportSizeLimit() bool`

HasDelegatedAdminReportSizeLimit returns a boolean if a field has been set.

### GetMembersColumnName

`func (o *AddGenericRestResourceTypeRequest) GetMembersColumnName() string`

GetMembersColumnName returns the MembersColumnName field if non-nil, zero value otherwise.

### GetMembersColumnNameOk

`func (o *AddGenericRestResourceTypeRequest) GetMembersColumnNameOk() (*string, bool)`

GetMembersColumnNameOk returns a tuple with the MembersColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersColumnName

`func (o *AddGenericRestResourceTypeRequest) SetMembersColumnName(v string)`

SetMembersColumnName sets MembersColumnName field to given value.

### HasMembersColumnName

`func (o *AddGenericRestResourceTypeRequest) HasMembersColumnName() bool`

HasMembersColumnName returns a boolean if a field has been set.

### GetNonmembersColumnName

`func (o *AddGenericRestResourceTypeRequest) GetNonmembersColumnName() string`

GetNonmembersColumnName returns the NonmembersColumnName field if non-nil, zero value otherwise.

### GetNonmembersColumnNameOk

`func (o *AddGenericRestResourceTypeRequest) GetNonmembersColumnNameOk() (*string, bool)`

GetNonmembersColumnNameOk returns a tuple with the NonmembersColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonmembersColumnName

`func (o *AddGenericRestResourceTypeRequest) SetNonmembersColumnName(v string)`

SetNonmembersColumnName sets NonmembersColumnName field to given value.

### HasNonmembersColumnName

`func (o *AddGenericRestResourceTypeRequest) HasNonmembersColumnName() bool`

HasNonmembersColumnName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


