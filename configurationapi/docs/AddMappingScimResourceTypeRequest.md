# AddMappingScimResourceTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnummappingScimResourceTypeSchemaUrn**](EnummappingScimResourceTypeSchemaUrn.md) |  | 
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

### NewAddMappingScimResourceTypeRequest

`func NewAddMappingScimResourceTypeRequest(schemas []EnummappingScimResourceTypeSchemaUrn, coreSchema string, enabled bool, endpoint string, typeName string, ) *AddMappingScimResourceTypeRequest`

NewAddMappingScimResourceTypeRequest instantiates a new AddMappingScimResourceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMappingScimResourceTypeRequestWithDefaults

`func NewAddMappingScimResourceTypeRequestWithDefaults() *AddMappingScimResourceTypeRequest`

NewAddMappingScimResourceTypeRequestWithDefaults instantiates a new AddMappingScimResourceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddMappingScimResourceTypeRequest) GetSchemas() []EnummappingScimResourceTypeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddMappingScimResourceTypeRequest) GetSchemasOk() (*[]EnummappingScimResourceTypeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddMappingScimResourceTypeRequest) SetSchemas(v []EnummappingScimResourceTypeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetCoreSchema

`func (o *AddMappingScimResourceTypeRequest) GetCoreSchema() string`

GetCoreSchema returns the CoreSchema field if non-nil, zero value otherwise.

### GetCoreSchemaOk

`func (o *AddMappingScimResourceTypeRequest) GetCoreSchemaOk() (*string, bool)`

GetCoreSchemaOk returns a tuple with the CoreSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreSchema

`func (o *AddMappingScimResourceTypeRequest) SetCoreSchema(v string)`

SetCoreSchema sets CoreSchema field to given value.


### GetRequiredSchemaExtension

`func (o *AddMappingScimResourceTypeRequest) GetRequiredSchemaExtension() []string`

GetRequiredSchemaExtension returns the RequiredSchemaExtension field if non-nil, zero value otherwise.

### GetRequiredSchemaExtensionOk

`func (o *AddMappingScimResourceTypeRequest) GetRequiredSchemaExtensionOk() (*[]string, bool)`

GetRequiredSchemaExtensionOk returns a tuple with the RequiredSchemaExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSchemaExtension

`func (o *AddMappingScimResourceTypeRequest) SetRequiredSchemaExtension(v []string)`

SetRequiredSchemaExtension sets RequiredSchemaExtension field to given value.

### HasRequiredSchemaExtension

`func (o *AddMappingScimResourceTypeRequest) HasRequiredSchemaExtension() bool`

HasRequiredSchemaExtension returns a boolean if a field has been set.

### GetOptionalSchemaExtension

`func (o *AddMappingScimResourceTypeRequest) GetOptionalSchemaExtension() []string`

GetOptionalSchemaExtension returns the OptionalSchemaExtension field if non-nil, zero value otherwise.

### GetOptionalSchemaExtensionOk

`func (o *AddMappingScimResourceTypeRequest) GetOptionalSchemaExtensionOk() (*[]string, bool)`

GetOptionalSchemaExtensionOk returns a tuple with the OptionalSchemaExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalSchemaExtension

`func (o *AddMappingScimResourceTypeRequest) SetOptionalSchemaExtension(v []string)`

SetOptionalSchemaExtension sets OptionalSchemaExtension field to given value.

### HasOptionalSchemaExtension

`func (o *AddMappingScimResourceTypeRequest) HasOptionalSchemaExtension() bool`

HasOptionalSchemaExtension returns a boolean if a field has been set.

### GetDescription

`func (o *AddMappingScimResourceTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddMappingScimResourceTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddMappingScimResourceTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddMappingScimResourceTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddMappingScimResourceTypeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddMappingScimResourceTypeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddMappingScimResourceTypeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEndpoint

`func (o *AddMappingScimResourceTypeRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AddMappingScimResourceTypeRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AddMappingScimResourceTypeRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetLookthroughLimit

`func (o *AddMappingScimResourceTypeRequest) GetLookthroughLimit() int64`

GetLookthroughLimit returns the LookthroughLimit field if non-nil, zero value otherwise.

### GetLookthroughLimitOk

`func (o *AddMappingScimResourceTypeRequest) GetLookthroughLimitOk() (*int64, bool)`

GetLookthroughLimitOk returns a tuple with the LookthroughLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookthroughLimit

`func (o *AddMappingScimResourceTypeRequest) SetLookthroughLimit(v int64)`

SetLookthroughLimit sets LookthroughLimit field to given value.

### HasLookthroughLimit

`func (o *AddMappingScimResourceTypeRequest) HasLookthroughLimit() bool`

HasLookthroughLimit returns a boolean if a field has been set.

### GetSchemaCheckingOption

`func (o *AddMappingScimResourceTypeRequest) GetSchemaCheckingOption() []EnumscimResourceTypeSchemaCheckingOptionProp`

GetSchemaCheckingOption returns the SchemaCheckingOption field if non-nil, zero value otherwise.

### GetSchemaCheckingOptionOk

`func (o *AddMappingScimResourceTypeRequest) GetSchemaCheckingOptionOk() (*[]EnumscimResourceTypeSchemaCheckingOptionProp, bool)`

GetSchemaCheckingOptionOk returns a tuple with the SchemaCheckingOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaCheckingOption

`func (o *AddMappingScimResourceTypeRequest) SetSchemaCheckingOption(v []EnumscimResourceTypeSchemaCheckingOptionProp)`

SetSchemaCheckingOption sets SchemaCheckingOption field to given value.

### HasSchemaCheckingOption

`func (o *AddMappingScimResourceTypeRequest) HasSchemaCheckingOption() bool`

HasSchemaCheckingOption returns a boolean if a field has been set.

### GetStructuralLDAPObjectclass

`func (o *AddMappingScimResourceTypeRequest) GetStructuralLDAPObjectclass() string`

GetStructuralLDAPObjectclass returns the StructuralLDAPObjectclass field if non-nil, zero value otherwise.

### GetStructuralLDAPObjectclassOk

`func (o *AddMappingScimResourceTypeRequest) GetStructuralLDAPObjectclassOk() (*string, bool)`

GetStructuralLDAPObjectclassOk returns a tuple with the StructuralLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuralLDAPObjectclass

`func (o *AddMappingScimResourceTypeRequest) SetStructuralLDAPObjectclass(v string)`

SetStructuralLDAPObjectclass sets StructuralLDAPObjectclass field to given value.

### HasStructuralLDAPObjectclass

`func (o *AddMappingScimResourceTypeRequest) HasStructuralLDAPObjectclass() bool`

HasStructuralLDAPObjectclass returns a boolean if a field has been set.

### GetAuxiliaryLDAPObjectclass

`func (o *AddMappingScimResourceTypeRequest) GetAuxiliaryLDAPObjectclass() []string`

GetAuxiliaryLDAPObjectclass returns the AuxiliaryLDAPObjectclass field if non-nil, zero value otherwise.

### GetAuxiliaryLDAPObjectclassOk

`func (o *AddMappingScimResourceTypeRequest) GetAuxiliaryLDAPObjectclassOk() (*[]string, bool)`

GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryLDAPObjectclass

`func (o *AddMappingScimResourceTypeRequest) SetAuxiliaryLDAPObjectclass(v []string)`

SetAuxiliaryLDAPObjectclass sets AuxiliaryLDAPObjectclass field to given value.

### HasAuxiliaryLDAPObjectclass

`func (o *AddMappingScimResourceTypeRequest) HasAuxiliaryLDAPObjectclass() bool`

HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.

### GetIncludeBaseDN

`func (o *AddMappingScimResourceTypeRequest) GetIncludeBaseDN() string`

GetIncludeBaseDN returns the IncludeBaseDN field if non-nil, zero value otherwise.

### GetIncludeBaseDNOk

`func (o *AddMappingScimResourceTypeRequest) GetIncludeBaseDNOk() (*string, bool)`

GetIncludeBaseDNOk returns a tuple with the IncludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBaseDN

`func (o *AddMappingScimResourceTypeRequest) SetIncludeBaseDN(v string)`

SetIncludeBaseDN sets IncludeBaseDN field to given value.

### HasIncludeBaseDN

`func (o *AddMappingScimResourceTypeRequest) HasIncludeBaseDN() bool`

HasIncludeBaseDN returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *AddMappingScimResourceTypeRequest) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *AddMappingScimResourceTypeRequest) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *AddMappingScimResourceTypeRequest) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *AddMappingScimResourceTypeRequest) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetIncludeOperationalAttribute

