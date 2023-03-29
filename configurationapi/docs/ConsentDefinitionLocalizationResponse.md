# ConsentDefinitionLocalizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Consent Definition | 
**Schemas** | Pointer to [**[]EnumconsentDefinitionLocalizationSchemaUrn**](EnumconsentDefinitionLocalizationSchemaUrn.md) |  | [optional] 
**Locale** | **string** | The locale of this Consent Definition Localization. | 
**Version** | **string** | The version of this Consent Definition Localization, using the format MAJOR.MINOR. | 
**TitleText** | Pointer to **string** | Localized text that may be used to provide a title or summary for a consent request or a granted consent. | [optional] 
**DataText** | **string** | Localized text describing the data to be shared. | 
**PurposeText** | **string** | Localized text describing how the data is to be used. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewConsentDefinitionLocalizationResponse

`func NewConsentDefinitionLocalizationResponse(id string, locale string, version string, dataText string, purposeText string, ) *ConsentDefinitionLocalizationResponse`

NewConsentDefinitionLocalizationResponse instantiates a new ConsentDefinitionLocalizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentDefinitionLocalizationResponseWithDefaults

`func NewConsentDefinitionLocalizationResponseWithDefaults() *ConsentDefinitionLocalizationResponse`

NewConsentDefinitionLocalizationResponseWithDefaults instantiates a new ConsentDefinitionLocalizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsentDefinitionLocalizationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsentDefinitionLocalizationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsentDefinitionLocalizationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ConsentDefinitionLocalizationResponse) GetSchemas() []EnumconsentDefinitionLocalizationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ConsentDefinitionLocalizationResponse) GetSchemasOk() (*[]EnumconsentDefinitionLocalizationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ConsentDefinitionLocalizationResponse) SetSchemas(v []EnumconsentDefinitionLocalizationSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ConsentDefinitionLocalizationResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetLocale

`func (o *ConsentDefinitionLocalizationResponse) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *ConsentDefinitionLocalizationResponse) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *ConsentDefinitionLocalizationResponse) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetVersion

`func (o *ConsentDefinitionLocalizationResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConsentDefinitionLocalizationResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConsentDefinitionLocalizationResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetTitleText

`func (o *ConsentDefinitionLocalizationResponse) GetTitleText() string`

GetTitleText returns the TitleText field if non-nil, zero value otherwise.

### GetTitleTextOk

`func (o *ConsentDefinitionLocalizationResponse) GetTitleTextOk() (*string, bool)`

GetTitleTextOk returns a tuple with the TitleText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleText

`func (o *ConsentDefinitionLocalizationResponse) SetTitleText(v string)`

SetTitleText sets TitleText field to given value.

### HasTitleText

`func (o *ConsentDefinitionLocalizationResponse) HasTitleText() bool`

HasTitleText returns a boolean if a field has been set.

### GetDataText

`func (o *ConsentDefinitionLocalizationResponse) GetDataText() string`

GetDataText returns the DataText field if non-nil, zero value otherwise.

### GetDataTextOk

`func (o *ConsentDefinitionLocalizationResponse) GetDataTextOk() (*string, bool)`

GetDataTextOk returns a tuple with the DataText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataText

`func (o *ConsentDefinitionLocalizationResponse) SetDataText(v string)`

SetDataText sets DataText field to given value.


### GetPurposeText

`func (o *ConsentDefinitionLocalizationResponse) GetPurposeText() string`

GetPurposeText returns the PurposeText field if non-nil, zero value otherwise.

### GetPurposeTextOk

`func (o *ConsentDefinitionLocalizationResponse) GetPurposeTextOk() (*string, bool)`

GetPurposeTextOk returns a tuple with the PurposeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeText

`func (o *ConsentDefinitionLocalizationResponse) SetPurposeText(v string)`

SetPurposeText sets PurposeText field to given value.


### GetMeta

`func (o *ConsentDefinitionLocalizationResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConsentDefinitionLocalizationResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConsentDefinitionLocalizationResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ConsentDefinitionLocalizationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ConsentDefinitionLocalizationResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ConsentDefinitionLocalizationResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ConsentDefinitionLocalizationResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ConsentDefinitionLocalizationResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


