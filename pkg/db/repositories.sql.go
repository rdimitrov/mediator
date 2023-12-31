// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: repositories.sql

package db

import (
	"context"
	"database/sql"
)

const createRepository = `-- name: CreateRepository :one
INSERT INTO repositories (
    provider,
    group_id,
    repo_owner, 
    repo_name,
    repo_id,
    is_private,
    is_fork,
    webhook_id,
    webhook_url,
    deploy_url,
    clone_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id, provider, group_id, repo_owner, repo_name, repo_id, is_private, is_fork, webhook_id, webhook_url, deploy_url, clone_url, created_at, updated_at
`

type CreateRepositoryParams struct {
	Provider   string        `json:"provider"`
	GroupID    int32         `json:"group_id"`
	RepoOwner  string        `json:"repo_owner"`
	RepoName   string        `json:"repo_name"`
	RepoID     int32         `json:"repo_id"`
	IsPrivate  bool          `json:"is_private"`
	IsFork     bool          `json:"is_fork"`
	WebhookID  sql.NullInt32 `json:"webhook_id"`
	WebhookUrl string        `json:"webhook_url"`
	DeployUrl  string        `json:"deploy_url"`
	CloneUrl   string        `json:"clone_url"`
}

func (q *Queries) CreateRepository(ctx context.Context, arg CreateRepositoryParams) (Repository, error) {
	row := q.db.QueryRowContext(ctx, createRepository,
		arg.Provider,
		arg.GroupID,
		arg.RepoOwner,
		arg.RepoName,
		arg.RepoID,
		arg.IsPrivate,
		arg.IsFork,
		arg.WebhookID,
		arg.WebhookUrl,
		arg.DeployUrl,
		arg.CloneUrl,
	)
	var i Repository
	err := row.Scan(
		&i.ID,
		&i.Provider,
		&i.GroupID,
		&i.RepoOwner,
		&i.RepoName,
		&i.RepoID,
		&i.IsPrivate,
		&i.IsFork,
		&i.WebhookID,
		&i.WebhookUrl,
		&i.DeployUrl,
		&i.CloneUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteRepository = `-- name: DeleteRepository :exec
DELETE FROM repositories
WHERE id = $1
`

func (q *Queries) DeleteRepository(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteRepository, id)
	return err
}

const getRepositoryByID = `-- name: GetRepositoryByID :one
SELECT id, provider, group_id, repo_owner, repo_name, repo_id, is_private, is_fork, webhook_id, webhook_url, deploy_url, clone_url, created_at, updated_at FROM repositories WHERE id = $1
`

func (q *Queries) GetRepositoryByID(ctx context.Context, id int32) (Repository, error) {
	row := q.db.QueryRowContext(ctx, getRepositoryByID, id)
	var i Repository
	err := row.Scan(
		&i.ID,
		&i.Provider,
		&i.GroupID,
		&i.RepoOwner,
		&i.RepoName,
		&i.RepoID,
		&i.IsPrivate,
		&i.IsFork,
		&i.WebhookID,
		&i.WebhookUrl,
		&i.DeployUrl,
		&i.CloneUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRepositoryByIDAndGroup = `-- name: GetRepositoryByIDAndGroup :one
SELECT id, provider, group_id, repo_owner, repo_name, repo_id, is_private, is_fork, webhook_id, webhook_url, deploy_url, clone_url, created_at, updated_at FROM repositories WHERE provider = $1 AND repo_id = $2 AND group_id = $3
`

type GetRepositoryByIDAndGroupParams struct {
	Provider string `json:"provider"`
	RepoID   int32  `json:"repo_id"`
	GroupID  int32  `json:"group_id"`
}

func (q *Queries) GetRepositoryByIDAndGroup(ctx context.Context, arg GetRepositoryByIDAndGroupParams) (Repository, error) {
	row := q.db.QueryRowContext(ctx, getRepositoryByIDAndGroup, arg.Provider, arg.RepoID, arg.GroupID)
	var i Repository
	err := row.Scan(
		&i.ID,
		&i.Provider,
		&i.GroupID,
		&i.RepoOwner,
		&i.RepoName,
		&i.RepoID,
		&i.IsPrivate,
		&i.IsFork,
		&i.WebhookID,
		&i.WebhookUrl,
		&i.DeployUrl,
		&i.CloneUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRepositoryByRepoID = `-- name: GetRepositoryByRepoID :one
SELECT id, provider, group_id, repo_owner, repo_name, repo_id, is_private, is_fork, webhook_id, webhook_url, deploy_url, clone_url, created_at, updated_at FROM repositories WHERE provider = $1 AND repo_id = $2
`

type GetRepositoryByRepoIDParams struct {
	Provider string `json:"provider"`
	RepoID   int32  `json:"repo_id"`
}

func (q *Queries) GetRepositoryByRepoID(ctx context.Context, arg GetRepositoryByRepoIDParams) (Repository, error) {
	row := q.db.QueryRowContext(ctx, getRepositoryByRepoID, arg.Provider, arg.RepoID)
	var i Repository
	err := row.Scan(
		&i.ID,
		&i.Provider,
		&i.GroupID,
		&i.RepoOwner,
		&i.RepoName,
		&i.RepoID,
		&i.IsPrivate,
		&i.IsFork,
		&i.WebhookID,
		&i.WebhookUrl,
		&i.DeployUrl,
		&i.CloneUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRepositoryByRepoName = `-- name: GetRepositoryByRepoName :one
SELECT id, provider, group_id, repo_owner, repo_name, repo_id, is_private, is_fork, webhook_id, webhook_url, deploy_url, clone_url, created_at, updated_at FROM repositories WHERE provider = $1 AND repo_owner = $2 AND repo_name = $3
`

type GetRepositoryByRepoNameParams struct {
	Provider  string `json:"provider"`
	RepoOwner string `json:"repo_owner"`
	RepoName  string `json:"repo_name"`
}

func (q *Queries) GetRepositoryByRepoName(ctx context.Context, arg GetRepositoryByRepoNameParams) (Repository, error) {
	row := q.db.QueryRowContext(ctx, getRepositoryByRepoName, arg.Provider, arg.RepoOwner, arg.RepoName)
	var i Repository
	err := row.Scan(
		&i.ID,
		&i.Provider,
		&i.GroupID,
		&i.RepoOwner,
		&i.RepoName,
		&i.RepoID,
		&i.IsPrivate,
		&i.IsFork,
		&i.WebhookID,
		&i.WebhookUrl,
		&i.DeployUrl,
		&i.CloneUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAllRepositories = `-- name: ListAllRepositories :many
SELECT id, provider, group_id, repo_owner, repo_name, repo_id, is_private, is_fork, webhook_id, webhook_url, deploy_url, clone_url, created_at, updated_at FROM repositories WHERE provider = $1
ORDER BY id
`

func (q *Queries) ListAllRepositories(ctx context.Context, provider string) ([]Repository, error) {
	rows, err := q.db.QueryContext(ctx, listAllRepositories, provider)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Repository{}
	for rows.Next() {
		var i Repository
		if err := rows.Scan(
			&i.ID,
			&i.Provider,
			&i.GroupID,
			&i.RepoOwner,
			&i.RepoName,
			&i.RepoID,
			&i.IsPrivate,
			&i.IsFork,
			&i.WebhookID,
			&i.WebhookUrl,
			&i.DeployUrl,
			&i.CloneUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRegisteredRepositoriesByGroupIDAndProvider = `-- name: ListRegisteredRepositoriesByGroupIDAndProvider :many
SELECT id, provider, group_id, repo_owner, repo_name, repo_id, is_private, is_fork, webhook_id, webhook_url, deploy_url, clone_url, created_at, updated_at FROM repositories
WHERE provider = $1 AND group_id = $2 AND webhook_id IS NOT NULL
ORDER BY id
`

type ListRegisteredRepositoriesByGroupIDAndProviderParams struct {
	Provider string `json:"provider"`
	GroupID  int32  `json:"group_id"`
}

func (q *Queries) ListRegisteredRepositoriesByGroupIDAndProvider(ctx context.Context, arg ListRegisteredRepositoriesByGroupIDAndProviderParams) ([]Repository, error) {
	rows, err := q.db.QueryContext(ctx, listRegisteredRepositoriesByGroupIDAndProvider, arg.Provider, arg.GroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Repository{}
	for rows.Next() {
		var i Repository
		if err := rows.Scan(
			&i.ID,
			&i.Provider,
			&i.GroupID,
			&i.RepoOwner,
			&i.RepoName,
			&i.RepoID,
			&i.IsPrivate,
			&i.IsFork,
			&i.WebhookID,
			&i.WebhookUrl,
			&i.DeployUrl,
			&i.CloneUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRepositoriesByGroupID = `-- name: ListRepositoriesByGroupID :many
SELECT id, provider, group_id, repo_owner, repo_name, repo_id, is_private, is_fork, webhook_id, webhook_url, deploy_url, clone_url, created_at, updated_at FROM repositories
WHERE provider = $1 AND group_id = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

type ListRepositoriesByGroupIDParams struct {
	Provider string `json:"provider"`
	GroupID  int32  `json:"group_id"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

func (q *Queries) ListRepositoriesByGroupID(ctx context.Context, arg ListRepositoriesByGroupIDParams) ([]Repository, error) {
	rows, err := q.db.QueryContext(ctx, listRepositoriesByGroupID,
		arg.Provider,
		arg.GroupID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Repository{}
	for rows.Next() {
		var i Repository
		if err := rows.Scan(
			&i.ID,
			&i.Provider,
			&i.GroupID,
			&i.RepoOwner,
			&i.RepoName,
			&i.RepoID,
			&i.IsPrivate,
			&i.IsFork,
			&i.WebhookID,
			&i.WebhookUrl,
			&i.DeployUrl,
			&i.CloneUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRepositoriesByOwner = `-- name: ListRepositoriesByOwner :many
SELECT id, provider, group_id, repo_owner, repo_name, repo_id, is_private, is_fork, webhook_id, webhook_url, deploy_url, clone_url, created_at, updated_at FROM repositories
WHERE provider = $1 AND repo_owner = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

type ListRepositoriesByOwnerParams struct {
	Provider  string `json:"provider"`
	RepoOwner string `json:"repo_owner"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
}

func (q *Queries) ListRepositoriesByOwner(ctx context.Context, arg ListRepositoriesByOwnerParams) ([]Repository, error) {
	rows, err := q.db.QueryContext(ctx, listRepositoriesByOwner,
		arg.Provider,
		arg.RepoOwner,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Repository{}
	for rows.Next() {
		var i Repository
		if err := rows.Scan(
			&i.ID,
			&i.Provider,
			&i.GroupID,
			&i.RepoOwner,
			&i.RepoName,
			&i.RepoID,
			&i.IsPrivate,
			&i.IsFork,
			&i.WebhookID,
			&i.WebhookUrl,
			&i.DeployUrl,
			&i.CloneUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRepository = `-- name: UpdateRepository :one
UPDATE repositories 
SET group_id = $2,
repo_owner = $3,
repo_name = $4,
repo_id = $5,
is_private = $6,
is_fork = $7,
webhook_id = $8,
webhook_url = $9,
deploy_url = $10, 
provider = $11,
clone_url = $12,
updated_at = NOW() 
WHERE id = $1 RETURNING id, provider, group_id, repo_owner, repo_name, repo_id, is_private, is_fork, webhook_id, webhook_url, deploy_url, clone_url, created_at, updated_at
`

type UpdateRepositoryParams struct {
	ID         int32         `json:"id"`
	GroupID    int32         `json:"group_id"`
	RepoOwner  string        `json:"repo_owner"`
	RepoName   string        `json:"repo_name"`
	RepoID     int32         `json:"repo_id"`
	IsPrivate  bool          `json:"is_private"`
	IsFork     bool          `json:"is_fork"`
	WebhookID  sql.NullInt32 `json:"webhook_id"`
	WebhookUrl string        `json:"webhook_url"`
	DeployUrl  string        `json:"deploy_url"`
	Provider   string        `json:"provider"`
	CloneUrl   string        `json:"clone_url"`
}

func (q *Queries) UpdateRepository(ctx context.Context, arg UpdateRepositoryParams) (Repository, error) {
	row := q.db.QueryRowContext(ctx, updateRepository,
		arg.ID,
		arg.GroupID,
		arg.RepoOwner,
		arg.RepoName,
		arg.RepoID,
		arg.IsPrivate,
		arg.IsFork,
		arg.WebhookID,
		arg.WebhookUrl,
		arg.DeployUrl,
		arg.Provider,
		arg.CloneUrl,
	)
	var i Repository
	err := row.Scan(
		&i.ID,
		&i.Provider,
		&i.GroupID,
		&i.RepoOwner,
		&i.RepoName,
		&i.RepoID,
		&i.IsPrivate,
		&i.IsFork,
		&i.WebhookID,
		&i.WebhookUrl,
		&i.DeployUrl,
		&i.CloneUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateRepositoryByID = `-- name: UpdateRepositoryByID :one
UPDATE repositories 
SET group_id = $2,
repo_owner = $3,
repo_name = $4,
is_private = $5,
is_fork = $6,
webhook_id = $7,
webhook_url = $8,
deploy_url = $9, 
provider = $10,
clone_url = $11,
updated_at = NOW() 
WHERE repo_id = $1 RETURNING id, provider, group_id, repo_owner, repo_name, repo_id, is_private, is_fork, webhook_id, webhook_url, deploy_url, clone_url, created_at, updated_at
`

type UpdateRepositoryByIDParams struct {
	RepoID     int32         `json:"repo_id"`
	GroupID    int32         `json:"group_id"`
	RepoOwner  string        `json:"repo_owner"`
	RepoName   string        `json:"repo_name"`
	IsPrivate  bool          `json:"is_private"`
	IsFork     bool          `json:"is_fork"`
	WebhookID  sql.NullInt32 `json:"webhook_id"`
	WebhookUrl string        `json:"webhook_url"`
	DeployUrl  string        `json:"deploy_url"`
	Provider   string        `json:"provider"`
	CloneUrl   string        `json:"clone_url"`
}

func (q *Queries) UpdateRepositoryByID(ctx context.Context, arg UpdateRepositoryByIDParams) (Repository, error) {
	row := q.db.QueryRowContext(ctx, updateRepositoryByID,
		arg.RepoID,
		arg.GroupID,
		arg.RepoOwner,
		arg.RepoName,
		arg.IsPrivate,
		arg.IsFork,
		arg.WebhookID,
		arg.WebhookUrl,
		arg.DeployUrl,
		arg.Provider,
		arg.CloneUrl,
	)
	var i Repository
	err := row.Scan(
		&i.ID,
		&i.Provider,
		&i.GroupID,
		&i.RepoOwner,
		&i.RepoName,
		&i.RepoID,
		&i.IsPrivate,
		&i.IsFork,
		&i.WebhookID,
		&i.WebhookUrl,
		&i.DeployUrl,
		&i.CloneUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
