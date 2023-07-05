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

// AddPlugin200Response - struct for AddPlugin200Response
type AddPlugin200Response struct {
	AttributeMapperPluginResponse                                  *AttributeMapperPluginResponse
	CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse
	CleanUpExpiredPingfederatePersistentSessionsPluginResponse     *CleanUpExpiredPingfederatePersistentSessionsPluginResponse
	CleanUpInactivePingfederatePersistentSessionsPluginResponse    *CleanUpInactivePingfederatePersistentSessionsPluginResponse
	CoalesceModificationsPluginResponse                            *CoalesceModificationsPluginResponse
	ComposedAttributePluginResponse                                *ComposedAttributePluginResponse
	DelayPluginResponse                                            *DelayPluginResponse
	DnMapperPluginResponse                                         *DnMapperPluginResponse
	GroovyScriptedPluginResponse                                   *GroovyScriptedPluginResponse
	InternalSearchRatePluginResponse                               *InternalSearchRatePluginResponse
	ModifiablePasswordPolicyStatePluginResponse                    *ModifiablePasswordPolicyStatePluginResponse
	PassThroughAuthenticationPluginResponse                        *PassThroughAuthenticationPluginResponse
	PeriodicGcPluginResponse                                       *PeriodicGcPluginResponse
	PeriodicStatsLoggerPluginResponse                              *PeriodicStatsLoggerPluginResponse
	PingOnePassThroughAuthenticationPluginResponse                 *PingOnePassThroughAuthenticationPluginResponse
	PluggablePassThroughAuthenticationPluginResponse               *PluggablePassThroughAuthenticationPluginResponse
	PurgeExpiredDataPluginResponse                                 *PurgeExpiredDataPluginResponse
	ReferentialIntegrityPluginResponse                             *ReferentialIntegrityPluginResponse
	ReferralOnUpdatePluginResponse                                 *ReferralOnUpdatePluginResponse
	SearchShutdownPluginResponse                                   *SearchShutdownPluginResponse
	SevenBitCleanPluginResponse                                    *SevenBitCleanPluginResponse
	SimpleToExternalBindPluginResponse                             *SimpleToExternalBindPluginResponse
	SnmpSubagentPluginResponse                                     *SnmpSubagentPluginResponse
	SubOperationTimingPluginResponse                               *SubOperationTimingPluginResponse
	ThirdPartyPluginResponse                                       *ThirdPartyPluginResponse
	UniqueAttributePluginResponse                                  *UniqueAttributePluginResponse
}

// AttributeMapperPluginResponseAsAddPlugin200Response is a convenience function that returns AttributeMapperPluginResponse wrapped in AddPlugin200Response
func AttributeMapperPluginResponseAsAddPlugin200Response(v *AttributeMapperPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		AttributeMapperPluginResponse: v,
	}
}

// CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponseAsAddPlugin200Response is a convenience function that returns CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse wrapped in AddPlugin200Response
func CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponseAsAddPlugin200Response(v *CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse: v,
	}
}

// CleanUpExpiredPingfederatePersistentSessionsPluginResponseAsAddPlugin200Response is a convenience function that returns CleanUpExpiredPingfederatePersistentSessionsPluginResponse wrapped in AddPlugin200Response
func CleanUpExpiredPingfederatePersistentSessionsPluginResponseAsAddPlugin200Response(v *CleanUpExpiredPingfederatePersistentSessionsPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		CleanUpExpiredPingfederatePersistentSessionsPluginResponse: v,
	}
}

// CleanUpInactivePingfederatePersistentSessionsPluginResponseAsAddPlugin200Response is a convenience function that returns CleanUpInactivePingfederatePersistentSessionsPluginResponse wrapped in AddPlugin200Response
func CleanUpInactivePingfederatePersistentSessionsPluginResponseAsAddPlugin200Response(v *CleanUpInactivePingfederatePersistentSessionsPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		CleanUpInactivePingfederatePersistentSessionsPluginResponse: v,
	}
}

// CoalesceModificationsPluginResponseAsAddPlugin200Response is a convenience function that returns CoalesceModificationsPluginResponse wrapped in AddPlugin200Response
func CoalesceModificationsPluginResponseAsAddPlugin200Response(v *CoalesceModificationsPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		CoalesceModificationsPluginResponse: v,
	}
}

