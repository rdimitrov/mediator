// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const cleanTokenIat = `-- name: CleanTokenIat :one
UPDATE users SET min_token_issued_time = NULL WHERE id = $1 RETURNING id, organization_id, email, username, password, needs_password_change, first_name, last_name, is_protected, created_at, updated_at, min_token_issued_time
`

func (q *Queries) CleanTokenIat(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, cleanTokenIat, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.NeedsPasswordChange,
		&i.FirstName,
		&i.LastName,
		&i.IsProtected,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MinTokenIssuedTime,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (organization_id, email, username, password, first_name, last_name, is_protected, needs_password_change) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, organization_id, email, username, password, needs_password_change, first_name, last_name, is_protected, created_at, updated_at, min_token_issued_time
`

type CreateUserParams struct {
	OrganizationID      int32          `json:"organization_id"`
	Email               sql.NullString `json:"email"`
	Username            string         `json:"username"`
	Password            string         `json:"password"`
	FirstName           sql.NullString `json:"first_name"`
	LastName            sql.NullString `json:"last_name"`
	IsProtected         bool           `json:"is_protected"`
	NeedsPasswordChange bool           `json:"needs_password_change"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.OrganizationID,
		arg.Email,
		arg.Username,
		arg.Password,
		arg.FirstName,
		arg.LastName,
		arg.IsProtected,
		arg.NeedsPasswordChange,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.NeedsPasswordChange,
		&i.FirstName,
		&i.LastName,
		&i.IsProtected,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MinTokenIssuedTime,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, organization_id, email, username, password, needs_password_change, first_name, last_name, is_protected, created_at, updated_at, min_token_issued_time FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.NeedsPasswordChange,
		&i.FirstName,
		&i.LastName,
		&i.IsProtected,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MinTokenIssuedTime,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, organization_id, email, username, password, needs_password_change, first_name, last_name, is_protected, created_at, updated_at, min_token_issued_time FROM users WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.NeedsPasswordChange,
		&i.FirstName,
		&i.LastName,
		&i.IsProtected,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MinTokenIssuedTime,
	)
	return i, err
}

const getUserByUserName = `-- name: GetUserByUserName :one
SELECT id, organization_id, email, username, password, needs_password_change, first_name, last_name, is_protected, created_at, updated_at, min_token_issued_time FROM users WHERE username = $1
`

func (q *Queries) GetUserByUserName(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUserName, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.NeedsPasswordChange,
		&i.FirstName,
		&i.LastName,
		&i.IsProtected,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MinTokenIssuedTime,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, organization_id, email, username, password, needs_password_change, first_name, last_name, is_protected, created_at, updated_at, min_token_issued_time FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.OrganizationID,
			&i.Email,
			&i.Username,
			&i.Password,
			&i.NeedsPasswordChange,
			&i.FirstName,
			&i.LastName,
			&i.IsProtected,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.MinTokenIssuedTime,
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

const listUsersByGroup = `-- name: ListUsersByGroup :many
SELECT users.id, users.organization_id, users.email, users.username, users.password, users.needs_password_change, users.first_name, users.last_name, users.is_protected, users.created_at, users.updated_at, users.min_token_issued_time FROM users
JOIN user_groups ON users.id = user_groups.user_id
WHERE user_groups.group_id = $1
ORDER BY users.id
LIMIT $2
OFFSET $3
`

type ListUsersByGroupParams struct {
	GroupID int32 `json:"group_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListUsersByGroup(ctx context.Context, arg ListUsersByGroupParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsersByGroup, arg.GroupID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.OrganizationID,
			&i.Email,
			&i.Username,
			&i.Password,
			&i.NeedsPasswordChange,
			&i.FirstName,
			&i.LastName,
			&i.IsProtected,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.MinTokenIssuedTime,
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

const listUsersByOrganization = `-- name: ListUsersByOrganization :many
SELECT id, organization_id, email, username, password, needs_password_change, first_name, last_name, is_protected, created_at, updated_at, min_token_issued_time FROM users
WHERE organization_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListUsersByOrganizationParams struct {
	OrganizationID int32 `json:"organization_id"`
	Limit          int32 `json:"limit"`
	Offset         int32 `json:"offset"`
}

func (q *Queries) ListUsersByOrganization(ctx context.Context, arg ListUsersByOrganizationParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsersByOrganization, arg.OrganizationID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.OrganizationID,
			&i.Email,
			&i.Username,
			&i.Password,
			&i.NeedsPasswordChange,
			&i.FirstName,
			&i.LastName,
			&i.IsProtected,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.MinTokenIssuedTime,
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

const revokeUserToken = `-- name: RevokeUserToken :one
UPDATE users SET min_token_issued_time = $2 WHERE id = $1 RETURNING id, organization_id, email, username, password, needs_password_change, first_name, last_name, is_protected, created_at, updated_at, min_token_issued_time
`

type RevokeUserTokenParams struct {
	ID                 int32        `json:"id"`
	MinTokenIssuedTime sql.NullTime `json:"min_token_issued_time"`
}

func (q *Queries) RevokeUserToken(ctx context.Context, arg RevokeUserTokenParams) (User, error) {
	row := q.db.QueryRowContext(ctx, revokeUserToken, arg.ID, arg.MinTokenIssuedTime)
	var i User
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.NeedsPasswordChange,
		&i.FirstName,
		&i.LastName,
		&i.IsProtected,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MinTokenIssuedTime,
	)
	return i, err
}

const revokeUsersTokens = `-- name: RevokeUsersTokens :one
UPDATE users SET min_token_issued_time = $1 RETURNING id, organization_id, email, username, password, needs_password_change, first_name, last_name, is_protected, created_at, updated_at, min_token_issued_time
`

func (q *Queries) RevokeUsersTokens(ctx context.Context, minTokenIssuedTime sql.NullTime) (User, error) {
	row := q.db.QueryRowContext(ctx, revokeUsersTokens, minTokenIssuedTime)
	var i User
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.NeedsPasswordChange,
		&i.FirstName,
		&i.LastName,
		&i.IsProtected,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MinTokenIssuedTime,
	)
	return i, err
}

const updatePassword = `-- name: UpdatePassword :one
UPDATE users SET password = $2, needs_password_change = FALSE, updated_at = NOW() WHERE id = $1 RETURNING id, organization_id, email, username, password, needs_password_change, first_name, last_name, is_protected, created_at, updated_at, min_token_issued_time
`

type UpdatePasswordParams struct {
	ID       int32  `json:"id"`
	Password string `json:"password"`
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updatePassword, arg.ID, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.NeedsPasswordChange,
		&i.FirstName,
		&i.LastName,
		&i.IsProtected,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MinTokenIssuedTime,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET email = $2, first_name = $3, last_name = $4, updated_at = NOW() WHERE id = $1 RETURNING id, organization_id, email, username, password, needs_password_change, first_name, last_name, is_protected, created_at, updated_at, min_token_issued_time
`

type UpdateUserParams struct {
	ID        int32          `json:"id"`
	Email     sql.NullString `json:"email"`
	FirstName sql.NullString `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Email,
		arg.FirstName,
		arg.LastName,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.NeedsPasswordChange,
		&i.FirstName,
		&i.LastName,
		&i.IsProtected,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MinTokenIssuedTime,
	)
	return i, err
}
