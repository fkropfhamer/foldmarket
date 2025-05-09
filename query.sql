-- name: GetAccount :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: CreateAccount :one
INSERT INTO accounts (id, balance) VALUES ($1, $2) RETURNING *;

-- name: DeleteAllAccounts :exec
DELETE FROM accounts;

