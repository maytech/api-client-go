# \AccountApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountInfoGet**](AccountApi.md#AccountInfoGet) | **Get** /account/info | Account info
[**AccountListGet**](AccountApi.md#AccountListGet) | **Get** /account/list | List user accounts
[**AccountLogoGet**](AccountApi.md#AccountLogoGet) | **Get** /account/logo | Account logo
[**AccountMetadataGet**](AccountApi.md#AccountMetadataGet) | **Get** /account/metadata | Account metadata
[**AccountRolesGet**](AccountApi.md#AccountRolesGet) | **Get** /account/roles | Account roles


# **AccountInfoGet**
> AccountInfoResp AccountInfoGet(ctx, )
Account info

Account usage info 

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

List user accounts 

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
Account logo

Returns account logo body 

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
Account metadata

Get Account Public Metadata 

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
Account roles

List account roles 

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

