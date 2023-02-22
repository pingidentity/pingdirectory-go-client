# AddIdentityMapperRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MapperName** | **string** | Name of the new Identity Mapper | 
**Schemas** | [**[]EnumthirdPartyIdentityMapperSchemaUrn**](EnumthirdPartyIdentityMapperSchemaUrn.md) |  | 
**MatchAttribute** | **[]string** | Specifies the name or OID of the attribute whose value should match the provided identifier string after it has been processed by the associated regular expression. | 
**MatchBaseDN** | Pointer to **[]string** | Specifies the base DN(s) that should be used when performing searches to map the provided ID string to a user entry. If multiple values are given, searches are performed below all the specified base DNs. | [optional] 
**MatchFilter** | Pointer to **string** | An optional filter that mapped users must match. | [optional] 
**Description** | Pointer to **string** | A description for this Identity Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Identity Mapper is enabled for use. | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Identity Mapper. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted Identity Mapper. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**MatchPattern** | **string** | Specifies the regular expression pattern that is used to identify portions of the ID string that will be replaced. | 
**ReplacePattern** | Pointer to **string** | Specifies the replacement pattern that should be used for substrings in the ID string that match the provided regular expression pattern. | [optional] 
**AllIncludedIdentityMapper** | Pointer to **[]string** | The set of identity mappers that must all match the target entry. Each identity mapper must uniquely match the same target entry. If any of the identity mappers match multiple entries, if any of them match zero entries, or if any of them match different entries, then the mapping will fail. | [optional] 
**AnyIncludedIdentityMapper** | Pointer to **[]string** | The set of identity mappers that will be used to identify the target entry. At least one identity mapper must uniquely match an entry. If multiple identity mappers match entries, then they must all uniquely match the same entry. If none of the identity mappers match any entries, if any of them match multiple entries, or if any of them match different entries, then the mapping will fail. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Identity Mapper. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Identity Mapper. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddIdentityMapperRequest

`func NewAddIdentityMapperRequest(mapperName string, schemas []EnumthirdPartyIdentityMapperSchemaUrn, matchAttribute []string, enabled bool, scriptClass string, matchPattern string, extensionClass string, ) *AddIdentityMapperRequest`

NewAddIdentityMapperRequest instantiates a new AddIdentityMapperRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIdentityMapperRequestWithDefaults

`func NewAddIdentityMapperRequestWithDefaults() *AddIdentityMapperRequest`

NewAddIdentityMapperRequestWithDefaults instantiates a new AddIdentityMapperRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMapperName

`func (o *AddIdentityMapperRequest) GetMapperName() string`

GetMapperName returns the MapperName field if non-nil, zero value otherwise.

### GetMapperNameOk

`func (o *AddIdentityMapperRequest) GetMapperNameOk() (*string, bool)`

GetMapperNameOk returns a tuple with the MapperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperName

`func (o *AddIdentityMapperRequest) SetMapperName(v string)`

SetMapperName sets MapperName field to given value.


### GetSchemas

