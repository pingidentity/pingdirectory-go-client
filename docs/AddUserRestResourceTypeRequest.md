# AddUserRestResourceTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | **string** | Name of the new REST Resource Type | 
**Schemas** | [**[]EnumuserRestResourceTypeSchemaUrn**](EnumuserRestResourceTypeSchemaUrn.md) |  | 
**PasswordAttributeCategory** | Pointer to **string** | Specifies which attribute category the password belongs to. | [optional] 
**PasswordDisplayOrderIndex** | Pointer to **int32** | This property determines the display order for the password within its attribute category. Attributes are ordered within their category based on this index from least to greatest. | [optional] 
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
**MembersColumnName** | Pointer to **string** | Specifies the name of the group member column that will be displayed in the Delegated Admin UI | [optional] 
**NonmembersColumnName** | Pointer to **string** | Specifies the name of the group nonmember column that will be displayed in the Delegated Admin UI | [optional] 

## Methods

### NewAddUserRestResourceTypeRequest

`func NewAddUserRestResourceTypeRequest(typeName string, schemas []EnumuserRestResourceTypeSchemaUrn, enabled bool, resourceEndpoint string, structuralLDAPObjectclass string, searchBaseDN string, ) *AddUserRestResourceTypeRequest`

NewAddUserRestResourceTypeRequest instantiates a new AddUserRestResourceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserRestResourceTypeRequestWithDefaults

`func NewAddUserRestResourceTypeRequestWithDefaults() *AddUserRestResourceTypeRequest`

NewAddUserRestResourceTypeRequestWithDefaults instantiates a new AddUserRestResourceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *AddUserRestResourceTypeRequest) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *AddUserRestResourceTypeRequest) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *AddUserRestResourceTypeRequest) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetSchemas

