-- name: CreateBeer :one
INSERT INTO beers (beer_name, brewery_id, abv, beer_type)
VALUES (?, ?, ?, ?)
RETURNING *;
--

-- name: GetBeers :many
SELECT * FROM beers;
--

-- name: GetBeer :one
SELECT * FROM beers WHERE beer_id = ?;
--

-- name: DeleteBeer :exec
DELETE FROM beers WHERE beer_id = ?;
--

-- name: UpdateBeerByID :one
UPDATE beers
SET beer_name = ?, brewery_id = ?, abv = ?, beer_type = ?
WHERE beer_id = ?
RETURNING *;
