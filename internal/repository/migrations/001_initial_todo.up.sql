CREATE TABLE todos (
    id varchar(36) NOT NULL,
    topic varchar NOT NULL,
    description varchar(255) NOT NULL,
    create_time TIMESTAMPTZ
);
