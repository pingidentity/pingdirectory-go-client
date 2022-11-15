# DelegatedAdminRightsShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumdelegatedAdminRightsSchemaUrn**](EnumdelegatedAdminRightsSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Delegated Admin Rights | [optional] 
**Enabled** | **bool** | Indicates whether the Delegated Admin Rights is enabled. | 
**AdminUserDN** | Pointer to **string** | Specifies the DN of an administrative user who has authority to manage resources. Either admin-user-dn or admin-group-dn must be specified, but not both. | [optional] 
**AdminGroupDN** | Pointer to **string** | Specifies the DN of a group of administrative users who have authority to manage resources. Either admin-user-dn or admin-group-dn must be specified, but not both. | [optional] 

## Methods

### NewDelegatedAdminRightsShared

`func NewDelegatedAdminRightsShared(enabled bool, ) *DelegatedAdminRightsShared`

NewDelegatedAdminRightsShared instantiates a new DelegatedAdminRightsShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegatedAdminRightsSharedWithDefaults

`func NewDelegatedAdminRightsSharedWithDefaults() *DelegatedAdminRightsShared`

NewDelegatedAdminRightsSharedWithDefaults instantiates a new DelegatedAdminRightsShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *DelegatedAdminRightsShared) GetSchemas() []EnumdelegatedAdminRightsSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DelegatedAdminRightsShared) GetSchemasOk() (*[]EnumdelegatedAdminRightsSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DelegatedAdminRightsShared) SetSchemas(v []EnumdelegatedAdminRightsSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *DelegatedAdminRightsShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *DelegatedAdminRightsShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DelegatedAdminRightsShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DelegatedAdminRightsShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DelegatedAdminRightsShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DelegatedAdminRightsShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DelegatedAdminRightsShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DelegatedAdminRightsShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAdminUserDN

`func (o *DelegatedAdminRightsShared) GetAdminUserDN() string`

GetAdminUserDN returns the AdminUserDN field if non-nil, zero value otherwise.

### GetAdminUserDNOk

`func (o *DelegatedAdminRightsShared) GetAdminUserDNOk() (*string, bool)`

GetAdminUserDNOk returns a tuple with the AdminUserDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUserDN

`func (o *DelegatedAdminRightsShared) SetAdminUserDN(v string)`

SetAdminUserDN sets AdminUserDN field to given value.

### HasAdminUserDN

`func (o *DelegatedAdminRightsShared) HasAdminUserDN() bool`

HasAdminUserDN returns a boolean if a field has been set.

### GetAdminGroupDN

`func (o *DelegatedAdminRightsShared) GetAdminGroupDN() string`

GetAdminGroupDN returns the AdminGroupDN field if non-nil, zero value otherwise.

### GetAdminGroupDNOk

`func (o *DelegatedAdminRightsShared) GetAdminGroupDNOk() (*string, bool)`

GetAdminGroupDNOk returns a tuple with the AdminGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminGroupDN

`func (o *DelegatedAdminRightsShared) SetAdminGroupDN(v string)`

SetAdminGroupDN sets AdminGroupDN field to given value.

### HasAdminGroupDN

`func (o *DelegatedAdminRightsShared) HasAdminGroupDN() bool`

HasAdminGroupDN returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


