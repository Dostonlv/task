DROP TABLE IF EXISTS blogs CASCADE;
DROP TABLE IF EXISTS news CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE blogs
(
    id          UUID                       PRIMARY KEY   DEFAULT uuid_generate_v4(),
    title       VARCHAR(255)                NOT NULL    CHECK (title <> ''),
    content     VARCHAR(512)                NOT NULL    CHECK (content <> ''),
    created_at  TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT CURRENT_TIMESTAMP 
);

CREATE TABLE news
(
    id          UUID                       PRIMARY KEY   DEFAULT uuid_generate_v4(),
    title       VARCHAR(255)                NOT NULL    CHECK (title <> ''),
    content     VARCHAR(512)                NOT NULL    CHECK (content <> ''),
    created_at  TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT CURRENT_TIMESTAMP 
);