# CorrelatedLdapDataViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the SCIM Resource Type | 
**Schemas** | Pointer to [**[]EnumcorrelatedLdapDataViewSchemaUrn**](EnumcorrelatedLdapDataViewSchemaUrn.md) |  | [optional] 
**StructuralLDAPObjectclass** | **string** | Specifies the LDAP structural object class that should be exposed by this Correlated LDAP Data View. | 
**AuxiliaryLDAPObjectclass** | Pointer to **[]string** | Specifies an auxiliary LDAP object class that should be exposed by this Correlated LDAP Data View. | [optional] 
**IncludeBaseDN** | **string** | Specifies the base DN of the branch of the LDAP directory that can be accessed by this Correlated LDAP Data View. | 
**IncludeFilter** | Pointer to **[]string** | The set of LDAP filters that define the LDAP entries that should be included in this Correlated LDAP Data View. | [optional] 
**IncludeOperationalAttribute** | Pointer to **[]string** | Specifies the set of operational LDAP attributes to be provided by this Correlated LDAP Data View. | [optional] 
**CreateDNPattern** | Pointer to **string** | Specifies the template to use for the DN when creating new entries. | [optional] 
**PrimaryCorrelationAttribute** | **string** | The LDAP attribute from the parent SCIM Resource Type whose value will be used to match objects in the Correlated LDAP Data View. If multiple correlation attributes are required they may be created using additional correlation-attribute-pairs. | 
**SecondaryCorrelationAttribute** | **string** | The LDAP attribute from the Correlated LDAP Data View whose value will be matched with the primary-correlation-attribute. If multiple correlation attributes are required they may be specified by creating additional correlation-attribute-pairs. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewCorrelatedLdapDataViewResponse

`func NewCorrelatedLdapDataViewResponse(id string, structuralLDAPObjectclass string, includeBaseDN string, primaryCorrelationAttribute string, secondaryCorrelationAttribute string, ) *CorrelatedLdapDataViewResponse`

NewCorrelatedLdapDataViewResponse instantiates a new CorrelatedLdapDataViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrelatedLdapDataViewResponseWithDefaults

`func NewCorrelatedLdapDataViewResponseWithDefaults() *CorrelatedLdapDataViewResponse`

