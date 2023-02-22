# DelegatedAdminCorrelatedRestResourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Delegated Admin Correlated REST Resource | 
**Schemas** | Pointer to [**[]EnumdelegatedAdminCorrelatedRestResourceSchemaUrn**](EnumdelegatedAdminCorrelatedRestResourceSchemaUrn.md) |  | [optional] 
**DisplayName** | **string** | A human readable display name for this Delegated Admin Correlated REST Resource. | 
**CorrelatedRESTResource** | **string** | The REST Resource Type that will be linked to this REST Resource Type. | 
**PrimaryRESTResourceCorrelationAttribute** | **string** | The LDAP attribute from the parent REST Resource Type whose value will be used to match objects in the Delegated Admin Correlated REST Resource. This attribute must be writeable when use-secondary-value-for-linking is enabled. | 
**SecondaryRESTResourceCorrelationAttribute** | **string** | The LDAP attribute from the Delegated Admin Correlated REST Resource whose value will be matched with the primary-rest-resource-correlation-attribute. This attribute must be writeable when use-secondary-value-for-linking is disabled. | 
**UseSecondaryValueForLinking** | Pointer to **bool** | Indicates whether links should be created using the secondary correlation attribute value. | [optional] 

## Methods

### NewDelegatedAdminCorrelatedRestResourceResponse

`func NewDelegatedAdminCorrelatedRestResourceResponse(id string, displayName string, correlatedRESTResource string, primaryRESTResourceCorrelationAttribute string, secondaryRESTResourceCorrelationAttribute string, ) *DelegatedAdminCorrelatedRestResourceResponse`

NewDelegatedAdminCorrelatedRestResourceResponse instantiates a new DelegatedAdminCorrelatedRestResourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegatedAdminCorrelatedRestResourceResponseWithDefaults

`func NewDelegatedAdminCorrelatedRestResourceResponseWithDefaults() *DelegatedAdminCorrelatedRestResourceResponse`

NewDelegatedAdminCorrelatedRestResourceResponseWithDefaults instantiates a new DelegatedAdminCorrelatedRestResourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DelegatedAdminCorrelatedRestResourceResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DelegatedAdminCorrelatedRestResourceResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *DelegatedAdminCorrelatedRestResourceResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *DelegatedAdminCorrelatedRestResourceResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DelegatedAdminCorrelatedRestResourceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetSchemas() []EnumdelegatedAdminCorrelatedRestResourceSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetSchemasOk() (*[]EnumdelegatedAdminCorrelatedRestResourceSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DelegatedAdminCorrelatedRestResourceResponse) SetSchemas(v []EnumdelegatedAdminCorrelatedRestResourceSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *DelegatedAdminCorrelatedRestResourceResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDisplayName

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DelegatedAdminCorrelatedRestResourceResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetCorrelatedRESTResource

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetCorrelatedRESTResource() string`

GetCorrelatedRESTResource returns the CorrelatedRESTResource field if non-nil, zero value otherwise.

### GetCorrelatedRESTResourceOk

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetCorrelatedRESTResourceOk() (*string, bool)`

GetCorrelatedRESTResourceOk returns a tuple with the CorrelatedRESTResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedRESTResource

`func (o *DelegatedAdminCorrelatedRestResourceResponse) SetCorrelatedRESTResource(v string)`

SetCorrelatedRESTResource sets CorrelatedRESTResource field to given value.


### GetPrimaryRESTResourceCorrelationAttribute

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetPrimaryRESTResourceCorrelationAttribute() string`

GetPrimaryRESTResourceCorrelationAttribute returns the PrimaryRESTResourceCorrelationAttribute field if non-nil, zero value otherwise.

### GetPrimaryRESTResourceCorrelationAttributeOk

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetPrimaryRESTResourceCorrelationAttributeOk() (*string, bool)`

GetPrimaryRESTResourceCorrelationAttributeOk returns a tuple with the PrimaryRESTResourceCorrelationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryRESTResourceCorrelationAttribute

`func (o *DelegatedAdminCorrelatedRestResourceResponse) SetPrimaryRESTResourceCorrelationAttribute(v string)`

SetPrimaryRESTResourceCorrelationAttribute sets PrimaryRESTResourceCorrelationAttribute field to given value.


### GetSecondaryRESTResourceCorrelationAttribute

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetSecondaryRESTResourceCorrelationAttribute() string`

GetSecondaryRESTResourceCorrelationAttribute returns the SecondaryRESTResourceCorrelationAttribute field if non-nil, zero value otherwise.

### GetSecondaryRESTResourceCorrelationAttributeOk

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetSecondaryRESTResourceCorrelationAttributeOk() (*string, bool)`

GetSecondaryRESTResourceCorrelationAttributeOk returns a tuple with the SecondaryRESTResourceCorrelationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRESTResourceCorrelationAttribute

`func (o *DelegatedAdminCorrelatedRestResourceResponse) SetSecondaryRESTResourceCorrelationAttribute(v string)`

SetSecondaryRESTResourceCorrelationAttribute sets SecondaryRESTResourceCorrelationAttribute field to given value.


### GetUseSecondaryValueForLinking

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetUseSecondaryValueForLinking() bool`

GetUseSecondaryValueForLinking returns the UseSecondaryValueForLinking field if non-nil, zero value otherwise.

### GetUseSecondaryValueForLinkingOk

`func (o *DelegatedAdminCorrelatedRestResourceResponse) GetUseSecondaryValueForLinkingOk() (*bool, bool)`

GetUseSecondaryValueForLinkingOk returns a tuple with the UseSecondaryValueForLinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSecondaryValueForLinking

`func (o *DelegatedAdminCorrelatedRestResourceResponse) SetUseSecondaryValueForLinking(v bool)`

SetUseSecondaryValueForLinking sets UseSecondaryValueForLinking field to given value.

### HasUseSecondaryValueForLinking

`func (o *DelegatedAdminCorrelatedRestResourceResponse) HasUseSecondaryValueForLinking() bool`

HasUseSecondaryValueForLinking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


