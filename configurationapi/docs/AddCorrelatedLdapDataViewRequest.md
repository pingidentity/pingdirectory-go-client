# AddCorrelatedLdapDataViewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewName** | **string** | Name of the new Correlated LDAP Data View | 
**Schemas** | Pointer to [**[]EnumcorrelatedLdapDataViewSchemaUrn**](EnumcorrelatedLdapDataViewSchemaUrn.md) |  | [optional] 
**StructuralLDAPObjectclass** | **string** | Specifies the LDAP structural object class that should be exposed by this Correlated LDAP Data View. | 
**AuxiliaryLDAPObjectclass** | Pointer to **[]string** | Specifies an auxiliary LDAP object class that should be exposed by this Correlated LDAP Data View. | [optional] 
**IncludeBaseDN** | **string** | Specifies the base DN of the branch of the LDAP directory that can be accessed by this Correlated LDAP Data View. | 
**IncludeFilter** | Pointer to **[]string** | The set of LDAP filters that define the LDAP entries that should be included in this Correlated LDAP Data View. | [optional] 
**IncludeOperationalAttribute** | Pointer to **[]string** | Specifies the set of operational LDAP attributes to be provided by this Correlated LDAP Data View. | [optional] 
**CreateDNPattern** | Pointer to **string** | Specifies the template to use for the DN when creating new entries. | [optional] 
**PrimaryCorrelationAttribute** | **string** | The LDAP attribute from the parent SCIM Resource Type whose value will be used to match objects in the Correlated LDAP Data View. If multiple correlation attributes are required they may be created using additional correlation-attribute-pairs. | 
**SecondaryCorrelationAttribute** | **string** | The LDAP attribute from the Correlated LDAP Data View whose value will be matched with the primary-correlation-attribute. If multiple correlation attributes are required they may be specified by creating additional correlation-attribute-pairs. | 

## Methods

### NewAddCorrelatedLdapDataViewRequest

`func NewAddCorrelatedLdapDataViewRequest(viewName string, structuralLDAPObjectclass string, includeBaseDN string, primaryCorrelationAttribute string, secondaryCorrelationAttribute string, ) *AddCorrelatedLdapDataViewRequest`

NewAddCorrelatedLdapDataViewRequest instantiates a new AddCorrelatedLdapDataViewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCorrelatedLdapDataViewRequestWithDefaults

`func NewAddCorrelatedLdapDataViewRequestWithDefaults() *AddCorrelatedLdapDataViewRequest`

NewAddCorrelatedLdapDataViewRequestWithDefaults instantiates a new AddCorrelatedLdapDataViewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewName

`func (o *AddCorrelatedLdapDataViewRequest) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *AddCorrelatedLdapDataViewRequest) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *AddCorrelatedLdapDataViewRequest) SetViewName(v string)`

SetViewName sets ViewName field to given value.


### GetSchemas

`func (o *AddCorrelatedLdapDataViewRequest) GetSchemas() []EnumcorrelatedLdapDataViewSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddCorrelatedLdapDataViewRequest) GetSchemasOk() (*[]EnumcorrelatedLdapDataViewSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddCorrelatedLdapDataViewRequest) SetSchemas(v []EnumcorrelatedLdapDataViewSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddCorrelatedLdapDataViewRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetStructuralLDAPObjectclass

`func (o *AddCorrelatedLdapDataViewRequest) GetStructuralLDAPObjectclass() string`

GetStructuralLDAPObjectclass returns the StructuralLDAPObjectclass field if non-nil, zero value otherwise.

### GetStructuralLDAPObjectclassOk

`func (o *AddCorrelatedLdapDataViewRequest) GetStructuralLDAPObjectclassOk() (*string, bool)`

GetStructuralLDAPObjectclassOk returns a tuple with the StructuralLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuralLDAPObjectclass

`func (o *AddCorrelatedLdapDataViewRequest) SetStructuralLDAPObjectclass(v string)`

SetStructuralLDAPObjectclass sets StructuralLDAPObjectclass field to given value.


### GetAuxiliaryLDAPObjectclass

`func (o *AddCorrelatedLdapDataViewRequest) GetAuxiliaryLDAPObjectclass() []string`

GetAuxiliaryLDAPObjectclass returns the AuxiliaryLDAPObjectclass field if non-nil, zero value otherwise.

### GetAuxiliaryLDAPObjectclassOk

`func (o *AddCorrelatedLdapDataViewRequest) GetAuxiliaryLDAPObjectclassOk() (*[]string, bool)`

GetAuxiliaryLDAPObjectclassOk returns a tuple with the AuxiliaryLDAPObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryLDAPObjectclass

`func (o *AddCorrelatedLdapDataViewRequest) SetAuxiliaryLDAPObjectclass(v []string)`

SetAuxiliaryLDAPObjectclass sets AuxiliaryLDAPObjectclass field to given value.

### HasAuxiliaryLDAPObjectclass

`func (o *AddCorrelatedLdapDataViewRequest) HasAuxiliaryLDAPObjectclass() bool`

HasAuxiliaryLDAPObjectclass returns a boolean if a field has been set.

### GetIncludeBaseDN

`func (o *AddCorrelatedLdapDataViewRequest) GetIncludeBaseDN() string`

GetIncludeBaseDN returns the IncludeBaseDN field if non-nil, zero value otherwise.

### GetIncludeBaseDNOk

`func (o *AddCorrelatedLdapDataViewRequest) GetIncludeBaseDNOk() (*string, bool)`

GetIncludeBaseDNOk returns a tuple with the IncludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBaseDN

`func (o *AddCorrelatedLdapDataViewRequest) SetIncludeBaseDN(v string)`

SetIncludeBaseDN sets IncludeBaseDN field to given value.


### GetIncludeFilter

`func (o *AddCorrelatedLdapDataViewRequest) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *AddCorrelatedLdapDataViewRequest) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *AddCorrelatedLdapDataViewRequest) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *AddCorrelatedLdapDataViewRequest) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetIncludeOperationalAttribute

`func (o *AddCorrelatedLdapDataViewRequest) GetIncludeOperationalAttribute() []string`

GetIncludeOperationalAttribute returns the IncludeOperationalAttribute field if non-nil, zero value otherwise.

### GetIncludeOperationalAttributeOk

`func (o *AddCorrelatedLdapDataViewRequest) GetIncludeOperationalAttributeOk() (*[]string, bool)`

GetIncludeOperationalAttributeOk returns a tuple with the IncludeOperationalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOperationalAttribute

`func (o *AddCorrelatedLdapDataViewRequest) SetIncludeOperationalAttribute(v []string)`

SetIncludeOperationalAttribute sets IncludeOperationalAttribute field to given value.

### HasIncludeOperationalAttribute

`func (o *AddCorrelatedLdapDataViewRequest) HasIncludeOperationalAttribute() bool`

HasIncludeOperationalAttribute returns a boolean if a field has been set.

### GetCreateDNPattern

`func (o *AddCorrelatedLdapDataViewRequest) GetCreateDNPattern() string`

GetCreateDNPattern returns the CreateDNPattern field if non-nil, zero value otherwise.

### GetCreateDNPatternOk

`func (o *AddCorrelatedLdapDataViewRequest) GetCreateDNPatternOk() (*string, bool)`

GetCreateDNPatternOk returns a tuple with the CreateDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDNPattern

`func (o *AddCorrelatedLdapDataViewRequest) SetCreateDNPattern(v string)`

SetCreateDNPattern sets CreateDNPattern field to given value.

### HasCreateDNPattern

`func (o *AddCorrelatedLdapDataViewRequest) HasCreateDNPattern() bool`

HasCreateDNPattern returns a boolean if a field has been set.

### GetPrimaryCorrelationAttribute

`func (o *AddCorrelatedLdapDataViewRequest) GetPrimaryCorrelationAttribute() string`

GetPrimaryCorrelationAttribute returns the PrimaryCorrelationAttribute field if non-nil, zero value otherwise.

### GetPrimaryCorrelationAttributeOk

`func (o *AddCorrelatedLdapDataViewRequest) GetPrimaryCorrelationAttributeOk() (*string, bool)`

GetPrimaryCorrelationAttributeOk returns a tuple with the PrimaryCorrelationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCorrelationAttribute

`func (o *AddCorrelatedLdapDataViewRequest) SetPrimaryCorrelationAttribute(v string)`

SetPrimaryCorrelationAttribute sets PrimaryCorrelationAttribute field to given value.


### GetSecondaryCorrelationAttribute

`func (o *AddCorrelatedLdapDataViewRequest) GetSecondaryCorrelationAttribute() string`

GetSecondaryCorrelationAttribute returns the SecondaryCorrelationAttribute field if non-nil, zero value otherwise.

### GetSecondaryCorrelationAttributeOk

`func (o *AddCorrelatedLdapDataViewRequest) GetSecondaryCorrelationAttributeOk() (*string, bool)`

GetSecondaryCorrelationAttributeOk returns a tuple with the SecondaryCorrelationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCorrelationAttribute

`func (o *AddCorrelatedLdapDataViewRequest) SetSecondaryCorrelationAttribute(v string)`

SetSecondaryCorrelationAttribute sets SecondaryCorrelationAttribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


