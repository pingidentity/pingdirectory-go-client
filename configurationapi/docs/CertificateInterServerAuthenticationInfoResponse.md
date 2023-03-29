# CertificateInterServerAuthenticationInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumcertificateInterServerAuthenticationInfoSchemaUrn**](EnumcertificateInterServerAuthenticationInfoSchemaUrn.md) |  | 
**Id** | **string** | Name of the Server Instance | 
**Purpose** | Pointer to [**[]EnuminterServerAuthenticationInfoPurposeProp**](EnuminterServerAuthenticationInfoPurposeProp.md) | Identifies the purpose of this Inter Server Authentication Info. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewCertificateInterServerAuthenticationInfoResponse

`func NewCertificateInterServerAuthenticationInfoResponse(schemas []EnumcertificateInterServerAuthenticationInfoSchemaUrn, id string, ) *CertificateInterServerAuthenticationInfoResponse`

NewCertificateInterServerAuthenticationInfoResponse instantiates a new CertificateInterServerAuthenticationInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateInterServerAuthenticationInfoResponseWithDefaults

`func NewCertificateInterServerAuthenticationInfoResponseWithDefaults() *CertificateInterServerAuthenticationInfoResponse`

NewCertificateInterServerAuthenticationInfoResponseWithDefaults instantiates a new CertificateInterServerAuthenticationInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *CertificateInterServerAuthenticationInfoResponse) GetSchemas() []EnumcertificateInterServerAuthenticationInfoSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CertificateInterServerAuthenticationInfoResponse) GetSchemasOk() (*[]EnumcertificateInterServerAuthenticationInfoSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CertificateInterServerAuthenticationInfoResponse) SetSchemas(v []EnumcertificateInterServerAuthenticationInfoSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *CertificateInterServerAuthenticationInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateInterServerAuthenticationInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateInterServerAuthenticationInfoResponse) SetId(v string)`

SetId sets Id field to given value.


### GetPurpose

`func (o *CertificateInterServerAuthenticationInfoResponse) GetPurpose() []EnuminterServerAuthenticationInfoPurposeProp`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *CertificateInterServerAuthenticationInfoResponse) GetPurposeOk() (*[]EnuminterServerAuthenticationInfoPurposeProp, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *CertificateInterServerAuthenticationInfoResponse) SetPurpose(v []EnuminterServerAuthenticationInfoPurposeProp)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *CertificateInterServerAuthenticationInfoResponse) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetMeta

`func (o *CertificateInterServerAuthenticationInfoResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CertificateInterServerAuthenticationInfoResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CertificateInterServerAuthenticationInfoResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CertificateInterServerAuthenticationInfoResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *CertificateInterServerAuthenticationInfoResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *CertificateInterServerAuthenticationInfoResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *CertificateInterServerAuthenticationInfoResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *CertificateInterServerAuthenticationInfoResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