`func (o *AddUserRestResourceTypeRequest) GetSchemas() []EnumuserRestResourceTypeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddUserRestResourceTypeRequest) GetSchemasOk() (*[]EnumuserRestResourceTypeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddUserRestResourceTypeRequest) SetSchemas(v []EnumuserRestResourceTypeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordAttributeCategory

`func (o *AddUserRestResourceTypeRequest) GetPasswordAttributeCategory() string`

GetPasswordAttributeCategory returns the PasswordAttributeCategory field if non-nil, zero value otherwise.

### GetPasswordAttributeCategoryOk

`func (o *AddUserRestResourceTypeRequest) GetPasswordAttributeCategoryOk() (*string, bool)`

GetPasswordAttributeCategoryOk returns a tuple with the PasswordAttributeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordAttributeCategory

`func (o *AddUserRestResourceTypeRequest) SetPasswordAttributeCategory(v string)`

SetPasswordAttributeCategory sets PasswordAttributeCategory field to given value.

### HasPasswordAttributeCategory

`func (o *AddUserRestResourceTypeRequest) HasPasswordAttributeCategory() bool`

HasPasswordAttributeCategory returns a boolean if a field has been set.

### GetPasswordDisplayOrderIndex

`func (o *AddUserRestResourceTypeRequest) GetPasswordDisplayOrderIndex() int32`

GetPasswordDisplayOrderIndex returns the PasswordDisplayOrderIndex field if non-nil, zero value otherwise.

### GetPasswordDisplayOrderIndexOk

`func (o *AddUserRestResourceTypeRequest) GetPasswordDisplayOrderIndexOk() (*int32, bool)`

GetPasswordDisplayOrderIndexOk returns a tuple with the PasswordDisplayOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordDisplayOrderIndex

`func (o *AddUserRestResourceTypeRequest) SetPasswordDisplayOrderIndex(v int32)`

SetPasswordDisplayOrderIndex sets PasswordDisplayOrderIndex field to given value.

### HasPasswordDisplayOrderIndex

`func (o *AddUserRestResourceTypeRequest) HasPasswordDisplayOrderIndex() bool`

HasPasswordDisplayOrderIndex returns a boolean if a field has been set.

### GetDescription

`func (o *AddUserRestResourceTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddUserRestResourceTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddUserRestResourceTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddUserRestResourceTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddUserRestResourceTypeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddUserRestResourceTypeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddUserRestResourceTypeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetResourceEndpoint

`func (o *AddUserRestResourceTypeRequest) GetResourceEndpoint() string`

GetResourceEndpoint returns the ResourceEndpoint field if non-nil, zero value otherwise.

### GetResourceEndpointOk

`func (o *AddUserRestResourceTypeRequest) GetResourceEndpointOk() (*string, bool)`

GetResourceEndpointOk returns a tuple with the ResourceEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceEndpoint

`func (o *AddUserRestResourceTypeRequest) SetResourceEndpoint(v string)`

SetResourceEndpoint sets ResourceEndpoint field to given value.


### GetStructuralLDAPObjectclass

`func (o *AddUserRestResourceTypeRequest) GetStructuralLDAPObjectclass() string`

GetStructuralLDAPObjectclass returns the StructuralLDAPObjectclass field if non-nil, zero value otherwise.

### GetStructuralLDAPObjectclassOk

`func (o *AddUserRestResourceTypeRequest) GetStructuralLDAPObjectclassOk() (*string, bool)`

GetStructuralLDAPObjectclassOk returns a tuple with the StructuralLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuralLDAPObjectclass

`func (o *AddUserRestResourceTypeRequest) SetStructuralLDAPObjectclass(v string)`

SetStructuralLDAPObjectclass sets StructuralLDAPObjectclass field to given value.


### GetAuxiliaryLDAPObjectclass

`func (o *AddUserRestResourceTypeRequest) GetAuxiliaryLDAPObjectclass() []string`

GetAuxiliaryLDAPObjectclass returns the AuxiliaryLDAPObjectclass field if non-nil, zero value otherwise.

### GetAuxiliaryLDAPObjectclassOk

`func (o *AddUserRestResourceTypeRequest) GetAuxiliaryLDAPObjectclassOk() (*[]string, bool)`

GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryLDAPObjectclass

`func (o *AddUserRestResourceTypeRequest) SetAuxiliaryLDAPObjectclass(v []string)`

SetAuxiliaryLDAPObjectclass sets AuxiliaryLDAPObjectclass field to given value.

### HasAuxiliaryLDAPObjectclass

`func (o *AddUserRestResourceTypeRequest) HasAuxiliaryLDAPObjectclass() bool`

HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.

### GetSearchBaseDN

`func (o *AddUserRestResourceTypeRequest) GetSearchBaseDN() string`

GetSearchBaseDN returns the SearchBaseDN field if non-nil, zero value otherwise.

### GetSearchBaseDNOk

`func (o *AddUserRestResourceTypeRequest) GetSearchBaseDNOk() (*string, bool)`

GetSearchBaseDNOk returns a tuple with the SearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBaseDN

`func (o *AddUserRestResourceTypeRequest) SetSearchBaseDN(v string)`

SetSearchBaseDN sets SearchBaseDN field to given value.


### GetIncludeFilter

`func (o *AddUserRestResourceTypeRequest) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *AddUserRestResourceTypeRequest) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *AddUserRestResourceTypeRequest) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *AddUserRestResourceTypeRequest) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetParentDN

`func (o *AddUserRestResourceTypeRequest) GetParentDN() string`

GetParentDN returns the ParentDN field if non-nil, zero value otherwise.

### GetParentDNOk

`func (o *AddUserRestResourceTypeRequest) GetParentDNOk() (*string, bool)`

GetParentDNOk returns a tuple with the ParentDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDN

`func (o *AddUserRestResourceTypeRequest) SetParentDN(v string)`

SetParentDN sets ParentDN field to given value.

### HasParentDN

`func (o *AddUserRestResourceTypeRequest) HasParentDN() bool`

HasParentDN returns a boolean if a field has been set.

### GetParentResourceType

`func (o *AddUserRestResourceTypeRequest) GetParentResourceType() string`

GetParentResourceType returns the ParentResourceType field if non-nil, zero value otherwise.

### GetParentResourceTypeOk

`func (o *AddUserRestResourceTypeRequest) GetParentResourceTypeOk() (*string, bool)`

GetParentResourceTypeOk returns a tuple with the ParentResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceType

`func (o *AddUserRestResourceTypeRequest) SetParentResourceType(v string)`

SetParentResourceType sets ParentResourceType field to given value.

### HasParentResourceType

`func (o *AddUserRestResourceTypeRequest) HasParentResourceType() bool`

HasParentResourceType returns a boolean if a field has been set.

### GetRelativeDNFromParentResource

`func (o *AddUserRestResourceTypeRequest) GetRelativeDNFromParentResource() string`

GetRelativeDNFromParentResource returns the RelativeDNFromParentResource field if non-nil, zero value otherwise.

### GetRelativeDNFromParentResourceOk

`func (o *AddUserRestResourceTypeRequest) GetRelativeDNFromParentResourceOk() (*string, bool)`

GetRelativeDNFromParentResourceOk returns a tuple with the RelativeDNFromParentResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeDNFromParentResource

`func (o *AddUserRestResourceTypeRequest) SetRelativeDNFromParentResource(v string)`

SetRelativeDNFromParentResource sets RelativeDNFromParentResource field to given value.

### HasRelativeDNFromParentResource

`func (o *AddUserRestResourceTypeRequest) HasRelativeDNFromParentResource() bool`

HasRelativeDNFromParentResource returns a boolean if a field has been set.

### GetCreateRDNAttributeType

`func (o *AddUserRestResourceTypeRequest) GetCreateRDNAttributeType() string`

GetCreateRDNAttributeType returns the CreateRDNAttributeType field if non-nil, zero value otherwise.

### GetCreateRDNAttributeTypeOk

`func (o *AddUserRestResourceTypeRequest) GetCreateRDNAttributeTypeOk() (*string, bool)`

GetCreateRDNAttributeTypeOk returns a tuple with the CreateRDNAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateRDNAttributeType

`func (o *AddUserRestResourceTypeRequest) SetCreateRDNAttributeType(v string)`

SetCreateRDNAttributeType sets CreateRDNAttributeType field to given value.

### HasCreateRDNAttributeType

`func (o *AddUserRestResourceTypeRequest) HasCreateRDNAttributeType() bool`

HasCreateRDNAttributeType returns a boolean if a field has been set.

### GetPostCreateConstructedAttribute

`func (o *AddUserRestResourceTypeRequest) GetPostCreateConstructedAttribute() []string`

GetPostCreateConstructedAttribute returns the PostCreateConstructedAttribute field if non-nil, zero value otherwise.

### GetPostCreateConstructedAttributeOk

`func (o *AddUserRestResourceTypeRequest) GetPostCreateConstructedAttributeOk() (*[]string, bool)`

GetPostCreateConstructedAttributeOk returns a tuple with the PostCreateConstructedAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCreateConstructedAttribute

`func (o *AddUserRestResourceTypeRequest) SetPostCreateConstructedAttribute(v []string)`

SetPostCreateConstructedAttribute sets PostCreateConstructedAttribute field to given value.

### HasPostCreateConstructedAttribute

`func (o *AddUserRestResourceTypeRequest) HasPostCreateConstructedAttribute() bool`

HasPostCreateConstructedAttribute returns a boolean if a field has been set.

### GetUpdateConstructedAttribute

`func (o *AddUserRestResourceTypeRequest) GetUpdateConstructedAttribute() []string`

GetUpdateConstructedAttribute returns the UpdateConstructedAttribute field if non-nil, zero value otherwise.

### GetUpdateConstructedAttributeOk

`func (o *AddUserRestResourceTypeRequest) GetUpdateConstructedAttributeOk() (*[]string, bool)`

GetUpdateConstructedAttributeOk returns a tuple with the UpdateConstructedAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateConstructedAttribute

`func (o *AddUserRestResourceTypeRequest) SetUpdateConstructedAttribute(v []string)`

SetUpdateConstructedAttribute sets UpdateConstructedAttribute field to given value.

### HasUpdateConstructedAttribute

`func (o *AddUserRestResourceTypeRequest) HasUpdateConstructedAttribute() bool`

HasUpdateConstructedAttribute returns a boolean if a field has been set.

### GetDisplayName

`func (o *AddUserRestResourceTypeRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddUserRestResourceTypeRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddUserRestResourceTypeRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AddUserRestResourceTypeRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSearchFilterPattern

`func (o *AddUserRestResourceTypeRequest) GetSearchFilterPattern() string`

GetSearchFilterPattern returns the SearchFilterPattern field if non-nil, zero value otherwise.

### GetSearchFilterPatternOk

`func (o *AddUserRestResourceTypeRequest) GetSearchFilterPatternOk() (*string, bool)`

GetSearchFilterPatternOk returns a tuple with the SearchFilterPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterPattern

`func (o *AddUserRestResourceTypeRequest) SetSearchFilterPattern(v string)`

SetSearchFilterPattern sets SearchFilterPattern field to given value.

### HasSearchFilterPattern

`func (o *AddUserRestResourceTypeRequest) HasSearchFilterPattern() bool`

HasSearchFilterPattern returns a boolean if a field has been set.

### GetPrimaryDisplayAttributeType

`func (o *AddUserRestResourceTypeRequest) GetPrimaryDisplayAttributeType() string`

GetPrimaryDisplayAttributeType returns the PrimaryDisplayAttributeType field if non-nil, zero value otherwise.

### GetPrimaryDisplayAttributeTypeOk

`func (o *AddUserRestResourceTypeRequest) GetPrimaryDisplayAttributeTypeOk() (*string, bool)`

GetPrimaryDisplayAttributeTypeOk returns a tuple with the PrimaryDisplayAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDisplayAttributeType

`func (o *AddUserRestResourceTypeRequest) SetPrimaryDisplayAttributeType(v string)`

SetPrimaryDisplayAttributeType sets PrimaryDisplayAttributeType field to given value.

### HasPrimaryDisplayAttributeType

`func (o *AddUserRestResourceTypeRequest) HasPrimaryDisplayAttributeType() bool`

HasPrimaryDisplayAttributeType returns a boolean if a field has been set.

### GetDelegatedAdminSearchSizeLimit

`func (o *AddUserRestResourceTypeRequest) GetDelegatedAdminSearchSizeLimit() int32`

GetDelegatedAdminSearchSizeLimit returns the DelegatedAdminSearchSizeLimit field if non-nil, zero value otherwise.

### GetDelegatedAdminSearchSizeLimitOk

`func (o *AddUserRestResourceTypeRequest) GetDelegatedAdminSearchSizeLimitOk() (*int32, bool)`

GetDelegatedAdminSearchSizeLimitOk returns a tuple with the DelegatedAdminSearchSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAdminSearchSizeLimit

`func (o *AddUserRestResourceTypeRequest) SetDelegatedAdminSearchSizeLimit(v int32)`

SetDelegatedAdminSearchSizeLimit sets DelegatedAdminSearchSizeLimit field to given value.

### HasDelegatedAdminSearchSizeLimit

`func (o *AddUserRestResourceTypeRequest) HasDelegatedAdminSearchSizeLimit() bool`

HasDelegatedAdminSearchSizeLimit returns a boolean if a field has been set.

### GetDelegatedAdminReportSizeLimit

`func (o *AddUserRestResourceTypeRequest) GetDelegatedAdminReportSizeLimit() int32`

GetDelegatedAdminReportSizeLimit returns the DelegatedAdminReportSizeLimit field if non-nil, zero value otherwise.

### GetDelegatedAdminReportSizeLimitOk

`func (o *AddUserRestResourceTypeRequest) GetDelegatedAdminReportSizeLimitOk() (*int32, bool)`

GetDelegatedAdminReportSizeLimitOk returns a tuple with the DelegatedAdminReportSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAdminReportSizeLimit

`func (o *AddUserRestResourceTypeRequest) SetDelegatedAdminReportSizeLimit(v int32)`

SetDelegatedAdminReportSizeLimit sets DelegatedAdminReportSizeLimit field to given value.

### HasDelegatedAdminReportSizeLimit

`func (o *AddUserRestResourceTypeRequest) HasDelegatedAdminReportSizeLimit() bool`

HasDelegatedAdminReportSizeLimit returns a boolean if a field has been set.

### GetMembersColumnName

`func (o *AddUserRestResourceTypeRequest) GetMembersColumnName() string`

GetMembersColumnName returns the MembersColumnName field if non-nil, zero value otherwise.

### GetMembersColumnNameOk

`func (o *AddUserRestResourceTypeRequest) GetMembersColumnNameOk() (*string, bool)`

GetMembersColumnNameOk returns a tuple with the MembersColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersColumnName

`func (o *AddUserRestResourceTypeRequest) SetMembersColumnName(v string)`

SetMembersColumnName sets MembersColumnName field to given value.

### HasMembersColumnName

`func (o *AddUserRestResourceTypeRequest) HasMembersColumnName() bool`

HasMembersColumnName returns a boolean if a field has been set.

### GetNonmembersColumnName

`func (o *AddUserRestResourceTypeRequest) GetNonmembersColumnName() string`

GetNonmembersColumnName returns the NonmembersColumnName field if non-nil, zero value otherwise.

### GetNonmembersColumnNameOk

`func (o *AddUserRestResourceTypeRequest) GetNonmembersColumnNameOk() (*string, bool)`

GetNonmembersColumnNameOk returns a tuple with the NonmembersColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonmembersColumnName

`func (o *AddUserRestResourceTypeRequest) SetNonmembersColumnName(v string)`

SetNonmembersColumnName sets NonmembersColumnName field to given value.

### HasNonmembersColumnName

`func (o *AddUserRestResourceTypeRequest) HasNonmembersColumnName() bool`

HasNonmembersColumnName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


