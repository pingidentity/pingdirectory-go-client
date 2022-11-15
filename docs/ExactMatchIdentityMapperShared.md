# ExactMatchIdentityMapperShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumexactMatchIdentityMapperSchemaUrn**](EnumexactMatchIdentityMapperSchemaUrn.md) |  | 
**MatchAttribute** | **[]string** |  | 
**MatchBaseDN** | Pointer to **[]string** |  | [optional] 
**MatchFilter** | Pointer to **string** | An optional filter that mapped users must match. | [optional] 
**Description** | Pointer to **string** | A description for this Identity Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Identity Mapper is enabled for use. | 

## Methods

### NewExactMatchIdentityMapperShared

`func NewExactMatchIdentityMapperShared(schemas []EnumexactMatchIdentityMapperSchemaUrn, matchAttribute []string, enabled bool, ) *ExactMatchIdentityMapperShared`

NewExactMatchIdentityMapperShared instantiates a new ExactMatchIdentityMapperShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExactMatchIdentityMapperSharedWithDefaults

`func NewExactMatchIdentityMapperSharedWithDefaults() *ExactMatchIdentityMapperShared`

NewExactMatchIdentityMapperSharedWithDefaults instantiates a new ExactMatchIdentityMapperShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ExactMatchIdentityMapperShared) GetSchemas() []EnumexactMatchIdentityMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ExactMatchIdentityMapperShared) GetSchemasOk() (*[]EnumexactMatchIdentityMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ExactMatchIdentityMapperShared) SetSchemas(v []EnumexactMatchIdentityMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetMatchAttribute

`func (o *ExactMatchIdentityMapperShared) GetMatchAttribute() []string`

GetMatchAttribute returns the MatchAttribute field if non-nil, zero value otherwise.

### GetMatchAttributeOk

`func (o *ExactMatchIdentityMapperShared) GetMatchAttributeOk() (*[]string, bool)`

GetMatchAttributeOk returns a tuple with the MatchAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAttribute

`func (o *ExactMatchIdentityMapperShared) SetMatchAttribute(v []string)`

SetMatchAttribute sets MatchAttribute field to given value.


### GetMatchBaseDN

`func (o *ExactMatchIdentityMapperShared) GetMatchBaseDN() []string`

GetMatchBaseDN returns the MatchBaseDN field if non-nil, zero value otherwise.

### GetMatchBaseDNOk

`func (o *ExactMatchIdentityMapperShared) GetMatchBaseDNOk() (*[]string, bool)`

GetMatchBaseDNOk returns a tuple with the MatchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchBaseDN

`func (o *ExactMatchIdentityMapperShared) SetMatchBaseDN(v []string)`

SetMatchBaseDN sets MatchBaseDN field to given value.

### HasMatchBaseDN

`func (o *ExactMatchIdentityMapperShared) HasMatchBaseDN() bool`

HasMatchBaseDN returns a boolean if a field has been set.

### GetMatchFilter

`func (o *ExactMatchIdentityMapperShared) GetMatchFilter() string`

GetMatchFilter returns the MatchFilter field if non-nil, zero value otherwise.

### GetMatchFilterOk

`func (o *ExactMatchIdentityMapperShared) GetMatchFilterOk() (*string, bool)`

GetMatchFilterOk returns a tuple with the MatchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchFilter

`func (o *ExactMatchIdentityMapperShared) SetMatchFilter(v string)`

SetMatchFilter sets MatchFilter field to given value.

### HasMatchFilter

`func (o *ExactMatchIdentityMapperShared) HasMatchFilter() bool`

HasMatchFilter returns a boolean if a field has been set.

### GetDescription

`func (o *ExactMatchIdentityMapperShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExactMatchIdentityMapperShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExactMatchIdentityMapperShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExactMatchIdentityMapperShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ExactMatchIdentityMapperShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExactMatchIdentityMapperShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExactMatchIdentityMapperShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


