# RegularExpressionIdentityMapperShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumregularExpressionIdentityMapperSchemaUrn**](EnumregularExpressionIdentityMapperSchemaUrn.md) |  | 
**MatchAttribute** | **[]string** |  | 
**MatchBaseDN** | Pointer to **[]string** |  | [optional] 
**MatchFilter** | Pointer to **string** | An optional filter that mapped users must match. | [optional] 
**MatchPattern** | **string** | Specifies the regular expression pattern that is used to identify portions of the ID string that will be replaced. | 
**ReplacePattern** | Pointer to **string** | Specifies the replacement pattern that should be used for substrings in the ID string that match the provided regular expression pattern. | [optional] 
**Description** | Pointer to **string** | A description for this Identity Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Identity Mapper is enabled for use. | 

## Methods

### NewRegularExpressionIdentityMapperShared

`func NewRegularExpressionIdentityMapperShared(schemas []EnumregularExpressionIdentityMapperSchemaUrn, matchAttribute []string, matchPattern string, enabled bool, ) *RegularExpressionIdentityMapperShared`

NewRegularExpressionIdentityMapperShared instantiates a new RegularExpressionIdentityMapperShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegularExpressionIdentityMapperSharedWithDefaults

`func NewRegularExpressionIdentityMapperSharedWithDefaults() *RegularExpressionIdentityMapperShared`

NewRegularExpressionIdentityMapperSharedWithDefaults instantiates a new RegularExpressionIdentityMapperShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *RegularExpressionIdentityMapperShared) GetSchemas() []EnumregularExpressionIdentityMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *RegularExpressionIdentityMapperShared) GetSchemasOk() (*[]EnumregularExpressionIdentityMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *RegularExpressionIdentityMapperShared) SetSchemas(v []EnumregularExpressionIdentityMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetMatchAttribute

`func (o *RegularExpressionIdentityMapperShared) GetMatchAttribute() []string`

GetMatchAttribute returns the MatchAttribute field if non-nil, zero value otherwise.

### GetMatchAttributeOk

`func (o *RegularExpressionIdentityMapperShared) GetMatchAttributeOk() (*[]string, bool)`

GetMatchAttributeOk returns a tuple with the MatchAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAttribute

`func (o *RegularExpressionIdentityMapperShared) SetMatchAttribute(v []string)`

SetMatchAttribute sets MatchAttribute field to given value.


### GetMatchBaseDN

`func (o *RegularExpressionIdentityMapperShared) GetMatchBaseDN() []string`

GetMatchBaseDN returns the MatchBaseDN field if non-nil, zero value otherwise.

### GetMatchBaseDNOk

`func (o *RegularExpressionIdentityMapperShared) GetMatchBaseDNOk() (*[]string, bool)`

GetMatchBaseDNOk returns a tuple with the MatchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchBaseDN

`func (o *RegularExpressionIdentityMapperShared) SetMatchBaseDN(v []string)`

SetMatchBaseDN sets MatchBaseDN field to given value.

### HasMatchBaseDN

`func (o *RegularExpressionIdentityMapperShared) HasMatchBaseDN() bool`

HasMatchBaseDN returns a boolean if a field has been set.

### GetMatchFilter

`func (o *RegularExpressionIdentityMapperShared) GetMatchFilter() string`

GetMatchFilter returns the MatchFilter field if non-nil, zero value otherwise.

### GetMatchFilterOk

`func (o *RegularExpressionIdentityMapperShared) GetMatchFilterOk() (*string, bool)`

GetMatchFilterOk returns a tuple with the MatchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchFilter

`func (o *RegularExpressionIdentityMapperShared) SetMatchFilter(v string)`

SetMatchFilter sets MatchFilter field to given value.

### HasMatchFilter

`func (o *RegularExpressionIdentityMapperShared) HasMatchFilter() bool`

HasMatchFilter returns a boolean if a field has been set.

### GetMatchPattern

`func (o *RegularExpressionIdentityMapperShared) GetMatchPattern() string`

GetMatchPattern returns the MatchPattern field if non-nil, zero value otherwise.

### GetMatchPatternOk

`func (o *RegularExpressionIdentityMapperShared) GetMatchPatternOk() (*string, bool)`

GetMatchPatternOk returns a tuple with the MatchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPattern

`func (o *RegularExpressionIdentityMapperShared) SetMatchPattern(v string)`

SetMatchPattern sets MatchPattern field to given value.


### GetReplacePattern

`func (o *RegularExpressionIdentityMapperShared) GetReplacePattern() string`

GetReplacePattern returns the ReplacePattern field if non-nil, zero value otherwise.

### GetReplacePatternOk

`func (o *RegularExpressionIdentityMapperShared) GetReplacePatternOk() (*string, bool)`

GetReplacePatternOk returns a tuple with the ReplacePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacePattern

`func (o *RegularExpressionIdentityMapperShared) SetReplacePattern(v string)`

SetReplacePattern sets ReplacePattern field to given value.

### HasReplacePattern

`func (o *RegularExpressionIdentityMapperShared) HasReplacePattern() bool`

HasReplacePattern returns a boolean if a field has been set.

### GetDescription

`func (o *RegularExpressionIdentityMapperShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RegularExpressionIdentityMapperShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RegularExpressionIdentityMapperShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RegularExpressionIdentityMapperShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *RegularExpressionIdentityMapperShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RegularExpressionIdentityMapperShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RegularExpressionIdentityMapperShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


