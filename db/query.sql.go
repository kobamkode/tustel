// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"
)

const createImageInfo = `-- name: CreateImageInfo :exec
insert into images (
	name, path
) values (
	$1, $2
)
returning id, name, path
`

type CreateImageInfoParams struct {
	Name string
	Path string
}

func (q *Queries) CreateImageInfo(ctx context.Context, arg CreateImageInfoParams) error {
	_, err := q.db.Exec(ctx, createImageInfo, arg.Name, arg.Path)
	return err
}
