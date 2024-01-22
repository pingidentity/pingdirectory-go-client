# AddLdapMappingScimResourceTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumldapMappingScimResourceTypeSchemaUrn**](EnumldapMappingScimResourceTypeSchemaUrn.md) |  | 
**CoreSchema** | **string** | The core schema enforced on core attributes at the top level of a SCIM resource representation exposed by thisMapping SCIM Resource Type. | 
**RequiredSchemaExtension** | Pointer to **[]string** | Required additive schemas that are enforced on extension attributes in a SCIM resource representation for this Mapping SCIM Resource Type. | [optional] 
**OptionalSchemaExtension** | Pointer to **[]string** | Optional additive schemas that are enforced on extension attributes in a SCIM resource representation for this Mapping SCIM Resource Type. | [optional] 
**Description** | Pointer to **string** | A description for this SCIM Resource Type | [optional] 
**Enabled** | **bool** | Indicates whether the SCIM Resource Type is enabled. | 
**Endpoint** | **string** | The HTTP addressable endpoint of this SCIM Resource Type relative to the &#39;/scim/v2&#39; base URL. Do not include a leading &#39;/&#39;. | 
**LookthroughLimit** | Pointer to **int64** | The maximum number of resources that the SCIM Resource Type should \&quot;look through\&quot; in the course of processing a search request. | [optional] 
**SchemaCheckingOption** | Pointer to [**[]EnumscimResourceTypeSchemaCheckingOptionProp**](EnumscimResourceTypeSchemaCheckingOptionProp.md) |  | [optional] 
**StructuralLDAPObjectclass** | Pointer to **string** | Specifies the LDAP structural object class that should be exposed by this SCIM Resource Type. | [optional] 
**AuxiliaryLDAPObjectclass** | Pointer to **[]string** | Specifies an auxiliary LDAP object class that should be exposed by this SCIM Resource Type. | [optional] 
**IncludeBaseDN** | Pointer to **string** | Specifies the base DN of the branch of the LDAP directory that can be accessed by this SCIM Resource Type. | [optional] 
**IncludeFilter** | Pointer to **[]string** | The set of LDAP filters that define the LDAP entries that should be included in this SCIM Resource Type. | [optional] 
**IncludeOperationalAttribute** | Pointer to **[]string** | Specifies the set of operational LDAP attributes to be provided by this SCIM Resource Type. | [optional] 
**CreateDNPattern** | Pointer to **string** | Specifies the template to use for the DN when creating new entries. | [optional] 
**TypeName** | **string** | Name of the new SCIM Resource Type | 

## Methods

### NewAddLdapMappingScimResourceTypeRequest

`func NewAddLdapMappingScimResourceTypeRequest(schemas []EnumldapMappingScimResourceTypeSchemaUrn, coreSchema string, enabled bool, endpoint string, typeName string, ) *AddLdapMappingScimResourceTypeRequest`

NewAddLdapMappingScimResourceTypeRequest instantiates a new AddLdapMappingScimResourceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLdapMappingScimResourceTypeRequestWithDefaults

`func NewAddLdapMappingScimResourceTypeRequestWithDefaults() *AddLdapMappingScimResourceTypeRequest`

