/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// OtpDeliveryMechanismAPIService OtpDeliveryMechanismAPI service
type OtpDeliveryMechanismAPIService service

type ApiAddOtpDeliveryMechanismRequest struct {
	ctx                            context.Context
	ApiService                     *OtpDeliveryMechanismAPIService
	addOtpDeliveryMechanismRequest *AddOtpDeliveryMechanismRequest
}

// Create a new OTP Delivery Mechanism in the config
func (r ApiAddOtpDeliveryMechanismRequest) AddOtpDeliveryMechanismRequest(addOtpDeliveryMechanismRequest AddOtpDeliveryMechanismRequest) ApiAddOtpDeliveryMechanismRequest {
	r.addOtpDeliveryMechanismRequest = &addOtpDeliveryMechanismRequest
	return r
}

func (r ApiAddOtpDeliveryMechanismRequest) Execute() (*AddOtpDeliveryMechanism200Response, *http.Response, error) {
	return r.ApiService.AddOtpDeliveryMechanismExecute(r)
}

/*
AddOtpDeliveryMechanism Add a new OTP Delivery Mechanism to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddOtpDeliveryMechanismRequest
*/
func (a *OtpDeliveryMechanismAPIService) AddOtpDeliveryMechanism(ctx context.Context) ApiAddOtpDeliveryMechanismRequest {
	return ApiAddOtpDeliveryMechanismRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddOtpDeliveryMechanism200Response
func (a *OtpDeliveryMechanismAPIService) AddOtpDeliveryMechanismExecute(r ApiAddOtpDeliveryMechanismRequest) (*AddOtpDeliveryMechanism200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddOtpDeliveryMechanism200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OtpDeliveryMechanismAPIService.AddOtpDeliveryMechanism")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/otp-delivery-mechanisms"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addOtpDeliveryMechanismRequest == nil {
		return localVarReturnValue, nil, reportError("addOtpDeliveryMechanismRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.addOtpDeliveryMechanismRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteOtpDeliveryMechanismRequest struct {
	ctx                      context.Context
	ApiService               *OtpDeliveryMechanismAPIService
	otpDeliveryMechanismName string
}

func (r ApiDeleteOtpDeliveryMechanismRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOtpDeliveryMechanismExecute(r)
}

/*
DeleteOtpDeliveryMechanism Delete a OTP Delivery Mechanism

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param otpDeliveryMechanismName Name of the OTP Delivery Mechanism
	@return ApiDeleteOtpDeliveryMechanismRequest
*/
func (a *OtpDeliveryMechanismAPIService) DeleteOtpDeliveryMechanism(ctx context.Context, otpDeliveryMechanismName string) ApiDeleteOtpDeliveryMechanismRequest {
	return ApiDeleteOtpDeliveryMechanismRequest{
		ApiService:               a,
		ctx:                      ctx,
		otpDeliveryMechanismName: otpDeliveryMechanismName,
	}
}

// Execute executes the request
func (a *OtpDeliveryMechanismAPIService) DeleteOtpDeliveryMechanismExecute(r ApiDeleteOtpDeliveryMechanismRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OtpDeliveryMechanismAPIService.DeleteOtpDeliveryMechanism")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/otp-delivery-mechanisms/{otp-delivery-mechanism-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"otp-delivery-mechanism-name"+"}", url.PathEscape(parameterValueToString(r.otpDeliveryMechanismName, "otpDeliveryMechanismName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetOtpDeliveryMechanismRequest struct {
	ctx                      context.Context
	ApiService               *OtpDeliveryMechanismAPIService
	otpDeliveryMechanismName string
}

func (r ApiGetOtpDeliveryMechanismRequest) Execute() (*AddOtpDeliveryMechanism200Response, *http.Response, error) {
	return r.ApiService.GetOtpDeliveryMechanismExecute(r)
}

/*
GetOtpDeliveryMechanism Returns a single OTP Delivery Mechanism

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param otpDeliveryMechanismName Name of the OTP Delivery Mechanism
	@return ApiGetOtpDeliveryMechanismRequest
*/
func (a *OtpDeliveryMechanismAPIService) GetOtpDeliveryMechanism(ctx context.Context, otpDeliveryMechanismName string) ApiGetOtpDeliveryMechanismRequest {
	return ApiGetOtpDeliveryMechanismRequest{
		ApiService:               a,
		ctx:                      ctx,
		otpDeliveryMechanismName: otpDeliveryMechanismName,
	}
}

// Execute executes the request
//
//	@return AddOtpDeliveryMechanism200Response
func (a *OtpDeliveryMechanismAPIService) GetOtpDeliveryMechanismExecute(r ApiGetOtpDeliveryMechanismRequest) (*AddOtpDeliveryMechanism200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddOtpDeliveryMechanism200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OtpDeliveryMechanismAPIService.GetOtpDeliveryMechanism")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/otp-delivery-mechanisms/{otp-delivery-mechanism-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"otp-delivery-mechanism-name"+"}", url.PathEscape(parameterValueToString(r.otpDeliveryMechanismName, "otpDeliveryMechanismName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListOtpDeliveryMechanismsRequest struct {
	ctx        context.Context
	ApiService *OtpDeliveryMechanismAPIService
	filter     *string
}

// SCIM filter
func (r ApiListOtpDeliveryMechanismsRequest) Filter(filter string) ApiListOtpDeliveryMechanismsRequest {
	r.filter = &filter
	return r
}

func (r ApiListOtpDeliveryMechanismsRequest) Execute() (*OtpDeliveryMechanismListResponse, *http.Response, error) {
	return r.ApiService.ListOtpDeliveryMechanismsExecute(r)
}

/*
ListOtpDeliveryMechanisms Returns a list of all OTP Delivery Mechanism objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListOtpDeliveryMechanismsRequest
*/
func (a *OtpDeliveryMechanismAPIService) ListOtpDeliveryMechanisms(ctx context.Context) ApiListOtpDeliveryMechanismsRequest {
	return ApiListOtpDeliveryMechanismsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OtpDeliveryMechanismListResponse
func (a *OtpDeliveryMechanismAPIService) ListOtpDeliveryMechanismsExecute(r ApiListOtpDeliveryMechanismsRequest) (*OtpDeliveryMechanismListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OtpDeliveryMechanismListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OtpDeliveryMechanismAPIService.ListOtpDeliveryMechanisms")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/otp-delivery-mechanisms"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateOtpDeliveryMechanismRequest struct {
	ctx                      context.Context
	ApiService               *OtpDeliveryMechanismAPIService
	otpDeliveryMechanismName string
	updateRequest            *UpdateRequest
}

// Update an existing OTP Delivery Mechanism
func (r ApiUpdateOtpDeliveryMechanismRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateOtpDeliveryMechanismRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateOtpDeliveryMechanismRequest) Execute() (*AddOtpDeliveryMechanism200Response, *http.Response, error) {
	return r.ApiService.UpdateOtpDeliveryMechanismExecute(r)
}

/*
UpdateOtpDeliveryMechanism Update an existing OTP Delivery Mechanism by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param otpDeliveryMechanismName Name of the OTP Delivery Mechanism
	@return ApiUpdateOtpDeliveryMechanismRequest
*/
func (a *OtpDeliveryMechanismAPIService) UpdateOtpDeliveryMechanism(ctx context.Context, otpDeliveryMechanismName string) ApiUpdateOtpDeliveryMechanismRequest {
	return ApiUpdateOtpDeliveryMechanismRequest{
		ApiService:               a,
		ctx:                      ctx,
		otpDeliveryMechanismName: otpDeliveryMechanismName,
	}
}

// Execute executes the request
//
//	@return AddOtpDeliveryMechanism200Response
func (a *OtpDeliveryMechanismAPIService) UpdateOtpDeliveryMechanismExecute(r ApiUpdateOtpDeliveryMechanismRequest) (*AddOtpDeliveryMechanism200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddOtpDeliveryMechanism200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OtpDeliveryMechanismAPIService.UpdateOtpDeliveryMechanism")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/otp-delivery-mechanisms/{otp-delivery-mechanism-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"otp-delivery-mechanism-name"+"}", url.PathEscape(parameterValueToString(r.otpDeliveryMechanismName, "otpDeliveryMechanismName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateRequest == nil {
		return localVarReturnValue, nil, reportError("updateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
