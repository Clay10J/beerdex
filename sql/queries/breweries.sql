-- name: CreateBrewery :one
INSERT INTO breweries (brewery_name, city, state)
VALUES (?, ?, ?)
RETURNING *;
--

-- name: GetBreweries :many
SELECT * FROM breweries;
--

-- name: GetBrewery :one
SELECT * FROM breweries WHERE brewery_id = ?;
--

-- name: DeleteBrewery :exec
DELETE FROM breweries WHERE brewery_id = ?;
--

-- name: UpdateBreweryByID :one
UPDATE breweries
SET brewery_name = ?, city = ?, state = ?
WHERE brewery_id = ?
RETURNING *;
