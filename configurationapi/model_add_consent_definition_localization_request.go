/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AddConsentDefinitionLocalizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddConsentDefinitionLocalizationRequest{}

// AddConsentDefinitionLocalizationRequest struct for AddConsentDefinitionLocalizationRequest
type AddConsentDefinitionLocalizationRequest struct {
	Schemas []EnumconsentDefinitionLocalizationSchemaUrn `json:"schemas,omitempty"`
	// The locale of this Consent Definition Localization.
	Locale string `json:"locale"`
	// The version of this Consent Definition Localization, using the format MAJOR.MINOR.
	Version string `json:"version"`
	// Localized text that may be used to provide a title or summary for a consent request or a granted consent.
	TitleText *string `json:"titleText,omitempty"`
	// Localized text describing the data to be shared.
	DataText string `json:"dataText"`
	// Localized text describing how the data is to be used.
	PurposeText string `json:"purposeText"`
	// Name of the new Consent Definition Localization
	LocalizationName string `json:"localizationName"`
}

// NewAddConsentDefinitionLocalizationRequest instantiates a new AddConsentDefinitionLocalizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddConsentDefinitionLocalizationRequest(locale string, version string, dataText string, purposeText string, localizationName string) *AddConsentDefinitionLocalizationRequest {
	this := AddConsentDefinitionLocalizationRequest{}
	this.Locale = locale
	this.Version = version
	this.DataText = dataText
	this.PurposeText = purposeText
	this.LocalizationName = localizationName
	return &this
}

// NewAddConsentDefinitionLocalizationRequestWithDefaults instantiates a new AddConsentDefinitionLocalizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddConsentDefinitionLocalizationRequestWithDefaults() *AddConsentDefinitionLocalizationRequest {
	this := AddConsentDefinitionLocalizationRequest{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddConsentDefinitionLocalizationRequest) GetSchemas() []EnumconsentDefinitionLocalizationSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumconsentDefinitionLocalizationSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddConsentDefinitionLocalizationRequest) GetSchemasOk() ([]EnumconsentDefinitionLocalizationSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddConsentDefinitionLocalizationRequest) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumconsentDefinitionLocalizationSchemaUrn and assigns it to the Schemas field.
func (o *AddConsentDefinitionLocalizationRequest) SetSchemas(v []EnumconsentDefinitionLocalizationSchemaUrn) {
	o.Schemas = v
}

// GetLocale returns the Locale field value
func (o *AddConsentDefinitionLocalizationRequest) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *AddConsentDefinitionLocalizationRequest) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *AddConsentDefinitionLocalizationRequest) SetLocale(v string) {
	o.Locale = v
}

// GetVersion returns the Version field value
func (o *AddConsentDefinitionLocalizationRequest) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AddConsentDefinitionLocalizationRequest) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *AddConsentDefinitionLocalizationRequest) SetVersion(v string) {
	o.Version = v
}

// GetTitleText returns the TitleText field value if set, zero value otherwise.
func (o *AddConsentDefinitionLocalizationRequest) GetTitleText() string {
	if o == nil || IsNil(o.TitleText) {
		var ret string
		return ret
	}
	return *o.TitleText
}

// GetTitleTextOk returns a tuple with the TitleText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddConsentDefinitionLocalizationRequest) GetTitleTextOk() (*string, bool) {
	if o == nil || IsNil(o.TitleText) {
		return nil, false
	}
	return o.TitleText, true
}

// HasTitleText returns a boolean if a field has been set.
func (o *AddConsentDefinitionLocalizationRequest) HasTitleText() bool {
	if o != nil && !IsNil(o.TitleText) {
		return true
	}

	return false
}

// SetTitleText gets a reference to the given string and assigns it to the TitleText field.
func (o *AddConsentDefinitionLocalizationRequest) SetTitleText(v string) {
	o.TitleText = &v
}

// GetDataText returns the DataText field value
func (o *AddConsentDefinitionLocalizationRequest) GetDataText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataText
}

// GetDataTextOk returns a tuple with the DataText field value
// and a boolean to check if the value has been set.
func (o *AddConsentDefinitionLocalizationRequest) GetDataTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataText, true
}

// SetDataText sets field value
func (o *AddConsentDefinitionLocalizationRequest) SetDataText(v string) {
	o.DataText = v
}

// GetPurposeText returns the PurposeText field value
func (o *AddConsentDefinitionLocalizationRequest) GetPurposeText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurposeText
}

// GetPurposeTextOk returns a tuple with the PurposeText field value
// and a boolean to check if the value has been set.
func (o *AddConsentDefinitionLocalizationRequest) GetPurposeTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurposeText, true
}

// SetPurposeText sets field value
func (o *AddConsentDefinitionLocalizationRequest) SetPurposeText(v string) {
	o.PurposeText = v
}

// GetLocalizationName returns the LocalizationName field value
func (o *AddConsentDefinitionLocalizationRequest) GetLocalizationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocalizationName
}

// GetLocalizationNameOk returns a tuple with the LocalizationName field value
// and a boolean to check if the value has been set.
func (o *AddConsentDefinitionLocalizationRequest) GetLocalizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocalizationName, true
}

// SetLocalizationName sets field value
func (o *AddConsentDefinitionLocalizationRequest) SetLocalizationName(v string) {
	o.LocalizationName = v
}

func (o AddConsentDefinitionLocalizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddConsentDefinitionLocalizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	toSerialize["locale"] = o.Locale
	toSerialize["version"] = o.Version
	if !IsNil(o.TitleText) {
		toSerialize["titleText"] = o.TitleText
	}
	toSerialize["dataText"] = o.DataText
	toSerialize["purposeText"] = o.PurposeText
	toSerialize["localizationName"] = o.LocalizationName
	return toSerialize, nil
}

type NullableAddConsentDefinitionLocalizationRequest struct {
	value *AddConsentDefinitionLocalizationRequest
	isSet bool
}

func (v NullableAddConsentDefinitionLocalizationRequest) Get() *AddConsentDefinitionLocalizationRequest {
	return v.value
}

func (v *NullableAddConsentDefinitionLocalizationRequest) Set(val *AddConsentDefinitionLocalizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddConsentDefinitionLocalizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddConsentDefinitionLocalizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddConsentDefinitionLocalizationRequest(val *AddConsentDefinitionLocalizationRequest) *NullableAddConsentDefinitionLocalizationRequest {
	return &NullableAddConsentDefinitionLocalizationRequest{value: val, isSet: true}
}

func (v NullableAddConsentDefinitionLocalizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddConsentDefinitionLocalizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
