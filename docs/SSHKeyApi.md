# \SSHKeyApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SshKeyCreatePost**](SSHKeyApi.md#SshKeyCreatePost) | **Post** /ssh-key/create | Create a new SSH key
[**SshKeyDeletePost**](SSHKeyApi.md#SshKeyDeletePost) | **Post** /ssh-key/delete | Delete SSH key
[**SshKeyEditPost**](SSHKeyApi.md#SshKeyEditPost) | **Post** /ssh-key/edit | Edit SSH key metadata
[**SshKeyGet**](SSHKeyApi.md#SshKeyGet) | **Get** /ssh-key | List available SSH keys
[**SshKeyMetadataIdGet**](SSHKeyApi.md#SshKeyMetadataIdGet) | **Get** /ssh-key/metadata/{id} | Get SSH key metadata


# **SshKeyCreatePost**
> SshKeyResp SshKeyCreatePost(ctx, body)
Create a new SSH key

Generate a new SSH key. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**SshKeyCreateReq**](SshKeyCreateReq.md)|  | 

### Return type

[**SshKeyResp**](SshKeyResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SshKeyDeletePost**
> IdResp SshKeyDeletePost(ctx, body)
Delete SSH key

Remove the key specified by SSH key ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**SshKeyDeleteReq**](SshKeyDeleteReq.md)|  | 

### Return type

[**IdResp**](IdResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SshKeyEditPost**
> SshKeyResp SshKeyEditPost(ctx, body)
Edit SSH key metadata

Update the metadata of SSH key. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**SshKeyEditReq**](SshKeyEditReq.md)|  | 

### Return type

[**SshKeyResp**](SshKeyResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SshKeyGet**
> []SshKeyResp SshKeyGet(ctx, )
List available SSH keys

Get a list of all SSH keys of the logged-in user. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]SshKeyResp**](SshKeyResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SshKeyMetadataIdGet**
> SshKeyResp SshKeyMetadataIdGet(ctx, id)
Get SSH key metadata

Retrieve the metadata of SSH key by the given ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| SSH key ID | 

### Return type

[**SshKeyResp**](SshKeyResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

