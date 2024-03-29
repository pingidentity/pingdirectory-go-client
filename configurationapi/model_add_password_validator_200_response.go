/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"fmt"
)

// AddPasswordValidator200Response - struct for AddPasswordValidator200Response
type AddPasswordValidator200Response struct {
	AttributeValuePasswordValidatorResponse       *AttributeValuePasswordValidatorResponse
	CharacterSetPasswordValidatorResponse         *CharacterSetPasswordValidatorResponse
	DictionaryPasswordValidatorResponse           *DictionaryPasswordValidatorResponse
	DisallowedCharactersPasswordValidatorResponse *DisallowedCharactersPasswordValidatorResponse
	GroovyScriptedPasswordValidatorResponse       *GroovyScriptedPasswordValidatorResponse
	HaystackPasswordValidatorResponse             *HaystackPasswordValidatorResponse
	LengthBasedPasswordValidatorResponse          *LengthBasedPasswordValidatorResponse
	PwnedPasswordsPasswordValidatorResponse       *PwnedPasswordsPasswordValidatorResponse
	RegularExpressionPasswordValidatorResponse    *RegularExpressionPasswordValidatorResponse
	RepeatedCharactersPasswordValidatorResponse   *RepeatedCharactersPasswordValidatorResponse
	SimilarityBasedPasswordValidatorResponse      *SimilarityBasedPasswordValidatorResponse
	ThirdPartyPasswordValidatorResponse           *ThirdPartyPasswordValidatorResponse
	UniqueCharactersPasswordValidatorResponse     *UniqueCharactersPasswordValidatorResponse
	Utf8PasswordValidatorResponse                 *Utf8PasswordValidatorResponse
}

// AttributeValuePasswordValidatorResponseAsAddPasswordValidator200Response is a convenience function that returns AttributeValuePasswordValidatorResponse wrapped in AddPasswordValidator200Response
func AttributeValuePasswordValidatorResponseAsAddPasswordValidator200Response(v *AttributeValuePasswordValidatorResponse) AddPasswordValidator200Response {
	return AddPasswordValidator200Response{
		AttributeValuePasswordValidatorResponse: v,
	}
}

// CharacterSetPasswordValidatorResponseAsAddPasswordValidator200Response is a convenience function that returns CharacterSetPasswordValidatorResponse wrapped in AddPasswordValidator200Response
func CharacterSetPasswordValidatorResponseAsAddPasswordValidator200Response(v *CharacterSetPasswordValidatorResponse) AddPasswordValidator200Response {
	return AddPasswordValidator200Response{
		CharacterSetPasswordValidatorResponse: v,
	}
}

// DictionaryPasswordValidatorResponseAsAddPasswordValidator200Response is a convenience function that returns DictionaryPasswordValidatorResponse wrapped in AddPasswordValidator200Response
func DictionaryPasswordValidatorResponseAsAddPasswordValidator200Response(v *DictionaryPasswordValidatorResponse) AddPasswordValidator200Response {
	return AddPasswordValidator200Response{
		DictionaryPasswordValidatorResponse: v,
	}
}

// DisallowedCharactersPasswordValidatorResponseAsAddPasswordValidator200Response is a convenience function that returns DisallowedCharactersPasswordValidatorResponse wrapped in AddPasswordValidator200Response
func DisallowedCharactersPasswordValidatorResponseAsAddPasswordValidator200Response(v *DisallowedCharactersPasswordValidatorResponse) AddPasswordValidator200Response {
	return AddPasswordValidator200Response{
		DisallowedCharactersPasswordValidatorResponse: v,
	}
}

// GroovyScriptedPasswordValidatorResponseAsAddPasswordValidator200Response is a convenience function that returns GroovyScriptedPasswordValidatorResponse wrapped in AddPasswordValidator200Response
func GroovyScriptedPasswordValidatorResponseAsAddPasswordValidator200Response(v *GroovyScriptedPasswordValidatorResponse) AddPasswordValidator200Response {
	return AddPasswordValidator200Response{
		GroovyScriptedPasswordValidatorResponse: v,
	}
}

// HaystackPasswordValidatorResponseAsAddPasswordValidator200Response is a convenience function that returns HaystackPasswordValidatorResponse wrapped in AddPasswordValidator200Response
func HaystackPasswordValidatorResponseAsAddPasswordValidator200Response(v *HaystackPasswordValidatorResponse) AddPasswordValidator200Response {
	return AddPasswordValidator200Response{
		HaystackPasswordValidatorResponse: v,
	}
}