`func (o *AddMappingScimResourceTypeRequest) GetIncludeOperationalAttribute() []string`

GetIncludeOperationalAttribute returns the IncludeOperationalAttribute field if non-nil, zero value otherwise.

### GetIncludeOperationalAttributeOk

`func (o *AddMappingScimResourceTypeRequest) GetIncludeOperationalAttributeOk() (*[]string, bool)`

GetIncludeOperationalAttributeOk returns a tuple with the IncludeOperationalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOperationalAttribute

`func (o *AddMappingScimResourceTypeRequest) SetIncludeOperationalAttribute(v []string)`

SetIncludeOperationalAttribute sets IncludeOperationalAttribute field to given value.

### HasIncludeOperationalAttribute

`func (o *AddMappingScimResourceTypeRequest) HasIncludeOperationalAttribute() bool`

HasIncludeOperationalAttribute returns a boolean if a field has been set.

### GetCreateDNPattern

`func (o *AddMappingScimResourceTypeRequest) GetCreateDNPattern() string`

GetCreateDNPattern returns the CreateDNPattern field if non-nil, zero value otherwise.

### GetCreateDNPatternOk

`func (o *AddMappingScimResourceTypeRequest) GetCreateDNPatternOk() (*string, bool)`

GetCreateDNPatternOk returns a tuple with the CreateDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDNPattern

`func (o *AddMappingScimResourceTypeRequest) SetCreateDNPattern(v string)`

SetCreateDNPattern sets CreateDNPattern field to given value.

### HasCreateDNPattern

`func (o *AddMappingScimResourceTypeRequest) HasCreateDNPattern() bool`

HasCreateDNPattern returns a boolean if a field has been set.

### GetTypeName

`func (o *AddMappingScimResourceTypeRequest) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *AddMappingScimResourceTypeRequest) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *AddMappingScimResourceTypeRequest) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


