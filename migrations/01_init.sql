-- +migrate Up
-- +migrate Up
CREATE TABLE "users" (
    "user_id" text PRIMARY KEY,
    "full_name" text,
    "email" text UNIQUE,
    "password" text,
    "role" text,
    "created_at" TIMESTAMPTZ NOT NULL,
    "update_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "images" (
    "id" text PRIMARY KEY,
    "urls_full" text,
    "urls_raw" text,
    "urls_regular" text,
    "updated_at" TIMESTAMPTZ NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "width" int,
    "height" int
);

-- +migrate Down