# AmazonAwsExternalServerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumamazonAwsExternalServerSchemaUrn**](EnumamazonAwsExternalServerSchemaUrn.md) |  | 
**AwsAccessKeyID** | Pointer to **string** | The access key ID that will be used if authentication should use an access key. If this is provided, then an aws-secret-access-key must also be provided. If this is not provided, then no aws-secret-access-key may be configured, and the server must be running in an EC2 instance that is configured with an IAM role with permission to perform the necessary operations. | [optional] 
**AwsSecretAccessKey** | Pointer to **string** | The secret access key that will be used if authentication should use an access key. If this is provided, then an aws-access-key-id must also be provided. If this is not provided, then no aws-access-key-id may be configured, and the server must be running in an EC2 instance that is configured with an IAM role with permission to perform the necessary operations. | [optional] 
**AwsRegionName** | **string** | The name of the AWS region containing the resources that will be accessed. | 
**Description** | Pointer to **string** | A description for this External Server | [optional] 

## Methods

### NewAmazonAwsExternalServerShared

`func NewAmazonAwsExternalServerShared(schemas []EnumamazonAwsExternalServerSchemaUrn, awsRegionName string, ) *AmazonAwsExternalServerShared`

NewAmazonAwsExternalServerShared instantiates a new AmazonAwsExternalServerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonAwsExternalServerSharedWithDefaults

`func NewAmazonAwsExternalServerSharedWithDefaults() *AmazonAwsExternalServerShared`

NewAmazonAwsExternalServerSharedWithDefaults instantiates a new AmazonAwsExternalServerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AmazonAwsExternalServerShared) GetSchemas() []EnumamazonAwsExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AmazonAwsExternalServerShared) GetSchemasOk() (*[]EnumamazonAwsExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AmazonAwsExternalServerShared) SetSchemas(v []EnumamazonAwsExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAwsAccessKeyID

`func (o *AmazonAwsExternalServerShared) GetAwsAccessKeyID() string`

GetAwsAccessKeyID returns the AwsAccessKeyID field if non-nil, zero value otherwise.

### GetAwsAccessKeyIDOk

`func (o *AmazonAwsExternalServerShared) GetAwsAccessKeyIDOk() (*string, bool)`

GetAwsAccessKeyIDOk returns a tuple with the AwsAccessKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyID

`func (o *AmazonAwsExternalServerShared) SetAwsAccessKeyID(v string)`

SetAwsAccessKeyID sets AwsAccessKeyID field to given value.

### HasAwsAccessKeyID

`func (o *AmazonAwsExternalServerShared) HasAwsAccessKeyID() bool`

HasAwsAccessKeyID returns a boolean if a field has been set.

### GetAwsSecretAccessKey

`func (o *AmazonAwsExternalServerShared) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *AmazonAwsExternalServerShared) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *AmazonAwsExternalServerShared) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *AmazonAwsExternalServerShared) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### GetAwsRegionName

`func (o *AmazonAwsExternalServerShared) GetAwsRegionName() string`

GetAwsRegionName returns the AwsRegionName field if non-nil, zero value otherwise.

### GetAwsRegionNameOk

`func (o *AmazonAwsExternalServerShared) GetAwsRegionNameOk() (*string, bool)`

GetAwsRegionNameOk returns a tuple with the AwsRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegionName

`func (o *AmazonAwsExternalServerShared) SetAwsRegionName(v string)`

SetAwsRegionName sets AwsRegionName field to given value.


### GetDescription

`func (o *AmazonAwsExternalServerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AmazonAwsExternalServerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AmazonAwsExternalServerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AmazonAwsExternalServerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


