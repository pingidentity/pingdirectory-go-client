# ConsentServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumconsentServiceSchemaUrn**](EnumconsentServiceSchemaUrn.md) |  | [optional] 
**Enabled** | **bool** | Indicates whether the Consent Service is enabled. | 
**BaseDN** | Pointer to **string** | The base DN under which consent records are stored. | [optional] 
**BindDN** | Pointer to **string** | The DN of an internal service account used by the Consent Service to make internal LDAP requests. | [optional] 
**SearchSizeLimit** | Pointer to **int32** | The maximum number of consent resources that may be returned from a search request. | [optional] 
**ConsentRecordIdentityMapper** | Pointer to **[]string** | If specified, the Identity Mapper(s) that may be used to map consent record subject and actor values to DNs. This is typically only needed if privileged API clients will be used. | [optional] 
**ServiceAccountDN** | Pointer to **[]string** | The set of account DNs that the Consent Service will consider to be privileged. | [optional] 
**UnprivilegedConsentScope** | Pointer to **string** | The name of a scope that must be present in an access token accepted by the Consent Service for unprivileged clients. | [optional] 
**PrivilegedConsentScope** | Pointer to **string** | The name of a scope that must be present in an access token accepted by the Consent Service if the client is to be considered privileged. | [optional] 
**Audience** | Pointer to **string** | A string or URI that identifies the Consent Service in the context of OAuth2 authorization. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewConsentServiceResponse

`func NewConsentServiceResponse(enabled bool, ) *ConsentServiceResponse`

NewConsentServiceResponse instantiates a new ConsentServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentServiceResponseWithDefaults

`func NewConsentServiceResponseWithDefaults() *ConsentServiceResponse`

NewConsentServiceResponseWithDefaults instantiates a new ConsentServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ConsentServiceResponse) GetSchemas() []EnumconsentServiceSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ConsentServiceResponse) GetSchemasOk() (*[]EnumconsentServiceSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ConsentServiceResponse) SetSchemas(v []EnumconsentServiceSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ConsentServiceResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetEnabled

`func (o *ConsentServiceResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConsentServiceResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConsentServiceResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetBaseDN

`func (o *ConsentServiceResponse) GetBaseDN() string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *ConsentServiceResponse) GetBaseDNOk() (*string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *ConsentServiceResponse) SetBaseDN(v string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *ConsentServiceResponse) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetBindDN

`func (o *ConsentServiceResponse) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *ConsentServiceResponse) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *ConsentServiceResponse) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.

### HasBindDN

`func (o *ConsentServiceResponse) HasBindDN() bool`

HasBindDN returns a boolean if a field has been set.

### GetSearchSizeLimit

`func (o *ConsentServiceResponse) GetSearchSizeLimit() int32`

GetSearchSizeLimit returns the SearchSizeLimit field if non-nil, zero value otherwise.

### GetSearchSizeLimitOk

`func (o *ConsentServiceResponse) GetSearchSizeLimitOk() (*int32, bool)`

GetSearchSizeLimitOk returns a tuple with the SearchSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchSizeLimit

`func (o *ConsentServiceResponse) SetSearchSizeLimit(v int32)`

SetSearchSizeLimit sets SearchSizeLimit field to given value.

### HasSearchSizeLimit

`func (o *ConsentServiceResponse) HasSearchSizeLimit() bool`

HasSearchSizeLimit returns a boolean if a field has been set.

### GetConsentRecordIdentityMapper

`func (o *ConsentServiceResponse) GetConsentRecordIdentityMapper() []string`

GetConsentRecordIdentityMapper returns the ConsentRecordIdentityMapper field if non-nil, zero value otherwise.

### GetConsentRecordIdentityMapperOk

`func (o *ConsentServiceResponse) GetConsentRecordIdentityMapperOk() (*[]string, bool)`

GetConsentRecordIdentityMapperOk returns a tuple with the ConsentRecordIdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentRecordIdentityMapper

`func (o *ConsentServiceResponse) SetConsentRecordIdentityMapper(v []string)`

SetConsentRecordIdentityMapper sets ConsentRecordIdentityMapper field to given value.

### HasConsentRecordIdentityMapper

`func (o *ConsentServiceResponse) HasConsentRecordIdentityMapper() bool`

HasConsentRecordIdentityMapper returns a boolean if a field has been set.

### GetServiceAccountDN

`func (o *ConsentServiceResponse) GetServiceAccountDN() []string`

GetServiceAccountDN returns the ServiceAccountDN field if non-nil, zero value otherwise.

### GetServiceAccountDNOk

`func (o *ConsentServiceResponse) GetServiceAccountDNOk() (*[]string, bool)`

GetServiceAccountDNOk returns a tuple with the ServiceAccountDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountDN

`func (o *ConsentServiceResponse) SetServiceAccountDN(v []string)`

SetServiceAccountDN sets ServiceAccountDN field to given value.

### HasServiceAccountDN

`func (o *ConsentServiceResponse) HasServiceAccountDN() bool`

HasServiceAccountDN returns a boolean if a field has been set.

### GetUnprivilegedConsentScope

`func (o *ConsentServiceResponse) GetUnprivilegedConsentScope() string`

GetUnprivilegedConsentScope returns the UnprivilegedConsentScope field if non-nil, zero value otherwise.

### GetUnprivilegedConsentScopeOk

`func (o *ConsentServiceResponse) GetUnprivilegedConsentScopeOk() (*string, bool)`

GetUnprivilegedConsentScopeOk returns a tuple with the UnprivilegedConsentScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprivilegedConsentScope

`func (o *ConsentServiceResponse) SetUnprivilegedConsentScope(v string)`

SetUnprivilegedConsentScope sets UnprivilegedConsentScope field to given value.

### HasUnprivilegedConsentScope

`func (o *ConsentServiceResponse) HasUnprivilegedConsentScope() bool`

HasUnprivilegedConsentScope returns a boolean if a field has been set.

### GetPrivilegedConsentScope

`func (o *ConsentServiceResponse) GetPrivilegedConsentScope() string`

GetPrivilegedConsentScope returns the PrivilegedConsentScope field if non-nil, zero value otherwise.

### GetPrivilegedConsentScopeOk

`func (o *ConsentServiceResponse) GetPrivilegedConsentScopeOk() (*string, bool)`

GetPrivilegedConsentScopeOk returns a tuple with the PrivilegedConsentScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedConsentScope

`func (o *ConsentServiceResponse) SetPrivilegedConsentScope(v string)`

SetPrivilegedConsentScope sets PrivilegedConsentScope field to given value.

### HasPrivilegedConsentScope

`func (o *ConsentServiceResponse) HasPrivilegedConsentScope() bool`

HasPrivilegedConsentScope returns a boolean if a field has been set.

### GetAudience

`func (o *ConsentServiceResponse) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *ConsentServiceResponse) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *ConsentServiceResponse) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *ConsentServiceResponse) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetMeta

`func (o *ConsentServiceResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConsentServiceResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConsentServiceResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ConsentServiceResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ConsentServiceResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ConsentServiceResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ConsentServiceResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ConsentServiceResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