NewCorrelatedLdapDataViewResponseWithDefaults instantiates a new CorrelatedLdapDataViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CorrelatedLdapDataViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CorrelatedLdapDataViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CorrelatedLdapDataViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *CorrelatedLdapDataViewResponse) GetSchemas() []EnumcorrelatedLdapDataViewSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CorrelatedLdapDataViewResponse) GetSchemasOk() (*[]EnumcorrelatedLdapDataViewSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CorrelatedLdapDataViewResponse) SetSchemas(v []EnumcorrelatedLdapDataViewSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *CorrelatedLdapDataViewResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetStructuralLDAPObjectclass

`func (o *CorrelatedLdapDataViewResponse) GetStructuralLDAPObjectclass() string`

GetStructuralLDAPObjectclass returns the StructuralLDAPObjectclass field if non-nil, zero value otherwise.

### GetStructuralLDAPObjectclassOk

`func (o *CorrelatedLdapDataViewResponse) GetStructuralLDAPObjectclassOk() (*string, bool)`

GetStructuralLDAPObjectclassOk returns a tuple with the StructuralLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuralLDAPObjectclass

`func (o *CorrelatedLdapDataViewResponse) SetStructuralLDAPObjectclass(v string)`

SetStructuralLDAPObjectclass sets StructuralLDAPObjectclass field to given value.


### GetAuxiliaryLDAPObjectclass

`func (o *CorrelatedLdapDataViewResponse) GetAuxiliaryLDAPObjectclass() []string`

GetAuxiliaryLDAPObjectclass returns the AuxiliaryLDAPObjectclass field if non-nil, zero value otherwise.

### GetAuxiliaryLDAPObjectclassOk

`func (o *CorrelatedLdapDataViewResponse) GetAuxiliaryLDAPObjectclassOk() (*[]string, bool)`

GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryLDAPObjectclass

`func (o *CorrelatedLdapDataViewResponse) SetAuxiliaryLDAPObjectclass(v []string)`

SetAuxiliaryLDAPObjectclass sets AuxiliaryLDAPObjectclass field to given value.

### HasAuxiliaryLDAPObjectclass

`func (o *CorrelatedLdapDataViewResponse) HasAuxiliaryLDAPObjectclass() bool`

HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.

### GetIncludeBaseDN

`func (o *CorrelatedLdapDataViewResponse) GetIncludeBaseDN() string`

GetIncludeBaseDN returns the IncludeBaseDN field if non-nil, zero value otherwise.

### GetIncludeBaseDNOk

`func (o *CorrelatedLdapDataViewResponse) GetIncludeBaseDNOk() (*string, bool)`

GetIncludeBaseDNOk returns a tuple with the IncludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBaseDN

`func (o *CorrelatedLdapDataViewResponse) SetIncludeBaseDN(v string)`

SetIncludeBaseDN sets IncludeBaseDN field to given value.


### GetIncludeFilter

`func (o *CorrelatedLdapDataViewResponse) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *CorrelatedLdapDataViewResponse) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *CorrelatedLdapDataViewResponse) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *CorrelatedLdapDataViewResponse) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetIncludeOperationalAttribute

`func (o *CorrelatedLdapDataViewResponse) GetIncludeOperationalAttribute() []string`

GetIncludeOperationalAttribute returns the IncludeOperationalAttribute field if non-nil, zero value otherwise.

### GetIncludeOperationalAttributeOk

`func (o *CorrelatedLdapDataViewResponse) GetIncludeOperationalAttributeOk() (*[]string, bool)`

GetIncludeOperationalAttributeOk returns a tuple with the IncludeOperationalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOperationalAttribute

`func (o *CorrelatedLdapDataViewResponse) SetIncludeOperationalAttribute(v []string)`

SetIncludeOperationalAttribute sets IncludeOperationalAttribute field to given value.

### HasIncludeOperationalAttribute

`func (o *CorrelatedLdapDataViewResponse) HasIncludeOperationalAttribute() bool`

HasIncludeOperationalAttribute returns a boolean if a field has been set.

### GetCreateDNPattern

`func (o *CorrelatedLdapDataViewResponse) GetCreateDNPattern() string`

GetCreateDNPattern returns the CreateDNPattern field if non-nil, zero value otherwise.

### GetCreateDNPatternOk

`func (o *CorrelatedLdapDataViewResponse) GetCreateDNPatternOk() (*string, bool)`

GetCreateDNPatternOk returns a tuple with the CreateDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDNPattern

`func (o *CorrelatedLdapDataViewResponse) SetCreateDNPattern(v string)`

SetCreateDNPattern sets CreateDNPattern field to given value.

### HasCreateDNPattern

`func (o *CorrelatedLdapDataViewResponse) HasCreateDNPattern() bool`

HasCreateDNPattern returns a boolean if a field has been set.

### GetPrimaryCorrelationAttribute

`func (o *CorrelatedLdapDataViewResponse) GetPrimaryCorrelationAttribute() string`

GetPrimaryCorrelationAttribute returns the PrimaryCorrelationAttribute field if non-nil, zero value otherwise.

### GetPrimaryCorrelationAttributeOk

`func (o *CorrelatedLdapDataViewResponse) GetPrimaryCorrelationAttributeOk() (*string, bool)`

GetPrimaryCorrelationAttributeOk returns a tuple with the PrimaryCorrelationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCorrelationAttribute

`func (o *CorrelatedLdapDataViewResponse) SetPrimaryCorrelationAttribute(v string)`

SetPrimaryCorrelationAttribute sets PrimaryCorrelationAttribute field to given value.


### GetSecondaryCorrelationAttribute

`func (o *CorrelatedLdapDataViewResponse) GetSecondaryCorrelationAttribute() string`

GetSecondaryCorrelationAttribute returns the SecondaryCorrelationAttribute field if non-nil, zero value otherwise.

### GetSecondaryCorrelationAttributeOk

`func (o *CorrelatedLdapDataViewResponse) GetSecondaryCorrelationAttributeOk() (*string, bool)`

GetSecondaryCorrelationAttributeOk returns a tuple with the SecondaryCorrelationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCorrelationAttribute

`func (o *CorrelatedLdapDataViewResponse) SetSecondaryCorrelationAttribute(v string)`

SetSecondaryCorrelationAttribute sets SecondaryCorrelationAttribute field to given value.


### GetMeta

`func (o *CorrelatedLdapDataViewResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CorrelatedLdapDataViewResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CorrelatedLdapDataViewResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CorrelatedLdapDataViewResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *CorrelatedLdapDataViewResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *CorrelatedLdapDataViewResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *CorrelatedLdapDataViewResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *CorrelatedLdapDataViewResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


