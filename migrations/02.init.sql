-- +migrate Up
ALTER TABLE images
ADD  description text;