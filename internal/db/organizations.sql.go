// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: organizations.sql

package db

import (
	"context"
)

const createOrganization = `-- name: CreateOrganization :one
INSERT INTO organizations (
    name,
    company
) VALUES (
    $1, $2
) RETURNING id, name, company, created_at, updated_at
`

type CreateOrganizationParams struct {
	Name    string `json:"name"`
	Company string `json:"company"`
}

func (q *Queries) CreateOrganization(ctx context.Context, arg CreateOrganizationParams) (Organization, error) {
	row := q.db.QueryRowContext(ctx, createOrganization, arg.Name, arg.Company)
	var i Organization
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Company,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteOrganization = `-- name: DeleteOrganization :exec
DELETE FROM organizations
WHERE id = $1
`

func (q *Queries) DeleteOrganization(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteOrganization, id)
	return err
}

const getOrganization = `-- name: GetOrganization :one
SELECT id, name, company, created_at, updated_at FROM organizations 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOrganization(ctx context.Context, id int32) (Organization, error) {
	row := q.db.QueryRowContext(ctx, getOrganization, id)
	var i Organization
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Company,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOrganizationByName = `-- name: GetOrganizationByName :one
SELECT id, name, company, created_at, updated_at FROM organizations 
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetOrganizationByName(ctx context.Context, name string) (Organization, error) {
	row := q.db.QueryRowContext(ctx, getOrganizationByName, name)
	var i Organization
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Company,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOrganizationForUpdate = `-- name: GetOrganizationForUpdate :one
SELECT id, name, company, created_at, updated_at FROM organizations
WHERE name = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetOrganizationForUpdate(ctx context.Context, name string) (Organization, error) {
	row := q.db.QueryRowContext(ctx, getOrganizationForUpdate, name)
	var i Organization
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Company,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listOrganizations = `-- name: ListOrganizations :many
SELECT id, name, company, created_at, updated_at FROM organizations
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListOrganizationsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListOrganizations(ctx context.Context, arg ListOrganizationsParams) ([]Organization, error) {
	rows, err := q.db.QueryContext(ctx, listOrganizations, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Organization{}
	for rows.Next() {
		var i Organization
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Company,
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

const updateOrganization = `-- name: UpdateOrganization :one
UPDATE organizations
SET name = $2, company = $3, updated_at = NOW()
WHERE id = $1 RETURNING id, name, company, created_at, updated_at
`

type UpdateOrganizationParams struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	Company string `json:"company"`
}

func (q *Queries) UpdateOrganization(ctx context.Context, arg UpdateOrganizationParams) (Organization, error) {
	row := q.db.QueryRowContext(ctx, updateOrganization, arg.ID, arg.Name, arg.Company)
	var i Organization
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Company,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}