// ComposedAttributePluginResponseAsAddPlugin200Response is a convenience function that returns ComposedAttributePluginResponse wrapped in AddPlugin200Response
func ComposedAttributePluginResponseAsAddPlugin200Response(v *ComposedAttributePluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		ComposedAttributePluginResponse: v,
	}
}

// DelayPluginResponseAsAddPlugin200Response is a convenience function that returns DelayPluginResponse wrapped in AddPlugin200Response
func DelayPluginResponseAsAddPlugin200Response(v *DelayPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		DelayPluginResponse: v,
	}
}

// DnMapperPluginResponseAsAddPlugin200Response is a convenience function that returns DnMapperPluginResponse wrapped in AddPlugin200Response
func DnMapperPluginResponseAsAddPlugin200Response(v *DnMapperPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		DnMapperPluginResponse: v,
	}
}

// GroovyScriptedPluginResponseAsAddPlugin200Response is a convenience function that returns GroovyScriptedPluginResponse wrapped in AddPlugin200Response
func GroovyScriptedPluginResponseAsAddPlugin200Response(v *GroovyScriptedPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		GroovyScriptedPluginResponse: v,
	}
}

// InternalSearchRatePluginResponseAsAddPlugin200Response is a convenience function that returns InternalSearchRatePluginResponse wrapped in AddPlugin200Response
func InternalSearchRatePluginResponseAsAddPlugin200Response(v *InternalSearchRatePluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		InternalSearchRatePluginResponse: v,
	}
}

// ModifiablePasswordPolicyStatePluginResponseAsAddPlugin200Response is a convenience function that returns ModifiablePasswordPolicyStatePluginResponse wrapped in AddPlugin200Response
func ModifiablePasswordPolicyStatePluginResponseAsAddPlugin200Response(v *ModifiablePasswordPolicyStatePluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		ModifiablePasswordPolicyStatePluginResponse: v,
	}
}

// PassThroughAuthenticationPluginResponseAsAddPlugin200Response is a convenience function that returns PassThroughAuthenticationPluginResponse wrapped in AddPlugin200Response
func PassThroughAuthenticationPluginResponseAsAddPlugin200Response(v *PassThroughAuthenticationPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		PassThroughAuthenticationPluginResponse: v,
	}
}

// PeriodicGcPluginResponseAsAddPlugin200Response is a convenience function that returns PeriodicGcPluginResponse wrapped in AddPlugin200Response
func PeriodicGcPluginResponseAsAddPlugin200Response(v *PeriodicGcPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		PeriodicGcPluginResponse: v,
	}
}

// PeriodicStatsLoggerPluginResponseAsAddPlugin200Response is a convenience function that returns PeriodicStatsLoggerPluginResponse wrapped in AddPlugin200Response
func PeriodicStatsLoggerPluginResponseAsAddPlugin200Response(v *PeriodicStatsLoggerPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		PeriodicStatsLoggerPluginResponse: v,
	}
}

// PingOnePassThroughAuthenticationPluginResponseAsAddPlugin200Response is a convenience function that returns PingOnePassThroughAuthenticationPluginResponse wrapped in AddPlugin200Response
func PingOnePassThroughAuthenticationPluginResponseAsAddPlugin200Response(v *PingOnePassThroughAuthenticationPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		PingOnePassThroughAuthenticationPluginResponse: v,
	}
}

// PluggablePassThroughAuthenticationPluginResponseAsAddPlugin200Response is a convenience function that returns PluggablePassThroughAuthenticationPluginResponse wrapped in AddPlugin200Response
func PluggablePassThroughAuthenticationPluginResponseAsAddPlugin200Response(v *PluggablePassThroughAuthenticationPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		PluggablePassThroughAuthenticationPluginResponse: v,
	}
}

// PurgeExpiredDataPluginResponseAsAddPlugin200Response is a convenience function that returns PurgeExpiredDataPluginResponse wrapped in AddPlugin200Response
func PurgeExpiredDataPluginResponseAsAddPlugin200Response(v *PurgeExpiredDataPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		PurgeExpiredDataPluginResponse: v,
	}
}

