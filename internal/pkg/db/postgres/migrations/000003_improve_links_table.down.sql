BEGIN
    ISOLATION LEVEL READ COMMITTED;

ALTER TABLE links DROP COLUMN created_at;
ALTER TABLE links DROP COLUMN updated_at;

COMMIT;