NewAddLdapMappingScimResourceTypeRequestWithDefaults instantiates a new AddLdapMappingScimResourceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddLdapMappingScimResourceTypeRequest) GetSchemas() []EnumldapMappingScimResourceTypeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLdapMappingScimResourceTypeRequest) GetSchemasOk() (*[]EnumldapMappingScimResourceTypeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLdapMappingScimResourceTypeRequest) SetSchemas(v []EnumldapMappingScimResourceTypeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetCoreSchema

`func (o *AddLdapMappingScimResourceTypeRequest) GetCoreSchema() string`

GetCoreSchema returns the CoreSchema field if non-nil, zero value otherwise.

### GetCoreSchemaOk

`func (o *AddLdapMappingScimResourceTypeRequest) GetCoreSchemaOk() (*string, bool)`

GetCoreSchemaOk returns a tuple with the CoreSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreSchema

`func (o *AddLdapMappingScimResourceTypeRequest) SetCoreSchema(v string)`

SetCoreSchema sets CoreSchema field to given value.


### GetRequiredSchemaExtension

`func (o *AddLdapMappingScimResourceTypeRequest) GetRequiredSchemaExtension() []string`

GetRequiredSchemaExtension returns the RequiredSchemaExtension field if non-nil, zero value otherwise.

### GetRequiredSchemaExtensionOk

`func (o *AddLdapMappingScimResourceTypeRequest) GetRequiredSchemaExtensionOk() (*[]string, bool)`

GetRequiredSchemaExtensionOk returns a tuple with the RequiredSchemaExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSchemaExtension

`func (o *AddLdapMappingScimResourceTypeRequest) SetRequiredSchemaExtension(v []string)`

SetRequiredSchemaExtension sets RequiredSchemaExtension field to given value.

### HasRequiredSchemaExtension

`func (o *AddLdapMappingScimResourceTypeRequest) HasRequiredSchemaExtension() bool`

HasRequiredSchemaExtension returns a boolean if a field has been set.

### GetOptionalSchemaExtension

`func (o *AddLdapMappingScimResourceTypeRequest) GetOptionalSchemaExtension() []string`

GetOptionalSchemaExtension returns the OptionalSchemaExtension field if non-nil, zero value otherwise.

### GetOptionalSchemaExtensionOk

`func (o *AddLdapMappingScimResourceTypeRequest) GetOptionalSchemaExtensionOk() (*[]string, bool)`

GetOptionalSchemaExtensionOk returns a tuple with the OptionalSchemaExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalSchemaExtension

`func (o *AddLdapMappingScimResourceTypeRequest) SetOptionalSchemaExtension(v []string)`

SetOptionalSchemaExtension sets OptionalSchemaExtension field to given value.

### HasOptionalSchemaExtension

`func (o *AddLdapMappingScimResourceTypeRequest) HasOptionalSchemaExtension() bool`

HasOptionalSchemaExtension returns a boolean if a field has been set.

### GetDescription

`func (o *AddLdapMappingScimResourceTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLdapMappingScimResourceTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLdapMappingScimResourceTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLdapMappingScimResourceTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddLdapMappingScimResourceTypeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddLdapMappingScimResourceTypeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddLdapMappingScimResourceTypeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEndpoint

`func (o *AddLdapMappingScimResourceTypeRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AddLdapMappingScimResourceTypeRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AddLdapMappingScimResourceTypeRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetLookthroughLimit

`func (o *AddLdapMappingScimResourceTypeRequest) GetLookthroughLimit() int64`

GetLookthroughLimit returns the LookthroughLimit field if non-nil, zero value otherwise.

### GetLookthroughLimitOk

`func (o *AddLdapMappingScimResourceTypeRequest) GetLookthroughLimitOk() (*int64, bool)`

GetLookthroughLimitOk returns a tuple with the LookthroughLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookthroughLimit

`func (o *AddLdapMappingScimResourceTypeRequest) SetLookthroughLimit(v int64)`

SetLookthroughLimit sets LookthroughLimit field to given value.

### HasLookthroughLimit

`func (o *AddLdapMappingScimResourceTypeRequest) HasLookthroughLimit() bool`

HasLookthroughLimit returns a boolean if a field has been set.

### GetSchemaCheckingOption

`func (o *AddLdapMappingScimResourceTypeRequest) GetSchemaCheckingOption() []EnumscimResourceTypeSchemaCheckingOptionProp`

GetSchemaCheckingOption returns the SchemaCheckingOption field if non-nil, zero value otherwise.

### GetSchemaCheckingOptionOk

`func (o *AddLdapMappingScimResourceTypeRequest) GetSchemaCheckingOptionOk() (*[]EnumscimResourceTypeSchemaCheckingOptionProp, bool)`

GetSchemaCheckingOptionOk returns a tuple with the SchemaCheckingOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaCheckingOption

`func (o *AddLdapMappingScimResourceTypeRequest) SetSchemaCheckingOption(v []EnumscimResourceTypeSchemaCheckingOptionProp)`

SetSchemaCheckingOption sets SchemaCheckingOption field to given value.

### HasSchemaCheckingOption

`func (o *AddLdapMappingScimResourceTypeRequest) HasSchemaCheckingOption() bool`

HasSchemaCheckingOption returns a boolean if a field has been set.

### GetStructuralLDAPObjectclass

`func (o *AddLdapMappingScimResourceTypeRequest) GetStructuralLDAPObjectclass() string`

GetStructuralLDAPObjectclass returns the StructuralLDAPObjectclass field if non-nil, zero value otherwise.

### GetStructuralLDAPObjectclassOk

`func (o *AddLdapMappingScimResourceTypeRequest) GetStructuralLDAPObjectclassOk() (*string, bool)`

GetStructuralLDAPObjectclassOk returns a tuple with the StructuralLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuralLDAPObjectclass

`func (o *AddLdapMappingScimResourceTypeRequest) SetStructuralLDAPObjectclass(v string)`

SetStructuralLDAPObjectclass sets StructuralLDAPObjectclass field to given value.

### HasStructuralLDAPObjectclass

`func (o *AddLdapMappingScimResourceTypeRequest) HasStructuralLDAPObjectclass() bool`

HasStructuralLDAPObjectclass returns a boolean if a field has been set.

### GetAuxiliaryLDAPObjectclass

`func (o *AddLdapMappingScimResourceTypeRequest) GetAuxiliaryLDAPObjectclass() []string`

GetAuxiliaryLDAPObjectclass returns the AuxiliaryLDAPObjectclass field if non-nil, zero value otherwise.

### GetAuxiliaryLDAPObjectclassOk

`func (o *AddLdapMappingScimResourceTypeRequest) GetAuxiliaryLDAPObjectclassOk() (*[]string, bool)`

GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryLDAPObjectclass

`func (o *AddLdapMappingScimResourceTypeRequest) SetAuxiliaryLDAPObjectclass(v []string)`

SetAuxiliaryLDAPObjectclass sets AuxiliaryLDAPObjectclass field to given value.

### HasAuxiliaryLDAPObjectclass

`func (o *AddLdapMappingScimResourceTypeRequest) HasAuxiliaryLDAPObjectclass() bool`

HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.

### GetIncludeBaseDN

`func (o *AddLdapMappingScimResourceTypeRequest) GetIncludeBaseDN() string`

GetIncludeBaseDN returns the IncludeBaseDN field if non-nil, zero value otherwise.

### GetIncludeBaseDNOk

`func (o *AddLdapMappingScimResourceTypeRequest) GetIncludeBaseDNOk() (*string, bool)`

GetIncludeBaseDNOk returns a tuple with the IncludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBaseDN

`func (o *AddLdapMappingScimResourceTypeRequest) SetIncludeBaseDN(v string)`

SetIncludeBaseDN sets IncludeBaseDN field to given value.

### HasIncludeBaseDN

`func (o *AddLdapMappingScimResourceTypeRequest) HasIncludeBaseDN() bool`

HasIncludeBaseDN returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *AddLdapMappingScimResourceTypeRequest) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *AddLdapMappingScimResourceTypeRequest) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *AddLdapMappingScimResourceTypeRequest) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *AddLdapMappingScimResourceTypeRequest) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetIncludeOperationalAttribute

