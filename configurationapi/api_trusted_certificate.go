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

// TrustedCertificateAPIService TrustedCertificateAPI service
type TrustedCertificateAPIService service

type ApiAddTrustedCertificateRequest struct {
	ctx                          context.Context
	ApiService                   *TrustedCertificateAPIService
	addTrustedCertificateRequest *AddTrustedCertificateRequest
}

// Create a new Trusted Certificate in the config
func (r ApiAddTrustedCertificateRequest) AddTrustedCertificateRequest(addTrustedCertificateRequest AddTrustedCertificateRequest) ApiAddTrustedCertificateRequest {
	r.addTrustedCertificateRequest = &addTrustedCertificateRequest
	return r
}

func (r ApiAddTrustedCertificateRequest) Execute() (*TrustedCertificateResponse, *http.Response, error) {
	return r.ApiService.AddTrustedCertificateExecute(r)
}

/*
AddTrustedCertificate Add a new Trusted Certificate to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddTrustedCertificateRequest
*/
func (a *TrustedCertificateAPIService) AddTrustedCertificate(ctx context.Context) ApiAddTrustedCertificateRequest {
	return ApiAddTrustedCertificateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TrustedCertificateResponse
func (a *TrustedCertificateAPIService) AddTrustedCertificateExecute(r ApiAddTrustedCertificateRequest) (*TrustedCertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TrustedCertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrustedCertificateAPIService.AddTrustedCertificate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trusted-certificates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addTrustedCertificateRequest == nil {
		return localVarReturnValue, nil, reportError("addTrustedCertificateRequest is required and must be specified")
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
	localVarPostBody = r.addTrustedCertificateRequest
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

type ApiDeleteTrustedCertificateRequest struct {
	ctx                    context.Context
	ApiService             *TrustedCertificateAPIService
	trustedCertificateName string
}

func (r ApiDeleteTrustedCertificateRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteTrustedCertificateExecute(r)
}

/*
DeleteTrustedCertificate Delete a Trusted Certificate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param trustedCertificateName Name of the Trusted Certificate
	@return ApiDeleteTrustedCertificateRequest
*/
func (a *TrustedCertificateAPIService) DeleteTrustedCertificate(ctx context.Context, trustedCertificateName string) ApiDeleteTrustedCertificateRequest {
	return ApiDeleteTrustedCertificateRequest{
		ApiService:             a,
		ctx:                    ctx,
		trustedCertificateName: trustedCertificateName,
	}
}

// Execute executes the request
func (a *TrustedCertificateAPIService) DeleteTrustedCertificateExecute(r ApiDeleteTrustedCertificateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrustedCertificateAPIService.DeleteTrustedCertificate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trusted-certificates/{trusted-certificate-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"trusted-certificate-name"+"}", url.PathEscape(parameterValueToString(r.trustedCertificateName, "trustedCertificateName")), -1)

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

type ApiGetTrustedCertificateRequest struct {
	ctx                    context.Context
	ApiService             *TrustedCertificateAPIService
	trustedCertificateName string
}

func (r ApiGetTrustedCertificateRequest) Execute() (*TrustedCertificateResponse, *http.Response, error) {
	return r.ApiService.GetTrustedCertificateExecute(r)
}

/*
GetTrustedCertificate Returns a single Trusted Certificate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param trustedCertificateName Name of the Trusted Certificate
	@return ApiGetTrustedCertificateRequest
*/
func (a *TrustedCertificateAPIService) GetTrustedCertificate(ctx context.Context, trustedCertificateName string) ApiGetTrustedCertificateRequest {
	return ApiGetTrustedCertificateRequest{
		ApiService:             a,
		ctx:                    ctx,
		trustedCertificateName: trustedCertificateName,
	}
}

// Execute executes the request
//
//	@return TrustedCertificateResponse
func (a *TrustedCertificateAPIService) GetTrustedCertificateExecute(r ApiGetTrustedCertificateRequest) (*TrustedCertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TrustedCertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrustedCertificateAPIService.GetTrustedCertificate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trusted-certificates/{trusted-certificate-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"trusted-certificate-name"+"}", url.PathEscape(parameterValueToString(r.trustedCertificateName, "trustedCertificateName")), -1)

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

type ApiListTrustedCertificatesRequest struct {
	ctx        context.Context
	ApiService *TrustedCertificateAPIService
	filter     *string
}

// SCIM filter
func (r ApiListTrustedCertificatesRequest) Filter(filter string) ApiListTrustedCertificatesRequest {
	r.filter = &filter
	return r
}

func (r ApiListTrustedCertificatesRequest) Execute() (*TrustedCertificateListResponse, *http.Response, error) {
	return r.ApiService.ListTrustedCertificatesExecute(r)
}

/*
ListTrustedCertificates Returns a list of all Trusted Certificate objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListTrustedCertificatesRequest
*/
func (a *TrustedCertificateAPIService) ListTrustedCertificates(ctx context.Context) ApiListTrustedCertificatesRequest {
	return ApiListTrustedCertificatesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TrustedCertificateListResponse
func (a *TrustedCertificateAPIService) ListTrustedCertificatesExecute(r ApiListTrustedCertificatesRequest) (*TrustedCertificateListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TrustedCertificateListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrustedCertificateAPIService.ListTrustedCertificates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trusted-certificates"

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

type ApiUpdateTrustedCertificateRequest struct {
	ctx                    context.Context
	ApiService             *TrustedCertificateAPIService
	trustedCertificateName string
	updateRequest          *UpdateRequest
}

// Update an existing Trusted Certificate
func (r ApiUpdateTrustedCertificateRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateTrustedCertificateRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateTrustedCertificateRequest) Execute() (*TrustedCertificateResponse, *http.Response, error) {
	return r.ApiService.UpdateTrustedCertificateExecute(r)
}

/*
UpdateTrustedCertificate Update an existing Trusted Certificate by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param trustedCertificateName Name of the Trusted Certificate
	@return ApiUpdateTrustedCertificateRequest
*/
func (a *TrustedCertificateAPIService) UpdateTrustedCertificate(ctx context.Context, trustedCertificateName string) ApiUpdateTrustedCertificateRequest {
	return ApiUpdateTrustedCertificateRequest{
		ApiService:             a,
		ctx:                    ctx,
		trustedCertificateName: trustedCertificateName,
	}
}

// Execute executes the request
//
//	@return TrustedCertificateResponse
func (a *TrustedCertificateAPIService) UpdateTrustedCertificateExecute(r ApiUpdateTrustedCertificateRequest) (*TrustedCertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TrustedCertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrustedCertificateAPIService.UpdateTrustedCertificate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trusted-certificates/{trusted-certificate-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"trusted-certificate-name"+"}", url.PathEscape(parameterValueToString(r.trustedCertificateName, "trustedCertificateName")), -1)

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
