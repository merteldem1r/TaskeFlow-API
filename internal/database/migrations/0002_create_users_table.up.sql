CREATE TABLE IF NOT EXISTS taskflow_db.users (
    id String,
    email String,
    password String,
    role String,
    created_at DateTime,
    updated_at DateTime
) ENGINE = MergeTree()
ORDER BY id;