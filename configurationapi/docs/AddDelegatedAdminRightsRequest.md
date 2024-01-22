# AddDelegatedAdminRightsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumdelegatedAdminRightsSchemaUrn**](EnumdelegatedAdminRightsSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Delegated Admin Rights | [optional] 
**Enabled** | **bool** | Indicates whether the Delegated Admin Rights is enabled. | 
**AdminUserDN** | Pointer to **string** | Specifies the DN of an administrative user who has authority to manage resources. Either admin-user-dn or admin-group-dn must be specified, but not both. | [optional] 
**AdminGroupDN** | Pointer to **string** | Specifies the DN of a group of administrative users who have authority to manage resources. Either admin-user-dn or admin-group-dn must be specified, but not both. | [optional] 
**RightsName** | **string** | Name of the new Delegated Admin Rights | 

## Methods

### NewAddDelegatedAdminRightsRequest

`func NewAddDelegatedAdminRightsRequest(enabled bool, rightsName string, ) *AddDelegatedAdminRightsRequest`

NewAddDelegatedAdminRightsRequest instantiates a new AddDelegatedAdminRightsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDelegatedAdminRightsRequestWithDefaults

`func NewAddDelegatedAdminRightsRequestWithDefaults() *AddDelegatedAdminRightsRequest`

NewAddDelegatedAdminRightsRequestWithDefaults instantiates a new AddDelegatedAdminRightsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddDelegatedAdminRightsRequest) GetSchemas() []EnumdelegatedAdminRightsSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDelegatedAdminRightsRequest) GetSchemasOk() (*[]EnumdelegatedAdminRightsSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDelegatedAdminRightsRequest) SetSchemas(v []EnumdelegatedAdminRightsSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddDelegatedAdminRightsRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddDelegatedAdminRightsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddDelegatedAdminRightsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddDelegatedAdminRightsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddDelegatedAdminRightsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddDelegatedAdminRightsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddDelegatedAdminRightsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddDelegatedAdminRightsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAdminUserDN

`func (o *AddDelegatedAdminRightsRequest) GetAdminUserDN() string`

GetAdminUserDN returns the AdminUserDN field if non-nil, zero value otherwise.

### GetAdminUserDNOk

`func (o *AddDelegatedAdminRightsRequest) GetAdminUserDNOk() (*string, bool)`

GetAdminUserDNOk returns a tuple with the AdminUserDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUserDN

`func (o *AddDelegatedAdminRightsRequest) SetAdminUserDN(v string)`

SetAdminUserDN sets AdminUserDN field to given value.

### HasAdminUserDN

`func (o *AddDelegatedAdminRightsRequest) HasAdminUserDN() bool`

HasAdminUserDN returns a boolean if a field has been set.

### GetAdminGroupDN

`func (o *AddDelegatedAdminRightsRequest) GetAdminGroupDN() string`

GetAdminGroupDN returns the AdminGroupDN field if non-nil, zero value otherwise.

### GetAdminGroupDNOk

`func (o *AddDelegatedAdminRightsRequest) GetAdminGroupDNOk() (*string, bool)`

GetAdminGroupDNOk returns a tuple with the AdminGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminGroupDN

`func (o *AddDelegatedAdminRightsRequest) SetAdminGroupDN(v string)`

SetAdminGroupDN sets AdminGroupDN field to given value.

### HasAdminGroupDN

`func (o *AddDelegatedAdminRightsRequest) HasAdminGroupDN() bool`

HasAdminGroupDN returns a boolean if a field has been set.

### GetRightsName

`func (o *AddDelegatedAdminRightsRequest) GetRightsName() string`

GetRightsName returns the RightsName field if non-nil, zero value otherwise.

### GetRightsNameOk

`func (o *AddDelegatedAdminRightsRequest) GetRightsNameOk() (*string, bool)`

GetRightsNameOk returns a tuple with the RightsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightsName

`func (o *AddDelegatedAdminRightsRequest) SetRightsName(v string)`

SetRightsName sets RightsName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


