
CREATE TABLE url (
    id SERIAL PRIMARY KEY,
    short_url VARCHAR(80) UNIQUE NOT NULL,
    url VARCHAR(255) NOT NULL
);

CREATE INDEX idx_short_url ON url(short_url);