-- +goose Up
CREATE TABLE beers (
  beer_id INTEGER PRIMARY KEY AUTOINCREMENT,
  beer_name VARCHAR(255) COLLATE NOCASE NOT NULL,
  brewery_id INT NOT NULL,
  abv DECIMAL(4, 1), -- ABV stored as a decimal number with precision 4 (total digits) and 1 decimal place
  -- date_consumed?
  created_at DATETIME
    DEFAULT CURRENT_TIMESTAMP, -- column for record creation time
  package_type VARCHAR(255) COLLATE NOCASE NOT NULL,
  FOREIGN KEY (brewery_id)
    REFERENCES breweries (brewery_id),
  UNIQUE (beer_name COLLATE NOCASE, brewery_id COLLATE NOCASE)
);

-- +goose Down
DROP TABLE beers;
