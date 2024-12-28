-- name: GetFamilyMeta :one
SELECT * FROM family 
WHERE id = ? LIMIT 1;

-- name: ListFamilies :many
SELECT * FROM family;

-- name: CreateFamily :one
INSERT INTO family(
	last_name,
	main_parent,
	sec_parent,
	phone1,
	phone2,
	phone3
) VALUES(?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateFamilyMeta :one
UPDATE family
set last_name= ?, 
main_parent= ?, 
sec_parent= ?, 
phone1= ?,
phone2= ?,
phone3= ?
WHERE id = ?
RETURNING *;

-- name: DeleteFamily :exec
DELETE FROM family WHERE id = ?;
