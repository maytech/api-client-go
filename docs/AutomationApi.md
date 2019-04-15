# \AutomationApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutomationCreatePost**](AutomationApi.md#AutomationCreatePost) | **Post** /automation/create | Create a new automation
[**AutomationDeletePost**](AutomationApi.md#AutomationDeletePost) | **Post** /automation/delete | Delete automations
[**AutomationEditPost**](AutomationApi.md#AutomationEditPost) | **Post** /automation/edit/ | Edit an existing automation
[**AutomationGet**](AutomationApi.md#AutomationGet) | **Get** /automation | List all automations
[**AutomationMetadataIdGet**](AutomationApi.md#AutomationMetadataIdGet) | **Get** /automation/metadata/{id} | Get automation metadata


# **AutomationCreatePost**
> AutomationResp AutomationCreatePost(ctx, body)
Create a new automation

Add a new automatic operation. 

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

Delete a rule set for automatic operations. 

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
Edit an existing automation

Update an existing rule for a specified automatic operation. 

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
List all automations

Retrieve a list of all automations of the user. 

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
Get automation metadata

Get the automation metadata by the specified automation ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| automation ID | 

### Return type

[**AutomationResp**](AutomationResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