// ReferentialIntegrityPluginResponseAsAddPlugin200Response is a convenience function that returns ReferentialIntegrityPluginResponse wrapped in AddPlugin200Response
func ReferentialIntegrityPluginResponseAsAddPlugin200Response(v *ReferentialIntegrityPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		ReferentialIntegrityPluginResponse: v,
	}
}

// ReferralOnUpdatePluginResponseAsAddPlugin200Response is a convenience function that returns ReferralOnUpdatePluginResponse wrapped in AddPlugin200Response
func ReferralOnUpdatePluginResponseAsAddPlugin200Response(v *ReferralOnUpdatePluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		ReferralOnUpdatePluginResponse: v,
	}
}

// SearchShutdownPluginResponseAsAddPlugin200Response is a convenience function that returns SearchShutdownPluginResponse wrapped in AddPlugin200Response
func SearchShutdownPluginResponseAsAddPlugin200Response(v *SearchShutdownPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		SearchShutdownPluginResponse: v,
	}
}

// SevenBitCleanPluginResponseAsAddPlugin200Response is a convenience function that returns SevenBitCleanPluginResponse wrapped in AddPlugin200Response
func SevenBitCleanPluginResponseAsAddPlugin200Response(v *SevenBitCleanPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		SevenBitCleanPluginResponse: v,
	}
}

// SimpleToExternalBindPluginResponseAsAddPlugin200Response is a convenience function that returns SimpleToExternalBindPluginResponse wrapped in AddPlugin200Response
func SimpleToExternalBindPluginResponseAsAddPlugin200Response(v *SimpleToExternalBindPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		SimpleToExternalBindPluginResponse: v,
	}
}

// SnmpSubagentPluginResponseAsAddPlugin200Response is a convenience function that returns SnmpSubagentPluginResponse wrapped in AddPlugin200Response
func SnmpSubagentPluginResponseAsAddPlugin200Response(v *SnmpSubagentPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		SnmpSubagentPluginResponse: v,
	}
}

// SubOperationTimingPluginResponseAsAddPlugin200Response is a convenience function that returns SubOperationTimingPluginResponse wrapped in AddPlugin200Response
func SubOperationTimingPluginResponseAsAddPlugin200Response(v *SubOperationTimingPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		SubOperationTimingPluginResponse: v,
	}
}

// ThirdPartyPluginResponseAsAddPlugin200Response is a convenience function that returns ThirdPartyPluginResponse wrapped in AddPlugin200Response
func ThirdPartyPluginResponseAsAddPlugin200Response(v *ThirdPartyPluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		ThirdPartyPluginResponse: v,
	}
}

