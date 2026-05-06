
ALTER TABLE url
    ADD CONSTRAINT short_url_unique UNIQUE (short_url);


ALTER TABLE url
    ADD CONSTRAINT url_unique UNIQUE (url);