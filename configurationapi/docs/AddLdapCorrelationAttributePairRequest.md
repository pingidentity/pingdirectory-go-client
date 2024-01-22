# AddLdapCorrelationAttributePairRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumldapCorrelationAttributePairSchemaUrn**](EnumldapCorrelationAttributePairSchemaUrn.md) |  | [optional] 
**PrimaryCorrelationAttribute** | **string** | The LDAP attribute from the base SCIM Resource Type whose value will be used to match objects in the Correlated LDAP Data View. | 
**SecondaryCorrelationAttribute** | **string** | The LDAP attribute from the Correlated LDAP Data View whose value will be matched. | 
**PairName** | **string** | Name of the new LDAP Correlation Attribute Pair | 

## Methods

### NewAddLdapCorrelationAttributePairRequest

`func NewAddLdapCorrelationAttributePairRequest(primaryCorrelationAttribute string, secondaryCorrelationAttribute string, pairName string, ) *AddLdapCorrelationAttributePairRequest`

NewAddLdapCorrelationAttributePairRequest instantiates a new AddLdapCorrelationAttributePairRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLdapCorrelationAttributePairRequestWithDefaults

`func NewAddLdapCorrelationAttributePairRequestWithDefaults() *AddLdapCorrelationAttributePairRequest`

NewAddLdapCorrelationAttributePairRequestWithDefaults instantiates a new AddLdapCorrelationAttributePairRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddLdapCorrelationAttributePairRequest) GetSchemas() []EnumldapCorrelationAttributePairSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLdapCorrelationAttributePairRequest) GetSchemasOk() (*[]EnumldapCorrelationAttributePairSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLdapCorrelationAttributePairRequest) SetSchemas(v []EnumldapCorrelationAttributePairSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddLdapCorrelationAttributePairRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetPrimaryCorrelationAttribute

`func (o *AddLdapCorrelationAttributePairRequest) GetPrimaryCorrelationAttribute() string`

GetPrimaryCorrelationAttribute returns the PrimaryCorrelationAttribute field if non-nil, zero value otherwise.

### GetPrimaryCorrelationAttributeOk

`func (o *AddLdapCorrelationAttributePairRequest) GetPrimaryCorrelationAttributeOk() (*string, bool)`

GetPrimaryCorrelationAttributeOk returns a tuple with the PrimaryCorrelationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCorrelationAttribute

`func (o *AddLdapCorrelationAttributePairRequest) SetPrimaryCorrelationAttribute(v string)`

SetPrimaryCorrelationAttribute sets PrimaryCorrelationAttribute field to given value.


### GetSecondaryCorrelationAttribute

`func (o *AddLdapCorrelationAttributePairRequest) GetSecondaryCorrelationAttribute() string`

GetSecondaryCorrelationAttribute returns the SecondaryCorrelationAttribute field if non-nil, zero value otherwise.

### GetSecondaryCorrelationAttributeOk

`func (o *AddLdapCorrelationAttributePairRequest) GetSecondaryCorrelationAttributeOk() (*string, bool)`

GetSecondaryCorrelationAttributeOk returns a tuple with the SecondaryCorrelationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCorrelationAttribute

`func (o *AddLdapCorrelationAttributePairRequest) SetSecondaryCorrelationAttribute(v string)`

SetSecondaryCorrelationAttribute sets SecondaryCorrelationAttribute field to given value.


### GetPairName

`func (o *AddLdapCorrelationAttributePairRequest) GetPairName() string`

GetPairName returns the PairName field if non-nil, zero value otherwise.

### GetPairNameOk

`func (o *AddLdapCorrelationAttributePairRequest) GetPairNameOk() (*string, bool)`

GetPairNameOk returns a tuple with the PairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairName

`func (o *AddLdapCorrelationAttributePairRequest) SetPairName(v string)`

SetPairName sets PairName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


