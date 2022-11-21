# RegularExpressionIdentityMapperResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Identity Mapper | 
**Schemas** | [**[]EnumregularExpressionIdentityMapperSchemaUrn**](EnumregularExpressionIdentityMapperSchemaUrn.md) |  | 
**MatchAttribute** | **[]string** | Specifies the name or OID of the attribute whose value should match the provided identifier string after it has been processed by the associated regular expression. | 
**MatchBaseDN** | Pointer to **[]string** | Specifies the base DN(s) that should be used when performing searches to map the provided ID string to a user entry. If multiple values are given, searches are performed below all the specified base DNs. | [optional] 
**MatchFilter** | Pointer to **string** | An optional filter that mapped users must match. | [optional] 
**MatchPattern** | **string** | Specifies the regular expression pattern that is used to identify portions of the ID string that will be replaced. | 
**ReplacePattern** | Pointer to **string** | Specifies the replacement pattern that should be used for substrings in the ID string that match the provided regular expression pattern. | [optional] 
**Description** | Pointer to **string** | A description for this Identity Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Identity Mapper is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewRegularExpressionIdentityMapperResponse

`func NewRegularExpressionIdentityMapperResponse(id string, schemas []EnumregularExpressionIdentityMapperSchemaUrn, matchAttribute []string, matchPattern string, enabled bool, ) *RegularExpressionIdentityMapperResponse`

NewRegularExpressionIdentityMapperResponse instantiates a new RegularExpressionIdentityMapperResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegularExpressionIdentityMapperResponseWithDefaults

`func NewRegularExpressionIdentityMapperResponseWithDefaults() *RegularExpressionIdentityMapperResponse`

NewRegularExpressionIdentityMapperResponseWithDefaults instantiates a new RegularExpressionIdentityMapperResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegularExpressionIdentityMapperResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegularExpressionIdentityMapperResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegularExpressionIdentityMapperResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *RegularExpressionIdentityMapperResponse) GetSchemas() []EnumregularExpressionIdentityMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *RegularExpressionIdentityMapperResponse) GetSchemasOk() (*[]EnumregularExpressionIdentityMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *RegularExpressionIdentityMapperResponse) SetSchemas(v []EnumregularExpressionIdentityMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetMatchAttribute

`func (o *RegularExpressionIdentityMapperResponse) GetMatchAttribute() []string`

GetMatchAttribute returns the MatchAttribute field if non-nil, zero value otherwise.

### GetMatchAttributeOk

`func (o *RegularExpressionIdentityMapperResponse) GetMatchAttributeOk() (*[]string, bool)`

GetMatchAttributeOk returns a tuple with the MatchAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAttribute

`func (o *RegularExpressionIdentityMapperResponse) SetMatchAttribute(v []string)`

SetMatchAttribute sets MatchAttribute field to given value.


### GetMatchBaseDN

`func (o *RegularExpressionIdentityMapperResponse) GetMatchBaseDN() []string`

GetMatchBaseDN returns the MatchBaseDN field if non-nil, zero value otherwise.

### GetMatchBaseDNOk

`func (o *RegularExpressionIdentityMapperResponse) GetMatchBaseDNOk() (*[]string, bool)`

GetMatchBaseDNOk returns a tuple with the MatchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchBaseDN

`func (o *RegularExpressionIdentityMapperResponse) SetMatchBaseDN(v []string)`

SetMatchBaseDN sets MatchBaseDN field to given value.

### HasMatchBaseDN

`func (o *RegularExpressionIdentityMapperResponse) HasMatchBaseDN() bool`

HasMatchBaseDN returns a boolean if a field has been set.

### GetMatchFilter

`func (o *RegularExpressionIdentityMapperResponse) GetMatchFilter() string`

GetMatchFilter returns the MatchFilter field if non-nil, zero value otherwise.

### GetMatchFilterOk

`func (o *RegularExpressionIdentityMapperResponse) GetMatchFilterOk() (*string, bool)`

GetMatchFilterOk returns a tuple with the MatchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchFilter

`func (o *RegularExpressionIdentityMapperResponse) SetMatchFilter(v string)`

SetMatchFilter sets MatchFilter field to given value.

### HasMatchFilter

`func (o *RegularExpressionIdentityMapperResponse) HasMatchFilter() bool`

HasMatchFilter returns a boolean if a field has been set.

### GetMatchPattern

`func (o *RegularExpressionIdentityMapperResponse) GetMatchPattern() string`

GetMatchPattern returns the MatchPattern field if non-nil, zero value otherwise.

### GetMatchPatternOk

`func (o *RegularExpressionIdentityMapperResponse) GetMatchPatternOk() (*string, bool)`

GetMatchPatternOk returns a tuple with the MatchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPattern

`func (o *RegularExpressionIdentityMapperResponse) SetMatchPattern(v string)`

SetMatchPattern sets MatchPattern field to given value.


### GetReplacePattern

`func (o *RegularExpressionIdentityMapperResponse) GetReplacePattern() string`

GetReplacePattern returns the ReplacePattern field if non-nil, zero value otherwise.

### GetReplacePatternOk

`func (o *RegularExpressionIdentityMapperResponse) GetReplacePatternOk() (*string, bool)`

GetReplacePatternOk returns a tuple with the ReplacePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacePattern

`func (o *RegularExpressionIdentityMapperResponse) SetReplacePattern(v string)`

SetReplacePattern sets ReplacePattern field to given value.

### HasReplacePattern

`func (o *RegularExpressionIdentityMapperResponse) HasReplacePattern() bool`

HasReplacePattern returns a boolean if a field has been set.

### GetDescription

`func (o *RegularExpressionIdentityMapperResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RegularExpressionIdentityMapperResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RegularExpressionIdentityMapperResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RegularExpressionIdentityMapperResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *RegularExpressionIdentityMapperResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RegularExpressionIdentityMapperResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RegularExpressionIdentityMapperResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *RegularExpressionIdentityMapperResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RegularExpressionIdentityMapperResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RegularExpressionIdentityMapperResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RegularExpressionIdentityMapperResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


