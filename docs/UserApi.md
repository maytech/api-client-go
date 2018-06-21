# \UserApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserCreatePost**](UserApi.md#UserCreatePost) | **Post** /user/create | Create user
[**UserDeletePost**](UserApi.md#UserDeletePost) | **Post** /user/delete | Delete users
[**UserEditPost**](UserApi.md#UserEditPost) | **Post** /user/edit | User metadata
[**UserGet**](UserApi.md#UserGet) | **Get** /user | List users
[**UserGroupGet**](UserApi.md#UserGroupGet) | **Get** /user/group | All user groups
[**UserMetadataIdGet**](UserApi.md#UserMetadataIdGet) | **Get** /user/metadata/{id} | User metadata
[**UserPgpKeyIdGet**](UserApi.md#UserPgpKeyIdGet) | **Get** /user/pgp-key/{id} | User PGP key
[**UserRemoveMfaPost**](UserApi.md#UserRemoveMfaPost) | **Post** /user/remove-mfa | Remove MFA for user
[**UserSetMfaPost**](UserApi.md#UserSetMfaPost) | **Post** /user/set-mfa | Set MFA enabled for user
[**UserSignupPost**](UserApi.md#UserSignupPost) | **Post** /user/signup | Signup existing user


# **UserCreatePost**
> UserResp UserCreatePost(ctx, body)
Create user

Create new user 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**UserCreateReq**](UserCreateReq.md)|  | 

### Return type

[**UserResp**](UserResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserDeletePost**
> JobResp UserDeletePost(ctx, body)
Delete users

Delete users 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**UserDeleteReq**](UserDeleteReq.md)|  | 

### Return type

[**JobResp**](JobResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserEditPost**
> []UserResp UserEditPost(ctx, body)
User metadata

Edit user 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**UserEditReq**](UserEditReq.md)|  | 

### Return type

[**[]UserResp**](UserResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGet**
> []UserResp UserGet(ctx, )
List users

List users 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UserResp**](UserResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGroupGet**
> []Group UserGroupGet(ctx, )
All user groups

List all user groups 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Group**](Group.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserMetadataIdGet**
> UserResp UserMetadataIdGet(ctx, id)
User metadata

Get user metadata 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a user | 

### Return type

[**UserResp**](UserResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPgpKeyIdGet**
> PgpKeyResp UserPgpKeyIdGet(ctx, id)
User PGP key

Get user PGP key 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a user | 

### Return type

[**PgpKeyResp**](PgpKeyResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserRemoveMfaPost**
> ProfileRemoveMfaResp UserRemoveMfaPost(ctx, body)
Remove MFA for user

Remove MFA for user 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**UserRemoveMfaReq**](UserRemoveMfaReq.md)|  | 

### Return type

[**ProfileRemoveMfaResp**](ProfileRemoveMfaResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserSetMfaPost**
> ProfileRemoveMfaResp UserSetMfaPost(ctx, body)
Set MFA enabled for user

Set multi factor autorization method (MFA) enabled for user 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**UserSetMfaReq**](UserSetMfaReq.md)|  | 

### Return type

[**ProfileRemoveMfaResp**](ProfileRemoveMfaResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserSignupPost**
> UserSignupPost(ctx, body)
Signup existing user

Signup existing user 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**UserSignupReq**](UserSignupReq.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: raw

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

