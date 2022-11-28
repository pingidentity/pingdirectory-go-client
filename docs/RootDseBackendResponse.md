# RootDseBackendResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | Pointer to [**[]EnumrootDseBackendSchemaUrn**](EnumrootDseBackendSchemaUrn.md) |  | [optional] 
**SubordinateBaseDN** | Pointer to **[]string** | Specifies the set of base DNs used for singleLevel, wholeSubtree, and subordinateSubtree searches based at the root DSE. | [optional] 
**AdditionalSupportedControlOID** | Pointer to **[]string** | Specifies an additional OID that should appear in the list of supportedControl values in the server&#39;s root DSE. | [optional] 
**ShowAllAttributes** | **bool** | Indicates whether all attributes in the root DSE are to be treated like user attributes (and therefore returned to clients by default) regardless of the Directory Server schema configuration. | 
**UseLegacyVendorVersion** | Pointer to **bool** | Indicates whether the server&#39;s root DSE should reflect current or legacy values for the vendorName and vendorVersion attributes. | [optional] 

## Methods

### NewRootDseBackendResponse

`func NewRootDseBackendResponse(showAllAttributes bool, ) *RootDseBackendResponse`

NewRootDseBackendResponse instantiates a new RootDseBackendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootDseBackendResponseWithDefaults

`func NewRootDseBackendResponseWithDefaults() *RootDseBackendResponse`

NewRootDseBackendResponseWithDefaults instantiates a new RootDseBackendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *RootDseBackendResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RootDseBackendResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RootDseBackendResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RootDseBackendResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *RootDseBackendResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *RootDseBackendResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *RootDseBackendResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *RootDseBackendResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *RootDseBackendResponse) GetSchemas() []EnumrootDseBackendSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *RootDseBackendResponse) GetSchemasOk() (*[]EnumrootDseBackendSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *RootDseBackendResponse) SetSchemas(v []EnumrootDseBackendSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *RootDseBackendResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetSubordinateBaseDN

`func (o *RootDseBackendResponse) GetSubordinateBaseDN() []string`

GetSubordinateBaseDN returns the SubordinateBaseDN field if non-nil, zero value otherwise.

### GetSubordinateBaseDNOk

`func (o *RootDseBackendResponse) GetSubordinateBaseDNOk() (*[]string, bool)`

GetSubordinateBaseDNOk returns a tuple with the SubordinateBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubordinateBaseDN

`func (o *RootDseBackendResponse) SetSubordinateBaseDN(v []string)`

SetSubordinateBaseDN sets SubordinateBaseDN field to given value.

### HasSubordinateBaseDN

`func (o *RootDseBackendResponse) HasSubordinateBaseDN() bool`

HasSubordinateBaseDN returns a boolean if a field has been set.

### GetAdditionalSupportedControlOID

`func (o *RootDseBackendResponse) GetAdditionalSupportedControlOID() []string`

GetAdditionalSupportedControlOID returns the AdditionalSupportedControlOID field if non-nil, zero value otherwise.

### GetAdditionalSupportedControlOIDOk

`func (o *RootDseBackendResponse) GetAdditionalSupportedControlOIDOk() (*[]string, bool)`

GetAdditionalSupportedControlOIDOk returns a tuple with the AdditionalSupportedControlOID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSupportedControlOID

`func (o *RootDseBackendResponse) SetAdditionalSupportedControlOID(v []string)`

SetAdditionalSupportedControlOID sets AdditionalSupportedControlOID field to given value.

### HasAdditionalSupportedControlOID

`func (o *RootDseBackendResponse) HasAdditionalSupportedControlOID() bool`

HasAdditionalSupportedControlOID returns a boolean if a field has been set.

### GetShowAllAttributes

`func (o *RootDseBackendResponse) GetShowAllAttributes() bool`

GetShowAllAttributes returns the ShowAllAttributes field if non-nil, zero value otherwise.

### GetShowAllAttributesOk

`func (o *RootDseBackendResponse) GetShowAllAttributesOk() (*bool, bool)`

GetShowAllAttributesOk returns a tuple with the ShowAllAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAllAttributes

`func (o *RootDseBackendResponse) SetShowAllAttributes(v bool)`

SetShowAllAttributes sets ShowAllAttributes field to given value.


### GetUseLegacyVendorVersion

`func (o *RootDseBackendResponse) GetUseLegacyVendorVersion() bool`

GetUseLegacyVendorVersion returns the UseLegacyVendorVersion field if non-nil, zero value otherwise.

### GetUseLegacyVendorVersionOk

`func (o *RootDseBackendResponse) GetUseLegacyVendorVersionOk() (*bool, bool)`

GetUseLegacyVendorVersionOk returns a tuple with the UseLegacyVendorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLegacyVendorVersion

`func (o *RootDseBackendResponse) SetUseLegacyVendorVersion(v bool)`

SetUseLegacyVendorVersion sets UseLegacyVendorVersion field to given value.

### HasUseLegacyVendorVersion

`func (o *RootDseBackendResponse) HasUseLegacyVendorVersion() bool`

HasUseLegacyVendorVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


