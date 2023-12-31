// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	AddUserGroup(ctx context.Context, arg AddUserGroupParams) (UserGroup, error)
	AddUserRole(ctx context.Context, arg AddUserRoleParams) (UserRole, error)
	CleanTokenIat(ctx context.Context, id int32) (User, error)
	CreateAccessToken(ctx context.Context, arg CreateAccessTokenParams) (ProviderAccessToken, error)
	CreateArtifact(ctx context.Context, arg CreateArtifactParams) (Artifact, error)
	CreateArtifactVersion(ctx context.Context, arg CreateArtifactVersionParams) (ArtifactVersion, error)
	CreateGroup(ctx context.Context, arg CreateGroupParams) (Group, error)
	CreateOrganization(ctx context.Context, arg CreateOrganizationParams) (Organization, error)
	CreatePolicy(ctx context.Context, arg CreatePolicyParams) (Policy, error)
	CreatePolicyForEntity(ctx context.Context, arg CreatePolicyForEntityParams) (EntityPolicy, error)
	CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error)
	CreateRepository(ctx context.Context, arg CreateRepositoryParams) (Repository, error)
	CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error)
	CreateRuleEvaluationStatusForRepository(ctx context.Context, arg CreateRuleEvaluationStatusForRepositoryParams) error
	CreateRuleType(ctx context.Context, arg CreateRuleTypeParams) (RuleType, error)
	CreateSessionState(ctx context.Context, arg CreateSessionStateParams) (SessionStore, error)
	CreateSigningKey(ctx context.Context, arg CreateSigningKeyParams) (SigningKey, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccessToken(ctx context.Context, arg DeleteAccessTokenParams) error
	DeleteArtifact(ctx context.Context, id int32) error
	DeleteArtifactVersion(ctx context.Context, id int32) error
	DeleteExpiredSessionStates(ctx context.Context) error
	DeleteGroup(ctx context.Context, id int32) error
	DeleteOldArtifactVersions(ctx context.Context, arg DeleteOldArtifactVersionsParams) error
	DeleteOrganization(ctx context.Context, id int32) error
	DeletePolicy(ctx context.Context, id int32) error
	DeleteProject(ctx context.Context, id uuid.UUID) ([]DeleteProjectRow, error)
	DeleteRepository(ctx context.Context, id int32) error
	DeleteRole(ctx context.Context, id int32) error
	DeleteRuleType(ctx context.Context, id int32) error
	DeleteSessionState(ctx context.Context, id int32) error
	DeleteSessionStateByGroupID(ctx context.Context, arg DeleteSessionStateByGroupIDParams) error
	DeleteSigningKey(ctx context.Context, arg DeleteSigningKeyParams) error
	DeleteUser(ctx context.Context, id int32) error
	GetAccessTokenByGroupID(ctx context.Context, arg GetAccessTokenByGroupIDParams) (ProviderAccessToken, error)
	GetAccessTokenByProvider(ctx context.Context, provider string) ([]ProviderAccessToken, error)
	GetAccessTokenSinceDate(ctx context.Context, arg GetAccessTokenSinceDateParams) (ProviderAccessToken, error)
	GetArtifactByID(ctx context.Context, id int32) (GetArtifactByIDRow, error)
	GetArtifactVersionByID(ctx context.Context, id int32) (ArtifactVersion, error)
	GetArtifactVersionBySha(ctx context.Context, arg GetArtifactVersionByShaParams) (ArtifactVersion, error)
	GetChildrenProjects(ctx context.Context, id uuid.UUID) ([]uuid.UUID, error)
	GetGroupByID(ctx context.Context, id int32) (Group, error)
	GetGroupByName(ctx context.Context, name string) (Group, error)
	GetGroupIDPortBySessionState(ctx context.Context, sessionState string) (GetGroupIDPortBySessionStateRow, error)
	GetOrganization(ctx context.Context, id int32) (Organization, error)
	GetOrganizationByName(ctx context.Context, name string) (Organization, error)
	GetOrganizationForUpdate(ctx context.Context, name string) (Organization, error)
	GetParentProjects(ctx context.Context, id uuid.UUID) ([]uuid.UUID, error)
	GetParentProjectsUntil(ctx context.Context, arg GetParentProjectsUntilParams) ([]uuid.UUID, error)
	GetPolicyByGroupAndID(ctx context.Context, arg GetPolicyByGroupAndIDParams) ([]GetPolicyByGroupAndIDRow, error)
	GetPolicyByGroupAndName(ctx context.Context, arg GetPolicyByGroupAndNameParams) ([]GetPolicyByGroupAndNameRow, error)
	GetPolicyByID(ctx context.Context, id int32) (Policy, error)
	GetPolicyStatusByGroup(ctx context.Context, groupID int32) ([]GetPolicyStatusByGroupRow, error)
	GetPolicyStatusByIdAndGroup(ctx context.Context, arg GetPolicyStatusByIdAndGroupParams) (GetPolicyStatusByIdAndGroupRow, error)
	GetProjectByID(ctx context.Context, id uuid.UUID) (Project, error)
	GetRepositoryByID(ctx context.Context, id int32) (Repository, error)
	GetRepositoryByIDAndGroup(ctx context.Context, arg GetRepositoryByIDAndGroupParams) (Repository, error)
	GetRepositoryByRepoID(ctx context.Context, arg GetRepositoryByRepoIDParams) (Repository, error)
	GetRepositoryByRepoName(ctx context.Context, arg GetRepositoryByRepoNameParams) (Repository, error)
	GetRoleByID(ctx context.Context, id int32) (Role, error)
	GetRoleByName(ctx context.Context, arg GetRoleByNameParams) (Role, error)
	GetRootProjects(ctx context.Context) ([]Project, error)
	GetRuleTypeByID(ctx context.Context, id int32) (RuleType, error)
	GetRuleTypeByName(ctx context.Context, arg GetRuleTypeByNameParams) (RuleType, error)
	GetSessionState(ctx context.Context, id int32) (SessionStore, error)
	GetSessionStateByGroupID(ctx context.Context, grpID sql.NullInt32) (SessionStore, error)
	GetSigningKeyByGroupID(ctx context.Context, groupID int32) (SigningKey, error)
	GetSigningKeyByIdentifier(ctx context.Context, keyIdentifier string) (SigningKey, error)
	GetUserByEmail(ctx context.Context, email sql.NullString) (User, error)
	GetUserByID(ctx context.Context, id int32) (User, error)
	GetUserByUserName(ctx context.Context, username string) (User, error)
	GetUserGroups(ctx context.Context, userID int32) ([]GetUserGroupsRow, error)
	GetUserRoles(ctx context.Context, userID int32) ([]GetUserRolesRow, error)
	ListAllRepositories(ctx context.Context, provider string) ([]Repository, error)
	ListArtifactVersionsByArtifactID(ctx context.Context, arg ListArtifactVersionsByArtifactIDParams) ([]ArtifactVersion, error)
	ListArtifactVersionsByArtifactIDAndTag(ctx context.Context, arg ListArtifactVersionsByArtifactIDAndTagParams) ([]ArtifactVersion, error)
	ListArtifactsByRepoID(ctx context.Context, repositoryID int32) ([]Artifact, error)
	ListGroups(ctx context.Context, arg ListGroupsParams) ([]Group, error)
	ListGroupsByOrganizationID(ctx context.Context, organizationID int32) ([]Group, error)
	ListOrganizations(ctx context.Context, arg ListOrganizationsParams) ([]Organization, error)
	ListPoliciesByGroupID(ctx context.Context, groupID int32) ([]ListPoliciesByGroupIDRow, error)
	ListRegisteredRepositoriesByGroupIDAndProvider(ctx context.Context, arg ListRegisteredRepositoriesByGroupIDAndProviderParams) ([]Repository, error)
	ListRepositoriesByGroupID(ctx context.Context, arg ListRepositoriesByGroupIDParams) ([]Repository, error)
	ListRepositoriesByOwner(ctx context.Context, arg ListRepositoriesByOwnerParams) ([]Repository, error)
	ListRoles(ctx context.Context, arg ListRolesParams) ([]Role, error)
	ListRolesByGroupID(ctx context.Context, arg ListRolesByGroupIDParams) ([]Role, error)
	ListRuleEvaluationStatusByPolicyId(ctx context.Context, arg ListRuleEvaluationStatusByPolicyIdParams) ([]ListRuleEvaluationStatusByPolicyIdRow, error)
	ListRuleTypesByProviderAndGroup(ctx context.Context, arg ListRuleTypesByProviderAndGroupParams) ([]RuleType, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	ListUsersByGroup(ctx context.Context, arg ListUsersByGroupParams) ([]User, error)
	ListUsersByOrganization(ctx context.Context, arg ListUsersByOrganizationParams) ([]User, error)
	ListUsersByRoleId(ctx context.Context, roleID int32) ([]int32, error)
	RevokeUserToken(ctx context.Context, arg RevokeUserTokenParams) (User, error)
	RevokeUsersTokens(ctx context.Context, minTokenIssuedTime sql.NullTime) (User, error)
	UpdateAccessToken(ctx context.Context, arg UpdateAccessTokenParams) (ProviderAccessToken, error)
	UpdateGroup(ctx context.Context, arg UpdateGroupParams) (Group, error)
	UpdateOrganization(ctx context.Context, arg UpdateOrganizationParams) (Organization, error)
	UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (User, error)
	UpdateRepository(ctx context.Context, arg UpdateRepositoryParams) (Repository, error)
	UpdateRepositoryByID(ctx context.Context, arg UpdateRepositoryByIDParams) (Repository, error)
	UpdateRole(ctx context.Context, arg UpdateRoleParams) (Role, error)
	UpdateRuleEvaluationStatusForRepository(ctx context.Context, arg UpdateRuleEvaluationStatusForRepositoryParams) error
	UpdateRuleType(ctx context.Context, arg UpdateRuleTypeParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpsertArtifact(ctx context.Context, arg UpsertArtifactParams) (Artifact, error)
	UpsertArtifactVersion(ctx context.Context, arg UpsertArtifactVersionParams) (ArtifactVersion, error)
	UpsertRuleEvaluationStatus(ctx context.Context, arg UpsertRuleEvaluationStatusParams) error
}

var _ Querier = (*Queries)(nil)
