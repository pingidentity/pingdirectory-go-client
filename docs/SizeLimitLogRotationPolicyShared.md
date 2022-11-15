# SizeLimitLogRotationPolicyShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsizeLimitLogRotationPolicySchemaUrn**](EnumsizeLimitLogRotationPolicySchemaUrn.md) |  | 
**FileSizeLimit** | **string** | Specifies the maximum size that a log file can reach before it is rotated. | 
**Description** | Pointer to **string** | A description for this Log Rotation Policy | [optional] 

## Methods

### NewSizeLimitLogRotationPolicyShared

`func NewSizeLimitLogRotationPolicyShared(schemas []EnumsizeLimitLogRotationPolicySchemaUrn, fileSizeLimit string, ) *SizeLimitLogRotationPolicyShared`

NewSizeLimitLogRotationPolicyShared instantiates a new SizeLimitLogRotationPolicyShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeLimitLogRotationPolicySharedWithDefaults

`func NewSizeLimitLogRotationPolicySharedWithDefaults() *SizeLimitLogRotationPolicyShared`

NewSizeLimitLogRotationPolicySharedWithDefaults instantiates a new SizeLimitLogRotationPolicyShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SizeLimitLogRotationPolicyShared) GetSchemas() []EnumsizeLimitLogRotationPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SizeLimitLogRotationPolicyShared) GetSchemasOk() (*[]EnumsizeLimitLogRotationPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SizeLimitLogRotationPolicyShared) SetSchemas(v []EnumsizeLimitLogRotationPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetFileSizeLimit

`func (o *SizeLimitLogRotationPolicyShared) GetFileSizeLimit() string`

GetFileSizeLimit returns the FileSizeLimit field if non-nil, zero value otherwise.

### GetFileSizeLimitOk

`func (o *SizeLimitLogRotationPolicyShared) GetFileSizeLimitOk() (*string, bool)`

GetFileSizeLimitOk returns a tuple with the FileSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeLimit

`func (o *SizeLimitLogRotationPolicyShared) SetFileSizeLimit(v string)`

SetFileSizeLimit sets FileSizeLimit field to given value.


### GetDescription

`func (o *SizeLimitLogRotationPolicyShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SizeLimitLogRotationPolicyShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SizeLimitLogRotationPolicyShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SizeLimitLogRotationPolicyShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


