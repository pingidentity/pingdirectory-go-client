# AddCertificateMapper200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Certificate Mapper | 
**Schemas** | [**[]EnumthirdPartyCertificateMapperSchemaUrn**](EnumthirdPartyCertificateMapperSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Certificate Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Certificate Mapper is enabled. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**SubjectAttribute** | **string** | Specifies the name or OID of the attribute whose value should exactly match the certificate subject DN. | 
**UserBaseDN** | Pointer to **[]string** | Specifies the set of base DNs below which to search for users. | [optional] 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Certificate Mapper. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted Certificate Mapper. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**SubjectAttributeMapping** | **[]string** | Specifies a mapping between certificate attributes and user attributes. | 
**FingerprintAttribute** | **string** | Specifies the attribute in which to look for the fingerprint. | 
**FingerprintAlgorithm** | [**EnumcertificateMapperFingerprintAlgorithmProp**](EnumcertificateMapperFingerprintAlgorithmProp.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Certificate Mapper. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Certificate Mapper. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddCertificateMapper200Response

`func NewAddCertificateMapper200Response(id string, schemas []EnumthirdPartyCertificateMapperSchemaUrn, enabled bool, subjectAttribute string, scriptClass string, subjectAttributeMapping []string, fingerprintAttribute string, fingerprintAlgorithm EnumcertificateMapperFingerprintAlgorithmProp, extensionClass string, ) *AddCertificateMapper200Response`

NewAddCertificateMapper200Response instantiates a new AddCertificateMapper200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCertificateMapper200ResponseWithDefaults

`func NewAddCertificateMapper200ResponseWithDefaults() *AddCertificateMapper200Response`

NewAddCertificateMapper200ResponseWithDefaults instantiates a new AddCertificateMapper200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddCertificateMapper200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddCertificateMapper200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddCertificateMapper200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddCertificateMapper200Response) GetSchemas() []EnumthirdPartyCertificateMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddCertificateMapper200Response) GetSchemasOk() (*[]EnumthirdPartyCertificateMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddCertificateMapper200Response) SetSchemas(v []EnumthirdPartyCertificateMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddCertificateMapper200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCertificateMapper200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCertificateMapper200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCertificateMapper200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddCertificateMapper200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCertificateMapper200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCertificateMapper200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *AddCertificateMapper200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddCertificateMapper200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddCertificateMapper200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddCertificateMapper200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddCertificateMapper200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddCertificateMapper200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddCertificateMapper200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddCertificateMapper200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSubjectAttribute

`func (o *AddCertificateMapper200Response) GetSubjectAttribute() string`

GetSubjectAttribute returns the SubjectAttribute field if non-nil, zero value otherwise.

### GetSubjectAttributeOk

`func (o *AddCertificateMapper200Response) GetSubjectAttributeOk() (*string, bool)`

GetSubjectAttributeOk returns a tuple with the SubjectAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAttribute

`func (o *AddCertificateMapper200Response) SetSubjectAttribute(v string)`

SetSubjectAttribute sets SubjectAttribute field to given value.


### GetUserBaseDN

`func (o *AddCertificateMapper200Response) GetUserBaseDN() []string`

GetUserBaseDN returns the UserBaseDN field if non-nil, zero value otherwise.

### GetUserBaseDNOk

`func (o *AddCertificateMapper200Response) GetUserBaseDNOk() (*[]string, bool)`

GetUserBaseDNOk returns a tuple with the UserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBaseDN

`func (o *AddCertificateMapper200Response) SetUserBaseDN(v []string)`

SetUserBaseDN sets UserBaseDN field to given value.

### HasUserBaseDN

`func (o *AddCertificateMapper200Response) HasUserBaseDN() bool`

HasUserBaseDN returns a boolean if a field has been set.

### GetScriptClass

`func (o *AddCertificateMapper200Response) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddCertificateMapper200Response) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddCertificateMapper200Response) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddCertificateMapper200Response) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddCertificateMapper200Response) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddCertificateMapper200Response) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddCertificateMapper200Response) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetSubjectAttributeMapping

`func (o *AddCertificateMapper200Response) GetSubjectAttributeMapping() []string`

GetSubjectAttributeMapping returns the SubjectAttributeMapping field if non-nil, zero value otherwise.

### GetSubjectAttributeMappingOk

`func (o *AddCertificateMapper200Response) GetSubjectAttributeMappingOk() (*[]string, bool)`

GetSubjectAttributeMappingOk returns a tuple with the SubjectAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAttributeMapping

`func (o *AddCertificateMapper200Response) SetSubjectAttributeMapping(v []string)`

SetSubjectAttributeMapping sets SubjectAttributeMapping field to given value.


### GetFingerprintAttribute

`func (o *AddCertificateMapper200Response) GetFingerprintAttribute() string`

GetFingerprintAttribute returns the FingerprintAttribute field if non-nil, zero value otherwise.

### GetFingerprintAttributeOk

`func (o *AddCertificateMapper200Response) GetFingerprintAttributeOk() (*string, bool)`

GetFingerprintAttributeOk returns a tuple with the FingerprintAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintAttribute

`func (o *AddCertificateMapper200Response) SetFingerprintAttribute(v string)`

SetFingerprintAttribute sets FingerprintAttribute field to given value.


### GetFingerprintAlgorithm

`func (o *AddCertificateMapper200Response) GetFingerprintAlgorithm() EnumcertificateMapperFingerprintAlgorithmProp`

GetFingerprintAlgorithm returns the FingerprintAlgorithm field if non-nil, zero value otherwise.

### GetFingerprintAlgorithmOk

`func (o *AddCertificateMapper200Response) GetFingerprintAlgorithmOk() (*EnumcertificateMapperFingerprintAlgorithmProp, bool)`

GetFingerprintAlgorithmOk returns a tuple with the FingerprintAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintAlgorithm

`func (o *AddCertificateMapper200Response) SetFingerprintAlgorithm(v EnumcertificateMapperFingerprintAlgorithmProp)`

SetFingerprintAlgorithm sets FingerprintAlgorithm field to given value.


### GetExtensionClass

`func (o *AddCertificateMapper200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddCertificateMapper200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddCertificateMapper200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddCertificateMapper200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddCertificateMapper200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddCertificateMapper200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddCertificateMapper200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


