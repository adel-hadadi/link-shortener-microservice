CREATE TABLE IF NOT EXISTS links
(
    id            SERIAL PRIMARY KEY,
    original_link VARCHAR(255) NOT NULL,
    short_link    VARCHAR(255) NOT NULL,
    created_at    TIMESTAMP DEFAULT now(),

    unique (original_link, short_link)
);

CREATE UNIQUE INDEX link_idx ON links (original_link);
