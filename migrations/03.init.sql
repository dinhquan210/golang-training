-- +migrate Up

ALTER TABLE images
ADD  user_creat text;