# AddAmazonSecretsManagerPasswordStorageSchemeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn**](EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn.md) |  | 
**AwsExternalServer** | **string** | The external server with information to use when interacting with the AWS Secrets Manager service. | 
**DefaultField** | Pointer to **string** | The default name of the field in JSON objects contained in the AWS Secrets Manager service that contains the password for the target user. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 
**SchemeName** | **string** | Name of the new Password Storage Scheme | 

## Methods

### NewAddAmazonSecretsManagerPasswordStorageSchemeRequest

`func NewAddAmazonSecretsManagerPasswordStorageSchemeRequest(schemas []EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn, awsExternalServer string, enabled bool, schemeName string, ) *AddAmazonSecretsManagerPasswordStorageSchemeRequest`

NewAddAmazonSecretsManagerPasswordStorageSchemeRequest instantiates a new AddAmazonSecretsManagerPasswordStorageSchemeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAmazonSecretsManagerPasswordStorageSchemeRequestWithDefaults

`func NewAddAmazonSecretsManagerPasswordStorageSchemeRequestWithDefaults() *AddAmazonSecretsManagerPasswordStorageSchemeRequest`

NewAddAmazonSecretsManagerPasswordStorageSchemeRequestWithDefaults instantiates a new AddAmazonSecretsManagerPasswordStorageSchemeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetSchemas() []EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetSchemasOk() (*[]EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) SetSchemas(v []EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAwsExternalServer

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.


### GetDefaultField

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetDefaultField() string`

GetDefaultField returns the DefaultField field if non-nil, zero value otherwise.

### GetDefaultFieldOk

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetDefaultFieldOk() (*string, bool)`

GetDefaultFieldOk returns a tuple with the DefaultField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultField

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) SetDefaultField(v string)`

SetDefaultField sets DefaultField field to given value.

### HasDefaultField

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) HasDefaultField() bool`

HasDefaultField returns a boolean if a field has been set.

### GetDescription

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSchemeName

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetSchemeName() string`

GetSchemeName returns the SchemeName field if non-nil, zero value otherwise.

### GetSchemeNameOk

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) GetSchemeNameOk() (*string, bool)`

GetSchemeNameOk returns a tuple with the SchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeName

`func (o *AddAmazonSecretsManagerPasswordStorageSchemeRequest) SetSchemeName(v string)`

SetSchemeName sets SchemeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