`func (o *AddIdentityMapperRequest) GetSchemas() []EnumthirdPartyIdentityMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddIdentityMapperRequest) GetSchemasOk() (*[]EnumthirdPartyIdentityMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddIdentityMapperRequest) SetSchemas(v []EnumthirdPartyIdentityMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetMatchAttribute

`func (o *AddIdentityMapperRequest) GetMatchAttribute() []string`

GetMatchAttribute returns the MatchAttribute field if non-nil, zero value otherwise.

### GetMatchAttributeOk

`func (o *AddIdentityMapperRequest) GetMatchAttributeOk() (*[]string, bool)`

GetMatchAttributeOk returns a tuple with the MatchAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAttribute

`func (o *AddIdentityMapperRequest) SetMatchAttribute(v []string)`

SetMatchAttribute sets MatchAttribute field to given value.


### GetMatchBaseDN

`func (o *AddIdentityMapperRequest) GetMatchBaseDN() []string`

GetMatchBaseDN returns the MatchBaseDN field if non-nil, zero value otherwise.

### GetMatchBaseDNOk

`func (o *AddIdentityMapperRequest) GetMatchBaseDNOk() (*[]string, bool)`

GetMatchBaseDNOk returns a tuple with the MatchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchBaseDN

`func (o *AddIdentityMapperRequest) SetMatchBaseDN(v []string)`

SetMatchBaseDN sets MatchBaseDN field to given value.

### HasMatchBaseDN

`func (o *AddIdentityMapperRequest) HasMatchBaseDN() bool`

HasMatchBaseDN returns a boolean if a field has been set.

### GetMatchFilter

`func (o *AddIdentityMapperRequest) GetMatchFilter() string`

GetMatchFilter returns the MatchFilter field if non-nil, zero value otherwise.

### GetMatchFilterOk

`func (o *AddIdentityMapperRequest) GetMatchFilterOk() (*string, bool)`

GetMatchFilterOk returns a tuple with the MatchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchFilter

`func (o *AddIdentityMapperRequest) SetMatchFilter(v string)`

SetMatchFilter sets MatchFilter field to given value.

### HasMatchFilter

`func (o *AddIdentityMapperRequest) HasMatchFilter() bool`

HasMatchFilter returns a boolean if a field has been set.

### GetDescription

`func (o *AddIdentityMapperRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddIdentityMapperRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddIdentityMapperRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddIdentityMapperRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddIdentityMapperRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddIdentityMapperRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddIdentityMapperRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetScriptClass

`func (o *AddIdentityMapperRequest) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddIdentityMapperRequest) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddIdentityMapperRequest) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddIdentityMapperRequest) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddIdentityMapperRequest) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddIdentityMapperRequest) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddIdentityMapperRequest) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetMatchPattern

`func (o *AddIdentityMapperRequest) GetMatchPattern() string`

GetMatchPattern returns the MatchPattern field if non-nil, zero value otherwise.

### GetMatchPatternOk

`func (o *AddIdentityMapperRequest) GetMatchPatternOk() (*string, bool)`

GetMatchPatternOk returns a tuple with the MatchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPattern

`func (o *AddIdentityMapperRequest) SetMatchPattern(v string)`

SetMatchPattern sets MatchPattern field to given value.


### GetReplacePattern

`func (o *AddIdentityMapperRequest) GetReplacePattern() string`

GetReplacePattern returns the ReplacePattern field if non-nil, zero value otherwise.

### GetReplacePatternOk

`func (o *AddIdentityMapperRequest) GetReplacePatternOk() (*string, bool)`

GetReplacePatternOk returns a tuple with the ReplacePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacePattern

`func (o *AddIdentityMapperRequest) SetReplacePattern(v string)`

SetReplacePattern sets ReplacePattern field to given value.

### HasReplacePattern

`func (o *AddIdentityMapperRequest) HasReplacePattern() bool`

HasReplacePattern returns a boolean if a field has been set.

### GetAllIncludedIdentityMapper

`func (o *AddIdentityMapperRequest) GetAllIncludedIdentityMapper() []string`

GetAllIncludedIdentityMapper returns the AllIncludedIdentityMapper field if non-nil, zero value otherwise.

### GetAllIncludedIdentityMapperOk

`func (o *AddIdentityMapperRequest) GetAllIncludedIdentityMapperOk() (*[]string, bool)`

GetAllIncludedIdentityMapperOk returns a tuple with the AllIncludedIdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedIdentityMapper

`func (o *AddIdentityMapperRequest) SetAllIncludedIdentityMapper(v []string)`

SetAllIncludedIdentityMapper sets AllIncludedIdentityMapper field to given value.

### HasAllIncludedIdentityMapper

`func (o *AddIdentityMapperRequest) HasAllIncludedIdentityMapper() bool`

HasAllIncludedIdentityMapper returns a boolean if a field has been set.

### GetAnyIncludedIdentityMapper

`func (o *AddIdentityMapperRequest) GetAnyIncludedIdentityMapper() []string`

GetAnyIncludedIdentityMapper returns the AnyIncludedIdentityMapper field if non-nil, zero value otherwise.

### GetAnyIncludedIdentityMapperOk

`func (o *AddIdentityMapperRequest) GetAnyIncludedIdentityMapperOk() (*[]string, bool)`

GetAnyIncludedIdentityMapperOk returns a tuple with the AnyIncludedIdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedIdentityMapper

`func (o *AddIdentityMapperRequest) SetAnyIncludedIdentityMapper(v []string)`

SetAnyIncludedIdentityMapper sets AnyIncludedIdentityMapper field to given value.

### HasAnyIncludedIdentityMapper

`func (o *AddIdentityMapperRequest) HasAnyIncludedIdentityMapper() bool`

HasAnyIncludedIdentityMapper returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddIdentityMapperRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddIdentityMapperRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddIdentityMapperRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddIdentityMapperRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddIdentityMapperRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddIdentityMapperRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddIdentityMapperRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


