# PasswordInterServerAuthenticationInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumpasswordInterServerAuthenticationInfoSchemaUrn**](EnumpasswordInterServerAuthenticationInfoSchemaUrn.md) |  | 
**Id** | **string** | Name of the Server Instance | 
**AuthenticationType** | Pointer to [**EnuminterServerAuthenticationInfoAuthenticationTypeProp**](EnuminterServerAuthenticationInfoAuthenticationTypeProp.md) |  | [optional] 
**BindDN** | Pointer to **string** | A DN of the username that should be used for the bind request. | [optional] 
**Username** | Pointer to **string** | The username that should be used for the bind request. | [optional] 
**Password** | **string** | The password for the username or bind-dn. | 
**Purpose** | Pointer to [**[]EnuminterServerAuthenticationInfoPurposeProp**](EnuminterServerAuthenticationInfoPurposeProp.md) |  | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewPasswordInterServerAuthenticationInfoResponse

`func NewPasswordInterServerAuthenticationInfoResponse(schemas []EnumpasswordInterServerAuthenticationInfoSchemaUrn, id string, password string, ) *PasswordInterServerAuthenticationInfoResponse`

NewPasswordInterServerAuthenticationInfoResponse instantiates a new PasswordInterServerAuthenticationInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordInterServerAuthenticationInfoResponseWithDefaults

`func NewPasswordInterServerAuthenticationInfoResponseWithDefaults() *PasswordInterServerAuthenticationInfoResponse`

NewPasswordInterServerAuthenticationInfoResponseWithDefaults instantiates a new PasswordInterServerAuthenticationInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *PasswordInterServerAuthenticationInfoResponse) GetSchemas() []EnumpasswordInterServerAuthenticationInfoSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PasswordInterServerAuthenticationInfoResponse) GetSchemasOk() (*[]EnumpasswordInterServerAuthenticationInfoSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PasswordInterServerAuthenticationInfoResponse) SetSchemas(v []EnumpasswordInterServerAuthenticationInfoSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *PasswordInterServerAuthenticationInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PasswordInterServerAuthenticationInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PasswordInterServerAuthenticationInfoResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAuthenticationType

`func (o *PasswordInterServerAuthenticationInfoResponse) GetAuthenticationType() EnuminterServerAuthenticationInfoAuthenticationTypeProp`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *PasswordInterServerAuthenticationInfoResponse) GetAuthenticationTypeOk() (*EnuminterServerAuthenticationInfoAuthenticationTypeProp, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *PasswordInterServerAuthenticationInfoResponse) SetAuthenticationType(v EnuminterServerAuthenticationInfoAuthenticationTypeProp)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *PasswordInterServerAuthenticationInfoResponse) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetBindDN

`func (o *PasswordInterServerAuthenticationInfoResponse) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *PasswordInterServerAuthenticationInfoResponse) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *PasswordInterServerAuthenticationInfoResponse) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.

### HasBindDN

`func (o *PasswordInterServerAuthenticationInfoResponse) HasBindDN() bool`

HasBindDN returns a boolean if a field has been set.

### GetUsername

`func (o *PasswordInterServerAuthenticationInfoResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PasswordInterServerAuthenticationInfoResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PasswordInterServerAuthenticationInfoResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PasswordInterServerAuthenticationInfoResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *PasswordInterServerAuthenticationInfoResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PasswordInterServerAuthenticationInfoResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PasswordInterServerAuthenticationInfoResponse) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPurpose

`func (o *PasswordInterServerAuthenticationInfoResponse) GetPurpose() []EnuminterServerAuthenticationInfoPurposeProp`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *PasswordInterServerAuthenticationInfoResponse) GetPurposeOk() (*[]EnuminterServerAuthenticationInfoPurposeProp, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *PasswordInterServerAuthenticationInfoResponse) SetPurpose(v []EnuminterServerAuthenticationInfoPurposeProp)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *PasswordInterServerAuthenticationInfoResponse) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetMeta

`func (o *PasswordInterServerAuthenticationInfoResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PasswordInterServerAuthenticationInfoResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PasswordInterServerAuthenticationInfoResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PasswordInterServerAuthenticationInfoResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *PasswordInterServerAuthenticationInfoResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *PasswordInterServerAuthenticationInfoResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *PasswordInterServerAuthenticationInfoResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *PasswordInterServerAuthenticationInfoResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


