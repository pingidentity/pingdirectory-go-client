# DelegatedAdminRightsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Delegated Admin Rights | 
**Schemas** | Pointer to [**[]EnumdelegatedAdminRightsSchemaUrn**](EnumdelegatedAdminRightsSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Delegated Admin Rights | [optional] 
**Enabled** | **bool** | Indicates whether the Delegated Admin Rights is enabled. | 
**AdminUserDN** | Pointer to **string** | Specifies the DN of an administrative user who has authority to manage resources. Either admin-user-dn or admin-group-dn must be specified, but not both. | [optional] 
**AdminGroupDN** | Pointer to **string** | Specifies the DN of a group of administrative users who have authority to manage resources. Either admin-user-dn or admin-group-dn must be specified, but not both. | [optional] 

## Methods

### NewDelegatedAdminRightsResponse

`func NewDelegatedAdminRightsResponse(id string, enabled bool, ) *DelegatedAdminRightsResponse`

NewDelegatedAdminRightsResponse instantiates a new DelegatedAdminRightsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegatedAdminRightsResponseWithDefaults

`func NewDelegatedAdminRightsResponseWithDefaults() *DelegatedAdminRightsResponse`

NewDelegatedAdminRightsResponseWithDefaults instantiates a new DelegatedAdminRightsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *DelegatedAdminRightsResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DelegatedAdminRightsResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DelegatedAdminRightsResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DelegatedAdminRightsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *DelegatedAdminRightsResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *DelegatedAdminRightsResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *DelegatedAdminRightsResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *DelegatedAdminRightsResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *DelegatedAdminRightsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DelegatedAdminRightsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DelegatedAdminRightsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *DelegatedAdminRightsResponse) GetSchemas() []EnumdelegatedAdminRightsSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DelegatedAdminRightsResponse) GetSchemasOk() (*[]EnumdelegatedAdminRightsSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DelegatedAdminRightsResponse) SetSchemas(v []EnumdelegatedAdminRightsSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *DelegatedAdminRightsResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *DelegatedAdminRightsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DelegatedAdminRightsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DelegatedAdminRightsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DelegatedAdminRightsResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DelegatedAdminRightsResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DelegatedAdminRightsResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DelegatedAdminRightsResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAdminUserDN

`func (o *DelegatedAdminRightsResponse) GetAdminUserDN() string`

GetAdminUserDN returns the AdminUserDN field if non-nil, zero value otherwise.

### GetAdminUserDNOk

`func (o *DelegatedAdminRightsResponse) GetAdminUserDNOk() (*string, bool)`

GetAdminUserDNOk returns a tuple with the AdminUserDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUserDN

`func (o *DelegatedAdminRightsResponse) SetAdminUserDN(v string)`

SetAdminUserDN sets AdminUserDN field to given value.

### HasAdminUserDN

`func (o *DelegatedAdminRightsResponse) HasAdminUserDN() bool`

HasAdminUserDN returns a boolean if a field has been set.

### GetAdminGroupDN

`func (o *DelegatedAdminRightsResponse) GetAdminGroupDN() string`

GetAdminGroupDN returns the AdminGroupDN field if non-nil, zero value otherwise.

### GetAdminGroupDNOk

`func (o *DelegatedAdminRightsResponse) GetAdminGroupDNOk() (*string, bool)`

GetAdminGroupDNOk returns a tuple with the AdminGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminGroupDN

`func (o *DelegatedAdminRightsResponse) SetAdminGroupDN(v string)`

SetAdminGroupDN sets AdminGroupDN field to given value.

### HasAdminGroupDN

`func (o *DelegatedAdminRightsResponse) HasAdminGroupDN() bool`

HasAdminGroupDN returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


