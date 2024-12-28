-- name: GetCourse :one
SELECT * FROM course 
WHERE id = ? LIMIT 1;

-- name: ListCourse :many
SELECT * FROM course;

-- name: CreateCourse :one
INSERT INTO course (
	name,
	desc,
	start_date,
	end_date
) VALUES(
	?, ?, ?, ?
) RETURNING *;

-- name: UpdateCourse :one
UPDATE course
set name = ?, desc = ?, start_date = ?, end_date = ?
WHERE id = ?
RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM course WHERE id = ?;
