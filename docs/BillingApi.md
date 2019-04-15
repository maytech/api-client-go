# \BillingApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingUpgradePost**](BillingApi.md#BillingUpgradePost) | **Post** /billing/upgrade | Upgrade the user&#39;s account


# **BillingUpgradePost**
> BillingUpgradeResp BillingUpgradePost(ctx, body)
Upgrade the user's account

Upgrade the user's account. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**BillingUpgradeReq**](BillingUpgradeReq.md)|  | 

### Return type

[**BillingUpgradeResp**](BillingUpgradeResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

