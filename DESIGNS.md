# DB

## Tables

### Beer

- beer_id [primary key]
- beer_name
- beer_type
- ABV?
- Date consumed?
- created_at?
- Package type (draft/can/bottle)
- brewery_id [foreign key]
```sql
CREATE TABLE beers (
  beer_id INT PRIMARY KEY,
  beer_name VARCHAR(255),
  brewery_id INT,
  abv DECIMAL(4, 1), -- ABV stored as a decimal number with precision 4 (total digits) and 1 decimal place
  -- date_consumed?
  created_at DATETIME
    DEFAULT CURRENT_TIMESTAMP, -- column for record creation time
  package_type VARCHAR(255),
  FOREIGN KEY (brewery_id)
    REFERENCES breweries (brewery_id)
);
```

### Brewery

- brewery_id [primary key]
- brewery_name
- created_at?
- city (not separating multiple locations out)
- state
```sql
CREATE TABLE breweries (
  brewery_id INT PRIMARY KEY,
  brewery_name VARCHAR(255),
  created_at DATETIME
    DEFAULT CURRENT_TIMESTAMP, -- column for record creation time
  city VARCHAR(255),
  state VARCHAR(255)
);
```

### Rating

- beer_id [foreign key]
- user_id [foreign key]
- rating (x/5 pints)
- (beer_id, user_id) [primary key]
```sql
CREATE TABLE ratings (
  beer_id INT,
  user_id INT,
  rating INT,
  created_at DATETIME
    DEFAULT CURRENT_TIMESTAMP, -- column for record creation time
  PRIMARY KEY (beer_id, user_id),
  FOREIGN KEY (beer_id)
    REFERENCES beers (beer_id),
  FOREIGN KEY (user_id)
    REFERENCES users (user_id)
);
```

### User (using Google sign-in)

- user_id [primary key]
- google_user_id [unique] (provided by Google)
- email [unique]
- display_name
```sql
CREATE TABLE users (
  user_id INT PRIMARY KEY,
  google_user_id VARCHAR(255) UNIQUE,
  email VARCHAR(255) UNIQUE,
  display_name VARCHAR(255),
  created_at DATETIME
    DEFAULT CURRENT_TIMESTAMP, -- column for record creation time
);
```
