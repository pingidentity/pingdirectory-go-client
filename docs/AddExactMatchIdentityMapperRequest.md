# AddExactMatchIdentityMapperRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MapperName** | **string** | Name of the new Identity Mapper | 
**Schemas** | [**[]EnumexactMatchIdentityMapperSchemaUrn**](EnumexactMatchIdentityMapperSchemaUrn.md) |  | 
**MatchAttribute** | **[]string** | Specifies the attribute whose value should exactly match the ID string provided to this identity mapper. | 
**MatchBaseDN** | Pointer to **[]string** | Specifies the set of base DNs below which to search for users. | [optional] 
**MatchFilter** | Pointer to **string** | An optional filter that mapped users must match. | [optional] 
**Description** | Pointer to **string** | A description for this Identity Mapper | [optional] 
**Enabled** | **bool** | Indicates whether the Identity Mapper is enabled for use. | 

## Methods

### NewAddExactMatchIdentityMapperRequest

`func NewAddExactMatchIdentityMapperRequest(mapperName string, schemas []EnumexactMatchIdentityMapperSchemaUrn, matchAttribute []string, enabled bool, ) *AddExactMatchIdentityMapperRequest`

NewAddExactMatchIdentityMapperRequest instantiates a new AddExactMatchIdentityMapperRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddExactMatchIdentityMapperRequestWithDefaults

`func NewAddExactMatchIdentityMapperRequestWithDefaults() *AddExactMatchIdentityMapperRequest`

NewAddExactMatchIdentityMapperRequestWithDefaults instantiates a new AddExactMatchIdentityMapperRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMapperName

`func (o *AddExactMatchIdentityMapperRequest) GetMapperName() string`

GetMapperName returns the MapperName field if non-nil, zero value otherwise.

### GetMapperNameOk

`func (o *AddExactMatchIdentityMapperRequest) GetMapperNameOk() (*string, bool)`

GetMapperNameOk returns a tuple with the MapperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperName

`func (o *AddExactMatchIdentityMapperRequest) SetMapperName(v string)`

SetMapperName sets MapperName field to given value.


### GetSchemas

`func (o *AddExactMatchIdentityMapperRequest) GetSchemas() []EnumexactMatchIdentityMapperSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddExactMatchIdentityMapperRequest) GetSchemasOk() (*[]EnumexactMatchIdentityMapperSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddExactMatchIdentityMapperRequest) SetSchemas(v []EnumexactMatchIdentityMapperSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetMatchAttribute

`func (o *AddExactMatchIdentityMapperRequest) GetMatchAttribute() []string`

GetMatchAttribute returns the MatchAttribute field if non-nil, zero value otherwise.

### GetMatchAttributeOk

`func (o *AddExactMatchIdentityMapperRequest) GetMatchAttributeOk() (*[]string, bool)`

GetMatchAttributeOk returns a tuple with the MatchAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAttribute

`func (o *AddExactMatchIdentityMapperRequest) SetMatchAttribute(v []string)`

SetMatchAttribute sets MatchAttribute field to given value.


### GetMatchBaseDN

`func (o *AddExactMatchIdentityMapperRequest) GetMatchBaseDN() []string`

GetMatchBaseDN returns the MatchBaseDN field if non-nil, zero value otherwise.

### GetMatchBaseDNOk

`func (o *AddExactMatchIdentityMapperRequest) GetMatchBaseDNOk() (*[]string, bool)`

GetMatchBaseDNOk returns a tuple with the MatchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchBaseDN

`func (o *AddExactMatchIdentityMapperRequest) SetMatchBaseDN(v []string)`

SetMatchBaseDN sets MatchBaseDN field to given value.

### HasMatchBaseDN

`func (o *AddExactMatchIdentityMapperRequest) HasMatchBaseDN() bool`

HasMatchBaseDN returns a boolean if a field has been set.

### GetMatchFilter

`func (o *AddExactMatchIdentityMapperRequest) GetMatchFilter() string`

GetMatchFilter returns the MatchFilter field if non-nil, zero value otherwise.

### GetMatchFilterOk

`func (o *AddExactMatchIdentityMapperRequest) GetMatchFilterOk() (*string, bool)`

GetMatchFilterOk returns a tuple with the MatchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchFilter

`func (o *AddExactMatchIdentityMapperRequest) SetMatchFilter(v string)`

SetMatchFilter sets MatchFilter field to given value.

### HasMatchFilter

`func (o *AddExactMatchIdentityMapperRequest) HasMatchFilter() bool`

HasMatchFilter returns a boolean if a field has been set.

### GetDescription

`func (o *AddExactMatchIdentityMapperRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddExactMatchIdentityMapperRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddExactMatchIdentityMapperRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddExactMatchIdentityMapperRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddExactMatchIdentityMapperRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddExactMatchIdentityMapperRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddExactMatchIdentityMapperRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


