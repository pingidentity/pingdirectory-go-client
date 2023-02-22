# LicenseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | Pointer to [**[]EnumlicenseSchemaUrn**](EnumlicenseSchemaUrn.md) |  | [optional] 
**DirectoryPlatformLicenseKey** | Pointer to **string** | License key enabling use of Directory Server, Directory Proxy Server, Data Sync Server, and Data Metrics Server products. | [optional] 

## Methods

### NewLicenseResponse

`func NewLicenseResponse() *LicenseResponse`

NewLicenseResponse instantiates a new LicenseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseResponseWithDefaults

`func NewLicenseResponseWithDefaults() *LicenseResponse`

NewLicenseResponseWithDefaults instantiates a new LicenseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *LicenseResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LicenseResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LicenseResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LicenseResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *LicenseResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *LicenseResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *LicenseResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *LicenseResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *LicenseResponse) GetSchemas() []EnumlicenseSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LicenseResponse) GetSchemasOk() (*[]EnumlicenseSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LicenseResponse) SetSchemas(v []EnumlicenseSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *LicenseResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDirectoryPlatformLicenseKey

`func (o *LicenseResponse) GetDirectoryPlatformLicenseKey() string`

GetDirectoryPlatformLicenseKey returns the DirectoryPlatformLicenseKey field if non-nil, zero value otherwise.

### GetDirectoryPlatformLicenseKeyOk

`func (o *LicenseResponse) GetDirectoryPlatformLicenseKeyOk() (*string, bool)`

GetDirectoryPlatformLicenseKeyOk returns a tuple with the DirectoryPlatformLicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryPlatformLicenseKey

`func (o *LicenseResponse) SetDirectoryPlatformLicenseKey(v string)`

SetDirectoryPlatformLicenseKey sets DirectoryPlatformLicenseKey field to given value.

### HasDirectoryPlatformLicenseKey

`func (o *LicenseResponse) HasDirectoryPlatformLicenseKey() bool`

HasDirectoryPlatformLicenseKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


