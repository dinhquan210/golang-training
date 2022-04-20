-- +migrate Up
CREATE TABLE "reacts" (
    "id_react" TEXT PRIMARY KEY,
    "id_image" TEXT,
    "react" TEXT,
    "id_user" TEXT
);
