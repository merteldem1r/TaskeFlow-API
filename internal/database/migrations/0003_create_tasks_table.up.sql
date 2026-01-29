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