# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [mediator/v1/mediator.proto](#mediator_v1_mediator-proto)
    - [AuthVerifyServiceVerifyRequest](#mediator-v1-AuthVerifyServiceVerifyRequest)
    - [AuthVerifyServiceVerifyResponse](#mediator-v1-AuthVerifyServiceVerifyResponse)
    - [BranchProtection](#mediator-v1-BranchProtection)
    - [CheckHealthRequest](#mediator-v1-CheckHealthRequest)
    - [CheckHealthResponse](#mediator-v1-CheckHealthResponse)
    - [CreateGroupRequest](#mediator-v1-CreateGroupRequest)
    - [CreateGroupResponse](#mediator-v1-CreateGroupResponse)
    - [CreateOrganizationRequest](#mediator-v1-CreateOrganizationRequest)
    - [CreateOrganizationResponse](#mediator-v1-CreateOrganizationResponse)
    - [CreateRoleRequest](#mediator-v1-CreateRoleRequest)
    - [CreateRoleResponse](#mediator-v1-CreateRoleResponse)
    - [CreateUserRequest](#mediator-v1-CreateUserRequest)
    - [CreateUserResponse](#mediator-v1-CreateUserResponse)
    - [DeleteGroupRequest](#mediator-v1-DeleteGroupRequest)
    - [DeleteGroupResponse](#mediator-v1-DeleteGroupResponse)
    - [DeleteOrganizationRequest](#mediator-v1-DeleteOrganizationRequest)
    - [DeleteOrganizationResponse](#mediator-v1-DeleteOrganizationResponse)
    - [DeleteRoleRequest](#mediator-v1-DeleteRoleRequest)
    - [DeleteRoleResponse](#mediator-v1-DeleteRoleResponse)
    - [DeleteUserRequest](#mediator-v1-DeleteUserRequest)
    - [DeleteUserResponse](#mediator-v1-DeleteUserResponse)
    - [ExchangeCodeForTokenCLIRequest](#mediator-v1-ExchangeCodeForTokenCLIRequest)
    - [ExchangeCodeForTokenCLIResponse](#mediator-v1-ExchangeCodeForTokenCLIResponse)
    - [ExchangeCodeForTokenWEBRequest](#mediator-v1-ExchangeCodeForTokenWEBRequest)
    - [ExchangeCodeForTokenWEBResponse](#mediator-v1-ExchangeCodeForTokenWEBResponse)
    - [GetAuthorizationURLRequest](#mediator-v1-GetAuthorizationURLRequest)
    - [GetAuthorizationURLResponse](#mediator-v1-GetAuthorizationURLResponse)
    - [GetBranchProtectionRequest](#mediator-v1-GetBranchProtectionRequest)
    - [GetBranchProtectionResponse](#mediator-v1-GetBranchProtectionResponse)
    - [GetGroupByIdRequest](#mediator-v1-GetGroupByIdRequest)
    - [GetGroupByIdResponse](#mediator-v1-GetGroupByIdResponse)
    - [GetGroupByNameRequest](#mediator-v1-GetGroupByNameRequest)
    - [GetGroupByNameResponse](#mediator-v1-GetGroupByNameResponse)
    - [GetGroupsRequest](#mediator-v1-GetGroupsRequest)
    - [GetGroupsResponse](#mediator-v1-GetGroupsResponse)
    - [GetOrganizationByNameRequest](#mediator-v1-GetOrganizationByNameRequest)
    - [GetOrganizationByNameResponse](#mediator-v1-GetOrganizationByNameResponse)
    - [GetOrganizationRequest](#mediator-v1-GetOrganizationRequest)
    - [GetOrganizationResponse](#mediator-v1-GetOrganizationResponse)
    - [GetOrganizationsRequest](#mediator-v1-GetOrganizationsRequest)
    - [GetOrganizationsResponse](#mediator-v1-GetOrganizationsResponse)
    - [GetRoleByIdRequest](#mediator-v1-GetRoleByIdRequest)
    - [GetRoleByIdResponse](#mediator-v1-GetRoleByIdResponse)
    - [GetRoleByNameRequest](#mediator-v1-GetRoleByNameRequest)
    - [GetRoleByNameResponse](#mediator-v1-GetRoleByNameResponse)
    - [GetRolesRequest](#mediator-v1-GetRolesRequest)
    - [GetRolesResponse](#mediator-v1-GetRolesResponse)
    - [GetSecretByIdRequest](#mediator-v1-GetSecretByIdRequest)
    - [GetSecretByIdResponse](#mediator-v1-GetSecretByIdResponse)
    - [GetSecretsRequest](#mediator-v1-GetSecretsRequest)
    - [GetSecretsResponse](#mediator-v1-GetSecretsResponse)
    - [GetUserByEmailRequest](#mediator-v1-GetUserByEmailRequest)
    - [GetUserByEmailResponse](#mediator-v1-GetUserByEmailResponse)
    - [GetUserByIdRequest](#mediator-v1-GetUserByIdRequest)
    - [GetUserByIdResponse](#mediator-v1-GetUserByIdResponse)
    - [GetUserByUserNameRequest](#mediator-v1-GetUserByUserNameRequest)
    - [GetUserByUserNameResponse](#mediator-v1-GetUserByUserNameResponse)
    - [GetUsersRequest](#mediator-v1-GetUsersRequest)
    - [GetUsersResponse](#mediator-v1-GetUsersResponse)
    - [GetVulnerabilitiesRequest](#mediator-v1-GetVulnerabilitiesRequest)
    - [GetVulnerabilitiesResponse](#mediator-v1-GetVulnerabilitiesResponse)
    - [GetVulnerabilityByIdRequest](#mediator-v1-GetVulnerabilityByIdRequest)
    - [GetVulnerabilityByIdResponse](#mediator-v1-GetVulnerabilityByIdResponse)
    - [GroupRecord](#mediator-v1-GroupRecord)
    - [HandleGitHubWebhookRequest](#mediator-v1-HandleGitHubWebhookRequest)
    - [HandleGitHubWebhookResponse](#mediator-v1-HandleGitHubWebhookResponse)
    - [LogInRequest](#mediator-v1-LogInRequest)
    - [LogInResponse](#mediator-v1-LogInResponse)
    - [LogOutRequest](#mediator-v1-LogOutRequest)
    - [LogOutResponse](#mediator-v1-LogOutResponse)
    - [OrganizationRecord](#mediator-v1-OrganizationRecord)
    - [RoleRecord](#mediator-v1-RoleRecord)
    - [UserRecord](#mediator-v1-UserRecord)
  
    - [AuthVerifyService](#mediator-v1-AuthVerifyService)
    - [BranchProtectionService](#mediator-v1-BranchProtectionService)
    - [GitHubWebhookService](#mediator-v1-GitHubWebhookService)
    - [GroupService](#mediator-v1-GroupService)
    - [HealthService](#mediator-v1-HealthService)
    - [LogInService](#mediator-v1-LogInService)
    - [LogOutService](#mediator-v1-LogOutService)
    - [OAuthService](#mediator-v1-OAuthService)
    - [OrganizationService](#mediator-v1-OrganizationService)
    - [RoleService](#mediator-v1-RoleService)
    - [SecretsService](#mediator-v1-SecretsService)
    - [UserService](#mediator-v1-UserService)
    - [VulnerabilitiesService](#mediator-v1-VulnerabilitiesService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="mediator_v1_mediator-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mediator/v1/mediator.proto



<a name="mediator-v1-AuthVerifyServiceVerifyRequest"></a>

### AuthVerifyServiceVerifyRequest







<a name="mediator-v1-AuthVerifyServiceVerifyResponse"></a>

### AuthVerifyServiceVerifyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  |  |






<a name="mediator-v1-BranchProtection"></a>

### BranchProtection



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch | [string](#string) |  |  |
| is_protected | [bool](#bool) |  | Add other relevant fields |






<a name="mediator-v1-CheckHealthRequest"></a>

### CheckHealthRequest







<a name="mediator-v1-CheckHealthResponse"></a>

### CheckHealthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  |  |






<a name="mediator-v1-CreateGroupRequest"></a>

### CreateGroupRequest
The CreateGroupRequest message represents a request to create a group


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| organization_id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| is_protected | [bool](#bool) | optional |  |






<a name="mediator-v1-CreateGroupResponse"></a>

### CreateGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| organization_id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| is_protected | [bool](#bool) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="mediator-v1-CreateOrganizationRequest"></a>

### CreateOrganizationRequest
Organization service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| company | [string](#string) |  |  |
| create_default_records | [bool](#bool) |  |  |






<a name="mediator-v1-CreateOrganizationResponse"></a>

### CreateOrganizationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| company | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| default_group | [GroupRecord](#mediator-v1-GroupRecord) | optional |  |
| default_role | [RoleRecord](#mediator-v1-RoleRecord) | optional |  |
| default_user | [UserRecord](#mediator-v1-UserRecord) | optional |  |






<a name="mediator-v1-CreateRoleRequest"></a>

### CreateRoleRequest
Role service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| is_admin | [bool](#bool) | optional |  |
| is_protected | [bool](#bool) | optional |  |






<a name="mediator-v1-CreateRoleResponse"></a>

### CreateRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| group_id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| is_admin | [bool](#bool) |  |  |
| is_protected | [bool](#bool) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="mediator-v1-CreateUserRequest"></a>

### CreateUserRequest
User service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role_id | [int32](#int32) |  |  |
| email | [string](#string) | optional |  |
| username | [string](#string) |  |  |
| password | [string](#string) | optional |  |
| first_name | [string](#string) | optional |  |
| last_name | [string](#string) | optional |  |
| is_protected | [bool](#bool) | optional |  |






<a name="mediator-v1-CreateUserResponse"></a>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| role_id | [int32](#int32) |  |  |
| email | [string](#string) | optional |  |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |
| first_name | [string](#string) | optional |  |
| last_name | [string](#string) | optional |  |
| is_protected | [bool](#bool) | optional |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="mediator-v1-DeleteGroupRequest"></a>

### DeleteGroupRequest
DeleteGroupRequest represents a request to delete a group


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| force | [bool](#bool) | optional |  |






<a name="mediator-v1-DeleteGroupResponse"></a>

### DeleteGroupResponse
DeleteGroupResponse represents a response to a delete group request






<a name="mediator-v1-DeleteOrganizationRequest"></a>

### DeleteOrganizationRequest
DeleteOrganizationRequest represents a request to delete a organization


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| force | [bool](#bool) | optional |  |






<a name="mediator-v1-DeleteOrganizationResponse"></a>

### DeleteOrganizationResponse
DeleteOrganizationResponse represents a response to a delete organization request






<a name="mediator-v1-DeleteRoleRequest"></a>

### DeleteRoleRequest
delete role


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| force | [bool](#bool) | optional |  |






<a name="mediator-v1-DeleteRoleResponse"></a>

### DeleteRoleResponse







<a name="mediator-v1-DeleteUserRequest"></a>

### DeleteUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| force | [bool](#bool) | optional |  |






<a name="mediator-v1-DeleteUserResponse"></a>

### DeleteUserResponse







<a name="mediator-v1-ExchangeCodeForTokenCLIRequest"></a>

### ExchangeCodeForTokenCLIRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider | [string](#string) |  |  |
| code | [string](#string) |  |  |
| redirect_uri | [string](#string) |  |  |






<a name="mediator-v1-ExchangeCodeForTokenCLIResponse"></a>

### ExchangeCodeForTokenCLIResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| html | [string](#string) |  |  |






<a name="mediator-v1-ExchangeCodeForTokenWEBRequest"></a>

### ExchangeCodeForTokenWEBRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider | [string](#string) |  |  |
| code | [string](#string) |  |  |
| redirect_uri | [string](#string) |  |  |






<a name="mediator-v1-ExchangeCodeForTokenWEBResponse"></a>

### ExchangeCodeForTokenWEBResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [string](#string) |  |  |
| token_type | [string](#string) |  |  |
| expires_in | [int64](#int64) |  |  |
| status | [string](#string) |  |  |






<a name="mediator-v1-GetAuthorizationURLRequest"></a>

### GetAuthorizationURLRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider | [string](#string) |  |  |
| cli | [bool](#bool) |  |  |






<a name="mediator-v1-GetAuthorizationURLResponse"></a>

### GetAuthorizationURLResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |






<a name="mediator-v1-GetBranchProtectionRequest"></a>

### GetBranchProtectionRequest







<a name="mediator-v1-GetBranchProtectionResponse"></a>

### GetBranchProtectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_protections | [BranchProtection](#mediator-v1-BranchProtection) | repeated |  |






<a name="mediator-v1-GetGroupByIdRequest"></a>

### GetGroupByIdRequest
The GetGroupByIdRequest message represents a request to get a group by ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |






<a name="mediator-v1-GetGroupByIdResponse"></a>

### GetGroupByIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group | [GroupRecord](#mediator-v1-GroupRecord) | optional |  |






<a name="mediator-v1-GetGroupByNameRequest"></a>

### GetGroupByNameRequest
The GetGroupByNameRequest message represents a request to get a group by name


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="mediator-v1-GetGroupByNameResponse"></a>

### GetGroupByNameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group | [GroupRecord](#mediator-v1-GroupRecord) | optional |  |






<a name="mediator-v1-GetGroupsRequest"></a>

### GetGroupsRequest
The GetGroupsRequest message represents a request to get an array of groups


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| organization_id | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |






<a name="mediator-v1-GetGroupsResponse"></a>

### GetGroupsResponse
The GetGroupsResponse message represents a response with an array of groups


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groups | [GroupRecord](#mediator-v1-GroupRecord) | repeated |  |






<a name="mediator-v1-GetOrganizationByNameRequest"></a>

### GetOrganizationByNameRequest
get organization by name


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="mediator-v1-GetOrganizationByNameResponse"></a>

### GetOrganizationByNameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| organization | [OrganizationRecord](#mediator-v1-OrganizationRecord) | optional |  |






<a name="mediator-v1-GetOrganizationRequest"></a>

### GetOrganizationRequest
get organization by id


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| organization_id | [int32](#int32) |  |  |






<a name="mediator-v1-GetOrganizationResponse"></a>

### GetOrganizationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| organization | [OrganizationRecord](#mediator-v1-OrganizationRecord) | optional |  |






<a name="mediator-v1-GetOrganizationsRequest"></a>

### GetOrganizationsRequest
list organizations


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int32](#int32) | optional |  |
| offset | [int32](#int32) | optional |  |






<a name="mediator-v1-GetOrganizationsResponse"></a>

### GetOrganizationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| organizations | [OrganizationRecord](#mediator-v1-OrganizationRecord) | repeated |  |






<a name="mediator-v1-GetRoleByIdRequest"></a>

### GetRoleByIdRequest
get role by id


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="mediator-v1-GetRoleByIdResponse"></a>

### GetRoleByIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role | [RoleRecord](#mediator-v1-RoleRecord) | optional |  |






<a name="mediator-v1-GetRoleByNameRequest"></a>

### GetRoleByNameRequest
get role by group and name


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |






<a name="mediator-v1-GetRoleByNameResponse"></a>

### GetRoleByNameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role | [RoleRecord](#mediator-v1-RoleRecord) | optional |  |






<a name="mediator-v1-GetRolesRequest"></a>

### GetRolesRequest
list roles


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| limit | [int32](#int32) | optional |  |
| offset | [int32](#int32) | optional |  |






<a name="mediator-v1-GetRolesResponse"></a>

### GetRolesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| roles | [RoleRecord](#mediator-v1-RoleRecord) | repeated |  |






<a name="mediator-v1-GetSecretByIdRequest"></a>

### GetSecretByIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="mediator-v1-GetSecretByIdResponse"></a>

### GetSecretByIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  | Add other relevant fields |






<a name="mediator-v1-GetSecretsRequest"></a>

### GetSecretsRequest







<a name="mediator-v1-GetSecretsResponse"></a>

### GetSecretsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secrets | [GetSecretByIdResponse](#mediator-v1-GetSecretByIdResponse) | repeated |  |






<a name="mediator-v1-GetUserByEmailRequest"></a>

### GetUserByEmailRequest
get user by email


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |






<a name="mediator-v1-GetUserByEmailResponse"></a>

### GetUserByEmailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [UserRecord](#mediator-v1-UserRecord) | optional |  |






<a name="mediator-v1-GetUserByIdRequest"></a>

### GetUserByIdRequest
get user by id


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="mediator-v1-GetUserByIdResponse"></a>

### GetUserByIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [UserRecord](#mediator-v1-UserRecord) | optional |  |






<a name="mediator-v1-GetUserByUserNameRequest"></a>

### GetUserByUserNameRequest
get user by username


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |






<a name="mediator-v1-GetUserByUserNameResponse"></a>

### GetUserByUserNameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [UserRecord](#mediator-v1-UserRecord) | optional |  |






<a name="mediator-v1-GetUsersRequest"></a>

### GetUsersRequest
list users


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role_id | [int32](#int32) |  |  |
| limit | [int32](#int32) | optional |  |
| offset | [int32](#int32) | optional |  |






<a name="mediator-v1-GetUsersResponse"></a>

### GetUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [UserRecord](#mediator-v1-UserRecord) | repeated |  |






<a name="mediator-v1-GetVulnerabilitiesRequest"></a>

### GetVulnerabilitiesRequest







<a name="mediator-v1-GetVulnerabilitiesResponse"></a>

### GetVulnerabilitiesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vulns | [GetVulnerabilityByIdResponse](#mediator-v1-GetVulnerabilityByIdResponse) | repeated |  |






<a name="mediator-v1-GetVulnerabilityByIdRequest"></a>

### GetVulnerabilityByIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="mediator-v1-GetVulnerabilityByIdResponse"></a>

### GetVulnerabilityByIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | May require adjustment, currently set up for GitHub Security Advisories only |
| github_id | [int64](#int64) |  |  |
| repo_id | [int64](#int64) |  |  |
| repo_name | [string](#string) |  |  |
| package_name | [string](#string) |  |  |
| severity | [string](#string) |  |  |
| version_affected | [string](#string) |  |  |
| upgrade_version | [string](#string) |  |  |
| ghsaid | [string](#string) |  |  |
| advisroy_url | [string](#string) |  |  |
| scanned_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="mediator-v1-GroupRecord"></a>

### GroupRecord
BUF does not allow grouping (which is a shame)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| organization_id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| is_protected | [bool](#bool) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="mediator-v1-HandleGitHubWebhookRequest"></a>

### HandleGitHubWebhookRequest







<a name="mediator-v1-HandleGitHubWebhookResponse"></a>

### HandleGitHubWebhookResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  |  |






<a name="mediator-v1-LogInRequest"></a>

### LogInRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="mediator-v1-LogInResponse"></a>

### LogInResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  |  |
| refresh_token | [string](#string) |  |  |
| access_token | [string](#string) |  |  |
| refresh_token_expires_in | [int64](#int64) |  |  |
| access_token_expires_in | [int64](#int64) |  |  |






<a name="mediator-v1-LogOutRequest"></a>

### LogOutRequest







<a name="mediator-v1-LogOutResponse"></a>

### LogOutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  |  |






<a name="mediator-v1-OrganizationRecord"></a>

### OrganizationRecord



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| company | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="mediator-v1-RoleRecord"></a>

### RoleRecord



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| group_id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| is_admin | [bool](#bool) |  |  |
| is_protected | [bool](#bool) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="mediator-v1-UserRecord"></a>

### UserRecord
user record to be returned


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| role_id | [int32](#int32) |  |  |
| email | [string](#string) | optional |  |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |
| first_name | [string](#string) | optional |  |
| last_name | [string](#string) | optional |  |
| is_protected | [bool](#bool) | optional |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 


<a name="mediator-v1-AuthVerifyService"></a>

### AuthVerifyService
Verify user has active session to Mediator

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Verify | [AuthVerifyServiceVerifyRequest](#mediator-v1-AuthVerifyServiceVerifyRequest) | [AuthVerifyServiceVerifyResponse](#mediator-v1-AuthVerifyServiceVerifyResponse) |  |


<a name="mediator-v1-BranchProtectionService"></a>

### BranchProtectionService
Get Branch Protection Settings

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetBranchProtection | [GetBranchProtectionRequest](#mediator-v1-GetBranchProtectionRequest) | [GetBranchProtectionResponse](#mediator-v1-GetBranchProtectionResponse) |  |


<a name="mediator-v1-GitHubWebhookService"></a>

### GitHubWebhookService
GitHub App Webhook Service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| HandleGitHubWebhook | [HandleGitHubWebhookRequest](#mediator-v1-HandleGitHubWebhookRequest) | [HandleGitHubWebhookResponse](#mediator-v1-HandleGitHubWebhookResponse) |  |


<a name="mediator-v1-GroupService"></a>

### GroupService
manage Groups CRUD

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateGroup | [CreateGroupRequest](#mediator-v1-CreateGroupRequest) | [CreateGroupResponse](#mediator-v1-CreateGroupResponse) |  |
| GetGroups | [GetGroupsRequest](#mediator-v1-GetGroupsRequest) | [GetGroupsResponse](#mediator-v1-GetGroupsResponse) |  |
| GetGroupByName | [GetGroupByNameRequest](#mediator-v1-GetGroupByNameRequest) | [GetGroupByNameResponse](#mediator-v1-GetGroupByNameResponse) |  |
| GetGroupById | [GetGroupByIdRequest](#mediator-v1-GetGroupByIdRequest) | [GetGroupByIdResponse](#mediator-v1-GetGroupByIdResponse) |  |
| DeleteGroup | [DeleteGroupRequest](#mediator-v1-DeleteGroupRequest) | [DeleteGroupResponse](#mediator-v1-DeleteGroupResponse) |  |


<a name="mediator-v1-HealthService"></a>

### HealthService
Simple Health Check Service
replies with OK

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CheckHealth | [CheckHealthRequest](#mediator-v1-CheckHealthRequest) | [CheckHealthResponse](#mediator-v1-CheckHealthResponse) |  |


<a name="mediator-v1-LogInService"></a>

### LogInService
Logout of Mediator

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| LogIn | [LogInRequest](#mediator-v1-LogInRequest) | [LogInResponse](#mediator-v1-LogInResponse) |  |


<a name="mediator-v1-LogOutService"></a>

### LogOutService
Logout of Mediator

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| LogOut | [LogOutRequest](#mediator-v1-LogOutRequest) | [LogOutResponse](#mediator-v1-LogOutResponse) |  |


<a name="mediator-v1-OAuthService"></a>

### OAuthService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetAuthorizationURL | [GetAuthorizationURLRequest](#mediator-v1-GetAuthorizationURLRequest) | [GetAuthorizationURLResponse](#mediator-v1-GetAuthorizationURLResponse) |  |
| ExchangeCodeForTokenCLI | [ExchangeCodeForTokenCLIRequest](#mediator-v1-ExchangeCodeForTokenCLIRequest) | [ExchangeCodeForTokenCLIResponse](#mediator-v1-ExchangeCodeForTokenCLIResponse) |  |
| ExchangeCodeForTokenWEB | [ExchangeCodeForTokenWEBRequest](#mediator-v1-ExchangeCodeForTokenWEBRequest) | [ExchangeCodeForTokenWEBResponse](#mediator-v1-ExchangeCodeForTokenWEBResponse) |  |


<a name="mediator-v1-OrganizationService"></a>

### OrganizationService
manage Organizations CRUD

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateOrganization | [CreateOrganizationRequest](#mediator-v1-CreateOrganizationRequest) | [CreateOrganizationResponse](#mediator-v1-CreateOrganizationResponse) |  |
| GetOrganizations | [GetOrganizationsRequest](#mediator-v1-GetOrganizationsRequest) | [GetOrganizationsResponse](#mediator-v1-GetOrganizationsResponse) |  |
| GetOrganization | [GetOrganizationRequest](#mediator-v1-GetOrganizationRequest) | [GetOrganizationResponse](#mediator-v1-GetOrganizationResponse) |  |
| GetOrganizationByName | [GetOrganizationByNameRequest](#mediator-v1-GetOrganizationByNameRequest) | [GetOrganizationByNameResponse](#mediator-v1-GetOrganizationByNameResponse) |  |
| DeleteOrganization | [DeleteOrganizationRequest](#mediator-v1-DeleteOrganizationRequest) | [DeleteOrganizationResponse](#mediator-v1-DeleteOrganizationResponse) |  |


<a name="mediator-v1-RoleService"></a>

### RoleService
manage Roles CRUD

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateRole | [CreateRoleRequest](#mediator-v1-CreateRoleRequest) | [CreateRoleResponse](#mediator-v1-CreateRoleResponse) |  |
| DeleteRole | [DeleteRoleRequest](#mediator-v1-DeleteRoleRequest) | [DeleteRoleResponse](#mediator-v1-DeleteRoleResponse) |  |
| GetRoles | [GetRolesRequest](#mediator-v1-GetRolesRequest) | [GetRolesResponse](#mediator-v1-GetRolesResponse) |  |
| GetRoleById | [GetRoleByIdRequest](#mediator-v1-GetRoleByIdRequest) | [GetRoleByIdResponse](#mediator-v1-GetRoleByIdResponse) |  |
| GetRoleByName | [GetRoleByNameRequest](#mediator-v1-GetRoleByNameRequest) | [GetRoleByNameResponse](#mediator-v1-GetRoleByNameResponse) |  |


<a name="mediator-v1-SecretsService"></a>

### SecretsService
Get Secrets
Note there are different APIs for enterprise or org secrets
https://docs.github.com/en/rest/secret-scanning?apiVersion=2022-11-28

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetSecrets | [GetSecretsRequest](#mediator-v1-GetSecretsRequest) | [GetSecretsResponse](#mediator-v1-GetSecretsResponse) |  |
| GetSecretById | [GetSecretByIdRequest](#mediator-v1-GetSecretByIdRequest) | [GetSecretByIdResponse](#mediator-v1-GetSecretByIdResponse) |  |


<a name="mediator-v1-UserService"></a>

### UserService
manage Users CRUD

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateUser | [CreateUserRequest](#mediator-v1-CreateUserRequest) | [CreateUserResponse](#mediator-v1-CreateUserResponse) |  |
| DeleteUser | [DeleteUserRequest](#mediator-v1-DeleteUserRequest) | [DeleteUserResponse](#mediator-v1-DeleteUserResponse) |  |
| GetUsers | [GetUsersRequest](#mediator-v1-GetUsersRequest) | [GetUsersResponse](#mediator-v1-GetUsersResponse) |  |
| GetUserById | [GetUserByIdRequest](#mediator-v1-GetUserByIdRequest) | [GetUserByIdResponse](#mediator-v1-GetUserByIdResponse) |  |
| GetUserByUserName | [GetUserByUserNameRequest](#mediator-v1-GetUserByUserNameRequest) | [GetUserByUserNameResponse](#mediator-v1-GetUserByUserNameResponse) |  |
| GetUserByEmail | [GetUserByEmailRequest](#mediator-v1-GetUserByEmailRequest) | [GetUserByEmailResponse](#mediator-v1-GetUserByEmailResponse) |  |


<a name="mediator-v1-VulnerabilitiesService"></a>

### VulnerabilitiesService
Get Vulnerabilities

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetVulnerabilities | [GetVulnerabilitiesRequest](#mediator-v1-GetVulnerabilitiesRequest) | [GetVulnerabilitiesResponse](#mediator-v1-GetVulnerabilitiesResponse) |  |
| GetVulnerabilityById | [GetVulnerabilityByIdRequest](#mediator-v1-GetVulnerabilityByIdRequest) | [GetVulnerabilityByIdResponse](#mediator-v1-GetVulnerabilityByIdResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