// LengthBasedPasswordValidatorResponseAsAddPasswordValidator200Response is a convenience function that returns LengthBasedPasswordValidatorResponse wrapped in AddPasswordValidator200Response
func LengthBasedPasswordValidatorResponseAsAddPasswordValidator200Response(v *LengthBasedPasswordValidatorResponse) AddPasswordValidator200Response {
	return AddPasswordValidator200Response{
		LengthBasedPasswordValidatorResponse: v,
	}
}

// PwnedPasswordsPasswordValidatorResponseAsAddPasswordValidator200Response is a convenience function that returns PwnedPasswordsPasswordValidatorResponse wrapped in AddPasswordValidator200Response
func PwnedPasswordsPasswordValidatorResponseAsAddPasswordValidator200Response(v *PwnedPasswordsPasswordValidatorResponse) AddPasswordValidator200Response {
	return AddPasswordValidator200Response{
		PwnedPasswordsPasswordValidatorResponse: v,
	}
}

// RegularExpressionPasswordValidatorResponseAsAddPasswordValidator200Response is a convenience function that returns RegularExpressionPasswordValidatorResponse wrapped in AddPasswordValidator200Response
func RegularExpressionPasswordValidatorResponseAsAddPasswordValidator200Response(v *RegularExpressionPasswordValidatorResponse) AddPasswordValidator200Response {
	return AddPasswordValidator200Response{
		RegularExpressionPasswordValidatorResponse: v,
	}
}

// RepeatedCharactersPasswordValidatorResponseAsAddPasswordValidator200Response is a convenience function that returns RepeatedCharactersPasswordValidatorResponse wrapped in AddPasswordValidator200Response
func RepeatedCharactersPasswordValidatorResponseAsAddPasswordValidator200Response(v *RepeatedCharactersPasswordValidatorResponse) AddPasswordValidator200Response {
	return AddPasswordValidator200Response{
		RepeatedCharactersPasswordValidatorResponse: v,
	}
}

// SimilarityBasedPasswordValidatorResponseAsAddPasswordValidator200Response is a convenience function that returns SimilarityBasedPasswordValidatorResponse wrapped in AddPasswordValidator200Response
func SimilarityBasedPasswordValidatorResponseAsAddPasswordValidator200Response(v *SimilarityBasedPasswordValidatorResponse) AddPasswordValidator200Response {
	return AddPasswordValidator200Response{
		SimilarityBasedPasswordValidatorResponse: v,
	}
}

// ThirdPartyPasswordValidatorResponseAsAddPasswordValidator200Response is a convenience function that returns ThirdPartyPasswordValidatorResponse wrapped in AddPasswordValidator200Response
func ThirdPartyPasswordValidatorResponseAsAddPasswordValidator200Response(v *ThirdPartyPasswordValidatorResponse) AddPasswordValidator200Response {
	return AddPasswordValidator200Response{
		ThirdPartyPasswordValidatorResponse: v,
	}
}

// UniqueCharactersPasswordValidatorResponseAsAddPasswordValidator200Response is a convenience function that returns UniqueCharactersPasswordValidatorResponse wrapped in AddPasswordValidator200Response
func UniqueCharactersPasswordValidatorResponseAsAddPasswordValidator200Response(v *UniqueCharactersPasswordValidatorResponse) AddPasswordValidator200Response {
	return AddPasswordValidator200Response{
		UniqueCharactersPasswordValidatorResponse: v,
	}
}

