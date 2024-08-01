-- name: CreateImageInfo :exec
insert into images (
	name, path
) values (
	$1, $2
)
returning *;