// UniqueAttributePluginResponseAsAddPlugin200Response is a convenience function that returns UniqueAttributePluginResponse wrapped in AddPlugin200Response
func UniqueAttributePluginResponseAsAddPlugin200Response(v *UniqueAttributePluginResponse) AddPlugin200Response {
	return AddPlugin200Response{
		UniqueAttributePluginResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddPlugin200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AttributeMapperPluginResponse
	err = newStrictDecoder(data).Decode(&dst.AttributeMapperPluginResponse)
	if err == nil {
		jsonAttributeMapperPluginResponse, _ := json.Marshal(dst.AttributeMapperPluginResponse)
		if string(jsonAttributeMapperPluginResponse) == "{}" { // empty struct
			dst.AttributeMapperPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.AttributeMapperPluginResponse = nil
	}

	// try to unmarshal data into CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse
	err = newStrictDecoder(data).Decode(&dst.CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse)
	if err == nil {
		jsonCleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse, _ := json.Marshal(dst.CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse)
		if string(jsonCleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse) == "{}" { // empty struct
			dst.CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse = nil
	}

	// try to unmarshal data into CleanUpExpiredPingfederatePersistentSessionsPluginResponse
	err = newStrictDecoder(data).Decode(&dst.CleanUpExpiredPingfederatePersistentSessionsPluginResponse)
	if err == nil {
		jsonCleanUpExpiredPingfederatePersistentSessionsPluginResponse, _ := json.Marshal(dst.CleanUpExpiredPingfederatePersistentSessionsPluginResponse)
		if string(jsonCleanUpExpiredPingfederatePersistentSessionsPluginResponse) == "{}" { // empty struct
			dst.CleanUpExpiredPingfederatePersistentSessionsPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.CleanUpExpiredPingfederatePersistentSessionsPluginResponse = nil
	}

	// try to unmarshal data into CleanUpInactivePingfederatePersistentSessionsPluginResponse
	err = newStrictDecoder(data).Decode(&dst.CleanUpInactivePingfederatePersistentSessionsPluginResponse)
	if err == nil {
		jsonCleanUpInactivePingfederatePersistentSessionsPluginResponse, _ := json.Marshal(dst.CleanUpInactivePingfederatePersistentSessionsPluginResponse)
		if string(jsonCleanUpInactivePingfederatePersistentSessionsPluginResponse) == "{}" { // empty struct
			dst.CleanUpInactivePingfederatePersistentSessionsPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.CleanUpInactivePingfederatePersistentSessionsPluginResponse = nil
	}

	// try to unmarshal data into CoalesceModificationsPluginResponse
	err = newStrictDecoder(data).Decode(&dst.CoalesceModificationsPluginResponse)
	if err == nil {
		jsonCoalesceModificationsPluginResponse, _ := json.Marshal(dst.CoalesceModificationsPluginResponse)
		if string(jsonCoalesceModificationsPluginResponse) == "{}" { // empty struct
			dst.CoalesceModificationsPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.CoalesceModificationsPluginResponse = nil
	}

	// try to unmarshal data into ComposedAttributePluginResponse
	err = newStrictDecoder(data).Decode(&dst.ComposedAttributePluginResponse)
	if err == nil {
		jsonComposedAttributePluginResponse, _ := json.Marshal(dst.ComposedAttributePluginResponse)
		if string(jsonComposedAttributePluginResponse) == "{}" { // empty struct
			dst.ComposedAttributePluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.ComposedAttributePluginResponse = nil
	}

	// try to unmarshal data into DelayPluginResponse
	err = newStrictDecoder(data).Decode(&dst.DelayPluginResponse)
	if err == nil {
		jsonDelayPluginResponse, _ := json.Marshal(dst.DelayPluginResponse)
		if string(jsonDelayPluginResponse) == "{}" { // empty struct
			dst.DelayPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.DelayPluginResponse = nil
	}

	// try to unmarshal data into DnMapperPluginResponse
	err = newStrictDecoder(data).Decode(&dst.DnMapperPluginResponse)
	if err == nil {
		jsonDnMapperPluginResponse, _ := json.Marshal(dst.DnMapperPluginResponse)
		if string(jsonDnMapperPluginResponse) == "{}" { // empty struct
			dst.DnMapperPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.DnMapperPluginResponse = nil
	}

	// try to unmarshal data into GroovyScriptedPluginResponse
	err = newStrictDecoder(data).Decode(&dst.GroovyScriptedPluginResponse)
	if err == nil {
		jsonGroovyScriptedPluginResponse, _ := json.Marshal(dst.GroovyScriptedPluginResponse)
		if string(jsonGroovyScriptedPluginResponse) == "{}" { // empty struct
			dst.GroovyScriptedPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.GroovyScriptedPluginResponse = nil
	}

	// try to unmarshal data into InternalSearchRatePluginResponse
	err = newStrictDecoder(data).Decode(&dst.InternalSearchRatePluginResponse)
	if err == nil {
		jsonInternalSearchRatePluginResponse, _ := json.Marshal(dst.InternalSearchRatePluginResponse)
		if string(jsonInternalSearchRatePluginResponse) == "{}" { // empty struct
			dst.InternalSearchRatePluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.InternalSearchRatePluginResponse = nil
	}

	// try to unmarshal data into ModifiablePasswordPolicyStatePluginResponse
	err = newStrictDecoder(data).Decode(&dst.ModifiablePasswordPolicyStatePluginResponse)
	if err == nil {
		jsonModifiablePasswordPolicyStatePluginResponse, _ := json.Marshal(dst.ModifiablePasswordPolicyStatePluginResponse)
		if string(jsonModifiablePasswordPolicyStatePluginResponse) == "{}" { // empty struct
			dst.ModifiablePasswordPolicyStatePluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.ModifiablePasswordPolicyStatePluginResponse = nil
	}

	// try to unmarshal data into PassThroughAuthenticationPluginResponse
	err = newStrictDecoder(data).Decode(&dst.PassThroughAuthenticationPluginResponse)
	if err == nil {
		jsonPassThroughAuthenticationPluginResponse, _ := json.Marshal(dst.PassThroughAuthenticationPluginResponse)
		if string(jsonPassThroughAuthenticationPluginResponse) == "{}" { // empty struct
			dst.PassThroughAuthenticationPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.PassThroughAuthenticationPluginResponse = nil
	}

	// try to unmarshal data into PeriodicGcPluginResponse
	err = newStrictDecoder(data).Decode(&dst.PeriodicGcPluginResponse)
	if err == nil {
		jsonPeriodicGcPluginResponse, _ := json.Marshal(dst.PeriodicGcPluginResponse)
		if string(jsonPeriodicGcPluginResponse) == "{}" { // empty struct
			dst.PeriodicGcPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.PeriodicGcPluginResponse = nil
	}

	// try to unmarshal data into PeriodicStatsLoggerPluginResponse
	err = newStrictDecoder(data).Decode(&dst.PeriodicStatsLoggerPluginResponse)
	if err == nil {
		jsonPeriodicStatsLoggerPluginResponse, _ := json.Marshal(dst.PeriodicStatsLoggerPluginResponse)
		if string(jsonPeriodicStatsLoggerPluginResponse) == "{}" { // empty struct
			dst.PeriodicStatsLoggerPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.PeriodicStatsLoggerPluginResponse = nil
	}

	// try to unmarshal data into PingOnePassThroughAuthenticationPluginResponse
	err = newStrictDecoder(data).Decode(&dst.PingOnePassThroughAuthenticationPluginResponse)
	if err == nil {
		jsonPingOnePassThroughAuthenticationPluginResponse, _ := json.Marshal(dst.PingOnePassThroughAuthenticationPluginResponse)
		if string(jsonPingOnePassThroughAuthenticationPluginResponse) == "{}" { // empty struct
			dst.PingOnePassThroughAuthenticationPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.PingOnePassThroughAuthenticationPluginResponse = nil
	}

	// try to unmarshal data into PluggablePassThroughAuthenticationPluginResponse
	err = newStrictDecoder(data).Decode(&dst.PluggablePassThroughAuthenticationPluginResponse)
	if err == nil {
		jsonPluggablePassThroughAuthenticationPluginResponse, _ := json.Marshal(dst.PluggablePassThroughAuthenticationPluginResponse)
		if string(jsonPluggablePassThroughAuthenticationPluginResponse) == "{}" { // empty struct
			dst.PluggablePassThroughAuthenticationPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.PluggablePassThroughAuthenticationPluginResponse = nil
	}

	// try to unmarshal data into PurgeExpiredDataPluginResponse
	err = newStrictDecoder(data).Decode(&dst.PurgeExpiredDataPluginResponse)
	if err == nil {
		jsonPurgeExpiredDataPluginResponse, _ := json.Marshal(dst.PurgeExpiredDataPluginResponse)
		if string(jsonPurgeExpiredDataPluginResponse) == "{}" { // empty struct
			dst.PurgeExpiredDataPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.PurgeExpiredDataPluginResponse = nil
	}

	// try to unmarshal data into ReferentialIntegrityPluginResponse
	err = newStrictDecoder(data).Decode(&dst.ReferentialIntegrityPluginResponse)
	if err == nil {
		jsonReferentialIntegrityPluginResponse, _ := json.Marshal(dst.ReferentialIntegrityPluginResponse)
		if string(jsonReferentialIntegrityPluginResponse) == "{}" { // empty struct
			dst.ReferentialIntegrityPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.ReferentialIntegrityPluginResponse = nil
	}

	// try to unmarshal data into ReferralOnUpdatePluginResponse
	err = newStrictDecoder(data).Decode(&dst.ReferralOnUpdatePluginResponse)
	if err == nil {
		jsonReferralOnUpdatePluginResponse, _ := json.Marshal(dst.ReferralOnUpdatePluginResponse)
		if string(jsonReferralOnUpdatePluginResponse) == "{}" { // empty struct
			dst.ReferralOnUpdatePluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.ReferralOnUpdatePluginResponse = nil
	}

	// try to unmarshal data into SearchShutdownPluginResponse
	err = newStrictDecoder(data).Decode(&dst.SearchShutdownPluginResponse)
	if err == nil {
		jsonSearchShutdownPluginResponse, _ := json.Marshal(dst.SearchShutdownPluginResponse)
		if string(jsonSearchShutdownPluginResponse) == "{}" { // empty struct
			dst.SearchShutdownPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.SearchShutdownPluginResponse = nil
	}

	// try to unmarshal data into SevenBitCleanPluginResponse
	err = newStrictDecoder(data).Decode(&dst.SevenBitCleanPluginResponse)
	if err == nil {
		jsonSevenBitCleanPluginResponse, _ := json.Marshal(dst.SevenBitCleanPluginResponse)
		if string(jsonSevenBitCleanPluginResponse) == "{}" { // empty struct
			dst.SevenBitCleanPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.SevenBitCleanPluginResponse = nil
	}

	// try to unmarshal data into SimpleToExternalBindPluginResponse
	err = newStrictDecoder(data).Decode(&dst.SimpleToExternalBindPluginResponse)
	if err == nil {
		jsonSimpleToExternalBindPluginResponse, _ := json.Marshal(dst.SimpleToExternalBindPluginResponse)
		if string(jsonSimpleToExternalBindPluginResponse) == "{}" { // empty struct
			dst.SimpleToExternalBindPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.SimpleToExternalBindPluginResponse = nil
	}

	// try to unmarshal data into SnmpSubagentPluginResponse
	err = newStrictDecoder(data).Decode(&dst.SnmpSubagentPluginResponse)
	if err == nil {
		jsonSnmpSubagentPluginResponse, _ := json.Marshal(dst.SnmpSubagentPluginResponse)
		if string(jsonSnmpSubagentPluginResponse) == "{}" { // empty struct
			dst.SnmpSubagentPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.SnmpSubagentPluginResponse = nil
	}

	// try to unmarshal data into SubOperationTimingPluginResponse
	err = newStrictDecoder(data).Decode(&dst.SubOperationTimingPluginResponse)
	if err == nil {
		jsonSubOperationTimingPluginResponse, _ := json.Marshal(dst.SubOperationTimingPluginResponse)
		if string(jsonSubOperationTimingPluginResponse) == "{}" { // empty struct
			dst.SubOperationTimingPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.SubOperationTimingPluginResponse = nil
	}

	// try to unmarshal data into ThirdPartyPluginResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartyPluginResponse)
	if err == nil {
		jsonThirdPartyPluginResponse, _ := json.Marshal(dst.ThirdPartyPluginResponse)
		if string(jsonThirdPartyPluginResponse) == "{}" { // empty struct
			dst.ThirdPartyPluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartyPluginResponse = nil
	}

	// try to unmarshal data into UniqueAttributePluginResponse
	err = newStrictDecoder(data).Decode(&dst.UniqueAttributePluginResponse)
	if err == nil {
		jsonUniqueAttributePluginResponse, _ := json.Marshal(dst.UniqueAttributePluginResponse)
		if string(jsonUniqueAttributePluginResponse) == "{}" { // empty struct
			dst.UniqueAttributePluginResponse = nil
		} else {
			match++
		}
	} else {
		dst.UniqueAttributePluginResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AttributeMapperPluginResponse = nil
		dst.CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse = nil
		dst.CleanUpExpiredPingfederatePersistentSessionsPluginResponse = nil
		dst.CleanUpInactivePingfederatePersistentSessionsPluginResponse = nil
		dst.CoalesceModificationsPluginResponse = nil
		dst.ComposedAttributePluginResponse = nil
		dst.DelayPluginResponse = nil
		dst.DnMapperPluginResponse = nil
		dst.GroovyScriptedPluginResponse = nil
		dst.InternalSearchRatePluginResponse = nil
		dst.ModifiablePasswordPolicyStatePluginResponse = nil
		dst.PassThroughAuthenticationPluginResponse = nil
		dst.PeriodicGcPluginResponse = nil
		dst.PeriodicStatsLoggerPluginResponse = nil
		dst.PingOnePassThroughAuthenticationPluginResponse = nil
		dst.PluggablePassThroughAuthenticationPluginResponse = nil
		dst.PurgeExpiredDataPluginResponse = nil
		dst.ReferentialIntegrityPluginResponse = nil
		dst.ReferralOnUpdatePluginResponse = nil
		dst.SearchShutdownPluginResponse = nil
		dst.SevenBitCleanPluginResponse = nil
		dst.SimpleToExternalBindPluginResponse = nil
		dst.SnmpSubagentPluginResponse = nil
		dst.SubOperationTimingPluginResponse = nil
		dst.ThirdPartyPluginResponse = nil
		dst.UniqueAttributePluginResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddPlugin200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddPlugin200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddPlugin200Response) MarshalJSON() ([]byte, error) {
	if src.AttributeMapperPluginResponse != nil {
		return json.Marshal(&src.AttributeMapperPluginResponse)
	}

	if src.CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse != nil {
		return json.Marshal(&src.CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse)
	}

	if src.CleanUpExpiredPingfederatePersistentSessionsPluginResponse != nil {
		return json.Marshal(&src.CleanUpExpiredPingfederatePersistentSessionsPluginResponse)
	}

	if src.CleanUpInactivePingfederatePersistentSessionsPluginResponse != nil {
		return json.Marshal(&src.CleanUpInactivePingfederatePersistentSessionsPluginResponse)
	}

	if src.CoalesceModificationsPluginResponse != nil {
		return json.Marshal(&src.CoalesceModificationsPluginResponse)
	}

	if src.ComposedAttributePluginResponse != nil {
		return json.Marshal(&src.ComposedAttributePluginResponse)
	}

	if src.DelayPluginResponse != nil {
		return json.Marshal(&src.DelayPluginResponse)
	}

	if src.DnMapperPluginResponse != nil {
		return json.Marshal(&src.DnMapperPluginResponse)
	}

	if src.GroovyScriptedPluginResponse != nil {
		return json.Marshal(&src.GroovyScriptedPluginResponse)
	}

	if src.InternalSearchRatePluginResponse != nil {
		return json.Marshal(&src.InternalSearchRatePluginResponse)
	}

	if src.ModifiablePasswordPolicyStatePluginResponse != nil {
		return json.Marshal(&src.ModifiablePasswordPolicyStatePluginResponse)
	}

	if src.PassThroughAuthenticationPluginResponse != nil {
		return json.Marshal(&src.PassThroughAuthenticationPluginResponse)
	}

	if src.PeriodicGcPluginResponse != nil {
		return json.Marshal(&src.PeriodicGcPluginResponse)
	}

	if src.PeriodicStatsLoggerPluginResponse != nil {
		return json.Marshal(&src.PeriodicStatsLoggerPluginResponse)
	}

	if src.PingOnePassThroughAuthenticationPluginResponse != nil {
		return json.Marshal(&src.PingOnePassThroughAuthenticationPluginResponse)
	}

	if src.PluggablePassThroughAuthenticationPluginResponse != nil {
		return json.Marshal(&src.PluggablePassThroughAuthenticationPluginResponse)
	}

	if src.PurgeExpiredDataPluginResponse != nil {
		return json.Marshal(&src.PurgeExpiredDataPluginResponse)
	}

	if src.ReferentialIntegrityPluginResponse != nil {
		return json.Marshal(&src.ReferentialIntegrityPluginResponse)
	}

	if src.ReferralOnUpdatePluginResponse != nil {
		return json.Marshal(&src.ReferralOnUpdatePluginResponse)
	}

	if src.SearchShutdownPluginResponse != nil {
		return json.Marshal(&src.SearchShutdownPluginResponse)
	}

	if src.SevenBitCleanPluginResponse != nil {
		return json.Marshal(&src.SevenBitCleanPluginResponse)
	}

	if src.SimpleToExternalBindPluginResponse != nil {
		return json.Marshal(&src.SimpleToExternalBindPluginResponse)
	}

	if src.SnmpSubagentPluginResponse != nil {
		return json.Marshal(&src.SnmpSubagentPluginResponse)
	}

	if src.SubOperationTimingPluginResponse != nil {
		return json.Marshal(&src.SubOperationTimingPluginResponse)
	}

	if src.ThirdPartyPluginResponse != nil {
		return json.Marshal(&src.ThirdPartyPluginResponse)
	}

	if src.UniqueAttributePluginResponse != nil {
		return json.Marshal(&src.UniqueAttributePluginResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddPlugin200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AttributeMapperPluginResponse != nil {
		return obj.AttributeMapperPluginResponse
	}

	if obj.CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse != nil {
		return obj.CleanUpExpiredPingfederatePersistentAccessGrantsPluginResponse
	}

	if obj.CleanUpExpiredPingfederatePersistentSessionsPluginResponse != nil {
		return obj.CleanUpExpiredPingfederatePersistentSessionsPluginResponse
	}

	if obj.CleanUpInactivePingfederatePersistentSessionsPluginResponse != nil {
		return obj.CleanUpInactivePingfederatePersistentSessionsPluginResponse
	}

	if obj.CoalesceModificationsPluginResponse != nil {
		return obj.CoalesceModificationsPluginResponse
	}

	if obj.ComposedAttributePluginResponse != nil {
		return obj.ComposedAttributePluginResponse
	}

	if obj.DelayPluginResponse != nil {
		return obj.DelayPluginResponse
	}

	if obj.DnMapperPluginResponse != nil {
		return obj.DnMapperPluginResponse
	}

	if obj.GroovyScriptedPluginResponse != nil {
		return obj.GroovyScriptedPluginResponse
	}

	if obj.InternalSearchRatePluginResponse != nil {
		return obj.InternalSearchRatePluginResponse
	}

	if obj.ModifiablePasswordPolicyStatePluginResponse != nil {
		return obj.ModifiablePasswordPolicyStatePluginResponse
	}

	if obj.PassThroughAuthenticationPluginResponse != nil {
		return obj.PassThroughAuthenticationPluginResponse
	}

	if obj.PeriodicGcPluginResponse != nil {
		return obj.PeriodicGcPluginResponse
	}

	if obj.PeriodicStatsLoggerPluginResponse != nil {
		return obj.PeriodicStatsLoggerPluginResponse
	}

	if obj.PingOnePassThroughAuthenticationPluginResponse != nil {
		return obj.PingOnePassThroughAuthenticationPluginResponse
	}

	if obj.PluggablePassThroughAuthenticationPluginResponse != nil {
		return obj.PluggablePassThroughAuthenticationPluginResponse
	}

	if obj.PurgeExpiredDataPluginResponse != nil {
		return obj.PurgeExpiredDataPluginResponse
	}

	if obj.ReferentialIntegrityPluginResponse != nil {
		return obj.ReferentialIntegrityPluginResponse
	}

	if obj.ReferralOnUpdatePluginResponse != nil {
		return obj.ReferralOnUpdatePluginResponse
	}

	if obj.SearchShutdownPluginResponse != nil {
		return obj.SearchShutdownPluginResponse
	}

	if obj.SevenBitCleanPluginResponse != nil {
		return obj.SevenBitCleanPluginResponse
	}

	if obj.SimpleToExternalBindPluginResponse != nil {
		return obj.SimpleToExternalBindPluginResponse
	}

	if obj.SnmpSubagentPluginResponse != nil {
		return obj.SnmpSubagentPluginResponse
	}

	if obj.SubOperationTimingPluginResponse != nil {
		return obj.SubOperationTimingPluginResponse
	}

	if obj.ThirdPartyPluginResponse != nil {
		return obj.ThirdPartyPluginResponse
	}

	if obj.UniqueAttributePluginResponse != nil {
		return obj.UniqueAttributePluginResponse
	}

	// all schemas are nil
	return nil
}

type NullableAddPlugin200Response struct {
	value *AddPlugin200Response
	isSet bool
}

func (v NullableAddPlugin200Response) Get() *AddPlugin200Response {
	return v.value
}

func (v *NullableAddPlugin200Response) Set(val *AddPlugin200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPlugin200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPlugin200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPlugin200Response(val *AddPlugin200Response) *NullableAddPlugin200Response {
	return &NullableAddPlugin200Response{value: val, isSet: true}
}

func (v NullableAddPlugin200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPlugin200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
