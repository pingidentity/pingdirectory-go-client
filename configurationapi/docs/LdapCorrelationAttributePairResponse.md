# LdapCorrelationAttributePairResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the LDAP Correlation Attribute Pair | 
**Schemas** | Pointer to [**[]EnumldapCorrelationAttributePairSchemaUrn**](EnumldapCorrelationAttributePairSchemaUrn.md) |  | [optional] 
**PrimaryCorrelationAttribute** | **string** | The LDAP attribute from the base SCIM Resource Type whose value will be used to match objects in the Correlated LDAP Data View. | 
**SecondaryCorrelationAttribute** | **string** | The LDAP attribute from the Correlated LDAP Data View whose value will be matched. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewLdapCorrelationAttributePairResponse

`func NewLdapCorrelationAttributePairResponse(id string, primaryCorrelationAttribute string, secondaryCorrelationAttribute string, ) *LdapCorrelationAttributePairResponse`

NewLdapCorrelationAttributePairResponse instantiates a new LdapCorrelationAttributePairResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapCorrelationAttributePairResponseWithDefaults

`func NewLdapCorrelationAttributePairResponseWithDefaults() *LdapCorrelationAttributePairResponse`

NewLdapCorrelationAttributePairResponseWithDefaults instantiates a new LdapCorrelationAttributePairResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LdapCorrelationAttributePairResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LdapCorrelationAttributePairResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LdapCorrelationAttributePairResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *LdapCorrelationAttributePairResponse) GetSchemas() []EnumldapCorrelationAttributePairSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LdapCorrelationAttributePairResponse) GetSchemasOk() (*[]EnumldapCorrelationAttributePairSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LdapCorrelationAttributePairResponse) SetSchemas(v []EnumldapCorrelationAttributePairSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *LdapCorrelationAttributePairResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetPrimaryCorrelationAttribute

`func (o *LdapCorrelationAttributePairResponse) GetPrimaryCorrelationAttribute() string`

GetPrimaryCorrelationAttribute returns the PrimaryCorrelationAttribute field if non-nil, zero value otherwise.

### GetPrimaryCorrelationAttributeOk

`func (o *LdapCorrelationAttributePairResponse) GetPrimaryCorrelationAttributeOk() (*string, bool)`

GetPrimaryCorrelationAttributeOk returns a tuple with the PrimaryCorrelationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCorrelationAttribute

`func (o *LdapCorrelationAttributePairResponse) SetPrimaryCorrelationAttribute(v string)`

SetPrimaryCorrelationAttribute sets PrimaryCorrelationAttribute field to given value.


### GetSecondaryCorrelationAttribute

`func (o *LdapCorrelationAttributePairResponse) GetSecondaryCorrelationAttribute() string`

GetSecondaryCorrelationAttribute returns the SecondaryCorrelationAttribute field if non-nil, zero value otherwise.

### GetSecondaryCorrelationAttributeOk

`func (o *LdapCorrelationAttributePairResponse) GetSecondaryCorrelationAttributeOk() (*string, bool)`

GetSecondaryCorrelationAttributeOk returns a tuple with the SecondaryCorrelationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCorrelationAttribute

`func (o *LdapCorrelationAttributePairResponse) SetSecondaryCorrelationAttribute(v string)`

SetSecondaryCorrelationAttribute sets SecondaryCorrelationAttribute field to given value.


### GetMeta

`func (o *LdapCorrelationAttributePairResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LdapCorrelationAttributePairResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LdapCorrelationAttributePairResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LdapCorrelationAttributePairResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *LdapCorrelationAttributePairResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *LdapCorrelationAttributePairResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *LdapCorrelationAttributePairResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *LdapCorrelationAttributePairResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


