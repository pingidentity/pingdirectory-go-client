# AddLdapPassThroughScimResourceTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumldapPassThroughScimResourceTypeSchemaUrn**](EnumldapPassThroughScimResourceTypeSchemaUrn.md) |  | 
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

### NewAddLdapPassThroughScimResourceTypeRequest

`func NewAddLdapPassThroughScimResourceTypeRequest(schemas []EnumldapPassThroughScimResourceTypeSchemaUrn, enabled bool, endpoint string, typeName string, ) *AddLdapPassThroughScimResourceTypeRequest`

NewAddLdapPassThroughScimResourceTypeRequest instantiates a new AddLdapPassThroughScimResourceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLdapPassThroughScimResourceTypeRequestWithDefaults

`func NewAddLdapPassThroughScimResourceTypeRequestWithDefaults() *AddLdapPassThroughScimResourceTypeRequest`

NewAddLdapPassThroughScimResourceTypeRequestWithDefaults instantiates a new AddLdapPassThroughScimResourceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetSchemas() []EnumldapPassThroughScimResourceTypeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetSchemasOk() (*[]EnumldapPassThroughScimResourceTypeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLdapPassThroughScimResourceTypeRequest) SetSchemas(v []EnumldapPassThroughScimResourceTypeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLdapPassThroughScimResourceTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLdapPassThroughScimResourceTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddLdapPassThroughScimResourceTypeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEndpoint

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AddLdapPassThroughScimResourceTypeRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetLookthroughLimit

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetLookthroughLimit() int64`

GetLookthroughLimit returns the LookthroughLimit field if non-nil, zero value otherwise.

### GetLookthroughLimitOk

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetLookthroughLimitOk() (*int64, bool)`

GetLookthroughLimitOk returns a tuple with the LookthroughLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookthroughLimit

`func (o *AddLdapPassThroughScimResourceTypeRequest) SetLookthroughLimit(v int64)`

SetLookthroughLimit sets LookthroughLimit field to given value.

### HasLookthroughLimit

`func (o *AddLdapPassThroughScimResourceTypeRequest) HasLookthroughLimit() bool`

HasLookthroughLimit returns a boolean if a field has been set.

### GetSchemaCheckingOption

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetSchemaCheckingOption() []EnumscimResourceTypeSchemaCheckingOptionProp`

GetSchemaCheckingOption returns the SchemaCheckingOption field if non-nil, zero value otherwise.

### GetSchemaCheckingOptionOk

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetSchemaCheckingOptionOk() (*[]EnumscimResourceTypeSchemaCheckingOptionProp, bool)`

GetSchemaCheckingOptionOk returns a tuple with the SchemaCheckingOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaCheckingOption

`func (o *AddLdapPassThroughScimResourceTypeRequest) SetSchemaCheckingOption(v []EnumscimResourceTypeSchemaCheckingOptionProp)`

SetSchemaCheckingOption sets SchemaCheckingOption field to given value.

### HasSchemaCheckingOption

`func (o *AddLdapPassThroughScimResourceTypeRequest) HasSchemaCheckingOption() bool`

HasSchemaCheckingOption returns a boolean if a field has been set.

### GetStructuralLDAPObjectclass

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetStructuralLDAPObjectclass() string`

GetStructuralLDAPObjectclass returns the StructuralLDAPObjectclass field if non-nil, zero value otherwise.

### GetStructuralLDAPObjectclassOk

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetStructuralLDAPObjectclassOk() (*string, bool)`

GetStructuralLDAPObjectclassOk returns a tuple with the StructuralLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuralLDAPObjectclass

`func (o *AddLdapPassThroughScimResourceTypeRequest) SetStructuralLDAPObjectclass(v string)`

SetStructuralLDAPObjectclass sets StructuralLDAPObjectclass field to given value.

### HasStructuralLDAPObjectclass

`func (o *AddLdapPassThroughScimResourceTypeRequest) HasStructuralLDAPObjectclass() bool`

HasStructuralLDAPObjectclass returns a boolean if a field has been set.

### GetAuxiliaryLDAPObjectclass

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetAuxiliaryLDAPObjectclass() []string`

GetAuxiliaryLDAPObjectclass returns the AuxiliaryLDAPObjectclass field if non-nil, zero value otherwise.

### GetAuxiliaryLDAPObjectclassOk

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetAuxiliaryLDAPObjectclassOk() (*[]string, bool)`

GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryLDAPObjectclass

`func (o *AddLdapPassThroughScimResourceTypeRequest) SetAuxiliaryLDAPObjectclass(v []string)`

SetAuxiliaryLDAPObjectclass sets AuxiliaryLDAPObjectclass field to given value.

### HasAuxiliaryLDAPObjectclass

`func (o *AddLdapPassThroughScimResourceTypeRequest) HasAuxiliaryLDAPObjectclass() bool`

HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.

### GetIncludeBaseDN

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetIncludeBaseDN() string`

GetIncludeBaseDN returns the IncludeBaseDN field if non-nil, zero value otherwise.

### GetIncludeBaseDNOk

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetIncludeBaseDNOk() (*string, bool)`

GetIncludeBaseDNOk returns a tuple with the IncludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBaseDN

`func (o *AddLdapPassThroughScimResourceTypeRequest) SetIncludeBaseDN(v string)`

SetIncludeBaseDN sets IncludeBaseDN field to given value.

### HasIncludeBaseDN

`func (o *AddLdapPassThroughScimResourceTypeRequest) HasIncludeBaseDN() bool`

HasIncludeBaseDN returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *AddLdapPassThroughScimResourceTypeRequest) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *AddLdapPassThroughScimResourceTypeRequest) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetIncludeOperationalAttribute

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetIncludeOperationalAttribute() []string`

GetIncludeOperationalAttribute returns the IncludeOperationalAttribute field if non-nil, zero value otherwise.

### GetIncludeOperationalAttributeOk

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetIncludeOperationalAttributeOk() (*[]string, bool)`

GetIncludeOperationalAttributeOk returns a tuple with the IncludeOperationalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOperationalAttribute

`func (o *AddLdapPassThroughScimResourceTypeRequest) SetIncludeOperationalAttribute(v []string)`

SetIncludeOperationalAttribute sets IncludeOperationalAttribute field to given value.

### HasIncludeOperationalAttribute

`func (o *AddLdapPassThroughScimResourceTypeRequest) HasIncludeOperationalAttribute() bool`

HasIncludeOperationalAttribute returns a boolean if a field has been set.

### GetCreateDNPattern

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetCreateDNPattern() string`

GetCreateDNPattern returns the CreateDNPattern field if non-nil, zero value otherwise.

### GetCreateDNPatternOk

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetCreateDNPatternOk() (*string, bool)`

GetCreateDNPatternOk returns a tuple with the CreateDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDNPattern

`func (o *AddLdapPassThroughScimResourceTypeRequest) SetCreateDNPattern(v string)`

SetCreateDNPattern sets CreateDNPattern field to given value.

### HasCreateDNPattern

`func (o *AddLdapPassThroughScimResourceTypeRequest) HasCreateDNPattern() bool`

HasCreateDNPattern returns a boolean if a field has been set.

### GetTypeName

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *AddLdapPassThroughScimResourceTypeRequest) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *AddLdapPassThroughScimResourceTypeRequest) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


