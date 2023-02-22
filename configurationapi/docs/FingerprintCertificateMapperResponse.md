# FingerprintCertificateMapperResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Certificate Mapper | 
**Schemas** | [**[]EnumfingerprintCertificateMapperSchemaUrn**](EnumfingerprintCertificateMapperSchemaUrn.md) |  | 
**FingerprintAttribute** | **string** | Specifies the attribute in which to look for the fingerprint. | 
**FingerprintAlgorithm** | [**EnumcertificateMapperFingerprintAlgorithmProp**](EnumcertificateMapperFingerprintAlgorithmProp.md) |  | 
**UserBaseDN** | Pointer to **[]string** | Specifies the set of base DNs below which to search for users. | [optional] 
**Description** | Pointer to **string** | A description for this Certificate Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Certificate Mapper is enabled. | 

## Methods

### NewFingerprintCertificateMapperResponse

`func NewFingerprintCertificateMapperResponse(id string, schemas []EnumfingerprintCertificateMapperSchemaUrn, fingerprintAttribute string, fingerprintAlgorithm EnumcertificateMapperFingerprintAlgorithmProp, enabled bool, ) *FingerprintCertificateMapperResponse`

NewFingerprintCertificateMapperResponse instantiates a new FingerprintCertificateMapperResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFingerprintCertificateMapperResponseWithDefaults

`func NewFingerprintCertificateMapperResponseWithDefaults() *FingerprintCertificateMapperResponse`

NewFingerprintCertificateMapperResponseWithDefaults instantiates a new FingerprintCertificateMapperResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *FingerprintCertificateMapperResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FingerprintCertificateMapperResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FingerprintCertificateMapperResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FingerprintCertificateMapperResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *FingerprintCertificateMapperResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *FingerprintCertificateMapperResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *FingerprintCertificateMapperResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *FingerprintCertificateMapperResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *FingerprintCertificateMapperResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FingerprintCertificateMapperResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FingerprintCertificateMapperResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *FingerprintCertificateMapperResponse) GetSchemas() []EnumfingerprintCertificateMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FingerprintCertificateMapperResponse) GetSchemasOk() (*[]EnumfingerprintCertificateMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FingerprintCertificateMapperResponse) SetSchemas(v []EnumfingerprintCertificateMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetFingerprintAttribute

`func (o *FingerprintCertificateMapperResponse) GetFingerprintAttribute() string`

GetFingerprintAttribute returns the FingerprintAttribute field if non-nil, zero value otherwise.

### GetFingerprintAttributeOk

`func (o *FingerprintCertificateMapperResponse) GetFingerprintAttributeOk() (*string, bool)`

GetFingerprintAttributeOk returns a tuple with the FingerprintAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintAttribute

`func (o *FingerprintCertificateMapperResponse) SetFingerprintAttribute(v string)`

SetFingerprintAttribute sets FingerprintAttribute field to given value.


### GetFingerprintAlgorithm

`func (o *FingerprintCertificateMapperResponse) GetFingerprintAlgorithm() EnumcertificateMapperFingerprintAlgorithmProp`

GetFingerprintAlgorithm returns the FingerprintAlgorithm field if non-nil, zero value otherwise.

### GetFingerprintAlgorithmOk

`func (o *FingerprintCertificateMapperResponse) GetFingerprintAlgorithmOk() (*EnumcertificateMapperFingerprintAlgorithmProp, bool)`

GetFingerprintAlgorithmOk returns a tuple with the FingerprintAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintAlgorithm

`func (o *FingerprintCertificateMapperResponse) SetFingerprintAlgorithm(v EnumcertificateMapperFingerprintAlgorithmProp)`

SetFingerprintAlgorithm sets FingerprintAlgorithm field to given value.


### GetUserBaseDN

`func (o *FingerprintCertificateMapperResponse) GetUserBaseDN() []string`

GetUserBaseDN returns the UserBaseDN field if non-nil, zero value otherwise.

### GetUserBaseDNOk

`func (o *FingerprintCertificateMapperResponse) GetUserBaseDNOk() (*[]string, bool)`

GetUserBaseDNOk returns a tuple with the UserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBaseDN

`func (o *FingerprintCertificateMapperResponse) SetUserBaseDN(v []string)`

SetUserBaseDN sets UserBaseDN field to given value.

### HasUserBaseDN

`func (o *FingerprintCertificateMapperResponse) HasUserBaseDN() bool`

HasUserBaseDN returns a boolean if a field has been set.

### GetDescription

`func (o *FingerprintCertificateMapperResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FingerprintCertificateMapperResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FingerprintCertificateMapperResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FingerprintCertificateMapperResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *FingerprintCertificateMapperResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FingerprintCertificateMapperResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FingerprintCertificateMapperResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


