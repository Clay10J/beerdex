-- +goose Up
CREATE TABLE breweries (
  brewery_id INTEGER PRIMARY KEY AUTOINCREMENT,
  brewery_name VARCHAR(255) COLLATE NOCASE NOT NULL,
  created_at DATETIME
    DEFAULT CURRENT_TIMESTAMP, -- column for record creation time
  city VARCHAR(255) COLLATE NOCASE NOT NULL,
  state VARCHAR(255) COLLATE NOCASE NOT NULL,
  UNIQUE (brewery_name COLLATE NOCASE, city COLLATE NOCASE, state COLLATE NOCASE)
);

-- +goose Down
DROP TABLE breweries;
