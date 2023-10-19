-- name: CreateTask :one
INSERT INTO tasks ("user_id", "title", "description")
VALUES
($1,$2,$3)
RETURNING *;

-- name: GetTaskByID :one
SELECT * FROM tasks
WHERE id = $1
LIMIT 1;

-- name: GetAllTask :many
SELECT * FROM tasks;

-- name: GetAllTaskByUser :many
SELECT * FROM tasks where user_id=$1;

-- name: UpdateTask :one
UPDATE tasks
SET
  user_id = $1,
  title = $2,
  description= $3,
  status=$4,
  updated_at=now()
WHERE id = $5
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id=$1;