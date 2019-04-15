# \InboxApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InboxGet**](InboxApi.md#InboxGet) | **Get** /inbox | Get contact&#39;s share tracking


# **InboxGet**
> []InboxRespItems InboxGet(ctx, optional)
Get contact's share tracking

Retrieve the tracking of actions where the contact was the recipient. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**string**](.md)| Inbox ID | 
 **limit** | **float32**| Rows per page | [default to 100]
 **from** | **float32**| UTC timestamp | [default to 0]
 **to** | **float32**| UTC timestamp | 

### Return type

[**[]InboxRespItems**](InboxRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

