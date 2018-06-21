# \AutomationApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutomationCreatePost**](AutomationApi.md#AutomationCreatePost) | **Post** /automation/create | Create automation
[**AutomationDeletePost**](AutomationApi.md#AutomationDeletePost) | **Post** /automation/delete | Delete automations
[**AutomationEditPost**](AutomationApi.md#AutomationEditPost) | **Post** /automation/edit/ | Edit automation
[**AutomationGet**](AutomationApi.md#AutomationGet) | **Get** /automation | List automations
[**AutomationMetadataIdGet**](AutomationApi.md#AutomationMetadataIdGet) | **Get** /automation/metadata/{id} | Automation metadata


# **AutomationCreatePost**
> AutomationResp AutomationCreatePost(ctx, body)
Create automation

Create new automation 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**AutomationCreateReq**](AutomationCreateReq.md)|  | 

### Return type

[**AutomationResp**](AutomationResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutomationDeletePost**
> AutomationDeleteResp AutomationDeletePost(ctx, body)
Delete automations

Delete automations 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**IdsReq**](IdsReq.md)|  | 

### Return type

[**AutomationDeleteResp**](AutomationDeleteResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutomationEditPost**
> []AutomationResp AutomationEditPost(ctx, body)
Edit automation

Edit automation 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**AutomationEditReq**](AutomationEditReq.md)|  | 

### Return type

[**[]AutomationResp**](AutomationResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutomationGet**
> []AutomationResp AutomationGet(ctx, )
List automations

List of all automations 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AutomationResp**](AutomationResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutomationMetadataIdGet**
> AutomationResp AutomationMetadataIdGet(ctx, id)
Automation metadata

Get automation metadata 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of an automation | 

### Return type

[**AutomationResp**](AutomationResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

