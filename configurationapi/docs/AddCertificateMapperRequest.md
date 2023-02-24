# AddCertificateMapperRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MapperName** | **string** | Name of the new Certificate Mapper | 
**Schemas** | [**[]EnumthirdPartyCertificateMapperSchemaUrn**](EnumthirdPartyCertificateMapperSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Certificate Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Certificate Mapper is enabled. | 
**SubjectAttribute** | Pointer to **string** | Specifies the name or OID of the attribute whose value should exactly match the certificate subject DN. | [optional] 
**UserBaseDN** | Pointer to **[]string** | Specifies the set of base DNs below which to search for users. | [optional] 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Certificate Mapper. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted Certificate Mapper. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**SubjectAttributeMapping** | **[]string** | Specifies a mapping between certificate attributes and user attributes. | 
**FingerprintAttribute** | Pointer to **string** | Specifies the attribute in which to look for the fingerprint. | [optional] 
**FingerprintAlgorithm** | [**EnumcertificateMapperFingerprintAlgorithmProp**](EnumcertificateMapperFingerprintAlgorithmProp.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Certificate Mapper. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Certificate Mapper. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddCertificateMapperRequest

`func NewAddCertificateMapperRequest(mapperName string, schemas []EnumthirdPartyCertificateMapperSchemaUrn, enabled bool, scriptClass string, subjectAttributeMapping []string, fingerprintAlgorithm EnumcertificateMapperFingerprintAlgorithmProp, extensionClass string, ) *AddCertificateMapperRequest`

NewAddCertificateMapperRequest instantiates a new AddCertificateMapperRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCertificateMapperRequestWithDefaults

`func NewAddCertificateMapperRequestWithDefaults() *AddCertificateMapperRequest`

NewAddCertificateMapperRequestWithDefaults instantiates a new AddCertificateMapperRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMapperName

`func (o *AddCertificateMapperRequest) GetMapperName() string`

GetMapperName returns the MapperName field if non-nil, zero value otherwise.

### GetMapperNameOk

`func (o *AddCertificateMapperRequest) GetMapperNameOk() (*string, bool)`

GetMapperNameOk returns a tuple with the MapperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperName

`func (o *AddCertificateMapperRequest) SetMapperName(v string)`

SetMapperName sets MapperName field to given value.


### GetSchemas

`func (o *AddCertificateMapperRequest) GetSchemas() []EnumthirdPartyCertificateMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddCertificateMapperRequest) GetSchemasOk() (*[]EnumthirdPartyCertificateMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddCertificateMapperRequest) SetSchemas(v []EnumthirdPartyCertificateMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddCertificateMapperRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCertificateMapperRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCertificateMapperRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCertificateMapperRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddCertificateMapperRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCertificateMapperRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCertificateMapperRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSubjectAttribute

`func (o *AddCertificateMapperRequest) GetSubjectAttribute() string`

GetSubjectAttribute returns the SubjectAttribute field if non-nil, zero value otherwise.

### GetSubjectAttributeOk

`func (o *AddCertificateMapperRequest) GetSubjectAttributeOk() (*string, bool)`

GetSubjectAttributeOk returns a tuple with the SubjectAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAttribute

`func (o *AddCertificateMapperRequest) SetSubjectAttribute(v string)`

SetSubjectAttribute sets SubjectAttribute field to given value.

### HasSubjectAttribute

`func (o *AddCertificateMapperRequest) HasSubjectAttribute() bool`

HasSubjectAttribute returns a boolean if a field has been set.

### GetUserBaseDN

`func (o *AddCertificateMapperRequest) GetUserBaseDN() []string`

GetUserBaseDN returns the UserBaseDN field if non-nil, zero value otherwise.

### GetUserBaseDNOk

`func (o *AddCertificateMapperRequest) GetUserBaseDNOk() (*[]string, bool)`

GetUserBaseDNOk returns a tuple with the UserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBaseDN

`func (o *AddCertificateMapperRequest) SetUserBaseDN(v []string)`

SetUserBaseDN sets UserBaseDN field to given value.

### HasUserBaseDN

`func (o *AddCertificateMapperRequest) HasUserBaseDN() bool`

HasUserBaseDN returns a boolean if a field has been set.

### GetScriptClass

`func (o *AddCertificateMapperRequest) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddCertificateMapperRequest) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddCertificateMapperRequest) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddCertificateMapperRequest) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddCertificateMapperRequest) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddCertificateMapperRequest) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddCertificateMapperRequest) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetSubjectAttributeMapping

`func (o *AddCertificateMapperRequest) GetSubjectAttributeMapping() []string`

GetSubjectAttributeMapping returns the SubjectAttributeMapping field if non-nil, zero value otherwise.

### GetSubjectAttributeMappingOk

`func (o *AddCertificateMapperRequest) GetSubjectAttributeMappingOk() (*[]string, bool)`

GetSubjectAttributeMappingOk returns a tuple with the SubjectAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAttributeMapping

`func (o *AddCertificateMapperRequest) SetSubjectAttributeMapping(v []string)`

SetSubjectAttributeMapping sets SubjectAttributeMapping field to given value.


### GetFingerprintAttribute

`func (o *AddCertificateMapperRequest) GetFingerprintAttribute() string`

GetFingerprintAttribute returns the FingerprintAttribute field if non-nil, zero value otherwise.

### GetFingerprintAttributeOk

`func (o *AddCertificateMapperRequest) GetFingerprintAttributeOk() (*string, bool)`

GetFingerprintAttributeOk returns a tuple with the FingerprintAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintAttribute

`func (o *AddCertificateMapperRequest) SetFingerprintAttribute(v string)`

SetFingerprintAttribute sets FingerprintAttribute field to given value.

### HasFingerprintAttribute

`func (o *AddCertificateMapperRequest) HasFingerprintAttribute() bool`

HasFingerprintAttribute returns a boolean if a field has been set.

### GetFingerprintAlgorithm

`func (o *AddCertificateMapperRequest) GetFingerprintAlgorithm() EnumcertificateMapperFingerprintAlgorithmProp`

GetFingerprintAlgorithm returns the FingerprintAlgorithm field if non-nil, zero value otherwise.

### GetFingerprintAlgorithmOk

`func (o *AddCertificateMapperRequest) GetFingerprintAlgorithmOk() (*EnumcertificateMapperFingerprintAlgorithmProp, bool)`

GetFingerprintAlgorithmOk returns a tuple with the FingerprintAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintAlgorithm

`func (o *AddCertificateMapperRequest) SetFingerprintAlgorithm(v EnumcertificateMapperFingerprintAlgorithmProp)`

SetFingerprintAlgorithm sets FingerprintAlgorithm field to given value.


### GetExtensionClass

`func (o *AddCertificateMapperRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddCertificateMapperRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddCertificateMapperRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddCertificateMapperRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddCertificateMapperRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddCertificateMapperRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddCertificateMapperRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


