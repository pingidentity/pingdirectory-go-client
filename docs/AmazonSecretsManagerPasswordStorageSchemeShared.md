# AmazonSecretsManagerPasswordStorageSchemeShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn**](EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn.md) |  | 
**AwsExternalServer** | **string** | The external server with information to use when interacting with the AWS Secrets Manager service. | 
**DefaultField** | Pointer to **string** | The default name of the field in JSON objects contained in the AWS Secrets Manager service that contains the password for the target user. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 

## Methods

### NewAmazonSecretsManagerPasswordStorageSchemeShared

`func NewAmazonSecretsManagerPasswordStorageSchemeShared(schemas []EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn, awsExternalServer string, enabled bool, ) *AmazonSecretsManagerPasswordStorageSchemeShared`

NewAmazonSecretsManagerPasswordStorageSchemeShared instantiates a new AmazonSecretsManagerPasswordStorageSchemeShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonSecretsManagerPasswordStorageSchemeSharedWithDefaults

`func NewAmazonSecretsManagerPasswordStorageSchemeSharedWithDefaults() *AmazonSecretsManagerPasswordStorageSchemeShared`

NewAmazonSecretsManagerPasswordStorageSchemeSharedWithDefaults instantiates a new AmazonSecretsManagerPasswordStorageSchemeShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) GetSchemas() []EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) GetSchemasOk() (*[]EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) SetSchemas(v []EnumamazonSecretsManagerPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAwsExternalServer

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.


### GetDefaultField

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) GetDefaultField() string`

GetDefaultField returns the DefaultField field if non-nil, zero value otherwise.

### GetDefaultFieldOk

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) GetDefaultFieldOk() (*string, bool)`

GetDefaultFieldOk returns a tuple with the DefaultField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultField

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) SetDefaultField(v string)`

SetDefaultField sets DefaultField field to given value.

### HasDefaultField

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) HasDefaultField() bool`

HasDefaultField returns a boolean if a field has been set.

### GetDescription

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AmazonSecretsManagerPasswordStorageSchemeShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


