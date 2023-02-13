# AddConsentDefinitionLocalizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalizationName** | **string** | Name of the new Consent Definition Localization | 
**Schemas** | Pointer to [**[]EnumconsentDefinitionLocalizationSchemaUrn**](EnumconsentDefinitionLocalizationSchemaUrn.md) |  | [optional] 
**Locale** | **string** | The locale of this Consent Definition Localization. | 
**Version** | **string** | The version of this Consent Definition Localization, using the format MAJOR.MINOR. | 
**TitleText** | Pointer to **string** | Localized text that may be used to provide a title or summary for a consent request or a granted consent. | [optional] 
**DataText** | **string** | Localized text describing the data to be shared. | 
**PurposeText** | **string** | Localized text describing how the data is to be used. | 

## Methods

### NewAddConsentDefinitionLocalizationRequest

`func NewAddConsentDefinitionLocalizationRequest(localizationName string, locale string, version string, dataText string, purposeText string, ) *AddConsentDefinitionLocalizationRequest`

NewAddConsentDefinitionLocalizationRequest instantiates a new AddConsentDefinitionLocalizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddConsentDefinitionLocalizationRequestWithDefaults

`func NewAddConsentDefinitionLocalizationRequestWithDefaults() *AddConsentDefinitionLocalizationRequest`

NewAddConsentDefinitionLocalizationRequestWithDefaults instantiates a new AddConsentDefinitionLocalizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalizationName

`func (o *AddConsentDefinitionLocalizationRequest) GetLocalizationName() string`

GetLocalizationName returns the LocalizationName field if non-nil, zero value otherwise.

### GetLocalizationNameOk

`func (o *AddConsentDefinitionLocalizationRequest) GetLocalizationNameOk() (*string, bool)`

GetLocalizationNameOk returns a tuple with the LocalizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizationName

`func (o *AddConsentDefinitionLocalizationRequest) SetLocalizationName(v string)`

SetLocalizationName sets LocalizationName field to given value.


### GetSchemas

`func (o *AddConsentDefinitionLocalizationRequest) GetSchemas() []EnumconsentDefinitionLocalizationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddConsentDefinitionLocalizationRequest) GetSchemasOk() (*[]EnumconsentDefinitionLocalizationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddConsentDefinitionLocalizationRequest) SetSchemas(v []EnumconsentDefinitionLocalizationSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddConsentDefinitionLocalizationRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetLocale

`func (o *AddConsentDefinitionLocalizationRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *AddConsentDefinitionLocalizationRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *AddConsentDefinitionLocalizationRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetVersion

`func (o *AddConsentDefinitionLocalizationRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AddConsentDefinitionLocalizationRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AddConsentDefinitionLocalizationRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetTitleText

`func (o *AddConsentDefinitionLocalizationRequest) GetTitleText() string`

GetTitleText returns the TitleText field if non-nil, zero value otherwise.

### GetTitleTextOk

`func (o *AddConsentDefinitionLocalizationRequest) GetTitleTextOk() (*string, bool)`

GetTitleTextOk returns a tuple with the TitleText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleText

`func (o *AddConsentDefinitionLocalizationRequest) SetTitleText(v string)`

SetTitleText sets TitleText field to given value.

### HasTitleText

`func (o *AddConsentDefinitionLocalizationRequest) HasTitleText() bool`

HasTitleText returns a boolean if a field has been set.

### GetDataText

`func (o *AddConsentDefinitionLocalizationRequest) GetDataText() string`

GetDataText returns the DataText field if non-nil, zero value otherwise.

### GetDataTextOk

`func (o *AddConsentDefinitionLocalizationRequest) GetDataTextOk() (*string, bool)`

GetDataTextOk returns a tuple with the DataText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataText

`func (o *AddConsentDefinitionLocalizationRequest) SetDataText(v string)`

SetDataText sets DataText field to given value.


### GetPurposeText

`func (o *AddConsentDefinitionLocalizationRequest) GetPurposeText() string`

GetPurposeText returns the PurposeText field if non-nil, zero value otherwise.

### GetPurposeTextOk

`func (o *AddConsentDefinitionLocalizationRequest) GetPurposeTextOk() (*string, bool)`

GetPurposeTextOk returns a tuple with the PurposeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeText

`func (o *AddConsentDefinitionLocalizationRequest) SetPurposeText(v string)`

SetPurposeText sets PurposeText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


