# ExactMatchIdentityMapperResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Identity Mapper | 
**Schemas** | [**[]EnumexactMatchIdentityMapperSchemaUrn**](EnumexactMatchIdentityMapperSchemaUrn.md) |  | 
**MatchAttribute** | **[]string** | Specifies the attribute whose value should exactly match the ID string provided to this identity mapper. | 
**MatchBaseDN** | Pointer to **[]string** | Specifies the set of base DNs below which to search for users. | [optional] 
**MatchFilter** | Pointer to **string** | An optional filter that mapped users must match. | [optional] 
**Description** | Pointer to **string** | A description for this Identity Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Identity Mapper is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewExactMatchIdentityMapperResponse

`func NewExactMatchIdentityMapperResponse(id string, schemas []EnumexactMatchIdentityMapperSchemaUrn, matchAttribute []string, enabled bool, ) *ExactMatchIdentityMapperResponse`

NewExactMatchIdentityMapperResponse instantiates a new ExactMatchIdentityMapperResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExactMatchIdentityMapperResponseWithDefaults

`func NewExactMatchIdentityMapperResponseWithDefaults() *ExactMatchIdentityMapperResponse`

NewExactMatchIdentityMapperResponseWithDefaults instantiates a new ExactMatchIdentityMapperResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExactMatchIdentityMapperResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExactMatchIdentityMapperResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExactMatchIdentityMapperResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ExactMatchIdentityMapperResponse) GetSchemas() []EnumexactMatchIdentityMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ExactMatchIdentityMapperResponse) GetSchemasOk() (*[]EnumexactMatchIdentityMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ExactMatchIdentityMapperResponse) SetSchemas(v []EnumexactMatchIdentityMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetMatchAttribute

`func (o *ExactMatchIdentityMapperResponse) GetMatchAttribute() []string`

GetMatchAttribute returns the MatchAttribute field if non-nil, zero value otherwise.

### GetMatchAttributeOk

`func (o *ExactMatchIdentityMapperResponse) GetMatchAttributeOk() (*[]string, bool)`

GetMatchAttributeOk returns a tuple with the MatchAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAttribute

`func (o *ExactMatchIdentityMapperResponse) SetMatchAttribute(v []string)`

SetMatchAttribute sets MatchAttribute field to given value.


### GetMatchBaseDN

`func (o *ExactMatchIdentityMapperResponse) GetMatchBaseDN() []string`

GetMatchBaseDN returns the MatchBaseDN field if non-nil, zero value otherwise.

### GetMatchBaseDNOk

`func (o *ExactMatchIdentityMapperResponse) GetMatchBaseDNOk() (*[]string, bool)`

GetMatchBaseDNOk returns a tuple with the MatchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchBaseDN

`func (o *ExactMatchIdentityMapperResponse) SetMatchBaseDN(v []string)`

SetMatchBaseDN sets MatchBaseDN field to given value.

### HasMatchBaseDN

`func (o *ExactMatchIdentityMapperResponse) HasMatchBaseDN() bool`

HasMatchBaseDN returns a boolean if a field has been set.

### GetMatchFilter

`func (o *ExactMatchIdentityMapperResponse) GetMatchFilter() string`

GetMatchFilter returns the MatchFilter field if non-nil, zero value otherwise.

### GetMatchFilterOk

`func (o *ExactMatchIdentityMapperResponse) GetMatchFilterOk() (*string, bool)`

GetMatchFilterOk returns a tuple with the MatchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchFilter

`func (o *ExactMatchIdentityMapperResponse) SetMatchFilter(v string)`

SetMatchFilter sets MatchFilter field to given value.

### HasMatchFilter

`func (o *ExactMatchIdentityMapperResponse) HasMatchFilter() bool`

HasMatchFilter returns a boolean if a field has been set.

### GetDescription

`func (o *ExactMatchIdentityMapperResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExactMatchIdentityMapperResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExactMatchIdentityMapperResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExactMatchIdentityMapperResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ExactMatchIdentityMapperResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExactMatchIdentityMapperResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExactMatchIdentityMapperResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *ExactMatchIdentityMapperResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ExactMatchIdentityMapperResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ExactMatchIdentityMapperResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ExactMatchIdentityMapperResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


