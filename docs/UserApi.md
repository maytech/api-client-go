# \UserApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserCreatePost**](UserApi.md#UserCreatePost) | **Post** /user/create | Create a user
[**UserDeletePost**](UserApi.md#UserDeletePost) | **Post** /user/delete | Delete users
[**UserEditPost**](UserApi.md#UserEditPost) | **Post** /user/edit | Update metadata of users
[**UserGet**](UserApi.md#UserGet) | **Get** /user | List users
[**UserGroupGet**](UserApi.md#UserGroupGet) | **Get** /user/group | List all user groups
[**UserMetadataIdGet**](UserApi.md#UserMetadataIdGet) | **Get** /user/metadata/{id} | Get user metadata
[**UserPgpKeyIdGet**](UserApi.md#UserPgpKeyIdGet) | **Get** /user/pgp-key/{id} | Get user&#39;s PGP key
[**UserRemoveMfaPost**](UserApi.md#UserRemoveMfaPost) | **Post** /user/remove-mfa | Disable MFA for users
[**UserResetMfaPost**](UserApi.md#UserResetMfaPost) | **Post** /user/reset-mfa | Update existing MFA settings for users
[**UserSetMfaPost**](UserApi.md#UserSetMfaPost) | **Post** /user/set-mfa | Enable MFA for users
[**UserSignupPost**](UserApi.md#UserSignupPost) | **Post** /user/signup | Register a new user


# **UserCreatePost**
> UserResp UserCreatePost(ctx, body)
Create a user

Add a new user to the account. The user receives an email where the password can be set. 

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

Remove users from the account by specified user IDs. 

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
Update metadata of users

Edit the existing users’ metadata. An email can be edited only for one user using one API call. In this case both old and new emails will be notified about the change. 

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

Retrieve a list of users 

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
List all user groups

Retrieve a list of all user groups that can be set for users. 

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
Get user metadata

Retrieve user’s metadata by given user ID. 

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
Get user's PGP key

Use to retrieve PGP key generated by the user. 

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
Disable MFA for users

Deactivate MFA for a user by specified user ID. If 2FA is forced for the account, the user won’t be able to deactivate it using this API call. 

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

# **UserResetMfaPost**
> ProfileRemoveMfaResp UserResetMfaPost(ctx, body)
Update existing MFA settings for users

Remove already set MFA code values, the MFA will be left activated. All logged-in sessions of users will be stopped. On the next login the user should set MFA again. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**UserResetMfaReq**](UserResetMfaReq.md)|  | 

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
Enable MFA for users

Activate MFA for a user by specified user ID. If auth type is 2FA, the user will be forced to use MFA. All logged-in sessions of users will be stopped. 

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
Register a new user

Set password for a new user. 

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

