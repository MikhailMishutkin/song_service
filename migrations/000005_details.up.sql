CREATE TABLE details (
    uniq_id integer PRIMARY KEY NOT NULL,
    release_date varchar,
    text varchar,
    link varchar,
    FOREIGN KEY (uniq_id) REFERENCES song_unique (id)
)