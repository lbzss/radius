//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20220315privatepreview

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DaprInvokeHTTPRoutesClient contains the methods for the DaprInvokeHTTPRoutes group.
// Don't use this type directly, use NewDaprInvokeHTTPRoutesClient() instead.
type DaprInvokeHTTPRoutesClient struct {
	ep string
	pl runtime.Pipeline
	subscriptionID string
}

// NewDaprInvokeHTTPRoutesClient creates a new instance of DaprInvokeHTTPRoutesClient with the specified values.
func NewDaprInvokeHTTPRoutesClient(con *arm.Connection, subscriptionID string) *DaprInvokeHTTPRoutesClient {
	return &DaprInvokeHTTPRoutesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or updates a DaprInvokeHttpRoute resource
// If the operation fails it returns the *ErrorResponse error type.
func (client *DaprInvokeHTTPRoutesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, daprInvokeHTTPRouteName string, daprInvokeHTTPRouteParameters DaprInvokeHTTPRouteResource, options *DaprInvokeHTTPRoutesCreateOrUpdateOptions) (DaprInvokeHTTPRoutesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, daprInvokeHTTPRouteName, daprInvokeHTTPRouteParameters, options)
	if err != nil {
		return DaprInvokeHTTPRoutesCreateOrUpdateResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return DaprInvokeHTTPRoutesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DaprInvokeHTTPRoutesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DaprInvokeHTTPRoutesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, daprInvokeHTTPRouteName string, daprInvokeHTTPRouteParameters DaprInvokeHTTPRouteResource, options *DaprInvokeHTTPRoutesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Applications.Connector/daprInvokeHttpRoutes/{daprInvokeHttpRouteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if daprInvokeHTTPRouteName == "" {
		return nil, errors.New("parameter daprInvokeHTTPRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprInvokeHttpRouteName}", url.PathEscape(daprInvokeHTTPRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, daprInvokeHTTPRouteParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DaprInvokeHTTPRoutesClient) createOrUpdateHandleResponse(resp *http.Response) (DaprInvokeHTTPRoutesCreateOrUpdateResponse, error) {
	result := DaprInvokeHTTPRoutesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprInvokeHTTPRouteResource); err != nil {
		return DaprInvokeHTTPRoutesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DaprInvokeHTTPRoutesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes an existing daprInvokeHttpRoute resource
// If the operation fails it returns the *ErrorResponse error type.
func (client *DaprInvokeHTTPRoutesClient) Delete(ctx context.Context, resourceGroupName string, daprInvokeHTTPRouteName string, options *DaprInvokeHTTPRoutesDeleteOptions) (DaprInvokeHTTPRoutesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, daprInvokeHTTPRouteName, options)
	if err != nil {
		return DaprInvokeHTTPRoutesDeleteResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return DaprInvokeHTTPRoutesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return DaprInvokeHTTPRoutesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return DaprInvokeHTTPRoutesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DaprInvokeHTTPRoutesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, daprInvokeHTTPRouteName string, options *DaprInvokeHTTPRoutesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Applications.Connector/daprInvokeHttpRoutes/{daprInvokeHttpRouteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if daprInvokeHTTPRouteName == "" {
		return nil, errors.New("parameter daprInvokeHTTPRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprInvokeHttpRouteName}", url.PathEscape(daprInvokeHTTPRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DaprInvokeHTTPRoutesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Retrieves information about a daprInvokeHttpRoute resource
// If the operation fails it returns the *ErrorResponse error type.
func (client *DaprInvokeHTTPRoutesClient) Get(ctx context.Context, resourceGroupName string, daprInvokeHTTPRouteName string, options *DaprInvokeHTTPRoutesGetOptions) (DaprInvokeHTTPRoutesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, daprInvokeHTTPRouteName, options)
	if err != nil {
		return DaprInvokeHTTPRoutesGetResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return DaprInvokeHTTPRoutesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DaprInvokeHTTPRoutesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DaprInvokeHTTPRoutesClient) getCreateRequest(ctx context.Context, resourceGroupName string, daprInvokeHTTPRouteName string, options *DaprInvokeHTTPRoutesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Applications.Connector/daprInvokeHttpRoutes/{daprInvokeHttpRouteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if daprInvokeHTTPRouteName == "" {
		return nil, errors.New("parameter daprInvokeHTTPRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprInvokeHttpRouteName}", url.PathEscape(daprInvokeHTTPRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DaprInvokeHTTPRoutesClient) getHandleResponse(resp *http.Response) (DaprInvokeHTTPRoutesGetResponse, error) {
	result := DaprInvokeHTTPRoutesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprInvokeHTTPRouteResource); err != nil {
		return DaprInvokeHTTPRoutesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DaprInvokeHTTPRoutesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Lists information about all daprInvokeHttpRoute resources in the given subscription and resource group
// If the operation fails it returns the *ErrorResponse error type.
func (client *DaprInvokeHTTPRoutesClient) List(resourceGroupName string, options *DaprInvokeHTTPRoutesListOptions) (*DaprInvokeHTTPRoutesListPager) {
	return &DaprInvokeHTTPRoutesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp DaprInvokeHTTPRoutesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DaprInvokeHTTPRouteList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *DaprInvokeHTTPRoutesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *DaprInvokeHTTPRoutesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Applications.Connector/daprInvokeHttpRoutes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DaprInvokeHTTPRoutesClient) listHandleResponse(resp *http.Response) (DaprInvokeHTTPRoutesListResponse, error) {
	result := DaprInvokeHTTPRoutesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprInvokeHTTPRouteList); err != nil {
		return DaprInvokeHTTPRoutesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *DaprInvokeHTTPRoutesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - Lists information about all daprInvokeHttpRoute resources in the given subscription
// If the operation fails it returns the *ErrorResponse error type.
func (client *DaprInvokeHTTPRoutesClient) ListBySubscription(options *DaprInvokeHTTPRoutesListBySubscriptionOptions) (*DaprInvokeHTTPRoutesListBySubscriptionPager) {
	return &DaprInvokeHTTPRoutesListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DaprInvokeHTTPRoutesListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DaprInvokeHTTPRouteList.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DaprInvokeHTTPRoutesClient) listBySubscriptionCreateRequest(ctx context.Context, options *DaprInvokeHTTPRoutesListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Applications.Connector/daprInvokeHttpRoutes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DaprInvokeHTTPRoutesClient) listBySubscriptionHandleResponse(resp *http.Response) (DaprInvokeHTTPRoutesListBySubscriptionResponse, error) {
	result := DaprInvokeHTTPRoutesListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprInvokeHTTPRouteList); err != nil {
		return DaprInvokeHTTPRoutesListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *DaprInvokeHTTPRoutesClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

