# AddScimAttributeMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumscimAttributeMappingSchemaUrn**](EnumscimAttributeMappingSchemaUrn.md) |  | [optional] 
**CorrelatedLDAPDataView** | Pointer to **string** | The Correlated LDAP Data View that persists the mapped SCIM Resource Type attribute(s). | [optional] 
**ScimResourceTypeAttribute** | **string** | The attribute path of SCIM Resource Type attributes to be mapped. | 
**LdapAttribute** | **string** | The LDAP attribute to be mapped, or the path to a specific field of an LDAP attribute with the JSON object attribute syntax. | 
**Readable** | Pointer to **bool** | Specifies whether the mapping is used to map from LDAP attribute to SCIM Resource Type attribute in a read operation. | [optional] 
**Writable** | Pointer to **bool** | Specifies that the mapping is used to map from SCIM Resource Type attribute to LDAP attribute in a write operation. | [optional] 
**Searchable** | Pointer to **bool** | Specifies that the mapping is used to map from SCIM Resource Type attribute to LDAP attribute in a search filter. | [optional] 
**Authoritative** | Pointer to **bool** | Specifies that the mapping is authoritative over other mappings for the same SCIM Resource Type attribute (for read operations). | [optional] 
**MappingName** | **string** | Name of the new SCIM Attribute Mapping | 

## Methods

### NewAddScimAttributeMappingRequest

`func NewAddScimAttributeMappingRequest(scimResourceTypeAttribute string, ldapAttribute string, mappingName string, ) *AddScimAttributeMappingRequest`

NewAddScimAttributeMappingRequest instantiates a new AddScimAttributeMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddScimAttributeMappingRequestWithDefaults

`func NewAddScimAttributeMappingRequestWithDefaults() *AddScimAttributeMappingRequest`

NewAddScimAttributeMappingRequestWithDefaults instantiates a new AddScimAttributeMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddScimAttributeMappingRequest) GetSchemas() []EnumscimAttributeMappingSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddScimAttributeMappingRequest) GetSchemasOk() (*[]EnumscimAttributeMappingSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddScimAttributeMappingRequest) SetSchemas(v []EnumscimAttributeMappingSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddScimAttributeMappingRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetCorrelatedLDAPDataView

`func (o *AddScimAttributeMappingRequest) GetCorrelatedLDAPDataView() string`

GetCorrelatedLDAPDataView returns the CorrelatedLDAPDataView field if non-nil, zero value otherwise.

### GetCorrelatedLDAPDataViewOk

`func (o *AddScimAttributeMappingRequest) GetCorrelatedLDAPDataViewOk() (*string, bool)`

GetCorrelatedLDAPDataViewOk returns a tuple with the CorrelatedLDAPDataView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedLDAPDataView

`func (o *AddScimAttributeMappingRequest) SetCorrelatedLDAPDataView(v string)`

SetCorrelatedLDAPDataView sets CorrelatedLDAPDataView field to given value.

### HasCorrelatedLDAPDataView

`func (o *AddScimAttributeMappingRequest) HasCorrelatedLDAPDataView() bool`

HasCorrelatedLDAPDataView returns a boolean if a field has been set.

### GetScimResourceTypeAttribute

`func (o *AddScimAttributeMappingRequest) GetScimResourceTypeAttribute() string`

GetScimResourceTypeAttribute returns the ScimResourceTypeAttribute field if non-nil, zero value otherwise.

### GetScimResourceTypeAttributeOk

`func (o *AddScimAttributeMappingRequest) GetScimResourceTypeAttributeOk() (*string, bool)`

GetScimResourceTypeAttributeOk returns a tuple with the ScimResourceTypeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimResourceTypeAttribute

`func (o *AddScimAttributeMappingRequest) SetScimResourceTypeAttribute(v string)`

SetScimResourceTypeAttribute sets ScimResourceTypeAttribute field to given value.


### GetLdapAttribute

`func (o *AddScimAttributeMappingRequest) GetLdapAttribute() string`

GetLdapAttribute returns the LdapAttribute field if non-nil, zero value otherwise.

### GetLdapAttributeOk

`func (o *AddScimAttributeMappingRequest) GetLdapAttributeOk() (*string, bool)`

GetLdapAttributeOk returns a tuple with the LdapAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAttribute

`func (o *AddScimAttributeMappingRequest) SetLdapAttribute(v string)`

SetLdapAttribute sets LdapAttribute field to given value.


### GetReadable

`func (o *AddScimAttributeMappingRequest) GetReadable() bool`

GetReadable returns the Readable field if non-nil, zero value otherwise.

### GetReadableOk

`func (o *AddScimAttributeMappingRequest) GetReadableOk() (*bool, bool)`

GetReadableOk returns a tuple with the Readable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadable

`func (o *AddScimAttributeMappingRequest) SetReadable(v bool)`

SetReadable sets Readable field to given value.

### HasReadable

`func (o *AddScimAttributeMappingRequest) HasReadable() bool`

HasReadable returns a boolean if a field has been set.

### GetWritable

`func (o *AddScimAttributeMappingRequest) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *AddScimAttributeMappingRequest) GetWritableOk() (*bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritable

`func (o *AddScimAttributeMappingRequest) SetWritable(v bool)`

SetWritable sets Writable field to given value.

### HasWritable

`func (o *AddScimAttributeMappingRequest) HasWritable() bool`

HasWritable returns a boolean if a field has been set.

### GetSearchable

`func (o *AddScimAttributeMappingRequest) GetSearchable() bool`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *AddScimAttributeMappingRequest) GetSearchableOk() (*bool, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *AddScimAttributeMappingRequest) SetSearchable(v bool)`

SetSearchable sets Searchable field to given value.

### HasSearchable

`func (o *AddScimAttributeMappingRequest) HasSearchable() bool`

HasSearchable returns a boolean if a field has been set.

### GetAuthoritative

`func (o *AddScimAttributeMappingRequest) GetAuthoritative() bool`

GetAuthoritative returns the Authoritative field if non-nil, zero value otherwise.

### GetAuthoritativeOk

`func (o *AddScimAttributeMappingRequest) GetAuthoritativeOk() (*bool, bool)`

GetAuthoritativeOk returns a tuple with the Authoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoritative

`func (o *AddScimAttributeMappingRequest) SetAuthoritative(v bool)`

SetAuthoritative sets Authoritative field to given value.

### HasAuthoritative

`func (o *AddScimAttributeMappingRequest) HasAuthoritative() bool`

HasAuthoritative returns a boolean if a field has been set.

### GetMappingName

`func (o *AddScimAttributeMappingRequest) GetMappingName() string`

GetMappingName returns the MappingName field if non-nil, zero value otherwise.

### GetMappingNameOk

`func (o *AddScimAttributeMappingRequest) GetMappingNameOk() (*string, bool)`

GetMappingNameOk returns a tuple with the MappingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingName

`func (o *AddScimAttributeMappingRequest) SetMappingName(v string)`

SetMappingName sets MappingName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


