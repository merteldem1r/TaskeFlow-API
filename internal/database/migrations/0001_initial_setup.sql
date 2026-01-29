CREATE DATABASE IF NOT EXISTS taskflow_db;

CREATE TABLE IF NOT EXISTS taskflow_db.users (
    id String,
    email String,
    password String,
    role String,
    created_at DateTime,
    updated_at DateTime
) ENGINE = MergeTree()
ORDER BY id;

CREATE TABLE IF NOT EXISTS taskflow_db.tasks (
    id String,
    title String,
    description String,
    status String,
    user_id String,
    created_at DateTime,
    updated_at DateTime
) ENGINE = MergeTree()
ORDER BY id;