// Utf8PasswordValidatorResponseAsAddPasswordValidator200Response is a convenience function that returns Utf8PasswordValidatorResponse wrapped in AddPasswordValidator200Response
func Utf8PasswordValidatorResponseAsAddPasswordValidator200Response(v *Utf8PasswordValidatorResponse) AddPasswordValidator200Response {
	return AddPasswordValidator200Response{
		Utf8PasswordValidatorResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddPasswordValidator200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AttributeValuePasswordValidatorResponse
	err = newStrictDecoder(data).Decode(&dst.AttributeValuePasswordValidatorResponse)
	if err == nil {
		jsonAttributeValuePasswordValidatorResponse, _ := json.Marshal(dst.AttributeValuePasswordValidatorResponse)
		if string(jsonAttributeValuePasswordValidatorResponse) == "{}" { // empty struct
			dst.AttributeValuePasswordValidatorResponse = nil
		} else {
			match++
		}
	} else {
		dst.AttributeValuePasswordValidatorResponse = nil
	}

	// try to unmarshal data into CharacterSetPasswordValidatorResponse
	err = newStrictDecoder(data).Decode(&dst.CharacterSetPasswordValidatorResponse)
	if err == nil {
		jsonCharacterSetPasswordValidatorResponse, _ := json.Marshal(dst.CharacterSetPasswordValidatorResponse)
		if string(jsonCharacterSetPasswordValidatorResponse) == "{}" { // empty struct
			dst.CharacterSetPasswordValidatorResponse = nil
		} else {
			match++
		}
	} else {
		dst.CharacterSetPasswordValidatorResponse = nil
	}

	// try to unmarshal data into DictionaryPasswordValidatorResponse
	err = newStrictDecoder(data).Decode(&dst.DictionaryPasswordValidatorResponse)
	if err == nil {
		jsonDictionaryPasswordValidatorResponse, _ := json.Marshal(dst.DictionaryPasswordValidatorResponse)
		if string(jsonDictionaryPasswordValidatorResponse) == "{}" { // empty struct
			dst.DictionaryPasswordValidatorResponse = nil
		} else {
			match++
		}
	} else {
		dst.DictionaryPasswordValidatorResponse = nil
	}

	// try to unmarshal data into DisallowedCharactersPasswordValidatorResponse
	err = newStrictDecoder(data).Decode(&dst.DisallowedCharactersPasswordValidatorResponse)
	if err == nil {
		jsonDisallowedCharactersPasswordValidatorResponse, _ := json.Marshal(dst.DisallowedCharactersPasswordValidatorResponse)
		if string(jsonDisallowedCharactersPasswordValidatorResponse) == "{}" { // empty struct
			dst.DisallowedCharactersPasswordValidatorResponse = nil
		} else {
			match++
		}
	} else {
		dst.DisallowedCharactersPasswordValidatorResponse = nil
	}

	// try to unmarshal data into GroovyScriptedPasswordValidatorResponse
	err = newStrictDecoder(data).Decode(&dst.GroovyScriptedPasswordValidatorResponse)
	if err == nil {
		jsonGroovyScriptedPasswordValidatorResponse, _ := json.Marshal(dst.GroovyScriptedPasswordValidatorResponse)
		if string(jsonGroovyScriptedPasswordValidatorResponse) == "{}" { // empty struct
			dst.GroovyScriptedPasswordValidatorResponse = nil
		} else {
			match++
		}
	} else {
		dst.GroovyScriptedPasswordValidatorResponse = nil
	}

	// try to unmarshal data into HaystackPasswordValidatorResponse
	err = newStrictDecoder(data).Decode(&dst.HaystackPasswordValidatorResponse)
	if err == nil {
		jsonHaystackPasswordValidatorResponse, _ := json.Marshal(dst.HaystackPasswordValidatorResponse)
		if string(jsonHaystackPasswordValidatorResponse) == "{}" { // empty struct
			dst.HaystackPasswordValidatorResponse = nil
		} else {
			match++
		}
	} else {
		dst.HaystackPasswordValidatorResponse = nil
	}

	// try to unmarshal data into LengthBasedPasswordValidatorResponse
	err = newStrictDecoder(data).Decode(&dst.LengthBasedPasswordValidatorResponse)
	if err == nil {
		jsonLengthBasedPasswordValidatorResponse, _ := json.Marshal(dst.LengthBasedPasswordValidatorResponse)
		if string(jsonLengthBasedPasswordValidatorResponse) == "{}" { // empty struct
			dst.LengthBasedPasswordValidatorResponse = nil
		} else {
			match++
		}
	} else {
		dst.LengthBasedPasswordValidatorResponse = nil
	}

	// try to unmarshal data into PwnedPasswordsPasswordValidatorResponse
	err = newStrictDecoder(data).Decode(&dst.PwnedPasswordsPasswordValidatorResponse)
	if err == nil {
		jsonPwnedPasswordsPasswordValidatorResponse, _ := json.Marshal(dst.PwnedPasswordsPasswordValidatorResponse)
		if string(jsonPwnedPasswordsPasswordValidatorResponse) == "{}" { // empty struct
			dst.PwnedPasswordsPasswordValidatorResponse = nil
		} else {
			match++
		}
	} else {
		dst.PwnedPasswordsPasswordValidatorResponse = nil
	}

	// try to unmarshal data into RegularExpressionPasswordValidatorResponse
	err = newStrictDecoder(data).Decode(&dst.RegularExpressionPasswordValidatorResponse)
	if err == nil {
		jsonRegularExpressionPasswordValidatorResponse, _ := json.Marshal(dst.RegularExpressionPasswordValidatorResponse)
		if string(jsonRegularExpressionPasswordValidatorResponse) == "{}" { // empty struct
			dst.RegularExpressionPasswordValidatorResponse = nil
		} else {
			match++
		}
	} else {
		dst.RegularExpressionPasswordValidatorResponse = nil
	}

	// try to unmarshal data into RepeatedCharactersPasswordValidatorResponse
	err = newStrictDecoder(data).Decode(&dst.RepeatedCharactersPasswordValidatorResponse)
	if err == nil {
		jsonRepeatedCharactersPasswordValidatorResponse, _ := json.Marshal(dst.RepeatedCharactersPasswordValidatorResponse)
		if string(jsonRepeatedCharactersPasswordValidatorResponse) == "{}" { // empty struct
			dst.RepeatedCharactersPasswordValidatorResponse = nil
		} else {
			match++
		}
	} else {
		dst.RepeatedCharactersPasswordValidatorResponse = nil
	}

	// try to unmarshal data into SimilarityBasedPasswordValidatorResponse
	err = newStrictDecoder(data).Decode(&dst.SimilarityBasedPasswordValidatorResponse)
	if err == nil {
		jsonSimilarityBasedPasswordValidatorResponse, _ := json.Marshal(dst.SimilarityBasedPasswordValidatorResponse)
		if string(jsonSimilarityBasedPasswordValidatorResponse) == "{}" { // empty struct
			dst.SimilarityBasedPasswordValidatorResponse = nil
		} else {
			match++
		}
	} else {
		dst.SimilarityBasedPasswordValidatorResponse = nil
	}

	// try to unmarshal data into ThirdPartyPasswordValidatorResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartyPasswordValidatorResponse)
	if err == nil {
		jsonThirdPartyPasswordValidatorResponse, _ := json.Marshal(dst.ThirdPartyPasswordValidatorResponse)
		if string(jsonThirdPartyPasswordValidatorResponse) == "{}" { // empty struct
			dst.ThirdPartyPasswordValidatorResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartyPasswordValidatorResponse = nil
	}

	// try to unmarshal data into UniqueCharactersPasswordValidatorResponse
	err = newStrictDecoder(data).Decode(&dst.UniqueCharactersPasswordValidatorResponse)
	if err == nil {
		jsonUniqueCharactersPasswordValidatorResponse, _ := json.Marshal(dst.UniqueCharactersPasswordValidatorResponse)
		if string(jsonUniqueCharactersPasswordValidatorResponse) == "{}" { // empty struct
			dst.UniqueCharactersPasswordValidatorResponse = nil
		} else {
			match++
		}
	} else {
		dst.UniqueCharactersPasswordValidatorResponse = nil
	}

	// try to unmarshal data into Utf8PasswordValidatorResponse
	err = newStrictDecoder(data).Decode(&dst.Utf8PasswordValidatorResponse)
	if err == nil {
		jsonUtf8PasswordValidatorResponse, _ := json.Marshal(dst.Utf8PasswordValidatorResponse)
		if string(jsonUtf8PasswordValidatorResponse) == "{}" { // empty struct
			dst.Utf8PasswordValidatorResponse = nil
		} else {
			match++
		}
	} else {
		dst.Utf8PasswordValidatorResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AttributeValuePasswordValidatorResponse = nil
		dst.CharacterSetPasswordValidatorResponse = nil
		dst.DictionaryPasswordValidatorResponse = nil
		dst.DisallowedCharactersPasswordValidatorResponse = nil
		dst.GroovyScriptedPasswordValidatorResponse = nil
		dst.HaystackPasswordValidatorResponse = nil
		dst.LengthBasedPasswordValidatorResponse = nil
		dst.PwnedPasswordsPasswordValidatorResponse = nil
		dst.RegularExpressionPasswordValidatorResponse = nil
		dst.RepeatedCharactersPasswordValidatorResponse = nil
		dst.SimilarityBasedPasswordValidatorResponse = nil
		dst.ThirdPartyPasswordValidatorResponse = nil
		dst.UniqueCharactersPasswordValidatorResponse = nil
		dst.Utf8PasswordValidatorResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddPasswordValidator200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddPasswordValidator200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddPasswordValidator200Response) MarshalJSON() ([]byte, error) {
	if src.AttributeValuePasswordValidatorResponse != nil {
		return json.Marshal(&src.AttributeValuePasswordValidatorResponse)
	}

	if src.CharacterSetPasswordValidatorResponse != nil {
		return json.Marshal(&src.CharacterSetPasswordValidatorResponse)
	}

	if src.DictionaryPasswordValidatorResponse != nil {
		return json.Marshal(&src.DictionaryPasswordValidatorResponse)
	}

	if src.DisallowedCharactersPasswordValidatorResponse != nil {
		return json.Marshal(&src.DisallowedCharactersPasswordValidatorResponse)
	}

	if src.GroovyScriptedPasswordValidatorResponse != nil {
		return json.Marshal(&src.GroovyScriptedPasswordValidatorResponse)
	}

	if src.HaystackPasswordValidatorResponse != nil {
		return json.Marshal(&src.HaystackPasswordValidatorResponse)
	}

	if src.LengthBasedPasswordValidatorResponse != nil {
		return json.Marshal(&src.LengthBasedPasswordValidatorResponse)
	}

	if src.PwnedPasswordsPasswordValidatorResponse != nil {
		return json.Marshal(&src.PwnedPasswordsPasswordValidatorResponse)
	}

	if src.RegularExpressionPasswordValidatorResponse != nil {
		return json.Marshal(&src.RegularExpressionPasswordValidatorResponse)
	}

	if src.RepeatedCharactersPasswordValidatorResponse != nil {
		return json.Marshal(&src.RepeatedCharactersPasswordValidatorResponse)
	}

	if src.SimilarityBasedPasswordValidatorResponse != nil {
		return json.Marshal(&src.SimilarityBasedPasswordValidatorResponse)
	}

	if src.ThirdPartyPasswordValidatorResponse != nil {
		return json.Marshal(&src.ThirdPartyPasswordValidatorResponse)
	}

	if src.UniqueCharactersPasswordValidatorResponse != nil {
		return json.Marshal(&src.UniqueCharactersPasswordValidatorResponse)
	}

	if src.Utf8PasswordValidatorResponse != nil {
		return json.Marshal(&src.Utf8PasswordValidatorResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddPasswordValidator200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AttributeValuePasswordValidatorResponse != nil {
		return obj.AttributeValuePasswordValidatorResponse
	}

	if obj.CharacterSetPasswordValidatorResponse != nil {
		return obj.CharacterSetPasswordValidatorResponse
	}

	if obj.DictionaryPasswordValidatorResponse != nil {
		return obj.DictionaryPasswordValidatorResponse
	}

	if obj.DisallowedCharactersPasswordValidatorResponse != nil {
		return obj.DisallowedCharactersPasswordValidatorResponse
	}

	if obj.GroovyScriptedPasswordValidatorResponse != nil {
		return obj.GroovyScriptedPasswordValidatorResponse
	}

	if obj.HaystackPasswordValidatorResponse != nil {
		return obj.HaystackPasswordValidatorResponse
	}

	if obj.LengthBasedPasswordValidatorResponse != nil {
		return obj.LengthBasedPasswordValidatorResponse
	}

	if obj.PwnedPasswordsPasswordValidatorResponse != nil {
		return obj.PwnedPasswordsPasswordValidatorResponse
	}

	if obj.RegularExpressionPasswordValidatorResponse != nil {
		return obj.RegularExpressionPasswordValidatorResponse
	}

	if obj.RepeatedCharactersPasswordValidatorResponse != nil {
		return obj.RepeatedCharactersPasswordValidatorResponse
	}

	if obj.SimilarityBasedPasswordValidatorResponse != nil {
		return obj.SimilarityBasedPasswordValidatorResponse
	}

	if obj.ThirdPartyPasswordValidatorResponse != nil {
		return obj.ThirdPartyPasswordValidatorResponse
	}

	if obj.UniqueCharactersPasswordValidatorResponse != nil {
		return obj.UniqueCharactersPasswordValidatorResponse
	}

	if obj.Utf8PasswordValidatorResponse != nil {
		return obj.Utf8PasswordValidatorResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddPasswordValidator200Response struct {
	value *AddPasswordValidator200Response
	isSet bool
}

func (v NullableAddPasswordValidator200Response) Get() *AddPasswordValidator200Response {
	return v.value
}

func (v *NullableAddPasswordValidator200Response) Set(val *AddPasswordValidator200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPasswordValidator200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPasswordValidator200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPasswordValidator200Response(val *AddPasswordValidator200Response) *NullableAddPasswordValidator200Response {
	return &NullableAddPasswordValidator200Response{value: val, isSet: true}
}

func (v NullableAddPasswordValidator200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPasswordValidator200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
