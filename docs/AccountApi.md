# \AccountApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountInfoGet**](AccountApi.md#AccountInfoGet) | **Get** /account/info | Get account usage info
[**AccountListGet**](AccountApi.md#AccountListGet) | **Get** /account/list | List user accounts
[**AccountLogoGet**](AccountApi.md#AccountLogoGet) | **Get** /account/logo | Get account logo
[**AccountMetadataGet**](AccountApi.md#AccountMetadataGet) | **Get** /account/metadata | Get account metadata
[**AccountRolesGet**](AccountApi.md#AccountRolesGet) | **Get** /account/roles | Call users of the account


# **AccountInfoGet**
> AccountInfoResp AccountInfoGet(ctx, )
Get account usage info

Retrieve details of the account e.g. service settings, virus scan details, available and used storage, set automations, etc. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AccountInfoResp**](AccountInfoResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountListGet**
> []AccountListRespItems AccountListGet(ctx, )
List user accounts

Get the list with user accounts displaying the ID, host name, plan and status. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AccountListRespItems**](AccountListRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountLogoGet**
> AccountLogoGet(ctx, )
Get account logo

Retrieve the logo set for the current account. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountMetadataGet**
> AccountMetadataResp AccountMetadataGet(ctx, )
Get account metadata

Retrieve general settings for the account e.g. status of the user, identity providers, language, set logo and banner text. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AccountMetadataResp**](AccountMetadataResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountRolesGet**
> []AccountRolesRespItems AccountRolesGet(ctx, )
Call users of the account

Get the list of the current account users with their IDs, names and emails 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AccountRolesRespItems**](AccountRolesRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