`func (o *AddLdapMappingScimResourceTypeRequest) GetIncludeOperationalAttribute() []string`

GetIncludeOperationalAttribute returns the IncludeOperationalAttribute field if non-nil, zero value otherwise.

### GetIncludeOperationalAttributeOk

`func (o *AddLdapMappingScimResourceTypeRequest) GetIncludeOperationalAttributeOk() (*[]string, bool)`

GetIncludeOperationalAttributeOk returns a tuple with the IncludeOperationalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOperationalAttribute

`func (o *AddLdapMappingScimResourceTypeRequest) SetIncludeOperationalAttribute(v []string)`

SetIncludeOperationalAttribute sets IncludeOperationalAttribute field to given value.

### HasIncludeOperationalAttribute

`func (o *AddLdapMappingScimResourceTypeRequest) HasIncludeOperationalAttribute() bool`

HasIncludeOperationalAttribute returns a boolean if a field has been set.

### GetCreateDNPattern

`func (o *AddLdapMappingScimResourceTypeRequest) GetCreateDNPattern() string`

GetCreateDNPattern returns the CreateDNPattern field if non-nil, zero value otherwise.

### GetCreateDNPatternOk

`func (o *AddLdapMappingScimResourceTypeRequest) GetCreateDNPatternOk() (*string, bool)`

GetCreateDNPatternOk returns a tuple with the CreateDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDNPattern

`func (o *AddLdapMappingScimResourceTypeRequest) SetCreateDNPattern(v string)`

SetCreateDNPattern sets CreateDNPattern field to given value.

### HasCreateDNPattern

`func (o *AddLdapMappingScimResourceTypeRequest) HasCreateDNPattern() bool`

HasCreateDNPattern returns a boolean if a field has been set.

### GetTypeName

`func (o *AddLdapMappingScimResourceTypeRequest) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *AddLdapMappingScimResourceTypeRequest) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *AddLdapMappingScimResourceTypeRequest) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


