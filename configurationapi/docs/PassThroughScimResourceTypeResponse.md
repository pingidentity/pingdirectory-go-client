# PassThroughScimResourceTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumpassThroughScimResourceTypeSchemaUrn**](EnumpassThroughScimResourceTypeSchemaUrn.md) |  | 
**Id** | **string** | Name of the SCIM Resource Type | 
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

## Methods

### NewPassThroughScimResourceTypeResponse

`func NewPassThroughScimResourceTypeResponse(schemas []EnumpassThroughScimResourceTypeSchemaUrn, id string, enabled bool, endpoint string, ) *PassThroughScimResourceTypeResponse`

NewPassThroughScimResourceTypeResponse instantiates a new PassThroughScimResourceTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassThroughScimResourceTypeResponseWithDefaults

`func NewPassThroughScimResourceTypeResponseWithDefaults() *PassThroughScimResourceTypeResponse`

NewPassThroughScimResourceTypeResponseWithDefaults instantiates a new PassThroughScimResourceTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *PassThroughScimResourceTypeResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PassThroughScimResourceTypeResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PassThroughScimResourceTypeResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PassThroughScimResourceTypeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *PassThroughScimResourceTypeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *PassThroughScimResourceTypeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *PassThroughScimResourceTypeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *PassThroughScimResourceTypeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *PassThroughScimResourceTypeResponse) GetSchemas() []EnumpassThroughScimResourceTypeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PassThroughScimResourceTypeResponse) GetSchemasOk() (*[]EnumpassThroughScimResourceTypeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PassThroughScimResourceTypeResponse) SetSchemas(v []EnumpassThroughScimResourceTypeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *PassThroughScimResourceTypeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PassThroughScimResourceTypeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PassThroughScimResourceTypeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *PassThroughScimResourceTypeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PassThroughScimResourceTypeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PassThroughScimResourceTypeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PassThroughScimResourceTypeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PassThroughScimResourceTypeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PassThroughScimResourceTypeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PassThroughScimResourceTypeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEndpoint

`func (o *PassThroughScimResourceTypeResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PassThroughScimResourceTypeResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PassThroughScimResourceTypeResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetLookthroughLimit

`func (o *PassThroughScimResourceTypeResponse) GetLookthroughLimit() int64`

GetLookthroughLimit returns the LookthroughLimit field if non-nil, zero value otherwise.

### GetLookthroughLimitOk

`func (o *PassThroughScimResourceTypeResponse) GetLookthroughLimitOk() (*int64, bool)`

GetLookthroughLimitOk returns a tuple with the LookthroughLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookthroughLimit

`func (o *PassThroughScimResourceTypeResponse) SetLookthroughLimit(v int64)`

SetLookthroughLimit sets LookthroughLimit field to given value.

### HasLookthroughLimit

`func (o *PassThroughScimResourceTypeResponse) HasLookthroughLimit() bool`

HasLookthroughLimit returns a boolean if a field has been set.

### GetSchemaCheckingOption

`func (o *PassThroughScimResourceTypeResponse) GetSchemaCheckingOption() []EnumscimResourceTypeSchemaCheckingOptionProp`

GetSchemaCheckingOption returns the SchemaCheckingOption field if non-nil, zero value otherwise.

### GetSchemaCheckingOptionOk

`func (o *PassThroughScimResourceTypeResponse) GetSchemaCheckingOptionOk() (*[]EnumscimResourceTypeSchemaCheckingOptionProp, bool)`

GetSchemaCheckingOptionOk returns a tuple with the SchemaCheckingOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaCheckingOption

`func (o *PassThroughScimResourceTypeResponse) SetSchemaCheckingOption(v []EnumscimResourceTypeSchemaCheckingOptionProp)`

SetSchemaCheckingOption sets SchemaCheckingOption field to given value.

### HasSchemaCheckingOption

`func (o *PassThroughScimResourceTypeResponse) HasSchemaCheckingOption() bool`

HasSchemaCheckingOption returns a boolean if a field has been set.

### GetStructuralLDAPObjectclass

`func (o *PassThroughScimResourceTypeResponse) GetStructuralLDAPObjectclass() string`

GetStructuralLDAPObjectclass returns the StructuralLDAPObjectclass field if non-nil, zero value otherwise.

### GetStructuralLDAPObjectclassOk

`func (o *PassThroughScimResourceTypeResponse) GetStructuralLDAPObjectclassOk() (*string, bool)`

GetStructuralLDAPObjectclassOk returns a tuple with the StructuralLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuralLDAPObjectclass

`func (o *PassThroughScimResourceTypeResponse) SetStructuralLDAPObjectclass(v string)`

SetStructuralLDAPObjectclass sets StructuralLDAPObjectclass field to given value.

### HasStructuralLDAPObjectclass

`func (o *PassThroughScimResourceTypeResponse) HasStructuralLDAPObjectclass() bool`

HasStructuralLDAPObjectclass returns a boolean if a field has been set.

### GetAuxiliaryLDAPObjectclass

`func (o *PassThroughScimResourceTypeResponse) GetAuxiliaryLDAPObjectclass() []string`

GetAuxiliaryLDAPObjectclass returns the AuxiliaryLDAPObjectclass field if non-nil, zero value otherwise.

### GetAuxiliaryLDAPObjectclassOk

`func (o *PassThroughScimResourceTypeResponse) GetAuxiliaryLDAPObjectclassOk() (*[]string, bool)`

GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryLDAPObjectclass

`func (o *PassThroughScimResourceTypeResponse) SetAuxiliaryLDAPObjectclass(v []string)`

SetAuxiliaryLDAPObjectclass sets AuxiliaryLDAPObjectclass field to given value.

### HasAuxiliaryLDAPObjectclass

`func (o *PassThroughScimResourceTypeResponse) HasAuxiliaryLDAPObjectclass() bool`

HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.

### GetIncludeBaseDN

`func (o *PassThroughScimResourceTypeResponse) GetIncludeBaseDN() string`

GetIncludeBaseDN returns the IncludeBaseDN field if non-nil, zero value otherwise.

### GetIncludeBaseDNOk

`func (o *PassThroughScimResourceTypeResponse) GetIncludeBaseDNOk() (*string, bool)`

GetIncludeBaseDNOk returns a tuple with the IncludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBaseDN

`func (o *PassThroughScimResourceTypeResponse) SetIncludeBaseDN(v string)`

SetIncludeBaseDN sets IncludeBaseDN field to given value.

### HasIncludeBaseDN

`func (o *PassThroughScimResourceTypeResponse) HasIncludeBaseDN() bool`

HasIncludeBaseDN returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *PassThroughScimResourceTypeResponse) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *PassThroughScimResourceTypeResponse) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *PassThroughScimResourceTypeResponse) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *PassThroughScimResourceTypeResponse) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetIncludeOperationalAttribute

`func (o *PassThroughScimResourceTypeResponse) GetIncludeOperationalAttribute() []string`

GetIncludeOperationalAttribute returns the IncludeOperationalAttribute field if non-nil, zero value otherwise.

### GetIncludeOperationalAttributeOk

`func (o *PassThroughScimResourceTypeResponse) GetIncludeOperationalAttributeOk() (*[]string, bool)`

GetIncludeOperationalAttributeOk returns a tuple with the IncludeOperationalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOperationalAttribute

`func (o *PassThroughScimResourceTypeResponse) SetIncludeOperationalAttribute(v []string)`

SetIncludeOperationalAttribute sets IncludeOperationalAttribute field to given value.

### HasIncludeOperationalAttribute

`func (o *PassThroughScimResourceTypeResponse) HasIncludeOperationalAttribute() bool`

HasIncludeOperationalAttribute returns a boolean if a field has been set.

### GetCreateDNPattern

`func (o *PassThroughScimResourceTypeResponse) GetCreateDNPattern() string`

GetCreateDNPattern returns the CreateDNPattern field if non-nil, zero value otherwise.

### GetCreateDNPatternOk

`func (o *PassThroughScimResourceTypeResponse) GetCreateDNPatternOk() (*string, bool)`

GetCreateDNPatternOk returns a tuple with the CreateDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDNPattern

`func (o *PassThroughScimResourceTypeResponse) SetCreateDNPattern(v string)`

SetCreateDNPattern sets CreateDNPattern field to given value.

### HasCreateDNPattern

`func (o *PassThroughScimResourceTypeResponse) HasCreateDNPattern() bool`

HasCreateDNPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


