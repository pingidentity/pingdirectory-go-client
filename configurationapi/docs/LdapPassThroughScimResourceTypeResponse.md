# LdapPassThroughScimResourceTypeResponse

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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the SCIM Resource Type | 

## Methods

### NewLdapPassThroughScimResourceTypeResponse

`func NewLdapPassThroughScimResourceTypeResponse(schemas []EnumldapPassThroughScimResourceTypeSchemaUrn, enabled bool, endpoint string, id string, ) *LdapPassThroughScimResourceTypeResponse`

NewLdapPassThroughScimResourceTypeResponse instantiates a new LdapPassThroughScimResourceTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapPassThroughScimResourceTypeResponseWithDefaults

`func NewLdapPassThroughScimResourceTypeResponseWithDefaults() *LdapPassThroughScimResourceTypeResponse`

NewLdapPassThroughScimResourceTypeResponseWithDefaults instantiates a new LdapPassThroughScimResourceTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *LdapPassThroughScimResourceTypeResponse) GetSchemas() []EnumldapPassThroughScimResourceTypeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LdapPassThroughScimResourceTypeResponse) GetSchemasOk() (*[]EnumldapPassThroughScimResourceTypeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LdapPassThroughScimResourceTypeResponse) SetSchemas(v []EnumldapPassThroughScimResourceTypeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *LdapPassThroughScimResourceTypeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LdapPassThroughScimResourceTypeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LdapPassThroughScimResourceTypeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LdapPassThroughScimResourceTypeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *LdapPassThroughScimResourceTypeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LdapPassThroughScimResourceTypeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LdapPassThroughScimResourceTypeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEndpoint

`func (o *LdapPassThroughScimResourceTypeResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *LdapPassThroughScimResourceTypeResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *LdapPassThroughScimResourceTypeResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetLookthroughLimit

`func (o *LdapPassThroughScimResourceTypeResponse) GetLookthroughLimit() int64`

GetLookthroughLimit returns the LookthroughLimit field if non-nil, zero value otherwise.

### GetLookthroughLimitOk

`func (o *LdapPassThroughScimResourceTypeResponse) GetLookthroughLimitOk() (*int64, bool)`

GetLookthroughLimitOk returns a tuple with the LookthroughLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookthroughLimit

`func (o *LdapPassThroughScimResourceTypeResponse) SetLookthroughLimit(v int64)`

SetLookthroughLimit sets LookthroughLimit field to given value.

### HasLookthroughLimit

`func (o *LdapPassThroughScimResourceTypeResponse) HasLookthroughLimit() bool`

HasLookthroughLimit returns a boolean if a field has been set.

### GetSchemaCheckingOption

`func (o *LdapPassThroughScimResourceTypeResponse) GetSchemaCheckingOption() []EnumscimResourceTypeSchemaCheckingOptionProp`

GetSchemaCheckingOption returns the SchemaCheckingOption field if non-nil, zero value otherwise.

### GetSchemaCheckingOptionOk

`func (o *LdapPassThroughScimResourceTypeResponse) GetSchemaCheckingOptionOk() (*[]EnumscimResourceTypeSchemaCheckingOptionProp, bool)`

GetSchemaCheckingOptionOk returns a tuple with the SchemaCheckingOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaCheckingOption

`func (o *LdapPassThroughScimResourceTypeResponse) SetSchemaCheckingOption(v []EnumscimResourceTypeSchemaCheckingOptionProp)`

SetSchemaCheckingOption sets SchemaCheckingOption field to given value.

### HasSchemaCheckingOption

`func (o *LdapPassThroughScimResourceTypeResponse) HasSchemaCheckingOption() bool`

HasSchemaCheckingOption returns a boolean if a field has been set.

### GetStructuralLDAPObjectclass

`func (o *LdapPassThroughScimResourceTypeResponse) GetStructuralLDAPObjectclass() string`

GetStructuralLDAPObjectclass returns the StructuralLDAPObjectclass field if non-nil, zero value otherwise.

### GetStructuralLDAPObjectclassOk

`func (o *LdapPassThroughScimResourceTypeResponse) GetStructuralLDAPObjectclassOk() (*string, bool)`

GetStructuralLDAPObjectclassOk returns a tuple with the StructuralLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuralLDAPObjectclass

`func (o *LdapPassThroughScimResourceTypeResponse) SetStructuralLDAPObjectclass(v string)`

SetStructuralLDAPObjectclass sets StructuralLDAPObjectclass field to given value.

### HasStructuralLDAPObjectclass

`func (o *LdapPassThroughScimResourceTypeResponse) HasStructuralLDAPObjectclass() bool`

HasStructuralLDAPObjectclass returns a boolean if a field has been set.

### GetAuxiliaryLDAPObjectclass

`func (o *LdapPassThroughScimResourceTypeResponse) GetAuxiliaryLDAPObjectclass() []string`

GetAuxiliaryLDAPObjectclass returns the AuxiliaryLDAPObjectclass field if non-nil, zero value otherwise.

### GetAuxiliaryLDAPObjectclassOk

`func (o *LdapPassThroughScimResourceTypeResponse) GetAuxiliaryLDAPObjectclassOk() (*[]string, bool)`

GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryLDAPObjectclass

`func (o *LdapPassThroughScimResourceTypeResponse) SetAuxiliaryLDAPObjectclass(v []string)`

SetAuxiliaryLDAPObjectclass sets AuxiliaryLDAPObjectclass field to given value.

### HasAuxiliaryLDAPObjectclass

`func (o *LdapPassThroughScimResourceTypeResponse) HasAuxiliaryLDAPObjectclass() bool`

HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.

### GetIncludeBaseDN

`func (o *LdapPassThroughScimResourceTypeResponse) GetIncludeBaseDN() string`

GetIncludeBaseDN returns the IncludeBaseDN field if non-nil, zero value otherwise.

### GetIncludeBaseDNOk

`func (o *LdapPassThroughScimResourceTypeResponse) GetIncludeBaseDNOk() (*string, bool)`

GetIncludeBaseDNOk returns a tuple with the IncludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBaseDN

`func (o *LdapPassThroughScimResourceTypeResponse) SetIncludeBaseDN(v string)`

SetIncludeBaseDN sets IncludeBaseDN field to given value.

### HasIncludeBaseDN

`func (o *LdapPassThroughScimResourceTypeResponse) HasIncludeBaseDN() bool`

HasIncludeBaseDN returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *LdapPassThroughScimResourceTypeResponse) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *LdapPassThroughScimResourceTypeResponse) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *LdapPassThroughScimResourceTypeResponse) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *LdapPassThroughScimResourceTypeResponse) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetIncludeOperationalAttribute

`func (o *LdapPassThroughScimResourceTypeResponse) GetIncludeOperationalAttribute() []string`

GetIncludeOperationalAttribute returns the IncludeOperationalAttribute field if non-nil, zero value otherwise.

### GetIncludeOperationalAttributeOk

`func (o *LdapPassThroughScimResourceTypeResponse) GetIncludeOperationalAttributeOk() (*[]string, bool)`

GetIncludeOperationalAttributeOk returns a tuple with the IncludeOperationalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOperationalAttribute

`func (o *LdapPassThroughScimResourceTypeResponse) SetIncludeOperationalAttribute(v []string)`

SetIncludeOperationalAttribute sets IncludeOperationalAttribute field to given value.

### HasIncludeOperationalAttribute

`func (o *LdapPassThroughScimResourceTypeResponse) HasIncludeOperationalAttribute() bool`

HasIncludeOperationalAttribute returns a boolean if a field has been set.

### GetCreateDNPattern

`func (o *LdapPassThroughScimResourceTypeResponse) GetCreateDNPattern() string`

GetCreateDNPattern returns the CreateDNPattern field if non-nil, zero value otherwise.

### GetCreateDNPatternOk

`func (o *LdapPassThroughScimResourceTypeResponse) GetCreateDNPatternOk() (*string, bool)`

GetCreateDNPatternOk returns a tuple with the CreateDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDNPattern

`func (o *LdapPassThroughScimResourceTypeResponse) SetCreateDNPattern(v string)`

SetCreateDNPattern sets CreateDNPattern field to given value.

### HasCreateDNPattern

`func (o *LdapPassThroughScimResourceTypeResponse) HasCreateDNPattern() bool`

HasCreateDNPattern returns a boolean if a field has been set.

### GetMeta

`func (o *LdapPassThroughScimResourceTypeResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LdapPassThroughScimResourceTypeResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LdapPassThroughScimResourceTypeResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LdapPassThroughScimResourceTypeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *LdapPassThroughScimResourceTypeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *LdapPassThroughScimResourceTypeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *LdapPassThroughScimResourceTypeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *LdapPassThroughScimResourceTypeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *LdapPassThroughScimResourceTypeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LdapPassThroughScimResourceTypeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LdapPassThroughScimResourceTypeResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


