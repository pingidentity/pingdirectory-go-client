# GetInterServerAuthenticationInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumcertificateInterServerAuthenticationInfoSchemaUrn**](EnumcertificateInterServerAuthenticationInfoSchemaUrn.md) |  | 
**Id** | **string** | Name of the Inter Server Authentication Info | 
**AuthenticationType** | Pointer to [**EnuminterServerAuthenticationInfoAuthenticationTypeProp**](EnuminterServerAuthenticationInfoAuthenticationTypeProp.md) |  | [optional] 
**BindDN** | Pointer to **string** | A DN of the username that should be used for the bind request. | [optional] 
**Username** | Pointer to **string** | The username that should be used for the bind request. | [optional] 
**Password** | **string** | The password for the username or bind-dn. | 
**Purpose** | Pointer to [**[]EnuminterServerAuthenticationInfoPurposeProp**](EnuminterServerAuthenticationInfoPurposeProp.md) |  | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewGetInterServerAuthenticationInfo200Response

`func NewGetInterServerAuthenticationInfo200Response(schemas []EnumcertificateInterServerAuthenticationInfoSchemaUrn, id string, password string, ) *GetInterServerAuthenticationInfo200Response`

NewGetInterServerAuthenticationInfo200Response instantiates a new GetInterServerAuthenticationInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInterServerAuthenticationInfo200ResponseWithDefaults

`func NewGetInterServerAuthenticationInfo200ResponseWithDefaults() *GetInterServerAuthenticationInfo200Response`

NewGetInterServerAuthenticationInfo200ResponseWithDefaults instantiates a new GetInterServerAuthenticationInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GetInterServerAuthenticationInfo200Response) GetSchemas() []EnumcertificateInterServerAuthenticationInfoSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetInterServerAuthenticationInfo200Response) GetSchemasOk() (*[]EnumcertificateInterServerAuthenticationInfoSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetInterServerAuthenticationInfo200Response) SetSchemas(v []EnumcertificateInterServerAuthenticationInfoSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GetInterServerAuthenticationInfo200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInterServerAuthenticationInfo200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInterServerAuthenticationInfo200Response) SetId(v string)`

SetId sets Id field to given value.


### GetAuthenticationType

`func (o *GetInterServerAuthenticationInfo200Response) GetAuthenticationType() EnuminterServerAuthenticationInfoAuthenticationTypeProp`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *GetInterServerAuthenticationInfo200Response) GetAuthenticationTypeOk() (*EnuminterServerAuthenticationInfoAuthenticationTypeProp, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *GetInterServerAuthenticationInfo200Response) SetAuthenticationType(v EnuminterServerAuthenticationInfoAuthenticationTypeProp)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *GetInterServerAuthenticationInfo200Response) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetBindDN

`func (o *GetInterServerAuthenticationInfo200Response) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *GetInterServerAuthenticationInfo200Response) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *GetInterServerAuthenticationInfo200Response) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.

### HasBindDN

`func (o *GetInterServerAuthenticationInfo200Response) HasBindDN() bool`

HasBindDN returns a boolean if a field has been set.

### GetUsername

`func (o *GetInterServerAuthenticationInfo200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetInterServerAuthenticationInfo200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetInterServerAuthenticationInfo200Response) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetInterServerAuthenticationInfo200Response) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *GetInterServerAuthenticationInfo200Response) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetInterServerAuthenticationInfo200Response) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetInterServerAuthenticationInfo200Response) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPurpose

`func (o *GetInterServerAuthenticationInfo200Response) GetPurpose() []EnuminterServerAuthenticationInfoPurposeProp`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *GetInterServerAuthenticationInfo200Response) GetPurposeOk() (*[]EnuminterServerAuthenticationInfoPurposeProp, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *GetInterServerAuthenticationInfo200Response) SetPurpose(v []EnuminterServerAuthenticationInfoPurposeProp)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *GetInterServerAuthenticationInfo200Response) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetMeta

`func (o *GetInterServerAuthenticationInfo200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetInterServerAuthenticationInfo200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetInterServerAuthenticationInfo200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetInterServerAuthenticationInfo200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GetInterServerAuthenticationInfo200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GetInterServerAuthenticationInfo200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GetInterServerAuthenticationInfo200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GetInterServerAuthenticationInfo200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


