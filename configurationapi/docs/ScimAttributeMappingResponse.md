# ScimAttributeMappingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the SCIM Resource Type | 
**Schemas** | Pointer to [**[]EnumscimAttributeMappingSchemaUrn**](EnumscimAttributeMappingSchemaUrn.md) |  | [optional] 
**CorrelatedLDAPDataView** | Pointer to **string** | The Correlated LDAP Data View that persists the mapped SCIM Resource Type attribute(s). | [optional] 
**ScimResourceTypeAttribute** | **string** | The attribute path of SCIM Resource Type attributes to be mapped. | 
**LdapAttribute** | **string** | The LDAP attribute to be mapped, or the path to a specific field of an LDAP attribute with the JSON object attribute syntax. | 
**Readable** | Pointer to **bool** | Specifies whether the mapping is used to map from LDAP attribute to SCIM Resource Type attribute in a read operation. | [optional] 
**Writable** | Pointer to **bool** | Specifies that the mapping is used to map from SCIM Resource Type attribute to LDAP attribute in a write operation. | [optional] 
**Searchable** | Pointer to **bool** | Specifies that the mapping is used to map from SCIM Resource Type attribute to LDAP attribute in a search filter. | [optional] 
**Authoritative** | Pointer to **bool** | Specifies that the mapping is authoritative over other mappings for the same SCIM Resource Type attribute (for read operations). | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewScimAttributeMappingResponse

`func NewScimAttributeMappingResponse(id string, scimResourceTypeAttribute string, ldapAttribute string, ) *ScimAttributeMappingResponse`

NewScimAttributeMappingResponse instantiates a new ScimAttributeMappingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimAttributeMappingResponseWithDefaults

`func NewScimAttributeMappingResponseWithDefaults() *ScimAttributeMappingResponse`

NewScimAttributeMappingResponseWithDefaults instantiates a new ScimAttributeMappingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScimAttributeMappingResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimAttributeMappingResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimAttributeMappingResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ScimAttributeMappingResponse) GetSchemas() []EnumscimAttributeMappingSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimAttributeMappingResponse) GetSchemasOk() (*[]EnumscimAttributeMappingSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimAttributeMappingResponse) SetSchemas(v []EnumscimAttributeMappingSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ScimAttributeMappingResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetCorrelatedLDAPDataView

`func (o *ScimAttributeMappingResponse) GetCorrelatedLDAPDataView() string`

GetCorrelatedLDAPDataView returns the CorrelatedLDAPDataView field if non-nil, zero value otherwise.

### GetCorrelatedLDAPDataViewOk

`func (o *ScimAttributeMappingResponse) GetCorrelatedLDAPDataViewOk() (*string, bool)`

GetCorrelatedLDAPDataViewOk returns a tuple with the CorrelatedLDAPDataView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedLDAPDataView

`func (o *ScimAttributeMappingResponse) SetCorrelatedLDAPDataView(v string)`

SetCorrelatedLDAPDataView sets CorrelatedLDAPDataView field to given value.

### HasCorrelatedLDAPDataView

`func (o *ScimAttributeMappingResponse) HasCorrelatedLDAPDataView() bool`

HasCorrelatedLDAPDataView returns a boolean if a field has been set.

### GetScimResourceTypeAttribute

`func (o *ScimAttributeMappingResponse) GetScimResourceTypeAttribute() string`

GetScimResourceTypeAttribute returns the ScimResourceTypeAttribute field if non-nil, zero value otherwise.

### GetScimResourceTypeAttributeOk

`func (o *ScimAttributeMappingResponse) GetScimResourceTypeAttributeOk() (*string, bool)`

GetScimResourceTypeAttributeOk returns a tuple with the ScimResourceTypeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimResourceTypeAttribute

`func (o *ScimAttributeMappingResponse) SetScimResourceTypeAttribute(v string)`

SetScimResourceTypeAttribute sets ScimResourceTypeAttribute field to given value.


### GetLdapAttribute

`func (o *ScimAttributeMappingResponse) GetLdapAttribute() string`

GetLdapAttribute returns the LdapAttribute field if non-nil, zero value otherwise.

### GetLdapAttributeOk

`func (o *ScimAttributeMappingResponse) GetLdapAttributeOk() (*string, bool)`

GetLdapAttributeOk returns a tuple with the LdapAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAttribute

`func (o *ScimAttributeMappingResponse) SetLdapAttribute(v string)`

SetLdapAttribute sets LdapAttribute field to given value.


### GetReadable

`func (o *ScimAttributeMappingResponse) GetReadable() bool`

GetReadable returns the Readable field if non-nil, zero value otherwise.

### GetReadableOk

`func (o *ScimAttributeMappingResponse) GetReadableOk() (*bool, bool)`

GetReadableOk returns a tuple with the Readable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadable

`func (o *ScimAttributeMappingResponse) SetReadable(v bool)`

SetReadable sets Readable field to given value.

### HasReadable

`func (o *ScimAttributeMappingResponse) HasReadable() bool`

HasReadable returns a boolean if a field has been set.

### GetWritable

`func (o *ScimAttributeMappingResponse) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *ScimAttributeMappingResponse) GetWritableOk() (*bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritable

`func (o *ScimAttributeMappingResponse) SetWritable(v bool)`

SetWritable sets Writable field to given value.

### HasWritable

`func (o *ScimAttributeMappingResponse) HasWritable() bool`

HasWritable returns a boolean if a field has been set.

### GetSearchable

`func (o *ScimAttributeMappingResponse) GetSearchable() bool`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *ScimAttributeMappingResponse) GetSearchableOk() (*bool, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *ScimAttributeMappingResponse) SetSearchable(v bool)`

SetSearchable sets Searchable field to given value.

### HasSearchable

`func (o *ScimAttributeMappingResponse) HasSearchable() bool`

HasSearchable returns a boolean if a field has been set.

### GetAuthoritative

`func (o *ScimAttributeMappingResponse) GetAuthoritative() bool`

GetAuthoritative returns the Authoritative field if non-nil, zero value otherwise.

### GetAuthoritativeOk

`func (o *ScimAttributeMappingResponse) GetAuthoritativeOk() (*bool, bool)`

GetAuthoritativeOk returns a tuple with the Authoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoritative

`func (o *ScimAttributeMappingResponse) SetAuthoritative(v bool)`

SetAuthoritative sets Authoritative field to given value.

### HasAuthoritative

`func (o *ScimAttributeMappingResponse) HasAuthoritative() bool`

HasAuthoritative returns a boolean if a field has been set.

### GetMeta

`func (o *ScimAttributeMappingResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScimAttributeMappingResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScimAttributeMappingResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ScimAttributeMappingResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ScimAttributeMappingResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ScimAttributeMappingResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ScimAttributeMappingResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ScimAttributeMappingResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


