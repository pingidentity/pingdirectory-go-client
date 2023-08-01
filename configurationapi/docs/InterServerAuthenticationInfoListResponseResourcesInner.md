# InterServerAuthenticationInfoListResponseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumpasswordInterServerAuthenticationInfoSchemaUrn**](EnumpasswordInterServerAuthenticationInfoSchemaUrn.md) |  | 
**Id** | **string** | Name of the Inter Server Authentication Info | 
**Purpose** | Pointer to [**[]EnuminterServerAuthenticationInfoPurposeProp**](EnuminterServerAuthenticationInfoPurposeProp.md) |  | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**AuthenticationType** | Pointer to [**EnuminterServerAuthenticationInfoAuthenticationTypeProp**](EnuminterServerAuthenticationInfoAuthenticationTypeProp.md) |  | [optional] 
**BindDN** | Pointer to **string** | A DN of the username that should be used for the bind request. | [optional] 
**Username** | Pointer to **string** | The username that should be used for the bind request. | [optional] 
**Password** | **string** | The password for the username or bind-dn. | 

## Methods

### NewInterServerAuthenticationInfoListResponseResourcesInner

`func NewInterServerAuthenticationInfoListResponseResourcesInner(schemas []EnumpasswordInterServerAuthenticationInfoSchemaUrn, id string, password string, ) *InterServerAuthenticationInfoListResponseResourcesInner`

NewInterServerAuthenticationInfoListResponseResourcesInner instantiates a new InterServerAuthenticationInfoListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterServerAuthenticationInfoListResponseResourcesInnerWithDefaults

`func NewInterServerAuthenticationInfoListResponseResourcesInnerWithDefaults() *InterServerAuthenticationInfoListResponseResourcesInner`

NewInterServerAuthenticationInfoListResponseResourcesInnerWithDefaults instantiates a new InterServerAuthenticationInfoListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetSchemas() []EnumpasswordInterServerAuthenticationInfoSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetSchemasOk() (*[]EnumpasswordInterServerAuthenticationInfoSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) SetSchemas(v []EnumpasswordInterServerAuthenticationInfoSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetPurpose

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetPurpose() []EnuminterServerAuthenticationInfoPurposeProp`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetPurposeOk() (*[]EnuminterServerAuthenticationInfoPurposeProp, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) SetPurpose(v []EnuminterServerAuthenticationInfoPurposeProp)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetMeta

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetAuthenticationType() EnuminterServerAuthenticationInfoAuthenticationTypeProp`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetAuthenticationTypeOk() (*EnuminterServerAuthenticationInfoAuthenticationTypeProp, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) SetAuthenticationType(v EnuminterServerAuthenticationInfoAuthenticationTypeProp)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetBindDN

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.

### HasBindDN

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) HasBindDN() bool`

HasBindDN returns a boolean if a field has been set.

### GetUsername

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InterServerAuthenticationInfoListResponseResourcesInner) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


